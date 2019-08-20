{{$messages := .Messages -}}
{{$fieldTypes := .FieldTypes -}}

syntax = "proto3";

package {{.Package}};
{{ range $messages }}
// Message - {{.Label}} ({{.Code}})
message {{.Name}} {
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
