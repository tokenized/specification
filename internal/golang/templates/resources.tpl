package {{ .Package }}


{{ range $i, $resource := .Resources }}
/************************************ {{ .Name }} ************************************/
const (
{{- range $j, $value := .Values}}
	// {{ $value.Name }} - {{ $value.Description }}
	{{ $resource.Name }}{{ stripwhitespace $value.Name }} = {{ if eq $resource.CodeType.BaseType "fixedchar" "varchar" }}"{{ $value.Code }}"{{ else }}{{ $value.Code }}{{ end }}
{{ end }}
)

type {{ .Name }}Code struct {
	Name string
	Label string
	Description string
}

// {{ $resource.Name }}Data holds a mapping of {{ $resource.Name }} codes.
func {{ $resource.Name }}Data(code {{ $resource.CodeType.GoType }}) *{{ $resource.Name }}Code {
	switch(code) {
{{- range $j, $value := $resource.Values }}
	case {{ $resource.Name }}{{ stripwhitespace $value.Name }}:
		return &{{ $resource.Name }}Code{
			Name: "{{ $value.Name }}",
			Label: "{{ $value.Label }}",
			Description: "{{ $value.Description }}",
		}
{{- end }}
	default:
		return nil
	}
}

// {{ $resource.Name }}Set returns a mapping of {{ $resource.Name }} codes.
func {{ $resource.Name }}Set() map[{{ $resource.CodeType.GoType }}]*{{ $resource.Name }}Code {
	return map[{{ $resource.CodeType.GoType }}]*{{ $resource.Name }}Code {
{{- range $j, $value := $resource.Values }}
		{{ $resource.Name }}{{ stripwhitespace $value.Name }} :
			&{{ $resource.Name }}Code{
				Name: "{{ $value.Name }}",
				Label: "{{ $value.Label }}",
				Description: "{{ $value.Description }}",
			},
{{- end }}
	}
}
{{ end }}
