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
	templateToFile(distPath, "field_types.tpl", "field_types.ts", types)
	templateToFile(distPath, "resources.tpl", "resources.ts", resources)
}

func templateToFile(distPath, tplFile, tsFile string, data interface{}) {

	tpl := "./internal/typescript/templates/" + tplFile

	path := distPath + "/typescript/src/" + tsFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
