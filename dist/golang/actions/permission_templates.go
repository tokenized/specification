package actions

import "github.com/tokenized/specification/dist/golang/permissions"

type PermissionConfig struct {
	VotingSystems       []VotingSystemField
	ContractPermissions permissions.Permissions
	AssetPermissions    map[string]permissions.Permissions
}

var PrivateCompany = PermissionConfig{
	VotingSystems: []VotingSystemField{
		VotingSystemField{
			Name:                    "Unanimous Resolution",
			VoteType:                "A",
			TallyLogic:              0,
			ThresholdPercentage:     100,
			VoteMultiplierPermitted: false,
			HolderProposalFee:       1000000,
		},
		VotingSystemField{
			Name:                    "Ordinary Resolution",
			VoteType:                "R",
			TallyLogic:              0,
			ThresholdPercentage:     50,
			VoteMultiplierPermitted: true,
			HolderProposalFee:       1000000,
		},
		VotingSystemField{
			Name:                    "Special Resolution",
			VoteType:                "R",
			TallyLogic:              0,
			ThresholdPercentage:     75,
			VoteMultiplierPermitted: false,
			HolderProposalFee:       1000000,
		},
		VotingSystemField{
			Name:                    "General Board Resolution",
			VoteType:                "R",
			TallyLogic:              0,
			ThresholdPercentage:     50,
			VoteMultiplierPermitted: false,
			HolderProposalFee:       1000000,
		},
		VotingSystemField{
			Name:                    "Special Board Resolution",
			VoteType:                "A",
			TallyLogic:              0,
			ThresholdPercentage:     67,
			VoteMultiplierPermitted: false,
			HolderProposalFee:       1000000,
		},
	},
	ContractPermissions: permissions.Permissions{
		permissions.Permission{ // Administration Proposal Unanimous Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{true, false, false, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{1}, // [ContractName]
				permissions.FieldIndexPath{2}, // [BodyOfAgreementType]
				permissions.FieldIndexPath{3}, // [BodyOfAgreement]
				permissions.FieldIndexPath{4}, // [ContractType]
				permissions.FieldIndexPath{5}, // [SupportingDocs]
				permissions.FieldIndexPath{6}, // [GoverningLaw]
				permissions.FieldIndexPath{7}, // [Jurisdiction]
			},
		},
		permissions.Permission{ // General Board Resolution
			Permitted:              false,
			AdministrationProposal: false,
			HolderProposal:         false,
			AdministrativeMatter:   true,
			VotingSystemsAllowed:   []bool{false, false, false, true, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{8},         // [ContractExpiration]
				permissions.FieldIndexPath{10, 3},     // [Issuer LEI]
				permissions.FieldIndexPath{10, 4},     // [Issuer UnitNumber]
				permissions.FieldIndexPath{10, 5},     // [Issuer BuildingNumber]
				permissions.FieldIndexPath{10, 6},     // [Issuer Street]
				permissions.FieldIndexPath{10, 7},     // [Issuer SuburbCity]
				permissions.FieldIndexPath{10, 8},     // [Issuer TerritoryStateProvinceCode]
				permissions.FieldIndexPath{10, 14, 1}, // [Issuer Management Type]
				permissions.FieldIndexPath{10, 14, 2}, // [Issuer Management Name]
				permissions.FieldIndexPath{18, 0, 6},  // [VotingSystems 0 HolderProposalFee]
				permissions.FieldIndexPath{23, 0, 1},  // [Oracles 0 Entity]
				permissions.FieldIndexPath{23, 0, 2},  // [Oracles 0 URL]
				permissions.FieldIndexPath{23, 0, 3},  // [Oracles 0 PublicKey]
				permissions.FieldIndexPath{28},        // [Services]
			},
		},
		permissions.Permission{ // Administrator Direct
			Permitted:              true,
			AdministrationProposal: false,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, false, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{9},      // [ContractURI]
				permissions.FieldIndexPath{10, 10}, // [Issuer PostalZIPCode]
				permissions.FieldIndexPath{10, 11}, // [Issuer EmailAddress]
				permissions.FieldIndexPath{10, 12}, // [Issuer PhoneNumber]
				permissions.FieldIndexPath{11},     // [IssuerLogoURL]
				permissions.FieldIndexPath{13},     // [ContractOperator]
				permissions.FieldIndexPath{29},     // [AdminIdentityCertificates]
				permissions.FieldIndexPath{17},     // [ContractFee]
			},
		},
		permissions.Permission{ // Administrator Proposal Ordinary Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, true, false, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{10, 1}, // [Issuer Name]
				permissions.FieldIndexPath{10, 2}, // [Issuer Type]
				permissions.FieldIndexPath{25},    // [EntityContract]
			},
		},
		permissions.Permission{ // Administrator Proposal Special Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, true, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{10, 9}, // [Issuer CountryCode]
				permissions.FieldIndexPath{20},    // [RestrictedQtyAssets]
			},
		},
		permissions.Permission{ // Administrator/Holder Proposal Special Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         true,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, true, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{10, 13, 1}, // [Issuer Administration Type]
				permissions.FieldIndexPath{10, 13, 2}, // [Issuer Administration Name]
				permissions.FieldIndexPath{21},        // [AdministrationProposal]
				permissions.FieldIndexPath{22},        // [HolderProposal]
			},
		},
		permissions.Permission{ // Administrator/Holder Proposal Unanimous Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         true,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{true, false, false, false, false},
			Fields: []permissions.FieldIndexPath{
				permissions.FieldIndexPath{18, 0, 1}, // [VotingSystems 0 Name]
				permissions.FieldIndexPath{18, 0, 2}, // [VotingSystems 0 VoteType]
				permissions.FieldIndexPath{18, 0, 3}, // [VotingSystems 0 TallyLogic]
				permissions.FieldIndexPath{18, 0, 4}, // [VotingSystems 0 ThresholdPercentage]
				permissions.FieldIndexPath{18, 0, 5}, // [VotingSystems 0 VoteMultiplierPermitted]
				permissions.FieldIndexPath{19},       // [ContractPermissions]
			},
		},
	},
	AssetPermissions: map[string]permissions.Permissions{
		"CCY": permissions.Permissions{
			permissions.Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},  // [AssetPermissions]
					permissions.FieldIndexPath{13}, // [TradeRestrictions]
					permissions.FieldIndexPath{4},  // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{5},  // [VotingRights]
					permissions.FieldIndexPath{6},  // [VoteMultiplier]
					permissions.FieldIndexPath{7},  // [AdministrationProposal]
					permissions.FieldIndexPath{8},  // [HolderProposal]
					permissions.FieldIndexPath{9},  // [AssetModificationGovernance]
				},
			},
			permissions.Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{10}, // [AuthorizedTokenQty]
				},
			},
			permissions.Permission{ // Administrator/Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{12, 1}, // [AssetPayload CurrencyCode]
					permissions.FieldIndexPath{12, 2}, // [AssetPayload MonetaryAuthority]
					permissions.FieldIndexPath{12, 3}, // [AssetPayload Description]
				},
			},
		},
		"CHP": permissions.Permissions{
			permissions.Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},     // [AssetPermissions]
					permissions.FieldIndexPath{12, 7}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{13},    // [TradeRestrictions]
					permissions.FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{5},     // [VotingRights]
					permissions.FieldIndexPath{6},     // [VoteMultiplier]
					permissions.FieldIndexPath{7},     // [AdministrationProposal]
					permissions.FieldIndexPath{8},     // [HolderProposal]
					permissions.FieldIndexPath{9},     // [AssetModificationGovernance]
				},
			},
			permissions.Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{10},    // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 9}, // [AssetPayload FaceValue]
					permissions.FieldIndexPath{12, 2}, // [AssetPayload UseType]
					permissions.FieldIndexPath{12, 3}, // [AssetPayload AgeRestriction]
					permissions.FieldIndexPath{12, 4}, // [AssetPayload ValidFrom]
					permissions.FieldIndexPath{12, 5}, // [AssetPayload ExpirationTimestamp]
				},
			},
		},
		"COU": permissions.Permissions{
			permissions.Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},     // [AssetPermissions]
					permissions.FieldIndexPath{12, 8}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{13},    // [TradeRestrictions]
					permissions.FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{5},     // [VotingRights]
					permissions.FieldIndexPath{6},     // [VoteMultiplier]
					permissions.FieldIndexPath{7},     // [AdministrationProposal]
					permissions.FieldIndexPath{8},     // [HolderProposal]
					permissions.FieldIndexPath{9},     // [AssetModificationGovernance]
				},
			},
			permissions.Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{10},     // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 1},  // [AssetPayload RedeemingEntity]
					permissions.FieldIndexPath{12, 3},  // [AssetPayload ExpirationTimestamp]
					permissions.FieldIndexPath{12, 9},  // [AssetPayload FaceValue]
					permissions.FieldIndexPath{12, 11}, // [AssetPayload Details]
				},
			},
		},
		"LOY": permissions.Permissions{
			permissions.Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},     // [AssetPermissions]
					permissions.FieldIndexPath{12, 6}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{13},    // [TradeRestrictions]
					permissions.FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{5},     // [VotingRights]
					permissions.FieldIndexPath{6},     // [VoteMultiplier]
					permissions.FieldIndexPath{7},     // [AdministrationProposal]
					permissions.FieldIndexPath{8},     // [HolderProposal]
					permissions.FieldIndexPath{9},     // [AssetModificationGovernance]
					permissions.FieldIndexPath{12, 1}, // [AssetPayload AgeRestriction]
					permissions.FieldIndexPath{12, 2}, // [AssetPayload ProgramName]
				},
			},
			permissions.Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{10},    // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 4}, // [AssetPayload ExpirationTimestamp]
					permissions.FieldIndexPath{12, 5}, // [AssetPayload Details]
				},
			},
		},
		"MBR": permissions.Permissions{
			permissions.Permission{ // Administrator/Holder Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},     // [AssetPermissions]
					permissions.FieldIndexPath{12, 9}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{5},     // [VotingRights]
					permissions.FieldIndexPath{6},     // [VoteMultiplier]
					permissions.FieldIndexPath{7},     // [AdministrationProposal]
					permissions.FieldIndexPath{8},     // [HolderProposal]
					permissions.FieldIndexPath{10},    // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 1}, // [AssetPayload AgeRestriction]
					permissions.FieldIndexPath{12, 2}, // [AssetPayload ValidFrom]
				},
			},
			permissions.Permission{ // Administrator Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{13},    // [TradeRestrictions]
					permissions.FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{9},     // [AssetModificationGovernance]
					permissions.FieldIndexPath{12, 3}, // [AssetPayload ExpirationTimestamp]
					permissions.FieldIndexPath{12, 4}, // [AssetPayload ID]
					permissions.FieldIndexPath{12, 5}, // [AssetPayload MembershipClass]
					permissions.FieldIndexPath{12, 6}, // [AssetPayload RoleType]
					permissions.FieldIndexPath{12, 7}, // [AssetPayload MembershipType]
				},
			},
			permissions.Permission{ // Administrative Vote General Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, true, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{12, 8}, // [AssetPayload Description]
				},
			},
		},
		"SHC": permissions.Permissions{
			permissions.Permission{ // Administrator/Holder Special Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, true, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},     // [AssetPermissions]
					permissions.FieldIndexPath{12, 4}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{5},     // [VotingRights]
					permissions.FieldIndexPath{6},     // [VoteMultiplier]
					permissions.FieldIndexPath{7},     // [AdministrationProposal]
					permissions.FieldIndexPath{8},     // [HolderProposal]
					permissions.FieldIndexPath{10},    // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 1}, // [AssetPayload Ticker]
					permissions.FieldIndexPath{12, 2}, // [AssetPayload ISIN]
					permissions.FieldIndexPath{12, 3}, // [AssetPayload Description]
				},
			},
			permissions.Permission{ // Administrator Special Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, true, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{13}, // [TradeRestrictions]
					permissions.FieldIndexPath{4},  // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{9},  // [AssetModificationGovernance]
				},
			},
		},
		"TIC": permissions.Permissions{
			permissions.Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{1},      // [AssetPermissions]
					permissions.FieldIndexPath{12, 11}, // [AssetPayload TransfersPermitted]
					permissions.FieldIndexPath{13},     // [TradeRestrictions]
					permissions.FieldIndexPath{4},      // [EnforcementOrdersPermitted]
					permissions.FieldIndexPath{5},      // [VotingRights]
					permissions.FieldIndexPath{6},      // [VoteMultiplier]
					permissions.FieldIndexPath{7},      // [AdministrationProposal]
					permissions.FieldIndexPath{8},      // [HolderProposal]
					permissions.FieldIndexPath{9},      // [AssetModificationGovernance]
				},
			},
			permissions.Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []permissions.FieldIndexPath{
					permissions.FieldIndexPath{10},     // [AuthorizedTokenQty]
					permissions.FieldIndexPath{12, 1},  // [AssetPayload AgeRestriction]
					permissions.FieldIndexPath{12, 3},  // [AssetPayload Venue]
					permissions.FieldIndexPath{12, 5},  // [AssetPayload Area]
					permissions.FieldIndexPath{12, 13}, // [AssetPayload Section]
					permissions.FieldIndexPath{12, 14}, // [AssetPayload Row]
					permissions.FieldIndexPath{12, 6},  // [AssetPayload Seat]
					permissions.FieldIndexPath{12, 7},  // [AssetPayload EventStartTimestamp]
					permissions.FieldIndexPath{12, 15}, // [AssetPayload EventEndTimestamp]
					permissions.FieldIndexPath{12, 12}, // [AssetPayload Details]
				},
			},
		},
	},
}
