package actions

import (
	"testing"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// ContractOffer identifies a payload as a ContractOffer message.
	actionContractOffer, err := Deserialize([]byte(CodeContractOffer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractOffer : %s", err)
	}
	_, ok = actionContractOffer.(*ContractOffer)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractOffer")
	}
	err = actionContractOffer.Validate()
	if err != nil {
		t.Fatalf("Failed validate for ContractOffer : %s", err)
	}

	// ContractFormation identifies a payload as a ContractFormation message.
	actionContractFormation, err := Deserialize([]byte(CodeContractFormation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractFormation : %s", err)
	}
	_, ok = actionContractFormation.(*ContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractFormation")
	}
	err = actionContractFormation.Validate()
	if err != nil {
		t.Fatalf("Failed validate for ContractFormation : %s", err)
	}

	// ContractAmendment identifies a payload as a ContractAmendment message.
	actionContractAmendment, err := Deserialize([]byte(CodeContractAmendment), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAmendment : %s", err)
	}
	_, ok = actionContractAmendment.(*ContractAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAmendment")
	}
	err = actionContractAmendment.Validate()
	if err != nil {
		t.Fatalf("Failed validate for ContractAmendment : %s", err)
	}

	// StaticContractFormation identifies a payload as a StaticContractFormation message.
	actionStaticContractFormation, err := Deserialize([]byte(CodeStaticContractFormation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for StaticContractFormation : %s", err)
	}
	_, ok = actionStaticContractFormation.(*StaticContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for StaticContractFormation")
	}
	err = actionStaticContractFormation.Validate()
	if err != nil {
		t.Fatalf("Failed validate for StaticContractFormation : %s", err)
	}

	// ContractAddressChange identifies a payload as a ContractAddressChange message.
	actionContractAddressChange, err := Deserialize([]byte(CodeContractAddressChange), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAddressChange : %s", err)
	}
	_, ok = actionContractAddressChange.(*ContractAddressChange)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAddressChange")
	}
	err = actionContractAddressChange.Validate()
	if err != nil {
		t.Fatalf("Failed validate for ContractAddressChange : %s", err)
	}

	// AssetDefinition identifies a payload as a AssetDefinition message.
	actionAssetDefinition, err := Deserialize([]byte(CodeAssetDefinition), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for AssetDefinition : %s", err)
	}
	_, ok = actionAssetDefinition.(*AssetDefinition)
	if !ok {
		t.Fatalf("Failed deserialize type check for AssetDefinition")
	}
	err = actionAssetDefinition.Validate()
	if err != nil {
		t.Fatalf("Failed validate for AssetDefinition : %s", err)
	}

	// AssetCreation identifies a payload as a AssetCreation message.
	actionAssetCreation, err := Deserialize([]byte(CodeAssetCreation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for AssetCreation : %s", err)
	}
	_, ok = actionAssetCreation.(*AssetCreation)
	if !ok {
		t.Fatalf("Failed deserialize type check for AssetCreation")
	}
	err = actionAssetCreation.Validate()
	if err != nil {
		t.Fatalf("Failed validate for AssetCreation : %s", err)
	}

	// AssetModification identifies a payload as a AssetModification message.
	actionAssetModification, err := Deserialize([]byte(CodeAssetModification), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for AssetModification : %s", err)
	}
	_, ok = actionAssetModification.(*AssetModification)
	if !ok {
		t.Fatalf("Failed deserialize type check for AssetModification")
	}
	err = actionAssetModification.Validate()
	if err != nil {
		t.Fatalf("Failed validate for AssetModification : %s", err)
	}

	// Transfer identifies a payload as a Transfer message.
	actionTransfer, err := Deserialize([]byte(CodeTransfer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Transfer : %s", err)
	}
	_, ok = actionTransfer.(*Transfer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Transfer")
	}
	err = actionTransfer.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Transfer : %s", err)
	}

	// Settlement identifies a payload as a Settlement message.
	actionSettlement, err := Deserialize([]byte(CodeSettlement), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Settlement : %s", err)
	}
	_, ok = actionSettlement.(*Settlement)
	if !ok {
		t.Fatalf("Failed deserialize type check for Settlement")
	}
	err = actionSettlement.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Settlement : %s", err)
	}

	// Proposal identifies a payload as a Proposal message.
	actionProposal, err := Deserialize([]byte(CodeProposal), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Proposal : %s", err)
	}
	_, ok = actionProposal.(*Proposal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Proposal")
	}
	err = actionProposal.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Proposal : %s", err)
	}

	// Vote identifies a payload as a Vote message.
	actionVote, err := Deserialize([]byte(CodeVote), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Vote : %s", err)
	}
	_, ok = actionVote.(*Vote)
	if !ok {
		t.Fatalf("Failed deserialize type check for Vote")
	}
	err = actionVote.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Vote : %s", err)
	}

	// BallotCast identifies a payload as a BallotCast message.
	actionBallotCast, err := Deserialize([]byte(CodeBallotCast), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCast : %s", err)
	}
	_, ok = actionBallotCast.(*BallotCast)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCast")
	}
	err = actionBallotCast.Validate()
	if err != nil {
		t.Fatalf("Failed validate for BallotCast : %s", err)
	}

	// BallotCounted identifies a payload as a BallotCounted message.
	actionBallotCounted, err := Deserialize([]byte(CodeBallotCounted), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCounted : %s", err)
	}
	_, ok = actionBallotCounted.(*BallotCounted)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCounted")
	}
	err = actionBallotCounted.Validate()
	if err != nil {
		t.Fatalf("Failed validate for BallotCounted : %s", err)
	}

	// Result identifies a payload as a Result message.
	actionResult, err := Deserialize([]byte(CodeResult), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Result : %s", err)
	}
	_, ok = actionResult.(*Result)
	if !ok {
		t.Fatalf("Failed deserialize type check for Result")
	}
	err = actionResult.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Result : %s", err)
	}

	// Order identifies a payload as a Order message.
	actionOrder, err := Deserialize([]byte(CodeOrder), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Order : %s", err)
	}
	_, ok = actionOrder.(*Order)
	if !ok {
		t.Fatalf("Failed deserialize type check for Order")
	}
	err = actionOrder.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Order : %s", err)
	}

	// Freeze identifies a payload as a Freeze message.
	actionFreeze, err := Deserialize([]byte(CodeFreeze), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Freeze : %s", err)
	}
	_, ok = actionFreeze.(*Freeze)
	if !ok {
		t.Fatalf("Failed deserialize type check for Freeze")
	}
	err = actionFreeze.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Freeze : %s", err)
	}

	// Thaw identifies a payload as a Thaw message.
	actionThaw, err := Deserialize([]byte(CodeThaw), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Thaw : %s", err)
	}
	_, ok = actionThaw.(*Thaw)
	if !ok {
		t.Fatalf("Failed deserialize type check for Thaw")
	}
	err = actionThaw.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Thaw : %s", err)
	}

	// Confiscation identifies a payload as a Confiscation message.
	actionConfiscation, err := Deserialize([]byte(CodeConfiscation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Confiscation : %s", err)
	}
	_, ok = actionConfiscation.(*Confiscation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Confiscation")
	}
	err = actionConfiscation.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Confiscation : %s", err)
	}

	// Reconciliation identifies a payload as a Reconciliation message.
	actionReconciliation, err := Deserialize([]byte(CodeReconciliation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Reconciliation : %s", err)
	}
	_, ok = actionReconciliation.(*Reconciliation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Reconciliation")
	}
	err = actionReconciliation.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Reconciliation : %s", err)
	}

	// Establishment identifies a payload as a Establishment message.
	actionEstablishment, err := Deserialize([]byte(CodeEstablishment), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Establishment : %s", err)
	}
	_, ok = actionEstablishment.(*Establishment)
	if !ok {
		t.Fatalf("Failed deserialize type check for Establishment")
	}
	err = actionEstablishment.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Establishment : %s", err)
	}

	// Addition identifies a payload as a Addition message.
	actionAddition, err := Deserialize([]byte(CodeAddition), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Addition : %s", err)
	}
	_, ok = actionAddition.(*Addition)
	if !ok {
		t.Fatalf("Failed deserialize type check for Addition")
	}
	err = actionAddition.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Addition : %s", err)
	}

	// Alteration identifies a payload as a Alteration message.
	actionAlteration, err := Deserialize([]byte(CodeAlteration), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Alteration : %s", err)
	}
	_, ok = actionAlteration.(*Alteration)
	if !ok {
		t.Fatalf("Failed deserialize type check for Alteration")
	}
	err = actionAlteration.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Alteration : %s", err)
	}

	// Removal identifies a payload as a Removal message.
	actionRemoval, err := Deserialize([]byte(CodeRemoval), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Removal : %s", err)
	}
	_, ok = actionRemoval.(*Removal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Removal")
	}
	err = actionRemoval.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Removal : %s", err)
	}

	// Message identifies a payload as a Message message.
	actionMessage, err := Deserialize([]byte(CodeMessage), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Message : %s", err)
	}
	_, ok = actionMessage.(*Message)
	if !ok {
		t.Fatalf("Failed deserialize type check for Message")
	}
	err = actionMessage.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Message : %s", err)
	}

	// Rejection identifies a payload as a Rejection message.
	actionRejection, err := Deserialize([]byte(CodeRejection), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Rejection : %s", err)
	}
	_, ok = actionRejection.(*Rejection)
	if !ok {
		t.Fatalf("Failed deserialize type check for Rejection")
	}
	err = actionRejection.Validate()
	if err != nil {
		t.Fatalf("Failed validate for Rejection : %s", err)
	}

}
