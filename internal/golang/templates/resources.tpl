package {{ .Package }}


{{ range $i, $resource := .Resources }}
/************************************ {{ .Name }} ************************************/
const (
{{- range $j, $value := .Values}}
	{{ $name := $value.Name }}
	{{ if eq (len $name) 0 -}}
	{{ $name = stripwhitespace $value.Label -}}
	{{ end -}}
	// {{ $name }} - {{ $value.Description }}
	{{ $resource.Name }}{{ stripwhitespace $name }} = {{ if eq $resource.CodeType.BaseType "fixedchar" "varchar" }}"{{ $value.Code }}"{{ else }}{{ $value.Code }}{{ end }}
{{ end }}
)

type {{ .Name }}Code struct {
	Name string
	Label string
	Description string
	MetaData string
}

// {{ $resource.Name }}Data holds a mapping of {{ $resource.Name }} codes.
func {{ $resource.Name }}Data(code {{ $resource.CodeType.GoType }}) *{{ $resource.Name }}Code {
	switch code {
{{- range $j, $value := $resource.Values }}
	{{ $name := $value.Name }}
	{{ if eq (len $name) 0 -}}
	{{ $name = stripwhitespace $value.Label -}}
	{{ end -}}
	case {{ $resource.Name }}{{ stripwhitespace $name }}:
		return &{{ $resource.Name }}Code{
			Name: "{{ $name }}",
			Label: "{{ $value.Label }}",
			Description: "{{ $value.Description }}",
			MetaData: `{{ $value.MetaDataJSON }}`,
		}
{{- end }}
	default:
		return nil
	}
}

// {{ $resource.Name }}Map returns a mapping of {{ $resource.Name }} objects with the code as the key.
func {{ $resource.Name }}Map() map[{{ $resource.CodeType.GoType }}]*{{ $resource.Name }}Code {
	return map[{{ $resource.CodeType.GoType }}]*{{ $resource.Name }}Code {
{{- range $j, $value := $resource.Values }}
		{{ $name := $value.Name }}
		{{ if eq (len $name) 0 -}}
		{{ $name = stripwhitespace $value.Label -}}
		{{ end -}}
		{{ $resource.Name }}{{ stripwhitespace $name }} :
			&{{ $resource.Name }}Code{
				Name: "{{ $name }}",
				Label: "{{ $value.Label }}",
				Description: "{{ $value.Description }}",
				MetaData: `{{ $value.MetaDataJSON }}`,
			},
{{- end }}
	}
}
{{ end }}
