package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bsor"

	"github.com/pkg/errors"
)

type Message interface {
	Code() uint16

	Validate() error
	Equal(interface{}) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error
}

const (
{{- range .Messages }}
	// Code{{.Name}} identifies a payload as a {{.Name}} message.
	Code{{.Name}} = uint16({{.Code}})
{{ end }}
)

// NewMessageFromCode returns a new object of the correct type for the code.
func NewMessageFromCode(code uint16) Message {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	default:
		return nil
	}
}

// DeserializeV1 reads a message from a byte slice.
func DeserializeV1(code uint16, payload []byte) (Message, error) {
	result := NewMessageFromCode(code)
	if result == nil {
		return nil, fmt.Errorf("Unknown message code : %d", code)
	}

	if len(payload) != 0 {
		if _, err := bsor.UnmarshalBinary(payload, result); err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
	}

	return result, nil
}

{{ range .Messages }}
func (a *{{.Name}}) Code() uint16 {
	return Code{{.Name}}
}

func (a *{{.Name}}) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *{{.Name}}) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize {{.Name}}")
	}

	_, err = buf.Write(data)
	return err
}
{{ end }}
