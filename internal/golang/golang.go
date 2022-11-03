package golang

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
	parser.TemplateToFile(actions, "internal/golang/templates/binary_marshal.tpl",
		distPath+"/golang/"+"actions/binary_marshal.go")
	parser.TemplateToFile(actions, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"actions/resources.go")
	parser.TemplateToFile(actions, "internal/golang/templates/actions_test.tpl",
		distPath+"/golang/"+"actions/actions_test.go")
	parser.TemplateToFile(actions, "internal/golang/templates/write_deterministic.tpl",
		distPath+"/golang/"+"actions/write_deterministic.go")

	// Instruments
	parser.TemplateToFile(instruments, "internal/golang/templates/instruments.tpl",
		distPath+"/golang/"+"instruments/instruments.go")
	parser.TemplateToFile(instruments,
		"internal/golang/templates/apply_amendments.tpl",  // Pull in ApplyAmendmentField
		"internal/golang/templates/create_amendments.tpl", // Pull in CreateAmendmentField
		"internal/golang/templates/instrument_amendments.tpl",
		distPath+"/golang/"+"instruments/amendments.go")
	parser.TemplateToFile(instruments, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"instruments/validate.go")
	parser.TemplateToFile(instruments, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"instruments/equal.go")
	parser.TemplateToFile(instruments, "internal/golang/templates/binary_marshal.tpl",
		distPath+"/golang/"+"instruments/binary_marshal.go")
	parser.TemplateToFile(instruments, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"instruments/resources.go")
	parser.TemplateToFile(instruments, "internal/golang/templates/instruments_test.tpl",
		distPath+"/golang/"+"instruments/instruments_test.go")

	// Messages
	parser.TemplateToFile(messages, "internal/golang/templates/messages.tpl",
		distPath+"/golang/"+"messages/messages.go")
	parser.TemplateToFile(messages, "internal/golang/templates/validate.tpl",
		distPath+"/golang/"+"messages/validate.go")
	parser.TemplateToFile(messages, "internal/golang/templates/equal.tpl",
		distPath+"/golang/"+"messages/equal.go")
	parser.TemplateToFile(messages, "internal/golang/templates/binary_marshal.tpl",
		distPath+"/golang/"+"messages/binary_marshal.go")
	parser.TemplateToFile(messages, "internal/golang/templates/resources.tpl",
		distPath+"/golang/"+"messages/resources.go")
	parser.TemplateToFile(messages, "internal/golang/templates/messages_test.tpl",
		distPath+"/golang/"+"messages/messages_test.go")

	// Templates
	parser.ProcessContractPermissionConfigs(actions, instruments, "src/templates/develop/",
		distPath+"/golang/"+"actions/permission_templates.go")
}
