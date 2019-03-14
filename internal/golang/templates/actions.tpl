package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
{{- range .}}
	{{.CodeNameComment}}
	{{.CodeName}} = "{{.Code}}"
{{ end -}}
)

// TypeMapping holds a mapping of message codes to message types.
func TypeMapping(code string) OpReturnMessage {
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

{{ range $action := . }}

{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{ range .Fields }}	{{ .FieldName }} {{ .FieldGoType }} // {{ .FieldDescription }}
{{ end -}}
}

{{ if .Constructor -}}
// New{{.Name}} returns a new {{.Name}} with specified values set.
func New{{.Name}}({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
	result := {{.Name}}{}
	{{ range .Constructor -}}
	result.{{ .ConstructorField -}}
	{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
	{{ end -}}
	return &result
}
{{ end -}}

{{ range .Functions }}// {{.FunctionName}} {{ .FunctionDescription }}.
func (action *{{ $action.Name }}) {{.FunctionName}}({{ range $i, $c := .FunctionParams }}{{if ne .ParamName "hidden"}}{{if $i}}, {{end}}{{ .ParamName }} {{ .ParamGoType }}{{ end -}}{{ end -}}) {
{{ if eq .FunctionType "set" -}}
	{{ range .FunctionParams -}}
	action.{{ .ParamField -}}
	{{ if eq .ParamSetMethod "=" }} = {{ .ParamValue }}{{ else }}.{{ .ParamSetMethod }}({{ .ParamValue }}){{ end }}
	{{ end -}}
{{ else if eq .FunctionType "append" -}}
	action.{{ .FunctionAppendType }}s = append(action.{{ .FunctionAppendType }}s, *New{{ .FunctionAppendType }}({{ range $i, $param := .FunctionParams }}{{if $i}}, {{end}}{{ .ParamName }}{{ end }}))
{{ end -}}
}
{{ end }}

// Type returns the type identifer for this message.
func (m {{.Name}}) Type() string {
	return Code{{.Name}}
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m *{{.Name}}) read(b []byte) (int, error) {
	data, err := m.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m *{{.Name}}) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
{{ range .PayloadFields }}

	// {{.FieldName}} ({{.FieldGoType}})
	// fmt.Printf("Serializing {{.FieldName}}\n")
{{- if .IsVarChar }}
	if err := WriteVarChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsFixedChar }}
	if err := WriteFixedChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsVarBin }}
	if err := WriteVarBin(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsInternalTypeArray }}
	if err := WriteVariableSize(buf, uint64(len(m.{{.FieldName}})), {{.Length}}, 8); err != nil {
		return nil, err
	}
	for _, value := range m.{{.FieldName}} {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}
{{- else if .IsNativeTypeArray }}
	if err := WriteVariableSize(buf, uint64(len(m.{{.FieldName}})), {{.Length}}, 8); err != nil {
		return nil, err
	}
	for _, value := range m.{{.FieldName}} {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}
{{- else if .IsInternalType }}
	{
		b, err := m.{{.Name}}.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
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
{{- end }}
	// fmt.Printf("Serialized {{.FieldName}} : buf len %d\n", buf.Len())
{{ end }}
	return buf.Bytes(), nil
}

// write populates the fields in {{.Name}} from the byte slice
func (m *{{.Name}}) write(b []byte) (int, error) {
	// fmt.Printf("Reading {{.Name}} : %d bytes\n", len(b))
	buf := bytes.NewBuffer(b)

{{- range .Fields }}

	// {{.FieldName}} ({{.FieldGoType}})
	// fmt.Printf("Reading {{.FieldName}} : %d bytes remaining\n", buf.Len())
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
{{- else if .IsPushDataLength }}
	{
		var err error
		m.{{.Name}}, err = ParsePushDataScript(buf)
		if err != nil {
			return err
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
{{- else if .IsInternalType }}
	if err := m.{{.FieldName}}.Write(buf); err != nil {
		return 0, err
	}
{{- else if or .IsBytes .IsData }}
	m.{{.FieldName}} = make([]byte, {{.Length}})
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return 0, err
	}
{{- else if .Trimmable }}
	m.{{.FieldName}} = bytes.Trim(m.{{.FieldName}}, "\x00")
{{- else }}
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return 0, err
	}
{{- end }}

	// fmt.Printf("Read {{.FieldName}} : %d bytes remaining\n%+v\n", buf.Len(), m.{{.FieldName}})
{{ end }}

	// fmt.Printf("Read {{.Name}} : %d bytes remaining\n", buf.Len())
	return len(b) - buf.Len(), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m {{.Name}}) PayloadMessage() (PayloadMessage, error) {
{{- if .HasPayloadMessage }}
	p, err := New([]byte(m.AssetType))
	if p == nil || err != nil {
		return nil, err
	}

	if _, err := p.write(m.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
{{ else }}
	return nil, nil
{{ end -}}
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
