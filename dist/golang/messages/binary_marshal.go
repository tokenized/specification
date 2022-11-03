package messages

import (
	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (a *PublicMessage) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *PublicMessage) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *PrivateMessage) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *PrivateMessage) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *RevertedTx) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *RevertedTx) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Offer) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Offer) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *SignatureRequest) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *SignatureRequest) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *SettlementRequest) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *SettlementRequest) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *OutputMetadata) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *OutputMetadata) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Distribution) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Distribution) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InitiateRelationship) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InitiateRelationship) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *PendingAcceptRelationship) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *PendingAcceptRelationship) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *AcceptRelationship) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *AcceptRelationship) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *RelationshipAmendment) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *RelationshipAmendment) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InitiateThread) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InitiateThread) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *AdministratorField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *AdministratorField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ChannelPartyField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ChannelPartyField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *DocumentField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *DocumentField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *EntityField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *EntityField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *IdentityOracleProofField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *IdentityOracleProofField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ManagerField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ManagerField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *OracleSignatureField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *OracleSignatureField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *OutpointField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *OutpointField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *OutputTagField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *OutputTagField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *PaymailProofField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *PaymailProofField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *PeriodField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *PeriodField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *TargetAddressField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *TargetAddressField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
