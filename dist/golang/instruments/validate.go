package instruments

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/pkg/errors"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295

	maxArticleDepth = 4
)

func (a *Membership) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field ID - varchar
	if len(a.ID) > max1ByteInteger {
		return fmt.Errorf("ID over max size : %d > %d", len(a.ID), max1ByteInteger)
	}

	// Field MembershipClass - varchar
	if len(a.MembershipClass) > max1ByteInteger {
		return fmt.Errorf("MembershipClass over max size : %d > %d", len(a.MembershipClass), max1ByteInteger)
	}

	// Field RoleType - varchar
	if len(a.RoleType) > max1ByteInteger {
		return fmt.Errorf("RoleType over max size : %d > %d", len(a.RoleType), max1ByteInteger)
	}

	// Field MembershipType - varchar
	if len(a.MembershipType) > max1ByteInteger {
		return fmt.Errorf("MembershipType over max size : %d > %d", len(a.MembershipType), max1ByteInteger)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	// Field TransfersPermitted - bool

	return nil
}

func (a *Currency) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field CurrencyCode - fixedchar  (Currencies Resource)
	if CurrenciesData(a.CurrencyCode) == nil {
		return fmt.Errorf("CurrencyCode resource Currencies value not defined : %v", a.CurrencyCode)
	}
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("CurrencyCode fixed width field wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}
	if len(a.CurrencyCode) == 0 {
		return fmt.Errorf("CurrencyCode required")
	}

	// Field MonetaryAuthority - varchar
	if len(a.MonetaryAuthority) > max1ByteInteger {
		return fmt.Errorf("MonetaryAuthority over max size : %d > %d", len(a.MonetaryAuthority), max1ByteInteger)
	}

	// Field Precision - uint
	if a.Precision == 0 {
		return fmt.Errorf("Precision required")
	}

	return nil
}

func (a *ShareCommon) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Ticker - fixedchar
	if len(a.Ticker) != 0 && len(a.Ticker) != 5 {
		return fmt.Errorf("Ticker fixed width field wrong size : %d should be %d",
			len(a.Ticker), 5)
	}

	// Field ISIN - fixedchar
	if len(a.ISIN) != 0 && len(a.ISIN) != 12 {
		return fmt.Errorf("ISIN fixed width field wrong size : %d should be %d",
			len(a.ISIN), 12)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	// Field TransfersPermitted - bool

	return nil
}

func (a *BondFixedRate) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	InterestRateFieldIsEmpty := a.InterestRate == nil

	LatePaymentPenaltyRateFieldIsEmpty := a.LatePaymentPenaltyRate == nil

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field BondType - fixedchar
	if len(a.BondType) != 0 && len(a.BondType) != 1 {
		return fmt.Errorf("BondType fixed width field wrong size : %d should be %d",
			len(a.BondType), 1)
	}
	foundBondType := false
	for _, o := range []string{"C", "M", "G"} {
		if a.BondType == o {
			foundBondType = true
			break
		}
	}
	if !foundBondType {
		return fmt.Errorf("BondType value not within options [C M G] : %s", a.BondType)
	}
	if len(a.BondType) == 0 {
		return fmt.Errorf("BondType required")
	}

	// Field ISIN - varchar
	if len(a.ISIN) > max1ByteInteger {
		return fmt.Errorf("ISIN over max size : %d > %d", len(a.ISIN), max1ByteInteger)
	}

	// Field Collateral - varchar
	if len(a.Collateral) > max2ByteInteger {
		return fmt.Errorf("Collateral over max size : %d > %d", len(a.Collateral), max2ByteInteger)
	}

	// Field ParValue - CurrencyValue
	if err := a.ParValue.Validate(); err != nil {
		return errors.Wrap(err, "ParValue")
	}
	if a.ParValue == nil {
		return fmt.Errorf("ParValue required")
	}

	// Field InterestRate - Rate
	if err := a.InterestRate.Validate(); err != nil {
		return errors.Wrap(err, "InterestRate")
	}

	// Field InterestPaymentInitialDate - uint
	if InterestRateFieldIsEmpty && a.InterestPaymentInitialDate != 0 {
		return fmt.Errorf("InterestPaymentInitialDate is only allowed when InterestRate is specified : %v", a.InterestRate)
	}
	if !InterestRateFieldIsEmpty && a.InterestPaymentInitialDate == 0 {
		return fmt.Errorf("InterestPaymentInitialDate is required when InterestRate is specified : %v", a.InterestRate)
	}

	// Field InterestPaymentDateDeltas - uint
	if len(a.InterestPaymentDateDeltas) > max2ByteInteger {
		return fmt.Errorf("InterestPaymentDateDeltas list over max length : %d > %d", len(a.InterestPaymentDateDeltas), max2ByteInteger)
	}
	if InterestRateFieldIsEmpty && len(a.InterestPaymentDateDeltas) != 0 {
		return fmt.Errorf("InterestPaymentDateDeltas is only allowed when InterestRate is specified : %v", a.InterestRate)
	}
	if !InterestRateFieldIsEmpty && len(a.InterestPaymentDateDeltas) == 0 {
		return fmt.Errorf("InterestPaymentDateDeltas is required when InterestRate is specified : %v", a.InterestRate)
	}

	// Field LatePaymentPenaltyRate - Rate
	if err := a.LatePaymentPenaltyRate.Validate(); err != nil {
		return errors.Wrap(err, "LatePaymentPenaltyRate")
	}

	// Field LatePaymentWindow - uint
	if LatePaymentPenaltyRateFieldIsEmpty && a.LatePaymentWindow != 0 {
		return fmt.Errorf("LatePaymentWindow is only allowed when LatePaymentPenaltyRate is specified : %v", a.LatePaymentPenaltyRate)
	}

	// Field LatePaymentPenaltyPeriod - uint
	if LatePaymentPenaltyRateFieldIsEmpty && a.LatePaymentPenaltyPeriod != 0 {
		return fmt.Errorf("LatePaymentPenaltyPeriod is only allowed when LatePaymentPenaltyRate is specified : %v", a.LatePaymentPenaltyRate)
	}

	// Field MaturityDate - uint
	if a.MaturityDate == 0 {
		return fmt.Errorf("MaturityDate required")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field TransfersPermitted - bool

	return nil
}

func (a *Coupon) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field RedeemingEntity - varchar
	if len(a.RedeemingEntity) > max1ByteInteger {
		return fmt.Errorf("RedeemingEntity over max size : %d > %d", len(a.RedeemingEntity), max1ByteInteger)
	}

	// Field ValidFromTimestamp - uint

	// Field ExpirationTimestamp - uint

	// Field CouponName - varchar
	if len(a.CouponName) > max1ByteInteger {
		return fmt.Errorf("CouponName over max size : %d > %d", len(a.CouponName), max1ByteInteger)
	}
	if len(a.CouponName) == 0 {
		return fmt.Errorf("CouponName required")
	}

	// Field TransfersPermitted - bool

	// Field FaceValue - CurrencyValue
	if err := a.FaceValue.Validate(); err != nil {
		return errors.Wrap(err, "FaceValue")
	}

	// Field RedemptionVenue - varchar
	if len(a.RedemptionVenue) > max1ByteInteger {
		return fmt.Errorf("RedemptionVenue over max size : %d > %d", len(a.RedemptionVenue), max1ByteInteger)
	}

	// Field Details - varchar
	if len(a.Details) > max2ByteInteger {
		return fmt.Errorf("Details over max size : %d > %d", len(a.Details), max2ByteInteger)
	}

	return nil
}

func (a *LoyaltyPoints) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ProgramName - varchar
	if len(a.ProgramName) > max1ByteInteger {
		return fmt.Errorf("ProgramName over max size : %d > %d", len(a.ProgramName), max1ByteInteger)
	}
	if len(a.ProgramName) == 0 {
		return fmt.Errorf("ProgramName required")
	}

	// Field ExpirationTimestamp - uint

	// Field Details - varchar
	if len(a.Details) > max2ByteInteger {
		return fmt.Errorf("Details over max size : %d > %d", len(a.Details), max2ByteInteger)
	}

	// Field TransfersPermitted - bool

	return nil
}

func (a *TicketAdmission) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field Venue - varchar
	if len(a.Venue) > max1ByteInteger {
		return fmt.Errorf("Venue over max size : %d > %d", len(a.Venue), max1ByteInteger)
	}

	// Field Area - varchar
	if len(a.Area) > max1ByteInteger {
		return fmt.Errorf("Area over max size : %d > %d", len(a.Area), max1ByteInteger)
	}

	// Field Seat - varchar
	if len(a.Seat) > max1ByteInteger {
		return fmt.Errorf("Seat over max size : %d > %d", len(a.Seat), max1ByteInteger)
	}

	// Field EventStartTimestamp - uint

	// Field EventName - varchar
	if len(a.EventName) > max1ByteInteger {
		return fmt.Errorf("EventName over max size : %d > %d", len(a.EventName), max1ByteInteger)
	}
	if len(a.EventName) == 0 {
		return fmt.Errorf("EventName required")
	}

	// Field TransfersPermitted - bool

	// Field Details - varchar
	if len(a.Details) > max2ByteInteger {
		return fmt.Errorf("Details over max size : %d > %d", len(a.Details), max2ByteInteger)
	}

	// Field Section - varchar
	if len(a.Section) > max1ByteInteger {
		return fmt.Errorf("Section over max size : %d > %d", len(a.Section), max1ByteInteger)
	}

	// Field Row - varchar
	if len(a.Row) > max1ByteInteger {
		return fmt.Errorf("Row over max size : %d > %d", len(a.Row), max1ByteInteger)
	}

	// Field EventEndTimestamp - uint

	return nil
}

func (a *CasinoChip) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field UseType - fixedchar
	if len(a.UseType) != 0 && len(a.UseType) != 1 {
		return fmt.Errorf("UseType fixed width field wrong size : %d should be %d",
			len(a.UseType), 1)
	}
	foundUseType := false
	for _, o := range []string{"R", "F"} {
		if a.UseType == o {
			foundUseType = true
			break
		}
	}
	if !foundUseType {
		return fmt.Errorf("UseType value not within options [R F] : %s", a.UseType)
	}
	if len(a.UseType) == 0 {
		return fmt.Errorf("UseType required")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ExpirationTimestamp - uint

	// Field TransfersPermitted - bool

	// Field CasinoName - varchar
	if len(a.CasinoName) > max1ByteInteger {
		return fmt.Errorf("CasinoName over max size : %d > %d", len(a.CasinoName), max1ByteInteger)
	}
	if len(a.CasinoName) == 0 {
		return fmt.Errorf("CasinoName required")
	}

	// Field FaceValue - CurrencyValue
	if err := a.FaceValue.Validate(); err != nil {
		return errors.Wrap(err, "FaceValue")
	}
	if a.FaceValue == nil {
		return fmt.Errorf("FaceValue required")
	}

	return nil
}

func (a *InformationServiceLicense) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ExpirationTimestamp - uint

	// Field ServiceName - varchar
	if len(a.ServiceName) > max1ByteInteger {
		return fmt.Errorf("ServiceName over max size : %d > %d", len(a.ServiceName), max1ByteInteger)
	}
	if len(a.ServiceName) == 0 {
		return fmt.Errorf("ServiceName required")
	}

	// Field TransfersPermitted - bool

	// Field URL - varchar
	if len(a.URL) > max2ByteInteger {
		return fmt.Errorf("URL over max size : %d > %d", len(a.URL), max2ByteInteger)
	}

	return nil
}

func (a *CreditNote) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field FaceValue - CurrencyValue
	if err := a.FaceValue.Validate(); err != nil {
		return errors.Wrap(err, "FaceValue")
	}

	// Field ExpirationTimestamp - uint

	return nil
}

func (a *AgeRestrictionField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Lower - uint
	if a.Lower > uint32(max1ByteInteger) {
		return fmt.Errorf("Lower over max value : %d > %d", a.Lower, max1ByteInteger)
	}

	// Field Upper - uint
	if a.Upper > uint32(max1ByteInteger) {
		return fmt.Errorf("Upper over max value : %d > %d", a.Upper, max1ByteInteger)
	}

	return nil
}

func (a *CurrencyValueField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Value - uint
	if a.Value == 0 {
		return fmt.Errorf("Value required")
	}

	// Field CurrencyCode - fixedchar  (Currencies Resource)
	if CurrenciesData(a.CurrencyCode) == nil {
		return fmt.Errorf("CurrencyCode resource Currencies value not defined : %v", a.CurrencyCode)
	}
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("CurrencyCode fixed width field wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}
	if len(a.CurrencyCode) == 0 {
		return fmt.Errorf("CurrencyCode required")
	}

	// Field Precision - uint
	if a.Precision > uint32(max1ByteInteger) {
		return fmt.Errorf("Precision over max value : %d > %d", a.Precision, max1ByteInteger)
	}
	if a.Precision == 0 {
		return fmt.Errorf("Precision required")
	}

	return nil
}

func (a *RateField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Precision - uint
	if a.Precision > uint32(max1ByteInteger) {
		return fmt.Errorf("Precision over max value : %d > %d", a.Precision, max1ByteInteger)
	}

	// Field Value - uint

	return nil
}

// AddressIsValid returns true if an "Address" alias field is valid.
func AddressIsValid(b []byte) error {
	_, err := bitcoin.DecodeRawAddress(b)
	return err
}

// PublicKeyIsValid returns true if a "PublicKey" alias field is valid.
func PublicKeyIsValid(b []byte) error {
	_, err := bitcoin.PublicKeyFromBytes(b)
	return err
}

// SignatureIsValid returns true if a "Signature" alias field is valid.
func SignatureIsValid(b []byte) error {
	_, err := bitcoin.SignatureFromBytes(b)
	return err
}
