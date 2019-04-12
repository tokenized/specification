package python

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
	templateToFile(distPath, "actions.tpl", "actions.py", actions)
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/python/templates/" + tplFile

	path := distPath + "/python/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
