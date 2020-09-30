package parser

import (
	"fmt"
	"os"
	"path/filepath"

	yamltojson "github.com/ghodss/yaml"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// NewSchema creates a Schema from a directory.
func NewSchema(path string) (*Schema, error) {

	// 1. Resulting Schema
	schemaFilePath := filepath.Join(path, "schema.yaml")
	var result Schema
	if err := unmarshalFile(schemaFilePath, &result); err != nil {
		return nil, err
	}

	// 2. Schema Messages
	var messages []Message
	for _, spath := range result.MessagePaths {
		messageFilePath := filepath.Join(path, spath+".yaml")
		var message Message
		unmarshalFile(messageFilePath, &message)
		messages = append(messages, message)
	}
	result.Messages = messages

	// 3. Schema Resources
	var resources []Resource
	for _, spath := range result.ResourcePaths {
		resourceFilePath := filepath.Join(path, spath+".yaml")
		var resource Resource
		unmarshalFile(resourceFilePath, &resource)
		for i, value := range resource.Values {
			y, err := yaml.Marshal(value.MetaData)
			if err != nil {
				return nil, errors.Wrap(err, "marshal resource metadata")
			}
			data, err := yamltojson.YAMLToJSON(y)
			if err != nil {
				return nil, errors.Wrap(err, "json resource metadata")
			}
			value.MetaDataJSON = string(data)
			resource.Values[i] = value
		}
		resources = append(resources, resource)
	}
	result.Resources = resources

	// 4. Schema Field Types
	var fieldTypes []FieldType
	for _, spath := range result.FieldTypePaths {
		typeFilePath := filepath.Join(path, spath+".yaml")
		var fieldType FieldType
		unmarshalFile(typeFilePath, &fieldType)
		fieldTypes = append(fieldTypes, fieldType)
	}
	result.FieldTypes = fieldTypes

	// 5. Post process fields
	for _, fieldType := range result.FieldTypes {
		postProcessFields(fieldType.Fields, result.FieldTypes, result.FieldAliases)
	}

	for _, message := range result.Messages {
		postProcessFields(message.Fields, result.FieldTypes, result.FieldAliases)
	}

	return &result, nil
}

// postProcessFields applies defaults, attaches alias information and detects compound fields.
func postProcessFields(fields []Field, fTypes []FieldType, aliases []Field) {
	for i, field := range fields {
		baseType := field.BaseType()

		// Field is primitive
		if field.IsPrimitive() {
			continue
		}

		// Field is an alias
		for _, alias := range aliases {
			if alias.Name == baseType {
				fields[i].AliasField = &alias
				break
			}
		}
		if fields[i].IsAlias() {
			continue
		}

		// Field must be compound
		for _, fType := range fTypes {
			if fType.Name == baseType {
				fields[i].IsCompoundType = true
				fields[i].CompoundField = &fType
				break
			}

			for i, field := range fType.Fields {
				if field.OnlyValidWhen != nil {
					for _, f := range fType.Fields {
						if f.Type != "deprecated" && f.Name == field.OnlyValidWhen.FieldName {
							field.OnlyValidWhen.FieldGoType = f.GoType()
							fType.Fields[i] = field
							break
						}
					}
					if len(field.OnlyValidWhen.FieldGoType) == 0 {
						fmt.Printf("Failed to find OnlyValidWhen field %s\n", field.OnlyValidWhen.FieldName)
					}
				}
				if field.RequiredWhen != nil {
					for _, f := range fType.Fields {
						if f.Type != "deprecated" && f.Name == field.RequiredWhen.FieldName {
							field.RequiredWhen.FieldGoType = f.GoType()
							fType.Fields[i] = field
							break
						}
					}
					if len(field.RequiredWhen.FieldGoType) == 0 {
						fmt.Printf("Failed to find RequiredWhen field %s\n", field.RequiredWhen.FieldName)
					}
				}
			}
		}

		for i, field := range fields {
			if field.OnlyValidWhen != nil {
				for _, f := range fields {
					if f.Type != "deprecated" && f.Name == field.OnlyValidWhen.FieldName {
						field.OnlyValidWhen.FieldGoType = f.GoType()
						fields[i] = field
						break
					}
				}
				if len(field.OnlyValidWhen.FieldGoType) == 0 {
					fmt.Printf("Failed to find OnlyValidWhen field %s\n", field.OnlyValidWhen.FieldName)
				}
			}
			if field.RequiredWhen != nil {
				for _, f := range fields {
					if f.Type != "deprecated" && f.Name == field.RequiredWhen.FieldName {
						field.RequiredWhen.FieldGoType = f.GoType()
						fields[i] = field
						break
					}
				}
				if len(field.RequiredWhen.FieldGoType) == 0 {
					fmt.Printf("Failed to find RequiredWhen field %s\n", field.RequiredWhen.FieldName)
				}
			}
		}
	}
}

// unmarshalFile
func unmarshalFile(filename string, v interface{}) error {
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
