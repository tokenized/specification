// Settlement : Settlement Action - Finalizes the transfer of bitcoins and
// tokens from send, exchange, and swap actions.
package protocol

type Settlement struct {
	Header Header
	TextEncoding uint8
	TransferType byte
	AssetType1 []byte
	AssetID1 []byte
	AssetType2 []byte
	AssetID2 []byte
	Asset1SettlementsCount uint8
	Asset1AddressesXQty []QuantityIndex
	Asset2SettlementsCount uint8
	Asset2AddressXQty []QuantityIndex
	Timestamp uint64
}

// NewSettlement returns a new Settlement with defaults set.
func NewSettlement() Settlement {
	return Settlement{}
}

// Type returns the type identifer for this message.
func (m Settlement) Type() string {
	return CodeSettlement
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Settlement) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Settlement) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransferType); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetType1, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID1, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetType2, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID2, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.Asset1SettlementsCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.Asset1AddressesXQty[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Asset2SettlementsCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.Asset2AddressXQty[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeSettlement, len(b))
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
func (m *Settlement) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransferType); err != nil {
		return 0, err
	}

	m.AssetType1 = make([]byte, 3)
	if err := readLen(buf, m.AssetType1); err != nil {
		return 0, err
	}

	m.AssetID1 = make([]byte, 32)
	if err := readLen(buf, m.AssetID1); err != nil {
		return 0, err
	}

	m.AssetType2 = make([]byte, 3)
	if err := readLen(buf, m.AssetType2); err != nil {
		return 0, err
	}

	m.AssetID2 = make([]byte, 32)
	if err := readLen(buf, m.AssetID2); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Asset1SettlementsCount); err != nil {
		return 0, err
	}

	m.Asset1AddressesXQty = make([]QuantityIndex, m., m.)

	if err := read(buf, &m.Asset1AddressesXQty); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Asset2SettlementsCount); err != nil {
		return 0, err
	}

	m.Asset2AddressXQty = make([]QuantityIndex, m., m.)

	if err := read(buf, &m.Asset2AddressXQty); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Settlement) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Settlement) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("TransferType:%#+v", m.TransferType))
	vals = append(vals, fmt.Sprintf("AssetType1:%#x", m.AssetType1))
	vals = append(vals, fmt.Sprintf("AssetID1:%#x", m.AssetID1))
	vals = append(vals, fmt.Sprintf("AssetType2:%#x", m.AssetType2))
	vals = append(vals, fmt.Sprintf("AssetID2:%#x", m.AssetID2))
	vals = append(vals, fmt.Sprintf("Asset1SettlementsCount:%v", m.Asset1SettlementsCount))
	vals = append(vals, fmt.Sprintf("Asset1AddressesXQty:%#+v", m.Asset1AddressesXQty))
	vals = append(vals, fmt.Sprintf("Asset2SettlementsCount:%v", m.Asset2SettlementsCount))
	vals = append(vals, fmt.Sprintf("Asset2AddressXQty:%#+v", m.Asset2AddressXQty))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
