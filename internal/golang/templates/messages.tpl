package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Message interface {
	proto.Message

	Code() uint32

	Validate() error
	Equal(proto.Message) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error
}

const (
{{- range .Messages }}
	// Code{{.Name}} identifies a payload as a {{.Name}} message.
	Code{{.Name}} = uint32({{.Code}})
{{ end }}
)

// NewMessageFromCode returns a new object of the correct type for the code.
func NewMessageFromCode(code uint32) Message {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	default:
		return nil
	}
}

// Deserialize reads a message from a byte slice.
func Deserialize(code uint32, payload []byte) (Message, error) {
	result := NewMessageFromCode(code)
	if result == nil {
		return nil, fmt.Errorf("Unknown message code : %d", code)
	}

	if len(payload) != 0 {
		if err := proto.Unmarshal(payload, result); err != nil {
			return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
		}
	}

	return result, nil
}

{{ range .Messages }}
func (a *{{.Name}}) Code() uint32 {
	return Code{{.Name}}
}

func (a *{{.Name}}) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
func (a *{{.Name}}) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize {{.Name}}")
	}

	_, err = buf.Write(data)
	return err
}
{{ end }}

// Helper functions for amendments

// ReadBase128VarInt reads a base 128 variable encoded integer from the reader.
func ReadBase128VarInt(r io.ByteReader) (int, error) {
	value := uint32(0)
	done := false
	bitOffset := uint32(0)
	for !done {
		subValue, err := r.ReadByte()
		if err != nil {
			return int(value), err
		}

		done = (subValue & 0x80) == 0 // High bit not set
		subValue = subValue & 0x7f    // Remove high bit

		value += uint32(subValue) << bitOffset
		bitOffset += 7
	}

	return int(value), nil
}

// WriteBase128VarInt writes a base 128 variable encoded integer to the writer.
func WriteBase128VarInt(w io.ByteWriter, value int) error {
	v := uint32(value)
	for {
		if v < 128 {
			return w.WriteByte(byte(v))
		}
		subValue := (byte(v&0x7f) | 0x80) // Get last 7 bits and set high bit
		if err := w.WriteByte(subValue); err != nil {
			return err
		}
		v = v >> 7
	}
}