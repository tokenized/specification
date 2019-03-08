// Establishment : Establishment Action - Establishes a register. The
// register is intended to be used primarily for whitelisting. However,
// other types of registers can be used.
package protocol

type Establishment struct {
	Header Header
	TextEncoding uint8
	Registrar Nvarchar8
	RegisterType byte
	Jurisdiction []byte
	SupportingDocumentationHash []byte
	Message Nvarchar16
}

// NewEstablishment returns a new Establishment with defaults set.
func NewEstablishment() Establishment {
	return Establishment{}
}

// Type returns the type identifer for this message.
func (m Establishment) Type() string {
	return CodeEstablishment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Establishment) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Establishment) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	{
		b, err := m.Registrar.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.RegisterType); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.SupportingDocumentationHash, 32)); err != nil {
		return nil, err
	}

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeEstablishment, len(b))
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
func (m *Establishment) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := m.Registrar.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RegisterType); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	m.SupportingDocumentationHash = make([]byte, 32)
	if err := readLen(buf, m.SupportingDocumentationHash); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Establishment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Establishment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("Registrar:%#+v", m.Registrar))
	vals = append(vals, fmt.Sprintf("RegisterType:%#+v", m.RegisterType))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("SupportingDocumentationHash:%#x", m.SupportingDocumentationHash))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
