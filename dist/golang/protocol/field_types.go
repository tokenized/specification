package protocol

import "bytes"

// Address Address represents a public address
type Address struct {
	Address []byte // Public address where the token balance will be changed.
}

// NewAddress returns a new Address with defaults set.
func NewAddress() *Address {
	result := Address{}
	return &result
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

	m.Address = make([]byte, 20)
	if err := readLen(buf, m.Address); err != nil {
		return err
	}

	return nil
}

// Amendment An Amendment is used to describe the modification of a single
// field in a Contract or Asset, as defined in the ContractFormation and
// AssetCreation messages.
type Amendment struct {
	FieldIndex    uint8  // Index of the field to be amended.
	Element       uint16 // Specifies the element of the complex array type to be amended. This only applies to array types, and has no meaning for a simple type such as uint64, string, byte or byte[]. Specifying a value > 0 for a simple type will result in a Rejection.
	SubfieldIndex uint8  // Index of the subfield to be amended. This only applies to specific fields of an element in an array. This is used to specify which field of the array element the amendment applies to.
	DeleteElement bool   // 1 = remove the element listed in the Element field, 0 means this is not a delete operation. The DeleteElement field only applies to a particilar element of a complex array type. For example, it could be used to remove a particular VotingSystem from a Contract.
	Data          []byte // New data for the amended subfield. Data type depends on the the type of the field being amended.
}

// NewAmendment returns a new Amendment with defaults set.
func NewAmendment() *Amendment {
	result := Amendment{}
	return &result
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

	if err := write(buf, m.DeleteElement); err != nil {
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

	if err := read(buf, &m.DeleteElement); err != nil {
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
	Name                       string        // Length 1-255 bytes (0 is not valid). Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	Type                       byte          // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	Address                    bool          // Registered/Physical/mailing address(HQ). Y-1/N-0, N means there is no issuer address.
	UnitNumber                 string        // Issuer/Entity/Contracting Party X Address Details (eg. HQ)
	BuildingNumber             string        //
	Street                     string        //
	SuburbCity                 string        //
	TerritoryStateProvinceCode string        //
	CountryCode                string        //
	PostalZIPCode              string        //
	EmailAddress               string        // Length 0-255 bytes. Address for text-based communication: eg. email address, Bitcoin address
	PhoneNumber                string        // Length 0-50 bytes. 0 is valid. Phone Number for Entity.
	KeyRolesCount              uint8         // Number of key roles associated with the issuing entity.  (eg. Directors, etc.) 0-255. 0 is valid.
	KeyRoles                   []KeyRole     // A list of Key Roles.
	NotableRolesCount          uint8         // Number of notable roles associated with the issuing entity.  (eg. Corporate Officers, Managers, etc.) 0-255. 0 is valid.
	NotableRoles               []NotableRole // A list of Notable Roles.
}

// NewEntity returns a new Entity with defaults set.
func NewEntity() *Entity {
	result := Entity{}
	return &result
}

// Serialize returns the byte representation of the message.
func (m Entity) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := WriteVarChar(buf, m.Name, 11); err != nil {
		return nil, err
	}

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	if err := write(buf, m.Address); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.UnitNumber, 2); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.BuildingNumber, 6); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Street, 14); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.SuburbCity, 8); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.CountryCode, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.PostalZIPCode, 12); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.EmailAddress, 20); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PhoneNumber, 11); err != nil {
		return nil, err
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

	if err := read(buf, &m.Name); err != nil {
		return err
	}

	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := read(buf, &m.Address); err != nil {
		return err
	}

	if err := read(buf, &m.UnitNumber); err != nil {
		return err
	}

	if err := read(buf, &m.BuildingNumber); err != nil {
		return err
	}

	if err := read(buf, &m.Street); err != nil {
		return err
	}

	if err := read(buf, &m.SuburbCity); err != nil {
		return err
	}

	if err := read(buf, &m.TerritoryStateProvinceCode); err != nil {
		return err
	}

	if err := read(buf, &m.CountryCode); err != nil {
		return err
	}

	if err := read(buf, &m.PostalZIPCode); err != nil {
		return err
	}

	if err := read(buf, &m.EmailAddress); err != nil {
		return err
	}

	if err := read(buf, &m.PhoneNumber); err != nil {
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
	ProtocolID       []byte // Tokenized ID Prefix.  tokenized.com
	OpPushdata       byte   // PACKET LENGTH, PUSHDATA1 (76), PUSHDATA2 (77), or PUSHDATA4 (78) depending on total size of action payload.
	LenActionPayload []byte // Length of the action message (0 - 65,535 bytes). 0 if pushdata length <76B, 1 byte if PUSHDATA1 is used, 2 bytes if PUSHDATA2 and 4 bytes if PUSHDATA4.
	Version          uint8  // 255 reserved for additional versions. Tokenized protocol versioning.
	ActionPrefix     []byte // Contract Offer: The Contract Offer Action allows the Issuer to initialize a smart contract by providing all the necessary information, including T&C's.  The Contract Offer Action can also be used to signal to a market actor that they want to buy/form a contract.
}

// NewHeader returns a new Header with defaults set.
func NewHeader() *Header {
	result := Header{}
	return &result
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
	Type uint8  // Chairman, Director. Found in 'Roles' in Specification/Resources
	Name string // Length 0-255 bytes. 0 is valid. Name (eg. John Alexander Smith)
}

// NewKeyRole returns a new KeyRole with defaults set.
func NewKeyRole(roleType uint8, name string) *KeyRole {
	result := KeyRole{}
	result.Type = roleType
	result.Name = name
	return &result
}

// Serialize returns the byte representation of the message.
func (m KeyRole) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Name, 14); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *KeyRole) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := read(buf, &m.Name); err != nil {
		return err
	}

	return nil
}

// NotableRole NotableRole is used to refer to a role of note in an Entity.
type NotableRole struct {
	Type uint8  // Found in 'Roles' in Specification/Resources
	Name string // Length 0-255 bytes. 0 is valid. Name (eg. John Alexander Smith)
}

// NewNotableRole returns a new NotableRole with defaults set.
func NewNotableRole(roleType uint8, name string) *NotableRole {
	result := NotableRole{}
	result.Type = roleType
	result.Name = name
	return &result
}

// Serialize returns the byte representation of the message.
func (m NotableRole) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Type); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Name, 14); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *NotableRole) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Type); err != nil {
		return err
	}

	if err := read(buf, &m.Name); err != nil {
		return err
	}

	return nil
}

// QuantityIndex A QuantityIndex contains a quantity, and an index. The
// quantity could be used to describe a number of tokens, or a value. The
// index is used to refer to an input index position.
type QuantityIndex struct {
	Index    uint16 // The index of the input sending the tokens
	Quantity uint64 // Number of tokens being sent
}

// NewQuantityIndex returns a new QuantityIndex with defaults set.
func NewQuantityIndex() *QuantityIndex {
	result := QuantityIndex{}
	return &result
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
	Name      string // Length 0-255 bytes. 0 is valid. Registry X Name (eg. Coinbase, Tokenized, etc.)
	URL       string // Length 0-255 bytes. 0 is valid. If applicable: URL for REST/RPC Endpoint
	PublicKey string // Length 0-255 bytes. 0 is not valid. Registry Public Key (eg. Bitcoin Public key), used to confirm digital signed proofs for transfers.  Can also be the same public address that controls a Tokenized Registry.
}

// NewRegistry returns a new Registry with defaults set.
func NewRegistry() *Registry {
	result := Registry{}
	return &result
}

// Serialize returns the byte representation of the message.
func (m Registry) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := WriteVarChar(buf, m.Name, 10); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.URL, 53); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PublicKey, 1); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Registry) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Name); err != nil {
		return err
	}

	if err := read(buf, &m.URL); err != nil {
		return err
	}

	if err := read(buf, &m.PublicKey); err != nil {
		return err
	}

	return nil
}

// TargetAddress A TargetAddress defines a public address and quantity.
type TargetAddress struct {
	Address  Address // Public address where the token balance will be changed.
	Quantity uint64  // Qty of tokens to be frozen, thawed, confiscated or reconciled. For Contract-wide freezes 0 will be used.
}

// NewTargetAddress returns a new TargetAddress with defaults set.
func NewTargetAddress() *TargetAddress {
	result := TargetAddress{}
	return &result
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

	if err := read(buf, &m.Address); err != nil {
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
	Index                        uint16 // The index of the output receiving the tokens
	Quantity                     uint64 // Number of tokens to be received by address at Output X
	RegistrySigAlgorithm         uint8  // 0 = No Registry-signed Message, 1 = ECDSA+secp256k1
	RegistryConfirmationSigToken string // Length 0-255 bytes. IF restricted to a registry (whitelist) or has transfer restrictions (age, location, investor status): ECDSA+secp256k1 (or the like) signed message provided by an approved/trusted registry through an API signature of [Contract Address + Asset Code + Public Address + Blockhash of the Latest Block + Block Height + Confirmed/Rejected Bool]. If no transfer restrictions(trade restriction/age restriction fields in the Asset Type payload. or restricted to a whitelist by the Contract Auth Flags, it is a NULL field.
}

// NewTokenReceiver returns a new TokenReceiver with defaults set.
func NewTokenReceiver() *TokenReceiver {
	result := TokenReceiver{}
	return &result
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

	if err := WriteVarChar(buf, m.RegistryConfirmationSigToken, 255); err != nil {
		return nil, err
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

	if err := read(buf, &m.RegistryConfirmationSigToken); err != nil {
		return err
	}

	return nil
}

// VotingSystem A VotingSystem defines all details of a Voting System.
type VotingSystem struct {
	Name                        string  // eg. Special Resolutions, Ordinary Resolutions, Fundamental Matters, General Matters, Directors' Vote, Poll, etc.
	System                      []byte  // Specifies which subfield is subject to this vote system's control.
	Method                      byte    // R - Relative Threshold, A - Absolute Threshold, P - Plurality,  (Relative Threshold means the number of counted votes must exceed the threshold % of total ballots cast.  Abstentations/spoiled votes do not detract from the liklihood of a vote passing as they are not included in the denominator.  Absolute Threshold requires the number of ballots counted to exceed the threshold value when compared to the total outstanding tokens.  Abstentations/spoiled votes detract from the liklihood of the vote passing.  For example, in an absolute threshold vote, if the threshold was 50% and 51% of the total outstanding tokens did not vote, then the vote cannot pass.  50% of all tokens would have had to vote for one vote option for the vote to be successful.
	Logic                       byte    // 0 - Standard Scoring (+1 * # of tokens owned), 1 - Weighted Scoring (1st choice * Vote Max * # of tokens held, 2nd choice * Vote Max-1 * # of tokens held,..etc.)
	ThresholdPercentage         uint8   // 1-100 is valid for relative threshold and absolute threshold. (eg. 75 means 75% and greater). 0 & >=101 is invalid and will be rejected by the smart contract.  Only applicable to Relative and Absolute Threshold vote methods.  The Plurality vote method requires no threshold value (NULL), as the successful vote option is simply selected on the basis of highest ballots cast for it.
	VoteMultiplierPermitted     byte    // Y - Yes, N - No. Where an asset has a vote multiplier, Y must be selected here for the vote multiplier to count, otherwise votes are simply treated as 1x per token.
	InitiativeThreshold         float32 // Token Owners must pay the threshold amount to broadcast a valid Initiative.  If the Initiative action is valid, the smart contract will start a vote. 0 is valid.
	InitiativeThresholdCurrency string  // Currency.  Always paid in BSV or a currency token (CUR) at current market valuations in the currency listed. NULL is valid.
}

// NewVotingSystem returns a new VotingSystem with defaults set.
func NewVotingSystem() *VotingSystem {
	result := VotingSystem{}
	return &result
}

// Serialize returns the byte representation of the message.
func (m VotingSystem) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := WriteVarChar(buf, m.Name, 20); err != nil {
		return nil, err
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

	if err := WriteFixedChar(buf, m.InitiativeThresholdCurrency, 3); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *VotingSystem) Write(buf *bytes.Buffer) error {

	if err := read(buf, &m.Name); err != nil {
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

	if err := read(buf, &m.InitiativeThresholdCurrency); err != nil {
		return err
	}

	return nil
}
