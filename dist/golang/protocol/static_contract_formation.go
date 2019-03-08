// StaticContract Formation : Static Contract Formation Action
package protocol

type StaticContractFormation struct {
	Header Header
	TextEncoding uint8
	ContractName Nvarchar8
	ContractType Nvarchar8
	ContractFileType uint8
	LenContractFile uint32
	ContractFile []byte
	ContractRevision uint16
	GoverningLaw []byte
	Jurisdiction []byte
	EffectiveDate uint64
	ContractExpiration uint64
	ContractURI Nvarchar8
	PrevRevTxID []byte
	EntityCount uint8
	Entities []Entity
}

// NewStaticContractFormation returns a new StaticContractFormation with defaults set.
func NewStaticContractFormation() StaticContractFormation {
	return StaticContractFormation{}
}

// Type returns the type identifer for this message.
func (m StaticContractFormation) Type() string {
	return CodeStaticContractFormation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m StaticContractFormation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m StaticContractFormation) Serialize() ([]byte, error) {
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

	{
		b, err := m.ContractType.Serialize()
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

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.GoverningLaw, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EffectiveDate); err != nil {
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

	if err := write(buf, pad(m.PrevRevTxID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EntityCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Entities[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeStaticContractFormation, len(b))
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
func (m *StaticContractFormation) Write(b []byte) (int, error) {
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

	if err := m.ContractType.Write(buf); err != nil {
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

	if err := read(buf, &m.ContractRevision); err != nil {
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

	if err := read(buf, &m.EffectiveDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	if err := m.ContractURI.Write(buf); err != nil {
		return 0, err
	}

	m.PrevRevTxID = make([]byte, 32)
	if err := readLen(buf, m.PrevRevTxID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EntityCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Entity{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Entities = append(m.Entities, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m StaticContractFormation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m StaticContractFormation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", m.ContractName))
	vals = append(vals, fmt.Sprintf("ContractType:%#+v", m.ContractType))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", m.ContractFileType))
	vals = append(vals, fmt.Sprintf("LenContractFile:%v", m.LenContractFile))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#x", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("EffectiveDate:%v", m.EffectiveDate))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("PrevRevTxID:%#x", m.PrevRevTxID))
	vals = append(vals, fmt.Sprintf("EntityCount:%v", m.EntityCount))
	vals = append(vals, fmt.Sprintf("Entities:%#+v", m.Entities))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
