package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/internal"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/pkg/errors"
)

const (
	// AmendmentOperationModify specifies an amendment is modifying a value.
	AmendmentOperationModify = uint32(0)

	// AmendmentOperationAddElement specifies an amendment is adding a new element to a list.
	AmendmentOperationAddElement = uint32(1)

	// AmendmentOperationRemoveElement specifies an amendment is removing an element from a list.
	AmendmentOperationRemoveElement = uint32(2)
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
func (a *{{ $message.Name }}) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

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
func (a *{{ $message.Name }}) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *{{ $message.Name }}) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

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
func (a *{{ $message.Name }}Field) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

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
func (a *{{ $message.Name }}Field) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *{{ $message.Name }}Field) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

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
func CreatePayloadAmendments(fip permissions.FieldIndexPath,
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
