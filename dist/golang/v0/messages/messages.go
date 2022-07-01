package messages

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Message interface {
	proto.Message

	Code() uint32

	Validate() error
	Equal(proto.Message) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error
}

const (
	// CodePublicMessage identifies a payload as a PublicMessage message.
	CodePublicMessage = uint32(0002)

	// CodePrivateMessage identifies a payload as a PrivateMessage message.
	CodePrivateMessage = uint32(0003)

	// CodeRevertedTx identifies a payload as a RevertedTx message.
	CodeRevertedTx = uint32(0004)

	// CodeOffer identifies a payload as a Offer message.
	CodeOffer = uint32(1001)

	// CodeSignatureRequest identifies a payload as a SignatureRequest message.
	CodeSignatureRequest = uint32(1002)

	// CodeSettlementRequest identifies a payload as a SettlementRequest message.
	CodeSettlementRequest = uint32(1003)

	// CodeOutputMetadata identifies a payload as a OutputMetadata message.
	CodeOutputMetadata = uint32(1004)

	// CodeDistribution identifies a payload as a Distribution message.
	CodeDistribution = uint32(1005)

	// CodeInitiateRelationship identifies a payload as a InitiateRelationship message.
	CodeInitiateRelationship = uint32(2001)

	// CodePendingAcceptRelationship identifies a payload as a PendingAcceptRelationship message.
	CodePendingAcceptRelationship = uint32(2002)

	// CodeAcceptRelationship identifies a payload as a AcceptRelationship message.
	CodeAcceptRelationship = uint32(2003)

	// CodeRelationshipAmendment identifies a payload as a RelationshipAmendment message.
	CodeRelationshipAmendment = uint32(2004)

	// CodeInitiateThread identifies a payload as a InitiateThread message.
	CodeInitiateThread = uint32(2005)
)

// NewMessageFromCode returns a new object of the correct type for the code.
func NewMessageFromCode(code uint32) Message {
	switch code {
	case CodePublicMessage:
		return &PublicMessage{}
	case CodePrivateMessage:
		return &PrivateMessage{}
	case CodeRevertedTx:
		return &RevertedTx{}
	case CodeOffer:
		return &Offer{}
	case CodeSignatureRequest:
		return &SignatureRequest{}
	case CodeSettlementRequest:
		return &SettlementRequest{}
	case CodeOutputMetadata:
		return &OutputMetadata{}
	case CodeDistribution:
		return &Distribution{}
	case CodeInitiateRelationship:
		return &InitiateRelationship{}
	case CodePendingAcceptRelationship:
		return &PendingAcceptRelationship{}
	case CodeAcceptRelationship:
		return &AcceptRelationship{}
	case CodeRelationshipAmendment:
		return &RelationshipAmendment{}
	case CodeInitiateThread:
		return &InitiateThread{}
	default:
		return nil
	}
}

// DeserializeV0 reads a message from a byte slice.
func DeserializeV0(code uint32, payload []byte) (Message, error) {
	result := NewMessageFromCode(code)
	if result == nil {
		return nil, fmt.Errorf("Unknown message code : %d", code)
	}

	if len(payload) != 0 {
		if err := proto.Unmarshal(payload, result); err != nil {
			return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
		}
	}

	return result, nil
}

func (a *PublicMessage) Code() uint32 {
	return CodePublicMessage
}

func (a *PublicMessage) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *PublicMessage) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize PublicMessage")
	}

	_, err = buf.Write(data)
	return err
}

func (a *PrivateMessage) Code() uint32 {
	return CodePrivateMessage
}

func (a *PrivateMessage) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *PrivateMessage) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize PrivateMessage")
	}

	_, err = buf.Write(data)
	return err
}

func (a *RevertedTx) Code() uint32 {
	return CodeRevertedTx
}

func (a *RevertedTx) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *RevertedTx) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize RevertedTx")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Offer) Code() uint32 {
	return CodeOffer
}

func (a *Offer) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Offer) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize Offer")
	}

	_, err = buf.Write(data)
	return err
}

func (a *SignatureRequest) Code() uint32 {
	return CodeSignatureRequest
}

func (a *SignatureRequest) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *SignatureRequest) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize SignatureRequest")
	}

	_, err = buf.Write(data)
	return err
}

func (a *SettlementRequest) Code() uint32 {
	return CodeSettlementRequest
}

func (a *SettlementRequest) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *SettlementRequest) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize SettlementRequest")
	}

	_, err = buf.Write(data)
	return err
}

func (a *OutputMetadata) Code() uint32 {
	return CodeOutputMetadata
}

func (a *OutputMetadata) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *OutputMetadata) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize OutputMetadata")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Distribution) Code() uint32 {
	return CodeDistribution
}

func (a *Distribution) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Distribution) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize Distribution")
	}

	_, err = buf.Write(data)
	return err
}

func (a *InitiateRelationship) Code() uint32 {
	return CodeInitiateRelationship
}

func (a *InitiateRelationship) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *InitiateRelationship) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize InitiateRelationship")
	}

	_, err = buf.Write(data)
	return err
}

func (a *PendingAcceptRelationship) Code() uint32 {
	return CodePendingAcceptRelationship
}

func (a *PendingAcceptRelationship) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *PendingAcceptRelationship) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize PendingAcceptRelationship")
	}

	_, err = buf.Write(data)
	return err
}

func (a *AcceptRelationship) Code() uint32 {
	return CodeAcceptRelationship
}

func (a *AcceptRelationship) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *AcceptRelationship) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize AcceptRelationship")
	}

	_, err = buf.Write(data)
	return err
}

func (a *RelationshipAmendment) Code() uint32 {
	return CodeRelationshipAmendment
}

func (a *RelationshipAmendment) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *RelationshipAmendment) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize RelationshipAmendment")
	}

	_, err = buf.Write(data)
	return err
}

func (a *InitiateThread) Code() uint32 {
	return CodeInitiateThread
}

func (a *InitiateThread) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *InitiateThread) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize InitiateThread")
	}

	_, err = buf.Write(data)
	return err
}
