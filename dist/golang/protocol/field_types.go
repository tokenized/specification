package protocol

import "bytes"

// Address Address represents a public address
type Address struct {
	Address []byte
}

// NewAddress return a new Address
func NewAddress() Address {
	return Address{}
}

// Serialize returns the byte representation of the message.
func (m Address) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Address); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Address) Write(buf *bytes.Buffer) error {

	m.Address = make([]byte, 34)
	if err := readLen(buf, m.Address); err != nil {
		return err
	}

	return nil
}

// Amendment An Amendment is used to describe the modification of a single
// field in a Contract or Asset, as defined in the ContractFormation and
// AssetCreation messages.
type Amendment struct {
	FieldIndex    uint8
	Element       uint16
	SubfieldIndex uint8
	Operation     byte
	Data          []byte
}

// NewAmendment return a new Amendment
func NewAmendment() Amendment {
	return Amendment{}
}

// Serialize returns the byte representation of the message.
func (m Amendment) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.FieldIndex); err != nil {
		return nil, err
	}

	if err := write(buf, m.Element); err != nil {
		return nil, err
	}

	if err := write(buf, m.SubfieldIndex); err != nil {
		return nil, err
	}

	if err := write(buf, m.Operation); err != nil {
		return nil, err
	}

	if err := write(buf, m.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Amendment) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.FieldIndex); err != nil {
		return err
	}

	if err := read(buf, &m.Element); err != nil {
		return err
	}

	if err := read(buf, &m.SubfieldIndex); err != nil {
		return err
	}

	if err := read(buf, &m.Operation); err != nil {
		return err
	}

	m.Data = make([]byte, 0)
	if err := readLen(buf, m.Data); err != nil {
		return err
	}

	return nil
}

// Entity Entity represents the details of a legal Entity, such as a public
// or private company, government agency, or and individual.
type Entity struct {
	Name                       Nvarchar8
	Type                       byte
	Address                    bool
	UnitNumber                 Nvarchar8
	BuildingNumber             Nvarchar8
	Street                     Nvarchar16
	SuburbCity                 Nvarchar8
	TerritoryStateProvinceCode []byte
	CountryCode                []byte
	PostalZIPCode              []byte
	EmailAddress               Nvarchar8
	PhoneNumber                Nvarchar8
	KeyRolesCount              uint8
	KeyRoles                   []KeyRole
	NotableRolesCount          uint8
	NotableRoles               []NotableRole
}

// NewEntity return a new Entity
func NewEntity() Entity {
	return Entity{}
}

// Serialize returns the byte representation of the message.
func (m Entity) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Name.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	if err := write(buf, m.Address); err != nil {
		return nil, err
	}

	{
		b, err := m.UnitNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.BuildingNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Street.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.SuburbCity.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.TerritoryStateProvinceCode); err != nil {
		return nil, err
	}

	if err := write(buf, m.CountryCode); err != nil {
		return nil, err
	}

	if err := write(buf, m.PostalZIPCode); err != nil {
		return nil, err
	}

	{
		b, err := m.EmailAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.PhoneNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.KeyRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		b, err := m.KeyRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.NotableRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		b, err := m.NotableRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (m *Entity) Write(buf *bytes.Buffer) error {

	if err := m.Name.Write(buf); err != nil {
		return err
	}
	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := read(buf, &m.Address); err != nil {
		return err
	}

	if err := m.UnitNumber.Write(buf); err != nil {
		return err
	}
	if err := m.BuildingNumber.Write(buf); err != nil {
		return err
	}
	if err := m.Street.Write(buf); err != nil {
		return err
	}
	if err := m.SuburbCity.Write(buf); err != nil {
		return err
	}
	m.TerritoryStateProvinceCode = make([]byte, 5)
	if err := readLen(buf, m.TerritoryStateProvinceCode); err != nil {
		return err
	}

	m.CountryCode = make([]byte, 3)
	if err := readLen(buf, m.CountryCode); err != nil {
		return err
	}

	m.PostalZIPCode = make([]byte, 12)
	if err := readLen(buf, m.PostalZIPCode); err != nil {
		return err
	}

	if err := m.EmailAddress.Write(buf); err != nil {
		return err
	}
	if err := m.PhoneNumber.Write(buf); err != nil {
		return err
	}
	if err := read(buf, &m.KeyRolesCount); err != nil {
		return err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		x := &KeyRole{}
		if err := x.Write(buf); err != nil {
			return err
		}

		m.KeyRoles = append(m.KeyRoles, *x)
	}
	if err := read(buf, &m.NotableRolesCount); err != nil {
		return err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		x := &NotableRole{}
		if err := x.Write(buf); err != nil {
			return err
		}

		m.NotableRoles = append(m.NotableRoles, *x)
	}
	return nil
}

// Header Header contains common details required by every Tokenized
// message.
type Header struct {
	ProtocolID       []byte
	OpPushdata       byte
	LenActionPayload []byte
	Version          uint8
	ActionPrefix     []byte
}

// NewHeader return a new Header
func NewHeader() Header {
	return Header{}
}

// Serialize returns the byte representation of the message.
func (m Header) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.ProtocolID); err != nil {
		return nil, err
	}

	if err := write(buf, m.OpPushdata); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenActionPayload); err != nil {
		return nil, err
	}

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, m.ActionPrefix); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Header) Write(buf *bytes.Buffer) error {

	m.ProtocolID = make([]byte, 13)
	if err := readLen(buf, m.ProtocolID); err != nil {
		return err
	}

	if err := read(buf, &m.OpPushdata); err != nil {
		return err
	}

	l := lenForOpPushdata(m.OpPushdata)
	m.LenActionPayload = make([]byte, l, l)
	if err := readLen(buf, m.LenActionPayload); err != nil {
		return err
	}
	if err := read(buf, &m.Version); err != nil {
		return err
	}

	m.ActionPrefix = make([]byte, 2)
	if err := readLen(buf, m.ActionPrefix); err != nil {
		return err
	}

	return nil
}

// KeyRole KeyRole is used to refer to a key role in an Entity.
type KeyRole struct {
	Type byte
	Name Nvarchar8
}

// NewKeyRole return a new KeyRole
func NewKeyRole() KeyRole {
	return KeyRole{}
}

// Serialize returns the byte representation of the message.
func (m KeyRole) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	{
		b, err := m.Name.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (m *KeyRole) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := m.Name.Write(buf); err != nil {
		return err
	}
	return nil
}

// NotableRole NotableRole is used to refer to a role of note in an Entity.
type NotableRole struct {
	Type byte
	Name Nvarchar8
}

// NewNotableRole return a new NotableRole
func NewNotableRole() NotableRole {
	return NotableRole{}
}

// Serialize returns the byte representation of the message.
func (m NotableRole) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	{
		b, err := m.Name.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (m *NotableRole) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := m.Name.Write(buf); err != nil {
		return err
	}
	return nil
}

// QuantityIndex A QuantityIndex contains a quantity, and an index. The
// quantity could be used to describe a number of tokens, or a value. The
// index is used to refer to an input index position.
type QuantityIndex struct {
	Index    uint16
	Quantity uint64
}

// NewQuantityIndex return a new QuantityIndex
func NewQuantityIndex() QuantityIndex {
	return QuantityIndex{}
}

// Serialize returns the byte representation of the message.
func (m QuantityIndex) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Index); err != nil {
		return nil, err
	}

	if err := write(buf, m.Quantity); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *QuantityIndex) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Index); err != nil {
		return err
	}

	if err := read(buf, &m.Quantity); err != nil {
		return err
	}

	return nil
}

// Registry A Registry defines the details of a public Registry.
type Registry struct {
	Name      Nvarchar8
	URL       Nvarchar8
	PublicKey Nvarchar8
}

// NewRegistry return a new Registry
func NewRegistry() Registry {
	return Registry{}
}

// Serialize returns the byte representation of the message.
func (m Registry) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Name.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.URL.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.PublicKey.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (m *Registry) Write(buf *bytes.Buffer) error {

	if err := m.Name.Write(buf); err != nil {
		return err
	}
	if err := m.URL.Write(buf); err != nil {
		return err
	}
	if err := m.PublicKey.Write(buf); err != nil {
		return err
	}
	return nil
}

// TargetAddress A TargetAddress defines a public address and quantity.
type TargetAddress struct {
	Address  []byte
	Quantity uint64
}

// NewTargetAddress return a new TargetAddress
func NewTargetAddress() TargetAddress {
	return TargetAddress{}
}

// Serialize returns the byte representation of the message.
func (m TargetAddress) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Address); err != nil {
		return nil, err
	}

	if err := write(buf, m.Quantity); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *TargetAddress) Write(buf *bytes.Buffer) error {

	m.Address = make([]byte, 34)
	if err := readLen(buf, m.Address); err != nil {
		return err
	}

	if err := read(buf, &m.Quantity); err != nil {
		return err
	}

	return nil
}

// TokenReceiver A TokenReceiver is contains a quantity, index, and
// registry token deails. The quantity could be used to describe a number
// of tokens, or a value. The index is used to refer to an input index
// position. The registry token details include the type of algorithm, and
// the token.
type TokenReceiver struct {
	Index                        uint16
	Quantity                     uint64
	RegistrySigAlgorithm         uint8
	RegistryConfirmationSigToken Nvarchar8
}

// NewTokenReceiver return a new TokenReceiver
func NewTokenReceiver() TokenReceiver {
	return TokenReceiver{}
}

// Serialize returns the byte representation of the message.
func (m TokenReceiver) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Index); err != nil {
		return nil, err
	}

	if err := write(buf, m.Quantity); err != nil {
		return nil, err
	}

	if err := write(buf, m.RegistrySigAlgorithm); err != nil {
		return nil, err
	}

	{
		b, err := m.RegistryConfirmationSigToken.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (m *TokenReceiver) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Index); err != nil {
		return err
	}

	if err := read(buf, &m.Quantity); err != nil {
		return err
	}

	if err := read(buf, &m.RegistrySigAlgorithm); err != nil {
		return err
	}

	if err := m.RegistryConfirmationSigToken.Write(buf); err != nil {
		return err
	}
	return nil
}

// VotingSystem A VotingSystem defines all details of a Voting System.
type VotingSystem struct {
	Name                        Nvarchar8
	System                      []byte
	Method                      byte
	Logic                       byte
	ThresholdPercentage         uint8
	VoteMultiplierPermitted     byte
	InitiativeThreshold         float32
	InitiativeThresholdCurrency []byte
}

// NewVotingSystem return a new VotingSystem
func NewVotingSystem() VotingSystem {
	return VotingSystem{}
}

// Serialize returns the byte representation of the message.
func (m VotingSystem) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Name.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.System); err != nil {
		return nil, err
	}

	if err := write(buf, m.Method); err != nil {
		return nil, err
	}

	if err := write(buf, m.Logic); err != nil {
		return nil, err
	}

	if err := write(buf, m.ThresholdPercentage); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMultiplierPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeThreshold); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeThresholdCurrency); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *VotingSystem) Write(buf *bytes.Buffer) error {

	if err := m.Name.Write(buf); err != nil {
		return err
	}
	m.System = make([]byte, 8)
	if err := readLen(buf, m.System); err != nil {
		return err
	}

	if err := read(buf, &m.Method); err != nil {
		return err
	}

	if err := read(buf, &m.Logic); err != nil {
		return err
	}

	if err := read(buf, &m.ThresholdPercentage); err != nil {
		return err
	}

	if err := read(buf, &m.VoteMultiplierPermitted); err != nil {
		return err
	}

	if err := read(buf, &m.InitiativeThreshold); err != nil {
		return err
	}

	m.InitiativeThresholdCurrency = make([]byte, 3)
	if err := readLen(buf, m.InitiativeThresholdCurrency); err != nil {
		return err
	}

	return nil
}
