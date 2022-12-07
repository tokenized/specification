package instruments

import (
	"testing"
)

func TestEmptyDeserialize(t *testing.T) {
	var ok bool
	// Membership identifies a payload as a Membership message.
	instrumentMembership, err := Deserialize([]byte(CodeMembership), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Membership : %s", err)
	}
	_, ok = instrumentMembership.(*Membership)
	if !ok {
		t.Fatalf("Failed deserialize type check for Membership")
	}

	// Currency identifies a payload as a Currency message.
	instrumentCurrency, err := Deserialize([]byte(CodeCurrency), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Currency : %s", err)
	}
	_, ok = instrumentCurrency.(*Currency)
	if !ok {
		t.Fatalf("Failed deserialize type check for Currency")
	}

	// ShareCommon identifies a payload as a ShareCommon message.
	instrumentShareCommon, err := Deserialize([]byte(CodeShareCommon), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for ShareCommon : %s", err)
	}
	_, ok = instrumentShareCommon.(*ShareCommon)
	if !ok {
		t.Fatalf("Failed deserialize type check for ShareCommon")
	}

	// BondFixedRate identifies a payload as a BondFixedRate message.
	instrumentBondFixedRate, err := Deserialize([]byte(CodeBondFixedRate), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for BondFixedRate : %s", err)
	}
	_, ok = instrumentBondFixedRate.(*BondFixedRate)
	if !ok {
		t.Fatalf("Failed deserialize type check for BondFixedRate")
	}

	// Coupon identifies a payload as a Coupon message.
	instrumentCoupon, err := Deserialize([]byte(CodeCoupon), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for Coupon : %s", err)
	}
	_, ok = instrumentCoupon.(*Coupon)
	if !ok {
		t.Fatalf("Failed deserialize type check for Coupon")
	}

	// LoyaltyPoints identifies a payload as a LoyaltyPoints message.
	instrumentLoyaltyPoints, err := Deserialize([]byte(CodeLoyaltyPoints), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for LoyaltyPoints : %s", err)
	}
	_, ok = instrumentLoyaltyPoints.(*LoyaltyPoints)
	if !ok {
		t.Fatalf("Failed deserialize type check for LoyaltyPoints")
	}

	// TicketAdmission identifies a payload as a TicketAdmission message.
	instrumentTicketAdmission, err := Deserialize([]byte(CodeTicketAdmission), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for TicketAdmission : %s", err)
	}
	_, ok = instrumentTicketAdmission.(*TicketAdmission)
	if !ok {
		t.Fatalf("Failed deserialize type check for TicketAdmission")
	}

	// CasinoChip identifies a payload as a CasinoChip message.
	instrumentCasinoChip, err := Deserialize([]byte(CodeCasinoChip), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for CasinoChip : %s", err)
	}
	_, ok = instrumentCasinoChip.(*CasinoChip)
	if !ok {
		t.Fatalf("Failed deserialize type check for CasinoChip")
	}

	// InformationServiceLicense identifies a payload as a InformationServiceLicense message.
	instrumentInformationServiceLicense, err := Deserialize([]byte(CodeInformationServiceLicense), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for InformationServiceLicense : %s", err)
	}
	_, ok = instrumentInformationServiceLicense.(*InformationServiceLicense)
	if !ok {
		t.Fatalf("Failed deserialize type check for InformationServiceLicense")
	}

	// CreditNote identifies a payload as a CreditNote message.
	instrumentCreditNote, err := Deserialize([]byte(CodeCreditNote), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for CreditNote : %s", err)
	}
	_, ok = instrumentCreditNote.(*CreditNote)
	if !ok {
		t.Fatalf("Failed deserialize type check for CreditNote")
	}

}
