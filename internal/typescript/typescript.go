package typescript

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
	templateToFile(distPath, "actions.tpl", "actions.ts", actions)
	templateToTestFile(distPath, "actions.test.tpl", "actions.test.ts", actions)
	templateToFile(distPath, "assets.tpl", "assets.ts", assets)
	templateToFile(distPath, "field_types.tpl", "field_types.ts", types)
	templateToFile(distPath, "resource_data.tpl", "resource_data.ts", resources)
	templateToFile(distPath, "rejection_codes.tpl", "rejection_codes.ts", rejectionCodes)
	templateToFile(distPath, "messages.tpl", "messages.ts", messages)
	templateToTestFile(distPath, "messages.test.tpl", "messages.test.ts", messages)
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
