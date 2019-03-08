// AssetDefinition : Asset Definition Action - This action is used by the
// issuer to define the properties/characteristics of the Asset (token)
// that it wants to create.
package protocol

type AssetDefinition struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	AssetAuthFlags []byte
	TransfersPermitted bool
	TradeRestrictions []byte
	EnforcementOrdersPermitted bool
	VoteMultiplier uint8
	ReferendumProposal bool
	InitiativeProposal bool
	AssetModificationGovernance bool
	TokenQty uint64
	ContractFeeCurrency []byte
	ContractFeeVar float32
	ContractFeeFixed float32
	AssetPayloadLen uint16
	AssetPayload []byte
}

// NewAssetDefinition returns a new AssetDefinition with defaults set.
func NewAssetDefinition() AssetDefinition {
	return AssetDefinition{}
}

// Type returns the type identifer for this message.
func (m AssetDefinition) Type() string {
	return CodeAssetDefinition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m AssetDefinition) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m AssetDefinition) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradeRestrictions, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMultiplier); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetModificationGovernance); err != nil {
		return nil, err
	}

	if err := write(buf, m.TokenQty); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFeeCurrency, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeFixed); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetPayloadLen); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.AssetPayload[i]); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAssetDefinition, len(b))
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
func (m *AssetDefinition) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	m.TradeRestrictions = make([]byte, 3)
	if err := readLen(buf, m.TradeRestrictions); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMultiplier); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetModificationGovernance); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TokenQty); err != nil {
		return 0, err
	}

	m.ContractFeeCurrency = make([]byte, 3)
	if err := readLen(buf, m.ContractFeeCurrency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeFixed); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetPayloadLen); err != nil {
		return 0, err
	}

	m.AssetPayload = make([]byte, m., m.)

	if err := read(buf, &m.AssetPayload); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m AssetDefinition) PayloadMessage() (PayloadMessage, error) {
	p, err := NewPayloadMessageFromCode(m.AssetType)
	if p == nil || err != nil {
		return nil, err
	}

	if _, err := p.Write(m.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (m AssetDefinition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#x", m.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", m.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#x", m.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", m.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", m.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", m.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", m.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#x", m.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", m.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", m.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayloadLen:%v", m.AssetPayloadLen))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", m.AssetPayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
