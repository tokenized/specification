package actions

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

	// CodeAssetDefinition identifies a payload as a AssetDefinition action message.
	CodeAssetDefinition = "A1"

	// CodeAssetCreation identifies a payload as a AssetCreation action message.
	CodeAssetCreation = "A2"

	// CodeAssetModification identifies a payload as a AssetModification action message.
	CodeAssetModification = "A3"

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
	case CodeAssetDefinition:
		return &AssetDefinition{}
	case CodeAssetCreation:
		return &AssetCreation{}
	case CodeAssetModification:
		return &AssetModification{}
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

func (a *AssetDefinition) Code() string {
	return CodeAssetDefinition
}

func (a *AssetCreation) Code() string {
	return CodeAssetCreation
}

func (a *AssetModification) Code() string {
	return CodeAssetModification
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
