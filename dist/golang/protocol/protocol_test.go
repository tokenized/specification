package protocol

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestAssetID(t *testing.T) {
	// Random asset code
	b := make([]byte, 32)
	rand.Read(b)
	code := *AssetCodeFromBytes(b)

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

	if !bytes.Equal(rAssetCode.Bytes(), b) {
		t.Fatalf("Incorrect asset code : got %x, want %x", rAssetCode.Bytes(), b)
	}
}

func TestAssetID_BSV(t *testing.T) {
	b := make([]byte, 32)
	code := *AssetCodeFromBytes(b)

	assetType := "BSV"

	id := AssetID(assetType, code)

	if id != assetType {
		t.Fatalf("got %v, want %v", id, assetType)
	}
}
