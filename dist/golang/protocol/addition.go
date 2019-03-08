// Addition : Addition Action - Adds a User's public address to a global
// distributed whitelist. Entities (eg. Issuer) can filter by the public
// address of known and trusted entities (eg. KYC Databases such as
// coinbase) and therefore are able to create sublists - or subsets - of
// the main global whitelist.
package protocol

type Addition struct {
	Header Header
	TextEncoding uint8
	Sublist []byte
	KYC byte
	Jurisdiction []byte
	DOB uint64
	CountryofResidence []byte
	SupportingDocumentationHash []byte
	Message Nvarchar16
}

// NewAddition returns a new Addition with defaults set.
func NewAddition() Addition {
	return Addition{}
}

// Type returns the type identifer for this message.
func (m Addition) Type() string {
	return CodeAddition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Addition) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Addition) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Sublist, 4)); err != nil {
		return nil, err
	}

	if err := write(buf, m.KYC); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.DOB); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.CountryofResidence, 3)); err != nil {
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

	header, err := NewHeaderForCode(CodeAddition, len(b))
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
func (m *Addition) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	m.Sublist = make([]byte, 4)
	if err := readLen(buf, m.Sublist); err != nil {
		return 0, err
	}

	if err := read(buf, &m.KYC); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.DOB); err != nil {
		return 0, err
	}

	m.CountryofResidence = make([]byte, 3)
	if err := readLen(buf, m.CountryofResidence); err != nil {
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
func (m Addition) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Addition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("Sublist:%#x", m.Sublist))
	vals = append(vals, fmt.Sprintf("KYC:%#+v", m.KYC))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("DOB:%v", m.DOB))
	vals = append(vals, fmt.Sprintf("CountryofResidence:%#x", m.CountryofResidence))
	vals = append(vals, fmt.Sprintf("SupportingDocumentationHash:%#x", m.SupportingDocumentationHash))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
