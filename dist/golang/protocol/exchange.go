// Exchange : Exchange Action - Tokens exchanged for Bitcoin (BSV).
// Example: Bob (Token Sender) to sell 21,000 tokens to Alice (Token
// Receiver) for 7 BSV. Both parties must sign the transaction for it to be
// valid.
package protocol

type Exchange struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	OfferExpiry uint64
	ExchangeFeeCurrency []byte
	ExchangeFeeVar float32
	ExchangeFeeFixed float32
	ExchangeFeeAddress []byte
	TokenSenderCount uint8
	TokenSenders []QuantityIndex
	TokenReceiverCount uint8
	TokenReceivers []TokenReceiver
}

// NewExchange returns a new Exchange with defaults set.
func NewExchange() Exchange {
	return Exchange{}
}

// Type returns the type identifer for this message.
func (m Exchange) Type() string {
	return CodeExchange
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Exchange) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Exchange) Serialize() ([]byte, error) {
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

	if err := write(buf, m.TokenSenderCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.TokenSenders[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.TokenReceiverCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.TokenReceivers[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeExchange, len(b))
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
func (m *Exchange) Write(b []byte) (int, error) {
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

	if err := read(buf, &m.TokenSenderCount); err != nil {
		return 0, err
	}

	m.TokenSenders = make([]QuantityIndex, m., m.)

	if err := read(buf, &m.TokenSenders); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TokenReceiverCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &TokenReceiver{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.TokenReceivers = append(m.TokenReceivers, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Exchange) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Exchange) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("OfferExpiry:%v", m.OfferExpiry))
	vals = append(vals, fmt.Sprintf("ExchangeFeeCurrency:%#x", m.ExchangeFeeCurrency))
	vals = append(vals, fmt.Sprintf("ExchangeFeeVar:%v", m.ExchangeFeeVar))
	vals = append(vals, fmt.Sprintf("ExchangeFeeFixed:%v", m.ExchangeFeeFixed))
	vals = append(vals, fmt.Sprintf("ExchangeFeeAddress:%#x", m.ExchangeFeeAddress))
	vals = append(vals, fmt.Sprintf("TokenSenderCount:%v", m.TokenSenderCount))
	vals = append(vals, fmt.Sprintf("TokenSenders:%#+v", m.TokenSenders))
	vals = append(vals, fmt.Sprintf("TokenReceiverCount:%v", m.TokenReceiverCount))
	vals = append(vals, fmt.Sprintf("TokenReceivers:%#+v", m.TokenReceivers))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
