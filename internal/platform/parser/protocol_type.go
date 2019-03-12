package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolTypes(filenames []string) ProtocolTypes {
	items := ProtocolTypes{}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := ProtocolType{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}

		items = append(items, m)
	}

	types := map[string]ProtocolType{}

	for _, v := range items {
		types[v.Metadata.Name] = v
	}

	// the protocol types may have protocol types as fields.
	for i, t := range items {
		if t.Metadata.Name != "Entity" {
			continue
		}

		for j, f := range t.Fields {
			if _, ok := types[f.GoTypeSingular()]; !ok {
				continue
			}

			// this is an internal type
			items[i].Fields[j].internalType = true
		}
	}

	return items
}

type ProtocolTypes []ProtocolType

type ProtocolType struct {
	Metadata Metadata
	Fields   []Field
	SubTypes []ProtocolType
}

func (m ProtocolType) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}
