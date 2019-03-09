package parser

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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

	return items
}

type ProtocolTypes []ProtocolType

type ProtocolType struct {
	Metadata Metadata
	Fields   []Field
}
