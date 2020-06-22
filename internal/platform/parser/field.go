package parser

import (
	"fmt"
	"strings"
)

// Field defines the expected properties of a field definition in the specification YAML.
type Field struct {
	Name           string     `yaml:"name"`
	Label          string     `yaml:"label"`
	Description    string     `yaml:"description"`
	Notes          string     `yaml:"notes"`
	Type           string     `yaml:"type"`
	Size           int        `yaml:"size"`
	ListSize       string     `yaml:"listSize"`
	VarSize        string     `yaml:"varSize"`
	Example        string     `yaml:"example"`
	Options        []string   `yaml:"options"`
	Resource       string     `yaml:"resource"`
	IsAlias        bool       `yaml:"is_alias"`
	IsCompoundType bool       `yaml:"is_compound_type"`
	AliasField     *Field     `yaml:"-"`
	CompoundField  *FieldType `yaml:"-"`
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

// BaseTypeRaw returns the raw base type
func (f *Field) BaseTypeRaw() string {
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

// BaseResource returns the base resource of the field.
func (f *Field) BaseResource() string {
	if f.AliasField != nil {
		return f.AliasField.BaseResource()
	}
	return f.Resource
}

func (f *Field) BaseOptions() []string {
	if f.AliasField != nil {
		return f.AliasField.BaseOptions()
	}
	return f.Options
}

func (f *Field) GoSingularType() string {
	gt := f.BaseType()

	if f.AliasField != nil {
		gt = f.AliasField.ProtobufType()
	} else {
		switch gt {
		case "varchar", "fixedchar":
			gt = "string"

		case "bin", "varbin":
			gt = "[]byte"

		case "uint":
			if f.Size > 4 {
				gt = "uint64"
			} else {
				gt = "uint32"
			}

		case "int":
			if f.Size > 4 {
				gt = "int64"
			} else {
				gt = "int32"
			}
		}

		if f.IsCompoundType {
			gt += "Field"
		}
	}

	return gt
}

func (f *Field) GoType() string {
	gt := f.GoSingularType()

	if f.IsList() {
		gt = "[]" + gt
	}
	return gt
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

func (f *Field) MarkdownType() string {
	baseType := f.BaseTypeRaw()

	switch baseType {
	case "fixedchar", "bin", "uint":
		baseType += "(" + fmt.Sprintf("%v", f.BaseSize()) + ")"
	case "varbin", "varchar":
		baseType += "(" + f.BaseVarSize() + ")"
	}

	if f.IsList() {
		baseType += "[" + f.BaseListSize() + "]"
	}

	return baseType
}
