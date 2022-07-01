package actions

import (
	"testing"
)

func TestEmptyDeserializeV0(t *testing.T) {
	var ok bool
	// ContractOffer identifies a payload as a ContractOffer message.
	actionContractOffer, err := DeserializeV0([]byte(CodeContractOffer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractOffer : %s", err)
	}
	_, ok = actionContractOffer.(*ContractOffer)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractOffer")
	}

	// ContractFormation identifies a payload as a ContractFormation message.
	actionContractFormation, err := DeserializeV0([]byte(CodeContractFormation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractFormation : %s", err)
	}
	_, ok = actionContractFormation.(*ContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractFormation")
	}

	// ContractAmendment identifies a payload as a ContractAmendment message.
	actionContractAmendment, err := DeserializeV0([]byte(CodeContractAmendment), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAmendment : %s", err)
	}
	_, ok = actionContractAmendment.(*ContractAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAmendment")
	}

	// StaticContractFormation identifies a payload as a StaticContractFormation message.
	actionStaticContractFormation, err := DeserializeV0([]byte(CodeStaticContractFormation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for StaticContractFormation : %s", err)
	}
	_, ok = actionStaticContractFormation.(*StaticContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for StaticContractFormation")
	}

	// ContractAddressChange identifies a payload as a ContractAddressChange message.
	actionContractAddressChange, err := DeserializeV0([]byte(CodeContractAddressChange), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAddressChange : %s", err)
	}
	_, ok = actionContractAddressChange.(*ContractAddressChange)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAddressChange")
	}

	// BodyOfAgreementOffer identifies a payload as a BodyOfAgreementOffer message.
	actionBodyOfAgreementOffer, err := DeserializeV0([]byte(CodeBodyOfAgreementOffer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementOffer : %s", err)
	}
	_, ok = actionBodyOfAgreementOffer.(*BodyOfAgreementOffer)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementOffer")
	}

	// BodyOfAgreementFormation identifies a payload as a BodyOfAgreementFormation message.
	actionBodyOfAgreementFormation, err := DeserializeV0([]byte(CodeBodyOfAgreementFormation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementFormation : %s", err)
	}
	_, ok = actionBodyOfAgreementFormation.(*BodyOfAgreementFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementFormation")
	}

	// BodyOfAgreementAmendment identifies a payload as a BodyOfAgreementAmendment message.
	actionBodyOfAgreementAmendment, err := DeserializeV0([]byte(CodeBodyOfAgreementAmendment), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementAmendment : %s", err)
	}
	_, ok = actionBodyOfAgreementAmendment.(*BodyOfAgreementAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementAmendment")
	}

	// InstrumentDefinition identifies a payload as a InstrumentDefinition message.
	actionInstrumentDefinition, err := DeserializeV0([]byte(CodeInstrumentDefinition), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentDefinition : %s", err)
	}
	_, ok = actionInstrumentDefinition.(*InstrumentDefinition)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentDefinition")
	}

	// InstrumentCreation identifies a payload as a InstrumentCreation message.
	actionInstrumentCreation, err := DeserializeV0([]byte(CodeInstrumentCreation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentCreation : %s", err)
	}
	_, ok = actionInstrumentCreation.(*InstrumentCreation)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentCreation")
	}

	// InstrumentModification identifies a payload as a InstrumentModification message.
	actionInstrumentModification, err := DeserializeV0([]byte(CodeInstrumentModification), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentModification : %s", err)
	}
	_, ok = actionInstrumentModification.(*InstrumentModification)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentModification")
	}

	// Transfer identifies a payload as a Transfer message.
	actionTransfer, err := DeserializeV0([]byte(CodeTransfer), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Transfer : %s", err)
	}
	_, ok = actionTransfer.(*Transfer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Transfer")
	}

	// Settlement identifies a payload as a Settlement message.
	actionSettlement, err := DeserializeV0([]byte(CodeSettlement), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Settlement : %s", err)
	}
	_, ok = actionSettlement.(*Settlement)
	if !ok {
		t.Fatalf("Failed deserialize type check for Settlement")
	}

	// Proposal identifies a payload as a Proposal message.
	actionProposal, err := DeserializeV0([]byte(CodeProposal), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Proposal : %s", err)
	}
	_, ok = actionProposal.(*Proposal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Proposal")
	}

	// Vote identifies a payload as a Vote message.
	actionVote, err := DeserializeV0([]byte(CodeVote), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Vote : %s", err)
	}
	_, ok = actionVote.(*Vote)
	if !ok {
		t.Fatalf("Failed deserialize type check for Vote")
	}

	// BallotCast identifies a payload as a BallotCast message.
	actionBallotCast, err := DeserializeV0([]byte(CodeBallotCast), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCast : %s", err)
	}
	_, ok = actionBallotCast.(*BallotCast)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCast")
	}

	// BallotCounted identifies a payload as a BallotCounted message.
	actionBallotCounted, err := DeserializeV0([]byte(CodeBallotCounted), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCounted : %s", err)
	}
	_, ok = actionBallotCounted.(*BallotCounted)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCounted")
	}

	// Result identifies a payload as a Result message.
	actionResult, err := DeserializeV0([]byte(CodeResult), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Result : %s", err)
	}
	_, ok = actionResult.(*Result)
	if !ok {
		t.Fatalf("Failed deserialize type check for Result")
	}

	// Order identifies a payload as a Order message.
	actionOrder, err := DeserializeV0([]byte(CodeOrder), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Order : %s", err)
	}
	_, ok = actionOrder.(*Order)
	if !ok {
		t.Fatalf("Failed deserialize type check for Order")
	}

	// Freeze identifies a payload as a Freeze message.
	actionFreeze, err := DeserializeV0([]byte(CodeFreeze), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Freeze : %s", err)
	}
	_, ok = actionFreeze.(*Freeze)
	if !ok {
		t.Fatalf("Failed deserialize type check for Freeze")
	}

	// Thaw identifies a payload as a Thaw message.
	actionThaw, err := DeserializeV0([]byte(CodeThaw), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Thaw : %s", err)
	}
	_, ok = actionThaw.(*Thaw)
	if !ok {
		t.Fatalf("Failed deserialize type check for Thaw")
	}

	// Confiscation identifies a payload as a Confiscation message.
	actionConfiscation, err := DeserializeV0([]byte(CodeConfiscation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Confiscation : %s", err)
	}
	_, ok = actionConfiscation.(*Confiscation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Confiscation")
	}

	// Reconciliation identifies a payload as a Reconciliation message.
	actionReconciliation, err := DeserializeV0([]byte(CodeReconciliation), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Reconciliation : %s", err)
	}
	_, ok = actionReconciliation.(*Reconciliation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Reconciliation")
	}

	// Establishment identifies a payload as a Establishment message.
	actionEstablishment, err := DeserializeV0([]byte(CodeEstablishment), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Establishment : %s", err)
	}
	_, ok = actionEstablishment.(*Establishment)
	if !ok {
		t.Fatalf("Failed deserialize type check for Establishment")
	}

	// Addition identifies a payload as a Addition message.
	actionAddition, err := DeserializeV0([]byte(CodeAddition), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Addition : %s", err)
	}
	_, ok = actionAddition.(*Addition)
	if !ok {
		t.Fatalf("Failed deserialize type check for Addition")
	}

	// Alteration identifies a payload as a Alteration message.
	actionAlteration, err := DeserializeV0([]byte(CodeAlteration), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Alteration : %s", err)
	}
	_, ok = actionAlteration.(*Alteration)
	if !ok {
		t.Fatalf("Failed deserialize type check for Alteration")
	}

	// Removal identifies a payload as a Removal message.
	actionRemoval, err := DeserializeV0([]byte(CodeRemoval), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Removal : %s", err)
	}
	_, ok = actionRemoval.(*Removal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Removal")
	}

	// Message identifies a payload as a Message message.
	actionMessage, err := DeserializeV0([]byte(CodeMessage), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Message : %s", err)
	}
	_, ok = actionMessage.(*Message)
	if !ok {
		t.Fatalf("Failed deserialize type check for Message")
	}

	// Rejection identifies a payload as a Rejection message.
	actionRejection, err := DeserializeV0([]byte(CodeRejection), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Rejection : %s", err)
	}
	_, ok = actionRejection.(*Rejection)
	if !ok {
		t.Fatalf("Failed deserialize type check for Rejection")
	}

}
