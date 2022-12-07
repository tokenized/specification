package instruments

import (
	"bytes"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Instrument interface {
	proto.Message

	Code() string
	TypeName() string

	Validate() error
	Equal(proto.Message) bool

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

	// CodeCreditNote identifies a payload as a CreditNote instrument message.
	CodeCreditNote = "CRD"

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
	case CodeCreditNote:
		return &CreditNote{}
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

	if len(payload) == 0 {
		return result, nil // empty result
	}

	if len(payload) == 1 && payload[0] == bitcoin.OP_FALSE {
		return result, nil // empty result
	}

	if err := proto.Unmarshal(payload, result); err != nil {
		return nil, errors.Wrap(err, "protobuf unmarshal")
	}

	return result, nil
}

func (a *Membership) Code() string {
	return CodeMembership
}

func (a *Membership) TypeName() string {
	return "Membership"
}

func (a *Membership) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Membership) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize Membership")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Currency) Code() string {
	return CodeCurrency
}

func (a *Currency) TypeName() string {
	return "Currency"
}

func (a *Currency) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Currency) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize Currency")
	}

	_, err = buf.Write(data)
	return err
}

func (a *ShareCommon) Code() string {
	return CodeShareCommon
}

func (a *ShareCommon) TypeName() string {
	return "ShareCommon"
}

func (a *ShareCommon) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *ShareCommon) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize ShareCommon")
	}

	_, err = buf.Write(data)
	return err
}

func (a *BondFixedRate) Code() string {
	return CodeBondFixedRate
}

func (a *BondFixedRate) TypeName() string {
	return "BondFixedRate"
}

func (a *BondFixedRate) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *BondFixedRate) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize BondFixedRate")
	}

	_, err = buf.Write(data)
	return err
}

func (a *Coupon) Code() string {
	return CodeCoupon
}

func (a *Coupon) TypeName() string {
	return "Coupon"
}

func (a *Coupon) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *Coupon) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize Coupon")
	}

	_, err = buf.Write(data)
	return err
}

func (a *LoyaltyPoints) Code() string {
	return CodeLoyaltyPoints
}

func (a *LoyaltyPoints) TypeName() string {
	return "LoyaltyPoints"
}

func (a *LoyaltyPoints) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *LoyaltyPoints) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize LoyaltyPoints")
	}

	_, err = buf.Write(data)
	return err
}

func (a *TicketAdmission) Code() string {
	return CodeTicketAdmission
}

func (a *TicketAdmission) TypeName() string {
	return "TicketAdmission"
}

func (a *TicketAdmission) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *TicketAdmission) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize TicketAdmission")
	}

	_, err = buf.Write(data)
	return err
}

func (a *CasinoChip) Code() string {
	return CodeCasinoChip
}

func (a *CasinoChip) TypeName() string {
	return "CasinoChip"
}

func (a *CasinoChip) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *CasinoChip) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize CasinoChip")
	}

	_, err = buf.Write(data)
	return err
}

func (a *InformationServiceLicense) Code() string {
	return CodeInformationServiceLicense
}

func (a *InformationServiceLicense) TypeName() string {
	return "InformationServiceLicense"
}

func (a *InformationServiceLicense) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *InformationServiceLicense) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize InformationServiceLicense")
	}

	_, err = buf.Write(data)
	return err
}

func (a *CreditNote) Code() string {
	return CodeCreditNote
}

func (a *CreditNote) TypeName() string {
	return "CreditNote"
}

func (a *CreditNote) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an instrument to a byte slice.
func (a *CreditNote) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize CreditNote")
	}

	_, err = buf.Write(data)
	return err
}
