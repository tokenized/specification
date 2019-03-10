package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileProtocol(distPath string, msgs parser.ProtocolActions, typs parser.ProtocolTypes) {

	templateToFile(distPath, msgs, "actions_test.tpl", "actions_test.go")
	templateToFile(distPath, msgs, "actions.tpl", "actions.go")
	templateToFile(distPath, typs, "field_types.tpl", "field_types.go")

}

func CompileAssets(distPath string, assts parser.Assets) {

	templateToFile(distPath, assts, "assets.tpl", "assets.go")

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/golang/templates/" + tplFile

	path := distPath + "/golang/protocol/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
