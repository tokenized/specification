package json

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tokenized/smart-contract/pkg/json"
	"github.com/tokenized/specification/internal/platform/parser"

	"github.com/ghodss/yaml"
)

// Compile converts a codec schema from YAML to JSON in raw format
func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	assets parser.Schema,
	messages parser.Schema,
) {

	schemaToFile(distPath, "actions/develop/schema.yaml", "actions.json")

	schemaToFile(distPath, "assets/develop/schema.yaml", "assets.json")

	schemaToFile(distPath, "messages/develop/schema.yaml", "messages.json")

	actionsToExample(distPath, actions, "action_examples.json")

}

// schemaToFile converts a codec schema file to its raw JSON form and writes
// it to the supplied file
func schemaToFile(distPath, schemaFile, jsonFile string) {

	inPath := "./src/" + schemaFile
	outPath := distPath + "/json/" + jsonFile

	// Convert schema to object
	schemaObj := yamlToObject(inPath)

	// Replace messages and field types with their inner contents
	basePath := filepath.Dir(inPath)
	schemaObj["messages"] = replaceFileRefs(basePath, schemaObj["messages"])
	schemaObj["fieldTypes"] = replaceFileRefs(basePath, schemaObj["fieldTypes"])
	schemaObj["resources"] = replaceFileRefs(basePath, schemaObj["resources"])

	// Beautify JSON
	body, err := json.MarshalIndent(schemaObj, "", "  ")
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

// replaceFileRefs will replace an array-interface with YAML files from a base path.
// For example, refs as:
//
//     [file1, file2, file3]
//
// Will return as:
//
//     [
//         <basePath/file1.yaml contents>,
//         <basePath/file2.yaml contents>,
//         <basePath/file3.yaml contents>,
//     ]
//
func replaceFileRefs(basePath string, refs interface{}) interface{} {

	refsArr := refs.([]interface{})
	var newRefs []interface{}

	for _, value := range refsArr {
		includeObj := yamlToObject(basePath + "/" + value.(string) + ".yaml")

		newRefs = append(newRefs, includeObj)
	}

	return newRefs
}

// yamlToObject converts a YAML file in to a general purpose Go struct
func yamlToObject(filePath string) map[string]interface{} {

	// Read YAML
	yamlContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	// Compile JSON
	body, err := yaml.YAMLToJSON(yamlContents)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	// Beautify
	var schemaObj map[string]interface{}

	if err := json.Unmarshal(body, &schemaObj); err != nil {
		panic(err)
	}

	return schemaObj
}
