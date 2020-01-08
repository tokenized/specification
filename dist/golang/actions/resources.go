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
	Name string
	Label string
	Description string
	MetaData string
}

// EntitiesData holds a mapping of Entities codes.
func EntitiesData(code string) *EntitiesCode {
	switch(code) {
	case EntitiesUnspecified:
		return &EntitiesCode{
			Name: "Unspecified",
			Label: "Unspecified",
			Description: "",
			MetaData: `{}`,
		}
	case EntitiesIndividual:
		return &EntitiesCode{
			Name: "Individual",
			Label: "Individual",
			Description: "",
			MetaData: `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"legalGuardian":[],"principal":null},"associates":{"accountant":[],"advisor":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"friend":[],"governmentAgency":[],"suppliers":[]},"managers":null,"members":{"principal":null}},"type":"Legal"}`,
		}
	case EntitiesPublicCompany:
		return &EntitiesCode{
			Name: "PublicCompany",
			Label: "Public Company Limited by Shares",
			Description: "",
			MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		}
	case EntitiesPrivateCompany:
		return &EntitiesCode{
			Name: "PrivateCompany",
			Label: "Private Company Limited by Shares",
			Description: "",
			MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
		}
	case EntitiesLimitedPartnership:
		return &EntitiesCode{
			Name: "LimitedPartnership",
			Label: "Limited Partnership",
			Description: "",
			MetaData: `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		}
	case EntitiesUnlimitedPartnership:
		return &EntitiesCode{
			Name: "UnlimitedPartnership",
			Label: "Unlimited Partnership",
			Description: "",
			MetaData: `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"contractor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
		}
	case EntitiesSoleProprietorship:
		return &EntitiesCode{
			Name: "SoleProprietorship",
			Label: "Sole Proprietorship",
			Description: "",
			MetaData: `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"proprietor":null},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"proprietor":null}},"type":"Ownership"}`,
		}
	case EntitiesStatutoryCompany:
		return &EntitiesCode{
			Name: "StatutoryCompany",
			Label: "Statutory Company",
			Description: "",
			MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"nationState":null}},"type":"Legal"}`,
		}
	case EntitiesNonProfitOrganization:
		return &EntitiesCode{
			Name: "NonProfitOrganization",
			Label: "Non-Profit Organization",
			Description: "",
			MetaData: `{"constitutionalDocument":{"organizationConstitution":null},"role":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null}},"type":"Legal"}`,
		}
	case EntitiesNationState:
		return &EntitiesCode{
			Name: "NationState",
			Label: "Nation State",
			Description: "",
			MetaData: `{"constitutionalDocument":{"nationalConstitution":null},"role":{"administrators":null,"associates":null,"collaborators":null,"managers":null,"members":{"citizen":[]}},"type":"Legal"}`,
		}
	case EntitiesGovernmentAgency:
		return &EntitiesCode{
			Name: "GovernmentAgency",
			Label: "Government Agency",
			Description: "",
			MetaData: `{"constitutionalDocument":{"charter":null},"role":{"administrators":null,"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":null},"type":"Legal"}`,
		}
	case EntitiesUnitTrust:
		return &EntitiesCode{
			Name: "UnitTrust",
			Label: "Unit Trust",
			Description: "",
			MetaData: `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"unitholder":[]}},"type":"Ownership"}`,
		}
	case EntitiesDiscretionaryTrust:
		return &EntitiesCode{
			Name: "DiscretionaryTrust",
			Label: "Discretionary Trust",
			Description: "",
			MetaData: `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"beneficiary":[]}},"type":"Ownership"}`,
		}
	default:
		return nil
	}
}

// EntitiesMap returns a mapping of Entities objects with the code as the key.
func EntitiesMap() map[string]*EntitiesCode {
	return map[string]*EntitiesCode {
		EntitiesUnspecified :
			&EntitiesCode{
				Name: "Unspecified",
				Label: "Unspecified",
				Description: "",
				MetaData: `{}`,
			},
		EntitiesIndividual :
			&EntitiesCode{
				Name: "Individual",
				Label: "Individual",
				Description: "",
				MetaData: `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"legalGuardian":[],"principal":null},"associates":{"accountant":[],"advisor":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"friend":[],"governmentAgency":[],"suppliers":[]},"managers":null,"members":{"principal":null}},"type":"Legal"}`,
			},
		EntitiesPublicCompany :
			&EntitiesCode{
				Name: "PublicCompany",
				Label: "Public Company Limited by Shares",
				Description: "",
				MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
			},
		EntitiesPrivateCompany :
			&EntitiesCode{
				Name: "PrivateCompany",
				Label: "Private Company Limited by Shares",
				Description: "",
				MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"shareholder":[],"significantShareholder":[]}},"type":"Legal"}`,
			},
		EntitiesLimitedPartnership :
			&EntitiesCode{
				Name: "LimitedPartnership",
				Label: "Limited Partnership",
				Description: "",
				MetaData: `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
			},
		EntitiesUnlimitedPartnership :
			&EntitiesCode{
				Name: "UnlimitedPartnership",
				Label: "Unlimited Partnership",
				Description: "",
				MetaData: `{"constitutionalDocument":{"partnershipAgreement":null},"roles":{"administrators":{"partner":[]},"associates":{"accountant":[],"advisor":[],"contractor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"managingPartner":[]},"members":{"partner":[]}},"type":"Ownership"}`,
			},
		EntitiesSoleProprietorship :
			&EntitiesCode{
				Name: "SoleProprietorship",
				Label: "Sole Proprietorship",
				Description: "",
				MetaData: `{"constitutionalDocument":null,"roles":{"administrators":{"agent":[],"proprietor":null},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"proprietor":null}},"type":"Ownership"}`,
			},
		EntitiesStatutoryCompany :
			&EntitiesCode{
				Name: "StatutoryCompany",
				Label: "Statutory Company",
				Description: "",
				MetaData: `{"constitutionalDocument":{"companyConstitution":null},"roles":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null},"members":{"nationState":null}},"type":"Legal"}`,
			},
		EntitiesNonProfitOrganization :
			&EntitiesCode{
				Name: "NonProfitOrganization",
				Label: "Non-Profit Organization",
				Description: "",
				MetaData: `{"constitutionalDocument":{"organizationConstitution":null},"role":{"administrators":{"chairman":null,"director":[]},"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"managers":{"ceo":null,"cfo":null,"coo":null,"cto":null,"executive":[],"secretary":null}},"type":"Legal"}`,
			},
		EntitiesNationState :
			&EntitiesCode{
				Name: "NationState",
				Label: "Nation State",
				Description: "",
				MetaData: `{"constitutionalDocument":{"nationalConstitution":null},"role":{"administrators":null,"associates":null,"collaborators":null,"managers":null,"members":{"citizen":[]}},"type":"Legal"}`,
			},
		EntitiesGovernmentAgency :
			&EntitiesCode{
				Name: "GovernmentAgency",
				Label: "Government Agency",
				Description: "",
				MetaData: `{"constitutionalDocument":{"charter":null},"role":{"administrators":null,"associates":{"accountant":[],"advisor":[],"employee":[],"lawyer":[],"manager":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":null},"type":"Legal"}`,
			},
		EntitiesUnitTrust :
			&EntitiesCode{
				Name: "UnitTrust",
				Label: "Unit Trust",
				Description: "",
				MetaData: `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"unitholder":[]}},"type":"Ownership"}`,
			},
		EntitiesDiscretionaryTrust :
			&EntitiesCode{
				Name: "DiscretionaryTrust",
				Label: "Discretionary Trust",
				Description: "",
				MetaData: `{"constitutionalDocument":{"trustDeed":null},"roles":{"administrators":{"protector":[],"trustee":[]},"associates":{"accountant":[],"advisor":[],"custodian":[],"employee":[],"lawyer":[],"manager":[],"settlor":[],"trader":[]},"collaborators":{"customer":[],"governmentAgency":[],"supplier":[]},"members":{"beneficiary":[]}},"type":"Ownership"}`,
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
	PolitiesCuracao(NetherlandsAntilles) = "CUW"

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
	PolitiesCoteD'Ivoire(FormerIvoryCoast) = "CIV"

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
	Name string
	Label string
	Description string
	MetaData string
}

// PolitiesData holds a mapping of Polities codes.
func PolitiesData(code string) *PolitiesCode {
	switch(code) {
	case PolitiesAalandIslands:
		return &PolitiesCode{
			Name: "Aaland Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/52/Flag_of_%C3%85land.svg","phone_code":"+358","states":null}`,
		}
	case PolitiesAfghanistan:
		return &PolitiesCode{
			Name: "Afghanistan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Afghanistan.svg","phone_code":"+93","states":null}`,
		}
	case PolitiesAlbania:
		return &PolitiesCode{
			Name: "Albania",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/36/Flag_of_Albania.svg","phone_code":"+355","states":null}`,
		}
	case PolitiesAlgeria:
		return &PolitiesCode{
			Name: "Algeria",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg","phone_code":"+213","states":null}`,
		}
	case PolitiesAmericanSamoa:
		return &PolitiesCode{
			Name: "American Samoa",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/87/Flag_of_American_Samoa.svg","phone_code":"+1684","states":null}`,
		}
	case PolitiesAndorra:
		return &PolitiesCode{
			Name: "Andorra",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Andorra.svg","phone_code":"+376","states":null}`,
		}
	case PolitiesAngola:
		return &PolitiesCode{
			Name: "Angola",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9d/Flag_of_Angola.svg","phone_code":"+244","states":null}`,
		}
	case PolitiesAnguilla:
		return &PolitiesCode{
			Name: "Anguilla",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Anguilla.svg","phone_code":"+1264","states":null}`,
		}
	case PolitiesAntarctica:
		return &PolitiesCode{
			Name: "Antarctica",
			Label: "",
			Description: "",
			MetaData: `{"flag":null,"phone_code":"+672","states":null}`,
		}
	case PolitiesAntiguaandBarbuda:
		return &PolitiesCode{
			Name: "Antigua and Barbuda",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Antigua_and_Barbuda.svg","phone_code":"+1268","states":null}`,
		}
	case PolitiesArgentina:
		return &PolitiesCode{
			Name: "Argentina",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1a/Flag_of_Argentina.svg","phone_code":"+54","states":null}`,
		}
	case PolitiesArmenia:
		return &PolitiesCode{
			Name: "Armenia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2f/Flag_of_Armenia.svg","phone_code":"+374","states":null}`,
		}
	case PolitiesAruba:
		return &PolitiesCode{
			Name: "Aruba",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Aruba.svg","phone_code":"+297","states":null}`,
		}
	case PolitiesAfricanUnion:
		return &PolitiesCode{
			Name: "African Union",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_the_African_Union.svg","states":{"AGO":null,"BDI":null,"BEN":null,"BFA":null,"BWA":null,"CAF":null,"CIV":null,"CMR":null,"COD":null,"COM":null,"CPV":null,"DJI":null,"DZA":null,"EGY":null,"ERI":null,"ETH":null,"GAB":null,"GHA":null,"GIN":null,"GMB":null,"GNB":null,"GNQ":null,"KEN":null,"LBR":null,"LBY":null,"LSO":null,"SWZ":null,"TCD":null}}`,
		}
	case PolitiesAustralia:
		return &PolitiesCode{
			Name: "Australia",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/b/b9/Flag_of_Australia.svg","gov_fiscal_year":"0701-0630","phone_code":"+61","states":{"AUACT":"Australian Capital Territory","AUCC":"Cocos (Keening) Island","AUCX":"Christmas Island","AUHM":"Heard Island and McDonalds Islands","AUJBT":"Jervis Bay Territory","AUNF":"Norfolk Island","AUNSW":"New South Wales","AUNT":"Northern Territory","AUQLD":"Queensland","AUSA":"South Australia","AUTAS":"Tasmania","AUVIC":"Victoria","AUWA":"Western Australia"}}`,
		}
	case PolitiesAustria:
		return &PolitiesCode{
			Name: "Austria",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/41/Flag_of_Austria.svg","gov_fiscal_year":"0101-1231","phone_code":"+43","states":null}`,
		}
	case PolitiesAzerbaijan:
		return &PolitiesCode{
			Name: "Azerbaijan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg","phone_code":"+994","states":null}`,
		}
	case PolitiesTheBahamas:
		return &PolitiesCode{
			Name: "The Bahamas",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/93/Flag_of_the_Bahamas.svg","phone_code":"+1242","states":null}`,
		}
	case PolitiesBahrain:
		return &PolitiesCode{
			Name: "Bahrain",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Bahrain.svg","phone_code":"+973","states":null}`,
		}
	case PolitiesBangladesh:
		return &PolitiesCode{
			Name: "Bangladesh",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/f9/Flag_of_Bangladesh.svg","gov_fiscal_year":"0701-0630","phone_code":"+880","states":null}`,
		}
	case PolitiesBarbados:
		return &PolitiesCode{
			Name: "Barbados",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ef/Flag_of_Barbados.svg","phone_code":"+1246","states":null}`,
		}
	case PolitiesBelarus:
		return &PolitiesCode{
			Name: "Belarus",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/85/Flag_of_Belarus.svg","phone_code":"+375","states":null}`,
		}
	case PolitiesBelgium:
		return &PolitiesCode{
			Name: "Belgium",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg","gov_fiscal_year":"0101-1231","phone_code":"+32","states":null}`,
		}
	case PolitiesBelize:
		return &PolitiesCode{
			Name: "Belize",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e7/Flag_of_Belize.svg","phone_code":"+501","states":null}`,
		}
	case PolitiesBenin:
		return &PolitiesCode{
			Name: "Benin",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Benin.svg","phone_code":"+229","states":null}`,
		}
	case PolitiesBermuda:
		return &PolitiesCode{
			Name: "Bermuda",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bermuda.svg","phone_code":"+1441","states":null}`,
		}
	case PolitiesBhutan:
		return &PolitiesCode{
			Name: "Bhutan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/91/Flag_of_Bhutan.svg","phone_code":"+975","states":null}`,
		}
	case PolitiesBolivia:
		return &PolitiesCode{
			Name: "Bolivia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/de/Flag_of_Bolivia_%28state%29.svg","phone_code":"+591","states":null}`,
		}
	case PolitiesBonaireStEustasuisandSaba:
		return &PolitiesCode{
			Name: "Bonaire, St Eustasuis and Saba",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","phone_code":"+599","states":null}`,
		}
	case PolitiesBosniaandHerzegovina:
		return &PolitiesCode{
			Name: "Bosnia and Herzegovina",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bosnia_and_Herzegovina.svg","phone_code":"+387","states":null}`,
		}
	case PolitiesBotswana:
		return &PolitiesCode{
			Name: "Botswana",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Botswana.svg","phone_code":"+267","states":null}`,
		}
	case PolitiesBouvetIsland:
		return &PolitiesCode{
			Name: "Bouvet Island",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		}
	case PolitiesBrazil:
		return &PolitiesCode{
			Name: "Brazil",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg","gov_fiscal_year":"0101-1231","phone_code":"+55","states":{"BRAC":"Acre","BRAL":"Alagoas","BRAM":"Amazonas","BRAP":"Amapá","BRBA":"Bahia","BRCE":"Ceará","BRDF":"Federal District","BRES":"Espírito Santo","BRGO":"Goiás","BRMA":"Maranhão","BRMG":"Minas Gerais","BRMS":"Mato Grosso do Sul","BRMT":"Mato Grosso","BRPA":"Pará","BRPB":"Paraíba","BRPE":"Pernambuco","BRPI":"Piauí","BRPR":"Paraná","BRRJ":"Rio de Janeiro","BRRN":"Rio Grande do Norte","BRRO":"Rondônia","BRRR":"Roraima","BRRS":"Rio Grande do Sul","BRSC":"Santa Catarina","BRSE":"Sergipe","BRSP":"São Paulo","BRTO":"Tocantins"}}`,
		}
	case PolitiesBritishIndianOceanTerritory:
		return &PolitiesCode{
			Name: "British Indian Ocean Territory",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_the_British_Indian_Ocean_Territory.svg","phone_code":"+246","states":null}`,
		}
	case PolitiesBritishVirginIslands:
		return &PolitiesCode{
			Name: "British Virgin Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/42/Flag_of_the_British_Virgin_Islands.svg","phone_code":"+1284","states":null}`,
		}
	case PolitiesBrunei:
		return &PolitiesCode{
			Name: "Brunei",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Brunei.svg","phone_code":"+673","states":null}`,
		}
	case PolitiesBulgaria:
		return &PolitiesCode{
			Name: "Bulgaria",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Bulgaria.svg","phone_code":"+359","states":null}`,
		}
	case PolitiesBurkinaFaso:
		return &PolitiesCode{
			Name: "Burkina Faso",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Burkina_Faso.svg","phone_code":"+226","states":null}`,
		}
	case PolitiesBurundi:
		return &PolitiesCode{
			Name: "Burundi",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/50/Flag_of_Burundi.svg","phone_code":"+257","states":null}`,
		}
	case PolitiesCambodia:
		return &PolitiesCode{
			Name: "Cambodia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_Cambodia.svg","phone_code":"+855","states":null}`,
		}
	case PolitiesCameroon:
		return &PolitiesCode{
			Name: "Cameroon",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Cameroon.svg","phone_code":"+237","states":null}`,
		}
	case PolitiesCanada:
		return &PolitiesCode{
			Name: "Canada",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg","gov_fiscal_year":"0401-0331","phone_code":"+1","states":{"CAAB":"Alberta","CABC":"British Columbia","CAMB":"Manitoba","CANB":"New Brunswick","CANL":"Newfoundland and Labrador","CANS":"Nova Scotia","CANT":"Northwest Territories","CANU":"Nunavut","CAON":"Ontario","CAPE":"Prince Edward Island","CAQC":"Quebec","CASK":"Saskatchewan","CAYT":"Yukon"}}`,
		}
	case PolitiesCapeVerde:
		return &PolitiesCode{
			Name: "Cape Verde",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Cape_Verde.svg","phone_code":"+238","states":null}`,
		}
	case PolitiesCaymanIslands:
		return &PolitiesCode{
			Name: "Cayman Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_the_Cayman_Islands.svg","phone_code":"+1345","states":null}`,
		}
	case PolitiesCentralAfricanRepublic:
		return &PolitiesCode{
			Name: "Central African Republic",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Central_African_Republic.svg","phone_code":"+236","states":null}`,
		}
	case PolitiesChad:
		return &PolitiesCode{
			Name: "Chad",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4b/Flag_of_Chad.svg","phone_code":"+235","states":null}`,
		}
	case PolitiesChile:
		return &PolitiesCode{
			Name: "Chile",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/78/Flag_of_Chile.svg","phone_code":"+56","states":null}`,
		}
	case PolitiesChina:
		return &PolitiesCode{
			Name: "China",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_the_People%27s_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+86","states":{"CNAH":"Anhui","CNBJ":"Beijing","CNCQ":"Chongqing","CNFJ":"Fujian","CNGD":"Guangdong","CNGS":"Gansu","CNGX":"Guangxi","CNGZ":"Guizhou","CNHA":"Henan","CNHB":"Hubei","CNHE":"Hebei","CNHI":"Hainan","CNHK":"Hong Kong (Xianggang)","CNHL":"Heilongjiang","CNHN":"Hunan","CNJL":"Jilin","CNJS":"Jiangsu","CNJX":"Jiangxi","CNLN":"Liaoning","CNMC":"Macao (Aomen)","CNNM":"Nei Mongol (mn)","CNNX":"Ningxia","CNQH":"Qinghai","CNSC":"Sichuan","CNSD":"Shandong","CNSH":"Shanghai","CNSN":"Shaanxi","CNSX":"Shanxi","CNTJ":"Tianjin","CNTW":"Taiwan","CNXJ":"Xinjiang","CNXZ":"Xizang","CNYN":"Yunnan","CNZJ":"Zhejiang"}}`,
		}
	case PolitiesChristmasIsland:
		return &PolitiesCode{
			Name: "Christmas Island",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/67/Flag_of_Christmas_Island.svg","phone_code":"+53","states":null}`,
		}
	case PolitiesCocosIslands:
		return &PolitiesCode{
			Name: "Cocos Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Cocos_%28Keeling%29_Islands.svg","phone_code":"+61","states":null}`,
		}
	case PolitiesColombia:
		return &PolitiesCode{
			Name: "Colombia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Colombia.svg","phone_code":"+57","states":null}`,
		}
	case PolitiesComoros:
		return &PolitiesCode{
			Name: "Comoros",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/94/Flag_of_the_Comoros.svg","phone_code":"+269","states":null}`,
		}
	case PolitiesCongoBrazzaville:
		return &PolitiesCode{
			Name: "Congo-Brazzaville",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_the_Republic_of_the_Congo.svg","phone_code":"+242","states":null}`,
		}
	case PolitiesCongoKinshasa:
		return &PolitiesCode{
			Name: "Congo-Kinshasa",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Democratic_Republic_of_the_Congo.svg","phone_code":"+243","states":null}`,
		}
	case PolitiesCookIslands:
		return &PolitiesCode{
			Name: "Cook Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/35/Flag_of_the_Cook_Islands.svg","phone_code":"+682","states":null}`,
		}
	case PolitiesCostaRica:
		return &PolitiesCode{
			Name: "Costa Rica",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"1001-0931","flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Costa_Rica_%28state%29.svg","gov_fiscal_year":"1001-0931","phone_code":"+506","states":null}`,
		}
	case PolitiesCroatia:
		return &PolitiesCode{
			Name: "Croatia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Croatia.svg","phone_code":"+385","states":null}`,
		}
	case PolitiesCuba:
		return &PolitiesCode{
			Name: "Cuba",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bd/Flag_of_Cuba.svg","phone_code":"+53","states":null}`,
		}
	case PolitiesCuracao(NetherlandsAntilles):
		return &PolitiesCode{
			Name: "Curacao (Netherlands Antilles)",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_Cura%C3%A7ao.svg","phone_code":"+599","states":null}`,
		}
	case PolitiesCyprus:
		return &PolitiesCode{
			Name: "Cyprus",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Cyprus.svg","phone_code":"+357","states":null}`,
		}
	case PolitiesCzechRepublic:
		return &PolitiesCode{
			Name: "Czech Republic",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_Czech_Republic.svg","phone_code":"+420","states":null}`,
		}
	case PolitiesDenmark:
		return &PolitiesCode{
			Name: "Denmark",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Denmark.svg","phone_code":"+45","states":null}`,
		}
	case PolitiesDjibouti:
		return &PolitiesCode{
			Name: "Djibouti",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_Djibouti.svg","phone_code":"+253","states":null}`,
		}
	case PolitiesDominica:
		return &PolitiesCode{
			Name: "Dominica",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c4/Flag_of_Dominica.svg","phone_code":"+1767","states":null}`,
		}
	case PolitiesDominicanRepublic:
		return &PolitiesCode{
			Name: "Dominican Republic",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_the_Dominican_Republic.svg","phone_code":"+1829","states":null}`,
		}
	case PolitiesEastTimor:
		return &PolitiesCode{
			Name: "East Timor",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/26/Flag_of_East_Timor.svg","phone_code":"+670","states":null}`,
		}
	case PolitiesEcuador:
		return &PolitiesCode{
			Name: "Ecuador",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e8/Flag_of_Ecuador.svg","phone_code":"+593","states":null}`,
		}
	case PolitiesEgypt:
		return &PolitiesCode{
			Name: "Egypt",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Egypt.svg","gov_fiscal_year":"0701-0630","phone_code":"+20","states":null}`,
		}
	case PolitiesElSalvador:
		return &PolitiesCode{
			Name: "El Salvador",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_El_Salvador.svg","phone_code":"+503","states":null}`,
		}
	case PolitiesEquatorialGuinea:
		return &PolitiesCode{
			Name: "Equatorial Guinea",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Equatorial_Guinea.svg","phone_code":"+240","states":null}`,
		}
	case PolitiesEritrea:
		return &PolitiesCode{
			Name: "Eritrea",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/29/Flag_of_Eritrea.svg","phone_code":"+291","states":null}`,
		}
	case PolitiesEstonia:
		return &PolitiesCode{
			Name: "Estonia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8f/Flag_of_Estonia.svg","phone_code":"+372","states":null}`,
		}
	case PolitiesEthiopia:
		return &PolitiesCode{
			Name: "Ethiopia",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0708-0707","flag":"https://upload.wikimedia.org/wikipedia/commons/7/71/Flag_of_Ethiopia.svg","gov_fiscal_year":"0708-0707","phone_code":"+251","states":null}`,
		}
	case PolitiesEuropeanUnion:
		return &PolitiesCode{
			Name: "European Union",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b7/Flag_of_Europe.svg","states":{"AUT":null,"BEL":null,"BGR":null,"CYP":null,"CZE":null,"DEU":null,"DNK":null,"ESP":null,"EST":null,"FIN":null,"FRA":null,"GBR":null,"GRC":null,"HRV":null,"HUN":null,"IRL":null,"ITA":null,"LTU":null,"LUX":null,"LVA":null,"MLT":null,"NLD":null,"POL":null,"PRT":null,"ROU":null,"SVK":null,"SVN":null,"SWE":null}}`,
		}
	case PolitiesFalklandIslands:
		return &PolitiesCode{
			Name: "Falkland Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_the_Falkland_Islands.svg","phone_code":"+500","states":null}`,
		}
	case PolitiesFaroeIslands:
		return &PolitiesCode{
			Name: "Faroe Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3c/Flag_of_the_Faroe_Islands.svg","phone_code":"+298","states":null}`,
		}
	case PolitiesFiji:
		return &PolitiesCode{
			Name: "Fiji",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/ba/Flag_of_Fiji.svg","phone_code":"+679","states":null}`,
		}
	case PolitiesFinland:
		return &PolitiesCode{
			Name: "Finland",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg","phone_code":"+358","states":null}`,
		}
	case PolitiesFrance:
		return &PolitiesCode{
			Name: "France",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+33","states":null}`,
		}
	case PolitiesFrenchGuiana:
		return &PolitiesCode{
			Name: "French Guiana",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+594","states":null}`,
		}
	case PolitiesFrenchPolynesia:
		return &PolitiesCode{
			Name: "French Polynesia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/db/Flag_of_French_Polynesia.svg","phone_code":"+689","states":null}`,
		}
	case PolitiesFrenchSouthernandAntarcticLands:
		return &PolitiesCode{
			Name: "French Southern and Antarctic Lands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a7/Flag_of_the_French_Southern_and_Antarctic_Lands.svg","phone_code":"+262","states":null}`,
		}
	case PolitiesGabon:
		return &PolitiesCode{
			Name: "Gabon",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/04/Flag_of_Gabon.svg","phone_code":"+241","states":null}`,
		}
	case PolitiesTheGambia:
		return &PolitiesCode{
			Name: "The Gambia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_The_Gambia.svg","phone_code":"+220","states":null}`,
		}
	case PolitiesGeorgia:
		return &PolitiesCode{
			Name: "Georgia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Georgia.svg","phone_code":"+995","states":null}`,
		}
	case PolitiesGermany:
		return &PolitiesCode{
			Name: "Germany",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg","gov_fiscal_year":"0101-1231","phone_code":"+49","states":null}`,
		}
	case PolitiesGhana:
		return &PolitiesCode{
			Name: "Ghana",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Ghana.svg","phone_code":"+233","states":null}`,
		}
	case PolitiesGibraltar:
		return &PolitiesCode{
			Name: "Gibraltar",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/02/Flag_of_Gibraltar.svg","phone_code":"+350","states":null}`,
		}
	case PolitiesGreece:
		return &PolitiesCode{
			Name: "Greece",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Greece.svg","gov_fiscal_year":"0101-1231","phone_code":"+30","states":null}`,
		}
	case PolitiesGreenland:
		return &PolitiesCode{
			Name: "Greenland",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_Greenland.svg","phone_code":"+299","states":null}`,
		}
	case PolitiesGrenada:
		return &PolitiesCode{
			Name: "Grenada",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Grenada.svg","phone_code":"+1473","states":null}`,
		}
	case PolitiesGuadeloupe:
		return &PolitiesCode{
			Name: "Guadeloupe",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+590","states":null}`,
		}
	case PolitiesGuam:
		return &PolitiesCode{
			Name: "Guam",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/07/Flag_of_Guam.svg","phone_code":"+1671","states":null}`,
		}
	case PolitiesGuatemala:
		return &PolitiesCode{
			Name: "Guatemala",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ec/Flag_of_Guatemala.svg","phone_code":"+502","states":null}`,
		}
	case PolitiesGuernsey:
		return &PolitiesCode{
			Name: "Guernsey",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Guernsey.svg","phone_code":"+44","states":null}`,
		}
	case PolitiesGuinea:
		return &PolitiesCode{
			Name: "Guinea",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_Guinea.svg","phone_code":"+224","states":null}`,
		}
	case PolitiesGuineaBissau:
		return &PolitiesCode{
			Name: "Guinea-Bissau",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Guinea-Bissau.svg","phone_code":"+245","states":null}`,
		}
	case PolitiesGuyana:
		return &PolitiesCode{
			Name: "Guyana",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_Guyana.svg","phone_code":"+592","states":null}`,
		}
	case PolitiesHaiti:
		return &PolitiesCode{
			Name: "Haiti",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Haiti.svg","phone_code":"+509","states":null}`,
		}
	case PolitiesHeardIslandandMcDonaldIslands:
		return &PolitiesCode{
			Name: "Heard Island and McDonald Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg","phone_code":"+672","states":null}`,
		}
	case PolitiesHonduras:
		return &PolitiesCode{
			Name: "Honduras",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/82/Flag_of_Honduras.svg","phone_code":"+504","states":null}`,
		}
	case PolitiesHongKong:
		return &PolitiesCode{
			Name: "Hong Kong",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg","gov_fiscal_year":"0401-0331","phone_code":"+852","states":null}`,
		}
	case PolitiesHungary:
		return &PolitiesCode{
			Name: "Hungary",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c1/Flag_of_Hungary.svg","phone_code":"+36","states":null}`,
		}
	case PolitiesIceland:
		return &PolitiesCode{
			Name: "Iceland",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Iceland.svg","phone_code":"+354","states":null}`,
		}
	case PolitiesIndia:
		return &PolitiesCode{
			Name: "India",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg","gov_fiscal_year":"0401-0331","phone_code":"+91","states":{"INAN":"Andaman and Nicobar Islands","INAP":"Andhra Pradesh","INAR":"Arunachal Pradesh","INAS":"Assam","INBR":"Bihar","INCH":"Chandigarh","INCT":"Chhattisgarh","INDD":"Daman and Diu","INDL":"Delhi","INDN":"Dadra and Nagar Haveli","INGA":"Goa","INGJ":"Gujarat","INHP":"Himachal Pradesh","INHR":"Haryana","INJH":"Jharkhand","INJK":"Jammu and Kashmir","INKA":"Karnataka","INKL":"Kerala","INLD":"Lakshadweep","INMH":"Maharashtra","INML":"Meghalaya","INMN":"Manipur","INMP":"Madhya Pradesh","INMZ":"Mizoram","INNL":"Nagaland","INOR":"Odisha (formerly known as Orissa)","INPB":"Punjab","INPY":"Puducherry (Pondicherry)","INRJ":"Rajasthan","INSK":"Sikkim","INTG":"Telangana","INTN":"Tamil Nadu","INTR":"Tripura","INUP":"Uttar Pradesh","INUT":"Uttarakhand","INWB":"West Bengal"}}`,
		}
	case PolitiesIndonesia:
		return &PolitiesCode{
			Name: "Indonesia",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg","gov_fiscal_year":"0101-1231","phone_code":"+62","states":null}`,
		}
	case PolitiesIran:
		return &PolitiesCode{
			Name: "Iran",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ca/Flag_of_Iran.svg","phone_code":"+98","states":null}`,
		}
	case PolitiesIraq:
		return &PolitiesCode{
			Name: "Iraq",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Iraq.svg","phone_code":"+964","states":null}`,
		}
	case PolitiesIreland:
		return &PolitiesCode{
			Name: "Ireland",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/45/Flag_of_Ireland.svg","gov_fiscal_year":"0101-1231","phone_code":"+353","states":null}`,
		}
	case PolitiesIsleofMan:
		return &PolitiesCode{
			Name: "Isle of Man",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5d/Flag_of_the_Isle_of_Mann.svg","phone_code":"+44","states":null}`,
		}
	case PolitiesIsrael:
		return &PolitiesCode{
			Name: "Israel",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Israel.svg","states":null}`,
		}
	case PolitiesItaly:
		return &PolitiesCode{
			Name: "Italy",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/0/03/Flag_of_Italy.svg","phone_code":"+39","states":null}`,
		}
	case PolitiesCoteD'Ivoire(FormerIvoryCoast):
		return &PolitiesCode{
			Name: "Cote D'Ivoire (Former Ivory Coast)",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_C%C3%B4te_d%27Ivoire.svg","phone_code":"+225","states":null}`,
		}
	case PolitiesJamaica:
		return &PolitiesCode{
			Name: "Jamaica",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Jamaica.svg","phone_code":"+1876","states":null}`,
		}
	case PolitiesJapan:
		return &PolitiesCode{
			Name: "Japan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg","states":null}`,
		}
	case PolitiesJersey:
		return &PolitiesCode{
			Name: "Jersey",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1c/Flag_of_Jersey.svg","phone_code":"+44","states":null}`,
		}
	case PolitiesJordan:
		return &PolitiesCode{
			Name: "Jordan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c0/Flag_of_Jordan.svg","phone_code":"+962","states":null}`,
		}
	case PolitiesKazakhstan:
		return &PolitiesCode{
			Name: "Kazakhstan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kazakhstan.svg","phone_code":"+7","states":null}`,
		}
	case PolitiesKenya:
		return &PolitiesCode{
			Name: "Kenya",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Kenya.svg","phone_code":"+254","states":null}`,
		}
	case PolitiesKiribati:
		return &PolitiesCode{
			Name: "Kiribati",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kiribati.svg","phone_code":"+686","states":null}`,
		}
	case PolitiesKuwait:
		return &PolitiesCode{
			Name: "Kuwait",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/aa/Flag_of_Kuwait.svg","phone_code":"+965","states":null}`,
		}
	case PolitiesKyrgyzstan:
		return &PolitiesCode{
			Name: "Kyrgyzstan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c7/Flag_of_Kyrgyzstan.svg","phone_code":"+996","states":null}`,
		}
	case PolitiesLaos:
		return &PolitiesCode{
			Name: "Laos",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Laos.svg","phone_code":"+856","states":null}`,
		}
	case PolitiesLatvia:
		return &PolitiesCode{
			Name: "Latvia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Latvia.svg","phone_code":"+371","states":null}`,
		}
	case PolitiesLebanon:
		return &PolitiesCode{
			Name: "Lebanon",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/59/Flag_of_Lebanon.svg","phone_code":"+961","states":null}`,
		}
	case PolitiesLesotho:
		return &PolitiesCode{
			Name: "Lesotho",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Lesotho.svg","phone_code":"+266","states":null}`,
		}
	case PolitiesLiberia:
		return &PolitiesCode{
			Name: "Liberia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b8/Flag_of_Liberia.svg","phone_code":"+231","states":null}`,
		}
	case PolitiesLibya:
		return &PolitiesCode{
			Name: "Libya",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/05/Flag_of_Libya.svg","phone_code":"+218","states":null}`,
		}
	case PolitiesLiechtenstein:
		return &PolitiesCode{
			Name: "Liechtenstein",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/47/Flag_of_Liechtenstein.svg","phone_code":"+423","states":null}`,
		}
	case PolitiesLithuania:
		return &PolitiesCode{
			Name: "Lithuania",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Lithuania.svg","phone_code":"+370","states":null}`,
		}
	case PolitiesLuxembourg:
		return &PolitiesCode{
			Name: "Luxembourg",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/da/Flag_of_Luxembourg.svg","phone_code":"+352","states":null}`,
		}
	case PolitiesMacau:
		return &PolitiesCode{
			Name: "Macau",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/63/Flag_of_Macau.svg","phone_code":"+853","states":null}`,
		}
	case PolitiesMacedonia:
		return &PolitiesCode{
			Name: "Macedonia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_Macedonia.svg","phone_code":"+389","states":null}`,
		}
	case PolitiesMadagascar:
		return &PolitiesCode{
			Name: "Madagascar",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Madagascar.svg","phone_code":"+261","states":null}`,
		}
	case PolitiesMalawi:
		return &PolitiesCode{
			Name: "Malawi",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d1/Flag_of_Malawi.svg","phone_code":"+265","states":null}`,
		}
	case PolitiesMalaysia:
		return &PolitiesCode{
			Name: "Malaysia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/66/Flag_of_Malaysia.svg","phone_code":"+60","states":null}`,
		}
	case PolitiesMaldives:
		return &PolitiesCode{
			Name: "Maldives",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Maldives.svg","phone_code":"+960","states":null}`,
		}
	case PolitiesMali:
		return &PolitiesCode{
			Name: "Mali",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Mali.svg","phone_code":"+223","states":null}`,
		}
	case PolitiesMalta:
		return &PolitiesCode{
			Name: "Malta",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Malta.svg","phone_code":"+356","states":null}`,
		}
	case PolitiesMarshallIslands:
		return &PolitiesCode{
			Name: "Marshall Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2e/Flag_of_the_Marshall_Islands.svg","phone_code":"+692","states":null}`,
		}
	case PolitiesMartinique:
		return &PolitiesCode{
			Name: "Martinique",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+596","states":null}`,
		}
	case PolitiesMauritania:
		return &PolitiesCode{
			Name: "Mauritania",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/New_National_Flag_of_Mauritania.png","phone_code":"+222","states":null}`,
		}
	case PolitiesMauritius:
		return &PolitiesCode{
			Name: "Mauritius",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Mauritius.svg","phone_code":"+230","states":null}`,
		}
	case PolitiesMayotte:
		return &PolitiesCode{
			Name: "Mayotte",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Mayotte_%28local%29.svg","phone_code":"+269","states":null}`,
		}
	case PolitiesMexico:
		return &PolitiesCode{
			Name: "Mexico",
			Label: "",
			Description: "",
			MetaData: `{"MXAGU":"Aguascalientes","MXBCN":"Baja California","MXBCS":"Baja California Sur","MXCAM":"Campeche","MXCHH":"Chihuahua","MXCHP":"Chiapas","MXCOA":"Coahuila","MXCOL":"Colima","MXDIF":"Distrito Federal","MXDUR":"Durango","MXGRO":"Guerrero","MXGUA":"Guanajuato","MXHID":"Hidalgo","MXJAL":"Jalisco","MXMEX":"Mexico (Federal District)","MXMIC":"Michoacán","MXMOR":"Morelos","MXNAY":"Nayarit","MXNLE":"Nuevo León","MXOAX":"Oaxaca","MXPUE":"Puebla","MXQUE":"Querétaro","MXROO":"Quintana Roo","MXSIN":"Sinaloa","MXSLP":"San Luis Potosí","MXSON":"Sonora","MXTAB":"Tabasco","MXTAM":"Tamaulipas","MXTLA":"Tlaxcala","MXVER":"Veracruz","MXYUC":"Yucatán","MXZAC":"Zacatecas","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Mexico.svg","phone_code":"+52","states":null}`,
		}
	case PolitiesMicronesia:
		return &PolitiesCode{
			Name: "Micronesia",
			Label: "",
			Description: "",
			MetaData: `{"flag":null,"phone_code":"+691","states":null}`,
		}
	case PolitiesMoldova:
		return &PolitiesCode{
			Name: "Moldova",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Moldova.svg","phone_code":"+373","states":null}`,
		}
	case PolitiesMonaco:
		return &PolitiesCode{
			Name: "Monaco",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ea/Flag_of_Monaco.svg","phone_code":"+377","states":null}`,
		}
	case PolitiesMongolia:
		return &PolitiesCode{
			Name: "Mongolia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4c/Flag_of_Mongolia.svg","phone_code":"+976","states":null}`,
		}
	case PolitiesMontenegro:
		return &PolitiesCode{
			Name: "Montenegro",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Montenegro.svg","phone_code":"+382","states":null}`,
		}
	case PolitiesMontserrat:
		return &PolitiesCode{
			Name: "Montserrat",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Montserrat.svg","phone_code":"+1664","states":null}`,
		}
	case PolitiesMorocco:
		return &PolitiesCode{
			Name: "Morocco",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Morocco.svg","phone_code":"+212","states":null}`,
		}
	case PolitiesMozambique:
		return &PolitiesCode{
			Name: "Mozambique",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Mozambique.svg","phone_code":"+258","states":null}`,
		}
	case PolitiesMyanmar:
		return &PolitiesCode{
			Name: "Myanmar",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8c/Flag_of_Myanmar.svg","phone_code":"+95","states":null}`,
		}
	case PolitiesNamibia:
		return &PolitiesCode{
			Name: "Namibia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Namibia.svg","phone_code":"+264","states":null}`,
		}
	case PolitiesNauru:
		return &PolitiesCode{
			Name: "Nauru",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/Flag_of_Nauru.svg","phone_code":"+674","states":null}`,
		}
	case PolitiesNepal:
		return &PolitiesCode{
			Name: "Nepal",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9b/Flag_of_Nepal.svg","gov_fiscal_year":"0101-1231","phone_code":"+977","states":null}`,
		}
	case PolitiesNetherlands:
		return &PolitiesCode{
			Name: "Netherlands",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","gov_fiscal_year":"0101-1231","phone_code":"+31","states":null}`,
		}
	case PolitiesNewCaledonia:
		return &PolitiesCode{
			Name: "New Caledonia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+687","states":null}`,
		}
	case PolitiesNewZealand:
		return &PolitiesCode{
			Name: "New Zealand",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3e/Flag_of_New_Zealand.svg","phone_code":"+64","states":null}`,
		}
	case PolitiesNicaragua:
		return &PolitiesCode{
			Name: "Nicaragua",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Nicaragua.svg","phone_code":"+505","states":null}`,
		}
	case PolitiesNiger:
		return &PolitiesCode{
			Name: "Niger",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f4/Flag_of_Niger.svg","phone_code":"+227","states":null}`,
		}
	case PolitiesNigeria:
		return &PolitiesCode{
			Name: "Nigeria",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/79/Flag_of_Nigeria.svg","phone_code":"+234","states":null}`,
		}
	case PolitiesNiue:
		return &PolitiesCode{
			Name: "Niue",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Niue.svg","phone_code":"+683","states":null}`,
		}
	case PolitiesNorfolkandPhilipIsland:
		return &PolitiesCode{
			Name: "Norfolk and Philip Island",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Norfolk_Island.svg","phone_code":"+672","states":null}`,
		}
	case PolitiesNorthKorea:
		return &PolitiesCode{
			Name: "North Korea",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_North_Korea.svg","phone_code":"+850","states":null}`,
		}
	case PolitiesNorthernMarianaIslands:
		return &PolitiesCode{
			Name: "Northern Mariana Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e0/Flag_of_the_Northern_Mariana_Islands.svg","phone_code":"+1670","states":null}`,
		}
	case PolitiesNorway:
		return &PolitiesCode{
			Name: "Norway",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
		}
	case PolitiesOman:
		return &PolitiesCode{
			Name: "Oman",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Oman.svg","phone_code":"+968","states":null}`,
		}
	case PolitiesPakistan:
		return &PolitiesCode{
			Name: "Pakistan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/32/Flag_of_Pakistan.svg","phone_code":"+92","states":null}`,
		}
	case PolitiesPalau:
		return &PolitiesCode{
			Name: "Palau",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Palau.svg","phone_code":"+680","states":null}`,
		}
	case PolitiesPalestinianTerritory:
		return &PolitiesCode{
			Name: "Palestinian Territory",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Palestine.svg","phone_code":"+970","states":null}`,
		}
	case PolitiesPanama:
		return &PolitiesCode{
			Name: "Panama",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/ab/Flag_of_Panama.svg","phone_code":"+507","states":null}`,
		}
	case PolitiesPapuaNewGuinea:
		return &PolitiesCode{
			Name: "Papua New Guinea",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e3/Flag_of_Papua_New_Guinea.svg","phone_code":"+675","states":null}`,
		}
	case PolitiesParaguay:
		return &PolitiesCode{
			Name: "Paraguay",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Paraguay.svg","phone_code":"+595","states":null}`,
		}
	case PolitiesPeru:
		return &PolitiesCode{
			Name: "Peru",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Peru_%28state%29.svg","phone_code":"+51","states":null}`,
		}
	case PolitiesPhilippines:
		return &PolitiesCode{
			Name: "Philippines",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_the_Philippines.svg","phone_code":"+63","states":null}`,
		}
	case PolitiesPitcairnIslands:
		return &PolitiesCode{
			Name: "Pitcairn Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_the_Pitcairn_Islands.svg","phone_code":"+64","states":null}`,
		}
	case PolitiesPoland:
		return &PolitiesCode{
			Name: "Poland",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg","phone_code":"+48","states":null}`,
		}
	case PolitiesPortugal:
		return &PolitiesCode{
			Name: "Portugal",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Portugal.svg","states":null}`,
		}
	case PolitiesPuertoRico:
		return &PolitiesCode{
			Name: "Puerto Rico",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/28/Flag_of_Puerto_Rico.svg","phone_code":"+1787","states":null}`,
		}
	case PolitiesQatar:
		return &PolitiesCode{
			Name: "Qatar",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Qatar.svg","states":null}`,
		}
	case PolitiesRéunion:
		return &PolitiesCode{
			Name: "Réunion",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5a/Flag_of_R%C3%A9union.svg","phone_code":"+262","states":null}`,
		}
	case PolitiesRomania:
		return &PolitiesCode{
			Name: "Romania",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Romania.svg","states":null}`,
		}
	case PolitiesRussia:
		return &PolitiesCode{
			Name: "Russia",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/f/f3/Flag_of_Russia.svg","gov_fiscal_year":"0101-1231","phone_code":"+7","states":{"RUAD":"Adygeya, Respublika","RUAL":"Altay, Respublika","RUALT":"Altayskiy kray","RUAMU":"Amurskaya oblast\"","RUARK":"Arkhangel\"skaya oblast\"","RUAST":"Astrakhanskaya oblast\"","RUBA":"Bashkortostan, Respublika","RUBEL":"Belgorodskaya oblast\"","RUBRY":"Bryanskaya oblast\"","RUBU":"Buryatiya, Respublika","RUCE":"Chechenskaya Respublika","RUCHE":"Chelyabinskaya oblast\"","RUCHU":"Chukotskiy avtonomnyy okrug","RUCU":"Chuvashskaya Respublika","RUDA":"Dagestan, Respublika","RUIN":"Ingushetiya, Respublika","RUIRK":"Irkutskaya oblast\"","RUIVA":"Ivanovskaya oblast\"","RUKAM":"Kamchatskiy kray","RUKB":"Kabardino-Balkarskaya Respublika","RUKC":"Karachayevo-Cherkesskaya Respubl.","RUKDA":"Krasnodarskiy kray","RUKEM":"Kemerovskaya oblast\"","RUKGD":"Kaliningradskaya oblast\"","RUKGN":"Kurganskaya oblast\"","RUKHA":"Khabarovskiy kray","RUKHM":"Khanty-Mansiyskiy avtonomnyy okrug-Yugra","RUKIR":"Kirovskaya oblast\"","RUKK":"Khakasiya, Respublika","RUKL":"Kalmykiya, Respublika","RUKLU":"Kaluzhskaya oblast\"","RUKO":"Komi, Respublika","RUKOS":"Kostromskaya oblast\"","RUKR":"Kareliya, Respublika","RUKRS":"Kurskaya oblast\"","RUKYA":"Krasnoyarskiy kray","RULEN":"Leningradskaya oblast\"","RULIP":"Lipetskaya oblast\"","RUMAG":"Magadanskaya oblast\"","RUME":"Mariy El, Respublika","RUMO":"Mordoviya, Respublika","RUMOS":"Moskovskaya oblast\"","RUMOW":"Moskva (autonomous city)","RUMUR":"Murmanskaya oblast\"","RUNEN":"Nenetskiy avtonomnyy okrug","RUNGR":"Novgorodskaya oblast\"","RUNIZ":"Nizhegorodskaya oblast\"","RUNVS":"Novosibirskaya oblast\"","RUOMS":"Omskaya oblast\"","RUORE":"Orenburgskaya oblast\"","RUORL":"Orlovskaya oblast\"","RUPER":"Permskiy kray","RUPNZ":"Penzenskaya oblast\"","RUPRI":"Primorskiy kray","RUPSK":"Pskovskaya oblast\"","RUROS":"Rostovskaya oblast\"","RURYA":"Ryazanskaya oblast\"","RUSA":"Sakha, Respublika","RUSAK":"Sakhalinskaya oblast\"","RUSAM":"Samarskaya oblast\"","RUSAR":"Saratovskaya oblast\"","RUSE":"Severnaya Osetiya-Alaniya, Respubl.","RUSMO":"Smolenskaya oblast\"","RUSPE":"Sankt-Peterburg (autonomous city)","RUSTA":"Stavropol\"skiy kray","RUSVE":"Sverdlovskaya oblast\"","RUTA":"Tatarstan, Respublika","RUTAM":"Tambovskaya oblast\"","RUTOM":"Tomskaya oblast\"","RUTUL":"Tul\"skaya oblast\"","RUTVE":"Tverskaya oblast\"","RUTY":"Tyva, Respublika","RUTYU":"Tyumenskaya oblast\"","RUUD":"Udmurtskaya Respublika","RUULY":"Ul\"yanovskaya oblast\"","RUVGG":"Volgogradskaya oblast\"","RUVLA":"Vladimirskaya oblast\"","RUVLG":"Vologodskaya oblast\"","RUVOR":"Voronezhskaya oblast\"","RUYAN":"Yamalo-Nenetskiy avtonomnyy okrug","RUYAR":"Yaroslavskaya oblast\"","RUYEV":"Yevreyskaya avtonomnaya oblast\"","RUZAB":"Zabaykal\"skiy kray"}}`,
		}
	case PolitiesRwanda:
		return &PolitiesCode{
			Name: "Rwanda",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Rwanda.svg","phone_code":"+250","states":null}`,
		}
	case PolitiesSaintHelenaAscensionandTristandaCunha:
		return &PolitiesCode{
			Name: "Saint Helena, Ascension and Tristan da Cunha",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","phone_code":"+290","states":null}`,
		}
	case PolitiesSaintKittsandNevis:
		return &PolitiesCode{
			Name: "Saint Kitts and Nevis",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Saint_Kitts_and_Nevis.svg","phone_code":"+1869","states":null}`,
		}
	case PolitiesSaintLucia:
		return &PolitiesCode{
			Name: "Saint Lucia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Saint_Lucia.svg","phone_code":"+1758","states":null}`,
		}
	case PolitiesSaintPierreandMiquelon:
		return &PolitiesCode{
			Name: "Saint Pierre and Miquelon",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+508","states":null}`,
		}
	case PolitiesSaintVincentandtheGrenadines:
		return &PolitiesCode{
			Name: "Saint Vincent and the Grenadines",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6d/Flag_of_Saint_Vincent_and_the_Grenadines.svg","phone_code":"+1784","states":null}`,
		}
	case PolitiesSaintBarthelemy:
		return &PolitiesCode{
			Name: "Saint-Barthelemy",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Saint_Barthelemy_%28local%29.svg","phone_code":"+590","states":null}`,
		}
	case PolitiesSaintMartin:
		return &PolitiesCode{
			Name: "Saint-Martin",
			Label: "",
			Description: "",
			MetaData: `{"flag":null,"phone_code":"+590","states":null}`,
		}
	case PolitiesSamoa:
		return &PolitiesCode{
			Name: "Samoa",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Samoa.svg","phone_code":"+685","states":null}`,
		}
	case PolitiesSanMarino:
		return &PolitiesCode{
			Name: "San Marino",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_San_Marino.svg","phone_code":"+378","states":null}`,
		}
	case PolitiesSãoToméandPríncipe:
		return &PolitiesCode{
			Name: "São Tomé and Príncipe",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Sao_Tome_and_Principe.svg","phone_code":"+239","states":null}`,
		}
	case PolitiesSaudiArabia:
		return &PolitiesCode{
			Name: "Saudi Arabia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0d/Flag_of_Saudi_Arabia.svg","phone_code":"+966","states":null}`,
		}
	case PolitiesSenegal:
		return &PolitiesCode{
			Name: "Senegal",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fd/Flag_of_Senegal.svg","phone_code":"+221","states":null}`,
		}
	case PolitiesSerbia:
		return &PolitiesCode{
			Name: "Serbia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/ff/Flag_of_Serbia.svg","phone_code":"+381","states":null}`,
		}
	case PolitiesSeychelles:
		return &PolitiesCode{
			Name: "Seychelles",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Seychelles.svg","phone_code":"+248","states":null}`,
		}
	case PolitiesSierraLeone:
		return &PolitiesCode{
			Name: "Sierra Leone",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Sierra_Leone.svg","phone_code":"+232","states":null}`,
		}
	case PolitiesSingapore:
		return &PolitiesCode{
			Name: "Singapore",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg","phone_code":"+65","states":null}`,
		}
	case PolitiesSintMaarten:
		return &PolitiesCode{
			Name: "Sint Maarten",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Sint_Maarten.svg","phone_code":"+1721","states":null}`,
		}
	case PolitiesSlovakia:
		return &PolitiesCode{
			Name: "Slovakia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e6/Flag_of_Slovakia.svg","phone_code":"+421","states":null}`,
		}
	case PolitiesSlovenia:
		return &PolitiesCode{
			Name: "Slovenia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f0/Flag_of_Slovenia.svg","phone_code":"+386","states":null}`,
		}
	case PolitiesSolomonIslands:
		return &PolitiesCode{
			Name: "Solomon Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Solomon_Islands.svg","phone_code":"+677","states":null}`,
		}
	case PolitiesSomalia:
		return &PolitiesCode{
			Name: "Somalia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_Somalia.svg","phone_code":"+252","states":null}`,
		}
	case PolitiesSouthAfrica:
		return &PolitiesCode{
			Name: "South Africa",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/af/Flag_of_South_Africa.svg","phone_code":"+27","states":null}`,
		}
	case PolitiesSouthGeorgiaandtheSouthSandwichIslands:
		return &PolitiesCode{
			Name: "South Georgia and the South Sandwich Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_South_Georgia_and_the_South_Sandwich_Islands.svg","phone_code":"+500","states":null}`,
		}
	case PolitiesSouthKorea:
		return &PolitiesCode{
			Name: "South Korea",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg","gov_fiscal_year":"0101-1231","phone_code":"+82","states":null}`,
		}
	case PolitiesSouthSudan:
		return &PolitiesCode{
			Name: "South Sudan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7a/Flag_of_South_Sudan.svg","phone_code":"+211","states":null}`,
		}
	case PolitiesSpain:
		return &PolitiesCode{
			Name: "Spain",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9a/Flag_of_Spain.svg","gov_fiscal_year":"0101-1231","phone_code":"+34","states":null}`,
		}
	case PolitiesSriLanka:
		return &PolitiesCode{
			Name: "Sri Lanka",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Sri_Lanka.svg","phone_code":"+94","states":null}`,
		}
	case PolitiesSudan:
		return &PolitiesCode{
			Name: "Sudan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Sudan.svg","phone_code":"+249","states":null}`,
		}
	case PolitiesSuriname:
		return &PolitiesCode{
			Name: "Suriname",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/60/Flag_of_Suriname.svg","phone_code":"+597","states":null}`,
		}
	case PolitiesSvalbardandJanMayen:
		return &PolitiesCode{
			Name: "Svalbard and Jan Mayen",
			Label: "",
			Description: "",
			MetaData: `{"flag":null,"phone_code":"+47","states":null}`,
		}
	case PolitiesSwaziland:
		return &PolitiesCode{
			Name: "Swaziland",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fb/Flag_of_Eswatini.svg","phone_code":"+268","states":null}`,
		}
	case PolitiesSweden:
		return &PolitiesCode{
			Name: "Sweden",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/4/4c/Flag_of_Sweden.svg","gov_fiscal_year":"0101-1231","phone_code":"+46","states":null}`,
		}
	case PolitiesSwitzerland:
		return &PolitiesCode{
			Name: "Switzerland",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/08/Flag_of_Switzerland_%28Pantone%29.svg","gov_fiscal_year":"0101-1231","phone_code":"+41","states":null}`,
		}
	case PolitiesSyria:
		return &PolitiesCode{
			Name: "Syria",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/53/Flag_of_Syria.svg","phone_code":"+963","states":null}`,
		}
	case PolitiesTaiwan:
		return &PolitiesCode{
			Name: "Taiwan",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+886","states":null}`,
		}
	case PolitiesTajikistan:
		return &PolitiesCode{
			Name: "Tajikistan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Tajikistan.svg","phone_code":"+992","states":null}`,
		}
	case PolitiesTanzania:
		return &PolitiesCode{
			Name: "Tanzania",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tanzania.svg","phone_code":"+255","states":null}`,
		}
	case PolitiesThailand:
		return &PolitiesCode{
			Name: "Thailand",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a9/Flag_of_Thailand.svg","phone_code":"+66","states":null}`,
		}
	case PolitiesTogo:
		return &PolitiesCode{
			Name: "Togo",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/68/Flag_of_Togo.svg","phone_code":"+228","states":null}`,
		}
	case PolitiesTokelau:
		return &PolitiesCode{
			Name: "Tokelau",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8e/Flag_of_Tokelau.svg","phone_code":"+690","states":null}`,
		}
	case PolitiesTonga:
		return &PolitiesCode{
			Name: "Tonga",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Tonga.svg","phone_code":"+676","states":null}`,
		}
	case PolitiesTrinidadandTobago:
		return &PolitiesCode{
			Name: "Trinidad and Tobago",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Trinidad_and_Tobago.svg","phone_code":"+1868","states":null}`,
		}
	case PolitiesTunisia:
		return &PolitiesCode{
			Name: "Tunisia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Tunisia.svg","phone_code":"+216","states":null}`,
		}
	case PolitiesTurkey:
		return &PolitiesCode{
			Name: "Turkey",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Turkey.svg","phone_code":"+90","states":null}`,
		}
	case PolitiesTurkmenistan:
		return &PolitiesCode{
			Name: "Turkmenistan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Turkmenistan.svg","phone_code":"+993","states":null}`,
		}
	case PolitiesTurksandCaicosIslands:
		return &PolitiesCode{
			Name: "Turks and Caicos Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_the_Turks_and_Caicos_Islands.svg","phone_code":"+1649","states":null}`,
		}
	case PolitiesTuvalu:
		return &PolitiesCode{
			Name: "Tuvalu",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tuvalu.svg","phone_code":"+688","states":null}`,
		}
	case PolitiesUganda:
		return &PolitiesCode{
			Name: "Uganda",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4e/Flag_of_Uganda.svg","phone_code":"+256","states":null}`,
		}
	case PolitiesUkraine:
		return &PolitiesCode{
			Name: "Ukraine",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Ukraine.svg","phone_code":"+380","states":null}`,
		}
	case PolitiesUnitedArabEmirates:
		return &PolitiesCode{
			Name: "United Arab Emirates",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_United_Arab_Emirates.svg","states":null}`,
		}
	case PolitiesUnitedKingdom:
		return &PolitiesCode{
			Name: "United Kingdom",
			Label: "",
			Description: "",
			MetaData: `{"fiscal_year":"0406-0405","flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","gov_fiscal_year":"0401-0331","phone_code":"+44","states":null}`,
		}
	case PolitiesUnitedStatesMinorOutlyingIslands:
		return &PolitiesCode{
			Name: "United States Minor Outlying Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+246","states":null}`,
		}
	case PolitiesUruguay:
		return &PolitiesCode{
			Name: "Uruguay",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Uruguay.svg","phone_code":"+598","states":null}`,
		}
	case PolitiesUSVirginIslands:
		return &PolitiesCode{
			Name: "US Virgin Islands",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_the_United_States_Virgin_Islands.svg","phone_code":"+1340","states":null}`,
		}
	case PolitiesUSA:
		return &PolitiesCode{
			Name: "USA",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+1","states":{"USAK":"Alaska","USAL":"Alabama","USAR":"Arkansas","USAS":"American Samoa","USAZ":"Arizona","USCA":"California","USCO":"Colorado","USCT":"Connecticut","USDC":"Washington D.C.","USDE":"Delaware","USFL":"Florida","USGA":"Georgia","USGU":"Guam","USHI":"Hawaii","USIA":"Iowa","USID":"Idaho","USIL":"Illinois","USIN":"Indiana","USKS":"Kansas","USKY":"Kentucky","USLA":"Louisiana","USMA":"Massachusetts","USMD":"Maryland","USME":"Maine","USMI":"Michigan","USMN":"Minnesota","USMO":"Missouri","USMP":"Northern Mariana Islands","USMS":"Mississippi","USMT":"Montana","USNC":"North Carolina","USND":"North Dakota","USNE":"Nebraska","USNH":"New Hampshire","USNJ":"New Jersey","USNM":"New Mexico","USNV":"Nevada","USNY":"New York","USOH":"Ohio","USOK":"Oklahoma","USOR":"Oregon","USPA":"Pennsylvania","USPR":"Puerto Rico","USRI":"Rhode Island","USSC":"South Carolina","USSD":"South Dakota","USTN":"Tennessee","USTX":"Texas","USUM":"United States Minor Outlying Islands","USUT":"Utah","USVA":"Virginia","USVI":"US Virgin Islands","USVT":"Vermont","USWA":"Washington","USWI":"Wisconsin","USWV":"West Virginia","USWY":"Wyoming"}}`,
		}
	case PolitiesUzbekistan:
		return &PolitiesCode{
			Name: "Uzbekistan",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Uzbekistan.svg","phone_code":"+998","states":null}`,
		}
	case PolitiesVanuatu:
		return &PolitiesCode{
			Name: "Vanuatu",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_Vanuatu_%28official%29.svg","phone_code":"+678","states":null}`,
		}
	case PolitiesVaticanCity:
		return &PolitiesCode{
			Name: "Vatican City",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_the_Vatican_City.svg","phone_code":"+418","states":null}`,
		}
	case PolitiesVenezuela:
		return &PolitiesCode{
			Name: "Venezuela",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7b/Flag_of_Venezuela_%28state%29.svg","phone_code":"+58","states":null}`,
		}
	case PolitiesVietnam:
		return &PolitiesCode{
			Name: "Vietnam",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Vietnam.svg","phone_code":"+84","states":null}`,
		}
	case PolitiesWallisandFutuna:
		return &PolitiesCode{
			Name: "Wallis and Futuna",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+681","states":null}`,
		}
	case PolitiesWesternSahara:
		return &PolitiesCode{
			Name: "Western Sahara",
			Label: "",
			Description: "",
			MetaData: `{"flag":null,"phone_code":"+212","states":null}`,
		}
	case PolitiesYemen:
		return &PolitiesCode{
			Name: "Yemen",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Yemen.svg","phone_code":"+967","states":null}`,
		}
	case PolitiesZambia:
		return &PolitiesCode{
			Name: "Zambia",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/06/Flag_of_Zambia.svg","phone_code":"+260","states":null}`,
		}
	case PolitiesZimbabwe:
		return &PolitiesCode{
			Name: "Zimbabwe",
			Label: "",
			Description: "",
			MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6a/Flag_of_Zimbabwe.svg","phone_code":"+263","states":null}`,
		}
	default:
		return nil
	}
}

// PolitiesMap returns a mapping of Polities objects with the code as the key.
func PolitiesMap() map[string]*PolitiesCode {
	return map[string]*PolitiesCode {
		PolitiesAalandIslands :
			&PolitiesCode{
				Name: "Aaland Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/52/Flag_of_%C3%85land.svg","phone_code":"+358","states":null}`,
			},
		PolitiesAfghanistan :
			&PolitiesCode{
				Name: "Afghanistan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Afghanistan.svg","phone_code":"+93","states":null}`,
			},
		PolitiesAlbania :
			&PolitiesCode{
				Name: "Albania",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/36/Flag_of_Albania.svg","phone_code":"+355","states":null}`,
			},
		PolitiesAlgeria :
			&PolitiesCode{
				Name: "Algeria",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg","phone_code":"+213","states":null}`,
			},
		PolitiesAmericanSamoa :
			&PolitiesCode{
				Name: "American Samoa",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/87/Flag_of_American_Samoa.svg","phone_code":"+1684","states":null}`,
			},
		PolitiesAndorra :
			&PolitiesCode{
				Name: "Andorra",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Andorra.svg","phone_code":"+376","states":null}`,
			},
		PolitiesAngola :
			&PolitiesCode{
				Name: "Angola",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9d/Flag_of_Angola.svg","phone_code":"+244","states":null}`,
			},
		PolitiesAnguilla :
			&PolitiesCode{
				Name: "Anguilla",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Anguilla.svg","phone_code":"+1264","states":null}`,
			},
		PolitiesAntarctica :
			&PolitiesCode{
				Name: "Antarctica",
				Label: "",
				Description: "",
				MetaData: `{"flag":null,"phone_code":"+672","states":null}`,
			},
		PolitiesAntiguaandBarbuda :
			&PolitiesCode{
				Name: "Antigua and Barbuda",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Antigua_and_Barbuda.svg","phone_code":"+1268","states":null}`,
			},
		PolitiesArgentina :
			&PolitiesCode{
				Name: "Argentina",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1a/Flag_of_Argentina.svg","phone_code":"+54","states":null}`,
			},
		PolitiesArmenia :
			&PolitiesCode{
				Name: "Armenia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2f/Flag_of_Armenia.svg","phone_code":"+374","states":null}`,
			},
		PolitiesAruba :
			&PolitiesCode{
				Name: "Aruba",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Aruba.svg","phone_code":"+297","states":null}`,
			},
		PolitiesAfricanUnion :
			&PolitiesCode{
				Name: "African Union",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_the_African_Union.svg","states":{"AGO":null,"BDI":null,"BEN":null,"BFA":null,"BWA":null,"CAF":null,"CIV":null,"CMR":null,"COD":null,"COM":null,"CPV":null,"DJI":null,"DZA":null,"EGY":null,"ERI":null,"ETH":null,"GAB":null,"GHA":null,"GIN":null,"GMB":null,"GNB":null,"GNQ":null,"KEN":null,"LBR":null,"LBY":null,"LSO":null,"SWZ":null,"TCD":null}}`,
			},
		PolitiesAustralia :
			&PolitiesCode{
				Name: "Australia",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/b/b9/Flag_of_Australia.svg","gov_fiscal_year":"0701-0630","phone_code":"+61","states":{"AUACT":"Australian Capital Territory","AUCC":"Cocos (Keening) Island","AUCX":"Christmas Island","AUHM":"Heard Island and McDonalds Islands","AUJBT":"Jervis Bay Territory","AUNF":"Norfolk Island","AUNSW":"New South Wales","AUNT":"Northern Territory","AUQLD":"Queensland","AUSA":"South Australia","AUTAS":"Tasmania","AUVIC":"Victoria","AUWA":"Western Australia"}}`,
			},
		PolitiesAustria :
			&PolitiesCode{
				Name: "Austria",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/41/Flag_of_Austria.svg","gov_fiscal_year":"0101-1231","phone_code":"+43","states":null}`,
			},
		PolitiesAzerbaijan :
			&PolitiesCode{
				Name: "Azerbaijan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg","phone_code":"+994","states":null}`,
			},
		PolitiesTheBahamas :
			&PolitiesCode{
				Name: "The Bahamas",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/93/Flag_of_the_Bahamas.svg","phone_code":"+1242","states":null}`,
			},
		PolitiesBahrain :
			&PolitiesCode{
				Name: "Bahrain",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Bahrain.svg","phone_code":"+973","states":null}`,
			},
		PolitiesBangladesh :
			&PolitiesCode{
				Name: "Bangladesh",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/f9/Flag_of_Bangladesh.svg","gov_fiscal_year":"0701-0630","phone_code":"+880","states":null}`,
			},
		PolitiesBarbados :
			&PolitiesCode{
				Name: "Barbados",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ef/Flag_of_Barbados.svg","phone_code":"+1246","states":null}`,
			},
		PolitiesBelarus :
			&PolitiesCode{
				Name: "Belarus",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/85/Flag_of_Belarus.svg","phone_code":"+375","states":null}`,
			},
		PolitiesBelgium :
			&PolitiesCode{
				Name: "Belgium",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg","gov_fiscal_year":"0101-1231","phone_code":"+32","states":null}`,
			},
		PolitiesBelize :
			&PolitiesCode{
				Name: "Belize",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e7/Flag_of_Belize.svg","phone_code":"+501","states":null}`,
			},
		PolitiesBenin :
			&PolitiesCode{
				Name: "Benin",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Benin.svg","phone_code":"+229","states":null}`,
			},
		PolitiesBermuda :
			&PolitiesCode{
				Name: "Bermuda",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bermuda.svg","phone_code":"+1441","states":null}`,
			},
		PolitiesBhutan :
			&PolitiesCode{
				Name: "Bhutan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/91/Flag_of_Bhutan.svg","phone_code":"+975","states":null}`,
			},
		PolitiesBolivia :
			&PolitiesCode{
				Name: "Bolivia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/de/Flag_of_Bolivia_%28state%29.svg","phone_code":"+591","states":null}`,
			},
		PolitiesBonaireStEustasuisandSaba :
			&PolitiesCode{
				Name: "Bonaire, St Eustasuis and Saba",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","phone_code":"+599","states":null}`,
			},
		PolitiesBosniaandHerzegovina :
			&PolitiesCode{
				Name: "Bosnia and Herzegovina",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bosnia_and_Herzegovina.svg","phone_code":"+387","states":null}`,
			},
		PolitiesBotswana :
			&PolitiesCode{
				Name: "Botswana",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Botswana.svg","phone_code":"+267","states":null}`,
			},
		PolitiesBouvetIsland :
			&PolitiesCode{
				Name: "Bouvet Island",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
			},
		PolitiesBrazil :
			&PolitiesCode{
				Name: "Brazil",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg","gov_fiscal_year":"0101-1231","phone_code":"+55","states":{"BRAC":"Acre","BRAL":"Alagoas","BRAM":"Amazonas","BRAP":"Amapá","BRBA":"Bahia","BRCE":"Ceará","BRDF":"Federal District","BRES":"Espírito Santo","BRGO":"Goiás","BRMA":"Maranhão","BRMG":"Minas Gerais","BRMS":"Mato Grosso do Sul","BRMT":"Mato Grosso","BRPA":"Pará","BRPB":"Paraíba","BRPE":"Pernambuco","BRPI":"Piauí","BRPR":"Paraná","BRRJ":"Rio de Janeiro","BRRN":"Rio Grande do Norte","BRRO":"Rondônia","BRRR":"Roraima","BRRS":"Rio Grande do Sul","BRSC":"Santa Catarina","BRSE":"Sergipe","BRSP":"São Paulo","BRTO":"Tocantins"}}`,
			},
		PolitiesBritishIndianOceanTerritory :
			&PolitiesCode{
				Name: "British Indian Ocean Territory",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_the_British_Indian_Ocean_Territory.svg","phone_code":"+246","states":null}`,
			},
		PolitiesBritishVirginIslands :
			&PolitiesCode{
				Name: "British Virgin Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/42/Flag_of_the_British_Virgin_Islands.svg","phone_code":"+1284","states":null}`,
			},
		PolitiesBrunei :
			&PolitiesCode{
				Name: "Brunei",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Brunei.svg","phone_code":"+673","states":null}`,
			},
		PolitiesBulgaria :
			&PolitiesCode{
				Name: "Bulgaria",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Bulgaria.svg","phone_code":"+359","states":null}`,
			},
		PolitiesBurkinaFaso :
			&PolitiesCode{
				Name: "Burkina Faso",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Burkina_Faso.svg","phone_code":"+226","states":null}`,
			},
		PolitiesBurundi :
			&PolitiesCode{
				Name: "Burundi",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/50/Flag_of_Burundi.svg","phone_code":"+257","states":null}`,
			},
		PolitiesCambodia :
			&PolitiesCode{
				Name: "Cambodia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_Cambodia.svg","phone_code":"+855","states":null}`,
			},
		PolitiesCameroon :
			&PolitiesCode{
				Name: "Cameroon",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Cameroon.svg","phone_code":"+237","states":null}`,
			},
		PolitiesCanada :
			&PolitiesCode{
				Name: "Canada",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg","gov_fiscal_year":"0401-0331","phone_code":"+1","states":{"CAAB":"Alberta","CABC":"British Columbia","CAMB":"Manitoba","CANB":"New Brunswick","CANL":"Newfoundland and Labrador","CANS":"Nova Scotia","CANT":"Northwest Territories","CANU":"Nunavut","CAON":"Ontario","CAPE":"Prince Edward Island","CAQC":"Quebec","CASK":"Saskatchewan","CAYT":"Yukon"}}`,
			},
		PolitiesCapeVerde :
			&PolitiesCode{
				Name: "Cape Verde",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Cape_Verde.svg","phone_code":"+238","states":null}`,
			},
		PolitiesCaymanIslands :
			&PolitiesCode{
				Name: "Cayman Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_the_Cayman_Islands.svg","phone_code":"+1345","states":null}`,
			},
		PolitiesCentralAfricanRepublic :
			&PolitiesCode{
				Name: "Central African Republic",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Central_African_Republic.svg","phone_code":"+236","states":null}`,
			},
		PolitiesChad :
			&PolitiesCode{
				Name: "Chad",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4b/Flag_of_Chad.svg","phone_code":"+235","states":null}`,
			},
		PolitiesChile :
			&PolitiesCode{
				Name: "Chile",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/78/Flag_of_Chile.svg","phone_code":"+56","states":null}`,
			},
		PolitiesChina :
			&PolitiesCode{
				Name: "China",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_the_People%27s_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+86","states":{"CNAH":"Anhui","CNBJ":"Beijing","CNCQ":"Chongqing","CNFJ":"Fujian","CNGD":"Guangdong","CNGS":"Gansu","CNGX":"Guangxi","CNGZ":"Guizhou","CNHA":"Henan","CNHB":"Hubei","CNHE":"Hebei","CNHI":"Hainan","CNHK":"Hong Kong (Xianggang)","CNHL":"Heilongjiang","CNHN":"Hunan","CNJL":"Jilin","CNJS":"Jiangsu","CNJX":"Jiangxi","CNLN":"Liaoning","CNMC":"Macao (Aomen)","CNNM":"Nei Mongol (mn)","CNNX":"Ningxia","CNQH":"Qinghai","CNSC":"Sichuan","CNSD":"Shandong","CNSH":"Shanghai","CNSN":"Shaanxi","CNSX":"Shanxi","CNTJ":"Tianjin","CNTW":"Taiwan","CNXJ":"Xinjiang","CNXZ":"Xizang","CNYN":"Yunnan","CNZJ":"Zhejiang"}}`,
			},
		PolitiesChristmasIsland :
			&PolitiesCode{
				Name: "Christmas Island",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/67/Flag_of_Christmas_Island.svg","phone_code":"+53","states":null}`,
			},
		PolitiesCocosIslands :
			&PolitiesCode{
				Name: "Cocos Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Cocos_%28Keeling%29_Islands.svg","phone_code":"+61","states":null}`,
			},
		PolitiesColombia :
			&PolitiesCode{
				Name: "Colombia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Colombia.svg","phone_code":"+57","states":null}`,
			},
		PolitiesComoros :
			&PolitiesCode{
				Name: "Comoros",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/94/Flag_of_the_Comoros.svg","phone_code":"+269","states":null}`,
			},
		PolitiesCongoBrazzaville :
			&PolitiesCode{
				Name: "Congo-Brazzaville",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_the_Republic_of_the_Congo.svg","phone_code":"+242","states":null}`,
			},
		PolitiesCongoKinshasa :
			&PolitiesCode{
				Name: "Congo-Kinshasa",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Democratic_Republic_of_the_Congo.svg","phone_code":"+243","states":null}`,
			},
		PolitiesCookIslands :
			&PolitiesCode{
				Name: "Cook Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/35/Flag_of_the_Cook_Islands.svg","phone_code":"+682","states":null}`,
			},
		PolitiesCostaRica :
			&PolitiesCode{
				Name: "Costa Rica",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"1001-0931","flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Costa_Rica_%28state%29.svg","gov_fiscal_year":"1001-0931","phone_code":"+506","states":null}`,
			},
		PolitiesCroatia :
			&PolitiesCode{
				Name: "Croatia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Croatia.svg","phone_code":"+385","states":null}`,
			},
		PolitiesCuba :
			&PolitiesCode{
				Name: "Cuba",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bd/Flag_of_Cuba.svg","phone_code":"+53","states":null}`,
			},
		PolitiesCuracao(NetherlandsAntilles) :
			&PolitiesCode{
				Name: "Curacao (Netherlands Antilles)",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_Cura%C3%A7ao.svg","phone_code":"+599","states":null}`,
			},
		PolitiesCyprus :
			&PolitiesCode{
				Name: "Cyprus",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Cyprus.svg","phone_code":"+357","states":null}`,
			},
		PolitiesCzechRepublic :
			&PolitiesCode{
				Name: "Czech Republic",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_Czech_Republic.svg","phone_code":"+420","states":null}`,
			},
		PolitiesDenmark :
			&PolitiesCode{
				Name: "Denmark",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Denmark.svg","phone_code":"+45","states":null}`,
			},
		PolitiesDjibouti :
			&PolitiesCode{
				Name: "Djibouti",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_Djibouti.svg","phone_code":"+253","states":null}`,
			},
		PolitiesDominica :
			&PolitiesCode{
				Name: "Dominica",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c4/Flag_of_Dominica.svg","phone_code":"+1767","states":null}`,
			},
		PolitiesDominicanRepublic :
			&PolitiesCode{
				Name: "Dominican Republic",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_the_Dominican_Republic.svg","phone_code":"+1829","states":null}`,
			},
		PolitiesEastTimor :
			&PolitiesCode{
				Name: "East Timor",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/26/Flag_of_East_Timor.svg","phone_code":"+670","states":null}`,
			},
		PolitiesEcuador :
			&PolitiesCode{
				Name: "Ecuador",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e8/Flag_of_Ecuador.svg","phone_code":"+593","states":null}`,
			},
		PolitiesEgypt :
			&PolitiesCode{
				Name: "Egypt",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Egypt.svg","gov_fiscal_year":"0701-0630","phone_code":"+20","states":null}`,
			},
		PolitiesElSalvador :
			&PolitiesCode{
				Name: "El Salvador",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_El_Salvador.svg","phone_code":"+503","states":null}`,
			},
		PolitiesEquatorialGuinea :
			&PolitiesCode{
				Name: "Equatorial Guinea",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Equatorial_Guinea.svg","phone_code":"+240","states":null}`,
			},
		PolitiesEritrea :
			&PolitiesCode{
				Name: "Eritrea",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/29/Flag_of_Eritrea.svg","phone_code":"+291","states":null}`,
			},
		PolitiesEstonia :
			&PolitiesCode{
				Name: "Estonia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8f/Flag_of_Estonia.svg","phone_code":"+372","states":null}`,
			},
		PolitiesEthiopia :
			&PolitiesCode{
				Name: "Ethiopia",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0708-0707","flag":"https://upload.wikimedia.org/wikipedia/commons/7/71/Flag_of_Ethiopia.svg","gov_fiscal_year":"0708-0707","phone_code":"+251","states":null}`,
			},
		PolitiesEuropeanUnion :
			&PolitiesCode{
				Name: "European Union",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b7/Flag_of_Europe.svg","states":{"AUT":null,"BEL":null,"BGR":null,"CYP":null,"CZE":null,"DEU":null,"DNK":null,"ESP":null,"EST":null,"FIN":null,"FRA":null,"GBR":null,"GRC":null,"HRV":null,"HUN":null,"IRL":null,"ITA":null,"LTU":null,"LUX":null,"LVA":null,"MLT":null,"NLD":null,"POL":null,"PRT":null,"ROU":null,"SVK":null,"SVN":null,"SWE":null}}`,
			},
		PolitiesFalklandIslands :
			&PolitiesCode{
				Name: "Falkland Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_the_Falkland_Islands.svg","phone_code":"+500","states":null}`,
			},
		PolitiesFaroeIslands :
			&PolitiesCode{
				Name: "Faroe Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3c/Flag_of_the_Faroe_Islands.svg","phone_code":"+298","states":null}`,
			},
		PolitiesFiji :
			&PolitiesCode{
				Name: "Fiji",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/ba/Flag_of_Fiji.svg","phone_code":"+679","states":null}`,
			},
		PolitiesFinland :
			&PolitiesCode{
				Name: "Finland",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg","phone_code":"+358","states":null}`,
			},
		PolitiesFrance :
			&PolitiesCode{
				Name: "France",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+33","states":null}`,
			},
		PolitiesFrenchGuiana :
			&PolitiesCode{
				Name: "French Guiana",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+594","states":null}`,
			},
		PolitiesFrenchPolynesia :
			&PolitiesCode{
				Name: "French Polynesia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/db/Flag_of_French_Polynesia.svg","phone_code":"+689","states":null}`,
			},
		PolitiesFrenchSouthernandAntarcticLands :
			&PolitiesCode{
				Name: "French Southern and Antarctic Lands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a7/Flag_of_the_French_Southern_and_Antarctic_Lands.svg","phone_code":"+262","states":null}`,
			},
		PolitiesGabon :
			&PolitiesCode{
				Name: "Gabon",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/04/Flag_of_Gabon.svg","phone_code":"+241","states":null}`,
			},
		PolitiesTheGambia :
			&PolitiesCode{
				Name: "The Gambia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_The_Gambia.svg","phone_code":"+220","states":null}`,
			},
		PolitiesGeorgia :
			&PolitiesCode{
				Name: "Georgia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Georgia.svg","phone_code":"+995","states":null}`,
			},
		PolitiesGermany :
			&PolitiesCode{
				Name: "Germany",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg","gov_fiscal_year":"0101-1231","phone_code":"+49","states":null}`,
			},
		PolitiesGhana :
			&PolitiesCode{
				Name: "Ghana",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Ghana.svg","phone_code":"+233","states":null}`,
			},
		PolitiesGibraltar :
			&PolitiesCode{
				Name: "Gibraltar",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/02/Flag_of_Gibraltar.svg","phone_code":"+350","states":null}`,
			},
		PolitiesGreece :
			&PolitiesCode{
				Name: "Greece",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Greece.svg","gov_fiscal_year":"0101-1231","phone_code":"+30","states":null}`,
			},
		PolitiesGreenland :
			&PolitiesCode{
				Name: "Greenland",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_Greenland.svg","phone_code":"+299","states":null}`,
			},
		PolitiesGrenada :
			&PolitiesCode{
				Name: "Grenada",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Grenada.svg","phone_code":"+1473","states":null}`,
			},
		PolitiesGuadeloupe :
			&PolitiesCode{
				Name: "Guadeloupe",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+590","states":null}`,
			},
		PolitiesGuam :
			&PolitiesCode{
				Name: "Guam",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/07/Flag_of_Guam.svg","phone_code":"+1671","states":null}`,
			},
		PolitiesGuatemala :
			&PolitiesCode{
				Name: "Guatemala",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ec/Flag_of_Guatemala.svg","phone_code":"+502","states":null}`,
			},
		PolitiesGuernsey :
			&PolitiesCode{
				Name: "Guernsey",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Guernsey.svg","phone_code":"+44","states":null}`,
			},
		PolitiesGuinea :
			&PolitiesCode{
				Name: "Guinea",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_Guinea.svg","phone_code":"+224","states":null}`,
			},
		PolitiesGuineaBissau :
			&PolitiesCode{
				Name: "Guinea-Bissau",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Guinea-Bissau.svg","phone_code":"+245","states":null}`,
			},
		PolitiesGuyana :
			&PolitiesCode{
				Name: "Guyana",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_Guyana.svg","phone_code":"+592","states":null}`,
			},
		PolitiesHaiti :
			&PolitiesCode{
				Name: "Haiti",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Haiti.svg","phone_code":"+509","states":null}`,
			},
		PolitiesHeardIslandandMcDonaldIslands :
			&PolitiesCode{
				Name: "Heard Island and McDonald Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg","phone_code":"+672","states":null}`,
			},
		PolitiesHonduras :
			&PolitiesCode{
				Name: "Honduras",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/82/Flag_of_Honduras.svg","phone_code":"+504","states":null}`,
			},
		PolitiesHongKong :
			&PolitiesCode{
				Name: "Hong Kong",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg","gov_fiscal_year":"0401-0331","phone_code":"+852","states":null}`,
			},
		PolitiesHungary :
			&PolitiesCode{
				Name: "Hungary",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c1/Flag_of_Hungary.svg","phone_code":"+36","states":null}`,
			},
		PolitiesIceland :
			&PolitiesCode{
				Name: "Iceland",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Iceland.svg","phone_code":"+354","states":null}`,
			},
		PolitiesIndia :
			&PolitiesCode{
				Name: "India",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg","gov_fiscal_year":"0401-0331","phone_code":"+91","states":{"INAN":"Andaman and Nicobar Islands","INAP":"Andhra Pradesh","INAR":"Arunachal Pradesh","INAS":"Assam","INBR":"Bihar","INCH":"Chandigarh","INCT":"Chhattisgarh","INDD":"Daman and Diu","INDL":"Delhi","INDN":"Dadra and Nagar Haveli","INGA":"Goa","INGJ":"Gujarat","INHP":"Himachal Pradesh","INHR":"Haryana","INJH":"Jharkhand","INJK":"Jammu and Kashmir","INKA":"Karnataka","INKL":"Kerala","INLD":"Lakshadweep","INMH":"Maharashtra","INML":"Meghalaya","INMN":"Manipur","INMP":"Madhya Pradesh","INMZ":"Mizoram","INNL":"Nagaland","INOR":"Odisha (formerly known as Orissa)","INPB":"Punjab","INPY":"Puducherry (Pondicherry)","INRJ":"Rajasthan","INSK":"Sikkim","INTG":"Telangana","INTN":"Tamil Nadu","INTR":"Tripura","INUP":"Uttar Pradesh","INUT":"Uttarakhand","INWB":"West Bengal"}}`,
			},
		PolitiesIndonesia :
			&PolitiesCode{
				Name: "Indonesia",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg","gov_fiscal_year":"0101-1231","phone_code":"+62","states":null}`,
			},
		PolitiesIran :
			&PolitiesCode{
				Name: "Iran",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ca/Flag_of_Iran.svg","phone_code":"+98","states":null}`,
			},
		PolitiesIraq :
			&PolitiesCode{
				Name: "Iraq",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Iraq.svg","phone_code":"+964","states":null}`,
			},
		PolitiesIreland :
			&PolitiesCode{
				Name: "Ireland",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/45/Flag_of_Ireland.svg","gov_fiscal_year":"0101-1231","phone_code":"+353","states":null}`,
			},
		PolitiesIsleofMan :
			&PolitiesCode{
				Name: "Isle of Man",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5d/Flag_of_the_Isle_of_Mann.svg","phone_code":"+44","states":null}`,
			},
		PolitiesIsrael :
			&PolitiesCode{
				Name: "Israel",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Israel.svg","states":null}`,
			},
		PolitiesItaly :
			&PolitiesCode{
				Name: "Italy",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/0/03/Flag_of_Italy.svg","phone_code":"+39","states":null}`,
			},
		PolitiesCoteD'Ivoire(FormerIvoryCoast) :
			&PolitiesCode{
				Name: "Cote D'Ivoire (Former Ivory Coast)",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_C%C3%B4te_d%27Ivoire.svg","phone_code":"+225","states":null}`,
			},
		PolitiesJamaica :
			&PolitiesCode{
				Name: "Jamaica",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Jamaica.svg","phone_code":"+1876","states":null}`,
			},
		PolitiesJapan :
			&PolitiesCode{
				Name: "Japan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg","states":null}`,
			},
		PolitiesJersey :
			&PolitiesCode{
				Name: "Jersey",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1c/Flag_of_Jersey.svg","phone_code":"+44","states":null}`,
			},
		PolitiesJordan :
			&PolitiesCode{
				Name: "Jordan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c0/Flag_of_Jordan.svg","phone_code":"+962","states":null}`,
			},
		PolitiesKazakhstan :
			&PolitiesCode{
				Name: "Kazakhstan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kazakhstan.svg","phone_code":"+7","states":null}`,
			},
		PolitiesKenya :
			&PolitiesCode{
				Name: "Kenya",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Kenya.svg","phone_code":"+254","states":null}`,
			},
		PolitiesKiribati :
			&PolitiesCode{
				Name: "Kiribati",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kiribati.svg","phone_code":"+686","states":null}`,
			},
		PolitiesKuwait :
			&PolitiesCode{
				Name: "Kuwait",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/aa/Flag_of_Kuwait.svg","phone_code":"+965","states":null}`,
			},
		PolitiesKyrgyzstan :
			&PolitiesCode{
				Name: "Kyrgyzstan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c7/Flag_of_Kyrgyzstan.svg","phone_code":"+996","states":null}`,
			},
		PolitiesLaos :
			&PolitiesCode{
				Name: "Laos",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Laos.svg","phone_code":"+856","states":null}`,
			},
		PolitiesLatvia :
			&PolitiesCode{
				Name: "Latvia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Latvia.svg","phone_code":"+371","states":null}`,
			},
		PolitiesLebanon :
			&PolitiesCode{
				Name: "Lebanon",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/59/Flag_of_Lebanon.svg","phone_code":"+961","states":null}`,
			},
		PolitiesLesotho :
			&PolitiesCode{
				Name: "Lesotho",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Lesotho.svg","phone_code":"+266","states":null}`,
			},
		PolitiesLiberia :
			&PolitiesCode{
				Name: "Liberia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b8/Flag_of_Liberia.svg","phone_code":"+231","states":null}`,
			},
		PolitiesLibya :
			&PolitiesCode{
				Name: "Libya",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/05/Flag_of_Libya.svg","phone_code":"+218","states":null}`,
			},
		PolitiesLiechtenstein :
			&PolitiesCode{
				Name: "Liechtenstein",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/47/Flag_of_Liechtenstein.svg","phone_code":"+423","states":null}`,
			},
		PolitiesLithuania :
			&PolitiesCode{
				Name: "Lithuania",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Lithuania.svg","phone_code":"+370","states":null}`,
			},
		PolitiesLuxembourg :
			&PolitiesCode{
				Name: "Luxembourg",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/da/Flag_of_Luxembourg.svg","phone_code":"+352","states":null}`,
			},
		PolitiesMacau :
			&PolitiesCode{
				Name: "Macau",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/63/Flag_of_Macau.svg","phone_code":"+853","states":null}`,
			},
		PolitiesMacedonia :
			&PolitiesCode{
				Name: "Macedonia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_Macedonia.svg","phone_code":"+389","states":null}`,
			},
		PolitiesMadagascar :
			&PolitiesCode{
				Name: "Madagascar",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Madagascar.svg","phone_code":"+261","states":null}`,
			},
		PolitiesMalawi :
			&PolitiesCode{
				Name: "Malawi",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d1/Flag_of_Malawi.svg","phone_code":"+265","states":null}`,
			},
		PolitiesMalaysia :
			&PolitiesCode{
				Name: "Malaysia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/66/Flag_of_Malaysia.svg","phone_code":"+60","states":null}`,
			},
		PolitiesMaldives :
			&PolitiesCode{
				Name: "Maldives",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Maldives.svg","phone_code":"+960","states":null}`,
			},
		PolitiesMali :
			&PolitiesCode{
				Name: "Mali",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Mali.svg","phone_code":"+223","states":null}`,
			},
		PolitiesMalta :
			&PolitiesCode{
				Name: "Malta",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Malta.svg","phone_code":"+356","states":null}`,
			},
		PolitiesMarshallIslands :
			&PolitiesCode{
				Name: "Marshall Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2e/Flag_of_the_Marshall_Islands.svg","phone_code":"+692","states":null}`,
			},
		PolitiesMartinique :
			&PolitiesCode{
				Name: "Martinique",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+596","states":null}`,
			},
		PolitiesMauritania :
			&PolitiesCode{
				Name: "Mauritania",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/New_National_Flag_of_Mauritania.png","phone_code":"+222","states":null}`,
			},
		PolitiesMauritius :
			&PolitiesCode{
				Name: "Mauritius",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Mauritius.svg","phone_code":"+230","states":null}`,
			},
		PolitiesMayotte :
			&PolitiesCode{
				Name: "Mayotte",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Mayotte_%28local%29.svg","phone_code":"+269","states":null}`,
			},
		PolitiesMexico :
			&PolitiesCode{
				Name: "Mexico",
				Label: "",
				Description: "",
				MetaData: `{"MXAGU":"Aguascalientes","MXBCN":"Baja California","MXBCS":"Baja California Sur","MXCAM":"Campeche","MXCHH":"Chihuahua","MXCHP":"Chiapas","MXCOA":"Coahuila","MXCOL":"Colima","MXDIF":"Distrito Federal","MXDUR":"Durango","MXGRO":"Guerrero","MXGUA":"Guanajuato","MXHID":"Hidalgo","MXJAL":"Jalisco","MXMEX":"Mexico (Federal District)","MXMIC":"Michoacán","MXMOR":"Morelos","MXNAY":"Nayarit","MXNLE":"Nuevo León","MXOAX":"Oaxaca","MXPUE":"Puebla","MXQUE":"Querétaro","MXROO":"Quintana Roo","MXSIN":"Sinaloa","MXSLP":"San Luis Potosí","MXSON":"Sonora","MXTAB":"Tabasco","MXTAM":"Tamaulipas","MXTLA":"Tlaxcala","MXVER":"Veracruz","MXYUC":"Yucatán","MXZAC":"Zacatecas","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Mexico.svg","phone_code":"+52","states":null}`,
			},
		PolitiesMicronesia :
			&PolitiesCode{
				Name: "Micronesia",
				Label: "",
				Description: "",
				MetaData: `{"flag":null,"phone_code":"+691","states":null}`,
			},
		PolitiesMoldova :
			&PolitiesCode{
				Name: "Moldova",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Moldova.svg","phone_code":"+373","states":null}`,
			},
		PolitiesMonaco :
			&PolitiesCode{
				Name: "Monaco",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ea/Flag_of_Monaco.svg","phone_code":"+377","states":null}`,
			},
		PolitiesMongolia :
			&PolitiesCode{
				Name: "Mongolia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4c/Flag_of_Mongolia.svg","phone_code":"+976","states":null}`,
			},
		PolitiesMontenegro :
			&PolitiesCode{
				Name: "Montenegro",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Montenegro.svg","phone_code":"+382","states":null}`,
			},
		PolitiesMontserrat :
			&PolitiesCode{
				Name: "Montserrat",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Montserrat.svg","phone_code":"+1664","states":null}`,
			},
		PolitiesMorocco :
			&PolitiesCode{
				Name: "Morocco",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Morocco.svg","phone_code":"+212","states":null}`,
			},
		PolitiesMozambique :
			&PolitiesCode{
				Name: "Mozambique",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Mozambique.svg","phone_code":"+258","states":null}`,
			},
		PolitiesMyanmar :
			&PolitiesCode{
				Name: "Myanmar",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8c/Flag_of_Myanmar.svg","phone_code":"+95","states":null}`,
			},
		PolitiesNamibia :
			&PolitiesCode{
				Name: "Namibia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Namibia.svg","phone_code":"+264","states":null}`,
			},
		PolitiesNauru :
			&PolitiesCode{
				Name: "Nauru",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/Flag_of_Nauru.svg","phone_code":"+674","states":null}`,
			},
		PolitiesNepal :
			&PolitiesCode{
				Name: "Nepal",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9b/Flag_of_Nepal.svg","gov_fiscal_year":"0101-1231","phone_code":"+977","states":null}`,
			},
		PolitiesNetherlands :
			&PolitiesCode{
				Name: "Netherlands",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","gov_fiscal_year":"0101-1231","phone_code":"+31","states":null}`,
			},
		PolitiesNewCaledonia :
			&PolitiesCode{
				Name: "New Caledonia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+687","states":null}`,
			},
		PolitiesNewZealand :
			&PolitiesCode{
				Name: "New Zealand",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3e/Flag_of_New_Zealand.svg","phone_code":"+64","states":null}`,
			},
		PolitiesNicaragua :
			&PolitiesCode{
				Name: "Nicaragua",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Nicaragua.svg","phone_code":"+505","states":null}`,
			},
		PolitiesNiger :
			&PolitiesCode{
				Name: "Niger",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f4/Flag_of_Niger.svg","phone_code":"+227","states":null}`,
			},
		PolitiesNigeria :
			&PolitiesCode{
				Name: "Nigeria",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/79/Flag_of_Nigeria.svg","phone_code":"+234","states":null}`,
			},
		PolitiesNiue :
			&PolitiesCode{
				Name: "Niue",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Niue.svg","phone_code":"+683","states":null}`,
			},
		PolitiesNorfolkandPhilipIsland :
			&PolitiesCode{
				Name: "Norfolk and Philip Island",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Norfolk_Island.svg","phone_code":"+672","states":null}`,
			},
		PolitiesNorthKorea :
			&PolitiesCode{
				Name: "North Korea",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_North_Korea.svg","phone_code":"+850","states":null}`,
			},
		PolitiesNorthernMarianaIslands :
			&PolitiesCode{
				Name: "Northern Mariana Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e0/Flag_of_the_Northern_Mariana_Islands.svg","phone_code":"+1670","states":null}`,
			},
		PolitiesNorway :
			&PolitiesCode{
				Name: "Norway",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","phone_code":"+47","states":null}`,
			},
		PolitiesOman :
			&PolitiesCode{
				Name: "Oman",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Oman.svg","phone_code":"+968","states":null}`,
			},
		PolitiesPakistan :
			&PolitiesCode{
				Name: "Pakistan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/32/Flag_of_Pakistan.svg","phone_code":"+92","states":null}`,
			},
		PolitiesPalau :
			&PolitiesCode{
				Name: "Palau",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Palau.svg","phone_code":"+680","states":null}`,
			},
		PolitiesPalestinianTerritory :
			&PolitiesCode{
				Name: "Palestinian Territory",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Palestine.svg","phone_code":"+970","states":null}`,
			},
		PolitiesPanama :
			&PolitiesCode{
				Name: "Panama",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/ab/Flag_of_Panama.svg","phone_code":"+507","states":null}`,
			},
		PolitiesPapuaNewGuinea :
			&PolitiesCode{
				Name: "Papua New Guinea",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e3/Flag_of_Papua_New_Guinea.svg","phone_code":"+675","states":null}`,
			},
		PolitiesParaguay :
			&PolitiesCode{
				Name: "Paraguay",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Paraguay.svg","phone_code":"+595","states":null}`,
			},
		PolitiesPeru :
			&PolitiesCode{
				Name: "Peru",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Peru_%28state%29.svg","phone_code":"+51","states":null}`,
			},
		PolitiesPhilippines :
			&PolitiesCode{
				Name: "Philippines",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_the_Philippines.svg","phone_code":"+63","states":null}`,
			},
		PolitiesPitcairnIslands :
			&PolitiesCode{
				Name: "Pitcairn Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_the_Pitcairn_Islands.svg","phone_code":"+64","states":null}`,
			},
		PolitiesPoland :
			&PolitiesCode{
				Name: "Poland",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg","phone_code":"+48","states":null}`,
			},
		PolitiesPortugal :
			&PolitiesCode{
				Name: "Portugal",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Portugal.svg","states":null}`,
			},
		PolitiesPuertoRico :
			&PolitiesCode{
				Name: "Puerto Rico",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/28/Flag_of_Puerto_Rico.svg","phone_code":"+1787","states":null}`,
			},
		PolitiesQatar :
			&PolitiesCode{
				Name: "Qatar",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Qatar.svg","states":null}`,
			},
		PolitiesRéunion :
			&PolitiesCode{
				Name: "Réunion",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5a/Flag_of_R%C3%A9union.svg","phone_code":"+262","states":null}`,
			},
		PolitiesRomania :
			&PolitiesCode{
				Name: "Romania",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Romania.svg","states":null}`,
			},
		PolitiesRussia :
			&PolitiesCode{
				Name: "Russia",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/f/f3/Flag_of_Russia.svg","gov_fiscal_year":"0101-1231","phone_code":"+7","states":{"RUAD":"Adygeya, Respublika","RUAL":"Altay, Respublika","RUALT":"Altayskiy kray","RUAMU":"Amurskaya oblast\"","RUARK":"Arkhangel\"skaya oblast\"","RUAST":"Astrakhanskaya oblast\"","RUBA":"Bashkortostan, Respublika","RUBEL":"Belgorodskaya oblast\"","RUBRY":"Bryanskaya oblast\"","RUBU":"Buryatiya, Respublika","RUCE":"Chechenskaya Respublika","RUCHE":"Chelyabinskaya oblast\"","RUCHU":"Chukotskiy avtonomnyy okrug","RUCU":"Chuvashskaya Respublika","RUDA":"Dagestan, Respublika","RUIN":"Ingushetiya, Respublika","RUIRK":"Irkutskaya oblast\"","RUIVA":"Ivanovskaya oblast\"","RUKAM":"Kamchatskiy kray","RUKB":"Kabardino-Balkarskaya Respublika","RUKC":"Karachayevo-Cherkesskaya Respubl.","RUKDA":"Krasnodarskiy kray","RUKEM":"Kemerovskaya oblast\"","RUKGD":"Kaliningradskaya oblast\"","RUKGN":"Kurganskaya oblast\"","RUKHA":"Khabarovskiy kray","RUKHM":"Khanty-Mansiyskiy avtonomnyy okrug-Yugra","RUKIR":"Kirovskaya oblast\"","RUKK":"Khakasiya, Respublika","RUKL":"Kalmykiya, Respublika","RUKLU":"Kaluzhskaya oblast\"","RUKO":"Komi, Respublika","RUKOS":"Kostromskaya oblast\"","RUKR":"Kareliya, Respublika","RUKRS":"Kurskaya oblast\"","RUKYA":"Krasnoyarskiy kray","RULEN":"Leningradskaya oblast\"","RULIP":"Lipetskaya oblast\"","RUMAG":"Magadanskaya oblast\"","RUME":"Mariy El, Respublika","RUMO":"Mordoviya, Respublika","RUMOS":"Moskovskaya oblast\"","RUMOW":"Moskva (autonomous city)","RUMUR":"Murmanskaya oblast\"","RUNEN":"Nenetskiy avtonomnyy okrug","RUNGR":"Novgorodskaya oblast\"","RUNIZ":"Nizhegorodskaya oblast\"","RUNVS":"Novosibirskaya oblast\"","RUOMS":"Omskaya oblast\"","RUORE":"Orenburgskaya oblast\"","RUORL":"Orlovskaya oblast\"","RUPER":"Permskiy kray","RUPNZ":"Penzenskaya oblast\"","RUPRI":"Primorskiy kray","RUPSK":"Pskovskaya oblast\"","RUROS":"Rostovskaya oblast\"","RURYA":"Ryazanskaya oblast\"","RUSA":"Sakha, Respublika","RUSAK":"Sakhalinskaya oblast\"","RUSAM":"Samarskaya oblast\"","RUSAR":"Saratovskaya oblast\"","RUSE":"Severnaya Osetiya-Alaniya, Respubl.","RUSMO":"Smolenskaya oblast\"","RUSPE":"Sankt-Peterburg (autonomous city)","RUSTA":"Stavropol\"skiy kray","RUSVE":"Sverdlovskaya oblast\"","RUTA":"Tatarstan, Respublika","RUTAM":"Tambovskaya oblast\"","RUTOM":"Tomskaya oblast\"","RUTUL":"Tul\"skaya oblast\"","RUTVE":"Tverskaya oblast\"","RUTY":"Tyva, Respublika","RUTYU":"Tyumenskaya oblast\"","RUUD":"Udmurtskaya Respublika","RUULY":"Ul\"yanovskaya oblast\"","RUVGG":"Volgogradskaya oblast\"","RUVLA":"Vladimirskaya oblast\"","RUVLG":"Vologodskaya oblast\"","RUVOR":"Voronezhskaya oblast\"","RUYAN":"Yamalo-Nenetskiy avtonomnyy okrug","RUYAR":"Yaroslavskaya oblast\"","RUYEV":"Yevreyskaya avtonomnaya oblast\"","RUZAB":"Zabaykal\"skiy kray"}}`,
			},
		PolitiesRwanda :
			&PolitiesCode{
				Name: "Rwanda",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Rwanda.svg","phone_code":"+250","states":null}`,
			},
		PolitiesSaintHelenaAscensionandTristandaCunha :
			&PolitiesCode{
				Name: "Saint Helena, Ascension and Tristan da Cunha",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","phone_code":"+290","states":null}`,
			},
		PolitiesSaintKittsandNevis :
			&PolitiesCode{
				Name: "Saint Kitts and Nevis",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Saint_Kitts_and_Nevis.svg","phone_code":"+1869","states":null}`,
			},
		PolitiesSaintLucia :
			&PolitiesCode{
				Name: "Saint Lucia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Saint_Lucia.svg","phone_code":"+1758","states":null}`,
			},
		PolitiesSaintPierreandMiquelon :
			&PolitiesCode{
				Name: "Saint Pierre and Miquelon",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+508","states":null}`,
			},
		PolitiesSaintVincentandtheGrenadines :
			&PolitiesCode{
				Name: "Saint Vincent and the Grenadines",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6d/Flag_of_Saint_Vincent_and_the_Grenadines.svg","phone_code":"+1784","states":null}`,
			},
		PolitiesSaintBarthelemy :
			&PolitiesCode{
				Name: "Saint-Barthelemy",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Saint_Barthelemy_%28local%29.svg","phone_code":"+590","states":null}`,
			},
		PolitiesSaintMartin :
			&PolitiesCode{
				Name: "Saint-Martin",
				Label: "",
				Description: "",
				MetaData: `{"flag":null,"phone_code":"+590","states":null}`,
			},
		PolitiesSamoa :
			&PolitiesCode{
				Name: "Samoa",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Samoa.svg","phone_code":"+685","states":null}`,
			},
		PolitiesSanMarino :
			&PolitiesCode{
				Name: "San Marino",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_San_Marino.svg","phone_code":"+378","states":null}`,
			},
		PolitiesSãoToméandPríncipe :
			&PolitiesCode{
				Name: "São Tomé and Príncipe",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Sao_Tome_and_Principe.svg","phone_code":"+239","states":null}`,
			},
		PolitiesSaudiArabia :
			&PolitiesCode{
				Name: "Saudi Arabia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0d/Flag_of_Saudi_Arabia.svg","phone_code":"+966","states":null}`,
			},
		PolitiesSenegal :
			&PolitiesCode{
				Name: "Senegal",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fd/Flag_of_Senegal.svg","phone_code":"+221","states":null}`,
			},
		PolitiesSerbia :
			&PolitiesCode{
				Name: "Serbia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/ff/Flag_of_Serbia.svg","phone_code":"+381","states":null}`,
			},
		PolitiesSeychelles :
			&PolitiesCode{
				Name: "Seychelles",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Seychelles.svg","phone_code":"+248","states":null}`,
			},
		PolitiesSierraLeone :
			&PolitiesCode{
				Name: "Sierra Leone",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Sierra_Leone.svg","phone_code":"+232","states":null}`,
			},
		PolitiesSingapore :
			&PolitiesCode{
				Name: "Singapore",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg","phone_code":"+65","states":null}`,
			},
		PolitiesSintMaarten :
			&PolitiesCode{
				Name: "Sint Maarten",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Sint_Maarten.svg","phone_code":"+1721","states":null}`,
			},
		PolitiesSlovakia :
			&PolitiesCode{
				Name: "Slovakia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e6/Flag_of_Slovakia.svg","phone_code":"+421","states":null}`,
			},
		PolitiesSlovenia :
			&PolitiesCode{
				Name: "Slovenia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f0/Flag_of_Slovenia.svg","phone_code":"+386","states":null}`,
			},
		PolitiesSolomonIslands :
			&PolitiesCode{
				Name: "Solomon Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Solomon_Islands.svg","phone_code":"+677","states":null}`,
			},
		PolitiesSomalia :
			&PolitiesCode{
				Name: "Somalia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_Somalia.svg","phone_code":"+252","states":null}`,
			},
		PolitiesSouthAfrica :
			&PolitiesCode{
				Name: "South Africa",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/af/Flag_of_South_Africa.svg","phone_code":"+27","states":null}`,
			},
		PolitiesSouthGeorgiaandtheSouthSandwichIslands :
			&PolitiesCode{
				Name: "South Georgia and the South Sandwich Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_South_Georgia_and_the_South_Sandwich_Islands.svg","phone_code":"+500","states":null}`,
			},
		PolitiesSouthKorea :
			&PolitiesCode{
				Name: "South Korea",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg","gov_fiscal_year":"0101-1231","phone_code":"+82","states":null}`,
			},
		PolitiesSouthSudan :
			&PolitiesCode{
				Name: "South Sudan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7a/Flag_of_South_Sudan.svg","phone_code":"+211","states":null}`,
			},
		PolitiesSpain :
			&PolitiesCode{
				Name: "Spain",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9a/Flag_of_Spain.svg","gov_fiscal_year":"0101-1231","phone_code":"+34","states":null}`,
			},
		PolitiesSriLanka :
			&PolitiesCode{
				Name: "Sri Lanka",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Sri_Lanka.svg","phone_code":"+94","states":null}`,
			},
		PolitiesSudan :
			&PolitiesCode{
				Name: "Sudan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Sudan.svg","phone_code":"+249","states":null}`,
			},
		PolitiesSuriname :
			&PolitiesCode{
				Name: "Suriname",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/60/Flag_of_Suriname.svg","phone_code":"+597","states":null}`,
			},
		PolitiesSvalbardandJanMayen :
			&PolitiesCode{
				Name: "Svalbard and Jan Mayen",
				Label: "",
				Description: "",
				MetaData: `{"flag":null,"phone_code":"+47","states":null}`,
			},
		PolitiesSwaziland :
			&PolitiesCode{
				Name: "Swaziland",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fb/Flag_of_Eswatini.svg","phone_code":"+268","states":null}`,
			},
		PolitiesSweden :
			&PolitiesCode{
				Name: "Sweden",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/4/4c/Flag_of_Sweden.svg","gov_fiscal_year":"0101-1231","phone_code":"+46","states":null}`,
			},
		PolitiesSwitzerland :
			&PolitiesCode{
				Name: "Switzerland",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/08/Flag_of_Switzerland_%28Pantone%29.svg","gov_fiscal_year":"0101-1231","phone_code":"+41","states":null}`,
			},
		PolitiesSyria :
			&PolitiesCode{
				Name: "Syria",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/53/Flag_of_Syria.svg","phone_code":"+963","states":null}`,
			},
		PolitiesTaiwan :
			&PolitiesCode{
				Name: "Taiwan",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg","gov_fiscal_year":"0101-1231","phone_code":"+886","states":null}`,
			},
		PolitiesTajikistan :
			&PolitiesCode{
				Name: "Tajikistan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Tajikistan.svg","phone_code":"+992","states":null}`,
			},
		PolitiesTanzania :
			&PolitiesCode{
				Name: "Tanzania",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tanzania.svg","phone_code":"+255","states":null}`,
			},
		PolitiesThailand :
			&PolitiesCode{
				Name: "Thailand",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a9/Flag_of_Thailand.svg","phone_code":"+66","states":null}`,
			},
		PolitiesTogo :
			&PolitiesCode{
				Name: "Togo",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/68/Flag_of_Togo.svg","phone_code":"+228","states":null}`,
			},
		PolitiesTokelau :
			&PolitiesCode{
				Name: "Tokelau",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8e/Flag_of_Tokelau.svg","phone_code":"+690","states":null}`,
			},
		PolitiesTonga :
			&PolitiesCode{
				Name: "Tonga",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Tonga.svg","phone_code":"+676","states":null}`,
			},
		PolitiesTrinidadandTobago :
			&PolitiesCode{
				Name: "Trinidad and Tobago",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Trinidad_and_Tobago.svg","phone_code":"+1868","states":null}`,
			},
		PolitiesTunisia :
			&PolitiesCode{
				Name: "Tunisia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Tunisia.svg","phone_code":"+216","states":null}`,
			},
		PolitiesTurkey :
			&PolitiesCode{
				Name: "Turkey",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Turkey.svg","phone_code":"+90","states":null}`,
			},
		PolitiesTurkmenistan :
			&PolitiesCode{
				Name: "Turkmenistan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Turkmenistan.svg","phone_code":"+993","states":null}`,
			},
		PolitiesTurksandCaicosIslands :
			&PolitiesCode{
				Name: "Turks and Caicos Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_the_Turks_and_Caicos_Islands.svg","phone_code":"+1649","states":null}`,
			},
		PolitiesTuvalu :
			&PolitiesCode{
				Name: "Tuvalu",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tuvalu.svg","phone_code":"+688","states":null}`,
			},
		PolitiesUganda :
			&PolitiesCode{
				Name: "Uganda",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4e/Flag_of_Uganda.svg","phone_code":"+256","states":null}`,
			},
		PolitiesUkraine :
			&PolitiesCode{
				Name: "Ukraine",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Ukraine.svg","phone_code":"+380","states":null}`,
			},
		PolitiesUnitedArabEmirates :
			&PolitiesCode{
				Name: "United Arab Emirates",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_United_Arab_Emirates.svg","states":null}`,
			},
		PolitiesUnitedKingdom :
			&PolitiesCode{
				Name: "United Kingdom",
				Label: "",
				Description: "",
				MetaData: `{"fiscal_year":"0406-0405","flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","gov_fiscal_year":"0401-0331","phone_code":"+44","states":null}`,
			},
		PolitiesUnitedStatesMinorOutlyingIslands :
			&PolitiesCode{
				Name: "United States Minor Outlying Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+246","states":null}`,
			},
		PolitiesUruguay :
			&PolitiesCode{
				Name: "Uruguay",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Uruguay.svg","phone_code":"+598","states":null}`,
			},
		PolitiesUSVirginIslands :
			&PolitiesCode{
				Name: "US Virgin Islands",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_the_United_States_Virgin_Islands.svg","phone_code":"+1340","states":null}`,
			},
		PolitiesUSA :
			&PolitiesCode{
				Name: "USA",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","phone_code":"+1","states":{"USAK":"Alaska","USAL":"Alabama","USAR":"Arkansas","USAS":"American Samoa","USAZ":"Arizona","USCA":"California","USCO":"Colorado","USCT":"Connecticut","USDC":"Washington D.C.","USDE":"Delaware","USFL":"Florida","USGA":"Georgia","USGU":"Guam","USHI":"Hawaii","USIA":"Iowa","USID":"Idaho","USIL":"Illinois","USIN":"Indiana","USKS":"Kansas","USKY":"Kentucky","USLA":"Louisiana","USMA":"Massachusetts","USMD":"Maryland","USME":"Maine","USMI":"Michigan","USMN":"Minnesota","USMO":"Missouri","USMP":"Northern Mariana Islands","USMS":"Mississippi","USMT":"Montana","USNC":"North Carolina","USND":"North Dakota","USNE":"Nebraska","USNH":"New Hampshire","USNJ":"New Jersey","USNM":"New Mexico","USNV":"Nevada","USNY":"New York","USOH":"Ohio","USOK":"Oklahoma","USOR":"Oregon","USPA":"Pennsylvania","USPR":"Puerto Rico","USRI":"Rhode Island","USSC":"South Carolina","USSD":"South Dakota","USTN":"Tennessee","USTX":"Texas","USUM":"United States Minor Outlying Islands","USUT":"Utah","USVA":"Virginia","USVI":"US Virgin Islands","USVT":"Vermont","USWA":"Washington","USWI":"Wisconsin","USWV":"West Virginia","USWY":"Wyoming"}}`,
			},
		PolitiesUzbekistan :
			&PolitiesCode{
				Name: "Uzbekistan",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Uzbekistan.svg","phone_code":"+998","states":null}`,
			},
		PolitiesVanuatu :
			&PolitiesCode{
				Name: "Vanuatu",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_Vanuatu_%28official%29.svg","phone_code":"+678","states":null}`,
			},
		PolitiesVaticanCity :
			&PolitiesCode{
				Name: "Vatican City",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_the_Vatican_City.svg","phone_code":"+418","states":null}`,
			},
		PolitiesVenezuela :
			&PolitiesCode{
				Name: "Venezuela",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7b/Flag_of_Venezuela_%28state%29.svg","phone_code":"+58","states":null}`,
			},
		PolitiesVietnam :
			&PolitiesCode{
				Name: "Vietnam",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Vietnam.svg","phone_code":"+84","states":null}`,
			},
		PolitiesWallisandFutuna :
			&PolitiesCode{
				Name: "Wallis and Futuna",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","phone_code":"+681","states":null}`,
			},
		PolitiesWesternSahara :
			&PolitiesCode{
				Name: "Western Sahara",
				Label: "",
				Description: "",
				MetaData: `{"flag":null,"phone_code":"+212","states":null}`,
			},
		PolitiesYemen :
			&PolitiesCode{
				Name: "Yemen",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Yemen.svg","phone_code":"+967","states":null}`,
			},
		PolitiesZambia :
			&PolitiesCode{
				Name: "Zambia",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/06/Flag_of_Zambia.svg","phone_code":"+260","states":null}`,
			},
		PolitiesZimbabwe :
			&PolitiesCode{
				Name: "Zimbabwe",
				Label: "",
				Description: "",
				MetaData: `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6a/Flag_of_Zimbabwe.svg","phone_code":"+263","states":null}`,
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
	Name string
	Label string
	Description string
	MetaData string
}

// RolesData holds a mapping of Roles codes.
func RolesData(code uint32) *RolesCode {
	switch(code) {
	case RolesAccountant:
		return &RolesCode{
			Name: "Accountant",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesAdvisor:
		return &RolesCode{
			Name: "Advisor",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesAgent:
		return &RolesCode{
			Name: "Agent",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesBeneficiary:
		return &RolesCode{
			Name: "Beneficiary",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesCEO:
		return &RolesCode{
			Name: "CEO",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesCFO:
		return &RolesCode{
			Name: "CFO",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesChair:
		return &RolesCode{
			Name: "Chair",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesCOO:
		return &RolesCode{
			Name: "COO",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesCTO:
		return &RolesCode{
			Name: "CTO",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesCustodian:
		return &RolesCode{
			Name: "Custodian",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesDirector:
		return &RolesCode{
			Name: "Director",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesExecutive:
		return &RolesCode{
			Name: "Executive",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesLawyer:
		return &RolesCode{
			Name: "Lawyer",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesLegalGuardian:
		return &RolesCode{
			Name: "Legal Guardian",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesLimitedPartner:
		return &RolesCode{
			Name: "Limited Partner",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesManager:
		return &RolesCode{
			Name: "Manager",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesManagingPartner:
		return &RolesCode{
			Name: "Managing Partner",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesMember:
		return &RolesCode{
			Name: "Member",
			Label: "",
			Description: "Shareholder",
			MetaData: `{}`,
		}
	case RolesPartner:
		return &RolesCode{
			Name: "Partner",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesPrincipal:
		return &RolesCode{
			Name: "Principal",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesProprietor:
		return &RolesCode{
			Name: "Proprietor",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesProtector:
		return &RolesCode{
			Name: "Protector",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesSecretary:
		return &RolesCode{
			Name: "Secretary",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesSettlor:
		return &RolesCode{
			Name: "Settlor",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesSignificantMember:
		return &RolesCode{
			Name: "Significant Member",
			Label: "",
			Description: "Major Shareholder",
			MetaData: `{}`,
		}
	case RolesSmartContractOperator:
		return &RolesCode{
			Name: "Smart Contract Operator",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesTrader:
		return &RolesCode{
			Name: "Trader",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesTrustee:
		return &RolesCode{
			Name: "Trustee",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case RolesUnitHolder:
		return &RolesCode{
			Name: "Unit Holder",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	default:
		return nil
	}
}

// RolesMap returns a mapping of Roles objects with the code as the key.
func RolesMap() map[uint32]*RolesCode {
	return map[uint32]*RolesCode {
		RolesAccountant :
			&RolesCode{
				Name: "Accountant",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesAdvisor :
			&RolesCode{
				Name: "Advisor",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesAgent :
			&RolesCode{
				Name: "Agent",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesBeneficiary :
			&RolesCode{
				Name: "Beneficiary",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesCEO :
			&RolesCode{
				Name: "CEO",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesCFO :
			&RolesCode{
				Name: "CFO",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesChair :
			&RolesCode{
				Name: "Chair",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesCOO :
			&RolesCode{
				Name: "COO",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesCTO :
			&RolesCode{
				Name: "CTO",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesCustodian :
			&RolesCode{
				Name: "Custodian",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesDirector :
			&RolesCode{
				Name: "Director",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesExecutive :
			&RolesCode{
				Name: "Executive",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesLawyer :
			&RolesCode{
				Name: "Lawyer",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesLegalGuardian :
			&RolesCode{
				Name: "Legal Guardian",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesLimitedPartner :
			&RolesCode{
				Name: "Limited Partner",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesManager :
			&RolesCode{
				Name: "Manager",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesManagingPartner :
			&RolesCode{
				Name: "Managing Partner",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesMember :
			&RolesCode{
				Name: "Member",
				Label: "",
				Description: "Shareholder",
				MetaData: `{}`,
			},
		RolesPartner :
			&RolesCode{
				Name: "Partner",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesPrincipal :
			&RolesCode{
				Name: "Principal",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesProprietor :
			&RolesCode{
				Name: "Proprietor",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesProtector :
			&RolesCode{
				Name: "Protector",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesSecretary :
			&RolesCode{
				Name: "Secretary",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesSettlor :
			&RolesCode{
				Name: "Settlor",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesSignificantMember :
			&RolesCode{
				Name: "Significant Member",
				Label: "",
				Description: "Major Shareholder",
				MetaData: `{}`,
			},
		RolesSmartContractOperator :
			&RolesCode{
				Name: "Smart Contract Operator",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesTrader :
			&RolesCode{
				Name: "Trader",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesTrustee :
			&RolesCode{
				Name: "Trustee",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		RolesUnitHolder :
			&RolesCode{
				Name: "Unit Holder",
				Label: "",
				Description: "",
				MetaData: `{}`,
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

	// ContractExists - The contract already exists and can't be recreated.
	RejectionsContractExists = 10

	// ContractAssetQtyReduction - Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.
	RejectionsContractAssetQtyReduction = 12

	// ContractFixedQuantity - Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.
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

	// AssetCodeExists - The asset code specified already exists and can't be reused.
	RejectionsAssetCodeExists = 20

	// AssetNotFound - The asset code is not found.
	RejectionsAssetNotFound = 21

	// AssetPermissions - The asset permissions prohibit the action requested.
	RejectionsAssetPermissions = 22

	// AssetFrozen - The asset is frozen and the request is not permitted while frozen.
	RejectionsAssetFrozen = 23

	// AssetRevision - The revision in an asset amendment is incorrect.
	RejectionsAssetRevision = 24

	// AssetNotPermitted - Action not permitted by asset.
	RejectionsAssetNotPermitted = 25

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

)

type RejectionsCode struct {
	Name string
	Label string
	Description string
	MetaData string
}

// RejectionsData holds a mapping of Rejections codes.
func RejectionsData(code uint32) *RejectionsCode {
	switch(code) {
	case RejectionsSuccess:
		return &RejectionsCode{
			Name: "Success",
			Label: "Success",
			Description: "No failure. This code should not be used in a reject message.",
			MetaData: `{}`,
		}
	case RejectionsMsgMalformed:
		return &RejectionsCode{
			Name: "MsgMalformed",
			Label: "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.",
			MetaData: `{}`,
		}
	case RejectionsTxMalformed:
		return &RejectionsCode{
			Name: "TxMalformed",
			Label: "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
			MetaData: `{}`,
		}
	case RejectionsTimeout:
		return &RejectionsCode{
			Name: "Timeout",
			Label: "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract's timeout.",
			MetaData: `{}`,
		}
	case RejectionsContractMoved:
		return &RejectionsCode{
			Name: "ContractMoved",
			Label: "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
			MetaData: `{}`,
		}
	case RejectionsDoubleSpend:
		return &RejectionsCode{
			Name: "DoubleSpend",
			Label: "Double Spend",
			Description: "A double spend attempt has been detected.",
			MetaData: `{}`,
		}
	case RejectionsContractExists:
		return &RejectionsCode{
			Name: "ContractExists",
			Label: "Contract Already Exists",
			Description: "The contract already exists and can't be recreated.",
			MetaData: `{}`,
		}
	case RejectionsContractAssetQtyReduction:
		return &RejectionsCode{
			Name: "ContractAssetQtyReduction",
			Label: "Contract Asset Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
			MetaData: `{}`,
		}
	case RejectionsContractFixedQuantity:
		return &RejectionsCode{
			Name: "ContractFixedQuantity",
			Label: "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
			MetaData: `{}`,
		}
	case RejectionsContractPermissions:
		return &RejectionsCode{
			Name: "ContractPermissions",
			Label: "Contract Permissions Prohibit",
			Description: "The contract permissions prohibit the action requested.",
			MetaData: `{}`,
		}
	case RejectionsContractExpired:
		return &RejectionsCode{
			Name: "ContractExpired",
			Label: "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
			MetaData: `{}`,
		}
	case RejectionsContractFrozen:
		return &RejectionsCode{
			Name: "ContractFrozen",
			Label: "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
			MetaData: `{}`,
		}
	case RejectionsContractRevision:
		return &RejectionsCode{
			Name: "ContractRevision",
			Label: "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
			MetaData: `{}`,
		}
	case RejectionsContractNotPermitted:
		return &RejectionsCode{
			Name: "ContractNotPermitted",
			Label: "Contract Not Permitted",
			Description: "Action not permitted by contract.",
			MetaData: `{}`,
		}
	case RejectionsContractBothOperatorsRequired:
		return &RejectionsCode{
			Name: "ContractBothOperatorsRequired",
			Label: "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
			MetaData: `{}`,
		}
	case RejectionsAssetCodeExists:
		return &RejectionsCode{
			Name: "AssetCodeExists",
			Label: "Asset Code Already Exists",
			Description: "The asset code specified already exists and can't be reused.",
			MetaData: `{}`,
		}
	case RejectionsAssetNotFound:
		return &RejectionsCode{
			Name: "AssetNotFound",
			Label: "Asset Not Found",
			Description: "The asset code is not found.",
			MetaData: `{}`,
		}
	case RejectionsAssetPermissions:
		return &RejectionsCode{
			Name: "AssetPermissions",
			Label: "Asset Permissions Prohibit",
			Description: "The asset permissions prohibit the action requested.",
			MetaData: `{}`,
		}
	case RejectionsAssetFrozen:
		return &RejectionsCode{
			Name: "AssetFrozen",
			Label: "Asset Frozen",
			Description: "The asset is frozen and the request is not permitted while frozen.",
			MetaData: `{}`,
		}
	case RejectionsAssetRevision:
		return &RejectionsCode{
			Name: "AssetRevision",
			Label: "Asset Revision Incorrect",
			Description: "The revision in an asset amendment is incorrect.",
			MetaData: `{}`,
		}
	case RejectionsAssetNotPermitted:
		return &RejectionsCode{
			Name: "AssetNotPermitted",
			Label: "Asset Not Permitted",
			Description: "Action not permitted by asset.",
			MetaData: `{}`,
		}
	case RejectionsTransferSelf:
		return &RejectionsCode{
			Name: "TransferSelf",
			Label: "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
			MetaData: `{}`,
		}
	case RejectionsTransferExpired:
		return &RejectionsCode{
			Name: "TransferExpired",
			Label: "Transfer Expired",
			Description: "The transfer has expired.",
			MetaData: `{}`,
		}
	case RejectionsHoldingsFrozen:
		return &RejectionsCode{
			Name: "HoldingsFrozen",
			Label: "Holdings Frozen",
			Description: "Holdings are frozen, so the request can't be completed.",
			MetaData: `{}`,
		}
	case RejectionsHoldingsLocked:
		return &RejectionsCode{
			Name: "HoldingsLocked",
			Label: "Holdings Locked",
			Description: "Holdings are locked by a multi-contract request, so the request can't be completed yet.",
			MetaData: `{}`,
		}
	case RejectionsHolderProposalProhibited:
		return &RejectionsCode{
			Name: "HolderProposalProhibited",
			Label: "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
			MetaData: `{}`,
		}
	case RejectionsProposalConflicts:
		return &RejectionsCode{
			Name: "ProposalConflicts",
			Label: "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
			MetaData: `{}`,
		}
	case RejectionsVoteNotFound:
		return &RejectionsCode{
			Name: "VoteNotFound",
			Label: "Vote Not Found",
			Description: "The vote ID referenced is not found.",
			MetaData: `{}`,
		}
	case RejectionsVoteClosed:
		return &RejectionsCode{
			Name: "VoteClosed",
			Label: "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
			MetaData: `{}`,
		}
	case RejectionsBallotAlreadyCounted:
		return &RejectionsCode{
			Name: "BallotAlreadyCounted",
			Label: "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
			MetaData: `{}`,
		}
	case RejectionsVoteSystemNotPermitted:
		return &RejectionsCode{
			Name: "VoteSystemNotPermitted",
			Label: "Vote System Not Permitted",
			Description: "The voting system isn't permitted for this request.",
			MetaData: `{}`,
		}
	case RejectionsInsufficientTxFeeFunding:
		return &RejectionsCode{
			Name: "InsufficientTxFeeFunding",
			Label: "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
			MetaData: `{}`,
		}
	case RejectionsInsufficientValue:
		return &RejectionsCode{
			Name: "InsufficientValue",
			Label: "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
			MetaData: `{}`,
		}
	case RejectionsInsufficientQuantity:
		return &RejectionsCode{
			Name: "InsufficientQuantity",
			Label: "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
			MetaData: `{}`,
		}
	case RejectionsNotAdministration:
		return &RejectionsCode{
			Name: "NotAdministration",
			Label: "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
			MetaData: `{}`,
		}
	case RejectionsNotOperator:
		return &RejectionsCode{
			Name: "NotOperator",
			Label: "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
			MetaData: `{}`,
		}
	case RejectionsUnauthorizedAddress:
		return &RejectionsCode{
			Name: "UnauthorizedAddress",
			Label: "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
			MetaData: `{}`,
		}
	case RejectionsInvalidSignature:
		return &RejectionsCode{
			Name: "InvalidSignature",
			Label: "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
			MetaData: `{}`,
		}
	default:
		return nil
	}
}

// RejectionsMap returns a mapping of Rejections objects with the code as the key.
func RejectionsMap() map[uint32]*RejectionsCode {
	return map[uint32]*RejectionsCode {
		RejectionsSuccess :
			&RejectionsCode{
				Name: "Success",
				Label: "Success",
				Description: "No failure. This code should not be used in a reject message.",
				MetaData: `{}`,
			},
		RejectionsMsgMalformed :
			&RejectionsCode{
				Name: "MsgMalformed",
				Label: "Message Malformed",
				Description: "The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.",
				MetaData: `{}`,
			},
		RejectionsTxMalformed :
			&RejectionsCode{
				Name: "TxMalformed",
				Label: "Transaction Malformed",
				Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
				MetaData: `{}`,
			},
		RejectionsTimeout :
			&RejectionsCode{
				Name: "Timeout",
				Label: "Time Out",
				Description: "A dependency, other contract/service, has failed to complete before the smart contract's timeout.",
				MetaData: `{}`,
			},
		RejectionsContractMoved :
			&RejectionsCode{
				Name: "ContractMoved",
				Label: "Contract Moved",
				Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
				MetaData: `{}`,
			},
		RejectionsDoubleSpend :
			&RejectionsCode{
				Name: "DoubleSpend",
				Label: "Double Spend",
				Description: "A double spend attempt has been detected.",
				MetaData: `{}`,
			},
		RejectionsContractExists :
			&RejectionsCode{
				Name: "ContractExists",
				Label: "Contract Already Exists",
				Description: "The contract already exists and can't be recreated.",
				MetaData: `{}`,
			},
		RejectionsContractAssetQtyReduction :
			&RejectionsCode{
				Name: "ContractAssetQtyReduction",
				Label: "Contract Asset Quantity Reduction",
				Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
				MetaData: `{}`,
			},
		RejectionsContractFixedQuantity :
			&RejectionsCode{
				Name: "ContractFixedQuantity",
				Label: "Contract Fixed Quantity",
				Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
				MetaData: `{}`,
			},
		RejectionsContractPermissions :
			&RejectionsCode{
				Name: "ContractPermissions",
				Label: "Contract Permissions Prohibit",
				Description: "The contract permissions prohibit the action requested.",
				MetaData: `{}`,
			},
		RejectionsContractExpired :
			&RejectionsCode{
				Name: "ContractExpired",
				Label: "Contract Expired",
				Description: "The contract is expired so can no longer accept requests.",
				MetaData: `{}`,
			},
		RejectionsContractFrozen :
			&RejectionsCode{
				Name: "ContractFrozen",
				Label: "Contract Frozen",
				Description: "The contract is frozen and the request is not permitted while frozen.",
				MetaData: `{}`,
			},
		RejectionsContractRevision :
			&RejectionsCode{
				Name: "ContractRevision",
				Label: "Contract Revision Incorrect",
				Description: "The revision in a contract amendment is incorrect.",
				MetaData: `{}`,
			},
		RejectionsContractNotPermitted :
			&RejectionsCode{
				Name: "ContractNotPermitted",
				Label: "Contract Not Permitted",
				Description: "Action not permitted by contract.",
				MetaData: `{}`,
			},
		RejectionsContractBothOperatorsRequired :
			&RejectionsCode{
				Name: "ContractBothOperatorsRequired",
				Label: "Contract BothOperatorsRequired",
				Description: "Both operators signatures are required to perform this action.",
				MetaData: `{}`,
			},
		RejectionsAssetCodeExists :
			&RejectionsCode{
				Name: "AssetCodeExists",
				Label: "Asset Code Already Exists",
				Description: "The asset code specified already exists and can't be reused.",
				MetaData: `{}`,
			},
		RejectionsAssetNotFound :
			&RejectionsCode{
				Name: "AssetNotFound",
				Label: "Asset Not Found",
				Description: "The asset code is not found.",
				MetaData: `{}`,
			},
		RejectionsAssetPermissions :
			&RejectionsCode{
				Name: "AssetPermissions",
				Label: "Asset Permissions Prohibit",
				Description: "The asset permissions prohibit the action requested.",
				MetaData: `{}`,
			},
		RejectionsAssetFrozen :
			&RejectionsCode{
				Name: "AssetFrozen",
				Label: "Asset Frozen",
				Description: "The asset is frozen and the request is not permitted while frozen.",
				MetaData: `{}`,
			},
		RejectionsAssetRevision :
			&RejectionsCode{
				Name: "AssetRevision",
				Label: "Asset Revision Incorrect",
				Description: "The revision in an asset amendment is incorrect.",
				MetaData: `{}`,
			},
		RejectionsAssetNotPermitted :
			&RejectionsCode{
				Name: "AssetNotPermitted",
				Label: "Asset Not Permitted",
				Description: "Action not permitted by asset.",
				MetaData: `{}`,
			},
		RejectionsTransferSelf :
			&RejectionsCode{
				Name: "TransferSelf",
				Label: "Transfer To Self Prohibited",
				Description: "Transfers with the sender and receiver addresses the same are not permitted.",
				MetaData: `{}`,
			},
		RejectionsTransferExpired :
			&RejectionsCode{
				Name: "TransferExpired",
				Label: "Transfer Expired",
				Description: "The transfer has expired.",
				MetaData: `{}`,
			},
		RejectionsHoldingsFrozen :
			&RejectionsCode{
				Name: "HoldingsFrozen",
				Label: "Holdings Frozen",
				Description: "Holdings are frozen, so the request can't be completed.",
				MetaData: `{}`,
			},
		RejectionsHoldingsLocked :
			&RejectionsCode{
				Name: "HoldingsLocked",
				Label: "Holdings Locked",
				Description: "Holdings are locked by a multi-contract request, so the request can't be completed yet.",
				MetaData: `{}`,
			},
		RejectionsHolderProposalProhibited :
			&RejectionsCode{
				Name: "HolderProposalProhibited",
				Label: "Holder Proposal Prohibited",
				Description: "Holders are not permitted to make proposals.",
				MetaData: `{}`,
			},
		RejectionsProposalConflicts :
			&RejectionsCode{
				Name: "ProposalConflicts",
				Label: "Proposal Conflicts",
				Description: "The proposal conflicts with an unapplied proposal.",
				MetaData: `{}`,
			},
		RejectionsVoteNotFound :
			&RejectionsCode{
				Name: "VoteNotFound",
				Label: "Vote Not Found",
				Description: "The vote ID referenced is not found.",
				MetaData: `{}`,
			},
		RejectionsVoteClosed :
			&RejectionsCode{
				Name: "VoteClosed",
				Label: "Vote Closed",
				Description: "The vote has closed and ballots are no longer permitted.",
				MetaData: `{}`,
			},
		RejectionsBallotAlreadyCounted :
			&RejectionsCode{
				Name: "BallotAlreadyCounted",
				Label: "Ballot Already Counted",
				Description: "The ballot has already been counted for this address.",
				MetaData: `{}`,
			},
		RejectionsVoteSystemNotPermitted :
			&RejectionsCode{
				Name: "VoteSystemNotPermitted",
				Label: "Vote System Not Permitted",
				Description: "The voting system isn't permitted for this request.",
				MetaData: `{}`,
			},
		RejectionsInsufficientTxFeeFunding :
			&RejectionsCode{
				Name: "InsufficientTxFeeFunding",
				Label: "Insufficient Transaction Fee Funding",
				Description: "Insufficient bitcoin quantities for response transaction fees.",
				MetaData: `{}`,
			},
		RejectionsInsufficientValue :
			&RejectionsCode{
				Name: "InsufficientValue",
				Label: "Insufficient Value",
				Description: "Insufficient bitcoin quantity in inputs to fund request.",
				MetaData: `{}`,
			},
		RejectionsInsufficientQuantity :
			&RejectionsCode{
				Name: "InsufficientQuantity",
				Label: "Insufficient Quantity",
				Description: "Insufficient token holdings to for request.",
				MetaData: `{}`,
			},
		RejectionsNotAdministration :
			&RejectionsCode{
				Name: "NotAdministration",
				Label: "Requestor Is Not Administration",
				Description: "The requestor is not the administration and is required to be for this request.",
				MetaData: `{}`,
			},
		RejectionsNotOperator :
			&RejectionsCode{
				Name: "NotOperator",
				Label: "Requestor Is Not Operator",
				Description: "The requestor is not the operator and is required to be for this request.",
				MetaData: `{}`,
			},
		RejectionsUnauthorizedAddress :
			&RejectionsCode{
				Name: "UnauthorizedAddress",
				Label: "Unauthorized Address",
				Description: "The address specified is not permitted for this request.",
				MetaData: `{}`,
			},
		RejectionsInvalidSignature :
			&RejectionsCode{
				Name: "InvalidSignature",
				Label: "Invalid Signature",
				Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
				MetaData: `{}`,
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
	Name string
	Label string
	Description string
	MetaData string
}

// TagsData holds a mapping of Tags codes.
func TagsData(code uint32) *TagsCode {
	switch(code) {
	case TagsHousing:
		return &TagsCode{
			Name: "Housing",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsUtilities:
		return &TagsCode{
			Name: "Utilities",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsFood:
		return &TagsCode{
			Name: "Food",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsMedical:
		return &TagsCode{
			Name: "Medical",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsFinancialServices:
		return &TagsCode{
			Name: "Financial Services",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsEntertainment:
		return &TagsCode{
			Name: "Entertainment",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsSales:
		return &TagsCode{
			Name: "Sales",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsAutomotive:
		return &TagsCode{
			Name: "Automotive",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsTransportation:
		return &TagsCode{
			Name: "Transportation",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsFitness:
		return &TagsCode{
			Name: "Fitness",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsElectricity:
		return &TagsCode{
			Name: "Electricity",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsWater:
		return &TagsCode{
			Name: "Water",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsInternet:
		return &TagsCode{
			Name: "Internet",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsMedicine:
		return &TagsCode{
			Name: "Medicine",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsService:
		return &TagsCode{
			Name: "Service",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsRepair:
		return &TagsCode{
			Name: "Repair",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsSupplies:
		return &TagsCode{
			Name: "Supplies",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsParts:
		return &TagsCode{
			Name: "Parts",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsLabor:
		return &TagsCode{
			Name: "Labor",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsTip:
		return &TagsCode{
			Name: "Tip",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsMedia:
		return &TagsCode{
			Name: "Media",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsMusic:
		return &TagsCode{
			Name: "Music",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsVideo:
		return &TagsCode{
			Name: "Video",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsPhoto:
		return &TagsCode{
			Name: "Photo",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsAudio:
		return &TagsCode{
			Name: "Audio",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsAlcohol:
		return &TagsCode{
			Name: "Alcohol",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsTobacco:
		return &TagsCode{
			Name: "Tobacco",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsDiscounted:
		return &TagsCode{
			Name: "Discounted",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	case TagsPromotional:
		return &TagsCode{
			Name: "Promotional",
			Label: "",
			Description: "",
			MetaData: `{}`,
		}
	default:
		return nil
	}
}

// TagsMap returns a mapping of Tags objects with the code as the key.
func TagsMap() map[uint32]*TagsCode {
	return map[uint32]*TagsCode {
		TagsHousing :
			&TagsCode{
				Name: "Housing",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsUtilities :
			&TagsCode{
				Name: "Utilities",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsFood :
			&TagsCode{
				Name: "Food",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsMedical :
			&TagsCode{
				Name: "Medical",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsFinancialServices :
			&TagsCode{
				Name: "Financial Services",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsEntertainment :
			&TagsCode{
				Name: "Entertainment",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsSales :
			&TagsCode{
				Name: "Sales",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsAutomotive :
			&TagsCode{
				Name: "Automotive",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsTransportation :
			&TagsCode{
				Name: "Transportation",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsFitness :
			&TagsCode{
				Name: "Fitness",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsElectricity :
			&TagsCode{
				Name: "Electricity",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsWater :
			&TagsCode{
				Name: "Water",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsInternet :
			&TagsCode{
				Name: "Internet",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsMedicine :
			&TagsCode{
				Name: "Medicine",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsService :
			&TagsCode{
				Name: "Service",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsRepair :
			&TagsCode{
				Name: "Repair",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsSupplies :
			&TagsCode{
				Name: "Supplies",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsParts :
			&TagsCode{
				Name: "Parts",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsLabor :
			&TagsCode{
				Name: "Labor",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsTip :
			&TagsCode{
				Name: "Tip",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsMedia :
			&TagsCode{
				Name: "Media",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsMusic :
			&TagsCode{
				Name: "Music",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsVideo :
			&TagsCode{
				Name: "Video",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsPhoto :
			&TagsCode{
				Name: "Photo",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsAudio :
			&TagsCode{
				Name: "Audio",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsAlcohol :
			&TagsCode{
				Name: "Alcohol",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsTobacco :
			&TagsCode{
				Name: "Tobacco",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsDiscounted :
			&TagsCode{
				Name: "Discounted",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
		TagsPromotional :
			&TagsCode{
				Name: "Promotional",
				Label: "",
				Description: "",
				MetaData: `{}`,
			},
	}
}

