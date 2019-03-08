// Order : Order Action - Issuer to signal to the smart contract that the
// tokens that a particular public address(es) owns are to be confiscated,
// frozen, thawed or reconciled.
package protocol

type Order struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	ComplianceAction byte
	TargetAddressCount uint16
	TargetAddresses []TargetAddress
	DepositAddress Nvarchar8
	EnforcementAuthorityName Nvarchar8
	SigAlgoAddressList uint8
	EnforcementAuthorityPublicKey Nvarchar8
	OrderSignature Nvarchar8
	SupportingEvidenceHash []byte
	RefTxnID []byte
	FreezePeriod uint64
	Message Nvarchar16
}

// NewOrder returns a new Order with defaults set.
func NewOrder() Order {
	return Order{}
}

// Type returns the type identifer for this message.
func (m Order) Type() string {
	return CodeOrder
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Order) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Order) Serialize() ([]byte, error) {
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

	if err := write(buf, m.ComplianceAction); err != nil {
		return nil, err
	}

	if err := write(buf, m.TargetAddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.TargetAddresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.DepositAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.EnforcementAuthorityName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.SigAlgoAddressList); err != nil {
		return nil, err
	}

	{
		b, err := m.EnforcementAuthorityPublicKey.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.OrderSignature.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.SupportingEvidenceHash, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.RefTxnID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.FreezePeriod); err != nil {
		return nil, err
	}

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeOrder, len(b))
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
func (m *Order) Write(b []byte) (int, error) {
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

	if err := read(buf, &m.ComplianceAction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TargetAddressCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &TargetAddress{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.TargetAddresses = append(m.TargetAddresses, *x)
	}

	if err := m.DepositAddress.Write(buf); err != nil {
		return 0, err
	}

	if err := m.EnforcementAuthorityName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.SigAlgoAddressList); err != nil {
		return 0, err
	}

	if err := m.EnforcementAuthorityPublicKey.Write(buf); err != nil {
		return 0, err
	}

	if err := m.OrderSignature.Write(buf); err != nil {
		return 0, err
	}

	m.SupportingEvidenceHash = make([]byte, 32)
	if err := readLen(buf, m.SupportingEvidenceHash); err != nil {
		return 0, err
	}

	m.RefTxnID = make([]byte, 32)
	if err := readLen(buf, m.RefTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.FreezePeriod); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Order) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Order) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("ComplianceAction:%#+v", m.ComplianceAction))
	vals = append(vals, fmt.Sprintf("TargetAddressCount:%v", m.TargetAddressCount))
	vals = append(vals, fmt.Sprintf("TargetAddresses:%#+v", m.TargetAddresses))
	vals = append(vals, fmt.Sprintf("DepositAddress:%#+v", m.DepositAddress))
	vals = append(vals, fmt.Sprintf("EnforcementAuthorityName:%#+v", m.EnforcementAuthorityName))
	vals = append(vals, fmt.Sprintf("SigAlgoAddressList:%v", m.SigAlgoAddressList))
	vals = append(vals, fmt.Sprintf("EnforcementAuthorityPublicKey:%#+v", m.EnforcementAuthorityPublicKey))
	vals = append(vals, fmt.Sprintf("OrderSignature:%#+v", m.OrderSignature))
	vals = append(vals, fmt.Sprintf("SupportingEvidenceHash:%#x", m.SupportingEvidenceHash))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#x", m.RefTxnID))
	vals = append(vals, fmt.Sprintf("FreezePeriod:%v", m.FreezePeriod))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
