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
	Code        string
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var currencyTypes map[[3]byte]CurrencyTypeData

func GetCurrencies() (map[[3]byte]CurrencyTypeData, error) {
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
	types, err := GetCurrencies()
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
	Code        string
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var polityTypes map[string]PolityType

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
	Code        uint8
	Name        string
	Label       string
	Description string
	// Metadata    string
}

var roleTypes map[uint8]RoleType

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

{{range .}}
{{comment (print .Name " - " .Metadata.Description) "//"}}
var yaml{{ .Name }} = `
{{ .Data }}
`
{{ end }}