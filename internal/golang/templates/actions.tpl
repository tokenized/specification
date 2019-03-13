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
var (
	TypeMapping = map[string]OpReturnMessage{
{{- range . }}
		Code{{.Name}}: &{{.Name}}{},
{{- end }}
	}

	// ProtocolID is the current protocol ID
	ProtocolID = []byte("tokenized.com")
)

{{ range $action := . }}

{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{ range .Fields }}	{{ .FieldName }} {{ .FieldGoType }} // {{ .FieldDescription }}
{{ end -}}
}

// New{{.Name}} returns a new empty {{.Name}}.
func NewEmpty{{.Name}}() *{{.Name}} {
	result := {{.Name}}{}
	return &result
}

// New{{.Name}} returns a new {{.Name}} with specified values set.
func New{{.Name}}({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
	result := {{.Name}}{}
	{{ range .Constructor -}}
	result.{{ .ConstructorField -}}
	{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
	{{ end -}}
	return &result
}

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
{{$last := ""}}{{range .PayloadFields}}{{ if .IsVarChar }}
	if err := WriteVarChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{ else if .IsFixedChar }}
	if err := WriteFixedChar(buf, m.{{.FieldName}}, {{.Length}}); err != nil {
		return nil, err
	}
{{ else if .IsInternalTypeArray }}
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

	header, err := NewHeaderForCode([]byte(Code{{.Name}}), len(b))
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
func (m *{{.Name}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
{{$last := ""}}{{range .Fields -}}{{- if .IsVarChar }}
	{
		var err error
		m.{{.FieldName}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
	}
{{ else if .IsFixedChar }}
	{
		var err error
		m.{{.FieldName}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
	}
{{ else if .IsInternalTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		x := &{{.SingularType}}{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.{{.Name}} = append(m.{{.Name}}, *x)
	}
{{ else if .IsNativeTypeArray }}
	m.{{.FieldName}} = make({{.FieldGoType}}, m.{{$last}}, m.{{$last}})
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

// PayloadMessage returns the PayloadMessage, if any.
func (m {{.Name}}) PayloadMessage() (PayloadMessage, error) {
{{- if .HasPayloadMessage }}
	p, err := NewPayloadMessageFromCode([]byte(m.AssetType))
	if p == nil || err != nil {
		return nil, err
	}

	if _, err := p.Write(m.AssetPayload); err != nil {
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
