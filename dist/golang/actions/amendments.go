package actions

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	proto "github.com/golang/protobuf/proto"
)

// Contract Permission / Amendment Field Indices
const (
	ContractFieldContractName              = uint32(1)
	ContractFieldBodyOfAgreementType       = uint32(2)
	ContractFieldBodyOfAgreement           = uint32(3)
	ContractFieldContractType              = uint32(4)
	ContractFieldSupportingDocs            = uint32(5)
	ContractFieldGoverningLaw              = uint32(6)
	ContractFieldJurisdiction              = uint32(7)
	ContractFieldContractExpiration        = uint32(8)
	ContractFieldContractURI               = uint32(9)
	ContractFieldIssuer                    = uint32(10)
	ContractFieldIssuerLogoURL             = uint32(11)
	ContractFieldContractOperator          = uint32(12)
	ContractFieldAdminOracle               = uint32(13)
	ContractFieldAdminOracleSignature      = uint32(14)
	ContractFieldAdminOracleSigBlockHeight = uint32(15)
	ContractFieldContractFee               = uint32(16)
	ContractFieldVotingSystems             = uint32(17)
	ContractFieldContractPermissions       = uint32(18)
	ContractFieldRestrictedQtyAssets       = uint32(19)
	ContractFieldAdministrationProposal    = uint32(20)
	ContractFieldHolderProposal            = uint32(21)
	ContractFieldOracles                   = uint32(22)
	ContractFieldMasterAddress             = uint32(23)
	ContractFieldContractRevision          = uint32(24)
	ContractFieldTimestamp                 = uint32(25)
)

// ApplyAmendment updates a ContractFormation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ContractFormation) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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
		if len(data) != 1 {
			return nil, fmt.Errorf("BodyOfAgreementType amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.BodyOfAgreementType); err != nil {
			return nil, fmt.Errorf("BodyOfAgreementType amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case ContractFieldBodyOfAgreement: // []byte
		a.BodyOfAgreement = data
		return fip[:], nil

	case ContractFieldContractType: // string
		a.ContractType = string(data)
		return fip[:], nil

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

	case ContractFieldGoverningLaw: // string
		a.GoverningLaw = string(data)
		return fip[:], nil

	case ContractFieldJurisdiction: // string
		a.Jurisdiction = string(data)
		return fip[:], nil

	case ContractFieldContractExpiration: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractExpiration : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ContractExpiration amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ContractExpiration); err != nil {
			return nil, fmt.Errorf("ContractExpiration amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case ContractFieldContractURI: // string
		a.ContractURI = string(data)
		return fip[:], nil

	case ContractFieldIssuer: // EntityField
		return a.Issuer.ApplyAmendment(fip[1:], operation, data)

	case ContractFieldIssuerLogoURL: // string
		a.IssuerLogoURL = string(data)
		return fip[:], nil

	case ContractFieldContractOperator: // EntityField
		return a.ContractOperator.ApplyAmendment(fip[1:], operation, data)

	case ContractFieldAdminOracle: // OracleField
		return a.AdminOracle.ApplyAmendment(fip[1:], operation, data)

	case ContractFieldAdminOracleSignature: // []byte
		a.AdminOracleSignature = data
		return fip[:], nil

	case ContractFieldAdminOracleSigBlockHeight: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AdminOracleSigBlockHeight : %v", fip)
		}
		if len(data) != 4 {
			return nil, fmt.Errorf("AdminOracleSigBlockHeight amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AdminOracleSigBlockHeight); err != nil {
			return nil, fmt.Errorf("AdminOracleSigBlockHeight amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case ContractFieldContractFee: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractFee : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("ContractFee amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ContractFee); err != nil {
			return nil, fmt.Errorf("ContractFee amendment value failed to deserialize : %s", err)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("RestrictedQtyAssets amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.RestrictedQtyAssets); err != nil {
			return nil, fmt.Errorf("RestrictedQtyAssets amendment value failed to deserialize : %s", err)
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

	case ContractFieldMasterAddress: // bytes
		a.MasterAddress = data
		return fip[:], nil

	case ContractFieldContractRevision: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractRevision : %v", fip)
		}
		if len(data) != 4 {
			return nil, fmt.Errorf("ContractRevision amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ContractRevision); err != nil {
			return nil, fmt.Errorf("ContractRevision amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case ContractFieldTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Timestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Timestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Timestamp); err != nil {
			return nil, fmt.Errorf("Timestamp amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown contract amendment field index : %v", fip)
}

// Asset Permission / Amendment Field Indices
const (
	AssetFieldAssetCode                   = uint32(1)
	AssetFieldAssetIndex                  = uint32(2)
	AssetFieldAssetPermissions            = uint32(3)
	AssetFieldTransfersPermitted          = uint32(4)
	AssetFieldTradeRestrictions           = uint32(5)
	AssetFieldEnforcementOrdersPermitted  = uint32(6)
	AssetFieldVotingRights                = uint32(7)
	AssetFieldVoteMultiplier              = uint32(8)
	AssetFieldAdministrationProposal      = uint32(9)
	AssetFieldHolderProposal              = uint32(10)
	AssetFieldAssetModificationGovernance = uint32(11)
	AssetFieldTokenQty                    = uint32(12)
	AssetFieldAssetType                   = uint32(13)
	AssetFieldAssetPayload                = uint32(14)
	AssetFieldAssetRevision               = uint32(15)
	AssetFieldTimestamp                   = uint32(16)
)

// ApplyAmendment updates a AssetCreation based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetCreation) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case AssetFieldAssetCode: // bytes
		if len(data) != 0 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 0)
		}
		copy(a.AssetCode, data)
		return fip[:], nil

	case AssetFieldAssetIndex: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AssetIndex : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("AssetIndex amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AssetIndex); err != nil {
			return nil, fmt.Errorf("AssetIndex amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetFieldAssetPermissions: // []byte
		a.AssetPermissions = data
		return fip[:], nil

	case AssetFieldTransfersPermitted: // bool
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
		if len(data) != 1 {
			return nil, fmt.Errorf("VoteMultiplier amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.VoteMultiplier); err != nil {
			return nil, fmt.Errorf("VoteMultiplier amendment value failed to deserialize : %s", err)
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
		if len(data) != 1 {
			return nil, fmt.Errorf("AssetModificationGovernance amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AssetModificationGovernance); err != nil {
			return nil, fmt.Errorf("AssetModificationGovernance amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetFieldTokenQty: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TokenQty : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("TokenQty amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TokenQty); err != nil {
			return nil, fmt.Errorf("TokenQty amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetFieldAssetPayload: // []byte
		a.AssetPayload = data
		return fip[:], nil

	case AssetFieldAssetRevision: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for AssetRevision : %v", fip)
		}
		if len(data) != 4 {
			return nil, fmt.Errorf("AssetRevision amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.AssetRevision); err != nil {
			return nil, fmt.Errorf("AssetRevision amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetFieldTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Timestamp : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Timestamp amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Timestamp); err != nil {
			return nil, fmt.Errorf("Timestamp amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

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
func (a *AdministratorField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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
		if len(data) != 1 {
			return nil, fmt.Errorf("Type amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Type); err != nil {
			return nil, fmt.Errorf("Type amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AdministratorFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Administrator amendment field index : %v", fip)
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
func (a *AmendmentField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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
		if len(data) != 1 {
			return nil, fmt.Errorf("Operation amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Operation); err != nil {
			return nil, fmt.Errorf("Operation amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AmendmentFieldData: // []byte
		a.Data = data
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Amendment amendment field index : %v", fip)
}

// AssetReceiverField Permission / Amendment Field Indices
const (
	AssetReceiverFieldAddress               = uint32(1)
	AssetReceiverFieldQuantity              = uint32(2)
	AssetReceiverFieldOracleSigAlgorithm    = uint32(3)
	AssetReceiverFieldOracleIndex           = uint32(4)
	AssetReceiverFieldOracleConfirmationSig = uint32(5)
	AssetReceiverFieldOracleSigBlockHeight  = uint32(6)
)

// ApplyAmendment updates a AssetReceiverField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AssetReceiverField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty AssetReceiver amendment field index path")
	}

	switch fip[0] {
	case AssetReceiverFieldAddress: // bytes
		a.Address = data
		return fip[:], nil

	case AssetReceiverFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Quantity amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Quantity); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleSigAlgorithm: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigAlgorithm : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("OracleSigAlgorithm amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.OracleSigAlgorithm); err != nil {
			return nil, fmt.Errorf("OracleSigAlgorithm amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleIndex : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("OracleIndex amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.OracleIndex); err != nil {
			return nil, fmt.Errorf("OracleIndex amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetReceiverFieldOracleConfirmationSig: // []byte
		a.OracleConfirmationSig = data
		return fip[:], nil

	case AssetReceiverFieldOracleSigBlockHeight: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for OracleSigBlockHeight : %v", fip)
		}
		if len(data) != 4 {
			return nil, fmt.Errorf("OracleSigBlockHeight amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.OracleSigBlockHeight); err != nil {
			return nil, fmt.Errorf("OracleSigBlockHeight amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown AssetReceiver amendment field index : %v", fip)
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
func (a *AssetSettlementField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty AssetSettlement amendment field index path")
	}

	switch fip[0] {
	case AssetSettlementFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		if len(data) != 2 {
			return nil, fmt.Errorf("ContractIndex amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ContractIndex); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetSettlementFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetSettlementFieldAssetCode: // bytes
		if len(data) != 0 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 0)
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
func (a *AssetTransferField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty AssetTransfer amendment field index path")
	}

	switch fip[0] {
	case AssetTransferFieldContractIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ContractIndex : %v", fip)
		}
		if len(data) != 2 {
			return nil, fmt.Errorf("ContractIndex amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ContractIndex); err != nil {
			return nil, fmt.Errorf("ContractIndex amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case AssetTransferFieldAssetType: // string
		a.AssetType = string(data)
		return fip[:], nil

	case AssetTransferFieldAssetCode: // bytes
		if len(data) != 0 {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), 0)
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

// DocumentField Permission / Amendment Field Indices
const (
	DocumentFieldName     = uint32(1)
	DocumentFieldType     = uint32(2)
	DocumentFieldContents = uint32(3)
)

// ApplyAmendment updates a DocumentField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *DocumentField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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

// EntityField Permission / Amendment Field Indices
const (
	EntityFieldName                       = uint32(1)
	EntityFieldType                       = uint32(2)
	EntityFieldLEI                        = uint32(3)
	EntityFieldUnitNumber                 = uint32(4)
	EntityFieldBuildingNumber             = uint32(5)
	EntityFieldStreet                     = uint32(6)
	EntityFieldSuburbCity                 = uint32(7)
	EntityFieldTerritoryStateProvinceCode = uint32(8)
	EntityFieldCountryCode                = uint32(9)
	EntityFieldPostalZIPCode              = uint32(10)
	EntityFieldEmailAddress               = uint32(11)
	EntityFieldPhoneNumber                = uint32(12)
	EntityFieldAdministration             = uint32(13)
	EntityFieldManagement                 = uint32(14)
)

// ApplyAmendment updates a EntityField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *EntityField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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

	}

	return nil, fmt.Errorf("Unknown Entity amendment field index : %v", fip)
}

// ManagerField Permission / Amendment Field Indices
const (
	ManagerFieldType = uint32(1)
	ManagerFieldName = uint32(2)
)

// ApplyAmendment updates a ManagerField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ManagerField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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
		if len(data) != 1 {
			return nil, fmt.Errorf("Type amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Type); err != nil {
			return nil, fmt.Errorf("Type amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case ManagerFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Manager amendment field index : %v", fip)
}

// OracleField Permission / Amendment Field Indices
const (
	OracleFieldName      = uint32(1)
	OracleFieldURL       = uint32(2)
	OracleFieldPublicKey = uint32(3)
)

// ApplyAmendment updates a OracleField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *OracleField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty Oracle amendment field index path")
	}

	switch fip[0] {
	case OracleFieldName: // string
		a.Name = string(data)
		return fip[:], nil

	case OracleFieldURL: // string
		a.URL = string(data)
		return fip[:], nil

	case OracleFieldPublicKey: // []byte
		a.PublicKey = data
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown Oracle amendment field index : %v", fip)
}

// QuantityIndexField Permission / Amendment Field Indices
const (
	QuantityIndexFieldIndex    = uint32(1)
	QuantityIndexFieldQuantity = uint32(2)
)

// ApplyAmendment updates a QuantityIndexField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *QuantityIndexField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty QuantityIndex amendment field index path")
	}

	switch fip[0] {
	case QuantityIndexFieldIndex: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Index : %v", fip)
		}
		if len(data) != 2 {
			return nil, fmt.Errorf("Index amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Index); err != nil {
			return nil, fmt.Errorf("Index amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case QuantityIndexFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Quantity amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Quantity); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown QuantityIndex amendment field index : %v", fip)
}

// TargetAddressField Permission / Amendment Field Indices
const (
	TargetAddressFieldAddress  = uint32(1)
	TargetAddressFieldQuantity = uint32(2)
)

// ApplyAmendment updates a TargetAddressField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TargetAddressField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty TargetAddress amendment field index path")
	}

	switch fip[0] {
	case TargetAddressFieldAddress: // bytes
		a.Address = data
		return fip[:], nil

	case TargetAddressFieldQuantity: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Quantity : %v", fip)
		}
		if len(data) != 8 {
			return nil, fmt.Errorf("Quantity amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.Quantity); err != nil {
			return nil, fmt.Errorf("Quantity amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown TargetAddress amendment field index : %v", fip)
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
func (a *VotingSystemField) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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
		if len(data) != 1 {
			return nil, fmt.Errorf("TallyLogic amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TallyLogic); err != nil {
			return nil, fmt.Errorf("TallyLogic amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	case VotingSystemFieldThresholdPercentage: // uint32
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ThresholdPercentage : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("ThresholdPercentage amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.ThresholdPercentage); err != nil {
			return nil, fmt.Errorf("ThresholdPercentage amendment value failed to deserialize : %s", err)
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
		if len(data) != 8 {
			return nil, fmt.Errorf("HolderProposalFee amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.HolderProposalFee); err != nil {
			return nil, fmt.Errorf("HolderProposalFee amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil

	}

	return nil, fmt.Errorf("Unknown VotingSystem amendment field index : %v", fip)
}
