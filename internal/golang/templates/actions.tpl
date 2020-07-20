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

	// OracleTypeIdentity identifies an identity oracle
	OracleTypeIdentity = uint32(0)

	// OracleTypeAuthority identifies an authority oracle
	OracleTypeAuthority = uint32(1)

	// OracleTypeEvent identifies an event oracle
	OracleTypeEvent = uint32(2)

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = "F"

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = "T"

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = "C"

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = "R"

	// ContractTypeEntity identifies an entity contract
	ContractTypeEntity = uint32(0)

	// ContractTypeAsset identifies an asset contract
	ContractTypeAsset = uint32(1)
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
