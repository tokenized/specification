package protocol

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type RejectionCodeData struct {
	Name string
	Text string
	Description string
}

var rejectionTypes map[uint8]RejectionCodeData

func GetRejectionCodes() (map[uint8]RejectionCodeData, error) {
	if rejectionTypes != nil {
		return rejectionTypes, nil
	}

	var load map[uint8]RejectionCodeData

	if err := yaml.Unmarshal([]byte(yamlRejections), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal rejection codes yaml")
	}

	rejectionTypes = load
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
	Name              string
	Label             string
	Symbol            string
	Fractions         uint64
	FractionalUnit    string `yaml:"fractional_unit"`
	Polity            string
	MonetaryAuthority string `yaml:"monetary_authority"`
}

var currencyTypes map[string]CurrencyTypeData

func GetCurrencies() (map[string]CurrencyTypeData, error) {
	if currencyTypes != nil {
		return currencyTypes, nil
	}

	var load map[string]CurrencyTypeData

	if err := yaml.Unmarshal([]byte(yamlCurrencies), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal currencyTypes yaml")
	}

	currencyTypes = load
	return currencyTypes, nil
}

func GetCurrency(cur string) *CurrencyTypeData {
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
	Name              string
	Label             string
	Symbol            string
}

var entityTypes map[byte]EntityTypeData

func GetEntityTypes() (map[byte]EntityTypeData, error) {
	if entityTypes != nil {
		return entityTypes, nil
	}

	var load map[string]EntityTypeData

	if err := yaml.Unmarshal([]byte(yamlEntities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal entities yaml")
	}

	entityTypes = make(map[byte]EntityTypeData)
	for key, value := range load {
		if len(key) > 1 {
			return nil, fmt.Errorf("Entity type too long : %s", key)
		}
		entityTypes[byte(key[0])] = value
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
	Name          string
	States        map[string]string
	Flag          string
	FiscalYear    string `yaml:"fractional_unit"`
	GovFiscalYear string `yaml:"gov_fiscal_year"`
}

var polityTypes map[string]PolityType

func GetPolityTypes() (map[string]PolityType, error) {
	if polityTypes != nil {
		return polityTypes, nil
	}

	var load map[string]PolityType

	if err := yaml.Unmarshal([]byte(yamlPolities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal polities yaml")
	}

	polityTypes = load
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
	Name string
}

var roleTypes map[uint8]RoleType

func GetRoleTypes() (map[uint8]RoleType, error) {
	if roleTypes != nil {
		return roleTypes, nil
	}

	var load map[uint8]RoleType

	if err := yaml.Unmarshal([]byte(yamlRoles), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal roles yaml")
	}

	roleTypes = load
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