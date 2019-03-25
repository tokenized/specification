package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152
{{range . }}
	// Code{{.Name}} identifies data as a {{.Name}} message.
	Code{{.Name}} = "{{.Code}}"
{{end}})

// AssetTypeMapping holds a mapping of asset codes to asset types.
func AssetTypeMapping(code string) PayloadMessage {
	switch(code) {
{{- range . }}
	case Code{{.Name}}:
		result := {{.Name}}{}
		return &result
{{- end }}
	default:
		return nil
	}
}

{{range .}}
// {{.Name}} asset type.
type {{.Name}} struct {
{{range .Fields }}	{{.Name}} {{.FieldGoType}} {{ .GoTags }} // {{ .FieldDescription }}
{{ end -}}
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
func (m *{{.Name}}) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m *{{.Name}}) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
{{range .Fields}}

	// {{.FieldName}} ({{.FieldGoType}})
{{- if .IsVarChar }}
	if err := WriteVarChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsFixedChar }}
	if err := WriteFixedChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsVarBin }}
	if err := WriteVarBin(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsInternalTypeArray }}
	if err := WriteVariableSize(buf, uint64(len(m.{{.Name}})), {{.Length}}, 8); err != nil {
		return nil, err
	}
	for _, value := range m.{{.Name}} {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
{{- else if .IsNativeTypeArray }}
	if err := WriteVariableSize(buf, uint64(len(m.{{.Name}})), {{.Length}}, 8); err != nil {
		return nil, err
	}
	for _, value := range m.{{.Name}} {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}
{{- else if .IsBytes }}
	if err := write(buf, pad(m.{{.FieldName}}, {{.Length}})); err != nil {
		return nil, err
	}
{{- else}}
	if err := write(buf, m.{{.FieldName}}); err != nil {
		return nil, err
	}
{{ end -}}{{ end }}
	return buf.Bytes(), nil
}

// Write populates the fields in {{.Name}} from the byte slice
func (m *{{.Name}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
{{- range .Fields}}

	// {{.Name}} ({{.FieldGoType}})
{{- if .IsVarChar }}
	{
		var err error
		m.{{.FieldName}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
	}
{{- else if .IsFixedChar }}
	{
		var err error
		m.{{.FieldName}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
	}
{{- else if .IsVarBin }}
	{
		var err error
		m.{{.FieldName}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
	}
{{- else if .IsInternalTypeArray }}
	{
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		m.{{.FieldName}} = make([]{{.SingularType}}, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue {{.SingularType}}
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.{{.FieldName}} = append(m.{{.FieldName}}, newValue)
		}
	}
{{- else if .IsNativeTypeArray }}
	{
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		m.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &m.{{.FieldName}}); err != nil {
			return 0, err
		}
	}
{{ else if .IsInternalType }}
	if err := m.{{.Name}}.Write(buf); err != nil {
		return 0, err
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
{{ end }}{{ end }}
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
	{{- else if eq .FieldGoType "[]byte" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#x", m.{{.FieldName}}))
	{{- else }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#+v", m.{{.FieldName}}))
	{{- end }}{{ end }}

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
{{ end -}}
