package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tokenized/specification/internal/platform/parser"
)

func actionsToExample(distPath string, schema parser.Schema, jsonFile string) {
	outPath := distPath + "/json/" + jsonFile

	var examples []ActionExample

	for _, message := range schema.Messages {

		fields := make(map[string]interface{})

		for _, field := range message.Fields {
			fields[field.Name] = field.Example
		}

		example := ActionExample{
			Code:   message.Code,
			Label:  message.Label,
			Name:   message.Name,
			Fields: fields,
		}

		examples = append(examples, example)
	}

	// Beautify JSON
	body, err := json.MarshalIndent(examples, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write file
	file, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write(body); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if err := file.Close(); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}

type ActionExample struct {
	Code   string                 `json:"code"`
	Label  string                 `json:"label"`
	Name   string                 `json:"name"`
	Fields map[string]interface{} `json:"fields"`
}
