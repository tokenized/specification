// Reconciliation : Reconciliation Action - to be used at the direction of
// the issuer to fix record keeping errors with bitcoin and token balances.
package protocol

type Reconciliation struct {
	Header Header
	AddressCount uint16
	Addresses []Address
	Timestamp uint64
}

// NewReconciliation returns a new Reconciliation with defaults set.
func NewReconciliation() Reconciliation {
	return Reconciliation{}
}

// Type returns the type identifer for this message.
func (m Reconciliation) Type() string {
	return CodeReconciliation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Reconciliation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Reconciliation) Serialize() ([]byte, error) {
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

	header, err := NewHeaderForCode(CodeReconciliation, len(b))
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
func (m *Reconciliation) Write(b []byte) (int, error) {
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
func (m Reconciliation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Reconciliation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
