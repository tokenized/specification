// Freeze : Freeze Action - To be used to comply with
// contractual/legal/issuer requirements. The target public address(es)
// will be marked as frozen. However the Freeze action publishes this fact
// to the public blockchain for transparency. The Contract will not respond
// to any actions requested by the frozen address.
package protocol

type Freeze struct {
	Header Header
	AddressCount uint16
	Addresses []Address
	Timestamp uint64
}

// NewFreeze returns a new Freeze with defaults set.
func NewFreeze() Freeze {
	return Freeze{}
}

// Type returns the type identifer for this message.
func (m Freeze) Type() string {
	return CodeFreeze
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Freeze) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Freeze) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Addresses[i].Serialize()
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

	header, err := NewHeaderForCode(CodeFreeze, len(b))
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
func (m *Freeze) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AddressCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Address{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Addresses = append(m.Addresses, *x)
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Freeze) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Freeze) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
