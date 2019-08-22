package parser

// Schema represents a group of messages and their dependencies
type Schema struct {
	Name           string      `yaml:"name"`
	Package        string      `yaml:"package"`
	Version        uint64      `yaml:"version"`
	MessagePaths   []string    `yaml:"messages"`
	ResourcePaths  []string    `yaml:"resources"`
	FieldTypePaths []string    `yaml:"fieldTypes"`
	Messages       []Message   `yaml:"-"`
	Resources      []Resource  `yaml:"-"`
	FieldTypes     []FieldType `yaml:"-"`
	FieldAliases   []Field     `yaml:"fieldAliases"`
}

// Message represents a protocol action, message or asset
type Message struct {
	Code        string  `yaml:"code"`
	Name        string  `yaml:"name"`
	Label       string  `yaml:"label"`
	Description string  `yaml:"description"`
	Fields      []Field `yaml:"fields"`

	MetaData struct {
		Validation string `yaml:"validation"`
		Rejection  string `yaml:"rejection"`
		Inputs     []struct {
			Name     string `yaml:"name"`
			Label    string `yaml:"label"`
			Comments string `yaml:"comments"`
		} `yaml:"inputs"`
		Outputs []struct {
			Name     string `yaml:"name"`
			Label    string `yaml:"label"`
			Comments string `yaml:"comments"`
		} `yaml:"outputs"`
	} `yaml:"metadata"`
}

// FieldType is a compound field type definition provided by the schema
type FieldType struct {
	Name        string                 `yaml:"name"`
	Label       string                 `yaml:"label"`
	Description string                 `yaml:"description"`
	Fields      []Field                `yaml:"fields"`
	MetaData    map[string]interface{} `yaml:"metadata"`
}

// Resource is a standalone group of resource definitions provided by the schema
type Resource struct {
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"description"`
	CodeType    Field                  `yaml:"codeType"`
	Values      []ResourceValue        `yaml:"values"`
	MetaData    map[string]interface{} `yaml:"metadata"`
}

// ResourceValue is the value of a resource inside a resource group
type ResourceValue struct {
	Code        string                 `yaml:"code"`
	Name        string                 `yaml:"name"`
	Label       string                 `yaml:"label"`
	Description string                 `yaml:"description"`
	MetaData    map[string]interface{} `yaml:"metadata"`
}
