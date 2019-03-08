package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(distPath string, msgs parser.Actions) {

	templateToFile(distPath, msgs, "actions_test.tpl", "actions_test.go")
	templateToFile(distPath, msgs, "actions.tpl", "actions.go")

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/golang/templates/" + tplFile

	path := distPath + "/golang/protocol/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
