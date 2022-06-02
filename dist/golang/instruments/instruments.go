package instruments

import (
	"bytes"
	"fmt"

	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/pkg/errors"
)

type Instrument interface {
	Code() string

	Validate() error
	Equal(interface{}) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error

	ApplyAmendment(fip permissions.FieldIndexPath, operation uint32, data []byte,
		permissions permissions.Permissions) (permissions.Permissions, error)
}

const (
	// CodeMembership identifies a payload as a Membership instrument message.
	CodeMembership = "MBR"

	// CodeCurrency identifies a payload as a Currency instrument message.
	CodeCurrency = "CCY"

	// CodeShareCommon identifies a payload as a ShareCommon instrument message.
	CodeShareCommon = "SHC"

	// CodeBondFixedRate identifies a payload as a BondFixedRate instrument message.
	CodeBondFixedRate = "BFR"

	// CodeCoupon identifies a payload as a Coupon instrument message.
	CodeCoupon = "COU"

	// CodeLoyaltyPoints identifies a payload as a LoyaltyPoints instrument message.
	CodeLoyaltyPoints = "LOY"

	// CodeTicketAdmission identifies a payload as a TicketAdmission instrument message.
	CodeTicketAdmission = "TIC"

	// CodeCasinoChip identifies a payload as a CasinoChip instrument message.
	CodeCasinoChip = "CHP"

	// CodeInformationServiceLicense identifies a payload as a InformationServiceLicense instrument message.
	CodeInformationServiceLicense = "ISL"

	// BondTypeCorporate specifies a corporate bond.
	BondTypeCorporate = "C"

	// BondTypeMunicipal specifies a municipal bond.
	BondTypeMunicipal = "M"

	// BondTypeGovernment specifies a government bond.
	BondTypeGovernment = "G"

	// CasinoChipUseTypeRealMoney specifies a casino chip is real money.
	CasinoChipUseTypeRealMoney = "R"

	// CasinoChipUseTypeSocial specifies a casino chip is social.
	CasinoChipUseTypeSocial = "S"

	// CasinoChipUseTypeFreePlay specifies a casino chip is free play.
	CasinoChipUseTypeFreePlay = "F"
)

// NewInstrumentFromCode returns a new object of the correct type for the code.
func NewInstrumentFromCode(code string) Instrument {
	switch code {
	case CodeMembership:
		return &Membership{}
	case CodeCurrency:
		return &Currency{}
	case CodeShareCommon:
		return &ShareCommon{}
	case CodeBondFixedRate:
		return &BondFixedRate{}
	case CodeCoupon:
		return &Coupon{}
	case CodeLoyaltyPoints:
		return &LoyaltyPoints{}
	case CodeTicketAdmission:
		return &TicketAdmission{}
	case CodeCasinoChip:
		return &CasinoChip{}
	case CodeInformationServiceLicense:
		return &InformationServiceLicense{}
	default:
		return nil
	}
}

// Deserialize reads an instrument from a byte slice.
func Deserialize(code []byte, payload []byte) (Instrument, error) {
	result := NewInstrumentFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown instrument code : %s", string(code))
	}

	if len(payload) != 0 {
		if _, err := bsor.UnmarshalBinary(payload, result); err != nil {
			return nil, errors.Wrap(err, "unmarshal bsor")
		}
	}

	return result, nil
}

func (a *Membership) Code() string {
	return CodeMembership
}

func (a *Membership) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Membership) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize Membership")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Currency) Code() string {
	return CodeCurrency
}

func (a *Currency) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Currency) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize Currency")
	}

	_, err = buf.Write(data)
	return err
}

func (a *ShareCommon) Code() string {
	return CodeShareCommon
}

func (a *ShareCommon) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *ShareCommon) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize ShareCommon")
	}

	_, err = buf.Write(data)
	return err
}

func (a *BondFixedRate) Code() string {
	return CodeBondFixedRate
}

func (a *BondFixedRate) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *BondFixedRate) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize BondFixedRate")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Coupon) Code() string {
	return CodeCoupon
}

func (a *Coupon) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Coupon) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize Coupon")
	}

	_, err = buf.Write(data)
	return err
}

func (a *LoyaltyPoints) Code() string {
	return CodeLoyaltyPoints
}

func (a *LoyaltyPoints) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *LoyaltyPoints) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize LoyaltyPoints")
	}

	_, err = buf.Write(data)
	return err
}

func (a *TicketAdmission) Code() string {
	return CodeTicketAdmission
}

func (a *TicketAdmission) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *TicketAdmission) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize TicketAdmission")
	}

	_, err = buf.Write(data)
	return err
}

func (a *CasinoChip) Code() string {
	return CodeCasinoChip
}

func (a *CasinoChip) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *CasinoChip) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize CasinoChip")
	}

	_, err = buf.Write(data)
	return err
}

func (a *InformationServiceLicense) Code() string {
	return CodeInformationServiceLicense
}

func (a *InformationServiceLicense) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *InformationServiceLicense) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize InformationServiceLicense")
	}

	_, err = buf.Write(data)
	return err
}
