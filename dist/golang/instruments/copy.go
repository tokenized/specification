package instruments

func (a *Membership) Copy() *Membership {
	result := &Membership{}

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field ValidFrom - uint
	result.ValidFrom = a.ValidFrom

	// Field ExpirationTimestamp - uint
	result.ExpirationTimestamp = a.ExpirationTimestamp

	// Field ID - varchar
	result.ID = a.ID

	// Field MembershipClass - varchar
	result.MembershipClass = a.MembershipClass

	// Field RoleType - varchar
	result.RoleType = a.RoleType

	// Field MembershipType - varchar
	result.MembershipType = a.MembershipType

	// Field Description - varchar
	result.Description = a.Description

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	return result
}

func (a *Currency) Copy() *Currency {
	result := &Currency{}

	// Field CurrencyCode - fixedchar
	result.CurrencyCode = a.CurrencyCode

	// Field MonetaryAuthority - varchar
	result.MonetaryAuthority = a.MonetaryAuthority

	// Field Precision - uint
	result.Precision = a.Precision

	return result
}

func (a *ShareCommon) Copy() *ShareCommon {
	result := &ShareCommon{}

	// Field Ticker - fixedchar
	result.Ticker = a.Ticker

	// Field ISIN - fixedchar
	result.ISIN = a.ISIN

	// Field Description - varchar
	result.Description = a.Description

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	return result
}

func (a *BondFixedRate) Copy() *BondFixedRate {
	result := &BondFixedRate{}

	// Field Name - varchar
	result.Name = a.Name

	// Field BondType - fixedchar
	result.BondType = a.BondType

	// Field ISIN - varchar
	result.ISIN = a.ISIN

	// Field Collateral - varchar
	result.Collateral = a.Collateral

	// Field ParValue - CurrencyValue
	result.ParValue = a.ParValue.Copy()

	// Field InterestRate - Rate
	result.InterestRate = a.InterestRate.Copy()

	// Field InterestPaymentInitialDate - uint
	result.InterestPaymentInitialDate = a.InterestPaymentInitialDate

	// Field InterestPaymentDateDeltas - uint
	result.InterestPaymentDateDeltas = make([]uint64, len(a.InterestPaymentDateDeltas))
	for i, v := range a.InterestPaymentDateDeltas {
		result.InterestPaymentDateDeltas[i] = v
	}

	// Field LatePaymentPenaltyRate - Rate
	result.LatePaymentPenaltyRate = a.LatePaymentPenaltyRate.Copy()

	// Field LatePaymentWindow - uint
	result.LatePaymentWindow = a.LatePaymentWindow

	// Field LatePaymentPenaltyPeriod - uint
	result.LatePaymentPenaltyPeriod = a.LatePaymentPenaltyPeriod

	// Field MaturityDate - uint
	result.MaturityDate = a.MaturityDate

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	return result
}

func (a *Coupon) Copy() *Coupon {
	result := &Coupon{}

	// Field RedeemingEntity - varchar
	result.RedeemingEntity = a.RedeemingEntity

	// Field ValidFromTimestamp - uint
	result.ValidFromTimestamp = a.ValidFromTimestamp

	// Field ExpirationTimestamp - uint
	result.ExpirationTimestamp = a.ExpirationTimestamp

	// Field CouponName - varchar
	result.CouponName = a.CouponName

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	// Field FaceValue - CurrencyValue
	result.FaceValue = a.FaceValue.Copy()

	// Field RedemptionVenue - varchar
	result.RedemptionVenue = a.RedemptionVenue

	// Field Details - varchar
	result.Details = a.Details

	return result
}

func (a *LoyaltyPoints) Copy() *LoyaltyPoints {
	result := &LoyaltyPoints{}

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field ProgramName - varchar
	result.ProgramName = a.ProgramName

	// Field ExpirationTimestamp - uint
	result.ExpirationTimestamp = a.ExpirationTimestamp

	// Field Details - varchar
	result.Details = a.Details

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	return result
}

func (a *TicketAdmission) Copy() *TicketAdmission {
	result := &TicketAdmission{}

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field Venue - varchar
	result.Venue = a.Venue

	// Field Area - varchar
	result.Area = a.Area

	// Field Seat - varchar
	result.Seat = a.Seat

	// Field EventStartTimestamp - uint
	result.EventStartTimestamp = a.EventStartTimestamp

	// Field EventName - varchar
	result.EventName = a.EventName

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	// Field Details - varchar
	result.Details = a.Details

	// Field Section - varchar
	result.Section = a.Section

	// Field Row - varchar
	result.Row = a.Row

	// Field EventEndTimestamp - uint
	result.EventEndTimestamp = a.EventEndTimestamp

	return result
}

func (a *CasinoChip) Copy() *CasinoChip {
	result := &CasinoChip{}

	// Field UseType - fixedchar
	result.UseType = a.UseType

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field ExpirationTimestamp - uint
	result.ExpirationTimestamp = a.ExpirationTimestamp

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	// Field CasinoName - varchar
	result.CasinoName = a.CasinoName

	// Field FaceValue - CurrencyValue
	result.FaceValue = a.FaceValue.Copy()

	return result
}

func (a *InformationServiceLicense) Copy() *InformationServiceLicense {
	result := &InformationServiceLicense{}

	// Field AgeRestriction - AgeRestriction
	result.AgeRestriction = a.AgeRestriction.Copy()

	// Field ExpirationTimestamp - uint
	result.ExpirationTimestamp = a.ExpirationTimestamp

	// Field ServiceName - varchar
	result.ServiceName = a.ServiceName

	// Field TransfersPermitted - bool
	result.TransfersPermitted = a.TransfersPermitted

	// Field URL - varchar
	result.URL = a.URL

	return result
}

func (a *AgeRestrictionField) Copy() *AgeRestrictionField {
	result := &AgeRestrictionField{}

	// Field Lower - uint
	result.Lower = a.Lower

	// Field Upper - uint
	result.Upper = a.Upper

	return result
}

func (a *CurrencyValueField) Copy() *CurrencyValueField {
	result := &CurrencyValueField{}

	// Field Value - uint
	result.Value = a.Value

	// Field CurrencyCode - fixedchar
	result.CurrencyCode = a.CurrencyCode

	// Field Precision - uint
	result.Precision = a.Precision

	return result
}

func (a *RateField) Copy() *RateField {
	result := &RateField{}

	// Field Precision - uint
	result.Precision = a.Precision

	// Field Value - uint
	result.Value = a.Value

	return result
}
