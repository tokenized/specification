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

// NewHeaderForCode returns a new Header with the given code and size.
func NewHeaderForCode(code []byte, size uint64) (*Header, error) {
	h := Header{
		ProtocolID:       ProtocolID,
		OpPushDataLength: size,
		Version:          Version,
		ActionPrefix:     code,
	}

	return &h, nil
}

func WriteVarChar(buf *bytes.Buffer, value string, sizeBits int) error {
	if err := WriteVariableSize(buf, uint64(len(value)), sizeBits, 0); err != nil {
		return err
	}

	if _, err := buf.Write([]byte(value)); err != nil {
		return err
	}
	return nil
}

func ReadVarChar(buf *bytes.Buffer, sizeBits int) (string, error) {
	size, err := ReadVariableSize(buf, sizeBits, 0)
	if err != nil {
		return "", err
	}

	data := make([]byte, size)
	err = readLen(buf, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFixedChar(buf *bytes.Buffer, value string, size uint64) error {
	if uint64(len(value)) > size {
		return errors.New(fmt.Sprintf("FixedChar too long %d > %d", len(value), size))
	}
	if _, err := buf.Write([]byte(value)); err != nil {
		return err
	}

	// Pad with zeroes
	if uint64(len(value)) < size {
		padCount := size - uint64(len(value))
		empty := make([]byte, padCount)
		for i := uint64(0); i < padCount; i++ {
			empty[i] = 0
		}

		if _, err := buf.Write(empty); err != nil {
			return err
		}
	}
	return nil
}

func ReadFixedChar(buf *bytes.Buffer, size uint64) (string, error) {
	var err error
	data := make([]byte, size)
	err = readLen(buf, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteVarBin(buf *bytes.Buffer, value []byte, sizeBits int) error {
	if err := WriteVariableSize(buf, uint64(len(value)), sizeBits, 0); err != nil {
		return err
	}

	if _, err := buf.Write(value); err != nil {
		return err
	}
	return nil
}

func ReadVarBin(buf *bytes.Buffer, sizeBits int) ([]byte, error) {
	size, err := ReadVariableSize(buf, sizeBits, 0)
	if err != nil {
		return []byte{}, err
	}

	data := make([]byte, size)
	err = readLen(buf, data)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func WriteFixedBin(buf *bytes.Buffer, value []byte, size uint64) error {
	if uint64(len(value)) > size {
		return errors.New(fmt.Sprintf("FixedBin too long %d > %d", len(value), size))
	}
	if _, err := buf.Write(value); err != nil {
		return err
	}

	// Pad with zeroes
	if uint64(len(value)) < size {
		padCount := size - uint64(len(value))
		empty := make([]byte, padCount)
		for i := uint64(0); i < padCount; i++ {
			empty[i] = 0
		}

		if _, err := buf.Write(empty); err != nil {
			return err
		}
	}
	return nil
}

// WriteVariableSize writes a size/length to the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
func WriteVariableSize(buf *bytes.Buffer, size uint64, sizeBits int, defaultSizeBits int) error {
	if sizeBits == 0 {
		sizeBits = defaultSizeBits
	}
	if size >= 2<<uint64(sizeBits) {
		return errors.New(fmt.Sprintf("Size beyond size bits limit (%d) : %d", (2<<uint64(sizeBits))-1, size))
	}

	var err error
	switch sizeBits {
	case 8:
		err = write(buf, uint8(size))
	case 16:
		err = write(buf, uint16(size))
	case 32:
		err = write(buf, uint32(size))
	case 64:
		err = write(buf, uint64(size))
	default:
		return errors.New(fmt.Sprintf("Invalid variable size bits : %d", sizeBits))
	}
	if err != nil {
		return err
	}

	return nil
}

// ReadVariableSize reads a size/length from the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
func ReadVariableSize(buf *bytes.Buffer, sizeBits int, defaultSizeBits int) (uint64, error) {
	if sizeBits == 0 {
		sizeBits = defaultSizeBits
	}

	var err error
	var size uint64
	switch sizeBits {
	case 8:
		var size8 uint8
		err = read(buf, &size8)
		size = uint64(size8)
	case 16:
		var size16 uint16
		err = read(buf, &size16)
		size = uint64(size16)
	case 32:
		var size32 uint32
		err = read(buf, &size32)
		size = uint64(size32)
	case 64:
		err = read(buf, &size)
	default:
		err = errors.New(fmt.Sprintf("Invalid variable size bits : %d", sizeBits))
	}
	return size, err
}
