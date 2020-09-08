package assets

import (
	"bytes"
	"fmt"

	"github.com/tokenized/specification/dist/golang/internal"

	"github.com/pkg/errors"
)

// Membership Permission / Amendment Field Indices
const (
	MembershipFieldAgeRestriction      = uint32(1)
	MembershipFieldValidFrom           = uint32(2)
	MembershipFieldExpirationTimestamp = uint32(3)
	MembershipFieldID                  = uint32(4)
	MembershipFieldMembershipClass     = uint32(5)
	MembershipFieldRoleType            = uint32(6)
	MembershipFieldMembershipType      = uint32(7)
	MembershipFieldDescription         = uint32(8)
)

// ApplyAmendment updates a Membership based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Membership) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case MembershipFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case MembershipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return fip[:], nil

	case MembershipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return fip[:], nil

	case MembershipFieldID: // string
		a.ID = string(data)
		return fip[:], nil

	case MembershipFieldMembershipClass: // string
		a.MembershipClass = string(data)
		return fip[:], nil

	case MembershipFieldRoleType: // string
		a.RoleType = string(data)
		return fip[:], nil

	case MembershipFieldMembershipType: // string
		a.MembershipType = string(data)
		return fip[:], nil

	case MembershipFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Membership amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Memberships and returns
// amendment data.
func (a *Membership) CreateAmendments(fip []uint32,
	newValue *Membership) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, MembershipFieldAgeRestriction)
	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ValidFrom uint64
	fip = append(ofip, MembershipFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, MembershipFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ID string
	fip = append(ofip, MembershipFieldID)
	if a.ID != newValue.ID {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ID),
		})
	}

	// MembershipClass string
	fip = append(ofip, MembershipFieldMembershipClass)
	if a.MembershipClass != newValue.MembershipClass {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MembershipClass),
		})
	}

	// RoleType string
	fip = append(ofip, MembershipFieldRoleType)
	if a.RoleType != newValue.RoleType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RoleType),
		})
	}

	// MembershipType string
	fip = append(ofip, MembershipFieldMembershipType)
	if a.MembershipType != newValue.MembershipType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MembershipType),
		})
	}

	// Description string
	fip = append(ofip, MembershipFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	return result, nil
}

// Currency Permission / Amendment Field Indices
const (
	CurrencyFieldCurrencyCode          = uint32(1)
	CurrencyFieldMonetaryAuthority     = uint32(2)
	DeprecatedCurrencyFieldDescription = uint32(3)
	CurrencyFieldPrecision             = uint32(4)
)

// ApplyAmendment updates a Currency based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Currency) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CurrencyFieldCurrencyCode: // string
		a.CurrencyCode = string(data)
		return fip[:], nil

	case CurrencyFieldMonetaryAuthority: // string
		a.MonetaryAuthority = string(data)
		return fip[:], nil

	case DeprecatedCurrencyFieldDescription: // deprecated

	case CurrencyFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Currency amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Currencys and returns
// amendment data.
func (a *Currency) CreateAmendments(fip []uint32,
	newValue *Currency) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// CurrencyCode string
	fip = append(ofip, CurrencyFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// MonetaryAuthority string
	fip = append(ofip, CurrencyFieldMonetaryAuthority)
	if a.MonetaryAuthority != newValue.MonetaryAuthority {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MonetaryAuthority),
		})
	}

	// deprecated Description deprecated

	// Precision uint64
	fip = append(ofip, CurrencyFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// ShareCommon Permission / Amendment Field Indices
const (
	ShareCommonFieldTicker      = uint32(1)
	ShareCommonFieldISIN        = uint32(2)
	ShareCommonFieldDescription = uint32(3)
)

// ApplyAmendment updates a ShareCommon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ShareCommon) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case ShareCommonFieldTicker: // string
		a.Ticker = string(data)
		return fip[:], nil

	case ShareCommonFieldISIN: // string
		a.ISIN = string(data)
		return fip[:], nil

	case ShareCommonFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown ShareCommon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two ShareCommons and returns
// amendment data.
func (a *ShareCommon) CreateAmendments(fip []uint32,
	newValue *ShareCommon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// Ticker string
	fip = append(ofip, ShareCommonFieldTicker)
	if a.Ticker != newValue.Ticker {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Ticker),
		})
	}

	// ISIN string
	fip = append(ofip, ShareCommonFieldISIN)
	if a.ISIN != newValue.ISIN {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ISIN),
		})
	}

	// Description string
	fip = append(ofip, ShareCommonFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	return result, nil
}

// Coupon Permission / Amendment Field Indices
const (
	CouponFieldRedeemingEntity = uint32(1)
	CouponFieldIssueDate       = uint32(2)
	CouponFieldExpiryDate      = uint32(3)
	CouponFieldValue           = uint32(4)
	CouponFieldCurrency        = uint32(5)
	CouponFieldDescription     = uint32(6)
	CouponFieldPrecision       = uint32(7)
)

// ApplyAmendment updates a Coupon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Coupon) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CouponFieldRedeemingEntity: // string
		a.RedeemingEntity = string(data)
		return fip[:], nil

	case CouponFieldIssueDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for IssueDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("IssueDate amendment value failed to deserialize : %s", err)
		} else {
			a.IssueDate = uint64(value)
		}
		return fip[:], nil

	case CouponFieldExpiryDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpiryDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpiryDate amendment value failed to deserialize : %s", err)
		} else {
			a.ExpiryDate = uint64(value)
		}
		return fip[:], nil

	case CouponFieldValue: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return fip[:], nil

	case CouponFieldCurrency: // string
		a.Currency = string(data)
		return fip[:], nil

	case CouponFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	case CouponFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Coupon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Coupons and returns
// amendment data.
func (a *Coupon) CreateAmendments(fip []uint32,
	newValue *Coupon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// RedeemingEntity string
	fip = append(ofip, CouponFieldRedeemingEntity)
	if a.RedeemingEntity != newValue.RedeemingEntity {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RedeemingEntity),
		})
	}

	// IssueDate uint64
	fip = append(ofip, CouponFieldIssueDate)
	if a.IssueDate != newValue.IssueDate {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.IssueDate)); err != nil {
			return nil, errors.Wrap(err, "IssueDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpiryDate uint64
	fip = append(ofip, CouponFieldExpiryDate)
	if a.ExpiryDate != newValue.ExpiryDate {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ExpiryDate)); err != nil {
			return nil, errors.Wrap(err, "ExpiryDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Value uint64
	fip = append(ofip, CouponFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Currency string
	fip = append(ofip, CouponFieldCurrency)
	if a.Currency != newValue.Currency {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Currency),
		})
	}

	// Description string
	fip = append(ofip, CouponFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// Precision uint64
	fip = append(ofip, CouponFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// LoyaltyPoints Permission / Amendment Field Indices
const (
	LoyaltyPointsFieldAgeRestriction      = uint32(1)
	LoyaltyPointsFieldOfferName           = uint32(2)
	LoyaltyPointsFieldValidFrom           = uint32(3)
	LoyaltyPointsFieldExpirationTimestamp = uint32(4)
	LoyaltyPointsFieldDescription         = uint32(5)
)

// ApplyAmendment updates a LoyaltyPoints based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *LoyaltyPoints) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case LoyaltyPointsFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case LoyaltyPointsFieldOfferName: // string
		a.OfferName = string(data)
		return fip[:], nil

	case LoyaltyPointsFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return fip[:], nil

	case LoyaltyPointsFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return fip[:], nil

	case LoyaltyPointsFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown LoyaltyPoints amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two LoyaltyPointss and returns
// amendment data.
func (a *LoyaltyPoints) CreateAmendments(fip []uint32,
	newValue *LoyaltyPoints) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, LoyaltyPointsFieldAgeRestriction)
	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// OfferName string
	fip = append(ofip, LoyaltyPointsFieldOfferName)
	if a.OfferName != newValue.OfferName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.OfferName),
		})
	}

	// ValidFrom uint64
	fip = append(ofip, LoyaltyPointsFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, LoyaltyPointsFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Description string
	fip = append(ofip, LoyaltyPointsFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	return result, nil
}

// TicketAdmission Permission / Amendment Field Indices
const (
	TicketAdmissionFieldAgeRestriction      = uint32(1)
	TicketAdmissionFieldAdmissionType       = uint32(2)
	TicketAdmissionFieldVenue               = uint32(3)
	TicketAdmissionFieldClass               = uint32(4)
	TicketAdmissionFieldArea                = uint32(5)
	TicketAdmissionFieldSeat                = uint32(6)
	TicketAdmissionFieldStartTimeDate       = uint32(7)
	TicketAdmissionFieldValidFrom           = uint32(8)
	TicketAdmissionFieldExpirationTimestamp = uint32(9)
	TicketAdmissionFieldDescription         = uint32(10)
)

// ApplyAmendment updates a TicketAdmission based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TicketAdmission) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case TicketAdmissionFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case TicketAdmissionFieldAdmissionType: // string
		a.AdmissionType = string(data)
		return fip[:], nil

	case TicketAdmissionFieldVenue: // string
		a.Venue = string(data)
		return fip[:], nil

	case TicketAdmissionFieldClass: // string
		a.Class = string(data)
		return fip[:], nil

	case TicketAdmissionFieldArea: // string
		a.Area = string(data)
		return fip[:], nil

	case TicketAdmissionFieldSeat: // string
		a.Seat = string(data)
		return fip[:], nil

	case TicketAdmissionFieldStartTimeDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for StartTimeDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("StartTimeDate amendment value failed to deserialize : %s", err)
		} else {
			a.StartTimeDate = uint64(value)
		}
		return fip[:], nil

	case TicketAdmissionFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return fip[:], nil

	case TicketAdmissionFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return fip[:], nil

	case TicketAdmissionFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown TicketAdmission amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two TicketAdmissions and returns
// amendment data.
func (a *TicketAdmission) CreateAmendments(fip []uint32,
	newValue *TicketAdmission) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, TicketAdmissionFieldAgeRestriction)
	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// AdmissionType string
	fip = append(ofip, TicketAdmissionFieldAdmissionType)
	if a.AdmissionType != newValue.AdmissionType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.AdmissionType),
		})
	}

	// Venue string
	fip = append(ofip, TicketAdmissionFieldVenue)
	if a.Venue != newValue.Venue {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Venue),
		})
	}

	// Class string
	fip = append(ofip, TicketAdmissionFieldClass)
	if a.Class != newValue.Class {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Class),
		})
	}

	// Area string
	fip = append(ofip, TicketAdmissionFieldArea)
	if a.Area != newValue.Area {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Area),
		})
	}

	// Seat string
	fip = append(ofip, TicketAdmissionFieldSeat)
	if a.Seat != newValue.Seat {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Seat),
		})
	}

	// StartTimeDate uint64
	fip = append(ofip, TicketAdmissionFieldStartTimeDate)
	if a.StartTimeDate != newValue.StartTimeDate {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.StartTimeDate)); err != nil {
			return nil, errors.Wrap(err, "StartTimeDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ValidFrom uint64
	fip = append(ofip, TicketAdmissionFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, TicketAdmissionFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Description string
	fip = append(ofip, TicketAdmissionFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	return result, nil
}

// CasinoChip Permission / Amendment Field Indices
const (
	CasinoChipFieldCurrencyCode        = uint32(1)
	CasinoChipFieldUseType             = uint32(2)
	CasinoChipFieldAgeRestriction      = uint32(3)
	CasinoChipFieldValidFrom           = uint32(4)
	CasinoChipFieldExpirationTimestamp = uint32(5)
	CasinoChipFieldPrecision           = uint32(6)
)

// ApplyAmendment updates a CasinoChip based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CasinoChip) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CasinoChipFieldCurrencyCode: // string
		a.CurrencyCode = string(data)
		return fip[:], nil

	case CasinoChipFieldUseType: // string
		a.UseType = string(data)
		return fip[:], nil

	case CasinoChipFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case CasinoChipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return fip[:], nil

	case CasinoChipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return fip[:], nil

	case CasinoChipFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown CasinoChip amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CasinoChips and returns
// amendment data.
func (a *CasinoChip) CreateAmendments(fip []uint32,
	newValue *CasinoChip) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// CurrencyCode string
	fip = append(ofip, CasinoChipFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// UseType string
	fip = append(ofip, CasinoChipFieldUseType)
	if a.UseType != newValue.UseType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.UseType),
		})
	}

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, CasinoChipFieldAgeRestriction)
	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ValidFrom uint64
	fip = append(ofip, CasinoChipFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, CasinoChipFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Precision uint64
	fip = append(ofip, CasinoChipFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// AgeRestrictionField Permission / Amendment Field Indices
const (
	AgeRestrictionFieldLower = uint32(1)
	AgeRestrictionFieldUpper = uint32(2)
)

// ApplyAmendment updates a AgeRestrictionField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AgeRestrictionField) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty AgeRestriction amendment field index path")
	}

	switch fip[0] {
	case AgeRestrictionFieldLower: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Lower : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Lower amendment value failed to deserialize : %s", err)
		} else {
			a.Lower = uint32(value)
		}
		return fip[:], nil
	case AgeRestrictionFieldUpper: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Upper : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Upper amendment value failed to deserialize : %s", err)
		} else {
			a.Upper = uint32(value)
		}
		return fip[:], nil
	}

	return nil, fmt.Errorf("Unknown AgeRestriction amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AgeRestrictions and returns
// amendment data.
func (a *AgeRestrictionField) CreateAmendments(fip []uint32,
	newValue *AgeRestrictionField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// Lower uint32
	fip = append(ofip, AgeRestrictionFieldLower)
	if a.Lower != newValue.Lower {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Lower)); err != nil {
			return nil, errors.Wrap(err, "Lower")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Upper uint32
	fip = append(ofip, AgeRestrictionFieldUpper)
	if a.Upper != newValue.Upper {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.Upper)); err != nil {
			return nil, errors.Wrap(err, "Upper")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}
