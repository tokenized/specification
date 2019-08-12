package parser2

import (
	"fmt"
	"html"
	"html/template"
	"math"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// NewSchema creates a Schema from a directory.
func NewSchema(path string) (*Schema, error) {

	// 1. Resulting Schema
	schemaFilePath := filepath.Join(path, "schema.yaml")
	var result Schema
	if err := UnmarshalFile(schemaFilePath, &result); err != nil {
		return nil, err
	}

	// 2. Schema Messages
	var messages []Message
	for _, spath := range result.MessagePaths {
		messageFilePath := filepath.Join(path, spath+".yaml")
		var message Message
		UnmarshalFile(messageFilePath, &message)
		messages = append(messages, message)
	}
	result.Messages = messages

	// 3. Schema Field Types
	var fieldTypes []FieldType
	for _, spath := range result.FieldTypePaths {
		typeFilePath := filepath.Join(path, spath+".yaml")
		var fieldType FieldType
		UnmarshalFile(typeFilePath, &fieldType)
		fieldTypes = append(fieldTypes, fieldType)
	}
	result.FieldTypes = fieldTypes

	return &result, nil
}

// UnmarshalFile
func UnmarshalFile(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to open file %s", filename))
	}

	// Decode message from yaml
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(v); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to yaml unmarshal file %s", filename))
	}

	return nil
}

func TemplateToFile(distPath string, data interface{}, inFile, outFile string) {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	tmplFuncs := template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"padding": func(str string, size int) string {
			return strings.Repeat(" ", int(math.Max(float64(size-len(str)), 0)))
		},
		"comment": func(str, chr string) string {
			return StrComment(html.UnescapeString(str), chr)
		},
		"snakecase": func(str string) string {
			return StrSnakeCase(str)
		},
		"camelcase": func(str string) string {
			return StrCamelCase(str)
		},
	}

	tmpl := template.Must(template.New(path.Base(inFile)).Funcs(tmplFuncs).ParseFiles(inFile))
	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}
}
