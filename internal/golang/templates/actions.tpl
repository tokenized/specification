package {{ .Package }}

import (
	"fmt"

	"github.com/tokenized/pkg/bsor"

	"github.com/pkg/errors"
)

type Action interface {
	Code() string

	Validate() error
	Equal(interface{}) bool
}

type Field interface {}

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
	ServiceTypeIdentityOracle = uint8(0)

	// ServiceTypeAuthorityOracle identifies an authority oracle
	ServiceTypeAuthorityOracle = uint8(1)

	// ServiceTypeEventOracle identifies an event oracle
	ServiceTypeEventOracle = uint8(2)

	// ServiceTypeContractOperator identifies a contract operator service
	ServiceTypeContractOperator = uint8(3)

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = "F"

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = "T"

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = "C"

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = "R"

	// ContractTypeEntity identifies an entity contract
	ContractTypeEntity = uint8(0)

	// ContractTypeInstrument identifies an instrument contract
	ContractTypeInstrument = uint8(1)

	// ContractBodyOfAgreementTypeNone identifies a contract with no body of agreement specified.
	ContractBodyOfAgreementTypeNone = uint8(0)

	// ContractBodyOfAgreementTypeHash identifies a contract with the hash of the body of agreement
	// document included.
	ContractBodyOfAgreementTypeHash = uint8(1)

	// ContractBodyOfAgreementTypeFull identifies a contract with the a full body of agreement
	// specified in a BodyOfAgreementFormation action.
	ContractBodyOfAgreementTypeFull = uint8(2)
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

// DeserializeV1 reads an action from a byte slice.
func DeserializeV1(payload bitcoin.ScriptItems) (Action, error) {
	if len(payload) < 1 {
		return nil, errors.New("No action code")
	}

	result := NewActionFromCode(string([]byte(payload[0].Data)))
	if result == nil {
		return nil, fmt.Errorf("Unknown action code : %s", string([]byte(payload[0].Data)))
	}

	if len(payload) > 1 {
		if _, err := bsor.Unmarshal(payload[1:], result); err != nil {
			return nil, errors.Wrap(err, "unmarshal")
		}
	}

	switch res := result.(type) {
	case *InstrumentDefinition:
		res.InstrumentPayloadVersion = 1
	case *InstrumentModification:
		res.InstrumentPayloadVersion = 1
	case *InstrumentCreation:
		res.InstrumentPayloadVersion = 1
	case *Message:
		res.MessagePayloadVersion = 1
	}

	return result, nil
}

{{ range .Messages }}
func (a *{{.Name}}) Code() string {
	return Code{{.Name}}
}
{{ end }}
