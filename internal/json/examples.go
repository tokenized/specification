package json

import (
	"encoding/json"
	"io/ioutil"

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
	assembleActionExamples(distPath, "action_examples.json", actions, types)
}

// TODO: Get this from YAML?
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

func assembleActionExamples(distPath, jsonFile string, actions parser.ProtocolActions, types parser.ProtocolTypes) {
	top := make([]example, 0)
	b, _ := json.MarshalIndent(top, "", "  ")

	path := distPath + "/json/" + jsonFile

	err := ioutil.WriteFile(path, b, 0644)
	if err != nil {
		panic(err)
	}
}
