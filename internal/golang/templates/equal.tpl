package {{ .Package }}

import (
	"bytes"
)

{{ define "EqualField" -}}
    // Field {{ .Name }} - {{ .BaseType }}
    {{- if .IsList }}
        if len(c.{{ .Name }}) != len(r.{{ .Name }}) {
            return false // fmt.Errorf("List length mismatched")
        }
        for i, v := range c.{{ .Name }} {
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
		if !bytes.Equal(c.{{ .Name }}, r.{{ .Name }}) {
			return false // fmt.Errorf("{{ .Name }} bytes mismatched")
		}
    {{- else if eq .BaseType "fixedchar" "varchar" }}
		if c.{{ .Name }} != r.{{ .Name }} {
			return false // fmt.Errorf("{{ .Name }} string mismatched")
		}
    {{- else if eq .BaseType "uint" }}
	if c.{{ .Name }} != r.{{ .Name }} {
		return false // fmt.Errorf("{{ .Name }} integer mismatched")
	}
    {{- else if eq .BaseType "bool" }}
	if c.{{ .Name }} != r.{{ .Name }} {
		return false // fmt.Errorf("{{ .Name }} boolean mismatched")
	}
    {{- else }}
        if !c.{{ .Name }}.Equal(r.{{ .Name }}) {
            return false // fmt.Errorf("{{ .Name }} : %s", err)
        }
    {{- end }}
{{ end }}

{{ range .Messages }}
func (l *{{.Name}}) Equal(right interface{}) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &{{.Name}}{}
	}
	cr := right
	if cr == nil {
		cr = &{{.Name}}{}
	}
	r, ok := cr.(*{{.Name}})
	if !ok {
		return false
	}

	{{ range .Fields }}
		{{- if ne .Type "deprecated" }}
		{{ template "EqualField" . -}}
		{{- end }}
	{{- end }}

	return true
}
{{ end }}

{{ range .FieldTypes }}
func (l *{{.Name}}Field) Equal(right interface{}) bool {
	c := l
	if c == nil {
		if right == nil {
			return true
		}
		c = &{{.Name}}Field{}
	}
	r, ok := right.(*{{.Name}}Field)
	if !ok {
		return false
	}

	if r == nil {
		r = &{{.Name}}Field{}
	}

	{{ range .Fields }}
		{{- if ne .Type "deprecated" }}
		{{ template "EqualField" . -}}
		{{- end }}
	{{- end }}

	return true
}
{{ end }}
