package {{ .Package }}

import (
	"bytes"
	"fmt"

	"github.com/tokenized/specification/dist/golang/internal"

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

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data.
func (a *{{ $message.Name }}) CreateAmendments(fip []uint32,
	newValue *{{ $message.Name }}) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = append(ofip, {{ $message.Name }}Field{{ $field.Name }})
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	return result, nil
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

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data.
func (a *{{ $message.Name }}Field) CreateAmendments(fip []uint32,
	newValue *{{ $message.Name }}Field) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip // save original to be appended for each field

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = append(ofip, {{ $message.Name }}Field{{ $field.Name }})
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	return result, nil
}

{{ end }}

// CreatePayloadAmendments deserializes asset payloads and create amendments for them.
func CreatePayloadAmendments(fip []uint32,
	assetType, payload, newPayload []byte) ([]*internal.Amendment, error) {

	current, err := Deserialize(assetType, payload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	new, err := Deserialize(assetType, newPayload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	var result []*internal.Amendment
	switch c := current.(type) {
{{- range $i, $message := .Messages }}
	case *{{ $message.Name }}:
		result, err = c.CreateAmendments(fip, new.(*{{ $message.Name }}))
{{ end }}
	}

	return result, nil
}
