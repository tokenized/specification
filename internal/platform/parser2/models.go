package parser2

type Schema struct {
	Name           string      `yaml:"name"`
	Package        string      `yaml:"package"`
	Version        uint64      `yaml:"version"`
	MessagePaths   []string    `yaml:"messages"`
	FieldTypePaths []string    `yaml:"fieldTypes"`
	Messages       []Message   `yaml:"-"`
	FieldTypes     []FieldType `yaml:"-"`
	FieldAliases   []Field     `yaml:"fieldAliases"`
}

type Message struct {
	Code        string                 `yaml:"code"`
	Name        string                 `yaml:"name"`
	Label       string                 `yaml:"label"`
	Description string                 `yaml:"description"`
	MetaData    map[string]interface{} `yaml:"metadata"`
	Fields      []Field                `yaml:"fields"`
}

type FieldType struct {
	Name        string                 `yaml:"name"`
	Label       string                 `yaml:"label"`
	Description string                 `yaml:"description"`
	MetaData    map[string]interface{} `yaml:"metadata"`
	Fields      []Field                `yaml:"fields"`
}

type Field struct {
	Name        string
	Label       string
	Description string
	Type        string
	Size        int
	Example     string
	Notes       string
	IsAlias     bool
	IsFullType  bool
}
