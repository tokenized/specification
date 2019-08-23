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
}

// EntitiesData holds a mapping of Entities codes.
func EntitiesData(code string) *EntitiesCode {
	switch code {
	case EntitiesIndividual:
		return &EntitiesCode{
			Name:        "Individual",
			Label:       "Individual",
			Description: "",
		}
	case EntitiesPublicCompany:
		return &EntitiesCode{
			Name:        "PublicCompany",
			Label:       "Public Company Limited by Shares",
			Description: "",
		}
	case EntitiesPrivateCompany:
		return &EntitiesCode{
			Name:        "PrivateCompany",
			Label:       "Private Company Limited by Shares",
			Description: "",
		}
	case EntitiesLimitedPartnership:
		return &EntitiesCode{
			Name:        "LimitedPartnership",
			Label:       "Limited Partnership",
			Description: "",
		}
	case EntitiesUnlimitedPartnership:
		return &EntitiesCode{
			Name:        "UnlimitedPartnership",
			Label:       "Unlimited Partnership",
			Description: "",
		}
	case EntitiesSoleProprietorship:
		return &EntitiesCode{
			Name:        "SoleProprietorship",
			Label:       "Sole Proprietorship",
			Description: "",
		}
	case EntitiesStatutoryCompany:
		return &EntitiesCode{
			Name:        "StatutoryCompany",
			Label:       "Statutory Company",
			Description: "",
		}
	case EntitiesNonProfitOrganization:
		return &EntitiesCode{
			Name:        "NonProfitOrganization",
			Label:       "Non-Profit Organization",
			Description: "",
		}
	case EntitiesNationState:
		return &EntitiesCode{
			Name:        "NationState",
			Label:       "Nation State",
			Description: "",
		}
	case EntitiesGovernmentAgency:
		return &EntitiesCode{
			Name:        "GovernmentAgency",
			Label:       "Government Agency",
			Description: "",
		}
	case EntitiesUnitTrust:
		return &EntitiesCode{
			Name:        "UnitTrust",
			Label:       "Unit Trust",
			Description: "",
		}
	case EntitiesDiscretionaryTrust:
		return &EntitiesCode{
			Name:        "DiscretionaryTrust",
			Label:       "Discretionary Trust",
			Description: "",
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
}

// PolitiesData holds a mapping of Polities codes.
func PolitiesData(code string) *PolitiesCode {
	switch code {
	case PolitiesAalandIslands:
		return &PolitiesCode{
			Name:        "Aaland Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesAfghanistan:
		return &PolitiesCode{
			Name:        "Afghanistan",
			Label:       "",
			Description: "",
		}
	case PolitiesAlbania:
		return &PolitiesCode{
			Name:        "Albania",
			Label:       "",
			Description: "",
		}
	case PolitiesAlgeria:
		return &PolitiesCode{
			Name:        "Algeria",
			Label:       "",
			Description: "",
		}
	case PolitiesAmericanSamoa:
		return &PolitiesCode{
			Name:        "American Samoa",
			Label:       "",
			Description: "",
		}
	case PolitiesAndorra:
		return &PolitiesCode{
			Name:        "Andorra",
			Label:       "",
			Description: "",
		}
	case PolitiesAngola:
		return &PolitiesCode{
			Name:        "Angola",
			Label:       "",
			Description: "",
		}
	case PolitiesAnguilla:
		return &PolitiesCode{
			Name:        "Anguilla",
			Label:       "",
			Description: "",
		}
	case PolitiesAntarctica:
		return &PolitiesCode{
			Name:        "Antarctica",
			Label:       "",
			Description: "",
		}
	case PolitiesAntiguaandBarbuda:
		return &PolitiesCode{
			Name:        "Antigua and Barbuda",
			Label:       "",
			Description: "",
		}
	case PolitiesArgentina:
		return &PolitiesCode{
			Name:        "Argentina",
			Label:       "",
			Description: "",
		}
	case PolitiesArmenia:
		return &PolitiesCode{
			Name:        "Armenia",
			Label:       "",
			Description: "",
		}
	case PolitiesAruba:
		return &PolitiesCode{
			Name:        "Aruba",
			Label:       "",
			Description: "",
		}
	case PolitiesAfricanUnion:
		return &PolitiesCode{
			Name:        "African Union",
			Label:       "",
			Description: "",
		}
	case PolitiesAustralia:
		return &PolitiesCode{
			Name:        "Australia",
			Label:       "",
			Description: "",
		}
	case PolitiesAustria:
		return &PolitiesCode{
			Name:        "Austria",
			Label:       "",
			Description: "",
		}
	case PolitiesAzerbaijan:
		return &PolitiesCode{
			Name:        "Azerbaijan",
			Label:       "",
			Description: "",
		}
	case PolitiesTheBahamas:
		return &PolitiesCode{
			Name:        "The Bahamas",
			Label:       "",
			Description: "",
		}
	case PolitiesBahrain:
		return &PolitiesCode{
			Name:        "Bahrain",
			Label:       "",
			Description: "",
		}
	case PolitiesBangladesh:
		return &PolitiesCode{
			Name:        "Bangladesh",
			Label:       "",
			Description: "",
		}
	case PolitiesBarbados:
		return &PolitiesCode{
			Name:        "Barbados",
			Label:       "",
			Description: "",
		}
	case PolitiesBelarus:
		return &PolitiesCode{
			Name:        "Belarus",
			Label:       "",
			Description: "",
		}
	case PolitiesBelgium:
		return &PolitiesCode{
			Name:        "Belgium",
			Label:       "",
			Description: "",
		}
	case PolitiesBelize:
		return &PolitiesCode{
			Name:        "Belize",
			Label:       "",
			Description: "",
		}
	case PolitiesBenin:
		return &PolitiesCode{
			Name:        "Benin",
			Label:       "",
			Description: "",
		}
	case PolitiesBermuda:
		return &PolitiesCode{
			Name:        "Bermuda",
			Label:       "",
			Description: "",
		}
	case PolitiesBhutan:
		return &PolitiesCode{
			Name:        "Bhutan",
			Label:       "",
			Description: "",
		}
	case PolitiesBolivia:
		return &PolitiesCode{
			Name:        "Bolivia",
			Label:       "",
			Description: "",
		}
	case PolitiesBonaireStEustasuisandSaba:
		return &PolitiesCode{
			Name:        "Bonaire, St Eustasuis and Saba",
			Label:       "",
			Description: "",
		}
	case PolitiesBosniaandHerzegovina:
		return &PolitiesCode{
			Name:        "Bosnia and Herzegovina",
			Label:       "",
			Description: "",
		}
	case PolitiesBotswana:
		return &PolitiesCode{
			Name:        "Botswana",
			Label:       "",
			Description: "",
		}
	case PolitiesBouvetIsland:
		return &PolitiesCode{
			Name:        "Bouvet Island",
			Label:       "",
			Description: "",
		}
	case PolitiesBrazil:
		return &PolitiesCode{
			Name:        "Brazil",
			Label:       "",
			Description: "",
		}
	case PolitiesBritishIndianOceanTerritory:
		return &PolitiesCode{
			Name:        "British Indian Ocean Territory",
			Label:       "",
			Description: "",
		}
	case PolitiesBritishVirginIslands:
		return &PolitiesCode{
			Name:        "British Virgin Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesBrunei:
		return &PolitiesCode{
			Name:        "Brunei",
			Label:       "",
			Description: "",
		}
	case PolitiesBulgaria:
		return &PolitiesCode{
			Name:        "Bulgaria",
			Label:       "",
			Description: "",
		}
	case PolitiesBurkinaFaso:
		return &PolitiesCode{
			Name:        "Burkina Faso",
			Label:       "",
			Description: "",
		}
	case PolitiesBurundi:
		return &PolitiesCode{
			Name:        "Burundi",
			Label:       "",
			Description: "",
		}
	case PolitiesCambodia:
		return &PolitiesCode{
			Name:        "Cambodia",
			Label:       "",
			Description: "",
		}
	case PolitiesCameroon:
		return &PolitiesCode{
			Name:        "Cameroon",
			Label:       "",
			Description: "",
		}
	case PolitiesCanada:
		return &PolitiesCode{
			Name:        "Canada",
			Label:       "",
			Description: "",
		}
	case PolitiesCapeVerde:
		return &PolitiesCode{
			Name:        "Cape Verde",
			Label:       "",
			Description: "",
		}
	case PolitiesCaymanIslands:
		return &PolitiesCode{
			Name:        "Cayman Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesCentralAfricanRepublic:
		return &PolitiesCode{
			Name:        "Central African Republic",
			Label:       "",
			Description: "",
		}
	case PolitiesChad:
		return &PolitiesCode{
			Name:        "Chad",
			Label:       "",
			Description: "",
		}
	case PolitiesChile:
		return &PolitiesCode{
			Name:        "Chile",
			Label:       "",
			Description: "",
		}
	case PolitiesChina:
		return &PolitiesCode{
			Name:        "China",
			Label:       "",
			Description: "",
		}
	case PolitiesChristmasIsland:
		return &PolitiesCode{
			Name:        "Christmas Island",
			Label:       "",
			Description: "",
		}
	case PolitiesCocosIslands:
		return &PolitiesCode{
			Name:        "Cocos Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesColombia:
		return &PolitiesCode{
			Name:        "Colombia",
			Label:       "",
			Description: "",
		}
	case PolitiesComoros:
		return &PolitiesCode{
			Name:        "Comoros",
			Label:       "",
			Description: "",
		}
	case PolitiesCongoBrazzaville:
		return &PolitiesCode{
			Name:        "Congo-Brazzaville",
			Label:       "",
			Description: "",
		}
	case PolitiesCongoKinshasa:
		return &PolitiesCode{
			Name:        "Congo-Kinshasa",
			Label:       "",
			Description: "",
		}
	case PolitiesCookIslands:
		return &PolitiesCode{
			Name:        "Cook Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesCostaRica:
		return &PolitiesCode{
			Name:        "Costa Rica",
			Label:       "",
			Description: "",
		}
	case PolitiesCroatia:
		return &PolitiesCode{
			Name:        "Croatia",
			Label:       "",
			Description: "",
		}
	case PolitiesCuba:
		return &PolitiesCode{
			Name:        "Cuba",
			Label:       "",
			Description: "",
		}
	case PolitiesCuracao:
		return &PolitiesCode{
			Name:        "Curacao",
			Label:       "",
			Description: "",
		}
	case PolitiesCyprus:
		return &PolitiesCode{
			Name:        "Cyprus",
			Label:       "",
			Description: "",
		}
	case PolitiesCzechRepublic:
		return &PolitiesCode{
			Name:        "Czech Republic",
			Label:       "",
			Description: "",
		}
	case PolitiesDenmark:
		return &PolitiesCode{
			Name:        "Denmark",
			Label:       "",
			Description: "",
		}
	case PolitiesDjibouti:
		return &PolitiesCode{
			Name:        "Djibouti",
			Label:       "",
			Description: "",
		}
	case PolitiesDominica:
		return &PolitiesCode{
			Name:        "Dominica",
			Label:       "",
			Description: "",
		}
	case PolitiesDominicanRepublic:
		return &PolitiesCode{
			Name:        "Dominican Republic",
			Label:       "",
			Description: "",
		}
	case PolitiesEastTimor:
		return &PolitiesCode{
			Name:        "East Timor",
			Label:       "",
			Description: "",
		}
	case PolitiesEcuador:
		return &PolitiesCode{
			Name:        "Ecuador",
			Label:       "",
			Description: "",
		}
	case PolitiesEgypt:
		return &PolitiesCode{
			Name:        "Egypt",
			Label:       "",
			Description: "",
		}
	case PolitiesElSalvador:
		return &PolitiesCode{
			Name:        "El Salvador",
			Label:       "",
			Description: "",
		}
	case PolitiesEquatorialGuinea:
		return &PolitiesCode{
			Name:        "Equatorial Guinea",
			Label:       "",
			Description: "",
		}
	case PolitiesEritrea:
		return &PolitiesCode{
			Name:        "Eritrea",
			Label:       "",
			Description: "",
		}
	case PolitiesEstonia:
		return &PolitiesCode{
			Name:        "Estonia",
			Label:       "",
			Description: "",
		}
	case PolitiesEthiopia:
		return &PolitiesCode{
			Name:        "Ethiopia",
			Label:       "",
			Description: "",
		}
	case PolitiesEuropeanUnion:
		return &PolitiesCode{
			Name:        "European Union",
			Label:       "",
			Description: "",
		}
	case PolitiesFalklandIslands:
		return &PolitiesCode{
			Name:        "Falkland Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesFaroeIslands:
		return &PolitiesCode{
			Name:        "Faroe Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesFiji:
		return &PolitiesCode{
			Name:        "Fiji",
			Label:       "",
			Description: "",
		}
	case PolitiesFinland:
		return &PolitiesCode{
			Name:        "Finland",
			Label:       "",
			Description: "",
		}
	case PolitiesFrance:
		return &PolitiesCode{
			Name:        "France",
			Label:       "",
			Description: "",
		}
	case PolitiesFrenchGuiana:
		return &PolitiesCode{
			Name:        "French Guiana",
			Label:       "",
			Description: "",
		}
	case PolitiesFrenchPolynesia:
		return &PolitiesCode{
			Name:        "French Polynesia",
			Label:       "",
			Description: "",
		}
	case PolitiesFrenchSouthernandAntarcticLands:
		return &PolitiesCode{
			Name:        "French Southern and Antarctic Lands",
			Label:       "",
			Description: "",
		}
	case PolitiesGabon:
		return &PolitiesCode{
			Name:        "Gabon",
			Label:       "",
			Description: "",
		}
	case PolitiesTheGambia:
		return &PolitiesCode{
			Name:        "The Gambia",
			Label:       "",
			Description: "",
		}
	case PolitiesGeorgia:
		return &PolitiesCode{
			Name:        "Georgia",
			Label:       "",
			Description: "",
		}
	case PolitiesGermany:
		return &PolitiesCode{
			Name:        "Germany",
			Label:       "",
			Description: "",
		}
	case PolitiesGhana:
		return &PolitiesCode{
			Name:        "Ghana",
			Label:       "",
			Description: "",
		}
	case PolitiesGibraltar:
		return &PolitiesCode{
			Name:        "Gibraltar",
			Label:       "",
			Description: "",
		}
	case PolitiesGreece:
		return &PolitiesCode{
			Name:        "Greece",
			Label:       "",
			Description: "",
		}
	case PolitiesGreenland:
		return &PolitiesCode{
			Name:        "Greenland",
			Label:       "",
			Description: "",
		}
	case PolitiesGrenada:
		return &PolitiesCode{
			Name:        "Grenada",
			Label:       "",
			Description: "",
		}
	case PolitiesGuadeloupe:
		return &PolitiesCode{
			Name:        "Guadeloupe",
			Label:       "",
			Description: "",
		}
	case PolitiesGuam:
		return &PolitiesCode{
			Name:        "Guam",
			Label:       "",
			Description: "",
		}
	case PolitiesGuatemala:
		return &PolitiesCode{
			Name:        "Guatemala",
			Label:       "",
			Description: "",
		}
	case PolitiesGuernsey:
		return &PolitiesCode{
			Name:        "Guernsey",
			Label:       "",
			Description: "",
		}
	case PolitiesGuinea:
		return &PolitiesCode{
			Name:        "Guinea",
			Label:       "",
			Description: "",
		}
	case PolitiesGuineaBissau:
		return &PolitiesCode{
			Name:        "Guinea-Bissau",
			Label:       "",
			Description: "",
		}
	case PolitiesGuyana:
		return &PolitiesCode{
			Name:        "Guyana",
			Label:       "",
			Description: "",
		}
	case PolitiesHaiti:
		return &PolitiesCode{
			Name:        "Haiti",
			Label:       "",
			Description: "",
		}
	case PolitiesHeardIslandandMcDonaldIslands:
		return &PolitiesCode{
			Name:        "Heard Island and McDonald Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesHonduras:
		return &PolitiesCode{
			Name:        "Honduras",
			Label:       "",
			Description: "",
		}
	case PolitiesHongKong:
		return &PolitiesCode{
			Name:        "Hong Kong",
			Label:       "",
			Description: "",
		}
	case PolitiesHungary:
		return &PolitiesCode{
			Name:        "Hungary",
			Label:       "",
			Description: "",
		}
	case PolitiesIceland:
		return &PolitiesCode{
			Name:        "Iceland",
			Label:       "",
			Description: "",
		}
	case PolitiesIndia:
		return &PolitiesCode{
			Name:        "India",
			Label:       "",
			Description: "",
		}
	case PolitiesIndonesia:
		return &PolitiesCode{
			Name:        "Indonesia",
			Label:       "",
			Description: "",
		}
	case PolitiesIran:
		return &PolitiesCode{
			Name:        "Iran",
			Label:       "",
			Description: "",
		}
	case PolitiesIraq:
		return &PolitiesCode{
			Name:        "Iraq",
			Label:       "",
			Description: "",
		}
	case PolitiesIreland:
		return &PolitiesCode{
			Name:        "Ireland",
			Label:       "",
			Description: "",
		}
	case PolitiesIsleofMan:
		return &PolitiesCode{
			Name:        "Isle of Man",
			Label:       "",
			Description: "",
		}
	case PolitiesIsrael:
		return &PolitiesCode{
			Name:        "Israel",
			Label:       "",
			Description: "",
		}
	case PolitiesItaly:
		return &PolitiesCode{
			Name:        "Italy",
			Label:       "",
			Description: "",
		}
	case PolitiesIvoryCoast:
		return &PolitiesCode{
			Name:        "Ivory Coast",
			Label:       "",
			Description: "",
		}
	case PolitiesJamaica:
		return &PolitiesCode{
			Name:        "Jamaica",
			Label:       "",
			Description: "",
		}
	case PolitiesJapan:
		return &PolitiesCode{
			Name:        "Japan",
			Label:       "",
			Description: "",
		}
	case PolitiesJersey:
		return &PolitiesCode{
			Name:        "Jersey",
			Label:       "",
			Description: "",
		}
	case PolitiesJordan:
		return &PolitiesCode{
			Name:        "Jordan",
			Label:       "",
			Description: "",
		}
	case PolitiesKazakhstan:
		return &PolitiesCode{
			Name:        "Kazakhstan",
			Label:       "",
			Description: "",
		}
	case PolitiesKenya:
		return &PolitiesCode{
			Name:        "Kenya",
			Label:       "",
			Description: "",
		}
	case PolitiesKiribati:
		return &PolitiesCode{
			Name:        "Kiribati",
			Label:       "",
			Description: "",
		}
	case PolitiesKuwait:
		return &PolitiesCode{
			Name:        "Kuwait",
			Label:       "",
			Description: "",
		}
	case PolitiesKyrgyzstan:
		return &PolitiesCode{
			Name:        "Kyrgyzstan",
			Label:       "",
			Description: "",
		}
	case PolitiesLaos:
		return &PolitiesCode{
			Name:        "Laos",
			Label:       "",
			Description: "",
		}
	case PolitiesLatvia:
		return &PolitiesCode{
			Name:        "Latvia",
			Label:       "",
			Description: "",
		}
	case PolitiesLebanon:
		return &PolitiesCode{
			Name:        "Lebanon",
			Label:       "",
			Description: "",
		}
	case PolitiesLesotho:
		return &PolitiesCode{
			Name:        "Lesotho",
			Label:       "",
			Description: "",
		}
	case PolitiesLiberia:
		return &PolitiesCode{
			Name:        "Liberia",
			Label:       "",
			Description: "",
		}
	case PolitiesLibya:
		return &PolitiesCode{
			Name:        "Libya",
			Label:       "",
			Description: "",
		}
	case PolitiesLiechtenstein:
		return &PolitiesCode{
			Name:        "Liechtenstein",
			Label:       "",
			Description: "",
		}
	case PolitiesLithuania:
		return &PolitiesCode{
			Name:        "Lithuania",
			Label:       "",
			Description: "",
		}
	case PolitiesLuxembourg:
		return &PolitiesCode{
			Name:        "Luxembourg",
			Label:       "",
			Description: "",
		}
	case PolitiesMacau:
		return &PolitiesCode{
			Name:        "Macau",
			Label:       "",
			Description: "",
		}
	case PolitiesMacedonia:
		return &PolitiesCode{
			Name:        "Macedonia",
			Label:       "",
			Description: "",
		}
	case PolitiesMadagascar:
		return &PolitiesCode{
			Name:        "Madagascar",
			Label:       "",
			Description: "",
		}
	case PolitiesMalawi:
		return &PolitiesCode{
			Name:        "Malawi",
			Label:       "",
			Description: "",
		}
	case PolitiesMalaysia:
		return &PolitiesCode{
			Name:        "Malaysia",
			Label:       "",
			Description: "",
		}
	case PolitiesMaldives:
		return &PolitiesCode{
			Name:        "Maldives",
			Label:       "",
			Description: "",
		}
	case PolitiesMali:
		return &PolitiesCode{
			Name:        "Mali",
			Label:       "",
			Description: "",
		}
	case PolitiesMalta:
		return &PolitiesCode{
			Name:        "Malta",
			Label:       "",
			Description: "",
		}
	case PolitiesMarshallIslands:
		return &PolitiesCode{
			Name:        "Marshall Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesMartinique:
		return &PolitiesCode{
			Name:        "Martinique",
			Label:       "",
			Description: "",
		}
	case PolitiesMauritania:
		return &PolitiesCode{
			Name:        "Mauritania",
			Label:       "",
			Description: "",
		}
	case PolitiesMauritius:
		return &PolitiesCode{
			Name:        "Mauritius",
			Label:       "",
			Description: "",
		}
	case PolitiesMayotte:
		return &PolitiesCode{
			Name:        "Mayotte",
			Label:       "",
			Description: "",
		}
	case PolitiesMexico:
		return &PolitiesCode{
			Name:        "Mexico",
			Label:       "",
			Description: "",
		}
	case PolitiesMicronesia:
		return &PolitiesCode{
			Name:        "Micronesia",
			Label:       "",
			Description: "",
		}
	case PolitiesMoldova:
		return &PolitiesCode{
			Name:        "Moldova",
			Label:       "",
			Description: "",
		}
	case PolitiesMonaco:
		return &PolitiesCode{
			Name:        "Monaco",
			Label:       "",
			Description: "",
		}
	case PolitiesMongolia:
		return &PolitiesCode{
			Name:        "Mongolia",
			Label:       "",
			Description: "",
		}
	case PolitiesMontenegro:
		return &PolitiesCode{
			Name:        "Montenegro",
			Label:       "",
			Description: "",
		}
	case PolitiesMontserrat:
		return &PolitiesCode{
			Name:        "Montserrat",
			Label:       "",
			Description: "",
		}
	case PolitiesMorocco:
		return &PolitiesCode{
			Name:        "Morocco",
			Label:       "",
			Description: "",
		}
	case PolitiesMozambique:
		return &PolitiesCode{
			Name:        "Mozambique",
			Label:       "",
			Description: "",
		}
	case PolitiesMyanmar:
		return &PolitiesCode{
			Name:        "Myanmar",
			Label:       "",
			Description: "",
		}
	case PolitiesNamibia:
		return &PolitiesCode{
			Name:        "Namibia",
			Label:       "",
			Description: "",
		}
	case PolitiesNauru:
		return &PolitiesCode{
			Name:        "Nauru",
			Label:       "",
			Description: "",
		}
	case PolitiesNepal:
		return &PolitiesCode{
			Name:        "Nepal",
			Label:       "",
			Description: "",
		}
	case PolitiesNetherlands:
		return &PolitiesCode{
			Name:        "Netherlands",
			Label:       "",
			Description: "",
		}
	case PolitiesNewCaledonia:
		return &PolitiesCode{
			Name:        "New Caledonia",
			Label:       "",
			Description: "",
		}
	case PolitiesNewZealand:
		return &PolitiesCode{
			Name:        "New Zealand",
			Label:       "",
			Description: "",
		}
	case PolitiesNicaragua:
		return &PolitiesCode{
			Name:        "Nicaragua",
			Label:       "",
			Description: "",
		}
	case PolitiesNiger:
		return &PolitiesCode{
			Name:        "Niger",
			Label:       "",
			Description: "",
		}
	case PolitiesNigeria:
		return &PolitiesCode{
			Name:        "Nigeria",
			Label:       "",
			Description: "",
		}
	case PolitiesNiue:
		return &PolitiesCode{
			Name:        "Niue",
			Label:       "",
			Description: "",
		}
	case PolitiesNorfolkandPhilipIsland:
		return &PolitiesCode{
			Name:        "Norfolk and Philip Island",
			Label:       "",
			Description: "",
		}
	case PolitiesNorthKorea:
		return &PolitiesCode{
			Name:        "North Korea",
			Label:       "",
			Description: "",
		}
	case PolitiesNorthernMarianaIslands:
		return &PolitiesCode{
			Name:        "Northern Mariana Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesNorway:
		return &PolitiesCode{
			Name:        "Norway",
			Label:       "",
			Description: "",
		}
	case PolitiesOman:
		return &PolitiesCode{
			Name:        "Oman",
			Label:       "",
			Description: "",
		}
	case PolitiesPakistan:
		return &PolitiesCode{
			Name:        "Pakistan",
			Label:       "",
			Description: "",
		}
	case PolitiesPalau:
		return &PolitiesCode{
			Name:        "Palau",
			Label:       "",
			Description: "",
		}
	case PolitiesPalestinianTerritory:
		return &PolitiesCode{
			Name:        "Palestinian Territory",
			Label:       "",
			Description: "",
		}
	case PolitiesPanama:
		return &PolitiesCode{
			Name:        "Panama",
			Label:       "",
			Description: "",
		}
	case PolitiesPapuaNewGuinea:
		return &PolitiesCode{
			Name:        "Papua New Guinea",
			Label:       "",
			Description: "",
		}
	case PolitiesParaguay:
		return &PolitiesCode{
			Name:        "Paraguay",
			Label:       "",
			Description: "",
		}
	case PolitiesPeru:
		return &PolitiesCode{
			Name:        "Peru",
			Label:       "",
			Description: "",
		}
	case PolitiesPhilippines:
		return &PolitiesCode{
			Name:        "Philippines",
			Label:       "",
			Description: "",
		}
	case PolitiesPitcairnIslands:
		return &PolitiesCode{
			Name:        "Pitcairn Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesPoland:
		return &PolitiesCode{
			Name:        "Poland",
			Label:       "",
			Description: "",
		}
	case PolitiesPortugal:
		return &PolitiesCode{
			Name:        "Portugal",
			Label:       "",
			Description: "",
		}
	case PolitiesPuertoRico:
		return &PolitiesCode{
			Name:        "Puerto Rico",
			Label:       "",
			Description: "",
		}
	case PolitiesQatar:
		return &PolitiesCode{
			Name:        "Qatar",
			Label:       "",
			Description: "",
		}
	case PolitiesRéunion:
		return &PolitiesCode{
			Name:        "Réunion",
			Label:       "",
			Description: "",
		}
	case PolitiesRomania:
		return &PolitiesCode{
			Name:        "Romania",
			Label:       "",
			Description: "",
		}
	case PolitiesRussia:
		return &PolitiesCode{
			Name:        "Russia",
			Label:       "",
			Description: "",
		}
	case PolitiesRwanda:
		return &PolitiesCode{
			Name:        "Rwanda",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintHelenaAscensionandTristandaCunha:
		return &PolitiesCode{
			Name:        "Saint Helena, Ascension and Tristan da Cunha",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintKittsandNevis:
		return &PolitiesCode{
			Name:        "Saint Kitts and Nevis",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintLucia:
		return &PolitiesCode{
			Name:        "Saint Lucia",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintPierreandMiquelon:
		return &PolitiesCode{
			Name:        "Saint Pierre and Miquelon",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintVincentandtheGrenadines:
		return &PolitiesCode{
			Name:        "Saint Vincent and the Grenadines",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintBarthelemy:
		return &PolitiesCode{
			Name:        "Saint-Barthelemy",
			Label:       "",
			Description: "",
		}
	case PolitiesSaintMartin:
		return &PolitiesCode{
			Name:        "Saint-Martin",
			Label:       "",
			Description: "",
		}
	case PolitiesSamoa:
		return &PolitiesCode{
			Name:        "Samoa",
			Label:       "",
			Description: "",
		}
	case PolitiesSanMarino:
		return &PolitiesCode{
			Name:        "San Marino",
			Label:       "",
			Description: "",
		}
	case PolitiesSãoToméandPríncipe:
		return &PolitiesCode{
			Name:        "São Tomé and Príncipe",
			Label:       "",
			Description: "",
		}
	case PolitiesSaudiArabia:
		return &PolitiesCode{
			Name:        "Saudi Arabia",
			Label:       "",
			Description: "",
		}
	case PolitiesSenegal:
		return &PolitiesCode{
			Name:        "Senegal",
			Label:       "",
			Description: "",
		}
	case PolitiesSerbia:
		return &PolitiesCode{
			Name:        "Serbia",
			Label:       "",
			Description: "",
		}
	case PolitiesSeychelles:
		return &PolitiesCode{
			Name:        "Seychelles",
			Label:       "",
			Description: "",
		}
	case PolitiesSierraLeone:
		return &PolitiesCode{
			Name:        "Sierra Leone",
			Label:       "",
			Description: "",
		}
	case PolitiesSingapore:
		return &PolitiesCode{
			Name:        "Singapore",
			Label:       "",
			Description: "",
		}
	case PolitiesSintMaarten:
		return &PolitiesCode{
			Name:        "Sint Maarten",
			Label:       "",
			Description: "",
		}
	case PolitiesSlovakia:
		return &PolitiesCode{
			Name:        "Slovakia",
			Label:       "",
			Description: "",
		}
	case PolitiesSlovenia:
		return &PolitiesCode{
			Name:        "Slovenia",
			Label:       "",
			Description: "",
		}
	case PolitiesSolomonIslands:
		return &PolitiesCode{
			Name:        "Solomon Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesSomalia:
		return &PolitiesCode{
			Name:        "Somalia",
			Label:       "",
			Description: "",
		}
	case PolitiesSouthAfrica:
		return &PolitiesCode{
			Name:        "South Africa",
			Label:       "",
			Description: "",
		}
	case PolitiesSouthGeorgiaandtheSouthSandwichIslands:
		return &PolitiesCode{
			Name:        "South Georgia and the South Sandwich Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesSouthKorea:
		return &PolitiesCode{
			Name:        "South Korea",
			Label:       "",
			Description: "",
		}
	case PolitiesSouthSudan:
		return &PolitiesCode{
			Name:        "South Sudan",
			Label:       "",
			Description: "",
		}
	case PolitiesSpain:
		return &PolitiesCode{
			Name:        "Spain",
			Label:       "",
			Description: "",
		}
	case PolitiesSriLanka:
		return &PolitiesCode{
			Name:        "Sri Lanka",
			Label:       "",
			Description: "",
		}
	case PolitiesSudan:
		return &PolitiesCode{
			Name:        "Sudan",
			Label:       "",
			Description: "",
		}
	case PolitiesSuriname:
		return &PolitiesCode{
			Name:        "Suriname",
			Label:       "",
			Description: "",
		}
	case PolitiesSvalbardandJanMayen:
		return &PolitiesCode{
			Name:        "Svalbard and Jan Mayen",
			Label:       "",
			Description: "",
		}
	case PolitiesSwaziland:
		return &PolitiesCode{
			Name:        "Swaziland",
			Label:       "",
			Description: "",
		}
	case PolitiesSweden:
		return &PolitiesCode{
			Name:        "Sweden",
			Label:       "",
			Description: "",
		}
	case PolitiesSwitzerland:
		return &PolitiesCode{
			Name:        "Switzerland",
			Label:       "",
			Description: "",
		}
	case PolitiesSyria:
		return &PolitiesCode{
			Name:        "Syria",
			Label:       "",
			Description: "",
		}
	case PolitiesTaiwan:
		return &PolitiesCode{
			Name:        "Taiwan",
			Label:       "",
			Description: "",
		}
	case PolitiesTajikistan:
		return &PolitiesCode{
			Name:        "Tajikistan",
			Label:       "",
			Description: "",
		}
	case PolitiesTanzania:
		return &PolitiesCode{
			Name:        "Tanzania",
			Label:       "",
			Description: "",
		}
	case PolitiesThailand:
		return &PolitiesCode{
			Name:        "Thailand",
			Label:       "",
			Description: "",
		}
	case PolitiesTogo:
		return &PolitiesCode{
			Name:        "Togo",
			Label:       "",
			Description: "",
		}
	case PolitiesTokelau:
		return &PolitiesCode{
			Name:        "Tokelau",
			Label:       "",
			Description: "",
		}
	case PolitiesTonga:
		return &PolitiesCode{
			Name:        "Tonga",
			Label:       "",
			Description: "",
		}
	case PolitiesTrinidadandTobago:
		return &PolitiesCode{
			Name:        "Trinidad and Tobago",
			Label:       "",
			Description: "",
		}
	case PolitiesTunisia:
		return &PolitiesCode{
			Name:        "Tunisia",
			Label:       "",
			Description: "",
		}
	case PolitiesTurkey:
		return &PolitiesCode{
			Name:        "Turkey",
			Label:       "",
			Description: "",
		}
	case PolitiesTurkmenistan:
		return &PolitiesCode{
			Name:        "Turkmenistan",
			Label:       "",
			Description: "",
		}
	case PolitiesTurksandCaicosIslands:
		return &PolitiesCode{
			Name:        "Turks and Caicos Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesTuvalu:
		return &PolitiesCode{
			Name:        "Tuvalu",
			Label:       "",
			Description: "",
		}
	case PolitiesUganda:
		return &PolitiesCode{
			Name:        "Uganda",
			Label:       "",
			Description: "",
		}
	case PolitiesUkraine:
		return &PolitiesCode{
			Name:        "Ukraine",
			Label:       "",
			Description: "",
		}
	case PolitiesUnitedArabEmirates:
		return &PolitiesCode{
			Name:        "United Arab Emirates",
			Label:       "",
			Description: "",
		}
	case PolitiesUnitedKingdom:
		return &PolitiesCode{
			Name:        "United Kingdom",
			Label:       "",
			Description: "",
		}
	case PolitiesUnitedStatesMinorOutlyingIslands:
		return &PolitiesCode{
			Name:        "United States Minor Outlying Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesUruguay:
		return &PolitiesCode{
			Name:        "Uruguay",
			Label:       "",
			Description: "",
		}
	case PolitiesUSVirginIslands:
		return &PolitiesCode{
			Name:        "US Virgin Islands",
			Label:       "",
			Description: "",
		}
	case PolitiesUSA:
		return &PolitiesCode{
			Name:        "USA",
			Label:       "",
			Description: "",
		}
	case PolitiesUzbekistan:
		return &PolitiesCode{
			Name:        "Uzbekistan",
			Label:       "",
			Description: "",
		}
	case PolitiesVanuatu:
		return &PolitiesCode{
			Name:        "Vanuatu",
			Label:       "",
			Description: "",
		}
	case PolitiesVaticanCity:
		return &PolitiesCode{
			Name:        "Vatican City",
			Label:       "",
			Description: "",
		}
	case PolitiesVenezuela:
		return &PolitiesCode{
			Name:        "Venezuela",
			Label:       "",
			Description: "",
		}
	case PolitiesVietnam:
		return &PolitiesCode{
			Name:        "Vietnam",
			Label:       "",
			Description: "",
		}
	case PolitiesWallisandFutuna:
		return &PolitiesCode{
			Name:        "Wallis and Futuna",
			Label:       "",
			Description: "",
		}
	case PolitiesWesternSahara:
		return &PolitiesCode{
			Name:        "Western Sahara",
			Label:       "",
			Description: "",
		}
	case PolitiesYemen:
		return &PolitiesCode{
			Name:        "Yemen",
			Label:       "",
			Description: "",
		}
	case PolitiesZambia:
		return &PolitiesCode{
			Name:        "Zambia",
			Label:       "",
			Description: "",
		}
	case PolitiesZimbabwe:
		return &PolitiesCode{
			Name:        "Zimbabwe",
			Label:       "",
			Description: "",
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
}

// RolesData holds a mapping of Roles codes.
func RolesData(code uint32) *RolesCode {
	switch code {
	case RolesAccountant:
		return &RolesCode{
			Name:        "Accountant",
			Label:       "",
			Description: "",
		}
	case RolesAdvisor:
		return &RolesCode{
			Name:        "Advisor",
			Label:       "",
			Description: "",
		}
	case RolesAgent:
		return &RolesCode{
			Name:        "Agent",
			Label:       "",
			Description: "",
		}
	case RolesBeneficiary:
		return &RolesCode{
			Name:        "Beneficiary",
			Label:       "",
			Description: "",
		}
	case RolesCEO:
		return &RolesCode{
			Name:        "CEO",
			Label:       "",
			Description: "",
		}
	case RolesCFO:
		return &RolesCode{
			Name:        "CFO",
			Label:       "",
			Description: "",
		}
	case RolesChair:
		return &RolesCode{
			Name:        "Chair",
			Label:       "",
			Description: "",
		}
	case RolesCOO:
		return &RolesCode{
			Name:        "COO",
			Label:       "",
			Description: "",
		}
	case RolesCTO:
		return &RolesCode{
			Name:        "CTO",
			Label:       "",
			Description: "",
		}
	case RolesCustodian:
		return &RolesCode{
			Name:        "Custodian",
			Label:       "",
			Description: "",
		}
	case RolesDirector:
		return &RolesCode{
			Name:        "Director",
			Label:       "",
			Description: "",
		}
	case RolesExecutive:
		return &RolesCode{
			Name:        "Executive",
			Label:       "",
			Description: "",
		}
	case RolesLawyer:
		return &RolesCode{
			Name:        "Lawyer",
			Label:       "",
			Description: "",
		}
	case RolesLegalGuardian:
		return &RolesCode{
			Name:        "Legal Guardian",
			Label:       "",
			Description: "",
		}
	case RolesLimitedPartner:
		return &RolesCode{
			Name:        "Limited Partner",
			Label:       "",
			Description: "",
		}
	case RolesManager:
		return &RolesCode{
			Name:        "Manager",
			Label:       "",
			Description: "",
		}
	case RolesManagingPartner:
		return &RolesCode{
			Name:        "Managing Partner",
			Label:       "",
			Description: "",
		}
	case RolesMember:
		return &RolesCode{
			Name:        "Member",
			Label:       "",
			Description: "Shareholder",
		}
	case RolesPartner:
		return &RolesCode{
			Name:        "Partner",
			Label:       "",
			Description: "",
		}
	case RolesPrincipal:
		return &RolesCode{
			Name:        "Principal",
			Label:       "",
			Description: "",
		}
	case RolesProprietor:
		return &RolesCode{
			Name:        "Proprietor",
			Label:       "",
			Description: "",
		}
	case RolesProtector:
		return &RolesCode{
			Name:        "Protector",
			Label:       "",
			Description: "",
		}
	case RolesSecretary:
		return &RolesCode{
			Name:        "Secretary",
			Label:       "",
			Description: "",
		}
	case RolesSettlor:
		return &RolesCode{
			Name:        "Settlor",
			Label:       "",
			Description: "",
		}
	case RolesSignificantMember:
		return &RolesCode{
			Name:        "Significant Member",
			Label:       "",
			Description: "Major Shareholder",
		}
	case RolesSmartContractOperator:
		return &RolesCode{
			Name:        "Smart Contract Operator",
			Label:       "",
			Description: "",
		}
	case RolesTrader:
		return &RolesCode{
			Name:        "Trader",
			Label:       "",
			Description: "",
		}
	case RolesTrustee:
		return &RolesCode{
			Name:        "Trustee",
			Label:       "",
			Description: "",
		}
	case RolesUnitHolder:
		return &RolesCode{
			Name:        "Unit Holder",
			Label:       "",
			Description: "",
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

	// MsgMalformed - The OP_RETURN message is malformed. It doesn&#39;t pass data validation or something about it isn&#39;t according to protocol.
	RejectionsMsgMalformed = 1

	// TxMalformed - The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.
	RejectionsTxMalformed = 2

	// Timeout - A dependency, other contract/service, has failed to complete before the smart contract&#39;s timeout.
	RejectionsTimeout = 3

	// ContractMoved - The contract has been moved to a different address. Please find the addres change message and send requests to new address.
	RejectionsContractMoved = 4

	// DoubleSpend - A double spend attempt has been detected.
	RejectionsDoubleSpend = 5

	// ContractExists - The contract already exists and can&#39;t be recreated.
	RejectionsContractExists = 10

	// ContractAssetQtyReduction - Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.
	RejectionsContractAssetQtyReduction = 12

	// ContractFixedQuantity - Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.
	RejectionsContractFixedQuantity = 13

	// ContractAuthFlags - The contract auth flags don&#39;t permit the action requested.
	RejectionsContractAuthFlags = 14

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

	// AssetCodeExists - The asset code specified already exists and can&#39;t be reused.
	RejectionsAssetCodeExists = 20

	// AssetNotFound - The asset code is not found.
	RejectionsAssetNotFound = 21

	// AssetAuthFlags - The asset auth flags don&#39;t permit the action requested.
	RejectionsAssetAuthFlags = 22

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

	// HoldingsFrozen - Holdings are frozen, so the request can&#39;t be completed.
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

	// VoteSystemNotPermitted - The voting system isn&#39;t permitted for this request.
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
}

// RejectionsData holds a mapping of Rejections codes.
func RejectionsData(code uint32) *RejectionsCode {
	switch code {
	case RejectionsSuccess:
		return &RejectionsCode{
			Name:        "Success",
			Label:       "Success",
			Description: "No failure. This code should not be used in a reject message.",
		}
	case RejectionsMsgMalformed:
		return &RejectionsCode{
			Name:        "MsgMalformed",
			Label:       "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn&#39;t pass data validation or something about it isn&#39;t according to protocol.",
		}
	case RejectionsTxMalformed:
		return &RejectionsCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
		}
	case RejectionsTimeout:
		return &RejectionsCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract&#39;s timeout.",
		}
	case RejectionsContractMoved:
		return &RejectionsCode{
			Name:        "ContractMoved",
			Label:       "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
		}
	case RejectionsDoubleSpend:
		return &RejectionsCode{
			Name:        "DoubleSpend",
			Label:       "Double Spend",
			Description: "A double spend attempt has been detected.",
		}
	case RejectionsContractExists:
		return &RejectionsCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can&#39;t be recreated.",
		}
	case RejectionsContractAssetQtyReduction:
		return &RejectionsCode{
			Name:        "ContractAssetQtyReduction",
			Label:       "Contract Asset Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
		}
	case RejectionsContractFixedQuantity:
		return &RejectionsCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
		}
	case RejectionsContractAuthFlags:
		return &RejectionsCode{
			Name:        "ContractAuthFlags",
			Label:       "Contract Auth Flags Prohibit",
			Description: "The contract auth flags don&#39;t permit the action requested.",
		}
	case RejectionsContractExpired:
		return &RejectionsCode{
			Name:        "ContractExpired",
			Label:       "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
		}
	case RejectionsContractFrozen:
		return &RejectionsCode{
			Name:        "ContractFrozen",
			Label:       "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
		}
	case RejectionsContractRevision:
		return &RejectionsCode{
			Name:        "ContractRevision",
			Label:       "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
		}
	case RejectionsContractNotPermitted:
		return &RejectionsCode{
			Name:        "ContractNotPermitted",
			Label:       "Contract Not Permitted",
			Description: "Action not permitted by contract.",
		}
	case RejectionsContractBothOperatorsRequired:
		return &RejectionsCode{
			Name:        "ContractBothOperatorsRequired",
			Label:       "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
		}
	case RejectionsAssetCodeExists:
		return &RejectionsCode{
			Name:        "AssetCodeExists",
			Label:       "Asset Code Already Exists",
			Description: "The asset code specified already exists and can&#39;t be reused.",
		}
	case RejectionsAssetNotFound:
		return &RejectionsCode{
			Name:        "AssetNotFound",
			Label:       "Asset Not Found",
			Description: "The asset code is not found.",
		}
	case RejectionsAssetAuthFlags:
		return &RejectionsCode{
			Name:        "AssetAuthFlags",
			Label:       "Asset Auth Flags Prohibit",
			Description: "The asset auth flags don&#39;t permit the action requested.",
		}
	case RejectionsAssetFrozen:
		return &RejectionsCode{
			Name:        "AssetFrozen",
			Label:       "Asset Frozen",
			Description: "The asset is frozen and the request is not permitted while frozen.",
		}
	case RejectionsAssetRevision:
		return &RejectionsCode{
			Name:        "AssetRevision",
			Label:       "Asset Revision Incorrect",
			Description: "The revision in an asset amendment is incorrect.",
		}
	case RejectionsAssetNotPermitted:
		return &RejectionsCode{
			Name:        "AssetNotPermitted",
			Label:       "Asset Not Permitted",
			Description: "Action not permitted by asset.",
		}
	case RejectionsTransferSelf:
		return &RejectionsCode{
			Name:        "TransferSelf",
			Label:       "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
		}
	case RejectionsTransferExpired:
		return &RejectionsCode{
			Name:        "TransferExpired",
			Label:       "Transfer Expired",
			Description: "The transfer has expired.",
		}
	case RejectionsHoldingsFrozen:
		return &RejectionsCode{
			Name:        "HoldingsFrozen",
			Label:       "Holdings Frozen",
			Description: "Holdings are frozen, so the request can&#39;t be completed.",
		}
	case RejectionsHolderProposalProhibited:
		return &RejectionsCode{
			Name:        "HolderProposalProhibited",
			Label:       "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
		}
	case RejectionsProposalConflicts:
		return &RejectionsCode{
			Name:        "ProposalConflicts",
			Label:       "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
		}
	case RejectionsVoteNotFound:
		return &RejectionsCode{
			Name:        "VoteNotFound",
			Label:       "Vote Not Found",
			Description: "The vote ID referenced is not found.",
		}
	case RejectionsVoteClosed:
		return &RejectionsCode{
			Name:        "VoteClosed",
			Label:       "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
		}
	case RejectionsBallotAlreadyCounted:
		return &RejectionsCode{
			Name:        "BallotAlreadyCounted",
			Label:       "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
		}
	case RejectionsVoteSystemNotPermitted:
		return &RejectionsCode{
			Name:        "VoteSystemNotPermitted",
			Label:       "Vote System Not Permitted",
			Description: "The voting system isn&#39;t permitted for this request.",
		}
	case RejectionsInsufficientTxFeeFunding:
		return &RejectionsCode{
			Name:        "InsufficientTxFeeFunding",
			Label:       "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
		}
	case RejectionsInsufficientValue:
		return &RejectionsCode{
			Name:        "InsufficientValue",
			Label:       "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
		}
	case RejectionsInsufficientQuantity:
		return &RejectionsCode{
			Name:        "InsufficientQuantity",
			Label:       "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
		}
	case RejectionsNotAdministration:
		return &RejectionsCode{
			Name:        "NotAdministration",
			Label:       "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
		}
	case RejectionsNotOperator:
		return &RejectionsCode{
			Name:        "NotOperator",
			Label:       "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
		}
	case RejectionsUnauthorizedAddress:
		return &RejectionsCode{
			Name:        "UnauthorizedAddress",
			Label:       "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
		}
	case RejectionsInvalidSignature:
		return &RejectionsCode{
			Name:        "InvalidSignature",
			Label:       "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
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
			Description: "The OP_RETURN message is malformed. It doesn&#39;t pass data validation or something about it isn&#39;t according to protocol.",
		},
		RejectionsTxMalformed: &RejectionsCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
		},
		RejectionsTimeout: &RejectionsCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract&#39;s timeout.",
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
			Description: "The contract already exists and can&#39;t be recreated.",
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
		RejectionsContractAuthFlags: &RejectionsCode{
			Name:        "ContractAuthFlags",
			Label:       "Contract Auth Flags Prohibit",
			Description: "The contract auth flags don&#39;t permit the action requested.",
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
			Description: "The asset code specified already exists and can&#39;t be reused.",
		},
		RejectionsAssetNotFound: &RejectionsCode{
			Name:        "AssetNotFound",
			Label:       "Asset Not Found",
			Description: "The asset code is not found.",
		},
		RejectionsAssetAuthFlags: &RejectionsCode{
			Name:        "AssetAuthFlags",
			Label:       "Asset Auth Flags Prohibit",
			Description: "The asset auth flags don&#39;t permit the action requested.",
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
			Description: "Holdings are frozen, so the request can&#39;t be completed.",
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
			Description: "The voting system isn&#39;t permitted for this request.",
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
}

// TagsData holds a mapping of Tags codes.
func TagsData(code uint32) *TagsCode {
	switch code {
	case TagsHousing:
		return &TagsCode{
			Name:        "Housing",
			Label:       "",
			Description: "",
		}
	case TagsUtilities:
		return &TagsCode{
			Name:        "Utilities",
			Label:       "",
			Description: "",
		}
	case TagsFood:
		return &TagsCode{
			Name:        "Food",
			Label:       "",
			Description: "",
		}
	case TagsMedical:
		return &TagsCode{
			Name:        "Medical",
			Label:       "",
			Description: "",
		}
	case TagsFinancialServices:
		return &TagsCode{
			Name:        "Financial Services",
			Label:       "",
			Description: "",
		}
	case TagsEntertainment:
		return &TagsCode{
			Name:        "Entertainment",
			Label:       "",
			Description: "",
		}
	case TagsSales:
		return &TagsCode{
			Name:        "Sales",
			Label:       "",
			Description: "",
		}
	case TagsAutomotive:
		return &TagsCode{
			Name:        "Automotive",
			Label:       "",
			Description: "",
		}
	case TagsTransportation:
		return &TagsCode{
			Name:        "Transportation",
			Label:       "",
			Description: "",
		}
	case TagsFitness:
		return &TagsCode{
			Name:        "Fitness",
			Label:       "",
			Description: "",
		}
	case TagsElectricity:
		return &TagsCode{
			Name:        "Electricity",
			Label:       "",
			Description: "",
		}
	case TagsWater:
		return &TagsCode{
			Name:        "Water",
			Label:       "",
			Description: "",
		}
	case TagsInternet:
		return &TagsCode{
			Name:        "Internet",
			Label:       "",
			Description: "",
		}
	case TagsMedicine:
		return &TagsCode{
			Name:        "Medicine",
			Label:       "",
			Description: "",
		}
	case TagsService:
		return &TagsCode{
			Name:        "Service",
			Label:       "",
			Description: "",
		}
	case TagsRepair:
		return &TagsCode{
			Name:        "Repair",
			Label:       "",
			Description: "",
		}
	case TagsSupplies:
		return &TagsCode{
			Name:        "Supplies",
			Label:       "",
			Description: "",
		}
	case TagsParts:
		return &TagsCode{
			Name:        "Parts",
			Label:       "",
			Description: "",
		}
	case TagsLabor:
		return &TagsCode{
			Name:        "Labor",
			Label:       "",
			Description: "",
		}
	case TagsTip:
		return &TagsCode{
			Name:        "Tip",
			Label:       "",
			Description: "",
		}
	case TagsMedia:
		return &TagsCode{
			Name:        "Media",
			Label:       "",
			Description: "",
		}
	case TagsMusic:
		return &TagsCode{
			Name:        "Music",
			Label:       "",
			Description: "",
		}
	case TagsVideo:
		return &TagsCode{
			Name:        "Video",
			Label:       "",
			Description: "",
		}
	case TagsPhoto:
		return &TagsCode{
			Name:        "Photo",
			Label:       "",
			Description: "",
		}
	case TagsAudio:
		return &TagsCode{
			Name:        "Audio",
			Label:       "",
			Description: "",
		}
	case TagsAlcohol:
		return &TagsCode{
			Name:        "Alcohol",
			Label:       "",
			Description: "",
		}
	case TagsTobacco:
		return &TagsCode{
			Name:        "Tobacco",
			Label:       "",
			Description: "",
		}
	case TagsDiscounted:
		return &TagsCode{
			Name:        "Discounted",
			Label:       "",
			Description: "",
		}
	case TagsPromotional:
		return &TagsCode{
			Name:        "Promotional",
			Label:       "",
			Description: "",
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
