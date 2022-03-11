package messages

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

func (l *PublicMessage) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &PublicMessage{}
	}
	cr := right
	if cr == nil {
		cr = &PublicMessage{}
	}
	r, ok := cr.(*PublicMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if c.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint
	if !c.Regarding.Equal(r.Regarding) {
		return false // fmt.Errorf("Regarding : %s", err)
	}

	// Field PublicMessage - Document
	if !c.PublicMessage.Equal(r.PublicMessage) {
		return false // fmt.Errorf("PublicMessage : %s", err)
	}

	// Field Attachments - Document
	if len(c.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *PrivateMessage) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &PrivateMessage{}
	}
	cr := right
	if cr == nil {
		cr = &PrivateMessage{}
	}
	r, ok := cr.(*PrivateMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if c.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint
	if !c.Regarding.Equal(r.Regarding) {
		return false // fmt.Errorf("Regarding : %s", err)
	}

	// Field PrivateMessage - Document
	if !c.PrivateMessage.Equal(r.PrivateMessage) {
		return false // fmt.Errorf("PrivateMessage : %s", err)
	}

	// Field Attachments - Document
	if len(c.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *RevertedTx) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &RevertedTx{}
	}
	cr := right
	if cr == nil {
		cr = &RevertedTx{}
	}
	r, ok := cr.(*RevertedTx)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Transaction - varbin
	if !bytes.Equal(c.Transaction, r.Transaction) {
		return false // fmt.Errorf("Transaction bytes mismatched")
	}

	return true
}

func (l *Offer) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Offer{}
	}
	cr := right
	if cr == nil {
		cr = &Offer{}
	}
	r, ok := cr.(*Offer)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(c.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SignatureRequest) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &SignatureRequest{}
	}
	cr := right
	if cr == nil {
		cr = &SignatureRequest{}
	}
	r, ok := cr.(*SignatureRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(c.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SettlementRequest) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &SettlementRequest{}
	}
	cr := right
	if cr == nil {
		cr = &SettlementRequest{}
	}
	r, ok := cr.(*SettlementRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field TransferTxId - bin
	if !bytes.Equal(c.TransferTxId, r.TransferTxId) {
		return false // fmt.Errorf("TransferTxId bytes mismatched")
	}

	// Field ContractFees - TargetAddress
	if len(c.ContractFees) != len(r.ContractFees) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ContractFees {
		if !v.Equal(r.ContractFees[i]) {
			return false // fmt.Errorf("ContractFees[%d] : %s", i, err)
		}
	}

	// Field Settlement - varbin
	if !bytes.Equal(c.Settlement, r.Settlement) {
		return false // fmt.Errorf("Settlement bytes mismatched")
	}

	return true
}

func (l *OutputMetadata) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &OutputMetadata{}
	}
	cr := right
	if cr == nil {
		cr = &OutputMetadata{}
	}
	r, ok := cr.(*OutputMetadata)
	if !ok {
		return false
	}

	// Field OutputDescription - varchar
	if c.OutputDescription != r.OutputDescription {
		return false // fmt.Errorf("OutputDescription string mismatched")
	}

	// Field Tags - uint
	if len(c.Tags) != len(r.Tags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.Tags {
		if v != r.Tags[i] {
			return false // fmt.Errorf("Element Tags integer mismatched")
		}
	}

	// Field CustomTags - OutputTag
	if len(c.CustomTags) != len(r.CustomTags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.CustomTags {
		if !v.Equal(r.CustomTags[i]) {
			return false // fmt.Errorf("CustomTags[%d] : %s", i, err)
		}
	}

	return true
}

func (l *Distribution) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &Distribution{}
	}
	cr := right
	if cr == nil {
		cr = &Distribution{}
	}
	r, ok := cr.(*Distribution)
	if !ok {
		return false
	}

	// Field InstrumentCode - bin
	if !bytes.Equal(c.InstrumentCode, r.InstrumentCode) {
		return false // fmt.Errorf("InstrumentCode bytes mismatched")
	}

	// Field Timestamp - uint
	if c.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	return true
}

func (l *InitiateRelationship) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InitiateRelationship{}
	}
	cr := right
	if cr == nil {
		cr = &InitiateRelationship{}
	}
	r, ok := cr.(*InitiateRelationship)
	if !ok {
		return false
	}

	// Field Type - uint
	if c.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field Seed - varbin
	if !bytes.Equal(c.Seed, r.Seed) {
		return false // fmt.Errorf("Seed bytes mismatched")
	}

	// Field Flag - varbin
	if !bytes.Equal(c.Flag, r.Flag) {
		return false // fmt.Errorf("Flag bytes mismatched")
	}

	// Field EncryptionType - uint
	if c.EncryptionType != r.EncryptionType {
		return false // fmt.Errorf("EncryptionType integer mismatched")
	}

	// Field ProofOfIdentityType - uint
	if c.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(c.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	// Field ChannelParties - ChannelParty
	if len(c.ChannelParties) != len(r.ChannelParties) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range c.ChannelParties {
		if !v.Equal(r.ChannelParties[i]) {
			return false // fmt.Errorf("ChannelParties[%d] : %s", i, err)
		}
	}

	return true
}

func (l *PendingAcceptRelationship) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &PendingAcceptRelationship{}
	}
	cr := right
	if cr == nil {
		cr = &PendingAcceptRelationship{}
	}
	r, ok := cr.(*PendingAcceptRelationship)
	if !ok {
		return false
	}

	// Field ProofOfIdentityType - uint
	if c.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(c.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	return true
}

func (l *AcceptRelationship) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &AcceptRelationship{}
	}
	cr := right
	if cr == nil {
		cr = &AcceptRelationship{}
	}
	r, ok := cr.(*AcceptRelationship)
	if !ok {
		return false
	}

	// Field ProofOfIdentityType - uint
	if c.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(c.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	return true
}

func (l *RelationshipAmendment) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &RelationshipAmendment{}
	}
	cr := right
	if cr == nil {
		cr = &RelationshipAmendment{}
	}
	r, ok := cr.(*RelationshipAmendment)
	if !ok {
		return false
	}

	// Field Seed - varbin
	if !bytes.Equal(c.Seed, r.Seed) {
		return false // fmt.Errorf("Seed bytes mismatched")
	}

	// Field BaseEncryptionSecret - varbin
	if !bytes.Equal(c.BaseEncryptionSecret, r.BaseEncryptionSecret) {
		return false // fmt.Errorf("BaseEncryptionSecret bytes mismatched")
	}

	// Field AddMemberIndexes - uint
	if c.AddMemberIndexes != r.AddMemberIndexes {
		return false // fmt.Errorf("AddMemberIndexes integer mismatched")
	}

	// Field DropMemberIndexes - uint
	if c.DropMemberIndexes != r.DropMemberIndexes {
		return false // fmt.Errorf("DropMemberIndexes integer mismatched")
	}

	return true
}

func (l *InitiateThread) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &InitiateThread{}
	}
	cr := right
	if cr == nil {
		cr = &InitiateThread{}
	}
	r, ok := cr.(*InitiateThread)
	if !ok {
		return false
	}

	// Field Flag - varbin
	if !bytes.Equal(c.Flag, r.Flag) {
		return false // fmt.Errorf("Flag bytes mismatched")
	}

	// Field Seed - varbin
	if !bytes.Equal(c.Seed, r.Seed) {
		return false // fmt.Errorf("Seed bytes mismatched")
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

func (l *ChannelPartyField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &ChannelPartyField{}
	}
	r, ok := right.(*ChannelPartyField)
	if !ok {
		return false
	}

	if r == nil {
		r = &ChannelPartyField{}
	}

	// Field AdministrativeAddress - varbin
	if !bytes.Equal(c.AdministrativeAddress, r.AdministrativeAddress) {
		return false // fmt.Errorf("AdministrativeAddress bytes mismatched")
	}

	// Field OutputIndex - uint
	if c.OutputIndex != r.OutputIndex {
		return false // fmt.Errorf("OutputIndex integer mismatched")
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

func (l *IdentityOracleProofField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &IdentityOracleProofField{}
	}
	r, ok := right.(*IdentityOracleProofField)
	if !ok {
		return false
	}

	if r == nil {
		r = &IdentityOracleProofField{}
	}

	// Field UserID - varbin
	if !bytes.Equal(c.UserID, r.UserID) {
		return false // fmt.Errorf("UserID bytes mismatched")
	}

	// Field Entity - Entity
	if !c.Entity.Equal(r.Entity) {
		return false // fmt.Errorf("Entity : %s", err)
	}

	// Field OracleSignature - OracleSignature
	if !c.OracleSignature.Equal(r.OracleSignature) {
		return false // fmt.Errorf("OracleSignature : %s", err)
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

func (l *OracleSignatureField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &OracleSignatureField{}
	}
	r, ok := right.(*OracleSignatureField)
	if !ok {
		return false
	}

	if r == nil {
		r = &OracleSignatureField{}
	}

	// Field OracleURL - varchar
	if c.OracleURL != r.OracleURL {
		return false // fmt.Errorf("OracleURL string mismatched")
	}

	// Field BlockHeight - uint
	if c.BlockHeight != r.BlockHeight {
		return false // fmt.Errorf("BlockHeight integer mismatched")
	}

	// Field ValidityPeriod - Period
	if !c.ValidityPeriod.Equal(r.ValidityPeriod) {
		return false // fmt.Errorf("ValidityPeriod : %s", err)
	}

	// Field SignatureAlgorithm - uint
	if c.SignatureAlgorithm != r.SignatureAlgorithm {
		return false // fmt.Errorf("SignatureAlgorithm integer mismatched")
	}

	// Field Signature - varbin
	if !bytes.Equal(c.Signature, r.Signature) {
		return false // fmt.Errorf("Signature bytes mismatched")
	}

	return true
}

func (l *OutpointField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &OutpointField{}
	}
	r, ok := right.(*OutpointField)
	if !ok {
		return false
	}

	if r == nil {
		r = &OutpointField{}
	}

	// Field TxId - bin
	if !bytes.Equal(c.TxId, r.TxId) {
		return false // fmt.Errorf("TxId bytes mismatched")
	}

	// Field OutputIndex - uint
	if c.OutputIndex != r.OutputIndex {
		return false // fmt.Errorf("OutputIndex integer mismatched")
	}

	return true
}

func (l *OutputTagField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &OutputTagField{}
	}
	r, ok := right.(*OutputTagField)
	if !ok {
		return false
	}

	if r == nil {
		r = &OutputTagField{}
	}

	// Field Tag - varchar
	if c.Tag != r.Tag {
		return false // fmt.Errorf("Tag string mismatched")
	}

	return true
}

func (l *PaymailProofField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &PaymailProofField{}
	}
	r, ok := right.(*PaymailProofField)
	if !ok {
		return false
	}

	if r == nil {
		r = &PaymailProofField{}
	}

	// Field UserID - varbin
	if !bytes.Equal(c.UserID, r.UserID) {
		return false // fmt.Errorf("UserID bytes mismatched")
	}

	// Field Handle - varchar
	if c.Handle != r.Handle {
		return false // fmt.Errorf("Handle string mismatched")
	}

	// Field OracleSignature - OracleSignature
	if !c.OracleSignature.Equal(r.OracleSignature) {
		return false // fmt.Errorf("OracleSignature : %s", err)
	}

	return true
}

func (l *PeriodField) Equal(right proto.Message) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &PeriodField{}
	}
	r, ok := right.(*PeriodField)
	if !ok {
		return false
	}

	if r == nil {
		r = &PeriodField{}
	}

	// Field Begin - uint
	if c.Begin != r.Begin {
		return false // fmt.Errorf("Begin integer mismatched")
	}

	// Field End - uint
	if c.End != r.End {
		return false // fmt.Errorf("End integer mismatched")
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
