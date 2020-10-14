package protocol

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/pkg/txbuilder"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/assets"
	"github.com/tokenized/specification/dist/golang/messages"

	"github.com/pkg/errors"
)

// EstimatedResponse calculates information about the contract's response to a request.
//   fees is the sum of all contract related fees including base contract fee, proposal fee, and
///  others. dustLimit is based on the expected P2PKH notification outputs.
// WARNING: This function is really only accurate for ContractOffer and AssetDefinition txs. Other
// actions might be wrong.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// For transfers use EstimatedTransferResponse.
// For contract amendments use EstimatedContractAmendmentResponse.
// For asset modifications use EstimatedAssetModificationResponse.
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
		action, err = Deserialize(output.PkScript, isTest)
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

	case *actions.AssetDefinition:
		var assetCode bitcoin.Hash32
		assetCreation := actions.AssetCreation{
			AssetCode:     assetCode.Bytes(), // Asset code is added by smart contract
			AssetRevision: 0,
			Timestamp:     uint64(now.UnixNano()),
		}
		response = &assetCreation

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
		for _, asset := range request.Assets {

			settleAsset := &actions.AssetSettlementField{
				AssetType: asset.AssetType,
				AssetCode: asset.AssetCode,
			}
			// No settlement needed for bitcoin transfers. Just outputs.
			if asset.AssetType != "BSV" {
				if len(contractScript) == 0 {
					contractScript = requestTx.TxOut[asset.ContractIndex].PkScript
				} else if !bytes.Equal(contractScript, requestTx.TxOut[asset.ContractIndex].PkScript) {
					return 0, 0, errors.New("More than one contract")
				}

				settlement.Assets = append(settlement.Assets, settleAsset)
			}

			// Sig script is probably still empty, so assume each sender is unique and the
			//   address is not reused. So each will get a notification output.
			if asset.AssetType != "BSV" {
				for _, _ = range asset.AssetSenders {
					// Use quantity that is enough for a 3 byte var 128 value, since we can't use a
					//   real value without knowing the resulting balance.
					settleAsset.Settlements = append(settleAsset.Settlements,
						&actions.QuantityIndexField{
							Index:    uint32(outputCount),
							Quantity: 100000,
						})
					outputCount += 1
					if asset.AssetType != "BSV" {
						value += dustLimit // Dust will be put in each notification output.
					}
				}
			}

			for _, receiver := range asset.AssetReceivers {
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
					settleAsset.Settlements = append(settleAsset.Settlements,
						&actions.QuantityIndexField{
							Index:    uint32(outputCount),
							Quantity: 100000,
						})
					outputCount += 1
					if asset.AssetType != "BSV" {
						value += dustLimit // Dust will be put in each notification output.
					}
				}
				if asset.AssetType == "BSV" {
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
func EstimatedContractAmendmentResponse(amendTx *wire.MsgTx, cf *actions.ContractFormation,
	dustFeeRate float32, isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range amendTx.TxOut {
		action, err = Deserialize(output.PkScript, isTest)
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

	cf.ContractRevision = amendment.ContractRevision + 1
	cf.Timestamp = uint64(now.UnixNano())

	for i, amendment := range amendment.Amendments {
		fip, err := actions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if _, err := cf.ApplyAmendment(fip, amendment.Operation, amendment.Data); err != nil {
			return 0, 0, errors.Wrapf(err, "apply amendment %d", i)
		}
	}

	// 1 input from contract
	size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

	// P2PKH dust output to contract, contract fee, and op return output
	if cf.ContractFee > 0 {
		size += wire.VarIntSerializeSize(uint64(3)) + (2 * outputSize)
		value += cf.ContractFee
	} else {
		size += wire.VarIntSerializeSize(uint64(2)) + outputSize
	}
	value += dustLimit

	script, err := Serialize(cf, isTest)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Failed to serialize contract formation envelope")
	}

	// OP_RETURN output size
	size += txbuilder.OutputBaseSize + wire.VarIntSerializeSize(uint64(len(script))) + len(script)

	return size, value, nil
}

// EstimatedAssetModificationResponse calculates information about the contract's response to an
// asset modification request.
// dustFeeRate is the fee rate used to calculate the dust limit for outputs.
// WARNING: This still makes assumptions that some of the addresses are P2PKH.
// Returns:
// estimated size of response tx in bytes.
// estimated funding needed.
// error if there were any
func EstimatedAssetModificationResponse(amendTx *wire.MsgTx, ac *actions.AssetCreation,
	contractFee uint64, dustFeeRate float32, isTest bool) (int, uint64, error) {

	// Find Tokenized OP_RETURN
	var err error
	var action actions.Action
	found := false
	for _, output := range amendTx.TxOut {
		action, err = Deserialize(output.PkScript, isTest)
		if err == nil {
			found = true
			break
		}
	}
	if !found {
		return 0, 0, errors.New("Tokenized OP_RETURN not found")
	}

	amendment, ok := action.(*actions.AssetModification)
	if !ok {
		return 0, 0, errors.New("Action is not a AssetModification")
	}

	outputSize := amendTx.TxOut[0].SerializeSize()
	dustLimit := txbuilder.DustLimitForOutput(amendTx.TxOut[0], dustFeeRate)

	// Build sample response and calculate values based on expected response inputs and outputs.
	size := txbuilder.BaseTxSize
	value := uint64(0)
	now := time.Now()

	ac.AssetRevision = amendment.AssetRevision + 1
	ac.Timestamp = uint64(now.UnixNano())

	var payload assets.Asset
	for i, amendment := range amendment.Amendments {
		fip, err := actions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return 0, 0, errors.Wrapf(err, "parse field index path %d", i)
		}
		if len(fip) == 0 {
			return 0, 0, fmt.Errorf("Amendment %d has no field specified", i)
		}

		if fip[0] == actions.AssetFieldAssetPayload {
			if payload == nil {
				// Get payload object
				payload, err = assets.Deserialize([]byte(ac.AssetType), ac.AssetPayload)
				if err != nil {
					return 0, 0, fmt.Errorf("Asset payload deserialize failed : %s %s",
						ac.AssetType, err)
				}
			}

			_, err = payload.ApplyAmendment(fip[1:], amendment.Operation, amendment.Data)
			if err != nil {
				return 0, 0, errors.Wrapf(err, "apply payload amendment %d", i)
			}

		} else {
			if _, err := ac.ApplyAmendment(fip, amendment.Operation, amendment.Data); err != nil {
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
		ac.AssetPayload = buf.Bytes()
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
//   request.
// fees is a list of the contract fee for the contract corresponding to each contract. The list
//   lines up with the "Assets" list in the transfer.
// dustLimit is the smallest amount of satoshis to make an output valid.
// feeRate is in satoshis per byte.
// WARNING: This function is inaccurate and incomplete!
// Returns
//   estimated funding per "Asset" object in transfer, including contract fees.
//   boomerang funding. needs to be added to a second output to the first contract address.
//     if zero then only one contract is involved and no second output is needed.
//   error if there were any
//
// First contract is master contract. If other contracts are involved it initializes and sends
//   settlement request to next contract. Then after each contract has completed the settlement
//   request, the last signs the tx and sends a signature request back to the previous. After the
//   first contract completes the signature request it broadcasts the completed settlement.
// Each settlement request contains the settlement action and contract fee data.
// Each signature request contains the full settlement tx minus some signatures.
// The first contract output must fund settlement tx miner fee.
// The boomerang must fund all settlement requests and signature requests.
// All other contract outputs can be dust.
// Attempts to use the maximum possible size of each element so the returned values are an
// overestimation, ensuring that a transfer funded in this manner can complete successfully.
func EstimatedTransferResponse(requestTx *wire.MsgTx, dustLimit uint64, feeRate float32,
	fees []uint64, isTest bool) ([]uint64, uint64, error) {

	// Find Tokenized Transfer
	var err error
	var action actions.Action
	var request *actions.Transfer
	var ok bool
	found := false
	for _, output := range requestTx.TxOut {
		action, err = Deserialize(output.PkScript, isTest)
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
	accumulatedSettlementData := make([][]byte, len(request.Assets))
	var previousSettlementData []byte
	fullSettlementSize := uint64(0)
	boomerang := uint64(0) // used for funding inter-contract communication
	now := time.Now()
	// Auditable estimation construction
	fundingForBSV := make([]uint64, len(request.Assets))
	fundingForTokens := make([]uint64, len(request.Assets))
	p2PKHInputCount := 0
	p2PKHOutputCount := 0

	settlement := &actions.Settlement{Timestamp: uint64(now.UnixNano())}

	// Calculate response tx and settlement payload size.
	// Indexes to addreses that already have outputs so they can be reused.
	addressIndexes := make(map[bitcoin.Hash20]uint32)
	var previousContractScript []byte
	multiContract := false
	masterContractIndex := uint32(0xffffffff) // First contract listed
	for _, asset := range request.Assets {
		if asset.AssetType != "BSV" {
			masterContractIndex = asset.ContractIndex
			break
		}
	}

	for i, asset := range request.Assets {
		settleAsset := &actions.AssetSettlementField{
			ContractIndex: uint32(i),
			AssetType:     asset.AssetType,
			AssetCode:     asset.AssetCode,
		}

		// No settlement needed for bitcoin transfers. Just outputs.
		if asset.AssetType != "BSV" {
			if !bytes.Equal(previousContractScript, requestTx.TxOut[asset.ContractIndex].PkScript) {
				multiContract = true
				p2PKHInputCount++ // input from each contract
			}
			previousContractScript = requestTx.TxOut[asset.ContractIndex].PkScript

			// Sig script is probably still empty, so assume each sender is unique and the address
			// is not reused. So each will get a notification output.
			for range asset.AssetSenders {
				// Use max quantity to ensure overestimation, since we can't use a
				//   real value without knowing the resulting balance.
				settleAsset.Settlements = append(settleAsset.Settlements,
					&actions.QuantityIndexField{
						Index:    127,
						Quantity: math.MaxUint64,
					})

				p2PKHOutputCount++
				fundingForTokens[masterContractIndex] += dustLimit // Dust will be put in each notification output.
			}

		} else {
			// Bitcoin senders don't need an output because they don't need a settlement entry.
			for _, sender := range asset.AssetSenders {
				// For the amount that needs to be sent
				fundingForBSV[masterContractIndex] += sender.Quantity
			}
		}

		for _, receiver := range asset.AssetReceivers {
			raddress, err := bitcoin.DecodeRawAddress(receiver.Address)
			if err != nil {
				return nil, 0, errors.Wrap(err, "parsing address")
			}
			hash, err := raddress.Hash()
			if err != nil {
				return nil, 0, errors.Wrap(err, "hashing address")
			}
			addressIndex, exists := addressIndexes[*hash]
			if !exists {
				addressIndex = uint32(p2PKHOutputCount)
				addressIndexes[*hash] = addressIndex
				p2PKHOutputCount++
			}

			addressDustLimit, err := txbuilder.DustLimitForAddress(raddress, feeRate)
			if err != nil {
				return nil, 0, errors.Wrap(err, "address dust limit")
			}

			if asset.AssetType != "BSV" {
				if !exists {
					// Dust will be put in each notification output.
					fundingForTokens[masterContractIndex] += addressDustLimit
				}

				// Use max quantity to ensure over estimation, since we can't use a real value
				// without knowing the resulting balance.
				settleAsset.Settlements = append(settleAsset.Settlements,
					&actions.QuantityIndexField{
						Index:    addressIndex,
						Quantity: math.MaxUint64,
					})
			}
		}

		if asset.AssetType == "BSV" {
			accumulatedSettlementData[i] = previousSettlementData
		} else {
			settlement.Assets = append(settlement.Assets, settleAsset)

			// Calculate settlement size at this contract
			if accumulatedSettlementData[i], err = Serialize(settlement, isTest); err != nil {
				return nil, 0, errors.Wrap(err, "Failed to serialize settlement")
			}

			fullSettlementSize = uint64(len(accumulatedSettlementData[i]))
			previousSettlementData = accumulatedSettlementData[i]
		}

		// An output for each contract's fee address, if they have a contract fee.
		if fees[i] > 0 {
			p2PKHOutputCount++
		}
	}

	// Calculate settlement tx funding from parts
	txSizeInputs := wire.VarIntSerializeSize(uint64(p2PKHInputCount)) +
		p2PKHInputCount*txbuilder.MaximumP2PKHInputSize
	txSizeP2PKHOutputs := p2PKHOutputCount * txbuilder.P2PKHOutputSize
	txSizeSettlements := int(txbuilder.OutputBaseSize+
		wire.VarIntSerializeSize(fullSettlementSize)) + int(fullSettlementSize)
	txSizeOutputCount := wire.VarIntSerializeSize(uint64(p2PKHOutputCount))
	settlementTxSize := txbuilder.BaseTxSize + txSizeInputs + txSizeOutputCount +
		txSizeP2PKHOutputs + txSizeSettlements
	settlementTxFee := int(math.Ceil(float64(settlementTxSize) * float64(feeRate)))

	// Sum funding for each contract's output.
	funding := make([]uint64, len(request.Assets))
	for i := range funding {
		funding[i] = fees[i] + fundingForBSV[i] + fundingForTokens[i]
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
	for i, asset := range request.Assets {
		// No boomerang for bitcoin transfers.
		if asset.AssetType == "BSV" {
			continue
		}

		if isMaster {
			requestContractFees = append(requestContractFees,
				&messages.TargetAddressField{
					Address:  make([]byte, 22),
					Quantity: fees[i],
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
		settleRequestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txHash, 2),
			make([]byte, txbuilder.MaximumP2PKHSigScriptSize)))

		// Output to this contract
		settleRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte,
			txbuilder.P2PKHOutputScriptSize)))

		// Output for op return
		settleRequestTx.AddTxOut(wire.NewTxOut(0, settleRequestScript))

		txHash = settleRequestTx.TxHash()

		var srTxBuf bytes.Buffer
		err = settleRequestTx.Serialize(&srTxBuf)
		if err != nil {
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
			make([]byte, txbuilder.MaximumP2PKHSigScriptSize)))

		// Output to this contract
		sigRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte,
			txbuilder.P2PKHOutputScriptSize)))

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
		boomerang += dustLimit // Dust output to previous contract

		requestContractFees = append(requestContractFees,
			&messages.TargetAddressField{
				Address:  make([]byte, 22),
				Quantity: fees[i],
			})
	}

	return funding, boomerang, nil
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
