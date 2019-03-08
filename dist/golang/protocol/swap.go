// Swap : Swap Action - Two (or more) parties want to swap a token (Atomic
// Swap) directly for another token. BSV is not used in the transaction
// other than for paying the necessary network/transaction fees.
package protocol

type Swap struct {
	Header Header
	TextEncoding uint8
	AssetType1 []byte
	AssetID1 []byte
	AssetType2 []byte
	AssetID2 []byte
	OfferExpiry uint64
	ExchangeFeeCurrency []byte
	ExchangeFeeVar float32
	ExchangeFeeFixed float32
	ExchangeFeeAddress []byte
	Asset1SenderCount uint8
	Asset1Senders []QuantityIndex
	Asset1ReceiverCount uint8
	Asset1Receivers []TokenReceiver
	Asset2SenderCount uint8
	Asset2Senders []QuantityIndex
	Asset2ReceiverCount uint8
	Asset2Receivers []TokenReceiver
}

// NewSwap returns a new Swap with defaults set.
func NewSwap() Swap {
	return Swap{}
}

// Type returns the type identifer for this message.
func (m Swap) Type() string {
	return CodeSwap
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Swap) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Swap) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
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

	if err := write(buf, m.OfferExpiry); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ExchangeFeeCurrency, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeFixed); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ExchangeFeeAddress, 34)); err != nil {
		return nil, err
	}

	if err := write(buf, m.Asset1SenderCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.Asset1Senders[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Asset1ReceiverCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Asset1Receivers[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Asset2SenderCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.Asset2Senders[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Asset2ReceiverCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Asset2Receivers[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeSwap, len(b))
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
func (m *Swap) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
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

	if err := read(buf, &m.OfferExpiry); err != nil {
		return 0, err
	}

	m.ExchangeFeeCurrency = make([]byte, 3)
	if err := readLen(buf, m.ExchangeFeeCurrency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExchangeFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExchangeFeeFixed); err != nil {
		return 0, err
	}

	m.ExchangeFeeAddress = make([]byte, 34)
	if err := readLen(buf, m.ExchangeFeeAddress); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Asset1SenderCount); err != nil {
		return 0, err
	}

	m.Asset1Senders = make([]QuantityIndex, m., m.)

	if err := read(buf, &m.Asset1Senders); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Asset1ReceiverCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &TokenReceiver{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Asset1Receivers = append(m.Asset1Receivers, *x)
	}

	if err := read(buf, &m.Asset2SenderCount); err != nil {
		return 0, err
	}

	m.Asset2Senders = make([]QuantityIndex, m., m.)

	if err := read(buf, &m.Asset2Senders); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Asset2ReceiverCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &TokenReceiver{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Asset2Receivers = append(m.Asset2Receivers, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Swap) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Swap) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType1:%#x", m.AssetType1))
	vals = append(vals, fmt.Sprintf("AssetID1:%#x", m.AssetID1))
	vals = append(vals, fmt.Sprintf("AssetType2:%#x", m.AssetType2))
	vals = append(vals, fmt.Sprintf("AssetID2:%#x", m.AssetID2))
	vals = append(vals, fmt.Sprintf("OfferExpiry:%v", m.OfferExpiry))
	vals = append(vals, fmt.Sprintf("ExchangeFeeCurrency:%#x", m.ExchangeFeeCurrency))
	vals = append(vals, fmt.Sprintf("ExchangeFeeVar:%v", m.ExchangeFeeVar))
	vals = append(vals, fmt.Sprintf("ExchangeFeeFixed:%v", m.ExchangeFeeFixed))
	vals = append(vals, fmt.Sprintf("ExchangeFeeAddress:%#x", m.ExchangeFeeAddress))
	vals = append(vals, fmt.Sprintf("Asset1SenderCount:%v", m.Asset1SenderCount))
	vals = append(vals, fmt.Sprintf("Asset1Senders:%#+v", m.Asset1Senders))
	vals = append(vals, fmt.Sprintf("Asset1ReceiverCount:%v", m.Asset1ReceiverCount))
	vals = append(vals, fmt.Sprintf("Asset1Receivers:%#+v", m.Asset1Receivers))
	vals = append(vals, fmt.Sprintf("Asset2SenderCount:%v", m.Asset2SenderCount))
	vals = append(vals, fmt.Sprintf("Asset2Senders:%#+v", m.Asset2Senders))
	vals = append(vals, fmt.Sprintf("Asset2ReceiverCount:%v", m.Asset2ReceiverCount))
	vals = append(vals, fmt.Sprintf("Asset2Receivers:%#+v", m.Asset2Receivers))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
