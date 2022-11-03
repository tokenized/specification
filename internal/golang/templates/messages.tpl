package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

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

	if len(payload) == 0 {
		return result, nil // empty result
	}

	if len(payload) == 1 && payload[0] == bitcoin.OP_FALSE {
		return result, nil // empty result
	}

	if err := proto.Unmarshal(payload, result); err != nil {
		return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
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

// Serialize writes an instrument to a byte slice.
func (a *{{.Name}}) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize {{.Name}}")
	}

	_, err = buf.Write(data)
	return err
}
{{ end }}
