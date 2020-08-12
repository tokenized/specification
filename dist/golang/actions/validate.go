package actions

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tokenized/pkg/bitcoin"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295
)

func (a *ContractOffer) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractName - varchar
	if len(a.ContractName) > max1ByteInteger {
		return fmt.Errorf("ContractName over max size : %d > %d", len(a.ContractName), max1ByteInteger)
	}

	// Field BodyOfAgreementType - uint
	foundBodyOfAgreementType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.BodyOfAgreementType == v {
			foundBodyOfAgreementType = true
			break
		}
	}
	if !foundBodyOfAgreementType {
		return fmt.Errorf("BodyOfAgreementType value not within options [0 1 2] : %d", a.BodyOfAgreementType)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max4ByteInteger {
		return fmt.Errorf("BodyOfAgreement over max size : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("SupportingDocs list over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("SupportingDocs[%d]", i))
		}
	}

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("GoverningLaw fixed width field wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Jurisdiction fixed width field wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("ContractURI over max size : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity
	validValueFoundIssuer := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundIssuer = true
			break
		}
	}
	if !validValueFoundIssuer && a.Issuer != nil {
		return fmt.Errorf("Issuer not allowed. ContractType value not within values [0] : %v", a.ContractType)
	}
	if err := a.Issuer.Validate(); err != nil {
		return errors.Wrap(err, "Issuer")
	}

	// Field ContractOperatorIncluded - bool

	// Field ContractFee - uint

	// Field VotingSystems - VotingSystem
	if len(a.VotingSystems) > max1ByteInteger {
		return fmt.Errorf("VotingSystems list over max length : %d > %d", len(a.VotingSystems), max1ByteInteger)
	}
	for i, v := range a.VotingSystems {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("VotingSystems[%d]", i))
		}
	}

	// Field ContractPermissions - varbin
	if len(a.ContractPermissions) > max2ByteInteger {
		return fmt.Errorf("ContractPermissions over max size : %d > %d", len(a.ContractPermissions), max2ByteInteger)
	}

	// Field RestrictedQtyAssets - uint

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field Oracles - Oracle
	if len(a.Oracles) > max1ByteInteger {
		return fmt.Errorf("Oracles list over max length : %d > %d", len(a.Oracles), max1ByteInteger)
	}
	for i, v := range a.Oracles {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Oracles[%d]", i))
		}
	}

	// Field MasterAddress - varbin
	if len(a.MasterAddress) > 0 {
		if err := AddressIsValid(a.MasterAddress); err != nil {
			return errors.Wrap(err, "MasterAddress")
		}
	}
	if len(a.MasterAddress) > max2ByteInteger {
		return fmt.Errorf("MasterAddress over max size : %d > %d", len(a.MasterAddress), max2ByteInteger)
	}

	// Field EntityContract - varbin
	if len(a.EntityContract) > 0 {
		if err := AddressIsValid(a.EntityContract); err != nil {
			return errors.Wrap(err, "EntityContract")
		}
	}
	validValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			validValueFoundEntityContract = true
			break
		}
	}
	if !validValueFoundEntityContract && len(a.EntityContract) != 0 {
		return fmt.Errorf("EntityContract not allowed. ContractType value not within values [1] : %v", a.ContractType)
	}
	requiredValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			requiredValueFoundEntityContract = true
			break
		}
	}
	if requiredValueFoundEntityContract && len(a.EntityContract) == 0 {
		return fmt.Errorf("EntityContract required. ContractType value within values [1] : %v", a.ContractType)
	}
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}

	// Field OperatorEntityContract - varbin
	if len(a.OperatorEntityContract) > 0 {
		if err := AddressIsValid(a.OperatorEntityContract); err != nil {
			return errors.Wrap(err, "OperatorEntityContract")
		}
	}
	validValueFoundOperatorEntityContract := false
	for _, v := range []bool{true} {
		if a.ContractOperatorIncluded == v {
			validValueFoundOperatorEntityContract = true
			break
		}
	}
	if !validValueFoundOperatorEntityContract && len(a.OperatorEntityContract) != 0 {
		return fmt.Errorf("OperatorEntityContract not allowed. ContractOperatorIncluded value not within values [true] : %v", a.ContractOperatorIncluded)
	}
	if len(a.OperatorEntityContract) > max2ByteInteger {
		return fmt.Errorf("OperatorEntityContract over max size : %d > %d", len(a.OperatorEntityContract), max2ByteInteger)
	}

	// Field ContractType - uint
	foundContractType := false
	for _, v := range []uint32{0, 1} {
		if a.ContractType == v {
			foundContractType = true
			break
		}
	}
	if !foundContractType {
		return fmt.Errorf("ContractType value not within options [0 1] : %d", a.ContractType)
	}

	// Field Services - Service
	validValueFoundServices := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundServices = true
			break
		}
	}
	if !validValueFoundServices && len(a.Services) != 0 {
		return fmt.Errorf("Services not allowed. ContractType value not within values [0] : %v", a.ContractType)
	}
	if len(a.Services) > max1ByteInteger {
		return fmt.Errorf("Services list over max length : %d > %d", len(a.Services), max1ByteInteger)
	}
	for i, v := range a.Services {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Services[%d]", i))
		}
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	if len(a.AdminIdentityCertificates) > max1ByteInteger {
		return fmt.Errorf("AdminIdentityCertificates list over max length : %d > %d", len(a.AdminIdentityCertificates), max1ByteInteger)
	}
	for i, v := range a.AdminIdentityCertificates {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("AdminIdentityCertificates[%d]", i))
		}
	}

	return nil
}

func (a *ContractFormation) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractName - varchar
	if len(a.ContractName) > max1ByteInteger {
		return fmt.Errorf("ContractName over max size : %d > %d", len(a.ContractName), max1ByteInteger)
	}

	// Field BodyOfAgreementType - uint
	foundBodyOfAgreementType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.BodyOfAgreementType == v {
			foundBodyOfAgreementType = true
			break
		}
	}
	if !foundBodyOfAgreementType {
		return fmt.Errorf("BodyOfAgreementType value not within options [0 1 2] : %d", a.BodyOfAgreementType)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max4ByteInteger {
		return fmt.Errorf("BodyOfAgreement over max size : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("SupportingDocs list over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("SupportingDocs[%d]", i))
		}
	}

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("GoverningLaw fixed width field wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Jurisdiction fixed width field wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("ContractURI over max size : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity
	validValueFoundIssuer := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundIssuer = true
			break
		}
	}
	if !validValueFoundIssuer && a.Issuer != nil {
		return fmt.Errorf("Issuer not allowed. ContractType value not within values [0] : %v", a.ContractType)
	}
	if err := a.Issuer.Validate(); err != nil {
		return errors.Wrap(err, "Issuer")
	}

	// Field ContractFee - uint

	// Field VotingSystems - VotingSystem
	if len(a.VotingSystems) > max1ByteInteger {
		return fmt.Errorf("VotingSystems list over max length : %d > %d", len(a.VotingSystems), max1ByteInteger)
	}
	for i, v := range a.VotingSystems {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("VotingSystems[%d]", i))
		}
	}

	// Field ContractPermissions - varbin
	if len(a.ContractPermissions) > max2ByteInteger {
		return fmt.Errorf("ContractPermissions over max size : %d > %d", len(a.ContractPermissions), max2ByteInteger)
	}

	// Field RestrictedQtyAssets - uint

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field Oracles - Oracle
	if len(a.Oracles) > max1ByteInteger {
		return fmt.Errorf("Oracles list over max length : %d > %d", len(a.Oracles), max1ByteInteger)
	}
	for i, v := range a.Oracles {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Oracles[%d]", i))
		}
	}

	// Field MasterAddress - varbin
	if len(a.MasterAddress) > 0 {
		if err := AddressIsValid(a.MasterAddress); err != nil {
			return errors.Wrap(err, "MasterAddress")
		}
	}
	if len(a.MasterAddress) > max2ByteInteger {
		return fmt.Errorf("MasterAddress over max size : %d > %d", len(a.MasterAddress), max2ByteInteger)
	}

	// Field ContractRevision - uint

	// Field Timestamp - uint

	// Field EntityContract - varbin
	if len(a.EntityContract) > 0 {
		if err := AddressIsValid(a.EntityContract); err != nil {
			return errors.Wrap(err, "EntityContract")
		}
	}
	validValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			validValueFoundEntityContract = true
			break
		}
	}
	if !validValueFoundEntityContract && len(a.EntityContract) != 0 {
		return fmt.Errorf("EntityContract not allowed. ContractType value not within values [1] : %v", a.ContractType)
	}
	requiredValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			requiredValueFoundEntityContract = true
			break
		}
	}
	if requiredValueFoundEntityContract && len(a.EntityContract) == 0 {
		return fmt.Errorf("EntityContract required. ContractType value within values [1] : %v", a.ContractType)
	}
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}

	// Field OperatorEntityContract - varbin
	if len(a.OperatorEntityContract) > 0 {
		if err := AddressIsValid(a.OperatorEntityContract); err != nil {
			return errors.Wrap(err, "OperatorEntityContract")
		}
	}
	if len(a.OperatorEntityContract) > max2ByteInteger {
		return fmt.Errorf("OperatorEntityContract over max size : %d > %d", len(a.OperatorEntityContract), max2ByteInteger)
	}

	// Field ContractType - uint
	foundContractType := false
	for _, v := range []uint32{0, 1} {
		if a.ContractType == v {
			foundContractType = true
			break
		}
	}
	if !foundContractType {
		return fmt.Errorf("ContractType value not within options [0 1] : %d", a.ContractType)
	}

	// Field Services - Service
	validValueFoundServices := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundServices = true
			break
		}
	}
	if !validValueFoundServices && len(a.Services) != 0 {
		return fmt.Errorf("Services not allowed. ContractType value not within values [0] : %v", a.ContractType)
	}
	if len(a.Services) > max1ByteInteger {
		return fmt.Errorf("Services list over max length : %d > %d", len(a.Services), max1ByteInteger)
	}
	for i, v := range a.Services {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Services[%d]", i))
		}
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	if len(a.AdminIdentityCertificates) > max1ByteInteger {
		return fmt.Errorf("AdminIdentityCertificates list over max length : %d > %d", len(a.AdminIdentityCertificates), max1ByteInteger)
	}
	for i, v := range a.AdminIdentityCertificates {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("AdminIdentityCertificates[%d]", i))
		}
	}

	// Field AdminAddress - varbin
	if len(a.AdminAddress) > 0 {
		if err := AddressIsValid(a.AdminAddress); err != nil {
			return errors.Wrap(err, "AdminAddress")
		}
	}
	if len(a.AdminAddress) > max2ByteInteger {
		return fmt.Errorf("AdminAddress over max size : %d > %d", len(a.AdminAddress), max2ByteInteger)
	}

	// Field OperatorAddress - varbin
	if len(a.OperatorAddress) > 0 {
		if err := AddressIsValid(a.OperatorAddress); err != nil {
			return errors.Wrap(err, "OperatorAddress")
		}
	}
	if len(a.OperatorAddress) > max2ByteInteger {
		return fmt.Errorf("OperatorAddress over max size : %d > %d", len(a.OperatorAddress), max2ByteInteger)
	}

	return nil
}

func (a *ContractAmendment) Validate() error {
	if a == nil {
		return nil
	}

	// Field ChangeAdministrationAddress - bool

	// Field ChangeOperatorAddress - bool

	// Field ContractRevision - uint

	// Field Amendments - Amendment
	if len(a.Amendments) > max1ByteInteger {
		return fmt.Errorf("Amendments list over max length : %d > %d", len(a.Amendments), max1ByteInteger)
	}
	for i, v := range a.Amendments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Amendments[%d]", i))
		}
	}

	// Field RefTxID - bin
	if len(a.RefTxID) != 0 && len(a.RefTxID) != 32 {
		return fmt.Errorf("RefTxID fixed width field wrong size : %d should be %d",
			len(a.RefTxID), 32)
	}

	return nil
}

func (a *StaticContractFormation) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractName - varchar
	if len(a.ContractName) > max1ByteInteger {
		return fmt.Errorf("ContractName over max size : %d > %d", len(a.ContractName), max1ByteInteger)
	}

	// Field ContractCode - bin
	if len(a.ContractCode) != 0 && len(a.ContractCode) != 32 {
		return fmt.Errorf("ContractCode fixed width field wrong size : %d should be %d",
			len(a.ContractCode), 32)
	}

	// Field BodyOfAgreementType - uint
	foundBodyOfAgreementType := false
	for _, v := range []uint32{1, 2} {
		if a.BodyOfAgreementType == v {
			foundBodyOfAgreementType = true
			break
		}
	}
	if !foundBodyOfAgreementType {
		return fmt.Errorf("BodyOfAgreementType value not within options [1 2] : %d", a.BodyOfAgreementType)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max4ByteInteger {
		return fmt.Errorf("BodyOfAgreement over max size : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field ContractType - varchar
	if len(a.ContractType) > max1ByteInteger {
		return fmt.Errorf("ContractType over max size : %d > %d", len(a.ContractType), max1ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("SupportingDocs list over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("SupportingDocs[%d]", i))
		}
	}

	// Field ContractRevision - uint

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("GoverningLaw fixed width field wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Jurisdiction fixed width field wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field EffectiveDate - uint

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("ContractURI over max size : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field PrevRevTxID - bin
	if len(a.PrevRevTxID) != 0 && len(a.PrevRevTxID) != 32 {
		return fmt.Errorf("PrevRevTxID fixed width field wrong size : %d should be %d",
			len(a.PrevRevTxID), 32)
	}

	// Field Entities - Entity
	if len(a.Entities) > max1ByteInteger {
		return fmt.Errorf("Entities list over max length : %d > %d", len(a.Entities), max1ByteInteger)
	}
	for i, v := range a.Entities {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Entities[%d]", i))
		}
	}

	// Field EntityOracle - Oracle
	if err := a.EntityOracle.Validate(); err != nil {
		return errors.Wrap(err, "EntityOracle")
	}

	// Field EntityOracleSignature - varbin
	if len(a.EntityOracleSignature) > max1ByteInteger {
		return fmt.Errorf("EntityOracleSignature over max size : %d > %d", len(a.EntityOracleSignature), max1ByteInteger)
	}

	// Field EntityOracleSigBlockHeight - uint

	return nil
}

func (a *ContractAddressChange) Validate() error {
	if a == nil {
		return nil
	}

	// Field NewContractAddress - varbin
	if len(a.NewContractAddress) > 0 {
		if err := AddressIsValid(a.NewContractAddress); err != nil {
			return errors.Wrap(err, "NewContractAddress")
		}
	}
	if len(a.NewContractAddress) > max2ByteInteger {
		return fmt.Errorf("NewContractAddress over max size : %d > %d", len(a.NewContractAddress), max2ByteInteger)
	}

	// Field Timestamp - uint

	return nil
}

func (a *AssetDefinition) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetPermissions - varbin
	if len(a.AssetPermissions) > max2ByteInteger {
		return fmt.Errorf("AssetPermissions over max size : %d > %d", len(a.AssetPermissions), max2ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("TradeRestrictions list over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if PolitiesData(v) == nil {
			return fmt.Errorf("TradeRestrictions[%d] resource Polities value not defined : %v", i, v)
		}
		if len(v) != 0 && len(v) != 3 {
			return fmt.Errorf("TradeRestrictions[%d] fixed width element wrong size : %d should be %d",
				i, len(v), 3)
		}
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteMultiplier over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
	}

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field AssetModificationGovernance - uint
	foundAssetModificationGovernance := false
	for _, v := range []uint32{0, 1} {
		if a.AssetModificationGovernance == v {
			foundAssetModificationGovernance = true
			break
		}
	}
	if !foundAssetModificationGovernance {
		return fmt.Errorf("AssetModificationGovernance value not within options [0 1] : %d", a.AssetModificationGovernance)
	}

	// Field TokenQty - uint

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max2ByteInteger {
		return fmt.Errorf("AssetPayload over max size : %d > %d", len(a.AssetPayload), max2ByteInteger)
	}

	return nil
}

func (a *AssetCreation) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetIndex - uint

	// Field AssetPermissions - varbin
	if len(a.AssetPermissions) > max2ByteInteger {
		return fmt.Errorf("AssetPermissions over max size : %d > %d", len(a.AssetPermissions), max2ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("TradeRestrictions list over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if PolitiesData(v) == nil {
			return fmt.Errorf("TradeRestrictions[%d] resource Polities value not defined : %v", i, v)
		}
		if len(v) != 0 && len(v) != 3 {
			return fmt.Errorf("TradeRestrictions[%d] fixed width element wrong size : %d should be %d",
				i, len(v), 3)
		}
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteMultiplier over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
	}

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field AssetModificationGovernance - uint
	foundAssetModificationGovernance := false
	for _, v := range []uint32{0, 1} {
		if a.AssetModificationGovernance == v {
			foundAssetModificationGovernance = true
			break
		}
	}
	if !foundAssetModificationGovernance {
		return fmt.Errorf("AssetModificationGovernance value not within options [0 1] : %d", a.AssetModificationGovernance)
	}

	// Field TokenQty - uint

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max2ByteInteger {
		return fmt.Errorf("AssetPayload over max size : %d > %d", len(a.AssetPayload), max2ByteInteger)
	}

	// Field AssetRevision - uint

	// Field Timestamp - uint

	return nil
}

func (a *AssetModification) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetRevision - uint

	// Field Amendments - Amendment
	if len(a.Amendments) > max1ByteInteger {
		return fmt.Errorf("Amendments list over max length : %d > %d", len(a.Amendments), max1ByteInteger)
	}
	for i, v := range a.Amendments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Amendments[%d]", i))
		}
	}

	// Field RefTxID - bin
	if len(a.RefTxID) != 0 && len(a.RefTxID) != 32 {
		return fmt.Errorf("RefTxID fixed width field wrong size : %d should be %d",
			len(a.RefTxID), 32)
	}

	return nil
}

func (a *Transfer) Validate() error {
	if a == nil {
		return nil
	}

	// Field Assets - AssetTransfer
	if len(a.Assets) > max1ByteInteger {
		return fmt.Errorf("Assets list over max length : %d > %d", len(a.Assets), max1ByteInteger)
	}
	for i, v := range a.Assets {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Assets[%d]", i))
		}
	}

	// Field OfferExpiry - uint

	// Field ExchangeFee - uint

	// Field ExchangeFeeAddress - varbin
	if len(a.ExchangeFeeAddress) > 0 {
		if err := AddressIsValid(a.ExchangeFeeAddress); err != nil {
			return errors.Wrap(err, "ExchangeFeeAddress")
		}
	}
	if len(a.ExchangeFeeAddress) > max2ByteInteger {
		return fmt.Errorf("ExchangeFeeAddress over max size : %d > %d", len(a.ExchangeFeeAddress), max2ByteInteger)
	}

	return nil
}

func (a *Settlement) Validate() error {
	if a == nil {
		return nil
	}

	// Field Assets - AssetSettlement
	if len(a.Assets) > max1ByteInteger {
		return fmt.Errorf("Assets list over max length : %d > %d", len(a.Assets), max1ByteInteger)
	}
	for i, v := range a.Assets {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Assets[%d]", i))
		}
	}

	// Field Timestamp - uint

	return nil
}

func (a *Proposal) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	foundType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.Type == v {
			foundType = true
			break
		}
	}
	if !foundType {
		return fmt.Errorf("Type value not within options [0 1 2] : %d", a.Type)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field VoteSystem - uint
	if a.VoteSystem > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteSystem over max value : %d > %d", a.VoteSystem, max1ByteInteger)
	}

	// Field ProposedAmendments - Amendment
	if len(a.ProposedAmendments) > max1ByteInteger {
		return fmt.Errorf("ProposedAmendments list over max length : %d > %d", len(a.ProposedAmendments), max1ByteInteger)
	}
	for i, v := range a.ProposedAmendments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("ProposedAmendments[%d]", i))
		}
	}

	// Field VoteOptions - varchar
	if len(a.VoteOptions) > max1ByteInteger {
		return fmt.Errorf("VoteOptions over max size : %d > %d", len(a.VoteOptions), max1ByteInteger)
	}

	// Field VoteMax - uint
	if a.VoteMax > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteMax over max value : %d > %d", a.VoteMax, max1ByteInteger)
	}

	// Field ProposalDescription - varchar
	if len(a.ProposalDescription) > max4ByteInteger {
		return fmt.Errorf("ProposalDescription over max size : %d > %d", len(a.ProposalDescription), max4ByteInteger)
	}

	// Field ProposalDocumentHash - bin
	if len(a.ProposalDocumentHash) != 0 && len(a.ProposalDocumentHash) != 32 {
		return fmt.Errorf("ProposalDocumentHash fixed width field wrong size : %d should be %d",
			len(a.ProposalDocumentHash), 32)
	}

	// Field VoteCutOffTimestamp - uint

	return nil
}

func (a *Vote) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	return nil
}

func (a *BallotCast) Validate() error {
	if a == nil {
		return nil
	}

	// Field VoteTxId - bin
	if len(a.VoteTxId) != 0 && len(a.VoteTxId) != 32 {
		return fmt.Errorf("VoteTxId fixed width field wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field Vote - varchar
	if len(a.Vote) > max1ByteInteger {
		return fmt.Errorf("Vote over max size : %d > %d", len(a.Vote), max1ByteInteger)
	}

	return nil
}

func (a *BallotCounted) Validate() error {
	if a == nil {
		return nil
	}

	// Field VoteTxId - bin
	if len(a.VoteTxId) != 0 && len(a.VoteTxId) != 32 {
		return fmt.Errorf("VoteTxId fixed width field wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field Vote - varchar
	if len(a.Vote) > max1ByteInteger {
		return fmt.Errorf("Vote over max size : %d > %d", len(a.Vote), max1ByteInteger)
	}

	// Field Quantity - uint

	// Field Timestamp - uint

	return nil
}

func (a *Result) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field ProposedAmendments - Amendment
	if len(a.ProposedAmendments) > max1ByteInteger {
		return fmt.Errorf("ProposedAmendments list over max length : %d > %d", len(a.ProposedAmendments), max1ByteInteger)
	}
	for i, v := range a.ProposedAmendments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("ProposedAmendments[%d]", i))
		}
	}

	// Field VoteTxId - bin
	if len(a.VoteTxId) != 0 && len(a.VoteTxId) != 32 {
		return fmt.Errorf("VoteTxId fixed width field wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field OptionTally - uint
	if len(a.OptionTally) > max1ByteInteger {
		return fmt.Errorf("OptionTally list over max length : %d > %d", len(a.OptionTally), max1ByteInteger)
	}

	// Field Result - varchar
	if len(a.Result) > max1ByteInteger {
		return fmt.Errorf("Result over max size : %d > %d", len(a.Result), max1ByteInteger)
	}

	// Field Timestamp - uint

	return nil
}

func (a *Order) Validate() error {
	if a == nil {
		return nil
	}

	// Field ComplianceAction - fixedchar
	if len(a.ComplianceAction) != 0 && len(a.ComplianceAction) != 1 {
		return fmt.Errorf("ComplianceAction fixed width field wrong size : %d should be %d",
			len(a.ComplianceAction), 1)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field TargetAddresses - TargetAddress
	if len(a.TargetAddresses) > max4ByteInteger {
		return fmt.Errorf("TargetAddresses list over max length : %d > %d", len(a.TargetAddresses), max4ByteInteger)
	}
	for i, v := range a.TargetAddresses {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("TargetAddresses[%d]", i))
		}
	}

	// Field FreezeTxId - bin
	if len(a.FreezeTxId) != 0 && len(a.FreezeTxId) != 32 {
		return fmt.Errorf("FreezeTxId fixed width field wrong size : %d should be %d",
			len(a.FreezeTxId), 32)
	}

	// Field FreezePeriod - uint

	// Field DepositAddress - varbin
	if len(a.DepositAddress) > 0 {
		if err := AddressIsValid(a.DepositAddress); err != nil {
			return errors.Wrap(err, "DepositAddress")
		}
	}
	if len(a.DepositAddress) > max2ByteInteger {
		return fmt.Errorf("DepositAddress over max size : %d > %d", len(a.DepositAddress), max2ByteInteger)
	}

	// Field AuthorityName - varchar
	if len(a.AuthorityName) > max1ByteInteger {
		return fmt.Errorf("AuthorityName over max size : %d > %d", len(a.AuthorityName), max1ByteInteger)
	}

	// Field AuthorityPublicKey - varbin
	if len(a.AuthorityPublicKey) > max1ByteInteger {
		return fmt.Errorf("AuthorityPublicKey over max size : %d > %d", len(a.AuthorityPublicKey), max1ByteInteger)
	}

	// Field SignatureAlgorithm - uint
	foundSignatureAlgorithm := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			foundSignatureAlgorithm = true
			break
		}
	}
	if !foundSignatureAlgorithm {
		return fmt.Errorf("SignatureAlgorithm value not within options [1] : %d", a.SignatureAlgorithm)
	}

	// Field OrderSignature - varbin
	if len(a.OrderSignature) > max1ByteInteger {
		return fmt.Errorf("OrderSignature over max size : %d > %d", len(a.OrderSignature), max1ByteInteger)
	}

	// Field BitcoinDispersions - QuantityIndex
	if len(a.BitcoinDispersions) > max2ByteInteger {
		return fmt.Errorf("BitcoinDispersions list over max length : %d > %d", len(a.BitcoinDispersions), max2ByteInteger)
	}
	for i, v := range a.BitcoinDispersions {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("BitcoinDispersions[%d]", i))
		}
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	// Field SupportingEvidenceFormat - uint
	foundSupportingEvidenceFormat := false
	for _, v := range []uint32{0, 1} {
		if a.SupportingEvidenceFormat == v {
			foundSupportingEvidenceFormat = true
			break
		}
	}
	if !foundSupportingEvidenceFormat {
		return fmt.Errorf("SupportingEvidenceFormat value not within options [0 1] : %d", a.SupportingEvidenceFormat)
	}

	// Field SupportingEvidence - varbin
	if len(a.SupportingEvidence) > max1ByteInteger {
		return fmt.Errorf("SupportingEvidence over max size : %d > %d", len(a.SupportingEvidence), max1ByteInteger)
	}

	// Field ReferenceTransactions - ReferenceTransaction
	if len(a.ReferenceTransactions) > max4ByteInteger {
		return fmt.Errorf("ReferenceTransactions list over max length : %d > %d", len(a.ReferenceTransactions), max4ByteInteger)
	}
	for i, v := range a.ReferenceTransactions {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("ReferenceTransactions[%d]", i))
		}
	}

	return nil
}

func (a *Freeze) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("Quantities list over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Quantities[%d]", i))
		}
	}

	// Field FreezePeriod - uint

	// Field Timestamp - uint

	return nil
}

func (a *Thaw) Validate() error {
	if a == nil {
		return nil
	}

	// Field FreezeTxId - bin
	if len(a.FreezeTxId) != 0 && len(a.FreezeTxId) != 32 {
		return fmt.Errorf("FreezeTxId fixed width field wrong size : %d should be %d",
			len(a.FreezeTxId), 32)
	}

	// Field Timestamp - uint

	return nil
}

func (a *Confiscation) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("Quantities list over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Quantities[%d]", i))
		}
	}

	// Field DepositQty - uint

	// Field Timestamp - uint

	return nil
}

func (a *Reconciliation) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("Quantities list over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Quantities[%d]", i))
		}
	}

	// Field Timestamp - uint

	return nil
}

func (a *Establishment) Validate() error {
	if a == nil {
		return nil
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Addition) Validate() error {
	if a == nil {
		return nil
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Alteration) Validate() error {
	if a == nil {
		return nil
	}

	// Field EntryTxID - bin
	if len(a.EntryTxID) != 0 && len(a.EntryTxID) != 32 {
		return fmt.Errorf("EntryTxID fixed width field wrong size : %d should be %d",
			len(a.EntryTxID), 32)
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Removal) Validate() error {
	if a == nil {
		return nil
	}

	// Field EntryTxID - bin
	if len(a.EntryTxID) != 0 && len(a.EntryTxID) != 32 {
		return fmt.Errorf("EntryTxID fixed width field wrong size : %d should be %d",
			len(a.EntryTxID), 32)
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Message) Validate() error {
	if a == nil {
		return nil
	}

	// Field SenderIndexes - uint
	if len(a.SenderIndexes) > max1ByteInteger {
		return fmt.Errorf("SenderIndexes list over max length : %d > %d", len(a.SenderIndexes), max1ByteInteger)
	}

	// Field ReceiverIndexes - uint
	if len(a.ReceiverIndexes) > max1ByteInteger {
		return fmt.Errorf("ReceiverIndexes list over max length : %d > %d", len(a.ReceiverIndexes), max1ByteInteger)
	}

	// Field MessageCode - uint
	if a.MessageCode > uint32(max2ByteInteger) {
		return fmt.Errorf("MessageCode over max value : %d > %d", a.MessageCode, max2ByteInteger)
	}

	// Field MessagePayload - varbin
	if len(a.MessagePayload) > max4ByteInteger {
		return fmt.Errorf("MessagePayload over max size : %d > %d", len(a.MessagePayload), max4ByteInteger)
	}

	return nil
}

func (a *Rejection) Validate() error {
	if a == nil {
		return nil
	}

	// Field AddressIndexes - uint
	if len(a.AddressIndexes) > max1ByteInteger {
		return fmt.Errorf("AddressIndexes list over max length : %d > %d", len(a.AddressIndexes), max1ByteInteger)
	}

	// Field RejectAddressIndex - uint
	if a.RejectAddressIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("RejectAddressIndex over max value : %d > %d", a.RejectAddressIndex, max2ByteInteger)
	}

	// Field RejectionCode - uint
	if RejectionsData(a.RejectionCode) == nil {
		return fmt.Errorf("RejectionCode resource Rejections value not defined : %v", a.RejectionCode)
	}
	if a.RejectionCode > uint32(max1ByteInteger) {
		return fmt.Errorf("RejectionCode over max value : %d > %d", a.RejectionCode, max1ByteInteger)
	}

	// Field Message - varchar
	if len(a.Message) > max2ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max2ByteInteger)
	}

	// Field Timestamp - uint

	return nil
}

func (a *AdministratorField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	if RolesData(a.Type) == nil {
		return fmt.Errorf("Type resource Roles value not defined : %v", a.Type)
	}
	if a.Type > uint32(max1ByteInteger) {
		return fmt.Errorf("Type over max value : %d > %d", a.Type, max1ByteInteger)
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	return nil
}

func (a *AdminIdentityCertificateField) Validate() error {
	if a == nil {
		return nil
	}

	// Field EntityContract - varbin
	if len(a.EntityContract) > 0 {
		if err := AddressIsValid(a.EntityContract); err != nil {
			return errors.Wrap(err, "EntityContract")
		}
	}
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}

	// Field Signature - varbin
	if len(a.Signature) > 0 {
		if err := SignatureIsValid(a.Signature); err != nil {
			return errors.Wrap(err, "Signature")
		}
	}
	if len(a.Signature) > max1ByteInteger {
		return fmt.Errorf("Signature over max size : %d > %d", len(a.Signature), max1ByteInteger)
	}

	// Field BlockHeight - uint

	// Field Expiration - uint

	return nil
}

func (a *AmendmentField) Validate() error {
	if a == nil {
		return nil
	}

	// Field FieldIndexPath - varbin
	if len(a.FieldIndexPath) > max1ByteInteger {
		return fmt.Errorf("FieldIndexPath over max size : %d > %d", len(a.FieldIndexPath), max1ByteInteger)
	}

	// Field Operation - uint
	foundOperation := false
	for _, v := range []uint32{0, 1, 2} {
		if a.Operation == v {
			foundOperation = true
			break
		}
	}
	if !foundOperation {
		return fmt.Errorf("Operation value not within options [0 1 2] : %d", a.Operation)
	}

	// Field Data - varbin
	if len(a.Data) > max4ByteInteger {
		return fmt.Errorf("Data over max size : %d > %d", len(a.Data), max4ByteInteger)
	}

	return nil
}

func (a *AssetReceiverField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > 0 {
		if err := AddressIsValid(a.Address); err != nil {
			return errors.Wrap(err, "Address")
		}
	}
	if len(a.Address) > max2ByteInteger {
		return fmt.Errorf("Address over max size : %d > %d", len(a.Address), max2ByteInteger)
	}

	// Field Quantity - uint

	// Field OracleSigAlgorithm - uint
	foundOracleSigAlgorithm := false
	for _, v := range []uint32{0, 1} {
		if a.OracleSigAlgorithm == v {
			foundOracleSigAlgorithm = true
			break
		}
	}
	if !foundOracleSigAlgorithm {
		return fmt.Errorf("OracleSigAlgorithm value not within options [0 1] : %d", a.OracleSigAlgorithm)
	}

	// Field OracleIndex - uint
	if a.OracleIndex > uint32(max1ByteInteger) {
		return fmt.Errorf("OracleIndex over max value : %d > %d", a.OracleIndex, max1ByteInteger)
	}

	// Field OracleConfirmationSig - varbin
	if len(a.OracleConfirmationSig) > 0 {
		if err := SignatureIsValid(a.OracleConfirmationSig); err != nil {
			return errors.Wrap(err, "OracleConfirmationSig")
		}
	}
	if len(a.OracleConfirmationSig) > max1ByteInteger {
		return fmt.Errorf("OracleConfirmationSig over max size : %d > %d", len(a.OracleConfirmationSig), max1ByteInteger)
	}

	// Field OracleSigBlockHeight - uint

	// Field OracleSigExpiry - uint

	return nil
}

func (a *AssetSettlementField) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractIndex - uint
	if a.ContractIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("ContractIndex over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Settlements - QuantityIndex
	if len(a.Settlements) > max1ByteInteger {
		return fmt.Errorf("Settlements list over max length : %d > %d", len(a.Settlements), max1ByteInteger)
	}
	for i, v := range a.Settlements {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Settlements[%d]", i))
		}
	}

	return nil
}

func (a *AssetTransferField) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractIndex - uint
	if a.ContractIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("ContractIndex over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("AssetType fixed width field wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("AssetCode fixed width field wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetSenders - QuantityIndex
	if len(a.AssetSenders) > max1ByteInteger {
		return fmt.Errorf("AssetSenders list over max length : %d > %d", len(a.AssetSenders), max1ByteInteger)
	}
	for i, v := range a.AssetSenders {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("AssetSenders[%d]", i))
		}
	}

	// Field AssetReceivers - AssetReceiver
	if len(a.AssetReceivers) > max1ByteInteger {
		return fmt.Errorf("AssetReceivers list over max length : %d > %d", len(a.AssetReceivers), max1ByteInteger)
	}
	for i, v := range a.AssetReceivers {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("AssetReceivers[%d]", i))
		}
	}

	return nil
}

func (a *DocumentField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field Type - varchar
	if len(a.Type) > max1ByteInteger {
		return fmt.Errorf("Type over max size : %d > %d", len(a.Type), max1ByteInteger)
	}

	// Field Contents - varbin
	if len(a.Contents) > max4ByteInteger {
		return fmt.Errorf("Contents over max size : %d > %d", len(a.Contents), max4ByteInteger)
	}

	return nil
}

func (a *EntityField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field Type - fixedchar
	if EntitiesData(a.Type) == nil {
		return fmt.Errorf("Type resource Entities value not defined : %v", a.Type)
	}
	if len(a.Type) != 0 && len(a.Type) != 1 {
		return fmt.Errorf("Type fixed width field wrong size : %d should be %d",
			len(a.Type), 1)
	}

	// Field LEI - fixedchar
	if len(a.LEI) != 0 && len(a.LEI) != 20 {
		return fmt.Errorf("LEI fixed width field wrong size : %d should be %d",
			len(a.LEI), 20)
	}

	// Field UnitNumber - varchar
	if len(a.UnitNumber) > max1ByteInteger {
		return fmt.Errorf("UnitNumber over max size : %d > %d", len(a.UnitNumber), max1ByteInteger)
	}

	// Field BuildingNumber - varchar
	if len(a.BuildingNumber) > max1ByteInteger {
		return fmt.Errorf("BuildingNumber over max size : %d > %d", len(a.BuildingNumber), max1ByteInteger)
	}

	// Field Street - varchar
	if len(a.Street) > max1ByteInteger {
		return fmt.Errorf("Street over max size : %d > %d", len(a.Street), max1ByteInteger)
	}

	// Field SuburbCity - varchar
	if len(a.SuburbCity) > max1ByteInteger {
		return fmt.Errorf("SuburbCity over max size : %d > %d", len(a.SuburbCity), max1ByteInteger)
	}

	// Field TerritoryStateProvinceCode - fixedchar
	if len(a.TerritoryStateProvinceCode) != 0 && len(a.TerritoryStateProvinceCode) != 5 {
		return fmt.Errorf("TerritoryStateProvinceCode fixed width field wrong size : %d should be %d",
			len(a.TerritoryStateProvinceCode), 5)
	}

	// Field CountryCode - fixedchar
	if len(a.CountryCode) != 0 && len(a.CountryCode) != 3 {
		return fmt.Errorf("CountryCode fixed width field wrong size : %d should be %d",
			len(a.CountryCode), 3)
	}

	// Field PostalZIPCode - fixedchar
	if len(a.PostalZIPCode) != 0 && len(a.PostalZIPCode) != 12 {
		return fmt.Errorf("PostalZIPCode fixed width field wrong size : %d should be %d",
			len(a.PostalZIPCode), 12)
	}

	// Field EmailAddress - varchar
	if len(a.EmailAddress) > max1ByteInteger {
		return fmt.Errorf("EmailAddress over max size : %d > %d", len(a.EmailAddress), max1ByteInteger)
	}

	// Field PhoneNumber - varchar
	if len(a.PhoneNumber) > max1ByteInteger {
		return fmt.Errorf("PhoneNumber over max size : %d > %d", len(a.PhoneNumber), max1ByteInteger)
	}

	// Field Administration - Administrator
	if len(a.Administration) > max1ByteInteger {
		return fmt.Errorf("Administration list over max length : %d > %d", len(a.Administration), max1ByteInteger)
	}
	for i, v := range a.Administration {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Administration[%d]", i))
		}
	}

	// Field Management - Manager
	if len(a.Management) > max1ByteInteger {
		return fmt.Errorf("Management list over max length : %d > %d", len(a.Management), max1ByteInteger)
	}
	for i, v := range a.Management {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Management[%d]", i))
		}
	}

	// Field DomainName - varchar
	if len(a.DomainName) > max1ByteInteger {
		return fmt.Errorf("DomainName over max size : %d > %d", len(a.DomainName), max1ByteInteger)
	}

	// Field PaymailHandle - varchar
	if len(a.PaymailHandle) > max1ByteInteger {
		return fmt.Errorf("PaymailHandle over max size : %d > %d", len(a.PaymailHandle), max1ByteInteger)
	}

	return nil
}

func (a *ManagerField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	if RolesData(a.Type) == nil {
		return fmt.Errorf("Type resource Roles value not defined : %v", a.Type)
	}
	if a.Type > uint32(max1ByteInteger) {
		return fmt.Errorf("Type over max value : %d > %d", a.Type, max1ByteInteger)
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	return nil
}

func (a *OracleField) Validate() error {
	if a == nil {
		return nil
	}

	// Field OracleTypes - uint
	if len(a.OracleTypes) > max1ByteInteger {
		return fmt.Errorf("OracleTypes list over max length : %d > %d", len(a.OracleTypes), max1ByteInteger)
	}
	for i, v := range a.OracleTypes {
		foundOracleTypes := false
		for _, o := range []uint32{0, 1, 2} {
			if v == o {
				foundOracleTypes = true
				break
			}
		}
		if !foundOracleTypes {
			return fmt.Errorf("OracleTypes[%d] value not within options [0 1 2] : %d", i, a.OracleTypes)
		}
	}

	// Field EntityContract - varbin
	if len(a.EntityContract) > 0 {
		if err := AddressIsValid(a.EntityContract); err != nil {
			return errors.Wrap(err, "EntityContract")
		}
	}
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}

	return nil
}

func (a *QuantityIndexField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Index - uint
	if a.Index > uint32(max2ByteInteger) {
		return fmt.Errorf("Index over max value : %d > %d", a.Index, max2ByteInteger)
	}

	// Field Quantity - uint

	return nil
}

func (a *ReferenceTransactionField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Transaction - varbin
	if len(a.Transaction) > max1ByteInteger {
		return fmt.Errorf("Transaction over max size : %d > %d", len(a.Transaction), max1ByteInteger)
	}

	// Field Outputs - varbin
	if len(a.Outputs) > max4ByteInteger {
		return fmt.Errorf("Outputs list over max length : %d > %d", len(a.Outputs), max4ByteInteger)
	}

	return nil
}

func (a *ServiceField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	foundType := false
	for _, v := range []uint32{0, 1, 2, 3} {
		if a.Type == v {
			foundType = true
			break
		}
	}
	if !foundType {
		return fmt.Errorf("Type value not within options [0 1 2 3] : %d", a.Type)
	}

	// Field URL - varchar
	if len(a.URL) > max1ByteInteger {
		return fmt.Errorf("URL over max size : %d > %d", len(a.URL), max1ByteInteger)
	}

	// Field PublicKey - bin
	if len(a.PublicKey) > 0 {
		if err := PublicKeyIsValid(a.PublicKey); err != nil {
			return errors.Wrap(err, "PublicKey")
		}
	}
	if len(a.PublicKey) != 0 && len(a.PublicKey) != 33 {
		return fmt.Errorf("PublicKey fixed width field wrong size : %d should be %d",
			len(a.PublicKey), 33)
	}

	return nil
}

func (a *TargetAddressField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > 0 {
		if err := AddressIsValid(a.Address); err != nil {
			return errors.Wrap(err, "Address")
		}
	}
	if len(a.Address) > max2ByteInteger {
		return fmt.Errorf("Address over max size : %d > %d", len(a.Address), max2ByteInteger)
	}

	// Field Quantity - uint

	return nil
}

func (a *VotingSystemField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("Name over max size : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field VoteType - fixedchar
	if len(a.VoteType) != 0 && len(a.VoteType) != 1 {
		return fmt.Errorf("VoteType fixed width field wrong size : %d should be %d",
			len(a.VoteType), 1)
	}

	// Field TallyLogic - uint
	if a.TallyLogic > uint32(max1ByteInteger) {
		return fmt.Errorf("TallyLogic over max value : %d > %d", a.TallyLogic, max1ByteInteger)
	}

	// Field ThresholdPercentage - uint
	if a.ThresholdPercentage > uint32(max1ByteInteger) {
		return fmt.Errorf("ThresholdPercentage over max value : %d > %d", a.ThresholdPercentage, max1ByteInteger)
	}

	// Field VoteMultiplierPermitted - bool

	// Field HolderProposalFee - uint

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
