package {{ .Package }}

import (
	"bytes"
	"fmt"

	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Instrument interface {
	proto.Message

	Code() string

	Validate() error
	Equal(proto.Message) bool

	Bytes() ([]byte, error)
	Serialize(buf *bytes.Buffer) error

	ApplyAmendment(fip permissions.FieldIndexPath, operation uint32, data []byte,
		permissions permissions.Permissions) (permissions.Permissions, error)
}

const (
{{- range .Messages }}
	// Code{{.Name}} identifies a payload as a {{.Name}} instrument message.
	Code{{.Name}} = "{{.Code}}"
{{ end }}

	// BondTypeCorporate specifies a corporate bond.
	BondTypeCorporate = "C"

	// BondTypeMunicipal specifies a municipal bond.
	BondTypeMunicipal = "M"

	// BondTypeGovernment specifies a government bond.
	BondTypeGovernment = "G"

	// CasinoChipUseTypeRealMoney specifies a casino chip is real money.
	CasinoChipUseTypeRealMoney = "R"

	// CasinoChipUseTypeSocial specifies a casino chip is social.
	CasinoChipUseTypeSocial = "S"

	// CasinoChipUseTypeFreePlay specifies a casino chip is free play.
	CasinoChipUseTypeFreePlay = "F"
)

// NewInstrumentFromCode returns a new object of the correct type for the code.
func NewInstrumentFromCode(code string) Instrument {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	default:
		return nil
	}
}

// Deserialize reads an instrument from a byte slice.
func Deserialize(code []byte, payload []byte) (Instrument, error) {
	result := NewInstrumentFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown instrument code : %s", string(code))
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

// Serialize writes an instrument to a byte slice.
func (a *{{ .Name }}) Serialize(buf *bytes.Buffer) error {
	data, err := proto.Marshal(a)
	if err != nil {
		return errors.Wrap(err, "Failed to serialize {{ .Name }}")
	}

	_, err = buf.Write(data)
	return err
}

{{ end }}
