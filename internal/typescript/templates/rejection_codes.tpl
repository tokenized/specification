
export enum RejectCodes {
{{- range .Values}}
	// {{.Name}} - {{.Description}}
	Reject{{.Name}} = {{ .Code }},
{{ end }}
}
