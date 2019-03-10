package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152
{{range . }}
	// Code{{.Name}} identifies data as a {{.Name}} message.
	Code{{.Name}} = "{{.Code}}"
{{end}})
{{range .}}
// {{.Name}} asset type.
type {{.Name}} struct {
{{range .Fields }}	{{.Name}} {{.GoType}}
{{ end -}}
}

// New{{.Name}} returns a new {{.Name}}.
func New{{.Name}}() {{.Name}} {
	return {{.Name}}{}
}

// Type returns the type identifer for this message.
func (m {{.Name}}) Type() string {
	return Code{{.Name}}
}

// Len returns the byte size of this message.
func (m {{.Name}}) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m {{.Name}}) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m {{.Name}}) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
{{$last := ""}}{{range .Fields}}{{ if .IsInternalTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		b, err := m.{{.Name}}[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
{{ else if .IsNativeTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		if err := write(buf, m.{{.FieldName}}[i]); err != nil {
			return nil, err
		}
	}
{{ else if .IsInternalType }}
	{
		b, err := m.{{.Name}}.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
{{ else if .IsNvarchar }}
	if _, err := m.{{.FieldName}}.Write(buf); err != nil {
		return nil, err
	}
{{else if .IsBytes }}
	if err := write(buf, pad(m.{{.FieldName}}, {{.Length}})); err != nil {
		return nil, err
	}
{{else}}
	if err := write(buf, m.{{.FieldName}}); err != nil {
		return nil, err
	}
{{ end -}}{{ $last = .Name }}{{ end }}
	b := buf.Bytes()

	header, err := NewHeaderForCode(Code{{.Name}}, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m {{.Name}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
{{$last := ""}}{{range .Fields -}}{{- if .IsInternalTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		x := &{{.GoTypeSingular}}{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.{{.Name}} = append(m.{{.Name}}, *x)
	}
{{ else if .IsNativeTypeArray }}
	m.{{.FieldName}} = make({{.GoType}}, m.{{$last}}, m.{{$last}})
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return 0, err
	}
{{ else if .IsInternalType }}
	if err := m.{{.Name}}.Write(buf); err != nil {
		return 0, err
	}
{{ else if eq .FieldName "AssetPayload" }}
	{
		b := make([]byte, m.{{$last}}, m.{{$last}})
		if err := readLen(buf, b); err != nil {
			return 0, err
		}

		m.{{.FieldName}} = b
	}
{{ else if or .IsBytes .IsData }}
	m.{{.FieldName}} = make([]byte, {{.Length}})
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return 0, err
	}
{{ else if .Trimmable }}
	m.{{.FieldName}} = bytes.Trim(m.{{.FieldName}}, "\x00")
{{ else }}
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return 0, err
	}
{{ end }}{{ $last = .FieldName }}{{ end }}
	return len(b), nil
}

func (m {{.Name}}) String() string {
	vals := []string{}
{{ range .Fields -}}
	{{- if eq .Type "STRING" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%v\"", string(m.{{.FieldName}})))
	{{- else if .IsNumeric }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%v", m.{{.FieldName}}))
	{{- else if eq .Type "SHA" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%x\"", m.{{.FieldName}}))
	{{- else if eq .GoType "[]byte" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#x", m.{{.FieldName}}))
	{{- else }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#+v", m.{{.FieldName}}))
	{{- end }}{{ end }}

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
{{ end -}}
