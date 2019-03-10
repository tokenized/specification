package protocol

import "bytes"

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

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
	BaseMessage
	Version byte
	TradingRestriction polity
	RedeemingEntity nvarchar8
	IssueDate uint64
	ExpiryDate uint64
	Value uint64
	Currency []byte
	Description nvarchar16
}

// NewCoupon returns a new Coupon.
func NewCoupon() *Coupon {
	return &Coupon{}
}

// Type returns the type identifer for this message.
func (m Coupon) Type() string {
	return CodeCoupon
}

// Len returns the byte size of this message.
func (m Coupon) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m Coupon) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.RedeemingEntity); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.IssueDate); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ExpiryDate); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Value); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.Currency, 3)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Coupon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.RedeemingEntity); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.IssueDate); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ExpiryDate); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Value); err != nil {
		return 0, err
	}

	m.Currency = make([]byte, 3)
	if err := m.readLen(buf, m.Currency); err != nil {
		return 0, err
	}

	m.Currency = bytes.Trim(m.Currency, "\x00")

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Coupon) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Currency asset type.
type Currency struct {
	BaseMessage
	Version byte
	TradingRestriction polity
	ISOCode []byte
	MonetaryAuthority nvarchar8
	Description nvarchar8
}

// NewCurrency returns a new Currency.
func NewCurrency() *Currency {
	return &Currency{}
}

// Type returns the type identifer for this message.
func (m Currency) Type() string {
	return CodeCurrency
}

// Len returns the byte size of this message.
func (m Currency) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m Currency) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.ISOCode, 3)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.MonetaryAuthority); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Currency) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	m.ISOCode = make([]byte, 3)
	if err := m.readLen(buf, m.ISOCode); err != nil {
		return 0, err
	}

	m.ISOCode = bytes.Trim(m.ISOCode, "\x00")

	if err := m.read(buf, &m.MonetaryAuthority); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Currency) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// LoyaltyPoints asset type.
type LoyaltyPoints struct {
	BaseMessage
	Version byte
	TradingRestriction polity
	AgeRestriction []byte
	Type byte
	Name nvarchar8
	ValidFrom timestamp
	ExpirationTimestamp timestamp
	Description nvarchar16
}

// NewLoyaltyPoints returns a new LoyaltyPoints.
func NewLoyaltyPoints() *LoyaltyPoints {
	return &LoyaltyPoints{}
}

// Type returns the type identifer for this message.
func (m LoyaltyPoints) Type() string {
	return CodeLoyaltyPoints
}

// Len returns the byte size of this message.
func (m LoyaltyPoints) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m LoyaltyPoints) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Type); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Name); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *LoyaltyPoints) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := m.readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = bytes.Trim(m.AgeRestriction, "\x00")

	if err := m.read(buf, &m.Type); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Name); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m LoyaltyPoints) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Membership asset type.
type Membership struct {
	BaseMessage
	Version byte
	TradingRestriction polity
	AgeRestriction []byte
	ValidFrom timestamp
	ExpirationTimestamp timestamp
	ID nvarchar8
	MembershipType nvarchar8
	Description nvarchar64
}

// NewMembership returns a new Membership.
func NewMembership() *Membership {
	return &Membership{}
}

// Type returns the type identifer for this message.
func (m Membership) Type() string {
	return CodeMembership
}

// Len returns the byte size of this message.
func (m Membership) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m Membership) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ID); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.MembershipType); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Membership) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := m.readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = bytes.Trim(m.AgeRestriction, "\x00")

	if err := m.read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ID); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.MembershipType); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Membership) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// ShareCommon asset type.
type ShareCommon struct {
	BaseMessage
	Version byte
	TradingRestriction polity
	TransferLockout timestamp
	Ticker []byte
	ISIN []byte
	Description nvarchar16
}

// NewShareCommon returns a new ShareCommon.
func NewShareCommon() *ShareCommon {
	return &ShareCommon{}
}

// Type returns the type identifer for this message.
func (m ShareCommon) Type() string {
	return CodeShareCommon
}

// Len returns the byte size of this message.
func (m ShareCommon) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m ShareCommon) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TransferLockout); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.Ticker, 5)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.ISIN, 12)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *ShareCommon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TransferLockout); err != nil {
		return 0, err
	}

	m.Ticker = make([]byte, 5)
	if err := m.readLen(buf, m.Ticker); err != nil {
		return 0, err
	}

	m.Ticker = bytes.Trim(m.Ticker, "\x00")

	m.ISIN = make([]byte, 12)
	if err := m.readLen(buf, m.ISIN); err != nil {
		return 0, err
	}

	m.ISIN = bytes.Trim(m.ISIN, "\x00")

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ShareCommon) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// TicketAdmission asset type.
type TicketAdmission struct {
	BaseMessage
	Version byte
	TradingRestriction polity
	AgeRestriction []byte
	Type []byte
	Venue nvarchar8
	Class nvarchar8
	Area nvarchar8
	Seat nvarchar8
	StartTimeDate timestamp
	ValidFrom timestamp
	ExpirationTimestamp timestamp
	Description nvarchar16
}

// NewTicketAdmission returns a new TicketAdmission.
func NewTicketAdmission() *TicketAdmission {
	return &TicketAdmission{}
}

// Type returns the type identifer for this message.
func (m TicketAdmission) Type() string {
	return CodeTicketAdmission
}

// Len returns the byte size of this message.
func (m TicketAdmission) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m TicketAdmission) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := m.write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.TradingRestriction); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.pad(m.Type, 3)); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Venue); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Class); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Area); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Seat); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.StartTimeDate); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := m.write(buf, m.Description); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *TicketAdmission) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.read(buf, &m.Version); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := m.readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = bytes.Trim(m.AgeRestriction, "\x00")

	m.Type = make([]byte, 3)
	if err := m.readLen(buf, m.Type); err != nil {
		return 0, err
	}

	m.Type = bytes.Trim(m.Type, "\x00")

	if err := m.read(buf, &m.Venue); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Class); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Area); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Seat); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.StartTimeDate); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := m.read(buf, &m.Description); err != nil {
		return 0, err
	}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m TicketAdmission) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}
