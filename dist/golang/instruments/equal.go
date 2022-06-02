package instruments

func (l *Membership) Equal(right interface{}) bool {
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

func (l *Currency) Equal(right interface{}) bool {
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

func (l *ShareCommon) Equal(right interface{}) bool {
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

func (l *BondFixedRate) Equal(right interface{}) bool {
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

	// Field BondType - fixedchar
	if c.BondType != r.BondType {
		return false // fmt.Errorf("BondType string mismatched")
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

	// Field LatePaymentPenaltyRate - Rate
	if !c.LatePaymentPenaltyRate.Equal(r.LatePaymentPenaltyRate) {
		return false // fmt.Errorf("LatePaymentPenaltyRate : %s", err)
	}

	// Field LatePaymentWindow - uint
	if c.LatePaymentWindow != r.LatePaymentWindow {
		return false // fmt.Errorf("LatePaymentWindow integer mismatched")
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

func (l *Coupon) Equal(right interface{}) bool {
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

	// Field ValidFromTimestamp - uint
	if c.ValidFromTimestamp != r.ValidFromTimestamp {
		return false // fmt.Errorf("ValidFromTimestamp integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field CouponName - varchar
	if c.CouponName != r.CouponName {
		return false // fmt.Errorf("CouponName string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field FaceValue - CurrencyValue
	if !c.FaceValue.Equal(r.FaceValue) {
		return false // fmt.Errorf("FaceValue : %s", err)
	}

	// Field RedemptionVenue - varchar
	if c.RedemptionVenue != r.RedemptionVenue {
		return false // fmt.Errorf("RedemptionVenue string mismatched")
	}

	// Field Details - varchar
	if c.Details != r.Details {
		return false // fmt.Errorf("Details string mismatched")
	}

	return true
}

func (l *LoyaltyPoints) Equal(right interface{}) bool {
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

	// Field ProgramName - varchar
	if c.ProgramName != r.ProgramName {
		return false // fmt.Errorf("ProgramName string mismatched")
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Details - varchar
	if c.Details != r.Details {
		return false // fmt.Errorf("Details string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	return true
}

func (l *TicketAdmission) Equal(right interface{}) bool {
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

	// Field Venue - varchar
	if c.Venue != r.Venue {
		return false // fmt.Errorf("Venue string mismatched")
	}

	// Field Area - varchar
	if c.Area != r.Area {
		return false // fmt.Errorf("Area string mismatched")
	}

	// Field Seat - varchar
	if c.Seat != r.Seat {
		return false // fmt.Errorf("Seat string mismatched")
	}

	// Field EventStartTimestamp - uint
	if c.EventStartTimestamp != r.EventStartTimestamp {
		return false // fmt.Errorf("EventStartTimestamp integer mismatched")
	}

	// Field EventName - varchar
	if c.EventName != r.EventName {
		return false // fmt.Errorf("EventName string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field Details - varchar
	if c.Details != r.Details {
		return false // fmt.Errorf("Details string mismatched")
	}

	// Field Section - varchar
	if c.Section != r.Section {
		return false // fmt.Errorf("Section string mismatched")
	}

	// Field Row - varchar
	if c.Row != r.Row {
		return false // fmt.Errorf("Row string mismatched")
	}

	// Field EventEndTimestamp - uint
	if c.EventEndTimestamp != r.EventEndTimestamp {
		return false // fmt.Errorf("EventEndTimestamp integer mismatched")
	}

	return true
}

func (l *CasinoChip) Equal(right interface{}) bool {
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

	// Field UseType - fixedchar
	if c.UseType != r.UseType {
		return false // fmt.Errorf("UseType string mismatched")
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field CasinoName - varchar
	if c.CasinoName != r.CasinoName {
		return false // fmt.Errorf("CasinoName string mismatched")
	}

	// Field FaceValue - CurrencyValue
	if !c.FaceValue.Equal(r.FaceValue) {
		return false // fmt.Errorf("FaceValue : %s", err)
	}

	return true
}

func (l *InformationServiceLicense) Equal(right interface{}) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InformationServiceLicense{}
	}
	cr := right
	if cr == nil {
		cr = &InformationServiceLicense{}
	}
	r, ok := cr.(*InformationServiceLicense)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction
	if !c.AgeRestriction.Equal(r.AgeRestriction) {
		return false // fmt.Errorf("AgeRestriction : %s", err)
	}

	// Field ExpirationTimestamp - uint
	if c.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field ServiceName - varchar
	if c.ServiceName != r.ServiceName {
		return false // fmt.Errorf("ServiceName string mismatched")
	}

	// Field TransfersPermitted - bool
	if c.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field URL - varchar
	if c.URL != r.URL {
		return false // fmt.Errorf("URL string mismatched")
	}

	return true
}

func (l *AgeRestrictionField) Equal(right interface{}) bool {
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

func (l *CurrencyValueField) Equal(right interface{}) bool {
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

func (l *RateField) Equal(right interface{}) bool {
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
