package parser

import (
	"strings"
)

// Field defines the expected properties of a field definition in the specification YAML.
type Field struct {
	Name           string   `yaml:"name"`
	Label          string   `yaml:"label"`
	Description    string   `yaml:"description"`
	Notes          string   `yaml:"notes"`
	Type           string   `yaml:"type"`
	Size           int      `yaml:"size"`
	ListSize       string   `yaml:"listSize"`
	VarSize        string   `yaml:"varSize"`
	Example        string   `yaml:"example"`
	Options        []string `yaml:"options`
	Resource       string   `yaml:"resource`
	IsAlias        bool     `yaml:"is_alias"`
	IsCompoundType bool     `yaml:"is_compound_type"`
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
	if f.AliasField != nil {
		return f.AliasField.BaseType()
	}
	return strings.Replace(f.Type, "[]", "", 1)
}

// BaseSize returns the size of the field's base type, the alias type if there is one.
func (f *Field) BaseSize() int {
	if f.Size != 0 {
		return f.Size
	}
	if f.AliasField != nil {
		return f.AliasField.BaseSize()
	}
	if f.Size == 0 {
		return 1
	}
	return f.Size
}

// BaseListSize returns the list size of the field's base type, the alias type if there is one.
func (f *Field) BaseListSize() string {
	if len(f.ListSize) != 0 {
		return f.ListSize
	}
	if f.AliasField != nil {
		return f.AliasField.BaseListSize()
	}
	if len(f.ListSize) == 0 {
		return "tiny"
	}
	return f.ListSize
}

// BaseVarSize returns the variable size of the field's base type, the alias type if there is one.
func (f *Field) BaseVarSize() string {
	if f.AliasField != nil {
		return f.AliasField.BaseVarSize()
	}
	if len(f.VarSize) == 0 {
		return "tiny"
	}
	return f.VarSize
}

func (f *Field) ProtobufType() string {
	pbt := ""
	if f.IsList() {
		pbt += "repeated "
	}

	baseType := f.BaseType()

	if f.AliasField != nil {
		return pbt + f.AliasField.ProtobufType()
	}

	switch baseType {
	case "varchar", "fixedchar":
		baseType = "string"

	case "bin", "varbin":
		baseType = "bytes"

	case "uint":
		if f.Size > 4 {
			baseType = "uint64"
		} else {
			baseType = "uint32"
		}

	case "int":
		if f.Size > 4 {
			baseType = "int64"
		} else {
			baseType = "int32"
		}
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
