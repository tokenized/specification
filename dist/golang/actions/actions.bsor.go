package actions

type ContractOffer struct {
	ContractName              string                           `bsor:"1" json:"ContractName"`
	BodyOfAgreementType       uint8                            `bsor:"2" json:"BodyOfAgreementType"`
	BodyOfAgreement           []byte                           `bsor:"3" json:"BodyOfAgreement"`
	SupportingDocs            []*DocumentField                 `bsor:"5" json:"SupportingDocs"`
	ContractExpiration        uint64                           `bsor:"8" json:"ContractExpiration"`
	ContractURI               string                           `bsor:"9" json:"ContractURI"`
	Issuer                    *EntityField                     `bsor:"10" json:"Issuer"`
	ContractOperatorIncluded  bool                             `bsor:"12" json:"ContractOperatorIncluded"`
	ContractFee               uint64                           `bsor:"17" json:"ContractFee"`
	VotingSystems             []*VotingSystemField             `bsor:"18" json:"VotingSystems"`
	ContractPermissions       []byte                           `bsor:"19" json:"ContractPermissions"`
	RestrictedQtyInstruments  uint64                           `bsor:"20" json:"RestrictedQtyInstruments"`
	AdministrationProposal    bool                             `bsor:"21" json:"AdministrationProposal"`
	HolderProposal            bool                             `bsor:"22" json:"HolderProposal"`
	Oracles                   []*OracleField                   `bsor:"23" json:"Oracles"`
	MasterAddress             []byte                           `bsor:"24" json:"MasterAddress"`
	EntityContract            []byte                           `bsor:"25" json:"EntityContract"`
	OperatorEntityContract    []byte                           `bsor:"26" json:"OperatorEntityContract"`
	ContractType              uint8                            `bsor:"27" json:"ContractType"`
	Services                  []*ServiceField                  `bsor:"28" json:"Services"`
	AdminIdentityCertificates []*AdminIdentityCertificateField `bsor:"29" json:"AdminIdentityCertificates"`
	GoverningLaw              string                           `bsor:"30" json:"GoverningLaw"`
	Jurisdiction              string                           `bsor:"31" json:"Jurisdiction"`
}

type ContractFormation struct {
	ContractName              string                           `bsor:"1" json:"ContractName"`
	BodyOfAgreementType       uint8                            `bsor:"2" json:"BodyOfAgreementType"`
	BodyOfAgreement           []byte                           `bsor:"3" json:"BodyOfAgreement"`
	SupportingDocs            []*DocumentField                 `bsor:"5" json:"SupportingDocs"`
	ContractExpiration        uint64                           `bsor:"8" json:"ContractExpiration"`
	ContractURI               string                           `bsor:"9" json:"ContractURI"`
	Issuer                    *EntityField                     `bsor:"10" json:"Issuer"`
	ContractFee               uint64                           `bsor:"16" json:"ContractFee"`
	VotingSystems             []*VotingSystemField             `bsor:"17" json:"VotingSystems"`
	ContractPermissions       []byte                           `bsor:"18" json:"ContractPermissions"`
	RestrictedQtyInstruments  uint64                           `bsor:"19" json:"RestrictedQtyInstruments"`
	AdministrationProposal    bool                             `bsor:"20" json:"AdministrationProposal"`
	HolderProposal            bool                             `bsor:"21" json:"HolderProposal"`
	Oracles                   []*OracleField                   `bsor:"22" json:"Oracles"`
	MasterAddress             []byte                           `bsor:"23" json:"MasterAddress"`
	ContractRevision          uint32                           `bsor:"24" json:"ContractRevision"`
	Timestamp                 uint64                           `bsor:"25" json:"Timestamp"`
	EntityContract            []byte                           `bsor:"26" json:"EntityContract"`
	OperatorEntityContract    []byte                           `bsor:"27" json:"OperatorEntityContract"`
	ContractType              uint8                            `bsor:"28" json:"ContractType"`
	Services                  []*ServiceField                  `bsor:"29" json:"Services"`
	AdminIdentityCertificates []*AdminIdentityCertificateField `bsor:"30" json:"AdminIdentityCertificates"`
	AdminAddress              []byte                           `bsor:"31" json:"AdminAddress"`
	OperatorAddress           []byte                           `bsor:"32" json:"OperatorAddress"`
	GoverningLaw              string                           `bsor:"33" json:"GoverningLaw"`
	Jurisdiction              string                           `bsor:"34" json:"Jurisdiction"`
}

type ContractAmendment struct {
	ChangeAdministrationAddress bool              `bsor:"1" json:"ChangeAdministrationAddress"`
	ChangeOperatorAddress       bool              `bsor:"2" json:"ChangeOperatorAddress"`
	ContractRevision            uint32            `bsor:"3" json:"ContractRevision"`
	Amendments                  []*AmendmentField `bsor:"4" json:"Amendments"`
	RefTxID                     []byte            `bsor:"5" json:"RefTxID"`
}

type StaticContractFormation struct {
	ContractName               string           `bsor:"1" json:"ContractName"`
	ContractCode               []byte           `bsor:"2" json:"ContractCode"`
	BodyOfAgreementType        uint8            `bsor:"3" json:"BodyOfAgreementType"`
	BodyOfAgreement            []byte           `bsor:"4" json:"BodyOfAgreement"`
	ContractType               string           `bsor:"5" json:"ContractType"`
	SupportingDocs             []*DocumentField `bsor:"6" json:"SupportingDocs"`
	ContractRevision           uint32           `bsor:"7" json:"ContractRevision"`
	EffectiveDate              uint64           `bsor:"10" json:"EffectiveDate"`
	ContractExpiration         uint64           `bsor:"11" json:"ContractExpiration"`
	ContractURI                string           `bsor:"12" json:"ContractURI"`
	PrevRevTxID                []byte           `bsor:"13" json:"PrevRevTxID"`
	Entities                   []*EntityField   `bsor:"14" json:"Entities"`
	EntityOracle               *OracleField     `bsor:"15" json:"EntityOracle"`
	EntityOracleSignature      []byte           `bsor:"16" json:"EntityOracleSignature"`
	EntityOracleSigBlockHeight uint32           `bsor:"17" json:"EntityOracleSigBlockHeight"`
	GoverningLaw               string           `bsor:"18" json:"GoverningLaw"`
	Jurisdiction               string           `bsor:"19" json:"Jurisdiction"`
}

type ContractAddressChange struct {
	NewContractAddress []byte `bsor:"1" json:"NewContractAddress"`
	Timestamp          uint64 `bsor:"2" json:"Timestamp"`
}

type BodyOfAgreementOffer struct {
	Chapters    []*ChapterField     `bsor:"1" json:"Chapters"`
	Definitions []*DefinedTermField `bsor:"2" json:"Definitions"`
}

type BodyOfAgreementFormation struct {
	Chapters    []*ChapterField     `bsor:"1" json:"Chapters"`
	Definitions []*DefinedTermField `bsor:"2" json:"Definitions"`
	Revision    uint32              `bsor:"3" json:"Revision"`
	Timestamp   uint64              `bsor:"4" json:"Timestamp"`
}

type BodyOfAgreementAmendment struct {
	Revision   uint32            `bsor:"1" json:"Revision"`
	Amendments []*AmendmentField `bsor:"2" json:"Amendments"`
	RefTxID    []byte            `bsor:"3" json:"RefTxID"`
}

type InstrumentDefinition struct {
	InstrumentPermissions            []byte   `bsor:"1" json:"InstrumentPermissions"`
	EnforcementOrdersPermitted       bool     `bsor:"4" json:"EnforcementOrdersPermitted"`
	VotingRights                     bool     `bsor:"5" json:"VotingRights"`
	VoteMultiplier                   uint8    `bsor:"6" json:"VoteMultiplier"`
	AdministrationProposal           bool     `bsor:"7" json:"AdministrationProposal"`
	HolderProposal                   bool     `bsor:"8" json:"HolderProposal"`
	InstrumentModificationGovernance uint8    `bsor:"9" json:"InstrumentModificationGovernance"`
	AuthorizedTokenQty               uint64   `bsor:"10" json:"AuthorizedTokenQty"`
	InstrumentType                   string   `bsor:"11" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentPayload                []byte   `bsor:"12" json:"InstrumentPayload"`
	TradeRestrictions                []string `bsor:"13" json:"TradeRestrictions"`
	InstrumentPayloadVersion         uint8    `bsor:"14" json:"InstrumentPayloadVersion"`
}

type InstrumentCreation struct {
	InstrumentCode                   []byte   `bsor:"1" json:"InstrumentCode"`
	InstrumentIndex                  uint64   `bsor:"2" json:"InstrumentIndex"`
	InstrumentPermissions            []byte   `bsor:"3" json:"InstrumentPermissions"`
	EnforcementOrdersPermitted       bool     `bsor:"6" json:"EnforcementOrdersPermitted"`
	VotingRights                     bool     `bsor:"7" json:"VotingRights"`
	VoteMultiplier                   uint8    `bsor:"8" json:"VoteMultiplier"`
	AdministrationProposal           bool     `bsor:"9" json:"AdministrationProposal"`
	HolderProposal                   bool     `bsor:"10" json:"HolderProposal"`
	InstrumentModificationGovernance uint8    `bsor:"11" json:"InstrumentModificationGovernance"`
	AuthorizedTokenQty               uint64   `bsor:"12" json:"AuthorizedTokenQty"`
	InstrumentType                   string   `bsor:"13" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentPayload                []byte   `bsor:"14" json:"InstrumentPayload"`
	InstrumentRevision               uint32   `bsor:"15" json:"InstrumentRevision"`
	Timestamp                        uint64   `bsor:"16" json:"Timestamp"`
	TradeRestrictions                []string `bsor:"17" json:"TradeRestrictions"`
	InstrumentPayloadVersion         uint8    `bsor:"18" json:"InstrumentPayloadVersion"`
}

type InstrumentModification struct {
	InstrumentType           string            `bsor:"1" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode           []byte            `bsor:"2" json:"InstrumentCode"`
	InstrumentRevision       uint32            `bsor:"3" json:"InstrumentRevision"`
	Amendments               []*AmendmentField `bsor:"4" json:"Amendments"`
	RefTxID                  []byte            `bsor:"5" json:"RefTxID"`
	InstrumentPayloadVersion uint8             `bsor:"6" json:"InstrumentPayloadVersion"`
}

type Transfer struct {
	Instruments        []*InstrumentTransferField `bsor:"1" json:"Instruments"`
	OfferExpiry        uint64                     `bsor:"2" json:"OfferExpiry"`
	ExchangeFee        uint64                     `bsor:"3" json:"ExchangeFee"`
	ExchangeFeeAddress []byte                     `bsor:"4" json:"ExchangeFeeAddress"`
}

type Settlement struct {
	Instruments []*InstrumentSettlementField `bsor:"1" json:"Instruments"`
	Timestamp   uint64                       `bsor:"2" json:"Timestamp"`
}

type Proposal struct {
	Type                 uint8             `bsor:"1" json:"Type"`
	InstrumentType       string            `bsor:"2" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode       []byte            `bsor:"3" json:"InstrumentCode"`
	VoteSystem           uint8             `bsor:"4" json:"VoteSystem"`
	ProposedAmendments   []*AmendmentField `bsor:"5" json:"ProposedAmendments"`
	VoteOptions          string            `bsor:"6" json:"VoteOptions"`
	VoteMax              uint8             `bsor:"7" json:"VoteMax"`
	ProposalDescription  string            `bsor:"8" json:"ProposalDescription"`
	ProposalDocumentHash []byte            `bsor:"9" json:"ProposalDocumentHash"`
	VoteCutOffTimestamp  uint64            `bsor:"10" json:"VoteCutOffTimestamp"`
}

type Vote struct {
	Timestamp uint64 `bsor:"1" json:"Timestamp"`
}

type BallotCast struct {
	VoteTxId []byte `bsor:"1" json:"VoteTxId"`
	Vote     string `bsor:"2" json:"Vote"`
}

type BallotCounted struct {
	VoteTxId  []byte `bsor:"1" json:"VoteTxId"`
	Vote      string `bsor:"2" json:"Vote"`
	Quantity  uint64 `bsor:"3" json:"Quantity"`
	Timestamp uint64 `bsor:"4" json:"Timestamp"`
}

type Result struct {
	InstrumentType     string            `bsor:"1" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode     []byte            `bsor:"2" json:"InstrumentCode"`
	ProposedAmendments []*AmendmentField `bsor:"3" json:"ProposedAmendments"`
	VoteTxId           []byte            `bsor:"4" json:"VoteTxId"`
	OptionTally        []uint64          `bsor:"5" json:"OptionTally"`
	Result             string            `bsor:"6" json:"Result"`
	Timestamp          uint64            `bsor:"7" json:"Timestamp"`
}

type Order struct {
	ComplianceAction         string                       `bsor:"1" bsor_fixed_size:"1" json:"ComplianceAction"`
	InstrumentType           string                       `bsor:"2" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode           []byte                       `bsor:"3" json:"InstrumentCode"`
	TargetAddresses          []*TargetAddressField        `bsor:"4" json:"TargetAddresses"`
	FreezeTxId               []byte                       `bsor:"5" json:"FreezeTxId"`
	FreezePeriod             uint64                       `bsor:"6" json:"FreezePeriod"`
	DepositAddress           []byte                       `bsor:"7" json:"DepositAddress"`
	AuthorityName            string                       `bsor:"8" json:"AuthorityName"`
	AuthorityPublicKey       []byte                       `bsor:"9" json:"AuthorityPublicKey"`
	SignatureAlgorithm       uint8                        `bsor:"10" json:"SignatureAlgorithm"`
	OrderSignature           []byte                       `bsor:"11" json:"OrderSignature"`
	BitcoinDispersions       []*QuantityIndexField        `bsor:"14" json:"BitcoinDispersions"`
	Message                  string                       `bsor:"15" json:"Message"`
	SupportingEvidenceFormat uint8                        `bsor:"16" json:"SupportingEvidenceFormat"`
	SupportingEvidence       []byte                       `bsor:"17" json:"SupportingEvidence"`
	ReferenceTransactions    []*ReferenceTransactionField `bsor:"18" json:"ReferenceTransactions"`
}

type Freeze struct {
	InstrumentType string                `bsor:"1" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode []byte                `bsor:"2" json:"InstrumentCode"`
	Quantities     []*QuantityIndexField `bsor:"3" json:"Quantities"`
	FreezePeriod   uint64                `bsor:"4" json:"FreezePeriod"`
	Timestamp      uint64                `bsor:"5" json:"Timestamp"`
}

type Thaw struct {
	FreezeTxId []byte `bsor:"1" json:"FreezeTxId"`
	Timestamp  uint64 `bsor:"2" json:"Timestamp"`
}

type Confiscation struct {
	InstrumentType string                `bsor:"1" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode []byte                `bsor:"2" json:"InstrumentCode"`
	Quantities     []*QuantityIndexField `bsor:"3" json:"Quantities"`
	DepositQty     uint64                `bsor:"4" json:"DepositQty"`
	Timestamp      uint64                `bsor:"5" json:"Timestamp"`
}

type Reconciliation struct {
	InstrumentType string                `bsor:"1" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode []byte                `bsor:"2" json:"InstrumentCode"`
	Quantities     []*QuantityIndexField `bsor:"3" json:"Quantities"`
	Timestamp      uint64                `bsor:"4" json:"Timestamp"`
}

type Establishment struct {
	Message string `bsor:"1" json:"Message"`
}

type Addition struct {
	Message string `bsor:"1" json:"Message"`
}

type Alteration struct {
	EntryTxID []byte `bsor:"1" json:"EntryTxID"`
	Message   string `bsor:"2" json:"Message"`
}

type Removal struct {
	EntryTxID []byte `bsor:"1" json:"EntryTxID"`
	Message   string `bsor:"2" json:"Message"`
}

type Message struct {
	SenderIndexes         []uint32 `bsor:"1" json:"SenderIndexes"`
	ReceiverIndexes       []uint32 `bsor:"2" json:"ReceiverIndexes"`
	MessageCode           uint16   `bsor:"3" json:"MessageCode"`
	MessagePayload        []byte   `bsor:"4" json:"MessagePayload"`
	MessagePayloadVersion uint8    `bsor:"5" json:"MessagePayloadVersion"`
}

type Rejection struct {
	AddressIndexes     []uint32 `bsor:"1" json:"AddressIndexes"`
	RejectAddressIndex uint32   `bsor:"2" json:"RejectAddressIndex"`
	RejectionCode      uint8    `bsor:"3" json:"RejectionCode"`
	Message            string   `bsor:"4" json:"Message"`
	Timestamp          uint64   `bsor:"5" json:"Timestamp"`
}

type AdministratorField struct {
	Type uint8  `bsor:"1" json:"Type"`
	Name string `bsor:"2" json:"Name"`
}

type AdminIdentityCertificateField struct {
	EntityContract []byte `bsor:"1" json:"EntityContract"`
	Signature      []byte `bsor:"2" json:"Signature"`
	BlockHeight    uint32 `bsor:"3" json:"BlockHeight"`
	Expiration     uint64 `bsor:"4" json:"Expiration"`
}

type AmendmentField struct {
	FieldIndexPath []byte `bsor:"1" json:"FieldIndexPath"`
	Operation      uint8  `bsor:"2" json:"Operation"`
	Data           []byte `bsor:"3" json:"Data"`
}

type InstrumentReceiverField struct {
	Address               []byte `bsor:"1" json:"Address"`
	Quantity              uint64 `bsor:"2" json:"Quantity"`
	OracleSigAlgorithm    uint8  `bsor:"3" json:"OracleSigAlgorithm"`
	OracleIndex           uint8  `bsor:"4" json:"OracleIndex"`
	OracleConfirmationSig []byte `bsor:"5" json:"OracleConfirmationSig"`
	OracleSigBlockHeight  uint32 `bsor:"6" json:"OracleSigBlockHeight"`
	OracleSigExpiry       uint64 `bsor:"7" json:"OracleSigExpiry"`
}

type InstrumentSettlementField struct {
	ContractIndex  uint32                `bsor:"1" json:"ContractIndex"`
	InstrumentType string                `bsor:"2" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode []byte                `bsor:"3" json:"InstrumentCode"`
	Settlements    []*QuantityIndexField `bsor:"4" json:"Settlements"`
}

type InstrumentTransferField struct {
	ContractIndex       uint32                     `bsor:"1" json:"ContractIndex"`
	InstrumentType      string                     `bsor:"2" bsor_fixed_size:"3" json:"InstrumentType"`
	InstrumentCode      []byte                     `bsor:"3" json:"InstrumentCode"`
	InstrumentSenders   []*QuantityIndexField      `bsor:"4" json:"InstrumentSenders"`
	InstrumentReceivers []*InstrumentReceiverField `bsor:"5" json:"InstrumentReceivers"`
}

type ChapterField struct {
	Title    string         `bsor:"1" json:"Title"`
	Preamble string         `bsor:"2" json:"Preamble"`
	Articles []*ClauseField `bsor:"3" json:"Articles"`
}

type ClauseField struct {
	Title    string         `bsor:"1" json:"Title"`
	Body     string         `bsor:"2" json:"Body"`
	Children []*ClauseField `bsor:"3" json:"Children"`
}

type DefinedTermField struct {
	Term       string `bsor:"1" json:"Term"`
	Definition string `bsor:"2" json:"Definition"`
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

type ManagerField struct {
	Type uint8  `bsor:"1" json:"Type"`
	Name string `bsor:"2" json:"Name"`
}

type OracleField struct {
	OracleTypes    []uint8 `bsor:"4" json:"OracleTypes"`
	EntityContract []byte  `bsor:"5" json:"EntityContract"`
}

type QuantityIndexField struct {
	Index    uint32 `bsor:"1" json:"Index"`
	Quantity uint64 `bsor:"2" json:"Quantity"`
}

type ReferenceTransactionField struct {
	Transaction []byte   `bsor:"1" json:"Transaction"`
	Outputs     [][]byte `bsor:"2" json:"Outputs"`
}

type ServiceField struct {
	Type      uint8  `bsor:"1" json:"Type"`
	URL       string `bsor:"2" json:"URL"`
	PublicKey []byte `bsor:"3" json:"PublicKey"`
}

type TargetAddressField struct {
	Address  []byte `bsor:"1" json:"Address"`
	Quantity uint64 `bsor:"2" json:"Quantity"`
}

type VotingSystemField struct {
	Name                    string `bsor:"1" json:"Name"`
	VoteType                string `bsor:"2" bsor_fixed_size:"1" json:"VoteType"`
	TallyLogic              uint8  `bsor:"3" json:"TallyLogic"`
	ThresholdPercentage     uint8  `bsor:"4" json:"ThresholdPercentage"`
	VoteMultiplierPermitted bool   `bsor:"5" json:"VoteMultiplierPermitted"`
	HolderProposalFee       uint64 `bsor:"6" json:"HolderProposalFee"`
}
