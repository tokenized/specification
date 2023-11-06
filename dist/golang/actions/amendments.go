package actions

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/internal"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/golang/protobuf/proto"
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

// Contract Permission / Amendment Field Indices
const (
	ContractFieldContractName                         = uint32(1)
	ContractFieldBodyOfAgreementType                  = uint32(2)
	ContractFieldBodyOfAgreement                      = uint32(3)
	DeprecatedContractFieldContractType               = uint32(4)
	ContractFieldSupportingDocs                       = uint32(5)
	DeprecatedContractFieldGoverningLaw               = uint32(6)
	DeprecatedContractFieldJurisdiction               = uint32(7)
	ContractFieldContractExpiration                   = uint32(8)
	ContractFieldContractURI                          = uint32(9)
	ContractFieldIssuer                               = uint32(10)
	DeprecatedContractFieldIssuerLogoURL              = uint32(11)
	NotAmendableContractFieldContractOperatorIncluded = uint32(12)
	DeprecatedContractFieldContractOperator           = uint32(13)
	DeprecatedContractFieldAdminOracle                = uint32(14)
	DeprecatedContractFieldAdminOracleSignature       = uint32(15)
	DeprecatedContractFieldAdminOracleSigBlockHeight  = uint32(16)
	ContractFieldContractFee                          = uint32(17)
	ContractFieldVotingSystems                        = uint32(18)
	ContractFieldContractPermissions                  = uint32(19)
	ContractFieldRestrictedQtyInstruments             = uint32(20)
	ContractFieldAdministrationProposal               = uint32(21)
	ContractFieldHolderProposal                       = uint32(22)
	ContractFieldOracles                              = uint32(23)
	NotAmendableContractFieldMasterAddress            = uint32(24)
	ContractFieldEntityContract                       = uint32(25)
	ContractFieldOperatorEntityContract               = uint32(26)
	ContractFieldContractType                         = uint32(27)
	ContractFieldServices                             = uint32(28)
	ContractFieldAdminIdentityCertificates            = uint32(29)
	ContractFieldGoverningLaw                         = uint32(30)
	ContractFieldJurisdiction                         = uint32(31)
)

// CreateAmendments determines the differences between two ContractOffers and returns
// amendment data. Use the current value of contract formation, and pass in the new values as a
// contract offer.
func (a *ContractFormation) CreateAmendments(newValue *ContractOffer) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

	// ContractName string
	fip = []uint32{ContractFieldContractName}
	if a.ContractName != newValue.ContractName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ContractName),
		})
	}

	// BodyOfAgreementType uint32
	fip = []uint32{ContractFieldBodyOfAgreementType}
	if a.BodyOfAgreementType != newValue.BodyOfAgreementType {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.BodyOfAgreementType)); err != nil {
			return nil, errors.Wrap(err, "BodyOfAgreementType")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// BodyOfAgreement []byte
	fip = []uint32{ContractFieldBodyOfAgreement}
	if !bytes.Equal(a.BodyOfAgreement, newValue.BodyOfAgreement) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.BodyOfAgreement,
		})
	}

	// deprecated ContractType deprecated

	// SupportingDocs []*DocumentField
	fip = []uint32{ContractFieldSupportingDocs}
	SupportingDocsMin := len(a.SupportingDocs)
	if SupportingDocsMin > len(newValue.SupportingDocs) {
		SupportingDocsMin = len(newValue.SupportingDocs)
	}

	// Compare values
	for i := 0; i < SupportingDocsMin; i++ {
		lfip := append(fip, uint32(i))
		SupportingDocsAmendments, err := a.SupportingDocs[i].CreateAmendments(lfip,
			newValue.SupportingDocs[i])
		if err != nil {
			return nil, errors.Wrapf(err, "SupportingDocs%d", i)
		}
		result = append(result, SupportingDocsAmendments...)
	}

	SupportingDocsMax := len(a.SupportingDocs)
	if SupportingDocsMax < len(newValue.SupportingDocs) {
		SupportingDocsMax = len(newValue.SupportingDocs)
	}

	// Add/Remove values
	for i := SupportingDocsMin; i < SupportingDocsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.SupportingDocs) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.SupportingDocs[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize SupportingDocs %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// deprecated GoverningLaw deprecated

	// deprecated Jurisdiction deprecated

	// ContractExpiration uint64
	fip = []uint32{ContractFieldContractExpiration}
	if a.ContractExpiration != newValue.ContractExpiration {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ContractExpiration)); err != nil {
			return nil, errors.Wrap(err, "ContractExpiration")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ContractURI string
	fip = []uint32{ContractFieldContractURI}
	if a.ContractURI != newValue.ContractURI {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ContractURI),
		})
	}

	// Issuer EntityField
	fip = []uint32{ContractFieldIssuer}

	IssuerAmendments, err := a.Issuer.CreateAmendments(fip, newValue.Issuer)
	if err != nil {
		return nil, errors.Wrap(err, "Issuer")
	}
	result = append(result, IssuerAmendments...)

	// deprecated IssuerLogoURL deprecated

	// deprecated ContractOperator deprecated

	// deprecated AdminOracle deprecated

	// deprecated AdminOracleSignature deprecated

	// deprecated AdminOracleSigBlockHeight deprecated

	// ContractFee uint64
	fip = []uint32{ContractFieldContractFee}
	if a.ContractFee != newValue.ContractFee {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ContractFee)); err != nil {
			return nil, errors.Wrap(err, "ContractFee")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// VotingSystems []*VotingSystemField
	fip = []uint32{ContractFieldVotingSystems}
	VotingSystemsMin := len(a.VotingSystems)
	if VotingSystemsMin > len(newValue.VotingSystems) {
		VotingSystemsMin = len(newValue.VotingSystems)
	}

	// Compare values
	for i := 0; i < VotingSystemsMin; i++ {
		lfip := append(fip, uint32(i))
		VotingSystemsAmendments, err := a.VotingSystems[i].CreateAmendments(lfip,
			newValue.VotingSystems[i])
		if err != nil {
			return nil, errors.Wrapf(err, "VotingSystems%d", i)
		}
		result = append(result, VotingSystemsAmendments...)
	}

	VotingSystemsMax := len(a.VotingSystems)
	if VotingSystemsMax < len(newValue.VotingSystems) {
		VotingSystemsMax = len(newValue.VotingSystems)
	}

	// Add/Remove values
	for i := VotingSystemsMin; i < VotingSystemsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.VotingSystems) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.VotingSystems[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize VotingSystems %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// ContractPermissions []byte
	fip = []uint32{ContractFieldContractPermissions}
	if !bytes.Equal(a.ContractPermissions, newValue.ContractPermissions) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.ContractPermissions,
		})
	}

	// RestrictedQtyInstruments uint64
	fip = []uint32{ContractFieldRestrictedQtyInstruments}
	if a.RestrictedQtyInstruments != newValue.RestrictedQtyInstruments {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.RestrictedQtyInstruments)); err != nil {
			return nil, errors.Wrap(err, "RestrictedQtyInstruments")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AdministrationProposal bool
	fip = []uint32{ContractFieldAdministrationProposal}
	if a.AdministrationProposal != newValue.AdministrationProposal {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.AdministrationProposal); err != nil {
			return nil, errors.Wrap(err, "AdministrationProposal")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// HolderProposal bool
	fip = []uint32{ContractFieldHolderProposal}
	if a.HolderProposal != newValue.HolderProposal {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.HolderProposal); err != nil {
			return nil, errors.Wrap(err, "HolderProposal")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Oracles []*OracleField
	fip = []uint32{ContractFieldOracles}
	OraclesMin := len(a.Oracles)
	if OraclesMin > len(newValue.Oracles) {
		OraclesMin = len(newValue.Oracles)
	}

	// Compare values
	for i := 0; i < OraclesMin; i++ {
		lfip := append(fip, uint32(i))
		OraclesAmendments, err := a.Oracles[i].CreateAmendments(lfip,
			newValue.Oracles[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Oracles%d", i)
		}
		result = append(result, OraclesAmendments...)
	}

	OraclesMax := len(a.Oracles)
	if OraclesMax < len(newValue.Oracles) {
		OraclesMax = len(newValue.Oracles)
	}

	// Add/Remove values
	for i := OraclesMin; i < OraclesMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Oracles) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Oracles[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Oracles %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// EntityContract []byte
	fip = []uint32{ContractFieldEntityContract}
	if !bytes.Equal(a.EntityContract, newValue.EntityContract) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.EntityContract,
		})
	}

	// OperatorEntityContract []byte
	fip = []uint32{ContractFieldOperatorEntityContract}
	if !bytes.Equal(a.OperatorEntityContract, newValue.OperatorEntityContract) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.OperatorEntityContract,
		})
	}

	// ContractType uint32
	fip = []uint32{ContractFieldContractType}
	if a.ContractType != newValue.ContractType {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ContractType)); err != nil {
			return nil, errors.Wrap(err, "ContractType")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Services []*ServiceField
	fip = []uint32{ContractFieldServices}
	ServicesMin := len(a.Services)
	if ServicesMin > len(newValue.Services) {
		ServicesMin = len(newValue.Services)
	}

	// Compare values
	for i := 0; i < ServicesMin; i++ {
		lfip := append(fip, uint32(i))
		ServicesAmendments, err := a.Services[i].CreateAmendments(lfip,
			newValue.Services[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Services%d", i)
		}
		result = append(result, ServicesAmendments...)
	}

	ServicesMax := len(a.Services)
	if ServicesMax < len(newValue.Services) {
		ServicesMax = len(newValue.Services)
	}

	// Add/Remove values
	for i := ServicesMin; i < ServicesMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Services) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Services[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Services %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// AdminIdentityCertificates []*AdminIdentityCertificateField
	fip = []uint32{ContractFieldAdminIdentityCertificates}
	AdminIdentityCertificatesMin := len(a.AdminIdentityCertificates)
	if AdminIdentityCertificatesMin > len(newValue.AdminIdentityCertificates) {
		AdminIdentityCertificatesMin = len(newValue.AdminIdentityCertificates)
	}

	// Compare values
	for i := 0; i < AdminIdentityCertificatesMin; i++ {
		lfip := append(fip, uint32(i))
		AdminIdentityCertificatesAmendments, err := a.AdminIdentityCertificates[i].CreateAmendments(lfip,
			newValue.AdminIdentityCertificates[i])
		if err != nil {
			return nil, errors.Wrapf(err, "AdminIdentityCertificates%d", i)
		}
		result = append(result, AdminIdentityCertificatesAmendments...)
	}

	AdminIdentityCertificatesMax := len(a.AdminIdentityCertificates)
	if AdminIdentityCertificatesMax < len(newValue.AdminIdentityCertificates) {
		AdminIdentityCertificatesMax = len(newValue.AdminIdentityCertificates)
	}

	// Add/Remove values
	for i := AdminIdentityCertificatesMin; i < AdminIdentityCertificatesMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.AdminIdentityCertificates) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.AdminIdentityCertificates[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize AdminIdentityCertificates %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// GoverningLaw string
	fip = []uint32{ContractFieldGoverningLaw}
	if a.GoverningLaw != newValue.GoverningLaw {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.GoverningLaw),
		})
	}

	// Jurisdiction string
	fip = []uint32{ContractFieldJurisdiction}
	if a.Jurisdiction != newValue.Jurisdiction {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Jurisdiction),
		})
	}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

// ApplyAmendment updates a ContractFormation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ContractFormation) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty contract amendment field index path")
	}

	switch fip[0] {
	case ContractFieldContractName: // string
		a.ContractName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldBodyOfAgreementType: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for BodyOfAgreementType : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("BodyOfAgreementType amendment value failed to deserialize : %s", err)
		} else {
			a.BodyOfAgreementType = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldBodyOfAgreement: // []byte
		a.BodyOfAgreement = data
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedContractFieldContractType: // deprecated

	case ContractFieldSupportingDocs: // []*DocumentField
		if len(fip) == 1 && len(data) == 0 {
			a.SupportingDocs = nil
			return permissions.SubPermissions(fip[1:], operation, true)
		}
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify SupportingDocs : %v",
					fip)
			}
			if int(fip[1]) >= len(a.SupportingDocs) {
				return nil, fmt.Errorf("Amendment element index out of range for modify SupportingDocs : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.SupportingDocs[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.SupportingDocs[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add SupportingDocs : %v",
					fip)
			}

			newValue := &DocumentField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to SupportingDocs failed to deserialize : %s",
						err)
				}
			}

			if len(a.SupportingDocs) <= int(fip[1]) {
				// Append item to the end
				a.SupportingDocs = append(a.SupportingDocs, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.SupportingDocs[:fip[1]]
				after := make([]*DocumentField, len(a.SupportingDocs)-int(fip[1]))
				copy(after, a.SupportingDocs[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.SupportingDocs = append(before, newValue)
				a.SupportingDocs = append(a.SupportingDocs, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.SupportingDocs[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete SupportingDocs : %v",
					fip)
			}

			// Remove item from list
			a.SupportingDocs = append(a.SupportingDocs[:fip[1]], a.SupportingDocs[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case DeprecatedContractFieldGoverningLaw: // deprecated

	case DeprecatedContractFieldJurisdiction: // deprecated

	case ContractFieldContractExpiration: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractExpiration : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractExpiration amendment value failed to deserialize : %s", err)
		} else {
			a.ContractExpiration = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldContractURI: // string
		a.ContractURI = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldIssuer: // EntityField
		if len(fip) == 1 && len(data) == 0 {
			a.Issuer = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.Issuer.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case DeprecatedContractFieldIssuerLogoURL: // deprecated

	case DeprecatedContractFieldContractOperator: // deprecated

	case DeprecatedContractFieldAdminOracle: // deprecated

	case DeprecatedContractFieldAdminOracleSignature: // deprecated

	case DeprecatedContractFieldAdminOracleSigBlockHeight: // deprecated

	case ContractFieldContractFee: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractFee : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractFee amendment value failed to deserialize : %s", err)
		} else {
			a.ContractFee = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldVotingSystems: // []*VotingSystemField
		if len(fip) == 1 && len(data) == 0 {
			a.VotingSystems = nil
			return permissions.SubPermissions(fip[1:], operation, true)
		}
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify VotingSystems : %v",
					fip)
			}
			if int(fip[1]) >= len(a.VotingSystems) {
				return nil, fmt.Errorf("Amendment element index out of range for modify VotingSystems : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.VotingSystems[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.VotingSystems[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add VotingSystems : %v",
					fip)
			}

			newValue := &VotingSystemField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to VotingSystems failed to deserialize : %s",
						err)
				}
			}

			if len(a.VotingSystems) <= int(fip[1]) {
				// Append item to the end
				a.VotingSystems = append(a.VotingSystems, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.VotingSystems[:fip[1]]
				after := make([]*VotingSystemField, len(a.VotingSystems)-int(fip[1]))
				copy(after, a.VotingSystems[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.VotingSystems = append(before, newValue)
				a.VotingSystems = append(a.VotingSystems, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.VotingSystems[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete VotingSystems : %v",
					fip)
			}

			// Remove item from list
			a.VotingSystems = append(a.VotingSystems[:fip[1]], a.VotingSystems[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case ContractFieldContractPermissions: // []byte
		a.ContractPermissions = data
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldRestrictedQtyInstruments: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for RestrictedQtyInstruments : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("RestrictedQtyInstruments amendment value failed to deserialize : %s", err)
		} else {
			a.RestrictedQtyInstruments = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldAdministrationProposal: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AdministrationProposal : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("AdministrationProposal amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AdministrationProposal); err != nil {
			return nil, fmt.Errorf("AdministrationProposal amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldHolderProposal: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for HolderProposal : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("HolderProposal amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.HolderProposal); err != nil {
			return nil, fmt.Errorf("HolderProposal amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldOracles: // []*OracleField
		if len(fip) == 1 && len(data) == 0 {
			a.Oracles = nil
			return permissions.SubPermissions(fip[1:], operation, true)
		}
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Oracles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Oracles) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Oracles : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Oracles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Oracles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Oracles : %v",
					fip)
			}

			newValue := &OracleField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Oracles failed to deserialize : %s",
						err)
				}
			}

			if len(a.Oracles) <= int(fip[1]) {
				// Append item to the end
				a.Oracles = append(a.Oracles, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Oracles[:fip[1]]
				after := make([]*OracleField, len(a.Oracles)-int(fip[1]))
				copy(after, a.Oracles[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Oracles = append(before, newValue)
				a.Oracles = append(a.Oracles, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Oracles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Oracles : %v",
					fip)
			}

			// Remove item from list
			a.Oracles = append(a.Oracles[:fip[1]], a.Oracles[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case NotAmendableContractFieldMasterAddress: // []byte
		return nil, fmt.Errorf("MasterAddress field not amendable")

	case ContractFieldEntityContract: // []byte
		a.EntityContract = data
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldOperatorEntityContract: // []byte
		a.OperatorEntityContract = data
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldContractType: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractType : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractType amendment value failed to deserialize : %s", err)
		} else {
			a.ContractType = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldServices: // []*ServiceField
		if len(fip) == 1 && len(data) == 0 {
			a.Services = nil
			return permissions.SubPermissions(fip[1:], operation, true)
		}
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Services : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Services) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Services : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Services[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Services[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Services : %v",
					fip)
			}

			newValue := &ServiceField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Services failed to deserialize : %s",
						err)
				}
			}

			if len(a.Services) <= int(fip[1]) {
				// Append item to the end
				a.Services = append(a.Services, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Services[:fip[1]]
				after := make([]*ServiceField, len(a.Services)-int(fip[1]))
				copy(after, a.Services[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Services = append(before, newValue)
				a.Services = append(a.Services, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Services[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Services : %v",
					fip)
			}

			// Remove item from list
			a.Services = append(a.Services[:fip[1]], a.Services[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case ContractFieldAdminIdentityCertificates: // []*AdminIdentityCertificateField
		if len(fip) == 1 && len(data) == 0 {
			a.AdminIdentityCertificates = nil
			return permissions.SubPermissions(fip[1:], operation, true)
		}
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify AdminIdentityCertificates : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AdminIdentityCertificates) {
				return nil, fmt.Errorf("Amendment element index out of range for modify AdminIdentityCertificates : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.AdminIdentityCertificates[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.AdminIdentityCertificates[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add AdminIdentityCertificates : %v",
					fip)
			}

			newValue := &AdminIdentityCertificateField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to AdminIdentityCertificates failed to deserialize : %s",
						err)
				}
			}

			if len(a.AdminIdentityCertificates) <= int(fip[1]) {
				// Append item to the end
				a.AdminIdentityCertificates = append(a.AdminIdentityCertificates, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.AdminIdentityCertificates[:fip[1]]
				after := make([]*AdminIdentityCertificateField, len(a.AdminIdentityCertificates)-int(fip[1]))
				copy(after, a.AdminIdentityCertificates[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.AdminIdentityCertificates = append(before, newValue)
				a.AdminIdentityCertificates = append(a.AdminIdentityCertificates, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.AdminIdentityCertificates[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete AdminIdentityCertificates : %v",
					fip)
			}

			// Remove item from list
			a.AdminIdentityCertificates = append(a.AdminIdentityCertificates[:fip[1]], a.AdminIdentityCertificates[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case ContractFieldGoverningLaw: // string
		a.GoverningLaw = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ContractFieldJurisdiction: // string
		a.Jurisdiction = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown contract amendment field index : %v", fip)
}

// BodyOfAgreement Permission / Modification Field Indices
const (
	BodyOfAgreementFieldChapters    = uint32(1)
	BodyOfAgreementFieldDefinitions = uint32(2)
)

// CreateAmendments determines the differences between two BodyOfAgreementOffers and returns
// amendment data. Use the current value of agreement, and pass in the new values as a
// agreement definition.
func (a *BodyOfAgreementFormation) CreateAmendments(newValue *BodyOfAgreementOffer) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

	// Chapters []*ChapterField
	fip = []uint32{BodyOfAgreementFieldChapters}
	ChaptersMin := len(a.Chapters)
	if ChaptersMin > len(newValue.Chapters) {
		ChaptersMin = len(newValue.Chapters)
	}

	// Compare values
	for i := 0; i < ChaptersMin; i++ {
		lfip := append(fip, uint32(i))
		ChaptersAmendments, err := a.Chapters[i].CreateAmendments(lfip,
			newValue.Chapters[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Chapters%d", i)
		}
		result = append(result, ChaptersAmendments...)
	}

	ChaptersMax := len(a.Chapters)
	if ChaptersMax < len(newValue.Chapters) {
		ChaptersMax = len(newValue.Chapters)
	}

	// Add/Remove values
	for i := ChaptersMin; i < ChaptersMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Chapters) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Chapters[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Chapters %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// Definitions []*DefinedTermField
	fip = []uint32{BodyOfAgreementFieldDefinitions}
	DefinitionsMin := len(a.Definitions)
	if DefinitionsMin > len(newValue.Definitions) {
		DefinitionsMin = len(newValue.Definitions)
	}

	// Compare values
	for i := 0; i < DefinitionsMin; i++ {
		lfip := append(fip, uint32(i))
		DefinitionsAmendments, err := a.Definitions[i].CreateAmendments(lfip,
			newValue.Definitions[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Definitions%d", i)
		}
		result = append(result, DefinitionsAmendments...)
	}

	DefinitionsMax := len(a.Definitions)
	if DefinitionsMax < len(newValue.Definitions) {
		DefinitionsMax = len(newValue.Definitions)
	}

	// Add/Remove values
	for i := DefinitionsMin; i < DefinitionsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Definitions) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Definitions[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Definitions %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

// ApplyAmendment updates a BodyOfAgreementFormation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *BodyOfAgreementFormation) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty agreement amendment field index path")
	}

	switch fip[0] {
	case BodyOfAgreementFieldChapters: // []*ChapterField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Chapters : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Chapters) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Chapters : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Chapters[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Chapters[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Chapters : %v",
					fip)
			}

			newValue := &ChapterField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Chapters failed to deserialize : %s",
						err)
				}
			}

			if len(a.Chapters) <= int(fip[1]) {
				// Append item to the end
				a.Chapters = append(a.Chapters, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Chapters[:fip[1]]
				after := make([]*ChapterField, len(a.Chapters)-int(fip[1]))
				copy(after, a.Chapters[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Chapters = append(before, newValue)
				a.Chapters = append(a.Chapters, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Chapters[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Chapters : %v",
					fip)
			}

			// Remove item from list
			a.Chapters = append(a.Chapters[:fip[1]], a.Chapters[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case BodyOfAgreementFieldDefinitions: // []*DefinedTermField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Definitions : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Definitions) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Definitions : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Definitions[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Definitions[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Definitions : %v",
					fip)
			}

			newValue := &DefinedTermField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Definitions failed to deserialize : %s",
						err)
				}
			}

			if len(a.Definitions) <= int(fip[1]) {
				// Append item to the end
				a.Definitions = append(a.Definitions, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Definitions[:fip[1]]
				after := make([]*DefinedTermField, len(a.Definitions)-int(fip[1]))
				copy(after, a.Definitions[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Definitions = append(before, newValue)
				a.Definitions = append(a.Definitions, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Definitions[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Definitions : %v",
					fip)
			}

			// Remove item from list
			a.Definitions = append(a.Definitions[:fip[1]], a.Definitions[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown agreement amendment field index : %v", fip)
}

// Instrument Permission / Amendment Field Indices
const (
	InstrumentFieldInstrumentPermissions                 = uint32(1)
	DeprecatedInstrumentFieldTransfersPermitted          = uint32(2)
	DeprecatedInstrumentFieldTradeRestrictionsDeprecated = uint32(3)
	InstrumentFieldEnforcementOrdersPermitted            = uint32(4)
	InstrumentFieldVotingRights                          = uint32(5)
	InstrumentFieldVoteMultiplier                        = uint32(6)
	InstrumentFieldAdministrationProposal                = uint32(7)
	InstrumentFieldHolderProposal                        = uint32(8)
	InstrumentFieldInstrumentModificationGovernance      = uint32(9)
	InstrumentFieldAuthorizedTokenQty                    = uint32(10)
	InstrumentFieldInstrumentType                        = uint32(11)
	InstrumentFieldInstrumentPayload                     = uint32(12)
	InstrumentFieldTradeRestrictions                     = uint32(13)
	InstrumentFieldTransferFee                           = uint32(14)
)

// CreateAmendments determines the differences between two InstrumentDefinitions and returns
// amendment data. Use the current value of instrument creation, and pass in the new values as an instrument
// definition.
func (a *InstrumentCreation) CreateAmendments(newValue *InstrumentDefinition) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

	// InstrumentPermissions []byte
	fip = []uint32{InstrumentFieldInstrumentPermissions}
	if !bytes.Equal(a.InstrumentPermissions, newValue.InstrumentPermissions) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.InstrumentPermissions,
		})
	}

	// deprecated TransfersPermitted deprecated

	// deprecated TradeRestrictionsDeprecated deprecated

	// EnforcementOrdersPermitted bool
	fip = []uint32{InstrumentFieldEnforcementOrdersPermitted}
	if a.EnforcementOrdersPermitted != newValue.EnforcementOrdersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.EnforcementOrdersPermitted); err != nil {
			return nil, errors.Wrap(err, "EnforcementOrdersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// VotingRights bool
	fip = []uint32{InstrumentFieldVotingRights}
	if a.VotingRights != newValue.VotingRights {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.VotingRights); err != nil {
			return nil, errors.Wrap(err, "VotingRights")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// VoteMultiplier uint32
	fip = []uint32{InstrumentFieldVoteMultiplier}
	if a.VoteMultiplier != newValue.VoteMultiplier {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.VoteMultiplier)); err != nil {
			return nil, errors.Wrap(err, "VoteMultiplier")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AdministrationProposal bool
	fip = []uint32{InstrumentFieldAdministrationProposal}
	if a.AdministrationProposal != newValue.AdministrationProposal {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.AdministrationProposal); err != nil {
			return nil, errors.Wrap(err, "AdministrationProposal")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// HolderProposal bool
	fip = []uint32{InstrumentFieldHolderProposal}
	if a.HolderProposal != newValue.HolderProposal {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.HolderProposal); err != nil {
			return nil, errors.Wrap(err, "HolderProposal")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// InstrumentModificationGovernance uint32
	fip = []uint32{InstrumentFieldInstrumentModificationGovernance}
	if a.InstrumentModificationGovernance != newValue.InstrumentModificationGovernance {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.InstrumentModificationGovernance)); err != nil {
			return nil, errors.Wrap(err, "InstrumentModificationGovernance")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AuthorizedTokenQty uint64
	fip = []uint32{InstrumentFieldAuthorizedTokenQty}
	if a.AuthorizedTokenQty != newValue.AuthorizedTokenQty {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.AuthorizedTokenQty)); err != nil {
			return nil, errors.Wrap(err, "AuthorizedTokenQty")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// InstrumentType string
	fip = []uint32{InstrumentFieldInstrumentType}
	// InstrumentType modifications not allowed

	// InstrumentPayload []byte
	fip = []uint32{InstrumentFieldInstrumentPayload}
	if a.InstrumentType != newValue.InstrumentType {
		return nil, fmt.Errorf("Instrument type modification not allowed : %s -> %s", a.InstrumentType,
			newValue.InstrumentType)
	}

	payloadAmendments, err := instruments.CreatePayloadAmendments(fip, []byte(a.InstrumentType),
		a.InstrumentPayload, newValue.InstrumentPayload)
	if err != nil {
		return nil, errors.Wrap(err, "InstrumentPayload")
	}
	result = append(result, payloadAmendments...)

	// TradeRestrictions []string
	fip = []uint32{InstrumentFieldTradeRestrictions}
	TradeRestrictionsMin := len(a.TradeRestrictions)
	if TradeRestrictionsMin > len(newValue.TradeRestrictions) {
		TradeRestrictionsMin = len(newValue.TradeRestrictions)
	}

	// Compare values
	for i := 0; i < TradeRestrictionsMin; i++ {
		lfip := append(fip, uint32(i))
		if a.TradeRestrictions[i] != newValue.TradeRestrictions[i] {
			result = append(result, &internal.Amendment{
				FIP:       lfip,
				Operation: 0, // Modify element
				Data:      []byte(newValue.TradeRestrictions[i]),
			})
		}
	}

	TradeRestrictionsMax := len(a.TradeRestrictions)
	if TradeRestrictionsMax < len(newValue.TradeRestrictions) {
		TradeRestrictionsMax = len(newValue.TradeRestrictions)
	}

	// Add/Remove values
	for i := TradeRestrictionsMin; i < TradeRestrictionsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.TradeRestrictions) {
			amendment.Operation = 1 // Add element
			amendment.Data = []byte(newValue.TradeRestrictions[i])
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// TransferFee FeeField
	fip = []uint32{InstrumentFieldTransferFee}

	TransferFeeAmendments, err := a.TransferFee.CreateAmendments(fip, newValue.TransferFee)
	if err != nil {
		return nil, errors.Wrap(err, "TransferFee")
	}
	result = append(result, TransferFeeAmendments...)

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

// ApplyAmendment updates a InstrumentCreation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *InstrumentCreation) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {

	case InstrumentFieldInstrumentPermissions: // []byte
		a.InstrumentPermissions = data
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedInstrumentFieldTransfersPermitted: // deprecated

	case DeprecatedInstrumentFieldTradeRestrictionsDeprecated: // deprecated

	case InstrumentFieldEnforcementOrdersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EnforcementOrdersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("EnforcementOrdersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.EnforcementOrdersPermitted); err != nil {
			return nil, fmt.Errorf("EnforcementOrdersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldVotingRights: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for VotingRights : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("VotingRights amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.VotingRights); err != nil {
			return nil, fmt.Errorf("VotingRights amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldVoteMultiplier: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for VoteMultiplier : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("VoteMultiplier amendment value failed to deserialize : %s", err)
		} else {
			a.VoteMultiplier = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldAdministrationProposal: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AdministrationProposal : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("AdministrationProposal amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AdministrationProposal); err != nil {
			return nil, fmt.Errorf("AdministrationProposal amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldHolderProposal: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for HolderProposal : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("HolderProposal amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.HolderProposal); err != nil {
			return nil, fmt.Errorf("HolderProposal amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldInstrumentModificationGovernance: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for InstrumentModificationGovernance : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("InstrumentModificationGovernance amendment value failed to deserialize : %s", err)
		} else {
			a.InstrumentModificationGovernance = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldAuthorizedTokenQty: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AuthorizedTokenQty : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("AuthorizedTokenQty amendment value failed to deserialize : %s", err)
		} else {
			a.AuthorizedTokenQty = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldInstrumentType: // string
		a.InstrumentType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldInstrumentPayload: // []byte
		a.InstrumentPayload = data
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentFieldTradeRestrictions: // []string
		switch operation {
		case 0: // Modify
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify TradeRestrictions : %v",
					fip)
			}
			if int(fip[1]) >= len(a.TradeRestrictions) {
				return nil, fmt.Errorf("Amendment element index out of range for modify TradeRestrictions : %d", fip[1])
			}
			a.TradeRestrictions[fip[1]] = string(data)
			return permissions.SubPermissions(fip, operation, true)

		case 1: // Add element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path wrong depth for add TradeRestrictions : %v",
					fip)
			}

			newValue := string(data)

			if len(a.TradeRestrictions) <= int(fip[1]) {
				// Append item to the end
				a.TradeRestrictions = append(a.TradeRestrictions, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.TradeRestrictions[:fip[1]]
				after := make([]string, len(a.TradeRestrictions)-int(fip[1]))
				copy(after, a.TradeRestrictions[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.TradeRestrictions = append(before, newValue)
				a.TradeRestrictions = append(a.TradeRestrictions, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element

			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete TradeRestrictions : %v",
					fip)
			}
			if int(fip[1]) >= len(a.TradeRestrictions) {
				return nil, fmt.Errorf("Amendment element index out of range for delete TradeRestrictions : %d", fip[1])
			}

			// Remove item from list
			a.TradeRestrictions = append(a.TradeRestrictions[:fip[1]], a.TradeRestrictions[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case InstrumentFieldTransferFee: // FeeField
		if len(fip) == 1 && len(data) == 0 {
			a.TransferFee = nil
			return permissions.SubPermissions(fip[1:], operation, false)
		}

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.TransferFee.ApplyAmendment(fip[1:], operation, data, subPermissions)

	}

	return nil, fmt.Errorf("Unknown instrument amendment field index : %v", fip)
}

// AdministratorField Permission / Amendment Field Indices
const (
	AdministratorFieldType = uint32(1)
	AdministratorFieldName = uint32(2)
)

// ApplyAmendment updates a AdministratorField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AdministratorField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Administrator amendment field index path")
	}

	switch fip[0] {
	case AdministratorFieldType: // uint32
		if RolesData(a.Type) == nil {
			return nil, fmt.Errorf("Roles resource value not defined : %v", a.Type)
		}
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Type : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Type amendment value failed to deserialize : %s", err)
		} else {
			a.Type = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case AdministratorFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Administrator amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Administrators and returns
// amendment data.
func (a *AdministratorField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *AdministratorField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Type uint32
	fip = append(ofip, AdministratorFieldType)
	if a.Type != newValue.Type {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Type)); err != nil {
			return nil, errors.Wrap(err, "Type")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Name string
	fip = append(ofip, AdministratorFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	return result, nil
}

// AdminIdentityCertificateField Permission / Amendment Field Indices
const (
	AdminIdentityCertificateFieldEntityContract = uint32(1)
	AdminIdentityCertificateFieldSignature      = uint32(2)
	AdminIdentityCertificateFieldBlockHeight    = uint32(3)
	AdminIdentityCertificateFieldExpiration     = uint32(4)
)

// ApplyAmendment updates a AdminIdentityCertificateField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AdminIdentityCertificateField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AdminIdentityCertificate amendment field index path")
	}

	switch fip[0] {
	case AdminIdentityCertificateFieldEntityContract: // []byte
		a.EntityContract = data
		return permissions.SubPermissions(fip, operation, false)

	case AdminIdentityCertificateFieldSignature: // []byte
		a.Signature = data
		return permissions.SubPermissions(fip, operation, false)

	case AdminIdentityCertificateFieldBlockHeight: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for BlockHeight : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("BlockHeight amendment value failed to deserialize : %s", err)
		} else {
			a.BlockHeight = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case AdminIdentityCertificateFieldExpiration: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Expiration : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Expiration amendment value failed to deserialize : %s", err)
		} else {
			a.Expiration = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown AdminIdentityCertificate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AdminIdentityCertificates and returns
// amendment data.
func (a *AdminIdentityCertificateField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *AdminIdentityCertificateField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// EntityContract []byte
	fip = append(ofip, AdminIdentityCertificateFieldEntityContract)
	if !bytes.Equal(a.EntityContract, newValue.EntityContract) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.EntityContract,
		})
	}

	// Signature []byte
	fip = append(ofip, AdminIdentityCertificateFieldSignature)
	if !bytes.Equal(a.Signature, newValue.Signature) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Signature,
		})
	}

	// BlockHeight uint32
	fip = append(ofip, AdminIdentityCertificateFieldBlockHeight)
	if a.BlockHeight != newValue.BlockHeight {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.BlockHeight)); err != nil {
			return nil, errors.Wrap(err, "BlockHeight")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Expiration uint64
	fip = append(ofip, AdminIdentityCertificateFieldExpiration)
	if a.Expiration != newValue.Expiration {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Expiration)); err != nil {
			return nil, errors.Wrap(err, "Expiration")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// AmendmentField Permission / Amendment Field Indices
const (
	AmendmentFieldFieldIndexPath = uint32(1)
	AmendmentFieldOperation      = uint32(2)
	AmendmentFieldData           = uint32(3)
)

// ApplyAmendment updates a AmendmentField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AmendmentField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Amendment amendment field index path")
	}

	switch fip[0] {
	case AmendmentFieldFieldIndexPath: // []byte
		a.FieldIndexPath = data
		return permissions.SubPermissions(fip, operation, false)

	case AmendmentFieldOperation: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Operation : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Operation amendment value failed to deserialize : %s", err)
		} else {
			a.Operation = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case AmendmentFieldData: // []byte
		a.Data = data
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Amendment amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Amendments and returns
// amendment data.
func (a *AmendmentField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *AmendmentField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// FieldIndexPath []byte
	fip = append(ofip, AmendmentFieldFieldIndexPath)
	if !bytes.Equal(a.FieldIndexPath, newValue.FieldIndexPath) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.FieldIndexPath,
		})
	}

	// Operation uint32
	fip = append(ofip, AmendmentFieldOperation)
	if a.Operation != newValue.Operation {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Operation)); err != nil {
			return nil, errors.Wrap(err, "Operation")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Data []byte
	fip = append(ofip, AmendmentFieldData)
	if !bytes.Equal(a.Data, newValue.Data) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Data,
		})
	}

	return result, nil
}

// InstrumentReceiverField Permission / Amendment Field Indices
const (
	InstrumentReceiverFieldAddress               = uint32(1)
	InstrumentReceiverFieldQuantity              = uint32(2)
	InstrumentReceiverFieldOracleSigAlgorithm    = uint32(3)
	InstrumentReceiverFieldOracleIndex           = uint32(4)
	InstrumentReceiverFieldOracleConfirmationSig = uint32(5)
	InstrumentReceiverFieldOracleSigBlockHeight  = uint32(6)
	InstrumentReceiverFieldOracleSigExpiry       = uint32(7)
)

// ApplyAmendment updates a InstrumentReceiverField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *InstrumentReceiverField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty InstrumentReceiver amendment field index path")
	}

	switch fip[0] {
	case InstrumentReceiverFieldAddress: // []byte
		a.Address = data
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		} else {
			a.Quantity = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldOracleSigAlgorithm: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigAlgorithm : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigAlgorithm amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigAlgorithm = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldOracleIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleIndex amendment value failed to deserialize : %s", err)
		} else {
			a.OracleIndex = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldOracleConfirmationSig: // []byte
		a.OracleConfirmationSig = data
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldOracleSigBlockHeight: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigBlockHeight : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigBlockHeight amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigBlockHeight = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentReceiverFieldOracleSigExpiry: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigExpiry : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigExpiry amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigExpiry = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown InstrumentReceiver amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two InstrumentReceivers and returns
// amendment data.
func (a *InstrumentReceiverField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *InstrumentReceiverField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Address []byte
	fip = append(ofip, InstrumentReceiverFieldAddress)
	if !bytes.Equal(a.Address, newValue.Address) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Address,
		})
	}

	// Quantity uint64
	fip = append(ofip, InstrumentReceiverFieldQuantity)
	if a.Quantity != newValue.Quantity {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Quantity)); err != nil {
			return nil, errors.Wrap(err, "Quantity")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// OracleSigAlgorithm uint32
	fip = append(ofip, InstrumentReceiverFieldOracleSigAlgorithm)
	if a.OracleSigAlgorithm != newValue.OracleSigAlgorithm {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleSigAlgorithm)); err != nil {
			return nil, errors.Wrap(err, "OracleSigAlgorithm")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// OracleIndex uint32
	fip = append(ofip, InstrumentReceiverFieldOracleIndex)
	if a.OracleIndex != newValue.OracleIndex {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleIndex)); err != nil {
			return nil, errors.Wrap(err, "OracleIndex")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// OracleConfirmationSig []byte
	fip = append(ofip, InstrumentReceiverFieldOracleConfirmationSig)
	if !bytes.Equal(a.OracleConfirmationSig, newValue.OracleConfirmationSig) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.OracleConfirmationSig,
		})
	}

	// OracleSigBlockHeight uint32
	fip = append(ofip, InstrumentReceiverFieldOracleSigBlockHeight)
	if a.OracleSigBlockHeight != newValue.OracleSigBlockHeight {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleSigBlockHeight)); err != nil {
			return nil, errors.Wrap(err, "OracleSigBlockHeight")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// OracleSigExpiry uint64
	fip = append(ofip, InstrumentReceiverFieldOracleSigExpiry)
	if a.OracleSigExpiry != newValue.OracleSigExpiry {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleSigExpiry)); err != nil {
			return nil, errors.Wrap(err, "OracleSigExpiry")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// InstrumentSettlementField Permission / Amendment Field Indices
const (
	InstrumentSettlementFieldContractIndex  = uint32(1)
	InstrumentSettlementFieldInstrumentType = uint32(2)
	InstrumentSettlementFieldInstrumentCode = uint32(3)
	InstrumentSettlementFieldSettlements    = uint32(4)
)

// ApplyAmendment updates a InstrumentSettlementField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *InstrumentSettlementField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty InstrumentSettlement amendment field index path")
	}

	switch fip[0] {
	case InstrumentSettlementFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		} else {
			a.ContractIndex = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentSettlementFieldInstrumentType: // string
		a.InstrumentType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentSettlementFieldInstrumentCode: // []byte
		if len(data) != 20 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 20)
		}
		copy(a.InstrumentCode, data)
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentSettlementFieldSettlements: // []*QuantityIndexField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Settlements : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Settlements) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Settlements : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Settlements[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Settlements[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Settlements : %v",
					fip)
			}

			newValue := &QuantityIndexField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Settlements failed to deserialize : %s",
						err)
				}
			}

			if len(a.Settlements) <= int(fip[1]) {
				// Append item to the end
				a.Settlements = append(a.Settlements, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Settlements[:fip[1]]
				after := make([]*QuantityIndexField, len(a.Settlements)-int(fip[1]))
				copy(after, a.Settlements[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Settlements = append(before, newValue)
				a.Settlements = append(a.Settlements, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Settlements[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Settlements : %v",
					fip)
			}

			// Remove item from list
			a.Settlements = append(a.Settlements[:fip[1]], a.Settlements[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown InstrumentSettlement amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two InstrumentSettlements and returns
// amendment data.
func (a *InstrumentSettlementField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *InstrumentSettlementField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// ContractIndex uint32
	fip = append(ofip, InstrumentSettlementFieldContractIndex)
	if a.ContractIndex != newValue.ContractIndex {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ContractIndex)); err != nil {
			return nil, errors.Wrap(err, "ContractIndex")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// InstrumentType string
	fip = append(ofip, InstrumentSettlementFieldInstrumentType)
	// InstrumentType modifications not allowed

	// InstrumentCode []byte
	fip = append(ofip, InstrumentSettlementFieldInstrumentCode)
	if !bytes.Equal(a.InstrumentCode[:], newValue.InstrumentCode[:]) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.InstrumentCode[:],
		})
	}

	// Settlements []*QuantityIndexField
	fip = append(ofip, InstrumentSettlementFieldSettlements)
	SettlementsMin := len(a.Settlements)
	if SettlementsMin > len(newValue.Settlements) {
		SettlementsMin = len(newValue.Settlements)
	}

	// Compare values
	for i := 0; i < SettlementsMin; i++ {
		lfip := append(fip, uint32(i))
		SettlementsAmendments, err := a.Settlements[i].CreateAmendments(lfip,
			newValue.Settlements[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Settlements%d", i)
		}
		result = append(result, SettlementsAmendments...)
	}

	SettlementsMax := len(a.Settlements)
	if SettlementsMax < len(newValue.Settlements) {
		SettlementsMax = len(newValue.Settlements)
	}

	// Add/Remove values
	for i := SettlementsMin; i < SettlementsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Settlements) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Settlements[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Settlements %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	return result, nil
}

// InstrumentTransferField Permission / Amendment Field Indices
const (
	InstrumentTransferFieldContractIndex       = uint32(1)
	InstrumentTransferFieldInstrumentType      = uint32(2)
	InstrumentTransferFieldInstrumentCode      = uint32(3)
	InstrumentTransferFieldInstrumentSenders   = uint32(4)
	InstrumentTransferFieldInstrumentReceivers = uint32(5)
)

// ApplyAmendment updates a InstrumentTransferField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *InstrumentTransferField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty InstrumentTransfer amendment field index path")
	}

	switch fip[0] {
	case InstrumentTransferFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		} else {
			a.ContractIndex = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentTransferFieldInstrumentType: // string
		a.InstrumentType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentTransferFieldInstrumentCode: // []byte
		if len(data) != 20 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 20)
		}
		copy(a.InstrumentCode, data)
		return permissions.SubPermissions(fip, operation, false)

	case InstrumentTransferFieldInstrumentSenders: // []*QuantityIndexField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify InstrumentSenders : %v",
					fip)
			}
			if int(fip[1]) >= len(a.InstrumentSenders) {
				return nil, fmt.Errorf("Amendment element index out of range for modify InstrumentSenders : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.InstrumentSenders[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.InstrumentSenders[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add InstrumentSenders : %v",
					fip)
			}

			newValue := &QuantityIndexField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to InstrumentSenders failed to deserialize : %s",
						err)
				}
			}

			if len(a.InstrumentSenders) <= int(fip[1]) {
				// Append item to the end
				a.InstrumentSenders = append(a.InstrumentSenders, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.InstrumentSenders[:fip[1]]
				after := make([]*QuantityIndexField, len(a.InstrumentSenders)-int(fip[1]))
				copy(after, a.InstrumentSenders[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.InstrumentSenders = append(before, newValue)
				a.InstrumentSenders = append(a.InstrumentSenders, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.InstrumentSenders[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete InstrumentSenders : %v",
					fip)
			}

			// Remove item from list
			a.InstrumentSenders = append(a.InstrumentSenders[:fip[1]], a.InstrumentSenders[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case InstrumentTransferFieldInstrumentReceivers: // []*InstrumentReceiverField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify InstrumentReceivers : %v",
					fip)
			}
			if int(fip[1]) >= len(a.InstrumentReceivers) {
				return nil, fmt.Errorf("Amendment element index out of range for modify InstrumentReceivers : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.InstrumentReceivers[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.InstrumentReceivers[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add InstrumentReceivers : %v",
					fip)
			}

			newValue := &InstrumentReceiverField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to InstrumentReceivers failed to deserialize : %s",
						err)
				}
			}

			if len(a.InstrumentReceivers) <= int(fip[1]) {
				// Append item to the end
				a.InstrumentReceivers = append(a.InstrumentReceivers, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.InstrumentReceivers[:fip[1]]
				after := make([]*InstrumentReceiverField, len(a.InstrumentReceivers)-int(fip[1]))
				copy(after, a.InstrumentReceivers[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.InstrumentReceivers = append(before, newValue)
				a.InstrumentReceivers = append(a.InstrumentReceivers, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.InstrumentReceivers[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete InstrumentReceivers : %v",
					fip)
			}

			// Remove item from list
			a.InstrumentReceivers = append(a.InstrumentReceivers[:fip[1]], a.InstrumentReceivers[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown InstrumentTransfer amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two InstrumentTransfers and returns
// amendment data.
func (a *InstrumentTransferField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *InstrumentTransferField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// ContractIndex uint32
	fip = append(ofip, InstrumentTransferFieldContractIndex)
	if a.ContractIndex != newValue.ContractIndex {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ContractIndex)); err != nil {
			return nil, errors.Wrap(err, "ContractIndex")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// InstrumentType string
	fip = append(ofip, InstrumentTransferFieldInstrumentType)
	// InstrumentType modifications not allowed

	// InstrumentCode []byte
	fip = append(ofip, InstrumentTransferFieldInstrumentCode)
	if !bytes.Equal(a.InstrumentCode[:], newValue.InstrumentCode[:]) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.InstrumentCode[:],
		})
	}

	// InstrumentSenders []*QuantityIndexField
	fip = append(ofip, InstrumentTransferFieldInstrumentSenders)
	InstrumentSendersMin := len(a.InstrumentSenders)
	if InstrumentSendersMin > len(newValue.InstrumentSenders) {
		InstrumentSendersMin = len(newValue.InstrumentSenders)
	}

	// Compare values
	for i := 0; i < InstrumentSendersMin; i++ {
		lfip := append(fip, uint32(i))
		InstrumentSendersAmendments, err := a.InstrumentSenders[i].CreateAmendments(lfip,
			newValue.InstrumentSenders[i])
		if err != nil {
			return nil, errors.Wrapf(err, "InstrumentSenders%d", i)
		}
		result = append(result, InstrumentSendersAmendments...)
	}

	InstrumentSendersMax := len(a.InstrumentSenders)
	if InstrumentSendersMax < len(newValue.InstrumentSenders) {
		InstrumentSendersMax = len(newValue.InstrumentSenders)
	}

	// Add/Remove values
	for i := InstrumentSendersMin; i < InstrumentSendersMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.InstrumentSenders) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.InstrumentSenders[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize InstrumentSenders %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// InstrumentReceivers []*InstrumentReceiverField
	fip = append(ofip, InstrumentTransferFieldInstrumentReceivers)
	InstrumentReceiversMin := len(a.InstrumentReceivers)
	if InstrumentReceiversMin > len(newValue.InstrumentReceivers) {
		InstrumentReceiversMin = len(newValue.InstrumentReceivers)
	}

	// Compare values
	for i := 0; i < InstrumentReceiversMin; i++ {
		lfip := append(fip, uint32(i))
		InstrumentReceiversAmendments, err := a.InstrumentReceivers[i].CreateAmendments(lfip,
			newValue.InstrumentReceivers[i])
		if err != nil {
			return nil, errors.Wrapf(err, "InstrumentReceivers%d", i)
		}
		result = append(result, InstrumentReceiversAmendments...)
	}

	InstrumentReceiversMax := len(a.InstrumentReceivers)
	if InstrumentReceiversMax < len(newValue.InstrumentReceivers) {
		InstrumentReceiversMax = len(newValue.InstrumentReceivers)
	}

	// Add/Remove values
	for i := InstrumentReceiversMin; i < InstrumentReceiversMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.InstrumentReceivers) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.InstrumentReceivers[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize InstrumentReceivers %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	return result, nil
}

// ChapterField Permission / Amendment Field Indices
const (
	ChapterFieldTitle    = uint32(1)
	ChapterFieldPreamble = uint32(2)
	ChapterFieldArticles = uint32(3)
)

// ApplyAmendment updates a ChapterField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ChapterField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Chapter amendment field index path")
	}

	switch fip[0] {
	case ChapterFieldTitle: // string
		a.Title = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ChapterFieldPreamble: // string
		a.Preamble = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ChapterFieldArticles: // []*ClauseField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Articles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Articles) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Articles : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Articles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Articles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Articles : %v",
					fip)
			}

			newValue := &ClauseField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Articles failed to deserialize : %s",
						err)
				}
			}

			if len(a.Articles) <= int(fip[1]) {
				// Append item to the end
				a.Articles = append(a.Articles, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Articles[:fip[1]]
				after := make([]*ClauseField, len(a.Articles)-int(fip[1]))
				copy(after, a.Articles[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Articles = append(before, newValue)
				a.Articles = append(a.Articles, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Articles[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Articles : %v",
					fip)
			}

			// Remove item from list
			a.Articles = append(a.Articles[:fip[1]], a.Articles[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown Chapter amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Chapters and returns
// amendment data.
func (a *ChapterField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ChapterField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Title string
	fip = append(ofip, ChapterFieldTitle)
	if a.Title != newValue.Title {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Title),
		})
	}

	// Preamble string
	fip = append(ofip, ChapterFieldPreamble)
	if a.Preamble != newValue.Preamble {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Preamble),
		})
	}

	// Articles []*ClauseField
	fip = append(ofip, ChapterFieldArticles)
	ArticlesMin := len(a.Articles)
	if ArticlesMin > len(newValue.Articles) {
		ArticlesMin = len(newValue.Articles)
	}

	// Compare values
	for i := 0; i < ArticlesMin; i++ {
		lfip := append(fip, uint32(i))
		ArticlesAmendments, err := a.Articles[i].CreateAmendments(lfip,
			newValue.Articles[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Articles%d", i)
		}
		result = append(result, ArticlesAmendments...)
	}

	ArticlesMax := len(a.Articles)
	if ArticlesMax < len(newValue.Articles) {
		ArticlesMax = len(newValue.Articles)
	}

	// Add/Remove values
	for i := ArticlesMin; i < ArticlesMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Articles) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Articles[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Articles %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	return result, nil
}

// ClauseField Permission / Amendment Field Indices
const (
	ClauseFieldTitle    = uint32(1)
	ClauseFieldBody     = uint32(2)
	ClauseFieldChildren = uint32(3)
)

// ApplyAmendment updates a ClauseField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ClauseField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Clause amendment field index path")
	}

	switch fip[0] {
	case ClauseFieldTitle: // string
		a.Title = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ClauseFieldBody: // string
		a.Body = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ClauseFieldChildren: // []*ClauseField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Children : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Children) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Children : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Children[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Children[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Children : %v",
					fip)
			}

			newValue := &ClauseField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Children failed to deserialize : %s",
						err)
				}
			}

			if len(a.Children) <= int(fip[1]) {
				// Append item to the end
				a.Children = append(a.Children, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Children[:fip[1]]
				after := make([]*ClauseField, len(a.Children)-int(fip[1]))
				copy(after, a.Children[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Children = append(before, newValue)
				a.Children = append(a.Children, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Children[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Children : %v",
					fip)
			}

			// Remove item from list
			a.Children = append(a.Children[:fip[1]], a.Children[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown Clause amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Clauses and returns
// amendment data.
func (a *ClauseField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ClauseField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Title string
	fip = append(ofip, ClauseFieldTitle)
	if a.Title != newValue.Title {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Title),
		})
	}

	// Body string
	fip = append(ofip, ClauseFieldBody)
	if a.Body != newValue.Body {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Body),
		})
	}

	// Children []*ClauseField
	fip = append(ofip, ClauseFieldChildren)
	ChildrenMin := len(a.Children)
	if ChildrenMin > len(newValue.Children) {
		ChildrenMin = len(newValue.Children)
	}

	// Compare values
	for i := 0; i < ChildrenMin; i++ {
		lfip := append(fip, uint32(i))
		ChildrenAmendments, err := a.Children[i].CreateAmendments(lfip,
			newValue.Children[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Children%d", i)
		}
		result = append(result, ChildrenAmendments...)
	}

	ChildrenMax := len(a.Children)
	if ChildrenMax < len(newValue.Children) {
		ChildrenMax = len(newValue.Children)
	}

	// Add/Remove values
	for i := ChildrenMin; i < ChildrenMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Children) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Children[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Children %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	return result, nil
}

// DefinedTermField Permission / Amendment Field Indices
const (
	DefinedTermFieldTerm       = uint32(1)
	DefinedTermFieldDefinition = uint32(2)
)

// ApplyAmendment updates a DefinedTermField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *DefinedTermField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty DefinedTerm amendment field index path")
	}

	switch fip[0] {
	case DefinedTermFieldTerm: // string
		a.Term = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DefinedTermFieldDefinition: // string
		a.Definition = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown DefinedTerm amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two DefinedTerms and returns
// amendment data.
func (a *DefinedTermField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *DefinedTermField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Term string
	fip = append(ofip, DefinedTermFieldTerm)
	if a.Term != newValue.Term {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Term),
		})
	}

	// Definition string
	fip = append(ofip, DefinedTermFieldDefinition)
	if a.Definition != newValue.Definition {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Definition),
		})
	}

	return result, nil
}

// DocumentField Permission / Amendment Field Indices
const (
	DocumentFieldName     = uint32(1)
	DocumentFieldType     = uint32(2)
	DocumentFieldContents = uint32(3)
)

// ApplyAmendment updates a DocumentField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *DocumentField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Document amendment field index path")
	}

	switch fip[0] {
	case DocumentFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DocumentFieldType: // string
		a.Type = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DocumentFieldContents: // []byte
		a.Contents = data
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Document amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Documents and returns
// amendment data.
func (a *DocumentField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *DocumentField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Name string
	fip = append(ofip, DocumentFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	// Type string
	fip = append(ofip, DocumentFieldType)
	if a.Type != newValue.Type {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Type),
		})
	}

	// Contents []byte
	fip = append(ofip, DocumentFieldContents)
	if !bytes.Equal(a.Contents, newValue.Contents) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Contents,
		})
	}

	return result, nil
}

// EntityField Permission / Amendment Field Indices
const (
	EntityFieldName                            = uint32(1)
	EntityFieldType                            = uint32(2)
	EntityFieldLEI                             = uint32(3)
	EntityFieldUnitNumber                      = uint32(4)
	EntityFieldBuildingNumber                  = uint32(5)
	EntityFieldStreet                          = uint32(6)
	EntityFieldSuburbCity                      = uint32(7)
	EntityFieldTerritoryStateProvinceCode      = uint32(8)
	EntityFieldCountryCode                     = uint32(9)
	EntityFieldPostalZIPCode                   = uint32(10)
	EntityFieldEmailAddress                    = uint32(11)
	EntityFieldPhoneNumber                     = uint32(12)
	EntityFieldAdministration                  = uint32(13)
	EntityFieldManagement                      = uint32(14)
	EntityFieldDomainName                      = uint32(15)
	DeprecatedEntityFieldEntityContractAddress = uint32(16)
	EntityFieldPaymailHandle                   = uint32(17)
)

// ApplyAmendment updates a EntityField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *EntityField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Entity amendment field index path")
	}

	switch fip[0] {
	case EntityFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldType: // string
		if EntitiesData(a.Type) == nil {
			return nil, fmt.Errorf("Entities resource value not defined : %v", a.Type)
		}
		a.Type = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldLEI: // string
		a.LEI = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldUnitNumber: // string
		a.UnitNumber = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldBuildingNumber: // string
		a.BuildingNumber = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldStreet: // string
		a.Street = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldSuburbCity: // string
		a.SuburbCity = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldTerritoryStateProvinceCode: // string
		a.TerritoryStateProvinceCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldCountryCode: // string
		a.CountryCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldPostalZIPCode: // string
		a.PostalZIPCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldEmailAddress: // string
		a.EmailAddress = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldPhoneNumber: // string
		a.PhoneNumber = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case EntityFieldAdministration: // []*AdministratorField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Administration : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Administration) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Administration : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Administration[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Administration[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Administration : %v",
					fip)
			}

			newValue := &AdministratorField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Administration failed to deserialize : %s",
						err)
				}
			}

			if len(a.Administration) <= int(fip[1]) {
				// Append item to the end
				a.Administration = append(a.Administration, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Administration[:fip[1]]
				after := make([]*AdministratorField, len(a.Administration)-int(fip[1]))
				copy(after, a.Administration[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Administration = append(before, newValue)
				a.Administration = append(a.Administration, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Administration[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Administration : %v",
					fip)
			}

			// Remove item from list
			a.Administration = append(a.Administration[:fip[1]], a.Administration[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case EntityFieldManagement: // []*ManagerField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Management : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Management) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Management : %d", fip[1])
			}

			subPermissions, err := permissions.SubPermissions(fip, operation, true)
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.Management[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

		case 1: // Add element
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Management[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Management : %v",
					fip)
			}

			newValue := &ManagerField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Management failed to deserialize : %s",
						err)
				}
			}

			if len(a.Management) <= int(fip[1]) {
				// Append item to the end
				a.Management = append(a.Management, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Management[:fip[1]]
				after := make([]*ManagerField, len(a.Management)-int(fip[1]))
				copy(after, a.Management[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Management = append(before, newValue)
				a.Management = append(a.Management, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element
			if len(fip) > 2 { // includes list index
				// Delete is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, true)
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.Management[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for delete Management : %v",
					fip)
			}

			// Remove item from list
			a.Management = append(a.Management[:fip[1]], a.Management[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case EntityFieldDomainName: // string
		a.DomainName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedEntityFieldEntityContractAddress: // deprecated

	case EntityFieldPaymailHandle: // string
		a.PaymailHandle = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Entity amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Entitys and returns
// amendment data.
func (a *EntityField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *EntityField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Name string
	fip = append(ofip, EntityFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	// Type string
	fip = append(ofip, EntityFieldType)
	if a.Type != newValue.Type {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Type),
		})
	}

	// LEI string
	fip = append(ofip, EntityFieldLEI)
	if a.LEI != newValue.LEI {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.LEI),
		})
	}

	// UnitNumber string
	fip = append(ofip, EntityFieldUnitNumber)
	if a.UnitNumber != newValue.UnitNumber {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.UnitNumber),
		})
	}

	// BuildingNumber string
	fip = append(ofip, EntityFieldBuildingNumber)
	if a.BuildingNumber != newValue.BuildingNumber {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.BuildingNumber),
		})
	}

	// Street string
	fip = append(ofip, EntityFieldStreet)
	if a.Street != newValue.Street {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Street),
		})
	}

	// SuburbCity string
	fip = append(ofip, EntityFieldSuburbCity)
	if a.SuburbCity != newValue.SuburbCity {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.SuburbCity),
		})
	}

	// TerritoryStateProvinceCode string
	fip = append(ofip, EntityFieldTerritoryStateProvinceCode)
	if a.TerritoryStateProvinceCode != newValue.TerritoryStateProvinceCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.TerritoryStateProvinceCode),
		})
	}

	// CountryCode string
	fip = append(ofip, EntityFieldCountryCode)
	if a.CountryCode != newValue.CountryCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CountryCode),
		})
	}

	// PostalZIPCode string
	fip = append(ofip, EntityFieldPostalZIPCode)
	if a.PostalZIPCode != newValue.PostalZIPCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.PostalZIPCode),
		})
	}

	// EmailAddress string
	fip = append(ofip, EntityFieldEmailAddress)
	if a.EmailAddress != newValue.EmailAddress {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.EmailAddress),
		})
	}

	// PhoneNumber string
	fip = append(ofip, EntityFieldPhoneNumber)
	if a.PhoneNumber != newValue.PhoneNumber {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.PhoneNumber),
		})
	}

	// Administration []*AdministratorField
	fip = append(ofip, EntityFieldAdministration)
	AdministrationMin := len(a.Administration)
	if AdministrationMin > len(newValue.Administration) {
		AdministrationMin = len(newValue.Administration)
	}

	// Compare values
	for i := 0; i < AdministrationMin; i++ {
		lfip := append(fip, uint32(i))
		AdministrationAmendments, err := a.Administration[i].CreateAmendments(lfip,
			newValue.Administration[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Administration%d", i)
		}
		result = append(result, AdministrationAmendments...)
	}

	AdministrationMax := len(a.Administration)
	if AdministrationMax < len(newValue.Administration) {
		AdministrationMax = len(newValue.Administration)
	}

	// Add/Remove values
	for i := AdministrationMin; i < AdministrationMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Administration) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Administration[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Administration %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// Management []*ManagerField
	fip = append(ofip, EntityFieldManagement)
	ManagementMin := len(a.Management)
	if ManagementMin > len(newValue.Management) {
		ManagementMin = len(newValue.Management)
	}

	// Compare values
	for i := 0; i < ManagementMin; i++ {
		lfip := append(fip, uint32(i))
		ManagementAmendments, err := a.Management[i].CreateAmendments(lfip,
			newValue.Management[i])
		if err != nil {
			return nil, errors.Wrapf(err, "Management%d", i)
		}
		result = append(result, ManagementAmendments...)
	}

	ManagementMax := len(a.Management)
	if ManagementMax < len(newValue.Management) {
		ManagementMax = len(newValue.Management)
	}

	// Add/Remove values
	for i := ManagementMin; i < ManagementMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Management) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.Management[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize Management %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// DomainName string
	fip = append(ofip, EntityFieldDomainName)
	if a.DomainName != newValue.DomainName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.DomainName),
		})
	}

	// deprecated EntityContractAddress deprecated

	// PaymailHandle string
	fip = append(ofip, EntityFieldPaymailHandle)
	if a.PaymailHandle != newValue.PaymailHandle {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.PaymailHandle),
		})
	}

	return result, nil
}

// FeeField Permission / Amendment Field Indices
const (
	FeeFieldQuantity = uint32(1)
)

// ApplyAmendment updates a FeeField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *FeeField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Fee amendment field index path")
	}

	switch fip[0] {
	case FeeFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		} else {
			a.Quantity = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Fee amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Fees and returns
// amendment data.
func (a *FeeField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *FeeField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Quantity uint64
	fip = append(ofip, FeeFieldQuantity)
	if a.Quantity != newValue.Quantity {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Quantity)); err != nil {
			return nil, errors.Wrap(err, "Quantity")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// ManagerField Permission / Amendment Field Indices
const (
	ManagerFieldType = uint32(1)
	ManagerFieldName = uint32(2)
)

// ApplyAmendment updates a ManagerField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ManagerField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Manager amendment field index path")
	}

	switch fip[0] {
	case ManagerFieldType: // uint32
		if RolesData(a.Type) == nil {
			return nil, fmt.Errorf("Roles resource value not defined : %v", a.Type)
		}
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Type : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Type amendment value failed to deserialize : %s", err)
		} else {
			a.Type = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ManagerFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Manager amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Managers and returns
// amendment data.
func (a *ManagerField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ManagerField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Type uint32
	fip = append(ofip, ManagerFieldType)
	if a.Type != newValue.Type {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Type)); err != nil {
			return nil, errors.Wrap(err, "Type")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Name string
	fip = append(ofip, ManagerFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	return result, nil
}

// OracleField Permission / Amendment Field Indices
const (
	DeprecatedOracleFieldEntity    = uint32(1)
	DeprecatedOracleFieldURL       = uint32(2)
	DeprecatedOracleFieldPublicKey = uint32(3)
	OracleFieldOracleTypes         = uint32(4)
	OracleFieldEntityContract      = uint32(5)
)

// ApplyAmendment updates a OracleField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *OracleField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Oracle amendment field index path")
	}

	switch fip[0] {
	case DeprecatedOracleFieldEntity: // deprecated

	case DeprecatedOracleFieldURL: // deprecated

	case DeprecatedOracleFieldPublicKey: // deprecated

	case OracleFieldOracleTypes: // []uint32
		switch operation {
		case 0: // Modify
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify OracleTypes : %v",
					fip)
			}
			if int(fip[1]) >= len(a.OracleTypes) {
				return nil, fmt.Errorf("Amendment element index out of range for modify OracleTypes : %d", fip[1])
			}
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("OracleTypes amendment value failed to deserialize : %s", err)
			} else {
				a.OracleTypes[fip[1]] = uint32(value)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 1: // Add element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path wrong depth for add OracleTypes : %v",
					fip)
			}

			var newValue uint32
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("OracleTypes amendment value failed to deserialize : %s",
					err)
			} else {
				newValue = uint32(value)
			}

			if len(a.OracleTypes) <= int(fip[1]) {
				// Append item to the end
				a.OracleTypes = append(a.OracleTypes, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.OracleTypes[:fip[1]]
				after := make([]uint32, len(a.OracleTypes)-int(fip[1]))
				copy(after, a.OracleTypes[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.OracleTypes = append(before, newValue)
				a.OracleTypes = append(a.OracleTypes, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element

			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete OracleTypes : %v",
					fip)
			}
			if int(fip[1]) >= len(a.OracleTypes) {
				return nil, fmt.Errorf("Amendment element index out of range for delete OracleTypes : %d", fip[1])
			}

			// Remove item from list
			a.OracleTypes = append(a.OracleTypes[:fip[1]], a.OracleTypes[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	case OracleFieldEntityContract: // []byte
		a.EntityContract = data
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Oracle amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Oracles and returns
// amendment data.
func (a *OracleField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *OracleField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// deprecated Entity deprecated

	// deprecated URL deprecated

	// deprecated PublicKey deprecated

	// OracleTypes []uint32
	fip = append(ofip, OracleFieldOracleTypes)
	OracleTypesMin := len(a.OracleTypes)
	if OracleTypesMin > len(newValue.OracleTypes) {
		OracleTypesMin = len(newValue.OracleTypes)
	}

	// Compare values
	for i := 0; i < OracleTypesMin; i++ {
		lfip := append(fip, uint32(i))
		if a.OracleTypes[i] != newValue.OracleTypes[i] {
			var buf bytes.Buffer
			if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleTypes[i])); err != nil {
				return nil, errors.Wrapf(err, "OracleTypes %d", i)
			}

			result = append(result, &internal.Amendment{
				FIP:       lfip,
				Operation: 0, // Modify element
				Data:      buf.Bytes(),
			})
		}
	}

	OracleTypesMax := len(a.OracleTypes)
	if OracleTypesMax < len(newValue.OracleTypes) {
		OracleTypesMax = len(newValue.OracleTypes)
	}

	// Add/Remove values
	for i := OracleTypesMin; i < OracleTypesMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.OracleTypes) {
			amendment.Operation = 1 // Add element
			var buf bytes.Buffer
			if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.OracleTypes[i])); err != nil {
				return nil, errors.Wrapf(err, "OracleTypes %d", i)
			}
			amendment.Data = buf.Bytes()
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// EntityContract []byte
	fip = append(ofip, OracleFieldEntityContract)
	if !bytes.Equal(a.EntityContract, newValue.EntityContract) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.EntityContract,
		})
	}

	return result, nil
}

// QuantityIndexField Permission / Amendment Field Indices
const (
	QuantityIndexFieldIndex    = uint32(1)
	QuantityIndexFieldQuantity = uint32(2)
)

// ApplyAmendment updates a QuantityIndexField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *QuantityIndexField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty QuantityIndex amendment field index path")
	}

	switch fip[0] {
	case QuantityIndexFieldIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Index : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Index amendment value failed to deserialize : %s", err)
		} else {
			a.Index = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case QuantityIndexFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		} else {
			a.Quantity = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown QuantityIndex amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two QuantityIndexs and returns
// amendment data.
func (a *QuantityIndexField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *QuantityIndexField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Index uint32
	fip = append(ofip, QuantityIndexFieldIndex)
	if a.Index != newValue.Index {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Index)); err != nil {
			return nil, errors.Wrap(err, "Index")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Quantity uint64
	fip = append(ofip, QuantityIndexFieldQuantity)
	if a.Quantity != newValue.Quantity {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Quantity)); err != nil {
			return nil, errors.Wrap(err, "Quantity")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// ReferenceTransactionField Permission / Amendment Field Indices
const (
	ReferenceTransactionFieldTransaction = uint32(1)
	ReferenceTransactionFieldOutputs     = uint32(2)
)

// ApplyAmendment updates a ReferenceTransactionField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ReferenceTransactionField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty ReferenceTransaction amendment field index path")
	}

	switch fip[0] {
	case ReferenceTransactionFieldTransaction: // []byte
		a.Transaction = data
		return permissions.SubPermissions(fip, operation, false)

	case ReferenceTransactionFieldOutputs: // [][]byte
		switch operation {
		case 0: // Modify
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Outputs : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Outputs) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Outputs : %d", fip[1])
			}
			a.Outputs[fip[1]] = data
			return permissions.SubPermissions(fip, operation, true)

		case 1: // Add element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path wrong depth for add Outputs : %v",
					fip)
			}

			newValue := data

			if len(a.Outputs) <= int(fip[1]) {
				// Append item to the end
				a.Outputs = append(a.Outputs, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.Outputs[:fip[1]]
				after := make([][]byte, len(a.Outputs)-int(fip[1]))
				copy(after, a.Outputs[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.Outputs = append(before, newValue)
				a.Outputs = append(a.Outputs, after...)
			}
			return permissions.SubPermissions(fip, operation, true)

		case 2: // Delete element

			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Outputs : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Outputs) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Outputs : %d", fip[1])
			}

			// Remove item from list
			a.Outputs = append(a.Outputs[:fip[1]], a.Outputs[fip[1]+1:]...)
			return permissions.SubPermissions(fip, operation, true)
		}

	}

	return nil, fmt.Errorf("Unknown ReferenceTransaction amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two ReferenceTransactions and returns
// amendment data.
func (a *ReferenceTransactionField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ReferenceTransactionField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Transaction []byte
	fip = append(ofip, ReferenceTransactionFieldTransaction)
	if !bytes.Equal(a.Transaction, newValue.Transaction) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Transaction,
		})
	}

	// Outputs [][]byte
	fip = append(ofip, ReferenceTransactionFieldOutputs)
	OutputsMin := len(a.Outputs)
	if OutputsMin > len(newValue.Outputs) {
		OutputsMin = len(newValue.Outputs)
	}

	// Compare values
	for i := 0; i < OutputsMin; i++ {
		lfip := append(fip, uint32(i))
		if !bytes.Equal(a.Outputs[i], newValue.Outputs[i]) {
			result = append(result, &internal.Amendment{
				FIP:       lfip,
				Operation: 0, // Modify element
				Data:      newValue.Outputs[i],
			})
		}
	}

	OutputsMax := len(a.Outputs)
	if OutputsMax < len(newValue.Outputs) {
		OutputsMax = len(newValue.Outputs)
	}

	// Add/Remove values
	for i := OutputsMin; i < OutputsMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.Outputs) {
			amendment.Operation = 1 // Add element
			amendment.Data = newValue.Outputs[i]
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	return result, nil
}

// ServiceField Permission / Amendment Field Indices
const (
	ServiceFieldType      = uint32(1)
	ServiceFieldURL       = uint32(2)
	ServiceFieldPublicKey = uint32(3)
)

// ApplyAmendment updates a ServiceField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ServiceField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Service amendment field index path")
	}

	switch fip[0] {
	case ServiceFieldType: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Type : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Type amendment value failed to deserialize : %s", err)
		} else {
			a.Type = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case ServiceFieldURL: // string
		a.URL = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ServiceFieldPublicKey: // []byte
		if len(data) != 33 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 33)
		}
		copy(a.PublicKey, data)
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Service amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Services and returns
// amendment data.
func (a *ServiceField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ServiceField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Type uint32
	fip = append(ofip, ServiceFieldType)
	if a.Type != newValue.Type {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Type)); err != nil {
			return nil, errors.Wrap(err, "Type")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// URL string
	fip = append(ofip, ServiceFieldURL)
	if a.URL != newValue.URL {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.URL),
		})
	}

	// PublicKey []byte
	fip = append(ofip, ServiceFieldPublicKey)
	if !bytes.Equal(a.PublicKey[:], newValue.PublicKey[:]) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.PublicKey[:],
		})
	}

	return result, nil
}

// TargetAddressField Permission / Amendment Field Indices
const (
	TargetAddressFieldAddress  = uint32(1)
	TargetAddressFieldQuantity = uint32(2)
)

// ApplyAmendment updates a TargetAddressField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TargetAddressField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty TargetAddress amendment field index path")
	}

	switch fip[0] {
	case TargetAddressFieldAddress: // []byte
		a.Address = data
		return permissions.SubPermissions(fip, operation, false)

	case TargetAddressFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		} else {
			a.Quantity = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown TargetAddress amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two TargetAddresss and returns
// amendment data.
func (a *TargetAddressField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *TargetAddressField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Address []byte
	fip = append(ofip, TargetAddressFieldAddress)
	if !bytes.Equal(a.Address, newValue.Address) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Address,
		})
	}

	// Quantity uint64
	fip = append(ofip, TargetAddressFieldQuantity)
	if a.Quantity != newValue.Quantity {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Quantity)); err != nil {
			return nil, errors.Wrap(err, "Quantity")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// VotingSystemField Permission / Amendment Field Indices
const (
	VotingSystemFieldName                    = uint32(1)
	VotingSystemFieldVoteType                = uint32(2)
	VotingSystemFieldTallyLogic              = uint32(3)
	VotingSystemFieldThresholdPercentage     = uint32(4)
	VotingSystemFieldVoteMultiplierPermitted = uint32(5)
	VotingSystemFieldHolderProposalFee       = uint32(6)
)

// ApplyAmendment updates a VotingSystemField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *VotingSystemField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty VotingSystem amendment field index path")
	}

	switch fip[0] {
	case VotingSystemFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case VotingSystemFieldVoteType: // string
		a.VoteType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case VotingSystemFieldTallyLogic: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TallyLogic : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("TallyLogic amendment value failed to deserialize : %s", err)
		} else {
			a.TallyLogic = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case VotingSystemFieldThresholdPercentage: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ThresholdPercentage : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ThresholdPercentage amendment value failed to deserialize : %s", err)
		} else {
			a.ThresholdPercentage = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case VotingSystemFieldVoteMultiplierPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for VoteMultiplierPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("VoteMultiplierPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.VoteMultiplierPermitted); err != nil {
			return nil, fmt.Errorf("VoteMultiplierPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case VotingSystemFieldHolderProposalFee: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for HolderProposalFee : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("HolderProposalFee amendment value failed to deserialize : %s", err)
		} else {
			a.HolderProposalFee = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown VotingSystem amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two VotingSystems and returns
// amendment data.
func (a *VotingSystemField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *VotingSystemField) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP:       fip,
			Operation: 0, // Modify element
			Data:      nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	// Name string
	fip = append(ofip, VotingSystemFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	// VoteType string
	fip = append(ofip, VotingSystemFieldVoteType)
	if a.VoteType != newValue.VoteType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.VoteType),
		})
	}

	// TallyLogic uint32
	fip = append(ofip, VotingSystemFieldTallyLogic)
	if a.TallyLogic != newValue.TallyLogic {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.TallyLogic)); err != nil {
			return nil, errors.Wrap(err, "TallyLogic")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ThresholdPercentage uint32
	fip = append(ofip, VotingSystemFieldThresholdPercentage)
	if a.ThresholdPercentage != newValue.ThresholdPercentage {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ThresholdPercentage)); err != nil {
			return nil, errors.Wrap(err, "ThresholdPercentage")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// VoteMultiplierPermitted bool
	fip = append(ofip, VotingSystemFieldVoteMultiplierPermitted)
	if a.VoteMultiplierPermitted != newValue.VoteMultiplierPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.VoteMultiplierPermitted); err != nil {
			return nil, errors.Wrap(err, "VoteMultiplierPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// HolderProposalFee uint64
	fip = append(ofip, VotingSystemFieldHolderProposalFee)
	if a.HolderProposalFee != newValue.HolderProposalFee {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.HolderProposalFee)); err != nil {
			return nil, errors.Wrap(err, "HolderProposalFee")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

func convertAmendments(amendments []*internal.Amendment) ([]*AmendmentField, error) {
	var result []*AmendmentField
	for _, am := range amendments {
		b, err := permissions.FieldIndexPath(am.FIP).Bytes()
		if err != nil {
			return nil, errors.Wrap(err, "fip")
		}

		result = append(result, &AmendmentField{
			FieldIndexPath: b,
			Operation:      am.Operation,
			Data:           am.Data,
		})
	}

	return result, nil
}
