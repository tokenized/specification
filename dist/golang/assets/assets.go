package assets

import (
	"bytes"
	"fmt"
	"io"

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

	ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error)
}

const (
	// CodeMembership identifies a payload as a Membership asset message.
	CodeMembership = "MEM"

	// CodeCurrency identifies a payload as a Currency asset message.
	CodeCurrency = "CUR"

	// CodeShareCommon identifies a payload as a ShareCommon asset message.
	CodeShareCommon = "SHC"

	// CodeCoupon identifies a payload as a Coupon asset message.
	CodeCoupon = "COU"

	// CodeLoyaltyPoints identifies a payload as a LoyaltyPoints asset message.
	CodeLoyaltyPoints = "LOY"

	// CodeTicketAdmission identifies a payload as a TicketAdmission asset message.
	CodeTicketAdmission = "TIC"

	// CodeCasinoChip identifies a payload as a CasinoChip asset message.
	CodeCasinoChip = "CHP"
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
	case CodeCoupon:
		return &Coupon{}
	case CodeLoyaltyPoints:
		return &LoyaltyPoints{}
	case CodeTicketAdmission:
		return &TicketAdmission{}
	case CodeCasinoChip:
		return &CasinoChip{}
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

// Helper functions for amendments

// ReadBase128VarInt reads a base 128 variable encoded integer from the reader.
func ReadBase128VarInt(r io.ByteReader) (int, error) {
	value := uint32(0)
	done := false
	bitOffset := uint32(0)
	for !done {
		subValue, err := r.ReadByte()
		if err != nil {
			return int(value), err
		}

		done = (subValue & 0x80) == 0 // High bit not set
		subValue = subValue & 0x7f    // Remove high bit

		value += uint32(subValue) << bitOffset
		bitOffset += 7
	}

	return int(value), nil
}

// WriteBase128VarInt writes a base 128 variable encoded integer to the writer.
func WriteBase128VarInt(w io.ByteWriter, value int) error {
	v := uint32(value)
	for {
		if v < 128 {
			return w.WriteByte(byte(v))
		}
		subValue := (byte(v&0x7f) | 0x80) // Get last 7 bits and set high bit
		if err := w.WriteByte(subValue); err != nil {
			return err
		}
		v = v >> 7
	}
}
