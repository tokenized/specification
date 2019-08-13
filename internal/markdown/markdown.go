package markdown

import (
	"github.com/tokenized/specification/internal/platform/parser2"
)

func Compile(
	distPath string,
	actions parser2.Schema,
	assets parser2.Schema,
	messages parser2.Schema,
) {

	templateToFile(distPath, "protocol-actions.tpl", "protocol-actions.md", actions)

	templateToFile(distPath, "protocol-assets.tpl", "protocol-assets.md", assets)

	templateToFile(distPath, "protocol-messages.tpl", "protocol-messages.md", messages)

	// templateToFile(distPath, "protocol-field-types.tpl", "protocol-field-types.md", types)

	// templateToFile(distPath, "protocol-resources.tpl", "protocol-resources.md", resources)
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser2.TemplateToFile(distPath, data, tpl, path)
}
