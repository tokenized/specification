package messages

func (a *PublicMessage) Copy() *PublicMessage {
	result := &PublicMessage{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field Subject - varchar
	result.Subject = a.Subject

	// Field Regarding - Outpoint
	result.Regarding = a.Regarding.Copy()

	// Field PublicMessage - Document
	result.PublicMessage = a.PublicMessage.Copy()

	// Field Attachments - Document
	result.Attachments = make([]*DocumentField, len(a.Attachments))
	for i, v := range a.Attachments {
		result.Attachments[i] = v.Copy()
	}

	return result
}

func (a *PrivateMessage) Copy() *PrivateMessage {
	result := &PrivateMessage{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field Subject - varchar
	result.Subject = a.Subject

	// Field Regarding - Outpoint
	result.Regarding = a.Regarding.Copy()

	// Field PrivateMessage - Document
	result.PrivateMessage = a.PrivateMessage.Copy()

	// Field Attachments - Document
	result.Attachments = make([]*DocumentField, len(a.Attachments))
	for i, v := range a.Attachments {
		result.Attachments[i] = v.Copy()
	}

	return result
}

func (a *RevertedTx) Copy() *RevertedTx {
	result := &RevertedTx{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field Transaction - varbin
	result.Transaction = make([]byte, len(a.Transaction))
	copy(result.Transaction, a.Transaction)

	return result
}

func (a *Offer) Copy() *Offer {
	result := &Offer{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field Payload - varbin
	result.Payload = make([]byte, len(a.Payload))
	copy(result.Payload, a.Payload)

	return result
}

func (a *SignatureRequest) Copy() *SignatureRequest {
	result := &SignatureRequest{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field Payload - varbin
	result.Payload = make([]byte, len(a.Payload))
	copy(result.Payload, a.Payload)

	return result
}

func (a *SettlementRequest) Copy() *SettlementRequest {
	result := &SettlementRequest{}

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	// Field TransferTxId - bin
	result.TransferTxId = make([]byte, len(a.TransferTxId))
	copy(result.TransferTxId, a.TransferTxId)

	// Field ContractFees - TargetAddress
	result.ContractFees = make([]*TargetAddressField, len(a.ContractFees))
	for i, v := range a.ContractFees {
		result.ContractFees[i] = v.Copy()
	}

	// Field Settlement - varbin
	result.Settlement = make([]byte, len(a.Settlement))
	copy(result.Settlement, a.Settlement)

	return result
}

func (a *OutputMetadata) Copy() *OutputMetadata {
	result := &OutputMetadata{}

	// Field OutputDescription - varchar
	result.OutputDescription = a.OutputDescription

	// Field Tags - uint
	result.Tags = make([]uint32, len(a.Tags))
	for i, v := range a.Tags {
		result.Tags[i] = v
	}

	// Field CustomTags - OutputTag
	result.CustomTags = make([]*OutputTagField, len(a.CustomTags))
	for i, v := range a.CustomTags {
		result.CustomTags[i] = v.Copy()
	}

	return result
}

func (a *Distribution) Copy() *Distribution {
	result := &Distribution{}

	// Field InstrumentCode - bin
	result.InstrumentCode = make([]byte, len(a.InstrumentCode))
	copy(result.InstrumentCode, a.InstrumentCode)

	// Field Timestamp - uint
	result.Timestamp = a.Timestamp

	return result
}

func (a *InitiateRelationship) Copy() *InitiateRelationship {
	result := &InitiateRelationship{}

	// Field Type - uint
	result.Type = a.Type

	// Field Seed - varbin
	result.Seed = make([]byte, len(a.Seed))
	copy(result.Seed, a.Seed)

	// Field Flag - varbin
	result.Flag = make([]byte, len(a.Flag))
	copy(result.Flag, a.Flag)

	// Field EncryptionType - uint
	result.EncryptionType = a.EncryptionType

	// Field ProofOfIdentityType - uint
	result.ProofOfIdentityType = a.ProofOfIdentityType

	// Field ProofOfIdentity - varbin
	result.ProofOfIdentity = make([]byte, len(a.ProofOfIdentity))
	copy(result.ProofOfIdentity, a.ProofOfIdentity)

	// Field ChannelParties - ChannelParty
	result.ChannelParties = make([]*ChannelPartyField, len(a.ChannelParties))
	for i, v := range a.ChannelParties {
		result.ChannelParties[i] = v.Copy()
	}

	return result
}

func (a *PendingAcceptRelationship) Copy() *PendingAcceptRelationship {
	result := &PendingAcceptRelationship{}

	// Field ProofOfIdentityType - uint
	result.ProofOfIdentityType = a.ProofOfIdentityType

	// Field ProofOfIdentity - varbin
	result.ProofOfIdentity = make([]byte, len(a.ProofOfIdentity))
	copy(result.ProofOfIdentity, a.ProofOfIdentity)

	return result
}

func (a *AcceptRelationship) Copy() *AcceptRelationship {
	result := &AcceptRelationship{}

	// Field ProofOfIdentityType - uint
	result.ProofOfIdentityType = a.ProofOfIdentityType

	// Field ProofOfIdentity - varbin
	result.ProofOfIdentity = make([]byte, len(a.ProofOfIdentity))
	copy(result.ProofOfIdentity, a.ProofOfIdentity)

	return result
}

func (a *RelationshipAmendment) Copy() *RelationshipAmendment {
	result := &RelationshipAmendment{}

	// Field Seed - varbin
	result.Seed = make([]byte, len(a.Seed))
	copy(result.Seed, a.Seed)

	// Field BaseEncryptionSecret - varbin
	result.BaseEncryptionSecret = make([]byte, len(a.BaseEncryptionSecret))
	copy(result.BaseEncryptionSecret, a.BaseEncryptionSecret)

	// Field AddMemberIndexes - uint
	result.AddMemberIndexes = a.AddMemberIndexes

	// Field DropMemberIndexes - uint
	result.DropMemberIndexes = a.DropMemberIndexes

	return result
}

func (a *InitiateThread) Copy() *InitiateThread {
	result := &InitiateThread{}

	// Field Flag - varbin
	result.Flag = make([]byte, len(a.Flag))
	copy(result.Flag, a.Flag)

	// Field Seed - varbin
	result.Seed = make([]byte, len(a.Seed))
	copy(result.Seed, a.Seed)

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

func (a *ChannelPartyField) Copy() *ChannelPartyField {
	result := &ChannelPartyField{}

	// Field AdministrativeAddress - varbin
	result.AdministrativeAddress = make([]byte, len(a.AdministrativeAddress))
	copy(result.AdministrativeAddress, a.AdministrativeAddress)

	// Field OutputIndex - uint
	result.OutputIndex = a.OutputIndex

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

func (a *IdentityOracleProofField) Copy() *IdentityOracleProofField {
	result := &IdentityOracleProofField{}

	// Field UserID - varbin
	result.UserID = make([]byte, len(a.UserID))
	copy(result.UserID, a.UserID)

	// Field Entity - Entity
	result.Entity = a.Entity.Copy()

	// Field OracleSignature - OracleSignature
	result.OracleSignature = a.OracleSignature.Copy()

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

func (a *OracleSignatureField) Copy() *OracleSignatureField {
	result := &OracleSignatureField{}

	// Field OracleURL - varchar
	result.OracleURL = a.OracleURL

	// Field BlockHeight - uint
	result.BlockHeight = a.BlockHeight

	// Field ValidityPeriod - Period
	result.ValidityPeriod = a.ValidityPeriod.Copy()

	// Field SignatureAlgorithm - uint
	result.SignatureAlgorithm = a.SignatureAlgorithm

	// Field Signature - varbin
	result.Signature = make([]byte, len(a.Signature))
	copy(result.Signature, a.Signature)

	return result
}

func (a *OutpointField) Copy() *OutpointField {
	result := &OutpointField{}

	// Field TxId - bin
	result.TxId = make([]byte, len(a.TxId))
	copy(result.TxId, a.TxId)

	// Field OutputIndex - uint
	result.OutputIndex = a.OutputIndex

	return result
}

func (a *OutputTagField) Copy() *OutputTagField {
	result := &OutputTagField{}

	// Field Tag - varchar
	result.Tag = a.Tag

	return result
}

func (a *PaymailProofField) Copy() *PaymailProofField {
	result := &PaymailProofField{}

	// Field UserID - varbin
	result.UserID = make([]byte, len(a.UserID))
	copy(result.UserID, a.UserID)

	// Field Handle - varchar
	result.Handle = a.Handle

	// Field OracleSignature - OracleSignature
	result.OracleSignature = a.OracleSignature.Copy()

	return result
}

func (a *PeriodField) Copy() *PeriodField {
	result := &PeriodField{}

	// Field Begin - uint
	result.Begin = a.Begin

	// Field End - uint
	result.End = a.End

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
