package assets

import (
	proto "github.com/golang/protobuf/proto"
)

func (l *Membership) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Membership{}
	}
	cr := right
	if cr == nil {
		cr = &Membership{}
	}
	r, ok := cr.(*Membership)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field ValidFrom - uint
	if c.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field ID - varchar
	if c.ID != r.ID {
		return false // fmt.Errorf("ID string mismatched")
	}

	// Field MembershipClass - varchar
	if c.MembershipClass != r.MembershipClass {
		return false // fmt.Errorf("MembershipClass string mismatched")
	}

	// Field RoleType - varchar
	if c.RoleType != r.RoleType {
		return false // fmt.Errorf("RoleType string mismatched")
	}

	// Field MembershipType - varchar
	if c.MembershipType != r.MembershipType {
		return false // fmt.Errorf("MembershipType string mismatched")
	}

	// Field Description - varchar
	if c.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *Currency) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Currency{}
	}
	cr := right
	if cr == nil {
		cr = &Currency{}
	}
	r, ok := cr.(*Currency)
	if !ok {
		return false
	}

	// Field CurrencyCode - fixedchar
	if c.CurrencyCode != r.CurrencyCode {
		return false // fmt.Errorf("CurrencyCode string mismatched")
	}

	// Field MonetaryAuthority - varchar
	if c.MonetaryAuthority != r.MonetaryAuthority {
		return false // fmt.Errorf("MonetaryAuthority string mismatched")
	}

	// Field Precision - uint
	if c.Precision != r.Precision {
		return false // fmt.Errorf("Precision integer mismatched")
	}

	return true
}

func (l *ShareCommon) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ShareCommon{}
	}
	cr := right
	if cr == nil {
		cr = &ShareCommon{}
	}
	r, ok := cr.(*ShareCommon)
	if !ok {
		return false
	}

	// Field Ticker - fixedchar
	if c.Ticker != r.Ticker {
		return false // fmt.Errorf("Ticker string mismatched")
	}

	// Field ISIN - fixedchar
	if c.ISIN != r.ISIN {
		return false // fmt.Errorf("ISIN string mismatched")
	}

	// Field Description - varchar
	if c.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *BondFixedRate) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BondFixedRate{}
	}
	cr := right
	if cr == nil {
		cr = &BondFixedRate{}
	}
	r, ok := cr.(*BondFixedRate)
	if !ok {
		return false
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - fixedchar
	if c.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field ISIN - varchar
	if c.ISIN != r.ISIN {
		return false // fmt.Errorf("ISIN string mismatched")
	}

	// Field Collateral - varchar
	if c.Collateral != r.Collateral {
		return false // fmt.Errorf("Collateral string mismatched")
	}

	// Field ParValue - CurrencyValue
	if !c.ParValue.Equal(r.ParValue) {
		return false // fmt.Errorf("ParValue : %s", err)
	}

	// Field InterestRate - Rate
	if !c.InterestRate.Equal(r.InterestRate) {
		return false // fmt.Errorf("InterestRate : %s", err)
	}

	// Field InterestPaymentInitialDate - uint
	if c.InterestPaymentInitialDate != r.InterestPaymentInitialDate {
		return false // fmt.Errorf("InterestPaymentInitialDate integer mismatched")
	}

	// Field InterestPaymentDateDeltas - uint
	if len(c.InterestPaymentDateDeltas) != len(r.InterestPaymentDateDeltas) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.InterestPaymentDateDeltas {
		if v != r.InterestPaymentDateDeltas[i] {
			return false // fmt.Errorf("Element InterestPaymentDateDeltas integer mismatched")
		}
	}

	// Field LatePaymentWindow - uint
	if c.LatePaymentWindow != r.LatePaymentWindow {
		return false // fmt.Errorf("LatePaymentWindow integer mismatched")
	}

	// Field LatePaymentPenaltyRate - Rate
	if !c.LatePaymentPenaltyRate.Equal(r.LatePaymentPenaltyRate) {
		return false // fmt.Errorf("LatePaymentPenaltyRate : %s", err)
	}

	// Field LatePaymentPenaltyPeriod - uint
	if c.LatePaymentPenaltyPeriod != r.LatePaymentPenaltyPeriod {
		return false // fmt.Errorf("LatePaymentPenaltyPeriod integer mismatched")
	}

	// Field MaturityDate - uint
	if c.MaturityDate != r.MaturityDate {
		return false // fmt.Errorf("MaturityDate integer mismatched")
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *Coupon) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Coupon{}
	}
	cr := right
	if cr == nil {
		cr = &Coupon{}
	}
	r, ok := cr.(*Coupon)
	if !ok {
		return false
	}

	// Field RedeemingEntity - varchar
	if c.RedeemingEntity != r.RedeemingEntity {
		return false // fmt.Errorf("RedeemingEntity string mismatched")
	}

	// Field IssueDate - uint
	if c.IssueDate != r.IssueDate {
		return false // fmt.Errorf("IssueDate integer mismatched")
	}

	// Field ExpiryDate - uint
	if c.ExpiryDate != r.ExpiryDate {
		return false // fmt.Errorf("ExpiryDate integer mismatched")
	}

	// Field Description - varchar
	if c.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field Value - CurrencyValue
	if !c.Value.Equal(r.Value) {
		return false // fmt.Errorf("Value : %s", err)
	}

	return true
}

func (l *LoyaltyPoints) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &LoyaltyPoints{}
	}
	cr := right
	if cr == nil {
		cr = &LoyaltyPoints{}
	}
	r, ok := cr.(*LoyaltyPoints)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field OfferName - varchar
	if c.OfferName != r.OfferName {
		return false // fmt.Errorf("OfferName string mismatched")
	}

	// Field ValidFrom - uint
	if c.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Description - varchar
	if c.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *TicketAdmission) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &TicketAdmission{}
	}
	cr := right
	if cr == nil {
		cr = &TicketAdmission{}
	}
	r, ok := cr.(*TicketAdmission)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field AdmissionType - fixedchar
	if c.AdmissionType != r.AdmissionType {
		return false // fmt.Errorf("AdmissionType string mismatched")
	}

	// Field Venue - varchar
	if c.Venue != r.Venue {
		return false // fmt.Errorf("Venue string mismatched")
	}

	// Field Class - varchar
	if c.Class != r.Class {
		return false // fmt.Errorf("Class string mismatched")
	}

	// Field Area - varchar
	if c.Area != r.Area {
		return false // fmt.Errorf("Area string mismatched")
	}

	// Field Seat - varchar
	if c.Seat != r.Seat {
		return false // fmt.Errorf("Seat string mismatched")
	}

	// Field StartTimeDate - uint
	if c.StartTimeDate != r.StartTimeDate {
		return false // fmt.Errorf("StartTimeDate integer mismatched")
	}

	// Field ValidFrom - uint
	if c.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Description - varchar
	if c.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *CasinoChip) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &CasinoChip{}
	}
	cr := right
	if cr == nil {
		cr = &CasinoChip{}
	}
	r, ok := cr.(*CasinoChip)
	if !ok {
		return false
	}

	// Field CurrencyCode - fixedchar
	if c.CurrencyCode != r.CurrencyCode {
		return false // fmt.Errorf("CurrencyCode string mismatched")
	}

	// Field UseType - fixedchar
	if c.UseType != r.UseType {
		return false // fmt.Errorf("UseType string mismatched")
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field ValidFrom - uint
	if c.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Precision - uint
	if c.Precision != r.Precision {
		return false // fmt.Errorf("Precision integer mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *AgeRestrictionField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &AgeRestrictionField{}
	}
	r, ok := right.(*AgeRestrictionField)
	if !ok {
		return false
	}

	if r == nil {
		r = &AgeRestrictionField{}
	}

	// Field Lower - uint
	if c.Lower != r.Lower {
		return false // fmt.Errorf("Lower integer mismatched")
	}

	// Field Upper - uint
	if c.Upper != r.Upper {
		return false // fmt.Errorf("Upper integer mismatched")
	}

	return true
}

func (l *CurrencyValueField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &CurrencyValueField{}
	}
	r, ok := right.(*CurrencyValueField)
	if !ok {
		return false
	}

	if r == nil {
		r = &CurrencyValueField{}
	}

	// Field Value - uint
	if c.Value != r.Value {
		return false // fmt.Errorf("Value integer mismatched")
	}

	// Field CurrencyCode - fixedchar
	if c.CurrencyCode != r.CurrencyCode {
		return false // fmt.Errorf("CurrencyCode string mismatched")
	}

	// Field Precision - uint
	if c.Precision != r.Precision {
		return false // fmt.Errorf("Precision integer mismatched")
	}

	return true
}

func (l *RateField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &RateField{}
	}
	r, ok := right.(*RateField)
	if !ok {
		return false
	}

	if r == nil {
		r = &RateField{}
	}

	// Field Precision - uint
	if c.Precision != r.Precision {
		return false // fmt.Errorf("Precision integer mismatched")
	}

	// Field Value - uint
	if c.Value != r.Value {
		return false // fmt.Errorf("Value integer mismatched")
	}

	return true
}
