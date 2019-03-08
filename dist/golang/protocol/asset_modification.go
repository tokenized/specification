// AssetModification : Asset Modification Action - Token Dilutions, Call
// Backs/Revocations, burning etc.
package protocol

type AssetModification struct {
	Header Header
	TextEncoding uint8
	AssetRevision uint64
	ModificationCount uint8
	Modifications []Amendment
	RefTxID []byte
}

// NewAssetModification returns a new AssetModification with defaults set.
func NewAssetModification() AssetModification {
	return AssetModification{}
}

// Type returns the type identifer for this message.
func (m AssetModification) Type() string {
	return CodeAssetModification
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m AssetModification) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m AssetModification) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.ModificationCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Modifications[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.RefTxID, 32)); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAssetModification, len(b))
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
func (m *AssetModification) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ModificationCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Modifications = append(m.Modifications, *x)
	}

	m.RefTxID = make([]byte, 32)
	if err := readLen(buf, m.RefTxID); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m AssetModification) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m AssetModification) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", m.AssetRevision))
	vals = append(vals, fmt.Sprintf("ModificationCount:%v", m.ModificationCount))
	vals = append(vals, fmt.Sprintf("Modifications:%#+v", m.Modifications))
	vals = append(vals, fmt.Sprintf("RefTxID:%#x", m.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
