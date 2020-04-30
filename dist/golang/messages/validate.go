package messages

import (
	"fmt"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295
)

func (a *PublicMessage) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint
	if err := a.Regarding.Validate(); err != nil {
		return fmt.Errorf("Regarding invalid : %s", err)
	}

	// Field PublicMessage - Document
	if err := a.PublicMessage.Validate(); err != nil {
		return fmt.Errorf("PublicMessage invalid : %s", err)
	}

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Attachments[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *PrivateMessage) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint
	if err := a.Regarding.Validate(); err != nil {
		return fmt.Errorf("Regarding invalid : %s", err)
	}

	// Field PrivateMessage - Document
	if err := a.PrivateMessage.Validate(); err != nil {
		return fmt.Errorf("PrivateMessage invalid : %s", err)
	}

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Attachments[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *RevertedTx) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Transaction - varbin
	if len(a.Transaction) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Transaction), max4ByteInteger)
	}

	return nil
}

func (a *Offer) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SignatureRequest) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SettlementRequest) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field TransferTxId - bin
	if len(a.TransferTxId) != 0 && len(a.TransferTxId) != 32 {
		return fmt.Errorf("Fixed width field TransferTxId wrong size : %d should be %d",
			len(a.TransferTxId), 32)
	}

	// Field ContractFees - TargetAddress
	if len(a.ContractFees) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ContractFees), max2ByteInteger)
	}
	for i, v := range a.ContractFees {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("ContractFees[%d] invalid : %s", i, err)
		}
	}

	// Field Settlement - varbin
	if len(a.Settlement) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Settlement), max4ByteInteger)
	}

	return nil
}

func (a *OutputMetadata) Validate() error {
	if a == nil {
		return nil
	}

	// Field OutputDescription - varchar
	if len(a.OutputDescription) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OutputDescription), max4ByteInteger)
	}

	// Field Tags - uint
	if len(a.Tags) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Tags), max1ByteInteger)
	}
	for i, v := range a.Tags {
		if TagsData(v) == nil {
			return fmt.Errorf("Tags resource value not defined : %v", v)
		}
		if v > uint32(max1ByteInteger) {
			return fmt.Errorf("Tags[%d] uint over max value : %d > %d", i, v, max1ByteInteger)
		}
	}

	// Field CustomTags - OutputTag
	if len(a.CustomTags) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.CustomTags), max1ByteInteger)
	}
	for i, v := range a.CustomTags {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("CustomTags[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *InitiateRelationship) Validate() error {
	if a == nil {
		return nil
	}

	// Field Type - uint
	foundType := false
	for _, v := range []uint32{0, 1} {
		if a.Type == v {
			foundType = true
			break
		}
	}
	if !foundType {
		return fmt.Errorf("Type value not within options [0 1] : %d", a.Type)
	}

	// Field Seed - varbin
	if len(a.Seed) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Seed), max1ByteInteger)
	}

	// Field Flag - varbin
	if len(a.Flag) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Flag), max1ByteInteger)
	}

	// Field EncryptionType - uint
	foundEncryptionType := false
	for _, v := range []uint32{0, 1} {
		if a.EncryptionType == v {
			foundEncryptionType = true
			break
		}
	}
	if !foundEncryptionType {
		return fmt.Errorf("EncryptionType value not within options [0 1] : %d", a.EncryptionType)
	}

	// Field ProofOfIdentityType - uint
	foundProofOfIdentityType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.ProofOfIdentityType == v {
			foundProofOfIdentityType = true
			break
		}
	}
	if !foundProofOfIdentityType {
		return fmt.Errorf("ProofOfIdentityType value not within options [0 1 2] : %d", a.ProofOfIdentityType)
	}

	// Field ProofOfIdentity - varbin
	if len(a.ProofOfIdentity) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	// Field ChannelParties - ChannelParty
	if len(a.ChannelParties) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ChannelParties), max1ByteInteger)
	}
	for i, v := range a.ChannelParties {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("ChannelParties[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *PendingAcceptRelationship) Validate() error {
	if a == nil {
		return nil
	}

	// Field ProofOfIdentityType - uint
	foundProofOfIdentityType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.ProofOfIdentityType == v {
			foundProofOfIdentityType = true
			break
		}
	}
	if !foundProofOfIdentityType {
		return fmt.Errorf("ProofOfIdentityType value not within options [0 1 2] : %d", a.ProofOfIdentityType)
	}

	// Field ProofOfIdentity - varbin
	if len(a.ProofOfIdentity) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	return nil
}

func (a *AcceptRelationship) Validate() error {
	if a == nil {
		return nil
	}

	// Field ProofOfIdentityType - uint
	foundProofOfIdentityType := false
	for _, v := range []uint32{0, 1, 2} {
		if a.ProofOfIdentityType == v {
			foundProofOfIdentityType = true
			break
		}
	}
	if !foundProofOfIdentityType {
		return fmt.Errorf("ProofOfIdentityType value not within options [0 1 2] : %d", a.ProofOfIdentityType)
	}

	// Field ProofOfIdentity - varbin
	if len(a.ProofOfIdentity) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	return nil
}

func (a *RelationshipAmendment) Validate() error {
	if a == nil {
		return nil
	}

	// Field Seed - varbin
	if len(a.Seed) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Seed), max1ByteInteger)
	}

	// Field BaseEncryptionSecret - varbin
	if len(a.BaseEncryptionSecret) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.BaseEncryptionSecret), max1ByteInteger)
	}

	// Field AddMemberIndexes - uint

	// Field DropMemberIndexes - uint

	return nil
}

func (a *InitiateThread) Validate() error {
	if a == nil {
		return nil
	}

	// Field Flag - varbin
	if len(a.Flag) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Flag), max1ByteInteger)
	}

	// Field Seed - varbin
	if len(a.Seed) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Seed), max1ByteInteger)
	}

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

func (a *ChannelPartyField) Validate() error {
	if a == nil {
		return nil
	}

	// Field AdministrativeAddress - varbin
	if len(a.AdministrativeAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.AdministrativeAddress), max2ByteInteger)
	}

	// Field OutputIndex - uint

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

	// Field DomainName - varchar
	if len(a.DomainName) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.DomainName), max1ByteInteger)
	}

	// Field EntityContractAddress - varbin
	if len(a.EntityContractAddress) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.EntityContractAddress), max2ByteInteger)
	}

	return nil
}

func (a *IdentityOracleProofField) Validate() error {
	if a == nil {
		return nil
	}

	// Field UserID - varbin
	if len(a.UserID) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.UserID), max1ByteInteger)
	}

	// Field Entity - Entity
	if err := a.Entity.Validate(); err != nil {
		return fmt.Errorf("Entity invalid : %s", err)
	}

	// Field OracleSignature - OracleSignature
	if err := a.OracleSignature.Validate(); err != nil {
		return fmt.Errorf("OracleSignature invalid : %s", err)
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

func (a *OracleSignatureField) Validate() error {
	if a == nil {
		return nil
	}

	// Field OracleURL - varchar
	if len(a.OracleURL) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OracleURL), max1ByteInteger)
	}

	// Field BlockHeight - uint

	// Field ValidityPeriod - Period
	if err := a.ValidityPeriod.Validate(); err != nil {
		return fmt.Errorf("ValidityPeriod invalid : %s", err)
	}

	// Field SignatureAlgorithm - uint
	if a.SignatureAlgorithm > uint32(max1ByteInteger) {
		return fmt.Errorf("uint over max value : %d > %d", a.SignatureAlgorithm, max1ByteInteger)
	}

	// Field Signature - varbin
	if len(a.Signature) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Signature), max1ByteInteger)
	}

	return nil
}

func (a *OutpointField) Validate() error {
	if a == nil {
		return nil
	}

	// Field TxId - bin
	if len(a.TxId) != 0 && len(a.TxId) != 32 {
		return fmt.Errorf("Fixed width field TxId wrong size : %d should be %d",
			len(a.TxId), 32)
	}

	// Field OutputIndex - uint

	return nil
}

func (a *OutputTagField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Tag - varchar
	if len(a.Tag) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Tag), max1ByteInteger)
	}

	return nil
}

func (a *PaymailProofField) Validate() error {
	if a == nil {
		return nil
	}

	// Field UserID - varbin
	if len(a.UserID) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.UserID), max1ByteInteger)
	}

	// Field Handle - varchar
	if len(a.Handle) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Handle), max1ByteInteger)
	}

	// Field OracleSignature - OracleSignature
	if err := a.OracleSignature.Validate(); err != nil {
		return fmt.Errorf("OracleSignature invalid : %s", err)
	}

	return nil
}

func (a *PeriodField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Begin - uint

	// Field End - uint

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
