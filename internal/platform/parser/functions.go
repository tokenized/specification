package parser

import (
	"strings"
)

type Param struct {
	Name      string
	Type      string
	Size      uint64
	Value     string
	Field     string
	SetMethod string `yaml:"set_method"`
}

type Function struct {
	Name        string
	Description string
	Type        string
	AppendType  string `yaml:"append_type"`
	Params      []Param
}

func (f Function) FunctionName() string {
	return strings.Replace(f.Name, " ", "", -1)
}

func (f Function) FunctionType() string {
	return strings.Replace(f.Type, " ", "", -1)
}

func (f Function) FunctionAppendType() string {
	return strings.Replace(f.AppendType, " ", "", -1)
}

func (f Function) FunctionDescription() string {
	return strings.Trim(f.Description, " ")
}

func (f Function) FunctionParams() []Param {
	return f.Params
}

func (p Param) ParamName() string {
	return strings.Replace(p.Name, " ", "", -1)
}

func (p Param) ParamValue() string {
	if len(p.Value) == 0 {
		return strings.Replace(p.Name, " ", "", -1)
	}
	return p.Value
}

func (p Param) ParamField() string {
	return strings.Replace(p.Field, " ", "", -1)
}

func (p Param) ParamSetMethod() string {
	result := strings.Replace(p.SetMethod, " ", "", -1)
	if len(result) == 0 {
		return "="
	}
	return result
}

func (p Param) ParamGoType() string {
	return GoType(p.Type, p.Size)
}
