// Send : Send Action - A Token Owner Sends a Token to a Receiver. The Send
// Action requires no sign-off by the Token Receiving Party and does not
// provide any on-chain consideration to the Token Sending Party. Can be
// used for redeeming a ticket, coupon, points, etc.
package protocol

type Send struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	TokenSenderCount uint8
	TokenSenders []QuantityIndex
	TokenReceiverCount uint8
	TokenReceivers []TokenReceiver
}

// NewSend returns a new Send with defaults set.
func NewSend() Send {
	return Send{}
}

// Type returns the type identifer for this message.
func (m Send) Type() string {
	return CodeSend
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Send) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Send) Serialize() ([]byte, error) {
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

	header, err := NewHeaderForCode(CodeSend, len(b))
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
func (m *Send) Write(b []byte) (int, error) {
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
func (m Send) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Send) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("TokenSenderCount:%v", m.TokenSenderCount))
	vals = append(vals, fmt.Sprintf("TokenSenders:%#+v", m.TokenSenders))
	vals = append(vals, fmt.Sprintf("TokenReceiverCount:%v", m.TokenReceiverCount))
	vals = append(vals, fmt.Sprintf("TokenReceivers:%#+v", m.TokenReceivers))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
