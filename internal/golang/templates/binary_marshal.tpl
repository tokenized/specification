package {{ .Package }}

import (
	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

{{ range .Messages }}
func (a *{{ .Name }}) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *{{ .Name }}) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
{{ end }}

{{ range .FieldTypes }}
func (a *{{ .Name }}Field) MarshalBinary() (data []byte, err error) {
	return proto.Marshal(a)
}

func (a *{{ .Name }}Field) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if err := proto.Unmarshal(data, a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
{{ end }}
