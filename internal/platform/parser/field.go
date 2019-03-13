package parser

import (
	"fmt"
	"strings"
)

type Field struct {
	Name         string
	Label        string
	Description  string
	Type         string
	Size         uint64
	Required     bool
	ExampleValue string `yaml:"example_value"`
	ExampleHex   string `yaml:"example_hex"`
	Notes        string
}

func (f Field) FieldName() string {
	return strings.Replace(f.Name, " ", "", -1)
}

func (f Field) FieldDescription() string {
	return strings.Trim(f.Description, " ")
}

func (f Field) FieldBytes() string {
	return fmt.Sprintf("%v", f.Size)
}

func (f Field) FormType() string {
	if f.Name == "Payload" {
		return "json.RawAction"
	}

	if f.Type == "[]byte" || f.Type == "byte" {
		return "string"
	}

	return f.Type
}

func (f Field) Tags() string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", f.JSONFriendly())
}

func (f Field) JSONFriendly() string {
	snake := matchFirstCap.ReplaceAllString(f.Name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (f Field) IsData() bool {
	return f.Type == "bin16"
}

func (f Field) IsBytes() bool {
	return f.FieldGoType() == "[]byte" && !f.IsData()
}

func (f Field) IsVarChar() bool {
	return f.Type == "varchar"
}

func (f Field) IsFixedChar() bool {
	return f.Type == "fixedchar" && f.Size != 1
}

func (f Field) IsVarBin() bool {
	return f.Type == "varbin"
}

func (f Field) IsNumeric() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(s, "uint") ||
		strings.HasPrefix(s, "float") ||
		strings.HasPrefix(s, "int") ||
		s == "time"
}

func (f Field) IsFloat() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(f.Type, s)
}

func (f Field) Length() uint64 {
	return f.Size
}

func (f Field) SingularType() string {
	return strings.Replace(f.FieldGoType(), "[]", "", 1)
}

func (f Field) FieldGoType() string {
	return GoType(f.Type, f.Size)
}

func (f Field) IsInternalTypeArray() bool {
	return f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsNativeTypeArray() bool {
	return !f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsInternalType() bool {
	return IsInternalType(f.Type, f.Size)
}

func (f Field) IsComplexType() bool {
	return f.IsInternalType()
}

func (f Field) Trimmable() bool {
	if f.Name == "Payload" {
		return false
	}

	if f.Type == "STRING" && f.Length() > 1 {
		return true
	}

	if f.Type == "SHA" {
		return true
	}

	return false
}
