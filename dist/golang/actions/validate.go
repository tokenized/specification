package actions

import (
	"fmt"
	"regexp"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/pkg/errors"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295

	maxArticleDepth = 4
)

func (a *ContractOffer) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
	validValueFoundBodyOfAgreement := false
	for _, v := range []uint32{1} {
		if a.BodyOfAgreementType == v {
			validValueFoundBodyOfAgreement = true
			break
		}
	}
	if !validValueFoundBodyOfAgreement && len(a.BodyOfAgreement) != 0 {
		return fmt.Errorf("BodyOfAgreement is only allowed when BodyOfAgreementType value is within values [1] : %v", a.BodyOfAgreementType)
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

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("ContractURI over max size : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity
	if err := a.Issuer.Validate(); err != nil {
		return errors.Wrap(err, "Issuer")
	}
	validValueFoundIssuer := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundIssuer = true
			break
		}
	}
	if !validValueFoundIssuer && a.Issuer != nil {
		return fmt.Errorf("Issuer is only allowed when ContractType value is within values [0] : %v", a.ContractType)
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

	// Field RestrictedQtyInstruments - uint

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
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}
	validValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			validValueFoundEntityContract = true
			break
		}
	}
	if !validValueFoundEntityContract && len(a.EntityContract) != 0 {
		return fmt.Errorf("EntityContract is only allowed when ContractType value is within values [1] : %v", a.ContractType)
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
	validValueFoundOperatorEntityContract := false
	for _, v := range []bool{true} {
		if a.ContractOperatorIncluded == v {
			validValueFoundOperatorEntityContract = true
			break
		}
	}
	if !validValueFoundOperatorEntityContract && len(a.OperatorEntityContract) != 0 {
		return fmt.Errorf("OperatorEntityContract is only allowed when ContractOperatorIncluded value is within values [true] : %v", a.ContractOperatorIncluded)
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
	if len(a.Services) > max1ByteInteger {
		return fmt.Errorf("Services list over max length : %d > %d", len(a.Services), max1ByteInteger)
	}
	for i, v := range a.Services {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Services[%d]", i))
		}
	}
	validValueFoundServices := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundServices = true
			break
		}
	}
	if !validValueFoundServices && len(a.Services) != 0 {
		return fmt.Errorf("Services is only allowed when ContractType value is within values [0] : %v", a.ContractType)
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

	// Field GoverningLaw - varchar  (LegalSystems Resource)
	if len(a.GoverningLaw) > max1ByteInteger {
		return fmt.Errorf("GoverningLaw over max size : %d > %d", len(a.GoverningLaw), max1ByteInteger)
	}

	// Field Jurisdiction - varchar  (Polities Resource)
	if len(a.Jurisdiction) > max1ByteInteger {
		return fmt.Errorf("Jurisdiction over max size : %d > %d", len(a.Jurisdiction), max1ByteInteger)
	}

	return nil
}

func (a *ContractFormation) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
	validValueFoundBodyOfAgreement := false
	for _, v := range []uint32{1} {
		if a.BodyOfAgreementType == v {
			validValueFoundBodyOfAgreement = true
			break
		}
	}
	if !validValueFoundBodyOfAgreement && len(a.BodyOfAgreement) != 0 {
		return fmt.Errorf("BodyOfAgreement is only allowed when BodyOfAgreementType value is within values [1] : %v", a.BodyOfAgreementType)
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

	// Field ContractExpiration - uint

	// Field ContractURI - varchar
	if len(a.ContractURI) > max1ByteInteger {
		return fmt.Errorf("ContractURI over max size : %d > %d", len(a.ContractURI), max1ByteInteger)
	}

	// Field Issuer - Entity
	if err := a.Issuer.Validate(); err != nil {
		return errors.Wrap(err, "Issuer")
	}
	validValueFoundIssuer := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundIssuer = true
			break
		}
	}
	if !validValueFoundIssuer && a.Issuer != nil {
		return fmt.Errorf("Issuer is only allowed when ContractType value is within values [0] : %v", a.ContractType)
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

	// Field RestrictedQtyInstruments - uint

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
	if len(a.EntityContract) > max2ByteInteger {
		return fmt.Errorf("EntityContract over max size : %d > %d", len(a.EntityContract), max2ByteInteger)
	}
	validValueFoundEntityContract := false
	for _, v := range []uint32{1} {
		if a.ContractType == v {
			validValueFoundEntityContract = true
			break
		}
	}
	if !validValueFoundEntityContract && len(a.EntityContract) != 0 {
		return fmt.Errorf("EntityContract is only allowed when ContractType value is within values [1] : %v", a.ContractType)
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
	if len(a.Services) > max1ByteInteger {
		return fmt.Errorf("Services list over max length : %d > %d", len(a.Services), max1ByteInteger)
	}
	for i, v := range a.Services {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Services[%d]", i))
		}
	}
	validValueFoundServices := false
	for _, v := range []uint32{0} {
		if a.ContractType == v {
			validValueFoundServices = true
			break
		}
	}
	if !validValueFoundServices && len(a.Services) != 0 {
		return fmt.Errorf("Services is only allowed when ContractType value is within values [0] : %v", a.ContractType)
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

	// Field GoverningLaw - varchar  (LegalSystems Resource)
	if len(a.GoverningLaw) > max1ByteInteger {
		return fmt.Errorf("GoverningLaw over max size : %d > %d", len(a.GoverningLaw), max1ByteInteger)
	}

	// Field Jurisdiction - varchar  (Polities Resource)
	if len(a.Jurisdiction) > max1ByteInteger {
		return fmt.Errorf("Jurisdiction over max size : %d > %d", len(a.Jurisdiction), max1ByteInteger)
	}

	// Field RequestPeerChannel - varchar
	if len(a.RequestPeerChannel) > max1ByteInteger {
		return fmt.Errorf("RequestPeerChannel over max size : %d > %d", len(a.RequestPeerChannel), max1ByteInteger)
	}

	return nil
}

func (a *ContractAmendment) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return errors.New("Empty")
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

	// Field GoverningLaw - varchar  (LegalSystems Resource)
	if len(a.GoverningLaw) > max1ByteInteger {
		return fmt.Errorf("GoverningLaw over max size : %d > %d", len(a.GoverningLaw), max1ByteInteger)
	}

	// Field Jurisdiction - varchar  (Polities Resource)
	if len(a.Jurisdiction) > max1ByteInteger {
		return fmt.Errorf("Jurisdiction over max size : %d > %d", len(a.Jurisdiction), max1ByteInteger)
	}

	return nil
}

func (a *ContractAddressChange) Validate() error {
	if a == nil {
		return errors.New("Empty")
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

func (a *BodyOfAgreementOffer) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}
	// Check depth or articles. Articles, Sections, Subsections, Paragraphs, Subparagraphs.
	terms := make(map[string]int)
	for _, chapter := range a.Chapters {
		terms = chapter.Terms(terms)
	}

	if err := checkTerms(a.Definitions, terms); err != nil {
		return err
	}

	// Field Chapters - Chapter
	if len(a.Chapters) > max1ByteInteger {
		return fmt.Errorf("Chapters list over max length : %d > %d", len(a.Chapters), max1ByteInteger)
	}
	for i, v := range a.Chapters {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Chapters[%d]", i))
		}
	}

	// Field Definitions - DefinedTerm
	if len(a.Definitions) > max1ByteInteger {
		return fmt.Errorf("Definitions list over max length : %d > %d", len(a.Definitions), max1ByteInteger)
	}
	for i, v := range a.Definitions {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Definitions[%d]", i))
		}
	}

	return nil
}

func (a *BodyOfAgreementFormation) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}
	// Check depth or articles. Articles, Sections, Subsections, Paragraphs, Subparagraphs.
	terms := make(map[string]int)
	for _, chapter := range a.Chapters {
		terms = chapter.Terms(terms)
	}

	if err := checkTerms(a.Definitions, terms); err != nil {
		return err
	}

	// Field Chapters - Chapter
	if len(a.Chapters) > max1ByteInteger {
		return fmt.Errorf("Chapters list over max length : %d > %d", len(a.Chapters), max1ByteInteger)
	}
	for i, v := range a.Chapters {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Chapters[%d]", i))
		}
	}

	// Field Definitions - DefinedTerm
	if len(a.Definitions) > max1ByteInteger {
		return fmt.Errorf("Definitions list over max length : %d > %d", len(a.Definitions), max1ByteInteger)
	}
	for i, v := range a.Definitions {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Definitions[%d]", i))
		}
	}

	// Field Revision - uint

	// Field Timestamp - uint

	return nil
}

func (a *BodyOfAgreementAmendment) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Revision - uint

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

func (a *InstrumentDefinition) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field InstrumentPermissions - varbin
	if len(a.InstrumentPermissions) > max2ByteInteger {
		return fmt.Errorf("InstrumentPermissions over max size : %d > %d", len(a.InstrumentPermissions), max2ByteInteger)
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteMultiplier over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
	}

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field InstrumentModificationGovernance - uint
	foundInstrumentModificationGovernance := false
	for _, v := range []uint32{0, 1} {
		if a.InstrumentModificationGovernance == v {
			foundInstrumentModificationGovernance = true
			break
		}
	}
	if !foundInstrumentModificationGovernance {
		return fmt.Errorf("InstrumentModificationGovernance value not within options [0 1] : %d", a.InstrumentModificationGovernance)
	}

	// Field AuthorizedTokenQty - uint

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}
	if len(a.InstrumentType) == 0 {
		return fmt.Errorf("InstrumentType required")
	}

	// Field InstrumentPayload - varbin
	if len(a.InstrumentPayload) > max2ByteInteger {
		return fmt.Errorf("InstrumentPayload over max size : %d > %d", len(a.InstrumentPayload), max2ByteInteger)
	}
	if len(a.InstrumentPayload) == 0 {
		return fmt.Errorf("InstrumentPayload required")
	}

	// Field TradeRestrictions - varchar  (Polities Resource)
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("TradeRestrictions list over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if len(v) > max1ByteInteger {
			return fmt.Errorf("[%d] TradeRestrictions size over max value : %d > %d", i, len(v), max1ByteInteger)
		}
	}

	// Field TransferFee - Fee
	if err := a.TransferFee.Validate(); err != nil {
		return errors.Wrap(err, "TransferFee")
	}

	return nil
}

func (a *InstrumentCreation) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
	}

	// Field InstrumentIndex - uint

	// Field InstrumentPermissions - varbin
	if len(a.InstrumentPermissions) > max2ByteInteger {
		return fmt.Errorf("InstrumentPermissions over max size : %d > %d", len(a.InstrumentPermissions), max2ByteInteger)
	}

	// Field EnforcementOrdersPermitted - bool

	// Field VotingRights - bool

	// Field VoteMultiplier - uint
	if a.VoteMultiplier > uint32(max1ByteInteger) {
		return fmt.Errorf("VoteMultiplier over max value : %d > %d", a.VoteMultiplier, max1ByteInteger)
	}

	// Field AdministrationProposal - bool

	// Field HolderProposal - bool

	// Field InstrumentModificationGovernance - uint
	foundInstrumentModificationGovernance := false
	for _, v := range []uint32{0, 1} {
		if a.InstrumentModificationGovernance == v {
			foundInstrumentModificationGovernance = true
			break
		}
	}
	if !foundInstrumentModificationGovernance {
		return fmt.Errorf("InstrumentModificationGovernance value not within options [0 1] : %d", a.InstrumentModificationGovernance)
	}

	// Field AuthorizedTokenQty - uint

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentPayload - varbin
	if len(a.InstrumentPayload) > max2ByteInteger {
		return fmt.Errorf("InstrumentPayload over max size : %d > %d", len(a.InstrumentPayload), max2ByteInteger)
	}

	// Field InstrumentRevision - uint

	// Field Timestamp - uint

	// Field TradeRestrictions - varchar  (Polities Resource)
	if len(a.TradeRestrictions) > max2ByteInteger {
		return fmt.Errorf("TradeRestrictions list over max length : %d > %d", len(a.TradeRestrictions), max2ByteInteger)
	}
	for i, v := range a.TradeRestrictions {
		if len(v) > max1ByteInteger {
			return fmt.Errorf("[%d] TradeRestrictions size over max value : %d > %d", i, len(v), max1ByteInteger)
		}
	}

	// Field TransferFee - Fee
	if err := a.TransferFee.Validate(); err != nil {
		return errors.Wrap(err, "TransferFee")
	}

	return nil
}

func (a *InstrumentModification) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
	}

	// Field InstrumentRevision - uint

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
		return errors.New("Empty")
	}

	// Field Instruments - InstrumentTransfer
	if len(a.Instruments) > max1ByteInteger {
		return fmt.Errorf("Instruments list over max length : %d > %d", len(a.Instruments), max1ByteInteger)
	}
	for i, v := range a.Instruments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Instruments[%d]", i))
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
		return errors.New("Empty")
	}

	// Field Instruments - InstrumentSettlement
	if len(a.Instruments) > max1ByteInteger {
		return fmt.Errorf("Instruments list over max length : %d > %d", len(a.Instruments), max1ByteInteger)
	}
	for i, v := range a.Instruments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Instruments[%d]", i))
		}
	}

	// Field Timestamp - uint

	return nil
}

func (a *RectificationSettlement) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Transfer - ReferenceTransaction
	if err := a.Transfer.Validate(); err != nil {
		return errors.Wrap(err, "Transfer")
	}

	// Field Instruments - InstrumentSettlement
	if len(a.Instruments) > max1ByteInteger {
		return fmt.Errorf("Instruments list over max length : %d > %d", len(a.Instruments), max1ByteInteger)
	}
	for i, v := range a.Instruments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Instruments[%d]", i))
		}
	}

	// Field Timestamp - uint

	return nil
}

func (a *Proposal) Validate() error {
	if a == nil {
		return errors.New("Empty")
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

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	return nil
}

func (a *BallotCast) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return errors.New("Empty")
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
		return errors.New("Empty")
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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
		return errors.New("Empty")
	}

	// Field ComplianceAction - fixedchar
	if len(a.ComplianceAction) != 0 && len(a.ComplianceAction) != 1 {
		return fmt.Errorf("ComplianceAction fixed width field wrong size : %d should be %d",
			len(a.ComplianceAction), 1)
	}
	foundComplianceAction := false
	for _, o := range []string{"F", "T", "C", "R"} {
		if a.ComplianceAction == o {
			foundComplianceAction = true
			break
		}
	}
	if !foundComplianceAction {
		return fmt.Errorf("ComplianceAction value not within options [F T C R] : %s", a.ComplianceAction)
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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
	validValueFoundTargetAddresses := false

	for _, v := range []string{"F", "C", "R"} {
		if a.ComplianceAction == v {
			validValueFoundTargetAddresses = true
			break
		}
	}
	if !validValueFoundTargetAddresses && len(a.TargetAddresses) != 0 {
		return fmt.Errorf("TargetAddresses is only allowed when ComplianceAction value is within values [F C R] : %v", a.ComplianceAction)
	}
	requiredValueFoundTargetAddresses := false

	for _, v := range []string{"F", "C", "R"} {
		if a.ComplianceAction == v {
			requiredValueFoundTargetAddresses = true
			break
		}
	}
	if requiredValueFoundTargetAddresses && len(a.TargetAddresses) == 0 {
		return fmt.Errorf("TargetAddresses is required when ComplianceAction value is within values [F C R] : %v", a.ComplianceAction)
	}

	// Field FreezeTxId - bin
	if len(a.FreezeTxId) != 0 && len(a.FreezeTxId) != 32 {
		return fmt.Errorf("FreezeTxId fixed width field wrong size : %d should be %d",
			len(a.FreezeTxId), 32)
	}
	validValueFoundFreezeTxId := false

	for _, v := range []string{"T"} {
		if a.ComplianceAction == v {
			validValueFoundFreezeTxId = true
			break
		}
	}
	if !validValueFoundFreezeTxId && len(a.FreezeTxId) != 0 {
		return fmt.Errorf("FreezeTxId is only allowed when ComplianceAction value is within values [T] : %v", a.ComplianceAction)
	}

	// Field FreezePeriod - uint
	validValueFoundFreezePeriod := false

	for _, v := range []string{"F"} {
		if a.ComplianceAction == v {
			validValueFoundFreezePeriod = true
			break
		}
	}
	if !validValueFoundFreezePeriod && a.FreezePeriod != 0 {
		return fmt.Errorf("FreezePeriod is only allowed when ComplianceAction value is within values [F] : %v", a.ComplianceAction)
	}

	// Field DepositAddress - varbin
	if len(a.DepositAddress) > 0 {
		if err := AddressIsValid(a.DepositAddress); err != nil {
			return errors.Wrap(err, "DepositAddress")
		}
	}
	if len(a.DepositAddress) > max2ByteInteger {
		return fmt.Errorf("DepositAddress over max size : %d > %d", len(a.DepositAddress), max2ByteInteger)
	}
	validValueFoundDepositAddress := false

	for _, v := range []string{"C"} {
		if a.ComplianceAction == v {
			validValueFoundDepositAddress = true
			break
		}
	}
	if !validValueFoundDepositAddress && len(a.DepositAddress) != 0 {
		return fmt.Errorf("DepositAddress is only allowed when ComplianceAction value is within values [C] : %v", a.ComplianceAction)
	}
	requiredValueFoundDepositAddress := false

	for _, v := range []string{"C"} {
		if a.ComplianceAction == v {
			requiredValueFoundDepositAddress = true
			break
		}
	}
	if requiredValueFoundDepositAddress && len(a.DepositAddress) == 0 {
		return fmt.Errorf("DepositAddress is required when ComplianceAction value is within values [C] : %v", a.ComplianceAction)
	}

	// Field AuthorityName - varchar
	if len(a.AuthorityName) > max1ByteInteger {
		return fmt.Errorf("AuthorityName over max size : %d > %d", len(a.AuthorityName), max1ByteInteger)
	}
	validValueFoundAuthorityName := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			validValueFoundAuthorityName = true
			break
		}
	}
	if !validValueFoundAuthorityName && len(a.AuthorityName) != 0 {
		return fmt.Errorf("AuthorityName is only allowed when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
	}
	requiredValueFoundAuthorityName := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			requiredValueFoundAuthorityName = true
			break
		}
	}
	if requiredValueFoundAuthorityName && len(a.AuthorityName) == 0 {
		return fmt.Errorf("AuthorityName is required when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
	}

	// Field AuthorityPublicKey - varbin
	if len(a.AuthorityPublicKey) > max1ByteInteger {
		return fmt.Errorf("AuthorityPublicKey over max size : %d > %d", len(a.AuthorityPublicKey), max1ByteInteger)
	}
	validValueFoundAuthorityPublicKey := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			validValueFoundAuthorityPublicKey = true
			break
		}
	}
	if !validValueFoundAuthorityPublicKey && len(a.AuthorityPublicKey) != 0 {
		return fmt.Errorf("AuthorityPublicKey is only allowed when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
	}
	requiredValueFoundAuthorityPublicKey := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			requiredValueFoundAuthorityPublicKey = true
			break
		}
	}
	if requiredValueFoundAuthorityPublicKey && len(a.AuthorityPublicKey) == 0 {
		return fmt.Errorf("AuthorityPublicKey is required when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
	}

	// Field SignatureAlgorithm - uint
	foundSignatureAlgorithm := false
	for _, v := range []uint32{0, 1} {
		if a.SignatureAlgorithm == v {
			foundSignatureAlgorithm = true
			break
		}
	}
	if !foundSignatureAlgorithm {
		return fmt.Errorf("SignatureAlgorithm value not within options [0 1] : %d", a.SignatureAlgorithm)
	}

	// Field OrderSignature - varbin
	if len(a.OrderSignature) > max1ByteInteger {
		return fmt.Errorf("OrderSignature over max size : %d > %d", len(a.OrderSignature), max1ByteInteger)
	}
	validValueFoundOrderSignature := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			validValueFoundOrderSignature = true
			break
		}
	}
	if !validValueFoundOrderSignature && len(a.OrderSignature) != 0 {
		return fmt.Errorf("OrderSignature is only allowed when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
	}
	requiredValueFoundOrderSignature := false
	for _, v := range []uint32{1} {
		if a.SignatureAlgorithm == v {
			requiredValueFoundOrderSignature = true
			break
		}
	}
	if requiredValueFoundOrderSignature && len(a.OrderSignature) == 0 {
		return fmt.Errorf("OrderSignature is required when SignatureAlgorithm value is within values [1] : %v", a.SignatureAlgorithm)
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
	validValueFoundBitcoinDispersions := false

	for _, v := range []string{"R"} {
		if a.ComplianceAction == v {
			validValueFoundBitcoinDispersions = true
			break
		}
	}
	if !validValueFoundBitcoinDispersions && len(a.BitcoinDispersions) != 0 {
		return fmt.Errorf("BitcoinDispersions is only allowed when ComplianceAction value is within values [R] : %v", a.ComplianceAction)
	}
	requiredValueFoundBitcoinDispersions := false

	for _, v := range []string{"R"} {
		if a.ComplianceAction == v {
			requiredValueFoundBitcoinDispersions = true
			break
		}
	}
	if requiredValueFoundBitcoinDispersions && len(a.BitcoinDispersions) == 0 {
		return fmt.Errorf("BitcoinDispersions is required when ComplianceAction value is within values [R] : %v", a.ComplianceAction)
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
	validValueFoundSupportingEvidence := false
	for _, v := range []uint32{1} {
		if a.SupportingEvidenceFormat == v {
			validValueFoundSupportingEvidence = true
			break
		}
	}
	if !validValueFoundSupportingEvidence && len(a.SupportingEvidence) != 0 {
		return fmt.Errorf("SupportingEvidence is only allowed when SupportingEvidenceFormat value is within values [1] : %v", a.SupportingEvidenceFormat)
	}
	requiredValueFoundSupportingEvidence := false
	for _, v := range []uint32{1} {
		if a.SupportingEvidenceFormat == v {
			requiredValueFoundSupportingEvidence = true
			break
		}
	}
	if requiredValueFoundSupportingEvidence && len(a.SupportingEvidence) == 0 {
		return fmt.Errorf("SupportingEvidence is required when SupportingEvidenceFormat value is within values [1] : %v", a.SupportingEvidenceFormat)
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
		return errors.New("Empty")
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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
		return errors.New("Empty")
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
		return errors.New("Empty")
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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

func (a *DeprecatedReconciliation) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
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
		return errors.New("Empty")
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Addition) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Message - varchar
	if len(a.Message) > max4ByteInteger {
		return fmt.Errorf("Message over max size : %d > %d", len(a.Message), max4ByteInteger)
	}

	return nil
}

func (a *Alteration) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return errors.New("Empty")
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
		return errors.New("Empty")
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
		return errors.New("Empty")
	}

	// Field AddressIndexes - uint
	if len(a.AddressIndexes) > max4ByteInteger {
		return fmt.Errorf("AddressIndexes list over max length : %d > %d", len(a.AddressIndexes), max4ByteInteger)
	}

	// Field RejectAddressIndex - uint
	if a.RejectAddressIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("RejectAddressIndex over max value : %d > %d", a.RejectAddressIndex, max2ByteInteger)
	}

	// Field RejectionCode - uint  (Rejections Resource)
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

	// Field Type - uint  (Roles Resource)
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

func (a *InstrumentReceiverField) Validate() error {
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

func (a *InstrumentSettlementField) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractIndex - uint
	if a.ContractIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("ContractIndex over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
	}

	// Field Settlements - QuantityIndex
	if len(a.Settlements) > max4ByteInteger {
		return fmt.Errorf("Settlements list over max length : %d > %d", len(a.Settlements), max4ByteInteger)
	}
	for i, v := range a.Settlements {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Settlements[%d]", i))
		}
	}

	return nil
}

func (a *InstrumentTransferField) Validate() error {
	if a == nil {
		return nil
	}

	// Field ContractIndex - uint
	if a.ContractIndex > uint32(max2ByteInteger) {
		return fmt.Errorf("ContractIndex over max value : %d > %d", a.ContractIndex, max2ByteInteger)
	}

	// Field InstrumentType - fixedchar
	if len(a.InstrumentType) != 0 && len(a.InstrumentType) != 3 {
		return fmt.Errorf("InstrumentType fixed width field wrong size : %d should be %d",
			len(a.InstrumentType), 3)
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
	}

	// Field InstrumentSenders - QuantityIndex
	if len(a.InstrumentSenders) > max2ByteInteger {
		return fmt.Errorf("InstrumentSenders list over max length : %d > %d", len(a.InstrumentSenders), max2ByteInteger)
	}
	for i, v := range a.InstrumentSenders {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("InstrumentSenders[%d]", i))
		}
	}

	// Field InstrumentReceivers - InstrumentReceiver
	if len(a.InstrumentReceivers) > max2ByteInteger {
		return fmt.Errorf("InstrumentReceivers list over max length : %d > %d", len(a.InstrumentReceivers), max2ByteInteger)
	}
	for i, v := range a.InstrumentReceivers {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("InstrumentReceivers[%d]", i))
		}
	}

	// Field RefundAddress - varbin
	if len(a.RefundAddress) > 0 {
		if err := AddressIsValid(a.RefundAddress); err != nil {
			return errors.Wrap(err, "RefundAddress")
		}
	}
	if len(a.RefundAddress) > max2ByteInteger {
		return fmt.Errorf("RefundAddress over max size : %d > %d", len(a.RefundAddress), max2ByteInteger)
	}

	return nil
}

func (a *ChapterField) Validate() error {
	if a == nil {
		return nil
	}
	// Check depth or articles. Articles, Sections, Subsections, Paragraphs, Subparagraphs.
	for i, article := range a.Articles {
		if article.Depth() > maxArticleDepth {
			return fmt.Errorf("Article %d over max depth : %d > %d", i, article.Depth(),
				maxArticleDepth)
		}
	}

	// Field Title - varchar
	if len(a.Title) > max1ByteInteger {
		return fmt.Errorf("Title over max size : %d > %d", len(a.Title), max1ByteInteger)
	}

	// Field Preamble - varchar
	if len(a.Preamble) > max2ByteInteger {
		return fmt.Errorf("Preamble over max size : %d > %d", len(a.Preamble), max2ByteInteger)
	}

	// Field Articles - Clause
	if len(a.Articles) > max1ByteInteger {
		return fmt.Errorf("Articles list over max length : %d > %d", len(a.Articles), max1ByteInteger)
	}
	for i, v := range a.Articles {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Articles[%d]", i))
		}
	}

	return nil
}

func (a *ClauseField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Title - varchar
	if len(a.Title) > max1ByteInteger {
		return fmt.Errorf("Title over max size : %d > %d", len(a.Title), max1ByteInteger)
	}

	// Field Body - varchar
	if len(a.Body) > max2ByteInteger {
		return fmt.Errorf("Body over max size : %d > %d", len(a.Body), max2ByteInteger)
	}

	// Field Children - Clause
	if len(a.Children) > max1ByteInteger {
		return fmt.Errorf("Children list over max length : %d > %d", len(a.Children), max1ByteInteger)
	}
	for i, v := range a.Children {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Children[%d]", i))
		}
	}

	return nil
}

func (a *DefinedTermField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Term - varchar
	if len(a.Term) > max1ByteInteger {
		return fmt.Errorf("Term over max size : %d > %d", len(a.Term), max1ByteInteger)
	}

	// Field Definition - varchar
	if len(a.Definition) > max2ByteInteger {
		return fmt.Errorf("Definition over max size : %d > %d", len(a.Definition), max2ByteInteger)
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

	// Field Type - fixedchar  (Entities Resource)
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

func (a *FeeField) Validate() error {
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
	if len(a.Address) == 0 {
		return fmt.Errorf("Address required")
	}

	// Field Quantity - uint

	return nil
}

func (a *ManagerField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint  (Roles Resource)
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
func (c *ClauseField) Depth() int {
	if len(c.Children) == 0 {
		return 0 // no children
	}

	depth := 0
	for _, child := range c.Children {
		childDepth := child.Depth()
		if childDepth > depth {
			depth = childDepth
		}
	}

	return depth + 1
}

// escape backslashes to get single backslashes in regex
var termRegEx = regexp.MustCompile("(?:\\[)(.+?)(?:\\]\\(\\))")

// findTerms adds any terms contained in the text to the map.
func findTerms(text string, terms map[string]int) map[string]int {
	matches := termRegEx.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		// index 1 for first regex group
		terms[match[1]] = 1
	}
	return terms
}

// Terms adds any terms in the chapter or its articles to the map.
func (c *ChapterField) Terms(terms map[string]int) map[string]int {
	terms = findTerms(c.Preamble, terms)

	for _, article := range c.Articles {
		terms = article.Terms(terms)
	}

	return terms
}

// Terms adds any terms in the clause or its children to the map.
func (c *ClauseField) Terms(terms map[string]int) map[string]int {
	terms = findTerms(c.Body, terms)

	for _, child := range c.Children {
		terms = child.Terms(terms)
	}

	return terms
}

// checkTerms returns an error if any of the defined terms are not found in the referenced terms
// map or if any of the referenced terms are not in the defined terms.
func checkTerms(defined []*DefinedTermField, referenced map[string]int) error {
	// Check that all referenced terms are defined.
	var undefined []string
	for term, _ := range referenced {
		found := false
		for _, defined := range defined {
			if term == defined.Term {
				found = true
				break
			}
		}

		if !found {
			undefined = append(undefined, term)
		}
	}

	if len(undefined) > 0 {
		return fmt.Errorf("Undefined terms : %v", undefined)
	}

	// Check that all defined terms are referenced
	var unreferenced []string
	for _, defined := range defined {
		if _, exists := referenced[defined.Term]; !exists {
			unreferenced = append(unreferenced, defined.Term)
		}
	}

	if len(unreferenced) > 0 {
		return fmt.Errorf("Unreferenced defined terms : %v", unreferenced)
	}

	return nil
}
