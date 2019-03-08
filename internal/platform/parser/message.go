package parser

import (
	"fmt"
	"html"
	"strings"
)

type Actions struct {
	Package string
	Actions []Action
}

func (m Actions) ProtocolActions() []Action {
	pm := []Action{}

	for _, v := range m.Actions {
		if len(v.Code) == 0 {
			continue
		}

		pm = append(pm, v)
	}

	return pm
}

func (m Actions) TypeActions() []Action {
	pm := []Action{}

	for _, v := range m.Actions {
		if len(v.Code) > 0 {
			continue
		}

		pm = append(pm, v)
	}

	return pm
}

type Action struct {
	Code     string
	Metadata Metadata
	Rules    Rules
	Fields   []Field
}

func (m Action) CodeNameComment() string {
	s := fmt.Sprintf("%v identifies data as a %v message.",
		m.CodeName(),
		m.Name())

	return reformat(s, "\t//")
}

func (m Action) CodeName() string {
	return fmt.Sprintf("Code%v", m.Name())
}

func (m Action) TypeLetter() string {
	code := strings.ToLower(m.Code[:1])
	return fmt.Sprintf(code)
}

func (m Action) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}

func (m Action) Label() string {
	return fmt.Sprintf(m.Metadata.Label)
}

func (m Action) Description() string {
	return fmt.Sprintf(m.Metadata.Description)
}

func (m Action) Hex() string {
	vals := []string{}

	for _, f := range m.PayloadFields() {
		s := strings.Replace(f.ExampleHex, "0x", "", 1)
		vals = append(vals, s)
	}

	return strings.Join(vals, "")
}

func (m Action) PayloadFields() []Field {
	fields := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" {
			continue

		}

		fields = append(fields, f)
	}

	return fields
}

func (m Action) HasPayloadAction() bool {
	switch m.Name() {
	case "AssetDefinition", "AssetCreation":
		return true
	}

	return false
}

func (m Action) CommentSlash() string {
	s := html.UnescapeString(m.Metadata.Description)
	s = m.Metadata.Name + " : " + s

	return reformat(s, "//")
}

func (m Action) CommentHash() string {
	s := html.UnescapeString(m.Metadata.Description)
	s = m.Metadata.Name + " : " + s

	return reformat(s, "#")
}

func (m Action) DataFields() []Field {
	d := []Field{}

	for _, f := range m.Fields {
		if f.Name == "Header" || f.Name == "ProtocolID" || f.Name == "ActionPrefix" {
			continue
		}

		d = append(d, f)
	}

	return d
}
