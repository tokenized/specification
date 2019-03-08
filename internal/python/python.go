package python

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(distPath string, msgs parser.Actions) {

	templateToFile(distPath, msgs, "actions.tpl", "actions.py")

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/python/templates/" + tplFile

	path := distPath + "/python/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
