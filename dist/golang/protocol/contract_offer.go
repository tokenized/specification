// ContractOffer : The Contract Offer action allows the Issuer to tell the
// smart contract what they want the details (labels, data, T&C's, etc.) of
// the Contract to be on-chain in a public and immutable way. The Contract
// Offer action 'initializes' a generic smart contract that has been spun
// up by either the Smart Contract Operator or the Issuer. This on-chain
// action allows for the positive response from the smart contract with
// either a Contract Formation Action or a Rejection Action.
package protocol

type ContractOffer struct {
	Header Header
	TextEncoding uint8
	ContractName Nvarchar8
	ContractFileType uint8
	LenContractFile uint32
	ContractFile []byte
	GoverningLaw []byte
	Jurisdiction []byte
	ContractExpiration uint64
	ContractURI Nvarchar8
	IssuerName Nvarchar8
	IssuerType byte
	IssuerLogoURL Nvarchar8
	ContractOperatorID Nvarchar8
	ContractAuthFlags []byte
	VotingSystemCount uint8
	VotingSystems []VotingSystem
	RestrictedQtyAssets uint64
	ReferendumProposal bool
	InitiativeProposal bool
	RegistryCount uint8
	Registries []Registry
	IssuerAddress bool
	UnitNumber Nvarchar8
	BuildingNumber Nvarchar8
	Street Nvarchar16
	SuburbCity Nvarchar8
	TerritoryStateProvinceCode []byte
	CountryCode []byte
	PostalZIPCode Nvarchar8
	EmailAddress Nvarchar8
	PhoneNumber Nvarchar8
	KeyRolesCount uint8
	KeyRoles []KeyRole
	NotableRolesCount uint8
	NotableRoles []NotableRole
}

// NewContractOffer returns a new ContractOffer with defaults set.
func NewContractOffer() ContractOffer {
	return ContractOffer{}
}

// Type returns the type identifer for this message.
func (m ContractOffer) Type() string {
	return CodeContractOffer
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ContractOffer) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ContractOffer) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	{
		b, err := m.ContractName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFile, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.GoverningLaw, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	{
		b, err := m.ContractURI.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.IssuerName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	{
		b, err := m.IssuerLogoURL.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.ContractOperatorID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ContractAuthFlags, 16)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VotingSystemCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.VotingSystems[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.RestrictedQtyAssets); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.RegistryCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Registries[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerAddress); err != nil {
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

	if err := write(buf, pad(m.TerritoryStateProvinceCode, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.CountryCode, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.PostalZIPCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
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

	for i := 0; i < int(m.); i++ {
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

	for i := 0; i < int(m.); i++ {
		b, err := m.NotableRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeContractOffer, len(b))
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
func (m *ContractOffer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := m.ContractName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	m.ContractFile = make([]byte, 32)
	if err := readLen(buf, m.ContractFile); err != nil {
		return 0, err
	}

	m.GoverningLaw = make([]byte, 5)
	if err := readLen(buf, m.GoverningLaw); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	if err := m.ContractURI.Write(buf); err != nil {
		return 0, err
	}

	if err := m.IssuerName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	if err := m.IssuerLogoURL.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractOperatorID.Write(buf); err != nil {
		return 0, err
	}

	m.ContractAuthFlags = make([]byte, 16)
	if err := readLen(buf, m.ContractAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VotingSystemCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &VotingSystem{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.VotingSystems = append(m.VotingSystems, *x)
	}

	if err := read(buf, &m.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RegistryCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Registry{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Registries = append(m.Registries, *x)
	}

	if err := read(buf, &m.IssuerAddress); err != nil {
		return 0, err
	}

	if err := m.UnitNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.BuildingNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Street.Write(buf); err != nil {
		return 0, err
	}

	if err := m.SuburbCity.Write(buf); err != nil {
		return 0, err
	}

	m.TerritoryStateProvinceCode = make([]byte, 5)
	if err := readLen(buf, m.TerritoryStateProvinceCode); err != nil {
		return 0, err
	}

	m.CountryCode = make([]byte, 3)
	if err := readLen(buf, m.CountryCode); err != nil {
		return 0, err
	}

	if err := m.PostalZIPCode.Write(buf); err != nil {
		return 0, err
	}

	if err := m.EmailAddress.Write(buf); err != nil {
		return 0, err
	}

	if err := m.PhoneNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.KeyRolesCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &KeyRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.KeyRoles = append(m.KeyRoles, *x)
	}

	if err := read(buf, &m.NotableRolesCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &NotableRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.NotableRoles = append(m.NotableRoles, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m ContractOffer) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m ContractOffer) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", m.ContractName))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", m.ContractFileType))
	vals = append(vals, fmt.Sprintf("LenContractFile:%v", m.LenContractFile))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#x", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", m.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", m.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", m.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", m.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#x", m.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystemCount:%v", m.VotingSystemCount))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", m.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", m.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("RegistryCount:%v", m.RegistryCount))
	vals = append(vals, fmt.Sprintf("Registries:%#+v", m.Registries))
	vals = append(vals, fmt.Sprintf("IssuerAddress:%#+v", m.IssuerAddress))
	vals = append(vals, fmt.Sprintf("UnitNumber:%#+v", m.UnitNumber))
	vals = append(vals, fmt.Sprintf("BuildingNumber:%#+v", m.BuildingNumber))
	vals = append(vals, fmt.Sprintf("Street:%#+v", m.Street))
	vals = append(vals, fmt.Sprintf("SuburbCity:%#+v", m.SuburbCity))
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#x", m.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#x", m.CountryCode))
	vals = append(vals, fmt.Sprintf("PostalZIPCode:%#+v", m.PostalZIPCode))
	vals = append(vals, fmt.Sprintf("EmailAddress:%#+v", m.EmailAddress))
	vals = append(vals, fmt.Sprintf("PhoneNumber:%#+v", m.PhoneNumber))
	vals = append(vals, fmt.Sprintf("KeyRolesCount:%v", m.KeyRolesCount))
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", m.KeyRoles))
	vals = append(vals, fmt.Sprintf("NotableRolesCount:%v", m.NotableRolesCount))
	vals = append(vals, fmt.Sprintf("NotableRoles:%#+v", m.NotableRoles))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
