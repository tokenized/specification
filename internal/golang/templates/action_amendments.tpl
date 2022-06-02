package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/instruments"
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
// amendment data. Use the current value of contract formation, and pass in the new values as a
// contract offer.
func (a *ContractFormation) CreateAmendments(newValue *{{ $message.Name }}) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

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
func (a *{{ $message.Name }}) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

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
				{{- if .IsCompoundType }}
		if len(fip) == 1 && len(data) == 0 {
			a.{{ $field.Name }} = nil
			return permissions.SubPermissions(fip[1:], operation, {{ $field.IsList }})
		}
				{{- end }}

	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown contract amendment field index : %v", fip)
}

	{{- else if eq $message.Name "BodyOfAgreementOffer" }}
// BodyOfAgreement Permission / Modification Field Indices
const (
		{{- range $i, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	DeprecatedBodyOfAgreementField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else }}
	BodyOfAgreementField{{ $field.Name }} = uint32({{add $i 1}})
			{{- end }}
		{{- end }}
)

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data. Use the current value of agreement, and pass in the new values as a
// agreement definition.
func (a *BodyOfAgreementFormation) CreateAmendments(newValue *{{ $message.Name }}) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = []uint32{BodyOfAgreementField{{ $field.Name }}}
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

	{{- else if eq $message.Name "BodyOfAgreementFormation" }}

// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty agreement amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedBodyOfAgreementField{{ $field.Name }}: // {{ $field.GoType }}
			{{- else if not (eq $field.Name "Revision" "Timestamp") }}
	case BodyOfAgreementField{{ $field.Name }}: // {{ $field.GoType }}
	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown agreement amendment field index : %v", fip)
}

	{{- else if eq $message.Name "InstrumentDefinition" }}
// Instrument Permission / Amendment Field Indices
const (
		{{- range $i, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	DeprecatedInstrumentField{{ $field.Name }} = uint32({{add $i 1}})
			{{- else }}
	InstrumentField{{ $field.Name }} = uint32({{add $i 1}})
			{{- end }}
		{{- end }}
)

// CreateAmendments determines the differences between two {{ $message.Name }}s and returns
// amendment data. Use the current value of instrument creation, and pass in the new values as an instrument
// definition.
func (a *InstrumentCreation) CreateAmendments(newValue *{{ $message.Name }}) ([]*AmendmentField, error) {
	if err := newValue.Validate(); err != nil {
		return nil, errors.Wrap(err, "new value invalid")
	}

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	var fip permissions.FieldIndexPath

	{{ range $offset, $field := .Fields }}
		{{- if eq $field.Type "deprecated" }}
	// deprecated {{ $field.Name }} {{ $field.GoType }}
		{{- else }}
	// {{ $field.Name }} {{ $field.GoType }}
	fip = []uint32{InstrumentField{{ $field.Name }}}
	{{- template "CreateAmendmentField" $field }}
		{{- end }}
	{{ end }}

	r, err := convertAmendments(result)
	if err != nil {
		return nil, errors.Wrap(err, "convert amendments")
	}

	return r, nil
}

	{{- else if eq $message.Name "InstrumentCreation" }}
// ApplyAmendment updates a {{ $message.Name }} based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *{{ $message.Name }}) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty instrument amendment field index path")
	}

	switch fip[0] {
		{{- range $offset, $field := .Fields }}
			{{- if eq $field.Type "deprecated" }}
	case DeprecatedInstrumentField{{ $field.Name }}: // {{ $field.GoType }}
		{{- else if not (eq $field.Name "InstrumentCode" "InstrumentIndex" "InstrumentRevision" "Timestamp") }}
	case InstrumentField{{ $field.Name }}: // {{ $field.GoType }}
				{{- if .IsCompoundType }}
		if len(fip) == 1 && len(data) == 0 {
			a.{{ $field.Name }} = nil
			return permissions.SubPermissions(fip[1:], operation, {{ $field.IsList }})
		}
				{{- end }}

	{{- template "ApplyAmendmentField" $field }}
			{{- end }}
		{{ end }}
	}

	return nil, fmt.Errorf("Unknown instrument amendment field index : %v", fip)
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
func (a *{{ $message.Name }}Field) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

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
func (a *{{ $message.Name }}Field) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *{{ $message.Name }}Field) ([]*internal.Amendment, error) {

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	if a != nil && newValue == nil {
		result = append(result, &internal.Amendment{
			FIP: fip,
			Operation: 0, // Modify element
			Data: nil,
		})

		return result, nil
	}

	if a.Equal(newValue) {
		return nil, nil
	}

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
		b, err := permissions.FieldIndexPath(am.FIP).Bytes()
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
