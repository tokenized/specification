package swagger

import (
	"fmt"

	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
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

	// Instruments
	allMessages = instruments.Messages
	for _, message := range allMessages {
		instruments.Messages = []parser.Message{message} // Each file only gets one message
		parser.TemplateToFile(instruments, "internal/swagger/templates/message.tpl",
			fmt.Sprintf("%s/swagger/instruments/%s.yaml", distPath, message.Name))
	}
	instruments.Messages = allMessages

	allFields = instruments.FieldTypes
	for _, fieldType := range allFields {
		instruments.FieldTypes = []parser.FieldType{fieldType} // Each file only gets one message
		parser.TemplateToFile(instruments, "internal/swagger/templates/field_type.tpl",
			fmt.Sprintf("%s/swagger/instruments/%s.yaml", distPath, fieldType.Name+"Field"))
	}
	instruments.FieldTypes = allFields

	parser.TemplateToFile(instruments, "internal/swagger/templates/components.tpl",
		fmt.Sprintf("%s/swagger/instruments/components.yaml", distPath))

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
