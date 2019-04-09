package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	distPath string,
	actions parser.ProtocolActions,
	messages parser.ProtocolMessages,
	types parser.ProtocolTypes,
	resources parser.ProtocolResources,
	rejectionCodes parser.ProtocolRejectionCodes,
	assets []parser.Asset,
) {
	templateToFile(distPath, "actions_test.tpl", "actions_test.go", actions)

	templateToFile(distPath, "actions.tpl", "actions.go", actions)

	templateToFile(distPath, "messages_test.tpl", "messages_test.go", messages)

	templateToFile(distPath, "messages.tpl", "messages.go", messages)

	templateToFile(distPath, "field_types.tpl", "field_types.go", types)

	templateToFile(distPath, "resources.tpl", "resources.go", resources)

	templateToFile(distPath, "rejection_codes.tpl", "rejection_codes.go", rejectionCodes)

	templateToFile(distPath, "assets.tpl", "assets.go", assets)
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/golang/templates/" + tplFile

	path := distPath + "/golang/protocol/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
