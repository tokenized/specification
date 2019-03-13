package protocol

import "bytes"

{{range .}}
{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{range .Fields}}	{{ .FieldName }} {{ .FieldGoType }} // {{ .FieldDescription }}
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
{{ $last := "" }}{{ range .Fields }}{{ if .IsVarChar }}
	if err := WriteVarChar(buf, m.{{.Name}}, {{.Length}}); err != nil {
		return nil, err
	}
{{ else if .IsFixedChar }}
	if err := WriteFixedChar(buf, m.{{.Name}}, {{.Length}}); err != nil {
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
{{else}}
	if err := write(buf, m.{{.FieldName}}); err != nil {
		return nil, err
	}
{{ end -}}{{ $last = .FieldName }}{{ end }}
	return buf.Bytes(), nil
}

func (m *{{.Name}}) Write(buf *bytes.Buffer) error {
{{ $last := "" }}{{- range .Fields }}{{ if .IsInternalTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		x := &{{.SingularType}}{}
		if err := x.Write(buf); err != nil {
				return err
		}

		m.{{.Name}} = append(m.{{.Name}}, *x)
	}
{{- else if .IsInternalType }}
	if err := m.{{.Name}}.Write(buf); err != nil {
		 return err
	}
{{- else if eq "LenActionPayload" .Name }}
	l := lenForOpPushdata(m.OpPushdata)
	m.LenActionPayload = make([]byte, l, l)
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return err
	}
{{- else if or .IsBytes .IsData }}
	m.{{.FieldName}} = make([]byte, {{.Length}})
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return err
	}
{{ else }}
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return err
	}
{{ end -}}{{ $last = .FieldName }}{{ end }}
	return nil
}
{{ end }}
