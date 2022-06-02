package {{ .Package }}


{{ range .Messages }}
type {{.Name}} struct {
	{{- $fields := .Fields -}}

	{{- $longestName := 8 }}
	{{- $longestType := 0 }}
	{{- range $i, $a := $fields }}
		{{- if gt (len .Name) $longestName }}
			{{- $longestName = len .Name }}
		{{- end }}
		{{- if gt (len .GoTypeWithPointer) $longestType }}
			{{- $longestType = len .GoTypeWithPointer }}
		{{- end }}
	{{- end }}
	{{- $longestName = add $longestName 1 }}
	{{- $longestType = add $longestType 1 }}

	{{- range $i, $a := $fields }}
		{{- if ne .Type "deprecated" }}
	{{ .Name }} {{ padding (print .Name) $longestName}}
			{{- if and (eq .BaseType "fixedchar") (ne .BaseSize 0) }}
	{{- .GoTypeWithPointer }} {{- padding  (print .GoTypeWithPointer) $longestType }} `bsor:"{{add $i 1}}" bsor_fixed_size:"{{ .BaseSize }}" json:"{{ .Name }}"`
			{{- else }}
	{{- .GoTypeWithPointer }} {{- padding  (print .GoTypeWithPointer) $longestType }} `bsor:"{{add $i 1}}" json:"{{ .Name }}"`
			{{- end }}
		{{- end }}
	{{- end }}
}
{{ end }}


{{ range .FieldTypes }}
type {{.Name}}Field struct {
	{{- $fields := .Fields }}

	{{- $longestName := 8 }}
	{{- $longestType := 0 }}
	{{- range $i, $a := $fields }}
		{{- if gt (len .Name) $longestName }}
			{{- $longestName = len .Name }}
		{{- end }}
		{{- if gt (len .GoTypeWithPointer) $longestType }}
			{{- $longestType = len .GoTypeWithPointer }}
		{{- end }}
	{{- end }}
	{{- $longestName = add $longestName 1 }}
	{{- $longestType = add $longestType 1 }}

	{{- range $i, $a := $fields }}
		{{- if ne .Type "deprecated" }}
	{{ .Name }} {{ padding (print .Name) $longestName}}
			{{- if and (eq .BaseType "fixedchar") (ne .BaseSize 0) }}
	{{- .GoTypeWithPointer }} {{- padding  (print .GoTypeWithPointer) $longestType }} `bsor:"{{add $i 1}}" bsor_fixed_size:"{{ .BaseSize }}" json:"{{ .Name }}"`
			{{- else }}
	{{- .GoTypeWithPointer }} {{- padding  (print .GoTypeWithPointer) $longestType }} `bsor:"{{add $i 1}}" json:"{{ .Name }}"`
			{{- end }}
		{{- end }}
	{{- end }}
}
{{ end }}
