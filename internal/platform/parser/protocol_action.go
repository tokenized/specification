package parser

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolActions(filenames []string) ProtocolActions {
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

		items = append(items, m)
	}

	// Order by action code
	sort.Slice(items, func(i, j int) bool {
		return items[i].Code < items[j].Code
	})

	return items
}

type ProtocolActions []ProtocolAction

type ProtocolAction struct {
	Code     string
	Metadata Metadata
	Rules    Rules
	Fields   []Field
	SubTypes []ProtocolType
}

func (m ProtocolAction) CodeNameComment() string {
	s := fmt.Sprintf("%v identifies data as a %v message.",
		m.CodeName(),
		m.Name())

	return reformat(s, "\t//")
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

func (m ProtocolAction) HasPayloadMessage() bool {
	switch m.Name() {
	case "AssetDefinition", "AssetCreation":
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

// AssociateActionsAndTypes sets the SubTypes of each ProtocolAction to be a
// slice of ProtocolType that the ProtocolAction references. This allows
// a ProtocolAction to reference these types where needed.
func AssociateActionsAndTypes(actions []ProtocolAction,
	fieldTypes []ProtocolType) []ProtocolAction {

	// put the protocol types in a map for easy lookup
	protocolTypes := map[string]ProtocolType{}
	for _, t := range fieldTypes {
		protocolTypes[t.Name()] = t
	}

	for i, a := range actions {
		subTypes := []ProtocolType{}

		for j, f := range a.Fields {
			k := f.GoTypeSingular()

			t, ok := protocolTypes[k]
			if !ok {
				continue
			}

			// the action knows about this type, so add it to the subTypes.
			subTypes = append(subTypes, t)

			// this field is an internal type
			actions[i].Fields[j].internalType = true
		}

		// set the action subtypes
		actions[i].SubTypes = subTypes
	}

	return actions
}
