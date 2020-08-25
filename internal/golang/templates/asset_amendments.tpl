package {{ .Package }}

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

{{- range $i, $message := .Messages }}
// {{ $message.Name }} Permission / Amendment Field Indices
const (
{{- range $offset, $field := .Fields }}
	{{- if eq .Type "deprecated" }}
	Deprecated{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
	{{- else }}
	{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
	{{- end}}
{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	{{- range $offset, $field := .Fields }}
		{{- if eq .Type "deprecated" }}
	case Deprecated{{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
		{{- else }}
	case {{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
		{{- end}}
	{{ end }}
	}

	return nil, fmt.Errorf("Unknown {{ $message.Name }} amendment field index : %v", fip)
}

{{ end }}


{{ range $i, $message := .FieldTypes }}
// {{ $message.Name }}Field Permission / Amendment Field Indices
const (
{{- range $offset, $field := .Fields }}
	{{- if eq .Type "deprecated" }}
	Deprecated{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
	{{- else }}
	{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
	{{- end}}
{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }}Field based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}Field) ApplyAmendment(fip []uint32, operation uint32, data []byte) ([]uint32, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty {{ $message.Name }} amendment field index path")
	}

	switch fip[0] {
	{{- range $offset, $field := .Fields }}
		{{- if eq .Type "deprecated" }}
	case Deprecated{{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
		{{- else }}
	case {{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
	{{ template "ApplyAmendmentField" $field }}
		{{- end}}
	{{- end }}
	}

	return nil, fmt.Errorf("Unknown {{ $message.Name }} amendment field index : %v", fip)
}

{{ end }}
