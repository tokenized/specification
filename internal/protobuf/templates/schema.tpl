{{$messages := .Messages -}}
{{$fieldTypes := .FieldTypes -}}

syntax = "proto3";

package {{.Package}};
{{ range $messages }}
// Action - {{.Label}} ({{.Code}})
message {{.Name}}Action {
{{- range $i, $a := .Fields}}
    {{ .ProtobufType }} {{.Name}}
    {{- padding (print .Name .ProtobufType) 45}} = {{add $i 1}};
    {{- padding (print $i) 3}} // {{.Type}}
{{- end}}
}
{{ end }}
{{range $i, $a := $fieldTypes}}
// Field - {{.Label}}
message {{.Name}}Field {
{{- range $i, $a := .Fields}}
    {{ .ProtobufType }} {{.Name}}
    {{- padding (print .Name .ProtobufType) 45}} = {{add $i 1}};
    {{- padding (print $i) 3}} // {{.Type}}
{{- end}}
}
{{ end }}
