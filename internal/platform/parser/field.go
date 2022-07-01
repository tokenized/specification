package parser

import (
	"fmt"
	"strings"
)

// Field defines the expected properties of a field definition in the specification YAML.
type Field struct {
	Name           string              `yaml:"name"`
	Label          string              `yaml:"label"`
	Description    string              `yaml:"description"`
	Notes          string              `yaml:"notes"`
	Type           string              `yaml:"type"`
	Size           int                 `yaml:"size"`
	ListSize       string              `yaml:"listSize"`
	VarSize        string              `yaml:"varSize"`
	Example        string              `yaml:"example"`
	Options        []string            `yaml:"options"`
	OnlyValidWhen  *RelatedFieldValues `yaml:"only_valid_when"` // specifies a field is only valid when another field has specific values
	RequiredWhen   *RelatedFieldValues `yaml:"required_when"`   // specifies a field is required when another field has specific values
	Required       bool                `yaml:"required"`
	Resource       string              `yaml:"resource"`
	IsCompoundType bool                `yaml:"is_compound_type"`
	AliasField     *Field              `yaml:"-"`
	CompoundField  *FieldType          `yaml:"-"`
}

// RelatedFieldValues specifies that a field is only valid to be specified when another field has one of
// the listed values.
type RelatedFieldValues struct {
	FieldName   string   `yaml:"field_name"`
	Values      []string `yaml:"values"`
	FieldGoType string
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

// IsAlias returns true if the field has an alias type.
func (f *Field) IsAlias() bool {
	return f.AliasField != nil
}

func (f *Field) HasOnlyValidWhen() bool {
	return f.OnlyValidWhen != nil
}

func (f *Field) HasRequiredWhen() bool {
	return f.RequiredWhen != nil
}

// BaseType returns the base type of the field, with no modifiers like list type.
func (f *Field) BaseType() string {
	if f.AliasField != nil {
		return f.AliasField.BaseType()
	}
	return strings.Replace(f.Type, "[]", "", 1)
}

// BaseExample returns the base example of the field.
func (f *Field) BaseExample() string {
	if f.AliasField != nil {
		return f.AliasField.BaseExample()
	}
	return f.Example
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
		gt = f.AliasField.GoSingularType()
	} else {
		switch gt {
		case "varchar", "fixedchar":
			gt = "string"

		case "bin", "varbin":
			gt = "[]byte"

		case "uint":
			if f.Size > 4 {
				gt = "uint64"
			} else if f.Size > 2 {
				gt = "uint32"
			} else if f.Size > 1 {
				gt = "uint16"
			} else {
				gt = "uint8"
			}

		case "int":
			if f.Size > 4 {
				gt = "int64"
			} else if f.Size > 2 {
				gt = "int32"
			} else if f.Size > 1 {
				gt = "int16"
			} else {
				gt = "int8"
			}
		}

		if f.IsCompoundType {
			gt += "Field"
		}
	}

	return gt
}

func (f *Field) GoSingularTypeWithPointer() string {
	gt := f.BaseType()

	if f.AliasField != nil {
		gt = f.AliasField.GoSingularType()
	} else {
		switch gt {
		case "varchar", "fixedchar":
			gt = "string"

		case "bin", "varbin":
			gt = "[]byte"

		case "uint":
			if f.Size > 4 {
				gt = "uint64"
			} else if f.Size > 2 {
				gt = "uint32"
			} else if f.Size > 1 {
				gt = "uint16"
			} else {
				gt = "uint8"
			}

		case "int":
			if f.Size > 4 {
				gt = "int64"
			} else if f.Size > 2 {
				gt = "int32"
			} else if f.Size > 1 {
				gt = "int16"
			} else {
				gt = "int8"
			}
		}

		if f.IsCompoundType {
			gt = "*" + gt + "Field"
		}
	}

	return gt
}

func (f *Field) GoType() string {
	gt := f.GoSingularType()

	if f.IsList() {
		gt = "[]" + f.GoSingularTypeWithPointer()
	}
	return gt
}

func (f *Field) GoTypeWithPointer() string {
	gt := f.GoSingularType()

	if f.IsList() {
		gt = "[]" + f.GoSingularTypeWithPointer()
	} else if f.IsCompoundType {
		gt = "*" + gt
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
		} else if f.Size > 2 {
			baseType = "uint32"
		} else if f.Size > 1 {
			baseType = "uint16"
		} else {
			baseType = "uint8"
		}

	case "int":
		if f.Size > 4 {
			baseType = "int64"
		} else if f.Size > 2 {
			baseType = "int32"
		} else if f.Size > 1 {
			baseType = "int16"
		} else {
			baseType = "int8"
		}
	}

	if f.IsCompoundType {
		pbt += baseType + "Field"
	} else {
		pbt += baseType
	}

	return pbt
}

func (f *Field) BSORType() string {
	result := ""
	if f.IsList() {
		if f.Size > 0 {
			result += fmt.Sprintf("[%d]", f.Size) // fixed size list
		} else {
			result += "[]" // variable size list
		}
	}

	baseType := f.BaseType()

	if f.AliasField != nil {
		return result + f.AliasField.BSORType()
	}

	switch baseType {
	case "varchar":
		baseType = "string"

	case "fixedchar":
		baseType = fmt.Sprintf("string(%d)", f.Size)

	case "varbin":
		baseType = "binary"

	case "bin":
		baseType = fmt.Sprintf("binary(%d)", f.Size)

	case "uint":
		if f.Size > 4 {
			baseType = "uint64"
		} else if f.Size > 2 {
			baseType = "uint32"
		} else if f.Size > 1 {
			baseType = "uint16"
		} else {
			baseType = "uint8"
		}

	case "int":
		if f.Size > 4 {
			baseType = "int64"
		} else if f.Size > 2 {
			baseType = "int32"
		} else if f.Size > 1 {
			baseType = "int16"
		} else {
			baseType = "int8"
		}
	}

	if f.IsCompoundType {
		result += baseType + "Field"
	} else {
		result += baseType
	}

	return result
}

// BSORFullType returns the BSOR base type name and the fixed size if it is fixed size.
func (f *Field) BSORFullType() (string, uint) {
	if f.AliasField != nil {
		return f.AliasField.BSORFullType()
	}

	baseType := f.BaseType()
	switch baseType {
	case "varchar":
		return "string", 0

	case "fixedchar":
		return "string", uint(f.Size)

	case "varbin":
		return "binary", 0

	case "bin":
		return "binary", uint(f.Size)

	case "uint":
		if f.Size > 4 {
			return "uint64", 0
		} else if f.Size > 2 {
			return "uint32", 0
		} else if f.Size > 1 {
			return "uint16", 0
		} else {
			return "uint8", 0
		}

	case "int":
		if f.Size > 4 {
			return "int64", 0
		} else if f.Size > 2 {
			return "int32", 0
		} else if f.Size > 1 {
			return "int16", 0
		} else {
			return "int8", 0
		}
	}

	if f.IsCompoundType {
		return baseType + "Field", 0
	}

	return baseType, 0
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
