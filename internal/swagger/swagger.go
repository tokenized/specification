package swagger

import (
	"fmt"

	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	assets parser.Schema,
	messages parser.Schema,
) {
	// Actions
	allMessages := actions.Messages
	for _, message := range allMessages {
		actions.Messages = []parser.Message{message} // Each file only gets one message
		parser.TemplateToFile(actions, "internal/swagger/templates/message.tpl",
			fmt.Sprintf("%s/swagger/actions/%s.yaml", distPath, message.Name))
	}
	actions.Messages = allMessages

	allFields := actions.FieldTypes
	for _, fieldType := range allFields {
		actions.FieldTypes = []parser.FieldType{fieldType} // Each file only gets one message
		parser.TemplateToFile(actions, "internal/swagger/templates/field_type.tpl",
			fmt.Sprintf("%s/swagger/actions/%s.yaml", distPath, fieldType.Name+"Field"))
	}
	actions.FieldTypes = allFields

	parser.TemplateToFile(actions, "internal/swagger/templates/components.tpl",
		fmt.Sprintf("%s/swagger/actions/components.yaml", distPath))

	// Assets
	allMessages = assets.Messages
	for _, message := range allMessages {
		assets.Messages = []parser.Message{message} // Each file only gets one message
		parser.TemplateToFile(assets, "internal/swagger/templates/message.tpl",
			fmt.Sprintf("%s/swagger/assets/%s.yaml", distPath, message.Name))
	}
	assets.Messages = allMessages

	allFields = assets.FieldTypes
	for _, fieldType := range allFields {
		assets.FieldTypes = []parser.FieldType{fieldType} // Each file only gets one message
		parser.TemplateToFile(assets, "internal/swagger/templates/field_type.tpl",
			fmt.Sprintf("%s/swagger/assets/%s.yaml", distPath, fieldType.Name+"Field"))
	}
	assets.FieldTypes = allFields

	parser.TemplateToFile(assets, "internal/swagger/templates/components.tpl",
		fmt.Sprintf("%s/swagger/assets/components.yaml", distPath))

	// Messages
	allMessages = messages.Messages
	for _, message := range allMessages {
		messages.Messages = []parser.Message{message} // Each file only gets one message
		parser.TemplateToFile(messages, "internal/swagger/templates/message.tpl",
			fmt.Sprintf("%s/swagger/messages/%s.yaml", distPath, message.Name))
	}
	messages.Messages = allMessages

	allFields = messages.FieldTypes
	for _, fieldType := range allFields {
		messages.FieldTypes = []parser.FieldType{fieldType} // Each file only gets one message
		parser.TemplateToFile(messages, "internal/swagger/templates/field_type.tpl",
			fmt.Sprintf("%s/swagger/messages/%s.yaml", distPath, fieldType.Name+"Field"))
	}
	messages.FieldTypes = allFields

	parser.TemplateToFile(messages, "internal/swagger/templates/components.tpl",
		fmt.Sprintf("%s/swagger/messages/components.yaml", distPath))
}
