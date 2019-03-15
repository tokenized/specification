package markdown

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func CompileProtocol(distPath string,
	msgs parser.ProtocolActions,
	typs parser.ProtocolTypes,
	assets []parser.Asset) {

	for _, pm := range msgs {
		outfile := "protocol-" + parser.KebabCase(pm.Name()) + ".md"
		templateToFile(distPath, pm, "action.tpl", outfile)
	}

	for _, a := range assets {
		outfile := "asset-" + parser.KebabCase(a.Name()) + ".md"
		templateToFile(distPath, a, "asset.tpl", outfile)
	}
}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
