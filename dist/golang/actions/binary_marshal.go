package actions

import (
	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (a *ContractOffer) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ContractOffer) UnmarshalBinary(data []byte) error {
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

func (a *ContractFormation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ContractFormation) UnmarshalBinary(data []byte) error {
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

func (a *ContractAmendment) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ContractAmendment) UnmarshalBinary(data []byte) error {
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

func (a *StaticContractFormation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *StaticContractFormation) UnmarshalBinary(data []byte) error {
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

func (a *ContractAddressChange) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ContractAddressChange) UnmarshalBinary(data []byte) error {
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

func (a *BodyOfAgreementOffer) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BodyOfAgreementOffer) UnmarshalBinary(data []byte) error {
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

func (a *BodyOfAgreementFormation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BodyOfAgreementFormation) UnmarshalBinary(data []byte) error {
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

func (a *BodyOfAgreementAmendment) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BodyOfAgreementAmendment) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentDefinition) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentDefinition) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentCreation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentCreation) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentModification) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentModification) UnmarshalBinary(data []byte) error {
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

func (a *Transfer) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Transfer) UnmarshalBinary(data []byte) error {
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

func (a *Settlement) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Settlement) UnmarshalBinary(data []byte) error {
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

func (a *Proposal) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Proposal) UnmarshalBinary(data []byte) error {
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

func (a *Vote) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Vote) UnmarshalBinary(data []byte) error {
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

func (a *BallotCast) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BallotCast) UnmarshalBinary(data []byte) error {
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

func (a *BallotCounted) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BallotCounted) UnmarshalBinary(data []byte) error {
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

func (a *Result) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Result) UnmarshalBinary(data []byte) error {
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

func (a *Order) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Order) UnmarshalBinary(data []byte) error {
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

func (a *Freeze) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Freeze) UnmarshalBinary(data []byte) error {
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

func (a *Thaw) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Thaw) UnmarshalBinary(data []byte) error {
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

func (a *Confiscation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Confiscation) UnmarshalBinary(data []byte) error {
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

func (a *Reconciliation) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Reconciliation) UnmarshalBinary(data []byte) error {
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

func (a *Establishment) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Establishment) UnmarshalBinary(data []byte) error {
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

func (a *Addition) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Addition) UnmarshalBinary(data []byte) error {
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

func (a *Alteration) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Alteration) UnmarshalBinary(data []byte) error {
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

func (a *Removal) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Removal) UnmarshalBinary(data []byte) error {
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

func (a *Message) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Message) UnmarshalBinary(data []byte) error {
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

func (a *Rejection) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Rejection) UnmarshalBinary(data []byte) error {
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

func (a *AdminIdentityCertificateField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *AdminIdentityCertificateField) UnmarshalBinary(data []byte) error {
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

func (a *AmendmentField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *AmendmentField) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentReceiverField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentReceiverField) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentSettlementField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentSettlementField) UnmarshalBinary(data []byte) error {
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

func (a *InstrumentTransferField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InstrumentTransferField) UnmarshalBinary(data []byte) error {
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

func (a *ChapterField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ChapterField) UnmarshalBinary(data []byte) error {
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

func (a *ClauseField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ClauseField) UnmarshalBinary(data []byte) error {
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

func (a *DefinedTermField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *DefinedTermField) UnmarshalBinary(data []byte) error {
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

func (a *OracleField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *OracleField) UnmarshalBinary(data []byte) error {
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

func (a *QuantityIndexField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *QuantityIndexField) UnmarshalBinary(data []byte) error {
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

func (a *ReferenceTransactionField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ReferenceTransactionField) UnmarshalBinary(data []byte) error {
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

func (a *ServiceField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ServiceField) UnmarshalBinary(data []byte) error {
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

func (a *VotingSystemField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *VotingSystemField) UnmarshalBinary(data []byte) error {
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
