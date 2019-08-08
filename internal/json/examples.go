package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileExamples(
	distPath string,
	actions parser.ProtocolActions,
	messages parser.ProtocolMessages,
	types parser.ProtocolTypes,
	resources parser.ProtocolResources,
	rejectionCodes parser.ProtocolRejectionCodes,
	assets []parser.Asset,
) {
	actionSchema := yamlToObject("./src/protocol/develop/schema.yaml")
	err := assembleActionExamples(distPath, "action_examples.json", actions, types, actionSchema)
	if err != nil {
		fmt.Println(err)
	}
}

// TODO: Get this from YAML? I'll need the schema
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
	Header         map[string]interface{} `json:"header"`
	ActionContents map[string]interface{} `json:"actionContents"`
}

func assembleActionExamples(distPath, jsonFile string, actions parser.ProtocolActions, types parser.ProtocolTypes, actionSchema map[string]interface{}) error {
	top := make([]example, len(actions))

	headerSchema := actionSchema["headerFields"]

	headerFields, ok := headerSchema.([]interface{})
	if !ok {
		return errors.New("failed to get header fields from schema")
	}

	for actionIndex, action := range actions {
		header, headerErr := generateHeader(action, headerFields)
		if headerErr != nil {
			return headerErr
		}
		contents := make(map[string]interface{})

		for _, field := range action.Fields {
			fieldName := lowerFirstCharacter(field.FieldName())
			fieldTypeName := field.Type
			var fieldErr error

			if field.Example != "" {
				if field.IsNativeType() {
					if contents[fieldName], fieldErr = exampleForNativeType(field.Example, fieldTypeName, field.Size); fieldErr != nil {
						return fieldErr
					}
				} else if field.IsNativeTypeArray() {
					if contents[fieldName], fieldErr = exampleForNativeTypeArray(field.Example, field.SingularType()); fieldErr != nil {
						return fieldErr
					}
				} else {
					contents[fieldName] = field.Example
				}
				continue
			}

			fmt.Println(fmt.Sprintf("Type %s IsInternalTypeArray(): %t", fieldTypeName, field.IsInternalTypeArray()))
			if field.IsInternalTypeArray() {
				if contents[fieldName], fieldErr = exampleForInternalTypeArray(field.SingularType(), types); fieldErr != nil {
					return fieldErr
				}
			} else if field.IsNativeType() {
				contents[fieldName] = fmt.Sprintf("NO EXAMPLE (%s) - native type", fieldTypeName)
			} else if field.IsInternalType() {
				contents[fieldName] = fmt.Sprintf("NO EXAMPLE (%s) - internal type", fieldTypeName)
			} else {
				contents[fieldName] = fmt.Sprintf("NO EXAMPLE (%s)", fieldTypeName)
			}
		}

		top[actionIndex] = example{
			Code: action.Code,
			Name: action.Metadata.Name,
			Sample: sampleJSON{
				Header:         header,
				ActionContents: contents,
			},
		}
	}

	b, _ := json.MarshalIndent(top, "", "  ")

	path := distPath + "/json/" + jsonFile

	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

func exampleForInternalType(example string, typeName string) (interface{}, error) {
	switch typeName {
	case "uint16":
		{
			first, err := strconv.ParseUint(example, 10, 16)
			if err != nil {
				return nil, err
			}
			return []uint16{uint16(first)}, nil
		}
	case "uint64":
		{
			first, err := strconv.ParseUint(example, 10, 64)
			if err != nil {
				return nil, err
			}
			return []uint64{first}, nil
		}
	default:
		return example, nil
	}
}

func exampleForInternalTypeArray(singularTypeName string, types parser.ProtocolTypes) (interface{}, error) {
	example := make(map[string]interface{})

	switch singularTypeName {
	case "AssetSettlement":
		{
			for _, protocolType := range types {
				if protocolType.Name() == singularTypeName {
					for _, field := range protocolType.Fields {
						camelCaseFieldName := lowerFirstCharacter(field.Name)
						var errExample error
						if field.IsInternalTypeArray() {
							if example[camelCaseFieldName], errExample = exampleForInternalTypeArray(field.SingularType(), types); errExample != nil {
								return nil, errExample
							}
						} else {
							if example[camelCaseFieldName], errExample = exampleForInternalType(field.Example, field.Type); errExample != nil {
								return nil, errExample
							}
						}

					}
					return []map[string]interface{}{example}, nil
				}
			}
			return nil, fmt.Errorf("NO EXAMPLE (%s) - not found for internal type array", singularTypeName)
		}

	case "QuantityIndex":
		{

			var errExample error
			for _, protocolType := range types {
				if protocolType.Name() == singularTypeName {
					for _, field := range protocolType.Fields {
						if example[lowerFirstCharacter(field.Name)], errExample = exampleForInternalType(field.Example, field.Type); errExample != nil {
							return nil, errExample
						}

					}
					return []map[string]interface{}{example}, nil
				}
			}

			return nil, fmt.Errorf("internal type (%s) not found for array", singularTypeName)
		}
	default:
		return fmt.Sprintf("NO EXAMPLE (%s) - internal type array", singularTypeName), nil
	}
}

func exampleForNativeType(example string, typeName string, size uint64) (interface{}, error) {
	switch typeName {
	case "bool":
		return (example == "true"), nil
	case "uint":
		{
			if size == 8 {
				return example, nil // So it can be read as a big number
			}

			exampleUint, err := strconv.ParseUint(example, 10, 32)
			if err != nil {
				return nil, err
			}
			return exampleUint, nil

		}

	default:
		return example, nil
	}
}

func exampleForNativeTypeArray(example string, singularTypeName string) (interface{}, error) {
	switch singularTypeName {
	case "uint16":
		{
			first, err := strconv.ParseUint(example, 10, 16)
			if err != nil {
				return nil, err
			}
			return []uint16{uint16(first)}, nil
		}
	case "uint64":
		{
			first, err := strconv.ParseUint(example, 10, 64)
			if err != nil {
				return nil, err
			}
			return []uint64{first}, nil
		}
	default:
		return example, nil
	}
}

func generateHeader(action parser.ProtocolAction, headerFields []interface{}) (map[string]interface{}, error) {
	header := make(map[string]interface{})

	for _, headerField := range headerFields {
		if headerFieldMap, ok := headerField.(map[string]interface{}); ok {

			if headerFieldName, ok := headerFieldMap["name"].(string); ok {
				switch headerFieldName {
				case "ActionCode":
					header[lowerFirstCharacter(headerFieldName)] = action.Code
				case "Version":
					header[lowerFirstCharacter(headerFieldName)] = action.Version
				default:
					return nil, fmt.Errorf("Unhandled header field: %s", headerFieldName)
				}
			}
		}
	}

	return header, nil
}

func lowerFirstCharacter(pascalCase string) string {
	camelCase := strings.ToLower(pascalCase[0:1]) + string(pascalCase[1:])
	return camelCase
}
