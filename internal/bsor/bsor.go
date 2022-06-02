package bsor

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
	messages parser.Schema,
) {
	templateToFile(distPath, "schema.tpl", "actions.bsor", actions)

	templateToFile(distPath, "schema.tpl", "instruments.bsor", instruments)

	templateToFile(distPath, "schema.tpl", "messages.bsor", messages)
}

func templateToFile(distPath, tplFile, bsorFile string, data interface{}) {

	tpl := "./internal/bsor/templates/" + tplFile

	path := distPath + "/bsor/" + bsorFile

	parser.TemplateToFile(data, tpl, path)
}
