package {{ .Package }}

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Asset interface {
	proto.Message

	Code() string

	Validate() error
	Equal(proto.Message) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error

	ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error)
}

const (
{{- range .Messages }}
	// Code{{.Name}} identifies a payload as a {{.Name}} asset message.
	Code{{.Name}} = "{{.Code}}"
{{ end }}
)

// NewAssetFromCode returns a new object of the correct type for the code.
func NewAssetFromCode(code string) Asset {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	default:
		return nil
	}
}

// Deserialize reads an asset from a byte slice.
func Deserialize(code []byte, payload []byte) (Asset, error) {
	result := NewAssetFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown asset code : %s", string(code))
	}

	if len(payload) != 0 {
		if err := proto.Unmarshal(payload, result); err != nil {
			return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
		}
	}

	return result, nil
}

{{ range $i, $message := .Messages }}
func (a *{{ .Name }}) Code() string {
	return Code{{ .Name }}
}

func (a *{{ .Name }}) Bytes() ([]byte, error) {
	return proto.Marshal(a)
}

// Serialize writes an asset to a byte slice.
func (a *{{ .Name }}) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize {{ .Name }}")
	}

	_, err = buf.Write(data)
	return err
}

{{ end }}