
import actions_pb2

{{ range .Messages }}
# Code{{.Name}} identifies a payload as a {{.Name}} action message.
Code{{.Name}} = b'{{.Code}}'
{{ end }}

def getActionObject(actionCode):
{{- range .Messages }}
	if actionCode == Code{{.Name}}:
		return actions_pb2.{{.Name}}()
{{- end }}

	raise Exception("Unsupported action code: ", str(actionCode))

def parseActionObject(actionCode, actionData):
	action = getActionObject(actionCode)
	action.ParseFromString(actionData)
	return action
