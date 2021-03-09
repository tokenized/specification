package assets

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

	// Field InterestPaymentDateDeltas - uint
	if len(a.InterestPaymentDateDeltas) > max2ByteInteger {
		return fmt.Errorf("InterestPaymentDateDeltas list over max length : %d > %d", len(a.InterestPaymentDateDeltas), max2ByteInteger)
	}

	// Field LatePaymentWindow - uint

	// Field LatePaymentPenaltyRate - Rate
	if err := a.LatePaymentPenaltyRate.Validate(); err != nil {
		return errors.Wrap(err, "LatePaymentPenaltyRate")
	}

	// Field LatePaymentPenaltyPeriod - uint

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

	// Field IssueDate - uint

	// Field ExpiryDate - uint

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	// Field TransfersPermitted - bool

	// Field Value - CurrencyValue
	if err := a.Value.Validate(); err != nil {
		return errors.Wrap(err, "Value")
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

	// Field OfferName - varchar
	if len(a.OfferName) > max1ByteInteger {
		return fmt.Errorf("OfferName over max size : %d > %d", len(a.OfferName), max1ByteInteger)
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

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

func (a *TicketAdmission) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field AdmissionType - fixedchar
	if len(a.AdmissionType) != 0 && len(a.AdmissionType) != 3 {
		return fmt.Errorf("AdmissionType fixed width field wrong size : %d should be %d",
			len(a.AdmissionType), 3)
	}

	// Field Venue - varchar
	if len(a.Venue) > max1ByteInteger {
		return fmt.Errorf("Venue over max size : %d > %d", len(a.Venue), max1ByteInteger)
	}

	// Field Class - varchar
	if len(a.Class) > max1ByteInteger {
		return fmt.Errorf("Class over max size : %d > %d", len(a.Class), max1ByteInteger)
	}

	// Field Area - varchar
	if len(a.Area) > max1ByteInteger {
		return fmt.Errorf("Area over max size : %d > %d", len(a.Area), max1ByteInteger)
	}

	// Field Seat - varchar
	if len(a.Seat) > max1ByteInteger {
		return fmt.Errorf("Seat over max size : %d > %d", len(a.Seat), max1ByteInteger)
	}

	// Field StartTimeDate - uint

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

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

func (a *CasinoChip) Validate() error {
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

	// Field UseType - fixedchar
	if len(a.UseType) != 0 && len(a.UseType) != 1 {
		return fmt.Errorf("UseType fixed width field wrong size : %d should be %d",
			len(a.UseType), 1)
	}
	foundUseType := false
	for _, o := range []string{"R", "S", "F"} {
		if a.UseType == o {
			foundUseType = true
			break
		}
	}
	if !foundUseType {
		return fmt.Errorf("UseType value not within options [R S F] : %s", a.UseType)
	}
	if len(a.UseType) == 0 {
		return fmt.Errorf("UseType required")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Precision - uint
	if a.Precision == 0 {
		return fmt.Errorf("Precision required")
	}

	// Field TransfersPermitted - bool

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
