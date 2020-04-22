package messages

import (
	"testing"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// PublicMessage identifies a payload as a PublicMessage message.
	messagePublicMessage, err := Deserialize(CodePublicMessage, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PublicMessage : %s", err)
	}
	_, ok = messagePublicMessage.(*PublicMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PublicMessage")
	}

	// PrivateMessage identifies a payload as a PrivateMessage message.
	messagePrivateMessage, err := Deserialize(CodePrivateMessage, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PrivateMessage : %s", err)
	}
	_, ok = messagePrivateMessage.(*PrivateMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PrivateMessage")
	}

	// RevertedTx identifies a payload as a RevertedTx message.
	messageRevertedTx, err := Deserialize(CodeRevertedTx, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for RevertedTx : %s", err)
	}
	_, ok = messageRevertedTx.(*RevertedTx)
	if !ok {
		t.Fatalf("Failed deserialize type check for RevertedTx")
	}

	// Offer identifies a payload as a Offer message.
	messageOffer, err := Deserialize(CodeOffer, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Offer : %s", err)
	}
	_, ok = messageOffer.(*Offer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Offer")
	}

	// SignatureRequest identifies a payload as a SignatureRequest message.
	messageSignatureRequest, err := Deserialize(CodeSignatureRequest, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SignatureRequest : %s", err)
	}
	_, ok = messageSignatureRequest.(*SignatureRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SignatureRequest")
	}

	// SettlementRequest identifies a payload as a SettlementRequest message.
	messageSettlementRequest, err := Deserialize(CodeSettlementRequest, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SettlementRequest : %s", err)
	}
	_, ok = messageSettlementRequest.(*SettlementRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SettlementRequest")
	}

	// OutputMetadata identifies a payload as a OutputMetadata message.
	messageOutputMetadata, err := Deserialize(CodeOutputMetadata, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for OutputMetadata : %s", err)
	}
	_, ok = messageOutputMetadata.(*OutputMetadata)
	if !ok {
		t.Fatalf("Failed deserialize type check for OutputMetadata")
	}

	// InitiateRelationship identifies a payload as a InitiateRelationship message.
	messageInitiateRelationship, err := Deserialize(CodeInitiateRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InitiateRelationship : %s", err)
	}
	_, ok = messageInitiateRelationship.(*InitiateRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for InitiateRelationship")
	}

	// PendingAcceptRelationship identifies a payload as a PendingAcceptRelationship message.
	messagePendingAcceptRelationship, err := Deserialize(CodePendingAcceptRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PendingAcceptRelationship : %s", err)
	}
	_, ok = messagePendingAcceptRelationship.(*PendingAcceptRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for PendingAcceptRelationship")
	}

	// AcceptRelationship identifies a payload as a AcceptRelationship message.
	messageAcceptRelationship, err := Deserialize(CodeAcceptRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for AcceptRelationship : %s", err)
	}
	_, ok = messageAcceptRelationship.(*AcceptRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for AcceptRelationship")
	}

	// RelationshipAmendment identifies a payload as a RelationshipAmendment message.
	messageRelationshipAmendment, err := Deserialize(CodeRelationshipAmendment, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for RelationshipAmendment : %s", err)
	}
	_, ok = messageRelationshipAmendment.(*RelationshipAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for RelationshipAmendment")
	}

	// InitiateThread identifies a payload as a InitiateThread message.
	messageInitiateThread, err := Deserialize(CodeInitiateThread, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InitiateThread : %s", err)
	}
	_, ok = messageInitiateThread.(*InitiateThread)
	if !ok {
		t.Fatalf("Failed deserialize type check for InitiateThread")
	}

}
