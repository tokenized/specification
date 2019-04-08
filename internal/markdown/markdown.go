package markdown

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileProtocol(distPath string,
	actions parser.ProtocolActions,
	typs parser.ProtocolTypes,
	assets []parser.Asset) {

	for _, action := range actions {
		outfile := "protocol-" + parser.KebabCase(action.Name()) + ".md"
		templateToFile(distPath, pm, "action.tpl", outfile)
	}

	for _, asset := range assets {
		outfile := "asset-" + parser.KebabCase(asset.Name()) + ".md"
		templateToFile(distPath, a, "asset.tpl", outfile)
	}

	for _, message := range messages {
		outfile := "message-" + parser.KebabCase(message.Name()) + ".md"
		templateToFile(distPath, m, "message.tpl", outfile)
	}
}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
