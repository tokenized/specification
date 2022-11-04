package instruments

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	binaryVersion = uint8(0)
)

func (a *Membership) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Membership) UnmarshalBinary(data []byte) error {
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

func (a *Currency) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Currency) UnmarshalBinary(data []byte) error {
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

func (a *ShareCommon) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *ShareCommon) UnmarshalBinary(data []byte) error {
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

func (a *BondFixedRate) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *BondFixedRate) UnmarshalBinary(data []byte) error {
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

func (a *Coupon) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *Coupon) UnmarshalBinary(data []byte) error {
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

func (a *LoyaltyPoints) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *LoyaltyPoints) UnmarshalBinary(data []byte) error {
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

func (a *TicketAdmission) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *TicketAdmission) UnmarshalBinary(data []byte) error {
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

func (a *CasinoChip) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *CasinoChip) UnmarshalBinary(data []byte) error {
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

func (a *InformationServiceLicense) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *InformationServiceLicense) UnmarshalBinary(data []byte) error {
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

func (a *AgeRestrictionField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *AgeRestrictionField) UnmarshalBinary(data []byte) error {
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

func (a *CurrencyValueField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *CurrencyValueField) UnmarshalBinary(data []byte) error {
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

func (a *RateField) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *RateField) UnmarshalBinary(data []byte) error {
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
