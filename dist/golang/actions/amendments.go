package actions

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/assets"
	"github.com/tokenized/specification/dist/golang/internal"

	proto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
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
	ContractFieldRestrictedQtyAssets                  = uint32(20)
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
	var fip []uint32

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

	// SupportingDocs []DocumentField
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

	// VotingSystems []VotingSystemField
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

	// RestrictedQtyAssets uint64
	fip = []uint32{ContractFieldRestrictedQtyAssets}
	if a.RestrictedQtyAssets != newValue.RestrictedQtyAssets {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.RestrictedQtyAssets)); err != nil {
			return nil, errors.Wrap(err, "RestrictedQtyAssets")
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

	// Oracles []OracleField
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

	// Services []ServiceField
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

	// AdminIdentityCertificates []AdminIdentityCertificateField
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
func (a *ContractFormation) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty contract amendment field index path")
	}

	switch fip[0] {
	case ContractFieldContractName: // string
		a.ContractName = string(data)
		return fip[:], nil

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
		return fip[:], nil

	case ContractFieldBodyOfAgreement: // []byte
		a.BodyOfAgreement = data
		return fip[:], nil

	case DeprecatedContractFieldContractType: // deprecated

	case ContractFieldSupportingDocs: // []DocumentField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify SupportingDocs : %v",
					fip)
			}
			if int(fip[1]) >= len(a.SupportingDocs) {
				return nil, fmt.Errorf("Amendment element index out of range for modify SupportingDocs : %d", fip[1])
			}
			result, err := a.SupportingDocs[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add SupportingDocs : %v",
					fip)
			}
			newValue := &DocumentField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to SupportingDocs failed to deserialize : %s",
						err)
				}
			}
			a.SupportingDocs = append(a.SupportingDocs, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete SupportingDocs : %v",
					fip)
			}
			if int(fip[1]) >= len(a.SupportingDocs) {
				return nil, fmt.Errorf("Amendment element index out of range for delete SupportingDocs : %d", fip[1])
			}

			// Remove item from list
			a.SupportingDocs = append(a.SupportingDocs[:fip[1]], a.SupportingDocs[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
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
		return fip[:], nil

	case ContractFieldContractURI: // string
		a.ContractURI = string(data)
		return fip[:], nil

	case ContractFieldIssuer: // EntityField
		return a.Issuer.ApplyAmendment(fip[1:], operation, data)

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
		return fip[:], nil

	case ContractFieldVotingSystems: // []VotingSystemField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify VotingSystems : %v",
					fip)
			}
			if int(fip[1]) >= len(a.VotingSystems) {
				return nil, fmt.Errorf("Amendment element index out of range for modify VotingSystems : %d", fip[1])
			}
			result, err := a.VotingSystems[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add VotingSystems : %v",
					fip)
			}
			newValue := &VotingSystemField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to VotingSystems failed to deserialize : %s",
						err)
				}
			}
			a.VotingSystems = append(a.VotingSystems, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete VotingSystems : %v",
					fip)
			}
			if int(fip[1]) >= len(a.VotingSystems) {
				return nil, fmt.Errorf("Amendment element index out of range for delete VotingSystems : %d", fip[1])
			}

			// Remove item from list
			a.VotingSystems = append(a.VotingSystems[:fip[1]], a.VotingSystems[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case ContractFieldContractPermissions: // []byte
		a.ContractPermissions = data
		return fip[:], nil

	case ContractFieldRestrictedQtyAssets: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for RestrictedQtyAssets : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("RestrictedQtyAssets amendment value failed to deserialize : %s", err)
		} else {
			a.RestrictedQtyAssets = uint64(value)
		}
		return fip[:], nil

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
		return fip[:], nil

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
		return fip[:], nil

	case ContractFieldOracles: // []OracleField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Oracles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Oracles) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Oracles : %d", fip[1])
			}
			result, err := a.Oracles[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Oracles : %v",
					fip)
			}
			newValue := &OracleField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Oracles failed to deserialize : %s",
						err)
				}
			}
			a.Oracles = append(a.Oracles, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Oracles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Oracles) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Oracles : %d", fip[1])
			}

			// Remove item from list
			a.Oracles = append(a.Oracles[:fip[1]], a.Oracles[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case NotAmendableContractFieldMasterAddress: // []byte
		return nil, fmt.Errorf("MasterAddress field not amendable")

	case ContractFieldEntityContract: // []byte
		a.EntityContract = data
		return fip[:], nil

	case ContractFieldOperatorEntityContract: // []byte
		a.OperatorEntityContract = data
		return fip[:], nil

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
		return fip[:], nil

	case ContractFieldServices: // []ServiceField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Services : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Services) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Services : %d", fip[1])
			}
			result, err := a.Services[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Services : %v",
					fip)
			}
			newValue := &ServiceField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Services failed to deserialize : %s",
						err)
				}
			}
			a.Services = append(a.Services, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Services : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Services) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Services : %d", fip[1])
			}

			// Remove item from list
			a.Services = append(a.Services[:fip[1]], a.Services[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case ContractFieldAdminIdentityCertificates: // []AdminIdentityCertificateField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify AdminIdentityCertificates : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AdminIdentityCertificates) {
				return nil, fmt.Errorf("Amendment element index out of range for modify AdminIdentityCertificates : %d", fip[1])
			}
			result, err := a.AdminIdentityCertificates[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add AdminIdentityCertificates : %v",
					fip)
			}
			newValue := &AdminIdentityCertificateField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to AdminIdentityCertificates failed to deserialize : %s",
						err)
				}
			}
			a.AdminIdentityCertificates = append(a.AdminIdentityCertificates, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete AdminIdentityCertificates : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AdminIdentityCertificates) {
				return nil, fmt.Errorf("Amendment element index out of range for delete AdminIdentityCertificates : %d", fip[1])
			}

			// Remove item from list
			a.AdminIdentityCertificates = append(a.AdminIdentityCertificates[:fip[1]], a.AdminIdentityCertificates[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case ContractFieldGoverningLaw: // string
		a.GoverningLaw = string(data)
		return fip[:], nil

	case ContractFieldJurisdiction: // string
		a.Jurisdiction = string(data)
		return fip[:], nil

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
	var fip []uint32

	// Chapters []ChapterField
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

	// Definitions []DefinedTermField
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
func (a *BodyOfAgreementFormation) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty agreement amendment field index path")
	}

	switch fip[0] {
	case BodyOfAgreementFieldChapters: // []ChapterField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Chapters : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Chapters) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Chapters : %d", fip[1])
			}
			result, err := a.Chapters[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Chapters : %v",
					fip)
			}
			newValue := &ChapterField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Chapters failed to deserialize : %s",
						err)
				}
			}
			a.Chapters = append(a.Chapters, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Chapters : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Chapters) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Chapters : %d", fip[1])
			}

			// Remove item from list
			a.Chapters = append(a.Chapters[:fip[1]], a.Chapters[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case BodyOfAgreementFieldDefinitions: // []DefinedTermField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Definitions : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Definitions) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Definitions : %d", fip[1])
			}
			result, err := a.Definitions[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Definitions : %v",
					fip)
			}
			newValue := &DefinedTermField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Definitions failed to deserialize : %s",
						err)
				}
			}
			a.Definitions = append(a.Definitions, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Definitions : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Definitions) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Definitions : %d", fip[1])
			}

			// Remove item from list
			a.Definitions = append(a.Definitions[:fip[1]], a.Definitions[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown agreement amendment field index : %v", fip)
}

// Asset Permission / Amendment Field Indices
const (
	AssetFieldAssetPermissions                      = uint32(1)
	DeprecatedAssetFieldTransfersPermitted          = uint32(2)
	DeprecatedAssetFieldTradeRestrictionsDeprecated = uint32(3)
	AssetFieldEnforcementOrdersPermitted            = uint32(4)
	AssetFieldVotingRights                          = uint32(5)
	AssetFieldVoteMultiplier                        = uint32(6)
	AssetFieldAdministrationProposal                = uint32(7)
	AssetFieldHolderProposal                        = uint32(8)
	AssetFieldAssetModificationGovernance           = uint32(9)
	AssetFieldAuthorizedTokenQty                    = uint32(10)
	AssetFieldAssetType                             = uint32(11)
	AssetFieldAssetPayload                          = uint32(12)
	AssetFieldTradeRestrictions                     = uint32(13)
)

// CreateAmendments determines the differences between two AssetDefinitions and returns
// amendment data. Use the current value of asset creation, and pass in the new values as an asset
// definition.
func (a *AssetCreation) CreateAmendments(newValue *AssetDefinition) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	var fip []uint32

	// AssetPermissions []byte
	fip = []uint32{AssetFieldAssetPermissions}
	if !bytes.Equal(a.AssetPermissions, newValue.AssetPermissions) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.AssetPermissions,
		})
	}

	// deprecated TransfersPermitted deprecated

	// deprecated TradeRestrictionsDeprecated deprecated

	// EnforcementOrdersPermitted bool
	fip = []uint32{AssetFieldEnforcementOrdersPermitted}
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
	fip = []uint32{AssetFieldVotingRights}
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
	fip = []uint32{AssetFieldVoteMultiplier}
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
	fip = []uint32{AssetFieldAdministrationProposal}
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
	fip = []uint32{AssetFieldHolderProposal}
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

	// AssetModificationGovernance uint32
	fip = []uint32{AssetFieldAssetModificationGovernance}
	if a.AssetModificationGovernance != newValue.AssetModificationGovernance {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.AssetModificationGovernance)); err != nil {
			return nil, errors.Wrap(err, "AssetModificationGovernance")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AuthorizedTokenQty uint64
	fip = []uint32{AssetFieldAuthorizedTokenQty}
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

	// AssetType string
	fip = []uint32{AssetFieldAssetType}
	// AssetType modifications not allowed

	// AssetPayload []byte
	fip = []uint32{AssetFieldAssetPayload}
	if a.AssetType != newValue.AssetType {
		return nil, fmt.Errorf("Asset type modification not allowed : %s -> %s", a.AssetType,
			newValue.AssetType)
	}

	payloadAmendments, err := assets.CreatePayloadAmendments(fip, []byte(a.AssetType),
		a.AssetPayload, newValue.AssetPayload)
	if err != nil {
		return nil, errors.Wrap(err, "AssetPayload")
	}
	result = append(result, payloadAmendments...)

	// TradeRestrictions []string
	fip = []uint32{AssetFieldTradeRestrictions}
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

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

// ApplyAmendment updates a AssetCreation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetCreation) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {

	case AssetFieldAssetPermissions: // []byte
		a.AssetPermissions = data
		return fip[:], nil

	case DeprecatedAssetFieldTransfersPermitted: // deprecated

	case DeprecatedAssetFieldTradeRestrictionsDeprecated: // deprecated

	case AssetFieldEnforcementOrdersPermitted: // bool
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
		return fip[:], nil

	case AssetFieldVotingRights: // bool
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
		return fip[:], nil

	case AssetFieldVoteMultiplier: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for VoteMultiplier : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("VoteMultiplier amendment value failed to deserialize : %s", err)
		} else {
			a.VoteMultiplier = uint32(value)
		}
		return fip[:], nil

	case AssetFieldAdministrationProposal: // bool
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
		return fip[:], nil

	case AssetFieldHolderProposal: // bool
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
		return fip[:], nil

	case AssetFieldAssetModificationGovernance: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AssetModificationGovernance : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("AssetModificationGovernance amendment value failed to deserialize : %s", err)
		} else {
			a.AssetModificationGovernance = uint32(value)
		}
		return fip[:], nil

	case AssetFieldAuthorizedTokenQty: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AuthorizedTokenQty : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("AuthorizedTokenQty amendment value failed to deserialize : %s", err)
		} else {
			a.AuthorizedTokenQty = uint64(value)
		}
		return fip[:], nil

	case AssetFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetFieldAssetPayload: // []byte
		a.AssetPayload = data
		return fip[:], nil

	case AssetFieldTradeRestrictions: // []string
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
			return append(fip[:1], fip[2:]...), nil

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add TradeRestrictions : %v",
					fip)
			}
			newValue := string(data)
			a.TradeRestrictions = append(a.TradeRestrictions, newValue)
			return fip[:], nil

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
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown asset amendment field index : %v", fip)
}

// AdministratorField Permission / Amendment Field Indices
const (
	AdministratorFieldType = uint32(1)
	AdministratorFieldName = uint32(2)
)

// ApplyAmendment updates a AdministratorField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AdministratorField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

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
		return fip[:], nil

	case AdministratorFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Administrator amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Administrators and returns
// amendment data.
func (a *AdministratorField) CreateAmendments(fip []uint32,
	newValue *AdministratorField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *AdminIdentityCertificateField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AdminIdentityCertificate amendment field index path")
	}

	switch fip[0] {
	case AdminIdentityCertificateFieldEntityContract: // []byte
		a.EntityContract = data
		return fip[:], nil

	case AdminIdentityCertificateFieldSignature: // []byte
		a.Signature = data
		return fip[:], nil

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
		return fip[:], nil

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
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown AdminIdentityCertificate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AdminIdentityCertificates and returns
// amendment data.
func (a *AdminIdentityCertificateField) CreateAmendments(fip []uint32,
	newValue *AdminIdentityCertificateField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *AmendmentField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Amendment amendment field index path")
	}

	switch fip[0] {
	case AmendmentFieldFieldIndexPath: // []byte
		a.FieldIndexPath = data
		return fip[:], nil

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
		return fip[:], nil

	case AmendmentFieldData: // []byte
		a.Data = data
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Amendment amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Amendments and returns
// amendment data.
func (a *AmendmentField) CreateAmendments(fip []uint32,
	newValue *AmendmentField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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

// AssetReceiverField Permission / Amendment Field Indices
const (
	AssetReceiverFieldAddress               = uint32(1)
	AssetReceiverFieldQuantity              = uint32(2)
	AssetReceiverFieldOracleSigAlgorithm    = uint32(3)
	AssetReceiverFieldOracleIndex           = uint32(4)
	AssetReceiverFieldOracleConfirmationSig = uint32(5)
	AssetReceiverFieldOracleSigBlockHeight  = uint32(6)
	AssetReceiverFieldOracleSigExpiry       = uint32(7)
)

// ApplyAmendment updates a AssetReceiverField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetReceiverField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AssetReceiver amendment field index path")
	}

	switch fip[0] {
	case AssetReceiverFieldAddress: // []byte
		a.Address = data
		return fip[:], nil

	case AssetReceiverFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		} else {
			a.Quantity = uint64(value)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleSigAlgorithm: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigAlgorithm : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigAlgorithm amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigAlgorithm = uint32(value)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleIndex amendment value failed to deserialize : %s", err)
		} else {
			a.OracleIndex = uint32(value)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleConfirmationSig: // []byte
		a.OracleConfirmationSig = data
		return fip[:], nil

	case AssetReceiverFieldOracleSigBlockHeight: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigBlockHeight : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigBlockHeight amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigBlockHeight = uint32(value)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleSigExpiry: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigExpiry : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("OracleSigExpiry amendment value failed to deserialize : %s", err)
		} else {
			a.OracleSigExpiry = uint64(value)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown AssetReceiver amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AssetReceivers and returns
// amendment data.
func (a *AssetReceiverField) CreateAmendments(fip []uint32,
	newValue *AssetReceiverField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// Address []byte
	fip = append(ofip, AssetReceiverFieldAddress)
	if !bytes.Equal(a.Address, newValue.Address) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.Address,
		})
	}

	// Quantity uint64
	fip = append(ofip, AssetReceiverFieldQuantity)
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
	fip = append(ofip, AssetReceiverFieldOracleSigAlgorithm)
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
	fip = append(ofip, AssetReceiverFieldOracleIndex)
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
	fip = append(ofip, AssetReceiverFieldOracleConfirmationSig)
	if !bytes.Equal(a.OracleConfirmationSig, newValue.OracleConfirmationSig) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.OracleConfirmationSig,
		})
	}

	// OracleSigBlockHeight uint32
	fip = append(ofip, AssetReceiverFieldOracleSigBlockHeight)
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
	fip = append(ofip, AssetReceiverFieldOracleSigExpiry)
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

// AssetSettlementField Permission / Amendment Field Indices
const (
	AssetSettlementFieldContractIndex = uint32(1)
	AssetSettlementFieldAssetType     = uint32(2)
	AssetSettlementFieldAssetCode     = uint32(3)
	AssetSettlementFieldSettlements   = uint32(4)
)

// ApplyAmendment updates a AssetSettlementField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetSettlementField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AssetSettlement amendment field index path")
	}

	switch fip[0] {
	case AssetSettlementFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		} else {
			a.ContractIndex = uint32(value)
		}
		return fip[:], nil

	case AssetSettlementFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetSettlementFieldAssetCode: // []byte
		if len(data) != 20 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 20)
		}
		copy(a.AssetCode, data)
		return fip[:], nil

	case AssetSettlementFieldSettlements: // []QuantityIndexField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Settlements : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Settlements) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Settlements : %d", fip[1])
			}
			result, err := a.Settlements[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Settlements : %v",
					fip)
			}
			newValue := &QuantityIndexField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Settlements failed to deserialize : %s",
						err)
				}
			}
			a.Settlements = append(a.Settlements, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Settlements : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Settlements) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Settlements : %d", fip[1])
			}

			// Remove item from list
			a.Settlements = append(a.Settlements[:fip[1]], a.Settlements[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown AssetSettlement amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AssetSettlements and returns
// amendment data.
func (a *AssetSettlementField) CreateAmendments(fip []uint32,
	newValue *AssetSettlementField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// ContractIndex uint32
	fip = append(ofip, AssetSettlementFieldContractIndex)
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

	// AssetType string
	fip = append(ofip, AssetSettlementFieldAssetType)
	// AssetType modifications not allowed

	// AssetCode []byte
	fip = append(ofip, AssetSettlementFieldAssetCode)
	if !bytes.Equal(a.AssetCode[:], newValue.AssetCode[:]) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.AssetCode[:],
		})
	}

	// Settlements []QuantityIndexField
	fip = append(ofip, AssetSettlementFieldSettlements)
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

// AssetTransferField Permission / Amendment Field Indices
const (
	AssetTransferFieldContractIndex  = uint32(1)
	AssetTransferFieldAssetType      = uint32(2)
	AssetTransferFieldAssetCode      = uint32(3)
	AssetTransferFieldAssetSenders   = uint32(4)
	AssetTransferFieldAssetReceivers = uint32(5)
)

// ApplyAmendment updates a AssetTransferField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetTransferField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AssetTransfer amendment field index path")
	}

	switch fip[0] {
	case AssetTransferFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		} else {
			a.ContractIndex = uint32(value)
		}
		return fip[:], nil

	case AssetTransferFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetTransferFieldAssetCode: // []byte
		if len(data) != 20 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 20)
		}
		copy(a.AssetCode, data)
		return fip[:], nil

	case AssetTransferFieldAssetSenders: // []QuantityIndexField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify AssetSenders : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AssetSenders) {
				return nil, fmt.Errorf("Amendment element index out of range for modify AssetSenders : %d", fip[1])
			}
			result, err := a.AssetSenders[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add AssetSenders : %v",
					fip)
			}
			newValue := &QuantityIndexField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to AssetSenders failed to deserialize : %s",
						err)
				}
			}
			a.AssetSenders = append(a.AssetSenders, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete AssetSenders : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AssetSenders) {
				return nil, fmt.Errorf("Amendment element index out of range for delete AssetSenders : %d", fip[1])
			}

			// Remove item from list
			a.AssetSenders = append(a.AssetSenders[:fip[1]], a.AssetSenders[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case AssetTransferFieldAssetReceivers: // []AssetReceiverField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify AssetReceivers : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AssetReceivers) {
				return nil, fmt.Errorf("Amendment element index out of range for modify AssetReceivers : %d", fip[1])
			}
			result, err := a.AssetReceivers[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add AssetReceivers : %v",
					fip)
			}
			newValue := &AssetReceiverField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to AssetReceivers failed to deserialize : %s",
						err)
				}
			}
			a.AssetReceivers = append(a.AssetReceivers, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete AssetReceivers : %v",
					fip)
			}
			if int(fip[1]) >= len(a.AssetReceivers) {
				return nil, fmt.Errorf("Amendment element index out of range for delete AssetReceivers : %d", fip[1])
			}

			// Remove item from list
			a.AssetReceivers = append(a.AssetReceivers[:fip[1]], a.AssetReceivers[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown AssetTransfer amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AssetTransfers and returns
// amendment data.
func (a *AssetTransferField) CreateAmendments(fip []uint32,
	newValue *AssetTransferField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	// ContractIndex uint32
	fip = append(ofip, AssetTransferFieldContractIndex)
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

	// AssetType string
	fip = append(ofip, AssetTransferFieldAssetType)
	// AssetType modifications not allowed

	// AssetCode []byte
	fip = append(ofip, AssetTransferFieldAssetCode)
	if !bytes.Equal(a.AssetCode[:], newValue.AssetCode[:]) {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: newValue.AssetCode[:],
		})
	}

	// AssetSenders []QuantityIndexField
	fip = append(ofip, AssetTransferFieldAssetSenders)
	AssetSendersMin := len(a.AssetSenders)
	if AssetSendersMin > len(newValue.AssetSenders) {
		AssetSendersMin = len(newValue.AssetSenders)
	}

	// Compare values
	for i := 0; i < AssetSendersMin; i++ {
		lfip := append(fip, uint32(i))
		AssetSendersAmendments, err := a.AssetSenders[i].CreateAmendments(lfip,
			newValue.AssetSenders[i])
		if err != nil {
			return nil, errors.Wrapf(err, "AssetSenders%d", i)
		}
		result = append(result, AssetSendersAmendments...)
	}

	AssetSendersMax := len(a.AssetSenders)
	if AssetSendersMax < len(newValue.AssetSenders) {
		AssetSendersMax = len(newValue.AssetSenders)
	}

	// Add/Remove values
	for i := AssetSendersMin; i < AssetSendersMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.AssetSenders) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.AssetSenders[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize AssetSenders %d", i)
			}
			amendment.Data = b
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	// AssetReceivers []AssetReceiverField
	fip = append(ofip, AssetTransferFieldAssetReceivers)
	AssetReceiversMin := len(a.AssetReceivers)
	if AssetReceiversMin > len(newValue.AssetReceivers) {
		AssetReceiversMin = len(newValue.AssetReceivers)
	}

	// Compare values
	for i := 0; i < AssetReceiversMin; i++ {
		lfip := append(fip, uint32(i))
		AssetReceiversAmendments, err := a.AssetReceivers[i].CreateAmendments(lfip,
			newValue.AssetReceivers[i])
		if err != nil {
			return nil, errors.Wrapf(err, "AssetReceivers%d", i)
		}
		result = append(result, AssetReceiversAmendments...)
	}

	AssetReceiversMax := len(a.AssetReceivers)
	if AssetReceiversMax < len(newValue.AssetReceivers) {
		AssetReceiversMax = len(newValue.AssetReceivers)
	}

	// Add/Remove values
	for i := AssetReceiversMin; i < AssetReceiversMax; i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.AssetReceivers) {
			amendment.Operation = 1 // Add element
			b, err := proto.Marshal(newValue.AssetReceivers[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize AssetReceivers %d", i)
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
func (a *ChapterField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Chapter amendment field index path")
	}

	switch fip[0] {
	case ChapterFieldTitle: // string
		a.Title = string(data)
		return fip[:], nil

	case ChapterFieldPreamble: // string
		a.Preamble = string(data)
		return fip[:], nil

	case ChapterFieldArticles: // []ClauseField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Articles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Articles) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Articles : %d", fip[1])
			}
			result, err := a.Articles[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Articles : %v",
					fip)
			}
			newValue := &ClauseField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Articles failed to deserialize : %s",
						err)
				}
			}
			a.Articles = append(a.Articles, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Articles : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Articles) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Articles : %d", fip[1])
			}

			// Remove item from list
			a.Articles = append(a.Articles[:fip[1]], a.Articles[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown Chapter amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Chapters and returns
// amendment data.
func (a *ChapterField) CreateAmendments(fip []uint32,
	newValue *ChapterField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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

	// Articles []ClauseField
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
func (a *ClauseField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Clause amendment field index path")
	}

	switch fip[0] {
	case ClauseFieldTitle: // string
		a.Title = string(data)
		return fip[:], nil

	case ClauseFieldBody: // string
		a.Body = string(data)
		return fip[:], nil

	case ClauseFieldChildren: // []ClauseField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Children : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Children) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Children : %d", fip[1])
			}
			result, err := a.Children[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Children : %v",
					fip)
			}
			newValue := &ClauseField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Children failed to deserialize : %s",
						err)
				}
			}
			a.Children = append(a.Children, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Children : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Children) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Children : %d", fip[1])
			}

			// Remove item from list
			a.Children = append(a.Children[:fip[1]], a.Children[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown Clause amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Clauses and returns
// amendment data.
func (a *ClauseField) CreateAmendments(fip []uint32,
	newValue *ClauseField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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

	// Children []ClauseField
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
func (a *DefinedTermField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty DefinedTerm amendment field index path")
	}

	switch fip[0] {
	case DefinedTermFieldTerm: // string
		a.Term = string(data)
		return fip[:], nil

	case DefinedTermFieldDefinition: // string
		a.Definition = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown DefinedTerm amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two DefinedTerms and returns
// amendment data.
func (a *DefinedTermField) CreateAmendments(fip []uint32,
	newValue *DefinedTermField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *DocumentField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Document amendment field index path")
	}

	switch fip[0] {
	case DocumentFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	case DocumentFieldType: // string
		a.Type = string(data)
		return fip[:], nil

	case DocumentFieldContents: // []byte
		a.Contents = data
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Document amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Documents and returns
// amendment data.
func (a *DocumentField) CreateAmendments(fip []uint32,
	newValue *DocumentField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *EntityField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Entity amendment field index path")
	}

	switch fip[0] {
	case EntityFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	case EntityFieldType: // string
		if EntitiesData(a.Type) == nil {
			return nil, fmt.Errorf("Entities resource value not defined : %v", a.Type)
		}
		a.Type = string(data)
		return fip[:], nil

	case EntityFieldLEI: // string
		a.LEI = string(data)
		return fip[:], nil

	case EntityFieldUnitNumber: // string
		a.UnitNumber = string(data)
		return fip[:], nil

	case EntityFieldBuildingNumber: // string
		a.BuildingNumber = string(data)
		return fip[:], nil

	case EntityFieldStreet: // string
		a.Street = string(data)
		return fip[:], nil

	case EntityFieldSuburbCity: // string
		a.SuburbCity = string(data)
		return fip[:], nil

	case EntityFieldTerritoryStateProvinceCode: // string
		a.TerritoryStateProvinceCode = string(data)
		return fip[:], nil

	case EntityFieldCountryCode: // string
		a.CountryCode = string(data)
		return fip[:], nil

	case EntityFieldPostalZIPCode: // string
		a.PostalZIPCode = string(data)
		return fip[:], nil

	case EntityFieldEmailAddress: // string
		a.EmailAddress = string(data)
		return fip[:], nil

	case EntityFieldPhoneNumber: // string
		a.PhoneNumber = string(data)
		return fip[:], nil

	case EntityFieldAdministration: // []AdministratorField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Administration : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Administration) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Administration : %d", fip[1])
			}
			result, err := a.Administration[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Administration : %v",
					fip)
			}
			newValue := &AdministratorField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Administration failed to deserialize : %s",
						err)
				}
			}
			a.Administration = append(a.Administration, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Administration : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Administration) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Administration : %d", fip[1])
			}

			// Remove item from list
			a.Administration = append(a.Administration[:fip[1]], a.Administration[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case EntityFieldManagement: // []ManagerField
		switch operation {
		case 0: // Modify
			if len(fip) < 3 { // includes list index and subfield index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify Management : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Management) {
				return nil, fmt.Errorf("Amendment element index out of range for modify Management : %d", fip[1])
			}
			result, err := a.Management[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Management : %v",
					fip)
			}
			newValue := &ManagerField{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to Management failed to deserialize : %s",
						err)
				}
			}
			a.Management = append(a.Management, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete Management : %v",
					fip)
			}
			if int(fip[1]) >= len(a.Management) {
				return nil, fmt.Errorf("Amendment element index out of range for delete Management : %d", fip[1])
			}

			// Remove item from list
			a.Management = append(a.Management[:fip[1]], a.Management[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}

	case EntityFieldDomainName: // string
		a.DomainName = string(data)
		return fip[:], nil

	case DeprecatedEntityFieldEntityContractAddress: // deprecated

	case EntityFieldPaymailHandle: // string
		a.PaymailHandle = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Entity amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Entitys and returns
// amendment data.
func (a *EntityField) CreateAmendments(fip []uint32,
	newValue *EntityField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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

	// Administration []AdministratorField
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

	// Management []ManagerField
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

// ManagerField Permission / Amendment Field Indices
const (
	ManagerFieldType = uint32(1)
	ManagerFieldName = uint32(2)
)

// ApplyAmendment updates a ManagerField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ManagerField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

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
		return fip[:], nil

	case ManagerFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Manager amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Managers and returns
// amendment data.
func (a *ManagerField) CreateAmendments(fip []uint32,
	newValue *ManagerField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *OracleField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

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
			return append(fip[:1], fip[2:]...), nil

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add OracleTypes : %v",
					fip)
			}
			var newValue uint32
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("OracleTypes amendment value failed to deserialize : %s", err)
			} else {
				newValue = uint32(value)
			}
			a.OracleTypes = append(a.OracleTypes, newValue)
			return fip[:], nil

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
			return append(fip[:1], fip[2:]...), nil
		}

	case OracleFieldEntityContract: // []byte
		a.EntityContract = data
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Oracle amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Oracles and returns
// amendment data.
func (a *OracleField) CreateAmendments(fip []uint32,
	newValue *OracleField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *QuantityIndexField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

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
		return fip[:], nil

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
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown QuantityIndex amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two QuantityIndexs and returns
// amendment data.
func (a *QuantityIndexField) CreateAmendments(fip []uint32,
	newValue *QuantityIndexField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *ReferenceTransactionField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty ReferenceTransaction amendment field index path")
	}

	switch fip[0] {
	case ReferenceTransactionFieldTransaction: // []byte
		a.Transaction = data
		return fip[:], nil

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
			return append(fip[:1], fip[2:]...), nil

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add Outputs : %v",
					fip)
			}
			newValue := data
			a.Outputs = append(a.Outputs, newValue)
			return fip[:], nil

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
			return append(fip[:1], fip[2:]...), nil
		}

	}

	return nil, fmt.Errorf("Unknown ReferenceTransaction amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two ReferenceTransactions and returns
// amendment data.
func (a *ReferenceTransactionField) CreateAmendments(fip []uint32,
	newValue *ReferenceTransactionField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *ServiceField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

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
		return fip[:], nil

	case ServiceFieldURL: // string
		a.URL = string(data)
		return fip[:], nil

	case ServiceFieldPublicKey: // []byte
		if len(data) != 33 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 33)
		}
		copy(a.PublicKey, data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Service amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Services and returns
// amendment data.
func (a *ServiceField) CreateAmendments(fip []uint32,
	newValue *ServiceField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *TargetAddressField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty TargetAddress amendment field index path")
	}

	switch fip[0] {
	case TargetAddressFieldAddress: // []byte
		a.Address = data
		return fip[:], nil

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
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown TargetAddress amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two TargetAddresss and returns
// amendment data.
func (a *TargetAddressField) CreateAmendments(fip []uint32,
	newValue *TargetAddressField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
func (a *VotingSystemField) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty VotingSystem amendment field index path")
	}

	switch fip[0] {
	case VotingSystemFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	case VotingSystemFieldVoteType: // string
		a.VoteType = string(data)
		return fip[:], nil

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
		return fip[:], nil

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
		return fip[:], nil

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
		return fip[:], nil

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
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown VotingSystem amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two VotingSystems and returns
// amendment data.
func (a *VotingSystemField) CreateAmendments(fip []uint32,
	newValue *VotingSystemField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

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
		b, err := FieldIndexPath(am.FIP).Bytes()
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
