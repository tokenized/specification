package cmd

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/messages"
	"github.com/tokenized/specification/dist/golang/protocol"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdDecode = &cobra.Command{
	Use:   "decode hex_or_asm",
	Short: "Decode a transaction or output script (hex or asm)",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Wrong argument count : requires 1")
		}

		b, hexErr := hex.DecodeString(args[0])
		if hexErr == nil {
			if err := decode(b); err != nil {
				return errors.Wrap(err, "decode")
			}

			return nil
		}

		script, scriptErr := bitcoin.StringToScript(args[0])
		if scriptErr == nil {
			if err := displayScript(script, false); err != nil {
				return errors.Wrap(err, "script")
			}

			return nil
		}

		return fmt.Errorf("Not hex or asm script: hex error (%s), asm error (%s)", hexErr,
			scriptErr)
	},
}

func decode(b []byte) error {
	tx := &wire.MsgTx{}
	if err := tx.Deserialize(bytes.NewReader(b)); err == nil {
		if err := displayTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

		return nil
	}

	if err := displayScript(bitcoin.Script(b), true); err != nil {
		return errors.Wrap(err, "script")
	}

	return nil
}

func displayTx(tx *wire.MsgTx) error {
	fmt.Printf("Tx: %s\n", tx)

	for _, txout := range tx.TxOut {
		if err := displayScript(txout.LockingScript, false); err != nil {
			return errors.Wrap(err, "decode")
		}
	}

	return nil
}

func displayScript(script bitcoin.Script, printASM bool) error {
	testAction, err := protocol.Deserialize(script, true)
	if err == nil {
		fmt.Printf("Test Action:\n")
		if err := displayAction(testAction); err != nil {
			return errors.Wrap(err, "test action")
		}

		return nil
	}

	action, err := protocol.Deserialize(script, false)
	if err == nil {
		fmt.Printf("Production Action:\n")
		if err := displayAction(action); err != nil {
			return errors.Wrapf(err, "action")
		}

		return nil
	}

	if printASM {
		fmt.Printf("ASM: %s\n", bitcoin.ScriptToString(script))
	}

	return nil
}

func displayAction(action actions.Action) error {
	js, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return errors.Wrap(err, "json")
	}

	fmt.Printf("%s %s\n", action.TypeName(), js)

	switch a := action.(type) {
	case *actions.InstrumentDefinition:
		if err := displayInstrument(a.InstrumentType, a.InstrumentPayload); err != nil {
			return errors.Wrap(err, "instrument")
		}

	case *actions.InstrumentCreation:
		instrumentID, err := protocol.InstrumentIDForRaw(a.InstrumentType, a.InstrumentCode)
		if err != nil {
			fmt.Printf("Instrument ID: Invalid: %s\n", err)
		} else {
			fmt.Printf("Instrument ID: %s\n", instrumentID)
		}

		if err := displayInstrument(a.InstrumentType, a.InstrumentPayload); err != nil {
			return errors.Wrap(err, "instrument")
		}

	case *actions.Transfer:
		for i, instrumentTransfer := range a.Instruments {
			instrumentID, err := protocol.InstrumentIDForTransfer(instrumentTransfer)
			if err != nil {
				fmt.Printf("Instrument ID %d: Invalid: %s\n", i, err)
			} else {
				fmt.Printf("Instrument ID %d: %s\n", i, instrumentID)
			}
		}

	case *actions.Settlement:
		for i, instrumentSettlement := range a.Instruments {
			instrumentID, err := protocol.InstrumentIDForSettlement(instrumentSettlement)
			if err != nil {
				fmt.Printf("Instrument ID %d: Invalid: %s\n", i, err)
			} else {
				fmt.Printf("Instrument ID %d: %s\n", i, instrumentID)
			}
		}

	case *actions.Message:
		if err := displayMessage(a.MessageCode, a.MessagePayload); err != nil {
			return errors.Wrap(err, "message")
		}

	case *actions.Rejection:
		rejectData := actions.RejectionsData(a.RejectionCode)
		if rejectData != nil {
			fmt.Printf("Reject Label: %s\n", rejectData.Label)
		}
	}

	return nil
}

func displayInstrument(typ string, payload []byte) error {
	p, err := instruments.Deserialize([]byte(typ), payload)
	if err != nil {
		return errors.Wrap(err, "deserialize")
	}

	js, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return errors.Wrap(err, "json")
	}

	fmt.Printf("%s %s\n", p.TypeName(), js)
	return nil
}

func displayMessage(code uint32, payload []byte) error {
	p, err := messages.Deserialize(code, payload)
	if err != nil {
		return errors.Wrap(err, "deserialize")
	}

	js, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return errors.Wrap(err, "json")
	}

	fmt.Printf("%s %s\n", p.TypeName(), js)

	switch m := p.(type) {
	case *messages.RevertedTx:
		tx := &wire.MsgTx{}
		if err := tx.Deserialize(bytes.NewReader(m.Transaction)); err != nil {
			return errors.Wrap(err, "deserialize tx")
		}

		if err := displayTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

	case *messages.Offer:
		tx := &wire.MsgTx{}
		if err := tx.Deserialize(bytes.NewReader(m.Payload)); err != nil {
			return errors.Wrap(err, "deserialize tx")
		}

		if err := displayTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

	case *messages.SignatureRequest:
		tx := &wire.MsgTx{}
		if err := tx.Deserialize(bytes.NewReader(m.Payload)); err != nil {
			return errors.Wrap(err, "deserialize tx")
		}

		if err := displayTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}
	}

	return nil
}
