// Initiative : Initiative Action - Allows Token Owners to propose an
// Initiative (aka Initiative/Shareholder vote). A significant cost -
// specified in the Contract Formation - can be attached to this action to
// reduce spam, as the resulting vote will be put to all token owners.
package protocol

type Initiative struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	VoteSystem uint8
	Proposal bool
	ProposedChangesCount uint8
	ProposedChanges []Amendment
	VoteOptions Nvarchar8
	VoteMax uint8
	ProposalDescription Nvarchar16
	ProposalDocumentHash []byte
	VoteCutOffTimestamp uint64
}

// NewInitiative returns a new Initiative with defaults set.
func NewInitiative() Initiative {
	return Initiative{}
}

// Type returns the type identifer for this message.
func (m Initiative) Type() string {
	return CodeInitiative
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Initiative) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Initiative) Serialize() ([]byte, error) {
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

	if err := write(buf, m.VoteSystem); err != nil {
		return nil, err
	}

	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposedChangesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.); i++ {
		b, err := m.ProposedChanges[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.VoteOptions.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	{
		b, err := m.ProposalDescription.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ProposalDocumentHash, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteCutOffTimestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeInitiative, len(b))
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
func (m *Initiative) Write(b []byte) (int, error) {
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

	if err := read(buf, &m.VoteSystem); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ProposedChangesCount); err != nil {
		return 0, err
	}
for i := 0; i < int(m.); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.ProposedChanges = append(m.ProposedChanges, *x)
	}

	if err := m.VoteOptions.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	if err := m.ProposalDescription.Write(buf); err != nil {
		return 0, err
	}

	m.ProposalDocumentHash = make([]byte, 32)
	if err := readLen(buf, m.ProposalDocumentHash); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteCutOffTimestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Initiative) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Initiative) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", m.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", m.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", m.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#x", m.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%v", m.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
