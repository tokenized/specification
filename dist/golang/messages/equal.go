package messages

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

func (l *PublicMessage) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PublicMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if l.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint
	if !l.Equal(r.Regarding) {
		return false // fmt.Errorf("Regarding : %s", err)
	}

	// Field PublicMessage - Document
	if !l.Equal(r.PublicMessage) {
		return false // fmt.Errorf("PublicMessage : %s", err)
	}

	// Field Attachments - Document
	if len(l.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *PrivateMessage) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PrivateMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if l.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint
	if !l.Equal(r.Regarding) {
		return false // fmt.Errorf("Regarding : %s", err)
	}

	// Field PrivateMessage - Document
	if !l.Equal(r.PrivateMessage) {
		return false // fmt.Errorf("PrivateMessage : %s", err)
	}

	// Field Attachments - Document
	if len(l.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *RevertedTx) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*RevertedTx)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Transaction - varbin
	if !bytes.Equal(l.Transaction, r.Transaction) {
		return false // fmt.Errorf("Transaction bytes mismatched")
	}

	return true
}

func (l *Offer) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Offer)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(l.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SignatureRequest) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*SignatureRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(l.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SettlementRequest) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*SettlementRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field TransferTxId - bin
	if !bytes.Equal(l.TransferTxId, r.TransferTxId) {
		return false // fmt.Errorf("TransferTxId bytes mismatched")
	}

	// Field ContractFees - TargetAddress
	if len(l.ContractFees) != len(r.ContractFees) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ContractFees {
		if !v.Equal(r.ContractFees[i]) {
			return false // fmt.Errorf("ContractFees[%d] : %s", i, err)
		}
	}

	// Field Settlement - varbin
	if !bytes.Equal(l.Settlement, r.Settlement) {
		return false // fmt.Errorf("Settlement bytes mismatched")
	}

	return true
}

func (l *OutputMetadata) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutputMetadata)
	if !ok {
		return false
	}

	// Field OutputDescription - varchar
	if l.OutputDescription != r.OutputDescription {
		return false // fmt.Errorf("OutputDescription string mismatched")
	}

	// Field Tags - uint
	if len(l.Tags) != len(r.Tags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Tags {
		if v != r.Tags[i] {
			return false // fmt.Errorf("Element Tags integer mismatched")
		}
	}

	// Field CustomTags - OutputTag
	if len(l.CustomTags) != len(r.CustomTags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.CustomTags {
		if !v.Equal(r.CustomTags[i]) {
			return false // fmt.Errorf("CustomTags[%d] : %s", i, err)
		}
	}

	return true
}

func (l *InitiateRelationship) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*InitiateRelationship)
	if !ok {
		return false
	}

	// Field Type - uint
	if l.Type != r.Type {
		return false // fmt.Errorf("Type integer mismatched")
	}

	// Field SeedValue - varbin
	if !bytes.Equal(l.SeedValue, r.SeedValue) {
		return false // fmt.Errorf("SeedValue bytes mismatched")
	}

	// Field FlagValue - varbin
	if !bytes.Equal(l.FlagValue, r.FlagValue) {
		return false // fmt.Errorf("FlagValue bytes mismatched")
	}

	// Field EncryptionType - uint
	if l.EncryptionType != r.EncryptionType {
		return false // fmt.Errorf("EncryptionType integer mismatched")
	}

	// Field ProofOfIdentityType - uint
	if l.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(l.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	// Field ChannelParties - ChannelParty
	if len(l.ChannelParties) != len(r.ChannelParties) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ChannelParties {
		if !v.Equal(r.ChannelParties[i]) {
			return false // fmt.Errorf("ChannelParties[%d] : %s", i, err)
		}
	}

	return true
}

func (l *PendingAcceptRelationship) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PendingAcceptRelationship)
	if !ok {
		return false
	}

	// Field ProofOfIdentityType - uint
	if l.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(l.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	return true
}

func (l *AcceptRelationship) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AcceptRelationship)
	if !ok {
		return false
	}

	// Field ProofOfIdentityType - uint
	if l.ProofOfIdentityType != r.ProofOfIdentityType {
		return false // fmt.Errorf("ProofOfIdentityType integer mismatched")
	}

	// Field ProofOfIdentity - varbin
	if !bytes.Equal(l.ProofOfIdentity, r.ProofOfIdentity) {
		return false // fmt.Errorf("ProofOfIdentity bytes mismatched")
	}

	return true
}

func (l *RelationshipAmendment) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*RelationshipAmendment)
	if !ok {
		return false
	}

	// Field SeedValue - varbin
	if !bytes.Equal(l.SeedValue, r.SeedValue) {
		return false // fmt.Errorf("SeedValue bytes mismatched")
	}

	// Field BaseEncryptionSecret - varbin
	if !bytes.Equal(l.BaseEncryptionSecret, r.BaseEncryptionSecret) {
		return false // fmt.Errorf("BaseEncryptionSecret bytes mismatched")
	}

	// Field AddMemberIndexes - uint
	if l.AddMemberIndexes != r.AddMemberIndexes {
		return false // fmt.Errorf("AddMemberIndexes integer mismatched")
	}

	// Field DropMemberIndexes - uint
	if l.DropMemberIndexes != r.DropMemberIndexes {
		return false // fmt.Errorf("DropMemberIndexes integer mismatched")
	}

	return true
}

func (l *InitiateThread) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*InitiateThread)
	if !ok {
		return false
	}

	// Field FlagValue - varbin
	if !bytes.Equal(l.FlagValue, r.FlagValue) {
		return false // fmt.Errorf("FlagValue bytes mismatched")
	}

	// Field SeedValue - varbin
	if !bytes.Equal(l.SeedValue, r.SeedValue) {
		return false // fmt.Errorf("SeedValue bytes mismatched")
	}

	return true
}

func (l *ChannelPartyField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ChannelPartyField)
	if !ok {
		return false
	}

	// Field AdministrativeAddress - varbin
	if !bytes.Equal(l.AdministrativeAddress, r.AdministrativeAddress) {
		return false // fmt.Errorf("AdministrativeAddress bytes mismatched")
	}

	// Field OutputIndex - uint
	if l.OutputIndex != r.OutputIndex {
		return false // fmt.Errorf("OutputIndex integer mismatched")
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

func (l *IdentityOracleProofField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*IdentityOracleProofField)
	if !ok {
		return false
	}

	// Field UserID - varbin
	if !bytes.Equal(l.UserID, r.UserID) {
		return false // fmt.Errorf("UserID bytes mismatched")
	}

	// Field EntityContractAddress - varbin
	if !bytes.Equal(l.EntityContractAddress, r.EntityContractAddress) {
		return false // fmt.Errorf("EntityContractAddress bytes mismatched")
	}

	// Field Entity - varbin
	if !bytes.Equal(l.Entity, r.Entity) {
		return false // fmt.Errorf("Entity bytes mismatched")
	}

	// Field OracleSignature - OracleSignature
	if !l.Equal(r.OracleSignature) {
		return false // fmt.Errorf("OracleSignature : %s", err)
	}

	// Field IdentitySignature - varbin
	if !bytes.Equal(l.IdentitySignature, r.IdentitySignature) {
		return false // fmt.Errorf("IdentitySignature bytes mismatched")
	}

	return true
}

func (l *OracleSignatureField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OracleSignatureField)
	if !ok {
		return false
	}

	// Field OracleURL - varchar
	if l.OracleURL != r.OracleURL {
		return false // fmt.Errorf("OracleURL string mismatched")
	}

	// Field BlockHeight - uint
	if l.BlockHeight != r.BlockHeight {
		return false // fmt.Errorf("BlockHeight integer mismatched")
	}

	// Field ValidityPeriod - Period
	if !l.Equal(r.ValidityPeriod) {
		return false // fmt.Errorf("ValidityPeriod : %s", err)
	}

	// Field SignatureAlgorithm - uint
	if l.SignatureAlgorithm != r.SignatureAlgorithm {
		return false // fmt.Errorf("SignatureAlgorithm integer mismatched")
	}

	// Field Signature - varbin
	if !bytes.Equal(l.Signature, r.Signature) {
		return false // fmt.Errorf("Signature bytes mismatched")
	}

	return true
}

func (l *OutpointField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutpointField)
	if !ok {
		return false
	}

	// Field TxId - bin
	if !bytes.Equal(l.TxId, r.TxId) {
		return false // fmt.Errorf("TxId bytes mismatched")
	}

	// Field OutputIndex - uint
	if l.OutputIndex != r.OutputIndex {
		return false // fmt.Errorf("OutputIndex integer mismatched")
	}

	return true
}

func (l *OutputTagField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutputTagField)
	if !ok {
		return false
	}

	// Field Tag - varchar
	if l.Tag != r.Tag {
		return false // fmt.Errorf("Tag string mismatched")
	}

	return true
}

func (l *PaymailProofField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PaymailProofField)
	if !ok {
		return false
	}

	// Field UserID - varbin
	if !bytes.Equal(l.UserID, r.UserID) {
		return false // fmt.Errorf("UserID bytes mismatched")
	}

	// Field Handle - varchar
	if l.Handle != r.Handle {
		return false // fmt.Errorf("Handle string mismatched")
	}

	// Field OracleSignature - OracleSignature
	if !l.Equal(r.OracleSignature) {
		return false // fmt.Errorf("OracleSignature : %s", err)
	}

	return true
}

func (l *PeriodField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PeriodField)
	if !ok {
		return false
	}

	// Field Begin - uint
	if l.Begin != r.Begin {
		return false // fmt.Errorf("Begin integer mismatched")
	}

	// Field End - uint
	if l.End != r.End {
		return false // fmt.Errorf("End integer mismatched")
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
