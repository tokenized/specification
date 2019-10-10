package actions

/************************************ Entities ************************************/
const (
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
		EntitiesIndividual: &EntitiesCode{
			Name:        "Individual",
			Label:       "Individual",
			Description: "",
		},
		EntitiesPublicCompany: &EntitiesCode{
			Name:        "PublicCompany",
			Label:       "Public Company Limited by Shares",
			Description: "",
		},
		EntitiesPrivateCompany: &EntitiesCode{
			Name:        "PrivateCompany",
			Label:       "Private Company Limited by Shares",
			Description: "",
		},
		EntitiesLimitedPartnership: &EntitiesCode{
			Name:        "LimitedPartnership",
			Label:       "Limited Partnership",
			Description: "",
		},
		EntitiesUnlimitedPartnership: &EntitiesCode{
			Name:        "UnlimitedPartnership",
			Label:       "Unlimited Partnership",
			Description: "",
		},
		EntitiesSoleProprietorship: &EntitiesCode{
			Name:        "SoleProprietorship",
			Label:       "Sole Proprietorship",
			Description: "",
		},
		EntitiesStatutoryCompany: &EntitiesCode{
			Name:        "StatutoryCompany",
			Label:       "Statutory Company",
			Description: "",
		},
		EntitiesNonProfitOrganization: &EntitiesCode{
			Name:        "NonProfitOrganization",
			Label:       "Non-Profit Organization",
			Description: "",
		},
		EntitiesNationState: &EntitiesCode{
			Name:        "NationState",
			Label:       "Nation State",
			Description: "",
		},
		EntitiesGovernmentAgency: &EntitiesCode{
			Name:        "GovernmentAgency",
			Label:       "Government Agency",
			Description: "",
		},
		EntitiesUnitTrust: &EntitiesCode{
			Name:        "UnitTrust",
			Label:       "Unit Trust",
			Description: "",
		},
		EntitiesDiscretionaryTrust: &EntitiesCode{
			Name:        "DiscretionaryTrust",
			Label:       "Discretionary Trust",
			Description: "",
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

	// Curacao -
	PolitiesCuracao = "CUW"

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

	// Ivory Coast -
	PolitiesIvoryCoast = "CIV"

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
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/52/Flag_of_%C3%85land.svg","states":null}`,
		}
	case PolitiesAfghanistan:
		return &PolitiesCode{
			Name:        "Afghanistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Afghanistan.svg","states":null}`,
		}
	case PolitiesAlbania:
		return &PolitiesCode{
			Name:        "Albania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/36/Flag_of_Albania.svg","states":null}`,
		}
	case PolitiesAlgeria:
		return &PolitiesCode{
			Name:        "Algeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Algeria.svg","states":null}`,
		}
	case PolitiesAmericanSamoa:
		return &PolitiesCode{
			Name:        "American Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/87/Flag_of_American_Samoa.svg","states":null}`,
		}
	case PolitiesAndorra:
		return &PolitiesCode{
			Name:        "Andorra",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Andorra.svg","states":null}`,
		}
	case PolitiesAngola:
		return &PolitiesCode{
			Name:        "Angola",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9d/Flag_of_Angola.svg","states":null}`,
		}
	case PolitiesAnguilla:
		return &PolitiesCode{
			Name:        "Anguilla",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Anguilla.svg","states":null}`,
		}
	case PolitiesAntarctica:
		return &PolitiesCode{
			Name:        "Antarctica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"states":null}`,
		}
	case PolitiesAntiguaandBarbuda:
		return &PolitiesCode{
			Name:        "Antigua and Barbuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Antigua_and_Barbuda.svg","states":null}`,
		}
	case PolitiesArgentina:
		return &PolitiesCode{
			Name:        "Argentina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1a/Flag_of_Argentina.svg","states":null}`,
		}
	case PolitiesArmenia:
		return &PolitiesCode{
			Name:        "Armenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2f/Flag_of_Armenia.svg","states":null}`,
		}
	case PolitiesAruba:
		return &PolitiesCode{
			Name:        "Aruba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Aruba.svg","states":null}`,
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
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/b/b9/Flag_of_Australia.svg","gov_fiscal_year":"0701-0630","states":{"AUACT":"Australian Capital Territory","AUCC":"Cocos (Keening) Island","AUCX":"Christmas Island","AUHM":"Heard Island and McDonalds Islands","AUJBT":"Jervis Bay Territory","AUNF":"Norfolk Island","AUNSW":"New South Wales","AUNT":"Northern Territory","AUQLD":"Queensland","AUSA":"South Australia","AUTAS":"Tasmania","AUVIC":"Victoria","AUWA":"Western Australia"}}`,
		}
	case PolitiesAustria:
		return &PolitiesCode{
			Name:        "Austria",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/41/Flag_of_Austria.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesAzerbaijan:
		return &PolitiesCode{
			Name:        "Azerbaijan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Azerbaijan.svg","states":null}`,
		}
	case PolitiesTheBahamas:
		return &PolitiesCode{
			Name:        "The Bahamas",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/93/Flag_of_the_Bahamas.svg","states":null}`,
		}
	case PolitiesBahrain:
		return &PolitiesCode{
			Name:        "Bahrain",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Bahrain.svg","states":null}`,
		}
	case PolitiesBangladesh:
		return &PolitiesCode{
			Name:        "Bangladesh",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/f9/Flag_of_Bangladesh.svg","gov_fiscal_year":"0701-0630","states":null}`,
		}
	case PolitiesBarbados:
		return &PolitiesCode{
			Name:        "Barbados",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ef/Flag_of_Barbados.svg","states":null}`,
		}
	case PolitiesBelarus:
		return &PolitiesCode{
			Name:        "Belarus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/85/Flag_of_Belarus.svg","states":null}`,
		}
	case PolitiesBelgium:
		return &PolitiesCode{
			Name:        "Belgium",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Belgium.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesBelize:
		return &PolitiesCode{
			Name:        "Belize",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e7/Flag_of_Belize.svg","states":null}`,
		}
	case PolitiesBenin:
		return &PolitiesCode{
			Name:        "Benin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Benin.svg","states":null}`,
		}
	case PolitiesBermuda:
		return &PolitiesCode{
			Name:        "Bermuda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bermuda.svg","states":null}`,
		}
	case PolitiesBhutan:
		return &PolitiesCode{
			Name:        "Bhutan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/91/Flag_of_Bhutan.svg","states":null}`,
		}
	case PolitiesBolivia:
		return &PolitiesCode{
			Name:        "Bolivia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/de/Flag_of_Bolivia_%28state%29.svg","states":null}`,
		}
	case PolitiesBonaireStEustasuisandSaba:
		return &PolitiesCode{
			Name:        "Bonaire, St Eustasuis and Saba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","states":null}`,
		}
	case PolitiesBosniaandHerzegovina:
		return &PolitiesCode{
			Name:        "Bosnia and Herzegovina",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bf/Flag_of_Bosnia_and_Herzegovina.svg","states":null}`,
		}
	case PolitiesBotswana:
		return &PolitiesCode{
			Name:        "Botswana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Botswana.svg","states":null}`,
		}
	case PolitiesBouvetIsland:
		return &PolitiesCode{
			Name:        "Bouvet Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","states":null}`,
		}
	case PolitiesBrazil:
		return &PolitiesCode{
			Name:        "Brazil",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg","gov_fiscal_year":"0101-1231","states":{"BRAC":"Acre","BRAL":"Alagoas","BRAM":"Amazonas","BRAP":"Amapá","BRBA":"Bahia","BRCE":"Ceará","BRDF":"Federal District","BRES":"Espírito Santo","BRGO":"Goiás","BRMA":"Maranhão","BRMG":"Minas Gerais","BRMS":"Mato Grosso do Sul","BRMT":"Mato Grosso","BRPA":"Pará","BRPB":"Paraíba","BRPE":"Pernambuco","BRPI":"Piauí","BRPR":"Paraná","BRRJ":"Rio de Janeiro","BRRN":"Rio Grande do Norte","BRRO":"Rondônia","BRRR":"Roraima","BRRS":"Rio Grande do Sul","BRSC":"Santa Catarina","BRSE":"Sergipe","BRSP":"São Paulo","BRTO":"Tocantins"}}`,
		}
	case PolitiesBritishIndianOceanTerritory:
		return &PolitiesCode{
			Name:        "British Indian Ocean Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_the_British_Indian_Ocean_Territory.svg","states":null}`,
		}
	case PolitiesBritishVirginIslands:
		return &PolitiesCode{
			Name:        "British Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/42/Flag_of_the_British_Virgin_Islands.svg","states":null}`,
		}
	case PolitiesBrunei:
		return &PolitiesCode{
			Name:        "Brunei",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Brunei.svg","states":null}`,
		}
	case PolitiesBulgaria:
		return &PolitiesCode{
			Name:        "Bulgaria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Bulgaria.svg","states":null}`,
		}
	case PolitiesBurkinaFaso:
		return &PolitiesCode{
			Name:        "Burkina Faso",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Burkina_Faso.svg","states":null}`,
		}
	case PolitiesBurundi:
		return &PolitiesCode{
			Name:        "Burundi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/50/Flag_of_Burundi.svg","states":null}`,
		}
	case PolitiesCambodia:
		return &PolitiesCode{
			Name:        "Cambodia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_Cambodia.svg","states":null}`,
		}
	case PolitiesCameroon:
		return &PolitiesCode{
			Name:        "Cameroon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Cameroon.svg","states":null}`,
		}
	case PolitiesCanada:
		return &PolitiesCode{
			Name:        "Canada",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg","gov_fiscal_year":"0401-0331","states":{"CAAB":"Alberta","CABC":"British Columbia","CAMB":"Manitoba","CANB":"New Brunswick","CANL":"Newfoundland and Labrador","CANS":"Nova Scotia","CANT":"Northwest Territories","CANU":"Nunavut","CAON":"Ontario","CAPE":"Prince Edward Island","CAQC":"Quebec","CASK":"Saskatchewan","CAYT":"Yukon"}}`,
		}
	case PolitiesCapeVerde:
		return &PolitiesCode{
			Name:        "Cape Verde",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Cape_Verde.svg","states":null}`,
		}
	case PolitiesCaymanIslands:
		return &PolitiesCode{
			Name:        "Cayman Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_the_Cayman_Islands.svg","states":null}`,
		}
	case PolitiesCentralAfricanRepublic:
		return &PolitiesCode{
			Name:        "Central African Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Central_African_Republic.svg","states":null}`,
		}
	case PolitiesChad:
		return &PolitiesCode{
			Name:        "Chad",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4b/Flag_of_Chad.svg","states":null}`,
		}
	case PolitiesChile:
		return &PolitiesCode{
			Name:        "Chile",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/78/Flag_of_Chile.svg","states":null}`,
		}
	case PolitiesChina:
		return &PolitiesCode{
			Name:        "China",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_the_People%27s_Republic_of_China.svg","gov_fiscal_year":"0101-1231","states":{"CNAH":"Anhui","CNBJ":"Beijing","CNCQ":"Chongqing","CNFJ":"Fujian","CNGD":"Guangdong","CNGS":"Gansu","CNGX":"Guangxi","CNGZ":"Guizhou","CNHA":"Henan","CNHB":"Hubei","CNHE":"Hebei","CNHI":"Hainan","CNHK":"Hong Kong (Xianggang)","CNHL":"Heilongjiang","CNHN":"Hunan","CNJL":"Jilin","CNJS":"Jiangsu","CNJX":"Jiangxi","CNLN":"Liaoning","CNMC":"Macao (Aomen)","CNNM":"Nei Mongol (mn)","CNNX":"Ningxia","CNQH":"Qinghai","CNSC":"Sichuan","CNSD":"Shandong","CNSH":"Shanghai","CNSN":"Shaanxi","CNSX":"Shanxi","CNTJ":"Tianjin","CNTW":"Taiwan","CNXJ":"Xinjiang","CNXZ":"Xizang","CNYN":"Yunnan","CNZJ":"Zhejiang"}}`,
		}
	case PolitiesChristmasIsland:
		return &PolitiesCode{
			Name:        "Christmas Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/67/Flag_of_Christmas_Island.svg","states":null}`,
		}
	case PolitiesCocosIslands:
		return &PolitiesCode{
			Name:        "Cocos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Cocos_%28Keeling%29_Islands.svg","states":null}`,
		}
	case PolitiesColombia:
		return &PolitiesCode{
			Name:        "Colombia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Colombia.svg","states":null}`,
		}
	case PolitiesComoros:
		return &PolitiesCode{
			Name:        "Comoros",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/94/Flag_of_the_Comoros.svg","states":null}`,
		}
	case PolitiesCongoBrazzaville:
		return &PolitiesCode{
			Name:        "Congo-Brazzaville",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_the_Republic_of_the_Congo.svg","states":null}`,
		}
	case PolitiesCongoKinshasa:
		return &PolitiesCode{
			Name:        "Congo-Kinshasa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6f/Flag_of_the_Democratic_Republic_of_the_Congo.svg","states":null}`,
		}
	case PolitiesCookIslands:
		return &PolitiesCode{
			Name:        "Cook Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/35/Flag_of_the_Cook_Islands.svg","states":null}`,
		}
	case PolitiesCostaRica:
		return &PolitiesCode{
			Name:        "Costa Rica",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"1001-0931","flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Costa_Rica_%28state%29.svg","gov_fiscal_year":"1001-0931","states":null}`,
		}
	case PolitiesCroatia:
		return &PolitiesCode{
			Name:        "Croatia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Croatia.svg","states":null}`,
		}
	case PolitiesCuba:
		return &PolitiesCode{
			Name:        "Cuba",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bd/Flag_of_Cuba.svg","states":null}`,
		}
	case PolitiesCuracao:
		return &PolitiesCode{
			Name:        "Curacao",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_Cura%C3%A7ao.svg","states":null}`,
		}
	case PolitiesCyprus:
		return &PolitiesCode{
			Name:        "Cyprus",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Cyprus.svg","states":null}`,
		}
	case PolitiesCzechRepublic:
		return &PolitiesCode{
			Name:        "Czech Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_Czech_Republic.svg","states":null}`,
		}
	case PolitiesDenmark:
		return &PolitiesCode{
			Name:        "Denmark",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9c/Flag_of_Denmark.svg","states":null}`,
		}
	case PolitiesDjibouti:
		return &PolitiesCode{
			Name:        "Djibouti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_Djibouti.svg","states":null}`,
		}
	case PolitiesDominica:
		return &PolitiesCode{
			Name:        "Dominica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c4/Flag_of_Dominica.svg","states":null}`,
		}
	case PolitiesDominicanRepublic:
		return &PolitiesCode{
			Name:        "Dominican Republic",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_the_Dominican_Republic.svg","states":null}`,
		}
	case PolitiesEastTimor:
		return &PolitiesCode{
			Name:        "East Timor",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/26/Flag_of_East_Timor.svg","states":null}`,
		}
	case PolitiesEcuador:
		return &PolitiesCode{
			Name:        "Ecuador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e8/Flag_of_Ecuador.svg","states":null}`,
		}
	case PolitiesEgypt:
		return &PolitiesCode{
			Name:        "Egypt",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0701-0630","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Egypt.svg","gov_fiscal_year":"0701-0630","states":null}`,
		}
	case PolitiesElSalvador:
		return &PolitiesCode{
			Name:        "El Salvador",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/34/Flag_of_El_Salvador.svg","states":null}`,
		}
	case PolitiesEquatorialGuinea:
		return &PolitiesCode{
			Name:        "Equatorial Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Equatorial_Guinea.svg","states":null}`,
		}
	case PolitiesEritrea:
		return &PolitiesCode{
			Name:        "Eritrea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/29/Flag_of_Eritrea.svg","states":null}`,
		}
	case PolitiesEstonia:
		return &PolitiesCode{
			Name:        "Estonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8f/Flag_of_Estonia.svg","states":null}`,
		}
	case PolitiesEthiopia:
		return &PolitiesCode{
			Name:        "Ethiopia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0708-0707","flag":"https://upload.wikimedia.org/wikipedia/commons/7/71/Flag_of_Ethiopia.svg","gov_fiscal_year":"0708-0707","states":null}`,
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
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/83/Flag_of_the_Falkland_Islands.svg","states":null}`,
		}
	case PolitiesFaroeIslands:
		return &PolitiesCode{
			Name:        "Faroe Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3c/Flag_of_the_Faroe_Islands.svg","states":null}`,
		}
	case PolitiesFiji:
		return &PolitiesCode{
			Name:        "Fiji",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/ba/Flag_of_Fiji.svg","states":null}`,
		}
	case PolitiesFinland:
		return &PolitiesCode{
			Name:        "Finland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg","states":null}`,
		}
	case PolitiesFrance:
		return &PolitiesCode{
			Name:        "France",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesFrenchGuiana:
		return &PolitiesCode{
			Name:        "French Guiana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesFrenchPolynesia:
		return &PolitiesCode{
			Name:        "French Polynesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/db/Flag_of_French_Polynesia.svg","states":null}`,
		}
	case PolitiesFrenchSouthernandAntarcticLands:
		return &PolitiesCode{
			Name:        "French Southern and Antarctic Lands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a7/Flag_of_the_French_Southern_and_Antarctic_Lands.svg","states":null}`,
		}
	case PolitiesGabon:
		return &PolitiesCode{
			Name:        "Gabon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/04/Flag_of_Gabon.svg","states":null}`,
		}
	case PolitiesTheGambia:
		return &PolitiesCode{
			Name:        "The Gambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_The_Gambia.svg","states":null}`,
		}
	case PolitiesGeorgia:
		return &PolitiesCode{
			Name:        "Georgia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Georgia.svg","states":null}`,
		}
	case PolitiesGermany:
		return &PolitiesCode{
			Name:        "Germany",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesGhana:
		return &PolitiesCode{
			Name:        "Ghana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Ghana.svg","states":null}`,
		}
	case PolitiesGibraltar:
		return &PolitiesCode{
			Name:        "Gibraltar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/02/Flag_of_Gibraltar.svg","states":null}`,
		}
	case PolitiesGreece:
		return &PolitiesCode{
			Name:        "Greece",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Greece.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesGreenland:
		return &PolitiesCode{
			Name:        "Greenland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_Greenland.svg","states":null}`,
		}
	case PolitiesGrenada:
		return &PolitiesCode{
			Name:        "Grenada",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Grenada.svg","states":null}`,
		}
	case PolitiesGuadeloupe:
		return &PolitiesCode{
			Name:        "Guadeloupe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesGuam:
		return &PolitiesCode{
			Name:        "Guam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/07/Flag_of_Guam.svg","states":null}`,
		}
	case PolitiesGuatemala:
		return &PolitiesCode{
			Name:        "Guatemala",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ec/Flag_of_Guatemala.svg","states":null}`,
		}
	case PolitiesGuernsey:
		return &PolitiesCode{
			Name:        "Guernsey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fa/Flag_of_Guernsey.svg","states":null}`,
		}
	case PolitiesGuinea:
		return &PolitiesCode{
			Name:        "Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_Guinea.svg","states":null}`,
		}
	case PolitiesGuineaBissau:
		return &PolitiesCode{
			Name:        "Guinea-Bissau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Guinea-Bissau.svg","states":null}`,
		}
	case PolitiesGuyana:
		return &PolitiesCode{
			Name:        "Guyana",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_Guyana.svg","states":null}`,
		}
	case PolitiesHaiti:
		return &PolitiesCode{
			Name:        "Haiti",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Haiti.svg","states":null}`,
		}
	case PolitiesHeardIslandandMcDonaldIslands:
		return &PolitiesCode{
			Name:        "Heard Island and McDonald Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg","states":null}`,
		}
	case PolitiesHonduras:
		return &PolitiesCode{
			Name:        "Honduras",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/82/Flag_of_Honduras.svg","states":null}`,
		}
	case PolitiesHongKong:
		return &PolitiesCode{
			Name:        "Hong Kong",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg","gov_fiscal_year":"0401-0331","states":null}`,
		}
	case PolitiesHungary:
		return &PolitiesCode{
			Name:        "Hungary",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c1/Flag_of_Hungary.svg","states":null}`,
		}
	case PolitiesIceland:
		return &PolitiesCode{
			Name:        "Iceland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Iceland.svg","states":null}`,
		}
	case PolitiesIndia:
		return &PolitiesCode{
			Name:        "India",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0401-0331","flag":"https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg","gov_fiscal_year":"0401-0331","states":{"INAN":"Andaman and Nicobar Islands","INAP":"Andhra Pradesh","INAR":"Arunachal Pradesh","INAS":"Assam","INBR":"Bihar","INCH":"Chandigarh","INCT":"Chhattisgarh","INDD":"Daman and Diu","INDL":"Delhi","INDN":"Dadra and Nagar Haveli","INGA":"Goa","INGJ":"Gujarat","INHP":"Himachal Pradesh","INHR":"Haryana","INJH":"Jharkhand","INJK":"Jammu and Kashmir","INKA":"Karnataka","INKL":"Kerala","INLD":"Lakshadweep","INMH":"Maharashtra","INML":"Meghalaya","INMN":"Manipur","INMP":"Madhya Pradesh","INMZ":"Mizoram","INNL":"Nagaland","INOR":"Odisha (formerly known as Orissa)","INPB":"Punjab","INPY":"Puducherry (Pondicherry)","INRJ":"Rajasthan","INSK":"Sikkim","INTG":"Telangana","INTN":"Tamil Nadu","INTR":"Tripura","INUP":"Uttar Pradesh","INUT":"Uttarakhand","INWB":"West Bengal"}}`,
		}
	case PolitiesIndonesia:
		return &PolitiesCode{
			Name:        "Indonesia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesIran:
		return &PolitiesCode{
			Name:        "Iran",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ca/Flag_of_Iran.svg","states":null}`,
		}
	case PolitiesIraq:
		return &PolitiesCode{
			Name:        "Iraq",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f6/Flag_of_Iraq.svg","states":null}`,
		}
	case PolitiesIreland:
		return &PolitiesCode{
			Name:        "Ireland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/4/45/Flag_of_Ireland.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesIsleofMan:
		return &PolitiesCode{
			Name:        "Isle of Man",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5d/Flag_of_the_Isle_of_Mann.svg","states":null}`,
		}
	case PolitiesIsrael:
		return &PolitiesCode{
			Name:        "Israel",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d4/Flag_of_Israel.svg","states":null}`,
		}
	case PolitiesItaly:
		return &PolitiesCode{
			Name:        "Italy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/0/03/Flag_of_Italy.svg","states":null}`,
		}
	case PolitiesIvoryCoast:
		return &PolitiesCode{
			Name:        "Ivory Coast",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_C%C3%B4te_d%27Ivoire.svg","states":null}`,
		}
	case PolitiesJamaica:
		return &PolitiesCode{
			Name:        "Jamaica",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0a/Flag_of_Jamaica.svg","states":null}`,
		}
	case PolitiesJapan:
		return &PolitiesCode{
			Name:        "Japan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg","states":null}`,
		}
	case PolitiesJersey:
		return &PolitiesCode{
			Name:        "Jersey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1c/Flag_of_Jersey.svg","states":null}`,
		}
	case PolitiesJordan:
		return &PolitiesCode{
			Name:        "Jordan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c0/Flag_of_Jordan.svg","states":null}`,
		}
	case PolitiesKazakhstan:
		return &PolitiesCode{
			Name:        "Kazakhstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kazakhstan.svg","states":null}`,
		}
	case PolitiesKenya:
		return &PolitiesCode{
			Name:        "Kenya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Kenya.svg","states":null}`,
		}
	case PolitiesKiribati:
		return &PolitiesCode{
			Name:        "Kiribati",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Kiribati.svg","states":null}`,
		}
	case PolitiesKuwait:
		return &PolitiesCode{
			Name:        "Kuwait",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/aa/Flag_of_Kuwait.svg","states":null}`,
		}
	case PolitiesKyrgyzstan:
		return &PolitiesCode{
			Name:        "Kyrgyzstan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/c7/Flag_of_Kyrgyzstan.svg","states":null}`,
		}
	case PolitiesLaos:
		return &PolitiesCode{
			Name:        "Laos",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/56/Flag_of_Laos.svg","states":null}`,
		}
	case PolitiesLatvia:
		return &PolitiesCode{
			Name:        "Latvia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Latvia.svg","states":null}`,
		}
	case PolitiesLebanon:
		return &PolitiesCode{
			Name:        "Lebanon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/59/Flag_of_Lebanon.svg","states":null}`,
		}
	case PolitiesLesotho:
		return &PolitiesCode{
			Name:        "Lesotho",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Lesotho.svg","states":null}`,
		}
	case PolitiesLiberia:
		return &PolitiesCode{
			Name:        "Liberia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b8/Flag_of_Liberia.svg","states":null}`,
		}
	case PolitiesLibya:
		return &PolitiesCode{
			Name:        "Libya",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/05/Flag_of_Libya.svg","states":null}`,
		}
	case PolitiesLiechtenstein:
		return &PolitiesCode{
			Name:        "Liechtenstein",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/47/Flag_of_Liechtenstein.svg","states":null}`,
		}
	case PolitiesLithuania:
		return &PolitiesCode{
			Name:        "Lithuania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Lithuania.svg","states":null}`,
		}
	case PolitiesLuxembourg:
		return &PolitiesCode{
			Name:        "Luxembourg",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/da/Flag_of_Luxembourg.svg","states":null}`,
		}
	case PolitiesMacau:
		return &PolitiesCode{
			Name:        "Macau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/63/Flag_of_Macau.svg","states":null}`,
		}
	case PolitiesMacedonia:
		return &PolitiesCode{
			Name:        "Macedonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_Macedonia.svg","states":null}`,
		}
	case PolitiesMadagascar:
		return &PolitiesCode{
			Name:        "Madagascar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Madagascar.svg","states":null}`,
		}
	case PolitiesMalawi:
		return &PolitiesCode{
			Name:        "Malawi",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d1/Flag_of_Malawi.svg","states":null}`,
		}
	case PolitiesMalaysia:
		return &PolitiesCode{
			Name:        "Malaysia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/66/Flag_of_Malaysia.svg","states":null}`,
		}
	case PolitiesMaldives:
		return &PolitiesCode{
			Name:        "Maldives",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0f/Flag_of_Maldives.svg","states":null}`,
		}
	case PolitiesMali:
		return &PolitiesCode{
			Name:        "Mali",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Mali.svg","states":null}`,
		}
	case PolitiesMalta:
		return &PolitiesCode{
			Name:        "Malta",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Malta.svg","states":null}`,
		}
	case PolitiesMarshallIslands:
		return &PolitiesCode{
			Name:        "Marshall Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2e/Flag_of_the_Marshall_Islands.svg","states":null}`,
		}
	case PolitiesMartinique:
		return &PolitiesCode{
			Name:        "Martinique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesMauritania:
		return &PolitiesCode{
			Name:        "Mauritania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/New_National_Flag_of_Mauritania.png","states":null}`,
		}
	case PolitiesMauritius:
		return &PolitiesCode{
			Name:        "Mauritius",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/77/Flag_of_Mauritius.svg","states":null}`,
		}
	case PolitiesMayotte:
		return &PolitiesCode{
			Name:        "Mayotte",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4a/Flag_of_Mayotte_%28local%29.svg","states":null}`,
		}
	case PolitiesMexico:
		return &PolitiesCode{
			Name:        "Mexico",
			Label:       "",
			Description: "",
			MetaData:    `{"MXAGU":"Aguascalientes","MXBCN":"Baja California","MXBCS":"Baja California Sur","MXCAM":"Campeche","MXCHH":"Chihuahua","MXCHP":"Chiapas","MXCOA":"Coahuila","MXCOL":"Colima","MXDIF":"Distrito Federal","MXDUR":"Durango","MXGRO":"Guerrero","MXGUA":"Guanajuato","MXHID":"Hidalgo","MXJAL":"Jalisco","MXMEX":"Mexico (Federal District)","MXMIC":"Michoacán","MXMOR":"Morelos","MXNAY":"Nayarit","MXNLE":"Nuevo León","MXOAX":"Oaxaca","MXPUE":"Puebla","MXQUE":"Querétaro","MXROO":"Quintana Roo","MXSIN":"Sinaloa","MXSLP":"San Luis Potosí","MXSON":"Sonora","MXTAB":"Tabasco","MXTAM":"Tamaulipas","MXTLA":"Tlaxcala","MXVER":"Veracruz","MXYUC":"Yucatán","MXZAC":"Zacatecas","flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Mexico.svg","states":null}`,
		}
	case PolitiesMicronesia:
		return &PolitiesCode{
			Name:        "Micronesia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"states":null}`,
		}
	case PolitiesMoldova:
		return &PolitiesCode{
			Name:        "Moldova",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Moldova.svg","states":null}`,
		}
	case PolitiesMonaco:
		return &PolitiesCode{
			Name:        "Monaco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ea/Flag_of_Monaco.svg","states":null}`,
		}
	case PolitiesMongolia:
		return &PolitiesCode{
			Name:        "Mongolia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4c/Flag_of_Mongolia.svg","states":null}`,
		}
	case PolitiesMontenegro:
		return &PolitiesCode{
			Name:        "Montenegro",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Montenegro.svg","states":null}`,
		}
	case PolitiesMontserrat:
		return &PolitiesCode{
			Name:        "Montserrat",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Montserrat.svg","states":null}`,
		}
	case PolitiesMorocco:
		return &PolitiesCode{
			Name:        "Morocco",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/2c/Flag_of_Morocco.svg","states":null}`,
		}
	case PolitiesMozambique:
		return &PolitiesCode{
			Name:        "Mozambique",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Mozambique.svg","states":null}`,
		}
	case PolitiesMyanmar:
		return &PolitiesCode{
			Name:        "Myanmar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8c/Flag_of_Myanmar.svg","states":null}`,
		}
	case PolitiesNamibia:
		return &PolitiesCode{
			Name:        "Namibia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Namibia.svg","states":null}`,
		}
	case PolitiesNauru:
		return &PolitiesCode{
			Name:        "Nauru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/30/Flag_of_Nauru.svg","states":null}`,
		}
	case PolitiesNepal:
		return &PolitiesCode{
			Name:        "Nepal",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/9/9b/Flag_of_Nepal.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesNetherlands:
		return &PolitiesCode{
			Name:        "Netherlands",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesNewCaledonia:
		return &PolitiesCode{
			Name:        "New Caledonia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesNewZealand:
		return &PolitiesCode{
			Name:        "New Zealand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/3e/Flag_of_New_Zealand.svg","states":null}`,
		}
	case PolitiesNicaragua:
		return &PolitiesCode{
			Name:        "Nicaragua",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/19/Flag_of_Nicaragua.svg","states":null}`,
		}
	case PolitiesNiger:
		return &PolitiesCode{
			Name:        "Niger",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f4/Flag_of_Niger.svg","states":null}`,
		}
	case PolitiesNigeria:
		return &PolitiesCode{
			Name:        "Nigeria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/79/Flag_of_Nigeria.svg","states":null}`,
		}
	case PolitiesNiue:
		return &PolitiesCode{
			Name:        "Niue",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Niue.svg","states":null}`,
		}
	case PolitiesNorfolkandPhilipIsland:
		return &PolitiesCode{
			Name:        "Norfolk and Philip Island",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Norfolk_Island.svg","states":null}`,
		}
	case PolitiesNorthKorea:
		return &PolitiesCode{
			Name:        "North Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/51/Flag_of_North_Korea.svg","states":null}`,
		}
	case PolitiesNorthernMarianaIslands:
		return &PolitiesCode{
			Name:        "Northern Mariana Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e0/Flag_of_the_Northern_Mariana_Islands.svg","states":null}`,
		}
	case PolitiesNorway:
		return &PolitiesCode{
			Name:        "Norway",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Norway.svg","states":null}`,
		}
	case PolitiesOman:
		return &PolitiesCode{
			Name:        "Oman",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/dd/Flag_of_Oman.svg","states":null}`,
		}
	case PolitiesPakistan:
		return &PolitiesCode{
			Name:        "Pakistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/32/Flag_of_Pakistan.svg","states":null}`,
		}
	case PolitiesPalau:
		return &PolitiesCode{
			Name:        "Palau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Palau.svg","states":null}`,
		}
	case PolitiesPalestinianTerritory:
		return &PolitiesCode{
			Name:        "Palestinian Territory",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_Palestine.svg","states":null}`,
		}
	case PolitiesPanama:
		return &PolitiesCode{
			Name:        "Panama",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/ab/Flag_of_Panama.svg","states":null}`,
		}
	case PolitiesPapuaNewGuinea:
		return &PolitiesCode{
			Name:        "Papua New Guinea",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e3/Flag_of_Papua_New_Guinea.svg","states":null}`,
		}
	case PolitiesParaguay:
		return &PolitiesCode{
			Name:        "Paraguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/27/Flag_of_Paraguay.svg","states":null}`,
		}
	case PolitiesPeru:
		return &PolitiesCode{
			Name:        "Peru",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Peru_%28state%29.svg","states":null}`,
		}
	case PolitiesPhilippines:
		return &PolitiesCode{
			Name:        "Philippines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/99/Flag_of_the_Philippines.svg","states":null}`,
		}
	case PolitiesPitcairnIslands:
		return &PolitiesCode{
			Name:        "Pitcairn Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_the_Pitcairn_Islands.svg","states":null}`,
		}
	case PolitiesPoland:
		return &PolitiesCode{
			Name:        "Poland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg","states":null}`,
		}
	case PolitiesPortugal:
		return &PolitiesCode{
			Name:        "Portugal",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5c/Flag_of_Portugal.svg","states":null}`,
		}
	case PolitiesPuertoRico:
		return &PolitiesCode{
			Name:        "Puerto Rico",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/28/Flag_of_Puerto_Rico.svg","states":null}`,
		}
	case PolitiesQatar:
		return &PolitiesCode{
			Name:        "Qatar",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/65/Flag_of_Qatar.svg","states":null}`,
		}
	case PolitiesRéunion:
		return &PolitiesCode{
			Name:        "Réunion",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/5a/Flag_of_R%C3%A9union.svg","states":null}`,
		}
	case PolitiesRomania:
		return &PolitiesCode{
			Name:        "Romania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/73/Flag_of_Romania.svg","states":null}`,
		}
	case PolitiesRussia:
		return &PolitiesCode{
			Name:        "Russia",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/f/f3/Flag_of_Russia.svg","gov_fiscal_year":"0101-1231","states":{"RUAD":"Adygeya, Respublika","RUAL":"Altay, Respublika","RUALT":"Altayskiy kray","RUAMU":"Amurskaya oblast\"","RUARK":"Arkhangel\"skaya oblast\"","RUAST":"Astrakhanskaya oblast\"","RUBA":"Bashkortostan, Respublika","RUBEL":"Belgorodskaya oblast\"","RUBRY":"Bryanskaya oblast\"","RUBU":"Buryatiya, Respublika","RUCE":"Chechenskaya Respublika","RUCHE":"Chelyabinskaya oblast\"","RUCHU":"Chukotskiy avtonomnyy okrug","RUCU":"Chuvashskaya Respublika","RUDA":"Dagestan, Respublika","RUIN":"Ingushetiya, Respublika","RUIRK":"Irkutskaya oblast\"","RUIVA":"Ivanovskaya oblast\"","RUKAM":"Kamchatskiy kray","RUKB":"Kabardino-Balkarskaya Respublika","RUKC":"Karachayevo-Cherkesskaya Respubl.","RUKDA":"Krasnodarskiy kray","RUKEM":"Kemerovskaya oblast\"","RUKGD":"Kaliningradskaya oblast\"","RUKGN":"Kurganskaya oblast\"","RUKHA":"Khabarovskiy kray","RUKHM":"Khanty-Mansiyskiy avtonomnyy okrug-Yugra","RUKIR":"Kirovskaya oblast\"","RUKK":"Khakasiya, Respublika","RUKL":"Kalmykiya, Respublika","RUKLU":"Kaluzhskaya oblast\"","RUKO":"Komi, Respublika","RUKOS":"Kostromskaya oblast\"","RUKR":"Kareliya, Respublika","RUKRS":"Kurskaya oblast\"","RUKYA":"Krasnoyarskiy kray","RULEN":"Leningradskaya oblast\"","RULIP":"Lipetskaya oblast\"","RUMAG":"Magadanskaya oblast\"","RUME":"Mariy El, Respublika","RUMO":"Mordoviya, Respublika","RUMOS":"Moskovskaya oblast\"","RUMOW":"Moskva (autonomous city)","RUMUR":"Murmanskaya oblast\"","RUNEN":"Nenetskiy avtonomnyy okrug","RUNGR":"Novgorodskaya oblast\"","RUNIZ":"Nizhegorodskaya oblast\"","RUNVS":"Novosibirskaya oblast\"","RUOMS":"Omskaya oblast\"","RUORE":"Orenburgskaya oblast\"","RUORL":"Orlovskaya oblast\"","RUPER":"Permskiy kray","RUPNZ":"Penzenskaya oblast\"","RUPRI":"Primorskiy kray","RUPSK":"Pskovskaya oblast\"","RUROS":"Rostovskaya oblast\"","RURYA":"Ryazanskaya oblast\"","RUSA":"Sakha, Respublika","RUSAK":"Sakhalinskaya oblast\"","RUSAM":"Samarskaya oblast\"","RUSAR":"Saratovskaya oblast\"","RUSE":"Severnaya Osetiya-Alaniya, Respubl.","RUSMO":"Smolenskaya oblast\"","RUSPE":"Sankt-Peterburg (autonomous city)","RUSTA":"Stavropol\"skiy kray","RUSVE":"Sverdlovskaya oblast\"","RUTA":"Tatarstan, Respublika","RUTAM":"Tambovskaya oblast\"","RUTOM":"Tomskaya oblast\"","RUTUL":"Tul\"skaya oblast\"","RUTVE":"Tverskaya oblast\"","RUTY":"Tyva, Respublika","RUTYU":"Tyumenskaya oblast\"","RUUD":"Udmurtskaya Respublika","RUULY":"Ul\"yanovskaya oblast\"","RUVGG":"Volgogradskaya oblast\"","RUVLA":"Vladimirskaya oblast\"","RUVLG":"Vologodskaya oblast\"","RUVOR":"Voronezhskaya oblast\"","RUYAN":"Yamalo-Nenetskiy avtonomnyy okrug","RUYAR":"Yaroslavskaya oblast\"","RUYEV":"Yevreyskaya avtonomnaya oblast\"","RUZAB":"Zabaykal\"skiy kray"}}`,
		}
	case PolitiesRwanda:
		return &PolitiesCode{
			Name:        "Rwanda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Rwanda.svg","states":null}`,
		}
	case PolitiesSaintHelenaAscensionandTristandaCunha:
		return &PolitiesCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","states":null}`,
		}
	case PolitiesSaintKittsandNevis:
		return &PolitiesCode{
			Name:        "Saint Kitts and Nevis",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Saint_Kitts_and_Nevis.svg","states":null}`,
		}
	case PolitiesSaintLucia:
		return &PolitiesCode{
			Name:        "Saint Lucia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Saint_Lucia.svg","states":null}`,
		}
	case PolitiesSaintPierreandMiquelon:
		return &PolitiesCode{
			Name:        "Saint Pierre and Miquelon",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesSaintVincentandtheGrenadines:
		return &PolitiesCode{
			Name:        "Saint Vincent and the Grenadines",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6d/Flag_of_Saint_Vincent_and_the_Grenadines.svg","states":null}`,
		}
	case PolitiesSaintBarthelemy:
		return &PolitiesCode{
			Name:        "Saint-Barthelemy",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/df/Flag_of_Saint_Barthelemy_%28local%29.svg","states":null}`,
		}
	case PolitiesSaintMartin:
		return &PolitiesCode{
			Name:        "Saint-Martin",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"states":null}`,
		}
	case PolitiesSamoa:
		return &PolitiesCode{
			Name:        "Samoa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/31/Flag_of_Samoa.svg","states":null}`,
		}
	case PolitiesSanMarino:
		return &PolitiesCode{
			Name:        "San Marino",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b1/Flag_of_San_Marino.svg","states":null}`,
		}
	case PolitiesSãoToméandPríncipe:
		return &PolitiesCode{
			Name:        "São Tomé and Príncipe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4f/Flag_of_Sao_Tome_and_Principe.svg","states":null}`,
		}
	case PolitiesSaudiArabia:
		return &PolitiesCode{
			Name:        "Saudi Arabia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/0d/Flag_of_Saudi_Arabia.svg","states":null}`,
		}
	case PolitiesSenegal:
		return &PolitiesCode{
			Name:        "Senegal",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fd/Flag_of_Senegal.svg","states":null}`,
		}
	case PolitiesSerbia:
		return &PolitiesCode{
			Name:        "Serbia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/ff/Flag_of_Serbia.svg","states":null}`,
		}
	case PolitiesSeychelles:
		return &PolitiesCode{
			Name:        "Seychelles",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fc/Flag_of_Seychelles.svg","states":null}`,
		}
	case PolitiesSierraLeone:
		return &PolitiesCode{
			Name:        "Sierra Leone",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/17/Flag_of_Sierra_Leone.svg","states":null}`,
		}
	case PolitiesSingapore:
		return &PolitiesCode{
			Name:        "Singapore",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg","states":null}`,
		}
	case PolitiesSintMaarten:
		return &PolitiesCode{
			Name:        "Sint Maarten",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d3/Flag_of_Sint_Maarten.svg","states":null}`,
		}
	case PolitiesSlovakia:
		return &PolitiesCode{
			Name:        "Slovakia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/e6/Flag_of_Slovakia.svg","states":null}`,
		}
	case PolitiesSlovenia:
		return &PolitiesCode{
			Name:        "Slovenia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f0/Flag_of_Slovenia.svg","states":null}`,
		}
	case PolitiesSolomonIslands:
		return &PolitiesCode{
			Name:        "Solomon Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/74/Flag_of_the_Solomon_Islands.svg","states":null}`,
		}
	case PolitiesSomalia:
		return &PolitiesCode{
			Name:        "Somalia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_Somalia.svg","states":null}`,
		}
	case PolitiesSouthAfrica:
		return &PolitiesCode{
			Name:        "South Africa",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/af/Flag_of_South_Africa.svg","states":null}`,
		}
	case PolitiesSouthGeorgiaandtheSouthSandwichIslands:
		return &PolitiesCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/e/ed/Flag_of_South_Georgia_and_the_South_Sandwich_Islands.svg","states":null}`,
		}
	case PolitiesSouthKorea:
		return &PolitiesCode{
			Name:        "South Korea",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesSouthSudan:
		return &PolitiesCode{
			Name:        "South Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7a/Flag_of_South_Sudan.svg","states":null}`,
		}
	case PolitiesSpain:
		return &PolitiesCode{
			Name:        "Spain",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/9/9a/Flag_of_Spain.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesSriLanka:
		return &PolitiesCode{
			Name:        "Sri Lanka",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/11/Flag_of_Sri_Lanka.svg","states":null}`,
		}
	case PolitiesSudan:
		return &PolitiesCode{
			Name:        "Sudan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/01/Flag_of_Sudan.svg","states":null}`,
		}
	case PolitiesSuriname:
		return &PolitiesCode{
			Name:        "Suriname",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/60/Flag_of_Suriname.svg","states":null}`,
		}
	case PolitiesSvalbardandJanMayen:
		return &PolitiesCode{
			Name:        "Svalbard and Jan Mayen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"states":null}`,
		}
	case PolitiesSwaziland:
		return &PolitiesCode{
			Name:        "Swaziland",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fb/Flag_of_Eswatini.svg","states":null}`,
		}
	case PolitiesSweden:
		return &PolitiesCode{
			Name:        "Sweden",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/en/4/4c/Flag_of_Sweden.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesSwitzerland:
		return &PolitiesCode{
			Name:        "Switzerland",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/0/08/Flag_of_Switzerland_%28Pantone%29.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesSyria:
		return &PolitiesCode{
			Name:        "Syria",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/5/53/Flag_of_Syria.svg","states":null}`,
		}
	case PolitiesTaiwan:
		return &PolitiesCode{
			Name:        "Taiwan",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0101-1231","flag":"https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg","gov_fiscal_year":"0101-1231","states":null}`,
		}
	case PolitiesTajikistan:
		return &PolitiesCode{
			Name:        "Tajikistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/d/d0/Flag_of_Tajikistan.svg","states":null}`,
		}
	case PolitiesTanzania:
		return &PolitiesCode{
			Name:        "Tanzania",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tanzania.svg","states":null}`,
		}
	case PolitiesThailand:
		return &PolitiesCode{
			Name:        "Thailand",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a9/Flag_of_Thailand.svg","states":null}`,
		}
	case PolitiesTogo:
		return &PolitiesCode{
			Name:        "Togo",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/68/Flag_of_Togo.svg","states":null}`,
		}
	case PolitiesTokelau:
		return &PolitiesCode{
			Name:        "Tokelau",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/8e/Flag_of_Tokelau.svg","states":null}`,
		}
	case PolitiesTonga:
		return &PolitiesCode{
			Name:        "Tonga",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/9/9a/Flag_of_Tonga.svg","states":null}`,
		}
	case PolitiesTrinidadandTobago:
		return &PolitiesCode{
			Name:        "Trinidad and Tobago",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/64/Flag_of_Trinidad_and_Tobago.svg","states":null}`,
		}
	case PolitiesTunisia:
		return &PolitiesCode{
			Name:        "Tunisia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/ce/Flag_of_Tunisia.svg","states":null}`,
		}
	case PolitiesTurkey:
		return &PolitiesCode{
			Name:        "Turkey",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/b/b4/Flag_of_Turkey.svg","states":null}`,
		}
	case PolitiesTurkmenistan:
		return &PolitiesCode{
			Name:        "Turkmenistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/1/1b/Flag_of_Turkmenistan.svg","states":null}`,
		}
	case PolitiesTurksandCaicosIslands:
		return &PolitiesCode{
			Name:        "Turks and Caicos Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/a/a0/Flag_of_the_Turks_and_Caicos_Islands.svg","states":null}`,
		}
	case PolitiesTuvalu:
		return &PolitiesCode{
			Name:        "Tuvalu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/3/38/Flag_of_Tuvalu.svg","states":null}`,
		}
	case PolitiesUganda:
		return &PolitiesCode{
			Name:        "Uganda",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/4e/Flag_of_Uganda.svg","states":null}`,
		}
	case PolitiesUkraine:
		return &PolitiesCode{
			Name:        "Ukraine",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/4/49/Flag_of_Ukraine.svg","states":null}`,
		}
	case PolitiesUnitedArabEmirates:
		return &PolitiesCode{
			Name:        "United Arab Emirates",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/c/cb/Flag_of_the_United_Arab_Emirates.svg","states":null}`,
		}
	case PolitiesUnitedKingdom:
		return &PolitiesCode{
			Name:        "United Kingdom",
			Label:       "",
			Description: "",
			MetaData:    `{"fiscal_year":"0406-0405","flag":"https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg","gov_fiscal_year":"0401-0331","states":null}`,
		}
	case PolitiesUnitedStatesMinorOutlyingIslands:
		return &PolitiesCode{
			Name:        "United States Minor Outlying Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","states":null}`,
		}
	case PolitiesUruguay:
		return &PolitiesCode{
			Name:        "Uruguay",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/fe/Flag_of_Uruguay.svg","states":null}`,
		}
	case PolitiesUSVirginIslands:
		return &PolitiesCode{
			Name:        "US Virgin Islands",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/f/f8/Flag_of_the_United_States_Virgin_Islands.svg","states":null}`,
		}
	case PolitiesUSA:
		return &PolitiesCode{
			Name:        "USA",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg","states":{"USAK":"Alaska","USAL":"Alabama","USAR":"Arkansas","USAS":"American Samoa","USAZ":"Arizona","USCA":"California","USCO":"Colorado","USCT":"Connecticut","USDC":"Washington D.C.","USDE":"Delaware","USFL":"Florida","USGA":"Georgia","USGU":"Guam","USHI":"Hawaii","USIA":"Iowa","USID":"Idaho","USIL":"Illinois","USIN":"Indiana","USKS":"Kansas","USKY":"Kentucky","USLA":"Louisiana","USMA":"Massachusetts","USMD":"Maryland","USME":"Maine","USMI":"Michigan","USMN":"Minnesota","USMO":"Missouri","USMP":"Northern Mariana Islands","USMS":"Mississippi","USMT":"Montana","USNC":"North Carolina","USND":"North Dakota","USNE":"Nebraska","USNH":"New Hampshire","USNJ":"New Jersey","USNM":"New Mexico","USNV":"Nevada","USNY":"New York","USOH":"Ohio","USOK":"Oklahoma","USOR":"Oregon","USPA":"Pennsylvania","USPR":"Puerto Rico","USRI":"Rhode Island","USSC":"South Carolina","USSD":"South Dakota","USTN":"Tennessee","USTX":"Texas","USUM":"United States Minor Outlying Islands","USUT":"Utah","USVA":"Virginia","USVI":"US Virgin Islands","USVT":"Vermont","USWA":"Washington","USWI":"Wisconsin","USWV":"West Virginia","USWY":"Wyoming"}}`,
		}
	case PolitiesUzbekistan:
		return &PolitiesCode{
			Name:        "Uzbekistan",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/84/Flag_of_Uzbekistan.svg","states":null}`,
		}
	case PolitiesVanuatu:
		return &PolitiesCode{
			Name:        "Vanuatu",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6e/Flag_of_Vanuatu_%28official%29.svg","states":null}`,
		}
	case PolitiesVaticanCity:
		return &PolitiesCode{
			Name:        "Vatican City",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/00/Flag_of_the_Vatican_City.svg","states":null}`,
		}
	case PolitiesVenezuela:
		return &PolitiesCode{
			Name:        "Venezuela",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/7/7b/Flag_of_Venezuela_%28state%29.svg","states":null}`,
		}
	case PolitiesVietnam:
		return &PolitiesCode{
			Name:        "Vietnam",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/2/21/Flag_of_Vietnam.svg","states":null}`,
		}
	case PolitiesWallisandFutuna:
		return &PolitiesCode{
			Name:        "Wallis and Futuna",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/en/c/c3/Flag_of_France.svg","states":null}`,
		}
	case PolitiesWesternSahara:
		return &PolitiesCode{
			Name:        "Western Sahara",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":null,"states":null}`,
		}
	case PolitiesYemen:
		return &PolitiesCode{
			Name:        "Yemen",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/8/89/Flag_of_Yemen.svg","states":null}`,
		}
	case PolitiesZambia:
		return &PolitiesCode{
			Name:        "Zambia",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/0/06/Flag_of_Zambia.svg","states":null}`,
		}
	case PolitiesZimbabwe:
		return &PolitiesCode{
			Name:        "Zimbabwe",
			Label:       "",
			Description: "",
			MetaData:    `{"flag":"https://upload.wikimedia.org/wikipedia/commons/6/6a/Flag_of_Zimbabwe.svg","states":null}`,
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
		},
		PolitiesAfghanistan: &PolitiesCode{
			Name:        "Afghanistan",
			Label:       "",
			Description: "",
		},
		PolitiesAlbania: &PolitiesCode{
			Name:        "Albania",
			Label:       "",
			Description: "",
		},
		PolitiesAlgeria: &PolitiesCode{
			Name:        "Algeria",
			Label:       "",
			Description: "",
		},
		PolitiesAmericanSamoa: &PolitiesCode{
			Name:        "American Samoa",
			Label:       "",
			Description: "",
		},
		PolitiesAndorra: &PolitiesCode{
			Name:        "Andorra",
			Label:       "",
			Description: "",
		},
		PolitiesAngola: &PolitiesCode{
			Name:        "Angola",
			Label:       "",
			Description: "",
		},
		PolitiesAnguilla: &PolitiesCode{
			Name:        "Anguilla",
			Label:       "",
			Description: "",
		},
		PolitiesAntarctica: &PolitiesCode{
			Name:        "Antarctica",
			Label:       "",
			Description: "",
		},
		PolitiesAntiguaandBarbuda: &PolitiesCode{
			Name:        "Antigua and Barbuda",
			Label:       "",
			Description: "",
		},
		PolitiesArgentina: &PolitiesCode{
			Name:        "Argentina",
			Label:       "",
			Description: "",
		},
		PolitiesArmenia: &PolitiesCode{
			Name:        "Armenia",
			Label:       "",
			Description: "",
		},
		PolitiesAruba: &PolitiesCode{
			Name:        "Aruba",
			Label:       "",
			Description: "",
		},
		PolitiesAfricanUnion: &PolitiesCode{
			Name:        "African Union",
			Label:       "",
			Description: "",
		},
		PolitiesAustralia: &PolitiesCode{
			Name:        "Australia",
			Label:       "",
			Description: "",
		},
		PolitiesAustria: &PolitiesCode{
			Name:        "Austria",
			Label:       "",
			Description: "",
		},
		PolitiesAzerbaijan: &PolitiesCode{
			Name:        "Azerbaijan",
			Label:       "",
			Description: "",
		},
		PolitiesTheBahamas: &PolitiesCode{
			Name:        "The Bahamas",
			Label:       "",
			Description: "",
		},
		PolitiesBahrain: &PolitiesCode{
			Name:        "Bahrain",
			Label:       "",
			Description: "",
		},
		PolitiesBangladesh: &PolitiesCode{
			Name:        "Bangladesh",
			Label:       "",
			Description: "",
		},
		PolitiesBarbados: &PolitiesCode{
			Name:        "Barbados",
			Label:       "",
			Description: "",
		},
		PolitiesBelarus: &PolitiesCode{
			Name:        "Belarus",
			Label:       "",
			Description: "",
		},
		PolitiesBelgium: &PolitiesCode{
			Name:        "Belgium",
			Label:       "",
			Description: "",
		},
		PolitiesBelize: &PolitiesCode{
			Name:        "Belize",
			Label:       "",
			Description: "",
		},
		PolitiesBenin: &PolitiesCode{
			Name:        "Benin",
			Label:       "",
			Description: "",
		},
		PolitiesBermuda: &PolitiesCode{
			Name:        "Bermuda",
			Label:       "",
			Description: "",
		},
		PolitiesBhutan: &PolitiesCode{
			Name:        "Bhutan",
			Label:       "",
			Description: "",
		},
		PolitiesBolivia: &PolitiesCode{
			Name:        "Bolivia",
			Label:       "",
			Description: "",
		},
		PolitiesBonaireStEustasuisandSaba: &PolitiesCode{
			Name:        "Bonaire, St Eustasuis and Saba",
			Label:       "",
			Description: "",
		},
		PolitiesBosniaandHerzegovina: &PolitiesCode{
			Name:        "Bosnia and Herzegovina",
			Label:       "",
			Description: "",
		},
		PolitiesBotswana: &PolitiesCode{
			Name:        "Botswana",
			Label:       "",
			Description: "",
		},
		PolitiesBouvetIsland: &PolitiesCode{
			Name:        "Bouvet Island",
			Label:       "",
			Description: "",
		},
		PolitiesBrazil: &PolitiesCode{
			Name:        "Brazil",
			Label:       "",
			Description: "",
		},
		PolitiesBritishIndianOceanTerritory: &PolitiesCode{
			Name:        "British Indian Ocean Territory",
			Label:       "",
			Description: "",
		},
		PolitiesBritishVirginIslands: &PolitiesCode{
			Name:        "British Virgin Islands",
			Label:       "",
			Description: "",
		},
		PolitiesBrunei: &PolitiesCode{
			Name:        "Brunei",
			Label:       "",
			Description: "",
		},
		PolitiesBulgaria: &PolitiesCode{
			Name:        "Bulgaria",
			Label:       "",
			Description: "",
		},
		PolitiesBurkinaFaso: &PolitiesCode{
			Name:        "Burkina Faso",
			Label:       "",
			Description: "",
		},
		PolitiesBurundi: &PolitiesCode{
			Name:        "Burundi",
			Label:       "",
			Description: "",
		},
		PolitiesCambodia: &PolitiesCode{
			Name:        "Cambodia",
			Label:       "",
			Description: "",
		},
		PolitiesCameroon: &PolitiesCode{
			Name:        "Cameroon",
			Label:       "",
			Description: "",
		},
		PolitiesCanada: &PolitiesCode{
			Name:        "Canada",
			Label:       "",
			Description: "",
		},
		PolitiesCapeVerde: &PolitiesCode{
			Name:        "Cape Verde",
			Label:       "",
			Description: "",
		},
		PolitiesCaymanIslands: &PolitiesCode{
			Name:        "Cayman Islands",
			Label:       "",
			Description: "",
		},
		PolitiesCentralAfricanRepublic: &PolitiesCode{
			Name:        "Central African Republic",
			Label:       "",
			Description: "",
		},
		PolitiesChad: &PolitiesCode{
			Name:        "Chad",
			Label:       "",
			Description: "",
		},
		PolitiesChile: &PolitiesCode{
			Name:        "Chile",
			Label:       "",
			Description: "",
		},
		PolitiesChina: &PolitiesCode{
			Name:        "China",
			Label:       "",
			Description: "",
		},
		PolitiesChristmasIsland: &PolitiesCode{
			Name:        "Christmas Island",
			Label:       "",
			Description: "",
		},
		PolitiesCocosIslands: &PolitiesCode{
			Name:        "Cocos Islands",
			Label:       "",
			Description: "",
		},
		PolitiesColombia: &PolitiesCode{
			Name:        "Colombia",
			Label:       "",
			Description: "",
		},
		PolitiesComoros: &PolitiesCode{
			Name:        "Comoros",
			Label:       "",
			Description: "",
		},
		PolitiesCongoBrazzaville: &PolitiesCode{
			Name:        "Congo-Brazzaville",
			Label:       "",
			Description: "",
		},
		PolitiesCongoKinshasa: &PolitiesCode{
			Name:        "Congo-Kinshasa",
			Label:       "",
			Description: "",
		},
		PolitiesCookIslands: &PolitiesCode{
			Name:        "Cook Islands",
			Label:       "",
			Description: "",
		},
		PolitiesCostaRica: &PolitiesCode{
			Name:        "Costa Rica",
			Label:       "",
			Description: "",
		},
		PolitiesCroatia: &PolitiesCode{
			Name:        "Croatia",
			Label:       "",
			Description: "",
		},
		PolitiesCuba: &PolitiesCode{
			Name:        "Cuba",
			Label:       "",
			Description: "",
		},
		PolitiesCuracao: &PolitiesCode{
			Name:        "Curacao",
			Label:       "",
			Description: "",
		},
		PolitiesCyprus: &PolitiesCode{
			Name:        "Cyprus",
			Label:       "",
			Description: "",
		},
		PolitiesCzechRepublic: &PolitiesCode{
			Name:        "Czech Republic",
			Label:       "",
			Description: "",
		},
		PolitiesDenmark: &PolitiesCode{
			Name:        "Denmark",
			Label:       "",
			Description: "",
		},
		PolitiesDjibouti: &PolitiesCode{
			Name:        "Djibouti",
			Label:       "",
			Description: "",
		},
		PolitiesDominica: &PolitiesCode{
			Name:        "Dominica",
			Label:       "",
			Description: "",
		},
		PolitiesDominicanRepublic: &PolitiesCode{
			Name:        "Dominican Republic",
			Label:       "",
			Description: "",
		},
		PolitiesEastTimor: &PolitiesCode{
			Name:        "East Timor",
			Label:       "",
			Description: "",
		},
		PolitiesEcuador: &PolitiesCode{
			Name:        "Ecuador",
			Label:       "",
			Description: "",
		},
		PolitiesEgypt: &PolitiesCode{
			Name:        "Egypt",
			Label:       "",
			Description: "",
		},
		PolitiesElSalvador: &PolitiesCode{
			Name:        "El Salvador",
			Label:       "",
			Description: "",
		},
		PolitiesEquatorialGuinea: &PolitiesCode{
			Name:        "Equatorial Guinea",
			Label:       "",
			Description: "",
		},
		PolitiesEritrea: &PolitiesCode{
			Name:        "Eritrea",
			Label:       "",
			Description: "",
		},
		PolitiesEstonia: &PolitiesCode{
			Name:        "Estonia",
			Label:       "",
			Description: "",
		},
		PolitiesEthiopia: &PolitiesCode{
			Name:        "Ethiopia",
			Label:       "",
			Description: "",
		},
		PolitiesEuropeanUnion: &PolitiesCode{
			Name:        "European Union",
			Label:       "",
			Description: "",
		},
		PolitiesFalklandIslands: &PolitiesCode{
			Name:        "Falkland Islands",
			Label:       "",
			Description: "",
		},
		PolitiesFaroeIslands: &PolitiesCode{
			Name:        "Faroe Islands",
			Label:       "",
			Description: "",
		},
		PolitiesFiji: &PolitiesCode{
			Name:        "Fiji",
			Label:       "",
			Description: "",
		},
		PolitiesFinland: &PolitiesCode{
			Name:        "Finland",
			Label:       "",
			Description: "",
		},
		PolitiesFrance: &PolitiesCode{
			Name:        "France",
			Label:       "",
			Description: "",
		},
		PolitiesFrenchGuiana: &PolitiesCode{
			Name:        "French Guiana",
			Label:       "",
			Description: "",
		},
		PolitiesFrenchPolynesia: &PolitiesCode{
			Name:        "French Polynesia",
			Label:       "",
			Description: "",
		},
		PolitiesFrenchSouthernandAntarcticLands: &PolitiesCode{
			Name:        "French Southern and Antarctic Lands",
			Label:       "",
			Description: "",
		},
		PolitiesGabon: &PolitiesCode{
			Name:        "Gabon",
			Label:       "",
			Description: "",
		},
		PolitiesTheGambia: &PolitiesCode{
			Name:        "The Gambia",
			Label:       "",
			Description: "",
		},
		PolitiesGeorgia: &PolitiesCode{
			Name:        "Georgia",
			Label:       "",
			Description: "",
		},
		PolitiesGermany: &PolitiesCode{
			Name:        "Germany",
			Label:       "",
			Description: "",
		},
		PolitiesGhana: &PolitiesCode{
			Name:        "Ghana",
			Label:       "",
			Description: "",
		},
		PolitiesGibraltar: &PolitiesCode{
			Name:        "Gibraltar",
			Label:       "",
			Description: "",
		},
		PolitiesGreece: &PolitiesCode{
			Name:        "Greece",
			Label:       "",
			Description: "",
		},
		PolitiesGreenland: &PolitiesCode{
			Name:        "Greenland",
			Label:       "",
			Description: "",
		},
		PolitiesGrenada: &PolitiesCode{
			Name:        "Grenada",
			Label:       "",
			Description: "",
		},
		PolitiesGuadeloupe: &PolitiesCode{
			Name:        "Guadeloupe",
			Label:       "",
			Description: "",
		},
		PolitiesGuam: &PolitiesCode{
			Name:        "Guam",
			Label:       "",
			Description: "",
		},
		PolitiesGuatemala: &PolitiesCode{
			Name:        "Guatemala",
			Label:       "",
			Description: "",
		},
		PolitiesGuernsey: &PolitiesCode{
			Name:        "Guernsey",
			Label:       "",
			Description: "",
		},
		PolitiesGuinea: &PolitiesCode{
			Name:        "Guinea",
			Label:       "",
			Description: "",
		},
		PolitiesGuineaBissau: &PolitiesCode{
			Name:        "Guinea-Bissau",
			Label:       "",
			Description: "",
		},
		PolitiesGuyana: &PolitiesCode{
			Name:        "Guyana",
			Label:       "",
			Description: "",
		},
		PolitiesHaiti: &PolitiesCode{
			Name:        "Haiti",
			Label:       "",
			Description: "",
		},
		PolitiesHeardIslandandMcDonaldIslands: &PolitiesCode{
			Name:        "Heard Island and McDonald Islands",
			Label:       "",
			Description: "",
		},
		PolitiesHonduras: &PolitiesCode{
			Name:        "Honduras",
			Label:       "",
			Description: "",
		},
		PolitiesHongKong: &PolitiesCode{
			Name:        "Hong Kong",
			Label:       "",
			Description: "",
		},
		PolitiesHungary: &PolitiesCode{
			Name:        "Hungary",
			Label:       "",
			Description: "",
		},
		PolitiesIceland: &PolitiesCode{
			Name:        "Iceland",
			Label:       "",
			Description: "",
		},
		PolitiesIndia: &PolitiesCode{
			Name:        "India",
			Label:       "",
			Description: "",
		},
		PolitiesIndonesia: &PolitiesCode{
			Name:        "Indonesia",
			Label:       "",
			Description: "",
		},
		PolitiesIran: &PolitiesCode{
			Name:        "Iran",
			Label:       "",
			Description: "",
		},
		PolitiesIraq: &PolitiesCode{
			Name:        "Iraq",
			Label:       "",
			Description: "",
		},
		PolitiesIreland: &PolitiesCode{
			Name:        "Ireland",
			Label:       "",
			Description: "",
		},
		PolitiesIsleofMan: &PolitiesCode{
			Name:        "Isle of Man",
			Label:       "",
			Description: "",
		},
		PolitiesIsrael: &PolitiesCode{
			Name:        "Israel",
			Label:       "",
			Description: "",
		},
		PolitiesItaly: &PolitiesCode{
			Name:        "Italy",
			Label:       "",
			Description: "",
		},
		PolitiesIvoryCoast: &PolitiesCode{
			Name:        "Ivory Coast",
			Label:       "",
			Description: "",
		},
		PolitiesJamaica: &PolitiesCode{
			Name:        "Jamaica",
			Label:       "",
			Description: "",
		},
		PolitiesJapan: &PolitiesCode{
			Name:        "Japan",
			Label:       "",
			Description: "",
		},
		PolitiesJersey: &PolitiesCode{
			Name:        "Jersey",
			Label:       "",
			Description: "",
		},
		PolitiesJordan: &PolitiesCode{
			Name:        "Jordan",
			Label:       "",
			Description: "",
		},
		PolitiesKazakhstan: &PolitiesCode{
			Name:        "Kazakhstan",
			Label:       "",
			Description: "",
		},
		PolitiesKenya: &PolitiesCode{
			Name:        "Kenya",
			Label:       "",
			Description: "",
		},
		PolitiesKiribati: &PolitiesCode{
			Name:        "Kiribati",
			Label:       "",
			Description: "",
		},
		PolitiesKuwait: &PolitiesCode{
			Name:        "Kuwait",
			Label:       "",
			Description: "",
		},
		PolitiesKyrgyzstan: &PolitiesCode{
			Name:        "Kyrgyzstan",
			Label:       "",
			Description: "",
		},
		PolitiesLaos: &PolitiesCode{
			Name:        "Laos",
			Label:       "",
			Description: "",
		},
		PolitiesLatvia: &PolitiesCode{
			Name:        "Latvia",
			Label:       "",
			Description: "",
		},
		PolitiesLebanon: &PolitiesCode{
			Name:        "Lebanon",
			Label:       "",
			Description: "",
		},
		PolitiesLesotho: &PolitiesCode{
			Name:        "Lesotho",
			Label:       "",
			Description: "",
		},
		PolitiesLiberia: &PolitiesCode{
			Name:        "Liberia",
			Label:       "",
			Description: "",
		},
		PolitiesLibya: &PolitiesCode{
			Name:        "Libya",
			Label:       "",
			Description: "",
		},
		PolitiesLiechtenstein: &PolitiesCode{
			Name:        "Liechtenstein",
			Label:       "",
			Description: "",
		},
		PolitiesLithuania: &PolitiesCode{
			Name:        "Lithuania",
			Label:       "",
			Description: "",
		},
		PolitiesLuxembourg: &PolitiesCode{
			Name:        "Luxembourg",
			Label:       "",
			Description: "",
		},
		PolitiesMacau: &PolitiesCode{
			Name:        "Macau",
			Label:       "",
			Description: "",
		},
		PolitiesMacedonia: &PolitiesCode{
			Name:        "Macedonia",
			Label:       "",
			Description: "",
		},
		PolitiesMadagascar: &PolitiesCode{
			Name:        "Madagascar",
			Label:       "",
			Description: "",
		},
		PolitiesMalawi: &PolitiesCode{
			Name:        "Malawi",
			Label:       "",
			Description: "",
		},
		PolitiesMalaysia: &PolitiesCode{
			Name:        "Malaysia",
			Label:       "",
			Description: "",
		},
		PolitiesMaldives: &PolitiesCode{
			Name:        "Maldives",
			Label:       "",
			Description: "",
		},
		PolitiesMali: &PolitiesCode{
			Name:        "Mali",
			Label:       "",
			Description: "",
		},
		PolitiesMalta: &PolitiesCode{
			Name:        "Malta",
			Label:       "",
			Description: "",
		},
		PolitiesMarshallIslands: &PolitiesCode{
			Name:        "Marshall Islands",
			Label:       "",
			Description: "",
		},
		PolitiesMartinique: &PolitiesCode{
			Name:        "Martinique",
			Label:       "",
			Description: "",
		},
		PolitiesMauritania: &PolitiesCode{
			Name:        "Mauritania",
			Label:       "",
			Description: "",
		},
		PolitiesMauritius: &PolitiesCode{
			Name:        "Mauritius",
			Label:       "",
			Description: "",
		},
		PolitiesMayotte: &PolitiesCode{
			Name:        "Mayotte",
			Label:       "",
			Description: "",
		},
		PolitiesMexico: &PolitiesCode{
			Name:        "Mexico",
			Label:       "",
			Description: "",
		},
		PolitiesMicronesia: &PolitiesCode{
			Name:        "Micronesia",
			Label:       "",
			Description: "",
		},
		PolitiesMoldova: &PolitiesCode{
			Name:        "Moldova",
			Label:       "",
			Description: "",
		},
		PolitiesMonaco: &PolitiesCode{
			Name:        "Monaco",
			Label:       "",
			Description: "",
		},
		PolitiesMongolia: &PolitiesCode{
			Name:        "Mongolia",
			Label:       "",
			Description: "",
		},
		PolitiesMontenegro: &PolitiesCode{
			Name:        "Montenegro",
			Label:       "",
			Description: "",
		},
		PolitiesMontserrat: &PolitiesCode{
			Name:        "Montserrat",
			Label:       "",
			Description: "",
		},
		PolitiesMorocco: &PolitiesCode{
			Name:        "Morocco",
			Label:       "",
			Description: "",
		},
		PolitiesMozambique: &PolitiesCode{
			Name:        "Mozambique",
			Label:       "",
			Description: "",
		},
		PolitiesMyanmar: &PolitiesCode{
			Name:        "Myanmar",
			Label:       "",
			Description: "",
		},
		PolitiesNamibia: &PolitiesCode{
			Name:        "Namibia",
			Label:       "",
			Description: "",
		},
		PolitiesNauru: &PolitiesCode{
			Name:        "Nauru",
			Label:       "",
			Description: "",
		},
		PolitiesNepal: &PolitiesCode{
			Name:        "Nepal",
			Label:       "",
			Description: "",
		},
		PolitiesNetherlands: &PolitiesCode{
			Name:        "Netherlands",
			Label:       "",
			Description: "",
		},
		PolitiesNewCaledonia: &PolitiesCode{
			Name:        "New Caledonia",
			Label:       "",
			Description: "",
		},
		PolitiesNewZealand: &PolitiesCode{
			Name:        "New Zealand",
			Label:       "",
			Description: "",
		},
		PolitiesNicaragua: &PolitiesCode{
			Name:        "Nicaragua",
			Label:       "",
			Description: "",
		},
		PolitiesNiger: &PolitiesCode{
			Name:        "Niger",
			Label:       "",
			Description: "",
		},
		PolitiesNigeria: &PolitiesCode{
			Name:        "Nigeria",
			Label:       "",
			Description: "",
		},
		PolitiesNiue: &PolitiesCode{
			Name:        "Niue",
			Label:       "",
			Description: "",
		},
		PolitiesNorfolkandPhilipIsland: &PolitiesCode{
			Name:        "Norfolk and Philip Island",
			Label:       "",
			Description: "",
		},
		PolitiesNorthKorea: &PolitiesCode{
			Name:        "North Korea",
			Label:       "",
			Description: "",
		},
		PolitiesNorthernMarianaIslands: &PolitiesCode{
			Name:        "Northern Mariana Islands",
			Label:       "",
			Description: "",
		},
		PolitiesNorway: &PolitiesCode{
			Name:        "Norway",
			Label:       "",
			Description: "",
		},
		PolitiesOman: &PolitiesCode{
			Name:        "Oman",
			Label:       "",
			Description: "",
		},
		PolitiesPakistan: &PolitiesCode{
			Name:        "Pakistan",
			Label:       "",
			Description: "",
		},
		PolitiesPalau: &PolitiesCode{
			Name:        "Palau",
			Label:       "",
			Description: "",
		},
		PolitiesPalestinianTerritory: &PolitiesCode{
			Name:        "Palestinian Territory",
			Label:       "",
			Description: "",
		},
		PolitiesPanama: &PolitiesCode{
			Name:        "Panama",
			Label:       "",
			Description: "",
		},
		PolitiesPapuaNewGuinea: &PolitiesCode{
			Name:        "Papua New Guinea",
			Label:       "",
			Description: "",
		},
		PolitiesParaguay: &PolitiesCode{
			Name:        "Paraguay",
			Label:       "",
			Description: "",
		},
		PolitiesPeru: &PolitiesCode{
			Name:        "Peru",
			Label:       "",
			Description: "",
		},
		PolitiesPhilippines: &PolitiesCode{
			Name:        "Philippines",
			Label:       "",
			Description: "",
		},
		PolitiesPitcairnIslands: &PolitiesCode{
			Name:        "Pitcairn Islands",
			Label:       "",
			Description: "",
		},
		PolitiesPoland: &PolitiesCode{
			Name:        "Poland",
			Label:       "",
			Description: "",
		},
		PolitiesPortugal: &PolitiesCode{
			Name:        "Portugal",
			Label:       "",
			Description: "",
		},
		PolitiesPuertoRico: &PolitiesCode{
			Name:        "Puerto Rico",
			Label:       "",
			Description: "",
		},
		PolitiesQatar: &PolitiesCode{
			Name:        "Qatar",
			Label:       "",
			Description: "",
		},
		PolitiesRéunion: &PolitiesCode{
			Name:        "Réunion",
			Label:       "",
			Description: "",
		},
		PolitiesRomania: &PolitiesCode{
			Name:        "Romania",
			Label:       "",
			Description: "",
		},
		PolitiesRussia: &PolitiesCode{
			Name:        "Russia",
			Label:       "",
			Description: "",
		},
		PolitiesRwanda: &PolitiesCode{
			Name:        "Rwanda",
			Label:       "",
			Description: "",
		},
		PolitiesSaintHelenaAscensionandTristandaCunha: &PolitiesCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Label:       "",
			Description: "",
		},
		PolitiesSaintKittsandNevis: &PolitiesCode{
			Name:        "Saint Kitts and Nevis",
			Label:       "",
			Description: "",
		},
		PolitiesSaintLucia: &PolitiesCode{
			Name:        "Saint Lucia",
			Label:       "",
			Description: "",
		},
		PolitiesSaintPierreandMiquelon: &PolitiesCode{
			Name:        "Saint Pierre and Miquelon",
			Label:       "",
			Description: "",
		},
		PolitiesSaintVincentandtheGrenadines: &PolitiesCode{
			Name:        "Saint Vincent and the Grenadines",
			Label:       "",
			Description: "",
		},
		PolitiesSaintBarthelemy: &PolitiesCode{
			Name:        "Saint-Barthelemy",
			Label:       "",
			Description: "",
		},
		PolitiesSaintMartin: &PolitiesCode{
			Name:        "Saint-Martin",
			Label:       "",
			Description: "",
		},
		PolitiesSamoa: &PolitiesCode{
			Name:        "Samoa",
			Label:       "",
			Description: "",
		},
		PolitiesSanMarino: &PolitiesCode{
			Name:        "San Marino",
			Label:       "",
			Description: "",
		},
		PolitiesSãoToméandPríncipe: &PolitiesCode{
			Name:        "São Tomé and Príncipe",
			Label:       "",
			Description: "",
		},
		PolitiesSaudiArabia: &PolitiesCode{
			Name:        "Saudi Arabia",
			Label:       "",
			Description: "",
		},
		PolitiesSenegal: &PolitiesCode{
			Name:        "Senegal",
			Label:       "",
			Description: "",
		},
		PolitiesSerbia: &PolitiesCode{
			Name:        "Serbia",
			Label:       "",
			Description: "",
		},
		PolitiesSeychelles: &PolitiesCode{
			Name:        "Seychelles",
			Label:       "",
			Description: "",
		},
		PolitiesSierraLeone: &PolitiesCode{
			Name:        "Sierra Leone",
			Label:       "",
			Description: "",
		},
		PolitiesSingapore: &PolitiesCode{
			Name:        "Singapore",
			Label:       "",
			Description: "",
		},
		PolitiesSintMaarten: &PolitiesCode{
			Name:        "Sint Maarten",
			Label:       "",
			Description: "",
		},
		PolitiesSlovakia: &PolitiesCode{
			Name:        "Slovakia",
			Label:       "",
			Description: "",
		},
		PolitiesSlovenia: &PolitiesCode{
			Name:        "Slovenia",
			Label:       "",
			Description: "",
		},
		PolitiesSolomonIslands: &PolitiesCode{
			Name:        "Solomon Islands",
			Label:       "",
			Description: "",
		},
		PolitiesSomalia: &PolitiesCode{
			Name:        "Somalia",
			Label:       "",
			Description: "",
		},
		PolitiesSouthAfrica: &PolitiesCode{
			Name:        "South Africa",
			Label:       "",
			Description: "",
		},
		PolitiesSouthGeorgiaandtheSouthSandwichIslands: &PolitiesCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Label:       "",
			Description: "",
		},
		PolitiesSouthKorea: &PolitiesCode{
			Name:        "South Korea",
			Label:       "",
			Description: "",
		},
		PolitiesSouthSudan: &PolitiesCode{
			Name:        "South Sudan",
			Label:       "",
			Description: "",
		},
		PolitiesSpain: &PolitiesCode{
			Name:        "Spain",
			Label:       "",
			Description: "",
		},
		PolitiesSriLanka: &PolitiesCode{
			Name:        "Sri Lanka",
			Label:       "",
			Description: "",
		},
		PolitiesSudan: &PolitiesCode{
			Name:        "Sudan",
			Label:       "",
			Description: "",
		},
		PolitiesSuriname: &PolitiesCode{
			Name:        "Suriname",
			Label:       "",
			Description: "",
		},
		PolitiesSvalbardandJanMayen: &PolitiesCode{
			Name:        "Svalbard and Jan Mayen",
			Label:       "",
			Description: "",
		},
		PolitiesSwaziland: &PolitiesCode{
			Name:        "Swaziland",
			Label:       "",
			Description: "",
		},
		PolitiesSweden: &PolitiesCode{
			Name:        "Sweden",
			Label:       "",
			Description: "",
		},
		PolitiesSwitzerland: &PolitiesCode{
			Name:        "Switzerland",
			Label:       "",
			Description: "",
		},
		PolitiesSyria: &PolitiesCode{
			Name:        "Syria",
			Label:       "",
			Description: "",
		},
		PolitiesTaiwan: &PolitiesCode{
			Name:        "Taiwan",
			Label:       "",
			Description: "",
		},
		PolitiesTajikistan: &PolitiesCode{
			Name:        "Tajikistan",
			Label:       "",
			Description: "",
		},
		PolitiesTanzania: &PolitiesCode{
			Name:        "Tanzania",
			Label:       "",
			Description: "",
		},
		PolitiesThailand: &PolitiesCode{
			Name:        "Thailand",
			Label:       "",
			Description: "",
		},
		PolitiesTogo: &PolitiesCode{
			Name:        "Togo",
			Label:       "",
			Description: "",
		},
		PolitiesTokelau: &PolitiesCode{
			Name:        "Tokelau",
			Label:       "",
			Description: "",
		},
		PolitiesTonga: &PolitiesCode{
			Name:        "Tonga",
			Label:       "",
			Description: "",
		},
		PolitiesTrinidadandTobago: &PolitiesCode{
			Name:        "Trinidad and Tobago",
			Label:       "",
			Description: "",
		},
		PolitiesTunisia: &PolitiesCode{
			Name:        "Tunisia",
			Label:       "",
			Description: "",
		},
		PolitiesTurkey: &PolitiesCode{
			Name:        "Turkey",
			Label:       "",
			Description: "",
		},
		PolitiesTurkmenistan: &PolitiesCode{
			Name:        "Turkmenistan",
			Label:       "",
			Description: "",
		},
		PolitiesTurksandCaicosIslands: &PolitiesCode{
			Name:        "Turks and Caicos Islands",
			Label:       "",
			Description: "",
		},
		PolitiesTuvalu: &PolitiesCode{
			Name:        "Tuvalu",
			Label:       "",
			Description: "",
		},
		PolitiesUganda: &PolitiesCode{
			Name:        "Uganda",
			Label:       "",
			Description: "",
		},
		PolitiesUkraine: &PolitiesCode{
			Name:        "Ukraine",
			Label:       "",
			Description: "",
		},
		PolitiesUnitedArabEmirates: &PolitiesCode{
			Name:        "United Arab Emirates",
			Label:       "",
			Description: "",
		},
		PolitiesUnitedKingdom: &PolitiesCode{
			Name:        "United Kingdom",
			Label:       "",
			Description: "",
		},
		PolitiesUnitedStatesMinorOutlyingIslands: &PolitiesCode{
			Name:        "United States Minor Outlying Islands",
			Label:       "",
			Description: "",
		},
		PolitiesUruguay: &PolitiesCode{
			Name:        "Uruguay",
			Label:       "",
			Description: "",
		},
		PolitiesUSVirginIslands: &PolitiesCode{
			Name:        "US Virgin Islands",
			Label:       "",
			Description: "",
		},
		PolitiesUSA: &PolitiesCode{
			Name:        "USA",
			Label:       "",
			Description: "",
		},
		PolitiesUzbekistan: &PolitiesCode{
			Name:        "Uzbekistan",
			Label:       "",
			Description: "",
		},
		PolitiesVanuatu: &PolitiesCode{
			Name:        "Vanuatu",
			Label:       "",
			Description: "",
		},
		PolitiesVaticanCity: &PolitiesCode{
			Name:        "Vatican City",
			Label:       "",
			Description: "",
		},
		PolitiesVenezuela: &PolitiesCode{
			Name:        "Venezuela",
			Label:       "",
			Description: "",
		},
		PolitiesVietnam: &PolitiesCode{
			Name:        "Vietnam",
			Label:       "",
			Description: "",
		},
		PolitiesWallisandFutuna: &PolitiesCode{
			Name:        "Wallis and Futuna",
			Label:       "",
			Description: "",
		},
		PolitiesWesternSahara: &PolitiesCode{
			Name:        "Western Sahara",
			Label:       "",
			Description: "",
		},
		PolitiesYemen: &PolitiesCode{
			Name:        "Yemen",
			Label:       "",
			Description: "",
		},
		PolitiesZambia: &PolitiesCode{
			Name:        "Zambia",
			Label:       "",
			Description: "",
		},
		PolitiesZimbabwe: &PolitiesCode{
			Name:        "Zimbabwe",
			Label:       "",
			Description: "",
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
		},
		RolesAdvisor: &RolesCode{
			Name:        "Advisor",
			Label:       "",
			Description: "",
		},
		RolesAgent: &RolesCode{
			Name:        "Agent",
			Label:       "",
			Description: "",
		},
		RolesBeneficiary: &RolesCode{
			Name:        "Beneficiary",
			Label:       "",
			Description: "",
		},
		RolesCEO: &RolesCode{
			Name:        "CEO",
			Label:       "",
			Description: "",
		},
		RolesCFO: &RolesCode{
			Name:        "CFO",
			Label:       "",
			Description: "",
		},
		RolesChair: &RolesCode{
			Name:        "Chair",
			Label:       "",
			Description: "",
		},
		RolesCOO: &RolesCode{
			Name:        "COO",
			Label:       "",
			Description: "",
		},
		RolesCTO: &RolesCode{
			Name:        "CTO",
			Label:       "",
			Description: "",
		},
		RolesCustodian: &RolesCode{
			Name:        "Custodian",
			Label:       "",
			Description: "",
		},
		RolesDirector: &RolesCode{
			Name:        "Director",
			Label:       "",
			Description: "",
		},
		RolesExecutive: &RolesCode{
			Name:        "Executive",
			Label:       "",
			Description: "",
		},
		RolesLawyer: &RolesCode{
			Name:        "Lawyer",
			Label:       "",
			Description: "",
		},
		RolesLegalGuardian: &RolesCode{
			Name:        "Legal Guardian",
			Label:       "",
			Description: "",
		},
		RolesLimitedPartner: &RolesCode{
			Name:        "Limited Partner",
			Label:       "",
			Description: "",
		},
		RolesManager: &RolesCode{
			Name:        "Manager",
			Label:       "",
			Description: "",
		},
		RolesManagingPartner: &RolesCode{
			Name:        "Managing Partner",
			Label:       "",
			Description: "",
		},
		RolesMember: &RolesCode{
			Name:        "Member",
			Label:       "",
			Description: "Shareholder",
		},
		RolesPartner: &RolesCode{
			Name:        "Partner",
			Label:       "",
			Description: "",
		},
		RolesPrincipal: &RolesCode{
			Name:        "Principal",
			Label:       "",
			Description: "",
		},
		RolesProprietor: &RolesCode{
			Name:        "Proprietor",
			Label:       "",
			Description: "",
		},
		RolesProtector: &RolesCode{
			Name:        "Protector",
			Label:       "",
			Description: "",
		},
		RolesSecretary: &RolesCode{
			Name:        "Secretary",
			Label:       "",
			Description: "",
		},
		RolesSettlor: &RolesCode{
			Name:        "Settlor",
			Label:       "",
			Description: "",
		},
		RolesSignificantMember: &RolesCode{
			Name:        "Significant Member",
			Label:       "",
			Description: "Major Shareholder",
		},
		RolesSmartContractOperator: &RolesCode{
			Name:        "Smart Contract Operator",
			Label:       "",
			Description: "",
		},
		RolesTrader: &RolesCode{
			Name:        "Trader",
			Label:       "",
			Description: "",
		},
		RolesTrustee: &RolesCode{
			Name:        "Trustee",
			Label:       "",
			Description: "",
		},
		RolesUnitHolder: &RolesCode{
			Name:        "Unit Holder",
			Label:       "",
			Description: "",
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
	case RejectionsContractExists:
		return &RejectionsCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can't be recreated.",
			MetaData:    `{}`,
		}
	case RejectionsContractAssetQtyReduction:
		return &RejectionsCode{
			Name:        "ContractAssetQtyReduction",
			Label:       "Contract Asset Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
			MetaData:    `{}`,
		}
	case RejectionsContractFixedQuantity:
		return &RejectionsCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
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
	case RejectionsAssetCodeExists:
		return &RejectionsCode{
			Name:        "AssetCodeExists",
			Label:       "Asset Code Already Exists",
			Description: "The asset code specified already exists and can't be reused.",
			MetaData:    `{}`,
		}
	case RejectionsAssetNotFound:
		return &RejectionsCode{
			Name:        "AssetNotFound",
			Label:       "Asset Not Found",
			Description: "The asset code is not found.",
			MetaData:    `{}`,
		}
	case RejectionsAssetPermissions:
		return &RejectionsCode{
			Name:        "AssetPermissions",
			Label:       "Asset Permissions Prohibit",
			Description: "The asset permissions prohibit the action requested.",
			MetaData:    `{}`,
		}
	case RejectionsAssetFrozen:
		return &RejectionsCode{
			Name:        "AssetFrozen",
			Label:       "Asset Frozen",
			Description: "The asset is frozen and the request is not permitted while frozen.",
			MetaData:    `{}`,
		}
	case RejectionsAssetRevision:
		return &RejectionsCode{
			Name:        "AssetRevision",
			Label:       "Asset Revision Incorrect",
			Description: "The revision in an asset amendment is incorrect.",
			MetaData:    `{}`,
		}
	case RejectionsAssetNotPermitted:
		return &RejectionsCode{
			Name:        "AssetNotPermitted",
			Label:       "Asset Not Permitted",
			Description: "Action not permitted by asset.",
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
		},
		RejectionsMsgMalformed: &RejectionsCode{
			Name:        "MsgMalformed",
			Label:       "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.",
		},
		RejectionsTxMalformed: &RejectionsCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
		},
		RejectionsTimeout: &RejectionsCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract's timeout.",
		},
		RejectionsContractMoved: &RejectionsCode{
			Name:        "ContractMoved",
			Label:       "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
		},
		RejectionsDoubleSpend: &RejectionsCode{
			Name:        "DoubleSpend",
			Label:       "Double Spend",
			Description: "A double spend attempt has been detected.",
		},
		RejectionsContractExists: &RejectionsCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can't be recreated.",
		},
		RejectionsContractAssetQtyReduction: &RejectionsCode{
			Name:        "ContractAssetQtyReduction",
			Label:       "Contract Asset Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
		},
		RejectionsContractFixedQuantity: &RejectionsCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
		},
		RejectionsContractPermissions: &RejectionsCode{
			Name:        "ContractPermissions",
			Label:       "Contract Permissions Prohibit",
			Description: "The contract permissions prohibit the action requested.",
		},
		RejectionsContractExpired: &RejectionsCode{
			Name:        "ContractExpired",
			Label:       "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
		},
		RejectionsContractFrozen: &RejectionsCode{
			Name:        "ContractFrozen",
			Label:       "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
		},
		RejectionsContractRevision: &RejectionsCode{
			Name:        "ContractRevision",
			Label:       "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
		},
		RejectionsContractNotPermitted: &RejectionsCode{
			Name:        "ContractNotPermitted",
			Label:       "Contract Not Permitted",
			Description: "Action not permitted by contract.",
		},
		RejectionsContractBothOperatorsRequired: &RejectionsCode{
			Name:        "ContractBothOperatorsRequired",
			Label:       "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
		},
		RejectionsAssetCodeExists: &RejectionsCode{
			Name:        "AssetCodeExists",
			Label:       "Asset Code Already Exists",
			Description: "The asset code specified already exists and can't be reused.",
		},
		RejectionsAssetNotFound: &RejectionsCode{
			Name:        "AssetNotFound",
			Label:       "Asset Not Found",
			Description: "The asset code is not found.",
		},
		RejectionsAssetPermissions: &RejectionsCode{
			Name:        "AssetPermissions",
			Label:       "Asset Permissions Prohibit",
			Description: "The asset permissions prohibit the action requested.",
		},
		RejectionsAssetFrozen: &RejectionsCode{
			Name:        "AssetFrozen",
			Label:       "Asset Frozen",
			Description: "The asset is frozen and the request is not permitted while frozen.",
		},
		RejectionsAssetRevision: &RejectionsCode{
			Name:        "AssetRevision",
			Label:       "Asset Revision Incorrect",
			Description: "The revision in an asset amendment is incorrect.",
		},
		RejectionsAssetNotPermitted: &RejectionsCode{
			Name:        "AssetNotPermitted",
			Label:       "Asset Not Permitted",
			Description: "Action not permitted by asset.",
		},
		RejectionsTransferSelf: &RejectionsCode{
			Name:        "TransferSelf",
			Label:       "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
		},
		RejectionsTransferExpired: &RejectionsCode{
			Name:        "TransferExpired",
			Label:       "Transfer Expired",
			Description: "The transfer has expired.",
		},
		RejectionsHoldingsFrozen: &RejectionsCode{
			Name:        "HoldingsFrozen",
			Label:       "Holdings Frozen",
			Description: "Holdings are frozen, so the request can't be completed.",
		},
		RejectionsHolderProposalProhibited: &RejectionsCode{
			Name:        "HolderProposalProhibited",
			Label:       "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
		},
		RejectionsProposalConflicts: &RejectionsCode{
			Name:        "ProposalConflicts",
			Label:       "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
		},
		RejectionsVoteNotFound: &RejectionsCode{
			Name:        "VoteNotFound",
			Label:       "Vote Not Found",
			Description: "The vote ID referenced is not found.",
		},
		RejectionsVoteClosed: &RejectionsCode{
			Name:        "VoteClosed",
			Label:       "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
		},
		RejectionsBallotAlreadyCounted: &RejectionsCode{
			Name:        "BallotAlreadyCounted",
			Label:       "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
		},
		RejectionsVoteSystemNotPermitted: &RejectionsCode{
			Name:        "VoteSystemNotPermitted",
			Label:       "Vote System Not Permitted",
			Description: "The voting system isn't permitted for this request.",
		},
		RejectionsInsufficientTxFeeFunding: &RejectionsCode{
			Name:        "InsufficientTxFeeFunding",
			Label:       "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
		},
		RejectionsInsufficientValue: &RejectionsCode{
			Name:        "InsufficientValue",
			Label:       "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
		},
		RejectionsInsufficientQuantity: &RejectionsCode{
			Name:        "InsufficientQuantity",
			Label:       "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
		},
		RejectionsNotAdministration: &RejectionsCode{
			Name:        "NotAdministration",
			Label:       "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
		},
		RejectionsNotOperator: &RejectionsCode{
			Name:        "NotOperator",
			Label:       "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
		},
		RejectionsUnauthorizedAddress: &RejectionsCode{
			Name:        "UnauthorizedAddress",
			Label:       "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
		},
		RejectionsInvalidSignature: &RejectionsCode{
			Name:        "InvalidSignature",
			Label:       "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
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
		},
		TagsUtilities: &TagsCode{
			Name:        "Utilities",
			Label:       "",
			Description: "",
		},
		TagsFood: &TagsCode{
			Name:        "Food",
			Label:       "",
			Description: "",
		},
		TagsMedical: &TagsCode{
			Name:        "Medical",
			Label:       "",
			Description: "",
		},
		TagsFinancialServices: &TagsCode{
			Name:        "Financial Services",
			Label:       "",
			Description: "",
		},
		TagsEntertainment: &TagsCode{
			Name:        "Entertainment",
			Label:       "",
			Description: "",
		},
		TagsSales: &TagsCode{
			Name:        "Sales",
			Label:       "",
			Description: "",
		},
		TagsAutomotive: &TagsCode{
			Name:        "Automotive",
			Label:       "",
			Description: "",
		},
		TagsTransportation: &TagsCode{
			Name:        "Transportation",
			Label:       "",
			Description: "",
		},
		TagsFitness: &TagsCode{
			Name:        "Fitness",
			Label:       "",
			Description: "",
		},
		TagsElectricity: &TagsCode{
			Name:        "Electricity",
			Label:       "",
			Description: "",
		},
		TagsWater: &TagsCode{
			Name:        "Water",
			Label:       "",
			Description: "",
		},
		TagsInternet: &TagsCode{
			Name:        "Internet",
			Label:       "",
			Description: "",
		},
		TagsMedicine: &TagsCode{
			Name:        "Medicine",
			Label:       "",
			Description: "",
		},
		TagsService: &TagsCode{
			Name:        "Service",
			Label:       "",
			Description: "",
		},
		TagsRepair: &TagsCode{
			Name:        "Repair",
			Label:       "",
			Description: "",
		},
		TagsSupplies: &TagsCode{
			Name:        "Supplies",
			Label:       "",
			Description: "",
		},
		TagsParts: &TagsCode{
			Name:        "Parts",
			Label:       "",
			Description: "",
		},
		TagsLabor: &TagsCode{
			Name:        "Labor",
			Label:       "",
			Description: "",
		},
		TagsTip: &TagsCode{
			Name:        "Tip",
			Label:       "",
			Description: "",
		},
		TagsMedia: &TagsCode{
			Name:        "Media",
			Label:       "",
			Description: "",
		},
		TagsMusic: &TagsCode{
			Name:        "Music",
			Label:       "",
			Description: "",
		},
		TagsVideo: &TagsCode{
			Name:        "Video",
			Label:       "",
			Description: "",
		},
		TagsPhoto: &TagsCode{
			Name:        "Photo",
			Label:       "",
			Description: "",
		},
		TagsAudio: &TagsCode{
			Name:        "Audio",
			Label:       "",
			Description: "",
		},
		TagsAlcohol: &TagsCode{
			Name:        "Alcohol",
			Label:       "",
			Description: "",
		},
		TagsTobacco: &TagsCode{
			Name:        "Tobacco",
			Label:       "",
			Description: "",
		},
		TagsDiscounted: &TagsCode{
			Name:        "Discounted",
			Label:       "",
			Description: "",
		},
		TagsPromotional: &TagsCode{
			Name:        "Promotional",
			Label:       "",
			Description: "",
		},
	}
}
