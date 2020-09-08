package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/specification/dist/golang/internal"

	proto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

{{- range $i, $message := .Messages }}
	{{- if eq $message.Name "ContractOffer" }}
// Contract Permission / Amendment Field Indices
const (
		{{- range $i, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	DeprecatedContractField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else if eq $field.Name "ContractOperatorIncluded" "MasterAddress" }}
	NotAmendableContractField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else }}
	ContractField{{ $field.Name }} = uint32({{add $i 1}})
			{{- end }}
		{{- end }}
)

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data.
func (a *{{ $message.Name }}) CreateAmendments(newValue *{{ $message.Name }}) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	var result []*internal.Amendment
	var fip []uint32

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else if not (eq $field.Name "ContractOperatorIncluded" "MasterAddress") }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = []uint32{ContractField{{ $field.Name }}}
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

	{{- else if eq $message.Name "ContractFormation" }}

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty contract amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedContractField{{ $field.Name }}: // {{ $field.GoType }}
			{{- else if eq $field.Name "ContractOperatorIncluded" "MasterAddress" }}
	case NotAmendableContractField{{ $field.Name }}: // {{ $field.GoType }}
		return nil, fmt.Errorf("{{ $field.Name }} field not amendable")
			{{- else if not (eq $field.Name "ContractRevision" "Timestamp" "AdminAddress" "OperatorAddress") }}
	case ContractField{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown contract amendment field index : %v", fip)
}

	{{- else if eq $message.Name "AssetDefinition" }}
// Asset Permission / Amendment Field Indices
const (
		{{- range $i, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	DeprecatedAssetField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else }}
	AssetField{{ $field.Name }} = uint32({{add $i 1}})
			{{- end }}
		{{- end }}
)

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data.
func (a *{{ $message.Name }}) CreateAmendments(newValue *{{ $message.Name }}) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	var fip []uint32

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = []uint32{AssetField{{ $field.Name }}}
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

	{{- else if eq $message.Name "AssetCreation" }}
// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedAssetField{{ $field.Name }}: // {{ $field.GoType }}
		{{- else if not (eq $field.Name "AssetCode" "AssetIndex" "AssetRevision" "Timestamp") }}
	case AssetField{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown asset amendment field index : %v", fip)
}
	{{- end }}
{{ end }}

{{ range $i, $message := .FieldTypes }}
// {{ $message.Name }}Field Permission / Amendment Field Indices
const (
	{{- range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	Deprecated{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
		{{- else }}
	{{ $message.Name }}Field{{ $field.Name }} = uint32({{add $offset 1}})
		{{- end }}
	{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }}Field based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}Field) ApplyAmendment(fip FieldIndexPath, operation uint32,
	data []byte) (FieldIndexPath, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty {{ $message.Name }} amendment field index path")
	}

	switch fip[0] {
	{{- range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	case Deprecated{{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
		{{- else }}
	case {{ $message.Name }}Field{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
		{{- end }}
	{{ end }}
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


func convertAmendments(amendments []*internal.Amendment) ([]*AmendmentField, error) {
	var result []*AmendmentField
	for _, am := range amendments {
		b, err := FieldIndexPath(am.FIP).Bytes()
		if err != nil {
			return nil, errors.Wrap(err, "fip")
		}

		result = append(result, &AmendmentField{
			FieldIndexPath: b,
			Operation: am.Operation,
			Data: am.Data,
		})
	}

	return result, nil
}
