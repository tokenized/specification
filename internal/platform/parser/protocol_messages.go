package parser

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolMessages(fieldTypes []ProtocolType, filenames []string) ProtocolMessages {
	items := ProtocolMessages{}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := ProtocolMessage{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}

		// This is not one of the action definitions
		if len(m.Metadata.Name) == 0 {
			continue
		}

		items = append(items, m)
	}

	// Order by message code
	sort.Slice(items, func(i, j int) bool {
		intI, err := strconv.Atoi(items[i].Code)
		if err != nil {
			return true
		}
		intJ, err := strconv.Atoi(items[j].Code)
		if err != nil {
			return false
		}
		return intI < intJ
	})

	return items
}

type ProtocolMessages []ProtocolMessage

type ProtocolMessage struct {
	Code        string
	Metadata    Metadata
	Rules       Rules
	Fields      []Field
	FieldTypes  []ProtocolType
	Constructor []Constructor
	Functions   []Function
}

func (m ProtocolMessage) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}

func (m ProtocolMessage) URLCode() string {
	return "message-" + KebabCase(m.Name())
}

func (m ProtocolMessage) CodeNameComment() string {
	s := fmt.Sprintf("%v identifies data as a %v message.",
		m.CodeName(),
		m.Name())

	return reformat(s, "\t//")
}

func (m ProtocolMessage) CodeComment() string {
	s := fmt.Sprintf("%s identifies data as a %v message.",
		m.ActionCode(),
		m.Name())

	return reformat(s, "\t//")
}

func (m ProtocolMessage) ActionCode() string {
	return m.Code
}

func (m ProtocolMessage) CodeName() string {
	return fmt.Sprintf("Code%v", m.Name())
}

func (m ProtocolMessage) TypeLetter() string {
	return strings.ToLower(m.Code[:1])
}

func (m ProtocolMessage) Label() string {
	return fmt.Sprintf(m.Metadata.Label)
}

func (m ProtocolMessage) Description() string {
	return fmt.Sprintf(m.Metadata.Description)
}

func (m ProtocolMessage) PayloadFields() []Field {
	fields := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" {
			continue
		}

		fields = append(fields, f)
	}

	return fields
}

func (m ProtocolMessage) HasPayloadMessage() bool {
	switch m.Name() {
	case "AssetDefinition", "AssetCreation":
		return true
	}

	return false
}

func (m ProtocolMessage) DataFields() []Field {
	d := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" || f.Name == "ProtocolID" || f.Name == "ActionPrefix" {
			continue
		}

		d = append(d, f)
	}

	return d
}

func (m ProtocolMessage) FieldCount() string {
	return fmt.Sprintf("%v", len(m.Fields))
}
