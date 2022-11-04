package {{ .Package }}

import (
	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	binaryVersion = uint8(0)
)

{{ range .Messages }}
func (a *{{ .Name }}) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *{{ .Name }}) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
{{ end }}

{{ range .FieldTypes }}
func (a *{{ .Name }}Field) MarshalBinary() (data []byte, err error) {
	b, err := proto.Marshal(a)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	return append([]byte{binaryVersion}, b...), nil
}

func (a *{{ .Name }}Field) UnmarshalBinary(data []byte) error {
	if len(data) == 0 {
		return nil // empty result
	}

	if len(data) == 1 && data[0] == bitcoin.OP_FALSE {
		return nil // empty result
	}

	if len(data) < 2 {
		return errors.New("Data too small, missing version or data")
	}

	if data[0] != binaryVersion {
		return fmt.Errorf("Wrong data version: got %d, want %d", data[0], binaryVersion)
	}

	if err := proto.Unmarshal(data[1:], a); err != nil {
		return errors.Wrap(err, "protobuf unmarshal")
	}

	return nil
}
{{ end }}
