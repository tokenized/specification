package messages

type PublicMessage struct {
	Timestamp     uint64           `bsor:"1" json:"Timestamp"`
	Subject       string           `bsor:"2" json:"Subject"`
	Regarding     *OutpointField   `bsor:"3" json:"Regarding"`
	PublicMessage *DocumentField   `bsor:"4" json:"PublicMessage"`
	Attachments   []*DocumentField `bsor:"5" json:"Attachments"`
}

type PrivateMessage struct {
	Timestamp      uint64           `bsor:"1" json:"Timestamp"`
	Subject        string           `bsor:"2" json:"Subject"`
	Regarding      *OutpointField   `bsor:"3" json:"Regarding"`
	PrivateMessage *DocumentField   `bsor:"4" json:"PrivateMessage"`
	Attachments    []*DocumentField `bsor:"5" json:"Attachments"`
}

type RevertedTx struct {
	Timestamp   uint64 `bsor:"1" json:"Timestamp"`
	Transaction []byte `bsor:"2" json:"Transaction"`
}

type Offer struct {
	Timestamp uint64 `bsor:"1" json:"Timestamp"`
	Payload   []byte `bsor:"2" json:"Payload"`
}

type SignatureRequest struct {
	Timestamp uint64 `bsor:"1" json:"Timestamp"`
	Payload   []byte `bsor:"2" json:"Payload"`
}

type SettlementRequest struct {
	Timestamp    uint64                `bsor:"1" json:"Timestamp"`
	TransferTxId []byte                `bsor:"2" json:"TransferTxId"`
	ContractFees []*TargetAddressField `bsor:"3" json:"ContractFees"`
	Settlement   []byte                `bsor:"4" json:"Settlement"`
}

type OutputMetadata struct {
	OutputDescription string            `bsor:"1" json:"OutputDescription"`
	Tags              []uint32          `bsor:"2" json:"Tags"`
	CustomTags        []*OutputTagField `bsor:"3" json:"CustomTags"`
}

type Distribution struct {
	InstrumentCode []byte `bsor:"1" json:"InstrumentCode"`
	Timestamp      uint64 `bsor:"2" json:"Timestamp"`
}

type InitiateRelationship struct {
	Type                uint32               `bsor:"1" json:"Type"`
	Seed                []byte               `bsor:"2" json:"Seed"`
	Flag                []byte               `bsor:"3" json:"Flag"`
	EncryptionType      uint32               `bsor:"4" json:"EncryptionType"`
	ProofOfIdentityType uint32               `bsor:"5" json:"ProofOfIdentityType"`
	ProofOfIdentity     []byte               `bsor:"6" json:"ProofOfIdentity"`
	ChannelParties      []*ChannelPartyField `bsor:"7" json:"ChannelParties"`
}

type PendingAcceptRelationship struct {
	ProofOfIdentityType uint32 `bsor:"1" json:"ProofOfIdentityType"`
	ProofOfIdentity     []byte `bsor:"2" json:"ProofOfIdentity"`
}

type AcceptRelationship struct {
	ProofOfIdentityType uint32 `bsor:"1" json:"ProofOfIdentityType"`
	ProofOfIdentity     []byte `bsor:"2" json:"ProofOfIdentity"`
}

type RelationshipAmendment struct {
	Seed                 []byte `bsor:"1" json:"Seed"`
	BaseEncryptionSecret []byte `bsor:"2" json:"BaseEncryptionSecret"`
	AddMemberIndexes     uint32 `bsor:"3" json:"AddMemberIndexes"`
	DropMemberIndexes    uint32 `bsor:"4" json:"DropMemberIndexes"`
}

type InitiateThread struct {
	Flag []byte `bsor:"1" json:"Flag"`
	Seed []byte `bsor:"2" json:"Seed"`
}

type AdministratorField struct {
	Type uint32 `bsor:"1" json:"Type"`
	Name string `bsor:"2" json:"Name"`
}

type ChannelPartyField struct {
	AdministrativeAddress []byte `bsor:"1" json:"AdministrativeAddress"`
	OutputIndex           uint32 `bsor:"2" json:"OutputIndex"`
}

type DocumentField struct {
	Name     string `bsor:"1" json:"Name"`
	Type     string `bsor:"2" json:"Type"`
	Contents []byte `bsor:"3" json:"Contents"`
}

type EntityField struct {
	Name                       string                `bsor:"1" json:"Name"`
	Type                       string                `bsor:"2" bsor_fixed_size:"1" json:"Type"`
	LEI                        string                `bsor:"3" bsor_fixed_size:"20" json:"LEI"`
	UnitNumber                 string                `bsor:"4" json:"UnitNumber"`
	BuildingNumber             string                `bsor:"5" json:"BuildingNumber"`
	Street                     string                `bsor:"6" json:"Street"`
	SuburbCity                 string                `bsor:"7" json:"SuburbCity"`
	TerritoryStateProvinceCode string                `bsor:"8" bsor_fixed_size:"5" json:"TerritoryStateProvinceCode"`
	CountryCode                string                `bsor:"9" bsor_fixed_size:"3" json:"CountryCode"`
	PostalZIPCode              string                `bsor:"10" bsor_fixed_size:"12" json:"PostalZIPCode"`
	EmailAddress               string                `bsor:"11" json:"EmailAddress"`
	PhoneNumber                string                `bsor:"12" json:"PhoneNumber"`
	Administration             []*AdministratorField `bsor:"13" json:"Administration"`
	Management                 []*ManagerField       `bsor:"14" json:"Management"`
	DomainName                 string                `bsor:"15" json:"DomainName"`
	PaymailHandle              string                `bsor:"17" json:"PaymailHandle"`
}

type IdentityOracleProofField struct {
	UserID          []byte                `bsor:"1" json:"UserID"`
	Entity          *EntityField          `bsor:"2" json:"Entity"`
	OracleSignature *OracleSignatureField `bsor:"3" json:"OracleSignature"`
}

type ManagerField struct {
	Type uint32 `bsor:"1" json:"Type"`
	Name string `bsor:"2" json:"Name"`
}

type OracleSignatureField struct {
	OracleURL          string       `bsor:"1" json:"OracleURL"`
	BlockHeight        uint32       `bsor:"2" json:"BlockHeight"`
	ValidityPeriod     *PeriodField `bsor:"3" json:"ValidityPeriod"`
	SignatureAlgorithm uint32       `bsor:"4" json:"SignatureAlgorithm"`
	Signature          []byte       `bsor:"5" json:"Signature"`
}

type OutpointField struct {
	TxId        []byte `bsor:"1" json:"TxId"`
	OutputIndex uint32 `bsor:"2" json:"OutputIndex"`
}

type OutputTagField struct {
	Tag string `bsor:"1" json:"Tag"`
}

type PaymailProofField struct {
	UserID          []byte                `bsor:"1" json:"UserID"`
	Handle          string                `bsor:"2" json:"Handle"`
	OracleSignature *OracleSignatureField `bsor:"3" json:"OracleSignature"`
}

type PeriodField struct {
	Begin uint64 `bsor:"1" json:"Begin"`
	End   uint64 `bsor:"2" json:"End"`
}

type TargetAddressField struct {
	Address  []byte `bsor:"1" json:"Address"`
	Quantity uint64 `bsor:"2" json:"Quantity"`
}
