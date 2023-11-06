package print

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/messages"
	"github.com/tokenized/specification/dist/golang/permissions"
	"github.com/tokenized/specification/dist/golang/protocol"

	"github.com/pkg/errors"
)

func PrintBytes(b []byte) error {
	tx := &wire.MsgTx{}
	if err := tx.Deserialize(bytes.NewReader(b)); err == nil && tx.Version == 1 || tx.Version == 2 {
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
	case *actions.ContractAmendment:
		if err := PrintAmendments(a.Amendments); err != nil {
			return errors.Wrap(err, "amendments")
		}

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

	case *actions.InstrumentModification:
		instrumentID, err := protocol.InstrumentIDForRaw(a.InstrumentType, a.InstrumentCode)
		if err != nil {
			fmt.Printf("Instrument ID: Invalid: %s\n", err)
		} else {
			fmt.Printf("Instrument ID: %s\n", instrumentID)
		}

		if err := PrintAmendments(a.Amendments); err != nil {
			return errors.Wrap(err, "amendments")
		}

	case *actions.Transfer:
		for i, instrumentTransfer := range a.Instruments {
			instrumentID, err := protocol.InstrumentIDForTransfer(instrumentTransfer)
			if err != nil {
				fmt.Printf("Instrument ID %d: Invalid: %s\n", i, err)
			} else {
				fmt.Printf("Instrument ID %d: %s\n", i, instrumentID)
			}

			for _, receiver := range instrumentTransfer.InstrumentReceivers {
				ra, err := bitcoin.DecodeRawAddress(receiver.Address)
				if err != nil {
					continue
				}

				fmt.Printf("Receiver:\n")
				fmt.Printf(" Address: %s\n", bitcoin.NewAddressFromRawAddress(ra, bitcoin.MainNet))

				ls, err := ra.LockingScript()
				if err == nil {
					fmt.Printf(" Script Hash: %s\n", bitcoin.Hash32(sha256.Sum256(ls)))
				}
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

func PrintAmendments(amendments []*actions.AmendmentField) error {
	for i, amendment := range amendments {
		fmt.Printf("Amendment %d:\n", i)
		switch amendment.Operation {
		case actions.AmendmentOperationModify:
			fmt.Printf("  Operation : modify\n")
		case actions.AmendmentOperationAddElement:
			fmt.Printf("  Operation : add\n")
		case actions.AmendmentOperationRemoveElement:
			fmt.Printf("  Operation : remove\n")
		}

		fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
		if err != nil {
			return errors.Wrap(err, "field index path")
		}

		js, _ := json.Marshal(fip)
		fmt.Printf("  Field Index Path : %s\n", js)

		fmt.Printf("  Data Hex : 0x%x\n", amendment.Data)

		if isText(amendment.Data) {
			fmt.Printf("  Data Text : %s\n", string(amendment.Data))
		}

		buf := bytes.NewBuffer(amendment.Data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err == nil {
			fmt.Printf("  Data Number : %d\n", value)
		}
	}

	return nil
}

func isText(bs []byte) bool {
	for _, b := range bs {
		if b < 0x20 { // ' ' space character
			return false
		}

		if b > 0x7e { // '~' tilde character
			return false
		}
	}

	return true
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

	case *messages.SettlementRequest:
		settlement, err := protocol.Deserialize(m.Settlement, true)
		if err != nil {
			settlement, err = protocol.Deserialize(m.Settlement, false)
			if err != nil {
				return errors.Wrap(err, "deserialize settlement")
			}
		}

		if err := PrintAction(settlement); err != nil {
			return errors.Wrap(err, "tx")
		}
	}

	return nil
}
