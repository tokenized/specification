package typescript

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tokenized/specification/internal/platform/parser"
)

const arraySuffix = "[]"

type actionHeader struct {
	Version    int64  `json:"version"`
	ActionCode string `json:"actionCode"`
}

type example struct {
	Code   string     `json:"code"`
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
	examplesToJsonFile(distPath, "actions.json", actions, types)
	templateToTestFile(distPath, "actions.test.tpl", "actions.test.ts", actions)
	templateToFile(distPath, "assets.tpl", "assets.ts", assets)
	templateToFile(distPath, "field_types.tpl", "field_types.ts", types)
	templateToFile(distPath, "resource_data.tpl", "resource_data.ts", resources)
	templateToFile(distPath, "rejection_codes.tpl", "rejection_codes.ts", rejectionCodes)
	templateToFile(distPath, "messages.tpl", "messages.ts", messages)
	templateToTestFile(distPath, "messages.test.tpl", "messages.test.ts", messages)
}

func exampleForInternalType(fieldTypeName string, size uint64) interface{} {
	switch fieldTypeName {
	case "AssetCode":
		return "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
	case "ContractCode":
		return "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
	default:
		return fmt.Sprintf("NO EXAMPLE - internal type (%s)", fieldTypeName)
	}
}

func exampleForInternalTypeArray(singularFieldTypeName string, size uint64, types parser.ProtocolTypes) interface{} {
	switch singularFieldTypeName {
	case "QuantityIndex":
		{
			count := 2
			q := make([]map[string]interface{}, count)

			for _, protocolType := range types {
				if protocolType.Name() == singularFieldTypeName {
					fmt.Println(singularFieldTypeName)
					for i := 0; i < count; i++ {
						item := make(map[string]interface{})
						for _, field := range protocolType.Fields {
							fmt.Println(field.Name)
							item[lowerFirstCharacter(field.Name)] = "hello"
						}
						q[i] = item
					}

					return q
				}
			}

		}
		return fmt.Sprintf("NO EXAMPLE - internal type array (%s) Example partially constructed.", singularFieldTypeName)
	default:
		return fmt.Sprintf("NO EXAMPLE - internal type array (%s)", singularFieldTypeName)
	}
}

func exampleForNativeType(fieldTypeName string, size uint64) interface{} {

	var nativeTypeExample interface{}

	switch fieldTypeName {
	case "bool":
		nativeTypeExample = true
	case "uint":
		switch size {
		case 2:
			nativeTypeExample = "2345"
		case 4:
			nativeTypeExample = "456789"
		case 8:
			nativeTypeExample = "8901234567"
		default:
			nativeTypeExample = fmt.Sprintf("NO EXAMPLE - native type (%s, size %d)", fieldTypeName, size)
		}
	case "varbin":
		{
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, size)
			if size < 8 {
				nativeTypeExample = hex.EncodeToString(b[0:size])
			} else {
				nativeTypeExample = hex.EncodeToString(b)
			}
		}
	case "varchar":
		{
			b := make([]byte, size)
			for i := 0; i < int(size); i++ {
				b[i] = byte(65 + i%26)
			}
			nativeTypeExample = string(b)
		}
	default:
		nativeTypeExample = fmt.Sprintf("NO EXAMPLE - native type (%s)", fieldTypeName)
	}
	return nativeTypeExample
}

func exampleForNativeTypeArray(singularFieldTypeName string, size uint64) interface{} {
	switch singularFieldTypeName {
	case "uint16":
		return [3]string{"0", "1", "2"}
	case "uint64":
		return [3]string{"8500000000", "8600000000", "8700000000"}
	default:
		return fmt.Sprintf("NO EXAMPLE - native type array (%s)", singularFieldTypeName)
	}
}

func examplesToJsonFile(distPath, jsonFile string, actions parser.ProtocolActions, types parser.ProtocolTypes) {

	top := make([]example, 0)

	missingExamples := 0

	for _, action := range actions {
		contents := make(map[string]interface{})

		// Depth is going to be finite
		for _, field := range action.Fields {
			fieldName := lowerFirstCharacter(field.FieldName())
			fieldTypeName := field.Type

			// TODO: Handle Arrays
			// Singular type

			if field.Example != "" {
				contents[fieldName] = field.Example
			} else {
				if field.IsNativeTypeArray() {
					//contents[fieldName] = fmt.Sprintf("NO EXAMPLE - native type array (%s)", fieldTypeName)
					contents[fieldName] = exampleForNativeTypeArray(field.SingularType(), field.Size)
				} else if field.IsNativeType() {
					contents[fieldName] = exampleForNativeType(fieldTypeName, field.Size)
				} else if field.IsInternalTypeArray() {
					contents[fieldName] = exampleForInternalTypeArray(field.SingularType(), field.Size, types)
					missingExamples++
				} else if field.IsInternalType() {
					contents[fieldName] = exampleForInternalType(fieldTypeName, field.Size)
					missingExamples++
				} else {
					contents[fieldName] = fmt.Sprintf("NO EXAMPLE (%s)", fieldTypeName)
					missingExamples++
				}
				/*
					switch field.Type {
					case "uint":
						contents[fieldName] = fmt.Sprintf("%v", field.Size)
					default:
						contents[fieldName] = "NO EXAMPLE"
						if strings.HasSuffix(field.Type, arraySuffix) {
							individualTypeName := field.Type[:len(field.Type)-len(arraySuffix)]
							for _, protocolType := range types {
								if protocolType.Name() == individualTypeName {
									contents[fieldName] = fmt.Sprintf("NO EXAMPLE, field count: %v", len(protocolType.Fields))
								}
							}

						} // else {
						//	contents[fieldName] = "NO EXAMPLE"
						//}

					}
				*/
			}

			//example := getExample(field, types)
			//contents[fieldName] = example[example] // Do this for every type
		}

		// TODO: Create header from header fields definition

		top = append(top, example{
			Code: action.Code,
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

	fmt.Println(fmt.Sprintf("Missing examples: %v", missingExamples))

	b, _ := json.MarshalIndent(top, "", "  ")

	path := distPath + "/typescript/src/" + jsonFile

	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}

}

/*
func getExample(field Field, protocolTypes ProtocolTypes) map[string]interface{} {
	contents := make(map[string]interface{})
	example := make(map[string]interface{})

	if field.Example != "" {
			example[fieldName] = field.Example
		} else {
			getExample(field.Type)
			switch field.Type {
			case "uint":
				contents[fieldName] = fmt.Sprintf("%v", field.Size)
			default:
				contents[fieldName] = "NO EXAMPLE"
				if strings.HasSuffix(field.Type, arraySuffix) {
					individualTypeName := field.Type[:len(field.Type)-len(arraySuffix)]
					for _, protocolType := range types {
						if protocolType.Name() == individualTypeName {
							contents[fieldName] = fmt.Sprintf("NO EXAMPLE, field count: %v", len(protocolType.Fields))
						}
					}

				} // else {
				//	contents[fieldName] = "NO EXAMPLE"
				//}
			}
		}
	}
	contents["example"] = example
	return contents
}
*/

func lowerFirstCharacter(pascalCase string) string {
	camelCase := strings.ToLower(pascalCase[0:1]) + string(pascalCase[1:])
	return camelCase
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
