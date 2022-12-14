package {{ .Package }}

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

{{ define "CopyField" -}}
	// Field {{ .Name }} - {{ .BaseType }}
	{{- if .IsList }}
		result.{{ .Name }} = make({{ .GoType }}, len(a.{{ .Name }}))
		for i, v := range a.{{ .Name }} {
		{{- if eq .BaseType "bin" "varbin" }}
			result.{{ .Name }}[i] = make({{ .GoSingularType }}, len(a.{{ .Name }}[i]))
			copy(result.{{ .Name }}[i], v)
		{{- else if eq .BaseType "fixedchar" "varchar" "uint" "bool" }}
			result.{{ .Name }}[i] = v
		{{- else }}
			result.{{ .Name }}[i] = v.Copy()
		{{- end }}
		}
	{{- else if eq .BaseType "bin" "varbin" }}
		result.{{ .Name }} = make({{ .GoType }}, len(a.{{ .Name }}))
		copy(result.{{ .Name }}, a.{{ .Name }})
	{{- else if eq .BaseType "fixedchar" "varchar" "uint" "bool" }}
		result.{{ .Name }} = a.{{ .Name }}
	{{- else }}
		result.{{ .Name }} = a.{{ .Name }}.Copy()
	{{- end }}
{{ end }}

{{ range .Messages }}
func (a *{{.Name}}) Copy() *{{.Name}} {
	if a == nil {
		return nil
	}

	result := &{{.Name}}{}

	{{ range .Fields }}
		{{- if ne .Type "deprecated" }}
		{{ template "CopyField" . -}}
		{{- end }}
	{{- end }}

	return result
}
{{ end }}

{{ range .FieldTypes }}
func (a *{{.Name}}Field) Copy() *{{.Name}}Field {
	if a == nil {
		return nil
	}

	result := &{{.Name}}Field{}

	{{ range .Fields }}
		{{- if ne .Type "deprecated" }}
		{{ template "CopyField" . -}}
		{{- end }}
	{{- end }}

	return result
}
{{ end }}
