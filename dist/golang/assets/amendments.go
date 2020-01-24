package assets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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
		if len(data) != 8 {
			return nil, fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case MembershipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
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

// Currency Permission / Amendment Field Indices
const (
	CurrencyFieldCurrencyCode      = uint32(1)
	CurrencyFieldMonetaryAuthority = uint32(2)
	CurrencyFieldDescription       = uint32(3)
	CurrencyFieldPrecision         = uint32(4)
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

	case CurrencyFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	case CurrencyFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Precision amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Precision); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Currency amendment field index : %v", fip)
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

// Coupon Permission / Amendment Field Indices
const (
	CouponFieldRedeemingEntity = uint32(1)
	CouponFieldIssueDate       = uint32(2)
	CouponFieldExpiryDate      = uint32(3)
	CouponFieldValue           = uint32(4)
	CouponFieldCurrency        = uint32(5)
	CouponFieldDescription     = uint32(6)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("IssueDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.IssueDate); err != nil {
			return nil, fmt.Errorf("IssueDate amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case CouponFieldExpiryDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpiryDate : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ExpiryDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpiryDate); err != nil {
			return nil, fmt.Errorf("ExpiryDate amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case CouponFieldValue: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Value amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Value); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case CouponFieldCurrency: // string
		a.Currency = string(data)
		return fip[:], nil

	case CouponFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Coupon amendment field index : %v", fip)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case LoyaltyPointsFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case LoyaltyPointsFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown LoyaltyPoints amendment field index : %v", fip)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("StartTimeDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.StartTimeDate); err != nil {
			return nil, fmt.Errorf("StartTimeDate amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case TicketAdmissionFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case TicketAdmissionFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case TicketAdmissionFieldDescription: // string
		a.Description = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown TicketAdmission amendment field index : %v", fip)
}

// CasinoChip Permission / Amendment Field Indices
const (
	CasinoChipFieldCurrencyCode        = uint32(1)
	CasinoChipFieldUseType             = uint32(2)
	CasinoChipFieldAgeRestriction      = uint32(3)
	CasinoChipFieldValidFrom           = uint32(4)
	CasinoChipFieldExpirationTimestamp = uint32(5)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case CasinoChipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown CasinoChip amendment field index : %v", fip)
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
		if len(data) != 1 {
			return nil, fmt.Errorf("Lower amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Lower); err != nil {
			return nil, fmt.Errorf("Lower amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil
	case AgeRestrictionFieldUpper: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Upper : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("Upper amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Upper); err != nil {
			return nil, fmt.Errorf("Upper amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil
	}

	return nil, fmt.Errorf("Unknown AgeRestriction amendment field index : %v", fip)
}
