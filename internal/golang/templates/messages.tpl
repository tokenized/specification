package {{ .Package }}

import (
	"bytes"
	"fmt"

	"github.com/tokenized/pkg/bsor"

	"github.com/pkg/errors"
)

type Message interface {
	Code() uint32

	Validate() error
	Equal(interface{}) bool

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
		if _, err := bsor.UnmarshalBinary(payload, result); err != nil {
			return nil, errors.Wrap(err, "unmarshal bsor")
		}
	}

	return result, nil
}

{{ range .Messages }}
func (a *{{.Name}}) Code() uint32 {
	return Code{{.Name}}
}

func (a *{{.Name}}) Bytes() ([]byte, error) {
	return bsor.MarshalBinary(a)
}

// Serialize writes an instrument to a byte slice.
func (a *{{.Name}}) Serialize(buf *bytes.Buffer) error {
	data, err := bsor.MarshalBinary(a)
	if err != nil {
		return errors.Wrap(err, "serialize {{.Name}}")
	}

	_, err = buf.Write(data)
	return err
}
{{ end }}
