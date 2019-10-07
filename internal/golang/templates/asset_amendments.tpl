package {{ .Package }}

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

{{- range $i, $message := .Messages }}
// {{ $message.Name }} Permission / Amendment Field Indices
const (
{{- range $offset, $field := .Fields }}
    {{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
    switch fip[0] {
    {{- range $offset, $field := .Fields }}
    case {{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
    {{- template "ApplyAmendmentField" $field }}
    {{ end }}
    }

    return nil
}

{{ end }}


{{ range $i, $message := .FieldTypes }}
// {{ $message.Name }}Field Permission / Amendment Field Indices
const (
{{- range $offset, $field := .Fields }}
    {{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }}Field based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}Field) ApplyAmendment(fip []uint32, operation uint32, data []byte) error {
    switch fip[0] {
    {{- range $offset, $field := .Fields }}
    case {{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
    {{ template "ApplyAmendmentField" $field }}
    {{- end }}
    }

    return nil
}

{{ end }}
