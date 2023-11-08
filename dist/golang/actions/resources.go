package actions

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

/************************************ Polities ************************************/
const (

	// Aaland Islands -
	PolitiesAalandIslands = "ALA"

	// Afghanistan -
	PolitiesAfghanistan = "AFG"

	// Albania -
	PolitiesAlbania = "ALB"

	// Algeria -
	PolitiesAlgeria = "DZA"

	// American Samoa -
	PolitiesAmericanSamoa = "ASM"

	// Andorra -
	PolitiesAndorra = "AND"

	// Angola -
	PolitiesAngola = "AGO"

	// Anguilla -
	PolitiesAnguilla = "AIA"

	// Antarctica -
	PolitiesAntarctica = "ATA"

	// Antigua and Barbuda -
	PolitiesAntiguaandBarbuda = "ATG"

	// Argentina -
	PolitiesArgentina = "ARG"

	// Armenia -
	PolitiesArmenia = "ARM"

	// Aruba -
	PolitiesAruba = "ABW"

	// African Union -
	PolitiesAfricanUnion = "AU"

	// Australia -
	PolitiesAustralia = "AUS"

	// Austria -
	PolitiesAustria = "AUT"

	// Azerbaijan -
	PolitiesAzerbaijan = "AZE"

	// The Bahamas -
	PolitiesTheBahamas = "BHS"

	// Bahrain -
	PolitiesBahrain = "BHR"

	// Bangladesh -
	PolitiesBangladesh = "BGD"

	// Barbados -
	PolitiesBarbados = "BRB"

	// Belarus -
	PolitiesBelarus = "BLR"

	// Belgium -
	PolitiesBelgium = "BEL"

	// Belize -
	PolitiesBelize = "BLZ"

	// Benin -
	PolitiesBenin = "BEN"

	// Bermuda -
	PolitiesBermuda = "BMU"

	// Bhutan -
	PolitiesBhutan = "BTN"

	// Bolivia -
	PolitiesBolivia = "BOL"

	// Bonaire, St Eustasuis and Saba -
	PolitiesBonaireStEustasuisandSaba = "BES"

	// Bosnia and Herzegovina -
	PolitiesBosniaandHerzegovina = "BIH"

	// Botswana -
	PolitiesBotswana = "BWA"

	// Bouvet Island -
	PolitiesBouvetIsland = "BVT"

	// Brazil -
	PolitiesBrazil = "BRA"

	// British Indian Ocean Territory -
	PolitiesBritishIndianOceanTerritory = "IOT"

	// British Virgin Islands -
	PolitiesBritishVirginIslands = "VGB"

	// Brunei -
	PolitiesBrunei = "BRN"

	// Bulgaria -
	PolitiesBulgaria = "BGR"

	// Burkina Faso -
	PolitiesBurkinaFaso = "BFA"

	// Burundi -
	PolitiesBurundi = "BDI"

	// Cambodia -
	PolitiesCambodia = "KHM"

	// Cameroon -
	PolitiesCameroon = "CMR"

	// Canada -
	PolitiesCanada = "CAN"

	// Cape Verde -
	PolitiesCapeVerde = "CPV"

	// Cayman Islands -
	PolitiesCaymanIslands = "CYM"

	// Central African Republic -
	PolitiesCentralAfricanRepublic = "CAF"

	// Chad -
	PolitiesChad = "TCD"

	// Chile -
	PolitiesChile = "CHL"

	// China -
	PolitiesChina = "CHN"

	// Christmas Island -
	PolitiesChristmasIsland = "CXR"

	// Cocos Islands -
	PolitiesCocosIslands = "CCK"

	// Colombia -
	PolitiesColombia = "COL"

	// Comoros -
	PolitiesComoros = "COM"

	// Congo-Brazzaville -
	PolitiesCongoBrazzaville = "COG"

	// Congo-Kinshasa -
	PolitiesCongoKinshasa = "COD"

	// Cook Islands -
	PolitiesCookIslands = "COK"

	// Costa Rica -
	PolitiesCostaRica = "CRI"

	// Croatia -
	PolitiesCroatia = "HRV"

	// Cuba -
	PolitiesCuba = "CUB"

	// Curacao (Netherlands Antilles) -
	PolitiesCuracaoNetherlandsAntilles = "CUW"

	// Cyprus -
	PolitiesCyprus = "CYP"

	// Czech Republic -
	PolitiesCzechRepublic = "CZE"

	// Denmark -
	PolitiesDenmark = "DNK"

	// Djibouti -
	PolitiesDjibouti = "DJI"

	// Dominica -
	PolitiesDominica = "DMA"

	// Dominican Republic -
	PolitiesDominicanRepublic = "DOM"

	// East Timor -
	PolitiesEastTimor = "TLS"

	// Ecuador -
	PolitiesEcuador = "ECU"

	// Egypt -
	PolitiesEgypt = "EGY"

	// El Salvador -
	PolitiesElSalvador = "SLV"

	// Equatorial Guinea -
	PolitiesEquatorialGuinea = "GNQ"

	// Eritrea -
	PolitiesEritrea = "ERI"

	// Estonia -
	PolitiesEstonia = "EST"

	// Ethiopia -
	PolitiesEthiopia = "ETH"

	// European Union -
	PolitiesEuropeanUnion = "EU"

	// Falkland Islands -
	PolitiesFalklandIslands = "FLK"

	// Faroe Islands -
	PolitiesFaroeIslands = "FRO"

	// Fiji -
	PolitiesFiji = "FJI"

	// Finland -
	PolitiesFinland = "FIN"

	// France -
	PolitiesFrance = "FRA"

	// French Guiana -
	PolitiesFrenchGuiana = "GUF"

	// French Polynesia -
	PolitiesFrenchPolynesia = "PYF"

	// French Southern and Antarctic Lands -
	PolitiesFrenchSouthernandAntarcticLands = "ATF"

	// Gabon -
	PolitiesGabon = "GAB"

	// The Gambia -
	PolitiesTheGambia = "GMB"

	// Georgia -
	PolitiesGeorgia = "GEO"

	// Germany -
	PolitiesGermany = "DEU"

	// Ghana -
	PolitiesGhana = "GHA"

	// Gibraltar -
	PolitiesGibraltar = "GIB"

	// Greece -
	PolitiesGreece = "GRC"

	// Greenland -
	PolitiesGreenland = "GRL"

	// Grenada -
	PolitiesGrenada = "GRD"

	// Guadeloupe -
	PolitiesGuadeloupe = "GLP"

	// Guam -
	PolitiesGuam = "GUM"

	// Guatemala -
	PolitiesGuatemala = "GTM"

	// Guernsey -
	PolitiesGuernsey = "GGY"

	// Guinea -
	PolitiesGuinea = "GIN"

	// Guinea-Bissau -
	PolitiesGuineaBissau = "GNB"

	// Guyana -
	PolitiesGuyana = "GUY"

	// Haiti -
	PolitiesHaiti = "HTI"

	// Heard Island and McDonald Islands -
	PolitiesHeardIslandandMcDonaldIslands = "HMD"

	// Honduras -
	PolitiesHonduras = "HND"

	// Hong Kong -
	PolitiesHongKong = "HKG"

	// Hungary -
	PolitiesHungary = "HUN"

	// Iceland -
	PolitiesIceland = "ISL"

	// India -
	PolitiesIndia = "IND"

	// Indonesia -
	PolitiesIndonesia = "IDN"

	// Iran -
	PolitiesIran = "IRN"

	// Iraq -
	PolitiesIraq = "IRQ"

	// Ireland -
	PolitiesIreland = "IRL"

	// Isle of Man -
	PolitiesIsleofMan = "IMN"

	// Israel -
	PolitiesIsrael = "ISR"

	// Italy -
	PolitiesItaly = "ITA"

	// Cote D'Ivoire (Former Ivory Coast) -
	PolitiesCoteDIvoireFormerIvoryCoast = "CIV"

	// Jamaica -
	PolitiesJamaica = "JAM"

	// Japan -
	PolitiesJapan = "JPN"

	// Jersey -
	PolitiesJersey = "JEY"

	// Jordan -
	PolitiesJordan = "JOR"

	// Kazakhstan -
	PolitiesKazakhstan = "KAZ"

	// Kenya -
	PolitiesKenya = "KEN"

	// Kiribati -
	PolitiesKiribati = "KIR"

	// Kuwait -
	PolitiesKuwait = "KWT"

	// Kyrgyzstan -
	PolitiesKyrgyzstan = "KGZ"

	// Laos -
	PolitiesLaos = "LAO"

	// Latvia -
	PolitiesLatvia = "LVA"

	// Lebanon -
	PolitiesLebanon = "LBN"

	// Lesotho -
	PolitiesLesotho = "LSO"

	// Liberia -
	PolitiesLiberia = "LBR"

	// Libya -
	PolitiesLibya = "LBY"

	// Liechtenstein -
	PolitiesLiechtenstein = "LIE"

	// Lithuania -
	PolitiesLithuania = "LTU"

	// Luxembourg -
	PolitiesLuxembourg = "LUX"

	// Macau -
	PolitiesMacau = "MAC"

	// Macedonia -
	PolitiesMacedonia = "MKD"

	// Madagascar -
	PolitiesMadagascar = "MDG"

	// Malawi -
	PolitiesMalawi = "MWI"

	// Malaysia -
	PolitiesMalaysia = "MYS"

	// Maldives -
	PolitiesMaldives = "MDV"

	// Mali -
	PolitiesMali = "MLI"

	// Malta -
	PolitiesMalta = "MLT"

	// Marshall Islands -
	PolitiesMarshallIslands = "MHL"

	// Martinique -
	PolitiesMartinique = "MTQ"

	// Mauritania -
	PolitiesMauritania = "MRT"

	// Mauritius -
	PolitiesMauritius = "MUS"

	// Mayotte -
	PolitiesMayotte = "MYT"

	// Mexico -
	PolitiesMexico = "MEX"

	// Micronesia -
	PolitiesMicronesia = "FSM"

	// Moldova -
	PolitiesMoldova = "MDA"

	// Monaco -
	PolitiesMonaco = "MCO"

	// Mongolia -
	PolitiesMongolia = "MNG"

	// Montenegro -
	PolitiesMontenegro = "MNE"

	// Montserrat -
	PolitiesMontserrat = "MSR"

	// Morocco -
	PolitiesMorocco = "MAR"

	// Mozambique -
	PolitiesMozambique = "MOZ"

	// Myanmar -
	PolitiesMyanmar = "MMR"

	// Namibia -
	PolitiesNamibia = "NAM"

	// Nauru -
	PolitiesNauru = "NRU"

	// Nepal -
	PolitiesNepal = "NPL"

	// Netherlands -
	PolitiesNetherlands = "NLD"

	// New Caledonia -
	PolitiesNewCaledonia = "NCL"

	// New Zealand -
	PolitiesNewZealand = "NZL"

	// Nicaragua -
	PolitiesNicaragua = "NIC"

	// Niger -
	PolitiesNiger = "NER"

	// Nigeria -
	PolitiesNigeria = "NGA"

	// Niue -
	PolitiesNiue = "NIU"

	// Norfolk and Philip Island -
	PolitiesNorfolkandPhilipIsland = "NFK"

	// North Korea -
	PolitiesNorthKorea = "PRK"

	// Northern Mariana Islands -
	PolitiesNorthernMarianaIslands = "MNP"

	// Norway -
	PolitiesNorway = "NOR"

	// Oman -
	PolitiesOman = "OMN"

	// Pakistan -
	PolitiesPakistan = "PAK"

	// Palau -
	PolitiesPalau = "PLW"

	// Palestinian Territory -
	PolitiesPalestinianTerritory = "PSE"

	// Panama -
	PolitiesPanama = "PAN"

	// Papua New Guinea -
	PolitiesPapuaNewGuinea = "PNG"

	// Paraguay -
	PolitiesParaguay = "PRY"

	// Peru -
	PolitiesPeru = "PER"

	// Philippines -
	PolitiesPhilippines = "PHL"

	// Pitcairn Islands -
	PolitiesPitcairnIslands = "PCN"

	// Poland -
	PolitiesPoland = "POL"

	// Portugal -
	PolitiesPortugal = "PRT"

	// Puerto Rico -
	PolitiesPuertoRico = "PRI"

	// Qatar -
	PolitiesQatar = "QAT"

	// Réunion -
	PolitiesRéunion = "REU"

	// Romania -
	PolitiesRomania = "ROU"

	// Russia -
	PolitiesRussia = "RUS"

	// Rwanda -
	PolitiesRwanda = "RWA"

	// Saint Helena, Ascension and Tristan da Cunha -
	PolitiesSaintHelenaAscensionandTristandaCunha = "SHN"

	// Saint Kitts and Nevis -
	PolitiesSaintKittsandNevis = "KNA"

	// Saint Lucia -
	PolitiesSaintLucia = "LCA"

	// Saint Pierre and Miquelon -
	PolitiesSaintPierreandMiquelon = "SPM"

	// Saint Vincent and the Grenadines -
	PolitiesSaintVincentandtheGrenadines = "VCT"

	// Saint-Barthelemy -
	PolitiesSaintBarthelemy = "BLM"

	// Saint-Martin -
	PolitiesSaintMartin = "MAF"

	// Samoa -
	PolitiesSamoa = "WSM"

	// San Marino -
	PolitiesSanMarino = "SMR"

	// São Tomé and Príncipe -
	PolitiesSãoToméandPríncipe = "STP"

	// Saudi Arabia -
	PolitiesSaudiArabia = "SAU"

	// Senegal -
	PolitiesSenegal = "SEN"

	// Serbia -
	PolitiesSerbia = "SRB"

	// Seychelles -
	PolitiesSeychelles = "SYC"

	// Sierra Leone -
	PolitiesSierraLeone = "SLE"

	// Singapore -
	PolitiesSingapore = "SGP"

	// Sint Maarten -
	PolitiesSintMaarten = "SXM"

	// Slovakia -
	PolitiesSlovakia = "SVK"

	// Slovenia -
	PolitiesSlovenia = "SVN"

	// Solomon Islands -
	PolitiesSolomonIslands = "SLB"

	// Somalia -
	PolitiesSomalia = "SOM"

	// South Africa -
	PolitiesSouthAfrica = "ZAF"

	// South Georgia and the South Sandwich Islands -
	PolitiesSouthGeorgiaandtheSouthSandwichIslands = "SGS"

	// South Korea -
	PolitiesSouthKorea = "KOR"

	// South Sudan -
	PolitiesSouthSudan = "SSD"

	// Spain -
	PolitiesSpain = "ESP"

	// Sri Lanka -
	PolitiesSriLanka = "LKA"

	// Sudan -
	PolitiesSudan = "SDN"

	// Suriname -
	PolitiesSuriname = "SUR"

	// Svalbard and Jan Mayen -
	PolitiesSvalbardandJanMayen = "SJM"

	// Swaziland -
	PolitiesSwaziland = "SWZ"

	// Sweden -
	PolitiesSweden = "SWE"

	// Switzerland -
	PolitiesSwitzerland = "CHE"

	// Syria -
	PolitiesSyria = "SYR"

	// Taiwan -
	PolitiesTaiwan = "TWN"

	// Tajikistan -
	PolitiesTajikistan = "TJK"

	// Tanzania -
	PolitiesTanzania = "TZA"

	// Thailand -
	PolitiesThailand = "THA"

	// Togo -
	PolitiesTogo = "TGO"

	// Tokelau -
	PolitiesTokelau = "TKL"

	// Tonga -
	PolitiesTonga = "TON"

	// Trinidad and Tobago -
	PolitiesTrinidadandTobago = "TTO"

	// Tunisia -
	PolitiesTunisia = "TUN"

	// Turkey -
	PolitiesTurkey = "TUR"

	// Turkmenistan -
	PolitiesTurkmenistan = "TKM"

	// Turks and Caicos Islands -
	PolitiesTurksandCaicosIslands = "TCA"

	// Tuvalu -
	PolitiesTuvalu = "TUV"

	// Uganda -
	PolitiesUganda = "UGA"

	// Ukraine -
	PolitiesUkraine = "UKR"

	// United Arab Emirates -
	PolitiesUnitedArabEmirates = "ARE"

	// United Kingdom -
	PolitiesUnitedKingdom = "GBR"

	// United States Minor Outlying Islands -
	PolitiesUnitedStatesMinorOutlyingIslands = "UMI"

	// Uruguay -
	PolitiesUruguay = "URY"

	// US Virgin Islands -
	PolitiesUSVirginIslands = "VIR"

	// USA -
	PolitiesUSA = "USA"

	// Uzbekistan -
	PolitiesUzbekistan = "UZB"

	// Vanuatu -
	PolitiesVanuatu = "VUT"

	// Vatican City -
	PolitiesVaticanCity = "VAT"

	// Venezuela -
	PolitiesVenezuela = "VEN"

	// Vietnam -
	PolitiesVietnam = "VNM"

	// Wallis and Futuna -
	PolitiesWallisandFutuna = "WLF"

	// Western Sahara -
	PolitiesWesternSahara = "ESH"

	// Yemen -
	PolitiesYemen = "YEM"

	// Zambia -
	PolitiesZambia = "ZMB"

	// Zimbabwe -
	PolitiesZimbabwe = "ZWE"
)

type PolitiesCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// PolitiesData holds a mapping of Polities codes.
func PolitiesData(code string) *PolitiesCode {
	switch code {

	case PolitiesAalandIslands:
		return &PolitiesCode{
			Name:        "Aaland Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/52/Flag_of_%C3%85land.svg","phone_code":"+358","states":null}`,
		}

	case PolitiesAfghanistan:
		return &PolitiesCode{
			Name:        "Afghanistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Afghanistan.svg","phone_code":"+93","states":null}`,
		}

	case PolitiesAlbania:
		return &PolitiesCode{
			Name:        "Albania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/36/Flag_of_Albania.svg","phone_code":"+355","states":null}`,
		}

	case PolitiesAlgeria:
		return &PolitiesCode{
			Name:        "Algeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg","phone_code":"+213","states":null}`,
		}

	case PolitiesAmericanSamoa:
		return &PolitiesCode{
			Name:        "American Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/87/Flag_of_American_Samoa.svg","phone_code":"+1684","states":null}`,
		}

	case PolitiesAndorra:
		return &PolitiesCode{
			Name:        "Andorra",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Andorra.svg","phone_code":"+376","states":null}`,
		}

	case PolitiesAngola:
		return &PolitiesCode{
			Name:        "Angola",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9d/Flag_of_Angola.svg","phone_code":"+244","states":null}`,
		}

	case PolitiesAnguilla:
		return &PolitiesCode{
			Name:        "Anguilla",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Anguilla.svg","phone_code":"+1264","states":null}`,
		}

	case PolitiesAntarctica:
		return &PolitiesCode{
			Name:        "Antarctica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+672","states":null}`,
		}

	case PolitiesAntiguaandBarbuda:
		return &PolitiesCode{
			Name:        "Antigua and Barbuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Antigua_and_Barbuda.svg","phone_code":"+1268","states":null}`,
		}

	case PolitiesArgentina:
		return &PolitiesCode{
			Name:        "Argentina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1a/Flag_of_Argentina.svg","phone_code":"+54","states":null}`,
		}

	case PolitiesArmenia:
		return &PolitiesCode{
			Name:        "Armenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2f/Flag_of_Armenia.svg","phone_code":"+374","states":null}`,
		}

	case PolitiesAruba:
		return &PolitiesCode{
			Name:        "Aruba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Aruba.svg","phone_code":"+297","states":null}`,
		}

	case PolitiesAfricanUnion:
		return &PolitiesCode{
			Name:        "African Union",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_the_African_Union.svg","states":{"AGO":null,"BDI":null,"BEN":null,"BFA":null,"BWA":null,"CAF":null,"CIV":null,"CMR":null,"COD":null,"COM":null,"CPV":null,"DJI":null,"DZA":null,"EGY":null,"ERI":null,"ETH":null,"GAB":null,"GHA":null,"GIN":null,"GMB":null,"GNB":null,"GNQ":null,"KEN":null,"LBR":null,"LBY":null,"LSO":null,"SWZ":null,"TCD":null}}`,
		}

	case PolitiesAustralia:
		return &PolitiesCode{
			Name:        "Australia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/b/b9/Flag_of_Australia.svg","gov_fiscal_year":"0701-0630","phone_code":"+61","states":{"AUACT":"Australian Capital Territory","AUCC":"Cocos (Keening) Island","AUCX":"Christmas Island","AUHM":"Heard Island and McDonalds Islands","AUJBT":"Jervis Bay Territory","AUNF":"Norfolk Island","AUNSW":"New South Wales","AUNT":"Northern Territory","AUQLD":"Queensland","AUSA":"South Australia","AUTAS":"Tasmania","AUVIC":"Victoria","AUWA":"Western Australia"}}`,
		}

	case PolitiesAustria:
		return &PolitiesCode{
			Name:        "Austria",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/41/Flag_of_Austria.svg","gov_fiscal_year":"0101-1231","phone_code":"+43","states":null}`,
		}

	case PolitiesAzerbaijan:
		return &PolitiesCode{
			Name:        "Azerbaijan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg","phone_code":"+994","states":null}`,
		}

	case PolitiesTheBahamas:
		return &PolitiesCode{
			Name:        "The Bahamas",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/93/Flag_of_the_Bahamas.svg","phone_code":"+1242","states":null}`,
		}

	case PolitiesBahrain:
		return &PolitiesCode{
			Name:        "Bahrain",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Bahrain.svg","phone_code":"+973","states":null}`,
		}

	case PolitiesBangladesh:
		return &PolitiesCode{
			Name:        "Bangladesh",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/f9/Flag_of_Bangladesh.svg","gov_fiscal_year":"0701-0630","phone_code":"+880","states":null}`,
		}

	case PolitiesBarbados:
		return &PolitiesCode{
			Name:        "Barbados",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ef/Flag_of_Barbados.svg","phone_code":"+1246","states":null}`,
		}

	case PolitiesBelarus:
		return &PolitiesCode{
			Name:        "Belarus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/85/Flag_of_Belarus.svg","phone_code":"+375","states":null}`,
		}

	case PolitiesBelgium:
		return &PolitiesCode{
			Name:        "Belgium",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg","gov_fiscal_year":"0101-1231","phone_code":"+32","states":null}`,
		}

	case PolitiesBelize:
		return &PolitiesCode{
			Name:        "Belize",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e7/Flag_of_Belize.svg","phone_code":"+501","states":null}`,
		}

	case PolitiesBenin:
		return &PolitiesCode{
			Name:        "Benin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Benin.svg","phone_code":"+229","states":null}`,
		}

	case PolitiesBermuda:
		return &PolitiesCode{
			Name:        "Bermuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bermuda.svg","phone_code":"+1441","states":null}`,
		}

	case PolitiesBhutan:
		return &PolitiesCode{
			Name:        "Bhutan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/91/Flag_of_Bhutan.svg","phone_code":"+975","states":null}`,
		}

	case PolitiesBolivia:
		return &PolitiesCode{
			Name:        "Bolivia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/de/Flag_of_Bolivia_%28state%29.svg","phone_code":"+591","states":null}`,
		}

	case PolitiesBonaireStEustasuisandSaba:
		return &PolitiesCode{
			Name:        "Bonaire, St Eustasuis and Saba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","phone_code":"+599","states":null}`,
		}

	case PolitiesBosniaandHerzegovina:
		return &PolitiesCode{
			Name:        "Bosnia and Herzegovina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bosnia_and_Herzegovina.svg","phone_code":"+387","states":null}`,
		}

	case PolitiesBotswana:
		return &PolitiesCode{
			Name:        "Botswana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Botswana.svg","phone_code":"+267","states":null}`,
		}

	case PolitiesBouvetIsland:
		return &PolitiesCode{
			Name:        "Bouvet Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		}

	case PolitiesBrazil:
		return &PolitiesCode{
			Name:        "Brazil",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg","gov_fiscal_year":"0101-1231","phone_code":"+55","states":{"BRAC":"Acre","BRAL":"Alagoas","BRAM":"Amazonas","BRAP":"Amapá","BRBA":"Bahia","BRCE":"Ceará","BRDF":"Federal District","BRES":"Espírito Santo","BRGO":"Goiás","BRMA":"Maranhão","BRMG":"Minas Gerais","BRMS":"Mato Grosso do Sul","BRMT":"Mato Grosso","BRPA":"Pará","BRPB":"Paraíba","BRPE":"Pernambuco","BRPI":"Piauí","BRPR":"Paraná","BRRJ":"Rio de Janeiro","BRRN":"Rio Grande do Norte","BRRO":"Rondônia","BRRR":"Roraima","BRRS":"Rio Grande do Sul","BRSC":"Santa Catarina","BRSE":"Sergipe","BRSP":"São Paulo","BRTO":"Tocantins"}}`,
		}

	case PolitiesBritishIndianOceanTerritory:
		return &PolitiesCode{
			Name:        "British Indian Ocean Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_the_British_Indian_Ocean_Territory.svg","phone_code":"+246","states":null}`,
		}

	case PolitiesBritishVirginIslands:
		return &PolitiesCode{
			Name:        "British Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/42/Flag_of_the_British_Virgin_Islands.svg","phone_code":"+1284","states":null}`,
		}

	case PolitiesBrunei:
		return &PolitiesCode{
			Name:        "Brunei",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Brunei.svg","phone_code":"+673","states":null}`,
		}

	case PolitiesBulgaria:
		return &PolitiesCode{
			Name:        "Bulgaria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Bulgaria.svg","phone_code":"+359","states":null}`,
		}

	case PolitiesBurkinaFaso:
		return &PolitiesCode{
			Name:        "Burkina Faso",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Burkina_Faso.svg","phone_code":"+226","states":null}`,
		}

	case PolitiesBurundi:
		return &PolitiesCode{
			Name:        "Burundi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/50/Flag_of_Burundi.svg","phone_code":"+257","states":null}`,
		}

	case PolitiesCambodia:
		return &PolitiesCode{
			Name:        "Cambodia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_Cambodia.svg","phone_code":"+855","states":null}`,
		}

	case PolitiesCameroon:
		return &PolitiesCode{
			Name:        "Cameroon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Cameroon.svg","phone_code":"+237","states":null}`,
		}

	case PolitiesCanada:
		return &PolitiesCode{
			Name:        "Canada",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg","gov_fiscal_year":"0401-0331","phone_code":"+1","states":{"CAAB":"Alberta","CABC":"British Columbia","CAMB":"Manitoba","CANB":"New Brunswick","CANL":"Newfoundland and Labrador","CANS":"Nova Scotia","CANT":"Northwest Territories","CANU":"Nunavut","CAON":"Ontario","CAPE":"Prince Edward Island","CAQC":"Quebec","CASK":"Saskatchewan","CAYT":"Yukon"}}`,
		}

	case PolitiesCapeVerde:
		return &PolitiesCode{
			Name:        "Cape Verde",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Cape_Verde.svg","phone_code":"+238","states":null}`,
		}

	case PolitiesCaymanIslands:
		return &PolitiesCode{
			Name:        "Cayman Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_the_Cayman_Islands.svg","phone_code":"+1345","states":null}`,
		}

	case PolitiesCentralAfricanRepublic:
		return &PolitiesCode{
			Name:        "Central African Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Central_African_Republic.svg","phone_code":"+236","states":null}`,
		}

	case PolitiesChad:
		return &PolitiesCode{
			Name:        "Chad",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4b/Flag_of_Chad.svg","phone_code":"+235","states":null}`,
		}

	case PolitiesChile:
		return &PolitiesCode{
			Name:        "Chile",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/78/Flag_of_Chile.svg","phone_code":"+56","states":null}`,
		}

	case PolitiesChina:
		return &PolitiesCode{
			Name:        "China",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_the_People%27s_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+86","states":{"CNAH":"Anhui","CNBJ":"Beijing","CNCQ":"Chongqing","CNFJ":"Fujian","CNGD":"Guangdong","CNGS":"Gansu","CNGX":"Guangxi","CNGZ":"Guizhou","CNHA":"Henan","CNHB":"Hubei","CNHE":"Hebei","CNHI":"Hainan","CNHK":"Hong Kong (Xianggang)","CNHL":"Heilongjiang","CNHN":"Hunan","CNJL":"Jilin","CNJS":"Jiangsu","CNJX":"Jiangxi","CNLN":"Liaoning","CNMC":"Macao (Aomen)","CNNM":"Nei Mongol (mn)","CNNX":"Ningxia","CNQH":"Qinghai","CNSC":"Sichuan","CNSD":"Shandong","CNSH":"Shanghai","CNSN":"Shaanxi","CNSX":"Shanxi","CNTJ":"Tianjin","CNTW":"Taiwan","CNXJ":"Xinjiang","CNXZ":"Xizang","CNYN":"Yunnan","CNZJ":"Zhejiang"}}`,
		}

	case PolitiesChristmasIsland:
		return &PolitiesCode{
			Name:        "Christmas Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/67/Flag_of_Christmas_Island.svg","phone_code":"+53","states":null}`,
		}

	case PolitiesCocosIslands:
		return &PolitiesCode{
			Name:        "Cocos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Cocos_%28Keeling%29_Islands.svg","phone_code":"+61","states":null}`,
		}

	case PolitiesColombia:
		return &PolitiesCode{
			Name:        "Colombia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Colombia.svg","phone_code":"+57","states":null}`,
		}

	case PolitiesComoros:
		return &PolitiesCode{
			Name:        "Comoros",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/94/Flag_of_the_Comoros.svg","phone_code":"+269","states":null}`,
		}

	case PolitiesCongoBrazzaville:
		return &PolitiesCode{
			Name:        "Congo-Brazzaville",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_the_Republic_of_the_Congo.svg","phone_code":"+242","states":null}`,
		}

	case PolitiesCongoKinshasa:
		return &PolitiesCode{
			Name:        "Congo-Kinshasa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Democratic_Republic_of_the_Congo.svg","phone_code":"+243","states":null}`,
		}

	case PolitiesCookIslands:
		return &PolitiesCode{
			Name:        "Cook Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/35/Flag_of_the_Cook_Islands.svg","phone_code":"+682","states":null}`,
		}

	case PolitiesCostaRica:
		return &PolitiesCode{
			Name:        "Costa Rica",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"1001-0931","flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Costa_Rica_%28state%29.svg","gov_fiscal_year":"1001-0931","phone_code":"+506","states":null}`,
		}

	case PolitiesCroatia:
		return &PolitiesCode{
			Name:        "Croatia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Croatia.svg","phone_code":"+385","states":null}`,
		}

	case PolitiesCuba:
		return &PolitiesCode{
			Name:        "Cuba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bd/Flag_of_Cuba.svg","phone_code":"+53","states":null}`,
		}

	case PolitiesCuracaoNetherlandsAntilles:
		return &PolitiesCode{
			Name:        "Curacao (Netherlands Antilles)",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_Cura%C3%A7ao.svg","phone_code":"+599","states":null}`,
		}

	case PolitiesCyprus:
		return &PolitiesCode{
			Name:        "Cyprus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Cyprus.svg","phone_code":"+357","states":null}`,
		}

	case PolitiesCzechRepublic:
		return &PolitiesCode{
			Name:        "Czech Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_Czech_Republic.svg","phone_code":"+420","states":null}`,
		}

	case PolitiesDenmark:
		return &PolitiesCode{
			Name:        "Denmark",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Denmark.svg","phone_code":"+45","states":null}`,
		}

	case PolitiesDjibouti:
		return &PolitiesCode{
			Name:        "Djibouti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_Djibouti.svg","phone_code":"+253","states":null}`,
		}

	case PolitiesDominica:
		return &PolitiesCode{
			Name:        "Dominica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c4/Flag_of_Dominica.svg","phone_code":"+1767","states":null}`,
		}

	case PolitiesDominicanRepublic:
		return &PolitiesCode{
			Name:        "Dominican Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_the_Dominican_Republic.svg","phone_code":"+1829","states":null}`,
		}

	case PolitiesEastTimor:
		return &PolitiesCode{
			Name:        "East Timor",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/26/Flag_of_East_Timor.svg","phone_code":"+670","states":null}`,
		}

	case PolitiesEcuador:
		return &PolitiesCode{
			Name:        "Ecuador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e8/Flag_of_Ecuador.svg","phone_code":"+593","states":null}`,
		}

	case PolitiesEgypt:
		return &PolitiesCode{
			Name:        "Egypt",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Egypt.svg","gov_fiscal_year":"0701-0630","phone_code":"+20","states":null}`,
		}

	case PolitiesElSalvador:
		return &PolitiesCode{
			Name:        "El Salvador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_El_Salvador.svg","phone_code":"+503","states":null}`,
		}

	case PolitiesEquatorialGuinea:
		return &PolitiesCode{
			Name:        "Equatorial Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Equatorial_Guinea.svg","phone_code":"+240","states":null}`,
		}

	case PolitiesEritrea:
		return &PolitiesCode{
			Name:        "Eritrea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/29/Flag_of_Eritrea.svg","phone_code":"+291","states":null}`,
		}

	case PolitiesEstonia:
		return &PolitiesCode{
			Name:        "Estonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8f/Flag_of_Estonia.svg","phone_code":"+372","states":null}`,
		}

	case PolitiesEthiopia:
		return &PolitiesCode{
			Name:        "Ethiopia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0708-0707","flag":"https://upload.wikimedia.org/wikipedia/commons/7/71/Flag_of_Ethiopia.svg","gov_fiscal_year":"0708-0707","phone_code":"+251","states":null}`,
		}

	case PolitiesEuropeanUnion:
		return &PolitiesCode{
			Name:        "European Union",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b7/Flag_of_Europe.svg","states":{"AUT":null,"BEL":null,"BGR":null,"CYP":null,"CZE":null,"DEU":null,"DNK":null,"ESP":null,"EST":null,"FIN":null,"FRA":null,"GBR":null,"GRC":null,"HRV":null,"HUN":null,"IRL":null,"ITA":null,"LTU":null,"LUX":null,"LVA":null,"MLT":null,"NLD":null,"POL":null,"PRT":null,"ROU":null,"SVK":null,"SVN":null,"SWE":null}}`,
		}

	case PolitiesFalklandIslands:
		return &PolitiesCode{
			Name:        "Falkland Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_the_Falkland_Islands.svg","phone_code":"+500","states":null}`,
		}

	case PolitiesFaroeIslands:
		return &PolitiesCode{
			Name:        "Faroe Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3c/Flag_of_the_Faroe_Islands.svg","phone_code":"+298","states":null}`,
		}

	case PolitiesFiji:
		return &PolitiesCode{
			Name:        "Fiji",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/ba/Flag_of_Fiji.svg","phone_code":"+679","states":null}`,
		}

	case PolitiesFinland:
		return &PolitiesCode{
			Name:        "Finland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg","phone_code":"+358","states":null}`,
		}

	case PolitiesFrance:
		return &PolitiesCode{
			Name:        "France",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+33","states":null}`,
		}

	case PolitiesFrenchGuiana:
		return &PolitiesCode{
			Name:        "French Guiana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+594","states":null}`,
		}

	case PolitiesFrenchPolynesia:
		return &PolitiesCode{
			Name:        "French Polynesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/db/Flag_of_French_Polynesia.svg","phone_code":"+689","states":null}`,
		}

	case PolitiesFrenchSouthernandAntarcticLands:
		return &PolitiesCode{
			Name:        "French Southern and Antarctic Lands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a7/Flag_of_the_French_Southern_and_Antarctic_Lands.svg","phone_code":"+262","states":null}`,
		}

	case PolitiesGabon:
		return &PolitiesCode{
			Name:        "Gabon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/04/Flag_of_Gabon.svg","phone_code":"+241","states":null}`,
		}

	case PolitiesTheGambia:
		return &PolitiesCode{
			Name:        "The Gambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_The_Gambia.svg","phone_code":"+220","states":null}`,
		}

	case PolitiesGeorgia:
		return &PolitiesCode{
			Name:        "Georgia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Georgia.svg","phone_code":"+995","states":null}`,
		}

	case PolitiesGermany:
		return &PolitiesCode{
			Name:        "Germany",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg","gov_fiscal_year":"0101-1231","phone_code":"+49","states":null}`,
		}

	case PolitiesGhana:
		return &PolitiesCode{
			Name:        "Ghana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Ghana.svg","phone_code":"+233","states":null}`,
		}

	case PolitiesGibraltar:
		return &PolitiesCode{
			Name:        "Gibraltar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/02/Flag_of_Gibraltar.svg","phone_code":"+350","states":null}`,
		}

	case PolitiesGreece:
		return &PolitiesCode{
			Name:        "Greece",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Greece.svg","gov_fiscal_year":"0101-1231","phone_code":"+30","states":null}`,
		}

	case PolitiesGreenland:
		return &PolitiesCode{
			Name:        "Greenland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_Greenland.svg","phone_code":"+299","states":null}`,
		}

	case PolitiesGrenada:
		return &PolitiesCode{
			Name:        "Grenada",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Grenada.svg","phone_code":"+1473","states":null}`,
		}

	case PolitiesGuadeloupe:
		return &PolitiesCode{
			Name:        "Guadeloupe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+590","states":null}`,
		}

	case PolitiesGuam:
		return &PolitiesCode{
			Name:        "Guam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/07/Flag_of_Guam.svg","phone_code":"+1671","states":null}`,
		}

	case PolitiesGuatemala:
		return &PolitiesCode{
			Name:        "Guatemala",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ec/Flag_of_Guatemala.svg","phone_code":"+502","states":null}`,
		}

	case PolitiesGuernsey:
		return &PolitiesCode{
			Name:        "Guernsey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Guernsey.svg","phone_code":"+44","states":null}`,
		}

	case PolitiesGuinea:
		return &PolitiesCode{
			Name:        "Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_Guinea.svg","phone_code":"+224","states":null}`,
		}

	case PolitiesGuineaBissau:
		return &PolitiesCode{
			Name:        "Guinea-Bissau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Guinea-Bissau.svg","phone_code":"+245","states":null}`,
		}

	case PolitiesGuyana:
		return &PolitiesCode{
			Name:        "Guyana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_Guyana.svg","phone_code":"+592","states":null}`,
		}

	case PolitiesHaiti:
		return &PolitiesCode{
			Name:        "Haiti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Haiti.svg","phone_code":"+509","states":null}`,
		}

	case PolitiesHeardIslandandMcDonaldIslands:
		return &PolitiesCode{
			Name:        "Heard Island and McDonald Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg","phone_code":"+672","states":null}`,
		}

	case PolitiesHonduras:
		return &PolitiesCode{
			Name:        "Honduras",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/82/Flag_of_Honduras.svg","phone_code":"+504","states":null}`,
		}

	case PolitiesHongKong:
		return &PolitiesCode{
			Name:        "Hong Kong",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg","gov_fiscal_year":"0401-0331","phone_code":"+852","states":null}`,
		}

	case PolitiesHungary:
		return &PolitiesCode{
			Name:        "Hungary",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c1/Flag_of_Hungary.svg","phone_code":"+36","states":null}`,
		}

	case PolitiesIceland:
		return &PolitiesCode{
			Name:        "Iceland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Iceland.svg","phone_code":"+354","states":null}`,
		}

	case PolitiesIndia:
		return &PolitiesCode{
			Name:        "India",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg","gov_fiscal_year":"0401-0331","phone_code":"+91","states":{"INAN":"Andaman and Nicobar Islands","INAP":"Andhra Pradesh","INAR":"Arunachal Pradesh","INAS":"Assam","INBR":"Bihar","INCH":"Chandigarh","INCT":"Chhattisgarh","INDD":"Daman and Diu","INDL":"Delhi","INDN":"Dadra and Nagar Haveli","INGA":"Goa","INGJ":"Gujarat","INHP":"Himachal Pradesh","INHR":"Haryana","INJH":"Jharkhand","INJK":"Jammu and Kashmir","INKA":"Karnataka","INKL":"Kerala","INLD":"Lakshadweep","INMH":"Maharashtra","INML":"Meghalaya","INMN":"Manipur","INMP":"Madhya Pradesh","INMZ":"Mizoram","INNL":"Nagaland","INOR":"Odisha (formerly known as Orissa)","INPB":"Punjab","INPY":"Puducherry (Pondicherry)","INRJ":"Rajasthan","INSK":"Sikkim","INTG":"Telangana","INTN":"Tamil Nadu","INTR":"Tripura","INUP":"Uttar Pradesh","INUT":"Uttarakhand","INWB":"West Bengal"}}`,
		}

	case PolitiesIndonesia:
		return &PolitiesCode{
			Name:        "Indonesia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg","gov_fiscal_year":"0101-1231","phone_code":"+62","states":null}`,
		}

	case PolitiesIran:
		return &PolitiesCode{
			Name:        "Iran",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ca/Flag_of_Iran.svg","phone_code":"+98","states":null}`,
		}

	case PolitiesIraq:
		return &PolitiesCode{
			Name:        "Iraq",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Iraq.svg","phone_code":"+964","states":null}`,
		}

	case PolitiesIreland:
		return &PolitiesCode{
			Name:        "Ireland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/45/Flag_of_Ireland.svg","gov_fiscal_year":"0101-1231","phone_code":"+353","states":null}`,
		}

	case PolitiesIsleofMan:
		return &PolitiesCode{
			Name:        "Isle of Man",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5d/Flag_of_the_Isle_of_Mann.svg","phone_code":"+44","states":null}`,
		}

	case PolitiesIsrael:
		return &PolitiesCode{
			Name:        "Israel",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Israel.svg","gov_fiscal_year":"0101-1231","phone_code":"+972","states":null}`,
		}

	case PolitiesItaly:
		return &PolitiesCode{
			Name:        "Italy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/0/03/Flag_of_Italy.svg","phone_code":"+39","states":null}`,
		}

	case PolitiesCoteDIvoireFormerIvoryCoast:
		return &PolitiesCode{
			Name:        "Cote D'Ivoire (Former Ivory Coast)",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_C%C3%B4te_d%27Ivoire.svg","phone_code":"+225","states":null}`,
		}

	case PolitiesJamaica:
		return &PolitiesCode{
			Name:        "Jamaica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Jamaica.svg","phone_code":"+1876","states":null}`,
		}

	case PolitiesJapan:
		return &PolitiesCode{
			Name:        "Japan",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg","gov_fiscal_year":"0401-0331","phone_code":"+81","states":null}`,
		}

	case PolitiesJersey:
		return &PolitiesCode{
			Name:        "Jersey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1c/Flag_of_Jersey.svg","phone_code":"+44","states":null}`,
		}

	case PolitiesJordan:
		return &PolitiesCode{
			Name:        "Jordan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c0/Flag_of_Jordan.svg","phone_code":"+962","states":null}`,
		}

	case PolitiesKazakhstan:
		return &PolitiesCode{
			Name:        "Kazakhstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kazakhstan.svg","phone_code":"+7","states":null}`,
		}

	case PolitiesKenya:
		return &PolitiesCode{
			Name:        "Kenya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Kenya.svg","phone_code":"+254","states":null}`,
		}

	case PolitiesKiribati:
		return &PolitiesCode{
			Name:        "Kiribati",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kiribati.svg","phone_code":"+686","states":null}`,
		}

	case PolitiesKuwait:
		return &PolitiesCode{
			Name:        "Kuwait",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/aa/Flag_of_Kuwait.svg","phone_code":"+965","states":null}`,
		}

	case PolitiesKyrgyzstan:
		return &PolitiesCode{
			Name:        "Kyrgyzstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c7/Flag_of_Kyrgyzstan.svg","phone_code":"+996","states":null}`,
		}

	case PolitiesLaos:
		return &PolitiesCode{
			Name:        "Laos",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Laos.svg","phone_code":"+856","states":null}`,
		}

	case PolitiesLatvia:
		return &PolitiesCode{
			Name:        "Latvia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Latvia.svg","phone_code":"+371","states":null}`,
		}

	case PolitiesLebanon:
		return &PolitiesCode{
			Name:        "Lebanon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/59/Flag_of_Lebanon.svg","phone_code":"+961","states":null}`,
		}

	case PolitiesLesotho:
		return &PolitiesCode{
			Name:        "Lesotho",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Lesotho.svg","phone_code":"+266","states":null}`,
		}

	case PolitiesLiberia:
		return &PolitiesCode{
			Name:        "Liberia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b8/Flag_of_Liberia.svg","phone_code":"+231","states":null}`,
		}

	case PolitiesLibya:
		return &PolitiesCode{
			Name:        "Libya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/05/Flag_of_Libya.svg","phone_code":"+218","states":null}`,
		}

	case PolitiesLiechtenstein:
		return &PolitiesCode{
			Name:        "Liechtenstein",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/47/Flag_of_Liechtenstein.svg","phone_code":"+423","states":null}`,
		}

	case PolitiesLithuania:
		return &PolitiesCode{
			Name:        "Lithuania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Lithuania.svg","phone_code":"+370","states":null}`,
		}

	case PolitiesLuxembourg:
		return &PolitiesCode{
			Name:        "Luxembourg",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/da/Flag_of_Luxembourg.svg","phone_code":"+352","states":null}`,
		}

	case PolitiesMacau:
		return &PolitiesCode{
			Name:        "Macau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/63/Flag_of_Macau.svg","phone_code":"+853","states":null}`,
		}

	case PolitiesMacedonia:
		return &PolitiesCode{
			Name:        "Macedonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_Macedonia.svg","phone_code":"+389","states":null}`,
		}

	case PolitiesMadagascar:
		return &PolitiesCode{
			Name:        "Madagascar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Madagascar.svg","phone_code":"+261","states":null}`,
		}

	case PolitiesMalawi:
		return &PolitiesCode{
			Name:        "Malawi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d1/Flag_of_Malawi.svg","phone_code":"+265","states":null}`,
		}

	case PolitiesMalaysia:
		return &PolitiesCode{
			Name:        "Malaysia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/66/Flag_of_Malaysia.svg","phone_code":"+60","states":null}`,
		}

	case PolitiesMaldives:
		return &PolitiesCode{
			Name:        "Maldives",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Maldives.svg","phone_code":"+960","states":null}`,
		}

	case PolitiesMali:
		return &PolitiesCode{
			Name:        "Mali",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Mali.svg","phone_code":"+223","states":null}`,
		}

	case PolitiesMalta:
		return &PolitiesCode{
			Name:        "Malta",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Malta.svg","phone_code":"+356","states":null}`,
		}

	case PolitiesMarshallIslands:
		return &PolitiesCode{
			Name:        "Marshall Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2e/Flag_of_the_Marshall_Islands.svg","phone_code":"+692","states":null}`,
		}

	case PolitiesMartinique:
		return &PolitiesCode{
			Name:        "Martinique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+596","states":null}`,
		}

	case PolitiesMauritania:
		return &PolitiesCode{
			Name:        "Mauritania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/New_National_Flag_of_Mauritania.png","phone_code":"+222","states":null}`,
		}

	case PolitiesMauritius:
		return &PolitiesCode{
			Name:        "Mauritius",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Mauritius.svg","phone_code":"+230","states":null}`,
		}

	case PolitiesMayotte:
		return &PolitiesCode{
			Name:        "Mayotte",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Mayotte_%28local%29.svg","phone_code":"+269","states":null}`,
		}

	case PolitiesMexico:
		return &PolitiesCode{
			Name:        "Mexico",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Mexico.svg","phone_code":"+52","states":{"MXAGU":"Aguascalientes","MXBCN":"Baja California","MXBCS":"Baja California Sur","MXCAM":"Campeche","MXCHH":"Chihuahua","MXCHP":"Chiapas","MXCOA":"Coahuila","MXCOL":"Colima","MXDIF":"Distrito Federal","MXDUR":"Durango","MXGRO":"Guerrero","MXGUA":"Guanajuato","MXHID":"Hidalgo","MXJAL":"Jalisco","MXMEX":"Mexico (Federal District)","MXMIC":"Michoacán","MXMOR":"Morelos","MXNAY":"Nayarit","MXNLE":"Nuevo León","MXOAX":"Oaxaca","MXPUE":"Puebla","MXQUE":"Querétaro","MXROO":"Quintana Roo","MXSIN":"Sinaloa","MXSLP":"San Luis Potosí","MXSON":"Sonora","MXTAB":"Tabasco","MXTAM":"Tamaulipas","MXTLA":"Tlaxcala","MXVER":"Veracruz","MXYUC":"Yucatán","MXZAC":"Zacatecas"}}`,
		}

	case PolitiesMicronesia:
		return &PolitiesCode{
			Name:        "Micronesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+691","states":null}`,
		}

	case PolitiesMoldova:
		return &PolitiesCode{
			Name:        "Moldova",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Moldova.svg","phone_code":"+373","states":null}`,
		}

	case PolitiesMonaco:
		return &PolitiesCode{
			Name:        "Monaco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ea/Flag_of_Monaco.svg","phone_code":"+377","states":null}`,
		}

	case PolitiesMongolia:
		return &PolitiesCode{
			Name:        "Mongolia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4c/Flag_of_Mongolia.svg","phone_code":"+976","states":null}`,
		}

	case PolitiesMontenegro:
		return &PolitiesCode{
			Name:        "Montenegro",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Montenegro.svg","phone_code":"+382","states":null}`,
		}

	case PolitiesMontserrat:
		return &PolitiesCode{
			Name:        "Montserrat",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Montserrat.svg","phone_code":"+1664","states":null}`,
		}

	case PolitiesMorocco:
		return &PolitiesCode{
			Name:        "Morocco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Morocco.svg","phone_code":"+212","states":null}`,
		}

	case PolitiesMozambique:
		return &PolitiesCode{
			Name:        "Mozambique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Mozambique.svg","phone_code":"+258","states":null}`,
		}

	case PolitiesMyanmar:
		return &PolitiesCode{
			Name:        "Myanmar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8c/Flag_of_Myanmar.svg","phone_code":"+95","states":null}`,
		}

	case PolitiesNamibia:
		return &PolitiesCode{
			Name:        "Namibia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Namibia.svg","phone_code":"+264","states":null}`,
		}

	case PolitiesNauru:
		return &PolitiesCode{
			Name:        "Nauru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/Flag_of_Nauru.svg","phone_code":"+674","states":null}`,
		}

	case PolitiesNepal:
		return &PolitiesCode{
			Name:        "Nepal",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9b/Flag_of_Nepal.svg","gov_fiscal_year":"0101-1231","phone_code":"+977","states":null}`,
		}

	case PolitiesNetherlands:
		return &PolitiesCode{
			Name:        "Netherlands",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","gov_fiscal_year":"0101-1231","phone_code":"+31","states":null}`,
		}

	case PolitiesNewCaledonia:
		return &PolitiesCode{
			Name:        "New Caledonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+687","states":null}`,
		}

	case PolitiesNewZealand:
		return &PolitiesCode{
			Name:        "New Zealand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3e/Flag_of_New_Zealand.svg","phone_code":"+64","states":null}`,
		}

	case PolitiesNicaragua:
		return &PolitiesCode{
			Name:        "Nicaragua",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Nicaragua.svg","phone_code":"+505","states":null}`,
		}

	case PolitiesNiger:
		return &PolitiesCode{
			Name:        "Niger",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f4/Flag_of_Niger.svg","phone_code":"+227","states":null}`,
		}

	case PolitiesNigeria:
		return &PolitiesCode{
			Name:        "Nigeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/79/Flag_of_Nigeria.svg","phone_code":"+234","states":null}`,
		}

	case PolitiesNiue:
		return &PolitiesCode{
			Name:        "Niue",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Niue.svg","phone_code":"+683","states":null}`,
		}

	case PolitiesNorfolkandPhilipIsland:
		return &PolitiesCode{
			Name:        "Norfolk and Philip Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Norfolk_Island.svg","phone_code":"+672","states":null}`,
		}

	case PolitiesNorthKorea:
		return &PolitiesCode{
			Name:        "North Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_North_Korea.svg","phone_code":"+850","states":null}`,
		}

	case PolitiesNorthernMarianaIslands:
		return &PolitiesCode{
			Name:        "Northern Mariana Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e0/Flag_of_the_Northern_Mariana_Islands.svg","phone_code":"+1670","states":null}`,
		}

	case PolitiesNorway:
		return &PolitiesCode{
			Name:        "Norway",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		}

	case PolitiesOman:
		return &PolitiesCode{
			Name:        "Oman",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Oman.svg","phone_code":"+968","states":null}`,
		}

	case PolitiesPakistan:
		return &PolitiesCode{
			Name:        "Pakistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/32/Flag_of_Pakistan.svg","phone_code":"+92","states":null}`,
		}

	case PolitiesPalau:
		return &PolitiesCode{
			Name:        "Palau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Palau.svg","phone_code":"+680","states":null}`,
		}

	case PolitiesPalestinianTerritory:
		return &PolitiesCode{
			Name:        "Palestinian Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Palestine.svg","phone_code":"+970","states":null}`,
		}

	case PolitiesPanama:
		return &PolitiesCode{
			Name:        "Panama",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/ab/Flag_of_Panama.svg","phone_code":"+507","states":null}`,
		}

	case PolitiesPapuaNewGuinea:
		return &PolitiesCode{
			Name:        "Papua New Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e3/Flag_of_Papua_New_Guinea.svg","phone_code":"+675","states":null}`,
		}

	case PolitiesParaguay:
		return &PolitiesCode{
			Name:        "Paraguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Paraguay.svg","phone_code":"+595","states":null}`,
		}

	case PolitiesPeru:
		return &PolitiesCode{
			Name:        "Peru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Peru_%28state%29.svg","phone_code":"+51","states":null}`,
		}

	case PolitiesPhilippines:
		return &PolitiesCode{
			Name:        "Philippines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_the_Philippines.svg","phone_code":"+63","states":null}`,
		}

	case PolitiesPitcairnIslands:
		return &PolitiesCode{
			Name:        "Pitcairn Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_the_Pitcairn_Islands.svg","phone_code":"+64","states":null}`,
		}

	case PolitiesPoland:
		return &PolitiesCode{
			Name:        "Poland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg","phone_code":"+48","states":null}`,
		}

	case PolitiesPortugal:
		return &PolitiesCode{
			Name:        "Portugal",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Portugal.svg","gov_fiscal_year":"0101-1231","phone_code":"+351","states":null}`,
		}

	case PolitiesPuertoRico:
		return &PolitiesCode{
			Name:        "Puerto Rico",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/28/Flag_of_Puerto_Rico.svg","phone_code":"+1787","states":null}`,
		}

	case PolitiesQatar:
		return &PolitiesCode{
			Name:        "Qatar",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Qatar.svg","gov_fiscal_year":"0101-1231","phone_code":"+974","states":null}`,
		}

	case PolitiesRéunion:
		return &PolitiesCode{
			Name:        "Réunion",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5a/Flag_of_R%C3%A9union.svg","phone_code":"+262","states":null}`,
		}

	case PolitiesRomania:
		return &PolitiesCode{
			Name:        "Romania",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Romania.svg","gov_fiscal_year":"0101-1231","phone_code":"+40","states":null}`,
		}

	case PolitiesRussia:
		return &PolitiesCode{
			Name:        "Russia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/f/f3/Flag_of_Russia.svg","gov_fiscal_year":"0101-1231","phone_code":"+7","states":{"RUAD":"Adygeya, Respublika","RUAL":"Altay, Respublika","RUALT":"Altayskiy kray","RUAMU":"Amurskaya oblast\"","RUARK":"Arkhangel\"skaya oblast\"","RUAST":"Astrakhanskaya oblast\"","RUBA":"Bashkortostan, Respublika","RUBEL":"Belgorodskaya oblast\"","RUBRY":"Bryanskaya oblast\"","RUBU":"Buryatiya, Respublika","RUCE":"Chechenskaya Respublika","RUCHE":"Chelyabinskaya oblast\"","RUCHU":"Chukotskiy avtonomnyy okrug","RUCU":"Chuvashskaya Respublika","RUDA":"Dagestan, Respublika","RUIN":"Ingushetiya, Respublika","RUIRK":"Irkutskaya oblast\"","RUIVA":"Ivanovskaya oblast\"","RUKAM":"Kamchatskiy kray","RUKB":"Kabardino-Balkarskaya Respublika","RUKC":"Karachayevo-Cherkesskaya Respubl.","RUKDA":"Krasnodarskiy kray","RUKEM":"Kemerovskaya oblast\"","RUKGD":"Kaliningradskaya oblast\"","RUKGN":"Kurganskaya oblast\"","RUKHA":"Khabarovskiy kray","RUKHM":"Khanty-Mansiyskiy avtonomnyy okrug-Yugra","RUKIR":"Kirovskaya oblast\"","RUKK":"Khakasiya, Respublika","RUKL":"Kalmykiya, Respublika","RUKLU":"Kaluzhskaya oblast\"","RUKO":"Komi, Respublika","RUKOS":"Kostromskaya oblast\"","RUKR":"Kareliya, Respublika","RUKRS":"Kurskaya oblast\"","RUKYA":"Krasnoyarskiy kray","RULEN":"Leningradskaya oblast\"","RULIP":"Lipetskaya oblast\"","RUMAG":"Magadanskaya oblast\"","RUME":"Mariy El, Respublika","RUMO":"Mordoviya, Respublika","RUMOS":"Moskovskaya oblast\"","RUMOW":"Moskva (autonomous city)","RUMUR":"Murmanskaya oblast\"","RUNEN":"Nenetskiy avtonomnyy okrug","RUNGR":"Novgorodskaya oblast\"","RUNIZ":"Nizhegorodskaya oblast\"","RUNVS":"Novosibirskaya oblast\"","RUOMS":"Omskaya oblast\"","RUORE":"Orenburgskaya oblast\"","RUORL":"Orlovskaya oblast\"","RUPER":"Permskiy kray","RUPNZ":"Penzenskaya oblast\"","RUPRI":"Primorskiy kray","RUPSK":"Pskovskaya oblast\"","RUROS":"Rostovskaya oblast\"","RURYA":"Ryazanskaya oblast\"","RUSA":"Sakha, Respublika","RUSAK":"Sakhalinskaya oblast\"","RUSAM":"Samarskaya oblast\"","RUSAR":"Saratovskaya oblast\"","RUSE":"Severnaya Osetiya-Alaniya, Respubl.","RUSMO":"Smolenskaya oblast\"","RUSPE":"Sankt-Peterburg (autonomous city)","RUSTA":"Stavropol\"skiy kray","RUSVE":"Sverdlovskaya oblast\"","RUTA":"Tatarstan, Respublika","RUTAM":"Tambovskaya oblast\"","RUTOM":"Tomskaya oblast\"","RUTUL":"Tul\"skaya oblast\"","RUTVE":"Tverskaya oblast\"","RUTY":"Tyva, Respublika","RUTYU":"Tyumenskaya oblast\"","RUUD":"Udmurtskaya Respublika","RUULY":"Ul\"yanovskaya oblast\"","RUVGG":"Volgogradskaya oblast\"","RUVLA":"Vladimirskaya oblast\"","RUVLG":"Vologodskaya oblast\"","RUVOR":"Voronezhskaya oblast\"","RUYAN":"Yamalo-Nenetskiy avtonomnyy okrug","RUYAR":"Yaroslavskaya oblast\"","RUYEV":"Yevreyskaya avtonomnaya oblast\"","RUZAB":"Zabaykal\"skiy kray"}}`,
		}

	case PolitiesRwanda:
		return &PolitiesCode{
			Name:        "Rwanda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Rwanda.svg","phone_code":"+250","states":null}`,
		}

	case PolitiesSaintHelenaAscensionandTristandaCunha:
		return &PolitiesCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","phone_code":"+290","states":null}`,
		}

	case PolitiesSaintKittsandNevis:
		return &PolitiesCode{
			Name:        "Saint Kitts and Nevis",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Saint_Kitts_and_Nevis.svg","phone_code":"+1869","states":null}`,
		}

	case PolitiesSaintLucia:
		return &PolitiesCode{
			Name:        "Saint Lucia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Saint_Lucia.svg","phone_code":"+1758","states":null}`,
		}

	case PolitiesSaintPierreandMiquelon:
		return &PolitiesCode{
			Name:        "Saint Pierre and Miquelon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+508","states":null}`,
		}

	case PolitiesSaintVincentandtheGrenadines:
		return &PolitiesCode{
			Name:        "Saint Vincent and the Grenadines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6d/Flag_of_Saint_Vincent_and_the_Grenadines.svg","phone_code":"+1784","states":null}`,
		}

	case PolitiesSaintBarthelemy:
		return &PolitiesCode{
			Name:        "Saint-Barthelemy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Saint_Barthelemy_%28local%29.svg","phone_code":"+590","states":null}`,
		}

	case PolitiesSaintMartin:
		return &PolitiesCode{
			Name:        "Saint-Martin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+590","states":null}`,
		}

	case PolitiesSamoa:
		return &PolitiesCode{
			Name:        "Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Samoa.svg","phone_code":"+685","states":null}`,
		}

	case PolitiesSanMarino:
		return &PolitiesCode{
			Name:        "San Marino",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_San_Marino.svg","phone_code":"+378","states":null}`,
		}

	case PolitiesSãoToméandPríncipe:
		return &PolitiesCode{
			Name:        "São Tomé and Príncipe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Sao_Tome_and_Principe.svg","phone_code":"+239","states":null}`,
		}

	case PolitiesSaudiArabia:
		return &PolitiesCode{
			Name:        "Saudi Arabia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0d/Flag_of_Saudi_Arabia.svg","phone_code":"+966","states":null}`,
		}

	case PolitiesSenegal:
		return &PolitiesCode{
			Name:        "Senegal",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fd/Flag_of_Senegal.svg","phone_code":"+221","states":null}`,
		}

	case PolitiesSerbia:
		return &PolitiesCode{
			Name:        "Serbia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/ff/Flag_of_Serbia.svg","phone_code":"+381","states":null}`,
		}

	case PolitiesSeychelles:
		return &PolitiesCode{
			Name:        "Seychelles",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Seychelles.svg","phone_code":"+248","states":null}`,
		}

	case PolitiesSierraLeone:
		return &PolitiesCode{
			Name:        "Sierra Leone",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Sierra_Leone.svg","phone_code":"+232","states":null}`,
		}

	case PolitiesSingapore:
		return &PolitiesCode{
			Name:        "Singapore",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg","phone_code":"+65","states":null}`,
		}

	case PolitiesSintMaarten:
		return &PolitiesCode{
			Name:        "Sint Maarten",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Sint_Maarten.svg","phone_code":"+1721","states":null}`,
		}

	case PolitiesSlovakia:
		return &PolitiesCode{
			Name:        "Slovakia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e6/Flag_of_Slovakia.svg","phone_code":"+421","states":null}`,
		}

	case PolitiesSlovenia:
		return &PolitiesCode{
			Name:        "Slovenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f0/Flag_of_Slovenia.svg","phone_code":"+386","states":null}`,
		}

	case PolitiesSolomonIslands:
		return &PolitiesCode{
			Name:        "Solomon Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Solomon_Islands.svg","phone_code":"+677","states":null}`,
		}

	case PolitiesSomalia:
		return &PolitiesCode{
			Name:        "Somalia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_Somalia.svg","phone_code":"+252","states":null}`,
		}

	case PolitiesSouthAfrica:
		return &PolitiesCode{
			Name:        "South Africa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/af/Flag_of_South_Africa.svg","phone_code":"+27","states":null}`,
		}

	case PolitiesSouthGeorgiaandtheSouthSandwichIslands:
		return &PolitiesCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_South_Georgia_and_the_South_Sandwich_Islands.svg","phone_code":"+500","states":null}`,
		}

	case PolitiesSouthKorea:
		return &PolitiesCode{
			Name:        "South Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg","gov_fiscal_year":"0101-1231","phone_code":"+82","states":null}`,
		}

	case PolitiesSouthSudan:
		return &PolitiesCode{
			Name:        "South Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7a/Flag_of_South_Sudan.svg","phone_code":"+211","states":null}`,
		}

	case PolitiesSpain:
		return &PolitiesCode{
			Name:        "Spain",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9a/Flag_of_Spain.svg","gov_fiscal_year":"0101-1231","phone_code":"+34","states":null}`,
		}

	case PolitiesSriLanka:
		return &PolitiesCode{
			Name:        "Sri Lanka",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Sri_Lanka.svg","phone_code":"+94","states":null}`,
		}

	case PolitiesSudan:
		return &PolitiesCode{
			Name:        "Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Sudan.svg","phone_code":"+249","states":null}`,
		}

	case PolitiesSuriname:
		return &PolitiesCode{
			Name:        "Suriname",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/60/Flag_of_Suriname.svg","phone_code":"+597","states":null}`,
		}

	case PolitiesSvalbardandJanMayen:
		return &PolitiesCode{
			Name:        "Svalbard and Jan Mayen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+47","states":null}`,
		}

	case PolitiesSwaziland:
		return &PolitiesCode{
			Name:        "Swaziland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fb/Flag_of_Eswatini.svg","phone_code":"+268","states":null}`,
		}

	case PolitiesSweden:
		return &PolitiesCode{
			Name:        "Sweden",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/4/4c/Flag_of_Sweden.svg","gov_fiscal_year":"0101-1231","phone_code":"+46","states":null}`,
		}

	case PolitiesSwitzerland:
		return &PolitiesCode{
			Name:        "Switzerland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/08/Flag_of_Switzerland_%28Pantone%29.svg","gov_fiscal_year":"0101-1231","phone_code":"+41","states":null}`,
		}

	case PolitiesSyria:
		return &PolitiesCode{
			Name:        "Syria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/53/Flag_of_Syria.svg","phone_code":"+963","states":null}`,
		}

	case PolitiesTaiwan:
		return &PolitiesCode{
			Name:        "Taiwan",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+886","states":null}`,
		}

	case PolitiesTajikistan:
		return &PolitiesCode{
			Name:        "Tajikistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Tajikistan.svg","phone_code":"+992","states":null}`,
		}

	case PolitiesTanzania:
		return &PolitiesCode{
			Name:        "Tanzania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tanzania.svg","phone_code":"+255","states":null}`,
		}

	case PolitiesThailand:
		return &PolitiesCode{
			Name:        "Thailand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a9/Flag_of_Thailand.svg","phone_code":"+66","states":null}`,
		}

	case PolitiesTogo:
		return &PolitiesCode{
			Name:        "Togo",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/68/Flag_of_Togo.svg","phone_code":"+228","states":null}`,
		}

	case PolitiesTokelau:
		return &PolitiesCode{
			Name:        "Tokelau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8e/Flag_of_Tokelau.svg","phone_code":"+690","states":null}`,
		}

	case PolitiesTonga:
		return &PolitiesCode{
			Name:        "Tonga",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Tonga.svg","phone_code":"+676","states":null}`,
		}

	case PolitiesTrinidadandTobago:
		return &PolitiesCode{
			Name:        "Trinidad and Tobago",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Trinidad_and_Tobago.svg","phone_code":"+1868","states":null}`,
		}

	case PolitiesTunisia:
		return &PolitiesCode{
			Name:        "Tunisia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Tunisia.svg","phone_code":"+216","states":null}`,
		}

	case PolitiesTurkey:
		return &PolitiesCode{
			Name:        "Turkey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Turkey.svg","phone_code":"+90","states":null}`,
		}

	case PolitiesTurkmenistan:
		return &PolitiesCode{
			Name:        "Turkmenistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Turkmenistan.svg","phone_code":"+993","states":null}`,
		}

	case PolitiesTurksandCaicosIslands:
		return &PolitiesCode{
			Name:        "Turks and Caicos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_the_Turks_and_Caicos_Islands.svg","phone_code":"+1649","states":null}`,
		}

	case PolitiesTuvalu:
		return &PolitiesCode{
			Name:        "Tuvalu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tuvalu.svg","phone_code":"+688","states":null}`,
		}

	case PolitiesUganda:
		return &PolitiesCode{
			Name:        "Uganda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4e/Flag_of_Uganda.svg","phone_code":"+256","states":null}`,
		}

	case PolitiesUkraine:
		return &PolitiesCode{
			Name:        "Ukraine",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Ukraine.svg","phone_code":"+380","states":null}`,
		}

	case PolitiesUnitedArabEmirates:
		return &PolitiesCode{
			Name:        "United Arab Emirates",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_United_Arab_Emirates.svg","gov_fiscal_year":"0101-1231","phone_code":"+971","states":null}`,
		}

	case PolitiesUnitedKingdom:
		return &PolitiesCode{
			Name:        "United Kingdom",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0406-0405","flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","gov_fiscal_year":"0401-0331","phone_code":"+44","states":null}`,
		}

	case PolitiesUnitedStatesMinorOutlyingIslands:
		return &PolitiesCode{
			Name:        "United States Minor Outlying Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+246","states":null}`,
		}

	case PolitiesUruguay:
		return &PolitiesCode{
			Name:        "Uruguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Uruguay.svg","phone_code":"+598","states":null}`,
		}

	case PolitiesUSVirginIslands:
		return &PolitiesCode{
			Name:        "US Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_the_United_States_Virgin_Islands.svg","phone_code":"+1340","states":null}`,
		}

	case PolitiesUSA:
		return &PolitiesCode{
			Name:        "USA",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+1","states":{"USAK":"Alaska","USAL":"Alabama","USAR":"Arkansas","USAS":"American Samoa","USAZ":"Arizona","USCA":"California","USCO":"Colorado","USCT":"Connecticut","USDC":"Washington D.C.","USDE":"Delaware","USFL":"Florida","USGA":"Georgia","USGU":"Guam","USHI":"Hawaii","USIA":"Iowa","USID":"Idaho","USIL":"Illinois","USIN":"Indiana","USKS":"Kansas","USKY":"Kentucky","USLA":"Louisiana","USMA":"Massachusetts","USMD":"Maryland","USME":"Maine","USMI":"Michigan","USMN":"Minnesota","USMO":"Missouri","USMP":"Northern Mariana Islands","USMS":"Mississippi","USMT":"Montana","USNC":"North Carolina","USND":"North Dakota","USNE":"Nebraska","USNH":"New Hampshire","USNJ":"New Jersey","USNM":"New Mexico","USNV":"Nevada","USNY":"New York","USOH":"Ohio","USOK":"Oklahoma","USOR":"Oregon","USPA":"Pennsylvania","USPR":"Puerto Rico","USRI":"Rhode Island","USSC":"South Carolina","USSD":"South Dakota","USTN":"Tennessee","USTX":"Texas","USUM":"United States Minor Outlying Islands","USUT":"Utah","USVA":"Virginia","USVI":"US Virgin Islands","USVT":"Vermont","USWA":"Washington","USWI":"Wisconsin","USWV":"West Virginia","USWY":"Wyoming"}}`,
		}

	case PolitiesUzbekistan:
		return &PolitiesCode{
			Name:        "Uzbekistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Uzbekistan.svg","phone_code":"+998","states":null}`,
		}

	case PolitiesVanuatu:
		return &PolitiesCode{
			Name:        "Vanuatu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_Vanuatu_%28official%29.svg","phone_code":"+678","states":null}`,
		}

	case PolitiesVaticanCity:
		return &PolitiesCode{
			Name:        "Vatican City",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_the_Vatican_City.svg","phone_code":"+418","states":null}`,
		}

	case PolitiesVenezuela:
		return &PolitiesCode{
			Name:        "Venezuela",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7b/Flag_of_Venezuela_%28state%29.svg","phone_code":"+58","states":null}`,
		}

	case PolitiesVietnam:
		return &PolitiesCode{
			Name:        "Vietnam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Vietnam.svg","phone_code":"+84","states":null}`,
		}

	case PolitiesWallisandFutuna:
		return &PolitiesCode{
			Name:        "Wallis and Futuna",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+681","states":null}`,
		}

	case PolitiesWesternSahara:
		return &PolitiesCode{
			Name:        "Western Sahara",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+212","states":null}`,
		}

	case PolitiesYemen:
		return &PolitiesCode{
			Name:        "Yemen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Yemen.svg","phone_code":"+967","states":null}`,
		}

	case PolitiesZambia:
		return &PolitiesCode{
			Name:        "Zambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/06/Flag_of_Zambia.svg","phone_code":"+260","states":null}`,
		}

	case PolitiesZimbabwe:
		return &PolitiesCode{
			Name:        "Zimbabwe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6a/Flag_of_Zimbabwe.svg","phone_code":"+263","states":null}`,
		}
	default:
		return nil
	}
}

// PolitiesMap returns a mapping of Polities objects with the code as the key.
func PolitiesMap() map[string]*PolitiesCode {
	return map[string]*PolitiesCode{

		PolitiesAalandIslands: &PolitiesCode{
			Name:        "Aaland Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/52/Flag_of_%C3%85land.svg","phone_code":"+358","states":null}`,
		},

		PolitiesAfghanistan: &PolitiesCode{
			Name:        "Afghanistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Afghanistan.svg","phone_code":"+93","states":null}`,
		},

		PolitiesAlbania: &PolitiesCode{
			Name:        "Albania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/36/Flag_of_Albania.svg","phone_code":"+355","states":null}`,
		},

		PolitiesAlgeria: &PolitiesCode{
			Name:        "Algeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg","phone_code":"+213","states":null}`,
		},

		PolitiesAmericanSamoa: &PolitiesCode{
			Name:        "American Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/87/Flag_of_American_Samoa.svg","phone_code":"+1684","states":null}`,
		},

		PolitiesAndorra: &PolitiesCode{
			Name:        "Andorra",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Andorra.svg","phone_code":"+376","states":null}`,
		},

		PolitiesAngola: &PolitiesCode{
			Name:        "Angola",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9d/Flag_of_Angola.svg","phone_code":"+244","states":null}`,
		},

		PolitiesAnguilla: &PolitiesCode{
			Name:        "Anguilla",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Anguilla.svg","phone_code":"+1264","states":null}`,
		},

		PolitiesAntarctica: &PolitiesCode{
			Name:        "Antarctica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+672","states":null}`,
		},

		PolitiesAntiguaandBarbuda: &PolitiesCode{
			Name:        "Antigua and Barbuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Antigua_and_Barbuda.svg","phone_code":"+1268","states":null}`,
		},

		PolitiesArgentina: &PolitiesCode{
			Name:        "Argentina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1a/Flag_of_Argentina.svg","phone_code":"+54","states":null}`,
		},

		PolitiesArmenia: &PolitiesCode{
			Name:        "Armenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2f/Flag_of_Armenia.svg","phone_code":"+374","states":null}`,
		},

		PolitiesAruba: &PolitiesCode{
			Name:        "Aruba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Aruba.svg","phone_code":"+297","states":null}`,
		},

		PolitiesAfricanUnion: &PolitiesCode{
			Name:        "African Union",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_the_African_Union.svg","states":{"AGO":null,"BDI":null,"BEN":null,"BFA":null,"BWA":null,"CAF":null,"CIV":null,"CMR":null,"COD":null,"COM":null,"CPV":null,"DJI":null,"DZA":null,"EGY":null,"ERI":null,"ETH":null,"GAB":null,"GHA":null,"GIN":null,"GMB":null,"GNB":null,"GNQ":null,"KEN":null,"LBR":null,"LBY":null,"LSO":null,"SWZ":null,"TCD":null}}`,
		},

		PolitiesAustralia: &PolitiesCode{
			Name:        "Australia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/b/b9/Flag_of_Australia.svg","gov_fiscal_year":"0701-0630","phone_code":"+61","states":{"AUACT":"Australian Capital Territory","AUCC":"Cocos (Keening) Island","AUCX":"Christmas Island","AUHM":"Heard Island and McDonalds Islands","AUJBT":"Jervis Bay Territory","AUNF":"Norfolk Island","AUNSW":"New South Wales","AUNT":"Northern Territory","AUQLD":"Queensland","AUSA":"South Australia","AUTAS":"Tasmania","AUVIC":"Victoria","AUWA":"Western Australia"}}`,
		},

		PolitiesAustria: &PolitiesCode{
			Name:        "Austria",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/41/Flag_of_Austria.svg","gov_fiscal_year":"0101-1231","phone_code":"+43","states":null}`,
		},

		PolitiesAzerbaijan: &PolitiesCode{
			Name:        "Azerbaijan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg","phone_code":"+994","states":null}`,
		},

		PolitiesTheBahamas: &PolitiesCode{
			Name:        "The Bahamas",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/93/Flag_of_the_Bahamas.svg","phone_code":"+1242","states":null}`,
		},

		PolitiesBahrain: &PolitiesCode{
			Name:        "Bahrain",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Bahrain.svg","phone_code":"+973","states":null}`,
		},

		PolitiesBangladesh: &PolitiesCode{
			Name:        "Bangladesh",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/f9/Flag_of_Bangladesh.svg","gov_fiscal_year":"0701-0630","phone_code":"+880","states":null}`,
		},

		PolitiesBarbados: &PolitiesCode{
			Name:        "Barbados",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ef/Flag_of_Barbados.svg","phone_code":"+1246","states":null}`,
		},

		PolitiesBelarus: &PolitiesCode{
			Name:        "Belarus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/85/Flag_of_Belarus.svg","phone_code":"+375","states":null}`,
		},

		PolitiesBelgium: &PolitiesCode{
			Name:        "Belgium",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg","gov_fiscal_year":"0101-1231","phone_code":"+32","states":null}`,
		},

		PolitiesBelize: &PolitiesCode{
			Name:        "Belize",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e7/Flag_of_Belize.svg","phone_code":"+501","states":null}`,
		},

		PolitiesBenin: &PolitiesCode{
			Name:        "Benin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Benin.svg","phone_code":"+229","states":null}`,
		},

		PolitiesBermuda: &PolitiesCode{
			Name:        "Bermuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bermuda.svg","phone_code":"+1441","states":null}`,
		},

		PolitiesBhutan: &PolitiesCode{
			Name:        "Bhutan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/91/Flag_of_Bhutan.svg","phone_code":"+975","states":null}`,
		},

		PolitiesBolivia: &PolitiesCode{
			Name:        "Bolivia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/de/Flag_of_Bolivia_%28state%29.svg","phone_code":"+591","states":null}`,
		},

		PolitiesBonaireStEustasuisandSaba: &PolitiesCode{
			Name:        "Bonaire, St Eustasuis and Saba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","phone_code":"+599","states":null}`,
		},

		PolitiesBosniaandHerzegovina: &PolitiesCode{
			Name:        "Bosnia and Herzegovina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bosnia_and_Herzegovina.svg","phone_code":"+387","states":null}`,
		},

		PolitiesBotswana: &PolitiesCode{
			Name:        "Botswana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Botswana.svg","phone_code":"+267","states":null}`,
		},

		PolitiesBouvetIsland: &PolitiesCode{
			Name:        "Bouvet Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		},

		PolitiesBrazil: &PolitiesCode{
			Name:        "Brazil",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg","gov_fiscal_year":"0101-1231","phone_code":"+55","states":{"BRAC":"Acre","BRAL":"Alagoas","BRAM":"Amazonas","BRAP":"Amapá","BRBA":"Bahia","BRCE":"Ceará","BRDF":"Federal District","BRES":"Espírito Santo","BRGO":"Goiás","BRMA":"Maranhão","BRMG":"Minas Gerais","BRMS":"Mato Grosso do Sul","BRMT":"Mato Grosso","BRPA":"Pará","BRPB":"Paraíba","BRPE":"Pernambuco","BRPI":"Piauí","BRPR":"Paraná","BRRJ":"Rio de Janeiro","BRRN":"Rio Grande do Norte","BRRO":"Rondônia","BRRR":"Roraima","BRRS":"Rio Grande do Sul","BRSC":"Santa Catarina","BRSE":"Sergipe","BRSP":"São Paulo","BRTO":"Tocantins"}}`,
		},

		PolitiesBritishIndianOceanTerritory: &PolitiesCode{
			Name:        "British Indian Ocean Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_the_British_Indian_Ocean_Territory.svg","phone_code":"+246","states":null}`,
		},

		PolitiesBritishVirginIslands: &PolitiesCode{
			Name:        "British Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/42/Flag_of_the_British_Virgin_Islands.svg","phone_code":"+1284","states":null}`,
		},

		PolitiesBrunei: &PolitiesCode{
			Name:        "Brunei",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Brunei.svg","phone_code":"+673","states":null}`,
		},

		PolitiesBulgaria: &PolitiesCode{
			Name:        "Bulgaria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Bulgaria.svg","phone_code":"+359","states":null}`,
		},

		PolitiesBurkinaFaso: &PolitiesCode{
			Name:        "Burkina Faso",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Burkina_Faso.svg","phone_code":"+226","states":null}`,
		},

		PolitiesBurundi: &PolitiesCode{
			Name:        "Burundi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/50/Flag_of_Burundi.svg","phone_code":"+257","states":null}`,
		},

		PolitiesCambodia: &PolitiesCode{
			Name:        "Cambodia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_Cambodia.svg","phone_code":"+855","states":null}`,
		},

		PolitiesCameroon: &PolitiesCode{
			Name:        "Cameroon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Cameroon.svg","phone_code":"+237","states":null}`,
		},

		PolitiesCanada: &PolitiesCode{
			Name:        "Canada",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg","gov_fiscal_year":"0401-0331","phone_code":"+1","states":{"CAAB":"Alberta","CABC":"British Columbia","CAMB":"Manitoba","CANB":"New Brunswick","CANL":"Newfoundland and Labrador","CANS":"Nova Scotia","CANT":"Northwest Territories","CANU":"Nunavut","CAON":"Ontario","CAPE":"Prince Edward Island","CAQC":"Quebec","CASK":"Saskatchewan","CAYT":"Yukon"}}`,
		},

		PolitiesCapeVerde: &PolitiesCode{
			Name:        "Cape Verde",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Cape_Verde.svg","phone_code":"+238","states":null}`,
		},

		PolitiesCaymanIslands: &PolitiesCode{
			Name:        "Cayman Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_the_Cayman_Islands.svg","phone_code":"+1345","states":null}`,
		},

		PolitiesCentralAfricanRepublic: &PolitiesCode{
			Name:        "Central African Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Central_African_Republic.svg","phone_code":"+236","states":null}`,
		},

		PolitiesChad: &PolitiesCode{
			Name:        "Chad",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4b/Flag_of_Chad.svg","phone_code":"+235","states":null}`,
		},

		PolitiesChile: &PolitiesCode{
			Name:        "Chile",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/78/Flag_of_Chile.svg","phone_code":"+56","states":null}`,
		},

		PolitiesChina: &PolitiesCode{
			Name:        "China",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_the_People%27s_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+86","states":{"CNAH":"Anhui","CNBJ":"Beijing","CNCQ":"Chongqing","CNFJ":"Fujian","CNGD":"Guangdong","CNGS":"Gansu","CNGX":"Guangxi","CNGZ":"Guizhou","CNHA":"Henan","CNHB":"Hubei","CNHE":"Hebei","CNHI":"Hainan","CNHK":"Hong Kong (Xianggang)","CNHL":"Heilongjiang","CNHN":"Hunan","CNJL":"Jilin","CNJS":"Jiangsu","CNJX":"Jiangxi","CNLN":"Liaoning","CNMC":"Macao (Aomen)","CNNM":"Nei Mongol (mn)","CNNX":"Ningxia","CNQH":"Qinghai","CNSC":"Sichuan","CNSD":"Shandong","CNSH":"Shanghai","CNSN":"Shaanxi","CNSX":"Shanxi","CNTJ":"Tianjin","CNTW":"Taiwan","CNXJ":"Xinjiang","CNXZ":"Xizang","CNYN":"Yunnan","CNZJ":"Zhejiang"}}`,
		},

		PolitiesChristmasIsland: &PolitiesCode{
			Name:        "Christmas Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/67/Flag_of_Christmas_Island.svg","phone_code":"+53","states":null}`,
		},

		PolitiesCocosIslands: &PolitiesCode{
			Name:        "Cocos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Cocos_%28Keeling%29_Islands.svg","phone_code":"+61","states":null}`,
		},

		PolitiesColombia: &PolitiesCode{
			Name:        "Colombia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Colombia.svg","phone_code":"+57","states":null}`,
		},

		PolitiesComoros: &PolitiesCode{
			Name:        "Comoros",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/94/Flag_of_the_Comoros.svg","phone_code":"+269","states":null}`,
		},

		PolitiesCongoBrazzaville: &PolitiesCode{
			Name:        "Congo-Brazzaville",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_the_Republic_of_the_Congo.svg","phone_code":"+242","states":null}`,
		},

		PolitiesCongoKinshasa: &PolitiesCode{
			Name:        "Congo-Kinshasa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Democratic_Republic_of_the_Congo.svg","phone_code":"+243","states":null}`,
		},

		PolitiesCookIslands: &PolitiesCode{
			Name:        "Cook Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/35/Flag_of_the_Cook_Islands.svg","phone_code":"+682","states":null}`,
		},

		PolitiesCostaRica: &PolitiesCode{
			Name:        "Costa Rica",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"1001-0931","flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Costa_Rica_%28state%29.svg","gov_fiscal_year":"1001-0931","phone_code":"+506","states":null}`,
		},

		PolitiesCroatia: &PolitiesCode{
			Name:        "Croatia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Croatia.svg","phone_code":"+385","states":null}`,
		},

		PolitiesCuba: &PolitiesCode{
			Name:        "Cuba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bd/Flag_of_Cuba.svg","phone_code":"+53","states":null}`,
		},

		PolitiesCuracaoNetherlandsAntilles: &PolitiesCode{
			Name:        "Curacao (Netherlands Antilles)",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_Cura%C3%A7ao.svg","phone_code":"+599","states":null}`,
		},

		PolitiesCyprus: &PolitiesCode{
			Name:        "Cyprus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Cyprus.svg","phone_code":"+357","states":null}`,
		},

		PolitiesCzechRepublic: &PolitiesCode{
			Name:        "Czech Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_Czech_Republic.svg","phone_code":"+420","states":null}`,
		},

		PolitiesDenmark: &PolitiesCode{
			Name:        "Denmark",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Denmark.svg","phone_code":"+45","states":null}`,
		},

		PolitiesDjibouti: &PolitiesCode{
			Name:        "Djibouti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_Djibouti.svg","phone_code":"+253","states":null}`,
		},

		PolitiesDominica: &PolitiesCode{
			Name:        "Dominica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c4/Flag_of_Dominica.svg","phone_code":"+1767","states":null}`,
		},

		PolitiesDominicanRepublic: &PolitiesCode{
			Name:        "Dominican Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_the_Dominican_Republic.svg","phone_code":"+1829","states":null}`,
		},

		PolitiesEastTimor: &PolitiesCode{
			Name:        "East Timor",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/26/Flag_of_East_Timor.svg","phone_code":"+670","states":null}`,
		},

		PolitiesEcuador: &PolitiesCode{
			Name:        "Ecuador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e8/Flag_of_Ecuador.svg","phone_code":"+593","states":null}`,
		},

		PolitiesEgypt: &PolitiesCode{
			Name:        "Egypt",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Egypt.svg","gov_fiscal_year":"0701-0630","phone_code":"+20","states":null}`,
		},

		PolitiesElSalvador: &PolitiesCode{
			Name:        "El Salvador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_El_Salvador.svg","phone_code":"+503","states":null}`,
		},

		PolitiesEquatorialGuinea: &PolitiesCode{
			Name:        "Equatorial Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Equatorial_Guinea.svg","phone_code":"+240","states":null}`,
		},

		PolitiesEritrea: &PolitiesCode{
			Name:        "Eritrea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/29/Flag_of_Eritrea.svg","phone_code":"+291","states":null}`,
		},

		PolitiesEstonia: &PolitiesCode{
			Name:        "Estonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8f/Flag_of_Estonia.svg","phone_code":"+372","states":null}`,
		},

		PolitiesEthiopia: &PolitiesCode{
			Name:        "Ethiopia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0708-0707","flag":"https://upload.wikimedia.org/wikipedia/commons/7/71/Flag_of_Ethiopia.svg","gov_fiscal_year":"0708-0707","phone_code":"+251","states":null}`,
		},

		PolitiesEuropeanUnion: &PolitiesCode{
			Name:        "European Union",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b7/Flag_of_Europe.svg","states":{"AUT":null,"BEL":null,"BGR":null,"CYP":null,"CZE":null,"DEU":null,"DNK":null,"ESP":null,"EST":null,"FIN":null,"FRA":null,"GBR":null,"GRC":null,"HRV":null,"HUN":null,"IRL":null,"ITA":null,"LTU":null,"LUX":null,"LVA":null,"MLT":null,"NLD":null,"POL":null,"PRT":null,"ROU":null,"SVK":null,"SVN":null,"SWE":null}}`,
		},

		PolitiesFalklandIslands: &PolitiesCode{
			Name:        "Falkland Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_the_Falkland_Islands.svg","phone_code":"+500","states":null}`,
		},

		PolitiesFaroeIslands: &PolitiesCode{
			Name:        "Faroe Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3c/Flag_of_the_Faroe_Islands.svg","phone_code":"+298","states":null}`,
		},

		PolitiesFiji: &PolitiesCode{
			Name:        "Fiji",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/ba/Flag_of_Fiji.svg","phone_code":"+679","states":null}`,
		},

		PolitiesFinland: &PolitiesCode{
			Name:        "Finland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg","phone_code":"+358","states":null}`,
		},

		PolitiesFrance: &PolitiesCode{
			Name:        "France",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+33","states":null}`,
		},

		PolitiesFrenchGuiana: &PolitiesCode{
			Name:        "French Guiana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+594","states":null}`,
		},

		PolitiesFrenchPolynesia: &PolitiesCode{
			Name:        "French Polynesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/db/Flag_of_French_Polynesia.svg","phone_code":"+689","states":null}`,
		},

		PolitiesFrenchSouthernandAntarcticLands: &PolitiesCode{
			Name:        "French Southern and Antarctic Lands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a7/Flag_of_the_French_Southern_and_Antarctic_Lands.svg","phone_code":"+262","states":null}`,
		},

		PolitiesGabon: &PolitiesCode{
			Name:        "Gabon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/04/Flag_of_Gabon.svg","phone_code":"+241","states":null}`,
		},

		PolitiesTheGambia: &PolitiesCode{
			Name:        "The Gambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_The_Gambia.svg","phone_code":"+220","states":null}`,
		},

		PolitiesGeorgia: &PolitiesCode{
			Name:        "Georgia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Georgia.svg","phone_code":"+995","states":null}`,
		},

		PolitiesGermany: &PolitiesCode{
			Name:        "Germany",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg","gov_fiscal_year":"0101-1231","phone_code":"+49","states":null}`,
		},

		PolitiesGhana: &PolitiesCode{
			Name:        "Ghana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Ghana.svg","phone_code":"+233","states":null}`,
		},

		PolitiesGibraltar: &PolitiesCode{
			Name:        "Gibraltar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/02/Flag_of_Gibraltar.svg","phone_code":"+350","states":null}`,
		},

		PolitiesGreece: &PolitiesCode{
			Name:        "Greece",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Greece.svg","gov_fiscal_year":"0101-1231","phone_code":"+30","states":null}`,
		},

		PolitiesGreenland: &PolitiesCode{
			Name:        "Greenland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_Greenland.svg","phone_code":"+299","states":null}`,
		},

		PolitiesGrenada: &PolitiesCode{
			Name:        "Grenada",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Grenada.svg","phone_code":"+1473","states":null}`,
		},

		PolitiesGuadeloupe: &PolitiesCode{
			Name:        "Guadeloupe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+590","states":null}`,
		},

		PolitiesGuam: &PolitiesCode{
			Name:        "Guam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/07/Flag_of_Guam.svg","phone_code":"+1671","states":null}`,
		},

		PolitiesGuatemala: &PolitiesCode{
			Name:        "Guatemala",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ec/Flag_of_Guatemala.svg","phone_code":"+502","states":null}`,
		},

		PolitiesGuernsey: &PolitiesCode{
			Name:        "Guernsey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Guernsey.svg","phone_code":"+44","states":null}`,
		},

		PolitiesGuinea: &PolitiesCode{
			Name:        "Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_Guinea.svg","phone_code":"+224","states":null}`,
		},

		PolitiesGuineaBissau: &PolitiesCode{
			Name:        "Guinea-Bissau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Guinea-Bissau.svg","phone_code":"+245","states":null}`,
		},

		PolitiesGuyana: &PolitiesCode{
			Name:        "Guyana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_Guyana.svg","phone_code":"+592","states":null}`,
		},

		PolitiesHaiti: &PolitiesCode{
			Name:        "Haiti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Haiti.svg","phone_code":"+509","states":null}`,
		},

		PolitiesHeardIslandandMcDonaldIslands: &PolitiesCode{
			Name:        "Heard Island and McDonald Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg","phone_code":"+672","states":null}`,
		},

		PolitiesHonduras: &PolitiesCode{
			Name:        "Honduras",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/82/Flag_of_Honduras.svg","phone_code":"+504","states":null}`,
		},

		PolitiesHongKong: &PolitiesCode{
			Name:        "Hong Kong",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg","gov_fiscal_year":"0401-0331","phone_code":"+852","states":null}`,
		},

		PolitiesHungary: &PolitiesCode{
			Name:        "Hungary",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c1/Flag_of_Hungary.svg","phone_code":"+36","states":null}`,
		},

		PolitiesIceland: &PolitiesCode{
			Name:        "Iceland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Iceland.svg","phone_code":"+354","states":null}`,
		},

		PolitiesIndia: &PolitiesCode{
			Name:        "India",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg","gov_fiscal_year":"0401-0331","phone_code":"+91","states":{"INAN":"Andaman and Nicobar Islands","INAP":"Andhra Pradesh","INAR":"Arunachal Pradesh","INAS":"Assam","INBR":"Bihar","INCH":"Chandigarh","INCT":"Chhattisgarh","INDD":"Daman and Diu","INDL":"Delhi","INDN":"Dadra and Nagar Haveli","INGA":"Goa","INGJ":"Gujarat","INHP":"Himachal Pradesh","INHR":"Haryana","INJH":"Jharkhand","INJK":"Jammu and Kashmir","INKA":"Karnataka","INKL":"Kerala","INLD":"Lakshadweep","INMH":"Maharashtra","INML":"Meghalaya","INMN":"Manipur","INMP":"Madhya Pradesh","INMZ":"Mizoram","INNL":"Nagaland","INOR":"Odisha (formerly known as Orissa)","INPB":"Punjab","INPY":"Puducherry (Pondicherry)","INRJ":"Rajasthan","INSK":"Sikkim","INTG":"Telangana","INTN":"Tamil Nadu","INTR":"Tripura","INUP":"Uttar Pradesh","INUT":"Uttarakhand","INWB":"West Bengal"}}`,
		},

		PolitiesIndonesia: &PolitiesCode{
			Name:        "Indonesia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg","gov_fiscal_year":"0101-1231","phone_code":"+62","states":null}`,
		},

		PolitiesIran: &PolitiesCode{
			Name:        "Iran",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ca/Flag_of_Iran.svg","phone_code":"+98","states":null}`,
		},

		PolitiesIraq: &PolitiesCode{
			Name:        "Iraq",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Iraq.svg","phone_code":"+964","states":null}`,
		},

		PolitiesIreland: &PolitiesCode{
			Name:        "Ireland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/45/Flag_of_Ireland.svg","gov_fiscal_year":"0101-1231","phone_code":"+353","states":null}`,
		},

		PolitiesIsleofMan: &PolitiesCode{
			Name:        "Isle of Man",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5d/Flag_of_the_Isle_of_Mann.svg","phone_code":"+44","states":null}`,
		},

		PolitiesIsrael: &PolitiesCode{
			Name:        "Israel",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Israel.svg","gov_fiscal_year":"0101-1231","phone_code":"+972","states":null}`,
		},

		PolitiesItaly: &PolitiesCode{
			Name:        "Italy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/0/03/Flag_of_Italy.svg","phone_code":"+39","states":null}`,
		},

		PolitiesCoteDIvoireFormerIvoryCoast: &PolitiesCode{
			Name:        "Cote D'Ivoire (Former Ivory Coast)",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_C%C3%B4te_d%27Ivoire.svg","phone_code":"+225","states":null}`,
		},

		PolitiesJamaica: &PolitiesCode{
			Name:        "Jamaica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Jamaica.svg","phone_code":"+1876","states":null}`,
		},

		PolitiesJapan: &PolitiesCode{
			Name:        "Japan",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg","gov_fiscal_year":"0401-0331","phone_code":"+81","states":null}`,
		},

		PolitiesJersey: &PolitiesCode{
			Name:        "Jersey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1c/Flag_of_Jersey.svg","phone_code":"+44","states":null}`,
		},

		PolitiesJordan: &PolitiesCode{
			Name:        "Jordan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c0/Flag_of_Jordan.svg","phone_code":"+962","states":null}`,
		},

		PolitiesKazakhstan: &PolitiesCode{
			Name:        "Kazakhstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kazakhstan.svg","phone_code":"+7","states":null}`,
		},

		PolitiesKenya: &PolitiesCode{
			Name:        "Kenya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Kenya.svg","phone_code":"+254","states":null}`,
		},

		PolitiesKiribati: &PolitiesCode{
			Name:        "Kiribati",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kiribati.svg","phone_code":"+686","states":null}`,
		},

		PolitiesKuwait: &PolitiesCode{
			Name:        "Kuwait",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/aa/Flag_of_Kuwait.svg","phone_code":"+965","states":null}`,
		},

		PolitiesKyrgyzstan: &PolitiesCode{
			Name:        "Kyrgyzstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c7/Flag_of_Kyrgyzstan.svg","phone_code":"+996","states":null}`,
		},

		PolitiesLaos: &PolitiesCode{
			Name:        "Laos",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Laos.svg","phone_code":"+856","states":null}`,
		},

		PolitiesLatvia: &PolitiesCode{
			Name:        "Latvia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Latvia.svg","phone_code":"+371","states":null}`,
		},

		PolitiesLebanon: &PolitiesCode{
			Name:        "Lebanon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/59/Flag_of_Lebanon.svg","phone_code":"+961","states":null}`,
		},

		PolitiesLesotho: &PolitiesCode{
			Name:        "Lesotho",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Lesotho.svg","phone_code":"+266","states":null}`,
		},

		PolitiesLiberia: &PolitiesCode{
			Name:        "Liberia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b8/Flag_of_Liberia.svg","phone_code":"+231","states":null}`,
		},

		PolitiesLibya: &PolitiesCode{
			Name:        "Libya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/05/Flag_of_Libya.svg","phone_code":"+218","states":null}`,
		},

		PolitiesLiechtenstein: &PolitiesCode{
			Name:        "Liechtenstein",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/47/Flag_of_Liechtenstein.svg","phone_code":"+423","states":null}`,
		},

		PolitiesLithuania: &PolitiesCode{
			Name:        "Lithuania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Lithuania.svg","phone_code":"+370","states":null}`,
		},

		PolitiesLuxembourg: &PolitiesCode{
			Name:        "Luxembourg",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/da/Flag_of_Luxembourg.svg","phone_code":"+352","states":null}`,
		},

		PolitiesMacau: &PolitiesCode{
			Name:        "Macau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/63/Flag_of_Macau.svg","phone_code":"+853","states":null}`,
		},

		PolitiesMacedonia: &PolitiesCode{
			Name:        "Macedonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_Macedonia.svg","phone_code":"+389","states":null}`,
		},

		PolitiesMadagascar: &PolitiesCode{
			Name:        "Madagascar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Madagascar.svg","phone_code":"+261","states":null}`,
		},

		PolitiesMalawi: &PolitiesCode{
			Name:        "Malawi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d1/Flag_of_Malawi.svg","phone_code":"+265","states":null}`,
		},

		PolitiesMalaysia: &PolitiesCode{
			Name:        "Malaysia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/66/Flag_of_Malaysia.svg","phone_code":"+60","states":null}`,
		},

		PolitiesMaldives: &PolitiesCode{
			Name:        "Maldives",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Maldives.svg","phone_code":"+960","states":null}`,
		},

		PolitiesMali: &PolitiesCode{
			Name:        "Mali",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Mali.svg","phone_code":"+223","states":null}`,
		},

		PolitiesMalta: &PolitiesCode{
			Name:        "Malta",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Malta.svg","phone_code":"+356","states":null}`,
		},

		PolitiesMarshallIslands: &PolitiesCode{
			Name:        "Marshall Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2e/Flag_of_the_Marshall_Islands.svg","phone_code":"+692","states":null}`,
		},

		PolitiesMartinique: &PolitiesCode{
			Name:        "Martinique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+596","states":null}`,
		},

		PolitiesMauritania: &PolitiesCode{
			Name:        "Mauritania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/New_National_Flag_of_Mauritania.png","phone_code":"+222","states":null}`,
		},

		PolitiesMauritius: &PolitiesCode{
			Name:        "Mauritius",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Mauritius.svg","phone_code":"+230","states":null}`,
		},

		PolitiesMayotte: &PolitiesCode{
			Name:        "Mayotte",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Mayotte_%28local%29.svg","phone_code":"+269","states":null}`,
		},

		PolitiesMexico: &PolitiesCode{
			Name:        "Mexico",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Mexico.svg","phone_code":"+52","states":{"MXAGU":"Aguascalientes","MXBCN":"Baja California","MXBCS":"Baja California Sur","MXCAM":"Campeche","MXCHH":"Chihuahua","MXCHP":"Chiapas","MXCOA":"Coahuila","MXCOL":"Colima","MXDIF":"Distrito Federal","MXDUR":"Durango","MXGRO":"Guerrero","MXGUA":"Guanajuato","MXHID":"Hidalgo","MXJAL":"Jalisco","MXMEX":"Mexico (Federal District)","MXMIC":"Michoacán","MXMOR":"Morelos","MXNAY":"Nayarit","MXNLE":"Nuevo León","MXOAX":"Oaxaca","MXPUE":"Puebla","MXQUE":"Querétaro","MXROO":"Quintana Roo","MXSIN":"Sinaloa","MXSLP":"San Luis Potosí","MXSON":"Sonora","MXTAB":"Tabasco","MXTAM":"Tamaulipas","MXTLA":"Tlaxcala","MXVER":"Veracruz","MXYUC":"Yucatán","MXZAC":"Zacatecas"}}`,
		},

		PolitiesMicronesia: &PolitiesCode{
			Name:        "Micronesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+691","states":null}`,
		},

		PolitiesMoldova: &PolitiesCode{
			Name:        "Moldova",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Moldova.svg","phone_code":"+373","states":null}`,
		},

		PolitiesMonaco: &PolitiesCode{
			Name:        "Monaco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ea/Flag_of_Monaco.svg","phone_code":"+377","states":null}`,
		},

		PolitiesMongolia: &PolitiesCode{
			Name:        "Mongolia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4c/Flag_of_Mongolia.svg","phone_code":"+976","states":null}`,
		},

		PolitiesMontenegro: &PolitiesCode{
			Name:        "Montenegro",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Montenegro.svg","phone_code":"+382","states":null}`,
		},

		PolitiesMontserrat: &PolitiesCode{
			Name:        "Montserrat",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Montserrat.svg","phone_code":"+1664","states":null}`,
		},

		PolitiesMorocco: &PolitiesCode{
			Name:        "Morocco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Morocco.svg","phone_code":"+212","states":null}`,
		},

		PolitiesMozambique: &PolitiesCode{
			Name:        "Mozambique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Mozambique.svg","phone_code":"+258","states":null}`,
		},

		PolitiesMyanmar: &PolitiesCode{
			Name:        "Myanmar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8c/Flag_of_Myanmar.svg","phone_code":"+95","states":null}`,
		},

		PolitiesNamibia: &PolitiesCode{
			Name:        "Namibia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Namibia.svg","phone_code":"+264","states":null}`,
		},

		PolitiesNauru: &PolitiesCode{
			Name:        "Nauru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/Flag_of_Nauru.svg","phone_code":"+674","states":null}`,
		},

		PolitiesNepal: &PolitiesCode{
			Name:        "Nepal",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9b/Flag_of_Nepal.svg","gov_fiscal_year":"0101-1231","phone_code":"+977","states":null}`,
		},

		PolitiesNetherlands: &PolitiesCode{
			Name:        "Netherlands",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","gov_fiscal_year":"0101-1231","phone_code":"+31","states":null}`,
		},

		PolitiesNewCaledonia: &PolitiesCode{
			Name:        "New Caledonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+687","states":null}`,
		},

		PolitiesNewZealand: &PolitiesCode{
			Name:        "New Zealand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3e/Flag_of_New_Zealand.svg","phone_code":"+64","states":null}`,
		},

		PolitiesNicaragua: &PolitiesCode{
			Name:        "Nicaragua",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Nicaragua.svg","phone_code":"+505","states":null}`,
		},

		PolitiesNiger: &PolitiesCode{
			Name:        "Niger",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f4/Flag_of_Niger.svg","phone_code":"+227","states":null}`,
		},

		PolitiesNigeria: &PolitiesCode{
			Name:        "Nigeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/79/Flag_of_Nigeria.svg","phone_code":"+234","states":null}`,
		},

		PolitiesNiue: &PolitiesCode{
			Name:        "Niue",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Niue.svg","phone_code":"+683","states":null}`,
		},

		PolitiesNorfolkandPhilipIsland: &PolitiesCode{
			Name:        "Norfolk and Philip Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Norfolk_Island.svg","phone_code":"+672","states":null}`,
		},

		PolitiesNorthKorea: &PolitiesCode{
			Name:        "North Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_North_Korea.svg","phone_code":"+850","states":null}`,
		},

		PolitiesNorthernMarianaIslands: &PolitiesCode{
			Name:        "Northern Mariana Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e0/Flag_of_the_Northern_Mariana_Islands.svg","phone_code":"+1670","states":null}`,
		},

		PolitiesNorway: &PolitiesCode{
			Name:        "Norway",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		},

		PolitiesOman: &PolitiesCode{
			Name:        "Oman",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Oman.svg","phone_code":"+968","states":null}`,
		},

		PolitiesPakistan: &PolitiesCode{
			Name:        "Pakistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/32/Flag_of_Pakistan.svg","phone_code":"+92","states":null}`,
		},

		PolitiesPalau: &PolitiesCode{
			Name:        "Palau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Palau.svg","phone_code":"+680","states":null}`,
		},

		PolitiesPalestinianTerritory: &PolitiesCode{
			Name:        "Palestinian Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Palestine.svg","phone_code":"+970","states":null}`,
		},

		PolitiesPanama: &PolitiesCode{
			Name:        "Panama",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/ab/Flag_of_Panama.svg","phone_code":"+507","states":null}`,
		},

		PolitiesPapuaNewGuinea: &PolitiesCode{
			Name:        "Papua New Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e3/Flag_of_Papua_New_Guinea.svg","phone_code":"+675","states":null}`,
		},

		PolitiesParaguay: &PolitiesCode{
			Name:        "Paraguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Paraguay.svg","phone_code":"+595","states":null}`,
		},

		PolitiesPeru: &PolitiesCode{
			Name:        "Peru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Peru_%28state%29.svg","phone_code":"+51","states":null}`,
		},

		PolitiesPhilippines: &PolitiesCode{
			Name:        "Philippines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_the_Philippines.svg","phone_code":"+63","states":null}`,
		},

		PolitiesPitcairnIslands: &PolitiesCode{
			Name:        "Pitcairn Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_the_Pitcairn_Islands.svg","phone_code":"+64","states":null}`,
		},

		PolitiesPoland: &PolitiesCode{
			Name:        "Poland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg","phone_code":"+48","states":null}`,
		},

		PolitiesPortugal: &PolitiesCode{
			Name:        "Portugal",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Portugal.svg","gov_fiscal_year":"0101-1231","phone_code":"+351","states":null}`,
		},

		PolitiesPuertoRico: &PolitiesCode{
			Name:        "Puerto Rico",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/28/Flag_of_Puerto_Rico.svg","phone_code":"+1787","states":null}`,
		},

		PolitiesQatar: &PolitiesCode{
			Name:        "Qatar",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Qatar.svg","gov_fiscal_year":"0101-1231","phone_code":"+974","states":null}`,
		},

		PolitiesRéunion: &PolitiesCode{
			Name:        "Réunion",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5a/Flag_of_R%C3%A9union.svg","phone_code":"+262","states":null}`,
		},

		PolitiesRomania: &PolitiesCode{
			Name:        "Romania",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Romania.svg","gov_fiscal_year":"0101-1231","phone_code":"+40","states":null}`,
		},

		PolitiesRussia: &PolitiesCode{
			Name:        "Russia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/f/f3/Flag_of_Russia.svg","gov_fiscal_year":"0101-1231","phone_code":"+7","states":{"RUAD":"Adygeya, Respublika","RUAL":"Altay, Respublika","RUALT":"Altayskiy kray","RUAMU":"Amurskaya oblast\"","RUARK":"Arkhangel\"skaya oblast\"","RUAST":"Astrakhanskaya oblast\"","RUBA":"Bashkortostan, Respublika","RUBEL":"Belgorodskaya oblast\"","RUBRY":"Bryanskaya oblast\"","RUBU":"Buryatiya, Respublika","RUCE":"Chechenskaya Respublika","RUCHE":"Chelyabinskaya oblast\"","RUCHU":"Chukotskiy avtonomnyy okrug","RUCU":"Chuvashskaya Respublika","RUDA":"Dagestan, Respublika","RUIN":"Ingushetiya, Respublika","RUIRK":"Irkutskaya oblast\"","RUIVA":"Ivanovskaya oblast\"","RUKAM":"Kamchatskiy kray","RUKB":"Kabardino-Balkarskaya Respublika","RUKC":"Karachayevo-Cherkesskaya Respubl.","RUKDA":"Krasnodarskiy kray","RUKEM":"Kemerovskaya oblast\"","RUKGD":"Kaliningradskaya oblast\"","RUKGN":"Kurganskaya oblast\"","RUKHA":"Khabarovskiy kray","RUKHM":"Khanty-Mansiyskiy avtonomnyy okrug-Yugra","RUKIR":"Kirovskaya oblast\"","RUKK":"Khakasiya, Respublika","RUKL":"Kalmykiya, Respublika","RUKLU":"Kaluzhskaya oblast\"","RUKO":"Komi, Respublika","RUKOS":"Kostromskaya oblast\"","RUKR":"Kareliya, Respublika","RUKRS":"Kurskaya oblast\"","RUKYA":"Krasnoyarskiy kray","RULEN":"Leningradskaya oblast\"","RULIP":"Lipetskaya oblast\"","RUMAG":"Magadanskaya oblast\"","RUME":"Mariy El, Respublika","RUMO":"Mordoviya, Respublika","RUMOS":"Moskovskaya oblast\"","RUMOW":"Moskva (autonomous city)","RUMUR":"Murmanskaya oblast\"","RUNEN":"Nenetskiy avtonomnyy okrug","RUNGR":"Novgorodskaya oblast\"","RUNIZ":"Nizhegorodskaya oblast\"","RUNVS":"Novosibirskaya oblast\"","RUOMS":"Omskaya oblast\"","RUORE":"Orenburgskaya oblast\"","RUORL":"Orlovskaya oblast\"","RUPER":"Permskiy kray","RUPNZ":"Penzenskaya oblast\"","RUPRI":"Primorskiy kray","RUPSK":"Pskovskaya oblast\"","RUROS":"Rostovskaya oblast\"","RURYA":"Ryazanskaya oblast\"","RUSA":"Sakha, Respublika","RUSAK":"Sakhalinskaya oblast\"","RUSAM":"Samarskaya oblast\"","RUSAR":"Saratovskaya oblast\"","RUSE":"Severnaya Osetiya-Alaniya, Respubl.","RUSMO":"Smolenskaya oblast\"","RUSPE":"Sankt-Peterburg (autonomous city)","RUSTA":"Stavropol\"skiy kray","RUSVE":"Sverdlovskaya oblast\"","RUTA":"Tatarstan, Respublika","RUTAM":"Tambovskaya oblast\"","RUTOM":"Tomskaya oblast\"","RUTUL":"Tul\"skaya oblast\"","RUTVE":"Tverskaya oblast\"","RUTY":"Tyva, Respublika","RUTYU":"Tyumenskaya oblast\"","RUUD":"Udmurtskaya Respublika","RUULY":"Ul\"yanovskaya oblast\"","RUVGG":"Volgogradskaya oblast\"","RUVLA":"Vladimirskaya oblast\"","RUVLG":"Vologodskaya oblast\"","RUVOR":"Voronezhskaya oblast\"","RUYAN":"Yamalo-Nenetskiy avtonomnyy okrug","RUYAR":"Yaroslavskaya oblast\"","RUYEV":"Yevreyskaya avtonomnaya oblast\"","RUZAB":"Zabaykal\"skiy kray"}}`,
		},

		PolitiesRwanda: &PolitiesCode{
			Name:        "Rwanda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Rwanda.svg","phone_code":"+250","states":null}`,
		},

		PolitiesSaintHelenaAscensionandTristandaCunha: &PolitiesCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","phone_code":"+290","states":null}`,
		},

		PolitiesSaintKittsandNevis: &PolitiesCode{
			Name:        "Saint Kitts and Nevis",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Saint_Kitts_and_Nevis.svg","phone_code":"+1869","states":null}`,
		},

		PolitiesSaintLucia: &PolitiesCode{
			Name:        "Saint Lucia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Saint_Lucia.svg","phone_code":"+1758","states":null}`,
		},

		PolitiesSaintPierreandMiquelon: &PolitiesCode{
			Name:        "Saint Pierre and Miquelon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+508","states":null}`,
		},

		PolitiesSaintVincentandtheGrenadines: &PolitiesCode{
			Name:        "Saint Vincent and the Grenadines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6d/Flag_of_Saint_Vincent_and_the_Grenadines.svg","phone_code":"+1784","states":null}`,
		},

		PolitiesSaintBarthelemy: &PolitiesCode{
			Name:        "Saint-Barthelemy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Saint_Barthelemy_%28local%29.svg","phone_code":"+590","states":null}`,
		},

		PolitiesSaintMartin: &PolitiesCode{
			Name:        "Saint-Martin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+590","states":null}`,
		},

		PolitiesSamoa: &PolitiesCode{
			Name:        "Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Samoa.svg","phone_code":"+685","states":null}`,
		},

		PolitiesSanMarino: &PolitiesCode{
			Name:        "San Marino",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_San_Marino.svg","phone_code":"+378","states":null}`,
		},

		PolitiesSãoToméandPríncipe: &PolitiesCode{
			Name:        "São Tomé and Príncipe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Sao_Tome_and_Principe.svg","phone_code":"+239","states":null}`,
		},

		PolitiesSaudiArabia: &PolitiesCode{
			Name:        "Saudi Arabia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0d/Flag_of_Saudi_Arabia.svg","phone_code":"+966","states":null}`,
		},

		PolitiesSenegal: &PolitiesCode{
			Name:        "Senegal",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fd/Flag_of_Senegal.svg","phone_code":"+221","states":null}`,
		},

		PolitiesSerbia: &PolitiesCode{
			Name:        "Serbia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/ff/Flag_of_Serbia.svg","phone_code":"+381","states":null}`,
		},

		PolitiesSeychelles: &PolitiesCode{
			Name:        "Seychelles",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Seychelles.svg","phone_code":"+248","states":null}`,
		},

		PolitiesSierraLeone: &PolitiesCode{
			Name:        "Sierra Leone",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Sierra_Leone.svg","phone_code":"+232","states":null}`,
		},

		PolitiesSingapore: &PolitiesCode{
			Name:        "Singapore",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg","phone_code":"+65","states":null}`,
		},

		PolitiesSintMaarten: &PolitiesCode{
			Name:        "Sint Maarten",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Sint_Maarten.svg","phone_code":"+1721","states":null}`,
		},

		PolitiesSlovakia: &PolitiesCode{
			Name:        "Slovakia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e6/Flag_of_Slovakia.svg","phone_code":"+421","states":null}`,
		},

		PolitiesSlovenia: &PolitiesCode{
			Name:        "Slovenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f0/Flag_of_Slovenia.svg","phone_code":"+386","states":null}`,
		},

		PolitiesSolomonIslands: &PolitiesCode{
			Name:        "Solomon Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Solomon_Islands.svg","phone_code":"+677","states":null}`,
		},

		PolitiesSomalia: &PolitiesCode{
			Name:        "Somalia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_Somalia.svg","phone_code":"+252","states":null}`,
		},

		PolitiesSouthAfrica: &PolitiesCode{
			Name:        "South Africa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/af/Flag_of_South_Africa.svg","phone_code":"+27","states":null}`,
		},

		PolitiesSouthGeorgiaandtheSouthSandwichIslands: &PolitiesCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_South_Georgia_and_the_South_Sandwich_Islands.svg","phone_code":"+500","states":null}`,
		},

		PolitiesSouthKorea: &PolitiesCode{
			Name:        "South Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg","gov_fiscal_year":"0101-1231","phone_code":"+82","states":null}`,
		},

		PolitiesSouthSudan: &PolitiesCode{
			Name:        "South Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7a/Flag_of_South_Sudan.svg","phone_code":"+211","states":null}`,
		},

		PolitiesSpain: &PolitiesCode{
			Name:        "Spain",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9a/Flag_of_Spain.svg","gov_fiscal_year":"0101-1231","phone_code":"+34","states":null}`,
		},

		PolitiesSriLanka: &PolitiesCode{
			Name:        "Sri Lanka",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Sri_Lanka.svg","phone_code":"+94","states":null}`,
		},

		PolitiesSudan: &PolitiesCode{
			Name:        "Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Sudan.svg","phone_code":"+249","states":null}`,
		},

		PolitiesSuriname: &PolitiesCode{
			Name:        "Suriname",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/60/Flag_of_Suriname.svg","phone_code":"+597","states":null}`,
		},

		PolitiesSvalbardandJanMayen: &PolitiesCode{
			Name:        "Svalbard and Jan Mayen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+47","states":null}`,
		},

		PolitiesSwaziland: &PolitiesCode{
			Name:        "Swaziland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fb/Flag_of_Eswatini.svg","phone_code":"+268","states":null}`,
		},

		PolitiesSweden: &PolitiesCode{
			Name:        "Sweden",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/4/4c/Flag_of_Sweden.svg","gov_fiscal_year":"0101-1231","phone_code":"+46","states":null}`,
		},

		PolitiesSwitzerland: &PolitiesCode{
			Name:        "Switzerland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/08/Flag_of_Switzerland_%28Pantone%29.svg","gov_fiscal_year":"0101-1231","phone_code":"+41","states":null}`,
		},

		PolitiesSyria: &PolitiesCode{
			Name:        "Syria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/53/Flag_of_Syria.svg","phone_code":"+963","states":null}`,
		},

		PolitiesTaiwan: &PolitiesCode{
			Name:        "Taiwan",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+886","states":null}`,
		},

		PolitiesTajikistan: &PolitiesCode{
			Name:        "Tajikistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Tajikistan.svg","phone_code":"+992","states":null}`,
		},

		PolitiesTanzania: &PolitiesCode{
			Name:        "Tanzania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tanzania.svg","phone_code":"+255","states":null}`,
		},

		PolitiesThailand: &PolitiesCode{
			Name:        "Thailand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a9/Flag_of_Thailand.svg","phone_code":"+66","states":null}`,
		},

		PolitiesTogo: &PolitiesCode{
			Name:        "Togo",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/68/Flag_of_Togo.svg","phone_code":"+228","states":null}`,
		},

		PolitiesTokelau: &PolitiesCode{
			Name:        "Tokelau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8e/Flag_of_Tokelau.svg","phone_code":"+690","states":null}`,
		},

		PolitiesTonga: &PolitiesCode{
			Name:        "Tonga",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Tonga.svg","phone_code":"+676","states":null}`,
		},

		PolitiesTrinidadandTobago: &PolitiesCode{
			Name:        "Trinidad and Tobago",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Trinidad_and_Tobago.svg","phone_code":"+1868","states":null}`,
		},

		PolitiesTunisia: &PolitiesCode{
			Name:        "Tunisia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Tunisia.svg","phone_code":"+216","states":null}`,
		},

		PolitiesTurkey: &PolitiesCode{
			Name:        "Turkey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Turkey.svg","phone_code":"+90","states":null}`,
		},

		PolitiesTurkmenistan: &PolitiesCode{
			Name:        "Turkmenistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Turkmenistan.svg","phone_code":"+993","states":null}`,
		},

		PolitiesTurksandCaicosIslands: &PolitiesCode{
			Name:        "Turks and Caicos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_the_Turks_and_Caicos_Islands.svg","phone_code":"+1649","states":null}`,
		},

		PolitiesTuvalu: &PolitiesCode{
			Name:        "Tuvalu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tuvalu.svg","phone_code":"+688","states":null}`,
		},

		PolitiesUganda: &PolitiesCode{
			Name:        "Uganda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4e/Flag_of_Uganda.svg","phone_code":"+256","states":null}`,
		},

		PolitiesUkraine: &PolitiesCode{
			Name:        "Ukraine",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Ukraine.svg","phone_code":"+380","states":null}`,
		},

		PolitiesUnitedArabEmirates: &PolitiesCode{
			Name:        "United Arab Emirates",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_United_Arab_Emirates.svg","gov_fiscal_year":"0101-1231","phone_code":"+971","states":null}`,
		},

		PolitiesUnitedKingdom: &PolitiesCode{
			Name:        "United Kingdom",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0406-0405","flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","gov_fiscal_year":"0401-0331","phone_code":"+44","states":null}`,
		},

		PolitiesUnitedStatesMinorOutlyingIslands: &PolitiesCode{
			Name:        "United States Minor Outlying Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+246","states":null}`,
		},

		PolitiesUruguay: &PolitiesCode{
			Name:        "Uruguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Uruguay.svg","phone_code":"+598","states":null}`,
		},

		PolitiesUSVirginIslands: &PolitiesCode{
			Name:        "US Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_the_United_States_Virgin_Islands.svg","phone_code":"+1340","states":null}`,
		},

		PolitiesUSA: &PolitiesCode{
			Name:        "USA",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+1","states":{"USAK":"Alaska","USAL":"Alabama","USAR":"Arkansas","USAS":"American Samoa","USAZ":"Arizona","USCA":"California","USCO":"Colorado","USCT":"Connecticut","USDC":"Washington D.C.","USDE":"Delaware","USFL":"Florida","USGA":"Georgia","USGU":"Guam","USHI":"Hawaii","USIA":"Iowa","USID":"Idaho","USIL":"Illinois","USIN":"Indiana","USKS":"Kansas","USKY":"Kentucky","USLA":"Louisiana","USMA":"Massachusetts","USMD":"Maryland","USME":"Maine","USMI":"Michigan","USMN":"Minnesota","USMO":"Missouri","USMP":"Northern Mariana Islands","USMS":"Mississippi","USMT":"Montana","USNC":"North Carolina","USND":"North Dakota","USNE":"Nebraska","USNH":"New Hampshire","USNJ":"New Jersey","USNM":"New Mexico","USNV":"Nevada","USNY":"New York","USOH":"Ohio","USOK":"Oklahoma","USOR":"Oregon","USPA":"Pennsylvania","USPR":"Puerto Rico","USRI":"Rhode Island","USSC":"South Carolina","USSD":"South Dakota","USTN":"Tennessee","USTX":"Texas","USUM":"United States Minor Outlying Islands","USUT":"Utah","USVA":"Virginia","USVI":"US Virgin Islands","USVT":"Vermont","USWA":"Washington","USWI":"Wisconsin","USWV":"West Virginia","USWY":"Wyoming"}}`,
		},

		PolitiesUzbekistan: &PolitiesCode{
			Name:        "Uzbekistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Uzbekistan.svg","phone_code":"+998","states":null}`,
		},

		PolitiesVanuatu: &PolitiesCode{
			Name:        "Vanuatu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_Vanuatu_%28official%29.svg","phone_code":"+678","states":null}`,
		},

		PolitiesVaticanCity: &PolitiesCode{
			Name:        "Vatican City",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_the_Vatican_City.svg","phone_code":"+418","states":null}`,
		},

		PolitiesVenezuela: &PolitiesCode{
			Name:        "Venezuela",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7b/Flag_of_Venezuela_%28state%29.svg","phone_code":"+58","states":null}`,
		},

		PolitiesVietnam: &PolitiesCode{
			Name:        "Vietnam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Vietnam.svg","phone_code":"+84","states":null}`,
		},

		PolitiesWallisandFutuna: &PolitiesCode{
			Name:        "Wallis and Futuna",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+681","states":null}`,
		},

		PolitiesWesternSahara: &PolitiesCode{
			Name:        "Western Sahara",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"phone_code":"+212","states":null}`,
		},

		PolitiesYemen: &PolitiesCode{
			Name:        "Yemen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Yemen.svg","phone_code":"+967","states":null}`,
		},

		PolitiesZambia: &PolitiesCode{
			Name:        "Zambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/06/Flag_of_Zambia.svg","phone_code":"+260","states":null}`,
		},

		PolitiesZimbabwe: &PolitiesCode{
			Name:        "Zimbabwe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6a/Flag_of_Zimbabwe.svg","phone_code":"+263","states":null}`,
		},
	}
}

/************************************ LegalSystems ************************************/
const (

	// LawofEnglandandWales -
	LegalSystemsLawofEnglandandWales = "GBREW"

	// LawofScotland -
	LegalSystemsLawofScotland = "GBRS"

	// LawofAfghanistan -
	LegalSystemsLawofAfghanistan = "AFG"

	// LawofAlbania -
	LegalSystemsLawofAlbania = "ALB"

	// LawofAlgeria -
	LegalSystemsLawofAlgeria = "DZA"

	// LawofAmericanSamoa -
	LegalSystemsLawofAmericanSamoa = "ASM"

	// LawofAndorra -
	LegalSystemsLawofAndorra = "AND"

	// LawofAngola -
	LegalSystemsLawofAngola = "AGO"

	// LawofAnguilla -
	LegalSystemsLawofAnguilla = "AIA"

	// LawofAntiguaandBarbuda -
	LegalSystemsLawofAntiguaandBarbuda = "ATG"

	// LawofArgentina -
	LegalSystemsLawofArgentina = "ARG"

	// LawofArmenia -
	LegalSystemsLawofArmenia = "ARM"

	// LawofAruba -
	LegalSystemsLawofAruba = "ABW"

	// LawofAustralia -
	LegalSystemsLawofAustralia = "AUS"

	// LawofNewSouthWales -
	LegalSystemsLawofNewSouthWales = "AUNSW"

	// LawofQueensland -
	LegalSystemsLawofQueensland = "AUQLD"

	// LawofSouthAustralia -
	LegalSystemsLawofSouthAustralia = "AUSA"

	// LawofTasmania -
	LegalSystemsLawofTasmania = "AUTAS"

	// LawofVictoria -
	LegalSystemsLawofVictoria = "AUVIC"

	// LawofWesternAustralia -
	LegalSystemsLawofWesternAustralia = "AUWA"

	// LawofAustralianCapitalTerritory -
	LegalSystemsLawofAustralianCapitalTerritory = "AUACT"

	// LawofNorthernTerritory -
	LegalSystemsLawofNorthernTerritory = "AUNT"

	// LawofJervisBayTerritoryAustralianCapitalTerritory -
	LegalSystemsLawofJervisBayTerritoryAustralianCapitalTerritory = "AUJBT"

	// LawofAustria -
	LegalSystemsLawofAustria = "AUT"

	// LawofAzerbaijan -
	LegalSystemsLawofAzerbaijan = "AZE"

	// LawofBahamas -
	LegalSystemsLawofBahamas = "BHS"

	// LawofBahrain -
	LegalSystemsLawofBahrain = "BHR"

	// LawofBangladesh -
	LegalSystemsLawofBangladesh = "BGD"

	// LawofBarbados -
	LegalSystemsLawofBarbados = "BRB"

	// LawofBelarus -
	LegalSystemsLawofBelarus = "BLR"

	// LawofBelgium -
	LegalSystemsLawofBelgium = "BEL"

	// LawofBelize -
	LegalSystemsLawofBelize = "BLZ"

	// LawofBenin -
	LegalSystemsLawofBenin = "BEN"

	// LawofBermuda -
	LegalSystemsLawofBermuda = "BMU"

	// LawofBhutan -
	LegalSystemsLawofBhutan = "BTN"

	// LawofBolivia -
	LegalSystemsLawofBolivia = "BOL"

	// LawofBonaireStEustasuisandSaba -
	LegalSystemsLawofBonaireStEustasuisandSaba = "BES"

	// LawofBosniaandHerzegovina -
	LegalSystemsLawofBosniaandHerzegovina = "BIH"

	// LawofBotswana -
	LegalSystemsLawofBotswana = "BWA"

	// LawofBrazil -
	LegalSystemsLawofBrazil = "BRA"

	// LawofBritishVirginIslands -
	LegalSystemsLawofBritishVirginIslands = "VGB"

	// LawofBrunei -
	LegalSystemsLawofBrunei = "BRN"

	// LawofBulgaria -
	LegalSystemsLawofBulgaria = "BGR"

	// LawofBurkinaFaso -
	LegalSystemsLawofBurkinaFaso = "BFA"

	// LawofBurundi -
	LegalSystemsLawofBurundi = "BDI"

	// LawofCambodia -
	LegalSystemsLawofCambodia = "KHM"

	// LawofCameroon -
	LegalSystemsLawofCameroon = "CMR"

	// LawofCanada -
	LegalSystemsLawofCanada = "CAN"

	// LawofAlberta -
	LegalSystemsLawofAlberta = "CAAB"

	// LawofBritishColumbia -
	LegalSystemsLawofBritishColumbia = "CABC"

	// LawofManitoba -
	LegalSystemsLawofManitoba = "CAMB"

	// LawofNewBrunswick -
	LegalSystemsLawofNewBrunswick = "CANB"

	// LawofNewfoundlandandLabrador -
	LegalSystemsLawofNewfoundlandandLabrador = "CANL"

	// LawofNovaScotia -
	LegalSystemsLawofNovaScotia = "CANS"

	// LawofOntario -
	LegalSystemsLawofOntario = "CAON"

	// LawofPrinceEdwardIsland -
	LegalSystemsLawofPrinceEdwardIsland = "CAPE"

	// LawofQuebec -
	LegalSystemsLawofQuebec = "CAQC"

	// LawofSaskatchewan -
	LegalSystemsLawofSaskatchewan = "CASK"

	// LawofNorthwestTerritories -
	LegalSystemsLawofNorthwestTerritories = "CANT"

	// LawofNunavut -
	LegalSystemsLawofNunavut = "CANU"

	// LawofYukon -
	LegalSystemsLawofYukon = "CAYT"

	// LawofCapeVerde -
	LegalSystemsLawofCapeVerde = "CPV"

	// LawofCaymanIslands -
	LegalSystemsLawofCaymanIslands = "CYM"

	// LawofCentralAfricanRepublic -
	LegalSystemsLawofCentralAfricanRepublic = "CAF"

	// LawofChad -
	LegalSystemsLawofChad = "TCD"

	// LawofChile -
	LegalSystemsLawofChile = "CHL"

	// LawofPeoplesRepublicofChina -
	LegalSystemsLawofPeoplesRepublicofChina = "CHN"

	// LawofChristmasIslandWesternAustralia -
	LegalSystemsLawofChristmasIslandWesternAustralia = "CXR"

	// LawofCocosIslandsWesternAustralia -
	LegalSystemsLawofCocosIslandsWesternAustralia = "CCK"

	// LawofColombia -
	LegalSystemsLawofColombia = "COL"

	// LawofComoros -
	LegalSystemsLawofComoros = "COM"

	// LawofRepublicoftheCongo -
	LegalSystemsLawofRepublicoftheCongo = "COG"

	// LawofDemocraticRepublicoftheCongo -
	LegalSystemsLawofDemocraticRepublicoftheCongo = "COD"

	// LawofCookIslands -
	LegalSystemsLawofCookIslands = "COK"

	// LawofCostaRica -
	LegalSystemsLawofCostaRica = "CRI"

	// LawofCroatia -
	LegalSystemsLawofCroatia = "HRV"

	// LawofCuba -
	LegalSystemsLawofCuba = "CUB"

	// LawofCuracao -
	LegalSystemsLawofCuracao = "CUW"

	// LawofCyprus -
	LegalSystemsLawofCyprus = "CYP"

	// LawofCzechRepublic -
	LegalSystemsLawofCzechRepublic = "CZE"

	// LawofDenmark -
	LegalSystemsLawofDenmark = "DNK"

	// LawofDjibouti -
	LegalSystemsLawofDjibouti = "DJI"

	// LawofDominica -
	LegalSystemsLawofDominica = "DMA"

	// LawofDominicanRepublic -
	LegalSystemsLawofDominicanRepublic = "DOM"

	// LawofTimorLeste -
	LegalSystemsLawofTimorLeste = "TLS"

	// LawofEcuador -
	LegalSystemsLawofEcuador = "ECU"

	// LawofEgypt -
	LegalSystemsLawofEgypt = "EGY"

	// LawofElSalvador -
	LegalSystemsLawofElSalvador = "SLV"

	// LawofEquatorialGuinea -
	LegalSystemsLawofEquatorialGuinea = "GNQ"

	// LawofEritrea -
	LegalSystemsLawofEritrea = "ERI"

	// LawofEstonia -
	LegalSystemsLawofEstonia = "EST"

	// LawofEthiopia -
	LegalSystemsLawofEthiopia = "ETH"

	// LawofFalklandIslandsEnglandandWales -
	LegalSystemsLawofFalklandIslandsEnglandandWales = "FLK"

	// LawofFaroeIslandsDenmark -
	LegalSystemsLawofFaroeIslandsDenmark = "FRO"

	// LawofFiji -
	LegalSystemsLawofFiji = "FJI"

	// LawofFinland -
	LegalSystemsLawofFinland = "FIN"

	// LawofFrance -
	LegalSystemsLawofFrance = "FRA"

	// LawofFrenchGuianaFrance -
	LegalSystemsLawofFrenchGuianaFrance = "GUF"

	// LawofFrenchPolynesiaFrance -
	LegalSystemsLawofFrenchPolynesiaFrance = "PYF"

	// LawofFrenchSouthernandAntarcticLandsFrance -
	LegalSystemsLawofFrenchSouthernandAntarcticLandsFrance = "ATF"

	// LawofGabon -
	LegalSystemsLawofGabon = "GAB"

	// LawofTheGambia -
	LegalSystemsLawofTheGambia = "GMB"

	// LawofGeorgia -
	LegalSystemsLawofGeorgia = "GEO"

	// LawofGermany -
	LegalSystemsLawofGermany = "DEU"

	// LawofGhana -
	LegalSystemsLawofGhana = "GHA"

	// LawofGibraltar -
	LegalSystemsLawofGibraltar = "GIB"

	// LawofGreece -
	LegalSystemsLawofGreece = "GRC"

	// LawofGreenlandDenmark -
	LegalSystemsLawofGreenlandDenmark = "GRL"

	// LawofGrenada -
	LegalSystemsLawofGrenada = "GRD"

	// LawofGuadeloupeFrance -
	LegalSystemsLawofGuadeloupeFrance = "GLP"

	// LawofGuamUnitedStates -
	LegalSystemsLawofGuamUnitedStates = "GUM"

	// LawofGuatemala -
	LegalSystemsLawofGuatemala = "GTM"

	// LawofGuernsey -
	LegalSystemsLawofGuernsey = "GGY"

	// LawofGuinea -
	LegalSystemsLawofGuinea = "GIN"

	// LawofGuineaBissau -
	LegalSystemsLawofGuineaBissau = "GNB"

	// LawofGuyana -
	LegalSystemsLawofGuyana = "GUY"

	// LawofHaiti -
	LegalSystemsLawofHaiti = "HTI"

	// LawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory -
	LegalSystemsLawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory = "HMD"

	// LawofHonduras -
	LegalSystemsLawofHonduras = "HND"

	// LawofHongKong -
	LegalSystemsLawofHongKong = "HKG"

	// LawofHungary -
	LegalSystemsLawofHungary = "HUN"

	// LawofIceland -
	LegalSystemsLawofIceland = "ISL"

	// LawofIndia -
	LegalSystemsLawofIndia = "IND"

	// LawofDadraandNagarHaveli -
	LegalSystemsLawofDadraandNagarHaveli = "INDN"

	// LawofDamanandDiu -
	LegalSystemsLawofDamanandDiu = "INDD"

	// LawofGoa -
	LegalSystemsLawofGoa = "INGA"

	// LawofIndonesia -
	LegalSystemsLawofIndonesia = "IDN"

	// LawofIran -
	LegalSystemsLawofIran = "IRN"

	// LawofIraq -
	LegalSystemsLawofIraq = "IRQ"

	// LawofIreland -
	LegalSystemsLawofIreland = "IRL"

	// IsleofMan -
	LegalSystemsIsleofMan = "IMN"

	// LawofIsrael -
	LegalSystemsLawofIsrael = "ISR"

	// LawofItaly -
	LegalSystemsLawofItaly = "ITA"

	// LawofCotedIvoire -
	LegalSystemsLawofCotedIvoire = "CIV"

	// LawofJamaica -
	LegalSystemsLawofJamaica = "JAM"

	// LawofJapan -
	LegalSystemsLawofJapan = "JPN"

	// LawofJersey -
	LegalSystemsLawofJersey = "JEY"

	// LawofJordan -
	LegalSystemsLawofJordan = "JOR"

	// LawofKazakhstan -
	LegalSystemsLawofKazakhstan = "KAZ"

	// LawofKenya -
	LegalSystemsLawofKenya = "KEN"

	// LawofKiribati -
	LegalSystemsLawofKiribati = "KIR"

	// LawofKuwait -
	LegalSystemsLawofKuwait = "KWT"

	// LawofKyrgyzstan -
	LegalSystemsLawofKyrgyzstan = "KGZ"

	// LawofLaos -
	LegalSystemsLawofLaos = "LAO"

	// LawofLatvia -
	LegalSystemsLawofLatvia = "LVA"

	// LawofLebanon -
	LegalSystemsLawofLebanon = "LBN"

	// LawofLesotho -
	LegalSystemsLawofLesotho = "LSO"

	// LawofLiberia -
	LegalSystemsLawofLiberia = "LBR"

	// LawofLibya -
	LegalSystemsLawofLibya = "LBY"

	// LawofLiechtenstein -
	LegalSystemsLawofLiechtenstein = "LIE"

	// LawofLithuania -
	LegalSystemsLawofLithuania = "LTU"

	// LawofLuxembourg -
	LegalSystemsLawofLuxembourg = "LUX"

	// LawofMacau -
	LegalSystemsLawofMacau = "MAC"

	// LawofMacedonia -
	LegalSystemsLawofMacedonia = "MKD"

	// LawofMadagascar -
	LegalSystemsLawofMadagascar = "MDG"

	// LawofMalawi -
	LegalSystemsLawofMalawi = "MWI"

	// LawofMalaysia -
	LegalSystemsLawofMalaysia = "MYS"

	// LawofMaldives -
	LegalSystemsLawofMaldives = "MDV"

	// LawofMali -
	LegalSystemsLawofMali = "MLI"

	// LawofMalta -
	LegalSystemsLawofMalta = "MLT"

	// LawofMarshallIslands -
	LegalSystemsLawofMarshallIslands = "MHL"

	// LawofMartiniqueFrance -
	LegalSystemsLawofMartiniqueFrance = "MTQ"

	// LawofMauritania -
	LegalSystemsLawofMauritania = "MRT"

	// LawofMauritius -
	LegalSystemsLawofMauritius = "MUS"

	// LawofMayotteFrance -
	LegalSystemsLawofMayotteFrance = "MYT"

	// LawofMexico -
	LegalSystemsLawofMexico = "MEX"

	// LawofMicronesia -
	LegalSystemsLawofMicronesia = "FSM"

	// LawofMoldova -
	LegalSystemsLawofMoldova = "MDA"

	// LawofMonaco -
	LegalSystemsLawofMonaco = "MCO"

	// LawofMongolia -
	LegalSystemsLawofMongolia = "MNG"

	// LawofMontenegro -
	LegalSystemsLawofMontenegro = "MNE"

	// LawofMontserratEnglandandWales -
	LegalSystemsLawofMontserratEnglandandWales = "MSR"

	// LawofMorocco -
	LegalSystemsLawofMorocco = "MAR"

	// LawofMozambique -
	LegalSystemsLawofMozambique = "MOZ"

	// LawofMyanmar -
	LegalSystemsLawofMyanmar = "MMR"

	// LawofNamibia -
	LegalSystemsLawofNamibia = "NAM"

	// LawofNauru -
	LegalSystemsLawofNauru = "NRU"

	// LawofNepal -
	LegalSystemsLawofNepal = "NPL"

	// LawofNetherlands -
	LegalSystemsLawofNetherlands = "NLD"

	// LawofNewCaledonia -
	LegalSystemsLawofNewCaledonia = "NCL"

	// LawofNewZealand -
	LegalSystemsLawofNewZealand = "NZL"

	// LawofNicaragua -
	LegalSystemsLawofNicaragua = "NIC"

	// LawofNiger -
	LegalSystemsLawofNiger = "NER"

	// LawofNigeria -
	LegalSystemsLawofNigeria = "NGA"

	// LawofNiue -
	LegalSystemsLawofNiue = "NIU"

	// LawofNorfolkIslandNewSouthWales -
	LegalSystemsLawofNorfolkIslandNewSouthWales = "NFK"

	// LawofNorthKorea -
	LegalSystemsLawofNorthKorea = "PRK"

	// LawofNorthernMarianaIslandsUnitedStates -
	LegalSystemsLawofNorthernMarianaIslandsUnitedStates = "MNP"

	// LawofNorway -
	LegalSystemsLawofNorway = "NOR"

	// LawofOman -
	LegalSystemsLawofOman = "OMN"

	// LawofPakistan -
	LegalSystemsLawofPakistan = "PAK"

	// LawofPalau -
	LegalSystemsLawofPalau = "PLW"

	// LawofPalestinianTerritory -
	LegalSystemsLawofPalestinianTerritory = "PSE"

	// LawofPanama -
	LegalSystemsLawofPanama = "PAN"

	// LawofPapuaNewGuinea -
	LegalSystemsLawofPapuaNewGuinea = "PNG"

	// LawofParaguay -
	LegalSystemsLawofParaguay = "PRY"

	// LawofPeru -
	LegalSystemsLawofPeru = "PER"

	// LawofPhilippines -
	LegalSystemsLawofPhilippines = "PHL"

	// LawofPitcairnIslandsEnglandandWales -
	LegalSystemsLawofPitcairnIslandsEnglandandWales = "PCN"

	// LawofPoland -
	LegalSystemsLawofPoland = "POL"

	// LawofPortugal -
	LegalSystemsLawofPortugal = "PRT"

	// LawofPuertoRico -
	LegalSystemsLawofPuertoRico = "PRI"

	// LawofQatar -
	LegalSystemsLawofQatar = "QAT"

	// LawofRéunionFrance -
	LegalSystemsLawofRéunionFrance = "REU"

	// LawofRomania -
	LegalSystemsLawofRomania = "ROU"

	// LawofRussia -
	LegalSystemsLawofRussia = "RUS"

	// LawofRwanda -
	LegalSystemsLawofRwanda = "RWA"

	// LawofSaintHelenaAscensionandTristandaCunhaEnglandandWales -
	LegalSystemsLawofSaintHelenaAscensionandTristandaCunhaEnglandandWales = "SHN"

	// LawofSaintKittsandNevis -
	LegalSystemsLawofSaintKittsandNevis = "KNA"

	// LawofSaintLucia -
	LegalSystemsLawofSaintLucia = "LCA"

	// LawofSaintPierreandMiquelonFrance -
	LegalSystemsLawofSaintPierreandMiquelonFrance = "SPM"

	// LawofSaintVincentandtheGrenadines -
	LegalSystemsLawofSaintVincentandtheGrenadines = "VCT"

	// LawofSaintBarthelemyFrance -
	LegalSystemsLawofSaintBarthelemyFrance = "BLM"

	// LawofSaintMartinFrance -
	LegalSystemsLawofSaintMartinFrance = "MAF"

	// LawofSamoa -
	LegalSystemsLawofSamoa = "WSM"

	// LawofSanMarino -
	LegalSystemsLawofSanMarino = "SMR"

	// LawofSãoToméandPríncipe -
	LegalSystemsLawofSãoToméandPríncipe = "STP"

	// LawofSaudiArabia -
	LegalSystemsLawofSaudiArabia = "SAU"

	// LawofSenegal -
	LegalSystemsLawofSenegal = "SEN"

	// LawofSerbia -
	LegalSystemsLawofSerbia = "SRB"

	// LawofSeychelles -
	LegalSystemsLawofSeychelles = "SYC"

	// LawofSierraLeoneEnglandandWales -
	LegalSystemsLawofSierraLeoneEnglandandWales = "SLE"

	// LawofSingapore -
	LegalSystemsLawofSingapore = "SGP"

	// LawofSintMaarten -
	LegalSystemsLawofSintMaarten = "SXM"

	// LawofSlovakia -
	LegalSystemsLawofSlovakia = "SVK"

	// LawofSlovenia -
	LegalSystemsLawofSlovenia = "SVN"

	// LawofSolomonIslands -
	LegalSystemsLawofSolomonIslands = "SLB"

	// LawofSomalia -
	LegalSystemsLawofSomalia = "SOM"

	// LawofSouthAfrica -
	LegalSystemsLawofSouthAfrica = "ZAF"

	// LawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales -
	LegalSystemsLawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales = "SGS"

	// LawofSouthKorea -
	LegalSystemsLawofSouthKorea = "KOR"

	// LawofSouthSudan -
	LegalSystemsLawofSouthSudan = "SSD"

	// LawofSpain -
	LegalSystemsLawofSpain = "ESP"

	// LawofSriLanka -
	LegalSystemsLawofSriLanka = "LKA"

	// LawofSudan -
	LegalSystemsLawofSudan = "SDN"

	// LawofSuriname -
	LegalSystemsLawofSuriname = "SUR"

	// LawofSvalbardandJanMayenNorway -
	LegalSystemsLawofSvalbardandJanMayenNorway = "SJM"

	// LawofEswatini -
	LegalSystemsLawofEswatini = "SWZ"

	// LawofSweden -
	LegalSystemsLawofSweden = "SWE"

	// LawofSwitzerland -
	LegalSystemsLawofSwitzerland = "CHE"

	// LawofSyria -
	LegalSystemsLawofSyria = "SYR"

	// LawofTaiwan -
	LegalSystemsLawofTaiwan = "TWN"

	// LawofTajikistan -
	LegalSystemsLawofTajikistan = "TJK"

	// LawofTanzania -
	LegalSystemsLawofTanzania = "TZA"

	// LawofThailand -
	LegalSystemsLawofThailand = "THA"

	// LawofTogo -
	LegalSystemsLawofTogo = "TGO"

	// LawofTokelau -
	LegalSystemsLawofTokelau = "TKL"

	// LawofTrinidadandTobago -
	LegalSystemsLawofTrinidadandTobago = "TON"

	// LawofTunisia -
	LegalSystemsLawofTunisia = "TUN"

	// LawofTurkey -
	LegalSystemsLawofTurkey = "TUR"

	// LawofTurkmenistan -
	LegalSystemsLawofTurkmenistan = "TKM"

	// LawofTurksandCaicosIsland -
	LegalSystemsLawofTurksandCaicosIsland = "TCA"

	// LawofTuvalu -
	LegalSystemsLawofTuvalu = "TUV"

	// LawofUganda -
	LegalSystemsLawofUganda = "UGA"

	// LawofUkraine -
	LegalSystemsLawofUkraine = "UKR"

	// LawofUnitedArabEmirates -
	LegalSystemsLawofUnitedArabEmirates = "ARE"

	// LawofNorthernIreland -
	LegalSystemsLawofNorthernIreland = "GBRNI"

	// LawofUruguay -
	LegalSystemsLawofUruguay = "URY"

	// LawofUSVirginIslandsUnitedStates -
	LegalSystemsLawofUSVirginIslandsUnitedStates = "VIR"

	// LawofUnitedStates -
	LegalSystemsLawofUnitedStates = "USA"

	// LawofAlaska -
	LegalSystemsLawofAlaska = "USAK"

	// LawofAlabama -
	LegalSystemsLawofAlabama = "USAL"

	// LawofArkansas -
	LegalSystemsLawofArkansas = "USAR"

	// LawofArizona -
	LegalSystemsLawofArizona = "USAZ"

	// LawofCalifornia -
	LegalSystemsLawofCalifornia = "USCA"

	// LawofColorado -
	LegalSystemsLawofColorado = "USCO"

	// LawofConnecticut -
	LegalSystemsLawofConnecticut = "USCT"

	// LawofWashingtonDC -
	LegalSystemsLawofWashingtonDC = "USDC"

	// LawofDelaware -
	LegalSystemsLawofDelaware = "USDE"

	// LawofFlorida -
	LegalSystemsLawofFlorida = "USFL"

	// LawofGeorgiaUSA -
	LegalSystemsLawofGeorgiaUSA = "USGA"

	// LawofHawaii -
	LegalSystemsLawofHawaii = "USHI"

	// LawofIowa -
	LegalSystemsLawofIowa = "USIA"

	// LawofIdaho -
	LegalSystemsLawofIdaho = "USID"

	// LawofIllinois -
	LegalSystemsLawofIllinois = "USIL"

	// LawofIndiana -
	LegalSystemsLawofIndiana = "USIN"

	// LawofKansas -
	LegalSystemsLawofKansas = "USKS"

	// LawofKentucky -
	LegalSystemsLawofKentucky = "USKY"

	// LawofLouisiana -
	LegalSystemsLawofLouisiana = "USLA"

	// LawofMassachusetts -
	LegalSystemsLawofMassachusetts = "USMA"

	// LawofMaryland -
	LegalSystemsLawofMaryland = "USMD"

	// LawofMaine -
	LegalSystemsLawofMaine = "USME"

	// LawofMichigan -
	LegalSystemsLawofMichigan = "USMI"

	// LawofMinnesota -
	LegalSystemsLawofMinnesota = "USMN"

	// LawofMissouri -
	LegalSystemsLawofMissouri = "USMO"

	// LawofMississippi -
	LegalSystemsLawofMississippi = "USMS"

	// LawofMontana -
	LegalSystemsLawofMontana = "USMT"

	// LawofNorthCarolina -
	LegalSystemsLawofNorthCarolina = "USNC"

	// LawofNorthDakota -
	LegalSystemsLawofNorthDakota = "USND"

	// LawofNebraska -
	LegalSystemsLawofNebraska = "USNE"

	// LawofNewHampshire -
	LegalSystemsLawofNewHampshire = "USNH"

	// LawofNewJersey -
	LegalSystemsLawofNewJersey = "USNJ"

	// LawofNewMexico -
	LegalSystemsLawofNewMexico = "USNM"

	// LawofNevada -
	LegalSystemsLawofNevada = "USNV"

	// LawofNewYork -
	LegalSystemsLawofNewYork = "USNY"

	// LawofOhio -
	LegalSystemsLawofOhio = "USOH"

	// LawofOklahoma -
	LegalSystemsLawofOklahoma = "USOK"

	// LawofOregon -
	LegalSystemsLawofOregon = "USOR"

	// LawofPennsylvania -
	LegalSystemsLawofPennsylvania = "USPA"

	// LawofRhodeIsland -
	LegalSystemsLawofRhodeIsland = "USRI"

	// LawofSouthCarolina -
	LegalSystemsLawofSouthCarolina = "USSC"

	// LawofSouthDakota -
	LegalSystemsLawofSouthDakota = "USSD"

	// LawofTennessee -
	LegalSystemsLawofTennessee = "USTN"

	// LawofTexas -
	LegalSystemsLawofTexas = "USTX"

	// LawofUtah -
	LegalSystemsLawofUtah = "USUT"

	// LawofVirginia -
	LegalSystemsLawofVirginia = "USVA"

	// LawofVermont -
	LegalSystemsLawofVermont = "USVT"

	// LawofWashington -
	LegalSystemsLawofWashington = "USWA"

	// LawofWisconsin -
	LegalSystemsLawofWisconsin = "USWI"

	// LawofWestVirginia -
	LegalSystemsLawofWestVirginia = "USWV"

	// LawofWyoming -
	LegalSystemsLawofWyoming = "USwy"

	// LawofUzbekistan -
	LegalSystemsLawofUzbekistan = "UZB"

	// LawofVanuatu -
	LegalSystemsLawofVanuatu = "VUT"

	// LawofVaticanCity -
	LegalSystemsLawofVaticanCity = "VAT"

	// LawofVenezuela -
	LegalSystemsLawofVenezuela = "VEN"

	// LawofVietnam -
	LegalSystemsLawofVietnam = "VNM"

	// LawofWallisandFutunaFrance -
	LegalSystemsLawofWallisandFutunaFrance = "WLF"

	// LawofYemen -
	LegalSystemsLawofYemen = "YEM"

	// LawofZambia -
	LegalSystemsLawofZambia = "ZMB"

	// LawofZimbabwe -
	LegalSystemsLawofZimbabwe = "ZWE"
)

type LegalSystemsCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// LegalSystemsData holds a mapping of LegalSystems codes.
func LegalSystemsData(code string) *LegalSystemsCode {
	switch code {

	case LegalSystemsLawofEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofEnglandandWales",
			Label:       "Law of England and Wales",
			Description: "",
			MetaData:    `{"description":"Primarily common law, with early Roman and some modern continental European influences.","type":"Common law"}`,
		}

	case LegalSystemsLawofScotland:
		return &LegalSystemsCode{
			Name:        "LawofScotland",
			Label:       "Law of Scotland",
			Description: "",
			MetaData:    `{"description":"Based on Roman and continental law, with common law elements dating back to the High Middle Ages.","type":"Common law"}`,
		}

	case LegalSystemsLawofAfghanistan:
		return &LegalSystemsCode{
			Name:        "LawofAfghanistan",
			Label:       "Law of Afghanistan",
			Description: "",
			MetaData:    `{"description":"Islamic law \u0026 American/British law after invasion.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofAlbania:
		return &LegalSystemsCode{
			Name:        "LawofAlbania",
			Label:       "Law of Albania",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law. The Civil Code of the Republic of Albania, 1991.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAlgeria:
		return &LegalSystemsCode{
			Name:        "LawofAlgeria",
			Label:       "Law of Algeria",
			Description: "",
			MetaData:    `{"description":"The Algerian judicial system is based on a civil law system with codes adapted from the French legal system. Personal status laws are based on Islamic law.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofAmericanSamoa:
		return &LegalSystemsCode{
			Name:        "LawofAmericanSamoa",
			Label:       "Law of American Samoa",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		}

	case LegalSystemsLawofAndorra:
		return &LegalSystemsCode{
			Name:        "LawofAndorra",
			Label:       "Law of Andorra",
			Description: "",
			MetaData:    `{"description":"Courts apply the customary laws of Andorra, supplemented with Roman law and customary Catalan law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAngola:
		return &LegalSystemsCode{
			Name:        "LawofAngola",
			Label:       "Law of Angola",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAnguilla:
		return &LegalSystemsCode{
			Name:        "LawofAnguilla",
			Label:       "Law of Anguilla",
			Description: "",
			MetaData:    `{"description":"A combination of common law and statute, and is based heavily upon English law.","type":"Common law and statute"}`,
		}

	case LegalSystemsLawofAntiguaandBarbuda:
		return &LegalSystemsCode{
			Name:        "LawofAntiguaandBarbuda",
			Label:       "Law of Antigua and Barbuda",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofArgentina:
		return &LegalSystemsCode{
			Name:        "LawofArgentina",
			Label:       "Law of Argentina",
			Description: "",
			MetaData:    `{"description":"The Legal system of Argentina is a Civil law legal system. The pillar of the Civil system is the Constitution of Argentina (1853).","type":"Civil law"}`,
		}

	case LegalSystemsLawofArmenia:
		return &LegalSystemsCode{
			Name:        "LawofArmenia",
			Label:       "Law of Armenia",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law and traditional Armenian law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAruba:
		return &LegalSystemsCode{
			Name:        "LawofAruba",
			Label:       "Law of Aruba",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAustralia:
		return &LegalSystemsCode{
			Name:        "LawofAustralia",
			Label:       "Law of Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewSouthWales:
		return &LegalSystemsCode{
			Name:        "LawofNewSouthWales",
			Label:       "Law of New South Wales",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofQueensland:
		return &LegalSystemsCode{
			Name:        "LawofQueensland",
			Label:       "Law of Queensland",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofSouthAustralia:
		return &LegalSystemsCode{
			Name:        "LawofSouthAustralia",
			Label:       "Law of South Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofTasmania:
		return &LegalSystemsCode{
			Name:        "LawofTasmania",
			Label:       "Law of Tasmania",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofVictoria:
		return &LegalSystemsCode{
			Name:        "LawofVictoria",
			Label:       "Law of Victoria",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWesternAustralia:
		return &LegalSystemsCode{
			Name:        "LawofWesternAustralia",
			Label:       "Law of Western Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofAustralianCapitalTerritory:
		return &LegalSystemsCode{
			Name:        "LawofAustralianCapitalTerritory",
			Label:       "Law of Australian Capital Territory",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNorthernTerritory:
		return &LegalSystemsCode{
			Name:        "LawofNorthernTerritory",
			Label:       "Law of Northern Territory",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofJervisBayTerritoryAustralianCapitalTerritory:
		return &LegalSystemsCode{
			Name:        "LawofJervisBayTerritoryAustralianCapitalTerritory",
			Label:       "Law of Jervis Bay Territory (Australian Capital Territory)",
			Description: "",
			MetaData:    `{"description":"Although the Jervis Bay Territory is not part of the Australian Capital Territory, the laws of the ACT apply.","type":"Common law"}`,
		}

	case LegalSystemsLawofAustria:
		return &LegalSystemsCode{
			Name:        "LawofAustria",
			Label:       "Law of Austria",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofAzerbaijan:
		return &LegalSystemsCode{
			Name:        "LawofAzerbaijan",
			Label:       "Law of Azerbaijan",
			Description: "",
			MetaData:    `{"description":"Based on German, French, Russian and traditional Azerbaijani Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBahamas:
		return &LegalSystemsCode{
			Name:        "LawofBahamas",
			Label:       "Law of Bahamas",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBahrain:
		return &LegalSystemsCode{
			Name:        "LawofBahrain",
			Label:       "Law of Bahrain",
			Description: "",
			MetaData:    `{"description":"Bahrain's legal system is based on civil law, customs, and Sharia'a (Islamic jurisprudence). The system mostly derives from the legal system of Egypt, which is based on the legal system of France.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofBangladesh:
		return &LegalSystemsCode{
			Name:        "LawofBangladesh",
			Label:       "Law of Bangladesh",
			Description: "",
			MetaData:    `{"description":"Based on English common law, with family law heavily based on Shar'iah law.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofBarbados:
		return &LegalSystemsCode{
			Name:        "LawofBarbados",
			Label:       "Law of Barbados",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBelarus:
		return &LegalSystemsCode{
			Name:        "LawofBelarus",
			Label:       "Law of Belarus",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBelgium:
		return &LegalSystemsCode{
			Name:        "LawofBelgium",
			Label:       "Law of Belgium",
			Description: "",
			MetaData:    `{"description":"The Napoleonic Code is still in use, although it is heavily modified (especially concerning family law).","type":"Civil law"}`,
		}

	case LegalSystemsLawofBelize:
		return &LegalSystemsCode{
			Name:        "LawofBelize",
			Label:       "Law of Belize",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBenin:
		return &LegalSystemsCode{
			Name:        "LawofBenin",
			Label:       "Law of Benin",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBermuda:
		return &LegalSystemsCode{
			Name:        "LawofBermuda",
			Label:       "Law of Bermuda",
			Description: "",
			MetaData:    `{"description":"Based on the common law legal system of Bermuda.","type":"Common law"}`,
		}

	case LegalSystemsLawofBhutan:
		return &LegalSystemsCode{
			Name:        "LawofBhutan",
			Label:       "Law of Bhutan",
			Description: "",
			MetaData:    `{"description":"Based on English common law, with Indian influence. Religious law influences personal law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBolivia:
		return &LegalSystemsCode{
			Name:        "LawofBolivia",
			Label:       "Law of Bolivia",
			Description: "",
			MetaData:    `{"description":"Influenced by the Napoleonic Code.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBonaireStEustasuisandSaba:
		return &LegalSystemsCode{
			Name:        "LawofBonaireStEustasuisandSaba",
			Label:       "Law of Bonaire St Eustasuis and Saba",
			Description: "",
			MetaData:    `{"description":"The main body of civil law is the Civil Code. In the years to come the applicable former Netherlands Antilles legislation on St. Eustatius will gradually be replaced by Dutch legislation.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBosniaandHerzegovina:
		return &LegalSystemsCode{
			Name:        "LawofBosniaandHerzegovina",
			Label:       "Law of Bosnia and Herzegovina",
			Description: "",
			MetaData:    `{"description":"Influenced by Austrian law. The Swiss civil law (Zivilgesetzbuch) was a model for the Law on Obligations of 1978.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBotswana:
		return &LegalSystemsCode{
			Name:        "LawofBotswana",
			Label:       "Law of Botswana",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1891 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to the Bechuanaland Protectorate (now Botswana).","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofBrazil:
		return &LegalSystemsCode{
			Name:        "LawofBrazil",
			Label:       "Law of Brazil",
			Description: "",
			MetaData:    `{"description":"Based on German, Italian, French and Portuguese law. However, in 2004 the Federal Constitution was amended to grant the Supreme Federal Court authority to issue binding precedents (súmulas vinculantes) to settle controversies involving constitutional law - a mechanism that echoes the stare decisis principle typically found in common law systems.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBritishVirginIslands:
		return &LegalSystemsCode{
			Name:        "LawofBritishVirginIslands",
			Label:       "Law of British Virgin Islands",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBrunei:
		return &LegalSystemsCode{
			Name:        "LawofBrunei",
			Label:       "Law of Brunei",
			Description: "",
			MetaData:    `{"description":"Most laws under Common Law and the Sharia Penal Code apply to all people in Brunei, regardless of nationality or religion.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofBulgaria:
		return &LegalSystemsCode{
			Name:        "LawofBulgaria",
			Label:       "Law of Bulgaria",
			Description: "",
			MetaData:    `{"description":"Civil Law system influenced by Germanic and Roman law systems.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBurkinaFaso:
		return &LegalSystemsCode{
			Name:        "LawofBurkinaFaso",
			Label:       "Law of Burkina Faso",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofBurundi:
		return &LegalSystemsCode{
			Name:        "LawofBurundi",
			Label:       "Law of Burundi",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of Belgian civil law and customary law.","type":"Civil law and customary law"}`,
		}

	case LegalSystemsLawofCambodia:
		return &LegalSystemsCode{
			Name:        "LawofCambodia",
			Label:       "Law of Cambodia",
			Description: "",
			MetaData:    `{"description":"The Cambodian legal system is based largely on the French civil system, and is statute based.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCameroon:
		return &LegalSystemsCode{
			Name:        "LawofCameroon",
			Label:       "Law of Cameroon",
			Description: "",
			MetaData:    `{"description":"Cameroon is a bijural system with the English Common Law operating in the two Anglophone regions of North West and South West and the French Civil Law operating in the eight francophone regions of Adamaoua, Centre, East, Far North, Littoral, North, West and South.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofCanada:
		return &LegalSystemsCode{
			Name:        "LawofCanada",
			Label:       "Law of Canada",
			Description: "",
			MetaData:    `{"description":"Based on English common law, except in Quebec Quebec, where a civil law system based on French law prevails in most matters of a civil nature, such as obligations (contract and delict), property law, family law and private matters. Federal statutes take into account the bijuridical nature of Canada and use both common law and civil law terms where appropriate.","type":"Common law"}`,
		}

	case LegalSystemsLawofAlberta:
		return &LegalSystemsCode{
			Name:        "LawofAlberta",
			Label:       "Law of Alberta",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofBritishColumbia:
		return &LegalSystemsCode{
			Name:        "LawofBritishColumbia",
			Label:       "Law of British Columbia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofManitoba:
		return &LegalSystemsCode{
			Name:        "LawofManitoba",
			Label:       "Law of Manitoba",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewBrunswick:
		return &LegalSystemsCode{
			Name:        "LawofNewBrunswick",
			Label:       "Law of New Brunswick",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewfoundlandandLabrador:
		return &LegalSystemsCode{
			Name:        "LawofNewfoundlandandLabrador",
			Label:       "Law of Newfoundland and Labrador",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNovaScotia:
		return &LegalSystemsCode{
			Name:        "LawofNovaScotia",
			Label:       "Law of Nova Scotia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofOntario:
		return &LegalSystemsCode{
			Name:        "LawofOntario",
			Label:       "Law of Ontario",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofPrinceEdwardIsland:
		return &LegalSystemsCode{
			Name:        "LawofPrinceEdwardIsland",
			Label:       "Law of Prince Edward Island",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofQuebec:
		return &LegalSystemsCode{
			Name:        "LawofQuebec",
			Label:       "Law of Quebec",
			Description: "",
			MetaData:    `{"description":"The Civil Code of Quebec is the civil code currently in force. Canadian (federal) criminal law in force in Quebec is based on common law, but federal statutes of or relating to private law take into account the bijuridical nature of Canada and use both common law and civil law terms where appropriate.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSaskatchewan:
		return &LegalSystemsCode{
			Name:        "LawofSaskatchewan",
			Label:       "Law of Saskatchewan",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNorthwestTerritories:
		return &LegalSystemsCode{
			Name:        "LawofNorthwestTerritories",
			Label:       "Law of Northwest Territories",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNunavut:
		return &LegalSystemsCode{
			Name:        "LawofNunavut",
			Label:       "Law of Nunavut",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofYukon:
		return &LegalSystemsCode{
			Name:        "LawofYukon",
			Label:       "Law of Yukon",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofCapeVerde:
		return &LegalSystemsCode{
			Name:        "LawofCapeVerde",
			Label:       "Law of Cape Verde",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCaymanIslands:
		return &LegalSystemsCode{
			Name:        "LawofCaymanIslands",
			Label:       "Law of Cayman Islands",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofCentralAfricanRepublic:
		return &LegalSystemsCode{
			Name:        "LawofCentralAfricanRepublic",
			Label:       "Law of Central African Republic",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofChad:
		return &LegalSystemsCode{
			Name:        "LawofChad",
			Label:       "Law of Chad",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofChile:
		return &LegalSystemsCode{
			Name:        "LawofChile",
			Label:       "Law of Chile",
			Description: "",
			MetaData:    `{"description":"Chile's legal system is civil law based. It is primarily based on the Chilean Civil Code of 1855, derived from Spanish law and other codes of Continental Europe of the 19th century.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPeoplesRepublicofChina:
		return &LegalSystemsCode{
			Name:        "LawofPeoplesRepublicofChina",
			Label:       "Law of People's Republic of China",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law with influences from the Soviet Socialist law from Soviet Union.","type":"Civil law"}`,
		}

	case LegalSystemsLawofChristmasIslandWesternAustralia:
		return &LegalSystemsCode{
			Name:        "LawofChristmasIslandWesternAustralia",
			Label:       "Law of Christmas Island (Western Australia)",
			Description: "",
			MetaData:    `{"description":"The Christmas Island Act - applies Western Australian laws on Christmas Island, including the Local Government Act 1995 (WA)","type":"Common law"}`,
		}

	case LegalSystemsLawofCocosIslandsWesternAustralia:
		return &LegalSystemsCode{
			Name:        "LawofCocosIslandsWesternAustralia",
			Label:       "Law of Cocos Islands (Western Australia)",
			Description: "",
			MetaData:    `{"description":"The Territories Law Reform Act 1992 (Cth) amended the Cocos (Keeling) Islands Act to introduce a modern body of Australian law to the Cocos (Keeling) Islands, including provisions to enliven the application of most Australian Government laws and regulations and to apply Western Australian laws to the territory.","type":"Common law"}`,
		}

	case LegalSystemsLawofColombia:
		return &LegalSystemsCode{
			Name:        "LawofColombia",
			Label:       "Law of Colombia",
			Description: "",
			MetaData:    `{"description":"Based on the Chilean Civil Law. Civil code introduced in 1873. Nearly faithful reproduction of the Chilean civil code.","type":"Civil law"}`,
		}

	case LegalSystemsLawofComoros:
		return &LegalSystemsCode{
			Name:        "LawofComoros",
			Label:       "Law of Comoros",
			Description: "",
			MetaData:    `{"description":"The Comorian legal system rests on Islamic law and an inherited French legal code. Village elders or civilian courts settle most disputes.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofRepublicoftheCongo:
		return &LegalSystemsCode{
			Name:        "LawofRepublicoftheCongo",
			Label:       "Law of Republic of the Congo",
			Description: "",
			MetaData:    `{"description":"Based on the Napoleonic Civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofDemocraticRepublicoftheCongo:
		return &LegalSystemsCode{
			Name:        "LawofDemocraticRepublicoftheCongo",
			Label:       "Law of Democratic Republic of the Congo",
			Description: "",
			MetaData:    `{"description":"Based on Belgian civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCookIslands:
		return &LegalSystemsCode{
			Name:        "LawofCookIslands",
			Label:       "Law of Cook Islands",
			Description: "",
			MetaData:    `{"description":"Constitution of the Cook Islands - the supreme law. English common law and equity - except so far as inconsistent with the Cook Islands Act 1915 (NZ) or inappropriate to the circumstances of the country or inconsistent with the Constitution.","type":"Common law"}`,
		}

	case LegalSystemsLawofCostaRica:
		return &LegalSystemsCode{
			Name:        "LawofCostaRica",
			Label:       "Law of Costa Rica",
			Description: "",
			MetaData:    `{"description":"The present Civil Code went into effect 1 January 1888, and was influenced by the Napoleonic Code and the Spanish Civil Code of 1889.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCroatia:
		return &LegalSystemsCode{
			Name:        "LawofCroatia",
			Label:       "Law of Croatia",
			Description: "",
			MetaData:    `{"description":"Based on the Germanic Civil Law. Croatian Law system is largely influenced by German and Austrian law systems.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCuba:
		return &LegalSystemsCode{
			Name:        "LawofCuba",
			Label:       "Law of Cuba",
			Description: "",
			MetaData:    `{"description":"Influenced by Spanish and American law with large elements of Communist legal theory.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCuracao:
		return &LegalSystemsCode{
			Name:        "LawofCuracao",
			Label:       "Law of Curacao",
			Description: "",
			MetaData:    `{"description":"Based on Dutch Civil Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCyprus:
		return &LegalSystemsCode{
			Name:        "LawofCyprus",
			Label:       "Law of Cyprus",
			Description: "",
			MetaData:    `{"description":"Based on English common law as inherited from British colonisation, with civil law influences, particularly in criminal law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofCzechRepublic:
		return &LegalSystemsCode{
			Name:        "LawofCzechRepublic",
			Label:       "Law of Czech Republic",
			Description: "",
			MetaData:    `{"description":"Based on Belgian civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofDenmark:
		return &LegalSystemsCode{
			Name:        "LawofDenmark",
			Label:       "Law of Denmark",
			Description: "",
			MetaData:    `{"description":"Based on North Germanic law. Scandinavian-North Germanic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofDjibouti:
		return &LegalSystemsCode{
			Name:        "LawofDjibouti",
			Label:       "Law of Djibouti",
			Description: "",
			MetaData:    `{"description":"Djibouti's legal and judicial system is largely inspired by the French legislation. Laws are codified in which the system of the country is based on the coexistence of Islamic law (Sharia), customary law and civil law inherited the French Napoleon Code.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofDominica:
		return &LegalSystemsCode{
			Name:        "LawofDominica",
			Label:       "Law of Dominica",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofDominicanRepublic:
		return &LegalSystemsCode{
			Name:        "LawofDominicanRepublic",
			Label:       "Law of Dominican Republic",
			Description: "",
			MetaData:    `{"description":"Based on the Napoleonic Code.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTimorLeste:
		return &LegalSystemsCode{
			Name:        "LawofTimorLeste",
			Label:       "Law of Timor-Leste",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEcuador:
		return &LegalSystemsCode{
			Name:        "LawofEcuador",
			Label:       "Law of Ecuador",
			Description: "",
			MetaData:    `{"description":"Based on the Chilean civil law. Civil code introduced in 1861.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEgypt:
		return &LegalSystemsCode{
			Name:        "LawofEgypt",
			Label:       "Law of Egypt",
			Description: "",
			MetaData:    `{"description":"Family Law (personal Statute) for Muslims based on Islamic Jurisprudence, Separate Personal Statute for non Muslims, and all other branches of Law are based on French civil law system.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofElSalvador:
		return &LegalSystemsCode{
			Name:        "LawofElSalvador",
			Label:       "Law of El Salvador",
			Description: "",
			MetaData:    `{"description":"Civil law system with minor common law influence; judicial review of legislative acts in the Supreme Court.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEquatorialGuinea:
		return &LegalSystemsCode{
			Name:        "LawofEquatorialGuinea",
			Label:       "Law of Equatorial Guinea",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEritrea:
		return &LegalSystemsCode{
			Name:        "LawofEritrea",
			Label:       "Law of Eritrea",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil, customary, and Islamic religious law.","type":"Civil law and sharia law and Customary law"}`,
		}

	case LegalSystemsLawofEstonia:
		return &LegalSystemsCode{
			Name:        "LawofEstonia",
			Label:       "Law of Estonia",
			Description: "",
			MetaData:    `{"description":"Based on German civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEthiopia:
		return &LegalSystemsCode{
			Name:        "LawofEthiopia",
			Label:       "Law of Ethiopia",
			Description: "",
			MetaData:    `{"description":"Ethiopia's legal system largely belongs in the civil law system. Conversely, the case law system is a distinguishing feature of the common law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofFalklandIslandsEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofFalklandIslandsEnglandandWales",
			Label:       "Law of Falkland Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law and local statutes.","type":"Common law"}`,
		}

	case LegalSystemsLawofFaroeIslandsDenmark:
		return &LegalSystemsCode{
			Name:        "LawofFaroeIslandsDenmark",
			Label:       "Law of Faroe Islands (Denmark)",
			Description: "",
			MetaData:    `{"description":"The laws of Denmark apply where applicable. The Faroe Islands have the exclusive right to legislate and govern independently in a wide range of areas.  The Faroese court system is under the jurisdiction of the high courts in Denmark.","type":"Civil law"}`,
		}

	case LegalSystemsLawofFiji:
		return &LegalSystemsCode{
			Name:        "LawofFiji",
			Label:       "Law of Fiji",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofFinland:
		return &LegalSystemsCode{
			Name:        "LawofFinland",
			Label:       "Law of Finland",
			Description: "",
			MetaData:    `{"description":"Based on Nordic law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofFrance:
		return &LegalSystemsCode{
			Name:        "LawofFrance",
			Label:       "Law of France",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic code (code civil of 1804).","type":"Civil law"}`,
		}

	case LegalSystemsLawofFrenchGuianaFrance:
		return &LegalSystemsCode{
			Name:        "LawofFrenchGuianaFrance",
			Label:       "Law of French Guiana (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		}

	case LegalSystemsLawofFrenchPolynesiaFrance:
		return &LegalSystemsCode{
			Name:        "LawofFrenchPolynesiaFrance",
			Label:       "Law of French Polynesia (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		}

	case LegalSystemsLawofFrenchSouthernandAntarcticLandsFrance:
		return &LegalSystemsCode{
			Name:        "LawofFrenchSouthernandAntarcticLandsFrance",
			Label:       "Law of French Southern and Antarctic Lands (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGabon:
		return &LegalSystemsCode{
			Name:        "LawofGabon",
			Label:       "Law of Gabon",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTheGambia:
		return &LegalSystemsCode{
			Name:        "LawofTheGambia",
			Label:       "Law of The Gambia",
			Description: "",
			MetaData:    `{"description":"English common law, Islamic law and customary law.","type":"Common law and sharia law and Customary law"}`,
		}

	case LegalSystemsLawofGeorgia:
		return &LegalSystemsCode{
			Name:        "LawofGeorgia",
			Label:       "Law of Georgia",
			Description: "",
			MetaData:    `{"description":"Civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGermany:
		return &LegalSystemsCode{
			Name:        "LawofGermany",
			Label:       "Law of Germany",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law. The Bürgerliches Gesetzbuch of 1900 (\"BGB\"). The BGB is influenced both by Roman and German law traditions.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGhana:
		return &LegalSystemsCode{
			Name:        "LawofGhana",
			Label:       "Law of Ghana",
			Description: "",
			MetaData:    `{"description":"Mixed system of English common law and customary law.","type":"Common law and customary law"}`,
		}

	case LegalSystemsLawofGibraltar:
		return &LegalSystemsCode{
			Name:        "LawofGibraltar",
			Label:       "Law of Gibraltar",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofGreece:
		return &LegalSystemsCode{
			Name:        "LawofGreece",
			Label:       "Law of Greece",
			Description: "",
			MetaData:    `{"description":"Greece is a civil law country.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGreenlandDenmark:
		return &LegalSystemsCode{
			Name:        "LawofGreenlandDenmark",
			Label:       "Law of Greenland (Denmark)",
			Description: "",
			MetaData:    `{"description":"The laws of Denmark apply where applicable and Greenlandic law applies to other areas.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGrenada:
		return &LegalSystemsCode{
			Name:        "LawofGrenada",
			Label:       "Law of Grenada",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofGuadeloupeFrance:
		return &LegalSystemsCode{
			Name:        "LawofGuadeloupeFrance",
			Label:       "Law of Guadeloupe (France)",
			Description: "",
			MetaData:    `{"description":"Guadeloupe honors the French Constitution of September 1958, and utilizes the French legal system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGuamUnitedStates:
		return &LegalSystemsCode{
			Name:        "LawofGuamUnitedStates",
			Label:       "Law of Guam (United States)",
			Description: "",
			MetaData:    `{"description":"The legal and judicial system in Guam is the same as in any other U.S. state or territory. All proceedings are in English, and based on the American common law legal system and US standards of integrity.","type":"Common law"}`,
		}

	case LegalSystemsLawofGuatemala:
		return &LegalSystemsCode{
			Name:        "LawofGuatemala",
			Label:       "Law of Guatemala",
			Description: "",
			MetaData:    `{"description":"Guatemala is a civil-law country.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGuernsey:
		return &LegalSystemsCode{
			Name:        "LawofGuernsey",
			Label:       "Law of Guernsey",
			Description: "",
			MetaData:    `{"description":"Customary legal system based on Norman customary law; includes elements of the French civil code and English common law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofGuinea:
		return &LegalSystemsCode{
			Name:        "LawofGuinea",
			Label:       "Law of Guinea",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system, customary law, and decree.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGuineaBissau:
		return &LegalSystemsCode{
			Name:        "LawofGuineaBissau",
			Label:       "Law of Guinea-Bissau",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGuyana:
		return &LegalSystemsCode{
			Name:        "LawofGuyana",
			Label:       "Law of Guyana",
			Description: "",
			MetaData:    `{"description":"Vestiges of a Dutch legal system remain, particularly in the area of land tenure. However, the common law of Britain is the basis for the legal system of Guyana.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofHaiti:
		return &LegalSystemsCode{
			Name:        "LawofHaiti",
			Label:       "Law of Haiti",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory:
		return &LegalSystemsCode{
			Name:        "LawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory",
			Label:       "Law of Heard Island and McDonalds Islands (Australian Capital Territory)",
			Description: "",
			MetaData:    `{"description":"The administration of the territory is established in the Heard Island and McDonald Islands Act 1953, which places it under the laws of the Australian Capital Territory and the jurisdiction of the Supreme Court of the Australian Capital Territory.","type":"Common law"}`,
		}

	case LegalSystemsLawofHonduras:
		return &LegalSystemsCode{
			Name:        "LawofHonduras",
			Label:       "Law of Honduras",
			Description: "",
			MetaData:    `{"description":"Because Honduras is a Civil law country, there is not much role for precedent or judicial review.","type":"Civil law"}`,
		}

	case LegalSystemsLawofHongKong:
		return &LegalSystemsCode{
			Name:        "LawofHongKong",
			Label:       "Law of Hong Kong",
			Description: "",
			MetaData:    `{"description":"Principally based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofHungary:
		return &LegalSystemsCode{
			Name:        "LawofHungary",
			Label:       "Law of Hungary",
			Description: "",
			MetaData:    `{"description":"Based on Germanic, codified Roman law with elements from Napoleonic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofIceland:
		return &LegalSystemsCode{
			Name:        "LawofIceland",
			Label:       "Law of Iceland",
			Description: "",
			MetaData:    `{"description":"Based on North Germanic law. Germanic traditional laws and influenced by Medieval Norwegian and Danish laws.","type":"Civil law"}`,
		}

	case LegalSystemsLawofIndia:
		return &LegalSystemsCode{
			Name:        "LawofIndia",
			Label:       "Law of India",
			Description: "",
			MetaData:    `{"description":"Based on English common law, except in Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law. Muslim personal law based on sharia law applies to Muslims. Exceptions for Muslims in Goa state, where the Goa Civil Code applies to all persons irrespective of religion, and for Muslims who marry under the Special Marriage Act, 1954.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofDadraandNagarHaveli:
		return &LegalSystemsCode{
			Name:        "LawofDadraandNagarHaveli",
			Label:       "Law of Dadra and Nagar Haveli",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofDamanandDiu:
		return &LegalSystemsCode{
			Name:        "LawofDamanandDiu",
			Label:       "Law of Daman and Diu",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofGoa:
		return &LegalSystemsCode{
			Name:        "LawofGoa",
			Label:       "Law of Goa",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofIndonesia:
		return &LegalSystemsCode{
			Name:        "LawofIndonesia",
			Label:       "Law of Indonesia",
			Description: "",
			MetaData:    `{"description":"Based on a civil law system, intermixed with customary law and the Roman Dutch law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofIran:
		return &LegalSystemsCode{
			Name:        "LawofIran",
			Label:       "Law of Iran",
			Description: "",
			MetaData:    `{"description":"Religious legal system based on secular and Islamic law.","type":"Religious law"}`,
		}

	case LegalSystemsLawofIraq:
		return &LegalSystemsCode{
			Name:        "LawofIraq",
			Label:       "Law of Iraq",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil and Islamic law.","type":"Civil law and Religious law"}`,
		}

	case LegalSystemsLawofIreland:
		return &LegalSystemsCode{
			Name:        "LawofIreland",
			Label:       "Law of Ireland",
			Description: "",
			MetaData:    `{"description":"Based on Irish law before 1922, which was itself based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsIsleofMan:
		return &LegalSystemsCode{
			Name:        "IsleofMan",
			Label:       "Isle of Man",
			Description: "",
			MetaData:    `{"description":"The legal system on the Isle of Man is Manx customary law, a form of common law. Manx law originally derived from Gaelic Brehon law and Norse Udal law.","type":"Common law"}`,
		}

	case LegalSystemsLawofIsrael:
		return &LegalSystemsCode{
			Name:        "LawofIsrael",
			Label:       "Law of Israel",
			Description: "",
			MetaData:    `{"description":"Based on English common law arising from the period of the British Mandate (which includes laws arising from previous Ottoman rule), also incorporating civil law and fragments of Halakha and Sharia for family law cases.","type":"Common law"}`,
		}

	case LegalSystemsLawofItaly:
		return &LegalSystemsCode{
			Name:        "LawofItaly",
			Label:       "Law of Italy",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law, with elements of the Napoleonic civil code; civil code of 1942 replaced the original one of 1865.","type":"Civil law"}`,
		}

	case LegalSystemsLawofCotedIvoire:
		return &LegalSystemsCode{
			Name:        "LawofCotedIvoire",
			Label:       "Law of Cote d'Ivoire",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofJamaica:
		return &LegalSystemsCode{
			Name:        "LawofJamaica",
			Label:       "Law of Jamaica",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofJapan:
		return &LegalSystemsCode{
			Name:        "LawofJapan",
			Label:       "Law of Japan",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law. Japanese civil code of 1895.","type":"Civil law"}`,
		}

	case LegalSystemsLawofJersey:
		return &LegalSystemsCode{
			Name:        "LawofJersey",
			Label:       "Law of Jersey",
			Description: "",
			MetaData:    `{"description":"The Bailiwick of Jersey's legal system draws on local legislation enacted by the States of Jersey, Norman customary law, English common law and modern French civil law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofJordan:
		return &LegalSystemsCode{
			Name:        "LawofJordan",
			Label:       "Law of Jordan",
			Description: "",
			MetaData:    `{"description":"Mainly based on French Civil Code and Ottoman Majalla, Islamic law applicable to family law.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofKazakhstan:
		return &LegalSystemsCode{
			Name:        "LawofKazakhstan",
			Label:       "Law of Kazakhstan",
			Description: "",
			MetaData:    `{"description":"The legal system of Kazakhstan is based on civil law and founded-on statutory legislation (the civil code), which provides for a hierarchy of legal acts.","type":"Civil law"}`,
		}

	case LegalSystemsLawofKenya:
		return &LegalSystemsCode{
			Name:        "LawofKenya",
			Label:       "Law of Kenya",
			Description: "",
			MetaData:    `{"description":"The Kenyan legal system is descended from the British Common Law system.","type":"Common law"}`,
		}

	case LegalSystemsLawofKiribati:
		return &LegalSystemsCode{
			Name:        "LawofKiribati",
			Label:       "Law of Kiribati",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofKuwait:
		return &LegalSystemsCode{
			Name:        "LawofKuwait",
			Label:       "Law of Kuwait",
			Description: "",
			MetaData:    `{"description":"Follows the \"civil law system\" modeled on the French legal system. Sharia law governs only family law for Muslim residents.","type":"Civil law and Sharia Law"}`,
		}

	case LegalSystemsLawofKyrgyzstan:
		return &LegalSystemsCode{
			Name:        "LawofKyrgyzstan",
			Label:       "Law of Kyrgyzstan",
			Description: "",
			MetaData:    `{"description":"Civil law system, which includes features of French civil law and Russian Federation laws.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLaos:
		return &LegalSystemsCode{
			Name:        "LawofLaos",
			Label:       "Law of Laos",
			Description: "",
			MetaData:    `{"description":"Inherited a typical civil law-based legal system from the French colonial administrators.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLatvia:
		return &LegalSystemsCode{
			Name:        "LawofLatvia",
			Label:       "Law of Latvia",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic and German civil law, as it was historically before the Soviet occupation.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLebanon:
		return &LegalSystemsCode{
			Name:        "LawofLebanon",
			Label:       "Law of Lebanon",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLesotho:
		return &LegalSystemsCode{
			Name:        "LawofLesotho",
			Label:       "Law of Lesotho",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1884 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to Basutoland (now Lesotho).","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofLiberia:
		return &LegalSystemsCode{
			Name:        "LawofLiberia",
			Label:       "Law of Liberia",
			Description: "",
			MetaData:    `{"description":"Based on Anglo-American and customary law.","type":"Common law"}`,
		}

	case LegalSystemsLawofLibya:
		return &LegalSystemsCode{
			Name:        "LawofLibya",
			Label:       "Law of Libya",
			Description: "",
			MetaData:    `{"description":"Islamic law. Based on Napoleonic civil law, with Ottoman, Italian, and Egyptian sources.","type":"Civil and Religious law"}`,
		}

	case LegalSystemsLawofLiechtenstein:
		return &LegalSystemsCode{
			Name:        "LawofLiechtenstein",
			Label:       "Law of Liechtenstein",
			Description: "",
			MetaData:    `{"description":"A civil law country. The source of the law is statute, which is enacted by the parliament.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLithuania:
		return &LegalSystemsCode{
			Name:        "LawofLithuania",
			Label:       "Law of Lithuania",
			Description: "",
			MetaData:    `{"description":"Modeled after Dutch civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofLuxembourg:
		return &LegalSystemsCode{
			Name:        "LawofLuxembourg",
			Label:       "Law of Luxembourg",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMacau:
		return &LegalSystemsCode{
			Name:        "LawofMacau",
			Label:       "Law of Macau",
			Description: "",
			MetaData:    `{"description":"Based on the Portuguese civil law; also influenced by the law of the PRC.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMacedonia:
		return &LegalSystemsCode{
			Name:        "LawofMacedonia",
			Label:       "Law of Macedonia",
			Description: "",
			MetaData:    `{"description":"Civil law system; judicial review of legislative acts.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMadagascar:
		return &LegalSystemsCode{
			Name:        "LawofMadagascar",
			Label:       "Law of Madagascar",
			Description: "",
			MetaData:    `{"description":"Civil law system based on the old French civil code and customary law in matters of marriage, family, and obligation.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMalawi:
		return &LegalSystemsCode{
			Name:        "LawofMalawi",
			Label:       "Law of Malawi",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law; judicial review of legislative acts in the Supreme Court of Appeal.","type":"Common law and customary law"}`,
		}

	case LegalSystemsLawofMalaysia:
		return &LegalSystemsCode{
			Name:        "LawofMalaysia",
			Label:       "Law of Malaysia",
			Description: "",
			MetaData:    `{"description":"Based on English common law, personal law based on sharia law applies to Muslims.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofMaldives:
		return &LegalSystemsCode{
			Name:        "LawofMaldives",
			Label:       "Law of Maldives",
			Description: "",
			MetaData:    `{"description":"Based on an admixture of Islamic Law and English common law.","type":"Common law and Religious law"}`,
		}

	case LegalSystemsLawofMali:
		return &LegalSystemsCode{
			Name:        "LawofMali",
			Label:       "Law of Mali",
			Description: "",
			MetaData:    `{"description":"Derives from French civil law and customary law.","type":"Civil law and customary law"}`,
		}

	case LegalSystemsLawofMalta:
		return &LegalSystemsCode{
			Name:        "LawofMalta",
			Label:       "Law of Malta",
			Description: "",
			MetaData:    `{"description":"Napoleonic Code with influences from Italian Civil Law. English common law however is also a source of Maltese Law, most notably in Public Law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofMarshallIslands:
		return &LegalSystemsCode{
			Name:        "LawofMarshallIslands",
			Label:       "Law of Marshall Islands",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		}

	case LegalSystemsLawofMartiniqueFrance:
		return &LegalSystemsCode{
			Name:        "LawofMartiniqueFrance",
			Label:       "Law of Martinique (France)",
			Description: "",
			MetaData:    `{"description":"The French system of justice is in force.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMauritania:
		return &LegalSystemsCode{
			Name:        "LawofMauritania",
			Label:       "Law of Mauritania",
			Description: "",
			MetaData:    `{"description":"Mix of Islamic law and French Civil Codes, Islamic law largely applicable to family law.","type":"Religious law"}`,
		}

	case LegalSystemsLawofMauritius:
		return &LegalSystemsCode{
			Name:        "LawofMauritius",
			Label:       "Law of Mauritius",
			Description: "",
			MetaData:    `{"description":"Laws governing the Mauritian penal system are derived partly from French civil law and British common law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofMayotteFrance:
		return &LegalSystemsCode{
			Name:        "LawofMayotteFrance",
			Label:       "Law of Mayotte (France)",
			Description: "",
			MetaData:    `{"description":"Local judicial, economic and social laws and customs were changed to conform with French law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMexico:
		return &LegalSystemsCode{
			Name:        "LawofMexico",
			Label:       "Law of Mexico",
			Description: "",
			MetaData:    `{"description":"The origins of Mexico's legal system are both ancient and classical, based on the Roman and French legal systems.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMicronesia:
		return &LegalSystemsCode{
			Name:        "LawofMicronesia",
			Label:       "Law of Micronesia",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of common and customary law.","type":"Common law and customary law"}`,
		}

	case LegalSystemsLawofMoldova:
		return &LegalSystemsCode{
			Name:        "LawofMoldova",
			Label:       "Law of Moldova",
			Description: "",
			MetaData:    `{"description":"Civil law system with Germanic law influences; Constitutional Court review of legislative acts.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMonaco:
		return &LegalSystemsCode{
			Name:        "LawofMonaco",
			Label:       "Law of Monaco",
			Description: "",
			MetaData:    `{"description":"Has its own legal system, Law Courts and Appeal Court. Its law is largely but not entirely based on the French Code Civil.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMongolia:
		return &LegalSystemsCode{
			Name:        "LawofMongolia",
			Label:       "Law of Mongolia",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMontenegro:
		return &LegalSystemsCode{
			Name:        "LawofMontenegro",
			Label:       "Law of Montenegro",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic and German civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMontserratEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofMontserratEnglandandWales",
			Label:       "Law of Montserrat (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMorocco:
		return &LegalSystemsCode{
			Name:        "LawofMorocco",
			Label:       "Law of Morocco",
			Description: "",
			MetaData:    `{"description":"Mix of Islamic law and French Civil Codes, Islamic law largely applicable to family law. Halakha recognized to family law cases for Jewish citizens.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofMozambique:
		return &LegalSystemsCode{
			Name:        "LawofMozambique",
			Label:       "Law of Mozambique",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofMyanmar:
		return &LegalSystemsCode{
			Name:        "LawofMyanmar",
			Label:       "Law of Myanmar",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNamibia:
		return &LegalSystemsCode{
			Name:        "LawofNamibia",
			Label:       "Law of Namibia",
			Description: "",
			MetaData:    `{"description":"Based on South African law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofNauru:
		return &LegalSystemsCode{
			Name:        "LawofNauru",
			Label:       "Law of Nauru",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNepal:
		return &LegalSystemsCode{
			Name:        "LawofNepal",
			Label:       "Law of Nepal",
			Description: "",
			MetaData:    `{"description":"English common law and Hindu legal concepts; note - new criminal and civil codes came into effect on 17 August 2018","type":"Common law, Civil law and Religous law"}`,
		}

	case LegalSystemsLawofNetherlands:
		return &LegalSystemsCode{
			Name:        "LawofNetherlands",
			Label:       "Law of Netherlands",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic code with German law influence.","type":"Civil law"}`,
		}

	case LegalSystemsLawofNewCaledonia:
		return &LegalSystemsCode{
			Name:        "LawofNewCaledonia",
			Label:       "Law of New Caledonia",
			Description: "",
			MetaData:    `{"description":"Civil law system based on French civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofNewZealand:
		return &LegalSystemsCode{
			Name:        "LawofNewZealand",
			Label:       "Law of New Zealand",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNicaragua:
		return &LegalSystemsCode{
			Name:        "LawofNicaragua",
			Label:       "Law of Nicaragua",
			Description: "",
			MetaData:    `{"description":"Civil law system; Supreme Court may review administrative acts.","type":"Civil law"}`,
		}

	case LegalSystemsLawofNiger:
		return &LegalSystemsCode{
			Name:        "LawofNiger",
			Label:       "Law of Niger",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil law, based on French civil law, Islamic law, and customary law.","type":"Civil law, Religious law and Customary law"}`,
		}

	case LegalSystemsLawofNigeria:
		return &LegalSystemsCode{
			Name:        "LawofNigeria",
			Label:       "Law of Nigeria",
			Description: "",
			MetaData:    `{"description":"Sharia in the northern states, common law in the south and at the federal level.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofNiue:
		return &LegalSystemsCode{
			Name:        "LawofNiue",
			Label:       "Law of Niue",
			Description: "",
			MetaData:    `{"description":"Certain British and New Zealand Acts; common law; customary law.","type":"Common law and customary law"}`,
		}

	case LegalSystemsLawofNorfolkIslandNewSouthWales:
		return &LegalSystemsCode{
			Name:        "LawofNorfolkIslandNewSouthWales",
			Label:       "Law of Norfolk Island (New South Wales)",
			Description: "",
			MetaData:    `{"description":"The Norfolk Island Regional Council operates under an applied version of NSW Local Government legislation, the Local Government Act 1993 (NSW)(NI).","type":"Common law"}`,
		}

	case LegalSystemsLawofNorthKorea:
		return &LegalSystemsCode{
			Name:        "LawofNorthKorea",
			Label:       "Law of North Korea",
			Description: "",
			MetaData:    `{"description":"Codified civil law system, which was inherited from colonial Japan and is similar to South Korea's system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofNorthernMarianaIslandsUnitedStates:
		return &LegalSystemsCode{
			Name:        "LawofNorthernMarianaIslandsUnitedStates",
			Label:       "Law of Northern Mariana Islands (United States)",
			Description: "",
			MetaData:    `{"description":"Laws of the US apply, except for customs and some aspects of taxation.","type":"Common law"}`,
		}

	case LegalSystemsLawofNorway:
		return &LegalSystemsCode{
			Name:        "LawofNorway",
			Label:       "Law of Norway",
			Description: "",
			MetaData:    `{"description":"Based on Civil Code.","type":"Civil law"}`,
		}

	case LegalSystemsLawofOman:
		return &LegalSystemsCode{
			Name:        "LawofOman",
			Label:       "Law of Oman",
			Description: "",
			MetaData:    `{"description":"Sharia and tribal custom laws.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofPakistan:
		return &LegalSystemsCode{
			Name:        "LawofPakistan",
			Label:       "Law of Pakistan",
			Description: "",
			MetaData:    `{"description":"Based on English Common Law, some Islamic law applications in inheritance. Tribal Law in FATA.","type":"Common law and sharia law"}`,
		}

	case LegalSystemsLawofPalau:
		return &LegalSystemsCode{
			Name:        "LawofPalau",
			Label:       "Law of Palau",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		}

	case LegalSystemsLawofPalestinianTerritory:
		return &LegalSystemsCode{
			Name:        "LawofPalestinianTerritory",
			Label:       "Law of Palestinian Territory",
			Description: "",
			MetaData:    `{"description":"Islamic customary law applies in Palestinian ruled territory.","type":"Islamic customary law"}`,
		}

	case LegalSystemsLawofPanama:
		return &LegalSystemsCode{
			Name:        "LawofPanama",
			Label:       "Law of Panama",
			Description: "",
			MetaData:    `{"description":"Based on civil law with influences from Spanish legal tradition and Roman laws.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPapuaNewGuinea:
		return &LegalSystemsCode{
			Name:        "LawofPapuaNewGuinea",
			Label:       "Law of Papua New Guinea",
			Description: "",
			MetaData:    `{"description":"The legal system is based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofParaguay:
		return &LegalSystemsCode{
			Name:        "LawofParaguay",
			Label:       "Law of Paraguay",
			Description: "",
			MetaData:    `{"description":"The Paraguayan Civil Code in force since 1987 is largely influenced by the Napoleonic Code and the Argentinian Code.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPeru:
		return &LegalSystemsCode{
			Name:        "LawofPeru",
			Label:       "Law of Peru",
			Description: "",
			MetaData:    `{"description":"Based on civil law system; accepts compulsory International Court of Justice ICJ jurisdiction with despotic and corrupting reservations.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPhilippines:
		return &LegalSystemsCode{
			Name:        "LawofPhilippines",
			Label:       "Law of Philippines",
			Description: "",
			MetaData:    `{"description":"Based on Spanish law; influenced by U.S. common law after 1898 Spanish– and Philippine–American Wars, personal law based on sharia law applies to Muslims.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofPitcairnIslandsEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofPitcairnIslandsEnglandandWales",
			Label:       "Law of Pitcairn Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"Comprises (a) laws (called \"ordinances\") made by the Governor, (b) United Kingdom legislation and Orders in Council extending to Pitcairn, and (c) the common law, rules of equity, and statutes of general application as in force in England for the time being, so far as local circumstances and the limits of local jurisdiction permit, and subject to any existing or future ordinance.","type":"Common law and Ordinances"}`,
		}

	case LegalSystemsLawofPoland:
		return &LegalSystemsCode{
			Name:        "LawofPoland",
			Label:       "Law of Poland",
			Description: "",
			MetaData:    `{"description":"The Polish Civil Code in force since 1965.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPortugal:
		return &LegalSystemsCode{
			Name:        "LawofPortugal",
			Label:       "Law of Portugal",
			Description: "",
			MetaData:    `{"description":"Influenced by the Napoleonic Code and later by the German civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofPuertoRico:
		return &LegalSystemsCode{
			Name:        "LawofPuertoRico",
			Label:       "Law of Puerto Rico",
			Description: "",
			MetaData:    `{"description":"Based on Spanish law; influenced by U.S. common law after 1898; federal laws (based on common law) are in effect because of federal Supremacy Clause.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofQatar:
		return &LegalSystemsCode{
			Name:        "LawofQatar",
			Label:       "Law of Qatar",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law and Egyptian civil law system (after the French civil law system).","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofRéunionFrance:
		return &LegalSystemsCode{
			Name:        "LawofRéunionFrance",
			Label:       "Law of Réunion (France)",
			Description: "",
			MetaData:    `{"description":"Governed by French law, and its constitution is the French constitution of 28 September 1958.","type":"Civil law"}`,
		}

	case LegalSystemsLawofRomania:
		return &LegalSystemsCode{
			Name:        "LawofRomania",
			Label:       "Law of  Romania",
			Description: "",
			MetaData:    `{"description":"Civil Code came into force in 2011. Based on the Civil Code of Quebec, but also influenced by the Napoleonic Code and other French-inspired codes (such as those of Italy, Spain and Switzerland).","type":"Civil law"}`,
		}

	case LegalSystemsLawofRussia:
		return &LegalSystemsCode{
			Name:        "LawofRussia",
			Label:       "Law of Russia",
			Description: "",
			MetaData:    `{"description":"Civil Law system descendant from Roman Law through Byzantine tradition. Heavily influenced by German and Dutch norms in 1700–1800s. Socialist-style modification in 1900s, and Continental European Law influences since 1990s.","type":"Civil law"}`,
		}

	case LegalSystemsLawofRwanda:
		return &LegalSystemsCode{
			Name:        "LawofRwanda",
			Label:       "Law of Rwanda",
			Description: "",
			MetaData:    `{"description":"Mixture of Belgian civil law and English common law.","type":"Civil and common law"}`,
		}

	case LegalSystemsLawofSaintHelenaAscensionandTristandaCunhaEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofSaintHelenaAscensionandTristandaCunhaEnglandandWales",
			Label:       "Law of Saint Helena Ascension and Tristan da Cunha (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law and local statutes.","type":"Common law and local statues"}`,
		}

	case LegalSystemsLawofSaintKittsandNevis:
		return &LegalSystemsCode{
			Name:        "LawofSaintKittsandNevis",
			Label:       "Law of Saint Kitts and Nevis",
			Description: "",
			MetaData:    `{"description":"Based on English common law and statutory acts of the House of Assembly.","type":"Common law"}`,
		}

	case LegalSystemsLawofSaintLucia:
		return &LegalSystemsCode{
			Name:        "LawofSaintLucia",
			Label:       "Law of Saint Lucia",
			Description: "",
			MetaData:    `{"description":"Because it was both a French and English colony before it gained independence on February 11, 1979, the legal system is a mix of civil and English common law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSaintPierreandMiquelonFrance:
		return &LegalSystemsCode{
			Name:        "LawofSaintPierreandMiquelonFrance",
			Label:       "Law of Saint Pierre and Miquelon (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSaintVincentandtheGrenadines:
		return &LegalSystemsCode{
			Name:        "LawofSaintVincentandtheGrenadines",
			Label:       "Law of Saint Vincent and the Grenadines",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofSaintBarthelemyFrance:
		return &LegalSystemsCode{
			Name:        "LawofSaintBarthelemyFrance",
			Label:       "Law of Saint-Barthelemy (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSaintMartinFrance:
		return &LegalSystemsCode{
			Name:        "LawofSaintMartinFrance",
			Label:       "Law of Saint-Martin (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSamoa:
		return &LegalSystemsCode{
			Name:        "LawofSamoa",
			Label:       "Law of Samoa",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law; judicial review of legislative acts with respect to fundamental rights of the citizen.","type":"Common law and customary law"}`,
		}

	case LegalSystemsLawofSanMarino:
		return &LegalSystemsCode{
			Name:        "LawofSanMarino",
			Label:       "Law of San Marino",
			Description: "",
			MetaData:    `{"description":"Civil law system with Italian civil law influences.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSãoToméandPríncipe:
		return &LegalSystemsCode{
			Name:        "LawofSãoToméandPríncipe",
			Label:       "Law of São Tomé and Príncipe",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSaudiArabia:
		return &LegalSystemsCode{
			Name:        "LawofSaudiArabia",
			Label:       "Law of Saudi Arabia",
			Description: "",
			MetaData:    `{"description":"Islamic law.","type":"Religious law"}`,
		}

	case LegalSystemsLawofSenegal:
		return &LegalSystemsCode{
			Name:        "LawofSenegal",
			Label:       "Law of Senegal",
			Description: "",
			MetaData:    `{"description":"Civil law system based on French law; judicial review of legislative acts in Constitutional Court.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSerbia:
		return &LegalSystemsCode{
			Name:        "LawofSerbia",
			Label:       "Law of Serbia",
			Description: "",
			MetaData:    `{"description":"The Swiss civil law (Zivilgesetzbuch) was a model for the Law on Obligations of 1978.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSeychelles:
		return &LegalSystemsCode{
			Name:        "LawofSeychelles",
			Label:       "Law of Seychelles",
			Description: "",
			MetaData:    `{"description":"The substantive civil law is based on the French Civil Code. Otherwise the criminal law and court procedure are based on the English common law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSierraLeoneEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofSierraLeoneEnglandandWales",
			Label:       "Law of Sierra Leone (England and Wales)",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law.","type":"Common law and Customary law"}`,
		}

	case LegalSystemsLawofSingapore:
		return &LegalSystemsCode{
			Name:        "LawofSingapore",
			Label:       "Law of Singapore",
			Description: "",
			MetaData:    `{"description":"Based on English common law, but Muslims are subject to the Administration of Muslim Law Act, which gives the Sharia Court jurisdiction over Muslim personal law.","type":"Common law"}`,
		}

	case LegalSystemsLawofSintMaarten:
		return &LegalSystemsCode{
			Name:        "LawofSintMaarten",
			Label:       "Law of Sint Maarten",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law system with some English common law influence.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSlovakia:
		return &LegalSystemsCode{
			Name:        "LawofSlovakia",
			Label:       "Law of Slovakia",
			Description: "",
			MetaData:    `{"description":"Descended from the Civil Code of the Austrian Empire (1811), influenced by German (1939–45) and Soviet (1947/68–89) legal codes during occupation periods, substantially reformed to remove Soviet influence and elements of socialist law after the Velvet Revolution (1989).","type":"Civil law"}`,
		}

	case LegalSystemsLawofSlovenia:
		return &LegalSystemsCode{
			Name:        "LawofSlovenia",
			Label:       "Law of Slovenia",
			Description: "",
			MetaData:    `{"description":"A Civil Law system influenced mostly by Germanic and Austro-Hungarian law systems.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSolomonIslands:
		return &LegalSystemsCode{
			Name:        "LawofSolomonIslands",
			Label:       "Law of Solomon Islands",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law.","type":"Common law and Customary law"}`,
		}

	case LegalSystemsLawofSomalia:
		return &LegalSystemsCode{
			Name:        "LawofSomalia",
			Label:       "Law of Somalia",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil law, Islamic (sharia) law, and customary law (referred to as Xeer).","type":"Civil law, Sharia law and Customary law"}`,
		}

	case LegalSystemsLawofSouthAfrica:
		return &LegalSystemsCode{
			Name:        "LawofSouthAfrica",
			Label:       "Law of South Africa",
			Description: "",
			MetaData:    `{"description":"An amalgam of Roman-Dutch civil law and English common law, as well as Customary Law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales:
		return &LegalSystemsCode{
			Name:        "LawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales",
			Label:       "Law of South Georgia and the South Sandwich Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"The laws of the UK, where applicable, apply; the senior magistrate from the Falkland Islands presides over the Magistrates Court.","type":"Common law"}`,
		}

	case LegalSystemsLawofSouthKorea:
		return &LegalSystemsCode{
			Name:        "LawofSouthKorea",
			Label:       "Law of South Korea",
			Description: "",
			MetaData:    `{"description":"Based on German civil law system. Also largely influenced by Japanese civil law which itself modelled after German one. Korean Civil Code was introduced 1958 and fully enacted by 1960.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSouthSudan:
		return &LegalSystemsCode{
			Name:        "LawofSouthSudan",
			Label:       "Law of South Sudan",
			Description: "",
			MetaData:    `{"description":"Legal system is built on the combination of statutory and customary laws. South Sudan has enacted dozens of laws since 2005, but their use in legal disputes and courts is limited.","type":"Statutory and Customary law"}`,
		}

	case LegalSystemsLawofSpain:
		return &LegalSystemsCode{
			Name:        "LawofSpain",
			Label:       "Law of Spain",
			Description: "",
			MetaData:    `{"description":"An amalgam of Roman-Dutch civil law and English common law, as well as Customary Law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSriLanka:
		return &LegalSystemsCode{
			Name:        "LawofSriLanka",
			Label:       "Law of Sri Lanka",
			Description: "",
			MetaData:    `{"description":"An amalgam of English common law, Roman-Dutch civil law and Customary Law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSudan:
		return &LegalSystemsCode{
			Name:        "LawofSudan",
			Label:       "Law of Sudan",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law.","type":"Religious law"}`,
		}

	case LegalSystemsLawofSuriname:
		return &LegalSystemsCode{
			Name:        "LawofSuriname",
			Label:       "Law of Suriname",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSvalbardandJanMayenNorway:
		return &LegalSystemsCode{
			Name:        "LawofSvalbardandJanMayenNorway",
			Label:       "Law of Svalbard and Jan Mayen (Norway)",
			Description: "",
			MetaData:    `{"description":"The laws of Norway where applicable apply; only the laws of Norway made explicitly applicable to Svalbard have effect there; the Svalbard Act and the Svalbard Environmental Protection Act, and certain regulations, apply only to Svalbard; the Spitsbergen Treaty and the Svalbard Treaty grant certain rights to citizens and corporations of signatory nations; as of June 2017, 45 nations had ratified the Svalbard Treaty.","type":"Civil law"}`,
		}

	case LegalSystemsLawofEswatini:
		return &LegalSystemsCode{
			Name:        "LawofEswatini",
			Label:       "Law of Eswatini",
			Description: "",
			MetaData:    `{"description":"Based on South African law. A 1907 proclamation by the High Commissioner for Southern Africa applied the Roman-Dutch common law of the Transvaal Colony (now part of South Africa) to the Swaziland Protectorate (now Eswatini).","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofSweden:
		return &LegalSystemsCode{
			Name:        "LawofSweden",
			Label:       "Law of Sweden",
			Description: "",
			MetaData:    `{"description":"Scandinavian-North Germanic civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofSwitzerland:
		return &LegalSystemsCode{
			Name:        "LawofSwitzerland",
			Label:       "Law of Switzerland",
			Description: "",
			MetaData:    `{"description":"The Swiss Civil Code of 1908 and 1912 (obligations; fifth book).","type":"Civil law"}`,
		}

	case LegalSystemsLawofSyria:
		return &LegalSystemsCode{
			Name:        "LawofSyria",
			Label:       "Law of Syria",
			Description: "",
			MetaData:    `{"description":"Mainly based on Napoleonic/French Civil Code. Islamic law applicable to family law. Non Muslims follow their own family laws.","type":"Civil law and sharia law"}`,
		}

	case LegalSystemsLawofTaiwan:
		return &LegalSystemsCode{
			Name:        "LawofTaiwan",
			Label:       "Law of Taiwan",
			Description: "",
			MetaData:    `{"description":"Influenced by German Civil Code. Enacted in 1931.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTajikistan:
		return &LegalSystemsCode{
			Name:        "LawofTajikistan",
			Label:       "Law of Tajikistan",
			Description: "",
			MetaData:    `{"description":"Civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTanzania:
		return &LegalSystemsCode{
			Name:        "LawofTanzania",
			Label:       "Law of Tanzania",
			Description: "",
			MetaData:    `{"description":"Tanzania’s legal system is based on the English Common Law system. It derived this system from its British colonial legacy, as it does the system of government, which is based to a large degree on the Westminster parliamentary model.","type":"Common law"}`,
		}

	case LegalSystemsLawofThailand:
		return &LegalSystemsCode{
			Name:        "LawofThailand",
			Label:       "Law of Thailand",
			Description: "",
			MetaData:    `{"description":"The Thai legal system became an amalgam of German, Swiss, French, English, Japanese, Italian, Indian and American laws and practices. Even today, Islamic laws and practices exist in four southern provinces.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofTogo:
		return &LegalSystemsCode{
			Name:        "LawofTogo",
			Label:       "Law of Togo",
			Description: "",
			MetaData:    `{"description":"Togo has a civil law system based on the French model mixed with customary law.","type":"Civil law and customary law"}`,
		}

	case LegalSystemsLawofTokelau:
		return &LegalSystemsCode{
			Name:        "LawofTokelau",
			Label:       "Law of Tokelau",
			Description: "",
			MetaData:    `{"description":"Certain British and NZ Acts of Parliament; Rules of the Tokelau General fono; British common law as at 14 January 1840.","type":"Common law"}`,
		}

	case LegalSystemsLawofTrinidadandTobago:
		return &LegalSystemsCode{
			Name:        "LawofTrinidadandTobago",
			Label:       "Law of Trinidad and Tobago",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofTunisia:
		return &LegalSystemsCode{
			Name:        "LawofTunisia",
			Label:       "Law of Tunisia",
			Description: "",
			MetaData:    `{"description":"The Tunisian legal system is based on French civil law system.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTurkey:
		return &LegalSystemsCode{
			Name:        "LawofTurkey",
			Label:       "Law of Turkey",
			Description: "",
			MetaData:    `{"description":"Modeled after the Swiss civil law (Zivilgesetzbuch) of 1907.","type":"Civil law"}`,
		}

	case LegalSystemsLawofTurkmenistan:
		return &LegalSystemsCode{
			Name:        "LawofTurkmenistan",
			Label:       "Law of Turkmenistan",
			Description: "",
			MetaData:    `{"description":"The legal system in Turkmenistan is a civil law system, with Islamic law influences.","type":"Civil law and Islamic law"}`,
		}

	case LegalSystemsLawofTurksandCaicosIsland:
		return &LegalSystemsCode{
			Name:        "LawofTurksandCaicosIsland",
			Label:       "Law of Turks and Caicos Island",
			Description: "",
			MetaData:    `{"description":"A British Overseas Territory and a common law jurisdiction which is modeled on the English legal system.","type":"Common Law"}`,
		}

	case LegalSystemsLawofTuvalu:
		return &LegalSystemsCode{
			Name:        "LawofTuvalu",
			Label:       "Law of Tuvalu",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofUganda:
		return &LegalSystemsCode{
			Name:        "LawofUganda",
			Label:       "Law of Uganda",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofUkraine:
		return &LegalSystemsCode{
			Name:        "LawofUkraine",
			Label:       "Law of Ukraine",
			Description: "",
			MetaData:    `{"description":"Civil Code of Ukraine of 2004.","type":"Civil law"}`,
		}

	case LegalSystemsLawofUnitedArabEmirates:
		return &LegalSystemsCode{
			Name:        "LawofUnitedArabEmirates",
			Label:       "Law of United Arab Emirates",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law and Egyptian civil law system (after the French civil law system). Based on Common law system in the Dubai International Financial Center (DIFC Courts) and Abu Dhabi Global Market (ADGM) Courts (after the English Common law system).","type":"Common Law, civil law and sharia law"}`,
		}

	case LegalSystemsLawofNorthernIreland:
		return &LegalSystemsCode{
			Name:        "LawofNorthernIreland",
			Label:       "Law of Northern Ireland",
			Description: "",
			MetaData:    `{"description":"Based on Irish law before 1921, in turn based on English common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofUruguay:
		return &LegalSystemsCode{
			Name:        "LawofUruguay",
			Label:       "Law of Uruguay",
			Description: "",
			MetaData:    `{"description":"Civil law system, with public law based on the 1967 Constitution, amended in 1989, 1994, 1997, and 2004.","type":"Civil law"}`,
		}

	case LegalSystemsLawofUSVirginIslandsUnitedStates:
		return &LegalSystemsCode{
			Name:        "LawofUSVirginIslandsUnitedStates",
			Label:       "Law of US Virgin Islands (United States)",
			Description: "",
			MetaData:    `{"description":"The virgin Islands legal system transitioned from danish law to the American common law tradition.","type":"Common law"}`,
		}

	case LegalSystemsLawofUnitedStates:
		return &LegalSystemsCode{
			Name:        "LawofUnitedStates",
			Label:       "Law of United States",
			Description: "",
			MetaData:    `{"description":"Federal courts and 49 states use the legal system based on English common law. Law in the state of Louisiana is based on French and Spanish civil law.","type":"Common law"}`,
		}

	case LegalSystemsLawofAlaska:
		return &LegalSystemsCode{
			Name:        "LawofAlaska",
			Label:       "Law of Alaska",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofAlabama:
		return &LegalSystemsCode{
			Name:        "LawofAlabama",
			Label:       "Law of Alabama",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofArkansas:
		return &LegalSystemsCode{
			Name:        "LawofArkansas",
			Label:       "Law of Arkansas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofArizona:
		return &LegalSystemsCode{
			Name:        "LawofArizona",
			Label:       "Law of Arizona",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofCalifornia:
		return &LegalSystemsCode{
			Name:        "LawofCalifornia",
			Label:       "Law of California",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofColorado:
		return &LegalSystemsCode{
			Name:        "LawofColorado",
			Label:       "Law of Colorado",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofConnecticut:
		return &LegalSystemsCode{
			Name:        "LawofConnecticut",
			Label:       "Law of Connecticut",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWashingtonDC:
		return &LegalSystemsCode{
			Name:        "LawofWashingtonDC",
			Label:       "Law of Washington D.C.",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofDelaware:
		return &LegalSystemsCode{
			Name:        "LawofDelaware",
			Label:       "Law of Delaware",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofFlorida:
		return &LegalSystemsCode{
			Name:        "LawofFlorida",
			Label:       "Law of Florida",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofGeorgiaUSA:
		return &LegalSystemsCode{
			Name:        "LawofGeorgiaUSA",
			Label:       "Law of Georgia (USA)",
			Description: "",
			MetaData:    `{"description":"Georgia's legal system is based on common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofHawaii:
		return &LegalSystemsCode{
			Name:        "LawofHawaii",
			Label:       "Law of Hawaii",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofIowa:
		return &LegalSystemsCode{
			Name:        "LawofIowa",
			Label:       "Law of Iowa",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofIdaho:
		return &LegalSystemsCode{
			Name:        "LawofIdaho",
			Label:       "Law of Idaho",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofIllinois:
		return &LegalSystemsCode{
			Name:        "LawofIllinois",
			Label:       "Law of Illinois",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofIndiana:
		return &LegalSystemsCode{
			Name:        "LawofIndiana",
			Label:       "Law of Indiana",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofKansas:
		return &LegalSystemsCode{
			Name:        "LawofKansas",
			Label:       "Law of Kansas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofKentucky:
		return &LegalSystemsCode{
			Name:        "LawofKentucky",
			Label:       "Law of Kentucky",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofLouisiana:
		return &LegalSystemsCode{
			Name:        "LawofLouisiana",
			Label:       "Law of Louisiana",
			Description: "",
			MetaData:    `{"description":"Based on French and Spanish civil law, but federal laws (based on common law) are also in effect in Louisiana because of federal Supremacy Clause.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofMassachusetts:
		return &LegalSystemsCode{
			Name:        "LawofMassachusetts",
			Label:       "Law of Massachusetts",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMaryland:
		return &LegalSystemsCode{
			Name:        "LawofMaryland",
			Label:       "Law of Maryland",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMaine:
		return &LegalSystemsCode{
			Name:        "LawofMaine",
			Label:       "Law of Maine",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMichigan:
		return &LegalSystemsCode{
			Name:        "LawofMichigan",
			Label:       "Law of Michigan",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMinnesota:
		return &LegalSystemsCode{
			Name:        "LawofMinnesota",
			Label:       "Law of Minnesota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMissouri:
		return &LegalSystemsCode{
			Name:        "LawofMissouri",
			Label:       "Law of Missouri",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMississippi:
		return &LegalSystemsCode{
			Name:        "LawofMississippi",
			Label:       "Law of Mississippi",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofMontana:
		return &LegalSystemsCode{
			Name:        "LawofMontana",
			Label:       "Law of Montana",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNorthCarolina:
		return &LegalSystemsCode{
			Name:        "LawofNorthCarolina",
			Label:       "Law of North Carolina",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNorthDakota:
		return &LegalSystemsCode{
			Name:        "LawofNorthDakota",
			Label:       "Law of North Dakota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNebraska:
		return &LegalSystemsCode{
			Name:        "LawofNebraska",
			Label:       "Law of Nebraska",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewHampshire:
		return &LegalSystemsCode{
			Name:        "LawofNewHampshire",
			Label:       "Law of New Hampshire",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewJersey:
		return &LegalSystemsCode{
			Name:        "LawofNewJersey",
			Label:       "Law of New Jersey",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewMexico:
		return &LegalSystemsCode{
			Name:        "LawofNewMexico",
			Label:       "Law of New Mexico",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNevada:
		return &LegalSystemsCode{
			Name:        "LawofNevada",
			Label:       "Law of Nevada",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofNewYork:
		return &LegalSystemsCode{
			Name:        "LawofNewYork",
			Label:       "Law of New York",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofOhio:
		return &LegalSystemsCode{
			Name:        "LawofOhio",
			Label:       "Law of Ohio",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofOklahoma:
		return &LegalSystemsCode{
			Name:        "LawofOklahoma",
			Label:       "Law of Oklahoma",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofOregon:
		return &LegalSystemsCode{
			Name:        "LawofOregon",
			Label:       "Law of Oregon",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofPennsylvania:
		return &LegalSystemsCode{
			Name:        "LawofPennsylvania",
			Label:       "Law of Pennsylvania",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofRhodeIsland:
		return &LegalSystemsCode{
			Name:        "LawofRhodeIsland",
			Label:       "Law of Rhode Island",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofSouthCarolina:
		return &LegalSystemsCode{
			Name:        "LawofSouthCarolina",
			Label:       "Law of South Carolina",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofSouthDakota:
		return &LegalSystemsCode{
			Name:        "LawofSouthDakota",
			Label:       "Law of South Dakota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofTennessee:
		return &LegalSystemsCode{
			Name:        "LawofTennessee",
			Label:       "Law of Tennessee",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofTexas:
		return &LegalSystemsCode{
			Name:        "LawofTexas",
			Label:       "Law of Texas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofUtah:
		return &LegalSystemsCode{
			Name:        "LawofUtah",
			Label:       "Law of Utah",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofVirginia:
		return &LegalSystemsCode{
			Name:        "LawofVirginia",
			Label:       "Law of Virginia",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofVermont:
		return &LegalSystemsCode{
			Name:        "LawofVermont",
			Label:       "Law of Vermont",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWashington:
		return &LegalSystemsCode{
			Name:        "LawofWashington",
			Label:       "Law of Washington",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWisconsin:
		return &LegalSystemsCode{
			Name:        "LawofWisconsin",
			Label:       "Law of Wisconsin",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWestVirginia:
		return &LegalSystemsCode{
			Name:        "LawofWestVirginia",
			Label:       "Law of West Virginia",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofWyoming:
		return &LegalSystemsCode{
			Name:        "LawofWyoming",
			Label:       "Law of Wyoming",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		}

	case LegalSystemsLawofUzbekistan:
		return &LegalSystemsCode{
			Name:        "LawofUzbekistan",
			Label:       "Law of Uzbekistan",
			Description: "",
			MetaData:    `{"description":"Represents an evolution of Soviet civil law. Overwhelmingly strong impact of the Communist legal theory is traceable.","type":"Civil law"}`,
		}

	case LegalSystemsLawofVanuatu:
		return &LegalSystemsCode{
			Name:        "LawofVanuatu",
			Label:       "Law of Vanuatu",
			Description: "",
			MetaData:    `{"description":"Consists of a mixed system combining the legacy of English common law, French civil law and indigenous customary law.","type":"Civil law and common law"}`,
		}

	case LegalSystemsLawofVaticanCity:
		return &LegalSystemsCode{
			Name:        "LawofVaticanCity",
			Label:       "Law of Vatican City",
			Description: "",
			MetaData:    `{"description":"Based on Roman \u0026 Italian civil law and Catholic canon law.","type":"Civil law and canon law"}`,
		}

	case LegalSystemsLawofVenezuela:
		return &LegalSystemsCode{
			Name:        "LawofVenezuela",
			Label:       "Law of Venezuela",
			Description: "",
			MetaData:    `{"description":"Venezuela's legal system has a legislative origin, grounded on \"written law\" (civil law).","type":"Civil law"}`,
		}

	case LegalSystemsLawofVietnam:
		return &LegalSystemsCode{
			Name:        "LawofVietnam",
			Label:       "Law of Vietnam",
			Description: "",
			MetaData:    `{"description":"Communist legal theory and French civil law.","type":"Civil law"}`,
		}

	case LegalSystemsLawofWallisandFutunaFrance:
		return &LegalSystemsCode{
			Name:        "LawofWallisandFutunaFrance",
			Label:       "Law of Wallis and Futuna (France)",
			Description: "",
			MetaData:    `{"description":"Uses both the French legal system and customary local laws.","type":"Civil law and Customary law"}`,
		}

	case LegalSystemsLawofYemen:
		return &LegalSystemsCode{
			Name:        "LawofYemen",
			Label:       "Law of Yemen",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of Islamic (sharia) law, Napoleonic law, English common law, and customary law.","type":"Civil, Common, Customary and Religious law"}`,
		}

	case LegalSystemsLawofZambia:
		return &LegalSystemsCode{
			Name:        "LawofZambia",
			Label:       "Law of Zambia",
			Description: "",
			MetaData:    `{"description":"Zambian legal system is largely based on English Common Law. Customary laws relating to particular communities around the country are also recognised by the constitution.","type":"Common law and Customary law"}`,
		}

	case LegalSystemsLawofZimbabwe:
		return &LegalSystemsCode{
			Name:        "LawofZimbabwe",
			Label:       "Law of Zimbabwe",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1891 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to Southern Rhodesia (now Zimbabwe).","type":"Civil law and common law"}`,
		}
	default:
		return nil
	}
}

// LegalSystemsMap returns a mapping of LegalSystems objects with the code as the key.
func LegalSystemsMap() map[string]*LegalSystemsCode {
	return map[string]*LegalSystemsCode{

		LegalSystemsLawofEnglandandWales: &LegalSystemsCode{
			Name:        "LawofEnglandandWales",
			Label:       "Law of England and Wales",
			Description: "",
			MetaData:    `{"description":"Primarily common law, with early Roman and some modern continental European influences.","type":"Common law"}`,
		},

		LegalSystemsLawofScotland: &LegalSystemsCode{
			Name:        "LawofScotland",
			Label:       "Law of Scotland",
			Description: "",
			MetaData:    `{"description":"Based on Roman and continental law, with common law elements dating back to the High Middle Ages.","type":"Common law"}`,
		},

		LegalSystemsLawofAfghanistan: &LegalSystemsCode{
			Name:        "LawofAfghanistan",
			Label:       "Law of Afghanistan",
			Description: "",
			MetaData:    `{"description":"Islamic law \u0026 American/British law after invasion.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofAlbania: &LegalSystemsCode{
			Name:        "LawofAlbania",
			Label:       "Law of Albania",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law. The Civil Code of the Republic of Albania, 1991.","type":"Civil law"}`,
		},

		LegalSystemsLawofAlgeria: &LegalSystemsCode{
			Name:        "LawofAlgeria",
			Label:       "Law of Algeria",
			Description: "",
			MetaData:    `{"description":"The Algerian judicial system is based on a civil law system with codes adapted from the French legal system. Personal status laws are based on Islamic law.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofAmericanSamoa: &LegalSystemsCode{
			Name:        "LawofAmericanSamoa",
			Label:       "Law of American Samoa",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		},

		LegalSystemsLawofAndorra: &LegalSystemsCode{
			Name:        "LawofAndorra",
			Label:       "Law of Andorra",
			Description: "",
			MetaData:    `{"description":"Courts apply the customary laws of Andorra, supplemented with Roman law and customary Catalan law.","type":"Civil law"}`,
		},

		LegalSystemsLawofAngola: &LegalSystemsCode{
			Name:        "LawofAngola",
			Label:       "Law of Angola",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofAnguilla: &LegalSystemsCode{
			Name:        "LawofAnguilla",
			Label:       "Law of Anguilla",
			Description: "",
			MetaData:    `{"description":"A combination of common law and statute, and is based heavily upon English law.","type":"Common law and statute"}`,
		},

		LegalSystemsLawofAntiguaandBarbuda: &LegalSystemsCode{
			Name:        "LawofAntiguaandBarbuda",
			Label:       "Law of Antigua and Barbuda",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofArgentina: &LegalSystemsCode{
			Name:        "LawofArgentina",
			Label:       "Law of Argentina",
			Description: "",
			MetaData:    `{"description":"The Legal system of Argentina is a Civil law legal system. The pillar of the Civil system is the Constitution of Argentina (1853).","type":"Civil law"}`,
		},

		LegalSystemsLawofArmenia: &LegalSystemsCode{
			Name:        "LawofArmenia",
			Label:       "Law of Armenia",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law and traditional Armenian law.","type":"Civil law"}`,
		},

		LegalSystemsLawofAruba: &LegalSystemsCode{
			Name:        "LawofAruba",
			Label:       "Law of Aruba",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofAustralia: &LegalSystemsCode{
			Name:        "LawofAustralia",
			Label:       "Law of Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewSouthWales: &LegalSystemsCode{
			Name:        "LawofNewSouthWales",
			Label:       "Law of New South Wales",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofQueensland: &LegalSystemsCode{
			Name:        "LawofQueensland",
			Label:       "Law of Queensland",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofSouthAustralia: &LegalSystemsCode{
			Name:        "LawofSouthAustralia",
			Label:       "Law of South Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofTasmania: &LegalSystemsCode{
			Name:        "LawofTasmania",
			Label:       "Law of Tasmania",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofVictoria: &LegalSystemsCode{
			Name:        "LawofVictoria",
			Label:       "Law of Victoria",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWesternAustralia: &LegalSystemsCode{
			Name:        "LawofWesternAustralia",
			Label:       "Law of Western Australia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofAustralianCapitalTerritory: &LegalSystemsCode{
			Name:        "LawofAustralianCapitalTerritory",
			Label:       "Law of Australian Capital Territory",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNorthernTerritory: &LegalSystemsCode{
			Name:        "LawofNorthernTerritory",
			Label:       "Law of Northern Territory",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofJervisBayTerritoryAustralianCapitalTerritory: &LegalSystemsCode{
			Name:        "LawofJervisBayTerritoryAustralianCapitalTerritory",
			Label:       "Law of Jervis Bay Territory (Australian Capital Territory)",
			Description: "",
			MetaData:    `{"description":"Although the Jervis Bay Territory is not part of the Australian Capital Territory, the laws of the ACT apply.","type":"Common law"}`,
		},

		LegalSystemsLawofAustria: &LegalSystemsCode{
			Name:        "LawofAustria",
			Label:       "Law of Austria",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofAzerbaijan: &LegalSystemsCode{
			Name:        "LawofAzerbaijan",
			Label:       "Law of Azerbaijan",
			Description: "",
			MetaData:    `{"description":"Based on German, French, Russian and traditional Azerbaijani Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofBahamas: &LegalSystemsCode{
			Name:        "LawofBahamas",
			Label:       "Law of Bahamas",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofBahrain: &LegalSystemsCode{
			Name:        "LawofBahrain",
			Label:       "Law of Bahrain",
			Description: "",
			MetaData:    `{"description":"Bahrain's legal system is based on civil law, customs, and Sharia'a (Islamic jurisprudence). The system mostly derives from the legal system of Egypt, which is based on the legal system of France.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofBangladesh: &LegalSystemsCode{
			Name:        "LawofBangladesh",
			Label:       "Law of Bangladesh",
			Description: "",
			MetaData:    `{"description":"Based on English common law, with family law heavily based on Shar'iah law.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofBarbados: &LegalSystemsCode{
			Name:        "LawofBarbados",
			Label:       "Law of Barbados",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofBelarus: &LegalSystemsCode{
			Name:        "LawofBelarus",
			Label:       "Law of Belarus",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofBelgium: &LegalSystemsCode{
			Name:        "LawofBelgium",
			Label:       "Law of Belgium",
			Description: "",
			MetaData:    `{"description":"The Napoleonic Code is still in use, although it is heavily modified (especially concerning family law).","type":"Civil law"}`,
		},

		LegalSystemsLawofBelize: &LegalSystemsCode{
			Name:        "LawofBelize",
			Label:       "Law of Belize",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofBenin: &LegalSystemsCode{
			Name:        "LawofBenin",
			Label:       "Law of Benin",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic Civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofBermuda: &LegalSystemsCode{
			Name:        "LawofBermuda",
			Label:       "Law of Bermuda",
			Description: "",
			MetaData:    `{"description":"Based on the common law legal system of Bermuda.","type":"Common law"}`,
		},

		LegalSystemsLawofBhutan: &LegalSystemsCode{
			Name:        "LawofBhutan",
			Label:       "Law of Bhutan",
			Description: "",
			MetaData:    `{"description":"Based on English common law, with Indian influence. Religious law influences personal law.","type":"Common law"}`,
		},

		LegalSystemsLawofBolivia: &LegalSystemsCode{
			Name:        "LawofBolivia",
			Label:       "Law of Bolivia",
			Description: "",
			MetaData:    `{"description":"Influenced by the Napoleonic Code.","type":"Civil law"}`,
		},

		LegalSystemsLawofBonaireStEustasuisandSaba: &LegalSystemsCode{
			Name:        "LawofBonaireStEustasuisandSaba",
			Label:       "Law of Bonaire St Eustasuis and Saba",
			Description: "",
			MetaData:    `{"description":"The main body of civil law is the Civil Code. In the years to come the applicable former Netherlands Antilles legislation on St. Eustatius will gradually be replaced by Dutch legislation.","type":"Civil law"}`,
		},

		LegalSystemsLawofBosniaandHerzegovina: &LegalSystemsCode{
			Name:        "LawofBosniaandHerzegovina",
			Label:       "Law of Bosnia and Herzegovina",
			Description: "",
			MetaData:    `{"description":"Influenced by Austrian law. The Swiss civil law (Zivilgesetzbuch) was a model for the Law on Obligations of 1978.","type":"Civil law"}`,
		},

		LegalSystemsLawofBotswana: &LegalSystemsCode{
			Name:        "LawofBotswana",
			Label:       "Law of Botswana",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1891 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to the Bechuanaland Protectorate (now Botswana).","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofBrazil: &LegalSystemsCode{
			Name:        "LawofBrazil",
			Label:       "Law of Brazil",
			Description: "",
			MetaData:    `{"description":"Based on German, Italian, French and Portuguese law. However, in 2004 the Federal Constitution was amended to grant the Supreme Federal Court authority to issue binding precedents (súmulas vinculantes) to settle controversies involving constitutional law - a mechanism that echoes the stare decisis principle typically found in common law systems.","type":"Civil law"}`,
		},

		LegalSystemsLawofBritishVirginIslands: &LegalSystemsCode{
			Name:        "LawofBritishVirginIslands",
			Label:       "Law of British Virgin Islands",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofBrunei: &LegalSystemsCode{
			Name:        "LawofBrunei",
			Label:       "Law of Brunei",
			Description: "",
			MetaData:    `{"description":"Most laws under Common Law and the Sharia Penal Code apply to all people in Brunei, regardless of nationality or religion.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofBulgaria: &LegalSystemsCode{
			Name:        "LawofBulgaria",
			Label:       "Law of Bulgaria",
			Description: "",
			MetaData:    `{"description":"Civil Law system influenced by Germanic and Roman law systems.","type":"Civil law"}`,
		},

		LegalSystemsLawofBurkinaFaso: &LegalSystemsCode{
			Name:        "LawofBurkinaFaso",
			Label:       "Law of Burkina Faso",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofBurundi: &LegalSystemsCode{
			Name:        "LawofBurundi",
			Label:       "Law of Burundi",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of Belgian civil law and customary law.","type":"Civil law and customary law"}`,
		},

		LegalSystemsLawofCambodia: &LegalSystemsCode{
			Name:        "LawofCambodia",
			Label:       "Law of Cambodia",
			Description: "",
			MetaData:    `{"description":"The Cambodian legal system is based largely on the French civil system, and is statute based.","type":"Civil law"}`,
		},

		LegalSystemsLawofCameroon: &LegalSystemsCode{
			Name:        "LawofCameroon",
			Label:       "Law of Cameroon",
			Description: "",
			MetaData:    `{"description":"Cameroon is a bijural system with the English Common Law operating in the two Anglophone regions of North West and South West and the French Civil Law operating in the eight francophone regions of Adamaoua, Centre, East, Far North, Littoral, North, West and South.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofCanada: &LegalSystemsCode{
			Name:        "LawofCanada",
			Label:       "Law of Canada",
			Description: "",
			MetaData:    `{"description":"Based on English common law, except in Quebec Quebec, where a civil law system based on French law prevails in most matters of a civil nature, such as obligations (contract and delict), property law, family law and private matters. Federal statutes take into account the bijuridical nature of Canada and use both common law and civil law terms where appropriate.","type":"Common law"}`,
		},

		LegalSystemsLawofAlberta: &LegalSystemsCode{
			Name:        "LawofAlberta",
			Label:       "Law of Alberta",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofBritishColumbia: &LegalSystemsCode{
			Name:        "LawofBritishColumbia",
			Label:       "Law of British Columbia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofManitoba: &LegalSystemsCode{
			Name:        "LawofManitoba",
			Label:       "Law of Manitoba",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewBrunswick: &LegalSystemsCode{
			Name:        "LawofNewBrunswick",
			Label:       "Law of New Brunswick",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewfoundlandandLabrador: &LegalSystemsCode{
			Name:        "LawofNewfoundlandandLabrador",
			Label:       "Law of Newfoundland and Labrador",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNovaScotia: &LegalSystemsCode{
			Name:        "LawofNovaScotia",
			Label:       "Law of Nova Scotia",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofOntario: &LegalSystemsCode{
			Name:        "LawofOntario",
			Label:       "Law of Ontario",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofPrinceEdwardIsland: &LegalSystemsCode{
			Name:        "LawofPrinceEdwardIsland",
			Label:       "Law of Prince Edward Island",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofQuebec: &LegalSystemsCode{
			Name:        "LawofQuebec",
			Label:       "Law of Quebec",
			Description: "",
			MetaData:    `{"description":"The Civil Code of Quebec is the civil code currently in force. Canadian (federal) criminal law in force in Quebec is based on common law, but federal statutes of or relating to private law take into account the bijuridical nature of Canada and use both common law and civil law terms where appropriate.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSaskatchewan: &LegalSystemsCode{
			Name:        "LawofSaskatchewan",
			Label:       "Law of Saskatchewan",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNorthwestTerritories: &LegalSystemsCode{
			Name:        "LawofNorthwestTerritories",
			Label:       "Law of Northwest Territories",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNunavut: &LegalSystemsCode{
			Name:        "LawofNunavut",
			Label:       "Law of Nunavut",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofYukon: &LegalSystemsCode{
			Name:        "LawofYukon",
			Label:       "Law of Yukon",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofCapeVerde: &LegalSystemsCode{
			Name:        "LawofCapeVerde",
			Label:       "Law of Cape Verde",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofCaymanIslands: &LegalSystemsCode{
			Name:        "LawofCaymanIslands",
			Label:       "Law of Cayman Islands",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofCentralAfricanRepublic: &LegalSystemsCode{
			Name:        "LawofCentralAfricanRepublic",
			Label:       "Law of Central African Republic",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofChad: &LegalSystemsCode{
			Name:        "LawofChad",
			Label:       "Law of Chad",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofChile: &LegalSystemsCode{
			Name:        "LawofChile",
			Label:       "Law of Chile",
			Description: "",
			MetaData:    `{"description":"Chile's legal system is civil law based. It is primarily based on the Chilean Civil Code of 1855, derived from Spanish law and other codes of Continental Europe of the 19th century.","type":"Civil law"}`,
		},

		LegalSystemsLawofPeoplesRepublicofChina: &LegalSystemsCode{
			Name:        "LawofPeoplesRepublicofChina",
			Label:       "Law of People's Republic of China",
			Description: "",
			MetaData:    `{"description":"Based on Germanic Civil law with influences from the Soviet Socialist law from Soviet Union.","type":"Civil law"}`,
		},

		LegalSystemsLawofChristmasIslandWesternAustralia: &LegalSystemsCode{
			Name:        "LawofChristmasIslandWesternAustralia",
			Label:       "Law of Christmas Island (Western Australia)",
			Description: "",
			MetaData:    `{"description":"The Christmas Island Act - applies Western Australian laws on Christmas Island, including the Local Government Act 1995 (WA)","type":"Common law"}`,
		},

		LegalSystemsLawofCocosIslandsWesternAustralia: &LegalSystemsCode{
			Name:        "LawofCocosIslandsWesternAustralia",
			Label:       "Law of Cocos Islands (Western Australia)",
			Description: "",
			MetaData:    `{"description":"The Territories Law Reform Act 1992 (Cth) amended the Cocos (Keeling) Islands Act to introduce a modern body of Australian law to the Cocos (Keeling) Islands, including provisions to enliven the application of most Australian Government laws and regulations and to apply Western Australian laws to the territory.","type":"Common law"}`,
		},

		LegalSystemsLawofColombia: &LegalSystemsCode{
			Name:        "LawofColombia",
			Label:       "Law of Colombia",
			Description: "",
			MetaData:    `{"description":"Based on the Chilean Civil Law. Civil code introduced in 1873. Nearly faithful reproduction of the Chilean civil code.","type":"Civil law"}`,
		},

		LegalSystemsLawofComoros: &LegalSystemsCode{
			Name:        "LawofComoros",
			Label:       "Law of Comoros",
			Description: "",
			MetaData:    `{"description":"The Comorian legal system rests on Islamic law and an inherited French legal code. Village elders or civilian courts settle most disputes.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofRepublicoftheCongo: &LegalSystemsCode{
			Name:        "LawofRepublicoftheCongo",
			Label:       "Law of Republic of the Congo",
			Description: "",
			MetaData:    `{"description":"Based on the Napoleonic Civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofDemocraticRepublicoftheCongo: &LegalSystemsCode{
			Name:        "LawofDemocraticRepublicoftheCongo",
			Label:       "Law of Democratic Republic of the Congo",
			Description: "",
			MetaData:    `{"description":"Based on Belgian civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofCookIslands: &LegalSystemsCode{
			Name:        "LawofCookIslands",
			Label:       "Law of Cook Islands",
			Description: "",
			MetaData:    `{"description":"Constitution of the Cook Islands - the supreme law. English common law and equity - except so far as inconsistent with the Cook Islands Act 1915 (NZ) or inappropriate to the circumstances of the country or inconsistent with the Constitution.","type":"Common law"}`,
		},

		LegalSystemsLawofCostaRica: &LegalSystemsCode{
			Name:        "LawofCostaRica",
			Label:       "Law of Costa Rica",
			Description: "",
			MetaData:    `{"description":"The present Civil Code went into effect 1 January 1888, and was influenced by the Napoleonic Code and the Spanish Civil Code of 1889.","type":"Civil law"}`,
		},

		LegalSystemsLawofCroatia: &LegalSystemsCode{
			Name:        "LawofCroatia",
			Label:       "Law of Croatia",
			Description: "",
			MetaData:    `{"description":"Based on the Germanic Civil Law. Croatian Law system is largely influenced by German and Austrian law systems.","type":"Civil law"}`,
		},

		LegalSystemsLawofCuba: &LegalSystemsCode{
			Name:        "LawofCuba",
			Label:       "Law of Cuba",
			Description: "",
			MetaData:    `{"description":"Influenced by Spanish and American law with large elements of Communist legal theory.","type":"Civil law"}`,
		},

		LegalSystemsLawofCuracao: &LegalSystemsCode{
			Name:        "LawofCuracao",
			Label:       "Law of Curacao",
			Description: "",
			MetaData:    `{"description":"Based on Dutch Civil Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofCyprus: &LegalSystemsCode{
			Name:        "LawofCyprus",
			Label:       "Law of Cyprus",
			Description: "",
			MetaData:    `{"description":"Based on English common law as inherited from British colonisation, with civil law influences, particularly in criminal law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofCzechRepublic: &LegalSystemsCode{
			Name:        "LawofCzechRepublic",
			Label:       "Law of Czech Republic",
			Description: "",
			MetaData:    `{"description":"Based on Belgian civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofDenmark: &LegalSystemsCode{
			Name:        "LawofDenmark",
			Label:       "Law of Denmark",
			Description: "",
			MetaData:    `{"description":"Based on North Germanic law. Scandinavian-North Germanic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofDjibouti: &LegalSystemsCode{
			Name:        "LawofDjibouti",
			Label:       "Law of Djibouti",
			Description: "",
			MetaData:    `{"description":"Djibouti's legal and judicial system is largely inspired by the French legislation. Laws are codified in which the system of the country is based on the coexistence of Islamic law (Sharia), customary law and civil law inherited the French Napoleon Code.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofDominica: &LegalSystemsCode{
			Name:        "LawofDominica",
			Label:       "Law of Dominica",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofDominicanRepublic: &LegalSystemsCode{
			Name:        "LawofDominicanRepublic",
			Label:       "Law of Dominican Republic",
			Description: "",
			MetaData:    `{"description":"Based on the Napoleonic Code.","type":"Civil law"}`,
		},

		LegalSystemsLawofTimorLeste: &LegalSystemsCode{
			Name:        "LawofTimorLeste",
			Label:       "Law of Timor-Leste",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofEcuador: &LegalSystemsCode{
			Name:        "LawofEcuador",
			Label:       "Law of Ecuador",
			Description: "",
			MetaData:    `{"description":"Based on the Chilean civil law. Civil code introduced in 1861.","type":"Civil law"}`,
		},

		LegalSystemsLawofEgypt: &LegalSystemsCode{
			Name:        "LawofEgypt",
			Label:       "Law of Egypt",
			Description: "",
			MetaData:    `{"description":"Family Law (personal Statute) for Muslims based on Islamic Jurisprudence, Separate Personal Statute for non Muslims, and all other branches of Law are based on French civil law system.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofElSalvador: &LegalSystemsCode{
			Name:        "LawofElSalvador",
			Label:       "Law of El Salvador",
			Description: "",
			MetaData:    `{"description":"Civil law system with minor common law influence; judicial review of legislative acts in the Supreme Court.","type":"Civil law"}`,
		},

		LegalSystemsLawofEquatorialGuinea: &LegalSystemsCode{
			Name:        "LawofEquatorialGuinea",
			Label:       "Law of Equatorial Guinea",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofEritrea: &LegalSystemsCode{
			Name:        "LawofEritrea",
			Label:       "Law of Eritrea",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil, customary, and Islamic religious law.","type":"Civil law and sharia law and Customary law"}`,
		},

		LegalSystemsLawofEstonia: &LegalSystemsCode{
			Name:        "LawofEstonia",
			Label:       "Law of Estonia",
			Description: "",
			MetaData:    `{"description":"Based on German civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofEthiopia: &LegalSystemsCode{
			Name:        "LawofEthiopia",
			Label:       "Law of Ethiopia",
			Description: "",
			MetaData:    `{"description":"Ethiopia's legal system largely belongs in the civil law system. Conversely, the case law system is a distinguishing feature of the common law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofFalklandIslandsEnglandandWales: &LegalSystemsCode{
			Name:        "LawofFalklandIslandsEnglandandWales",
			Label:       "Law of Falkland Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law and local statutes.","type":"Common law"}`,
		},

		LegalSystemsLawofFaroeIslandsDenmark: &LegalSystemsCode{
			Name:        "LawofFaroeIslandsDenmark",
			Label:       "Law of Faroe Islands (Denmark)",
			Description: "",
			MetaData:    `{"description":"The laws of Denmark apply where applicable. The Faroe Islands have the exclusive right to legislate and govern independently in a wide range of areas.  The Faroese court system is under the jurisdiction of the high courts in Denmark.","type":"Civil law"}`,
		},

		LegalSystemsLawofFiji: &LegalSystemsCode{
			Name:        "LawofFiji",
			Label:       "Law of Fiji",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofFinland: &LegalSystemsCode{
			Name:        "LawofFinland",
			Label:       "Law of Finland",
			Description: "",
			MetaData:    `{"description":"Based on Nordic law.","type":"Civil law"}`,
		},

		LegalSystemsLawofFrance: &LegalSystemsCode{
			Name:        "LawofFrance",
			Label:       "Law of France",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic code (code civil of 1804).","type":"Civil law"}`,
		},

		LegalSystemsLawofFrenchGuianaFrance: &LegalSystemsCode{
			Name:        "LawofFrenchGuianaFrance",
			Label:       "Law of French Guiana (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		},

		LegalSystemsLawofFrenchPolynesiaFrance: &LegalSystemsCode{
			Name:        "LawofFrenchPolynesiaFrance",
			Label:       "Law of French Polynesia (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		},

		LegalSystemsLawofFrenchSouthernandAntarcticLandsFrance: &LegalSystemsCode{
			Name:        "LawofFrenchSouthernandAntarcticLandsFrance",
			Label:       "Law of French Southern and Antarctic Lands (France)",
			Description: "",
			MetaData:    `{"description":"Is an overseas department and region of France.","type":"Civil law"}`,
		},

		LegalSystemsLawofGabon: &LegalSystemsCode{
			Name:        "LawofGabon",
			Label:       "Law of Gabon",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofTheGambia: &LegalSystemsCode{
			Name:        "LawofTheGambia",
			Label:       "Law of The Gambia",
			Description: "",
			MetaData:    `{"description":"English common law, Islamic law and customary law.","type":"Common law and sharia law and Customary law"}`,
		},

		LegalSystemsLawofGeorgia: &LegalSystemsCode{
			Name:        "LawofGeorgia",
			Label:       "Law of Georgia",
			Description: "",
			MetaData:    `{"description":"Civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofGermany: &LegalSystemsCode{
			Name:        "LawofGermany",
			Label:       "Law of Germany",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law. The Bürgerliches Gesetzbuch of 1900 (\"BGB\"). The BGB is influenced both by Roman and German law traditions.","type":"Civil law"}`,
		},

		LegalSystemsLawofGhana: &LegalSystemsCode{
			Name:        "LawofGhana",
			Label:       "Law of Ghana",
			Description: "",
			MetaData:    `{"description":"Mixed system of English common law and customary law.","type":"Common law and customary law"}`,
		},

		LegalSystemsLawofGibraltar: &LegalSystemsCode{
			Name:        "LawofGibraltar",
			Label:       "Law of Gibraltar",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofGreece: &LegalSystemsCode{
			Name:        "LawofGreece",
			Label:       "Law of Greece",
			Description: "",
			MetaData:    `{"description":"Greece is a civil law country.","type":"Civil law"}`,
		},

		LegalSystemsLawofGreenlandDenmark: &LegalSystemsCode{
			Name:        "LawofGreenlandDenmark",
			Label:       "Law of Greenland (Denmark)",
			Description: "",
			MetaData:    `{"description":"The laws of Denmark apply where applicable and Greenlandic law applies to other areas.","type":"Civil law"}`,
		},

		LegalSystemsLawofGrenada: &LegalSystemsCode{
			Name:        "LawofGrenada",
			Label:       "Law of Grenada",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofGuadeloupeFrance: &LegalSystemsCode{
			Name:        "LawofGuadeloupeFrance",
			Label:       "Law of Guadeloupe (France)",
			Description: "",
			MetaData:    `{"description":"Guadeloupe honors the French Constitution of September 1958, and utilizes the French legal system.","type":"Civil law"}`,
		},

		LegalSystemsLawofGuamUnitedStates: &LegalSystemsCode{
			Name:        "LawofGuamUnitedStates",
			Label:       "Law of Guam (United States)",
			Description: "",
			MetaData:    `{"description":"The legal and judicial system in Guam is the same as in any other U.S. state or territory. All proceedings are in English, and based on the American common law legal system and US standards of integrity.","type":"Common law"}`,
		},

		LegalSystemsLawofGuatemala: &LegalSystemsCode{
			Name:        "LawofGuatemala",
			Label:       "Law of Guatemala",
			Description: "",
			MetaData:    `{"description":"Guatemala is a civil-law country.","type":"Civil law"}`,
		},

		LegalSystemsLawofGuernsey: &LegalSystemsCode{
			Name:        "LawofGuernsey",
			Label:       "Law of Guernsey",
			Description: "",
			MetaData:    `{"description":"Customary legal system based on Norman customary law; includes elements of the French civil code and English common law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofGuinea: &LegalSystemsCode{
			Name:        "LawofGuinea",
			Label:       "Law of Guinea",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system, customary law, and decree.","type":"Civil law"}`,
		},

		LegalSystemsLawofGuineaBissau: &LegalSystemsCode{
			Name:        "LawofGuineaBissau",
			Label:       "Law of Guinea-Bissau",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofGuyana: &LegalSystemsCode{
			Name:        "LawofGuyana",
			Label:       "Law of Guyana",
			Description: "",
			MetaData:    `{"description":"Vestiges of a Dutch legal system remain, particularly in the area of land tenure. However, the common law of Britain is the basis for the legal system of Guyana.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofHaiti: &LegalSystemsCode{
			Name:        "LawofHaiti",
			Label:       "Law of Haiti",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory: &LegalSystemsCode{
			Name:        "LawofHeardIslandandMcDonaldsIslandsAustralianCapitalTerritory",
			Label:       "Law of Heard Island and McDonalds Islands (Australian Capital Territory)",
			Description: "",
			MetaData:    `{"description":"The administration of the territory is established in the Heard Island and McDonald Islands Act 1953, which places it under the laws of the Australian Capital Territory and the jurisdiction of the Supreme Court of the Australian Capital Territory.","type":"Common law"}`,
		},

		LegalSystemsLawofHonduras: &LegalSystemsCode{
			Name:        "LawofHonduras",
			Label:       "Law of Honduras",
			Description: "",
			MetaData:    `{"description":"Because Honduras is a Civil law country, there is not much role for precedent or judicial review.","type":"Civil law"}`,
		},

		LegalSystemsLawofHongKong: &LegalSystemsCode{
			Name:        "LawofHongKong",
			Label:       "Law of Hong Kong",
			Description: "",
			MetaData:    `{"description":"Principally based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofHungary: &LegalSystemsCode{
			Name:        "LawofHungary",
			Label:       "Law of Hungary",
			Description: "",
			MetaData:    `{"description":"Based on Germanic, codified Roman law with elements from Napoleonic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofIceland: &LegalSystemsCode{
			Name:        "LawofIceland",
			Label:       "Law of Iceland",
			Description: "",
			MetaData:    `{"description":"Based on North Germanic law. Germanic traditional laws and influenced by Medieval Norwegian and Danish laws.","type":"Civil law"}`,
		},

		LegalSystemsLawofIndia: &LegalSystemsCode{
			Name:        "LawofIndia",
			Label:       "Law of India",
			Description: "",
			MetaData:    `{"description":"Based on English common law, except in Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law. Muslim personal law based on sharia law applies to Muslims. Exceptions for Muslims in Goa state, where the Goa Civil Code applies to all persons irrespective of religion, and for Muslims who marry under the Special Marriage Act, 1954.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofDadraandNagarHaveli: &LegalSystemsCode{
			Name:        "LawofDadraandNagarHaveli",
			Label:       "Law of Dadra and Nagar Haveli",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofDamanandDiu: &LegalSystemsCode{
			Name:        "LawofDamanandDiu",
			Label:       "Law of Daman and Diu",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofGoa: &LegalSystemsCode{
			Name:        "LawofGoa",
			Label:       "Law of Goa",
			Description: "",
			MetaData:    `{"description":"Goa, Daman and Diu and Dadra and Nagar Haveli which follow a Civil law system based on the Portuguese Civil Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofIndonesia: &LegalSystemsCode{
			Name:        "LawofIndonesia",
			Label:       "Law of Indonesia",
			Description: "",
			MetaData:    `{"description":"Based on a civil law system, intermixed with customary law and the Roman Dutch law.","type":"Civil law"}`,
		},

		LegalSystemsLawofIran: &LegalSystemsCode{
			Name:        "LawofIran",
			Label:       "Law of Iran",
			Description: "",
			MetaData:    `{"description":"Religious legal system based on secular and Islamic law.","type":"Religious law"}`,
		},

		LegalSystemsLawofIraq: &LegalSystemsCode{
			Name:        "LawofIraq",
			Label:       "Law of Iraq",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil and Islamic law.","type":"Civil law and Religious law"}`,
		},

		LegalSystemsLawofIreland: &LegalSystemsCode{
			Name:        "LawofIreland",
			Label:       "Law of Ireland",
			Description: "",
			MetaData:    `{"description":"Based on Irish law before 1922, which was itself based on English common law.","type":"Common law"}`,
		},

		LegalSystemsIsleofMan: &LegalSystemsCode{
			Name:        "IsleofMan",
			Label:       "Isle of Man",
			Description: "",
			MetaData:    `{"description":"The legal system on the Isle of Man is Manx customary law, a form of common law. Manx law originally derived from Gaelic Brehon law and Norse Udal law.","type":"Common law"}`,
		},

		LegalSystemsLawofIsrael: &LegalSystemsCode{
			Name:        "LawofIsrael",
			Label:       "Law of Israel",
			Description: "",
			MetaData:    `{"description":"Based on English common law arising from the period of the British Mandate (which includes laws arising from previous Ottoman rule), also incorporating civil law and fragments of Halakha and Sharia for family law cases.","type":"Common law"}`,
		},

		LegalSystemsLawofItaly: &LegalSystemsCode{
			Name:        "LawofItaly",
			Label:       "Law of Italy",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law, with elements of the Napoleonic civil code; civil code of 1942 replaced the original one of 1865.","type":"Civil law"}`,
		},

		LegalSystemsLawofCotedIvoire: &LegalSystemsCode{
			Name:        "LawofCotedIvoire",
			Label:       "Law of Cote d'Ivoire",
			Description: "",
			MetaData:    `{"description":"Based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofJamaica: &LegalSystemsCode{
			Name:        "LawofJamaica",
			Label:       "Law of Jamaica",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofJapan: &LegalSystemsCode{
			Name:        "LawofJapan",
			Label:       "Law of Japan",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law. Japanese civil code of 1895.","type":"Civil law"}`,
		},

		LegalSystemsLawofJersey: &LegalSystemsCode{
			Name:        "LawofJersey",
			Label:       "Law of Jersey",
			Description: "",
			MetaData:    `{"description":"The Bailiwick of Jersey's legal system draws on local legislation enacted by the States of Jersey, Norman customary law, English common law and modern French civil law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofJordan: &LegalSystemsCode{
			Name:        "LawofJordan",
			Label:       "Law of Jordan",
			Description: "",
			MetaData:    `{"description":"Mainly based on French Civil Code and Ottoman Majalla, Islamic law applicable to family law.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofKazakhstan: &LegalSystemsCode{
			Name:        "LawofKazakhstan",
			Label:       "Law of Kazakhstan",
			Description: "",
			MetaData:    `{"description":"The legal system of Kazakhstan is based on civil law and founded-on statutory legislation (the civil code), which provides for a hierarchy of legal acts.","type":"Civil law"}`,
		},

		LegalSystemsLawofKenya: &LegalSystemsCode{
			Name:        "LawofKenya",
			Label:       "Law of Kenya",
			Description: "",
			MetaData:    `{"description":"The Kenyan legal system is descended from the British Common Law system.","type":"Common law"}`,
		},

		LegalSystemsLawofKiribati: &LegalSystemsCode{
			Name:        "LawofKiribati",
			Label:       "Law of Kiribati",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofKuwait: &LegalSystemsCode{
			Name:        "LawofKuwait",
			Label:       "Law of Kuwait",
			Description: "",
			MetaData:    `{"description":"Follows the \"civil law system\" modeled on the French legal system. Sharia law governs only family law for Muslim residents.","type":"Civil law and Sharia Law"}`,
		},

		LegalSystemsLawofKyrgyzstan: &LegalSystemsCode{
			Name:        "LawofKyrgyzstan",
			Label:       "Law of Kyrgyzstan",
			Description: "",
			MetaData:    `{"description":"Civil law system, which includes features of French civil law and Russian Federation laws.","type":"Civil law"}`,
		},

		LegalSystemsLawofLaos: &LegalSystemsCode{
			Name:        "LawofLaos",
			Label:       "Law of Laos",
			Description: "",
			MetaData:    `{"description":"Inherited a typical civil law-based legal system from the French colonial administrators.","type":"Civil law"}`,
		},

		LegalSystemsLawofLatvia: &LegalSystemsCode{
			Name:        "LawofLatvia",
			Label:       "Law of Latvia",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic and German civil law, as it was historically before the Soviet occupation.","type":"Civil law"}`,
		},

		LegalSystemsLawofLebanon: &LegalSystemsCode{
			Name:        "LawofLebanon",
			Label:       "Law of Lebanon",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofLesotho: &LegalSystemsCode{
			Name:        "LawofLesotho",
			Label:       "Law of Lesotho",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1884 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to Basutoland (now Lesotho).","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofLiberia: &LegalSystemsCode{
			Name:        "LawofLiberia",
			Label:       "Law of Liberia",
			Description: "",
			MetaData:    `{"description":"Based on Anglo-American and customary law.","type":"Common law"}`,
		},

		LegalSystemsLawofLibya: &LegalSystemsCode{
			Name:        "LawofLibya",
			Label:       "Law of Libya",
			Description: "",
			MetaData:    `{"description":"Islamic law. Based on Napoleonic civil law, with Ottoman, Italian, and Egyptian sources.","type":"Civil and Religious law"}`,
		},

		LegalSystemsLawofLiechtenstein: &LegalSystemsCode{
			Name:        "LawofLiechtenstein",
			Label:       "Law of Liechtenstein",
			Description: "",
			MetaData:    `{"description":"A civil law country. The source of the law is statute, which is enacted by the parliament.","type":"Civil law"}`,
		},

		LegalSystemsLawofLithuania: &LegalSystemsCode{
			Name:        "LawofLithuania",
			Label:       "Law of Lithuania",
			Description: "",
			MetaData:    `{"description":"Modeled after Dutch civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofLuxembourg: &LegalSystemsCode{
			Name:        "LawofLuxembourg",
			Label:       "Law of Luxembourg",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofMacau: &LegalSystemsCode{
			Name:        "LawofMacau",
			Label:       "Law of Macau",
			Description: "",
			MetaData:    `{"description":"Based on the Portuguese civil law; also influenced by the law of the PRC.","type":"Civil law"}`,
		},

		LegalSystemsLawofMacedonia: &LegalSystemsCode{
			Name:        "LawofMacedonia",
			Label:       "Law of Macedonia",
			Description: "",
			MetaData:    `{"description":"Civil law system; judicial review of legislative acts.","type":"Civil law"}`,
		},

		LegalSystemsLawofMadagascar: &LegalSystemsCode{
			Name:        "LawofMadagascar",
			Label:       "Law of Madagascar",
			Description: "",
			MetaData:    `{"description":"Civil law system based on the old French civil code and customary law in matters of marriage, family, and obligation.","type":"Civil law"}`,
		},

		LegalSystemsLawofMalawi: &LegalSystemsCode{
			Name:        "LawofMalawi",
			Label:       "Law of Malawi",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law; judicial review of legislative acts in the Supreme Court of Appeal.","type":"Common law and customary law"}`,
		},

		LegalSystemsLawofMalaysia: &LegalSystemsCode{
			Name:        "LawofMalaysia",
			Label:       "Law of Malaysia",
			Description: "",
			MetaData:    `{"description":"Based on English common law, personal law based on sharia law applies to Muslims.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofMaldives: &LegalSystemsCode{
			Name:        "LawofMaldives",
			Label:       "Law of Maldives",
			Description: "",
			MetaData:    `{"description":"Based on an admixture of Islamic Law and English common law.","type":"Common law and Religious law"}`,
		},

		LegalSystemsLawofMali: &LegalSystemsCode{
			Name:        "LawofMali",
			Label:       "Law of Mali",
			Description: "",
			MetaData:    `{"description":"Derives from French civil law and customary law.","type":"Civil law and customary law"}`,
		},

		LegalSystemsLawofMalta: &LegalSystemsCode{
			Name:        "LawofMalta",
			Label:       "Law of Malta",
			Description: "",
			MetaData:    `{"description":"Napoleonic Code with influences from Italian Civil Law. English common law however is also a source of Maltese Law, most notably in Public Law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofMarshallIslands: &LegalSystemsCode{
			Name:        "LawofMarshallIslands",
			Label:       "Law of Marshall Islands",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		},

		LegalSystemsLawofMartiniqueFrance: &LegalSystemsCode{
			Name:        "LawofMartiniqueFrance",
			Label:       "Law of Martinique (France)",
			Description: "",
			MetaData:    `{"description":"The French system of justice is in force.","type":"Civil law"}`,
		},

		LegalSystemsLawofMauritania: &LegalSystemsCode{
			Name:        "LawofMauritania",
			Label:       "Law of Mauritania",
			Description: "",
			MetaData:    `{"description":"Mix of Islamic law and French Civil Codes, Islamic law largely applicable to family law.","type":"Religious law"}`,
		},

		LegalSystemsLawofMauritius: &LegalSystemsCode{
			Name:        "LawofMauritius",
			Label:       "Law of Mauritius",
			Description: "",
			MetaData:    `{"description":"Laws governing the Mauritian penal system are derived partly from French civil law and British common law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofMayotteFrance: &LegalSystemsCode{
			Name:        "LawofMayotteFrance",
			Label:       "Law of Mayotte (France)",
			Description: "",
			MetaData:    `{"description":"Local judicial, economic and social laws and customs were changed to conform with French law.","type":"Civil law"}`,
		},

		LegalSystemsLawofMexico: &LegalSystemsCode{
			Name:        "LawofMexico",
			Label:       "Law of Mexico",
			Description: "",
			MetaData:    `{"description":"The origins of Mexico's legal system are both ancient and classical, based on the Roman and French legal systems.","type":"Civil law"}`,
		},

		LegalSystemsLawofMicronesia: &LegalSystemsCode{
			Name:        "LawofMicronesia",
			Label:       "Law of Micronesia",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of common and customary law.","type":"Common law and customary law"}`,
		},

		LegalSystemsLawofMoldova: &LegalSystemsCode{
			Name:        "LawofMoldova",
			Label:       "Law of Moldova",
			Description: "",
			MetaData:    `{"description":"Civil law system with Germanic law influences; Constitutional Court review of legislative acts.","type":"Civil law"}`,
		},

		LegalSystemsLawofMonaco: &LegalSystemsCode{
			Name:        "LawofMonaco",
			Label:       "Law of Monaco",
			Description: "",
			MetaData:    `{"description":"Has its own legal system, Law Courts and Appeal Court. Its law is largely but not entirely based on the French Code Civil.","type":"Civil law"}`,
		},

		LegalSystemsLawofMongolia: &LegalSystemsCode{
			Name:        "LawofMongolia",
			Label:       "Law of Mongolia",
			Description: "",
			MetaData:    `{"description":"Based on Germanic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofMontenegro: &LegalSystemsCode{
			Name:        "LawofMontenegro",
			Label:       "Law of Montenegro",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic and German civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofMontserratEnglandandWales: &LegalSystemsCode{
			Name:        "LawofMontserratEnglandandWales",
			Label:       "Law of Montserrat (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMorocco: &LegalSystemsCode{
			Name:        "LawofMorocco",
			Label:       "Law of Morocco",
			Description: "",
			MetaData:    `{"description":"Mix of Islamic law and French Civil Codes, Islamic law largely applicable to family law. Halakha recognized to family law cases for Jewish citizens.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofMozambique: &LegalSystemsCode{
			Name:        "LawofMozambique",
			Label:       "Law of Mozambique",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofMyanmar: &LegalSystemsCode{
			Name:        "LawofMyanmar",
			Label:       "Law of Myanmar",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNamibia: &LegalSystemsCode{
			Name:        "LawofNamibia",
			Label:       "Law of Namibia",
			Description: "",
			MetaData:    `{"description":"Based on South African law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofNauru: &LegalSystemsCode{
			Name:        "LawofNauru",
			Label:       "Law of Nauru",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNepal: &LegalSystemsCode{
			Name:        "LawofNepal",
			Label:       "Law of Nepal",
			Description: "",
			MetaData:    `{"description":"English common law and Hindu legal concepts; note - new criminal and civil codes came into effect on 17 August 2018","type":"Common law, Civil law and Religous law"}`,
		},

		LegalSystemsLawofNetherlands: &LegalSystemsCode{
			Name:        "LawofNetherlands",
			Label:       "Law of Netherlands",
			Description: "",
			MetaData:    `{"description":"Based on Napoleonic code with German law influence.","type":"Civil law"}`,
		},

		LegalSystemsLawofNewCaledonia: &LegalSystemsCode{
			Name:        "LawofNewCaledonia",
			Label:       "Law of New Caledonia",
			Description: "",
			MetaData:    `{"description":"Civil law system based on French civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofNewZealand: &LegalSystemsCode{
			Name:        "LawofNewZealand",
			Label:       "Law of New Zealand",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNicaragua: &LegalSystemsCode{
			Name:        "LawofNicaragua",
			Label:       "Law of Nicaragua",
			Description: "",
			MetaData:    `{"description":"Civil law system; Supreme Court may review administrative acts.","type":"Civil law"}`,
		},

		LegalSystemsLawofNiger: &LegalSystemsCode{
			Name:        "LawofNiger",
			Label:       "Law of Niger",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil law, based on French civil law, Islamic law, and customary law.","type":"Civil law, Religious law and Customary law"}`,
		},

		LegalSystemsLawofNigeria: &LegalSystemsCode{
			Name:        "LawofNigeria",
			Label:       "Law of Nigeria",
			Description: "",
			MetaData:    `{"description":"Sharia in the northern states, common law in the south and at the federal level.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofNiue: &LegalSystemsCode{
			Name:        "LawofNiue",
			Label:       "Law of Niue",
			Description: "",
			MetaData:    `{"description":"Certain British and New Zealand Acts; common law; customary law.","type":"Common law and customary law"}`,
		},

		LegalSystemsLawofNorfolkIslandNewSouthWales: &LegalSystemsCode{
			Name:        "LawofNorfolkIslandNewSouthWales",
			Label:       "Law of Norfolk Island (New South Wales)",
			Description: "",
			MetaData:    `{"description":"The Norfolk Island Regional Council operates under an applied version of NSW Local Government legislation, the Local Government Act 1993 (NSW)(NI).","type":"Common law"}`,
		},

		LegalSystemsLawofNorthKorea: &LegalSystemsCode{
			Name:        "LawofNorthKorea",
			Label:       "Law of North Korea",
			Description: "",
			MetaData:    `{"description":"Codified civil law system, which was inherited from colonial Japan and is similar to South Korea's system.","type":"Civil law"}`,
		},

		LegalSystemsLawofNorthernMarianaIslandsUnitedStates: &LegalSystemsCode{
			Name:        "LawofNorthernMarianaIslandsUnitedStates",
			Label:       "Law of Northern Mariana Islands (United States)",
			Description: "",
			MetaData:    `{"description":"Laws of the US apply, except for customs and some aspects of taxation.","type":"Common law"}`,
		},

		LegalSystemsLawofNorway: &LegalSystemsCode{
			Name:        "LawofNorway",
			Label:       "Law of Norway",
			Description: "",
			MetaData:    `{"description":"Based on Civil Code.","type":"Civil law"}`,
		},

		LegalSystemsLawofOman: &LegalSystemsCode{
			Name:        "LawofOman",
			Label:       "Law of Oman",
			Description: "",
			MetaData:    `{"description":"Sharia and tribal custom laws.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofPakistan: &LegalSystemsCode{
			Name:        "LawofPakistan",
			Label:       "Law of Pakistan",
			Description: "",
			MetaData:    `{"description":"Based on English Common Law, some Islamic law applications in inheritance. Tribal Law in FATA.","type":"Common law and sharia law"}`,
		},

		LegalSystemsLawofPalau: &LegalSystemsCode{
			Name:        "LawofPalau",
			Label:       "Law of Palau",
			Description: "",
			MetaData:    `{"description":"Based on law of the United States.","type":"Common law"}`,
		},

		LegalSystemsLawofPalestinianTerritory: &LegalSystemsCode{
			Name:        "LawofPalestinianTerritory",
			Label:       "Law of Palestinian Territory",
			Description: "",
			MetaData:    `{"description":"Islamic customary law applies in Palestinian ruled territory.","type":"Islamic customary law"}`,
		},

		LegalSystemsLawofPanama: &LegalSystemsCode{
			Name:        "LawofPanama",
			Label:       "Law of Panama",
			Description: "",
			MetaData:    `{"description":"Based on civil law with influences from Spanish legal tradition and Roman laws.","type":"Civil law"}`,
		},

		LegalSystemsLawofPapuaNewGuinea: &LegalSystemsCode{
			Name:        "LawofPapuaNewGuinea",
			Label:       "Law of Papua New Guinea",
			Description: "",
			MetaData:    `{"description":"The legal system is based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofParaguay: &LegalSystemsCode{
			Name:        "LawofParaguay",
			Label:       "Law of Paraguay",
			Description: "",
			MetaData:    `{"description":"The Paraguayan Civil Code in force since 1987 is largely influenced by the Napoleonic Code and the Argentinian Code.","type":"Civil law"}`,
		},

		LegalSystemsLawofPeru: &LegalSystemsCode{
			Name:        "LawofPeru",
			Label:       "Law of Peru",
			Description: "",
			MetaData:    `{"description":"Based on civil law system; accepts compulsory International Court of Justice ICJ jurisdiction with despotic and corrupting reservations.","type":"Civil law"}`,
		},

		LegalSystemsLawofPhilippines: &LegalSystemsCode{
			Name:        "LawofPhilippines",
			Label:       "Law of Philippines",
			Description: "",
			MetaData:    `{"description":"Based on Spanish law; influenced by U.S. common law after 1898 Spanish– and Philippine–American Wars, personal law based on sharia law applies to Muslims.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofPitcairnIslandsEnglandandWales: &LegalSystemsCode{
			Name:        "LawofPitcairnIslandsEnglandandWales",
			Label:       "Law of Pitcairn Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"Comprises (a) laws (called \"ordinances\") made by the Governor, (b) United Kingdom legislation and Orders in Council extending to Pitcairn, and (c) the common law, rules of equity, and statutes of general application as in force in England for the time being, so far as local circumstances and the limits of local jurisdiction permit, and subject to any existing or future ordinance.","type":"Common law and Ordinances"}`,
		},

		LegalSystemsLawofPoland: &LegalSystemsCode{
			Name:        "LawofPoland",
			Label:       "Law of Poland",
			Description: "",
			MetaData:    `{"description":"The Polish Civil Code in force since 1965.","type":"Civil law"}`,
		},

		LegalSystemsLawofPortugal: &LegalSystemsCode{
			Name:        "LawofPortugal",
			Label:       "Law of Portugal",
			Description: "",
			MetaData:    `{"description":"Influenced by the Napoleonic Code and later by the German civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofPuertoRico: &LegalSystemsCode{
			Name:        "LawofPuertoRico",
			Label:       "Law of Puerto Rico",
			Description: "",
			MetaData:    `{"description":"Based on Spanish law; influenced by U.S. common law after 1898; federal laws (based on common law) are in effect because of federal Supremacy Clause.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofQatar: &LegalSystemsCode{
			Name:        "LawofQatar",
			Label:       "Law of Qatar",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law and Egyptian civil law system (after the French civil law system).","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofRéunionFrance: &LegalSystemsCode{
			Name:        "LawofRéunionFrance",
			Label:       "Law of Réunion (France)",
			Description: "",
			MetaData:    `{"description":"Governed by French law, and its constitution is the French constitution of 28 September 1958.","type":"Civil law"}`,
		},

		LegalSystemsLawofRomania: &LegalSystemsCode{
			Name:        "LawofRomania",
			Label:       "Law of  Romania",
			Description: "",
			MetaData:    `{"description":"Civil Code came into force in 2011. Based on the Civil Code of Quebec, but also influenced by the Napoleonic Code and other French-inspired codes (such as those of Italy, Spain and Switzerland).","type":"Civil law"}`,
		},

		LegalSystemsLawofRussia: &LegalSystemsCode{
			Name:        "LawofRussia",
			Label:       "Law of Russia",
			Description: "",
			MetaData:    `{"description":"Civil Law system descendant from Roman Law through Byzantine tradition. Heavily influenced by German and Dutch norms in 1700–1800s. Socialist-style modification in 1900s, and Continental European Law influences since 1990s.","type":"Civil law"}`,
		},

		LegalSystemsLawofRwanda: &LegalSystemsCode{
			Name:        "LawofRwanda",
			Label:       "Law of Rwanda",
			Description: "",
			MetaData:    `{"description":"Mixture of Belgian civil law and English common law.","type":"Civil and common law"}`,
		},

		LegalSystemsLawofSaintHelenaAscensionandTristandaCunhaEnglandandWales: &LegalSystemsCode{
			Name:        "LawofSaintHelenaAscensionandTristandaCunhaEnglandandWales",
			Label:       "Law of Saint Helena Ascension and Tristan da Cunha (England and Wales)",
			Description: "",
			MetaData:    `{"description":"English common law and local statutes.","type":"Common law and local statues"}`,
		},

		LegalSystemsLawofSaintKittsandNevis: &LegalSystemsCode{
			Name:        "LawofSaintKittsandNevis",
			Label:       "Law of Saint Kitts and Nevis",
			Description: "",
			MetaData:    `{"description":"Based on English common law and statutory acts of the House of Assembly.","type":"Common law"}`,
		},

		LegalSystemsLawofSaintLucia: &LegalSystemsCode{
			Name:        "LawofSaintLucia",
			Label:       "Law of Saint Lucia",
			Description: "",
			MetaData:    `{"description":"Because it was both a French and English colony before it gained independence on February 11, 1979, the legal system is a mix of civil and English common law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSaintPierreandMiquelonFrance: &LegalSystemsCode{
			Name:        "LawofSaintPierreandMiquelonFrance",
			Label:       "Law of Saint Pierre and Miquelon (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSaintVincentandtheGrenadines: &LegalSystemsCode{
			Name:        "LawofSaintVincentandtheGrenadines",
			Label:       "Law of Saint Vincent and the Grenadines",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofSaintBarthelemyFrance: &LegalSystemsCode{
			Name:        "LawofSaintBarthelemyFrance",
			Label:       "Law of Saint-Barthelemy (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSaintMartinFrance: &LegalSystemsCode{
			Name:        "LawofSaintMartinFrance",
			Label:       "Law of Saint-Martin (France)",
			Description: "",
			MetaData:    `{"description":"French civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSamoa: &LegalSystemsCode{
			Name:        "LawofSamoa",
			Label:       "Law of Samoa",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law; judicial review of legislative acts with respect to fundamental rights of the citizen.","type":"Common law and customary law"}`,
		},

		LegalSystemsLawofSanMarino: &LegalSystemsCode{
			Name:        "LawofSanMarino",
			Label:       "Law of San Marino",
			Description: "",
			MetaData:    `{"description":"Civil law system with Italian civil law influences.","type":"Civil law"}`,
		},

		LegalSystemsLawofSãoToméandPríncipe: &LegalSystemsCode{
			Name:        "LawofSãoToméandPríncipe",
			Label:       "Law of São Tomé and Príncipe",
			Description: "",
			MetaData:    `{"description":"Based on Portuguese civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSaudiArabia: &LegalSystemsCode{
			Name:        "LawofSaudiArabia",
			Label:       "Law of Saudi Arabia",
			Description: "",
			MetaData:    `{"description":"Islamic law.","type":"Religious law"}`,
		},

		LegalSystemsLawofSenegal: &LegalSystemsCode{
			Name:        "LawofSenegal",
			Label:       "Law of Senegal",
			Description: "",
			MetaData:    `{"description":"Civil law system based on French law; judicial review of legislative acts in Constitutional Court.","type":"Civil law"}`,
		},

		LegalSystemsLawofSerbia: &LegalSystemsCode{
			Name:        "LawofSerbia",
			Label:       "Law of Serbia",
			Description: "",
			MetaData:    `{"description":"The Swiss civil law (Zivilgesetzbuch) was a model for the Law on Obligations of 1978.","type":"Civil law"}`,
		},

		LegalSystemsLawofSeychelles: &LegalSystemsCode{
			Name:        "LawofSeychelles",
			Label:       "Law of Seychelles",
			Description: "",
			MetaData:    `{"description":"The substantive civil law is based on the French Civil Code. Otherwise the criminal law and court procedure are based on the English common law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSierraLeoneEnglandandWales: &LegalSystemsCode{
			Name:        "LawofSierraLeoneEnglandandWales",
			Label:       "Law of Sierra Leone (England and Wales)",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law.","type":"Common law and Customary law"}`,
		},

		LegalSystemsLawofSingapore: &LegalSystemsCode{
			Name:        "LawofSingapore",
			Label:       "Law of Singapore",
			Description: "",
			MetaData:    `{"description":"Based on English common law, but Muslims are subject to the Administration of Muslim Law Act, which gives the Sharia Court jurisdiction over Muslim personal law.","type":"Common law"}`,
		},

		LegalSystemsLawofSintMaarten: &LegalSystemsCode{
			Name:        "LawofSintMaarten",
			Label:       "Law of Sint Maarten",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law system with some English common law influence.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSlovakia: &LegalSystemsCode{
			Name:        "LawofSlovakia",
			Label:       "Law of Slovakia",
			Description: "",
			MetaData:    `{"description":"Descended from the Civil Code of the Austrian Empire (1811), influenced by German (1939–45) and Soviet (1947/68–89) legal codes during occupation periods, substantially reformed to remove Soviet influence and elements of socialist law after the Velvet Revolution (1989).","type":"Civil law"}`,
		},

		LegalSystemsLawofSlovenia: &LegalSystemsCode{
			Name:        "LawofSlovenia",
			Label:       "Law of Slovenia",
			Description: "",
			MetaData:    `{"description":"A Civil Law system influenced mostly by Germanic and Austro-Hungarian law systems.","type":"Civil law"}`,
		},

		LegalSystemsLawofSolomonIslands: &LegalSystemsCode{
			Name:        "LawofSolomonIslands",
			Label:       "Law of Solomon Islands",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of English common law and customary law.","type":"Common law and Customary law"}`,
		},

		LegalSystemsLawofSomalia: &LegalSystemsCode{
			Name:        "LawofSomalia",
			Label:       "Law of Somalia",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of civil law, Islamic (sharia) law, and customary law (referred to as Xeer).","type":"Civil law, Sharia law and Customary law"}`,
		},

		LegalSystemsLawofSouthAfrica: &LegalSystemsCode{
			Name:        "LawofSouthAfrica",
			Label:       "Law of South Africa",
			Description: "",
			MetaData:    `{"description":"An amalgam of Roman-Dutch civil law and English common law, as well as Customary Law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales: &LegalSystemsCode{
			Name:        "LawofSouthGeorgiaandtheSouthSandwichIslandsEnglandandWales",
			Label:       "Law of South Georgia and the South Sandwich Islands (England and Wales)",
			Description: "",
			MetaData:    `{"description":"The laws of the UK, where applicable, apply; the senior magistrate from the Falkland Islands presides over the Magistrates Court.","type":"Common law"}`,
		},

		LegalSystemsLawofSouthKorea: &LegalSystemsCode{
			Name:        "LawofSouthKorea",
			Label:       "Law of South Korea",
			Description: "",
			MetaData:    `{"description":"Based on German civil law system. Also largely influenced by Japanese civil law which itself modelled after German one. Korean Civil Code was introduced 1958 and fully enacted by 1960.","type":"Civil law"}`,
		},

		LegalSystemsLawofSouthSudan: &LegalSystemsCode{
			Name:        "LawofSouthSudan",
			Label:       "Law of South Sudan",
			Description: "",
			MetaData:    `{"description":"Legal system is built on the combination of statutory and customary laws. South Sudan has enacted dozens of laws since 2005, but their use in legal disputes and courts is limited.","type":"Statutory and Customary law"}`,
		},

		LegalSystemsLawofSpain: &LegalSystemsCode{
			Name:        "LawofSpain",
			Label:       "Law of Spain",
			Description: "",
			MetaData:    `{"description":"An amalgam of Roman-Dutch civil law and English common law, as well as Customary Law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSriLanka: &LegalSystemsCode{
			Name:        "LawofSriLanka",
			Label:       "Law of Sri Lanka",
			Description: "",
			MetaData:    `{"description":"An amalgam of English common law, Roman-Dutch civil law and Customary Law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSudan: &LegalSystemsCode{
			Name:        "LawofSudan",
			Label:       "Law of Sudan",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law.","type":"Religious law"}`,
		},

		LegalSystemsLawofSuriname: &LegalSystemsCode{
			Name:        "LawofSuriname",
			Label:       "Law of Suriname",
			Description: "",
			MetaData:    `{"description":"Based on Dutch civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSvalbardandJanMayenNorway: &LegalSystemsCode{
			Name:        "LawofSvalbardandJanMayenNorway",
			Label:       "Law of Svalbard and Jan Mayen (Norway)",
			Description: "",
			MetaData:    `{"description":"The laws of Norway where applicable apply; only the laws of Norway made explicitly applicable to Svalbard have effect there; the Svalbard Act and the Svalbard Environmental Protection Act, and certain regulations, apply only to Svalbard; the Spitsbergen Treaty and the Svalbard Treaty grant certain rights to citizens and corporations of signatory nations; as of June 2017, 45 nations had ratified the Svalbard Treaty.","type":"Civil law"}`,
		},

		LegalSystemsLawofEswatini: &LegalSystemsCode{
			Name:        "LawofEswatini",
			Label:       "Law of Eswatini",
			Description: "",
			MetaData:    `{"description":"Based on South African law. A 1907 proclamation by the High Commissioner for Southern Africa applied the Roman-Dutch common law of the Transvaal Colony (now part of South Africa) to the Swaziland Protectorate (now Eswatini).","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofSweden: &LegalSystemsCode{
			Name:        "LawofSweden",
			Label:       "Law of Sweden",
			Description: "",
			MetaData:    `{"description":"Scandinavian-North Germanic civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofSwitzerland: &LegalSystemsCode{
			Name:        "LawofSwitzerland",
			Label:       "Law of Switzerland",
			Description: "",
			MetaData:    `{"description":"The Swiss Civil Code of 1908 and 1912 (obligations; fifth book).","type":"Civil law"}`,
		},

		LegalSystemsLawofSyria: &LegalSystemsCode{
			Name:        "LawofSyria",
			Label:       "Law of Syria",
			Description: "",
			MetaData:    `{"description":"Mainly based on Napoleonic/French Civil Code. Islamic law applicable to family law. Non Muslims follow their own family laws.","type":"Civil law and sharia law"}`,
		},

		LegalSystemsLawofTaiwan: &LegalSystemsCode{
			Name:        "LawofTaiwan",
			Label:       "Law of Taiwan",
			Description: "",
			MetaData:    `{"description":"Influenced by German Civil Code. Enacted in 1931.","type":"Civil law"}`,
		},

		LegalSystemsLawofTajikistan: &LegalSystemsCode{
			Name:        "LawofTajikistan",
			Label:       "Law of Tajikistan",
			Description: "",
			MetaData:    `{"description":"Civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofTanzania: &LegalSystemsCode{
			Name:        "LawofTanzania",
			Label:       "Law of Tanzania",
			Description: "",
			MetaData:    `{"description":"Tanzania’s legal system is based on the English Common Law system. It derived this system from its British colonial legacy, as it does the system of government, which is based to a large degree on the Westminster parliamentary model.","type":"Common law"}`,
		},

		LegalSystemsLawofThailand: &LegalSystemsCode{
			Name:        "LawofThailand",
			Label:       "Law of Thailand",
			Description: "",
			MetaData:    `{"description":"The Thai legal system became an amalgam of German, Swiss, French, English, Japanese, Italian, Indian and American laws and practices. Even today, Islamic laws and practices exist in four southern provinces.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofTogo: &LegalSystemsCode{
			Name:        "LawofTogo",
			Label:       "Law of Togo",
			Description: "",
			MetaData:    `{"description":"Togo has a civil law system based on the French model mixed with customary law.","type":"Civil law and customary law"}`,
		},

		LegalSystemsLawofTokelau: &LegalSystemsCode{
			Name:        "LawofTokelau",
			Label:       "Law of Tokelau",
			Description: "",
			MetaData:    `{"description":"Certain British and NZ Acts of Parliament; Rules of the Tokelau General fono; British common law as at 14 January 1840.","type":"Common law"}`,
		},

		LegalSystemsLawofTrinidadandTobago: &LegalSystemsCode{
			Name:        "LawofTrinidadandTobago",
			Label:       "Law of Trinidad and Tobago",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofTunisia: &LegalSystemsCode{
			Name:        "LawofTunisia",
			Label:       "Law of Tunisia",
			Description: "",
			MetaData:    `{"description":"The Tunisian legal system is based on French civil law system.","type":"Civil law"}`,
		},

		LegalSystemsLawofTurkey: &LegalSystemsCode{
			Name:        "LawofTurkey",
			Label:       "Law of Turkey",
			Description: "",
			MetaData:    `{"description":"Modeled after the Swiss civil law (Zivilgesetzbuch) of 1907.","type":"Civil law"}`,
		},

		LegalSystemsLawofTurkmenistan: &LegalSystemsCode{
			Name:        "LawofTurkmenistan",
			Label:       "Law of Turkmenistan",
			Description: "",
			MetaData:    `{"description":"The legal system in Turkmenistan is a civil law system, with Islamic law influences.","type":"Civil law and Islamic law"}`,
		},

		LegalSystemsLawofTurksandCaicosIsland: &LegalSystemsCode{
			Name:        "LawofTurksandCaicosIsland",
			Label:       "Law of Turks and Caicos Island",
			Description: "",
			MetaData:    `{"description":"A British Overseas Territory and a common law jurisdiction which is modeled on the English legal system.","type":"Common Law"}`,
		},

		LegalSystemsLawofTuvalu: &LegalSystemsCode{
			Name:        "LawofTuvalu",
			Label:       "Law of Tuvalu",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofUganda: &LegalSystemsCode{
			Name:        "LawofUganda",
			Label:       "Law of Uganda",
			Description: "",
			MetaData:    `{"description":"Based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofUkraine: &LegalSystemsCode{
			Name:        "LawofUkraine",
			Label:       "Law of Ukraine",
			Description: "",
			MetaData:    `{"description":"Civil Code of Ukraine of 2004.","type":"Civil law"}`,
		},

		LegalSystemsLawofUnitedArabEmirates: &LegalSystemsCode{
			Name:        "LawofUnitedArabEmirates",
			Label:       "Law of United Arab Emirates",
			Description: "",
			MetaData:    `{"description":"Based on Islamic law and Egyptian civil law system (after the French civil law system). Based on Common law system in the Dubai International Financial Center (DIFC Courts) and Abu Dhabi Global Market (ADGM) Courts (after the English Common law system).","type":"Common Law, civil law and sharia law"}`,
		},

		LegalSystemsLawofNorthernIreland: &LegalSystemsCode{
			Name:        "LawofNorthernIreland",
			Label:       "Law of Northern Ireland",
			Description: "",
			MetaData:    `{"description":"Based on Irish law before 1921, in turn based on English common law.","type":"Common law"}`,
		},

		LegalSystemsLawofUruguay: &LegalSystemsCode{
			Name:        "LawofUruguay",
			Label:       "Law of Uruguay",
			Description: "",
			MetaData:    `{"description":"Civil law system, with public law based on the 1967 Constitution, amended in 1989, 1994, 1997, and 2004.","type":"Civil law"}`,
		},

		LegalSystemsLawofUSVirginIslandsUnitedStates: &LegalSystemsCode{
			Name:        "LawofUSVirginIslandsUnitedStates",
			Label:       "Law of US Virgin Islands (United States)",
			Description: "",
			MetaData:    `{"description":"The virgin Islands legal system transitioned from danish law to the American common law tradition.","type":"Common law"}`,
		},

		LegalSystemsLawofUnitedStates: &LegalSystemsCode{
			Name:        "LawofUnitedStates",
			Label:       "Law of United States",
			Description: "",
			MetaData:    `{"description":"Federal courts and 49 states use the legal system based on English common law. Law in the state of Louisiana is based on French and Spanish civil law.","type":"Common law"}`,
		},

		LegalSystemsLawofAlaska: &LegalSystemsCode{
			Name:        "LawofAlaska",
			Label:       "Law of Alaska",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofAlabama: &LegalSystemsCode{
			Name:        "LawofAlabama",
			Label:       "Law of Alabama",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofArkansas: &LegalSystemsCode{
			Name:        "LawofArkansas",
			Label:       "Law of Arkansas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofArizona: &LegalSystemsCode{
			Name:        "LawofArizona",
			Label:       "Law of Arizona",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofCalifornia: &LegalSystemsCode{
			Name:        "LawofCalifornia",
			Label:       "Law of California",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofColorado: &LegalSystemsCode{
			Name:        "LawofColorado",
			Label:       "Law of Colorado",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofConnecticut: &LegalSystemsCode{
			Name:        "LawofConnecticut",
			Label:       "Law of Connecticut",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWashingtonDC: &LegalSystemsCode{
			Name:        "LawofWashingtonDC",
			Label:       "Law of Washington D.C.",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofDelaware: &LegalSystemsCode{
			Name:        "LawofDelaware",
			Label:       "Law of Delaware",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofFlorida: &LegalSystemsCode{
			Name:        "LawofFlorida",
			Label:       "Law of Florida",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofGeorgiaUSA: &LegalSystemsCode{
			Name:        "LawofGeorgiaUSA",
			Label:       "Law of Georgia (USA)",
			Description: "",
			MetaData:    `{"description":"Georgia's legal system is based on common law.","type":"Common law"}`,
		},

		LegalSystemsLawofHawaii: &LegalSystemsCode{
			Name:        "LawofHawaii",
			Label:       "Law of Hawaii",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofIowa: &LegalSystemsCode{
			Name:        "LawofIowa",
			Label:       "Law of Iowa",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofIdaho: &LegalSystemsCode{
			Name:        "LawofIdaho",
			Label:       "Law of Idaho",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofIllinois: &LegalSystemsCode{
			Name:        "LawofIllinois",
			Label:       "Law of Illinois",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofIndiana: &LegalSystemsCode{
			Name:        "LawofIndiana",
			Label:       "Law of Indiana",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofKansas: &LegalSystemsCode{
			Name:        "LawofKansas",
			Label:       "Law of Kansas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofKentucky: &LegalSystemsCode{
			Name:        "LawofKentucky",
			Label:       "Law of Kentucky",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofLouisiana: &LegalSystemsCode{
			Name:        "LawofLouisiana",
			Label:       "Law of Louisiana",
			Description: "",
			MetaData:    `{"description":"Based on French and Spanish civil law, but federal laws (based on common law) are also in effect in Louisiana because of federal Supremacy Clause.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofMassachusetts: &LegalSystemsCode{
			Name:        "LawofMassachusetts",
			Label:       "Law of Massachusetts",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMaryland: &LegalSystemsCode{
			Name:        "LawofMaryland",
			Label:       "Law of Maryland",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMaine: &LegalSystemsCode{
			Name:        "LawofMaine",
			Label:       "Law of Maine",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMichigan: &LegalSystemsCode{
			Name:        "LawofMichigan",
			Label:       "Law of Michigan",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMinnesota: &LegalSystemsCode{
			Name:        "LawofMinnesota",
			Label:       "Law of Minnesota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMissouri: &LegalSystemsCode{
			Name:        "LawofMissouri",
			Label:       "Law of Missouri",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMississippi: &LegalSystemsCode{
			Name:        "LawofMississippi",
			Label:       "Law of Mississippi",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofMontana: &LegalSystemsCode{
			Name:        "LawofMontana",
			Label:       "Law of Montana",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNorthCarolina: &LegalSystemsCode{
			Name:        "LawofNorthCarolina",
			Label:       "Law of North Carolina",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNorthDakota: &LegalSystemsCode{
			Name:        "LawofNorthDakota",
			Label:       "Law of North Dakota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNebraska: &LegalSystemsCode{
			Name:        "LawofNebraska",
			Label:       "Law of Nebraska",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewHampshire: &LegalSystemsCode{
			Name:        "LawofNewHampshire",
			Label:       "Law of New Hampshire",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewJersey: &LegalSystemsCode{
			Name:        "LawofNewJersey",
			Label:       "Law of New Jersey",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewMexico: &LegalSystemsCode{
			Name:        "LawofNewMexico",
			Label:       "Law of New Mexico",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNevada: &LegalSystemsCode{
			Name:        "LawofNevada",
			Label:       "Law of Nevada",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofNewYork: &LegalSystemsCode{
			Name:        "LawofNewYork",
			Label:       "Law of New York",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofOhio: &LegalSystemsCode{
			Name:        "LawofOhio",
			Label:       "Law of Ohio",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofOklahoma: &LegalSystemsCode{
			Name:        "LawofOklahoma",
			Label:       "Law of Oklahoma",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofOregon: &LegalSystemsCode{
			Name:        "LawofOregon",
			Label:       "Law of Oregon",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofPennsylvania: &LegalSystemsCode{
			Name:        "LawofPennsylvania",
			Label:       "Law of Pennsylvania",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofRhodeIsland: &LegalSystemsCode{
			Name:        "LawofRhodeIsland",
			Label:       "Law of Rhode Island",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofSouthCarolina: &LegalSystemsCode{
			Name:        "LawofSouthCarolina",
			Label:       "Law of South Carolina",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofSouthDakota: &LegalSystemsCode{
			Name:        "LawofSouthDakota",
			Label:       "Law of South Dakota",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofTennessee: &LegalSystemsCode{
			Name:        "LawofTennessee",
			Label:       "Law of Tennessee",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofTexas: &LegalSystemsCode{
			Name:        "LawofTexas",
			Label:       "Law of Texas",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofUtah: &LegalSystemsCode{
			Name:        "LawofUtah",
			Label:       "Law of Utah",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofVirginia: &LegalSystemsCode{
			Name:        "LawofVirginia",
			Label:       "Law of Virginia",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofVermont: &LegalSystemsCode{
			Name:        "LawofVermont",
			Label:       "Law of Vermont",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWashington: &LegalSystemsCode{
			Name:        "LawofWashington",
			Label:       "Law of Washington",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWisconsin: &LegalSystemsCode{
			Name:        "LawofWisconsin",
			Label:       "Law of Wisconsin",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWestVirginia: &LegalSystemsCode{
			Name:        "LawofWestVirginia",
			Label:       "Law of West Virginia",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofWyoming: &LegalSystemsCode{
			Name:        "LawofWyoming",
			Label:       "Law of Wyoming",
			Description: "",
			MetaData:    `{"description":"Based on english common law.","type":"Common law"}`,
		},

		LegalSystemsLawofUzbekistan: &LegalSystemsCode{
			Name:        "LawofUzbekistan",
			Label:       "Law of Uzbekistan",
			Description: "",
			MetaData:    `{"description":"Represents an evolution of Soviet civil law. Overwhelmingly strong impact of the Communist legal theory is traceable.","type":"Civil law"}`,
		},

		LegalSystemsLawofVanuatu: &LegalSystemsCode{
			Name:        "LawofVanuatu",
			Label:       "Law of Vanuatu",
			Description: "",
			MetaData:    `{"description":"Consists of a mixed system combining the legacy of English common law, French civil law and indigenous customary law.","type":"Civil law and common law"}`,
		},

		LegalSystemsLawofVaticanCity: &LegalSystemsCode{
			Name:        "LawofVaticanCity",
			Label:       "Law of Vatican City",
			Description: "",
			MetaData:    `{"description":"Based on Roman \u0026 Italian civil law and Catholic canon law.","type":"Civil law and canon law"}`,
		},

		LegalSystemsLawofVenezuela: &LegalSystemsCode{
			Name:        "LawofVenezuela",
			Label:       "Law of Venezuela",
			Description: "",
			MetaData:    `{"description":"Venezuela's legal system has a legislative origin, grounded on \"written law\" (civil law).","type":"Civil law"}`,
		},

		LegalSystemsLawofVietnam: &LegalSystemsCode{
			Name:        "LawofVietnam",
			Label:       "Law of Vietnam",
			Description: "",
			MetaData:    `{"description":"Communist legal theory and French civil law.","type":"Civil law"}`,
		},

		LegalSystemsLawofWallisandFutunaFrance: &LegalSystemsCode{
			Name:        "LawofWallisandFutunaFrance",
			Label:       "Law of Wallis and Futuna (France)",
			Description: "",
			MetaData:    `{"description":"Uses both the French legal system and customary local laws.","type":"Civil law and Customary law"}`,
		},

		LegalSystemsLawofYemen: &LegalSystemsCode{
			Name:        "LawofYemen",
			Label:       "Law of Yemen",
			Description: "",
			MetaData:    `{"description":"Mixed legal system of Islamic (sharia) law, Napoleonic law, English common law, and customary law.","type":"Civil, Common, Customary and Religious law"}`,
		},

		LegalSystemsLawofZambia: &LegalSystemsCode{
			Name:        "LawofZambia",
			Label:       "Law of Zambia",
			Description: "",
			MetaData:    `{"description":"Zambian legal system is largely based on English Common Law. Customary laws relating to particular communities around the country are also recognised by the constitution.","type":"Common law and Customary law"}`,
		},

		LegalSystemsLawofZimbabwe: &LegalSystemsCode{
			Name:        "LawofZimbabwe",
			Label:       "Law of Zimbabwe",
			Description: "",
			MetaData:    `{"description":"Based on South African law. An 1891 proclamation by the High Commissioner for Southern Africa applied the law of the Cape Colony (now part of South Africa) to Southern Rhodesia (now Zimbabwe).","type":"Civil law and common law"}`,
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

/************************************ Rejections ************************************/
const (

	// Success - No failure. This code should not be used in a reject message.
	RejectionsSuccess = 0

	// MsgMalformed - The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.
	RejectionsMsgMalformed = 1

	// TxMalformed - The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.
	RejectionsTxMalformed = 2

	// Timeout - A dependency, other contract/service, has failed to complete before the smart contract's timeout.
	RejectionsTimeout = 3

	// ContractMoved - The contract has been moved to a different address. Please find the addres change message and send requests to new address.
	RejectionsContractMoved = 4

	// DoubleSpend - A double spend attempt has been detected.
	RejectionsDoubleSpend = 5

	// Deprecated - Request or part of request has been deprecated and is no longer supported.
	RejectionsDeprecated = 6

	// Inactive - The smart contract agent is inactive due to non payment or other reason.
	RejectionsInactive = 7

	// NotConfigured - The smart contract agent is missing necessary configuration information.
	RejectionsNotConfigured = 8

	// ContractExists - The contract already exists and can't be recreated.
	RejectionsContractExists = 10

	// ContractDoesNotExist - The contract address specified does not have an active contract.
	RejectionsContractDoesNotExist = 11

	// ContractInstrumentQtyReduction - Sent when a CA tries to reduce the number of allowed instruments below the number of instruments that already exist for this contract.
	RejectionsContractInstrumentQtyReduction = 12

	// ContractFixedQuantity - Sent when the administration attempted to increase the quantity of instruments in a contract beyond the maximum number allowed.
	RejectionsContractFixedQuantity = 13

	// ContractPermissions - The contract permissions prohibit the action requested.
	RejectionsContractPermissions = 14

	// ContractExpired - The contract is expired so can no longer accept requests.
	RejectionsContractExpired = 15

	// ContractFrozen - The contract is frozen and the request is not permitted while frozen.
	RejectionsContractFrozen = 16

	// ContractRevision - The revision in a contract amendment is incorrect.
	RejectionsContractRevision = 17

	// ContractNotPermitted - Action not permitted by contract.
	RejectionsContractNotPermitted = 18

	// ContractBothOperatorsRequired - Both operators signatures are required to perform this action.
	RejectionsContractBothOperatorsRequired = 19

	// InstrumentCodeExists - The instrument code specified already exists and can't be reused.
	RejectionsInstrumentCodeExists = 20

	// InstrumentNotFound - The instrument code is not found.
	RejectionsInstrumentNotFound = 21

	// InstrumentPermissions - The instrument permissions prohibit the action requested.
	RejectionsInstrumentPermissions = 22

	// InstrumentFrozen - The instrument is frozen and the request is not permitted while frozen.
	RejectionsInstrumentFrozen = 23

	// InstrumentRevision - The revision in an instrument amendment is incorrect.
	RejectionsInstrumentRevision = 24

	// InstrumentNotPermitted - Action not permitted by instrument.
	RejectionsInstrumentNotPermitted = 25

	// TransferSelf - Transfers with the sender and receiver addresses the same are not permitted.
	RejectionsTransferSelf = 30

	// TransferExpired - The transfer has expired.
	RejectionsTransferExpired = 31

	// HoldingsFrozen - Holdings are frozen, so the request can't be completed.
	RejectionsHoldingsFrozen = 32

	// HoldingsLocked - Holdings are locked by a multi-contract request, so the request can't be completed yet.
	RejectionsHoldingsLocked = 33

	// HolderProposalProhibited - Holders are not permitted to make proposals.
	RejectionsHolderProposalProhibited = 40

	// ProposalConflicts - The proposal conflicts with an unapplied proposal.
	RejectionsProposalConflicts = 41

	// VoteNotFound - The vote ID referenced is not found.
	RejectionsVoteNotFound = 42

	// VoteClosed - The vote has closed and ballots are no longer permitted.
	RejectionsVoteClosed = 43

	// BallotAlreadyCounted - The ballot has already been counted for this address.
	RejectionsBallotAlreadyCounted = 44

	// VoteSystemNotPermitted - The voting system isn't permitted for this request.
	RejectionsVoteSystemNotPermitted = 45

	// InsufficientTxFeeFunding - Insufficient bitcoin quantities for response transaction fees.
	RejectionsInsufficientTxFeeFunding = 60

	// InsufficientValue - Insufficient bitcoin quantity in inputs to fund request.
	RejectionsInsufficientValue = 61

	// InsufficientQuantity - Insufficient token holdings to for request.
	RejectionsInsufficientQuantity = 62

	// NotAdministration - The requestor is not the administration and is required to be for this request.
	RejectionsNotAdministration = 70

	// NotOperator - The requestor is not the operator and is required to be for this request.
	RejectionsNotOperator = 71

	// UnauthorizedAddress - The address specified is not permitted for this request.
	RejectionsUnauthorizedAddress = 72

	// InvalidSignature - The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.
	RejectionsInvalidSignature = 80

	// SignatureNotSigHashAll - The transaction authorization signature doesn't sign the entire transaction so doesn't necessarily sign the action requested.
	RejectionsSignatureNotSigHashAll = 81

	// AgreementExists - The agreement already exists and can't be recreated.
	RejectionsAgreementExists = 90

	// AgreementDoesNotExist - The agreement address specified does not have an active agreement.
	RejectionsAgreementDoesNotExist = 91

	// AgreementRevision - The agreement revision in the request is not current.
	RejectionsAgreementRevision = 93
)

type RejectionsCode struct {
	Name        string
	Label       string
	Description string
	MetaData    string
}

// RejectionsData holds a mapping of Rejections codes.
func RejectionsData(code uint32) *RejectionsCode {
	switch code {

	case RejectionsSuccess:
		return &RejectionsCode{
			Name:        "Success",
			Label:       "Success",
			Description: "No failure. This code should not be used in a reject message.",
			MetaData:    `{}`,
		}

	case RejectionsMsgMalformed:
		return &RejectionsCode{
			Name:        "MsgMalformed",
			Label:       "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.",
			MetaData:    `{}`,
		}

	case RejectionsTxMalformed:
		return &RejectionsCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
			MetaData:    `{}`,
		}

	case RejectionsTimeout:
		return &RejectionsCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract's timeout.",
			MetaData:    `{}`,
		}

	case RejectionsContractMoved:
		return &RejectionsCode{
			Name:        "ContractMoved",
			Label:       "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
			MetaData:    `{}`,
		}

	case RejectionsDoubleSpend:
		return &RejectionsCode{
			Name:        "DoubleSpend",
			Label:       "Double Spend",
			Description: "A double spend attempt has been detected.",
			MetaData:    `{}`,
		}

	case RejectionsDeprecated:
		return &RejectionsCode{
			Name:        "Deprecated",
			Label:       "Deprecated",
			Description: "Request or part of request has been deprecated and is no longer supported.",
			MetaData:    `{}`,
		}

	case RejectionsInactive:
		return &RejectionsCode{
			Name:        "Inactive",
			Label:       "Inactive",
			Description: "The smart contract agent is inactive due to non payment or other reason.",
			MetaData:    `{}`,
		}

	case RejectionsNotConfigured:
		return &RejectionsCode{
			Name:        "NotConfigured",
			Label:       "Not Configured",
			Description: "The smart contract agent is missing necessary configuration information.",
			MetaData:    `{}`,
		}

	case RejectionsContractExists:
		return &RejectionsCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can't be recreated.",
			MetaData:    `{}`,
		}

	case RejectionsContractDoesNotExist:
		return &RejectionsCode{
			Name:        "ContractDoesNotExist",
			Label:       "Contract Does Not Exist",
			Description: "The contract address specified does not have an active contract.",
			MetaData:    `{}`,
		}

	case RejectionsContractInstrumentQtyReduction:
		return &RejectionsCode{
			Name:        "ContractInstrumentQtyReduction",
			Label:       "Contract Instrument Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed instruments below the number of instruments that already exist for this contract.",
			MetaData:    `{}`,
		}

	case RejectionsContractFixedQuantity:
		return &RejectionsCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of instruments in a contract beyond the maximum number allowed.",
			MetaData:    `{}`,
		}

	case RejectionsContractPermissions:
		return &RejectionsCode{
			Name:        "ContractPermissions",
			Label:       "Contract Permissions Prohibit",
			Description: "The contract permissions prohibit the action requested.",
			MetaData:    `{}`,
		}

	case RejectionsContractExpired:
		return &RejectionsCode{
			Name:        "ContractExpired",
			Label:       "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
			MetaData:    `{}`,
		}

	case RejectionsContractFrozen:
		return &RejectionsCode{
			Name:        "ContractFrozen",
			Label:       "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
			MetaData:    `{}`,
		}

	case RejectionsContractRevision:
		return &RejectionsCode{
			Name:        "ContractRevision",
			Label:       "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
			MetaData:    `{}`,
		}

	case RejectionsContractNotPermitted:
		return &RejectionsCode{
			Name:        "ContractNotPermitted",
			Label:       "Contract Not Permitted",
			Description: "Action not permitted by contract.",
			MetaData:    `{}`,
		}

	case RejectionsContractBothOperatorsRequired:
		return &RejectionsCode{
			Name:        "ContractBothOperatorsRequired",
			Label:       "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentCodeExists:
		return &RejectionsCode{
			Name:        "InstrumentCodeExists",
			Label:       "Instrument Code Already Exists",
			Description: "The instrument code specified already exists and can't be reused.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentNotFound:
		return &RejectionsCode{
			Name:        "InstrumentNotFound",
			Label:       "Instrument Not Found",
			Description: "The instrument code is not found.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentPermissions:
		return &RejectionsCode{
			Name:        "InstrumentPermissions",
			Label:       "Instrument Permissions Prohibit",
			Description: "The instrument permissions prohibit the action requested.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentFrozen:
		return &RejectionsCode{
			Name:        "InstrumentFrozen",
			Label:       "Instrument Frozen",
			Description: "The instrument is frozen and the request is not permitted while frozen.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentRevision:
		return &RejectionsCode{
			Name:        "InstrumentRevision",
			Label:       "Instrument Revision Incorrect",
			Description: "The revision in an instrument amendment is incorrect.",
			MetaData:    `{}`,
		}

	case RejectionsInstrumentNotPermitted:
		return &RejectionsCode{
			Name:        "InstrumentNotPermitted",
			Label:       "Instrument Not Permitted",
			Description: "Action not permitted by instrument.",
			MetaData:    `{}`,
		}

	case RejectionsTransferSelf:
		return &RejectionsCode{
			Name:        "TransferSelf",
			Label:       "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
			MetaData:    `{}`,
		}

	case RejectionsTransferExpired:
		return &RejectionsCode{
			Name:        "TransferExpired",
			Label:       "Transfer Expired",
			Description: "The transfer has expired.",
			MetaData:    `{}`,
		}

	case RejectionsHoldingsFrozen:
		return &RejectionsCode{
			Name:        "HoldingsFrozen",
			Label:       "Holdings Frozen",
			Description: "Holdings are frozen, so the request can't be completed.",
			MetaData:    `{}`,
		}

	case RejectionsHoldingsLocked:
		return &RejectionsCode{
			Name:        "HoldingsLocked",
			Label:       "Holdings Locked",
			Description: "Holdings are locked by a multi-contract request, so the request can't be completed yet.",
			MetaData:    `{}`,
		}

	case RejectionsHolderProposalProhibited:
		return &RejectionsCode{
			Name:        "HolderProposalProhibited",
			Label:       "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
			MetaData:    `{}`,
		}

	case RejectionsProposalConflicts:
		return &RejectionsCode{
			Name:        "ProposalConflicts",
			Label:       "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
			MetaData:    `{}`,
		}

	case RejectionsVoteNotFound:
		return &RejectionsCode{
			Name:        "VoteNotFound",
			Label:       "Vote Not Found",
			Description: "The vote ID referenced is not found.",
			MetaData:    `{}`,
		}

	case RejectionsVoteClosed:
		return &RejectionsCode{
			Name:        "VoteClosed",
			Label:       "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
			MetaData:    `{}`,
		}

	case RejectionsBallotAlreadyCounted:
		return &RejectionsCode{
			Name:        "BallotAlreadyCounted",
			Label:       "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
			MetaData:    `{}`,
		}

	case RejectionsVoteSystemNotPermitted:
		return &RejectionsCode{
			Name:        "VoteSystemNotPermitted",
			Label:       "Vote System Not Permitted",
			Description: "The voting system isn't permitted for this request.",
			MetaData:    `{}`,
		}

	case RejectionsInsufficientTxFeeFunding:
		return &RejectionsCode{
			Name:        "InsufficientTxFeeFunding",
			Label:       "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
			MetaData:    `{}`,
		}

	case RejectionsInsufficientValue:
		return &RejectionsCode{
			Name:        "InsufficientValue",
			Label:       "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
			MetaData:    `{}`,
		}

	case RejectionsInsufficientQuantity:
		return &RejectionsCode{
			Name:        "InsufficientQuantity",
			Label:       "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
			MetaData:    `{}`,
		}

	case RejectionsNotAdministration:
		return &RejectionsCode{
			Name:        "NotAdministration",
			Label:       "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
			MetaData:    `{}`,
		}

	case RejectionsNotOperator:
		return &RejectionsCode{
			Name:        "NotOperator",
			Label:       "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
			MetaData:    `{}`,
		}

	case RejectionsUnauthorizedAddress:
		return &RejectionsCode{
			Name:        "UnauthorizedAddress",
			Label:       "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
			MetaData:    `{}`,
		}

	case RejectionsInvalidSignature:
		return &RejectionsCode{
			Name:        "InvalidSignature",
			Label:       "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
			MetaData:    `{}`,
		}

	case RejectionsSignatureNotSigHashAll:
		return &RejectionsCode{
			Name:        "SignatureNotSigHashAll",
			Label:       "Signature Not Sig Hash All",
			Description: "The transaction authorization signature doesn't sign the entire transaction so doesn't necessarily sign the action requested.",
			MetaData:    `{}`,
		}

	case RejectionsAgreementExists:
		return &RejectionsCode{
			Name:        "AgreementExists",
			Label:       "Agreement Already Exists",
			Description: "The agreement already exists and can't be recreated.",
			MetaData:    `{}`,
		}

	case RejectionsAgreementDoesNotExist:
		return &RejectionsCode{
			Name:        "AgreementDoesNotExist",
			Label:       "Agreement Does Not Exist",
			Description: "The agreement address specified does not have an active agreement.",
			MetaData:    `{}`,
		}

	case RejectionsAgreementRevision:
		return &RejectionsCode{
			Name:        "AgreementRevision",
			Label:       "Agreement Revision",
			Description: "The agreement revision in the request is not current.",
			MetaData:    `{}`,
		}
	default:
		return nil
	}
}

// RejectionsMap returns a mapping of Rejections objects with the code as the key.
func RejectionsMap() map[uint32]*RejectionsCode {
	return map[uint32]*RejectionsCode{

		RejectionsSuccess: &RejectionsCode{
			Name:        "Success",
			Label:       "Success",
			Description: "No failure. This code should not be used in a reject message.",
			MetaData:    `{}`,
		},

		RejectionsMsgMalformed: &RejectionsCode{
			Name:        "MsgMalformed",
			Label:       "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.",
			MetaData:    `{}`,
		},

		RejectionsTxMalformed: &RejectionsCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
			MetaData:    `{}`,
		},

		RejectionsTimeout: &RejectionsCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract's timeout.",
			MetaData:    `{}`,
		},

		RejectionsContractMoved: &RejectionsCode{
			Name:        "ContractMoved",
			Label:       "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
			MetaData:    `{}`,
		},

		RejectionsDoubleSpend: &RejectionsCode{
			Name:        "DoubleSpend",
			Label:       "Double Spend",
			Description: "A double spend attempt has been detected.",
			MetaData:    `{}`,
		},

		RejectionsDeprecated: &RejectionsCode{
			Name:        "Deprecated",
			Label:       "Deprecated",
			Description: "Request or part of request has been deprecated and is no longer supported.",
			MetaData:    `{}`,
		},

		RejectionsInactive: &RejectionsCode{
			Name:        "Inactive",
			Label:       "Inactive",
			Description: "The smart contract agent is inactive due to non payment or other reason.",
			MetaData:    `{}`,
		},

		RejectionsNotConfigured: &RejectionsCode{
			Name:        "NotConfigured",
			Label:       "Not Configured",
			Description: "The smart contract agent is missing necessary configuration information.",
			MetaData:    `{}`,
		},

		RejectionsContractExists: &RejectionsCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can't be recreated.",
			MetaData:    `{}`,
		},

		RejectionsContractDoesNotExist: &RejectionsCode{
			Name:        "ContractDoesNotExist",
			Label:       "Contract Does Not Exist",
			Description: "The contract address specified does not have an active contract.",
			MetaData:    `{}`,
		},

		RejectionsContractInstrumentQtyReduction: &RejectionsCode{
			Name:        "ContractInstrumentQtyReduction",
			Label:       "Contract Instrument Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed instruments below the number of instruments that already exist for this contract.",
			MetaData:    `{}`,
		},

		RejectionsContractFixedQuantity: &RejectionsCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of instruments in a contract beyond the maximum number allowed.",
			MetaData:    `{}`,
		},

		RejectionsContractPermissions: &RejectionsCode{
			Name:        "ContractPermissions",
			Label:       "Contract Permissions Prohibit",
			Description: "The contract permissions prohibit the action requested.",
			MetaData:    `{}`,
		},

		RejectionsContractExpired: &RejectionsCode{
			Name:        "ContractExpired",
			Label:       "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
			MetaData:    `{}`,
		},

		RejectionsContractFrozen: &RejectionsCode{
			Name:        "ContractFrozen",
			Label:       "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
			MetaData:    `{}`,
		},

		RejectionsContractRevision: &RejectionsCode{
			Name:        "ContractRevision",
			Label:       "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
			MetaData:    `{}`,
		},

		RejectionsContractNotPermitted: &RejectionsCode{
			Name:        "ContractNotPermitted",
			Label:       "Contract Not Permitted",
			Description: "Action not permitted by contract.",
			MetaData:    `{}`,
		},

		RejectionsContractBothOperatorsRequired: &RejectionsCode{
			Name:        "ContractBothOperatorsRequired",
			Label:       "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentCodeExists: &RejectionsCode{
			Name:        "InstrumentCodeExists",
			Label:       "Instrument Code Already Exists",
			Description: "The instrument code specified already exists and can't be reused.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentNotFound: &RejectionsCode{
			Name:        "InstrumentNotFound",
			Label:       "Instrument Not Found",
			Description: "The instrument code is not found.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentPermissions: &RejectionsCode{
			Name:        "InstrumentPermissions",
			Label:       "Instrument Permissions Prohibit",
			Description: "The instrument permissions prohibit the action requested.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentFrozen: &RejectionsCode{
			Name:        "InstrumentFrozen",
			Label:       "Instrument Frozen",
			Description: "The instrument is frozen and the request is not permitted while frozen.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentRevision: &RejectionsCode{
			Name:        "InstrumentRevision",
			Label:       "Instrument Revision Incorrect",
			Description: "The revision in an instrument amendment is incorrect.",
			MetaData:    `{}`,
		},

		RejectionsInstrumentNotPermitted: &RejectionsCode{
			Name:        "InstrumentNotPermitted",
			Label:       "Instrument Not Permitted",
			Description: "Action not permitted by instrument.",
			MetaData:    `{}`,
		},

		RejectionsTransferSelf: &RejectionsCode{
			Name:        "TransferSelf",
			Label:       "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
			MetaData:    `{}`,
		},

		RejectionsTransferExpired: &RejectionsCode{
			Name:        "TransferExpired",
			Label:       "Transfer Expired",
			Description: "The transfer has expired.",
			MetaData:    `{}`,
		},

		RejectionsHoldingsFrozen: &RejectionsCode{
			Name:        "HoldingsFrozen",
			Label:       "Holdings Frozen",
			Description: "Holdings are frozen, so the request can't be completed.",
			MetaData:    `{}`,
		},

		RejectionsHoldingsLocked: &RejectionsCode{
			Name:        "HoldingsLocked",
			Label:       "Holdings Locked",
			Description: "Holdings are locked by a multi-contract request, so the request can't be completed yet.",
			MetaData:    `{}`,
		},

		RejectionsHolderProposalProhibited: &RejectionsCode{
			Name:        "HolderProposalProhibited",
			Label:       "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
			MetaData:    `{}`,
		},

		RejectionsProposalConflicts: &RejectionsCode{
			Name:        "ProposalConflicts",
			Label:       "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
			MetaData:    `{}`,
		},

		RejectionsVoteNotFound: &RejectionsCode{
			Name:        "VoteNotFound",
			Label:       "Vote Not Found",
			Description: "The vote ID referenced is not found.",
			MetaData:    `{}`,
		},

		RejectionsVoteClosed: &RejectionsCode{
			Name:        "VoteClosed",
			Label:       "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
			MetaData:    `{}`,
		},

		RejectionsBallotAlreadyCounted: &RejectionsCode{
			Name:        "BallotAlreadyCounted",
			Label:       "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
			MetaData:    `{}`,
		},

		RejectionsVoteSystemNotPermitted: &RejectionsCode{
			Name:        "VoteSystemNotPermitted",
			Label:       "Vote System Not Permitted",
			Description: "The voting system isn't permitted for this request.",
			MetaData:    `{}`,
		},

		RejectionsInsufficientTxFeeFunding: &RejectionsCode{
			Name:        "InsufficientTxFeeFunding",
			Label:       "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
			MetaData:    `{}`,
		},

		RejectionsInsufficientValue: &RejectionsCode{
			Name:        "InsufficientValue",
			Label:       "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
			MetaData:    `{}`,
		},

		RejectionsInsufficientQuantity: &RejectionsCode{
			Name:        "InsufficientQuantity",
			Label:       "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
			MetaData:    `{}`,
		},

		RejectionsNotAdministration: &RejectionsCode{
			Name:        "NotAdministration",
			Label:       "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
			MetaData:    `{}`,
		},

		RejectionsNotOperator: &RejectionsCode{
			Name:        "NotOperator",
			Label:       "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
			MetaData:    `{}`,
		},

		RejectionsUnauthorizedAddress: &RejectionsCode{
			Name:        "UnauthorizedAddress",
			Label:       "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
			MetaData:    `{}`,
		},

		RejectionsInvalidSignature: &RejectionsCode{
			Name:        "InvalidSignature",
			Label:       "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
			MetaData:    `{}`,
		},

		RejectionsSignatureNotSigHashAll: &RejectionsCode{
			Name:        "SignatureNotSigHashAll",
			Label:       "Signature Not Sig Hash All",
			Description: "The transaction authorization signature doesn't sign the entire transaction so doesn't necessarily sign the action requested.",
			MetaData:    `{}`,
		},

		RejectionsAgreementExists: &RejectionsCode{
			Name:        "AgreementExists",
			Label:       "Agreement Already Exists",
			Description: "The agreement already exists and can't be recreated.",
			MetaData:    `{}`,
		},

		RejectionsAgreementDoesNotExist: &RejectionsCode{
			Name:        "AgreementDoesNotExist",
			Label:       "Agreement Does Not Exist",
			Description: "The agreement address specified does not have an active agreement.",
			MetaData:    `{}`,
		},

		RejectionsAgreementRevision: &RejectionsCode{
			Name:        "AgreementRevision",
			Label:       "Agreement Revision",
			Description: "The agreement revision in the request is not current.",
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
