package bsor

import (
	"reflect"

	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/actions"
)

func Generateactions() (bsor.Definitions, error) {
	return bsor.BuildDefinitions(
		reflect.TypeOf(actions.ContractOffer{}),
		reflect.TypeOf(actions.ContractFormation{}),
		reflect.TypeOf(actions.ContractAmendment{}),
		reflect.TypeOf(actions.StaticContractFormation{}),
		reflect.TypeOf(actions.ContractAddressChange{}),
		reflect.TypeOf(actions.BodyOfAgreementOffer{}),
		reflect.TypeOf(actions.BodyOfAgreementFormation{}),
		reflect.TypeOf(actions.BodyOfAgreementAmendment{}),
		reflect.TypeOf(actions.InstrumentDefinition{}),
		reflect.TypeOf(actions.InstrumentCreation{}),
		reflect.TypeOf(actions.InstrumentModification{}),
		reflect.TypeOf(actions.Transfer{}),
		reflect.TypeOf(actions.Settlement{}),
		reflect.TypeOf(actions.Proposal{}),
		reflect.TypeOf(actions.Vote{}),
		reflect.TypeOf(actions.BallotCast{}),
		reflect.TypeOf(actions.BallotCounted{}),
		reflect.TypeOf(actions.Result{}),
		reflect.TypeOf(actions.Order{}),
		reflect.TypeOf(actions.Freeze{}),
		reflect.TypeOf(actions.Thaw{}),
		reflect.TypeOf(actions.Confiscation{}),
		reflect.TypeOf(actions.Reconciliation{}),
		reflect.TypeOf(actions.Establishment{}),
		reflect.TypeOf(actions.Addition{}),
		reflect.TypeOf(actions.Alteration{}),
		reflect.TypeOf(actions.Removal{}),
		reflect.TypeOf(actions.Message{}),
		reflect.TypeOf(actions.Rejection{}),
	)
}
