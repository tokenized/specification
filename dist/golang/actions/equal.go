package actions

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

func (l *ContractOffer) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ContractOffer)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if l.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field BodyOfAgreementType - uint
	if l.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(l.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field ContractType - varchar
	if l.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType string mismatched")
	}

	// Field SupportingDocs - Document
	if len(l.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field GoverningLaw - fixedchar
	if l.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - fixedchar
	if l.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	// Field ContractExpiration - uint
	if l.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if l.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field Issuer - Entity
	if !l.Equal(r.Issuer) {
		return false // fmt.Errorf("Issuer : %s", err)
	}

	// Field IssuerLogoURL - varchar
	if l.IssuerLogoURL != r.IssuerLogoURL {
		return false // fmt.Errorf("IssuerLogoURL string mismatched")
	}

	// Field ContractOperatorIncluded - bool
	if l.ContractOperatorIncluded != r.ContractOperatorIncluded {
		return false // fmt.Errorf("ContractOperatorIncluded boolean mismatched")
	}

	// Field ContractOperator - Entity
	if !l.Equal(r.ContractOperator) {
		return false // fmt.Errorf("ContractOperator : %s", err)
	}

	// Field AdminOracle - Oracle
	if !l.Equal(r.AdminOracle) {
		return false // fmt.Errorf("AdminOracle : %s", err)
	}

	// Field AdminOracleSignature - varbin
	if !bytes.Equal(l.AdminOracleSignature, r.AdminOracleSignature) {
		return false // fmt.Errorf("AdminOracleSignature bytes mismatched")
	}

	// Field AdminOracleSigBlockHeight - uint
	if l.AdminOracleSigBlockHeight != r.AdminOracleSigBlockHeight {
		return false // fmt.Errorf("AdminOracleSigBlockHeight integer mismatched")
	}

	// Field ContractFee - uint
	if l.ContractFee != r.ContractFee {
		return false // fmt.Errorf("ContractFee integer mismatched")
	}

	// Field VotingSystems - VotingSystem
	if len(l.VotingSystems) != len(r.VotingSystems) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.VotingSystems {
		if !v.Equal(r.VotingSystems[i]) {
			return false // fmt.Errorf("VotingSystems[%d] : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if !bytes.Equal(l.ContractPermissions, r.ContractPermissions) {
		return false // fmt.Errorf("ContractPermissions bytes mismatched")
	}

	// Field RestrictedQtyAssets - uint
	if l.RestrictedQtyAssets != r.RestrictedQtyAssets {
		return false // fmt.Errorf("RestrictedQtyAssets integer mismatched")
	}

	// Field AdministrationProposal - bool
	if l.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if l.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field Oracles - Oracle
	if len(l.Oracles) != len(r.Oracles) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Oracles {
		if !v.Equal(r.Oracles[i]) {
			return false // fmt.Errorf("Oracles[%d] : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if !bytes.Equal(l.MasterAddress, r.MasterAddress) {
		return false // fmt.Errorf("MasterAddress bytes mismatched")
	}

	return true
}

func (l *ContractFormation) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ContractFormation)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if l.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field BodyOfAgreementType - uint
	if l.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(l.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field ContractType - varchar
	if l.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType string mismatched")
	}

	// Field SupportingDocs - Document
	if len(l.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field GoverningLaw - fixedchar
	if l.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - fixedchar
	if l.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	// Field ContractExpiration - uint
	if l.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if l.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field Issuer - Entity
	if !l.Equal(r.Issuer) {
		return false // fmt.Errorf("Issuer : %s", err)
	}

	// Field IssuerLogoURL - varchar
	if l.IssuerLogoURL != r.IssuerLogoURL {
		return false // fmt.Errorf("IssuerLogoURL string mismatched")
	}

	// Field ContractOperator - Entity
	if !l.Equal(r.ContractOperator) {
		return false // fmt.Errorf("ContractOperator : %s", err)
	}

	// Field AdminOracle - Oracle
	if !l.Equal(r.AdminOracle) {
		return false // fmt.Errorf("AdminOracle : %s", err)
	}

	// Field AdminOracleSignature - varbin
	if !bytes.Equal(l.AdminOracleSignature, r.AdminOracleSignature) {
		return false // fmt.Errorf("AdminOracleSignature bytes mismatched")
	}

	// Field AdminOracleSigBlockHeight - uint
	if l.AdminOracleSigBlockHeight != r.AdminOracleSigBlockHeight {
		return false // fmt.Errorf("AdminOracleSigBlockHeight integer mismatched")
	}

	// Field ContractFee - uint
	if l.ContractFee != r.ContractFee {
		return false // fmt.Errorf("ContractFee integer mismatched")
	}

	// Field VotingSystems - VotingSystem
	if len(l.VotingSystems) != len(r.VotingSystems) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.VotingSystems {
		if !v.Equal(r.VotingSystems[i]) {
			return false // fmt.Errorf("VotingSystems[%d] : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if !bytes.Equal(l.ContractPermissions, r.ContractPermissions) {
		return false // fmt.Errorf("ContractPermissions bytes mismatched")
	}

	// Field RestrictedQtyAssets - uint
	if l.RestrictedQtyAssets != r.RestrictedQtyAssets {
		return false // fmt.Errorf("RestrictedQtyAssets integer mismatched")
	}

	// Field AdministrationProposal - bool
	if l.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if l.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field Oracles - Oracle
	if len(l.Oracles) != len(r.Oracles) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Oracles {
		if !v.Equal(r.Oracles[i]) {
			return false // fmt.Errorf("Oracles[%d] : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if !bytes.Equal(l.MasterAddress, r.MasterAddress) {
		return false // fmt.Errorf("MasterAddress bytes mismatched")
	}

	// Field ContractRevision - uint
	if l.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *ContractAmendment) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ContractAmendment)
	if !ok {
		return false
	}

	// Field ChangeAdministrationAddress - bool
	if l.ChangeAdministrationAddress != r.ChangeAdministrationAddress {
		return false // fmt.Errorf("ChangeAdministrationAddress boolean mismatched")
	}

	// Field ChangeOperatorAddress - bool
	if l.ChangeOperatorAddress != r.ChangeOperatorAddress {
		return false // fmt.Errorf("ChangeOperatorAddress boolean mismatched")
	}

	// Field ContractRevision - uint
	if l.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field Amendments - Amendment
	if len(l.Amendments) != len(r.Amendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Amendments {
		if !v.Equal(r.Amendments[i]) {
			return false // fmt.Errorf("Amendments[%d] : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if !bytes.Equal(l.RefTxID, r.RefTxID) {
		return false // fmt.Errorf("RefTxID bytes mismatched")
	}

	return true
}

func (l *StaticContractFormation) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*StaticContractFormation)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if l.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field ContractCode - bin
	if !bytes.Equal(l.ContractCode, r.ContractCode) {
		return false // fmt.Errorf("ContractCode bytes mismatched")
	}

	// Field BodyOfAgreementType - uint
	if l.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(l.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field ContractType - varchar
	if l.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType string mismatched")
	}

	// Field SupportingDocs - Document
	if len(l.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field ContractRevision - uint
	if l.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field GoverningLaw - fixedchar
	if l.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - fixedchar
	if l.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	// Field EffectiveDate - uint
	if l.EffectiveDate != r.EffectiveDate {
		return false // fmt.Errorf("EffectiveDate integer mismatched")
	}

	// Field ContractExpiration - uint
	if l.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if l.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field PrevRevTxID - bin
	if !bytes.Equal(l.PrevRevTxID, r.PrevRevTxID) {
		return false // fmt.Errorf("PrevRevTxID bytes mismatched")
	}

	// Field Entities - Entity
	if len(l.Entities) != len(r.Entities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Entities {
		if !v.Equal(r.Entities[i]) {
			return false // fmt.Errorf("Entities[%d] : %s", i, err)
		}
	}

	// Field EntityOracle - Oracle
	if !l.Equal(r.EntityOracle) {
		return false // fmt.Errorf("EntityOracle : %s", err)
	}

	// Field EntityOracleSignature - varbin
	if !bytes.Equal(l.EntityOracleSignature, r.EntityOracleSignature) {
		return false // fmt.Errorf("EntityOracleSignature bytes mismatched")
	}

	// Field EntityOracleSigBlockHeight - uint
	if l.EntityOracleSigBlockHeight != r.EntityOracleSigBlockHeight {
		return false // fmt.Errorf("EntityOracleSigBlockHeight integer mismatched")
	}

	return true
}

func (l *ContractAddressChange) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ContractAddressChange)
	if !ok {
		return false
	}

	// Field NewContractAddress - varbin
	if !bytes.Equal(l.NewContractAddress, r.NewContractAddress) {
		return false // fmt.Errorf("NewContractAddress bytes mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *AssetDefinition) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetDefinition)
	if !ok {
		return false
	}

	// Field AssetPermissions - varbin
	if !bytes.Equal(l.AssetPermissions, r.AssetPermissions) {
		return false // fmt.Errorf("AssetPermissions bytes mismatched")
	}

	// Field TransfersPermitted - bool
	if l.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field TradeRestrictions - fixedchar
	if len(l.TradeRestrictions) != len(r.TradeRestrictions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.TradeRestrictions {
		if v != r.TradeRestrictions[i] {
			return false // fmt.Errorf("Element TradeRestrictions string mismatched")
		}
	}

	// Field EnforcementOrdersPermitted - bool
	if l.EnforcementOrdersPermitted != r.EnforcementOrdersPermitted {
		return false // fmt.Errorf("EnforcementOrdersPermitted boolean mismatched")
	}

	// Field VotingRights - bool
	if l.VotingRights != r.VotingRights {
		return false // fmt.Errorf("VotingRights boolean mismatched")
	}

	// Field VoteMultiplier - uint
	if l.VoteMultiplier != r.VoteMultiplier {
		return false // fmt.Errorf("VoteMultiplier integer mismatched")
	}

	// Field AdministrationProposal - bool
	if l.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if l.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field AssetModificationGovernance - uint
	if l.AssetModificationGovernance != r.AssetModificationGovernance {
		return false // fmt.Errorf("AssetModificationGovernance integer mismatched")
	}

	// Field TokenQty - uint
	if l.TokenQty != r.TokenQty {
		return false // fmt.Errorf("TokenQty integer mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetPayload - varbin
	if !bytes.Equal(l.AssetPayload, r.AssetPayload) {
		return false // fmt.Errorf("AssetPayload bytes mismatched")
	}

	return true
}

func (l *AssetCreation) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetCreation)
	if !ok {
		return false
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field AssetIndex - uint
	if l.AssetIndex != r.AssetIndex {
		return false // fmt.Errorf("AssetIndex integer mismatched")
	}

	// Field AssetPermissions - varbin
	if !bytes.Equal(l.AssetPermissions, r.AssetPermissions) {
		return false // fmt.Errorf("AssetPermissions bytes mismatched")
	}

	// Field TransfersPermitted - bool
	if l.TransfersPermitted != r.TransfersPermitted {
		return false // fmt.Errorf("TransfersPermitted boolean mismatched")
	}

	// Field TradeRestrictions - fixedchar
	if len(l.TradeRestrictions) != len(r.TradeRestrictions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.TradeRestrictions {
		if v != r.TradeRestrictions[i] {
			return false // fmt.Errorf("Element TradeRestrictions string mismatched")
		}
	}

	// Field EnforcementOrdersPermitted - bool
	if l.EnforcementOrdersPermitted != r.EnforcementOrdersPermitted {
		return false // fmt.Errorf("EnforcementOrdersPermitted boolean mismatched")
	}

	// Field VotingRights - bool
	if l.VotingRights != r.VotingRights {
		return false // fmt.Errorf("VotingRights boolean mismatched")
	}

	// Field VoteMultiplier - uint
	if l.VoteMultiplier != r.VoteMultiplier {
		return false // fmt.Errorf("VoteMultiplier integer mismatched")
	}

	// Field AdministrationProposal - bool
	if l.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if l.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field AssetModificationGovernance - uint
	if l.AssetModificationGovernance != r.AssetModificationGovernance {
		return false // fmt.Errorf("AssetModificationGovernance integer mismatched")
	}

	// Field TokenQty - uint
	if l.TokenQty != r.TokenQty {
		return false // fmt.Errorf("TokenQty integer mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetPayload - varbin
	if !bytes.Equal(l.AssetPayload, r.AssetPayload) {
		return false // fmt.Errorf("AssetPayload bytes mismatched")
	}

	// Field AssetRevision - uint
	if l.AssetRevision != r.AssetRevision {
		return false // fmt.Errorf("AssetRevision integer mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *AssetModification) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetModification)
	if !ok {
		return false
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field AssetRevision - uint
	if l.AssetRevision != r.AssetRevision {
		return false // fmt.Errorf("AssetRevision integer mismatched")
	}

	// Field Amendments - Amendment
	if len(l.Amendments) != len(r.Amendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Amendments {
		if !v.Equal(r.Amendments[i]) {
			return false // fmt.Errorf("Amendments[%d] : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if !bytes.Equal(l.RefTxID, r.RefTxID) {
		return false // fmt.Errorf("RefTxID bytes mismatched")
	}

	return true
}

func (l *Transfer) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Transfer)
	if !ok {
		return false
	}

	// Field Assets - AssetTransfer
	if len(l.Assets) != len(r.Assets) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Assets {
		if !v.Equal(r.Assets[i]) {
			return false // fmt.Errorf("Assets[%d] : %s", i, err)
		}
	}

	// Field OfferExpiry - uint
	if l.OfferExpiry != r.OfferExpiry {
		return false // fmt.Errorf("OfferExpiry integer mismatched")
	}

	// Field ExchangeFee - uint
	if l.ExchangeFee != r.ExchangeFee {
		return false // fmt.Errorf("ExchangeFee integer mismatched")
	}

	// Field ExchangeFeeAddress - varbin
	if !bytes.Equal(l.ExchangeFeeAddress, r.ExchangeFeeAddress) {
		return false // fmt.Errorf("ExchangeFeeAddress bytes mismatched")
	}

	return true
}

func (l *Settlement) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Settlement)
	if !ok {
		return false
	}

	// Field Assets - AssetSettlement
	if len(l.Assets) != len(r.Assets) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Assets {
		if !v.Equal(r.Assets[i]) {
			return false // fmt.Errorf("Assets[%d] : %s", i, err)
		}
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Proposal) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Proposal)
	if !ok {
		return false
	}

	// Field Type - uint
	if l.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field VoteSystem - uint
	if l.VoteSystem != r.VoteSystem {
		return false // fmt.Errorf("VoteSystem integer mismatched")
	}

	// Field ProposedAmendments - Amendment
	if len(l.ProposedAmendments) != len(r.ProposedAmendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ProposedAmendments {
		if !v.Equal(r.ProposedAmendments[i]) {
			return false // fmt.Errorf("ProposedAmendments[%d] : %s", i, err)
		}
	}

	// Field VoteOptions - varchar
	if l.VoteOptions != r.VoteOptions {
		return false // fmt.Errorf("VoteOptions string mismatched")
	}

	// Field VoteMax - uint
	if l.VoteMax != r.VoteMax {
		return false // fmt.Errorf("VoteMax integer mismatched")
	}

	// Field ProposalDescription - varchar
	if l.ProposalDescription != r.ProposalDescription {
		return false // fmt.Errorf("ProposalDescription string mismatched")
	}

	// Field ProposalDocumentHash - bin
	if !bytes.Equal(l.ProposalDocumentHash, r.ProposalDocumentHash) {
		return false // fmt.Errorf("ProposalDocumentHash bytes mismatched")
	}

	// Field VoteCutOffTimestamp - uint
	if l.VoteCutOffTimestamp != r.VoteCutOffTimestamp {
		return false // fmt.Errorf("VoteCutOffTimestamp integer mismatched")
	}

	return true
}

func (l *Vote) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Vote)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *BallotCast) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*BallotCast)
	if !ok {
		return false
	}

	// Field VoteTxId - bin
	if !bytes.Equal(l.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field Vote - varchar
	if l.Vote != r.Vote {
		return false // fmt.Errorf("Vote string mismatched")
	}

	return true
}

func (l *BallotCounted) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*BallotCounted)
	if !ok {
		return false
	}

	// Field VoteTxId - bin
	if !bytes.Equal(l.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field Vote - varchar
	if l.Vote != r.Vote {
		return false // fmt.Errorf("Vote string mismatched")
	}

	// Field Quantity - uint
	if l.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Result) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Result)
	if !ok {
		return false
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field ProposedAmendments - Amendment
	if len(l.ProposedAmendments) != len(r.ProposedAmendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ProposedAmendments {
		if !v.Equal(r.ProposedAmendments[i]) {
			return false // fmt.Errorf("ProposedAmendments[%d] : %s", i, err)
		}
	}

	// Field VoteTxId - bin
	if !bytes.Equal(l.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field OptionTally - uint
	if len(l.OptionTally) != len(r.OptionTally) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.OptionTally {
		if v != r.OptionTally[i] {
			return false // fmt.Errorf("Element OptionTally integer mismatched")
		}
	}

	// Field Result - varchar
	if l.Result != r.Result {
		return false // fmt.Errorf("Result string mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Order) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Order)
	if !ok {
		return false
	}

	// Field ComplianceAction - fixedchar
	if l.ComplianceAction != r.ComplianceAction {
		return false // fmt.Errorf("ComplianceAction string mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field TargetAddresses - TargetAddress
	if len(l.TargetAddresses) != len(r.TargetAddresses) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.TargetAddresses {
		if !v.Equal(r.TargetAddresses[i]) {
			return false // fmt.Errorf("TargetAddresses[%d] : %s", i, err)
		}
	}

	// Field FreezeTxId - bin
	if !bytes.Equal(l.FreezeTxId, r.FreezeTxId) {
		return false // fmt.Errorf("FreezeTxId bytes mismatched")
	}

	// Field FreezePeriod - uint
	if l.FreezePeriod != r.FreezePeriod {
		return false // fmt.Errorf("FreezePeriod integer mismatched")
	}

	// Field DepositAddress - varbin
	if !bytes.Equal(l.DepositAddress, r.DepositAddress) {
		return false // fmt.Errorf("DepositAddress bytes mismatched")
	}

	// Field AuthorityName - varchar
	if l.AuthorityName != r.AuthorityName {
		return false // fmt.Errorf("AuthorityName string mismatched")
	}

	// Field AuthorityPublicKey - varbin
	if !bytes.Equal(l.AuthorityPublicKey, r.AuthorityPublicKey) {
		return false // fmt.Errorf("AuthorityPublicKey bytes mismatched")
	}

	// Field SignatureAlgorithm - uint
	if l.SignatureAlgorithm != r.SignatureAlgorithm {
		return false // fmt.Errorf("SignatureAlgorithm integer mismatched")
	}

	// Field OrderSignature - varbin
	if !bytes.Equal(l.OrderSignature, r.OrderSignature) {
		return false // fmt.Errorf("OrderSignature bytes mismatched")
	}

	// Field DeprecatedSupportingEvidenceHash - bin
	if !bytes.Equal(l.DeprecatedSupportingEvidenceHash, r.DeprecatedSupportingEvidenceHash) {
		return false // fmt.Errorf("DeprecatedSupportingEvidenceHash bytes mismatched")
	}

	// Field DeprecatedRefTxs - varbin
	if !bytes.Equal(l.DeprecatedRefTxs, r.DeprecatedRefTxs) {
		return false // fmt.Errorf("DeprecatedRefTxs bytes mismatched")
	}

	// Field BitcoinDispersions - QuantityIndex
	if len(l.BitcoinDispersions) != len(r.BitcoinDispersions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.BitcoinDispersions {
		if !v.Equal(r.BitcoinDispersions[i]) {
			return false // fmt.Errorf("BitcoinDispersions[%d] : %s", i, err)
		}
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	// Field SupportingEvidenceFormat - uint
	if l.SupportingEvidenceFormat != r.SupportingEvidenceFormat {
		return false // fmt.Errorf("SupportingEvidenceFormat integer mismatched")
	}

	// Field SupportingEvidence - varbin
	if !bytes.Equal(l.SupportingEvidence, r.SupportingEvidence) {
		return false // fmt.Errorf("SupportingEvidence bytes mismatched")
	}

	// Field ReferenceTransactions - ReferenceTransaction
	if len(l.ReferenceTransactions) != len(r.ReferenceTransactions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ReferenceTransactions {
		if !v.Equal(r.ReferenceTransactions[i]) {
			return false // fmt.Errorf("ReferenceTransactions[%d] : %s", i, err)
		}
	}

	return true
}

func (l *Freeze) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Freeze)
	if !ok {
		return false
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(l.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field FreezePeriod - uint
	if l.FreezePeriod != r.FreezePeriod {
		return false // fmt.Errorf("FreezePeriod integer mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Thaw) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Thaw)
	if !ok {
		return false
	}

	// Field FreezeTxId - bin
	if !bytes.Equal(l.FreezeTxId, r.FreezeTxId) {
		return false // fmt.Errorf("FreezeTxId bytes mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Confiscation) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Confiscation)
	if !ok {
		return false
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(l.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field DepositQty - uint
	if l.DepositQty != r.DepositQty {
		return false // fmt.Errorf("DepositQty integer mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Reconciliation) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Reconciliation)
	if !ok {
		return false
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(l.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Establishment) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Establishment)
	if !ok {
		return false
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Addition) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Addition)
	if !ok {
		return false
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Alteration) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Alteration)
	if !ok {
		return false
	}

	// Field EntryTxID - bin
	if !bytes.Equal(l.EntryTxID, r.EntryTxID) {
		return false // fmt.Errorf("EntryTxID bytes mismatched")
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Removal) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Removal)
	if !ok {
		return false
	}

	// Field EntryTxID - bin
	if !bytes.Equal(l.EntryTxID, r.EntryTxID) {
		return false // fmt.Errorf("EntryTxID bytes mismatched")
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Message) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Message)
	if !ok {
		return false
	}

	// Field SenderIndexes - uint
	if len(l.SenderIndexes) != len(r.SenderIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.SenderIndexes {
		if v != r.SenderIndexes[i] {
			return false // fmt.Errorf("Element SenderIndexes integer mismatched")
		}
	}

	// Field ReceiverIndexes - uint
	if len(l.ReceiverIndexes) != len(r.ReceiverIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ReceiverIndexes {
		if v != r.ReceiverIndexes[i] {
			return false // fmt.Errorf("Element ReceiverIndexes integer mismatched")
		}
	}

	// Field MessageCode - uint
	if l.MessageCode != r.MessageCode {
		return false // fmt.Errorf("MessageCode integer mismatched")
	}

	// Field MessagePayload - varbin
	if !bytes.Equal(l.MessagePayload, r.MessagePayload) {
		return false // fmt.Errorf("MessagePayload bytes mismatched")
	}

	return true
}

func (l *Rejection) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Rejection)
	if !ok {
		return false
	}

	// Field AddressIndexes - uint
	if len(l.AddressIndexes) != len(r.AddressIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.AddressIndexes {
		if v != r.AddressIndexes[i] {
			return false // fmt.Errorf("Element AddressIndexes integer mismatched")
		}
	}

	// Field RejectAddressIndex - uint
	if l.RejectAddressIndex != r.RejectAddressIndex {
		return false // fmt.Errorf("RejectAddressIndex integer mismatched")
	}

	// Field RejectionCode - uint
	if l.RejectionCode != r.RejectionCode {
		return false // fmt.Errorf("RejectionCode integer mismatched")
	}

	// Field Message - varchar
	if l.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *AdministratorField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AdministratorField)
	if !ok {
		return false
	}

	// Field Type - uint
	if l.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	return true
}

func (l *AmendmentField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AmendmentField)
	if !ok {
		return false
	}

	// Field FieldIndexPath - varbin
	if !bytes.Equal(l.FieldIndexPath, r.FieldIndexPath) {
		return false // fmt.Errorf("FieldIndexPath bytes mismatched")
	}

	// Field Operation - uint
	if l.Operation != r.Operation {
		return false // fmt.Errorf("Operation integer mismatched")
	}

	// Field Data - varbin
	if !bytes.Equal(l.Data, r.Data) {
		return false // fmt.Errorf("Data bytes mismatched")
	}

	return true
}

func (l *AssetReceiverField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetReceiverField)
	if !ok {
		return false
	}

	// Field Address - varbin
	if !bytes.Equal(l.Address, r.Address) {
		return false // fmt.Errorf("Address bytes mismatched")
	}

	// Field Quantity - uint
	if l.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	// Field OracleSigAlgorithm - uint
	if l.OracleSigAlgorithm != r.OracleSigAlgorithm {
		return false // fmt.Errorf("OracleSigAlgorithm integer mismatched")
	}

	// Field OracleIndex - uint
	if l.OracleIndex != r.OracleIndex {
		return false // fmt.Errorf("OracleIndex integer mismatched")
	}

	// Field OracleConfirmationSig - varbin
	if !bytes.Equal(l.OracleConfirmationSig, r.OracleConfirmationSig) {
		return false // fmt.Errorf("OracleConfirmationSig bytes mismatched")
	}

	// Field OracleSigBlockHeight - uint
	if l.OracleSigBlockHeight != r.OracleSigBlockHeight {
		return false // fmt.Errorf("OracleSigBlockHeight integer mismatched")
	}

	return true
}

func (l *AssetSettlementField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetSettlementField)
	if !ok {
		return false
	}

	// Field ContractIndex - uint
	if l.ContractIndex != r.ContractIndex {
		return false // fmt.Errorf("ContractIndex integer mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field Settlements - QuantityIndex
	if len(l.Settlements) != len(r.Settlements) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Settlements {
		if !v.Equal(r.Settlements[i]) {
			return false // fmt.Errorf("Settlements[%d] : %s", i, err)
		}
	}

	return true
}

func (l *AssetTransferField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AssetTransferField)
	if !ok {
		return false
	}

	// Field ContractIndex - uint
	if l.ContractIndex != r.ContractIndex {
		return false // fmt.Errorf("ContractIndex integer mismatched")
	}

	// Field AssetType - fixedchar
	if l.AssetType != r.AssetType {
		return false // fmt.Errorf("AssetType string mismatched")
	}

	// Field AssetCode - bin
	if !bytes.Equal(l.AssetCode, r.AssetCode) {
		return false // fmt.Errorf("AssetCode bytes mismatched")
	}

	// Field AssetSenders - QuantityIndex
	if len(l.AssetSenders) != len(r.AssetSenders) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.AssetSenders {
		if !v.Equal(r.AssetSenders[i]) {
			return false // fmt.Errorf("AssetSenders[%d] : %s", i, err)
		}
	}

	// Field AssetReceivers - AssetReceiver
	if len(l.AssetReceivers) != len(r.AssetReceivers) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.AssetReceivers {
		if !v.Equal(r.AssetReceivers[i]) {
			return false // fmt.Errorf("AssetReceivers[%d] : %s", i, err)
		}
	}

	return true
}

func (l *DocumentField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*DocumentField)
	if !ok {
		return false
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - varchar
	if l.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field Contents - varbin
	if !bytes.Equal(l.Contents, r.Contents) {
		return false // fmt.Errorf("Contents bytes mismatched")
	}

	return true
}

func (l *EntityField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*EntityField)
	if !ok {
		return false
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - fixedchar
	if l.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field LEI - fixedchar
	if l.LEI != r.LEI {
		return false // fmt.Errorf("LEI string mismatched")
	}

	// Field UnitNumber - varchar
	if l.UnitNumber != r.UnitNumber {
		return false // fmt.Errorf("UnitNumber string mismatched")
	}

	// Field BuildingNumber - varchar
	if l.BuildingNumber != r.BuildingNumber {
		return false // fmt.Errorf("BuildingNumber string mismatched")
	}

	// Field Street - varchar
	if l.Street != r.Street {
		return false // fmt.Errorf("Street string mismatched")
	}

	// Field SuburbCity - varchar
	if l.SuburbCity != r.SuburbCity {
		return false // fmt.Errorf("SuburbCity string mismatched")
	}

	// Field TerritoryStateProvinceCode - fixedchar
	if l.TerritoryStateProvinceCode != r.TerritoryStateProvinceCode {
		return false // fmt.Errorf("TerritoryStateProvinceCode string mismatched")
	}

	// Field CountryCode - fixedchar
	if l.CountryCode != r.CountryCode {
		return false // fmt.Errorf("CountryCode string mismatched")
	}

	// Field PostalZIPCode - fixedchar
	if l.PostalZIPCode != r.PostalZIPCode {
		return false // fmt.Errorf("PostalZIPCode string mismatched")
	}

	// Field EmailAddress - varchar
	if l.EmailAddress != r.EmailAddress {
		return false // fmt.Errorf("EmailAddress string mismatched")
	}

	// Field PhoneNumber - varchar
	if l.PhoneNumber != r.PhoneNumber {
		return false // fmt.Errorf("PhoneNumber string mismatched")
	}

	// Field Administration - Administrator
	if len(l.Administration) != len(r.Administration) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Administration {
		if !v.Equal(r.Administration[i]) {
			return false // fmt.Errorf("Administration[%d] : %s", i, err)
		}
	}

	// Field Management - Manager
	if len(l.Management) != len(r.Management) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Management {
		if !v.Equal(r.Management[i]) {
			return false // fmt.Errorf("Management[%d] : %s", i, err)
		}
	}

	// Field DomainName - varchar
	if l.DomainName != r.DomainName {
		return false // fmt.Errorf("DomainName string mismatched")
	}

	// Field ParentAddress - varbin
	if !bytes.Equal(l.ParentAddress, r.ParentAddress) {
		return false // fmt.Errorf("ParentAddress bytes mismatched")
	}

	return true
}

func (l *ManagerField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ManagerField)
	if !ok {
		return false
	}

	// Field Type - uint
	if l.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	return true
}

func (l *OracleField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OracleField)
	if !ok {
		return false
	}

	// Field Entity - Entity
	if !l.Equal(r.Entity) {
		return false // fmt.Errorf("Entity : %s", err)
	}

	// Field URL - varchar
	if l.URL != r.URL {
		return false // fmt.Errorf("URL string mismatched")
	}

	// Field PublicKey - varbin
	if !bytes.Equal(l.PublicKey, r.PublicKey) {
		return false // fmt.Errorf("PublicKey bytes mismatched")
	}

	// Field OracleType - uint
	if len(l.OracleType) != len(r.OracleType) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.OracleType {
		if v != r.OracleType[i] {
			return false // fmt.Errorf("Element OracleType integer mismatched")
		}
	}

	return true
}

func (l *QuantityIndexField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*QuantityIndexField)
	if !ok {
		return false
	}

	// Field Index - uint
	if l.Index != r.Index {
		return false // fmt.Errorf("Index integer mismatched")
	}

	// Field Quantity - uint
	if l.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	return true
}

func (l *ReferenceTransactionField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ReferenceTransactionField)
	if !ok {
		return false
	}

	// Field Transaction - varbin
	if !bytes.Equal(l.Transaction, r.Transaction) {
		return false // fmt.Errorf("Transaction bytes mismatched")
	}

	// Field Outputs - varbin
	if len(l.Outputs) != len(r.Outputs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Outputs {
		if !bytes.Equal(v, r.Outputs[i]) {
			return false // fmt.Errorf("Element Outputs bytes mismatched")
		}
	}

	return true
}

func (l *TargetAddressField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*TargetAddressField)
	if !ok {
		return false
	}

	// Field Address - varbin
	if !bytes.Equal(l.Address, r.Address) {
		return false // fmt.Errorf("Address bytes mismatched")
	}

	// Field Quantity - uint
	if l.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	return true
}

func (l *VotingSystemField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*VotingSystemField)
	if !ok {
		return false
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field VoteType - fixedchar
	if l.VoteType != r.VoteType {
		return false // fmt.Errorf("VoteType string mismatched")
	}

	// Field TallyLogic - uint
	if l.TallyLogic != r.TallyLogic {
		return false // fmt.Errorf("TallyLogic integer mismatched")
	}

	// Field ThresholdPercentage - uint
	if l.ThresholdPercentage != r.ThresholdPercentage {
		return false // fmt.Errorf("ThresholdPercentage integer mismatched")
	}

	// Field VoteMultiplierPermitted - bool
	if l.VoteMultiplierPermitted != r.VoteMultiplierPermitted {
		return false // fmt.Errorf("VoteMultiplierPermitted boolean mismatched")
	}

	// Field HolderProposalFee - uint
	if l.HolderProposalFee != r.HolderProposalFee {
		return false // fmt.Errorf("HolderProposalFee integer mismatched")
	}

	return true
}
