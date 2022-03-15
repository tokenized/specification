package messages

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/pkg/errors"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295

	maxArticleDepth = 4
)

func (a *PublicMessage) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("Subject over max size : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint
	if err := a.Regarding.Validate(); err != nil {
		return errors.Wrap(err, "Regarding")
	}

	// Field PublicMessage - Document
	if err := a.PublicMessage.Validate(); err != nil {
		return errors.Wrap(err, "PublicMessage")
	}

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("Attachments list over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Attachments[%d]", i))
		}
	}

	return nil
}

func (a *PrivateMessage) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("Subject over max size : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint
	if err := a.Regarding.Validate(); err != nil {
		return errors.Wrap(err, "Regarding")
	}

	// Field PrivateMessage - Document
	if err := a.PrivateMessage.Validate(); err != nil {
		return errors.Wrap(err, "PrivateMessage")
	}

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("Attachments list over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Attachments[%d]", i))
		}
	}

	return nil
}

func (a *RevertedTx) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field Transaction - varbin
	if len(a.Transaction) > max4ByteInteger {
		return fmt.Errorf("Transaction over max size : %d > %d", len(a.Transaction), max4ByteInteger)
	}

	return nil
}

func (a *Offer) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("Payload over max size : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SignatureRequest) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("Payload over max size : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SettlementRequest) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Timestamp - uint

	// Field TransferTxId - bin
	if len(a.TransferTxId) != 0 && len(a.TransferTxId) != 32 {
		return fmt.Errorf("TransferTxId fixed width field wrong size : %d should be %d",
			len(a.TransferTxId), 32)
	}

	// Field ContractFees - TargetAddress
	if len(a.ContractFees) > max2ByteInteger {
		return fmt.Errorf("ContractFees list over max length : %d > %d", len(a.ContractFees), max2ByteInteger)
	}
	for i, v := range a.ContractFees {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("ContractFees[%d]", i))
		}
	}

	// Field Settlement - varbin
	if len(a.Settlement) > max4ByteInteger {
		return fmt.Errorf("Settlement over max size : %d > %d", len(a.Settlement), max4ByteInteger)
	}

	return nil
}

func (a *OutputMetadata) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field OutputDescription - varchar
	if len(a.OutputDescription) > max4ByteInteger {
		return fmt.Errorf("OutputDescription over max size : %d > %d", len(a.OutputDescription), max4ByteInteger)
	}

	// Field Tags - uint  (Tags Resource)
	if len(a.Tags) > max1ByteInteger {
		return fmt.Errorf("Tags list over max length : %d > %d", len(a.Tags), max1ByteInteger)
	}
	for i, v := range a.Tags {
		if TagsData(v) == nil {
			return fmt.Errorf("Tags[%d] resource Tags value not defined : %v", i, v)
		}
		if v > uint32(max1ByteInteger) {
			return fmt.Errorf("Tags[%d] over max value : %d > %d", i, v, max1ByteInteger)
		}
	}

	// Field CustomTags - OutputTag
	if len(a.CustomTags) > max1ByteInteger {
		return fmt.Errorf("CustomTags list over max length : %d > %d", len(a.CustomTags), max1ByteInteger)
	}
	for i, v := range a.CustomTags {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("CustomTags[%d]", i))
		}
	}

	return nil
}

func (a *Distribution) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field InstrumentCode - bin
	if len(a.InstrumentCode) != 0 && len(a.InstrumentCode) != 20 {
		return fmt.Errorf("InstrumentCode fixed width field wrong size : %d should be %d",
			len(a.InstrumentCode), 20)
	}
	if len(a.InstrumentCode) == 0 {
		return fmt.Errorf("InstrumentCode required")
	}

	// Field Timestamp - uint

	return nil
}

func (a *InitiateRelationship) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return fmt.Errorf("Seed over max size : %d > %d", len(a.Seed), max1ByteInteger)
	}

	// Field Flag - varbin
	if len(a.Flag) > max1ByteInteger {
		return fmt.Errorf("Flag over max size : %d > %d", len(a.Flag), max1ByteInteger)
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
		return fmt.Errorf("ProofOfIdentity over max size : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	// Field ChannelParties - ChannelParty
	if len(a.ChannelParties) > max1ByteInteger {
		return fmt.Errorf("ChannelParties list over max length : %d > %d", len(a.ChannelParties), max1ByteInteger)
	}
	for i, v := range a.ChannelParties {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("ChannelParties[%d]", i))
		}
	}

	return nil
}

func (a *PendingAcceptRelationship) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return fmt.Errorf("ProofOfIdentity over max size : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	return nil
}

func (a *AcceptRelationship) Validate() error {
	if a == nil {
		return errors.New("Empty")
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
		return fmt.Errorf("ProofOfIdentity over max size : %d > %d", len(a.ProofOfIdentity), max1ByteInteger)
	}

	return nil
}

func (a *RelationshipAmendment) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Seed - varbin
	if len(a.Seed) > max1ByteInteger {
		return fmt.Errorf("Seed over max size : %d > %d", len(a.Seed), max1ByteInteger)
	}

	// Field BaseEncryptionSecret - varbin
	if len(a.BaseEncryptionSecret) > max1ByteInteger {
		return fmt.Errorf("BaseEncryptionSecret over max size : %d > %d", len(a.BaseEncryptionSecret), max1ByteInteger)
	}

	// Field AddMemberIndexes - uint

	// Field DropMemberIndexes - uint

	return nil
}

func (a *InitiateThread) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Flag - varbin
	if len(a.Flag) > max1ByteInteger {
		return fmt.Errorf("Flag over max size : %d > %d", len(a.Flag), max1ByteInteger)
	}

	// Field Seed - varbin
	if len(a.Seed) > max1ByteInteger {
		return fmt.Errorf("Seed over max size : %d > %d", len(a.Seed), max1ByteInteger)
	}

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

func (a *ChannelPartyField) Validate() error {
	if a == nil {
		return nil
	}

	// Field AdministrativeAddress - varbin
	if len(a.AdministrativeAddress) > 0 {
		if err := AddressIsValid(a.AdministrativeAddress); err != nil {
			return errors.Wrap(err, "AdministrativeAddress")
		}
	}
	if len(a.AdministrativeAddress) > max2ByteInteger {
		return fmt.Errorf("AdministrativeAddress over max size : %d > %d", len(a.AdministrativeAddress), max2ByteInteger)
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

func (a *IdentityOracleProofField) Validate() error {
	if a == nil {
		return nil
	}

	// Field UserID - varbin
	if len(a.UserID) > max1ByteInteger {
		return fmt.Errorf("UserID over max size : %d > %d", len(a.UserID), max1ByteInteger)
	}

	// Field Entity - Entity
	if err := a.Entity.Validate(); err != nil {
		return errors.Wrap(err, "Entity")
	}

	// Field OracleSignature - OracleSignature
	if err := a.OracleSignature.Validate(); err != nil {
		return errors.Wrap(err, "OracleSignature")
	}

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

func (a *OracleSignatureField) Validate() error {
	if a == nil {
		return nil
	}

	// Field OracleURL - varchar
	if len(a.OracleURL) > max1ByteInteger {
		return fmt.Errorf("OracleURL over max size : %d > %d", len(a.OracleURL), max1ByteInteger)
	}

	// Field BlockHeight - uint

	// Field ValidityPeriod - Period
	if err := a.ValidityPeriod.Validate(); err != nil {
		return errors.Wrap(err, "ValidityPeriod")
	}

	// Field SignatureAlgorithm - uint
	if a.SignatureAlgorithm > uint32(max1ByteInteger) {
		return fmt.Errorf("SignatureAlgorithm over max value : %d > %d", a.SignatureAlgorithm, max1ByteInteger)
	}

	// Field Signature - varbin
	if len(a.Signature) > max1ByteInteger {
		return fmt.Errorf("Signature over max size : %d > %d", len(a.Signature), max1ByteInteger)
	}

	return nil
}

func (a *OutpointField) Validate() error {
	if a == nil {
		return nil
	}

	// Field TxId - bin
	if len(a.TxId) != 0 && len(a.TxId) != 32 {
		return fmt.Errorf("TxId fixed width field wrong size : %d should be %d",
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
		return fmt.Errorf("Tag over max size : %d > %d", len(a.Tag), max1ByteInteger)
	}

	return nil
}

func (a *PaymailProofField) Validate() error {
	if a == nil {
		return nil
	}

	// Field UserID - varbin
	if len(a.UserID) > max1ByteInteger {
		return fmt.Errorf("UserID over max size : %d > %d", len(a.UserID), max1ByteInteger)
	}

	// Field Handle - varchar
	if len(a.Handle) > max1ByteInteger {
		return fmt.Errorf("Handle over max size : %d > %d", len(a.Handle), max1ByteInteger)
	}

	// Field OracleSignature - OracleSignature
	if err := a.OracleSignature.Validate(); err != nil {
		return errors.Wrap(err, "OracleSignature")
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
