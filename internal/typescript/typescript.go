package typescript

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/tokenized/specification/internal/platform/parser"
)

type actionHeader struct {
	Version    int64  `json:"version"`
	ActionCode string `json:"actionCode"`
}

type example struct {
	Name   string     `json:"name"`
	Sample sampleJSON `json:"sample"`
}

type sampleJSON struct {
	Header         actionHeader           `json:"header"`
	ActionContents map[string]interface{} `json:"actionContents"`
}

func Compile(
	distPath string,
	actions parser.ProtocolActions,
	messages parser.ProtocolMessages,
	types parser.ProtocolTypes,
	resources parser.ProtocolResources,
	rejectionCodes parser.ProtocolRejectionCodes,
	assets []parser.Asset,
) {
	templateToFile(distPath, "actions.tpl", "actions.ts", actions)
	jsonToFile(distPath, "actions.json", actions)
	templateToTestFile(distPath, "actions.test.tpl", "actions.test.ts", actions)
	templateToFile(distPath, "assets.tpl", "assets.ts", assets)
	templateToFile(distPath, "field_types.tpl", "field_types.ts", types)
	templateToFile(distPath, "resource_data.tpl", "resource_data.ts", resources)
	templateToFile(distPath, "rejection_codes.tpl", "rejection_codes.ts", rejectionCodes)
	templateToFile(distPath, "messages.tpl", "messages.ts", messages)
	templateToTestFile(distPath, "messages.test.tpl", "messages.test.ts", messages)
}

func jsonToFile(distPath, jsonFile string, actions parser.ProtocolActions) {

	top := make([]example, 0)

	for _, action := range actions {
		contents := make(map[string]interface{})

		for _, field := range action.Fields {
			// TODO: Handle Arrays
			if field.Example != "" {
				contents[field.FieldName()] = field.Example
			} else if field.ExampleHex != "" {
				contents[field.FieldName()] = field.Example
			} else {

				switch field.Type {
				case "uint":
					contents[field.FieldName()] = fmt.Sprintf("%v", field.Size)
				default:
					contents[field.FieldName()] = "NO EXAMPLE"
				}
			}

		}

		// TODO: Create header from header fields definition

		top = append(top, example{
			Name: action.Metadata.Name,
			Sample: sampleJSON{
				Header: actionHeader{
					Version:    1,
					ActionCode: action.Code,
				},
				ActionContents: contents,
			},
		})
	}

	b, _ := json.MarshalIndent(top, "", "  ")

	path := distPath + "/typescript/src/" + jsonFile

	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}

}

func templateToFile(distPath, tplFile, tsFile string, data interface{}) {

	tpl := "./internal/typescript/templates/" + tplFile

	path := distPath + "/typescript/src/" + tsFile

	parser.TemplateToFile(distPath, data, tpl, path)
}

func templateToTestFile(distPath, tplFile, tsFile string, data interface{}) {

	tpl := "./internal/typescript/templates/" + tplFile

	path := distPath + "/typescript/test/" + tsFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
