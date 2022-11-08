package actions

func (a *ContractOffer) Copy() *ContractOffer {
	result := &ContractOffer{}

	// Field ContractName - varchar
	result.ContractName = a.ContractName

	// Field BodyOfAgreementType - uint
	result.BodyOfAgreementType = a.BodyOfAgreementType

	// Field BodyOfAgreement - varbin
	result.BodyOfAgreement = make([]byte, len(a.BodyOfAgreement))
	copy(result.BodyOfAgreement, a.BodyOfAgreement)

	// Field SupportingDocs - Document
	result.SupportingDocs = make([]*DocumentField, len(a.SupportingDocs))
	for i, v := range a.SupportingDocs {
		result.SupportingDocs[i] = v.Copy()
	}

	// Field ContractExpiration - uint
	result.ContractExpiration = a.ContractExpiration

	// Field ContractURI - varchar
	result.ContractURI = a.ContractURI

	// Field Issuer - Entity
	result.Issuer = a.Issuer.Copy()

	// Field ContractOperatorIncluded - bool
	result.ContractOperatorIncluded = a.ContractOperatorIncluded

	// Field ContractFee - uint
	result.ContractFee = a.ContractFee

	// Field VotingSystems - VotingSystem
	result.VotingSystems = make([]*VotingSystemField, len(a.VotingSystems))
	for i, v := range a.VotingSystems {
		result.VotingSystems[i] = v.Copy()
	}

	// Field ContractPermissions - varbin
	result.ContractPermissions = make([]byte, len(a.ContractPermissions))
	copy(result.ContractPermissions, a.ContractPermissions)

	// Field RestrictedQtyInstruments - uint
	result.RestrictedQtyInstruments = a.RestrictedQtyInstruments

	// Field AdministrationProposal - bool
	result.AdministrationProposal = a.AdministrationProposal

	// Field HolderProposal - bool
	result.HolderProposal = a.HolderProposal

	// Field Oracles - Oracle
	result.Oracles = make([]*OracleField, len(a.Oracles))
	for i, v := range a.Oracles {
		result.Oracles[i] = v.Copy()
	}

	// Field MasterAddress - varbin
	result.MasterAddress = make([]byte, len(a.MasterAddress))
	copy(result.MasterAddress, a.MasterAddress)

	// Field EntityContract - varbin
	result.EntityContract = make([]byte, len(a.EntityContract))
	copy(result.EntityContract, a.EntityContract)

	// Field OperatorEntityContract - varbin
	result.OperatorEntityContract = make([]byte, len(a.OperatorEntityContract))
	copy(result.OperatorEntityContract, a.OperatorEntityContract)

	// Field ContractType - uint
	result.ContractType = a.ContractType

	// Field Services - Service
	result.Services = make([]*ServiceField, len(a.Services))
	for i, v := range a.Services {
		result.Services[i] = v.Copy()
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	result.AdminIdentityCertificates = make([]*AdminIdentityCertificateField, len(a.AdminIdentityCertificates))
	for i, v := range a.AdminIdentityCertificates {
		result.AdminIdentityCertificates[i] = v.Copy()
	}

	// Field GoverningLaw - varchar
	result.GoverningLaw = a.GoverningLaw

	// Field Jurisdiction - varchar
	result.Jurisdiction = a.Jurisdiction

	return result
}

func (a *ContractFormation) Copy() *ContractFormation {
	result := &ContractFormation{}

	// Field ContractName - varchar
	result.ContractName = a.ContractName

	// Field BodyOfAgreementType - uint
	result.BodyOfAgreementType = a.BodyOfAgreementType

	// Field BodyOfAgreement - varbin
	result.BodyOfAgreement = make([]byte, len(a.BodyOfAgreement))
	copy(result.BodyOfAgreement, a.BodyOfAgreement)

	// Field SupportingDocs - Document
	result.SupportingDocs = make([]*DocumentField, len(a.SupportingDocs))
	for i, v := range a.SupportingDocs {
		result.SupportingDocs[i] = v.Copy()
	}

	// Field ContractExpiration - uint
	result.ContractExpiration = a.ContractExpiration

	// Field ContractURI - varchar
	result.ContractURI = a.ContractURI

	// Field Issuer - Entity
	result.Issuer = a.Issuer.Copy()

	// Field ContractFee - uint
	result.ContractFee = a.ContractFee

	// Field VotingSystems - VotingSystem
	result.VotingSystems = make([]*VotingSystemField, len(a.VotingSystems))
	for i, v := range a.VotingSystems {
		result.VotingSystems[i] = v.Copy()
	}

	// Field ContractPermissions - varbin
	result.ContractPermissions = make([]byte, len(a.ContractPermissions))
	copy(result.ContractPermissions, a.ContractPermissions)

	// Field RestrictedQtyInstruments - uint
	result.RestrictedQtyInstruments = a.RestrictedQtyInstruments

	// Field AdministrationProposal - bool
	result.AdministrationProposal = a.AdministrationProposal

	// Field HolderProposal - bool
	result.HolderProposal = a.HolderProposal

	// Field Oracles - Oracle
	result.Oracles = make([]*OracleField, len(a.Oracles))
	for i, v := range a.Oracles {
		result.Oracles[i] = v.Copy()
	}

	// Field MasterAddress - varbin
	result.MasterAddress = make([]byte, len(a.MasterAddress))
	copy(result.MasterAddress, a.MasterAddress)

	// Field ContractRevision - uint
	result.ContractRevision = a.ContractRevision

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field EntityContract - varbin
	result.EntityContract = make([]byte, len(a.EntityContract))
	copy(result.EntityContract, a.EntityContract)

	// Field OperatorEntityContract - varbin
	result.OperatorEntityContract = make([]byte, len(a.OperatorEntityContract))
	copy(result.OperatorEntityContract, a.OperatorEntityContract)

	// Field ContractType - uint
	result.ContractType = a.ContractType

	// Field Services - Service
	result.Services = make([]*ServiceField, len(a.Services))
	for i, v := range a.Services {
		result.Services[i] = v.Copy()
	}

	// Field AdminIdentityCertificates - AdminIdentityCertificate
	result.AdminIdentityCertificates = make([]*AdminIdentityCertificateField, len(a.AdminIdentityCertificates))
	for i, v := range a.AdminIdentityCertificates {
		result.AdminIdentityCertificates[i] = v.Copy()
	}

	// Field AdminAddress - varbin
	result.AdminAddress = make([]byte, len(a.AdminAddress))
	copy(result.AdminAddress, a.AdminAddress)

	// Field OperatorAddress - varbin
	result.OperatorAddress = make([]byte, len(a.OperatorAddress))
	copy(result.OperatorAddress, a.OperatorAddress)

	// Field GoverningLaw - varchar
	result.GoverningLaw = a.GoverningLaw

	// Field Jurisdiction - varchar
	result.Jurisdiction = a.Jurisdiction

	return result
}

func (a *ContractAmendment) Copy() *ContractAmendment {
	result := &ContractAmendment{}

	// Field ChangeAdministrationAddress - bool
	result.ChangeAdministrationAddress = a.ChangeAdministrationAddress

	// Field ChangeOperatorAddress - bool
	result.ChangeOperatorAddress = a.ChangeOperatorAddress

	// Field ContractRevision - uint
	result.ContractRevision = a.ContractRevision

	// Field Amendments - Amendment
	result.Amendments = make([]*AmendmentField, len(a.Amendments))
	for i, v := range a.Amendments {
		result.Amendments[i] = v.Copy()
	}

	// Field RefTxID - bin
	result.RefTxID = make([]byte, len(a.RefTxID))
	copy(result.RefTxID, a.RefTxID)

	return result
}

func (a *StaticContractFormation) Copy() *StaticContractFormation {
	result := &StaticContractFormation{}

	// Field ContractName - varchar
	result.ContractName = a.ContractName

	// Field ContractCode - bin
	result.ContractCode = make([]byte, len(a.ContractCode))
	copy(result.ContractCode, a.ContractCode)

	// Field BodyOfAgreementType - uint
	result.BodyOfAgreementType = a.BodyOfAgreementType

	// Field BodyOfAgreement - varbin
	result.BodyOfAgreement = make([]byte, len(a.BodyOfAgreement))
	copy(result.BodyOfAgreement, a.BodyOfAgreement)

	// Field ContractType - varchar
	result.ContractType = a.ContractType

	// Field SupportingDocs - Document
	result.SupportingDocs = make([]*DocumentField, len(a.SupportingDocs))
	for i, v := range a.SupportingDocs {
		result.SupportingDocs[i] = v.Copy()
	}

	// Field ContractRevision - uint
	result.ContractRevision = a.ContractRevision

	// Field EffectiveDate - uint
	result.EffectiveDate = a.EffectiveDate

	// Field ContractExpiration - uint
	result.ContractExpiration = a.ContractExpiration

	// Field ContractURI - varchar
	result.ContractURI = a.ContractURI

	// Field PrevRevTxID - bin
	result.PrevRevTxID = make([]byte, len(a.PrevRevTxID))
	copy(result.PrevRevTxID, a.PrevRevTxID)

	// Field Entities - Entity
	result.Entities = make([]*EntityField, len(a.Entities))
	for i, v := range a.Entities {
		result.Entities[i] = v.Copy()
	}

	// Field EntityOracle - Oracle
	result.EntityOracle = a.EntityOracle.Copy()

	// Field EntityOracleSignature - varbin
	result.EntityOracleSignature = make([]byte, len(a.EntityOracleSignature))
	copy(result.EntityOracleSignature, a.EntityOracleSignature)

	// Field EntityOracleSigBlockHeight - uint
	result.EntityOracleSigBlockHeight = a.EntityOracleSigBlockHeight

	// Field GoverningLaw - varchar
	result.GoverningLaw = a.GoverningLaw

	// Field Jurisdiction - varchar
	result.Jurisdiction = a.Jurisdiction

	return result
}

func (a *ContractAddressChange) Copy() *ContractAddressChange {
	result := &ContractAddressChange{}

	// Field NewContractAddress - varbin
	result.NewContractAddress = make([]byte, len(a.NewContractAddress))
	copy(result.NewContractAddress, a.NewContractAddress)

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *BodyOfAgreementOffer) Copy() *BodyOfAgreementOffer {
	result := &BodyOfAgreementOffer{}

	// Field Chapters - Chapter
	result.Chapters = make([]*ChapterField, len(a.Chapters))
	for i, v := range a.Chapters {
		result.Chapters[i] = v.Copy()
	}

	// Field Definitions - DefinedTerm
	result.Definitions = make([]*DefinedTermField, len(a.Definitions))
	for i, v := range a.Definitions {
		result.Definitions[i] = v.Copy()
	}

	return result
}

func (a *BodyOfAgreementFormation) Copy() *BodyOfAgreementFormation {
	result := &BodyOfAgreementFormation{}

	// Field Chapters - Chapter
	result.Chapters = make([]*ChapterField, len(a.Chapters))
	for i, v := range a.Chapters {
		result.Chapters[i] = v.Copy()
	}

	// Field Definitions - DefinedTerm
	result.Definitions = make([]*DefinedTermField, len(a.Definitions))
	for i, v := range a.Definitions {
		result.Definitions[i] = v.Copy()
	}

	// Field Revision - uint
	result.Revision = a.Revision

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *BodyOfAgreementAmendment) Copy() *BodyOfAgreementAmendment {
	result := &BodyOfAgreementAmendment{}

	// Field Revision - uint
	result.Revision = a.Revision

	// Field Amendments - Amendment
	result.Amendments = make([]*AmendmentField, len(a.Amendments))
	for i, v := range a.Amendments {
		result.Amendments[i] = v.Copy()
	}

	// Field RefTxID - bin
	result.RefTxID = make([]byte, len(a.RefTxID))
	copy(result.RefTxID, a.RefTxID)

	return result
}

func (a *InstrumentDefinition) Copy() *InstrumentDefinition {
	result := &InstrumentDefinition{}

	// Field InstrumentPermissions - varbin
	result.InstrumentPermissions = make([]byte, len(a.InstrumentPermissions))
	copy(result.InstrumentPermissions, a.InstrumentPermissions)

	// Field EnforcementOrdersPermitted - bool
	result.EnforcementOrdersPermitted = a.EnforcementOrdersPermitted

	// Field VotingRights - bool
	result.VotingRights = a.VotingRights

	// Field VoteMultiplier - uint
	result.VoteMultiplier = a.VoteMultiplier

	// Field AdministrationProposal - bool
	result.AdministrationProposal = a.AdministrationProposal

	// Field HolderProposal - bool
	result.HolderProposal = a.HolderProposal

	// Field InstrumentModificationGovernance - uint
	result.InstrumentModificationGovernance = a.InstrumentModificationGovernance

	// Field AuthorizedTokenQty - uint
	result.AuthorizedTokenQty = a.AuthorizedTokenQty

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentPayload - varbin
	result.InstrumentPayload = make([]byte, len(a.InstrumentPayload))
	copy(result.InstrumentPayload, a.InstrumentPayload)

	// Field TradeRestrictions - varchar
	result.TradeRestrictions = make([]string, len(a.TradeRestrictions))
	for i, v := range a.TradeRestrictions {
		result.TradeRestrictions[i] = v
	}

	return result
}

func (a *InstrumentCreation) Copy() *InstrumentCreation {
	result := &InstrumentCreation{}

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field InstrumentIndex - uint
	result.InstrumentIndex = a.InstrumentIndex

	// Field InstrumentPermissions - varbin
	result.InstrumentPermissions = make([]byte, len(a.InstrumentPermissions))
	copy(result.InstrumentPermissions, a.InstrumentPermissions)

	// Field EnforcementOrdersPermitted - bool
	result.EnforcementOrdersPermitted = a.EnforcementOrdersPermitted

	// Field VotingRights - bool
	result.VotingRights = a.VotingRights

	// Field VoteMultiplier - uint
	result.VoteMultiplier = a.VoteMultiplier

	// Field AdministrationProposal - bool
	result.AdministrationProposal = a.AdministrationProposal

	// Field HolderProposal - bool
	result.HolderProposal = a.HolderProposal

	// Field InstrumentModificationGovernance - uint
	result.InstrumentModificationGovernance = a.InstrumentModificationGovernance

	// Field AuthorizedTokenQty - uint
	result.AuthorizedTokenQty = a.AuthorizedTokenQty

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentPayload - varbin
	result.InstrumentPayload = make([]byte, len(a.InstrumentPayload))
	copy(result.InstrumentPayload, a.InstrumentPayload)

	// Field InstrumentRevision - uint
	result.InstrumentRevision = a.InstrumentRevision

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field TradeRestrictions - varchar
	result.TradeRestrictions = make([]string, len(a.TradeRestrictions))
	for i, v := range a.TradeRestrictions {
		result.TradeRestrictions[i] = v
	}

	return result
}

func (a *InstrumentModification) Copy() *InstrumentModification {
	result := &InstrumentModification{}

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field InstrumentRevision - uint
	result.InstrumentRevision = a.InstrumentRevision

	// Field Amendments - Amendment
	result.Amendments = make([]*AmendmentField, len(a.Amendments))
	for i, v := range a.Amendments {
		result.Amendments[i] = v.Copy()
	}

	// Field RefTxID - bin
	result.RefTxID = make([]byte, len(a.RefTxID))
	copy(result.RefTxID, a.RefTxID)

	return result
}

func (a *Transfer) Copy() *Transfer {
	result := &Transfer{}

	// Field Instruments - InstrumentTransfer
	result.Instruments = make([]*InstrumentTransferField, len(a.Instruments))
	for i, v := range a.Instruments {
		result.Instruments[i] = v.Copy()
	}

	// Field OfferExpiry - uint
	result.OfferExpiry = a.OfferExpiry

	// Field ExchangeFee - uint
	result.ExchangeFee = a.ExchangeFee

	// Field ExchangeFeeAddress - varbin
	result.ExchangeFeeAddress = make([]byte, len(a.ExchangeFeeAddress))
	copy(result.ExchangeFeeAddress, a.ExchangeFeeAddress)

	return result
}

func (a *Settlement) Copy() *Settlement {
	result := &Settlement{}

	// Field Instruments - InstrumentSettlement
	result.Instruments = make([]*InstrumentSettlementField, len(a.Instruments))
	for i, v := range a.Instruments {
		result.Instruments[i] = v.Copy()
	}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Proposal) Copy() *Proposal {
	result := &Proposal{}

	// Field Type - uint
	result.Type = a.Type

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field VoteSystem - uint
	result.VoteSystem = a.VoteSystem

	// Field ProposedAmendments - Amendment
	result.ProposedAmendments = make([]*AmendmentField, len(a.ProposedAmendments))
	for i, v := range a.ProposedAmendments {
		result.ProposedAmendments[i] = v.Copy()
	}

	// Field VoteOptions - varchar
	result.VoteOptions = a.VoteOptions

	// Field VoteMax - uint
	result.VoteMax = a.VoteMax

	// Field ProposalDescription - varchar
	result.ProposalDescription = a.ProposalDescription

	// Field ProposalDocumentHash - bin
	result.ProposalDocumentHash = make([]byte, len(a.ProposalDocumentHash))
	copy(result.ProposalDocumentHash, a.ProposalDocumentHash)

	// Field VoteCutOffTimestamp - uint
	result.VoteCutOffTimestamp = a.VoteCutOffTimestamp

	return result
}

func (a *Vote) Copy() *Vote {
	result := &Vote{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *BallotCast) Copy() *BallotCast {
	result := &BallotCast{}

	// Field VoteTxId - bin
	result.VoteTxId = make([]byte, len(a.VoteTxId))
	copy(result.VoteTxId, a.VoteTxId)

	// Field Vote - varchar
	result.Vote = a.Vote

	return result
}

func (a *BallotCounted) Copy() *BallotCounted {
	result := &BallotCounted{}

	// Field VoteTxId - bin
	result.VoteTxId = make([]byte, len(a.VoteTxId))
	copy(result.VoteTxId, a.VoteTxId)

	// Field Vote - varchar
	result.Vote = a.Vote

	// Field Quantity - uint
	result.Quantity = a.Quantity

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Result) Copy() *Result {
	result := &Result{}

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field ProposedAmendments - Amendment
	result.ProposedAmendments = make([]*AmendmentField, len(a.ProposedAmendments))
	for i, v := range a.ProposedAmendments {
		result.ProposedAmendments[i] = v.Copy()
	}

	// Field VoteTxId - bin
	result.VoteTxId = make([]byte, len(a.VoteTxId))
	copy(result.VoteTxId, a.VoteTxId)

	// Field OptionTally - uint
	result.OptionTally = make([]uint64, len(a.OptionTally))
	for i, v := range a.OptionTally {
		result.OptionTally[i] = v
	}

	// Field Result - varchar
	result.Result = a.Result

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Order) Copy() *Order {
	result := &Order{}

	// Field ComplianceAction - fixedchar
	result.ComplianceAction = a.ComplianceAction

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field TargetAddresses - TargetAddress
	result.TargetAddresses = make([]*TargetAddressField, len(a.TargetAddresses))
	for i, v := range a.TargetAddresses {
		result.TargetAddresses[i] = v.Copy()
	}

	// Field FreezeTxId - bin
	result.FreezeTxId = make([]byte, len(a.FreezeTxId))
	copy(result.FreezeTxId, a.FreezeTxId)

	// Field FreezePeriod - uint
	result.FreezePeriod = a.FreezePeriod

	// Field DepositAddress - varbin
	result.DepositAddress = make([]byte, len(a.DepositAddress))
	copy(result.DepositAddress, a.DepositAddress)

	// Field AuthorityName - varchar
	result.AuthorityName = a.AuthorityName

	// Field AuthorityPublicKey - varbin
	result.AuthorityPublicKey = make([]byte, len(a.AuthorityPublicKey))
	copy(result.AuthorityPublicKey, a.AuthorityPublicKey)

	// Field SignatureAlgorithm - uint
	result.SignatureAlgorithm = a.SignatureAlgorithm

	// Field OrderSignature - varbin
	result.OrderSignature = make([]byte, len(a.OrderSignature))
	copy(result.OrderSignature, a.OrderSignature)

	// Field BitcoinDispersions - QuantityIndex
	result.BitcoinDispersions = make([]*QuantityIndexField, len(a.BitcoinDispersions))
	for i, v := range a.BitcoinDispersions {
		result.BitcoinDispersions[i] = v.Copy()
	}

	// Field Message - varchar
	result.Message = a.Message

	// Field SupportingEvidenceFormat - uint
	result.SupportingEvidenceFormat = a.SupportingEvidenceFormat

	// Field SupportingEvidence - varbin
	result.SupportingEvidence = make([]byte, len(a.SupportingEvidence))
	copy(result.SupportingEvidence, a.SupportingEvidence)

	// Field ReferenceTransactions - ReferenceTransaction
	result.ReferenceTransactions = make([]*ReferenceTransactionField, len(a.ReferenceTransactions))
	for i, v := range a.ReferenceTransactions {
		result.ReferenceTransactions[i] = v.Copy()
	}

	return result
}

func (a *Freeze) Copy() *Freeze {
	result := &Freeze{}

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field Quantities - QuantityIndex
	result.Quantities = make([]*QuantityIndexField, len(a.Quantities))
	for i, v := range a.Quantities {
		result.Quantities[i] = v.Copy()
	}

	// Field FreezePeriod - uint
	result.FreezePeriod = a.FreezePeriod

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Thaw) Copy() *Thaw {
	result := &Thaw{}

	// Field FreezeTxId - bin
	result.FreezeTxId = make([]byte, len(a.FreezeTxId))
	copy(result.FreezeTxId, a.FreezeTxId)

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Confiscation) Copy() *Confiscation {
	result := &Confiscation{}

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field Quantities - QuantityIndex
	result.Quantities = make([]*QuantityIndexField, len(a.Quantities))
	for i, v := range a.Quantities {
		result.Quantities[i] = v.Copy()
	}

	// Field DepositQty - uint
	result.DepositQty = a.DepositQty

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Reconciliation) Copy() *Reconciliation {
	result := &Reconciliation{}

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field Quantities - QuantityIndex
	result.Quantities = make([]*QuantityIndexField, len(a.Quantities))
	for i, v := range a.Quantities {
		result.Quantities[i] = v.Copy()
	}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *Establishment) Copy() *Establishment {
	result := &Establishment{}

	// Field Message - varchar
	result.Message = a.Message

	return result
}

func (a *Addition) Copy() *Addition {
	result := &Addition{}

	// Field Message - varchar
	result.Message = a.Message

	return result
}

func (a *Alteration) Copy() *Alteration {
	result := &Alteration{}

	// Field EntryTxID - bin
	result.EntryTxID = make([]byte, len(a.EntryTxID))
	copy(result.EntryTxID, a.EntryTxID)

	// Field Message - varchar
	result.Message = a.Message

	return result
}

func (a *Removal) Copy() *Removal {
	result := &Removal{}

	// Field EntryTxID - bin
	result.EntryTxID = make([]byte, len(a.EntryTxID))
	copy(result.EntryTxID, a.EntryTxID)

	// Field Message - varchar
	result.Message = a.Message

	return result
}

func (a *Message) Copy() *Message {
	result := &Message{}

	// Field SenderIndexes - uint
	result.SenderIndexes = make([]uint32, len(a.SenderIndexes))
	for i, v := range a.SenderIndexes {
		result.SenderIndexes[i] = v
	}

	// Field ReceiverIndexes - uint
	result.ReceiverIndexes = make([]uint32, len(a.ReceiverIndexes))
	for i, v := range a.ReceiverIndexes {
		result.ReceiverIndexes[i] = v
	}

	// Field MessageCode - uint
	result.MessageCode = a.MessageCode

	// Field MessagePayload - varbin
	result.MessagePayload = make([]byte, len(a.MessagePayload))
	copy(result.MessagePayload, a.MessagePayload)

	return result
}

func (a *Rejection) Copy() *Rejection {
	result := &Rejection{}

	// Field AddressIndexes - uint
	result.AddressIndexes = make([]uint32, len(a.AddressIndexes))
	for i, v := range a.AddressIndexes {
		result.AddressIndexes[i] = v
	}

	// Field RejectAddressIndex - uint
	result.RejectAddressIndex = a.RejectAddressIndex

	// Field RejectionCode - uint
	result.RejectionCode = a.RejectionCode

	// Field Message - varchar
	result.Message = a.Message

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *AdministratorField) Copy() *AdministratorField {
	result := &AdministratorField{}

	// Field Type - uint
	result.Type = a.Type

	// Field Name - varchar
	result.Name = a.Name

	return result
}

func (a *AdminIdentityCertificateField) Copy() *AdminIdentityCertificateField {
	result := &AdminIdentityCertificateField{}

	// Field EntityContract - varbin
	result.EntityContract = make([]byte, len(a.EntityContract))
	copy(result.EntityContract, a.EntityContract)

	// Field Signature - varbin
	result.Signature = make([]byte, len(a.Signature))
	copy(result.Signature, a.Signature)

	// Field BlockHeight - uint
	result.BlockHeight = a.BlockHeight

	// Field Expiration - uint
	result.Expiration = a.Expiration

	return result
}

func (a *AmendmentField) Copy() *AmendmentField {
	result := &AmendmentField{}

	// Field FieldIndexPath - varbin
	result.FieldIndexPath = make([]byte, len(a.FieldIndexPath))
	copy(result.FieldIndexPath, a.FieldIndexPath)

	// Field Operation - uint
	result.Operation = a.Operation

	// Field Data - varbin
	result.Data = make([]byte, len(a.Data))
	copy(result.Data, a.Data)

	return result
}

func (a *InstrumentReceiverField) Copy() *InstrumentReceiverField {
	result := &InstrumentReceiverField{}

	// Field Address - varbin
	result.Address = make([]byte, len(a.Address))
	copy(result.Address, a.Address)

	// Field Quantity - uint
	result.Quantity = a.Quantity

	// Field OracleSigAlgorithm - uint
	result.OracleSigAlgorithm = a.OracleSigAlgorithm

	// Field OracleIndex - uint
	result.OracleIndex = a.OracleIndex

	// Field OracleConfirmationSig - varbin
	result.OracleConfirmationSig = make([]byte, len(a.OracleConfirmationSig))
	copy(result.OracleConfirmationSig, a.OracleConfirmationSig)

	// Field OracleSigBlockHeight - uint
	result.OracleSigBlockHeight = a.OracleSigBlockHeight

	// Field OracleSigExpiry - uint
	result.OracleSigExpiry = a.OracleSigExpiry

	return result
}

func (a *InstrumentSettlementField) Copy() *InstrumentSettlementField {
	result := &InstrumentSettlementField{}

	// Field ContractIndex - uint
	result.ContractIndex = a.ContractIndex

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field Settlements - QuantityIndex
	result.Settlements = make([]*QuantityIndexField, len(a.Settlements))
	for i, v := range a.Settlements {
		result.Settlements[i] = v.Copy()
	}

	return result
}

func (a *InstrumentTransferField) Copy() *InstrumentTransferField {
	result := &InstrumentTransferField{}

	// Field ContractIndex - uint
	result.ContractIndex = a.ContractIndex

	// Field InstrumentType - fixedchar
	result.InstrumentType = a.InstrumentType

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field InstrumentSenders - QuantityIndex
	result.InstrumentSenders = make([]*QuantityIndexField, len(a.InstrumentSenders))
	for i, v := range a.InstrumentSenders {
		result.InstrumentSenders[i] = v.Copy()
	}

	// Field InstrumentReceivers - InstrumentReceiver
	result.InstrumentReceivers = make([]*InstrumentReceiverField, len(a.InstrumentReceivers))
	for i, v := range a.InstrumentReceivers {
		result.InstrumentReceivers[i] = v.Copy()
	}

	return result
}

func (a *ChapterField) Copy() *ChapterField {
	result := &ChapterField{}

	// Field Title - varchar
	result.Title = a.Title

	// Field Preamble - varchar
	result.Preamble = a.Preamble

	// Field Articles - Clause
	result.Articles = make([]*ClauseField, len(a.Articles))
	for i, v := range a.Articles {
		result.Articles[i] = v.Copy()
	}

	return result
}

func (a *ClauseField) Copy() *ClauseField {
	result := &ClauseField{}

	// Field Title - varchar
	result.Title = a.Title

	// Field Body - varchar
	result.Body = a.Body

	// Field Children - Clause
	result.Children = make([]*ClauseField, len(a.Children))
	for i, v := range a.Children {
		result.Children[i] = v.Copy()
	}

	return result
}

func (a *DefinedTermField) Copy() *DefinedTermField {
	result := &DefinedTermField{}

	// Field Term - varchar
	result.Term = a.Term

	// Field Definition - varchar
	result.Definition = a.Definition

	return result
}

func (a *DocumentField) Copy() *DocumentField {
	result := &DocumentField{}

	// Field Name - varchar
	result.Name = a.Name

	// Field Type - varchar
	result.Type = a.Type

	// Field Contents - varbin
	result.Contents = make([]byte, len(a.Contents))
	copy(result.Contents, a.Contents)

	return result
}

func (a *EntityField) Copy() *EntityField {
	result := &EntityField{}

	// Field Name - varchar
	result.Name = a.Name

	// Field Type - fixedchar
	result.Type = a.Type

	// Field LEI - fixedchar
	result.LEI = a.LEI

	// Field UnitNumber - varchar
	result.UnitNumber = a.UnitNumber

	// Field BuildingNumber - varchar
	result.BuildingNumber = a.BuildingNumber

	// Field Street - varchar
	result.Street = a.Street

	// Field SuburbCity - varchar
	result.SuburbCity = a.SuburbCity

	// Field TerritoryStateProvinceCode - fixedchar
	result.TerritoryStateProvinceCode = a.TerritoryStateProvinceCode

	// Field CountryCode - fixedchar
	result.CountryCode = a.CountryCode

	// Field PostalZIPCode - fixedchar
	result.PostalZIPCode = a.PostalZIPCode

	// Field EmailAddress - varchar
	result.EmailAddress = a.EmailAddress

	// Field PhoneNumber - varchar
	result.PhoneNumber = a.PhoneNumber

	// Field Administration - Administrator
	result.Administration = make([]*AdministratorField, len(a.Administration))
	for i, v := range a.Administration {
		result.Administration[i] = v.Copy()
	}

	// Field Management - Manager
	result.Management = make([]*ManagerField, len(a.Management))
	for i, v := range a.Management {
		result.Management[i] = v.Copy()
	}

	// Field DomainName - varchar
	result.DomainName = a.DomainName

	// Field PaymailHandle - varchar
	result.PaymailHandle = a.PaymailHandle

	return result
}

func (a *ManagerField) Copy() *ManagerField {
	result := &ManagerField{}

	// Field Type - uint
	result.Type = a.Type

	// Field Name - varchar
	result.Name = a.Name

	return result
}

func (a *OracleField) Copy() *OracleField {
	result := &OracleField{}

	// Field OracleTypes - uint
	result.OracleTypes = make([]uint32, len(a.OracleTypes))
	for i, v := range a.OracleTypes {
		result.OracleTypes[i] = v
	}

	// Field EntityContract - varbin
	result.EntityContract = make([]byte, len(a.EntityContract))
	copy(result.EntityContract, a.EntityContract)

	return result
}

func (a *QuantityIndexField) Copy() *QuantityIndexField {
	result := &QuantityIndexField{}

	// Field Index - uint
	result.Index = a.Index

	// Field Quantity - uint
	result.Quantity = a.Quantity

	return result
}

func (a *ReferenceTransactionField) Copy() *ReferenceTransactionField {
	result := &ReferenceTransactionField{}

	// Field Transaction - varbin
	result.Transaction = make([]byte, len(a.Transaction))
	copy(result.Transaction, a.Transaction)

	// Field Outputs - varbin
	result.Outputs = make([][]byte, len(a.Outputs))
	for i, v := range a.Outputs {
		result.Outputs[i] = make([]byte, len(a.Outputs[i]))
		copy(result.Outputs[i], v)
	}

	return result
}

func (a *ServiceField) Copy() *ServiceField {
	result := &ServiceField{}

	// Field Type - uint
	result.Type = a.Type

	// Field URL - varchar
	result.URL = a.URL

	// Field PublicKey - bin
	result.PublicKey = make([]byte, len(a.PublicKey))
	copy(result.PublicKey, a.PublicKey)

	return result
}

func (a *TargetAddressField) Copy() *TargetAddressField {
	result := &TargetAddressField{}

	// Field Address - varbin
	result.Address = make([]byte, len(a.Address))
	copy(result.Address, a.Address)

	// Field Quantity - uint
	result.Quantity = a.Quantity

	return result
}

func (a *VotingSystemField) Copy() *VotingSystemField {
	result := &VotingSystemField{}

	// Field Name - varchar
	result.Name = a.Name

	// Field VoteType - fixedchar
	result.VoteType = a.VoteType

	// Field TallyLogic - uint
	result.TallyLogic = a.TallyLogic

	// Field ThresholdPercentage - uint
	result.ThresholdPercentage = a.ThresholdPercentage

	// Field VoteMultiplierPermitted - bool
	result.VoteMultiplierPermitted = a.VoteMultiplierPermitted

	// Field HolderProposalFee - uint
	result.HolderProposalFee = a.HolderProposalFee

	return result
}
