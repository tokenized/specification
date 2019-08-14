package protobuf

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	assets parser.Schema,
	messages parser.Schema,
) {

	templateToFile(distPath, "schema.tpl", "actions.proto", actions)

	templateToFile(distPath, "schema.tpl", "assets.proto", assets)

	templateToFile(distPath, "schema.tpl", "messages.proto", messages)
}

func templateToFile(distPath, tplFile, protoFile string, data interface{}) {

	tpl := "./internal/protobuf/templates/" + tplFile

	path := distPath + "/protobuf/" + protoFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
