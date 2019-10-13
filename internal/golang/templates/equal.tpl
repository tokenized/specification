package {{ .Package }}

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

{{ define "EqualField" -}}
    // Field {{ .Name }} - {{ .BaseType }}
    {{- if .IsList }}
        if len(l.{{ .Name }}) != len(r.{{ .Name }}) {
            return false // fmt.Errorf("List length mismatched")
        }
        for i, v := range l.{{ .Name }} {
        {{- if eq .BaseType "bin" "varbin" }}
            if !bytes.Equal(v, r.{{ .Name }}[i]) {
                return false // fmt.Errorf("Element {{ .Name }} bytes mismatched")
            }
        {{- else if eq .BaseType "fixedchar" "varchar" }}
            if v != r.{{ .Name }}[i] {
                return false // fmt.Errorf("Element {{ .Name }} string mismatched")
            }
        {{- else if eq .BaseType "uint" }}
            if v != r.{{ .Name }}[i] {
                return false // fmt.Errorf("Element {{ .Name }} integer mismatched")
            }
        {{- else }}
            if !v.Equal(r.{{ .Name }}[i]) {
                return false // fmt.Errorf("{{ .Name }}[%d] : %s", i, err)
            }
        {{- end }}
        }
    {{- else if eq .BaseType "bin" "varbin" }}
		if !bytes.Equal(l.{{ .Name }}, r.{{ .Name }}) {
			return false // fmt.Errorf("{{ .Name }} bytes mismatched")
		}
    {{- else if eq .BaseType "fixedchar" "varchar" }}
		if l.{{ .Name }} != r.{{ .Name }} {
			return false // fmt.Errorf("{{ .Name }} string mismatched")
		}
    {{- else if eq .BaseType "uint" }}
	if l.{{ .Name }} != r.{{ .Name }} {
		return false // fmt.Errorf("{{ .Name }} integer mismatched")
	}
    {{- else if eq .BaseType "bool" }}
	if l.{{ .Name }} != r.{{ .Name }} {
		return false // fmt.Errorf("{{ .Name }} boolean mismatched")
	}
    {{- else }}
        if !l.Equal(r.{{ .Name }}) {
            return false // fmt.Errorf("{{ .Name }} : %s", err)
        }
    {{- end }}
{{ end }}

{{ range .Messages }}
func (l *{{.Name}}) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*{{.Name}})
	if !ok {
		return false
	}

	{{ range .Fields }}
		{{ template "EqualField" . -}}
	{{- end }}

	return true
}
{{ end }}

{{ range .FieldTypes }}
func (l *{{.Name}}Field) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*{{.Name}}Field)
	if !ok {
		return false
	}

	{{ range .Fields }}
		{{ template "EqualField" . -}}
	{{- end }}

	return true
}
{{ end }}
