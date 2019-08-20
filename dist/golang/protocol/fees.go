package protocol

import (
	"encoding/json"
	"fmt"

	"github.com/tokenized/smart-contract/pkg/txbuilder"
	"github.com/tokenized/smart-contract/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"

	"github.com/pkg/errors"
)

// EstimatedResponse calculates information about the contract's response to a request.
//   fees is the sum of all contract related fees including base contract fee, proposal fee, and others.
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

	switch request := action.(type) {
	case *actions.ContractOffer:
		contractFormation := actions.ContractFormation{}
		response = &contractFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

		// P2PKH dust output to contract, contract fee, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

	case *actions.ContractAmendment:
		contractFormation := actions.ContractFormation{}
		response = &contractFormation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

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
		assetCreation := actions.AssetCreation{}
		response = &assetCreation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

		// P2PKH dust output to contract, and op return output
		if fees > 0 {
			size += wire.VarIntSerializeSize(uint64(3)) + (3 * txbuilder.P2PKHOutputSize)
			value += fees
		} else {
			size += wire.VarIntSerializeSize(uint64(2)) + (2 * txbuilder.P2PKHOutputSize)
		}
		value += dustLimit

	case *actions.AssetModification:
		assetCreation := actions.AssetCreation{}
		response = &assetCreation

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

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
		settlement := actions.Settlement{}
		response = &settlement

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

		// P2PKH dust output to contract, and op return output
		outputCount := 2
		size += 2 * txbuilder.P2PKHOutputSize
		value += dustLimit
		if fees > 0 {
			outputCount += 1
			size += txbuilder.P2PKHOutputSize
			value += fees
		}

		// Add receiver outputs needed
		used := make(map[[20]byte]bool)
		for _, asset := range request.Assets {
			for _, _ = range asset.AssetSenders {
				// hash, err := txbuilder.PubKeyHashFromP2PKHSigScript(requestTx.TxIn[sender.Index].SignatureScript)
				// if err != nil {
				// return 0, 0, errors.Wrap(err, "Invalid request input")
				// }
				// var fixedHash [20]byte
				// copy(fixedHash[:], hash)
				// _, exists := used[fixedHash]
				// if !exists {
				// used[fixedHash] = true
				outputCount += 1
				// }
			}

			for _, receiver := range asset.AssetReceivers {
				var fixedHash [20]byte
				copy(fixedHash[:], receiver.Address)
				_, exists := used[fixedHash]
				if !exists {
					used[fixedHash] = true
					outputCount += 1
				}
			}
		}

		size += wire.VarIntSerializeSize(uint64(outputCount)) + (outputCount * txbuilder.P2PKHOutputSize)

	case *actions.Proposal:
		if inputIndex == 0 {
			// First input funds vote (initiation) message
			vote := actions.Vote{}
			response = &vote

			// 1 input from contract
			size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

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
			voteResult := actions.Result{}
			response = &voteResult

			voteResult.OptionTally = make([]uint64, len(request.VoteOptions))
			voteResult.Result = " "
			for len(voteResult.Result) < len(request.VoteOptions) {
				voteResult.Result += " "
			}

			// 1 input from contract
			size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

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
		counted := actions.BallotCounted{}
		response = &counted

		// 1 input from contract
		size += wire.VarIntSerializeSize(uint64(1)) + txbuilder.EstimatedInputSize

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
