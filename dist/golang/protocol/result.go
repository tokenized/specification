// Result : Result Action - Once a vote has been completed the results are
// published.
package protocol

type Result struct {
	Header Header
	TextEncoding uint8
	AssetType []byte
	AssetID []byte
	Proposal bool
	ProposedChangesCount uint8
	ProposedChanges []Amendment
	VoteTxnID []byte
	VoteOptionsCount uint8
	Option1Tally uint64
	Result Nvarchar8
	Timestamp uint64
}

// NewResult returns a new Result with defaults set.
func NewResult() Result {
	return Result{}
}

// Type returns the type identifer for this message.
func (m Result) Type() string {
	return CodeResult
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Result) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Result) Serialize() ([]byte, error) {
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

	if err := write(buf, pad(m.VoteTxnID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteOptionsCount); err != nil {
		return nil, err
	}

	if err := write(buf, m.Option1Tally); err != nil {
		return nil, err
	}

	{
		b, err := m.Result.Serialize()
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

	header, err := NewHeaderForCode(CodeResult, len(b))
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
func (m *Result) Write(b []byte) (int, error) {
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

	m.VoteTxnID = make([]byte, 32)
	if err := readLen(buf, m.VoteTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteOptionsCount); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Option1Tally); err != nil {
		return 0, err
	}

	if err := m.Result.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Result) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Result) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteTxnID:%#x", m.VoteTxnID))
	vals = append(vals, fmt.Sprintf("VoteOptionsCount:%v", m.VoteOptionsCount))
	vals = append(vals, fmt.Sprintf("Option1Tally:%v", m.Option1Tally))
	vals = append(vals, fmt.Sprintf("Result:%#+v", m.Result))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
