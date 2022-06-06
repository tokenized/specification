package bsor

import (
	"encoding/json"
	"os"

	"github.com/tokenized/specification/internal/platform/parser"

	"github.com/pkg/errors"
)

type BSORDefinition struct {
	Version  uint                            `json:"version"`
	Messages map[string]BSORStructDefinition `json:"messages"`
}

type BSORStructDefinition []BSORFieldDefinition

type BSORFieldDefinition struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	ID        uint   `json:"id"`
	IsArray   bool   `json:"is_array,omitempty"`
	FixedSize uint   `json:"fixed_size,omitempty"`
}

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
	messages parser.Schema,
) {
	toFile(distPath, "actions.bsor", actions)

	toFile(distPath, "instruments.bsor", instruments)

	toFile(distPath, "messages.bsor", messages)
}

func toFile(distPath, filename string, data parser.Schema) error {
	definition := BSORDefinition{
		Version:  0,
		Messages: make(map[string]BSORStructDefinition),
	}

	for _, message := range data.Messages {
		fields := make(BSORStructDefinition, 0, len(message.Fields))
		for i, field := range message.Fields {
			name, fixedSize := field.BSORFullType()
			if name == "deprecated" {
				continue
			}
			fld := BSORFieldDefinition{
				Name:      field.Name,
				Type:      name,
				ID:        uint(i + 1),
				IsArray:   field.IsList(),
				FixedSize: fixedSize,
			}

			fields = append(fields, fld)
		}

		definition.Messages[message.Name] = fields
	}

	for _, fieldType := range data.FieldTypes {
		fields := make(BSORStructDefinition, 0, len(fieldType.Fields))
		for i, field := range fieldType.Fields {
			name, fixedSize := field.BSORFullType()
			if name == "deprecated" {
				continue
			}
			fld := BSORFieldDefinition{
				Name:      field.Name,
				Type:      name,
				ID:        uint(i + 1),
				IsArray:   field.IsList(),
				FixedSize: fixedSize,
			}

			fields = append(fields, fld)
		}

		definition.Messages[fieldType.Name+"Field"] = fields
	}

	outputPath := distPath + "/bsor/" + filename
	file, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(definition); err != nil {
		return errors.Wrap(err, "marshal")
	}

	return nil
}
