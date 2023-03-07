package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/print"
	"github.com/tokenized/specification/dist/golang/protocol"
	"github.com/tokenized/txbuilder"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdCreateInstrument = &cobra.Command{
	Use:   "create_instrument admin_key outpoint outpoint_value definition_json_filename payload_json_filename contract_address contract_fee change_address",
	Short: "Creates a transaction that contains an InstrumentDefinition.",
	Long: `admin_key: WIF format key for contract administrator.
outpoint: UTXO to spend. P2PKH output to admin key.
  Example: "4ff6ef3dba153837bf1f0bc789e953be37255d5d515695b40a7d47bfd8ecd7e9:1"
outpoint_value: Satoshi value of outpoint being spent.
definition_json_filename: File name containing JSON format InstrumentDefinition action. Leave the payload empty.
payload_json_filename: File name containing JSON format instrument payload depending on instrument type.
contract_address: The bitcoin address of the smart contract agent.
contract_fee: the fee from the contract formation.
change_address: The bitcoin address to return any remaining bitcoin to.`,
	Args: cobra.ExactArgs(8),
	RunE: func(c *cobra.Command, args []string) error {
		adminKey, err := bitcoin.KeyFromStr(args[0])
		if err != nil {
			return errors.Wrap(err, "admin key")
		}

		adminLockingScript, err := adminKey.LockingScript()
		if err != nil {
			return errors.Wrap(err, "admin key locking script")
		}

		outpoint, err := wire.OutPointFromStr(args[1])
		if err != nil {
			return errors.Wrap(err, "outpoint")
		}

		outpointValue, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return errors.Wrap(err, "outpoint value")
		}

		jsFile, err := os.Open(args[3])
		if err != nil {
			return errors.Wrap(err, "open definition file")
		}
		defer jsFile.Close()

		jsFileSize, err := jsFile.Seek(0, 2)
		if err != nil {
			return errors.Wrap(err, "seek definition file end")
		}

		if _, err := jsFile.Seek(0, 0); err != nil {
			return errors.Wrap(err, "seek definition file begin")
		}

		js := make([]byte, jsFileSize)
		if readSize, err := jsFile.Read(js); err != nil {
			return errors.Wrap(err, "read file")
		} else if readSize != int(jsFileSize) {
			return fmt.Errorf("Failed to read full definition file: read %d, size %d", readSize,
				jsFileSize)
		}

		instrumentDefinition := &actions.InstrumentDefinition{}
		if err := json.Unmarshal(js, instrumentDefinition); err != nil {
			return errors.Wrap(err, "unmarshal instrument definition")
		}

		jsFile, err = os.Open(args[4])
		if err != nil {
			return errors.Wrap(err, "open payload file")
		}
		defer jsFile.Close()

		jsFileSize, err = jsFile.Seek(0, 2)
		if err != nil {
			return errors.Wrap(err, "seek payload file end")
		}

		if _, err = jsFile.Seek(0, 0); err != nil {
			return errors.Wrap(err, "seek payload file begin")
		}

		js = make([]byte, jsFileSize)
		if readSize, err := jsFile.Read(js); err != nil {
			return errors.Wrap(err, "read payload file")
		} else if readSize != int(jsFileSize) {
			return fmt.Errorf("Failed to read full payload file: read %d, size %d", readSize, jsFileSize)
		}

		payload := instruments.NewInstrumentFromCode(instrumentDefinition.InstrumentType)

		if err := json.Unmarshal(js, payload); err != nil {
			return errors.Wrap(err, "unmarshal instrument payload")
		}

		payloadBytes, err := payload.Bytes()
		if err != nil {
			return errors.Wrap(err, "serialize payload")
		}

		instrumentDefinition.InstrumentPayload = payloadBytes

		if err := instrumentDefinition.Validate(); err != nil {
			return errors.Wrap(err, "validate instrument definition")
		}

		isTest, err := IsTest()
		if err != nil {
			return errors.Wrap(err, "is test")
		}

		instrumentDefinitionScript, err := protocol.Serialize(instrumentDefinition, isTest)
		if err != nil {
			return errors.Wrap(err, "serialize instrument definition")
		}

		contractAddress, err := bitcoin.DecodeAddress(args[5])
		if err != nil {
			return errors.Wrap(err, "contract address")
		}

		contractRA := bitcoin.NewRawAddressFromAddress(contractAddress)
		contractLockingScript, err := contractRA.LockingScript()
		if err != nil {
			return errors.Wrap(err, "contract locking script")
		}

		contractFee, err := strconv.ParseUint(args[6], 10, 64)
		if err != nil {
			return errors.Wrap(err, "contract fee")
		}

		changeAddress, err := bitcoin.DecodeAddress(args[7])
		if err != nil {
			return errors.Wrap(err, "change address")
		}

		changeRA := bitcoin.NewRawAddressFromAddress(changeAddress)
		changeLockingScript, err := changeRA.LockingScript()
		if err != nil {
			return errors.Wrap(err, "change locking script")
		}

		feeRate, err := FeeRate()
		if err != nil {
			return errors.Wrap(err, "fee rate")
		}

		dustFeeRate, err := DustFeeRate()
		if err != nil {
			return errors.Wrap(err, "dust fee rate")
		}

		tx := txbuilder.NewTxBuilder(float32(feeRate), float32(dustFeeRate))

		if err := tx.AddInput(*outpoint, adminLockingScript, outpointValue); err != nil {
			return errors.Wrap(err, "add input")
		}

		if err := tx.AddOutput(contractLockingScript, 0, false, true); err != nil {
			return errors.Wrap(err, "add contract output")
		}

		if err := tx.AddOutput(instrumentDefinitionScript, 0, false, false); err != nil {
			return errors.Wrap(err, "add instrument definition output")
		}

		if err := tx.SetChangeLockingScript(changeLockingScript, ""); err != nil {
			return errors.Wrap(err, "set change locking script")
		}

		// Estimate contract request funding
		dustLimit := txbuilder.DustLimitForLockingScript(contractLockingScript,
			float32(dustFeeRate))
		size, funding, err := protocol.EstimatedResponse(tx.MsgTx, 0, dustLimit,
			contractFee, isTest)
		if err != nil {
			return errors.Wrap(err, "estimate funding")
		}

		responseFee := txbuilder.EstimatedFeeValue(uint64(size), feeRate)
		fmt.Printf("Estimated contract fee + response dust outputs: %d\n", funding)
		fmt.Printf("Estimated response tx size: %d\n", size)
		fmt.Printf("Estimated response tx fee: %d\n", responseFee)

		// Add response tx fee
		funding += responseFee

		if err := tx.AddValueToOutput(0, funding); err != nil {
			return errors.Wrap(err, "add funding value")
		}

		if _, err := tx.Sign([]bitcoin.Key{adminKey}); err != nil {
			return errors.Wrap(err, "sign tx")
		}

		finalFee := tx.Fee()
		finalSize := tx.MsgTx.SerializeSize()
		finalRate := float64(finalFee) / float64(finalSize)
		fmt.Printf("Request Tx Fee: %d sats for %d bytes (%.4f sats/byte)\n", finalFee, finalSize,
			finalRate)

		fmt.Printf("Request ")
		print.PrintTx(tx.MsgTx)

		txBuf := &bytes.Buffer{}
		if err := tx.MsgTx.Serialize(txBuf); err != nil {
			return errors.Wrap(err, "serialize tx")
		}

		fmt.Printf("Hex Tx: %x\n", txBuf.Bytes())

		return nil
	},
}
