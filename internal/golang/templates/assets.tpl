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

// AssetPayload is the interface for payloads within asset actions.
type AssetPayload interface {
	Type() string
	Serialize() ([]byte, error)
	Write(b []byte) (int, error)
	Validate() error
}

// AssetTypeMapping holds a mapping of asset codes to asset types.
func AssetTypeMapping(code string) AssetPayload {
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
{{range $f, $field := .Fields}}

	// {{.FieldName}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
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
{{ end -}}
	}
{{ end }}
	return buf.Bytes(), nil
}

// Write populates the fields in {{.Name}} from the byte slice
func (m *{{.Name}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
{{- range $f, $field := .Fields}}

	// {{.Name}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		var err error
		m.{{.FieldName}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsFixedChar }}
		var err error
		m.{{.FieldName}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsVarBin }}
		var err error
		m.{{.FieldName}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsInternalTypeArray }}
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
{{- else if .IsNativeTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		m.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &m.{{.FieldName}}); err != nil {
			return 0, err
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
{{ end }}
	}
{{ end }}
	return len(b), nil
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
{{- else if .IsResourceTypeArray }}
		if len(m.{{.Name}}) > (2 << {{.Length}}) - 1 {
			return fmt.Errorf("list field {{.Name}} has too many items %d/%d", len(m.{{.Name}}), (2 << {{.Length}}) - 1)
		}

		for _, value := range m.{{.Name}} {
	{{- if eq .Type "RejectionCode[]" }}
			if GetRejectionCode(value) == nil {
				return fmt.Errorf("Invalid rejection code value : %d", m.{{.Name}})
			}
	{{- else if eq .Type "Role[]" }}
			if GetRoleType(value) == nil {
				return fmt.Errorf("Invalid role value : %d", m.{{.Name}})
			}
	{{- else if eq .Type "CurrencyType[]" }}
			if GetCurrency(value) == nil {
				return fmt.Errorf("Invalid currency value : %d", m.{{.Name}})
			}
	{{- else if eq .Type "Polity[]" }}
			if GetPolityType(string(value[:])) == nil {
				return fmt.Errorf("Invalid polity value : %d", m.{{.Name}})
			}
	{{- else if eq .Type "EntityType[]" }}
			if GetEntityType(value) == nil {
				return fmt.Errorf("Invalid entity type value : %c", m.{{.Name}})
			}
	{{- end }}
		}
{{- else if eq .Type "RejectionCode" }}
		if GetRejectionCode(m.{{.Name}}) == nil {
			return fmt.Errorf("Invalid rejection code value : %d", m.{{.Name}})
		}
{{- else if eq .Type "Role" }}
		if GetRoleType(m.{{.Name}}) == nil {
			return fmt.Errorf("Invalid role value : %d", m.{{.Name}})
		}
{{- else if eq .Type "CurrencyType" }}
		if GetCurrency(m.{{.Name}}) == nil {
			return fmt.Errorf("Invalid currency value : %d", m.{{.Name}})
		}
{{- else if eq .Type "Polity" }}
		if GetPolityType(string(m.{{.Name}}[:])) == nil {
			return fmt.Errorf("Invalid polity value : %d", m.{{.Name}})
		}
{{- else if eq .Type "EntityType" }}
		if GetEntityType(m.{{.Name}}) == nil {
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
{{ else if ne (len $field.IntValues) 0 }}
		if {{ range $j, $value := $field.IntValues }}{{ if (ne $j 0) }} &&{{ end }} m.{{$field.Name}} != {{ $value }}{{ end }} {
			return fmt.Errorf("field {{$field.Name}} value is invalid : %d", m.{{$field.Name}})
		}
{{ else if ne (len $field.CharValues) 0 }}
		if {{ range $j, $value := $field.CharValues }}{{ if (ne $j 0) }} &&{{ end }} m.{{$field.Name}} != '{{ $value }}'{{ end }} {
			return fmt.Errorf("field {{$field.Name}} value is invalid : %d", m.{{$field.Name}})
		}
{{- end }}
	}
{{ end }}
	return nil
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
