package parser

import "strings"

// Field defines the expected properties of a field definition in the specification YAML.
type Field struct {
	Name           string   `yaml:"name"`
	Label          string   `yaml:"label"`
	Description    string   `yaml:"description"`
	Notes          string   `yaml:"notes"`
	Type           string   `yaml:"type"`
	Size           int      `yaml: "size"`
	Example        string   `yaml: "example"`
	Options        []string `yaml: "options`
	Resource       string   `yaml: "resource`
	IsAlias        bool     `yaml: "is_alias"`
	IsCompoundType bool     `yaml: "is_compound_type"`
	AliasField     *Field
}

// IsPrimitive returns true if the field is "primitive". Not a complex type. i.e. fieldtype or
//   message.
func (f *Field) IsPrimitive() bool {
	switch f.BaseType() {
	case "uint":
		return true
	case "int":
		return true
	case "float":
		return true
	case "bool":
		return true
	case "bin":
		return true
	case "varbin":
		return true
	case "fixedchar":
		return true
	case "varchar":
		return true
	}

	return false
}

// IsList returns true if the field is a list of objects.
func (f *Field) IsList() bool {
	return strings.HasSuffix(f.Type, "[]")
}

// BaseType returns the base type of the field, with no modifiers like list type.
func (f *Field) BaseType() string {
	return strings.Replace(f.Type, "[]", "", 1)
}

func (f *Field) ProtobufType() string {
	pbt := ""
	if f.IsList() {
		pbt += "repeated "
	}

	// TODO: Convert more to proto
	baseType := f.BaseType()

	switch baseType {
	case "varchar", "fixedchar":
		baseType = "string"

	case "bin", "varbin":
		baseType = "bytes"
	}

	if f.IsCompoundType {
		pbt += baseType + "Field"
	} else {
		pbt += baseType
	}

	return pbt
}

/*
// HasVariableSize returns true if the field is variable in length.
func (f *Field) HasVariableSize() bool {
	return f.IsList() || f.Type == "varbin" || f.Type == "varchar"
}

// MaxVariableSize returns the maximum length, in bytes, of a variable length field.
func (f *Field) MaxVariableSize() uint64 {
	return (1 << (uint64(f.Size) * 8)) - 1
}

// VariableSizeType returns a field that represents the type used to encode the size of a
//   variable length field.
func (f *Field) VariableSizeType() *Field {
	return &Field{
		Type: "uint",
		Size: f.Size,
	}
}
*/
