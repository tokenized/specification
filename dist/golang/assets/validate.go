package assets

import (
	"fmt"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295
)

func (a *Membership) Validate() error {
	if a == nil {
		return nil
	}

	// Field AgeRestriction - AgeRestriction

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field ID - varchar
	if len(a.ID) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ID), max1ByteInteger)
	}

	// Field MembershipClass - varchar
	if len(a.MembershipClass) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MembershipClass), max1ByteInteger)
	}

	// Field RoleType - varchar
	if len(a.RoleType) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.RoleType), max1ByteInteger)
	}

	// Field MembershipType - varchar
	if len(a.MembershipType) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MembershipType), max1ByteInteger)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *Currency) Validate() error {
	if a == nil {
		return nil
	}

	// Field CurrencyCode - fixedchar
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("Fixed width field CurrencyCode wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}

	// Field MonetaryAuthority - varchar
	if len(a.MonetaryAuthority) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MonetaryAuthority), max1ByteInteger)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *ShareCommon) Validate() error {
	if a == nil {
		return nil
	}

	// Field TransferLockout - uint

	// Field Ticker - fixedchar
	if len(a.Ticker) != 0 && len(a.Ticker) != 5 {
		return fmt.Errorf("Fixed width field Ticker wrong size : %d should be %d",
			len(a.Ticker), 5)
	}

	// Field ISIN - fixedchar
	if len(a.ISIN) != 0 && len(a.ISIN) != 12 {
		return fmt.Errorf("Fixed width field ISIN wrong size : %d should be %d",
			len(a.ISIN), 12)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *Coupon) Validate() error {
	if a == nil {
		return nil
	}

	// Field RedeemingEntity - varchar
	if len(a.RedeemingEntity) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.RedeemingEntity), max1ByteInteger)
	}

	// Field IssueDate - uint

	// Field ExpiryDate - uint

	// Field Value - uint

	// Field Currency - fixedchar
	if len(a.Currency) != 0 && len(a.Currency) != 3 {
		return fmt.Errorf("Fixed width field Currency wrong size : %d should be %d",
			len(a.Currency), 3)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *LoyaltyPoints) Validate() error {
	if a == nil {
		return nil
	}

	// Field AgeRestriction - AgeRestriction

	// Field OfferName - varchar
	if len(a.OfferName) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OfferName), max1ByteInteger)
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *TicketAdmission) Validate() error {
	if a == nil {
		return nil
	}

	// Field AgeRestriction - AgeRestriction

	// Field AdmissionType - fixedchar
	if len(a.AdmissionType) != 0 && len(a.AdmissionType) != 3 {
		return fmt.Errorf("Fixed width field AdmissionType wrong size : %d should be %d",
			len(a.AdmissionType), 3)
	}

	// Field Venue - varchar
	if len(a.Venue) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Venue), max1ByteInteger)
	}

	// Field Class - varchar
	if len(a.Class) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Class), max1ByteInteger)
	}

	// Field Area - varchar
	if len(a.Area) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Area), max1ByteInteger)
	}

	// Field Seat - varchar
	if len(a.Seat) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Seat), max1ByteInteger)
	}

	// Field StartTimeDate - uint

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Description), max2ByteInteger)
	}

	return nil
}

func (a *CasinoChip) Validate() error {
	if a == nil {
		return nil
	}

	// Field CurrencyCode - fixedchar
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("Fixed width field CurrencyCode wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}

	// Field UseType - fixedchar
	if len(a.UseType) != 0 && len(a.UseType) != 1 {
		return fmt.Errorf("Fixed width field UseType wrong size : %d should be %d",
			len(a.UseType), 1)
	}

	// Field AgeRestriction - AgeRestriction

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	return nil
}

func (a *AgeRestrictionField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Lower - uint
	if a.Lower > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Lower, max1ByteInteger)
	}

	// Field Upper - uint
	if a.Upper > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Upper, max1ByteInteger)
	}

	return nil
}
