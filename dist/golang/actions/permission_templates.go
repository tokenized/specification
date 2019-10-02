package actions

type PermissionConfig struct {
	VotingSystems       []VotingSystemField
	ContractPermissions Permissions
	AssetPermissions    map[string]Permissions
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
	ContractPermissions: Permissions{
		Permission{ // Administration Proposal Unanimous Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{true, false, false, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{1}, // [ContractName]
				FieldIndexPath{2}, // [BodyOfAgreementType]
				FieldIndexPath{3}, // [BodyOfAgreement]
				FieldIndexPath{4}, // [ContractType]
				FieldIndexPath{5}, // [SupportingDocs]
				FieldIndexPath{6}, // [GoverningLaw]
				FieldIndexPath{7}, // [Jurisdiction]
			},
		},
		Permission{ // General Board Resolution
			Permitted:              false,
			AdministrationProposal: false,
			HolderProposal:         false,
			AdministrativeMatter:   true,
			VotingSystemsAllowed:   []bool{false, false, false, true, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{8},         // [ContractExpiration]
				FieldIndexPath{10, 3},     // [Issuer LEI]
				FieldIndexPath{10, 4},     // [Issuer UnitNumber]
				FieldIndexPath{10, 5},     // [Issuer BuildingNumber]
				FieldIndexPath{10, 6},     // [Issuer Street]
				FieldIndexPath{10, 7},     // [Issuer SuburbCity]
				FieldIndexPath{10, 8},     // [Issuer TerritoryStateProvinceCode]
				FieldIndexPath{10, 14, 1}, // [Issuer Management Type]
				FieldIndexPath{10, 14, 2}, // [Issuer Management Name]
				FieldIndexPath{18, 6},     // [VotingSystems HolderProposalFee]
				FieldIndexPath{23, 1},     // [Oracles Name]
				FieldIndexPath{23, 2},     // [Oracles URL]
				FieldIndexPath{23, 3},     // [Oracles PublicKey]
			},
		},
		Permission{ // Administrator Direct
			Permitted:              true,
			AdministrationProposal: false,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, false, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{9},      // [ContractURI]
				FieldIndexPath{10, 10}, // [Issuer PostalZIPCode]
				FieldIndexPath{10, 11}, // [Issuer EmailAddress]
				FieldIndexPath{10, 12}, // [Issuer PhoneNumber]
				FieldIndexPath{11},     // [IssuerLogoURL]
				FieldIndexPath{13},     // [ContractOperator]
				FieldIndexPath{14},     // [AdminOracle]
				FieldIndexPath{15},     // [AdminOracleSignature]
				FieldIndexPath{16},     // [AdminOracleSigBlockHeight]
				FieldIndexPath{17},     // [ContractFee]
			},
		},
		Permission{ // Administrator Proposal Ordinary Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, true, false, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{10, 1}, // [Issuer Name]
				FieldIndexPath{10, 2}, // [Issuer Type]
			},
		},
		Permission{ // Administrator Proposal Special Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         false,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, true, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{10, 9}, // [Issuer CountryCode]
				FieldIndexPath{20},    // [RestrictedQtyAssets]
			},
		},
		Permission{ // Administrator/Holder Proposal Special Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         true,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{false, false, true, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{10, 13, 1}, // [Issuer Administration Type]
				FieldIndexPath{10, 13, 2}, // [Issuer Administration Name]
				FieldIndexPath{21},        // [AdministrationProposal]
				FieldIndexPath{22},        // [HolderProposal]
			},
		},
		Permission{ // Administrator/Holder Proposal Unanimous Resolution
			Permitted:              false,
			AdministrationProposal: true,
			HolderProposal:         true,
			AdministrativeMatter:   false,
			VotingSystemsAllowed:   []bool{true, false, false, false, false},
			Fields: []FieldIndexPath{
				FieldIndexPath{18, 1}, // [VotingSystems Name]
				FieldIndexPath{18, 2}, // [VotingSystems VoteType]
				FieldIndexPath{18, 3}, // [VotingSystems TallyLogic]
				FieldIndexPath{18, 4}, // [VotingSystems ThresholdPercentage]
				FieldIndexPath{18, 5}, // [VotingSystems VoteMultiplierPermitted]
				FieldIndexPath{19},    // [ContractPermissions]
			},
		},
	},
	AssetPermissions: map[string]Permissions{
		"SHC": Permissions{
			Permission{ // Administrator/Holder Special Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, true, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},     // [AssetPermissions]
					FieldIndexPath{2},     // [TransfersPermitted]
					FieldIndexPath{5},     // [VotingRights]
					FieldIndexPath{6},     // [VoteMultiplier]
					FieldIndexPath{7},     // [AdministrationProposal]
					FieldIndexPath{8},     // [HolderProposal]
					FieldIndexPath{10},    // [TokenQty]
					FieldIndexPath{12, 1}, // [AssetPayload TransferLockout]
					FieldIndexPath{12, 2}, // [AssetPayload Ticker]
					FieldIndexPath{12, 3}, // [AssetPayload ISIN]
					FieldIndexPath{12, 4}, // [AssetPayload Description]
				},
			},
			Permission{ // Administrator Special Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, true, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{3}, // [TradeRestrictions]
					FieldIndexPath{4}, // [EnforcementOrdersPermitted]
					FieldIndexPath{9}, // [AssetModificationGovernance]
				},
			},
		},
		"COU": Permissions{
			Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1}, // [AssetPermissions]
					FieldIndexPath{2}, // [TransfersPermitted]
					FieldIndexPath{3}, // [TradeRestrictions]
					FieldIndexPath{4}, // [EnforcementOrdersPermitted]
					FieldIndexPath{5}, // [VotingRights]
					FieldIndexPath{6}, // [VoteMultiplier]
					FieldIndexPath{7}, // [AdministrationProposal]
					FieldIndexPath{8}, // [HolderProposal]
					FieldIndexPath{9}, // [AssetModificationGovernance]
				},
			},
			Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{10},    // [TokenQty]
					FieldIndexPath{12, 1}, // [AssetPayload RedeemingEntity]
					FieldIndexPath{12, 2}, // [AssetPayload IssueDate]
					FieldIndexPath{12, 3}, // [AssetPayload ExpiryDate]
					FieldIndexPath{12, 4}, // [AssetPayload Value]
					FieldIndexPath{12, 5}, // [AssetPayload Currency]
					FieldIndexPath{12, 6}, // [AssetPayload Description]
				},
			},
		},
		"CUR": Permissions{
			Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1}, // [AssetPermissions]
					FieldIndexPath{2}, // [TransfersPermitted]
					FieldIndexPath{3}, // [TradeRestrictions]
					FieldIndexPath{4}, // [EnforcementOrdersPermitted]
					FieldIndexPath{5}, // [VotingRights]
					FieldIndexPath{6}, // [VoteMultiplier]
					FieldIndexPath{7}, // [AdministrationProposal]
					FieldIndexPath{8}, // [HolderProposal]
					FieldIndexPath{9}, // [AssetModificationGovernance]
				},
			},
			Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{10}, // [TokenQty]
				},
			},
			Permission{ // Administrator/Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{12, 1}, // [AssetPayload CurrencyCode]
					FieldIndexPath{12, 2}, // [AssetPayload MonetaryAuthority]
					FieldIndexPath{12, 3}, // [AssetPayload Description]
				},
			},
		},
		"LOY": Permissions{
			Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},     // [AssetPermissions]
					FieldIndexPath{2},     // [TransfersPermitted]
					FieldIndexPath{3},     // [TradeRestrictions]
					FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					FieldIndexPath{5},     // [VotingRights]
					FieldIndexPath{6},     // [VoteMultiplier]
					FieldIndexPath{7},     // [AdministrationProposal]
					FieldIndexPath{8},     // [HolderProposal]
					FieldIndexPath{9},     // [AssetModificationGovernance]
					FieldIndexPath{12, 1}, // [AssetPayload AgeRestriction]
					FieldIndexPath{12, 2}, // [AssetPayload OfferName]
				},
			},
			Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{10},    // [TokenQty]
					FieldIndexPath{12, 3}, // [AssetPayload ValidFrom]
					FieldIndexPath{12, 4}, // [AssetPayload ExpirationTimestamp]
					FieldIndexPath{12, 5}, // [AssetPayload Description]
				},
			},
		},
		"MEM": Permissions{
			Permission{ // Administrator/Holder Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},     // [AssetPermissions]
					FieldIndexPath{2},     // [TransfersPermitted]
					FieldIndexPath{5},     // [VotingRights]
					FieldIndexPath{6},     // [VoteMultiplier]
					FieldIndexPath{7},     // [AdministrationProposal]
					FieldIndexPath{8},     // [HolderProposal]
					FieldIndexPath{10},    // [TokenQty]
					FieldIndexPath{12, 1}, // [AssetPayload AgeRestriction]
					FieldIndexPath{12, 2}, // [AssetPayload ValidFrom]
				},
			},
			Permission{ // Administrator Special Board Resolution
				Permitted:              false,
				AdministrationProposal: true,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{3},     // [TradeRestrictions]
					FieldIndexPath{4},     // [EnforcementOrdersPermitted]
					FieldIndexPath{9},     // [AssetModificationGovernance]
					FieldIndexPath{12, 3}, // [AssetPayload ExpirationTimestamp]
					FieldIndexPath{12, 4}, // [AssetPayload ID]
					FieldIndexPath{12, 5}, // [AssetPayload MembershipClass]
					FieldIndexPath{12, 6}, // [AssetPayload RoleType]
					FieldIndexPath{12, 7}, // [AssetPayload MembershipType]
					FieldIndexPath{12},    // [AssetPayload]
					FieldIndexPath{12},    // [AssetPayload]
				},
			},
			Permission{ // Administrative Vote General Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, true, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{12, 8}, // [AssetPayload Description]
				},
			},
		},
		"TIC": Permissions{
			Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1}, // [AssetPermissions]
					FieldIndexPath{2}, // [TransfersPermitted]
					FieldIndexPath{3}, // [TradeRestrictions]
					FieldIndexPath{4}, // [EnforcementOrdersPermitted]
					FieldIndexPath{5}, // [VotingRights]
					FieldIndexPath{6}, // [VoteMultiplier]
					FieldIndexPath{7}, // [AdministrationProposal]
					FieldIndexPath{8}, // [HolderProposal]
					FieldIndexPath{9}, // [AssetModificationGovernance]
				},
			},
			Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{10},     // [TokenQty]
					FieldIndexPath{12, 1},  // [AssetPayload AgeRestriction]
					FieldIndexPath{12, 2},  // [AssetPayload AdmissionType]
					FieldIndexPath{12, 3},  // [AssetPayload Venue]
					FieldIndexPath{12, 4},  // [AssetPayload Class]
					FieldIndexPath{12, 5},  // [AssetPayload Area]
					FieldIndexPath{12, 6},  // [AssetPayload Seat]
					FieldIndexPath{12, 7},  // [AssetPayload StartTimeDate]
					FieldIndexPath{12, 8},  // [AssetPayload ValidFrom]
					FieldIndexPath{12, 9},  // [AssetPayload ExpirationTimestamp]
					FieldIndexPath{12, 10}, // [AssetPayload Description]
				},
			},
		},
		"CHP": Permissions{
			Permission{ // Administrative Vote Special Board Resolution
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{false, false, false, false, true},
				Fields: []FieldIndexPath{
					FieldIndexPath{1}, // [AssetPermissions]
					FieldIndexPath{2}, // [TransfersPermitted]
					FieldIndexPath{3}, // [TradeRestrictions]
					FieldIndexPath{4}, // [EnforcementOrdersPermitted]
					FieldIndexPath{5}, // [VotingRights]
					FieldIndexPath{6}, // [VoteMultiplier]
					FieldIndexPath{7}, // [AdministrationProposal]
					FieldIndexPath{8}, // [HolderProposal]
					FieldIndexPath{9}, // [AssetModificationGovernance]
				},
			},
			Permission{ // Administrator Direct
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{false, false, false, false, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{10},    // [TokenQty]
					FieldIndexPath{12, 1}, // [AssetPayload CurrencyCode]
					FieldIndexPath{12, 2}, // [AssetPayload UseType]
					FieldIndexPath{12, 3}, // [AssetPayload AgeRestriction]
					FieldIndexPath{12, 4}, // [AssetPayload ValidFrom]
					FieldIndexPath{12, 5}, // [AssetPayload ExpirationTimestamp]
				},
			},
		},
	},
}
