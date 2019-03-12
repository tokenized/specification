package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152

	// CodeCoupon identifies data as a Coupon message.
	CodeCoupon = "COU"

	// CodeCurrency identifies data as a Currency message.
	CodeCurrency = "CUR"

	// CodeLoyaltyPoints identifies data as a LoyaltyPoints message.
	CodeLoyaltyPoints = "LOY"

	// CodeMembership identifies data as a Membership message.
	CodeMembership = "MEM"

	// CodeShareCommon identifies data as a ShareCommon message.
	CodeShareCommon = "SHC"

	// CodeTicketAdmission identifies data as a TicketAdmission message.
	CodeTicketAdmission = "TIC"
)

// Coupon asset type.
type Coupon struct {
	Version            uint8
	TradingRestriction []byte
	RedeemingEntity    Nvarchar8
	IssueDate          uint64
	ExpiryDate         uint64
	Value              uint64
	Currency           []byte
	Description        Nvarchar16
}

// NewCoupon returns a new Coupon.
func NewCoupon() Coupon {
	return Coupon{}
}

// Type returns the type identifer for this message.
func (m Coupon) Type() string {
	return CodeCoupon
}

// Len returns the byte size of this message.
func (m Coupon) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Coupon) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Coupon) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	{
		b, err := m.RedeemingEntity.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssueDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpiryDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.Value); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Currency, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeCoupon, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Coupon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := m.RedeemingEntity.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.IssueDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpiryDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Value); err != nil {
		return 0, err
	}

	m.Currency = make([]byte, 3)
	if err := readLen(buf, m.Currency); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Coupon) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("RedeemingEntity:%#+v", m.RedeemingEntity))
	vals = append(vals, fmt.Sprintf("IssueDate:%v", m.IssueDate))
	vals = append(vals, fmt.Sprintf("ExpiryDate:%v", m.ExpiryDate))
	vals = append(vals, fmt.Sprintf("Value:%v", m.Value))
	vals = append(vals, fmt.Sprintf("Currency:%#x", m.Currency))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Currency asset type.
type Currency struct {
	Version            uint8
	TradingRestriction []byte
	ISOCode            []byte
	MonetaryAuthority  Nvarchar8
	Description        Nvarchar8
}

// NewCurrency returns a new Currency.
func NewCurrency() Currency {
	return Currency{}
}

// Type returns the type identifer for this message.
func (m Currency) Type() string {
	return CodeCurrency
}

// Len returns the byte size of this message.
func (m Currency) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Currency) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Currency) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ISOCode, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.MonetaryAuthority.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeCurrency, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Currency) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.ISOCode = make([]byte, 3)
	if err := readLen(buf, m.ISOCode); err != nil {
		return 0, err
	}

	if err := m.MonetaryAuthority.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Currency) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("ISOCode:%#x", m.ISOCode))
	vals = append(vals, fmt.Sprintf("MonetaryAuthority:%#+v", m.MonetaryAuthority))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// LoyaltyPoints asset type.
type LoyaltyPoints struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	OfferType           byte
	OfferName           Nvarchar8
	ValidFrom           uint64
	ExpirationTimestamp uint64
	Description         Nvarchar16
}

// NewLoyaltyPoints returns a new LoyaltyPoints.
func NewLoyaltyPoints() LoyaltyPoints {
	return LoyaltyPoints{}
}

// Type returns the type identifer for this message.
func (m LoyaltyPoints) Type() string {
	return CodeLoyaltyPoints
}

// Len returns the byte size of this message.
func (m LoyaltyPoints) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m LoyaltyPoints) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m LoyaltyPoints) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.OfferType); err != nil {
		return nil, err
	}

	{
		b, err := m.OfferName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeLoyaltyPoints, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m LoyaltyPoints) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.OfferType); err != nil {
		return 0, err
	}

	if err := m.OfferName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m LoyaltyPoints) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("OfferType:%#+v", m.OfferType))
	vals = append(vals, fmt.Sprintf("OfferName:%#+v", m.OfferName))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Membership asset type.
type Membership struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	ValidFrom           uint64
	ExpirationTimestamp uint64
	ID                  Nvarchar8
	MembershipType      Nvarchar8
	Description         Nvarchar64
}

// NewMembership returns a new Membership.
func NewMembership() Membership {
	return Membership{}
}

// Type returns the type identifer for this message.
func (m Membership) Type() string {
	return CodeMembership
}

// Len returns the byte size of this message.
func (m Membership) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Membership) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Membership) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	{
		b, err := m.ID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.MembershipType.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeMembership, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Membership) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.ID.Write(buf); err != nil {
		return 0, err
	}

	if err := m.MembershipType.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Membership) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("ID:%#+v", m.ID))
	vals = append(vals, fmt.Sprintf("MembershipType:%#+v", m.MembershipType))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ShareCommon asset type.
type ShareCommon struct {
	Version            uint8
	TradingRestriction []byte
	TransferLockout    uint64
	Ticker             []byte
	ISIN               []byte
	Description        Nvarchar16
}

// NewShareCommon returns a new ShareCommon.
func NewShareCommon() ShareCommon {
	return ShareCommon{}
}

// Type returns the type identifer for this message.
func (m ShareCommon) Type() string {
	return CodeShareCommon
}

// Len returns the byte size of this message.
func (m ShareCommon) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ShareCommon) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ShareCommon) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransferLockout); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Ticker, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ISIN, 12)); err != nil {
		return nil, err
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeShareCommon, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m ShareCommon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransferLockout); err != nil {
		return 0, err
	}

	m.Ticker = make([]byte, 5)
	if err := readLen(buf, m.Ticker); err != nil {
		return 0, err
	}

	m.ISIN = make([]byte, 12)
	if err := readLen(buf, m.ISIN); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m ShareCommon) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("TransferLockout:%#+v", m.TransferLockout))
	vals = append(vals, fmt.Sprintf("Ticker:%#x", m.Ticker))
	vals = append(vals, fmt.Sprintf("ISIN:%#x", m.ISIN))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// TicketAdmission asset type.
type TicketAdmission struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	AdmissionType       []byte
	Venue               Nvarchar8
	Class               Nvarchar8
	Area                Nvarchar8
	Seat                Nvarchar8
	StartTimeDate       uint64
	ValidFrom           uint64
	ExpirationTimestamp uint64
	Description         Nvarchar16
}

// NewTicketAdmission returns a new TicketAdmission.
func NewTicketAdmission() TicketAdmission {
	return TicketAdmission{}
}

// Type returns the type identifer for this message.
func (m TicketAdmission) Type() string {
	return CodeTicketAdmission
}

// Len returns the byte size of this message.
func (m TicketAdmission) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m TicketAdmission) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m TicketAdmission) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AdmissionType, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.Venue.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Class.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Area.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Seat.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.StartTimeDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	{
		b, err := m.Description.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeTicketAdmission, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m TicketAdmission) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	m.AdmissionType = make([]byte, 3)
	if err := readLen(buf, m.AdmissionType); err != nil {
		return 0, err
	}

	if err := m.Venue.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Class.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Area.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Seat.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.StartTimeDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.Description.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m TicketAdmission) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("AdmissionType:%#x", m.AdmissionType))
	vals = append(vals, fmt.Sprintf("Venue:%#+v", m.Venue))
	vals = append(vals, fmt.Sprintf("Class:%#+v", m.Class))
	vals = append(vals, fmt.Sprintf("Area:%#+v", m.Area))
	vals = append(vals, fmt.Sprintf("Seat:%#+v", m.Seat))
	vals = append(vals, fmt.Sprintf("StartTimeDate:%#+v", m.StartTimeDate))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
