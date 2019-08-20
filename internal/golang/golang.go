package golang

import (
	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	assets parser.Schema,
	messages parser.Schema,
) {
	// Actions
	parser.TemplateToFile(actions, "internal/golang/templates/actions.tpl",
		distPath+"/golang/"+"actions/actions.go")
	parser.TemplateToFile(actions, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"actions/validate.go")
	parser.TemplateToFile(actions, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"actions/equal.go")
	parser.TemplateToFile(actions, "internal/golang/templates/actions_test.tpl",
		distPath+"/golang/"+"actions/actions_test.go")

	// Rejection Codes
	var rejections *parser.Resource
	for i, resource := range actions.Resources {
		if resource.Name == "Rejections" {
			rejections = &actions.Resources[i]
			break
		}
	}
	if rejections == nil {
		panic("Couldn't find actions rejections resource")
	}
	parser.TemplateToFile(rejections, "internal/golang/templates/rejections.tpl",
		distPath+"/golang/"+"actions/rejections.go")

	// Assets
	parser.TemplateToFile(assets, "internal/golang/templates/assets.tpl",
		distPath+"/golang/"+"assets/assets.go")
	parser.TemplateToFile(assets, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"assets/validate.go")
	parser.TemplateToFile(assets, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"assets/equal.go")
	parser.TemplateToFile(assets, "internal/golang/templates/assets_test.tpl",
		distPath+"/golang/"+"assets/assets_test.go")

	// Messages
	parser.TemplateToFile(messages, "internal/golang/templates/messages.tpl",
		distPath+"/golang/"+"messages/messages.go")
	parser.TemplateToFile(messages, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"messages/validate.go")
	parser.TemplateToFile(messages, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"messages/equal.go")
	parser.TemplateToFile(messages, "internal/golang/templates/messages_test.tpl",
		distPath+"/golang/"+"messages/messages_test.go")
}
