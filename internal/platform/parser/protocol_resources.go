package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolResources(filenames []string) ProtocolResources {
	items := ProtocolResources{}

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := ProtocolResource{}
		if err := yaml.Unmarshal(data, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}
		m.Data = string(data)

		items = append(items, m)
	}

	return items
}

type ProtocolResources []ProtocolResource

type ProtocolResource struct {
	Metadata Metadata
	Data string
}

func (m ProtocolResource) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}
