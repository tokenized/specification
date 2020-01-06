package messages

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
