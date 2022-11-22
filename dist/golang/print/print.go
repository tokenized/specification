package print

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/messages"
	"github.com/tokenized/specification/dist/golang/protocol"

	"github.com/pkg/errors"
)

func PrintBytes(b []byte) error {
	tx := &wire.MsgTx{}
	if err := tx.Deserialize(bytes.NewReader(b)); err == nil {
		if err := PrintTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

		return nil
	}

	if err := PrintScript(bitcoin.Script(b), true); err != nil {
		return errors.Wrap(err, "script")
	}

	return nil
}

func PrintTx(tx *wire.MsgTx) error {
	fmt.Printf("Tx: %s\n", tx)

	for _, txout := range tx.TxOut {
		if err := PrintScript(txout.LockingScript, false); err != nil {
			return errors.Wrap(err, "decode")
		}
	}

	return nil
}

func PrintActions(tx *wire.MsgTx) error {
	for _, txout := range tx.TxOut {
		testAction, err := protocol.Deserialize(txout.LockingScript, true)
		if err == nil {
			fmt.Printf("Test Action:\n")
			if err := PrintAction(testAction); err != nil {
				return errors.Wrap(err, "test action")
			}

			continue
		}

		action, err := protocol.Deserialize(txout.LockingScript, false)
		if err == nil {
			fmt.Printf("Production Action:\n")
			if err := PrintAction(action); err != nil {
				return errors.Wrapf(err, "action")
			}

			continue
		}
	}

	return nil
}

func PrintScript(script bitcoin.Script, printASM bool) error {
	testAction, err := protocol.Deserialize(script, true)
	if err == nil {
		fmt.Printf("Test Action:\n")
		if err := PrintAction(testAction); err != nil {
			return errors.Wrap(err, "test action")
		}

		return nil
	}

	action, err := protocol.Deserialize(script, false)
	if err == nil {
		fmt.Printf("Production Action:\n")
		if err := PrintAction(action); err != nil {
			return errors.Wrapf(err, "action")
		}

		return nil
	}

	if printASM {
		fmt.Printf("ASM: %s\n", bitcoin.ScriptToString(script))
	}

	return nil
}

func PrintAction(action actions.Action) error {
	js, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return errors.Wrap(err, "json")
	}

	fmt.Printf("%s %s\n", action.TypeName(), js)

	switch a := action.(type) {
	case *actions.InstrumentDefinition:
		if err := PrintInstrument(a.InstrumentType, a.InstrumentPayload); err != nil {
			return errors.Wrap(err, "instrument")
		}

	case *actions.InstrumentCreation:
		instrumentID, err := protocol.InstrumentIDForRaw(a.InstrumentType, a.InstrumentCode)
		if err != nil {
			fmt.Printf("Instrument ID: Invalid: %s\n", err)
		} else {
			fmt.Printf("Instrument ID: %s\n", instrumentID)
		}

		if err := PrintInstrument(a.InstrumentType, a.InstrumentPayload); err != nil {
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
		if err := PrintMessage(a.MessageCode, a.MessagePayload); err != nil {
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

func PrintInstrument(typ string, payload []byte) error {
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

func PrintMessage(code uint32, payload []byte) error {
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

		if err := PrintTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

	case *messages.Offer:
		tx := &wire.MsgTx{}
		if err := tx.Deserialize(bytes.NewReader(m.Payload)); err != nil {
			return errors.Wrap(err, "deserialize tx")
		}

		if err := PrintTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}

	case *messages.SignatureRequest:
		tx := &wire.MsgTx{}
		if err := tx.Deserialize(bytes.NewReader(m.Payload)); err != nil {
			return errors.Wrap(err, "deserialize tx")
		}

		if err := PrintTx(tx); err != nil {
			return errors.Wrap(err, "tx")
		}
	}

	return nil
}
