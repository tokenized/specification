package bsor

import (
	"reflect"

	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/instruments"
)

func Generateinstruments() (bsor.Definitions, error) {
	return bsor.BuildDefinitions(
		reflect.TypeOf(instruments.Membership{}),
		reflect.TypeOf(instruments.Currency{}),
		reflect.TypeOf(instruments.ShareCommon{}),
		reflect.TypeOf(instruments.BondFixedRate{}),
		reflect.TypeOf(instruments.Coupon{}),
		reflect.TypeOf(instruments.LoyaltyPoints{}),
		reflect.TypeOf(instruments.TicketAdmission{}),
		reflect.TypeOf(instruments.CasinoChip{}),
		reflect.TypeOf(instruments.InformationServiceLicense{}),
	)
}
