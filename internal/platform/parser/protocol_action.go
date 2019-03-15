package parser

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolActions(fieldTypes []ProtocolType, filenames []string) ProtocolActions {
	items := ProtocolActions{}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := ProtocolAction{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}

		// This is not one of the action definitions
		if len(m.Metadata.Name) == 0 {
			continue
		}

		// Each action knowns about 0 or more protocol types
		m = setFieldTypes(m, fieldTypes)

		items = append(items, m)
	}

	// Order by action code
	sort.Slice(items, func(i, j int) bool {
		return items[i].Code < items[j].Code
	})

	return items
}

// setFieldTypes populates the available field types on the action
func setFieldTypes(action ProtocolAction, fieldTypes []ProtocolType) ProtocolAction {

	// Out the protocol types in a map for easy lookup
	protocolTypes := map[string]ProtocolType{}
	for _, t := range fieldTypes {
		protocolTypes[t.Name()] = t
	}

	// Set the action subtypes
	pts := []ProtocolType{}

	for _, f := range action.Fields {
		k := f.SingularType()

		t, ok := protocolTypes[k]
		if !ok {
			continue
		}

		// The action knows about this type, so add it to the field types
		pts = append(pts, t)

		// This field is an internal type
		// action.Fields[j].internalType = true
	}

	action.FieldTypes = pts

	return action
}

type ProtocolActions []ProtocolAction

type ProtocolAction struct {
	Code        string
	Metadata    Metadata
	Rules       Rules
	Fields      []Field
	FieldTypes  []ProtocolType
	Constructor []Constructor
	Functions   []Function
}

func (m ProtocolAction) CodeNameComment() string {
	s := fmt.Sprintf("%v identifies data as a %v message.",
		m.CodeName(),
		m.Name())

	return reformat(s, "\t//")
}

func (m ProtocolAction) CodeComment() string {
	s := fmt.Sprintf("%s identifies data as a %v message.",
		m.ActionCode(),
		m.Name())

	return reformat(s, "\t//")
}

func (m ProtocolAction) ActionCode() string {
	return m.Code
}

func (m ProtocolAction) CodeName() string {
	return fmt.Sprintf("Code%v", m.Name())
}

func (m ProtocolAction) TypeLetter() string {
	code := strings.ToLower(m.Code[:1])
	return fmt.Sprintf(code)
}

func (m ProtocolAction) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}

func (m ProtocolAction) Label() string {
	return fmt.Sprintf(m.Metadata.Label)
}

func (m ProtocolAction) Description() string {
	return fmt.Sprintf(m.Metadata.Description)
}

func (m ProtocolAction) Hex() string {
	vals := []string{}

	for _, f := range m.PayloadFields() {
		s := strings.Replace(f.ExampleHex, "0x", "", 1)
		vals = append(vals, s)
	}

	return strings.Join(vals, "")
}

func (m ProtocolAction) PayloadFields() []Field {
	fields := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" {
			continue
		}

		fields = append(fields, f)
	}

	return fields
}

func (m ProtocolAction) HasAssetPayload() bool {
	switch m.Name() {
	case "AssetDefinition", "AssetCreation":
		return true
	}

	return false
}

func (m ProtocolAction) HasMessagePayload() bool {
	switch m.Name() {
	case "Message":
		return true
	}

	return false
}

func (m ProtocolAction) DataFields() []Field {
	d := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" || f.Name == "ProtocolID" || f.Name == "ActionPrefix" {
			continue
		}

		d = append(d, f)
	}

	return d
}

func (m ProtocolAction) FieldCount() string {
	return fmt.Sprintf("%v", len(m.Fields))
}
