package protocol

const (
{{- range .Values}}
	// {{.Name}} - {{.Description}}
	Reject{{.Name}} = {{ .Code }}
{{ end }}
)