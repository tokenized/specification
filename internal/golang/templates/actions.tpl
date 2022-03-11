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

	// CodeAssetDefinition (Deprecated) identifies a payload as a InstrumentDefinition action message.
	CodeAssetDefinition = "A1"

	// CodeAssetCreation (Deprecated) identifies a payload as a InstrumentCreation action message.
	CodeAssetCreation = "A2"

	// CodeAssetModification (Deprecated) identifies a payload as a InstrumentModification action message.
	CodeAssetModification = "A3"

	// ServiceTypeIdentityOracle identifies an identity oracle
	ServiceTypeIdentityOracle = uint32(0)

	// ServiceTypeAuthorityOracle identifies an authority oracle
	ServiceTypeAuthorityOracle = uint32(1)

	// ServiceTypeEventOracle identifies an event oracle
	ServiceTypeEventOracle = uint32(2)

	// ServiceTypeContractOperator identifies a contract operator service
	ServiceTypeContractOperator = uint32(3)

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

	// ContractTypeInstrument identifies an instrument contract
	ContractTypeInstrument = uint32(1)

	// ContractBodyOfAgreementTypeNone identifies a contract with no body of agreement specified.
	ContractBodyOfAgreementTypeNone = uint32(0)

	// ContractBodyOfAgreementTypeHash identifies a contract with the hash of the body of agreement
	// document included.
	ContractBodyOfAgreementTypeHash = uint32(1)

	// ContractBodyOfAgreementTypeFull identifies a contract with the a full body of agreement
	// specified in a BodyOfAgreementFormation action.
	ContractBodyOfAgreementTypeFull = uint32(2)
)

// NewActionFromCode returns a new object of the correct type for the code.
func NewActionFromCode(code string) Action {
	switch code {
{{- range .Messages }}
	case Code{{.Name}}:
		return &{{.Name}}{}
{{- end }}
	case CodeAssetDefinition: // Deprecated, but backwards compatible
		return &InstrumentDefinition{}
	case CodeAssetCreation: // Deprecated, but backwards compatible
		return &InstrumentCreation{}
	case CodeAssetModification: // Deprecated, but backwards compatible
		return &InstrumentModification{}
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
