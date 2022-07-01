package messages

import (
	"testing"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// PublicMessage identifies a payload as a PublicMessage message.
	messagePublicMessage, err := DeserializeV1(CodePublicMessage, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PublicMessage : %s", err)
	}
	_, ok = messagePublicMessage.(*PublicMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PublicMessage")
	}

	// PrivateMessage identifies a payload as a PrivateMessage message.
	messagePrivateMessage, err := DeserializeV1(CodePrivateMessage, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PrivateMessage : %s", err)
	}
	_, ok = messagePrivateMessage.(*PrivateMessage)
	if !ok {
		t.Fatalf("Failed deserialize type check for PrivateMessage")
	}

	// RevertedTx identifies a payload as a RevertedTx message.
	messageRevertedTx, err := DeserializeV1(CodeRevertedTx, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for RevertedTx : %s", err)
	}
	_, ok = messageRevertedTx.(*RevertedTx)
	if !ok {
		t.Fatalf("Failed deserialize type check for RevertedTx")
	}

	// Offer identifies a payload as a Offer message.
	messageOffer, err := DeserializeV1(CodeOffer, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Offer : %s", err)
	}
	_, ok = messageOffer.(*Offer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Offer")
	}

	// SignatureRequest identifies a payload as a SignatureRequest message.
	messageSignatureRequest, err := DeserializeV1(CodeSignatureRequest, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SignatureRequest : %s", err)
	}
	_, ok = messageSignatureRequest.(*SignatureRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SignatureRequest")
	}

	// SettlementRequest identifies a payload as a SettlementRequest message.
	messageSettlementRequest, err := DeserializeV1(CodeSettlementRequest, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for SettlementRequest : %s", err)
	}
	_, ok = messageSettlementRequest.(*SettlementRequest)
	if !ok {
		t.Fatalf("Failed deserialize type check for SettlementRequest")
	}

	// OutputMetadata identifies a payload as a OutputMetadata message.
	messageOutputMetadata, err := DeserializeV1(CodeOutputMetadata, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for OutputMetadata : %s", err)
	}
	_, ok = messageOutputMetadata.(*OutputMetadata)
	if !ok {
		t.Fatalf("Failed deserialize type check for OutputMetadata")
	}

	// Distribution identifies a payload as a Distribution message.
	messageDistribution, err := DeserializeV1(CodeDistribution, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Distribution : %s", err)
	}
	_, ok = messageDistribution.(*Distribution)
	if !ok {
		t.Fatalf("Failed deserialize type check for Distribution")
	}

	// InitiateRelationship identifies a payload as a InitiateRelationship message.
	messageInitiateRelationship, err := DeserializeV1(CodeInitiateRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InitiateRelationship : %s", err)
	}
	_, ok = messageInitiateRelationship.(*InitiateRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for InitiateRelationship")
	}

	// PendingAcceptRelationship identifies a payload as a PendingAcceptRelationship message.
	messagePendingAcceptRelationship, err := DeserializeV1(CodePendingAcceptRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for PendingAcceptRelationship : %s", err)
	}
	_, ok = messagePendingAcceptRelationship.(*PendingAcceptRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for PendingAcceptRelationship")
	}

	// AcceptRelationship identifies a payload as a AcceptRelationship message.
	messageAcceptRelationship, err := DeserializeV1(CodeAcceptRelationship, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for AcceptRelationship : %s", err)
	}
	_, ok = messageAcceptRelationship.(*AcceptRelationship)
	if !ok {
		t.Fatalf("Failed deserialize type check for AcceptRelationship")
	}

	// RelationshipAmendment identifies a payload as a RelationshipAmendment message.
	messageRelationshipAmendment, err := DeserializeV1(CodeRelationshipAmendment, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for RelationshipAmendment : %s", err)
	}
	_, ok = messageRelationshipAmendment.(*RelationshipAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for RelationshipAmendment")
	}

	// InitiateThread identifies a payload as a InitiateThread message.
	messageInitiateThread, err := DeserializeV1(CodeInitiateThread, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InitiateThread : %s", err)
	}
	_, ok = messageInitiateThread.(*InitiateThread)
	if !ok {
		t.Fatalf("Failed deserialize type check for InitiateThread")
	}

}
