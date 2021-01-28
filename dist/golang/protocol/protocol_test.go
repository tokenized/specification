package protocol

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/tokenized/pkg/bitcoin"
)

func TestAssetID(t *testing.T) {
	// Random asset code
	var code bitcoin.Hash20
	rand.Read(code[:])

	assetType := "SHR"

	id := AssetID(assetType, code)
	t.Logf("Generated Asset ID : %s", id)

	rAssetType, rAssetCode, err := DecodeAssetID(id)
	if err != nil {
		t.Fatalf("Failed to decode asset id : %s", err)
	}

	if rAssetType != assetType {
		t.Fatalf("Incorrect asset type : got %s, want %s", rAssetType, assetType)
	}

	if !bytes.Equal(rAssetCode.Bytes(), code[:]) {
		t.Fatalf("Incorrect asset code : got %x, want %x", rAssetCode.Bytes(), code[:])
	}
}

func TestAssetID_BSV(t *testing.T) {
	var code bitcoin.Hash20

	assetType := "BSV"

	id := AssetID(assetType, code)

	if id != assetType {
		t.Fatalf("got %v, want %v", id, assetType)
	}
}
