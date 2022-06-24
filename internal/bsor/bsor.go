package bsor

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
	messages parser.Schema,
) {
	parser.TemplateToFile(actions, "internal/bsor/templates/golang.tpl",
		distPath+"/bsor/golang/"+"actions.go")
	parser.TemplateToFile(instruments, "internal/bsor/templates/golang.tpl",
		distPath+"/bsor/golang/"+"instruments.go")
	parser.TemplateToFile(messages, "internal/bsor/templates/golang.tpl",
		distPath+"/bsor/golang/"+"messages.go")
}
