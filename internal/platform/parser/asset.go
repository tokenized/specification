package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func NewAssets(filenames []string) Assets {
	items := Assets{}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := Asset{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}

		items = append(items, m)
	}

	return items
}

type Assets []Asset

type Asset struct {
	Code     string
	Metadata Metadata
	Fields   []AssetField
}

func (m Asset) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}

type AssetField struct {
	Type string
	Name string
	Size uint
}

func (m AssetField) numeric() bool {
	if m.Type == "float" || m.Type == "int" || m.Type == "uint" {
		return true
	}

	return false
}

func (m AssetField) GoType() string {
	if !m.numeric() && m.Size == 1 {
		return "byte"
	}

	switch m.Type {
	case "country":
		return "[]byte"
	case "dropdown":
		return "[]byte"
	case "string":
		return "[]byte"
	case "time":
		return "uint64"
	}

	if m.Type == "uint" || m.Type == "float" {
		i := m.Size * 8
		return fmt.Sprintf("%v%v", m.Type, i)
	}

	return m.Type
}

// Bytes returns the code to write the field into a bytes.Buffer.
func (m AssetField) Bytes() string {
	if m.GoType() == "[]byte" {
		return fmt.Sprintf(`
	if err := m.write(buf, m.pad(m.%v, %v)); err != nil {
		return nil, err
	}`, m.Name, m.Size)
	}

	return fmt.Sprintf(`
	if err := m.write(buf, m.%v); err != nil {
		return nil, err
	}`, m.Name)
}

func (m AssetField) Write() string {
	if m.GoType() == "[]byte" {
		return fmt.Sprintf(`
	m.%v = make([]byte, %v)
	if err := m.readLen(buf, m.%v); err != nil {
		return 0, err
	}

	m.%v = bytes.Trim(m.%v, "\x00")`, m.Name, m.Size, m.Name, m.Name, m.Name)
	}

	return fmt.Sprintf(`
	if err := m.read(buf, &m.%v); err != nil {
		return 0, err
	}`, m.Name)
}
