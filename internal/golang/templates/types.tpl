// Package {{.Package}} provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.
package {{.Package}}

{{range .TypeMessages}}
{{.GoComment}}
type {{.Name}} struct {
	{{range .Fields}}{{ .FieldName }} {{ .GoType }}
	{{ end -}}
}

func New{{.Name}}() {{.Name}} {
	return {{.Name}}{}
}

// Type returns the type identifer for this message.
func (m {{.Name}}) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	{{$last := ""}}
	{{range .Fields}}
	{{ if .IsInternalTypeArray }}
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
	{{ end -}}
	{{ $last = .FieldName }}
	{{ end }}

	return buf.Bytes(), nil
}

func (m *{{.Name}}) Write(buf *bytes.Buffer) error {
	{{$last := ""}}
	{{- range .Fields }}
	{{ if .IsInternalTypeArray }}
	for i := 0; i < int(m.{{$last}}); i++ {
		x := &{{.GoTypeSingular}}{}
		if err := x.Write(buf); err != nil {
			return err
		}

		m.{{.Name}} = append(m.{{.Name}}, *x)
	}
	{{- else if .IsInternalType -}}
	if err := m.{{.Name}}.Write(buf); err != nil {
		 return err
	}
	{{- else if eq "LenActionPayload" .Name -}}
	l := lenForOpPushdata(m.OpPushdata)
	m.LenActionPayload = make([]byte, l, l)
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return err
	}
	{{- else if or .IsBytes .IsData -}}
	m.{{.FieldName}} = make([]byte, {{.Length}})
	if err := readLen(buf, m.{{.FieldName}}); err != nil {
		return err
	}

	{{else}}
	if err := read(buf, &m.{{.FieldName}}); err != nil {
		return err
	}
	{{ end -}}
	{{ $last = .FieldName }}
	{{ end }}

	return nil
}

{{ end }}
