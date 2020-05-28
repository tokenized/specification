package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/txbuilder"
	"github.com/tokenized/pkg/wire"

	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/messages"

	"github.com/pkg/errors"
)

// EstimatedResponse calculates information about the contract's response to a request.
//   fees is the sum of all contract related fees including base contract fee, proposal fee, and
///  others. dustLimit is based on the expected P2PKH notification outputs.
// WARNING: This function is inaccurate and incomplete!
// Returns
//   estimated size of response tx in bytes.
//   estimated funding needed, not including contract/proposal fees.
//   error if there were any
func EstimatedResponse(requestTx *wire.MsgTx, inputIndex int, dustLimit, fees uint64, isTest bool) (int, uint64, error) {
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
		contractFormation := actions.ContractFormation{Timestamp: uint64(now.UnixNano())}
		response = &contractFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, contract fee, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

	case *actions.ContractAmendment:
		contractFormation := actions.ContractFormation{Timestamp: uint64(now.UnixNano())}
		response = &contractFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, contract fee, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

		// TODO Need last asset creation to know size of response. Determine change in size by applying amendments.

	case *actions.AssetDefinition:
		assetCreation := actions.AssetCreation{Timestamp: uint64(now.UnixNano())}
		response = &assetCreation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

	case *actions.AssetModification:
		assetCreation := actions.AssetCreation{Timestamp: uint64(now.UnixNano())}
		response = &assetCreation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		// P2PKH dust output to contract, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

		// TODO Need last asset creation to know size of response. Determine change in size by applying amendments.

	case *actions.Transfer:
		settlement := &actions.Settlement{Timestamp: uint64(now.UnixNano())}
		response = settlement

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.MaximumP2PKHInputSize

		outputCount := 0
		if fees > 0 {
			// TODO Update to support multi-contract transfers where there will be more than one fee.
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
				size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
				value += fees
			} else {
				size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
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
				size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
				value += fees
			} else {
				size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
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
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
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
	//   and outputs.
	settlementSize := make([]uint64, len(request.Assets))
	boomerang := uint64(0) // used for funding inter-contract communication
	now := time.Now()
	// Auditable estimation construction
	fundingForBSV := make([]uint64, len(request.Assets))
	fundingForTokens := make([]uint64, len(request.Assets))
	p2PKHInputCount := 0

	// 1 input from contract
	p2PKHInputCount++

	p2PKHOutputCount := 0
	for _, fee := range fees {
		if fee > 0 {
			p2PKHOutputCount++
		}
	}

	settlement := &actions.Settlement{Timestamp: uint64(now.UnixNano())}

	// Calculate response tx and settlement payload size.
	used := make(map[bitcoin.Hash20]bool)
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
			AssetType: asset.AssetType,
			AssetCode: asset.AssetCode,
		}

		// No settlement needed for bitcoin transfers. Just outputs.
		if asset.AssetType != "BSV" {
			if !bytes.Equal(previousContractScript, requestTx.TxOut[asset.ContractIndex].PkScript) {
				multiContract = true
			}
			previousContractScript = requestTx.TxOut[asset.ContractIndex].PkScript

			// Bitcoin senders don't get an output
			// Sig script is probably still empty, so assume each sender is unique and the address is
			//   not reused. So each will get a notification output.
			for range asset.AssetSenders {
				// Use max quantity to ensure overestimation, since we can't use a
				//   real value without knowing the resulting balance.
				settleAsset.Settlements = append(settleAsset.Settlements,
					&actions.QuantityIndexField{
						Index:    uint32(p2PKHOutputCount),
						Quantity: math.MaxUint64,
					})

				p2PKHOutputCount++
				fundingForTokens[i] += dustLimit // Dust will be put in each notification output.
			}

		} else {
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
			_, exists := used[*hash]

			if asset.AssetType == "BSV" {
				if !exists {
					p2PKHOutputCount++
					// Do not add funding since we are receiving BSV.
				}
				used[*hash] = false

			} else {
				// Use max quantity to ensure overestimation, since we can't use a
				//   real value without knowing the resulting balance.
				settleAsset.Settlements = append(settleAsset.Settlements,
					&actions.QuantityIndexField{
						Index:    uint32(p2PKHOutputCount),
						Quantity: math.MaxUint64,
					})
				if !exists {
					used[*hash] = true
					p2PKHOutputCount++

					// Dust will be put in each notification output.
					fundingForTokens[masterContractIndex] += dustLimit
				}
			}
		}

		if len(asset.AssetReceivers) == 0 { // Needs to be at least 1 receiver
			if asset.AssetType != "BSV" {
				fundingForTokens[masterContractIndex] += dustLimit // Needed for output

				settleAsset.Settlements = append(settleAsset.Settlements,
					&actions.QuantityIndexField{
						Index:    uint32(p2PKHOutputCount),
						Quantity: math.MaxUint64,
					})

			}
			p2PKHOutputCount++
		}

		if asset.AssetType != "BSV" {
			settlement.Assets = append(settlement.Assets, settleAsset)

			// Calculate settlement size at this contract
			script, err := Serialize(settlement, isTest)
			if err != nil {
				return nil, 0, errors.Wrap(err, "Failed to serialize settlement")
			}
			settlementSize[i] = uint64(len(script))
		}
	}

	// Calculate funding from parts
	txSizeInputs := wire.VarIntSerializeSize(uint64(p2PKHInputCount)) +
		p2PKHInputCount*txbuilder.MaximumP2PKHInputSize
	txSizeP2PKHOutputs := p2PKHOutputCount * txbuilder.P2PKHOutputSize
	txSizeSettlements := 0
	for _, size := range settlementSize {
		if size > 0 {
			txSizeSettlements += int(txbuilder.OutputBaseSize+wire.VarIntSerializeSize(size)) +
				int(size)
		}
	}
	txSizeOutputCount := wire.VarIntSerializeSize(uint64(p2PKHOutputCount + 1)) // +1 for contract
	size := txbuilder.BaseTxSize + txSizeInputs + txSizeOutputCount + txSizeP2PKHOutputs + txSizeSettlements
	txFee := int(math.Ceil(float64(size) * float64(feeRate)))

	funding := make([]uint64, len(request.Assets))
	for i := range funding {
		funding[i] = fees[i] + fundingForBSV[i] + fundingForTokens[i]
	}
	funding[masterContractIndex] += uint64(txFee)

	// Calculate boomerang funding
	isMaster := true
	requestContractFees := []*messages.TargetAddressField{}
	if multiContract {
		for i, asset := range request.Assets {
			// No boomerang for bitcoin transfers.
			if asset.AssetType == "BSV" {
				continue
			}

			if isMaster { // first contract (master) receives original response
				requestContractFees = append(requestContractFees,
					&messages.TargetAddressField{
						Address:  make([]byte, 22),
						Quantity: fees[i],
					})
				isMaster = false
				continue
			}

			// Add boomerang funding to send settlement request to this contract
			settleRequest := &messages.SettlementRequest{
				Timestamp:    uint64(now.UnixNano()),
				TransferTxId: make([]byte, 32),
				ContractFees: requestContractFees,
				Settlement:   make([]byte, settlementSize[i]),
			}
			var srBuf bytes.Buffer
			err = settleRequest.Serialize(&srBuf)
			if err != nil {
				return nil, 0, errors.Wrap(err, "serialize settle request payload")
			}

			settleRequestMessage := &actions.Message{
				SenderIndexes:   []uint32{0},
				ReceiverIndexes: []uint32{0},
				MessageCode:     messages.CodeSettlementRequest,
				MessagePayload:  srBuf.Bytes(),
			}

			// Wrapped in tx
			settleRequestTx := wire.NewMsgTx(1)

			// Input from previous contract
			hash, err := bitcoin.NewHash32(settleRequest.TransferTxId)
			if err != nil {
				return nil, 0, errors.Wrap(err, "create hash")
			}
			settleRequestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(hash, 2),
				make([]byte, txbuilder.MaximumP2PKHInputSize)))

			// Output to this contract
			settleRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte, txbuilder.P2PKHOutputSize)))

			// Output for op return
			settleRequestScript, err := Serialize(settleRequestMessage, isTest)
			if err != nil {
				return nil, 0, errors.Wrap(err, "serialize settlement request action")
			}
			settleRequestTx.AddTxOut(wire.NewTxOut(0, settleRequestScript))

			var srTxBuf bytes.Buffer
			err = settleRequestTx.Serialize(&srTxBuf)
			if err != nil {
				return nil, 0, errors.Wrap(err, "serialize settlement request tx")
			}
			boomerang += uint64(float32(srTxBuf.Len()) * feeRate)

			// Add boomerang funding to send signature request back from this contract
			sigRequest := &messages.SignatureRequest{
				Timestamp: uint64(now.UnixNano()),
				Payload:   make([]byte, size), // Full settlement tx
			}
			var sigBuf bytes.Buffer
			err = sigRequest.Serialize(&sigBuf)
			if err != nil {
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
			sigRequestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(hash, 2),
				make([]byte, txbuilder.MaximumP2PKHInputSize)))

			// Output to this contract
			sigRequestTx.AddTxOut(wire.NewTxOut(100000, make([]byte, txbuilder.P2PKHOutputSize)))

			// Output for op return
			sigRequestScript, err := Serialize(sigRequestMessage, isTest)
			if err != nil {
				return nil, 0, errors.Wrap(err, "serialize signature request action")
			}
			sigRequestTx.AddTxOut(wire.NewTxOut(0, sigRequestScript))

			var sigTxBuf bytes.Buffer
			err = sigRequestTx.Serialize(&sigTxBuf)
			if err != nil {
				return nil, 0, errors.Wrap(err, "serialize signature request tx")
			}
			boomerang += uint64(float32(sigTxBuf.Len()) * feeRate)

			requestContractFees = append(requestContractFees,
				&messages.TargetAddressField{
					Address:  make([]byte, 22),
					Quantity: fees[i],
				})
		}
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
