package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	proto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

{{- range $i, $message := .Messages }}
	{{- if eq $message.Name "ContractFormation" }}
// Contract Permission / Amendment Field Indices
const (
		{{- range $i, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	DeprecatedContractField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else }}
	ContractField{{ $field.Name }} = uint32({{add $i 1}})
			{{- end }}
		{{- end }}
)

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty contract amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedContractField{{ $field.Name }}: // {{ $field.GoType }}
			{{- else }}
	case ContractField{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown contract amendment field index : %v", fip)
}
	{{- else if eq $message.Name "AssetCreation" }}
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

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedAssetField{{ $field.Name }}: // {{ $field.GoType }}
			{{- else }}
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
func (a *{{ $message.Name }}Field) ApplyAmendment(fip FieldIndexPath, operation uint32, data []byte) (FieldIndexPath, error) {
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

{{ end }}
