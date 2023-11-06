package actions

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	binaryVersion = uint8(0)
)

func (a *ContractOffer) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ContractOffer) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ContractFormation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ContractFormation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ContractAmendment) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ContractAmendment) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *StaticContractFormation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *StaticContractFormation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ContractAddressChange) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ContractAddressChange) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *BodyOfAgreementOffer) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BodyOfAgreementOffer) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *BodyOfAgreementFormation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BodyOfAgreementFormation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *BodyOfAgreementAmendment) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BodyOfAgreementAmendment) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentDefinition) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentDefinition) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentCreation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentCreation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentModification) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentModification) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Transfer) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Transfer) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Settlement) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Settlement) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *RectificationSettlement) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *RectificationSettlement) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Proposal) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Proposal) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Vote) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Vote) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *BallotCast) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BallotCast) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *BallotCounted) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BallotCounted) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Result) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Result) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Order) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Order) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Freeze) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Freeze) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Thaw) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Thaw) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Confiscation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Confiscation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *DeprecatedReconciliation) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *DeprecatedReconciliation) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Establishment) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Establishment) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Addition) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Addition) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Alteration) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Alteration) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Removal) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Removal) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Message) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Message) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *Rejection) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Rejection) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *AdministratorField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *AdministratorField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *AdminIdentityCertificateField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *AdminIdentityCertificateField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *AmendmentField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *AmendmentField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentReceiverField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentReceiverField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentSettlementField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentSettlementField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *InstrumentTransferField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InstrumentTransferField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ChapterField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ChapterField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ClauseField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ClauseField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *DefinedTermField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *DefinedTermField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *DocumentField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *DocumentField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *EntityField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *EntityField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *FeeField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *FeeField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ManagerField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ManagerField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *OracleField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *OracleField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *QuantityIndexField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *QuantityIndexField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ReferenceTransactionField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ReferenceTransactionField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *ServiceField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ServiceField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *TargetAddressField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *TargetAddressField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}

func (a *VotingSystemField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *VotingSystemField) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
