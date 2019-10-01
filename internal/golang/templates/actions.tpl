package {{ .Package }}

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Action interface {
	proto.Message

	Code() string

	Validate() error
	Equal(proto.Message) bool
}

type Field interface {
	proto.Message
}

const (
{{- range .Messages }}
	// Code{{.Name}} identifies a payload as a {{.Name}} action message.
	Code{{.Name}} = "{{.Code}}"
{{ end }}

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = "F"

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = "T"

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = "C"

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = "R"
)

// NewActionFromCode returns a new object of the correct type for the code.
func NewActionFromCode(code string) Action {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	default:
		return nil
	}
}

// Deserialize reads an action from a byte slice.
func Deserialize(code []byte, payload []byte) (Action, error) {
	result := NewActionFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown action code : %s", string(code))
	}

	if len(payload) != 0 {
		if err := proto.Unmarshal(payload, result); err != nil {
			return nil, errors.Wrap(err, "Failed protobuf unmarshaling")
		}
	}

	return result, nil
}

{{ range .Messages }}
func (a *{{.Name}}) Code() string {
	return Code{{.Name}}
}
{{ end }}

{{- range .Messages }}
	{{- if eq .Name "ContractOffer" }}
// Contract Permission / Amendment Field Indices
const (
	{{- range $i, $field := .Fields }}
	ContractField{{ $field.Name }} = uint32({{add $i 1}})
	{{- end }}
)
	{{- end }}
{{ end }}

{{- range .Messages }}
	{{- if eq .Name "AssetDefinition" }}
// Asset Permission / Amendment Field Indices
const (
	{{- range $i, $field := .Fields }}
	AssetField{{ $field.Name }} = uint32({{add $i 1}})
	{{- end }}
)
	{{- end }}
{{ end }}

{{- range $i, $fieldType := .FieldTypes }}
	{{- if eq $fieldType.Name "Entity" "Oracle" }}
// {{ $fieldType.Name }} Permission / Amendment Field Indices
const (
	{{- range $offset, $field := .Fields }}
	{{ $fieldType.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
	{{- end }}
)
	{{- end }}
{{ end }}
