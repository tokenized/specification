package assets

import (
	"testing"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// Membership identifies a payload as a Membership message.
	assetMembership, err := Deserialize([]byte(CodeMembership), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Membership : %s", err)
	}
	_, ok = assetMembership.(*Membership)
	if !ok {
		t.Fatalf("Failed deserialize type check for Membership")
	}

	// Currency identifies a payload as a Currency message.
	assetCurrency, err := Deserialize([]byte(CodeCurrency), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Currency : %s", err)
	}
	_, ok = assetCurrency.(*Currency)
	if !ok {
		t.Fatalf("Failed deserialize type check for Currency")
	}

	// ShareCommon identifies a payload as a ShareCommon message.
	assetShareCommon, err := Deserialize([]byte(CodeShareCommon), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ShareCommon : %s", err)
	}
	_, ok = assetShareCommon.(*ShareCommon)
	if !ok {
		t.Fatalf("Failed deserialize type check for ShareCommon")
	}

	// Coupon identifies a payload as a Coupon message.
	assetCoupon, err := Deserialize([]byte(CodeCoupon), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Coupon : %s", err)
	}
	_, ok = assetCoupon.(*Coupon)
	if !ok {
		t.Fatalf("Failed deserialize type check for Coupon")
	}

	// LoyaltyPoints identifies a payload as a LoyaltyPoints message.
	assetLoyaltyPoints, err := Deserialize([]byte(CodeLoyaltyPoints), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for LoyaltyPoints : %s", err)
	}
	_, ok = assetLoyaltyPoints.(*LoyaltyPoints)
	if !ok {
		t.Fatalf("Failed deserialize type check for LoyaltyPoints")
	}

	// TicketAdmission identifies a payload as a TicketAdmission message.
	assetTicketAdmission, err := Deserialize([]byte(CodeTicketAdmission), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for TicketAdmission : %s", err)
	}
	_, ok = assetTicketAdmission.(*TicketAdmission)
	if !ok {
		t.Fatalf("Failed deserialize type check for TicketAdmission")
	}

}
