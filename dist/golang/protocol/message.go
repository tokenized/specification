// Message : Message Action - the message action is a general purpose
// communication action. 'Twitter/sms' for Issuers/Investors/Users. The
// message txn can also be used for passing partially signed txns on-chain,
// establishing private communication channels including receipting,
// invoices, PO, and private offers/bids. The messages are broken down by
// type for easy filtering in the a user’s wallet. The Message Types are
// listed in the Message Types table.
package protocol

type Message struct {
	Header Header
	TextEncoding uint8
	QtyReceivingAddresses uint8
	AddressIndexes []uint16
	MessageType []byte
	MessagePayload Nvarchar16
	Timestamp uint64
}

// NewMessage returns a new Message with defaults set.
func NewMessage() Message {
	return Message{}
}

// Type returns the type identifer for this message.
func (m Message) Type() string {
	return CodeMessage
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Message) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Message) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, m.QtyReceivingAddresses); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		if err := write(buf, m.AddressIndexes[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.MessageType, 2)); err != nil {
		return nil, err
	}

	{
		b, err := m.MessagePayload.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeMessage, len(b))
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
func (m *Message) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := read(buf, &m.QtyReceivingAddresses); err != nil {
		return 0, err
	}

	m.AddressIndexes = make([]uint16, m., m.)

	if err := read(buf, &m.AddressIndexes); err != nil {
		return 0, err
	}

	m.MessageType = make([]byte, 2)
	if err := readLen(buf, m.MessageType); err != nil {
		return 0, err
	}

	if err := m.MessagePayload.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Message) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Message) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("QtyReceivingAddresses:%v", m.QtyReceivingAddresses))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", m.AddressIndexes))
	vals = append(vals, fmt.Sprintf("MessageType:%#x", m.MessageType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#+v", m.MessagePayload))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
