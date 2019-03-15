package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileProtocol(distPath string, actions parser.ProtocolActions, messages parser.ProtocolMessages, typs parser.ProtocolTypes) {
	templateToFile(distPath, actions, "actions_test.tpl", "actions_test.go")
	templateToFile(distPath, actions, "actions.tpl", "actions.go")
	templateToFile(distPath, messages, "messages_test.tpl", "messages_test.go")
	templateToFile(distPath, messages, "messages.tpl", "messages.go")
	templateToFile(distPath, typs, "field_types.tpl", "field_types.go")
}

func CompileAssets(distPath string, assts []parser.Asset) {
	templateToFile(distPath, assts, "assets.tpl", "assets.go")
}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/golang/templates/" + tplFile
	path := distPath + "/golang/protocol/" + goFile
	parser.TemplateToFile(distPath, data, tpl, path)
}
