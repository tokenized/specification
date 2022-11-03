package instruments

import (
	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func (a *Membership) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Membership) UnmarshalBinary(data []byte) error {
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

func (a *Currency) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Currency) UnmarshalBinary(data []byte) error {
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

func (a *ShareCommon) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *ShareCommon) UnmarshalBinary(data []byte) error {
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

func (a *BondFixedRate) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *BondFixedRate) UnmarshalBinary(data []byte) error {
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

func (a *Coupon) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *Coupon) UnmarshalBinary(data []byte) error {
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

func (a *LoyaltyPoints) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *LoyaltyPoints) UnmarshalBinary(data []byte) error {
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

func (a *TicketAdmission) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *TicketAdmission) UnmarshalBinary(data []byte) error {
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

func (a *CasinoChip) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *CasinoChip) UnmarshalBinary(data []byte) error {
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

func (a *InformationServiceLicense) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *InformationServiceLicense) UnmarshalBinary(data []byte) error {
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

func (a *AgeRestrictionField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *AgeRestrictionField) UnmarshalBinary(data []byte) error {
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

func (a *CurrencyValueField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *CurrencyValueField) UnmarshalBinary(data []byte) error {
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

func (a *RateField) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *RateField) UnmarshalBinary(data []byte) error {
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
