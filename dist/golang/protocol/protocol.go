// Package protocol provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.
package protocol

const (
	// CodeAssetDefinition identifies data as a AssetDefinition message.
	CodeAssetDefinition = "A1"

	// CodeAssetCreation identifies data as a AssetCreation message.
	CodeAssetCreation = "A2"

	// CodeAssetModification identifies data as a AssetModification message.
	CodeAssetModification = "A3"

	// CodeContractOffer identifies data as a ContractOffer message.
	CodeContractOffer = "C1"

	// CodeContractFormation identifies data as a ContractFormation message.
	CodeContractFormation = "C2"

	// CodeContractAmendment identifies data as a ContractAmendment message.
	CodeContractAmendment = "C3"

	// CodeStaticContractFormation identifies data as a
	// StaticContractFormation message.
	CodeStaticContractFormation = "C4"

	// CodeOrder identifies data as a Order message.
	CodeOrder = "E1"

	// CodeFreeze identifies data as a Freeze message.
	CodeFreeze = "E2"

	// CodeThaw identifies data as a Thaw message.
	CodeThaw = "E3"

	// CodeConfiscation identifies data as a Confiscation message.
	CodeConfiscation = "E4"

	// CodeReconciliation identifies data as a Reconciliation message.
	CodeReconciliation = "E5"

	// CodeInitiative identifies data as a Initiative message.
	CodeInitiative = "G1"

	// CodeReferendum identifies data as a Referendum message.
	CodeReferendum = "G2"

	// CodeVote identifies data as a Vote message.
	CodeVote = "G3"

	// CodeBallotCast identifies data as a BallotCast message.
	CodeBallotCast = "G4"

	// CodeBallotCounted identifies data as a BallotCounted message.
	CodeBallotCounted = "G5"

	// CodeResult identifies data as a Result message.
	CodeResult = "G6"

	// CodeMessage identifies data as a Message message.
	CodeMessage = "M1"

	// CodeRejection identifies data as a Rejection message.
	CodeRejection = "M2"

	// CodeEstablishment identifies data as a Establishment message.
	CodeEstablishment = "R1"

	// CodeAddition identifies data as a Addition message.
	CodeAddition = "R2"

	// CodeAlteration identifies data as a Alteration message.
	CodeAlteration = "R3"

	// CodeRemoval identifies data as a Removal message.
	CodeRemoval = "R4"

	// CodeSend identifies data as a Send message.
	CodeSend = "T1"

	// CodeExchange identifies data as a Exchange message.
	CodeExchange = "T2"

	// CodeSwap identifies data as a Swap message.
	CodeSwap = "T3"

	// CodeSettlement identifies data as a Settlement message.
	CodeSettlement = "T4"
)

// TypeMapping holds a mapping of message codes to message types.
var (
	TypeMapping = map[string]OpReturnMessage{
		CodeAssetDefinition: &AssetDefinition{},
		CodeAssetCreation: &AssetCreation{},
		CodeAssetModification: &AssetModification{},
		CodeContractOffer: &ContractOffer{},
		CodeContractFormation: &ContractFormation{},
		CodeContractAmendment: &ContractAmendment{},
		CodeStaticContractFormation: &StaticContractFormation{},
		CodeOrder: &Order{},
		CodeFreeze: &Freeze{},
		CodeThaw: &Thaw{},
		CodeConfiscation: &Confiscation{},
		CodeReconciliation: &Reconciliation{},
		CodeInitiative: &Initiative{},
		CodeReferendum: &Referendum{},
		CodeVote: &Vote{},
		CodeBallotCast: &BallotCast{},
		CodeBallotCounted: &BallotCounted{},
		CodeResult: &Result{},
		CodeMessage: &Message{},
		CodeRejection: &Rejection{},
		CodeEstablishment: &Establishment{},
		CodeAddition: &Addition{},
		CodeAlteration: &Alteration{},
		CodeRemoval: &Removal{},
		CodeSend: &Send{},
		CodeExchange: &Exchange{},
		CodeSwap: &Swap{},
		CodeSettlement: &Settlement{},
	}

	// ProtocolID is the current protocol ID
	ProtocolID = []byte("tokenized.com")
)
