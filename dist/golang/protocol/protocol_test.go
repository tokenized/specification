package protocol

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/actions"
)

func Test_SerializeAction(t *testing.T) {
	var tests = []struct {
		action actions.Action
		code   string
	}{
		{
			action: &actions.ContractOffer{
				ContractName: "Test Contract Name",
			},
			code: actions.CodeContractOffer,
		},
		{ // Empty action
			action: &actions.ContractOffer{},
			code:   actions.CodeContractOffer,
		},
	}

	for i, test := range tests {
		script, err := Serialize(test.action, true)
		if err != nil {
			t.Fatalf("Failed to serialize action : %s", err)
		}

		t.Logf("Script %d : %s", i, bitcoin.Script(script))

		action, err := Deserialize(script, true)
		if err != nil {
			t.Fatalf("Failed to deserialize action : %s", err)
		}

		js, _ := json.MarshalIndent(action, "", "  ")
		t.Logf("Action %d : %s : %s", i, action.Code(), js)

		if action.Code() != test.code {
			t.Errorf("Wrong action code : got %s, want %s", action.Code(), test.code)
		}

		code, err := ActionCodeForScript(script, true)
		if err != nil {
			t.Fatalf("Failed to get action code : %s", err)
		}

		if code != test.code {
			t.Errorf("Wrong action code : got %s, want %s", code, test.code)
		}
	}
}

func Test_SerializeFlag(t *testing.T) {
	var tests = []struct {
		flag []byte
	}{
		{
			flag: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
		},
	}

	for i, test := range tests {
		script, err := SerializeFlagOutputScript(test.flag)
		if err != nil {
			t.Fatalf("Failed to serialize flag : %s", err)
		}

		t.Logf("Script %d : %s", i, bitcoin.Script(script))

		flag, err := DeserializeFlagOutputScript(script)
		if err != nil {
			t.Fatalf("Failed to serialize flag : %s", err)
		}

		t.Logf("Flag %d : %x", i, flag)

		if !bytes.Equal(flag, test.flag) {
			t.Errorf("Wrong flag %d : got %x, want %x", i, flag, test.flag)
		}
	}
}

func TestInstrumentID(t *testing.T) {
	// Random instrument code
	var code bitcoin.Hash20
	rand.Read(code[:])

	instrumentType := "SHR"

	id := InstrumentID(instrumentType, code)
	t.Logf("Generated Instrument ID : %s", id)

	rInstrumentType, rInstrumentCode, err := DecodeInstrumentID(id)
	if err != nil {
		t.Fatalf("Failed to decode instrument id : %s", err)
	}

	if rInstrumentType != instrumentType {
		t.Fatalf("Incorrect instrument type : got %s, want %s", rInstrumentType, instrumentType)
	}

	if !bytes.Equal(rInstrumentCode.Bytes(), code[:]) {
		t.Fatalf("Incorrect instrument code : got %x, want %x", rInstrumentCode.Bytes(), code[:])
	}
}

func TestInstrumentID_BSV(t *testing.T) {
	var code bitcoin.Hash20

	instrumentType := BSVInstrumentID

	id := InstrumentID(instrumentType, code)

	if id != instrumentType {
		t.Fatalf("got %v, want %v", id, instrumentType)
	}
}

func Test_Deserialize_Empty(t *testing.T) {
	transfer := &actions.Transfer{}
	script, err := Serialize(transfer, true)
	if err != nil {
		t.Fatalf("Failed to serialize empty transfer : %s", err)
	}

	script[len(script)-6] = bitcoin.OP_3
	script = append(script, bitcoin.OP_FALSE)
	t.Logf("Script : %s", script)

	action, err := Deserialize(script, true)
	if err != nil {
		t.Fatalf("Failed to deserialize empty transfer : %s", err)
	}

	if _, ok := action.(*actions.Transfer); !ok {
		t.Errorf("Action not a transfer")
	}

	js, _ := json.MarshalIndent(action, "", "  ")
	t.Logf("Action : %s", js)
}
