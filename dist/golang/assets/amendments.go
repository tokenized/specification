package assets

import (
	"bytes"
	"encoding/binary"
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
func (a *Membership) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case MembershipFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case MembershipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}

	case MembershipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}

	case MembershipFieldID: // string
		a.ID = string(data)

	case MembershipFieldMembershipClass: // string
		a.MembershipClass = string(data)

	case MembershipFieldRoleType: // string
		a.RoleType = string(data)

	case MembershipFieldMembershipType: // string
		a.MembershipType = string(data)

	case MembershipFieldDescription: // string
		a.Description = string(data)

	}

	return nil
}

// Currency Permission / Amendment Field Indices
const (
	CurrencyFieldCurrencyCode      = uint32(1)
	CurrencyFieldMonetaryAuthority = uint32(2)
	CurrencyFieldDescription       = uint32(3)
)

// ApplyAmendment updates a Currency based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Currency) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case CurrencyFieldCurrencyCode: // string
		a.CurrencyCode = string(data)

	case CurrencyFieldMonetaryAuthority: // string
		a.MonetaryAuthority = string(data)

	case CurrencyFieldDescription: // string
		a.Description = string(data)

	}

	return nil
}

// ShareCommon Permission / Amendment Field Indices
const (
	ShareCommonFieldTransferLockout = uint32(1)
	ShareCommonFieldTicker          = uint32(2)
	ShareCommonFieldISIN            = uint32(3)
	ShareCommonFieldDescription     = uint32(4)
)

// ApplyAmendment updates a ShareCommon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ShareCommon) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case ShareCommonFieldTransferLockout: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for TransferLockout : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("TransferLockout amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransferLockout); err != nil {
			return fmt.Errorf("TransferLockout amendment value failed to deserialize : %s", err)
		}

	case ShareCommonFieldTicker: // string
		a.Ticker = string(data)

	case ShareCommonFieldISIN: // string
		a.ISIN = string(data)

	case ShareCommonFieldDescription: // string
		a.Description = string(data)

	}

	return nil
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
func (a *Coupon) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case CouponFieldRedeemingEntity: // string
		a.RedeemingEntity = string(data)

	case CouponFieldIssueDate: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for IssueDate : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("IssueDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.IssueDate); err != nil {
			return fmt.Errorf("IssueDate amendment value failed to deserialize : %s", err)
		}

	case CouponFieldExpiryDate: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ExpiryDate : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ExpiryDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpiryDate); err != nil {
			return fmt.Errorf("ExpiryDate amendment value failed to deserialize : %s", err)
		}

	case CouponFieldValue: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("Value amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Value); err != nil {
			return fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		}

	case CouponFieldCurrency: // string
		a.Currency = string(data)

	case CouponFieldDescription: // string
		a.Description = string(data)

	}

	return nil
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
func (a *LoyaltyPoints) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case LoyaltyPointsFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case LoyaltyPointsFieldOfferName: // string
		a.OfferName = string(data)

	case LoyaltyPointsFieldValidFrom: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}

	case LoyaltyPointsFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}

	case LoyaltyPointsFieldDescription: // string
		a.Description = string(data)

	}

	return nil
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
func (a *TicketAdmission) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case TicketAdmissionFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case TicketAdmissionFieldAdmissionType: // string
		a.AdmissionType = string(data)

	case TicketAdmissionFieldVenue: // string
		a.Venue = string(data)

	case TicketAdmissionFieldClass: // string
		a.Class = string(data)

	case TicketAdmissionFieldArea: // string
		a.Area = string(data)

	case TicketAdmissionFieldSeat: // string
		a.Seat = string(data)

	case TicketAdmissionFieldStartTimeDate: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for StartTimeDate : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("StartTimeDate amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.StartTimeDate); err != nil {
			return fmt.Errorf("StartTimeDate amendment value failed to deserialize : %s", err)
		}

	case TicketAdmissionFieldValidFrom: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}

	case TicketAdmissionFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}

	case TicketAdmissionFieldDescription: // string
		a.Description = string(data)

	}

	return nil
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
func (a *CasinoChip) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case CasinoChipFieldCurrencyCode: // string
		a.CurrencyCode = string(data)

	case CasinoChipFieldUseType: // string
		a.UseType = string(data)

	case CasinoChipFieldAgeRestriction: // AgeRestrictionField
		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data)

	case CasinoChipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ValidFrom amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ValidFrom); err != nil {
			return fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		}

	case CasinoChipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		if len(data) != 8 {
			return fmt.Errorf("ExpirationTimestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ExpirationTimestamp); err != nil {
			return fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		}

	}

	return nil
}

// AgeRestrictionField Permission / Amendment Field Indices
const (
	AgeRestrictionFieldLower = uint32(1)
	AgeRestrictionFieldUpper = uint32(2)
)

// ApplyAmendment updates a AgeRestrictionField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AgeRestrictionField) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
	switch fip[0] {
	case AgeRestrictionFieldLower: // uint32

		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for Lower : %v", fip)
		}
		if len(data) != 1 {
			return fmt.Errorf("Lower amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Lower); err != nil {
			return fmt.Errorf("Lower amendment value failed to deserialize : %s", err)
		}
	case AgeRestrictionFieldUpper: // uint32

		if len(fip) > 1 {
			return fmt.Errorf("Amendment field index path too deep for Upper : %v", fip)
		}
		if len(data) != 1 {
			return fmt.Errorf("Upper amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Upper); err != nil {
			return fmt.Errorf("Upper amendment value failed to deserialize : %s", err)
		}
	}

	return nil
}
