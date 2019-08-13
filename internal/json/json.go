package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tokenized/specification/internal/platform/parser2"

	"github.com/ghodss/yaml"
)

// Compile converts a codec schema from YAML to JSON in raw format
func Compile(
	distPath string,
	actions parser2.Schema,
	assets parser2.Schema,
	messages parser2.Schema,
) {

	schemaToFile(distPath, "actions/develop/schema.yaml", "protocol.json")

	schemaToFile(distPath, "assets/develop/schema.yaml", "assets.json")

	schemaToFile(distPath, "messages/develop/schema.yaml", "messages.json")

	// TODO handle resources
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
