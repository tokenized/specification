{{$messages := .Messages -}}
{{$fieldTypes := .FieldTypes -}}

syntax = "proto3";

package protoc;
{{ range $messages }}
// Action - {{.Label}} ({{.Code}})
message {{.Name}}Action {
{{- range $i, $a := .Fields}}
    {{ .ProtobufType }} {{.Name}}
    {{- padding (print .Name .ProtobufType) 45}} = {{$i}};
    {{- padding (print $i) 3}} // {{.Type}}
{{- end}}
}
{{ end }}
{{range $i, $a := $fieldTypes}}
// Field - {{.Label}}
message {{.Name}}Field {
{{- range $i, $a := .Fields}}
    {{ .ProtobufType }} {{.Name}}
    {{- padding (print .Name .ProtobufType) 45}} = {{$i}};
    {{- padding (print $i) 3}} // {{.Type}}
{{- end}}
}
{{ end }}
