package protocol

import "bytes"

{{range .}}
{{comment (print .Name " " .Metadata.Description) "//"}}
type {{.Name}} struct {
{{range .Fields}}	{{ .FieldName }} {{ .FieldGoType }} {{ .GoTags }} // {{ .FieldDescription }}
{{ end -}}
}

// Serialize returns the byte representation of the message.
func (m {{.Name}}) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
{{- range $i, $field := .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if m.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !m.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} m.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
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
		b, err := m.{{.Name}}.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
{{- else if or .IsBytes .IsData }}
		if err := WriteFixedBin(buf, m.{{.Name}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else }}
		if err := write(buf, m.{{.FieldName}}); err != nil {
			return nil, err
		}
{{- end }}
	}
{{ end }}
	return buf.Bytes(), nil
}

func (m *{{.Name}}) Write(buf *bytes.Buffer) error {
{{- range $i, $field := .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if m.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} m.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		var err error
		m.{{.Name}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return err
		}
{{- else if .IsFixedChar }}
		var err error
		m.{{.Name}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return err
		}
{{- else if .IsVarBin }}
		var err error
		m.{{.Name}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return err
		}
{{- else if .IsInternalTypeArray }}
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
{{- else if .IsNativeTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return err
		}
		m.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &m.{{.FieldName}}); err != nil {
			return err
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
{{- end }}
	}
{{ end }}
	return nil
}

func (m *{{.Name}}) Validate() error {
{{- range $i, $field := .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if m.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} m.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		if len(m.{{.Name}}) > (2 << {{.Length}}) - 1 {
			return fmt.Errorf("varchar field {{.Name}} too long %d/%d", len(m.{{.Name}}), (2 << {{.Length}}) - 1)
		}
{{- else if .IsFixedChar }}
		if len(m.{{.Name}}) > {{.Length}} {
			return fmt.Errorf("fixedchar field {{.Name}} too long %d/%d", len(m.{{.Name}}), {{.Length}})
		}
{{- else if .IsVarBin }}
		if len(m.{{.Name}}) > (2 << {{.Length}}) - 1 {
			return fmt.Errorf("varbin field {{.Name}} too long %d/%d", len(m.{{.Name}}), (2 << {{.Length}}) - 1)
		}
{{- else if eq .Type "Role" }}
		roles, err := GetRoles()
		if err != nil {
			return err
		}
		_, exists := roles[m.{{.Name}}]
		if !exists {
			return fmt.Errorf("Invalid role value : %d", m.{{.Name}})
		}
{{- else if eq .Type "MessageType" }}
		messages, err := GetMessages()
		if err != nil {
			return err
		}
		_, exists := messages[m.{{.Name}}]
		if !exists {
			return fmt.Errorf("Invalid message value : %d", m.{{.Name}})
		}
{{- else if eq .Type "Currency" }}
		currencies, err := GetCurrencies()
		if err != nil {
			return err
		}
		_, exists := currencies[string(m.{{.Name}}[:])]
		if !exists {
			return fmt.Errorf("Invalid currency value : %d", m.{{.Name}})
		}
{{- else if eq .Type "Polity" }}
		polities, err := GetPolities()
		if err != nil {
			return err
		}
		_, exists := polities[string(m.{{.Name}}[:])]
		if !exists {
			return fmt.Errorf("Invalid polity value : %d", m.{{.Name}})
		}
{{- else if eq .Type "EntityType" }}
		entities, err := GetEntities()
		if err != nil {
			return err
		}
		_, exists := entities[string{m.{{.Name}}}]
		if !exists {
			return fmt.Errorf("Invalid entity type value : %c", m.{{.Name}})
		}
{{- else if .IsInternalTypeArray }}
		if len(m.{{.Name}}) > (2 << {{.Length}}) - 1 {
			return fmt.Errorf("list field {{.Name}} has too many items %d/%d", len(m.{{.Name}}), (2 << {{.Length}}) - 1)
		}

		for i, value := range m.{{.Name}} {
			err := value.Validate()
			if err != nil {
				return fmt.Errorf("list field {{.Name}}[%d] is invalid : %s", i, err)
			}
		}
{{- else if .IsNativeTypeArray }}
		if len(m.{{.Name}}) > (2 << {{.Length}}) - 1 {
			return fmt.Errorf("list field {{.Name}} has too many items %d/%d", len(m.{{.Name}}), (2 << {{.Length}}) - 1)
		}
{{- else if .IsInternalType }}
		if err := m.{{.Name}}.Validate(); err != nil {
			return fmt.Errorf("field {{.Name}} is invalid : %s", err)
		}
{{- end }}
	}
{{ end }}
	return nil
}

func (m *{{.Name}}) Equal(other {{.Name}}) bool {
{{- range .Fields }}

	// {{.Name}} ({{.FieldGoType}})
{{- if .IsVarBin }}
	if !bytes.Equal(m.{{.Name}}, other.{{.Name}}) {
		return false
	}
{{- else if or .IsBytes .IsData }}
	if !bytes.Equal(m.{{.Name}}, other.{{.Name}}) {
		return false
	}
{{- else if .IsInternalTypeArray }}
	if len(m.{{.Name}}) != 0 || len(other.{{.Name}}) != 0 {
		if len(m.{{.Name}}) != len(other.{{.Name}}) {
			return false
		}

		for i, value := range m.{{.Name}} {
			if !value.Equal(other.{{.Name}}[i]) {
				return false
			}
		}
	}
{{- else if .IsInternalType }}
	if !m.{{.Name}}.Equal(other.{{.Name}}) {
		return false
	}
{{- else }}
	if m.{{.Name}} != other.{{.Name}} {
		return false
	}
{{- end }}
{{- end }}
	return true
}
{{ end }}
