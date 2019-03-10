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
	Fields   []Field
}

func (m Asset) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}
