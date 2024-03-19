package instruments

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/internal"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/pkg/errors"
)

const (
	// AmendmentOperationModify specifies an amendment is modifying a value.
	AmendmentOperationModify = uint32(0)

	// AmendmentOperationAddElement specifies an amendment is adding a new element to a list.
	AmendmentOperationAddElement = uint32(1)

	// AmendmentOperationRemoveElement specifies an amendment is removing an element from a list.
	AmendmentOperationRemoveElement = uint32(2)
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
	MembershipFieldTransfersPermitted  = uint32(9)
)

// ApplyAmendment updates a Membership based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Membership) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case MembershipFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case MembershipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldID: // string
		a.ID = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldMembershipClass: // string
		a.MembershipClass = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldRoleType: // string
		a.RoleType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldMembershipType: // string
		a.MembershipType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Membership amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Memberships and returns
// amendment data.
func (a *Membership) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *Membership) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

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
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFrom)); err != nil {
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
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
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

	// TransfersPermitted bool
	fip = append(ofip, MembershipFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
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
func (a *Currency) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case CurrencyFieldCurrencyCode: // string
		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CurrencyFieldMonetaryAuthority: // string
		a.MonetaryAuthority = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedCurrencyFieldDescription: // deprecated

	case CurrencyFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Currency amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Currencys and returns
// amendment data.
func (a *Currency) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *Currency) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

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
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
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
	ShareCommonFieldTicker             = uint32(1)
	ShareCommonFieldISIN               = uint32(2)
	ShareCommonFieldDescription        = uint32(3)
	ShareCommonFieldTransfersPermitted = uint32(4)
)

// ApplyAmendment updates a ShareCommon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ShareCommon) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case ShareCommonFieldTicker: // string
		a.Ticker = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldISIN: // string
		a.ISIN = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown ShareCommon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two ShareCommons and returns
// amendment data.
func (a *ShareCommon) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ShareCommon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

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

	// TransfersPermitted bool
	fip = append(ofip, ShareCommonFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// BondFixedRate Permission / Amendment Field Indices
const (
	BondFixedRateFieldName                       = uint32(1)
	BondFixedRateFieldBondType                   = uint32(2)
	BondFixedRateFieldISIN                       = uint32(3)
	BondFixedRateFieldCollateral                 = uint32(4)
	BondFixedRateFieldParValue                   = uint32(5)
	BondFixedRateFieldInterestRate               = uint32(6)
	BondFixedRateFieldInterestPaymentInitialDate = uint32(7)
	BondFixedRateFieldInterestPaymentDateDeltas  = uint32(8)
	BondFixedRateFieldLatePaymentPenaltyRate     = uint32(9)
	BondFixedRateFieldLatePaymentWindow          = uint32(10)
	BondFixedRateFieldLatePaymentPenaltyPeriod   = uint32(11)
	BondFixedRateFieldMaturityDate               = uint32(12)
	BondFixedRateFieldAgeRestriction             = uint32(13)
	BondFixedRateFieldTransfersPermitted         = uint32(14)
)

// ApplyAmendment updates a BondFixedRate based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *BondFixedRate) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case BondFixedRateFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldBondType: // string
		a.BondType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldISIN: // string
		a.ISIN = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldCollateral: // string
		a.Collateral = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldParValue: // CurrencyValueField
		if len(fip) == 1 && len(data) == 0 {
			a.ParValue = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.ParValue.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *CurrencyValueField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &CurrencyValueField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for ParValue failed to deserialize : %s",
					err)
			}
		}

		a.ParValue = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case BondFixedRateFieldInterestRate: // RateField
		if len(fip) == 1 && len(data) == 0 {
			a.InterestRate = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.InterestRate.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *RateField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &RateField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for InterestRate failed to deserialize : %s",
					err)
			}
		}

		a.InterestRate = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case BondFixedRateFieldInterestPaymentInitialDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for InterestPaymentInitialDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("InterestPaymentInitialDate amendment value failed to deserialize : %s", err)
		} else {
			a.InterestPaymentInitialDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldInterestPaymentDateDeltas: // []uint64
		switch operation {
		case 0: // Modify
			if len(fip) == 1 && len(data) == 0 {
				a.InterestPaymentDateDeltas = nil
				return permissions.SubPermissions(fip[1:], operation, true)
			}
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify InterestPaymentDateDeltas : %v",
					fip)
			}
			if int(fip[1]) >= len(a.InterestPaymentDateDeltas) {
				return nil, fmt.Errorf("Amendment element index out of range for modify InterestPaymentDateDeltas : %d", fip[1])
			}
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("InterestPaymentDateDeltas amendment value failed to deserialize : %s", err)
			} else {
				a.InterestPaymentDateDeltas[fip[1]] = uint64(value)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 1: // Add element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path wrong depth for add InterestPaymentDateDeltas : %v",
					fip)
			}

			var newValue uint64
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("InterestPaymentDateDeltas amendment value failed to deserialize : %s",
					err)
			} else {
				newValue = uint64(value)
			}

			if len(a.InterestPaymentDateDeltas) <= int(fip[1]) {
				// Append item to the end
				a.InterestPaymentDateDeltas = append(a.InterestPaymentDateDeltas, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.InterestPaymentDateDeltas[:fip[1]]
				after := make([]uint64, len(a.InterestPaymentDateDeltas)-int(fip[1]))
				copy(after, a.InterestPaymentDateDeltas[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.InterestPaymentDateDeltas = append(before, newValue)
				a.InterestPaymentDateDeltas = append(a.InterestPaymentDateDeltas, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element

			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete InterestPaymentDateDeltas : %v",
					fip)
			}
			if int(fip[1]) >= len(a.InterestPaymentDateDeltas) {
				return nil, fmt.Errorf("Amendment element index out of range for delete InterestPaymentDateDeltas : %d", fip[1])
			}

			// Remove item from list
			a.InterestPaymentDateDeltas = append(a.InterestPaymentDateDeltas[:fip[1]], a.InterestPaymentDateDeltas[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case BondFixedRateFieldLatePaymentPenaltyRate: // RateField
		if len(fip) == 1 && len(data) == 0 {
			a.LatePaymentPenaltyRate = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.LatePaymentPenaltyRate.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *RateField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &RateField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for LatePaymentPenaltyRate failed to deserialize : %s",
					err)
			}
		}

		a.LatePaymentPenaltyRate = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case BondFixedRateFieldLatePaymentWindow: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for LatePaymentWindow : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("LatePaymentWindow amendment value failed to deserialize : %s", err)
		} else {
			a.LatePaymentWindow = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldLatePaymentPenaltyPeriod: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for LatePaymentPenaltyPeriod : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("LatePaymentPenaltyPeriod amendment value failed to deserialize : %s", err)
		} else {
			a.LatePaymentPenaltyPeriod = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldMaturityDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for MaturityDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("MaturityDate amendment value failed to deserialize : %s", err)
		} else {
			a.MaturityDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case BondFixedRateFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown BondFixedRate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two BondFixedRates and returns
// amendment data.
func (a *BondFixedRate) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *BondFixedRate) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Name string
	fip = append(ofip, BondFixedRateFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	// BondType string
	fip = append(ofip, BondFixedRateFieldBondType)
	if a.BondType != newValue.BondType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.BondType),
		})
	}

	// ISIN string
	fip = append(ofip, BondFixedRateFieldISIN)
	if a.ISIN != newValue.ISIN {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ISIN),
		})
	}

	// Collateral string
	fip = append(ofip, BondFixedRateFieldCollateral)
	if a.Collateral != newValue.Collateral {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Collateral),
		})
	}

	// ParValue CurrencyValueField
	fip = append(ofip, BondFixedRateFieldParValue)

	ParValueAmendments, err := a.ParValue.CreateAmendments(fip, newValue.ParValue)
	if err != nil {
		return nil, errors.Wrap(err, "ParValue")
	}
	result = append(result, ParValueAmendments...)

	// InterestRate RateField
	fip = append(ofip, BondFixedRateFieldInterestRate)

	InterestRateAmendments, err := a.InterestRate.CreateAmendments(fip, newValue.InterestRate)
	if err != nil {
		return nil, errors.Wrap(err, "InterestRate")
	}
	result = append(result, InterestRateAmendments...)

	// InterestPaymentInitialDate uint64
	fip = append(ofip, BondFixedRateFieldInterestPaymentInitialDate)
	if a.InterestPaymentInitialDate != newValue.InterestPaymentInitialDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.InterestPaymentInitialDate)); err != nil {
			return nil, errors.Wrap(err, "InterestPaymentInitialDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// InterestPaymentDateDeltas []uint64
	fip = append(ofip, BondFixedRateFieldInterestPaymentDateDeltas)
	InterestPaymentDateDeltasMin := len(a.InterestPaymentDateDeltas)
	if InterestPaymentDateDeltasMin > len(newValue.InterestPaymentDateDeltas) {
		InterestPaymentDateDeltasMin = len(newValue.InterestPaymentDateDeltas)
	}

	// Compare values
	for i := 0; i < InterestPaymentDateDeltasMin; i++ {
		lfip := append(fip, uint32(i))
		if a.InterestPaymentDateDeltas[i] != newValue.InterestPaymentDateDeltas[i] {
			var buf bytes.Buffer
			if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.InterestPaymentDateDeltas[i])); err != nil {
				return nil, errors.Wrapf(err, "InterestPaymentDateDeltas %d", i)
			}

			result = append(result, &internal.Amendment{
				FIP:       lfip,
				Operation: 0, // Modify element
				Data:      buf.Bytes(),
			})
		}
	}

	InterestPaymentDateDeltasMax := len(a.InterestPaymentDateDeltas)
	if InterestPaymentDateDeltasMax < len(newValue.InterestPaymentDateDeltas) {
		InterestPaymentDateDeltasMax = len(newValue.InterestPaymentDateDeltas)
	}

	// Add/Remove values
	for i := InterestPaymentDateDeltasMin; i < InterestPaymentDateDeltasMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.InterestPaymentDateDeltas) {
			amendment.Operation = 1 // Add element
			var buf bytes.Buffer
			if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.InterestPaymentDateDeltas[i])); err != nil {
				return nil, errors.Wrapf(err, "InterestPaymentDateDeltas %d", i)
			}
			amendment.Data = buf.Bytes()
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// LatePaymentPenaltyRate RateField
	fip = append(ofip, BondFixedRateFieldLatePaymentPenaltyRate)

	LatePaymentPenaltyRateAmendments, err := a.LatePaymentPenaltyRate.CreateAmendments(fip, newValue.LatePaymentPenaltyRate)
	if err != nil {
		return nil, errors.Wrap(err, "LatePaymentPenaltyRate")
	}
	result = append(result, LatePaymentPenaltyRateAmendments...)

	// LatePaymentWindow uint64
	fip = append(ofip, BondFixedRateFieldLatePaymentWindow)
	if a.LatePaymentWindow != newValue.LatePaymentWindow {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.LatePaymentWindow)); err != nil {
			return nil, errors.Wrap(err, "LatePaymentWindow")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// LatePaymentPenaltyPeriod uint64
	fip = append(ofip, BondFixedRateFieldLatePaymentPenaltyPeriod)
	if a.LatePaymentPenaltyPeriod != newValue.LatePaymentPenaltyPeriod {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.LatePaymentPenaltyPeriod)); err != nil {
			return nil, errors.Wrap(err, "LatePaymentPenaltyPeriod")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// MaturityDate uint64
	fip = append(ofip, BondFixedRateFieldMaturityDate)
	if a.MaturityDate != newValue.MaturityDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.MaturityDate)); err != nil {
			return nil, errors.Wrap(err, "MaturityDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, BondFixedRateFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// TransfersPermitted bool
	fip = append(ofip, BondFixedRateFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// DiscountCoupon Permission / Amendment Field Indices
const (
	DiscountCouponFieldRedeemingEntity     = uint32(1)
	DiscountCouponFieldValidFromTimestamp  = uint32(2)
	DiscountCouponFieldExpirationTimestamp = uint32(3)
	DeprecatedDiscountCouponFieldValue     = uint32(4)
	DeprecatedDiscountCouponFieldCurrency  = uint32(5)
	DiscountCouponFieldCouponName          = uint32(6)
	DeprecatedDiscountCouponFieldPrecision = uint32(7)
	DiscountCouponFieldTransfersPermitted  = uint32(8)
	DiscountCouponFieldFaceValue           = uint32(9)
	DiscountCouponFieldRedemptionVenue     = uint32(10)
	DiscountCouponFieldDetails             = uint32(11)
)

// ApplyAmendment updates a DiscountCoupon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *DiscountCoupon) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case DiscountCouponFieldRedeemingEntity: // string
		a.RedeemingEntity = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DiscountCouponFieldValidFromTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFromTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFromTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFromTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DiscountCouponFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedDiscountCouponFieldValue: // deprecated

	case DeprecatedDiscountCouponFieldCurrency: // deprecated

	case DiscountCouponFieldCouponName: // string
		a.CouponName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedDiscountCouponFieldPrecision: // deprecated

	case DiscountCouponFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DiscountCouponFieldFaceValue: // CurrencyValueField
		if len(fip) == 1 && len(data) == 0 {
			a.FaceValue = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.FaceValue.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *CurrencyValueField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &CurrencyValueField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for FaceValue failed to deserialize : %s",
					err)
			}
		}

		a.FaceValue = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case DiscountCouponFieldRedemptionVenue: // string
		a.RedemptionVenue = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DiscountCouponFieldDetails: // string
		a.Details = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown DiscountCoupon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two DiscountCoupons and returns
// amendment data.
func (a *DiscountCoupon) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *DiscountCoupon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// RedeemingEntity string
	fip = append(ofip, DiscountCouponFieldRedeemingEntity)
	if a.RedeemingEntity != newValue.RedeemingEntity {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RedeemingEntity),
		})
	}

	// ValidFromTimestamp uint64
	fip = append(ofip, DiscountCouponFieldValidFromTimestamp)
	if a.ValidFromTimestamp != newValue.ValidFromTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFromTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ValidFromTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, DiscountCouponFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// deprecated Value deprecated

	// deprecated Currency deprecated

	// CouponName string
	fip = append(ofip, DiscountCouponFieldCouponName)
	if a.CouponName != newValue.CouponName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CouponName),
		})
	}

	// deprecated Precision deprecated

	// TransfersPermitted bool
	fip = append(ofip, DiscountCouponFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// FaceValue CurrencyValueField
	fip = append(ofip, DiscountCouponFieldFaceValue)

	FaceValueAmendments, err := a.FaceValue.CreateAmendments(fip, newValue.FaceValue)
	if err != nil {
		return nil, errors.Wrap(err, "FaceValue")
	}
	result = append(result, FaceValueAmendments...)

	// RedemptionVenue string
	fip = append(ofip, DiscountCouponFieldRedemptionVenue)
	if a.RedemptionVenue != newValue.RedemptionVenue {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RedemptionVenue),
		})
	}

	// Details string
	fip = append(ofip, DiscountCouponFieldDetails)
	if a.Details != newValue.Details {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Details),
		})
	}

	return result, nil
}

// DeprecatedLoyaltyPoints Permission / Amendment Field Indices
const (
	DeprecatedLoyaltyPointsFieldAgeRestriction      = uint32(1)
	DeprecatedLoyaltyPointsFieldProgramName         = uint32(2)
	DeprecatedDeprecatedLoyaltyPointsFieldValidFrom = uint32(3)
	DeprecatedLoyaltyPointsFieldExpirationTimestamp = uint32(4)
	DeprecatedLoyaltyPointsFieldDetails             = uint32(5)
	DeprecatedLoyaltyPointsFieldTransfersPermitted  = uint32(6)
)

// ApplyAmendment updates a DeprecatedLoyaltyPoints based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *DeprecatedLoyaltyPoints) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case DeprecatedLoyaltyPointsFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case DeprecatedLoyaltyPointsFieldProgramName: // string
		a.ProgramName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedDeprecatedLoyaltyPointsFieldValidFrom: // deprecated

	case DeprecatedLoyaltyPointsFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedLoyaltyPointsFieldDetails: // string
		a.Details = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedLoyaltyPointsFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown DeprecatedLoyaltyPoints amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two DeprecatedLoyaltyPointss and returns
// amendment data.
func (a *DeprecatedLoyaltyPoints) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *DeprecatedLoyaltyPoints) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, DeprecatedLoyaltyPointsFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ProgramName string
	fip = append(ofip, DeprecatedLoyaltyPointsFieldProgramName)
	if a.ProgramName != newValue.ProgramName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ProgramName),
		})
	}

	// deprecated ValidFrom deprecated

	// ExpirationTimestamp uint64
	fip = append(ofip, DeprecatedLoyaltyPointsFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Details string
	fip = append(ofip, DeprecatedLoyaltyPointsFieldDetails)
	if a.Details != newValue.Details {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Details),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, DeprecatedLoyaltyPointsFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// TicketAdmission Permission / Amendment Field Indices
const (
	TicketAdmissionFieldAgeRestriction                = uint32(1)
	DeprecatedTicketAdmissionFieldAdmissionType       = uint32(2)
	TicketAdmissionFieldVenue                         = uint32(3)
	DeprecatedTicketAdmissionFieldClass               = uint32(4)
	TicketAdmissionFieldArea                          = uint32(5)
	TicketAdmissionFieldSeat                          = uint32(6)
	TicketAdmissionFieldEventStartTimestamp           = uint32(7)
	DeprecatedTicketAdmissionFieldValidFrom           = uint32(8)
	DeprecatedTicketAdmissionFieldExpirationTimestamp = uint32(9)
	TicketAdmissionFieldEventName                     = uint32(10)
	TicketAdmissionFieldTransfersPermitted            = uint32(11)
	TicketAdmissionFieldDetails                       = uint32(12)
	TicketAdmissionFieldSection                       = uint32(13)
	TicketAdmissionFieldRow                           = uint32(14)
	TicketAdmissionFieldEventEndTimestamp             = uint32(15)
)

// ApplyAmendment updates a TicketAdmission based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TicketAdmission) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case TicketAdmissionFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case DeprecatedTicketAdmissionFieldAdmissionType: // deprecated

	case TicketAdmissionFieldVenue: // string
		a.Venue = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedTicketAdmissionFieldClass: // deprecated

	case TicketAdmissionFieldArea: // string
		a.Area = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldSeat: // string
		a.Seat = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldEventStartTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EventStartTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("EventStartTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.EventStartTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedTicketAdmissionFieldValidFrom: // deprecated

	case DeprecatedTicketAdmissionFieldExpirationTimestamp: // deprecated

	case TicketAdmissionFieldEventName: // string
		a.EventName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldDetails: // string
		a.Details = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldSection: // string
		a.Section = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldRow: // string
		a.Row = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldEventEndTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EventEndTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("EventEndTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.EventEndTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown TicketAdmission amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two TicketAdmissions and returns
// amendment data.
func (a *TicketAdmission) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *TicketAdmission) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, TicketAdmissionFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// deprecated AdmissionType deprecated

	// Venue string
	fip = append(ofip, TicketAdmissionFieldVenue)
	if a.Venue != newValue.Venue {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Venue),
		})
	}

	// deprecated Class deprecated

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

	// EventStartTimestamp uint64
	fip = append(ofip, TicketAdmissionFieldEventStartTimestamp)
	if a.EventStartTimestamp != newValue.EventStartTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.EventStartTimestamp)); err != nil {
			return nil, errors.Wrap(err, "EventStartTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// deprecated ValidFrom deprecated

	// deprecated ExpirationTimestamp deprecated

	// EventName string
	fip = append(ofip, TicketAdmissionFieldEventName)
	if a.EventName != newValue.EventName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.EventName),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, TicketAdmissionFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Details string
	fip = append(ofip, TicketAdmissionFieldDetails)
	if a.Details != newValue.Details {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Details),
		})
	}

	// Section string
	fip = append(ofip, TicketAdmissionFieldSection)
	if a.Section != newValue.Section {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Section),
		})
	}

	// Row string
	fip = append(ofip, TicketAdmissionFieldRow)
	if a.Row != newValue.Row {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Row),
		})
	}

	// EventEndTimestamp uint64
	fip = append(ofip, TicketAdmissionFieldEventEndTimestamp)
	if a.EventEndTimestamp != newValue.EventEndTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.EventEndTimestamp)); err != nil {
			return nil, errors.Wrap(err, "EventEndTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CasinoChip Permission / Amendment Field Indices
const (
	DeprecatedCasinoChipFieldCurrencyCode = uint32(1)
	CasinoChipFieldUseType                = uint32(2)
	CasinoChipFieldAgeRestriction         = uint32(3)
	DeprecatedCasinoChipFieldValidFrom    = uint32(4)
	CasinoChipFieldExpirationTimestamp    = uint32(5)
	DeprecatedCasinoChipFieldPrecision    = uint32(6)
	CasinoChipFieldTransfersPermitted     = uint32(7)
	CasinoChipFieldCasinoName             = uint32(8)
	CasinoChipFieldFaceValue              = uint32(9)
)

// ApplyAmendment updates a CasinoChip based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CasinoChip) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case DeprecatedCasinoChipFieldCurrencyCode: // deprecated

	case CasinoChipFieldUseType: // string
		a.UseType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case DeprecatedCasinoChipFieldValidFrom: // deprecated

	case CasinoChipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedCasinoChipFieldPrecision: // deprecated

	case CasinoChipFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldCasinoName: // string
		a.CasinoName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldFaceValue: // CurrencyValueField
		if len(fip) == 1 && len(data) == 0 {
			a.FaceValue = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.FaceValue.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *CurrencyValueField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &CurrencyValueField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for FaceValue failed to deserialize : %s",
					err)
			}
		}

		a.FaceValue = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	}

	return nil, fmt.Errorf("Unknown CasinoChip amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CasinoChips and returns
// amendment data.
func (a *CasinoChip) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *CasinoChip) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// deprecated CurrencyCode deprecated

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

	// deprecated ValidFrom deprecated

	// ExpirationTimestamp uint64
	fip = append(ofip, CasinoChipFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// deprecated Precision deprecated

	// TransfersPermitted bool
	fip = append(ofip, CasinoChipFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// CasinoName string
	fip = append(ofip, CasinoChipFieldCasinoName)
	if a.CasinoName != newValue.CasinoName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CasinoName),
		})
	}

	// FaceValue CurrencyValueField
	fip = append(ofip, CasinoChipFieldFaceValue)

	FaceValueAmendments, err := a.FaceValue.CreateAmendments(fip, newValue.FaceValue)
	if err != nil {
		return nil, errors.Wrap(err, "FaceValue")
	}
	result = append(result, FaceValueAmendments...)

	return result, nil
}

// InformationServiceLicense Permission / Amendment Field Indices
const (
	InformationServiceLicenseFieldAgeRestriction      = uint32(1)
	InformationServiceLicenseFieldExpirationTimestamp = uint32(2)
	InformationServiceLicenseFieldServiceName         = uint32(3)
	InformationServiceLicenseFieldTransfersPermitted  = uint32(4)
	InformationServiceLicenseFieldURL                 = uint32(5)
)

// ApplyAmendment updates a InformationServiceLicense based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *InformationServiceLicense) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case InformationServiceLicenseFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case InformationServiceLicenseFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InformationServiceLicenseFieldServiceName: // string
		a.ServiceName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case InformationServiceLicenseFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InformationServiceLicenseFieldURL: // string
		a.URL = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown InformationServiceLicense amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two InformationServiceLicenses and returns
// amendment data.
func (a *InformationServiceLicense) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *InformationServiceLicense) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, InformationServiceLicenseFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ExpirationTimestamp uint64
	fip = append(ofip, InformationServiceLicenseFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ServiceName string
	fip = append(ofip, InformationServiceLicenseFieldServiceName)
	if a.ServiceName != newValue.ServiceName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ServiceName),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, InformationServiceLicenseFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// URL string
	fip = append(ofip, InformationServiceLicenseFieldURL)
	if a.URL != newValue.URL {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.URL),
		})
	}

	return result, nil
}

// CreditNote Permission / Amendment Field Indices
const (
	DeprecatedCreditNoteFieldName      = uint32(1)
	CreditNoteFieldFaceValue           = uint32(2)
	CreditNoteFieldExpirationTimestamp = uint32(3)
	CreditNoteFieldTransfersPermitted  = uint32(4)
)

// ApplyAmendment updates a CreditNote based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CreditNote) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case DeprecatedCreditNoteFieldName: // deprecated

	case CreditNoteFieldFaceValue: // FixedCurrencyValueField
		if len(fip) == 1 && len(data) == 0 {
			a.FaceValue = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.FaceValue.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *FixedCurrencyValueField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &FixedCurrencyValueField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for FaceValue failed to deserialize : %s",
					err)
			}
		}

		a.FaceValue = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case CreditNoteFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CreditNoteFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown CreditNote amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CreditNotes and returns
// amendment data.
func (a *CreditNote) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *CreditNote) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// deprecated Name deprecated

	// FaceValue FixedCurrencyValueField
	fip = append(ofip, CreditNoteFieldFaceValue)

	FaceValueAmendments, err := a.FaceValue.CreateAmendments(fip, newValue.FaceValue)
	if err != nil {
		return nil, errors.Wrap(err, "FaceValue")
	}
	result = append(result, FaceValueAmendments...)

	// ExpirationTimestamp uint64
	fip = append(ofip, CreditNoteFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, CreditNoteFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// RewardPoint Permission / Amendment Field Indices
const (
	RewardPointFieldAgeRestriction      = uint32(1)
	RewardPointFieldProgramName         = uint32(2)
	DeprecatedRewardPointFieldValidFrom = uint32(3)
	RewardPointFieldExpirationTimestamp = uint32(4)
	RewardPointFieldDetails             = uint32(5)
	RewardPointFieldTransfersPermitted  = uint32(6)
)

// ApplyAmendment updates a RewardPoint based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *RewardPoint) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
	case RewardPointFieldAgeRestriction: // AgeRestrictionField
		if len(fip) == 1 && len(data) == 0 {
			a.AgeRestriction = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		if len(fip) > 1 {
			// For sub-field
			subPermissions, err := permissions.SubPermissions(fip, operation, false)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)
		}

		var newValue *AgeRestrictionField
		if len(data) != 0 { // Leave default values if data is empty
			newValue = &AgeRestrictionField{}
			if err := proto.Unmarshal(data, newValue); err != nil {
				return nil, fmt.Errorf("Amendment for AgeRestriction failed to deserialize : %s",
					err)
			}
		}

		a.AgeRestriction = newValue
		return permissions.SubPermissions(fip[1:], operation, false)

	case RewardPointFieldProgramName: // string
		a.ProgramName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedRewardPointFieldValidFrom: // deprecated

	case RewardPointFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case RewardPointFieldDetails: // string
		a.Details = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case RewardPointFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown RewardPoint amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two RewardPoints and returns
// amendment data.
func (a *RewardPoint) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *RewardPoint) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, RewardPointFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ProgramName string
	fip = append(ofip, RewardPointFieldProgramName)
	if a.ProgramName != newValue.ProgramName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ProgramName),
		})
	}

	// deprecated ValidFrom deprecated

	// ExpirationTimestamp uint64
	fip = append(ofip, RewardPointFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Details string
	fip = append(ofip, RewardPointFieldDetails)
	if a.Details != newValue.Details {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Details),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, RewardPointFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
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
func (a *AgeRestrictionField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AgeRestriction amendment field index path")
	}

	switch fip[0] {
	case AgeRestrictionFieldLower: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Lower : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Lower amendment value failed to deserialize : %s", err)
		} else {
			a.Lower = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case AgeRestrictionFieldUpper: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Upper : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Upper amendment value failed to deserialize : %s", err)
		} else {
			a.Upper = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown AgeRestriction amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AgeRestrictions and returns
// amendment data.
func (a *AgeRestrictionField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *AgeRestrictionField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Lower uint32
	fip = append(ofip, AgeRestrictionFieldLower)
	if a.Lower != newValue.Lower {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Lower)); err != nil {
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
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Upper)); err != nil {
			return nil, errors.Wrap(err, "Upper")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CurrencyValueField Permission / Amendment Field Indices
const (
	CurrencyValueFieldValue        = uint32(1)
	CurrencyValueFieldCurrencyCode = uint32(2)
	CurrencyValueFieldPrecision    = uint32(3)
)

// ApplyAmendment updates a CurrencyValueField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CurrencyValueField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty CurrencyValue amendment field index path")
	}

	switch fip[0] {
	case CurrencyValueFieldValue: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case CurrencyValueFieldCurrencyCode: // string

		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)
	case CurrencyValueFieldPrecision: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown CurrencyValue amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CurrencyValues and returns
// amendment data.
func (a *CurrencyValueField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *CurrencyValueField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Value uint64
	fip = append(ofip, CurrencyValueFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// CurrencyCode string
	fip = append(ofip, CurrencyValueFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// Precision uint32
	fip = append(ofip, CurrencyValueFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// FixedCurrencyValueField Permission / Amendment Field Indices
const (
	FixedCurrencyValueFieldValue        = uint32(1)
	FixedCurrencyValueFieldCurrencyCode = uint32(2)
	FixedCurrencyValueFieldPrecision    = uint32(3)
)

// ApplyAmendment updates a FixedCurrencyValueField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *FixedCurrencyValueField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty FixedCurrencyValue amendment field index path")
	}

	switch fip[0] {
	case FixedCurrencyValueFieldValue: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case FixedCurrencyValueFieldCurrencyCode: // string

		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)
	case FixedCurrencyValueFieldPrecision: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown FixedCurrencyValue amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two FixedCurrencyValues and returns
// amendment data.
func (a *FixedCurrencyValueField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *FixedCurrencyValueField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Value uint64
	fip = append(ofip, FixedCurrencyValueFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// CurrencyCode string
	fip = append(ofip, FixedCurrencyValueFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// Precision uint32
	fip = append(ofip, FixedCurrencyValueFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// RateField Permission / Amendment Field Indices
const (
	RateFieldPrecision = uint32(1)
	RateFieldValue     = uint32(2)
)

// ApplyAmendment updates a RateField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *RateField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Rate amendment field index path")
	}

	switch fip[0] {
	case RateFieldPrecision: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case RateFieldValue: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown Rate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Rates and returns
// amendment data.
func (a *RateField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *RateField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Precision uint32
	fip = append(ofip, RateFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Value uint64
	fip = append(ofip, RateFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CreatePayloadAmendments deserializes instrument payloads and create amendments for them.
func CreatePayloadAmendments(fip permissions.FieldIndexPath,
	instrumentType, payload, newPayload []byte) ([]*internal.Amendment, error) {

	current, err := Deserialize(instrumentType, payload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	new, err := Deserialize(instrumentType, newPayload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	var result []*internal.Amendment
	switch c := current.(type) {
	case *Membership:
		result, err = c.CreateAmendments(fip, new.(*Membership))

	case *Currency:
		result, err = c.CreateAmendments(fip, new.(*Currency))

	case *ShareCommon:
		result, err = c.CreateAmendments(fip, new.(*ShareCommon))

	case *BondFixedRate:
		result, err = c.CreateAmendments(fip, new.(*BondFixedRate))

	case *DiscountCoupon:
		result, err = c.CreateAmendments(fip, new.(*DiscountCoupon))

	case *DeprecatedLoyaltyPoints:
		result, err = c.CreateAmendments(fip, new.(*DeprecatedLoyaltyPoints))

	case *TicketAdmission:
		result, err = c.CreateAmendments(fip, new.(*TicketAdmission))

	case *CasinoChip:
		result, err = c.CreateAmendments(fip, new.(*CasinoChip))

	case *InformationServiceLicense:
		result, err = c.CreateAmendments(fip, new.(*InformationServiceLicense))

	case *CreditNote:
		result, err = c.CreateAmendments(fip, new.(*CreditNote))

	case *RewardPoint:
		result, err = c.CreateAmendments(fip, new.(*RewardPoint))

	}

	return result, nil
}
