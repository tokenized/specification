package protocol

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type RejectionCodeData struct {
	Code        uint8
	Name        string
	Label       string
	Description string
	// Metadata string
}

var rejectionTypes map[uint8]RejectionCodeData
var rejectionLock sync.Mutex

func getRejectionCodes() (map[uint8]RejectionCodeData, error) {
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

	types, err := getRejectionCodes()
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
	Code        string
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var currencyTypes map[[3]byte]CurrencyTypeData
var currencyLock sync.Mutex

func getCurrencies() (map[[3]byte]CurrencyTypeData, error) {
	if currencyTypes != nil {
		return currencyTypes, nil
	}

	load := make([]CurrencyTypeData, 0)

	if err := yaml.Unmarshal([]byte(yamlCurrencies), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal currencyTypes yaml")
	}

	currencyTypes = make(map[[3]byte]CurrencyTypeData)
	for _, value := range load {
		if len(value.Code) != 3 {
			return nil, fmt.Errorf("Currency type incorrect length : %s", value.Code)
		}
		var code [3]byte
		copy(code[:], []byte(value.Code))
		currencyTypes[code] = value
	}
	return currencyTypes, nil
}

func GetCurrency(cur [3]byte) *CurrencyTypeData {
	currencyLock.Lock()
	defer currencyLock.Unlock()

	types, err := getCurrencies()
	if err != nil {
		return nil
	}
	result, exists := types[cur]
	if !exists {
		return nil
	}
	return &result
}

type EntityTypeData struct {
	Code        string
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var entityTypes map[byte]EntityTypeData
var entityLock sync.Mutex

func getEntityTypes() (map[byte]EntityTypeData, error) {
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

	types, err := getEntityTypes()
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
	Code        string
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var polityTypes map[string]PolityType
var polityLock sync.Mutex

func getPolityTypes() (map[string]PolityType, error) {
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

	types, err := getPolityTypes()
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
	Code        uint8
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var roleTypes map[uint8]RoleType
var roleLock sync.Mutex

func getRoleTypes() (map[uint8]RoleType, error) {
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

	types, err := getRoleTypes()
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
	Code  uint8
	Label string
}

var tagTypes map[uint8]TagType
var tagLock sync.Mutex

func getTagTypes() (map[uint8]TagType, error) {
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

	types, err := getTagTypes()
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
