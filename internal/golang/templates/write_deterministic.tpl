
{{- define "WriteDeterministicField" -}}
	{{- if .IsList }}
	for i, item := range a.{{ .Name }} {
		{{- if .IsCompoundType }}
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else if eq .BaseType "fixedchar" "varchar" }}
		if _, err := w.Write([]byte(item)); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else if eq .BaseType "bin" }}
		if _, err := w.Write(item[:]); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else if eq .BaseType "varbin" }}
		if _, err := w.Write(item); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else if eq .BaseType "uint" }}
		if err := WriteBase128VarInt64(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else if eq .BaseType "bool" }}
		if err := binary.Write(w, binary.LittleEndian, item); err != nil {
			return errors.Wrapf(err, "{{ .Name }} %d", i)
		}
		{{- else }}
		Unsupported Base Type : {{ .BaseType }}
		{{- end }}
	}
	{{- else if .IsCompoundType }}
	if err := a.{{ .Name }}.WriteDeterministic(w); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else if eq .BaseType "fixedchar" "varchar" }}
	if _, err := w.Write([]byte(a.{{ .Name }})); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else if eq .BaseType "bin" }}
	if _, err := w.Write(a.{{ .Name }}[:]); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else if eq .BaseType "varbin" }}
	if _, err := w.Write(a.{{ .Name }}); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else if eq .BaseType "uint" }}
	if err := WriteBase128VarInt64(w, uint64(a.{{ .Name }})); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else if eq .BaseType "bool" }}
	if err := binary.Write(w, binary.LittleEndian, a.{{ .Name }}); err != nil {
		return errors.Wrap(err, "{{ .Name }}")
	}
	{{- else }}
	Unsupported Base Type : {{ .BaseType }}
	{{- end }}
{{- end }}

package {{ .Package }}

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

{{- range $i, $message := .Messages }}
// WriteDeterministic writes data from a {{ $message.Name }} in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *{{ $message.Name }}) WriteDeterministic(w io.Writer) error {
	{{- range $offset, $field := .Fields }}
		{{- if ne $field.Type "deprecated" }}
		{{- template "WriteDeterministicField" $field }}
		{{- end }}
	{{ end }}

	return nil
}
{{- end }}

{{ range .FieldTypes }}
// WriteDeterministic writes data from a {{ .Name }} in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *{{ .Name }}Field) WriteDeterministic(w io.Writer) error {
	{{- range $offset, $field := .Fields }}
		{{- if ne $field.Type "deprecated" }}
		{{- template "WriteDeterministicField" $field }}
		{{- end }}
	{{ end }}

	return nil
}
{{- end }}
