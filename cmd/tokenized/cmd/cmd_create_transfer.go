package cmd

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/txbuilder"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/print"
	"github.com/tokenized/specification/dist/golang/protocol"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdCreateTransfer = &cobra.Command{
	Use:   "create_transfer sender_key outpoint outpoint_value instrument_id recipient_address recipient_token_quantity contract_address contract_fee change_address",
	Short: "Creates a transaction that contains an InstrumentDefinition.",
	Long: `admin_key: WIF format key for contract administrator.
outpoint: UTXO to spend. P2PKH output to admin key.
  Example: "4ff6ef3dba153837bf1f0bc789e953be37255d5d515695b40a7d47bfd8ecd7e9:1"
outpoint_value: Satoshi value of outpoint being spent.
instrument_id: Identifier for instrument to send.
recipient_address: Address of token recipient.
recipient_token_quantity: Quantity of tokens to send.
contract_address: The bitcoin address of the smart contract agent.
contract_fee: the fee from the contract formation.
change_address: The bitcoin address to return any remaining bitcoin to.`,
	Args: cobra.ExactArgs(9),
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

		instrumentType, instrumentCode, err := protocol.DecodeInstrumentID(args[3])
		if err != nil {
			return errors.Wrap(err, "instrument id")
		}

		recipientAddress, err := bitcoin.DecodeAddress(args[4])
		if err != nil {
			return errors.Wrap(err, "recipient address")
		}

		recipientRA := bitcoin.NewRawAddressFromAddress(recipientAddress)

		recipientQuantity, err := strconv.ParseUint(args[5], 10, 64)
		if err != nil {
			return errors.Wrap(err, "recipient quantity")
		}

		transfer := &actions.Transfer{
			Instruments: []*actions.InstrumentTransferField{
				{
					ContractIndex:  0,
					InstrumentType: instrumentType,
					InstrumentCode: instrumentCode.Bytes(),
					InstrumentSenders: []*actions.QuantityIndexField{
						{
							Quantity: recipientQuantity,
							Index:    0, // first input
						},
					},
					InstrumentReceivers: []*actions.InstrumentReceiverField{
						{
							Address:  recipientRA.Bytes(),
							Quantity: recipientQuantity,
						},
					},
				},
			},
		}

		if err := transfer.Validate(); err != nil {
			return errors.Wrap(err, "validate transfer")
		}

		isTest, err := IsTest()
		if err != nil {
			return errors.Wrap(err, "is test")
		}

		transferScript, err := protocol.Serialize(transfer, isTest)
		if err != nil {
			return errors.Wrap(err, "serialize transfer")
		}

		contractAddress, err := bitcoin.DecodeAddress(args[6])
		if err != nil {
			return errors.Wrap(err, "contract address")
		}

		contractRA := bitcoin.NewRawAddressFromAddress(contractAddress)
		contractLockingScript, err := contractRA.LockingScript()
		if err != nil {
			return errors.Wrap(err, "contract locking script")
		}

		contractFee, err := strconv.ParseUint(args[7], 10, 64)
		if err != nil {
			return errors.Wrap(err, "contract fee")
		}

		changeAddress, err := bitcoin.DecodeAddress(args[8])
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

		if err := tx.AddOutput(transferScript, 0, false, false); err != nil {
			return errors.Wrap(err, "add transfer output")
		}

		if err := tx.SetChangeLockingScript(changeLockingScript, ""); err != nil {
			return errors.Wrap(err, "set change locking script")
		}

		// Estimate contract request funding
		fundings, _, err := protocol.EstimatedTransferResponse(tx.MsgTx,
			[]bitcoin.Script{adminLockingScript}, float32(feeRate), float32(dustFeeRate),
			[]uint64{contractFee}, isTest)
		if err != nil {
			return errors.Wrap(err, "estimate funding")
		}

		fmt.Printf("Estimated response funding: %d\n", fundings[0])

		// Add response tx fee
		if err := tx.AddValueToOutput(0, fundings[0]); err != nil {
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
