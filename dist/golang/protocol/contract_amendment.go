// ContractAmendment : Contract Amendment Action - the issuer can initiate
// an amendment to the contract establishment metadata. The ability to make
// an amendment to the contract is restricted by the Authorization Flag set
// on the current revision of Contract Formation action.
package protocol

type ContractAmendment struct {
	Header Header
	TextEncoding uint8
	ChangeIssuerAddress bool
	ChangeOperatorAddress bool
	ContractRevision uint16
	AmendmentsCount uint8
	Amendments []Amendment
	RefTxID []byte
}

// NewContractAmendment returns a new ContractAmendment with defaults set.
func NewContractAmendment() ContractAmendment {
	return ContractAmendment{}
}

// Type returns the type identifer for this message.
func (m ContractAmendment) Type() string {
	return CodeContractAmendment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ContractAmendment) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ContractAmendment) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, m.ChangeIssuerAddress); err != nil {
		return nil, err
	}

	if err := write(buf, m.ChangeOperatorAddress); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.AmendmentsCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.Amendments[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.RefTxID, 32)); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeContractAmendment, len(b))
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
func (m *ContractAmendment) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ChangeIssuerAddress); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ChangeOperatorAddress); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AmendmentsCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Amendments = append(m.Amendments, *x)
	}

	m.RefTxID = make([]byte, 32)
	if err := readLen(buf, m.RefTxID); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m ContractAmendment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m ContractAmendment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("ChangeIssuerAddress:%#+v", m.ChangeIssuerAddress))
	vals = append(vals, fmt.Sprintf("ChangeOperatorAddress:%#+v", m.ChangeOperatorAddress))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("AmendmentsCount:%v", m.AmendmentsCount))
	vals = append(vals, fmt.Sprintf("Amendments:%#+v", m.Amendments))
	vals = append(vals, fmt.Sprintf("RefTxID:%#x", m.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
