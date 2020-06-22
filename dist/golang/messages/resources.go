package messages

/************************************ Entities ************************************/
const (
	// Unspecified -
	EntitiesUnspecified = ""

	// Individual -
	EntitiesIndividual = "I"

	// PublicCompany -
	EntitiesPublicCompany = "P"

	// PrivateCompany -
	EntitiesPrivateCompany = "C"

	// LimitedPartnership -
	EntitiesLimitedPartnership = "L"

	// UnlimitedPartnership -
	EntitiesUnlimitedPartnership = "X"

	// SoleProprietorship -
	EntitiesSoleProprietorship = "T"

	// StatutoryCompany -
	EntitiesStatutoryCompany = "S"

	// NonProfitOrganization -
	EntitiesNonProfitOrganization = "O"

	// NationState -
	EntitiesNationState = "N"

	// GovernmentAgency -
	EntitiesGovernmentAgency = "G"

	// UnitTrust -
	EntitiesUnitTrust = "U"

	// DiscretionaryTrust -
	EntitiesDiscretionaryTrust = "D"
)

type EntitiesCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// EntitiesData holds a mapping of Entities codes.
func EntitiesData(code string) *EntitiesCode {
	switch code {
	case EntitiesUnspecified:
		return &EntitiesCode{
			Name:        "Unspecified",
			Label:       "Unspecified",
			Description: "",
			MetaData:    `{}`,
		}
	case EntitiesIndividual:
		return &EntitiesCode{
			Name:        "Individual",
			Label:       "Individual",
			Description: "",
			MetaData:    `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"legalGuardian":[],"principal":null},"associates":{"accountant":[],"advisor":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"friend":[],"governmentAgency":[],"suppliers":[]},"managers":null,"members":{"principal":null}},"type":"Legal"}`,
		}
	case EntitiesPublicCompany:
		return &EntitiesCode{
			Name:        "PublicCompany",
			Label:       "Public Company Limited by Shares",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		}
	case EntitiesPrivateCompany:
		return &EntitiesCode{
			Name:        "PrivateCompany",
			Label:       "Private Company Limited by Shares",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		}
	case EntitiesLimitedPartnership:
		return &EntitiesCode{
			Name:        "LimitedPartnership",
			Label:       "Limited Partnership",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		}
	case EntitiesUnlimitedPartnership:
		return &EntitiesCode{
			Name:        "UnlimitedPartnership",
			Label:       "Unlimited Partnership",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"contractor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		}
	case EntitiesSoleProprietorship:
		return &EntitiesCode{
			Name:        "SoleProprietorship",
			Label:       "Sole Proprietorship",
			Description: "",
			MetaData:    `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"proprietor":null},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"proprietor":null}},"type":"Ownership"}`,
		}
	case EntitiesStatutoryCompany:
		return &EntitiesCode{
			Name:        "StatutoryCompany",
			Label:       "Statutory Company",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"nationState":null}},"type":"Legal"}`,
		}
	case EntitiesNonProfitOrganization:
		return &EntitiesCode{
			Name:        "NonProfitOrganization",
			Label:       "Non-Profit Organization",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"organizationConstitution":null},"role":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null}},"type":"Legal"}`,
		}
	case EntitiesNationState:
		return &EntitiesCode{
			Name:        "NationState",
			Label:       "Nation State",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"nationalConstitution":null},"role":{"administrators":null,"associates":null,"collaborators":null,"managers":null,"members":{"citizen":[]}},"type":"Legal"}`,
		}
	case EntitiesGovernmentAgency:
		return &EntitiesCode{
			Name:        "GovernmentAgency",
			Label:       "Government Agency",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"charter":null},"role":{"administrators":null,"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":null},"type":"Legal"}`,
		}
	case EntitiesUnitTrust:
		return &EntitiesCode{
			Name:        "UnitTrust",
			Label:       "Unit Trust",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"unitholder":[]}},"type":"Ownership"}`,
		}
	case EntitiesDiscretionaryTrust:
		return &EntitiesCode{
			Name:        "DiscretionaryTrust",
			Label:       "Discretionary Trust",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"beneficiary":[]}},"type":"Ownership"}`,
		}
	default:
		return nil
	}
}

// EntitiesMap returns a mapping of Entities objects with the code as the key.
func EntitiesMap() map[string]*EntitiesCode {
	return map[string]*EntitiesCode{
		EntitiesUnspecified: &EntitiesCode{
			Name:        "Unspecified",
			Label:       "Unspecified",
			Description: "",
			MetaData:    `{}`,
		},
		EntitiesIndividual: &EntitiesCode{
			Name:        "Individual",
			Label:       "Individual",
			Description: "",
			MetaData:    `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"legalGuardian":[],"principal":null},"associates":{"accountant":[],"advisor":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"friend":[],"governmentAgency":[],"suppliers":[]},"managers":null,"members":{"principal":null}},"type":"Legal"}`,
		},
		EntitiesPublicCompany: &EntitiesCode{
			Name:        "PublicCompany",
			Label:       "Public Company Limited by Shares",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		},
		EntitiesPrivateCompany: &EntitiesCode{
			Name:        "PrivateCompany",
			Label:       "Private Company Limited by Shares",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		},
		EntitiesLimitedPartnership: &EntitiesCode{
			Name:        "LimitedPartnership",
			Label:       "Limited Partnership",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		},
		EntitiesUnlimitedPartnership: &EntitiesCode{
			Name:        "UnlimitedPartnership",
			Label:       "Unlimited Partnership",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"contractor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		},
		EntitiesSoleProprietorship: &EntitiesCode{
			Name:        "SoleProprietorship",
			Label:       "Sole Proprietorship",
			Description: "",
			MetaData:    `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"proprietor":null},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"proprietor":null}},"type":"Ownership"}`,
		},
		EntitiesStatutoryCompany: &EntitiesCode{
			Name:        "StatutoryCompany",
			Label:       "Statutory Company",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"nationState":null}},"type":"Legal"}`,
		},
		EntitiesNonProfitOrganization: &EntitiesCode{
			Name:        "NonProfitOrganization",
			Label:       "Non-Profit Organization",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"organizationConstitution":null},"role":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null}},"type":"Legal"}`,
		},
		EntitiesNationState: &EntitiesCode{
			Name:        "NationState",
			Label:       "Nation State",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"nationalConstitution":null},"role":{"administrators":null,"associates":null,"collaborators":null,"managers":null,"members":{"citizen":[]}},"type":"Legal"}`,
		},
		EntitiesGovernmentAgency: &EntitiesCode{
			Name:        "GovernmentAgency",
			Label:       "Government Agency",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"charter":null},"role":{"administrators":null,"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":null},"type":"Legal"}`,
		},
		EntitiesUnitTrust: &EntitiesCode{
			Name:        "UnitTrust",
			Label:       "Unit Trust",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"unitholder":[]}},"type":"Ownership"}`,
		},
		EntitiesDiscretionaryTrust: &EntitiesCode{
			Name:        "DiscretionaryTrust",
			Label:       "Discretionary Trust",
			Description: "",
			MetaData:    `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"beneficiary":[]}},"type":"Ownership"}`,
		},
	}
}

/************************************ Roles ************************************/
const (
	// Accountant -
	RolesAccountant = 1

	// Advisor -
	RolesAdvisor = 2

	// Agent -
	RolesAgent = 3

	// Beneficiary -
	RolesBeneficiary = 4

	// CEO -
	RolesCEO = 5

	// CFO -
	RolesCFO = 6

	// Chair -
	RolesChair = 7

	// COO -
	RolesCOO = 8

	// CTO -
	RolesCTO = 9

	// Custodian -
	RolesCustodian = 10

	// Director -
	RolesDirector = 11

	// Executive -
	RolesExecutive = 12

	// Lawyer -
	RolesLawyer = 13

	// Legal Guardian -
	RolesLegalGuardian = 14

	// Limited Partner -
	RolesLimitedPartner = 15

	// Manager -
	RolesManager = 16

	// Managing Partner -
	RolesManagingPartner = 17

	// Member - Shareholder
	RolesMember = 18

	// Partner -
	RolesPartner = 19

	// Principal -
	RolesPrincipal = 20

	// Proprietor -
	RolesProprietor = 21

	// Protector -
	RolesProtector = 22

	// Secretary -
	RolesSecretary = 23

	// Settlor -
	RolesSettlor = 24

	// Significant Member - Major Shareholder
	RolesSignificantMember = 25

	// Smart Contract Operator -
	RolesSmartContractOperator = 26

	// Trader -
	RolesTrader = 27

	// Trustee -
	RolesTrustee = 28

	// Unit Holder -
	RolesUnitHolder = 29
)

type RolesCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// RolesData holds a mapping of Roles codes.
func RolesData(code uint32) *RolesCode {
	switch code {
	case RolesAccountant:
		return &RolesCode{
			Name:        "Accountant",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesAdvisor:
		return &RolesCode{
			Name:        "Advisor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesAgent:
		return &RolesCode{
			Name:        "Agent",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesBeneficiary:
		return &RolesCode{
			Name:        "Beneficiary",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesCEO:
		return &RolesCode{
			Name:        "CEO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesCFO:
		return &RolesCode{
			Name:        "CFO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesChair:
		return &RolesCode{
			Name:        "Chair",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesCOO:
		return &RolesCode{
			Name:        "COO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesCTO:
		return &RolesCode{
			Name:        "CTO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesCustodian:
		return &RolesCode{
			Name:        "Custodian",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesDirector:
		return &RolesCode{
			Name:        "Director",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesExecutive:
		return &RolesCode{
			Name:        "Executive",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesLawyer:
		return &RolesCode{
			Name:        "Lawyer",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesLegalGuardian:
		return &RolesCode{
			Name:        "Legal Guardian",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesLimitedPartner:
		return &RolesCode{
			Name:        "Limited Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesManager:
		return &RolesCode{
			Name:        "Manager",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesManagingPartner:
		return &RolesCode{
			Name:        "Managing Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesMember:
		return &RolesCode{
			Name:        "Member",
			Label:       "",
			Description: "Shareholder",
			MetaData:    `{}`,
		}
	case RolesPartner:
		return &RolesCode{
			Name:        "Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesPrincipal:
		return &RolesCode{
			Name:        "Principal",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesProprietor:
		return &RolesCode{
			Name:        "Proprietor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesProtector:
		return &RolesCode{
			Name:        "Protector",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesSecretary:
		return &RolesCode{
			Name:        "Secretary",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesSettlor:
		return &RolesCode{
			Name:        "Settlor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesSignificantMember:
		return &RolesCode{
			Name:        "Significant Member",
			Label:       "",
			Description: "Major Shareholder",
			MetaData:    `{}`,
		}
	case RolesSmartContractOperator:
		return &RolesCode{
			Name:        "Smart Contract Operator",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesTrader:
		return &RolesCode{
			Name:        "Trader",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesTrustee:
		return &RolesCode{
			Name:        "Trustee",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case RolesUnitHolder:
		return &RolesCode{
			Name:        "Unit Holder",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	default:
		return nil
	}
}

// RolesMap returns a mapping of Roles objects with the code as the key.
func RolesMap() map[uint32]*RolesCode {
	return map[uint32]*RolesCode{
		RolesAccountant: &RolesCode{
			Name:        "Accountant",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesAdvisor: &RolesCode{
			Name:        "Advisor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesAgent: &RolesCode{
			Name:        "Agent",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesBeneficiary: &RolesCode{
			Name:        "Beneficiary",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesCEO: &RolesCode{
			Name:        "CEO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesCFO: &RolesCode{
			Name:        "CFO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesChair: &RolesCode{
			Name:        "Chair",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesCOO: &RolesCode{
			Name:        "COO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesCTO: &RolesCode{
			Name:        "CTO",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesCustodian: &RolesCode{
			Name:        "Custodian",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesDirector: &RolesCode{
			Name:        "Director",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesExecutive: &RolesCode{
			Name:        "Executive",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesLawyer: &RolesCode{
			Name:        "Lawyer",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesLegalGuardian: &RolesCode{
			Name:        "Legal Guardian",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesLimitedPartner: &RolesCode{
			Name:        "Limited Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesManager: &RolesCode{
			Name:        "Manager",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesManagingPartner: &RolesCode{
			Name:        "Managing Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesMember: &RolesCode{
			Name:        "Member",
			Label:       "",
			Description: "Shareholder",
			MetaData:    `{}`,
		},
		RolesPartner: &RolesCode{
			Name:        "Partner",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesPrincipal: &RolesCode{
			Name:        "Principal",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesProprietor: &RolesCode{
			Name:        "Proprietor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesProtector: &RolesCode{
			Name:        "Protector",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesSecretary: &RolesCode{
			Name:        "Secretary",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesSettlor: &RolesCode{
			Name:        "Settlor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesSignificantMember: &RolesCode{
			Name:        "Significant Member",
			Label:       "",
			Description: "Major Shareholder",
			MetaData:    `{}`,
		},
		RolesSmartContractOperator: &RolesCode{
			Name:        "Smart Contract Operator",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesTrader: &RolesCode{
			Name:        "Trader",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesTrustee: &RolesCode{
			Name:        "Trustee",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		RolesUnitHolder: &RolesCode{
			Name:        "Unit Holder",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
	}
}

/************************************ Tags ************************************/
const (
	// Housing -
	TagsHousing = 1

	// Utilities -
	TagsUtilities = 2

	// Food -
	TagsFood = 3

	// Medical -
	TagsMedical = 4

	// Financial Services -
	TagsFinancialServices = 5

	// Entertainment -
	TagsEntertainment = 6

	// Sales -
	TagsSales = 7

	// Automotive -
	TagsAutomotive = 8

	// Transportation -
	TagsTransportation = 9

	// Fitness -
	TagsFitness = 10

	// Electricity -
	TagsElectricity = 20

	// Water -
	TagsWater = 21

	// Internet -
	TagsInternet = 22

	// Medicine -
	TagsMedicine = 23

	// Service -
	TagsService = 24

	// Repair -
	TagsRepair = 25

	// Supplies -
	TagsSupplies = 26

	// Parts -
	TagsParts = 27

	// Labor -
	TagsLabor = 28

	// Tip -
	TagsTip = 29

	// Media -
	TagsMedia = 30

	// Music -
	TagsMusic = 40

	// Video -
	TagsVideo = 41

	// Photo -
	TagsPhoto = 42

	// Audio -
	TagsAudio = 43

	// Alcohol -
	TagsAlcohol = 100

	// Tobacco -
	TagsTobacco = 101

	// Discounted -
	TagsDiscounted = 120

	// Promotional -
	TagsPromotional = 121
)

type TagsCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// TagsData holds a mapping of Tags codes.
func TagsData(code uint32) *TagsCode {
	switch code {
	case TagsHousing:
		return &TagsCode{
			Name:        "Housing",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsUtilities:
		return &TagsCode{
			Name:        "Utilities",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsFood:
		return &TagsCode{
			Name:        "Food",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsMedical:
		return &TagsCode{
			Name:        "Medical",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsFinancialServices:
		return &TagsCode{
			Name:        "Financial Services",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsEntertainment:
		return &TagsCode{
			Name:        "Entertainment",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsSales:
		return &TagsCode{
			Name:        "Sales",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsAutomotive:
		return &TagsCode{
			Name:        "Automotive",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsTransportation:
		return &TagsCode{
			Name:        "Transportation",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsFitness:
		return &TagsCode{
			Name:        "Fitness",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsElectricity:
		return &TagsCode{
			Name:        "Electricity",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsWater:
		return &TagsCode{
			Name:        "Water",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsInternet:
		return &TagsCode{
			Name:        "Internet",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsMedicine:
		return &TagsCode{
			Name:        "Medicine",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsService:
		return &TagsCode{
			Name:        "Service",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsRepair:
		return &TagsCode{
			Name:        "Repair",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsSupplies:
		return &TagsCode{
			Name:        "Supplies",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsParts:
		return &TagsCode{
			Name:        "Parts",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsLabor:
		return &TagsCode{
			Name:        "Labor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsTip:
		return &TagsCode{
			Name:        "Tip",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsMedia:
		return &TagsCode{
			Name:        "Media",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsMusic:
		return &TagsCode{
			Name:        "Music",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsVideo:
		return &TagsCode{
			Name:        "Video",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsPhoto:
		return &TagsCode{
			Name:        "Photo",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsAudio:
		return &TagsCode{
			Name:        "Audio",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsAlcohol:
		return &TagsCode{
			Name:        "Alcohol",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsTobacco:
		return &TagsCode{
			Name:        "Tobacco",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsDiscounted:
		return &TagsCode{
			Name:        "Discounted",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	case TagsPromotional:
		return &TagsCode{
			Name:        "Promotional",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		}
	default:
		return nil
	}
}

// TagsMap returns a mapping of Tags objects with the code as the key.
func TagsMap() map[uint32]*TagsCode {
	return map[uint32]*TagsCode{
		TagsHousing: &TagsCode{
			Name:        "Housing",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsUtilities: &TagsCode{
			Name:        "Utilities",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsFood: &TagsCode{
			Name:        "Food",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsMedical: &TagsCode{
			Name:        "Medical",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsFinancialServices: &TagsCode{
			Name:        "Financial Services",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsEntertainment: &TagsCode{
			Name:        "Entertainment",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsSales: &TagsCode{
			Name:        "Sales",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsAutomotive: &TagsCode{
			Name:        "Automotive",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsTransportation: &TagsCode{
			Name:        "Transportation",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsFitness: &TagsCode{
			Name:        "Fitness",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsElectricity: &TagsCode{
			Name:        "Electricity",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsWater: &TagsCode{
			Name:        "Water",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsInternet: &TagsCode{
			Name:        "Internet",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsMedicine: &TagsCode{
			Name:        "Medicine",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsService: &TagsCode{
			Name:        "Service",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsRepair: &TagsCode{
			Name:        "Repair",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsSupplies: &TagsCode{
			Name:        "Supplies",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsParts: &TagsCode{
			Name:        "Parts",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsLabor: &TagsCode{
			Name:        "Labor",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsTip: &TagsCode{
			Name:        "Tip",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsMedia: &TagsCode{
			Name:        "Media",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsMusic: &TagsCode{
			Name:        "Music",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsVideo: &TagsCode{
			Name:        "Video",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsPhoto: &TagsCode{
			Name:        "Photo",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsAudio: &TagsCode{
			Name:        "Audio",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsAlcohol: &TagsCode{
			Name:        "Alcohol",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsTobacco: &TagsCode{
			Name:        "Tobacco",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsDiscounted: &TagsCode{
			Name:        "Discounted",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
		TagsPromotional: &TagsCode{
			Name:        "Promotional",
			Label:       "",
			Description: "",
			MetaData:    `{}`,
		},
	}
}
