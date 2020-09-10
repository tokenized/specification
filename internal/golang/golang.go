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
	parser.TemplateToFile(actions,
		"internal/golang/templates/apply_amendments.tpl",  // Pull in ApplyAmendmentField
		"internal/golang/templates/create_amendments.tpl", // Pull in CreateAmendmentField
		"internal/golang/templates/action_amendments.tpl",
		distPath+"/golang/"+"actions/amendments.go")
	parser.TemplateToFile(actions, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"actions/validate.go")
	parser.TemplateToFile(actions, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"actions/equal.go")
	parser.TemplateToFile(actions, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"actions/resources.go")
	parser.TemplateToFile(actions, "internal/golang/templates/actions_test.tpl",
		distPath+"/golang/"+"actions/actions_test.go")

	// Assets
	parser.TemplateToFile(assets, "internal/golang/templates/assets.tpl",
		distPath+"/golang/"+"assets/assets.go")
	parser.TemplateToFile(assets,
		"internal/golang/templates/apply_amendments.tpl",  // Pull in ApplyAmendmentField
		"internal/golang/templates/create_amendments.tpl", // Pull in CreateAmendmentField
		"internal/golang/templates/asset_amendments.tpl",
		distPath+"/golang/"+"assets/amendments.go")
	parser.TemplateToFile(assets, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"assets/validate.go")
	parser.TemplateToFile(assets, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"assets/equal.go")
	parser.TemplateToFile(assets, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"assets/resources.go")
	parser.TemplateToFile(assets, "internal/golang/templates/assets_test.tpl",
		distPath+"/golang/"+"assets/assets_test.go")

	// Messages
	parser.TemplateToFile(messages, "internal/golang/templates/messages.tpl",
		distPath+"/golang/"+"messages/messages.go")
	parser.TemplateToFile(messages, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"messages/validate.go")
	parser.TemplateToFile(messages, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"messages/equal.go")
	parser.TemplateToFile(messages, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"messages/resources.go")
	parser.TemplateToFile(messages, "internal/golang/templates/messages_test.tpl",
		distPath+"/golang/"+"messages/messages_test.go")

	// Templates
	parser.ProcessContractPermissionConfigs(actions, assets, "src/templates/develop/",
		distPath+"/golang/"+"actions/permission_templates.go")
}
