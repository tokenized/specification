package python

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
	messages parser.Schema,
) {
	// Actions
	parser.TemplateToFile(actions, "internal/python/templates/actions.tpl",
		distPath+"/python/"+"actions_generated.py")
}
