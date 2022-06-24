package instruments

type Membership struct {
	AgeRestriction      *AgeRestrictionField `bsor:"1" json:"AgeRestriction"`
	ValidFrom           uint64               `bsor:"2" json:"ValidFrom"`
	ExpirationTimestamp uint64               `bsor:"3" json:"ExpirationTimestamp"`
	ID                  string               `bsor:"4" json:"ID"`
	MembershipClass     string               `bsor:"5" json:"MembershipClass"`
	RoleType            string               `bsor:"6" json:"RoleType"`
	MembershipType      string               `bsor:"7" json:"MembershipType"`
	Description         string               `bsor:"8" json:"Description"`
	TransfersPermitted  bool                 `bsor:"9" json:"TransfersPermitted"`
}

type Currency struct {
	CurrencyCode      string `bsor:"1" bsor_fixed_size:"3" json:"CurrencyCode"`
	MonetaryAuthority string `bsor:"2" json:"MonetaryAuthority"`
	Precision         uint64 `bsor:"4" json:"Precision"`
}

type ShareCommon struct {
	Ticker             string `bsor:"1" bsor_fixed_size:"5" json:"Ticker"`
	ISIN               string `bsor:"2" bsor_fixed_size:"12" json:"ISIN"`
	Description        string `bsor:"3" json:"Description"`
	TransfersPermitted bool   `bsor:"4" json:"TransfersPermitted"`
}

type BondFixedRate struct {
	Name                       string               `bsor:"1" json:"Name"`
	BondType                   string               `bsor:"2" bsor_fixed_size:"1" json:"BondType"`
	ISIN                       string               `bsor:"3" json:"ISIN"`
	Collateral                 string               `bsor:"4" json:"Collateral"`
	ParValue                   *CurrencyValueField  `bsor:"5" json:"ParValue"`
	InterestRate               *RateField           `bsor:"6" json:"InterestRate"`
	InterestPaymentInitialDate uint64               `bsor:"7" json:"InterestPaymentInitialDate"`
	InterestPaymentDateDeltas  []uint64             `bsor:"8" json:"InterestPaymentDateDeltas"`
	LatePaymentPenaltyRate     *RateField           `bsor:"9" json:"LatePaymentPenaltyRate"`
	LatePaymentWindow          uint64               `bsor:"10" json:"LatePaymentWindow"`
	LatePaymentPenaltyPeriod   uint64               `bsor:"11" json:"LatePaymentPenaltyPeriod"`
	MaturityDate               uint64               `bsor:"12" json:"MaturityDate"`
	AgeRestriction             *AgeRestrictionField `bsor:"13" json:"AgeRestriction"`
	TransfersPermitted         bool                 `bsor:"14" json:"TransfersPermitted"`
}

type Coupon struct {
	RedeemingEntity     string              `bsor:"1" json:"RedeemingEntity"`
	ValidFromTimestamp  uint64              `bsor:"2" json:"ValidFromTimestamp"`
	ExpirationTimestamp uint64              `bsor:"3" json:"ExpirationTimestamp"`
	CouponName          string              `bsor:"6" json:"CouponName"`
	TransfersPermitted  bool                `bsor:"8" json:"TransfersPermitted"`
	FaceValue           *CurrencyValueField `bsor:"9" json:"FaceValue"`
	RedemptionVenue     string              `bsor:"10" json:"RedemptionVenue"`
	Details             string              `bsor:"11" json:"Details"`
}

type LoyaltyPoints struct {
	AgeRestriction      *AgeRestrictionField `bsor:"1" json:"AgeRestriction"`
	ProgramName         string               `bsor:"2" json:"ProgramName"`
	ExpirationTimestamp uint64               `bsor:"4" json:"ExpirationTimestamp"`
	Details             string               `bsor:"5" json:"Details"`
	TransfersPermitted  bool                 `bsor:"6" json:"TransfersPermitted"`
}

type TicketAdmission struct {
	AgeRestriction      *AgeRestrictionField `bsor:"1" json:"AgeRestriction"`
	Venue               string               `bsor:"3" json:"Venue"`
	Area                string               `bsor:"5" json:"Area"`
	Seat                string               `bsor:"6" json:"Seat"`
	EventStartTimestamp uint64               `bsor:"7" json:"EventStartTimestamp"`
	EventName           string               `bsor:"10" json:"EventName"`
	TransfersPermitted  bool                 `bsor:"11" json:"TransfersPermitted"`
	Details             string               `bsor:"12" json:"Details"`
	Section             string               `bsor:"13" json:"Section"`
	Row                 string               `bsor:"14" json:"Row"`
	EventEndTimestamp   uint64               `bsor:"15" json:"EventEndTimestamp"`
}

type CasinoChip struct {
	UseType             string               `bsor:"2" bsor_fixed_size:"1" json:"UseType"`
	AgeRestriction      *AgeRestrictionField `bsor:"3" json:"AgeRestriction"`
	ExpirationTimestamp uint64               `bsor:"5" json:"ExpirationTimestamp"`
	TransfersPermitted  bool                 `bsor:"7" json:"TransfersPermitted"`
	CasinoName          string               `bsor:"8" json:"CasinoName"`
	FaceValue           *CurrencyValueField  `bsor:"9" json:"FaceValue"`
}

type InformationServiceLicense struct {
	AgeRestriction      *AgeRestrictionField `bsor:"1" json:"AgeRestriction"`
	ExpirationTimestamp uint64               `bsor:"2" json:"ExpirationTimestamp"`
	ServiceName         string               `bsor:"3" json:"ServiceName"`
	TransfersPermitted  bool                 `bsor:"4" json:"TransfersPermitted"`
	URL                 string               `bsor:"5" json:"URL"`
}

type AgeRestrictionField struct {
	Lower uint32 `bsor:"1" json:"Lower"`
	Upper uint32 `bsor:"2" json:"Upper"`
}

type CurrencyValueField struct {
	Value        uint64 `bsor:"1" json:"Value"`
	CurrencyCode string `bsor:"2" bsor_fixed_size:"3" json:"CurrencyCode"`
	Precision    uint32 `bsor:"3" json:"Precision"`
}

type RateField struct {
	Precision uint32 `bsor:"1" json:"Precision"`
	Value     uint64 `bsor:"2" json:"Value"`
}
