package parser

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolRejectionCodes(filename string) ProtocolRejectionCodes {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	m := ProtocolRejectionCodes{}
	if err := yaml.Unmarshal(data, &m); err != nil {
		panic(fmt.Errorf("file %v : %v", filename, err))
	}
	return m
}

type ProtocolRejectionCode struct {
	Name        string
	Text        string
	Description string
}

type ProtocolRejectionCodes struct {
	Metadata Metadata
	Values   map[uint8]ProtocolRejectionCode
}
