// Rejection : Rejection Action - used to reject Exchange, Send,
// Initiative, Referendum, Order, and Ballot Cast actions that do not
// comply with the Contract. If money is to be returned to a User then it
// is used in lieu of the Settlement Action to properly account for token
// balances. All Issuer/User Actions must be responded to by the Contract
// with an Action. The only exception to this rule is when there is not
// enough fees in the first Action for the Contract response action to
// remain revenue neutral. If not enough fees are attached to pay for the
// Contract response then the Contract will not respond. For example: Send
// and Exchange Actions must be responded to by the Contract with either a
// Settlement Action or a Rejection Action.
package protocol

type Rejection struct {
	Header Header
	TextEncoding uint8
	QtyReceivingAddresses uint8
	AddressIndexes []uint16
	RejectionType uint8
	MessagePayload Nvarchar16
	Timestamp uint64
}

// NewRejection returns a new Rejection with defaults set.
func NewRejection() Rejection {
	return Rejection{}
}

// Type returns the type identifer for this message.
func (m Rejection) Type() string {
	return CodeRejection
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Rejection) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Rejection) Serialize() ([]byte, error) {
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

	if err := write(buf, m.RejectionType); err != nil {
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

	header, err := NewHeaderForCode(CodeRejection, len(b))
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
func (m *Rejection) Write(b []byte) (int, error) {
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

	if err := read(buf, &m.RejectionType); err != nil {
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
func (m Rejection) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Rejection) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("QtyReceivingAddresses:%v", m.QtyReceivingAddresses))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", m.AddressIndexes))
	vals = append(vals, fmt.Sprintf("RejectionType:%v", m.RejectionType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#+v", m.MessagePayload))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
