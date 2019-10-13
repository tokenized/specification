package actions

import (
	"fmt"
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractName), max1ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field ContractType - varchar
	if len(a.ContractType) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractType), max1ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("SupportingDocs[%d] invalid : %s", i, err)
		}
	}

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("Fixed width field GoverningLaw wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Fixed width field Jurisdiction wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity

	// Field IssuerLogoURL - varchar
	if len(a.IssuerLogoURL) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.IssuerLogoURL), max1ByteInteger)
	}

	// Field ContractOperatorIncluded - bool

	// Field ContractOperator - Entity

	// Field AdminOracle - Oracle

	// Field AdminOracleSignature - varbin
	if len(a.AdminOracleSignature) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AdminOracleSignature), max1ByteInteger)
	}

	// Field AdminOracleSigBlockHeight - uint

	// Field ContractFee - uint

	// Field VotingSystems - VotingSystem
	if len(a.VotingSystems) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.VotingSystems), max1ByteInteger)
	}
	for i, v := range a.VotingSystems {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("VotingSystems[%d] invalid : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if len(a.ContractPermissions) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractPermissions), max2ByteInteger)
	}

	// Field RestrictedQtyAssets - uint

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field Oracles - Oracle
	if len(a.Oracles) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Oracles), max1ByteInteger)
	}
	for i, v := range a.Oracles {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Oracles[%d] invalid : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if len(a.MasterAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MasterAddress), max2ByteInteger)
	}

	return nil
}

func (a *ContractFormation) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractName - varchar
	if len(a.ContractName) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractName), max1ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field ContractType - varchar
	if len(a.ContractType) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractType), max1ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("SupportingDocs[%d] invalid : %s", i, err)
		}
	}

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("Fixed width field GoverningLaw wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Fixed width field Jurisdiction wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity

	// Field IssuerLogoURL - varchar
	if len(a.IssuerLogoURL) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.IssuerLogoURL), max1ByteInteger)
	}

	// Field ContractOperator - Entity

	// Field AdminOracle - Oracle

	// Field AdminOracleSignature - varbin
	if len(a.AdminOracleSignature) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AdminOracleSignature), max1ByteInteger)
	}

	// Field AdminOracleSigBlockHeight - uint

	// Field ContractFee - uint

	// Field VotingSystems - VotingSystem
	if len(a.VotingSystems) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.VotingSystems), max1ByteInteger)
	}
	for i, v := range a.VotingSystems {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("VotingSystems[%d] invalid : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if len(a.ContractPermissions) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractPermissions), max2ByteInteger)
	}

	// Field RestrictedQtyAssets - uint

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field Oracles - Oracle
	if len(a.Oracles) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Oracles), max1ByteInteger)
	}
	for i, v := range a.Oracles {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Oracles[%d] invalid : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if len(a.MasterAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MasterAddress), max2ByteInteger)
	}

	// Field ContractRevision - uint

	// Field Timestamp - uint

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
		return fmt.Errorf("List over max length : %d > %d", len(a.Amendments), max1ByteInteger)
	}
	for i, v := range a.Amendments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Amendments[%d] invalid : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if len(a.RefTxID) != 0 && len(a.RefTxID) != 32 {
		return fmt.Errorf("Fixed width field RefTxID wrong size : %d should be %d",
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractName), max1ByteInteger)
	}

	// Field ContractCode - bin
	if len(a.ContractCode) != 0 && len(a.ContractCode) != 32 {
		return fmt.Errorf("Fixed width field ContractCode wrong size : %d should be %d",
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max4ByteInteger)
	}

	// Field ContractType - varchar
	if len(a.ContractType) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractType), max1ByteInteger)
	}

	// Field SupportingDocs - Document
	if len(a.SupportingDocs) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.SupportingDocs), max1ByteInteger)
	}
	for i, v := range a.SupportingDocs {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("SupportingDocs[%d] invalid : %s", i, err)
		}
	}

	// Field ContractRevision - uint

	// Field GoverningLaw - fixedchar
	if len(a.GoverningLaw) != 0 && len(a.GoverningLaw) != 5 {
		return fmt.Errorf("Fixed width field GoverningLaw wrong size : %d should be %d",
			len(a.GoverningLaw), 5)
	}

	// Field Jurisdiction - fixedchar
	if len(a.Jurisdiction) != 0 && len(a.Jurisdiction) != 5 {
		return fmt.Errorf("Fixed width field Jurisdiction wrong size : %d should be %d",
			len(a.Jurisdiction), 5)
	}

	// Field EffectiveDate - uint

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field PrevRevTxID - bin
	if len(a.PrevRevTxID) != 0 && len(a.PrevRevTxID) != 32 {
		return fmt.Errorf("Fixed width field PrevRevTxID wrong size : %d should be %d",
			len(a.PrevRevTxID), 32)
	}

	// Field Entities - Entity
	if len(a.Entities) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Entities), max1ByteInteger)
	}
	for i, v := range a.Entities {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Entities[%d] invalid : %s", i, err)
		}
	}

	// Field EntityOracle - Oracle

	// Field EntityOracleSignature - varbin
	if len(a.EntityOracleSignature) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.EntityOracleSignature), max1ByteInteger)
	}

	// Field EntityOracleSigBlockHeight - uint

	return nil
}

func (a *ContractAddressChange) Validate() error {
	if a == nil {
		return nil
	}

	// Field NewContractAddress - varbin
	if len(a.NewContractAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.NewContractAddress), max2ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPermissions), max2ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if PolitiesData(v) == nil {
			return fmt.Errorf("Polities resource value not defined : %v", v)
		}
		if len(v) != 0 && len(v) != 3 {
			return fmt.Errorf("Fixed width element TradeRestrictions[%d] wrong size : %d should be %d",
				i, len(v), 3)
		}
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPayload), max2ByteInteger)
	}

	return nil
}

func (a *AssetCreation) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetIndex - uint

	// Field AssetPermissions - varbin
	if len(a.AssetPermissions) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPermissions), max2ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if PolitiesData(v) == nil {
			return fmt.Errorf("Polities resource value not defined : %v", v)
		}
		if len(v) != 0 && len(v) != 3 {
			return fmt.Errorf("Fixed width element TradeRestrictions[%d] wrong size : %d should be %d",
				i, len(v), 3)
		}
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPayload), max2ByteInteger)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetRevision - uint

	// Field Amendments - Amendment
	if len(a.Amendments) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Amendments), max1ByteInteger)
	}
	for i, v := range a.Amendments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Amendments[%d] invalid : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if len(a.RefTxID) != 0 && len(a.RefTxID) != 32 {
		return fmt.Errorf("Fixed width field RefTxID wrong size : %d should be %d",
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
		return fmt.Errorf("List over max length : %d > %d", len(a.Assets), max1ByteInteger)
	}
	for i, v := range a.Assets {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Assets[%d] invalid : %s", i, err)
		}
	}

	// Field OfferExpiry - uint

	// Field ExchangeFee - uint

	// Field ExchangeFeeAddress - varbin
	if len(a.ExchangeFeeAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ExchangeFeeAddress), max2ByteInteger)
	}

	return nil
}

func (a *Settlement) Validate() error {
	if a == nil {
		return nil
	}

	// Field Assets - AssetSettlement
	if len(a.Assets) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Assets), max1ByteInteger)
	}
	for i, v := range a.Assets {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Assets[%d] invalid : %s", i, err)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field VoteSystem - uint
	if a.VoteSystem > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.VoteSystem, max1ByteInteger)
	}

	// Field ProposedAmendments - Amendment
	if len(a.ProposedAmendments) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ProposedAmendments), max1ByteInteger)
	}
	for i, v := range a.ProposedAmendments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("ProposedAmendments[%d] invalid : %s", i, err)
		}
	}

	// Field VoteOptions - varchar
	if len(a.VoteOptions) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.VoteOptions), max1ByteInteger)
	}

	// Field VoteMax - uint
	if a.VoteMax > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.VoteMax, max1ByteInteger)
	}

	// Field ProposalDescription - varchar
	if len(a.ProposalDescription) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ProposalDescription), max4ByteInteger)
	}

	// Field ProposalDocumentHash - bin
	if len(a.ProposalDocumentHash) != 0 && len(a.ProposalDocumentHash) != 32 {
		return fmt.Errorf("Fixed width field ProposalDocumentHash wrong size : %d should be %d",
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
		return fmt.Errorf("Fixed width field VoteTxId wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field Vote - varchar
	if len(a.Vote) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Vote), max1ByteInteger)
	}

	return nil
}

func (a *BallotCounted) Validate() error {
	if a == nil {
		return nil
	}

	// Field VoteTxId - bin
	if len(a.VoteTxId) != 0 && len(a.VoteTxId) != 32 {
		return fmt.Errorf("Fixed width field VoteTxId wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field Vote - varchar
	if len(a.Vote) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Vote), max1ByteInteger)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field ProposedAmendments - Amendment
	if len(a.ProposedAmendments) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ProposedAmendments), max1ByteInteger)
	}
	for i, v := range a.ProposedAmendments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("ProposedAmendments[%d] invalid : %s", i, err)
		}
	}

	// Field VoteTxId - bin
	if len(a.VoteTxId) != 0 && len(a.VoteTxId) != 32 {
		return fmt.Errorf("Fixed width field VoteTxId wrong size : %d should be %d",
			len(a.VoteTxId), 32)
	}

	// Field OptionTally - uint
	if len(a.OptionTally) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.OptionTally), max1ByteInteger)
	}

	// Field Result - varchar
	if len(a.Result) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Result), max1ByteInteger)
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
		return fmt.Errorf("Fixed width field ComplianceAction wrong size : %d should be %d",
			len(a.ComplianceAction), 1)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field TargetAddresses - TargetAddress
	if len(a.TargetAddresses) > max4ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TargetAddresses), max4ByteInteger)
	}
	for i, v := range a.TargetAddresses {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("TargetAddresses[%d] invalid : %s", i, err)
		}
	}

	// Field FreezeTxId - bin
	if len(a.FreezeTxId) != 0 && len(a.FreezeTxId) != 32 {
		return fmt.Errorf("Fixed width field FreezeTxId wrong size : %d should be %d",
			len(a.FreezeTxId), 32)
	}

	// Field FreezePeriod - uint

	// Field DepositAddress - varbin
	if len(a.DepositAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.DepositAddress), max2ByteInteger)
	}

	// Field AuthorityName - varchar
	if len(a.AuthorityName) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AuthorityName), max1ByteInteger)
	}

	// Field AuthorityPublicKey - varbin
	if len(a.AuthorityPublicKey) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AuthorityPublicKey), max1ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OrderSignature), max1ByteInteger)
	}

	// Field SupportingEvidenceHash - bin
	if len(a.SupportingEvidenceHash) != 0 && len(a.SupportingEvidenceHash) != 32 {
		return fmt.Errorf("Fixed width field SupportingEvidenceHash wrong size : %d should be %d",
			len(a.SupportingEvidenceHash), 32)
	}

	// Field RefTxs - varbin
	if len(a.RefTxs) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.RefTxs), max4ByteInteger)
	}

	// Field BitcoinDispersions - QuantityIndex
	if len(a.BitcoinDispersions) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.BitcoinDispersions), max2ByteInteger)
	}
	for i, v := range a.BitcoinDispersions {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("BitcoinDispersions[%d] invalid : %s", i, err)
		}
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Freeze) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Quantities[%d] invalid : %s", i, err)
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
		return fmt.Errorf("Fixed width field FreezeTxId wrong size : %d should be %d",
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Quantities[%d] invalid : %s", i, err)
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
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Quantities - QuantityIndex
	if len(a.Quantities) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max2ByteInteger)
	}
	for i, v := range a.Quantities {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Quantities[%d] invalid : %s", i, err)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Addition) Validate() error {
	if a == nil {
		return nil
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Alteration) Validate() error {
	if a == nil {
		return nil
	}

	// Field EntryTxID - bin
	if len(a.EntryTxID) != 0 && len(a.EntryTxID) != 32 {
		return fmt.Errorf("Fixed width field EntryTxID wrong size : %d should be %d",
			len(a.EntryTxID), 32)
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Removal) Validate() error {
	if a == nil {
		return nil
	}

	// Field EntryTxID - bin
	if len(a.EntryTxID) != 0 && len(a.EntryTxID) != 32 {
		return fmt.Errorf("Fixed width field EntryTxID wrong size : %d should be %d",
			len(a.EntryTxID), 32)
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Message) Validate() error {
	if a == nil {
		return nil
	}

	// Field SenderIndexes - uint
	if len(a.SenderIndexes) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.SenderIndexes), max1ByteInteger)
	}

	// Field ReceiverIndexes - uint
	if len(a.ReceiverIndexes) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ReceiverIndexes), max1ByteInteger)
	}

	// Field MessageCode - uint
	if a.MessageCode > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.MessageCode, max2ByteInteger)
	}

	// Field MessagePayload - varbin
	if len(a.MessagePayload) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MessagePayload), max4ByteInteger)
	}

	return nil
}

func (a *Rejection) Validate() error {
	if a == nil {
		return nil
	}

	// Field AddressIndexes - uint
	if len(a.AddressIndexes) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.AddressIndexes), max1ByteInteger)
	}

	// Field RejectAddressIndex - uint
	if a.RejectAddressIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.RejectAddressIndex, max2ByteInteger)
	}

	// Field RejectionCode - uint
	if RejectionsData(a.RejectionCode) == nil {
		return fmt.Errorf("Rejections resource value not defined : %v", a.RejectionCode)
	}
	if a.RejectionCode > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.RejectionCode, max1ByteInteger)
	}

	// Field Message - varchar
	if len(a.Message) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max2ByteInteger)
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
		return fmt.Errorf("Roles resource value not defined : %v", a.Type)
	}
	if a.Type > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Type, max1ByteInteger)
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	return nil
}

func (a *AmendmentField) Validate() error {
	if a == nil {
		return nil
	}

	// Field FieldIndexPath - varbin
	if len(a.FieldIndexPath) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.FieldIndexPath), max1ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Data), max4ByteInteger)
	}

	return nil
}

func (a *AssetReceiverField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Address), max2ByteInteger)
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
		return fmt.Errorf("uint over max value : %d > %d", a.OracleIndex, max1ByteInteger)
	}

	// Field OracleConfirmationSig - varbin
	if len(a.OracleConfirmationSig) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OracleConfirmationSig), max1ByteInteger)
	}

	// Field OracleSigBlockHeight - uint

	return nil
}

func (a *AssetSettlementField) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractIndex - uint
	if a.ContractIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field Settlements - QuantityIndex
	if len(a.Settlements) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Settlements), max1ByteInteger)
	}
	for i, v := range a.Settlements {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Settlements[%d] invalid : %s", i, err)
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
		return fmt.Errorf("uint over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetCode - bin
	if len(a.AssetCode) != 0 && len(a.AssetCode) != 32 {
		return fmt.Errorf("Fixed width field AssetCode wrong size : %d should be %d",
			len(a.AssetCode), 32)
	}

	// Field AssetSenders - QuantityIndex
	if len(a.AssetSenders) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.AssetSenders), max1ByteInteger)
	}
	for i, v := range a.AssetSenders {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("AssetSenders[%d] invalid : %s", i, err)
		}
	}

	// Field AssetReceivers - AssetReceiver
	if len(a.AssetReceivers) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.AssetReceivers), max1ByteInteger)
	}
	for i, v := range a.AssetReceivers {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("AssetReceivers[%d] invalid : %s", i, err)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field Type - varchar
	if len(a.Type) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Type), max1ByteInteger)
	}

	// Field Contents - varbin
	if len(a.Contents) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Contents), max4ByteInteger)
	}

	return nil
}

func (a *EntityField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field Type - fixedchar
	if EntitiesData(a.Type) == nil {
		return fmt.Errorf("Entities resource value not defined : %v", a.Type)
	}
	if len(a.Type) != 0 && len(a.Type) != 1 {
		return fmt.Errorf("Fixed width field Type wrong size : %d should be %d",
			len(a.Type), 1)
	}

	// Field LEI - fixedchar
	if len(a.LEI) != 0 && len(a.LEI) != 20 {
		return fmt.Errorf("Fixed width field LEI wrong size : %d should be %d",
			len(a.LEI), 20)
	}

	// Field UnitNumber - varchar
	if len(a.UnitNumber) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.UnitNumber), max1ByteInteger)
	}

	// Field BuildingNumber - varchar
	if len(a.BuildingNumber) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BuildingNumber), max1ByteInteger)
	}

	// Field Street - varchar
	if len(a.Street) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Street), max2ByteInteger)
	}

	// Field SuburbCity - varchar
	if len(a.SuburbCity) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.SuburbCity), max1ByteInteger)
	}

	// Field TerritoryStateProvinceCode - fixedchar
	if len(a.TerritoryStateProvinceCode) != 0 && len(a.TerritoryStateProvinceCode) != 5 {
		return fmt.Errorf("Fixed width field TerritoryStateProvinceCode wrong size : %d should be %d",
			len(a.TerritoryStateProvinceCode), 5)
	}

	// Field CountryCode - fixedchar
	if len(a.CountryCode) != 0 && len(a.CountryCode) != 3 {
		return fmt.Errorf("Fixed width field CountryCode wrong size : %d should be %d",
			len(a.CountryCode), 3)
	}

	// Field PostalZIPCode - fixedchar
	if len(a.PostalZIPCode) != 0 && len(a.PostalZIPCode) != 12 {
		return fmt.Errorf("Fixed width field PostalZIPCode wrong size : %d should be %d",
			len(a.PostalZIPCode), 12)
	}

	// Field EmailAddress - varchar
	if len(a.EmailAddress) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.EmailAddress), max1ByteInteger)
	}

	// Field PhoneNumber - varchar
	if len(a.PhoneNumber) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.PhoneNumber), max1ByteInteger)
	}

	// Field Administration - Administrator
	if len(a.Administration) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Administration), max1ByteInteger)
	}
	for i, v := range a.Administration {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Administration[%d] invalid : %s", i, err)
		}
	}

	// Field Management - Manager
	if len(a.Management) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Management), max1ByteInteger)
	}
	for i, v := range a.Management {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Management[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *ManagerField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	if RolesData(a.Type) == nil {
		return fmt.Errorf("Roles resource value not defined : %v", a.Type)
	}
	if a.Type > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Type, max1ByteInteger)
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	return nil
}

func (a *OracleField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Entity - Entity

	// Field URL - varchar
	if len(a.URL) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.URL), max1ByteInteger)
	}

	// Field PublicKey - varbin
	if len(a.PublicKey) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.PublicKey), max1ByteInteger)
	}

	return nil
}

func (a *QuantityIndexField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Index - uint
	if a.Index > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Index, max2ByteInteger)
	}

	// Field Quantity - uint

	return nil
}

func (a *TargetAddressField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Address), max2ByteInteger)
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
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field VoteType - fixedchar
	if len(a.VoteType) != 0 && len(a.VoteType) != 1 {
		return fmt.Errorf("Fixed width field VoteType wrong size : %d should be %d",
			len(a.VoteType), 1)
	}

	// Field TallyLogic - uint
	if a.TallyLogic > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.TallyLogic, max1ByteInteger)
	}

	// Field ThresholdPercentage - uint
	if a.ThresholdPercentage > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.ThresholdPercentage, max1ByteInteger)
	}

	// Field VoteMultiplierPermitted - bool

	// Field HolderProposalFee - uint

	return nil
}
