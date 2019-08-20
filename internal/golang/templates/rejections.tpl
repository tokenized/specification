package actions

const (
{{- range .Values}}
	// {{.Name}} - {{.Description}}
	Reject{{.Name}} = {{ .Code }}
{{ end }}
)

type RejectionCode struct {
	Name string
	Label string
	Description string
}

// TypeMapping holds a mapping of action codes to action types.
func RejectionData(code uint32) *RejectionCode {
	switch(code) {
{{- range .Values }}
	case Reject{{.Name}}:
		return &RejectionCode{
			Name: "{{.Name}}",
			Label: "{{.Label}}",
			Description: "{{.Description}}",
		}
{{- end }}
	default:
		return nil
	}
}
