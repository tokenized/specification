package protocol

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/tokenized/bitcoin_interpreter/agent_bitcoin_transfer"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/messages"
	"github.com/tokenized/specification/dist/golang/permissions"
	"github.com/tokenized/txbuilder"
	"github.com/tokenized/txbuilder/fees"

	"github.com/pkg/errors"
)

var (
	ErrMissingInputScripts = errors.New("Missing Input Scripts")

	SamplePeerChannelSize = 125
)

// EstimatedContractOfferResponseTxFee estimates the satoshi amount to add to the contract fee in
// the contract output of a contract offer transaction.
// inputLockingScripts are needed to calculate the admin and operator admin fields of the contract
// formation. There must be at least one when ContractOperatorIncluded is false and at least two
// when it is true.
func EstimatedContractOfferResponseTxFee(offer *actions.ContractOffer,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	inputLockingScripts []bitcoin.Script, feeRate, dustFeeRate float64, peerChannelSize int,
	isTest bool) (uint64, error) {

	formation, err := offer.Formation()
	if err != nil {
		return 0, errors.Wrap(err, "formation")
	}

	formation.Timestamp = uint64(time.Now().UnixNano())
	formation.ContractRevision = 0
	formation.RequestPeerChannel = string(make([]byte, peerChannelSize))

	responseTxSize := txbuilder.BaseTxSize

	// Input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	// output count serialization. Dust output to contract, contract fee, and op return output
	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	if len(inputLockingScripts) == 0 {
		return 0, errors.Wrap(ErrMissingInputScripts, "admin required")
	}
	adminAddress, err := bitcoin.RawAddressFromLockingScript(inputLockingScripts[0])
	if err != nil {
		return 0, errors.Wrap(err, "admin address")
	}
	formation.AdminAddress = adminAddress.Bytes()

	if offer.ContractOperatorIncluded {
		if len(inputLockingScripts) < 2 {
			return 0, errors.Wrap(ErrMissingInputScripts, "operator required")
		}
		operatorAddress, err := bitcoin.RawAddressFromLockingScript(inputLockingScripts[1])
		if err != nil {
			return 0, errors.Wrap(err, "operator address")
		}
		formation.OperatorAddress = operatorAddress.Bytes()
	}

	formationScript, err := Serialize(formation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize formation")
	}

	// Contract formation output
	responseTxSize += txbuilder.OutputSize(formationScript)

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

func EstimatedContractAmendmentResponseTxFee(amendment *actions.ContractAmendment,
	currentFormation *actions.ContractFormation,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	inputLockingScripts []bitcoin.Script, feeRate, dustFeeRate float64,
	isTest bool) (uint64, error) {

	newFormation := currentFormation.Copy()
	newFormation.ContractRevision++
	newFormation.Timestamp = uint64(time.Now().UnixNano())

	for i, amendment := range amendment.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if _, err := newFormation.ApplyAmendment(fip, amendment.Operation, amendment.Data,
			nil); err != nil {
			return 0, errors.Wrapf(err, "apply amendment %d", i)
		}
	}

	responseTxSize := txbuilder.BaseTxSize

	// Input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	inputIndex := 1 // first input is the administrator
	if len(currentFormation.OperatorAddress) > 0 {
		inputIndex++ // second input is the contract operator
	}

	if amendment.ChangeAdministrationAddress {
		// Figure out new admin address to populate formation properly.
		if len(inputLockingScripts) <= inputIndex {
			return 0, errors.Wrapf(ErrMissingInputScripts, "new admin required: %d", inputIndex)
		}

		adminAddress, err := bitcoin.RawAddressFromLockingScript(inputLockingScripts[inputIndex])
		if err != nil {
			return 0, errors.Wrap(err, "new admin address")
		}
		newFormation.AdminAddress = adminAddress.Bytes()
		inputIndex++
	}

	if amendment.ChangeOperatorAddress {
		// Figure out new operator address to populate formation properly.
		if len(inputLockingScripts) <= inputIndex {
			return 0, errors.Wrapf(ErrMissingInputScripts, "new operator required: %d", inputIndex)
		}

		operatorAddress, err := bitcoin.RawAddressFromLockingScript(inputLockingScripts[inputIndex])
		if err != nil {
			return 0, errors.Wrap(err, "new operator address")
		}
		newFormation.OperatorAddress = operatorAddress.Bytes()
	}

	formationScript, err := Serialize(newFormation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize formation")
	}

	responseTxSize += txbuilder.OutputSize(formationScript)

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

// EstimatedBodyOfAgreementOfferResponseTxFee estimates the satoshi amount to add to the contract
// fee in the contract output of a body of agreement offer transaction.
func EstimatedBodyOfAgreementOfferResponseTxFee(offer *actions.BodyOfAgreementOffer,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	feeRate, dustFeeRate float64, isTest bool) (uint64, error) {

	formation, err := offer.Formation()
	if err != nil {
		return 0, errors.Wrap(err, "formation")
	}

	formation.Timestamp = uint64(time.Now().UnixNano())
	formation.Revision = 0

	responseTxSize := txbuilder.BaseTxSize

	// input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	// output count serialization. Dust output to contract, contract fee, and op return output
	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	formationScript, err := Serialize(formation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize formation")
	}

	// Body of agreement formation output
	responseTxSize += txbuilder.OutputSize(formationScript)

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

func EstimatedBodyOfAgreementAmendmentResponseTxFee(amendment *actions.BodyOfAgreementAmendment,
	currentFormation *actions.BodyOfAgreementFormation,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	feeRate, dustFeeRate float64, isTest bool) (uint64, error) {

	newFormation := currentFormation.Copy()
	newFormation.Revision++
	newFormation.Timestamp = uint64(time.Now().UnixNano())

	for i, amendment := range amendment.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if _, err := newFormation.ApplyAmendment(fip, amendment.Operation, amendment.Data,
			nil); err != nil {
			return 0, errors.Wrapf(err, "apply amendment %d", i)
		}
	}

	responseTxSize := txbuilder.BaseTxSize

	// Input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	formationScript, err := Serialize(newFormation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize formation")
	}

	responseTxSize += txbuilder.OutputSize(formationScript)

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

// EstimatedInstrumentDefinitionResponseTxFee estimates the satoshi amount to add to the contract
// fee in the contract output of a instrument definition transaction.
func EstimatedInstrumentDefinitionResponseTxFee(definition *actions.InstrumentDefinition,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	feeRate, dustFeeRate float64, isTest bool) (uint64, error) {

	creation, err := definition.Creation()
	if err != nil {
		return 0, errors.Wrap(err, "creation")
	}

	creation.InstrumentCode = make([]byte, InstrumentCodeSize)
	creation.Timestamp = uint64(time.Now().UnixNano())
	creation.InstrumentRevision = 0

	responseTxSize := txbuilder.BaseTxSize

	// input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	// output count serialization. Dust output to contract, contract fee, and op return output
	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	creationScript, err := Serialize(creation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize creation")
	}

	// Instrument creation output
	responseTxSize += txbuilder.OutputSize(creationScript) + 4 // add 4 for large instrument indexes

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

func EstimatedInstrumentModificationResponseTxFee(modification *actions.InstrumentModification,
	currentCreation *actions.InstrumentCreation,
	contractAgentLockingScript, contractFeeLockingScript bitcoin.Script,
	feeRate, dustFeeRate float64, isTest bool) (uint64, error) {

	newCreation := currentCreation.Copy()
	newCreation.InstrumentRevision++
	newCreation.Timestamp = uint64(time.Now().UnixNano())

	var payload instruments.Instrument
	for i, amendment := range modification.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if fip[0] == actions.InstrumentFieldInstrumentPayload {
			if payload == nil {
				// Get payload object
				payload, err = instruments.Deserialize([]byte(newCreation.InstrumentType),
					newCreation.InstrumentPayload)
				if err != nil {
					return 0, errors.Wrapf(err, "payload deserialize : %s",
						newCreation.InstrumentType)
				}
			}

			_, err = payload.ApplyAmendment(fip[1:], amendment.Operation, amendment.Data, nil)
			if err != nil {
				return 0, errors.Wrapf(err, "apply payload amendment %d", i)
			}

		} else {
			if _, err := newCreation.ApplyAmendment(fip, amendment.Operation, amendment.Data,
				nil); err != nil {
				return 0, errors.Wrapf(err, "apply amendment %d", i)
			}
		}
	}

	if payload != nil {
		// Serialize updated payload
		var buf bytes.Buffer
		if err := payload.Serialize(&buf); err != nil {
			return 0, errors.Wrap(err, "serialize payload")
		}
		newCreation.InstrumentPayload = buf.Bytes()
	}

	responseTxSize := txbuilder.BaseTxSize

	// Input count serialization. 1 input from contract
	responseTxSize += wire.VarIntSerializeSize(uint64(1))

	agentInputSize, err := txbuilder.InputSize(contractAgentLockingScript)
	if err != nil {
		return 0, errors.Wrap(err, "contract input size")
	}
	responseTxSize += agentInputSize

	if len(contractFeeLockingScript) > 0 {
		responseTxSize += wire.VarIntSerializeSize(uint64(3))

		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
		responseTxSize += txbuilder.OutputSize(contractFeeLockingScript)
	} else {
		responseTxSize += wire.VarIntSerializeSize(uint64(2))
		responseTxSize += txbuilder.OutputSize(contractAgentLockingScript)
	}

	creationScript, err := Serialize(newCreation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "serialize creation")
	}

	responseTxSize += txbuilder.OutputSize(creationScript)

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	// Dust for contract output
	responseTxFee += txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	return responseTxFee, nil
}

// EstimatedResponse calculates information about the contract's response to a request.
//   fees is the sum of all contract related fees including base contract fee, proposal fee, and
///  others. dustLimit is based on the expected P2PKH notification outputs.
// WARNING: This function is really only accurate for ContractOffer and InstrumentDefinition txs. Other
// actions might be wrong.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// For transfers use EstimatedTransferResponse.
// For contract amendments use EstimatedContractAmendmentResponse.
// For instrument modifications use EstimatedInstrumentModificationResponse.
// Returns:
// estimated size of response tx in bytes.
// estimated funding needed.
// error if there were any
func EstimatedResponse(requestTx *wire.MsgTx, inputIndex int, dustLimit, fees uint64,
	isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range requestTx.TxOut {
		action, err = Deserialize(output.LockingScript, isTest)
		if err == nil {
			found = true
			break
		}
	}
	if !found {
		return 0, 0, errors.New("Tokenized OP_RETURN not found")
	}

	// Build sample response and calculate values based on expected response inputs and outputs.
	var response actions.Action
	size := txbuilder.BaseTxSize
	value := uint64(0)
	now := time.Now()

	switch request := action.(type) {
	case *actions.ContractOffer:
		contractFormation := actions.ContractFormation{
			ContractRevision: 0,
			Timestamp:        uint64(now.UnixNano()),
		}
		response = &contractFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, contract fee, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
		}
		value += dustLimit

		contractFormation.AdminAddress = make([]byte, 21) // P2PKH Address
		if request.ContractOperatorIncluded {
			contractFormation.OperatorAddress = make([]byte, 21) // P2PKH Address
		}

	case *actions.BodyOfAgreementOffer:
		bodyOfAgreementFormation := actions.BodyOfAgreementFormation{
			Revision:  0,
			Timestamp: uint64(now.UnixNano()),
		}
		response = &bodyOfAgreementFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, contract fee, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
		}
		value += dustLimit

	case *actions.InstrumentDefinition:
		var instrumentCode bitcoin.Hash32
		instrumentCreation := actions.InstrumentCreation{
			InstrumentCode:     instrumentCode.Bytes(), // Instrument code is added by smart contract
			InstrumentRevision: 0,
			Timestamp:          uint64(now.UnixNano()),
		}
		response = &instrumentCreation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
		}
		value += dustLimit

	case *actions.Transfer:
		settlement := &actions.Settlement{Timestamp: uint64(now.UnixNano())}
		response = settlement

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		outputCount := 0
		if fees > 0 {
			// Fort multi-contract transfers use EstimatedTransferResponse below.
			outputCount += 1
			size += txbuilder.P2PKHOutputSize
			value += fees
		}

		// Add receiver outputs needed
		used := make(map[bitcoin.Hash20]bool)
		var contractScript []byte
		for _, instrument := range request.Instruments {

			settleInstrument := &actions.InstrumentSettlementField{
				InstrumentType: instrument.InstrumentType,
				InstrumentCode: instrument.InstrumentCode,
			}
			// No settlement needed for bitcoin transfers. Just outputs.
			if instrument.InstrumentType != BSVInstrumentID {
				if len(contractScript) == 0 {
					contractScript = requestTx.TxOut[instrument.ContractIndex].LockingScript
				} else if !bytes.Equal(contractScript, requestTx.TxOut[instrument.ContractIndex].LockingScript) {
					return 0, 0, errors.New("More than one contract")
				}

				settlement.Instruments = append(settlement.Instruments, settleInstrument)
			}

			// Sig script is probably still empty, so assume each sender is unique and the
			//   address is not reused. So each will get a notification output.
			if instrument.InstrumentType != BSVInstrumentID {
				for _, _ = range instrument.InstrumentSenders {
					// Use quantity that is enough for a 3 byte var 128 value, since we can't use a
					//   real value without knowing the resulting balance.
					settleInstrument.Settlements = append(settleInstrument.Settlements,
						&actions.QuantityIndexField{
							Index:    uint32(outputCount),
							Quantity: 100000,
						})
					outputCount += 1
					if instrument.InstrumentType != BSVInstrumentID {
						value += dustLimit // Dust will be put in each notification output.
					}
				}
			}

			for _, receiver := range instrument.InstrumentReceivers {
				raddress, err := bitcoin.DecodeRawAddress(receiver.Address)
				if err != nil {
					return 0, 0, errors.Wrap(err, "parsing address")
				}
				hash, err := raddress.Hash()
				if err != nil {
					return 0, 0, errors.Wrap(err, "hashing address")
				}
				isDust, exists := used[*hash]
				if !exists {
					used[*hash] = true
					// Use quantity that is enough for a 3 byte var 128 value, since we can't use a
					//   real value without knowing the resulting balance.
					settleInstrument.Settlements = append(settleInstrument.Settlements,
						&actions.QuantityIndexField{
							Index:    uint32(outputCount),
							Quantity: 100000,
						})
					outputCount += 1
					if instrument.InstrumentType != BSVInstrumentID {
						value += dustLimit // Dust will be put in each notification output.
					}
				}
				if instrument.InstrumentType == BSVInstrumentID {
					if isDust {
						value -= dustLimit
					}
					used[*hash] = false
					value += receiver.Quantity
				}
			}
		}

		// +1 for OP_RETURN
		size += wire.VarIntSerializeSize(uint64(outputCount+1)) + (outputCount * txbuilder.P2PKHOutputSize)

	case *actions.Proposal:
		if inputIndex == 0 {
			// First input funds vote (initiation) message
			vote := actions.Vote{Timestamp: uint64(now.UnixNano())}
			response = &vote

			// 1 input from contract
			size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

			// P2PKH dust output to contract, and op return output
			if fees > 0 {
				size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
				value += fees
			} else {
				size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
			}
			value += dustLimit
		} else {
			// Second input funds vote result message
			voteResult := actions.Result{Timestamp: uint64(now.UnixNano())}
			response = &voteResult

			voteResult.OptionTally = make([]uint64, len(request.VoteOptions))
			voteResult.Result = " "
			for len(voteResult.Result) < len(request.VoteOptions) {
				voteResult.Result += " "
			}

			// 1 input from contract
			size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

			// P2PKH dust output to contract, and op return output
			if fees > 0 {
				size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
				value += fees
			} else {
				size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
			}
			value += dustLimit
		}

	case *actions.BallotCast:
		counted := actions.BallotCounted{Timestamp: uint64(now.UnixNano())}
		response = &counted

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (2 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + txbuilder.P2PKHOutputSize
		}
		value += dustLimit

	default:
		return 0, 0, errors.New("Unsupported request type")
	}

	if err = convert(action, response); err != nil {
		return 0, 0, errors.Wrap(err, "Failed to convert request to response")
	}

	script, err := Serialize(response, isTest)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to serialize response envelope")
	}

	// OP_RETURN output size
	size += txbuilder.OutputBaseSize + wire.VarIntSerializeSize(uint64(len(script))) + len(script)

	return size, value, nil
}

// EstimatedContractAmendmentResponse calculates information about the contract's response to an
// amendment request.
// dustFeeRate is the fee rate used to calculate the dust limit for outputs.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// Returns:
// estimated size of response tx in bytes.
// estimated funding needed.
// error if there were any
func EstimatedContractAmendmentResponse(amendTx *wire.MsgTx, formation *actions.ContractFormation,
	dustFeeRate float32, isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range amendTx.TxOut {
		action, err = Deserialize(output.LockingScript, isTest)
		if err == nil {
			found = true
			break
		}
	}
	if !found {
		return 0, 0, errors.New("Tokenized OP_RETURN not found")
	}

	amendment, ok := action.(*actions.ContractAmendment)
	if !ok {
		return 0, 0, errors.New("Action is not a ContractAmendment")
	}

	outputSize := amendTx.TxOut[0].SerializeSize()
	dustLimit := txbuilder.DustLimitForOutput(amendTx.TxOut[0], dustFeeRate)

	// Build sample response and calculate values based on expected response inputs and outputs.
	size := txbuilder.BaseTxSize
	value := uint64(0)
	now := time.Now()

	formation.ContractRevision = amendment.ContractRevision + 1
	formation.Timestamp = uint64(now.UnixNano())

	for i, amendment := range amendment.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if _, err := formation.ApplyAmendment(fip, amendment.Operation, amendment.Data,
			nil); err != nil {
			return 0, 0, errors.Wrapf(err, "apply amendment %d", i)
		}
	}

	// 1 input from contract
	size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

	// P2PKH dust output to contract, contract fee, and op return output
	if formation.ContractFee > 0 {
		size += wire.VarIntSerializeSize(uint64(3)) + (2 * outputSize)
		value += formation.ContractFee
	} else {
		size += wire.VarIntSerializeSize(uint64(2)) + outputSize
	}
	value += dustLimit

	script, err := Serialize(formation, isTest)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to serialize contract formation envelope")
	}

	// OP_RETURN output size
	size += txbuilder.OutputBaseSize + wire.VarIntSerializeSize(uint64(len(script))) + len(script)

	return size, value, nil
}

// EstimatedBodyOfAgreementAmendmentResponse calculates information about the contract's response
// to an amendment request.
// dustFeeRate is the fee rate used to calculate the dust limit for outputs.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// Returns:
// estimated size of response tx in bytes.
// estimated funding needed.
// error if there were any
func EstimatedBodyOfAgreementAmendmentResponse(amendTx *wire.MsgTx,
	formation *actions.BodyOfAgreementFormation, contractFee uint64, dustFeeRate float32,
	isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range amendTx.TxOut {
		action, err = Deserialize(output.LockingScript, isTest)
		if err == nil {
			found = true
			break
		}
	}
	if !found {
		return 0, 0, errors.New("Tokenized OP_RETURN not found")
	}

	amendment, ok := action.(*actions.BodyOfAgreementAmendment)
	if !ok {
		return 0, 0, errors.New("Action is not a BodyOfAgreementAmendment")
	}

	outputSize := amendTx.TxOut[0].SerializeSize()
	dustLimit := txbuilder.DustLimitForOutput(amendTx.TxOut[0], dustFeeRate)

	// Build sample response and calculate values based on expected response inputs and outputs.
	size := txbuilder.BaseTxSize
	value := uint64(0)
	now := time.Now()

	formation.Revision = amendment.Revision + 1
	formation.Timestamp = uint64(now.UnixNano())

	for i, amendment := range amendment.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if _, err := formation.ApplyAmendment(fip, amendment.Operation,
			amendment.Data, nil); err != nil {
			return 0, 0, errors.Wrapf(err, "apply amendment %d", i)
		}
	}

	// 1 input from contract
	size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

	// P2PKH dust output to contract, contract fee, and op return output
	if contractFee > 0 {
		size += wire.VarIntSerializeSize(uint64(3)) + (2 * outputSize)
		value += contractFee
	} else {
		size += wire.VarIntSerializeSize(uint64(2)) + outputSize
	}
	value += dustLimit

	script, err := Serialize(formation, isTest)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to serialize body of agreement formation envelope")
	}

	// OP_RETURN output size
	size += txbuilder.OutputBaseSize + wire.VarIntSerializeSize(uint64(len(script))) + len(script)

	return size, value, nil
}

// EstimatedInstrumentModificationResponse calculates information about the contract's response to an
// instrument modification request.
// dustFeeRate is the fee rate used to calculate the dust limit for outputs.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// Returns:
// estimated size of response tx in bytes.
// estimated funding needed.
// error if there were any
func EstimatedInstrumentModificationResponse(amendTx *wire.MsgTx, ac *actions.InstrumentCreation,
	contractFee uint64, dustFeeRate float32, isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range amendTx.TxOut {
		action, err = Deserialize(output.LockingScript, isTest)
		if err == nil {
			found = true
			break
		}
	}
	if !found {
		return 0, 0, errors.New("Tokenized OP_RETURN not found")
	}

	amendment, ok := action.(*actions.InstrumentModification)
	if !ok {
		return 0, 0, errors.New("Action is not a InstrumentModification")
	}

	outputSize := amendTx.TxOut[0].SerializeSize()
	dustLimit := txbuilder.DustLimitForOutput(amendTx.TxOut[0], dustFeeRate)

	// Build sample response and calculate values based on expected response inputs and outputs.
	size := txbuilder.BaseTxSize
	value := uint64(0)
	now := time.Now()

	ac.InstrumentRevision = amendment.InstrumentRevision + 1
	ac.Timestamp = uint64(now.UnixNano())

	var payload instruments.Instrument
	for i, amendment := range amendment.Amendments {
		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if fip[0] == actions.InstrumentFieldInstrumentPayload {
			if payload == nil {
				// Get payload object
				payload, err = instruments.Deserialize([]byte(ac.InstrumentType), ac.InstrumentPayload)
				if err != nil {
					return 0, 0, fmt.Errorf("Instrument payload deserialize failed : %s %s",
						ac.InstrumentType, err)
				}
			}

			_, err = payload.ApplyAmendment(fip[1:], amendment.Operation, amendment.Data, nil)
			if err != nil {
				return 0, 0, errors.Wrapf(err, "apply payload amendment %d", i)
			}

		} else {
			if _, err := ac.ApplyAmendment(fip, amendment.Operation, amendment.Data,
				nil); err != nil {
				return 0, 0, errors.Wrapf(err, "apply amendment %d", i)
			}
		}
	}

	// 1 input from contract
	size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

	// P2PKH dust output to contract, contract fee, and op return output
	if contractFee > 0 {
		size += wire.VarIntSerializeSize(uint64(3)) + (2 * outputSize)
		value += contractFee
	} else {
		size += wire.VarIntSerializeSize(uint64(2)) + outputSize
	}
	value += dustLimit

	if payload != nil {
		// Serialize updated payload
		var buf bytes.Buffer
		if err := payload.Serialize(&buf); err != nil {
			return 0, 0, errors.Wrap(err, "serialize payload")
		}
		ac.InstrumentPayload = buf.Bytes()
	}

	script, err := Serialize(ac, isTest)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to serialize contract formation envelope")
	}

	// OP_RETURN output size
	size += txbuilder.OutputBaseSize + wire.VarIntSerializeSize(uint64(len(script))) + len(script)

	return size, value, nil
}

// EstimatedTransferResponse calculates information about the contract's response to a transfer
// request.
// fees is a list of the contract fee for the contract corresponding to each contract. The list
// lines up with the "Instruments" list in the transfer.
// dustLimit is the smallest amount of satoshis to make an output valid.
// feeRate is in satoshis per byte.
// WARNING: This function is inaccurate and incomplete!
// Returns
//   estimated funding per "Instrument" object in transfer, including contract fees.
//   boomerang funding. needs to be added to a second output to the first contract address.
//     if zero then only one contract is involved and no second output is needed.
//   error if there were any
//
// First contract is master contract. If other contracts are involved it initializes and sends
// settlement request to next contract. Then after each contract has completed the settlement
// request, the last signs the tx and sends a signature request back to the previous. After the
// first contract completes the signature request it broadcasts the completed settlement.
// Each settlement request contains the settlement action and contract fee data.
// Each signature request contains the full settlement tx minus some signatures.
// The first contract output must fund settlement tx miner fee.
// The boomerang must fund all settlement requests and signature requests.
// All other contract outputs can be dust.
// Attempts to use the maximum possible size of each element so the returned values are an
// overestimation, ensuring that a transfer funded in this manner can complete successfully.
func EstimatedTransferResponse(requestTx *wire.MsgTx, inputLockingScripts []bitcoin.Script,
	feeRate, dustFeeRate float32, contractFees []uint64, isTest bool) ([]uint64, uint64, error) {

	// Find Tokenized Transfer
	var err error
	var action actions.Action
	var request *actions.Transfer
	var ok bool
	found := false
	for _, output := range requestTx.TxOut {
		action, err = Deserialize(output.LockingScript, isTest)
		if err != nil {
			continue
		}

		request, ok = action.(*actions.Transfer)
		if ok {
			found = true
			break
		}
	}
	if !found {
		return nil, 0, errors.New("Tokenized Transfer OP_RETURN not found")
	}

	// Build sample response tx and payload and calculate values based on expected response inputs
	// and outputs.
	accumulatedSettlementData := make([][]byte, len(request.Instruments))
	var previousSettlementData []byte
	fullSettlementSize := uint64(0)
	boomerang := uint64(0) // used for funding inter-contract communication
	now := time.Now()
	// Auditable estimation construction
	fundingForTokens := make([]uint64, len(request.Instruments))
	txSizeInputs := 0
	inputCount := 0
	txSizeOutputs := 0
	outputCount := 0

	settlement := &actions.Settlement{Timestamp: uint64(now.UnixNano())}

	// Calculate response tx and settlement payload size.
	// Indexes to addreses that already have outputs so they can be reused.
	addressIndexes := make(map[bitcoin.Hash20]uint32)
	var previousContractScript, masterContractScript bitcoin.Script
	multiContract := false
	masterContractIndex := uint32(0) // First smart contract agent listed, indexed by instruments.
	for _, instrument := range request.Instruments {
		if instrument.InstrumentType != BSVInstrumentID {
			if len(masterContractScript) == 0 {
				masterContractScript = requestTx.TxOut[instrument.ContractIndex].LockingScript
			}
			break
		}
		masterContractIndex++
	}

	for i, instrument := range request.Instruments {
		settleInstrument := &actions.InstrumentSettlementField{
			ContractIndex:  uint32(i),
			InstrumentType: instrument.InstrumentType,
			InstrumentCode: instrument.InstrumentCode,
		}

		// No settlement needed for bitcoin transfers. Just outputs.
		if instrument.InstrumentType != BSVInstrumentID {
			if !bytes.Equal(previousContractScript, requestTx.TxOut[instrument.ContractIndex].LockingScript) {
				multiContract = true
				// input from each contract
				inputSize, err := txbuilder.InputSize(requestTx.TxOut[instrument.ContractIndex].LockingScript)
				if err != nil {
					return nil, 0, errors.Wrapf(err, "input size")
				}
				txSizeInputs += inputSize
				inputCount++
			}
			previousContractScript = requestTx.TxOut[instrument.ContractIndex].LockingScript

			// Sig script is probably still empty, so assume each sender is unique and the address
			// is not reused. So each will get a notification output.
			for _, sender := range instrument.InstrumentSenders {
				// Use max quantity to ensure overestimation, since we can't use a real value
				// without knowing the resulting balance.
				settleInstrument.Settlements = append(settleInstrument.Settlements,
					&actions.QuantityIndexField{
						Index:    127,
						Quantity: math.MaxUint64,
					})

				txSizeOutputs += txbuilder.OutputSize(inputLockingScripts[sender.Index])
				outputCount++
				dustLimit := txbuilder.DustLimitForLockingScript(inputLockingScripts[sender.Index],
					dustFeeRate)
				fundingForTokens[masterContractIndex] += dustLimit // Dust will be put in each notification output.
			}

		} else {
			// Bitcoin senders don't need an output because they don't need a settlement entry.
		}

		for _, receiver := range instrument.InstrumentReceivers {
			raddress, err := bitcoin.DecodeRawAddress(receiver.Address)
			if err != nil {
				return nil, 0, errors.Wrap(err, "parsing address")
			}
			hash, err := raddress.Hash()
			if err != nil {
				return nil, 0, errors.Wrap(err, "hashing address")
			}
			addressIndex, exists := addressIndexes[*hash]
			var addressDustLimit uint64
			if !exists {
				addressIndex = uint32(outputCount)
				addressIndexes[*hash] = addressIndex
				lockingScript, err := raddress.LockingScript()
				if err != nil {
					return nil, 0, errors.Wrap(err, "address locking script")
				}
				txSizeOutputs += txbuilder.OutputSize(lockingScript)
				outputCount++
				addressDustLimit = txbuilder.DustLimitForLockingScript(lockingScript, dustFeeRate)
			}

			if instrument.InstrumentType == BSVInstrumentID {
				// Input that spends agent bitcoin transfer output where bitcoin is transfered.
				masterUnlockingSize, err := fees.EstimateUnlockingSize(masterContractScript)
				if err != nil {
					return nil, 0, errors.Wrap(err, "master unlocking size")
				}

				bitcoinTransferUnlockingSize := agent_bitcoin_transfer.ApproveUnlockingSize(masterUnlockingSize)

				txSizeInputs += fees.InputBaseSize + // outpoint + sequence
					wire.VarIntSerializeSize(uint64(bitcoinTransferUnlockingSize)) +
					bitcoinTransferUnlockingSize // unlocking script
				inputCount++
			} else {
				if !exists {
					// Dust will be put in each notification output.
					fundingForTokens[masterContractIndex] += addressDustLimit
				}

				// Use max quantity to ensure over estimation, since we can't use a real value
				// without knowing the resulting balance.
				settleInstrument.Settlements = append(settleInstrument.Settlements,
					&actions.QuantityIndexField{
						Index:    addressIndex,
						Quantity: math.MaxUint64,
					})
			}
		}

		if instrument.InstrumentType == BSVInstrumentID {
			accumulatedSettlementData[i] = previousSettlementData
		} else {
			settlement.Instruments = append(settlement.Instruments, settleInstrument)

			// Calculate settlement size at this contract
			if accumulatedSettlementData[i], err = Serialize(settlement, isTest); err != nil {
				return nil, 0, errors.Wrap(err, "Failed to serialize settlement")
			}

			fullSettlementSize = uint64(len(accumulatedSettlementData[i]))
			previousSettlementData = accumulatedSettlementData[i]
		}

		// An output for each contract's fee address, if they have a contract fee.
		if contractFees[i] > 0 {
			// TODO use contract's actual fee address --ce
			txSizeOutputs += txbuilder.P2PKHOutputSize
			outputCount++
		}
	}

	// Calculate settlement tx funding from parts
	txSizeSettlements := int(txbuilder.OutputBaseSize+
		wire.VarIntSerializeSize(fullSettlementSize)) + int(fullSettlementSize)
	txSizeOutputCount := wire.VarIntSerializeSize(uint64(outputCount))
	txSizeInputCount := wire.VarIntSerializeSize(uint64(inputCount))
	settlementTxSize := txbuilder.BaseTxSize + txSizeInputCount + txSizeInputs + txSizeOutputCount +
		txSizeOutputs + txSizeSettlements
	settlementTxFee := txbuilder.EstimatedFeeValue(uint64(settlementTxSize), float64(feeRate))

	// Sum funding for each contract's output.
	funding := make([]uint64, len(request.Instruments))
	for i := range funding {
		funding[i] = contractFees[i] + fundingForTokens[i]
	}
	funding[masterContractIndex] += uint64(settlementTxFee)

	// Calculate boomerang funding for inter-contract settlement.
	if !multiContract {
		return funding, boomerang, nil
	}

	isMaster := true
	requestContractFees := []*messages.TargetAddressField{}
	previousSettlementData = nil
	txHash := requestTx.TxHash()
	previousLockingScript := requestTx.TxOut[masterContractIndex].LockingScript
	previousDust := txbuilder.DustLimitForLockingScript(previousLockingScript, dustFeeRate)
	for i, instrument := range request.Instruments {
		// No boomerang for bitcoin transfers.
		if instrument.InstrumentType == BSVInstrumentID {
			continue
		}

		if isMaster {
			requestContractFees = append(requestContractFees,
				&messages.TargetAddressField{
					Address:  make([]byte, 22),
					Quantity: contractFees[i],
				})

			previousSettlementData = accumulatedSettlementData[i]

			// The first (master) contract receives original response and needs no funding.
			isMaster = false
			continue
		}

		// Add boomerang funding to send settlement request to this contract
		settleRequest := &messages.SettlementRequest{
			Timestamp:    uint64(now.UnixNano()),
			TransferTxId: requestTx.TxHash().Bytes(),
			ContractFees: requestContractFees,
			Settlement:   previousSettlementData,
		}

		var srBuf bytes.Buffer
		if err := settleRequest.Serialize(&srBuf); err != nil {
			return nil, 0, errors.Wrap(err, "serialize settle request payload")
		}

		settleRequestMessage := &actions.Message{
			SenderIndexes:   []uint32{0},
			ReceiverIndexes: []uint32{0},
			MessageCode:     messages.CodeSettlementRequest,
			MessagePayload:  srBuf.Bytes(),
		}

		settleRequestScript, err := Serialize(settleRequestMessage, isTest)
		if err != nil {
			return nil, 0, errors.Wrap(err, "serialize settlement request action")
		}

		// Wrapped in tx
		settleRequestTx := wire.NewMsgTx(1)

		// Input from previous contract
		scriptSize, err := fees.EstimateUnlockingSize(previousLockingScript)
		if err != nil {
			return nil, 0, errors.Wrap(err, "unlocking script size")
		}

		settleRequestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txHash, 2),
			make([]byte, scriptSize)))

		// Output to this contract
		settleRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte, len(previousLockingScript))))

		// Output for op return
		settleRequestTx.AddTxOut(wire.NewTxOut(0, settleRequestScript))

		txHash = settleRequestTx.TxHash()

		var srTxBuf bytes.Buffer
		if err := settleRequestTx.Serialize(&srTxBuf); err != nil {
			return nil, 0, errors.Wrap(err, "serialize settlement request tx")
		}
		boomerang += uint64(math.Ceil(float64(srTxBuf.Len()) * float64(feeRate)))
		// We don't need to add dust here because the remaining funding is passed along and must
		// include the final dust so will always be at least dust.
		// boomerang += dustLimit // Dust output to next contract

		// Add boomerang funding to send signature request to the previous contract.

		// Subtract out missing signature scripts from previous contracts that haven't signed yet.
		adjustedSettleTxSize := settlementTxSize - (i * int(txbuilder.MaximumP2PKHSigScriptSize))

		sigRequest := &messages.SignatureRequest{
			Timestamp: uint64(now.UnixNano()),
			Payload:   make([]byte, adjustedSettleTxSize), // Full settlement tx
		}

		var sigBuf bytes.Buffer
		if err := sigRequest.Serialize(&sigBuf); err != nil {
			return nil, 0, errors.Wrap(err, "serialize signature request payload")
		}

		sigRequestMessage := &actions.Message{
			SenderIndexes:   []uint32{0},
			ReceiverIndexes: []uint32{0},
			MessageCode:     messages.CodeSignatureRequest,
			MessagePayload:  sigBuf.Bytes(),
		}

		// Wrapped in tx
		sigRequestTx := wire.NewMsgTx(1)

		// Input from previous contract
		sigRequestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txHash, 2),
			make([]byte, scriptSize)))

		// Output to this contract
		sigRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte, len(previousLockingScript))))

		// Output for op return
		sigRequestScript, err := Serialize(sigRequestMessage, isTest)
		if err != nil {
			return nil, 0, errors.Wrap(err, "serialize signature request action")
		}
		sigRequestTx.AddTxOut(wire.NewTxOut(0, sigRequestScript))

		txHash = sigRequestTx.TxHash()

		var sigTxBuf bytes.Buffer
		if err := sigRequestTx.Serialize(&sigTxBuf); err != nil {
			return nil, 0, errors.Wrap(err, "serialize signature request tx")
		}
		boomerang += uint64(math.Ceil(float64(sigTxBuf.Len()) * float64(feeRate)))
		boomerang += previousDust // Dust output to previous contract

		requestContractFees = append(requestContractFees,
			&messages.TargetAddressField{
				Address:  make([]byte, 22),
				Quantity: contractFees[i],
			})

		previousLockingScript := requestTx.TxOut[instrument.ContractIndex].LockingScript
		previousDust = txbuilder.DustLimitForLockingScript(previousLockingScript, dustFeeRate)
	}

	return funding, boomerang, nil
}

func EstimatedConfiscationResponse(requestTx *wire.MsgTx, feeRate, dustFeeRate float32,
	contractFee uint64, isTest bool) (uint64, error) {

	// Find Tokenized Order
	var request *actions.Order
	for _, output := range requestTx.TxOut {
		action, err := Deserialize(output.LockingScript, isTest)
		if err != nil {
			continue
		}

		confiscation, ok := action.(*actions.Order)
		if ok {
			request = confiscation
			break
		}
	}

	if request == nil {
		return 0, errors.New("Tokenized Order OP_RETURN not found")
	}
	if request.ComplianceAction != actions.ComplianceActionConfiscation {
		return 0, errors.New("Not a Tokenized Confiscation Order")
	}

	confiscation := &actions.Confiscation{}
	if err := convert(request, confiscation); err != nil {
		return 0, errors.Wrap(err, "Failed to convert confiscation order to confiscation")
	}

	confiscation.Timestamp = uint64(time.Now().UnixNano())
	confiscation.Quantities = make([]*actions.QuantityIndexField, len(request.TargetAddresses)+1)

	quantityIndex := uint32(0)
	outputsSize := 0
	outputCount := uint64(0)
	outputValue := uint64(0)

	for i, target := range request.TargetAddresses {
		targetAddress, err := bitcoin.DecodeRawAddress(target.Address)
		if err != nil {
			return 0, errors.Wrap(err, "decode target address")
		}

		confiscation.Quantities[quantityIndex] = &actions.QuantityIndexField{
			Index:    quantityIndex,
			Quantity: math.MaxUint64,
		}
		confiscation.DepositQty += target.Quantity

		lockingScript, err := targetAddress.LockingScript()
		if err != nil {
			return 0, errors.Wrapf(err, "target %d locking script", i)
		}
		outputsSize += txbuilder.OutputSize(lockingScript)
		outputValue += txbuilder.DustLimitForLockingScript(lockingScript, dustFeeRate)
		outputCount++
		quantityIndex++
	}

	depositAddress, err := bitcoin.DecodeRawAddress(request.DepositAddress)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to read deposit address")
	}

	confiscation.Quantities[quantityIndex] = &actions.QuantityIndexField{
		Index:    quantityIndex,
		Quantity: math.MaxUint64,
	}

	lockingScript, err := depositAddress.LockingScript()
	if err != nil {
		return 0, errors.Wrap(err, "deposit locking script")
	}
	outputsSize += txbuilder.OutputSize(lockingScript)
	outputValue += txbuilder.DustLimitForLockingScript(lockingScript, dustFeeRate)
	outputCount++

	if contractFee != 0 {
		// Add fee output
		// TODO use contract's actual fee address --ce
		outputsSize += txbuilder.P2PKHOutputSize
		dustLimit := txbuilder.DustLimit(txbuilder.P2PKHOutputSize, dustFeeRate)
		if contractFee < dustLimit {
			outputValue += dustLimit
		} else {
			outputValue += contractFee
		}
		outputCount++
	}

	// Calculate confiscation action output
	confiscationScript, err := Serialize(confiscation, isTest)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to serialize confiscation")
	}
	confiscationScriptSize := len(confiscationScript)
	confiscationOutputSize := txbuilder.OutputBaseSize +
		wire.VarIntSerializeSize(uint64(confiscationScriptSize)) + confiscationScriptSize
	outputCount++

	// Input from contract
	inputsSize, err := txbuilder.InputSize(requestTx.TxOut[0].LockingScript)
	if err != nil {
		return 0, errors.Wrapf(err, "input size")
	}

	txSize := txbuilder.BaseTxSize + inputsSize + wire.VarIntSerializeSize(outputCount) +
		outputsSize + confiscationOutputSize
	txFee := int(math.Ceil(float64(txSize) * float64(feeRate)))

	return uint64(txFee) + outputValue, nil
}

// Convert assigns all available compatible values with matching member names from one object to
//   another.
// The dst object needs to be a pointer so that it can be written to.
// Members of these objects that are "specialized", like a struct containing only a string, need
//   to have json.Marshaler and json.UnMarshaler interfaces implemented.
func convert(src interface{}, dst interface{}) error {
	// Marshal source object to json.
	var data []byte
	var err error
	data, err = json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed json marshal"))
	}

	// Unmarshal json back into destination object.
	err = json.Unmarshal(data, dst)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed json unmarshal"))
	}

	return nil
}
