{{$messages := .Messages -}}
{{$fieldTypes := .FieldTypes -}}

version 0

{{ range $messages }}
// {{.Label}} ({{.Code}})
{{.Name}} {

{{- $longestName := 8 }}
{{- $longestType := 0 }}
{{- range .Fields }}
    {{- if gt (len .Name) $longestName }}
        {{- $longestName = len .Name }}
    {{- end }}
    {{- if gt (len .BSORType) $longestType }}
        {{- $longestType = len .BSORType }}
    {{- end }}
{{- end }}
{{- $longestName = add $longestName 1 }}
{{- $longestType = add $longestType 1 }}

{{- range $i, $a := .Fields}}
    {{- if ne .Type "deprecated" }}
    {{.Name}} {{- padding (print .Name) $longestName}}
    {{- .BSORType }} {{- padding (print .BSORType) $longestType}}
    {{- "=" }} {{add $i 1}}; {{- padding (print (add $i 1)) 3}} // {{.Type}}
    {{- end }}
{{- end }}
}
{{ end }}

{{- range $i, $a := $fieldTypes}}
// {{.Label}} Field
{{.Name}}Field {

{{- $longestName := 8 }}
{{- $longestType := 0 }}
{{- range .Fields }}
    {{- if gt (len .Name) $longestName }}
        {{- $longestName = len .Name }}
    {{- end }}
    {{- if gt (len .BSORType) $longestType }}
        {{- $longestType = len .BSORType }}
    {{- end }}
{{- end }}
{{- $longestName = add $longestName 1 }}
{{- $longestType = add $longestType 1 }}

{{- range $i, $a := .Fields}}
    {{- if ne .Type "deprecated" }}
    {{.Name}} {{- padding (print .Name) $longestName}}
    {{- .BSORType }} {{- padding (print .BSORType) $longestType}}
    {{- "=" }} {{add $i 1}}; {{- padding (print (add $i 1)) 3}} // {{.Type}}
    {{- end }}
{{- end }}
}
{{ end }}
