package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

const (
	// Version of the Tokenized protocol.
	Version = uint8(0)

	// OpReturn (OP_RETURN) is a script opcode is used to mark a transaction
	// output as invalid, and can be used to add data to a TX.
	OpReturn = 0x6a

	OP_PUSHDATA1 = byte(0x4c)
	OP_PUSHDATA2 = byte(0x4d)
	OP_PUSHDATA4 = byte(0x4e)

	OP_PUSHDATA1_MAX = 255
	OP_PUSHDATA2_MAX = 65535
)

// PayloadMessage is the interface for messages that are derived from
// payloads, such as asset types.
type PayloadMessage interface {
	io.Writer
	Type() string
	Serialize() ([]byte, error)
}

// NewPayloadMessageFromCode returns the approriate PayloadMessage for the
// given code.
func NewPayloadMessageFromCode(code []byte) (PayloadMessage, error) {
	s := string(code)

	switch s {
	case CodeShareCommon:
		return NewShareCommon(), nil
	}

	return nil, fmt.Errorf("No asset type for code %s", code)
}

// OpReturnMessage implements a base interface for all message types.
type OpReturnMessage interface {
	PayloadMessage
	String() string
	PayloadMessage() (PayloadMessage, error)
}

// New returns a new message, as an OpReturnMessage, from the OP_RETURN
// payload.
func New(b []byte) (OpReturnMessage, error) {
	code, err := Code(b)
	if err != nil {
		return nil, err
	}

	t, ok := TypeMapping[code]
	if !ok {
		return nil, fmt.Errorf("Unknown code :  %v", code)
	}

	if _, err := t.Write(b); err != nil {
		return nil, err
	}

	return t, nil
}

// Code returns the identifying code from the OP_RETURN payload.
func Code(b []byte) (string, error) {
	if len(b) < 9 || b[0] != OpReturn {
		return "", errors.New("Not an OP_RETURN payload")
	}

	offset := 7

	if b[1] < 0x4c {
		offset = 6
	}

	return string(b[offset : offset+2]), nil
}

func NewHeaderForCode(code string, size int) (*Header, error) {
	// work out which opcode to use depending on size of the data.
	opcode := OP_PUSHDATA1

	if size > OP_PUSHDATA2_MAX {
		opcode = OP_PUSHDATA4
	} else if size > OP_PUSHDATA1_MAX {
		opcode = OP_PUSHDATA2
	}

	lenPayload, err := uintToBytes(uint64(size))
	if err != nil {
		return nil, err
	}

	h := Header{
		ProtocolID:       ProtocolID,
		OpPushdata:       opcode,
		LenActionPayload: lenPayload,
		Version:          Version,
		ActionPrefix:     []byte(code),
	}

	return &h, nil
}

// nchar and nvarchar can store utf-8 characters.
// char and varchar store ascii data
//
// 1st byte specifies length
// 2nd byte onwards is the data, length of which was specified in the first
// byte. If 1st byte was 0, there would be no data at all.

// Nvarchar is subject to the encoding specified in the message

// this could be a "factory". is it needed? return correct type based on size.

type Nvarchar interface {
	String() string
	Write(*bytes.Buffer) error
	Serialize() ([]byte, error)
}

func NewNvarchar(b []byte) Nvarchar {
	n := len(b)

	if n <= 0xff {
		return NewNvarchar8(b)
	} else if n <= 0xffff {
		return NewNvarchar16(b)
	} else if n <= 0xffffffff {
		return NewNvarchar32(b)
	}

	return NewNvarchar64(b)
}

type Nvarchar8 struct {
	Len  uint8
	Data []byte
}

func NewNvarchar8(b []byte) *Nvarchar8 {
	return &Nvarchar8{
		Len:  uint8(len(b)),
		Data: b,
	}
}

// Write writes the contents of the io.Writer to the struct.
func (t *Nvarchar8) Write(buf *bytes.Buffer) error {
	if len(t.Data) > 0xff {
		return errors.New("Data exceeds limit of type")
	}

	// write uint8
	if err := read(buf, &t.Len); err != nil {
		return err
	}

	// write data
	t.Data = make([]byte, t.Len, t.Len)
	if err := readLen(buf, t.Data); err != nil {
		return err
	}

	return nil
}

func (t Nvarchar8) Size() int {
	return 1 + len(t.Data)
}

func (t Nvarchar8) Serialize() ([]byte, error) {
	if len(t.Data) > 0xff {
		return nil, errors.New("Data exceeds limit of type")
	}

	buf := new(bytes.Buffer)

	if err := write(buf, t.Len); err != nil {
		return nil, err
	}

	if err := write(buf, t.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t Nvarchar8) String() string {
	return string(t.Data)
}

type Nvarchar16 struct {
	Len  uint16
	Data []byte
}

func NewNvarchar16(b []byte) *Nvarchar16 {
	return &Nvarchar16{
		Len:  uint16(len(b)),
		Data: b,
	}
}

// Write writes the contents of the io.Writer to the struct.
func (t *Nvarchar16) Write(buf *bytes.Buffer) error {
	if len(t.Data) > 0xffff {
		return errors.New("Data exceeds limit of type")
	}

	// write uint16
	if err := read(buf, &t.Len); err != nil {
		return err
	}

	// write data
	t.Data = make([]byte, t.Len, t.Len)
	if err := readLen(buf, t.Data); err != nil {
		return err
	}

	return nil
}

func (t Nvarchar16) Size() int {
	return 2 + len(t.Data)
}

func (t Nvarchar16) Serialize() ([]byte, error) {
	if len(t.Data) > 0xffff {
		return nil, errors.New("Data exceeds limit of type")
	}

	buf := new(bytes.Buffer)

	if err := write(buf, t.Len); err != nil {
		return nil, err
	}

	if err := write(buf, t.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t Nvarchar16) String() string {
	return string(t.Data)
}

type Nvarchar32 struct {
	Len  uint32
	Data []byte
}

func NewNvarchar32(b []byte) *Nvarchar32 {
	return &Nvarchar32{
		Len:  uint32(len(b)),
		Data: b,
	}
}

// Write writes the contents of the io.Writer to the struct.
func (t *Nvarchar32) Write(buf *bytes.Buffer) error {
	if len(t.Data) > 0xffffffff {
		return errors.New("Data exceeds limit of type")
	}

	// write uint32
	if err := read(buf, &t.Len); err != nil {
		return err
	}

	// write data
	t.Data = make([]byte, t.Len, t.Len)
	if err := readLen(buf, t.Data); err != nil {
		return err
	}

	return nil
}

func (t Nvarchar32) Size() int {
	return 4 + len(t.Data)
}

func (t Nvarchar32) Serialize() ([]byte, error) {
	if len(t.Data) > 0xffffffff {
		return nil, errors.New("Data exceeds limit of type")
	}

	buf := new(bytes.Buffer)

	if err := write(buf, t.Len); err != nil {
		return nil, err
	}

	if err := write(buf, t.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t Nvarchar32) String() string {
	return string(t.Data)
}

type Nvarchar64 struct {
	Len  uint64
	Data []byte
}

func NewNvarchar64(b []byte) *Nvarchar64 {
	return &Nvarchar64{
		Len:  uint64(len(b)),
		Data: b,
	}
}

// Write writes the contents of the io.Writer to the struct.
func (t *Nvarchar64) Write(buf *bytes.Buffer) error {
	if len(t.Data) > 0xffffffff {
		return errors.New("Data exceeds limit of type")
	}

	// write uint64
	if err := read(buf, &t.Len); err != nil {
		return err
	}

	// write data
	t.Data = make([]byte, t.Len, t.Len)
	if err := readLen(buf, t.Data); err != nil {
		return err
	}

	return nil
}

func (t Nvarchar64) Size() int {
	return 8 + len(t.Data)
}

func (t Nvarchar64) Serialize() ([]byte, error) {
	if len(t.Data) > 0xffffffff {
		return nil, errors.New("Data exceeds limit of type")
	}

	buf := new(bytes.Buffer)

	if err := write(buf, t.Len); err != nil {
		return nil, err
	}

	if err := write(buf, t.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t Nvarchar64) String() string {
	return string(t.Data)
}
