package protocol

import "bytes"

{{range .}}
{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{range .Fields}}	{{ .FieldName }} {{ .FieldGoType }} {{ .GoTags }} // {{ .FieldDescription }}
{{ end -}}
}

// New{{.Name}} returns a new {{.Name}} with defaults set.
func New{{.Name}}({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
	result := {{.Name}}{}
	{{ range .Constructor -}}
	result.{{ .ConstructorField -}}
	{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
	{{ end -}}
	return &result
}

// Serialize returns the byte representation of the message.
func (m {{.Name}}) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
{{- range .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if .IsVarChar }}
	if err := WriteVarChar(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsFixedChar }}
	if err := WriteFixedChar(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsVarBin }}
	if err := WriteVarBin(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else if .IsPushDataLength }}
	if _, err := buf.Write(PushDataScript(m.{{.Name}})); err != nil {
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
{{- else if or .IsBytes .IsData }}
	if err := WriteFixedBin(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{- else }}
	if err := write(buf, m.{{.FieldName}}); err != nil {
		return nil, err
	}
{{- end }}{{ end }}
	return buf.Bytes(), nil
}

func (m *{{.Name}}) Write(buf *bytes.Buffer) error {
{{- range .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if .IsVarChar }}
	{
		var err error
		m.{{.Name}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return err
		}
	}
{{- else if .IsFixedChar }}
	{
		var err error
		m.{{.Name}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return err
		}
	}
{{- else if .IsVarBin }}
	{
		var err error
		m.{{.Name}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return err
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
			return err
		}
		m.{{.FieldName}} = make([]{{.SingularType}}, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue {{.SingularType}}
			if err := newValue.Write(buf); err != nil {
				return err
			}

			m.{{.FieldName}} = append(m.{{.FieldName}}, newValue)
		}
	}
{{- else if .IsNativeTypeArray }}
	{
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return err
		}
		m.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &m.{{.FieldName}}); err != nil {
			return err
		}
	}
{{- else if .IsInternalType }}
	if err := m.{{.Name}}.Write(buf); err != nil {
		 return err
	}
{{- else if or .IsBytes .IsData }}
	m.{{.FieldName}} = make([]byte, {{.Length}})
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return err
	}
{{- else }}
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return err
	}
{{- end }}{{ end }}
	return nil
}
{{ end }}
