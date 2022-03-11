package protocol

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/tokenized/pkg/bitcoin"
)

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
