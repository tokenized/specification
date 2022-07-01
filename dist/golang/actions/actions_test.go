package actions

import (
	"testing"

	"github.com/tokenized/pkg/bitcoin"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// ContractOffer identifies a payload as a ContractOffer message.
	actionContractOfferScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeContractOffer)),
	}
	actionContractOffer, err := DeserializeV1(actionContractOfferScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractOffer : %s", err)
	}
	_, ok = actionContractOffer.(*ContractOffer)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractOffer")
	}

	// ContractFormation identifies a payload as a ContractFormation message.
	actionContractFormationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeContractFormation)),
	}
	actionContractFormation, err := DeserializeV1(actionContractFormationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractFormation : %s", err)
	}
	_, ok = actionContractFormation.(*ContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractFormation")
	}

	// ContractAmendment identifies a payload as a ContractAmendment message.
	actionContractAmendmentScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeContractAmendment)),
	}
	actionContractAmendment, err := DeserializeV1(actionContractAmendmentScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAmendment : %s", err)
	}
	_, ok = actionContractAmendment.(*ContractAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAmendment")
	}

	// StaticContractFormation identifies a payload as a StaticContractFormation message.
	actionStaticContractFormationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeStaticContractFormation)),
	}
	actionStaticContractFormation, err := DeserializeV1(actionStaticContractFormationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for StaticContractFormation : %s", err)
	}
	_, ok = actionStaticContractFormation.(*StaticContractFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for StaticContractFormation")
	}

	// ContractAddressChange identifies a payload as a ContractAddressChange message.
	actionContractAddressChangeScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeContractAddressChange)),
	}
	actionContractAddressChange, err := DeserializeV1(actionContractAddressChangeScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for ContractAddressChange : %s", err)
	}
	_, ok = actionContractAddressChange.(*ContractAddressChange)
	if !ok {
		t.Fatalf("Failed deserialize type check for ContractAddressChange")
	}

	// BodyOfAgreementOffer identifies a payload as a BodyOfAgreementOffer message.
	actionBodyOfAgreementOfferScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeBodyOfAgreementOffer)),
	}
	actionBodyOfAgreementOffer, err := DeserializeV1(actionBodyOfAgreementOfferScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementOffer : %s", err)
	}
	_, ok = actionBodyOfAgreementOffer.(*BodyOfAgreementOffer)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementOffer")
	}

	// BodyOfAgreementFormation identifies a payload as a BodyOfAgreementFormation message.
	actionBodyOfAgreementFormationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeBodyOfAgreementFormation)),
	}
	actionBodyOfAgreementFormation, err := DeserializeV1(actionBodyOfAgreementFormationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementFormation : %s", err)
	}
	_, ok = actionBodyOfAgreementFormation.(*BodyOfAgreementFormation)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementFormation")
	}

	// BodyOfAgreementAmendment identifies a payload as a BodyOfAgreementAmendment message.
	actionBodyOfAgreementAmendmentScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeBodyOfAgreementAmendment)),
	}
	actionBodyOfAgreementAmendment, err := DeserializeV1(actionBodyOfAgreementAmendmentScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for BodyOfAgreementAmendment : %s", err)
	}
	_, ok = actionBodyOfAgreementAmendment.(*BodyOfAgreementAmendment)
	if !ok {
		t.Fatalf("Failed deserialize type check for BodyOfAgreementAmendment")
	}

	// InstrumentDefinition identifies a payload as a InstrumentDefinition message.
	actionInstrumentDefinitionScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeInstrumentDefinition)),
	}
	actionInstrumentDefinition, err := DeserializeV1(actionInstrumentDefinitionScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentDefinition : %s", err)
	}
	_, ok = actionInstrumentDefinition.(*InstrumentDefinition)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentDefinition")
	}

	// InstrumentCreation identifies a payload as a InstrumentCreation message.
	actionInstrumentCreationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeInstrumentCreation)),
	}
	actionInstrumentCreation, err := DeserializeV1(actionInstrumentCreationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentCreation : %s", err)
	}
	_, ok = actionInstrumentCreation.(*InstrumentCreation)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentCreation")
	}

	// InstrumentModification identifies a payload as a InstrumentModification message.
	actionInstrumentModificationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeInstrumentModification)),
	}
	actionInstrumentModification, err := DeserializeV1(actionInstrumentModificationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for InstrumentModification : %s", err)
	}
	_, ok = actionInstrumentModification.(*InstrumentModification)
	if !ok {
		t.Fatalf("Failed deserialize type check for InstrumentModification")
	}

	// Transfer identifies a payload as a Transfer message.
	actionTransferScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeTransfer)),
	}
	actionTransfer, err := DeserializeV1(actionTransferScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Transfer : %s", err)
	}
	_, ok = actionTransfer.(*Transfer)
	if !ok {
		t.Fatalf("Failed deserialize type check for Transfer")
	}

	// Settlement identifies a payload as a Settlement message.
	actionSettlementScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeSettlement)),
	}
	actionSettlement, err := DeserializeV1(actionSettlementScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Settlement : %s", err)
	}
	_, ok = actionSettlement.(*Settlement)
	if !ok {
		t.Fatalf("Failed deserialize type check for Settlement")
	}

	// Proposal identifies a payload as a Proposal message.
	actionProposalScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeProposal)),
	}
	actionProposal, err := DeserializeV1(actionProposalScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Proposal : %s", err)
	}
	_, ok = actionProposal.(*Proposal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Proposal")
	}

	// Vote identifies a payload as a Vote message.
	actionVoteScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeVote)),
	}
	actionVote, err := DeserializeV1(actionVoteScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Vote : %s", err)
	}
	_, ok = actionVote.(*Vote)
	if !ok {
		t.Fatalf("Failed deserialize type check for Vote")
	}

	// BallotCast identifies a payload as a BallotCast message.
	actionBallotCastScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeBallotCast)),
	}
	actionBallotCast, err := DeserializeV1(actionBallotCastScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCast : %s", err)
	}
	_, ok = actionBallotCast.(*BallotCast)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCast")
	}

	// BallotCounted identifies a payload as a BallotCounted message.
	actionBallotCountedScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeBallotCounted)),
	}
	actionBallotCounted, err := DeserializeV1(actionBallotCountedScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for BallotCounted : %s", err)
	}
	_, ok = actionBallotCounted.(*BallotCounted)
	if !ok {
		t.Fatalf("Failed deserialize type check for BallotCounted")
	}

	// Result identifies a payload as a Result message.
	actionResultScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeResult)),
	}
	actionResult, err := DeserializeV1(actionResultScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Result : %s", err)
	}
	_, ok = actionResult.(*Result)
	if !ok {
		t.Fatalf("Failed deserialize type check for Result")
	}

	// Order identifies a payload as a Order message.
	actionOrderScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeOrder)),
	}
	actionOrder, err := DeserializeV1(actionOrderScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Order : %s", err)
	}
	_, ok = actionOrder.(*Order)
	if !ok {
		t.Fatalf("Failed deserialize type check for Order")
	}

	// Freeze identifies a payload as a Freeze message.
	actionFreezeScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeFreeze)),
	}
	actionFreeze, err := DeserializeV1(actionFreezeScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Freeze : %s", err)
	}
	_, ok = actionFreeze.(*Freeze)
	if !ok {
		t.Fatalf("Failed deserialize type check for Freeze")
	}

	// Thaw identifies a payload as a Thaw message.
	actionThawScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeThaw)),
	}
	actionThaw, err := DeserializeV1(actionThawScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Thaw : %s", err)
	}
	_, ok = actionThaw.(*Thaw)
	if !ok {
		t.Fatalf("Failed deserialize type check for Thaw")
	}

	// Confiscation identifies a payload as a Confiscation message.
	actionConfiscationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeConfiscation)),
	}
	actionConfiscation, err := DeserializeV1(actionConfiscationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Confiscation : %s", err)
	}
	_, ok = actionConfiscation.(*Confiscation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Confiscation")
	}

	// Reconciliation identifies a payload as a Reconciliation message.
	actionReconciliationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeReconciliation)),
	}
	actionReconciliation, err := DeserializeV1(actionReconciliationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Reconciliation : %s", err)
	}
	_, ok = actionReconciliation.(*Reconciliation)
	if !ok {
		t.Fatalf("Failed deserialize type check for Reconciliation")
	}

	// Establishment identifies a payload as a Establishment message.
	actionEstablishmentScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeEstablishment)),
	}
	actionEstablishment, err := DeserializeV1(actionEstablishmentScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Establishment : %s", err)
	}
	_, ok = actionEstablishment.(*Establishment)
	if !ok {
		t.Fatalf("Failed deserialize type check for Establishment")
	}

	// Addition identifies a payload as a Addition message.
	actionAdditionScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeAddition)),
	}
	actionAddition, err := DeserializeV1(actionAdditionScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Addition : %s", err)
	}
	_, ok = actionAddition.(*Addition)
	if !ok {
		t.Fatalf("Failed deserialize type check for Addition")
	}

	// Alteration identifies a payload as a Alteration message.
	actionAlterationScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeAlteration)),
	}
	actionAlteration, err := DeserializeV1(actionAlterationScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Alteration : %s", err)
	}
	_, ok = actionAlteration.(*Alteration)
	if !ok {
		t.Fatalf("Failed deserialize type check for Alteration")
	}

	// Removal identifies a payload as a Removal message.
	actionRemovalScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeRemoval)),
	}
	actionRemoval, err := DeserializeV1(actionRemovalScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Removal : %s", err)
	}
	_, ok = actionRemoval.(*Removal)
	if !ok {
		t.Fatalf("Failed deserialize type check for Removal")
	}

	// Message identifies a payload as a Message message.
	actionMessageScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeMessage)),
	}
	actionMessage, err := DeserializeV1(actionMessageScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Message : %s", err)
	}
	_, ok = actionMessage.(*Message)
	if !ok {
		t.Fatalf("Failed deserialize type check for Message")
	}

	// Rejection identifies a payload as a Rejection message.
	actionRejectionScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(CodeRejection)),
	}
	actionRejection, err := DeserializeV1(actionRejectionScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for Rejection : %s", err)
	}
	_, ok = actionRejection.(*Rejection)
	if !ok {
		t.Fatalf("Failed deserialize type check for Rejection")
	}

}
