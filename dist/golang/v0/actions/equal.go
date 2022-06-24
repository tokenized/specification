package actions

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

func (l *ContractOffer) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ContractOffer{}
	}
	cr := right
	if cr == nil {
		cr = &ContractOffer{}
	}
	r, ok := cr.(*ContractOffer)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if c.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field BodyOfAgreementType - uint
	if c.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(c.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field SupportingDocs - Document
	if len(c.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field ContractExpiration - uint
	if c.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if c.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field Issuer - Entity
	if !c.Issuer.Equal(r.Issuer) {
		return false // fmt.Errorf("Issuer : %s", err)
	}

	// Field ContractOperatorIncluded - bool
	if c.ContractOperatorIncluded != r.ContractOperatorIncluded {
		return false // fmt.Errorf("ContractOperatorIncluded boolean mismatched")
	}

	// Field ContractFee - uint
	if c.ContractFee != r.ContractFee {
		return false // fmt.Errorf("ContractFee integer mismatched")
	}

	// Field VotingSystems - VotingSystem
	if len(c.VotingSystems) != len(r.VotingSystems) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.VotingSystems {
		if !v.Equal(r.VotingSystems[i]) {
			return false // fmt.Errorf("VotingSystems[%d] : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if !bytes.Equal(c.ContractPermissions, r.ContractPermissions) {
		return false // fmt.Errorf("ContractPermissions bytes mismatched")
	}

	// Field RestrictedQtyInstruments - uint
	if c.RestrictedQtyInstruments != r.RestrictedQtyInstruments {
		return false // fmt.Errorf("RestrictedQtyInstruments integer mismatched")
	}

	// Field AdministrationProposal - bool
	if c.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if c.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field Oracles - Oracle
	if len(c.Oracles) != len(r.Oracles) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Oracles {
		if !v.Equal(r.Oracles[i]) {
			return false // fmt.Errorf("Oracles[%d] : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if !bytes.Equal(c.MasterAddress, r.MasterAddress) {
		return false // fmt.Errorf("MasterAddress bytes mismatched")
	}

	// Field EntityContract - varbin
	if !bytes.Equal(c.EntityContract, r.EntityContract) {
		return false // fmt.Errorf("EntityContract bytes mismatched")
	}

	// Field OperatorEntityContract - varbin
	if !bytes.Equal(c.OperatorEntityContract, r.OperatorEntityContract) {
		return false // fmt.Errorf("OperatorEntityContract bytes mismatched")
	}

	// Field ContractType - uint
	if c.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType integer mismatched")
	}

	// Field Services - Service
	if len(c.Services) != len(r.Services) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Services {
		if !v.Equal(r.Services[i]) {
			return false // fmt.Errorf("Services[%d] : %s", i, err)
		}
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	if len(c.AdminIdentityCertificates) != len(r.AdminIdentityCertificates) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.AdminIdentityCertificates {
		if !v.Equal(r.AdminIdentityCertificates[i]) {
			return false // fmt.Errorf("AdminIdentityCertificates[%d] : %s", i, err)
		}
	}

	// Field GoverningLaw - varchar
	if c.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - varchar
	if c.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	return true
}

func (l *ContractFormation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ContractFormation{}
	}
	cr := right
	if cr == nil {
		cr = &ContractFormation{}
	}
	r, ok := cr.(*ContractFormation)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if c.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field BodyOfAgreementType - uint
	if c.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(c.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field SupportingDocs - Document
	if len(c.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field ContractExpiration - uint
	if c.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if c.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field Issuer - Entity
	if !c.Issuer.Equal(r.Issuer) {
		return false // fmt.Errorf("Issuer : %s", err)
	}

	// Field ContractFee - uint
	if c.ContractFee != r.ContractFee {
		return false // fmt.Errorf("ContractFee integer mismatched")
	}

	// Field VotingSystems - VotingSystem
	if len(c.VotingSystems) != len(r.VotingSystems) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.VotingSystems {
		if !v.Equal(r.VotingSystems[i]) {
			return false // fmt.Errorf("VotingSystems[%d] : %s", i, err)
		}
	}

	// Field ContractPermissions - varbin
	if !bytes.Equal(c.ContractPermissions, r.ContractPermissions) {
		return false // fmt.Errorf("ContractPermissions bytes mismatched")
	}

	// Field RestrictedQtyInstruments - uint
	if c.RestrictedQtyInstruments != r.RestrictedQtyInstruments {
		return false // fmt.Errorf("RestrictedQtyInstruments integer mismatched")
	}

	// Field AdministrationProposal - bool
	if c.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if c.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field Oracles - Oracle
	if len(c.Oracles) != len(r.Oracles) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Oracles {
		if !v.Equal(r.Oracles[i]) {
			return false // fmt.Errorf("Oracles[%d] : %s", i, err)
		}
	}

	// Field MasterAddress - varbin
	if !bytes.Equal(c.MasterAddress, r.MasterAddress) {
		return false // fmt.Errorf("MasterAddress bytes mismatched")
	}

	// Field ContractRevision - uint
	if c.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field EntityContract - varbin
	if !bytes.Equal(c.EntityContract, r.EntityContract) {
		return false // fmt.Errorf("EntityContract bytes mismatched")
	}

	// Field OperatorEntityContract - varbin
	if !bytes.Equal(c.OperatorEntityContract, r.OperatorEntityContract) {
		return false // fmt.Errorf("OperatorEntityContract bytes mismatched")
	}

	// Field ContractType - uint
	if c.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType integer mismatched")
	}

	// Field Services - Service
	if len(c.Services) != len(r.Services) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Services {
		if !v.Equal(r.Services[i]) {
			return false // fmt.Errorf("Services[%d] : %s", i, err)
		}
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	if len(c.AdminIdentityCertificates) != len(r.AdminIdentityCertificates) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.AdminIdentityCertificates {
		if !v.Equal(r.AdminIdentityCertificates[i]) {
			return false // fmt.Errorf("AdminIdentityCertificates[%d] : %s", i, err)
		}
	}

	// Field AdminAddress - varbin
	if !bytes.Equal(c.AdminAddress, r.AdminAddress) {
		return false // fmt.Errorf("AdminAddress bytes mismatched")
	}

	// Field OperatorAddress - varbin
	if !bytes.Equal(c.OperatorAddress, r.OperatorAddress) {
		return false // fmt.Errorf("OperatorAddress bytes mismatched")
	}

	// Field GoverningLaw - varchar
	if c.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - varchar
	if c.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	return true
}

func (l *ContractAmendment) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ContractAmendment{}
	}
	cr := right
	if cr == nil {
		cr = &ContractAmendment{}
	}
	r, ok := cr.(*ContractAmendment)
	if !ok {
		return false
	}

	// Field ChangeAdministrationAddress - bool
	if c.ChangeAdministrationAddress != r.ChangeAdministrationAddress {
		return false // fmt.Errorf("ChangeAdministrationAddress boolean mismatched")
	}

	// Field ChangeOperatorAddress - bool
	if c.ChangeOperatorAddress != r.ChangeOperatorAddress {
		return false // fmt.Errorf("ChangeOperatorAddress boolean mismatched")
	}

	// Field ContractRevision - uint
	if c.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field Amendments - Amendment
	if len(c.Amendments) != len(r.Amendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Amendments {
		if !v.Equal(r.Amendments[i]) {
			return false // fmt.Errorf("Amendments[%d] : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if !bytes.Equal(c.RefTxID, r.RefTxID) {
		return false // fmt.Errorf("RefTxID bytes mismatched")
	}

	return true
}

func (l *StaticContractFormation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &StaticContractFormation{}
	}
	cr := right
	if cr == nil {
		cr = &StaticContractFormation{}
	}
	r, ok := cr.(*StaticContractFormation)
	if !ok {
		return false
	}

	// Field ContractName - varchar
	if c.ContractName != r.ContractName {
		return false // fmt.Errorf("ContractName string mismatched")
	}

	// Field ContractCode - bin
	if !bytes.Equal(c.ContractCode, r.ContractCode) {
		return false // fmt.Errorf("ContractCode bytes mismatched")
	}

	// Field BodyOfAgreementType - uint
	if c.BodyOfAgreementType != r.BodyOfAgreementType {
		return false // fmt.Errorf("BodyOfAgreementType integer mismatched")
	}

	// Field BodyOfAgreement - varbin
	if !bytes.Equal(c.BodyOfAgreement, r.BodyOfAgreement) {
		return false // fmt.Errorf("BodyOfAgreement bytes mismatched")
	}

	// Field ContractType - varchar
	if c.ContractType != r.ContractType {
		return false // fmt.Errorf("ContractType string mismatched")
	}

	// Field SupportingDocs - Document
	if len(c.SupportingDocs) != len(r.SupportingDocs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.SupportingDocs {
		if !v.Equal(r.SupportingDocs[i]) {
			return false // fmt.Errorf("SupportingDocs[%d] : %s", i, err)
		}
	}

	// Field ContractRevision - uint
	if c.ContractRevision != r.ContractRevision {
		return false // fmt.Errorf("ContractRevision integer mismatched")
	}

	// Field EffectiveDate - uint
	if c.EffectiveDate != r.EffectiveDate {
		return false // fmt.Errorf("EffectiveDate integer mismatched")
	}

	// Field ContractExpiration - uint
	if c.ContractExpiration != r.ContractExpiration {
		return false // fmt.Errorf("ContractExpiration integer mismatched")
	}

	// Field ContractURI - varchar
	if c.ContractURI != r.ContractURI {
		return false // fmt.Errorf("ContractURI string mismatched")
	}

	// Field PrevRevTxID - bin
	if !bytes.Equal(c.PrevRevTxID, r.PrevRevTxID) {
		return false // fmt.Errorf("PrevRevTxID bytes mismatched")
	}

	// Field Entities - Entity
	if len(c.Entities) != len(r.Entities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Entities {
		if !v.Equal(r.Entities[i]) {
			return false // fmt.Errorf("Entities[%d] : %s", i, err)
		}
	}

	// Field EntityOracle - Oracle
	if !c.EntityOracle.Equal(r.EntityOracle) {
		return false // fmt.Errorf("EntityOracle : %s", err)
	}

	// Field EntityOracleSignature - varbin
	if !bytes.Equal(c.EntityOracleSignature, r.EntityOracleSignature) {
		return false // fmt.Errorf("EntityOracleSignature bytes mismatched")
	}

	// Field EntityOracleSigBlockHeight - uint
	if c.EntityOracleSigBlockHeight != r.EntityOracleSigBlockHeight {
		return false // fmt.Errorf("EntityOracleSigBlockHeight integer mismatched")
	}

	// Field GoverningLaw - varchar
	if c.GoverningLaw != r.GoverningLaw {
		return false // fmt.Errorf("GoverningLaw string mismatched")
	}

	// Field Jurisdiction - varchar
	if c.Jurisdiction != r.Jurisdiction {
		return false // fmt.Errorf("Jurisdiction string mismatched")
	}

	return true
}

func (l *ContractAddressChange) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ContractAddressChange{}
	}
	cr := right
	if cr == nil {
		cr = &ContractAddressChange{}
	}
	r, ok := cr.(*ContractAddressChange)
	if !ok {
		return false
	}

	// Field NewContractAddress - varbin
	if !bytes.Equal(c.NewContractAddress, r.NewContractAddress) {
		return false // fmt.Errorf("NewContractAddress bytes mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *BodyOfAgreementOffer) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BodyOfAgreementOffer{}
	}
	cr := right
	if cr == nil {
		cr = &BodyOfAgreementOffer{}
	}
	r, ok := cr.(*BodyOfAgreementOffer)
	if !ok {
		return false
	}

	// Field Chapters - Chapter
	if len(c.Chapters) != len(r.Chapters) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Chapters {
		if !v.Equal(r.Chapters[i]) {
			return false // fmt.Errorf("Chapters[%d] : %s", i, err)
		}
	}

	// Field Definitions - DefinedTerm
	if len(c.Definitions) != len(r.Definitions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Definitions {
		if !v.Equal(r.Definitions[i]) {
			return false // fmt.Errorf("Definitions[%d] : %s", i, err)
		}
	}

	return true
}

func (l *BodyOfAgreementFormation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BodyOfAgreementFormation{}
	}
	cr := right
	if cr == nil {
		cr = &BodyOfAgreementFormation{}
	}
	r, ok := cr.(*BodyOfAgreementFormation)
	if !ok {
		return false
	}

	// Field Chapters - Chapter
	if len(c.Chapters) != len(r.Chapters) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Chapters {
		if !v.Equal(r.Chapters[i]) {
			return false // fmt.Errorf("Chapters[%d] : %s", i, err)
		}
	}

	// Field Definitions - DefinedTerm
	if len(c.Definitions) != len(r.Definitions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Definitions {
		if !v.Equal(r.Definitions[i]) {
			return false // fmt.Errorf("Definitions[%d] : %s", i, err)
		}
	}

	// Field Revision - uint
	if c.Revision != r.Revision {
		return false // fmt.Errorf("Revision integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *BodyOfAgreementAmendment) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BodyOfAgreementAmendment{}
	}
	cr := right
	if cr == nil {
		cr = &BodyOfAgreementAmendment{}
	}
	r, ok := cr.(*BodyOfAgreementAmendment)
	if !ok {
		return false
	}

	// Field Revision - uint
	if c.Revision != r.Revision {
		return false // fmt.Errorf("Revision integer mismatched")
	}

	// Field Amendments - Amendment
	if len(c.Amendments) != len(r.Amendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Amendments {
		if !v.Equal(r.Amendments[i]) {
			return false // fmt.Errorf("Amendments[%d] : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if !bytes.Equal(c.RefTxID, r.RefTxID) {
		return false // fmt.Errorf("RefTxID bytes mismatched")
	}

	return true
}

func (l *InstrumentDefinition) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentDefinition{}
	}
	cr := right
	if cr == nil {
		cr = &InstrumentDefinition{}
	}
	r, ok := cr.(*InstrumentDefinition)
	if !ok {
		return false
	}

	// Field InstrumentPermissions - varbin
	if !bytes.Equal(c.InstrumentPermissions, r.InstrumentPermissions) {
		return false // fmt.Errorf("InstrumentPermissions bytes mismatched")
	}

	// Field EnforcementOrdersPermitted - bool
	if c.EnforcementOrdersPermitted != r.EnforcementOrdersPermitted {
		return false // fmt.Errorf("EnforcementOrdersPermitted boolean mismatched")
	}

	// Field VotingRights - bool
	if c.VotingRights != r.VotingRights {
		return false // fmt.Errorf("VotingRights boolean mismatched")
	}

	// Field VoteMultiplier - uint
	if c.VoteMultiplier != r.VoteMultiplier {
		return false // fmt.Errorf("VoteMultiplier integer mismatched")
	}

	// Field AdministrationProposal - bool
	if c.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if c.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field InstrumentModificationGovernance - uint
	if c.InstrumentModificationGovernance != r.InstrumentModificationGovernance {
		return false // fmt.Errorf("InstrumentModificationGovernance integer mismatched")
	}

	// Field AuthorizedTokenQty - uint
	if c.AuthorizedTokenQty != r.AuthorizedTokenQty {
		return false // fmt.Errorf("AuthorizedTokenQty integer mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentPayload - varbin
	if !bytes.Equal(c.InstrumentPayload, r.InstrumentPayload) {
		return false // fmt.Errorf("InstrumentPayload bytes mismatched")
	}

	// Field TradeRestrictions - varchar
	if len(c.TradeRestrictions) != len(r.TradeRestrictions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.TradeRestrictions {
		if v != r.TradeRestrictions[i] {
			return false // fmt.Errorf("Element TradeRestrictions string mismatched")
		}
	}

	return true
}

func (l *InstrumentCreation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentCreation{}
	}
	cr := right
	if cr == nil {
		cr = &InstrumentCreation{}
	}
	r, ok := cr.(*InstrumentCreation)
	if !ok {
		return false
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field InstrumentIndex - uint
	if c.InstrumentIndex != r.InstrumentIndex {
		return false // fmt.Errorf("InstrumentIndex integer mismatched")
	}

	// Field InstrumentPermissions - varbin
	if !bytes.Equal(c.InstrumentPermissions, r.InstrumentPermissions) {
		return false // fmt.Errorf("InstrumentPermissions bytes mismatched")
	}

	// Field EnforcementOrdersPermitted - bool
	if c.EnforcementOrdersPermitted != r.EnforcementOrdersPermitted {
		return false // fmt.Errorf("EnforcementOrdersPermitted boolean mismatched")
	}

	// Field VotingRights - bool
	if c.VotingRights != r.VotingRights {
		return false // fmt.Errorf("VotingRights boolean mismatched")
	}

	// Field VoteMultiplier - uint
	if c.VoteMultiplier != r.VoteMultiplier {
		return false // fmt.Errorf("VoteMultiplier integer mismatched")
	}

	// Field AdministrationProposal - bool
	if c.AdministrationProposal != r.AdministrationProposal {
		return false // fmt.Errorf("AdministrationProposal boolean mismatched")
	}

	// Field HolderProposal - bool
	if c.HolderProposal != r.HolderProposal {
		return false // fmt.Errorf("HolderProposal boolean mismatched")
	}

	// Field InstrumentModificationGovernance - uint
	if c.InstrumentModificationGovernance != r.InstrumentModificationGovernance {
		return false // fmt.Errorf("InstrumentModificationGovernance integer mismatched")
	}

	// Field AuthorizedTokenQty - uint
	if c.AuthorizedTokenQty != r.AuthorizedTokenQty {
		return false // fmt.Errorf("AuthorizedTokenQty integer mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentPayload - varbin
	if !bytes.Equal(c.InstrumentPayload, r.InstrumentPayload) {
		return false // fmt.Errorf("InstrumentPayload bytes mismatched")
	}

	// Field InstrumentRevision - uint
	if c.InstrumentRevision != r.InstrumentRevision {
		return false // fmt.Errorf("InstrumentRevision integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field TradeRestrictions - varchar
	if len(c.TradeRestrictions) != len(r.TradeRestrictions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.TradeRestrictions {
		if v != r.TradeRestrictions[i] {
			return false // fmt.Errorf("Element TradeRestrictions string mismatched")
		}
	}

	return true
}

func (l *InstrumentModification) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentModification{}
	}
	cr := right
	if cr == nil {
		cr = &InstrumentModification{}
	}
	r, ok := cr.(*InstrumentModification)
	if !ok {
		return false
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field InstrumentRevision - uint
	if c.InstrumentRevision != r.InstrumentRevision {
		return false // fmt.Errorf("InstrumentRevision integer mismatched")
	}

	// Field Amendments - Amendment
	if len(c.Amendments) != len(r.Amendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Amendments {
		if !v.Equal(r.Amendments[i]) {
			return false // fmt.Errorf("Amendments[%d] : %s", i, err)
		}
	}

	// Field RefTxID - bin
	if !bytes.Equal(c.RefTxID, r.RefTxID) {
		return false // fmt.Errorf("RefTxID bytes mismatched")
	}

	return true
}

func (l *Transfer) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Transfer{}
	}
	cr := right
	if cr == nil {
		cr = &Transfer{}
	}
	r, ok := cr.(*Transfer)
	if !ok {
		return false
	}

	// Field Instruments - InstrumentTransfer
	if len(c.Instruments) != len(r.Instruments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Instruments {
		if !v.Equal(r.Instruments[i]) {
			return false // fmt.Errorf("Instruments[%d] : %s", i, err)
		}
	}

	// Field OfferExpiry - uint
	if c.OfferExpiry != r.OfferExpiry {
		return false // fmt.Errorf("OfferExpiry integer mismatched")
	}

	// Field ExchangeFee - uint
	if c.ExchangeFee != r.ExchangeFee {
		return false // fmt.Errorf("ExchangeFee integer mismatched")
	}

	// Field ExchangeFeeAddress - varbin
	if !bytes.Equal(c.ExchangeFeeAddress, r.ExchangeFeeAddress) {
		return false // fmt.Errorf("ExchangeFeeAddress bytes mismatched")
	}

	return true
}

func (l *Settlement) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Settlement{}
	}
	cr := right
	if cr == nil {
		cr = &Settlement{}
	}
	r, ok := cr.(*Settlement)
	if !ok {
		return false
	}

	// Field Instruments - InstrumentSettlement
	if len(c.Instruments) != len(r.Instruments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Instruments {
		if !v.Equal(r.Instruments[i]) {
			return false // fmt.Errorf("Instruments[%d] : %s", i, err)
		}
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Proposal) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Proposal{}
	}
	cr := right
	if cr == nil {
		cr = &Proposal{}
	}
	r, ok := cr.(*Proposal)
	if !ok {
		return false
	}

	// Field Type - uint
	if c.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field VoteSystem - uint
	if c.VoteSystem != r.VoteSystem {
		return false // fmt.Errorf("VoteSystem integer mismatched")
	}

	// Field ProposedAmendments - Amendment
	if len(c.ProposedAmendments) != len(r.ProposedAmendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ProposedAmendments {
		if !v.Equal(r.ProposedAmendments[i]) {
			return false // fmt.Errorf("ProposedAmendments[%d] : %s", i, err)
		}
	}

	// Field VoteOptions - varchar
	if c.VoteOptions != r.VoteOptions {
		return false // fmt.Errorf("VoteOptions string mismatched")
	}

	// Field VoteMax - uint
	if c.VoteMax != r.VoteMax {
		return false // fmt.Errorf("VoteMax integer mismatched")
	}

	// Field ProposalDescription - varchar
	if c.ProposalDescription != r.ProposalDescription {
		return false // fmt.Errorf("ProposalDescription string mismatched")
	}

	// Field ProposalDocumentHash - bin
	if !bytes.Equal(c.ProposalDocumentHash, r.ProposalDocumentHash) {
		return false // fmt.Errorf("ProposalDocumentHash bytes mismatched")
	}

	// Field VoteCutOffTimestamp - uint
	if c.VoteCutOffTimestamp != r.VoteCutOffTimestamp {
		return false // fmt.Errorf("VoteCutOffTimestamp integer mismatched")
	}

	return true
}

func (l *Vote) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Vote{}
	}
	cr := right
	if cr == nil {
		cr = &Vote{}
	}
	r, ok := cr.(*Vote)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *BallotCast) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BallotCast{}
	}
	cr := right
	if cr == nil {
		cr = &BallotCast{}
	}
	r, ok := cr.(*BallotCast)
	if !ok {
		return false
	}

	// Field VoteTxId - bin
	if !bytes.Equal(c.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field Vote - varchar
	if c.Vote != r.Vote {
		return false // fmt.Errorf("Vote string mismatched")
	}

	return true
}

func (l *BallotCounted) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &BallotCounted{}
	}
	cr := right
	if cr == nil {
		cr = &BallotCounted{}
	}
	r, ok := cr.(*BallotCounted)
	if !ok {
		return false
	}

	// Field VoteTxId - bin
	if !bytes.Equal(c.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field Vote - varchar
	if c.Vote != r.Vote {
		return false // fmt.Errorf("Vote string mismatched")
	}

	// Field Quantity - uint
	if c.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Result) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Result{}
	}
	cr := right
	if cr == nil {
		cr = &Result{}
	}
	r, ok := cr.(*Result)
	if !ok {
		return false
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field ProposedAmendments - Amendment
	if len(c.ProposedAmendments) != len(r.ProposedAmendments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ProposedAmendments {
		if !v.Equal(r.ProposedAmendments[i]) {
			return false // fmt.Errorf("ProposedAmendments[%d] : %s", i, err)
		}
	}

	// Field VoteTxId - bin
	if !bytes.Equal(c.VoteTxId, r.VoteTxId) {
		return false // fmt.Errorf("VoteTxId bytes mismatched")
	}

	// Field OptionTally - uint
	if len(c.OptionTally) != len(r.OptionTally) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.OptionTally {
		if v != r.OptionTally[i] {
			return false // fmt.Errorf("Element OptionTally integer mismatched")
		}
	}

	// Field Result - varchar
	if c.Result != r.Result {
		return false // fmt.Errorf("Result string mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Order) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Order{}
	}
	cr := right
	if cr == nil {
		cr = &Order{}
	}
	r, ok := cr.(*Order)
	if !ok {
		return false
	}

	// Field ComplianceAction - fixedchar
	if c.ComplianceAction != r.ComplianceAction {
		return false // fmt.Errorf("ComplianceAction string mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field TargetAddresses - TargetAddress
	if len(c.TargetAddresses) != len(r.TargetAddresses) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.TargetAddresses {
		if !v.Equal(r.TargetAddresses[i]) {
			return false // fmt.Errorf("TargetAddresses[%d] : %s", i, err)
		}
	}

	// Field FreezeTxId - bin
	if !bytes.Equal(c.FreezeTxId, r.FreezeTxId) {
		return false // fmt.Errorf("FreezeTxId bytes mismatched")
	}

	// Field FreezePeriod - uint
	if c.FreezePeriod != r.FreezePeriod {
		return false // fmt.Errorf("FreezePeriod integer mismatched")
	}

	// Field DepositAddress - varbin
	if !bytes.Equal(c.DepositAddress, r.DepositAddress) {
		return false // fmt.Errorf("DepositAddress bytes mismatched")
	}

	// Field AuthorityName - varchar
	if c.AuthorityName != r.AuthorityName {
		return false // fmt.Errorf("AuthorityName string mismatched")
	}

	// Field AuthorityPublicKey - varbin
	if !bytes.Equal(c.AuthorityPublicKey, r.AuthorityPublicKey) {
		return false // fmt.Errorf("AuthorityPublicKey bytes mismatched")
	}

	// Field SignatureAlgorithm - uint
	if c.SignatureAlgorithm != r.SignatureAlgorithm {
		return false // fmt.Errorf("SignatureAlgorithm integer mismatched")
	}

	// Field OrderSignature - varbin
	if !bytes.Equal(c.OrderSignature, r.OrderSignature) {
		return false // fmt.Errorf("OrderSignature bytes mismatched")
	}

	// Field BitcoinDispersions - QuantityIndex
	if len(c.BitcoinDispersions) != len(r.BitcoinDispersions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.BitcoinDispersions {
		if !v.Equal(r.BitcoinDispersions[i]) {
			return false // fmt.Errorf("BitcoinDispersions[%d] : %s", i, err)
		}
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	// Field SupportingEvidenceFormat - uint
	if c.SupportingEvidenceFormat != r.SupportingEvidenceFormat {
		return false // fmt.Errorf("SupportingEvidenceFormat integer mismatched")
	}

	// Field SupportingEvidence - varbin
	if !bytes.Equal(c.SupportingEvidence, r.SupportingEvidence) {
		return false // fmt.Errorf("SupportingEvidence bytes mismatched")
	}

	// Field ReferenceTransactions - ReferenceTransaction
	if len(c.ReferenceTransactions) != len(r.ReferenceTransactions) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ReferenceTransactions {
		if !v.Equal(r.ReferenceTransactions[i]) {
			return false // fmt.Errorf("ReferenceTransactions[%d] : %s", i, err)
		}
	}

	return true
}

func (l *Freeze) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Freeze{}
	}
	cr := right
	if cr == nil {
		cr = &Freeze{}
	}
	r, ok := cr.(*Freeze)
	if !ok {
		return false
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(c.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field FreezePeriod - uint
	if c.FreezePeriod != r.FreezePeriod {
		return false // fmt.Errorf("FreezePeriod integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Thaw) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Thaw{}
	}
	cr := right
	if cr == nil {
		cr = &Thaw{}
	}
	r, ok := cr.(*Thaw)
	if !ok {
		return false
	}

	// Field FreezeTxId - bin
	if !bytes.Equal(c.FreezeTxId, r.FreezeTxId) {
		return false // fmt.Errorf("FreezeTxId bytes mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Confiscation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Confiscation{}
	}
	cr := right
	if cr == nil {
		cr = &Confiscation{}
	}
	r, ok := cr.(*Confiscation)
	if !ok {
		return false
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(c.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field DepositQty - uint
	if c.DepositQty != r.DepositQty {
		return false // fmt.Errorf("DepositQty integer mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Reconciliation) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Reconciliation{}
	}
	cr := right
	if cr == nil {
		cr = &Reconciliation{}
	}
	r, ok := cr.(*Reconciliation)
	if !ok {
		return false
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field Quantities - QuantityIndex
	if len(c.Quantities) != len(r.Quantities) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Quantities {
		if !v.Equal(r.Quantities[i]) {
			return false // fmt.Errorf("Quantities[%d] : %s", i, err)
		}
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *Establishment) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Establishment{}
	}
	cr := right
	if cr == nil {
		cr = &Establishment{}
	}
	r, ok := cr.(*Establishment)
	if !ok {
		return false
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Addition) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Addition{}
	}
	cr := right
	if cr == nil {
		cr = &Addition{}
	}
	r, ok := cr.(*Addition)
	if !ok {
		return false
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Alteration) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Alteration{}
	}
	cr := right
	if cr == nil {
		cr = &Alteration{}
	}
	r, ok := cr.(*Alteration)
	if !ok {
		return false
	}

	// Field EntryTxID - bin
	if !bytes.Equal(c.EntryTxID, r.EntryTxID) {
		return false // fmt.Errorf("EntryTxID bytes mismatched")
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Removal) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Removal{}
	}
	cr := right
	if cr == nil {
		cr = &Removal{}
	}
	r, ok := cr.(*Removal)
	if !ok {
		return false
	}

	// Field EntryTxID - bin
	if !bytes.Equal(c.EntryTxID, r.EntryTxID) {
		return false // fmt.Errorf("EntryTxID bytes mismatched")
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	return true
}

func (l *Message) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Message{}
	}
	cr := right
	if cr == nil {
		cr = &Message{}
	}
	r, ok := cr.(*Message)
	if !ok {
		return false
	}

	// Field SenderIndexes - uint
	if len(c.SenderIndexes) != len(r.SenderIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.SenderIndexes {
		if v != r.SenderIndexes[i] {
			return false // fmt.Errorf("Element SenderIndexes integer mismatched")
		}
	}

	// Field ReceiverIndexes - uint
	if len(c.ReceiverIndexes) != len(r.ReceiverIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ReceiverIndexes {
		if v != r.ReceiverIndexes[i] {
			return false // fmt.Errorf("Element ReceiverIndexes integer mismatched")
		}
	}

	// Field MessageCode - uint
	if c.MessageCode != r.MessageCode {
		return false // fmt.Errorf("MessageCode integer mismatched")
	}

	// Field MessagePayload - varbin
	if !bytes.Equal(c.MessagePayload, r.MessagePayload) {
		return false // fmt.Errorf("MessagePayload bytes mismatched")
	}

	return true
}

func (l *Rejection) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Rejection{}
	}
	cr := right
	if cr == nil {
		cr = &Rejection{}
	}
	r, ok := cr.(*Rejection)
	if !ok {
		return false
	}

	// Field AddressIndexes - uint
	if len(c.AddressIndexes) != len(r.AddressIndexes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.AddressIndexes {
		if v != r.AddressIndexes[i] {
			return false // fmt.Errorf("Element AddressIndexes integer mismatched")
		}
	}

	// Field RejectAddressIndex - uint
	if c.RejectAddressIndex != r.RejectAddressIndex {
		return false // fmt.Errorf("RejectAddressIndex integer mismatched")
	}

	// Field RejectionCode - uint
	if c.RejectionCode != r.RejectionCode {
		return false // fmt.Errorf("RejectionCode integer mismatched")
	}

	// Field Message - varchar
	if c.Message != r.Message {
		return false // fmt.Errorf("Message string mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *AdministratorField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &AdministratorField{}
	}
	r, ok := right.(*AdministratorField)
	if !ok {
		return false
	}

	if r == nil {
		r = &AdministratorField{}
	}

	// Field Type - uint
	if c.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	return true
}

func (l *AdminIdentityCertificateField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &AdminIdentityCertificateField{}
	}
	r, ok := right.(*AdminIdentityCertificateField)
	if !ok {
		return false
	}

	if r == nil {
		r = &AdminIdentityCertificateField{}
	}

	// Field EntityContract - varbin
	if !bytes.Equal(c.EntityContract, r.EntityContract) {
		return false // fmt.Errorf("EntityContract bytes mismatched")
	}

	// Field Signature - varbin
	if !bytes.Equal(c.Signature, r.Signature) {
		return false // fmt.Errorf("Signature bytes mismatched")
	}

	// Field BlockHeight - uint
	if c.BlockHeight != r.BlockHeight {
		return false // fmt.Errorf("BlockHeight integer mismatched")
	}

	// Field Expiration - uint
	if c.Expiration != r.Expiration {
		return false // fmt.Errorf("Expiration integer mismatched")
	}

	return true
}

func (l *AmendmentField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &AmendmentField{}
	}
	r, ok := right.(*AmendmentField)
	if !ok {
		return false
	}

	if r == nil {
		r = &AmendmentField{}
	}

	// Field FieldIndexPath - varbin
	if !bytes.Equal(c.FieldIndexPath, r.FieldIndexPath) {
		return false // fmt.Errorf("FieldIndexPath bytes mismatched")
	}

	// Field Operation - uint
	if c.Operation != r.Operation {
		return false // fmt.Errorf("Operation integer mismatched")
	}

	// Field Data - varbin
	if !bytes.Equal(c.Data, r.Data) {
		return false // fmt.Errorf("Data bytes mismatched")
	}

	return true
}

func (l *InstrumentReceiverField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentReceiverField{}
	}
	r, ok := right.(*InstrumentReceiverField)
	if !ok {
		return false
	}

	if r == nil {
		r = &InstrumentReceiverField{}
	}

	// Field Address - varbin
	if !bytes.Equal(c.Address, r.Address) {
		return false // fmt.Errorf("Address bytes mismatched")
	}

	// Field Quantity - uint
	if c.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	// Field OracleSigAlgorithm - uint
	if c.OracleSigAlgorithm != r.OracleSigAlgorithm {
		return false // fmt.Errorf("OracleSigAlgorithm integer mismatched")
	}

	// Field OracleIndex - uint
	if c.OracleIndex != r.OracleIndex {
		return false // fmt.Errorf("OracleIndex integer mismatched")
	}

	// Field OracleConfirmationSig - varbin
	if !bytes.Equal(c.OracleConfirmationSig, r.OracleConfirmationSig) {
		return false // fmt.Errorf("OracleConfirmationSig bytes mismatched")
	}

	// Field OracleSigBlockHeight - uint
	if c.OracleSigBlockHeight != r.OracleSigBlockHeight {
		return false // fmt.Errorf("OracleSigBlockHeight integer mismatched")
	}

	// Field OracleSigExpiry - uint
	if c.OracleSigExpiry != r.OracleSigExpiry {
		return false // fmt.Errorf("OracleSigExpiry integer mismatched")
	}

	return true
}

func (l *InstrumentSettlementField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentSettlementField{}
	}
	r, ok := right.(*InstrumentSettlementField)
	if !ok {
		return false
	}

	if r == nil {
		r = &InstrumentSettlementField{}
	}

	// Field ContractIndex - uint
	if c.ContractIndex != r.ContractIndex {
		return false // fmt.Errorf("ContractIndex integer mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field Settlements - QuantityIndex
	if len(c.Settlements) != len(r.Settlements) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Settlements {
		if !v.Equal(r.Settlements[i]) {
			return false // fmt.Errorf("Settlements[%d] : %s", i, err)
		}
	}

	return true
}

func (l *InstrumentTransferField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InstrumentTransferField{}
	}
	r, ok := right.(*InstrumentTransferField)
	if !ok {
		return false
	}

	if r == nil {
		r = &InstrumentTransferField{}
	}

	// Field ContractIndex - uint
	if c.ContractIndex != r.ContractIndex {
		return false // fmt.Errorf("ContractIndex integer mismatched")
	}

	// Field InstrumentType - fixedchar
	if c.InstrumentType != r.InstrumentType {
		return false // fmt.Errorf("InstrumentType string mismatched")
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field InstrumentSenders - QuantityIndex
	if len(c.InstrumentSenders) != len(r.InstrumentSenders) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.InstrumentSenders {
		if !v.Equal(r.InstrumentSenders[i]) {
			return false // fmt.Errorf("InstrumentSenders[%d] : %s", i, err)
		}
	}

	// Field InstrumentReceivers - InstrumentReceiver
	if len(c.InstrumentReceivers) != len(r.InstrumentReceivers) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.InstrumentReceivers {
		if !v.Equal(r.InstrumentReceivers[i]) {
			return false // fmt.Errorf("InstrumentReceivers[%d] : %s", i, err)
		}
	}

	return true
}

func (l *ChapterField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ChapterField{}
	}
	r, ok := right.(*ChapterField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ChapterField{}
	}

	// Field Title - varchar
	if c.Title != r.Title {
		return false // fmt.Errorf("Title string mismatched")
	}

	// Field Preamble - varchar
	if c.Preamble != r.Preamble {
		return false // fmt.Errorf("Preamble string mismatched")
	}

	// Field Articles - Clause
	if len(c.Articles) != len(r.Articles) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Articles {
		if !v.Equal(r.Articles[i]) {
			return false // fmt.Errorf("Articles[%d] : %s", i, err)
		}
	}

	return true
}

func (l *ClauseField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ClauseField{}
	}
	r, ok := right.(*ClauseField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ClauseField{}
	}

	// Field Title - varchar
	if c.Title != r.Title {
		return false // fmt.Errorf("Title string mismatched")
	}

	// Field Body - varchar
	if c.Body != r.Body {
		return false // fmt.Errorf("Body string mismatched")
	}

	// Field Children - Clause
	if len(c.Children) != len(r.Children) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Children {
		if !v.Equal(r.Children[i]) {
			return false // fmt.Errorf("Children[%d] : %s", i, err)
		}
	}

	return true
}

func (l *DefinedTermField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &DefinedTermField{}
	}
	r, ok := right.(*DefinedTermField)
	if !ok {
		return false
	}

	if r == nil {
		r = &DefinedTermField{}
	}

	// Field Term - varchar
	if c.Term != r.Term {
		return false // fmt.Errorf("Term string mismatched")
	}

	// Field Definition - varchar
	if c.Definition != r.Definition {
		return false // fmt.Errorf("Definition string mismatched")
	}

	return true
}

func (l *DocumentField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &DocumentField{}
	}
	r, ok := right.(*DocumentField)
	if !ok {
		return false
	}

	if r == nil {
		r = &DocumentField{}
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - varchar
	if c.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field Contents - varbin
	if !bytes.Equal(c.Contents, r.Contents) {
		return false // fmt.Errorf("Contents bytes mismatched")
	}

	return true
}

func (l *EntityField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &EntityField{}
	}
	r, ok := right.(*EntityField)
	if !ok {
		return false
	}

	if r == nil {
		r = &EntityField{}
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - fixedchar
	if c.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field LEI - fixedchar
	if c.LEI != r.LEI {
		return false // fmt.Errorf("LEI string mismatched")
	}

	// Field UnitNumber - varchar
	if c.UnitNumber != r.UnitNumber {
		return false // fmt.Errorf("UnitNumber string mismatched")
	}

	// Field BuildingNumber - varchar
	if c.BuildingNumber != r.BuildingNumber {
		return false // fmt.Errorf("BuildingNumber string mismatched")
	}

	// Field Street - varchar
	if c.Street != r.Street {
		return false // fmt.Errorf("Street string mismatched")
	}

	// Field SuburbCity - varchar
	if c.SuburbCity != r.SuburbCity {
		return false // fmt.Errorf("SuburbCity string mismatched")
	}

	// Field TerritoryStateProvinceCode - fixedchar
	if c.TerritoryStateProvinceCode != r.TerritoryStateProvinceCode {
		return false // fmt.Errorf("TerritoryStateProvinceCode string mismatched")
	}

	// Field CountryCode - fixedchar
	if c.CountryCode != r.CountryCode {
		return false // fmt.Errorf("CountryCode string mismatched")
	}

	// Field PostalZIPCode - fixedchar
	if c.PostalZIPCode != r.PostalZIPCode {
		return false // fmt.Errorf("PostalZIPCode string mismatched")
	}

	// Field EmailAddress - varchar
	if c.EmailAddress != r.EmailAddress {
		return false // fmt.Errorf("EmailAddress string mismatched")
	}

	// Field PhoneNumber - varchar
	if c.PhoneNumber != r.PhoneNumber {
		return false // fmt.Errorf("PhoneNumber string mismatched")
	}

	// Field Administration - Administrator
	if len(c.Administration) != len(r.Administration) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Administration {
		if !v.Equal(r.Administration[i]) {
			return false // fmt.Errorf("Administration[%d] : %s", i, err)
		}
	}

	// Field Management - Manager
	if len(c.Management) != len(r.Management) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Management {
		if !v.Equal(r.Management[i]) {
			return false // fmt.Errorf("Management[%d] : %s", i, err)
		}
	}

	// Field DomainName - varchar
	if c.DomainName != r.DomainName {
		return false // fmt.Errorf("DomainName string mismatched")
	}

	// Field PaymailHandle - varchar
	if c.PaymailHandle != r.PaymailHandle {
		return false // fmt.Errorf("PaymailHandle string mismatched")
	}

	return true
}

func (l *ManagerField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ManagerField{}
	}
	r, ok := right.(*ManagerField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ManagerField{}
	}

	// Field Type - uint
	if c.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	return true
}

func (l *OracleField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &OracleField{}
	}
	r, ok := right.(*OracleField)
	if !ok {
		return false
	}

	if r == nil {
		r = &OracleField{}
	}

	// Field OracleTypes - uint
	if len(c.OracleTypes) != len(r.OracleTypes) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.OracleTypes {
		if v != r.OracleTypes[i] {
			return false // fmt.Errorf("Element OracleTypes integer mismatched")
		}
	}

	// Field EntityContract - varbin
	if !bytes.Equal(c.EntityContract, r.EntityContract) {
		return false // fmt.Errorf("EntityContract bytes mismatched")
	}

	return true
}

func (l *QuantityIndexField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &QuantityIndexField{}
	}
	r, ok := right.(*QuantityIndexField)
	if !ok {
		return false
	}

	if r == nil {
		r = &QuantityIndexField{}
	}

	// Field Index - uint
	if c.Index != r.Index {
		return false // fmt.Errorf("Index integer mismatched")
	}

	// Field Quantity - uint
	if c.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	return true
}

func (l *ReferenceTransactionField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ReferenceTransactionField{}
	}
	r, ok := right.(*ReferenceTransactionField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ReferenceTransactionField{}
	}

	// Field Transaction - varbin
	if !bytes.Equal(c.Transaction, r.Transaction) {
		return false // fmt.Errorf("Transaction bytes mismatched")
	}

	// Field Outputs - varbin
	if len(c.Outputs) != len(r.Outputs) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Outputs {
		if !bytes.Equal(v, r.Outputs[i]) {
			return false // fmt.Errorf("Element Outputs bytes mismatched")
		}
	}

	return true
}

func (l *ServiceField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ServiceField{}
	}
	r, ok := right.(*ServiceField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ServiceField{}
	}

	// Field Type - uint
	if c.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field URL - varchar
	if c.URL != r.URL {
		return false // fmt.Errorf("URL string mismatched")
	}

	// Field PublicKey - bin
	if !bytes.Equal(c.PublicKey, r.PublicKey) {
		return false // fmt.Errorf("PublicKey bytes mismatched")
	}

	return true
}

func (l *TargetAddressField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &TargetAddressField{}
	}
	r, ok := right.(*TargetAddressField)
	if !ok {
		return false
	}

	if r == nil {
		r = &TargetAddressField{}
	}

	// Field Address - varbin
	if !bytes.Equal(c.Address, r.Address) {
		return false // fmt.Errorf("Address bytes mismatched")
	}

	// Field Quantity - uint
	if c.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	return true
}

func (l *VotingSystemField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &VotingSystemField{}
	}
	r, ok := right.(*VotingSystemField)
	if !ok {
		return false
	}

	if r == nil {
		r = &VotingSystemField{}
	}

	// Field Name - varchar
	if c.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field VoteType - fixedchar
	if c.VoteType != r.VoteType {
		return false // fmt.Errorf("VoteType string mismatched")
	}

	// Field TallyLogic - uint
	if c.TallyLogic != r.TallyLogic {
		return false // fmt.Errorf("TallyLogic integer mismatched")
	}

	// Field ThresholdPercentage - uint
	if c.ThresholdPercentage != r.ThresholdPercentage {
		return false // fmt.Errorf("ThresholdPercentage integer mismatched")
	}

	// Field VoteMultiplierPermitted - bool
	if c.VoteMultiplierPermitted != r.VoteMultiplierPermitted {
		return false // fmt.Errorf("VoteMultiplierPermitted boolean mismatched")
	}

	// Field HolderProposalFee - uint
	if c.HolderProposalFee != r.HolderProposalFee {
		return false // fmt.Errorf("HolderProposalFee integer mismatched")
	}

	return true
}
