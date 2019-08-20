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
	if a.BodyOfAgreementType > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.BodyOfAgreementType, max1ByteInteger)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max1ByteInteger)
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

	// Field ContractAuthFlags - varbin
	if len(a.ContractAuthFlags) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractAuthFlags), max1ByteInteger)
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

	// Field MasterPKH - varbin
	if len(a.MasterPKH) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MasterPKH), max1ByteInteger)
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
	if a.BodyOfAgreementType > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.BodyOfAgreementType, max1ByteInteger)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max1ByteInteger)
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

	// Field ContractAuthFlags - varbin
	if len(a.ContractAuthFlags) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ContractAuthFlags), max1ByteInteger)
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

	// Field MasterPKH - varbin
	if len(a.MasterPKH) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MasterPKH), max1ByteInteger)
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
	if a.BodyOfAgreementType > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.BodyOfAgreementType, max1ByteInteger)
	}

	// Field BodyOfAgreement - varbin
	if len(a.BodyOfAgreement) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BodyOfAgreement), max1ByteInteger)
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

	return nil
}

func (a *ContractAddressChange) Validate() error {
	if a == nil {
		return nil
	}

	// Field NewContractAddress - varbin
	if len(a.NewContractAddress) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.NewContractAddress), max1ByteInteger)
	}

	// Field Timestamp - uint

	return nil
}

func (a *AssetDefinition) Validate() error {
	if a == nil {
		return nil
	}

	// Field AssetAuthFlags - varbin
	if len(a.AssetAuthFlags) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetAuthFlags), max1ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TradeRestrictions), max1ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
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
	if a.AssetModificationGovernance > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.AssetModificationGovernance, max1ByteInteger)
	}

	// Field TokenQty - uint

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPayload), max1ByteInteger)
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

	// Field AssetAuthFlags - varbin
	if len(a.AssetAuthFlags) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetAuthFlags), max1ByteInteger)
	}

	// Field TransfersPermitted - bool

	// Field TradeRestrictions - fixedchar
	if len(a.TradeRestrictions) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TradeRestrictions), max1ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
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
	if a.AssetModificationGovernance > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.AssetModificationGovernance, max1ByteInteger)
	}

	// Field TokenQty - uint

	// Field AssetType - fixedchar
	if len(a.AssetType) != 0 && len(a.AssetType) != 3 {
		return fmt.Errorf("Fixed width field AssetType wrong size : %d should be %d",
			len(a.AssetType), 3)
	}

	// Field AssetPayload - varbin
	if len(a.AssetPayload) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AssetPayload), max1ByteInteger)
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
	if len(a.ExchangeFeeAddress) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ExchangeFeeAddress), max1ByteInteger)
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

	// Field Initiator - uint
	if a.Initiator > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Initiator, max1ByteInteger)
	}

	// Field AssetSpecificVote - bool

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

	// Field Specific - bool

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
	if len(a.ProposalDescription) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ProposalDescription), max1ByteInteger)
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

	// Field AssetSpecificVote - bool

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

	// Field Specific - bool

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
	if len(a.TargetAddresses) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.TargetAddresses), max1ByteInteger)
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
	if len(a.DepositAddress) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.DepositAddress), max1ByteInteger)
	}

	// Field AuthorityIncluded - bool

	// Field AuthorityName - varchar
	if len(a.AuthorityName) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AuthorityName), max1ByteInteger)
	}

	// Field AuthorityPublicKey - varbin
	if len(a.AuthorityPublicKey) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AuthorityPublicKey), max1ByteInteger)
	}

	// Field SignatureAlgorithm - uint
	if a.SignatureAlgorithm > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.SignatureAlgorithm, max1ByteInteger)
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
	if len(a.RefTxs) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.RefTxs), max1ByteInteger)
	}

	// Field BitcoinDispersions - QuantityIndex
	if len(a.BitcoinDispersions) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.BitcoinDispersions), max1ByteInteger)
	}
	for i, v := range a.BitcoinDispersions {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("BitcoinDispersions[%d] invalid : %s", i, err)
		}
	}

	// Field Message - varchar
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
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
	if len(a.Quantities) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max1ByteInteger)
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
	if len(a.Quantities) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max1ByteInteger)
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
	if len(a.Quantities) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Quantities), max1ByteInteger)
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
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
	}

	return nil
}

func (a *Addition) Validate() error {
	if a == nil {
		return nil
	}

	// Field Message - varchar
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
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
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
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
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
	}

	return nil
}

func (a *Message) Validate() error {
	if a == nil {
		return nil
	}

	// Field AddressIndexes - uint
	if len(a.AddressIndexes) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.AddressIndexes), max1ByteInteger)
	}

	// Field MessageCode - uint
	if a.MessageCode > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.MessageCode, max2ByteInteger)
	}

	// Field MessagePayload - varbin
	if len(a.MessagePayload) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.MessagePayload), max1ByteInteger)
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
	if a.RejectionCode > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.RejectionCode, max1ByteInteger)
	}

	// Field Message - varchar
	if len(a.Message) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Message), max1ByteInteger)
	}

	// Field Timestamp - uint

	return nil
}

func (a *AdministratorField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
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

	// Field FieldIndex - uint
	if a.FieldIndex > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.FieldIndex, max1ByteInteger)
	}

	// Field Element - uint
	if a.Element > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Element, max2ByteInteger)
	}

	// Field SubfieldIndex - uint
	if a.SubfieldIndex > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.SubfieldIndex, max1ByteInteger)
	}

	// Field SubfieldElement - uint
	if a.SubfieldElement > uint32(max2ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.SubfieldElement, max2ByteInteger)
	}

	// Field Operation - uint
	if a.Operation > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.Operation, max1ByteInteger)
	}

	// Field Data - varbin
	if len(a.Data) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Data), max1ByteInteger)
	}

	return nil
}

func (a *AssetReceiverField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Address), max1ByteInteger)
	}

	// Field Quantity - uint

	// Field OracleSigAlgorithm - uint
	if a.OracleSigAlgorithm > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.OracleSigAlgorithm, max1ByteInteger)
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
	if len(a.Contents) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Contents), max1ByteInteger)
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
	if len(a.Type) != 0 && len(a.Type) != 1 {
		return fmt.Errorf("Fixed width field Type wrong size : %d should be %d",
			len(a.Type), 1)
	}

	// Field LEI - fixedchar
	if len(a.LEI) != 0 && len(a.LEI) != 20 {
		return fmt.Errorf("Fixed width field LEI wrong size : %d should be %d",
			len(a.LEI), 20)
	}

	// Field AddressIncluded - bool

	// Field UnitNumber - varchar
	if len(a.UnitNumber) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.UnitNumber), max1ByteInteger)
	}

	// Field BuildingNumber - varchar
	if len(a.BuildingNumber) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BuildingNumber), max1ByteInteger)
	}

	// Field Street - varchar
	if len(a.Street) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Street), max1ByteInteger)
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

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

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
	if len(a.Address) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Address), max1ByteInteger)
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
