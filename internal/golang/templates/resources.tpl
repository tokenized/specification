package protocol

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type RejectionCodeData struct {
	Code        uint8  `json:"code"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// Metadata string
}

var rejectionTypes map[uint8]RejectionCodeData
var rejectionLock sync.Mutex

func GetRejectionCodes() (map[uint8]RejectionCodeData, error) {
	if rejectionTypes != nil {
		return rejectionTypes, nil
	}

	load := make([]RejectionCodeData, 0)

	if err := yaml.Unmarshal([]byte(yamlRejections), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal rejection codes yaml")
	}

	rejectionTypes = make(map[uint8]RejectionCodeData)
	for _, value := range load {
		rejectionTypes[value.Code] = value
	}
	return rejectionTypes, nil
}

func GetRejectionCode(code uint8) *RejectionCodeData {
	rejectionLock.Lock()
	defer rejectionLock.Unlock()

	types, err := GetRejectionCodes()
	if err != nil {
		return nil
	}
	result, exists := types[code]
	if !exists {
		return nil
	}
	return &result
}

type CurrencyTypeData struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Metadata    struct {
		Symbol          string `json:"symbol"`
		Precision       int    `json:"precision"`
		Fractionals     int    `json:"fractionals"`
		Fractional_Unit string `json:"fractional_unit"`
	} `json:"metadata"`
}

var currencyTypes map[string]CurrencyTypeData
var currencyLock sync.Mutex

func GetCurrencies() (map[string]CurrencyTypeData, error) {
	if currencyTypes != nil {
		return currencyTypes, nil
	}

	load := make([]CurrencyTypeData, 0)

	if err := yaml.Unmarshal([]byte(yamlCurrencies), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal currencyTypes yaml")
	}

	currencyTypes = make(map[string]CurrencyTypeData)
	for _, value := range load {
		if len(value.Code) != 3 {
			return nil, fmt.Errorf("Currency type incorrect length : %s", value.Code)
		}
		currencyTypes[value.Code] = value
	}
	return currencyTypes, nil
}

func GetCurrency(cur [3]byte) *CurrencyTypeData {
	currencyLock.Lock()
	defer currencyLock.Unlock()

	types, err := GetCurrencies()
	if err != nil {
		return nil
	}

	curs := string(cur[:])
	result, exists := types[curs]
	if !exists {
		return nil
	}
	return &result
}

type EntityTypeData struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// Metadata    string
}

var entityTypes map[byte]EntityTypeData
var entityLock sync.Mutex

func GetEntityTypes() (map[byte]EntityTypeData, error) {
	if entityTypes != nil {
		return entityTypes, nil
	}

	load := make([]EntityTypeData, 0)

	if err := yaml.Unmarshal([]byte(yamlEntities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal entities yaml")
	}

	entityTypes = make(map[byte]EntityTypeData)
	for _, value := range load {
		if len(value.Code) > 1 {
			return nil, fmt.Errorf("Entity type too long : %s", value.Code)
		}
		entityTypes[byte(value.Code[0])] = value
	}
	return entityTypes, nil
}

func GetEntityType(ent byte) *EntityTypeData {
	entityLock.Lock()
	defer entityLock.Unlock()

	types, err := GetEntityTypes()
	if err != nil {
		return nil
	}
	result, exists := types[ent]
	if !exists {
		return nil
	}
	return &result
}

type PolityType struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// Metadata    string
}

var polityTypes map[string]PolityType
var polityLock sync.Mutex

func GetPolityTypes() (map[string]PolityType, error) {
	if polityTypes != nil {
		return polityTypes, nil
	}

	load := make([]PolityType, 0)

	if err := yaml.Unmarshal([]byte(yamlPolities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal polities yaml")
	}

	polityTypes = make(map[string]PolityType)
	for _, value := range load {
		polityTypes[value.Code] = value
	}
	return polityTypes, nil
}

func GetPolityType(pol string) *PolityType {
	polityLock.Lock()
	defer polityLock.Unlock()

	types, err := GetPolityTypes()
	if err != nil {
		return nil
	}
	result, exists := types[pol]
	if !exists {
		return nil
	}
	return &result
}

type RoleType struct {
	Code        uint8  `json:"code"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// Metadata    string
}

var roleTypes map[uint8]RoleType
var roleLock sync.Mutex

func GetRoleTypes() (map[uint8]RoleType, error) {
	if roleTypes != nil {
		return roleTypes, nil
	}

	load := make([]RoleType, 0)

	if err := yaml.Unmarshal([]byte(yamlRoles), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal roles yaml")
	}

	roleTypes = make(map[uint8]RoleType)
	for _, value := range load {
		roleTypes[value.Code] = value
	}
	return roleTypes, nil
}

func GetRoleType(role uint8) *RoleType {
	roleLock.Lock()
	defer roleLock.Unlock()

	types, err := GetRoleTypes()
	if err != nil {
		return nil
	}
	result, exists := types[role]
	if !exists {
		return nil
	}
	return &result
}

type TagType struct {
	Code  uint8  `json:"code"`
	Label string `json:"name"`
}

var tagTypes map[uint8]TagType
var tagLock sync.Mutex

func GetTagTypes() (map[uint8]TagType, error) {
	if tagTypes != nil {
		return tagTypes, nil
	}

	load := make([]TagType, 0)

	if err := yaml.Unmarshal([]byte(yamlTags), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal tags yaml")
	}

	tagTypes = make(map[uint8]TagType)
	for _, value := range load {
		tagTypes[value.Code] = value
	}
	return tagTypes, nil
}

func GetTagType(tag uint8) *TagType {
	tagLock.Lock()
	defer tagLock.Unlock()

	types, err := GetTagTypes()
	if err != nil {
		return nil
	}
	result, exists := types[tag]
	if !exists {
		return nil
	}
	return &result
}

{{range .}}
{{comment (print .Name " - " .Metadata.Description) "//"}}
var yaml{{ .Name }} = `
{{ .Data }}
`
{{ end }}
