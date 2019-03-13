package parser

import (
	"strings"
)

type Constructor struct {
	Name      string
	Type      string
	Size      uint64
	Field     string
	SetMethod string `yaml:"set_method"`
}

func (c Constructor) ConstructorName() string {
	return strings.Replace(c.Name, " ", "", -1)
}

func (c Constructor) ConstructorField() string {
	return strings.Replace(c.Field, " ", "", -1)
}

func (c Constructor) ConstructorSetMethod() string {
	result := strings.Replace(c.SetMethod, " ", "", -1)
	if len(result) == 0 {
		return "="
	}
	return result
}

func (c Constructor) ConstructorGoType() string {
	return GoType(c.Type, c.Size)
}
