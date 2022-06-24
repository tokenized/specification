package bsor

import (
	"reflect"

	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/messages"
)

func Generatemessages() (bsor.Definitions, error) {
	return bsor.BuildDefinitions(
		reflect.TypeOf(messages.PublicMessage{}),
		reflect.TypeOf(messages.PrivateMessage{}),
		reflect.TypeOf(messages.RevertedTx{}),
		reflect.TypeOf(messages.Offer{}),
		reflect.TypeOf(messages.SignatureRequest{}),
		reflect.TypeOf(messages.SettlementRequest{}),
		reflect.TypeOf(messages.OutputMetadata{}),
		reflect.TypeOf(messages.Distribution{}),
		reflect.TypeOf(messages.InitiateRelationship{}),
		reflect.TypeOf(messages.PendingAcceptRelationship{}),
		reflect.TypeOf(messages.AcceptRelationship{}),
		reflect.TypeOf(messages.RelationshipAmendment{}),
		reflect.TypeOf(messages.InitiateThread{}),
	)
}
