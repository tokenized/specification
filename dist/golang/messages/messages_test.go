package messages

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func codeBytes(code int) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, uint16(code))
	return buf.Bytes()
}

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// PublicMessage identifies a payload as a PublicMessage message.
	messagePublicMessage, err := Deserialize(codeBytes(CodePublicMessage), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PublicMessage : %s", err)
	}
	_, ok = messagePublicMessage.(*PublicMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PublicMessage")
	}

	// PrivateMessage identifies a payload as a PrivateMessage message.
	messagePrivateMessage, err := Deserialize(codeBytes(CodePrivateMessage), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PrivateMessage : %s", err)
	}
	_, ok = messagePrivateMessage.(*PrivateMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PrivateMessage")
	}

	// RevertedTx identifies a payload as a RevertedTx message.
	messageRevertedTx, err := Deserialize(codeBytes(CodeRevertedTx), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for RevertedTx : %s", err)
	}
	_, ok = messageRevertedTx.(*RevertedTx)
	if !ok {
		t.Fatalf("Failed deserialize type check for RevertedTx")
	}

	// Offer identifies a payload as a Offer message.
	messageOffer, err := Deserialize(codeBytes(CodeOffer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Offer : %s", err)
	}
	_, ok = messageOffer.(*Offer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Offer")
	}

	// SignatureRequest identifies a payload as a SignatureRequest message.
	messageSignatureRequest, err := Deserialize(codeBytes(CodeSignatureRequest), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SignatureRequest : %s", err)
	}
	_, ok = messageSignatureRequest.(*SignatureRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SignatureRequest")
	}

	// SettlementRequest identifies a payload as a SettlementRequest message.
	messageSettlementRequest, err := Deserialize(codeBytes(CodeSettlementRequest), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SettlementRequest : %s", err)
	}
	_, ok = messageSettlementRequest.(*SettlementRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SettlementRequest")
	}

	// OutputMetadata identifies a payload as a OutputMetadata message.
	messageOutputMetadata, err := Deserialize(codeBytes(CodeOutputMetadata), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for OutputMetadata : %s", err)
	}
	_, ok = messageOutputMetadata.(*OutputMetadata)
	if !ok {
		t.Fatalf("Failed deserialize type check for OutputMetadata")
	}

}
