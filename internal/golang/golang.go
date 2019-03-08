package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(distPath string, msgs parser.Messages) {

	templateToFile(distPath, msgs, "messages_test.tpl", "messages_test.go")
	templateToFile(distPath, msgs, "messages.tpl", "messages.go")

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/golang/templates/" + tplFile

	path := distPath + "/golang/protocol/" + goFile

	parser.TemplateToFile(distPath, data, tpl, path)
}
