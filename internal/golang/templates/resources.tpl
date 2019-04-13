package protocol

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type RejectionCode struct {
	Name string
	Text string
	Description string
}

var rejectionTypes map[uint8]RejectionCode

func GetRejectionCodes() (map[uint8]RejectionCode, error) {
	if rejectionTypes != nil {
		return rejectionTypes, nil
	}

	load := struct {
		Values map[uint8]RejectionCode
	}{}

	if err := yaml.Unmarshal([]byte(yamlRejectionCodes), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal rejection codes yaml")
	}

	rejectionTypes = load.Values
	return rejectionTypes, nil
}

func GetRejectionCode(code uint8) *RejectionCode {
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

type CurrencyType struct {
	Name              string
	Label             string
	Symbol            string
	Fractions         uint64
	FractionalUnit    string `yaml:"fractional_unit"`
	Polity            string
	MonetaryAuthority string `yaml:"monetary_authority"`
}

var currencyTypes map[string]CurrencyType

func GetCurrencies() (map[string]CurrencyType, error) {
	if currencyTypes != nil {
		return currencyTypes, nil
	}

	load := struct {
		Values map[string]CurrencyType
	}{}

	if err := yaml.Unmarshal([]byte(yamlCurrencies), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal currencyTypes yaml")
	}

	currencyTypes = load.Values
	return currencyTypes, nil
}

func GetCurrency(cur string) *CurrencyType {
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

type EntityType struct {
	Name              string
	Label             string
	Symbol            string
}

var entityTypes map[string]EntityType

func GetEntityTypes() (map[string]EntityType, error) {
	if entityTypes != nil {
		return entityTypes, nil
	}

	load := struct {
		Values map[string]EntityType
	}{}

	if err := yaml.Unmarshal([]byte(yamlEntities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal entities yaml")
	}

	entityTypes = load.Values
	return entityTypes, nil
}

func GetEntityType(ent string) *EntityType {
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

	load := struct {
		Values map[string]PolityType
	}{}

	if err := yaml.Unmarshal([]byte(yamlPolities), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal polities yaml")
	}

	polityTypes = load.Values
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

	load := struct {
		Values map[uint8]RoleType
	}{}

	if err := yaml.Unmarshal([]byte(yamlRoles), &load); err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal roles yaml")
	}

	roleTypes = load.Values
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