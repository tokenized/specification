package parser

import (
	"fmt"
	"html"
	"strings"
)

type Messages struct {
	Package  string
	Messages []Message
}

func (m Messages) ProtocolMessages() []Message {
	pm := []Message{}

	for _, v := range m.Messages {
		if len(v.Code) == 0 {
			continue
		}

		pm = append(pm, v)
	}

	return pm
}

func (m Messages) TypeMessages() []Message {
	pm := []Message{}

	for _, v := range m.Messages {
		if len(v.Code) > 0 {
			continue
		}

		pm = append(pm, v)
	}

	return pm
}

type Message struct {
	Code     string
	Metadata Metadata
	Rules    Rules
	Fields   []Field
}

func (m Message) CodeNameComment() string {
	s := fmt.Sprintf("%v identifies data as a %v message.",
		m.CodeName(),
		m.Name())

	return reformat(s, "\t//")
}

func (m Message) CodeName() string {
	return fmt.Sprintf("Code%v", m.Name())
}

func (m Message) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}

func (m Message) Hex() string {
	vals := []string{}

	for _, f := range m.PayloadFields() {
		s := strings.Replace(f.ExampleHex, "0x", "", 1)
		vals = append(vals, s)
	}

	return strings.Join(vals, "")
}

func (m Message) PayloadFields() []Field {
	fields := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" {
			continue

		}

		fields = append(fields, f)
	}

	return fields
}

func (m Message) HasPayloadMessage() bool {
	switch m.Name() {
	case "AssetDefinition", "AssetCreation":
		return true
	}

	return false
}

func (m Message) GoComment() string {
	s := html.UnescapeString(m.Metadata.Description)
	s = m.Metadata.Name + " : " + s

	return reformat(s, "//")
}

func (m Message) DataFields() []Field {
	d := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" || f.Name == "ProtocolID" || f.Name == "ActionPrefix" {
			continue
		}

		d = append(d, f)
	}

	return d
}
