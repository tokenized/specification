package markdown

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

	templateToFile(distPath, "protocol-actions.tpl", "protocol-actions.md", actions)

	templateToFile(distPath, "protocol-assets.tpl", "protocol-assets.md", assets)

	templateToFile(distPath, "protocol-messages.tpl", "protocol-messages.md", messages)

	templateToFile(distPath, "protocol-field-types.tpl", "protocol-field-types.md", types)

	// for _, action := range actions {
	// 	outfile := "protocol-" + parser.KebabCase(action.Name()) + ".md"
	// 	templateToFile(distPath, "action.tpl", outfile, action)
	// }

	// for _, asset := range assets {
	// 	outfile := "asset-" + parser.KebabCase(asset.Name()) + ".md"
	// 	templateToFile(distPath, "asset.tpl", outfile, asset)
	// }

	// for _, message := range messages {
	// 	outfile := "message-" + parser.KebabCase(message.Name()) + ".md"
	// 	templateToFile(distPath, "message.tpl", outfile, message)
	// }
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
