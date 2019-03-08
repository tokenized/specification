// BallotCounted : Ballot Counted Action - The smart contract will respond
// to a Ballot Cast action with a Ballot Counted action if the Ballot Cast
// is valid. If the Ballot Cast is not valid, then the smart contract will
// respond with a Rejection Action.
package protocol

type BallotCounted struct {
	Header Header
	Timestamp uint64
}

// NewBallotCounted returns a new BallotCounted with defaults set.
func NewBallotCounted() BallotCounted {
	return BallotCounted{}
}

// Type returns the type identifer for this message.
func (m BallotCounted) Type() string {
	return CodeBallotCounted
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m BallotCounted) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m BallotCounted) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeBallotCounted, len(b))
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
func (m *BallotCounted) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m BallotCounted) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m BallotCounted) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
