package assets

import (
	"bytes"
	"fmt"

	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Asset interface {
	proto.Message

	Code() string

	Validate() error
	Equal(proto.Message) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error

	ApplyAmendment(fip permissions.FieldIndexPath, operation uint32, data []byte,
		permissions permissions.Permissions) (permissions.Permissions, error)
}

const (
	// CodeMembership identifies a payload as a Membership asset message.
	CodeMembership = "MBR"

	// CodeCurrency identifies a payload as a Currency asset message.
	CodeCurrency = "CCY"

	// CodeShareCommon identifies a payload as a ShareCommon asset message.
	CodeShareCommon = "SHC"

	// CodeBondFixedRate identifies a payload as a BondFixedRate asset message.
	CodeBondFixedRate = "BFR"

	// CodeCoupon identifies a payload as a Coupon asset message.
	CodeCoupon = "COU"

	// CodeLoyaltyPoints identifies a payload as a LoyaltyPoints asset message.
	CodeLoyaltyPoints = "LOY"

	// CodeTicketAdmission identifies a payload as a TicketAdmission asset message.
	CodeTicketAdmission = "TIC"

	// CodeCasinoChip identifies a payload as a CasinoChip asset message.
	CodeCasinoChip = "CHP"

	// CodeInformationServiceLicense identifies a payload as a InformationServiceLicense asset message.
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

// NewAssetFromCode returns a new object of the correct type for the code.
func NewAssetFromCode(code string) Asset {
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

// Deserialize reads an asset from a byte slice.
func Deserialize(code []byte, payload []byte) (Asset, error) {
	result := NewAssetFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown asset code : %s", string(code))
	}

	if len(payload) != 0 {
		if err := proto.Unmarshal(payload, result); err != nil {
			return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
		}
	}

	return result, nil
}

func (a *Membership) Code() string {
	return CodeMembership
}

func (a *Membership) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *Currency) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *ShareCommon) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *BondFixedRate) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *Coupon) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *LoyaltyPoints) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *TicketAdmission) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *CasinoChip) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
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

func (a *InformationServiceLicense) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
func (a *InformationServiceLicense) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize InformationServiceLicense")
	}

	_, err = buf.Write(data)
	return err
}
