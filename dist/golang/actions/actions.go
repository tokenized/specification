package actions

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/bsor"

	"github.com/pkg/errors"
)

type Action interface {
	Code() string

	Validate() error
	Equal(interface{}) bool
}

type Field interface{}

const (
	// CodeContractOffer identifies a payload as a ContractOffer action message.
	CodeContractOffer = "C1"

	// CodeContractFormation identifies a payload as a ContractFormation action message.
	CodeContractFormation = "C2"

	// CodeContractAmendment identifies a payload as a ContractAmendment action message.
	CodeContractAmendment = "C3"

	// CodeStaticContractFormation identifies a payload as a StaticContractFormation action message.
	CodeStaticContractFormation = "C4"

	// CodeContractAddressChange identifies a payload as a ContractAddressChange action message.
	CodeContractAddressChange = "C5"

	// CodeBodyOfAgreementOffer identifies a payload as a BodyOfAgreementOffer action message.
	CodeBodyOfAgreementOffer = "C6"

	// CodeBodyOfAgreementFormation identifies a payload as a BodyOfAgreementFormation action message.
	CodeBodyOfAgreementFormation = "C7"

	// CodeBodyOfAgreementAmendment identifies a payload as a BodyOfAgreementAmendment action message.
	CodeBodyOfAgreementAmendment = "C8"

	// CodeInstrumentDefinition identifies a payload as a InstrumentDefinition action message.
	CodeInstrumentDefinition = "I1"

	// CodeInstrumentCreation identifies a payload as a InstrumentCreation action message.
	CodeInstrumentCreation = "I2"

	// CodeInstrumentModification identifies a payload as a InstrumentModification action message.
	CodeInstrumentModification = "I3"

	// CodeTransfer identifies a payload as a Transfer action message.
	CodeTransfer = "T1"

	// CodeSettlement identifies a payload as a Settlement action message.
	CodeSettlement = "T2"

	// CodeProposal identifies a payload as a Proposal action message.
	CodeProposal = "G1"

	// CodeVote identifies a payload as a Vote action message.
	CodeVote = "G2"

	// CodeBallotCast identifies a payload as a BallotCast action message.
	CodeBallotCast = "G3"

	// CodeBallotCounted identifies a payload as a BallotCounted action message.
	CodeBallotCounted = "G4"

	// CodeResult identifies a payload as a Result action message.
	CodeResult = "G5"

	// CodeOrder identifies a payload as a Order action message.
	CodeOrder = "E1"

	// CodeFreeze identifies a payload as a Freeze action message.
	CodeFreeze = "E2"

	// CodeThaw identifies a payload as a Thaw action message.
	CodeThaw = "E3"

	// CodeConfiscation identifies a payload as a Confiscation action message.
	CodeConfiscation = "E4"

	// CodeReconciliation identifies a payload as a Reconciliation action message.
	CodeReconciliation = "E5"

	// CodeEstablishment identifies a payload as a Establishment action message.
	CodeEstablishment = "R1"

	// CodeAddition identifies a payload as a Addition action message.
	CodeAddition = "R2"

	// CodeAlteration identifies a payload as a Alteration action message.
	CodeAlteration = "R3"

	// CodeRemoval identifies a payload as a Removal action message.
	CodeRemoval = "R4"

	// CodeMessage identifies a payload as a Message action message.
	CodeMessage = "M1"

	// CodeRejection identifies a payload as a Rejection action message.
	CodeRejection = "M2"

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
	case CodeContractOffer:
		return &ContractOffer{}
	case CodeContractFormation:
		return &ContractFormation{}
	case CodeContractAmendment:
		return &ContractAmendment{}
	case CodeStaticContractFormation:
		return &StaticContractFormation{}
	case CodeContractAddressChange:
		return &ContractAddressChange{}
	case CodeBodyOfAgreementOffer:
		return &BodyOfAgreementOffer{}
	case CodeBodyOfAgreementFormation:
		return &BodyOfAgreementFormation{}
	case CodeBodyOfAgreementAmendment:
		return &BodyOfAgreementAmendment{}
	case CodeInstrumentDefinition:
		return &InstrumentDefinition{}
	case CodeInstrumentCreation:
		return &InstrumentCreation{}
	case CodeInstrumentModification:
		return &InstrumentModification{}
	case CodeTransfer:
		return &Transfer{}
	case CodeSettlement:
		return &Settlement{}
	case CodeProposal:
		return &Proposal{}
	case CodeVote:
		return &Vote{}
	case CodeBallotCast:
		return &BallotCast{}
	case CodeBallotCounted:
		return &BallotCounted{}
	case CodeResult:
		return &Result{}
	case CodeOrder:
		return &Order{}
	case CodeFreeze:
		return &Freeze{}
	case CodeThaw:
		return &Thaw{}
	case CodeConfiscation:
		return &Confiscation{}
	case CodeReconciliation:
		return &Reconciliation{}
	case CodeEstablishment:
		return &Establishment{}
	case CodeAddition:
		return &Addition{}
	case CodeAlteration:
		return &Alteration{}
	case CodeRemoval:
		return &Removal{}
	case CodeMessage:
		return &Message{}
	case CodeRejection:
		return &Rejection{}
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

func (a *ContractOffer) Code() string {
	return CodeContractOffer
}

func (a *ContractFormation) Code() string {
	return CodeContractFormation
}

func (a *ContractAmendment) Code() string {
	return CodeContractAmendment
}

func (a *StaticContractFormation) Code() string {
	return CodeStaticContractFormation
}

func (a *ContractAddressChange) Code() string {
	return CodeContractAddressChange
}

func (a *BodyOfAgreementOffer) Code() string {
	return CodeBodyOfAgreementOffer
}

func (a *BodyOfAgreementFormation) Code() string {
	return CodeBodyOfAgreementFormation
}

func (a *BodyOfAgreementAmendment) Code() string {
	return CodeBodyOfAgreementAmendment
}

func (a *InstrumentDefinition) Code() string {
	return CodeInstrumentDefinition
}

func (a *InstrumentCreation) Code() string {
	return CodeInstrumentCreation
}

func (a *InstrumentModification) Code() string {
	return CodeInstrumentModification
}

func (a *Transfer) Code() string {
	return CodeTransfer
}

func (a *Settlement) Code() string {
	return CodeSettlement
}

func (a *Proposal) Code() string {
	return CodeProposal
}

func (a *Vote) Code() string {
	return CodeVote
}

func (a *BallotCast) Code() string {
	return CodeBallotCast
}

func (a *BallotCounted) Code() string {
	return CodeBallotCounted
}

func (a *Result) Code() string {
	return CodeResult
}

func (a *Order) Code() string {
	return CodeOrder
}

func (a *Freeze) Code() string {
	return CodeFreeze
}

func (a *Thaw) Code() string {
	return CodeThaw
}

func (a *Confiscation) Code() string {
	return CodeConfiscation
}

func (a *Reconciliation) Code() string {
	return CodeReconciliation
}

func (a *Establishment) Code() string {
	return CodeEstablishment
}

func (a *Addition) Code() string {
	return CodeAddition
}

func (a *Alteration) Code() string {
	return CodeAlteration
}

func (a *Removal) Code() string {
	return CodeRemoval
}

func (a *Message) Code() string {
	return CodeMessage
}

func (a *Rejection) Code() string {
	return CodeRejection
}
