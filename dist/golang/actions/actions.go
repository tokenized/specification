package actions

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Action interface {
	proto.Message

	Code() string
	TypeName() string

	// IsClientRequest returns true when this action is a client request to a smart contract agent.
	IsClientRequest() bool

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

	// CodeRectificationSettlement identifies a payload as a RectificationSettlement action message.
	CodeRectificationSettlement = "T3"

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

	// CodeDeprecatedReconciliation identifies a payload as a DeprecatedReconciliation action message.
	CodeDeprecatedReconciliation = "E5"

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

	// ComplianceActionDeprecatedReconciliation identifies a reconcilation type
	ComplianceActionDeprecatedReconciliation = "R"

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
	case CodeRectificationSettlement:
		return &RectificationSettlement{}
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
	case CodeDeprecatedReconciliation:
		return &DeprecatedReconciliation{}
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

// Deserialize reads an action from a byte slice.
func Deserialize(code []byte, payload []byte) (Action, error) {
	result := NewActionFromCode(string(code))
	if result == nil {
		return nil, fmt.Errorf("Unknown action code : %s", string(code))
	}

	if len(payload) == 0 {
		return result, nil // empty result
	}

	if len(payload) == 1 && payload[0] == bitcoin.OP_FALSE {
		return result, nil // empty result
	}

	if err := proto.Unmarshal(payload, result); err != nil {
		return nil, errors.Wrap(err, "protobuf unmarshal")
	}

	return result, nil
}

func (a *ContractOffer) Code() string {
	return CodeContractOffer
}

func (a *ContractOffer) TypeName() string {
	return "ContractOffer"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *ContractOffer) IsClientRequest() bool {
	return true
}

func (a *ContractFormation) Code() string {
	return CodeContractFormation
}

func (a *ContractFormation) TypeName() string {
	return "ContractFormation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *ContractFormation) IsClientRequest() bool {
	return false
}

func (a *ContractAmendment) Code() string {
	return CodeContractAmendment
}

func (a *ContractAmendment) TypeName() string {
	return "ContractAmendment"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *ContractAmendment) IsClientRequest() bool {
	return true
}

func (a *StaticContractFormation) Code() string {
	return CodeStaticContractFormation
}

func (a *StaticContractFormation) TypeName() string {
	return "StaticContractFormation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *StaticContractFormation) IsClientRequest() bool {
	return false
}

func (a *ContractAddressChange) Code() string {
	return CodeContractAddressChange
}

func (a *ContractAddressChange) TypeName() string {
	return "ContractAddressChange"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *ContractAddressChange) IsClientRequest() bool {
	return false
}

func (a *BodyOfAgreementOffer) Code() string {
	return CodeBodyOfAgreementOffer
}

func (a *BodyOfAgreementOffer) TypeName() string {
	return "BodyOfAgreementOffer"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *BodyOfAgreementOffer) IsClientRequest() bool {
	return true
}

func (a *BodyOfAgreementFormation) Code() string {
	return CodeBodyOfAgreementFormation
}

func (a *BodyOfAgreementFormation) TypeName() string {
	return "BodyOfAgreementFormation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *BodyOfAgreementFormation) IsClientRequest() bool {
	return false
}

func (a *BodyOfAgreementAmendment) Code() string {
	return CodeBodyOfAgreementAmendment
}

func (a *BodyOfAgreementAmendment) TypeName() string {
	return "BodyOfAgreementAmendment"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *BodyOfAgreementAmendment) IsClientRequest() bool {
	return true
}

func (a *InstrumentDefinition) Code() string {
	return CodeInstrumentDefinition
}

func (a *InstrumentDefinition) TypeName() string {
	return "InstrumentDefinition"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *InstrumentDefinition) IsClientRequest() bool {
	return true
}

func (a *InstrumentCreation) Code() string {
	return CodeInstrumentCreation
}

func (a *InstrumentCreation) TypeName() string {
	return "InstrumentCreation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *InstrumentCreation) IsClientRequest() bool {
	return false
}

func (a *InstrumentModification) Code() string {
	return CodeInstrumentModification
}

func (a *InstrumentModification) TypeName() string {
	return "InstrumentModification"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *InstrumentModification) IsClientRequest() bool {
	return true
}

func (a *Transfer) Code() string {
	return CodeTransfer
}

func (a *Transfer) TypeName() string {
	return "Transfer"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Transfer) IsClientRequest() bool {
	return true
}

func (a *Settlement) Code() string {
	return CodeSettlement
}

func (a *Settlement) TypeName() string {
	return "Settlement"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Settlement) IsClientRequest() bool {
	return false
}

func (a *RectificationSettlement) Code() string {
	return CodeRectificationSettlement
}

func (a *RectificationSettlement) TypeName() string {
	return "RectificationSettlement"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *RectificationSettlement) IsClientRequest() bool {
	return false
}

func (a *Proposal) Code() string {
	return CodeProposal
}

func (a *Proposal) TypeName() string {
	return "Proposal"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Proposal) IsClientRequest() bool {
	return true
}

func (a *Vote) Code() string {
	return CodeVote
}

func (a *Vote) TypeName() string {
	return "Vote"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Vote) IsClientRequest() bool {
	return false
}

func (a *BallotCast) Code() string {
	return CodeBallotCast
}

func (a *BallotCast) TypeName() string {
	return "BallotCast"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *BallotCast) IsClientRequest() bool {
	return true
}

func (a *BallotCounted) Code() string {
	return CodeBallotCounted
}

func (a *BallotCounted) TypeName() string {
	return "BallotCounted"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *BallotCounted) IsClientRequest() bool {
	return false
}

func (a *Result) Code() string {
	return CodeResult
}

func (a *Result) TypeName() string {
	return "Result"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Result) IsClientRequest() bool {
	return false
}

func (a *Order) Code() string {
	return CodeOrder
}

func (a *Order) TypeName() string {
	return "Order"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Order) IsClientRequest() bool {
	return true
}

func (a *Freeze) Code() string {
	return CodeFreeze
}

func (a *Freeze) TypeName() string {
	return "Freeze"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Freeze) IsClientRequest() bool {
	return false
}

func (a *Thaw) Code() string {
	return CodeThaw
}

func (a *Thaw) TypeName() string {
	return "Thaw"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Thaw) IsClientRequest() bool {
	return false
}

func (a *Confiscation) Code() string {
	return CodeConfiscation
}

func (a *Confiscation) TypeName() string {
	return "Confiscation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Confiscation) IsClientRequest() bool {
	return false
}

func (a *DeprecatedReconciliation) Code() string {
	return CodeDeprecatedReconciliation
}

func (a *DeprecatedReconciliation) TypeName() string {
	return "DeprecatedReconciliation"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *DeprecatedReconciliation) IsClientRequest() bool {
	return false
}

func (a *Establishment) Code() string {
	return CodeEstablishment
}

func (a *Establishment) TypeName() string {
	return "Establishment"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Establishment) IsClientRequest() bool {
	return false
}

func (a *Addition) Code() string {
	return CodeAddition
}

func (a *Addition) TypeName() string {
	return "Addition"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Addition) IsClientRequest() bool {
	return false
}

func (a *Alteration) Code() string {
	return CodeAlteration
}

func (a *Alteration) TypeName() string {
	return "Alteration"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Alteration) IsClientRequest() bool {
	return false
}

func (a *Removal) Code() string {
	return CodeRemoval
}

func (a *Removal) TypeName() string {
	return "Removal"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Removal) IsClientRequest() bool {
	return false
}

func (a *Message) Code() string {
	return CodeMessage
}

func (a *Message) TypeName() string {
	return "Message"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Message) IsClientRequest() bool {
	return false
}

func (a *Rejection) Code() string {
	return CodeRejection
}

func (a *Rejection) TypeName() string {
	return "Rejection"
}

// IsClientRequest returns true when this action is a client request to a smart contract agent.
func (a *Rejection) IsClientRequest() bool {
	return false
}

// Formation creates a contract formation with all the values from the contract offer.
func (a *ContractOffer) Formation() (*ContractFormation, error) {
	return &ContractFormation{
		ContractName:        a.ContractName,
		BodyOfAgreementType: a.BodyOfAgreementType,
		BodyOfAgreement:     a.BodyOfAgreement,
		// deprecated ContractType deprecated
		SupportingDocs: a.SupportingDocs,
		// deprecated GoverningLaw deprecated
		// deprecated Jurisdiction deprecated
		ContractExpiration: a.ContractExpiration,
		ContractURI:        a.ContractURI,
		Issuer:             a.Issuer,
		// deprecated IssuerLogoURL deprecated
		// deprecated ContractOperator deprecated
		// deprecated AdminOracle deprecated
		// deprecated AdminOracleSignature deprecated
		// deprecated AdminOracleSigBlockHeight deprecated
		ContractFee:               a.ContractFee,
		VotingSystems:             a.VotingSystems,
		ContractPermissions:       a.ContractPermissions,
		RestrictedQtyInstruments:  a.RestrictedQtyInstruments,
		AdministrationProposal:    a.AdministrationProposal,
		HolderProposal:            a.HolderProposal,
		Oracles:                   a.Oracles,
		MasterAddress:             a.MasterAddress,
		EntityContract:            a.EntityContract,
		OperatorEntityContract:    a.OperatorEntityContract,
		ContractType:              a.ContractType,
		Services:                  a.Services,
		AdminIdentityCertificates: a.AdminIdentityCertificates,
		GoverningLaw:              a.GoverningLaw,
		Jurisdiction:              a.Jurisdiction,
	}, nil
}

// Formation creates a body of agreement formation with all the values from the body of agreement
// offer.
func (a *BodyOfAgreementOffer) Formation() (*BodyOfAgreementFormation, error) {
	return &BodyOfAgreementFormation{
		Chapters:    a.Chapters,
		Definitions: a.Definitions,
	}, nil
}

// Creation creates an instrument creation with all the values from the instrument definition.
func (a *InstrumentDefinition) Creation() (*InstrumentCreation, error) {
	return &InstrumentCreation{
		InstrumentPermissions: a.InstrumentPermissions,
		// deprecated TransfersPermitted deprecated
		// deprecated TradeRestrictionsDeprecated deprecated
		EnforcementOrdersPermitted:       a.EnforcementOrdersPermitted,
		VotingRights:                     a.VotingRights,
		VoteMultiplier:                   a.VoteMultiplier,
		AdministrationProposal:           a.AdministrationProposal,
		HolderProposal:                   a.HolderProposal,
		InstrumentModificationGovernance: a.InstrumentModificationGovernance,
		AuthorizedTokenQty:               a.AuthorizedTokenQty,
		InstrumentType:                   a.InstrumentType,
		InstrumentPayload:                a.InstrumentPayload,
		TradeRestrictions:                a.TradeRestrictions,
		TransferFee:                      a.TransferFee,
	}, nil
}
