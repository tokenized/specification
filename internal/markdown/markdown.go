package markdown

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileProtocol(distPath string, msgs parser.ProtocolActions, typs parser.ProtocolTypes) {

	for _, pm := range msgs {
		templateToFile(distPath, pm, "action.tpl", "protocol-"+parser.KebabCase(pm.Name())+".md")
	}

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
