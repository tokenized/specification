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
{{ end }}

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = byte('F')

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = byte('T')

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = byte('C')

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = byte('R')
)

// TypeMapping holds a mapping of action codes to action types.
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
{{ range .Fields }}	{{ .FieldName }} {{ .FieldGoType }} {{ .GoTags }} // {{ .FieldDescription }}
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
func (action {{.Name}}) Type() string {
	return Code{{.Name}}
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (action *{{.Name}}) read(b []byte) (int, error) {
	data, err := action.serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// serialize returns the full OP_RETURN payload bytes.
func (action *{{.Name}}) serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

{{- range $f, $field := .PayloadFields }}
{{- $fieldName := .FieldName }}
	// {{.FieldName}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIfInt.Field}} == {{ $include }}{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		if err := WriteVarChar(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsFixedChar }}
		if err := WriteFixedChar(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsVarBin }}
		if err := WriteVarBin(buf, action.{{.FieldName}}, {{.Length}}); err != nil {
			return nil, err
		}
{{- else if .IsInternalTypeArray }}
		if err := WriteVariableSize(buf, uint64(len(action.{{.FieldName}})), {{.Length}}, 8); err != nil {
			return nil, err
		}
		for _, value := range action.{{.FieldName}} {
			b, err := value.Serialize()
			if err != nil {
				return nil, err
			}

			if err := write(buf, b); err != nil {
				return nil, err
			}
		}
{{- else if .IsNativeTypeArray }}
		if err := WriteVariableSize(buf, uint64(len(action.{{.FieldName}})), {{.Length}}, 8); err != nil {
			return nil, err
		}
		for _, value := range action.{{.FieldName}} {
			if err := write(buf, value); err != nil {
				return nil, err
			}
		}
{{- else if .IsInternalType }}
		b, err := action.{{.Name}}.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
{{- else if .IsBytes }}
		if err := write(buf, pad(action.{{.FieldName}}, {{.Length}})); err != nil {
			return nil, err
		}
{{- else}}
		if err := write(buf, action.{{.FieldName}}); err != nil {
			return nil, err
		}
{{- end }}
	}
{{ end }}
	return buf.Bytes(), nil
}

// write populates the fields in {{.Name}} from the byte slice
func (action *{{.Name}}) write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

{{- range $f, $field := .PayloadFields }}
	// {{.FieldName}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
	if action.{{ $field.IncludeIfTrue }} {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
	if !action.{{ $field.IncludeIfFalse }} {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
	if {{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} action.{{$field.IncludeIfInt.Field}} == {{ $include }}{{ end }} {
{{- else }}
	{
{{- end }}
{{- if .IsVarChar }}
		var err error
		action.{{.FieldName}}, err = ReadVarChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsFixedChar }}
		var err error
		action.{{.FieldName}}, err = ReadFixedChar(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsVarBin }}
		var err error
		action.{{.FieldName}}, err = ReadVarBin(buf, {{.Length}})
		if err != nil {
			return 0, err
		}
{{- else if .IsInternalTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		action.{{.FieldName}} = make([]{{.SingularType}}, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue {{.SingularType}}
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			action.{{.FieldName}} = append(action.{{.FieldName}}, newValue)
		}
{{- else if .IsNativeTypeArray }}
		size, err := ReadVariableSize(buf, {{.Length}}, 8)
		if err != nil {
			return 0, err
		}
		action.{{.FieldName}} = make({{.FieldGoType}}, size, size)
		if err := read(buf, &action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- else if .IsInternalType }}
		if err := action.{{.FieldName}}.Write(buf); err != nil {
			return 0, err
		}
{{- else if or .IsBytes .IsData }}
		action.{{.FieldName}} = make([]byte, {{.Length}})
		if err := readLen(buf, action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- else if .Trimmable }}
		action.{{.FieldName}} = bytes.Trim(action.{{.FieldName}}, "\x00")
{{- else }}
		if err := read(buf, &action.{{.FieldName}}); err != nil {
			return 0, err
		}
{{- end }}
	}
{{ end }}
	return len(b) - buf.Len(), nil
}

func (m *{{.Name}}) Validate() error {
{{- range $i, $field := .PayloadFields }}

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
			if GetCurrency(value[:]) == nil {
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
		if GetCurrency(string(m.{{.Name}}[:])) == nil {
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

func (action {{.Name}}) String() string {
	vals := []string{}
{{ range .Fields -}}
	{{- if eq .Type "STRING" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%v\"", string(action.{{.FieldName}})))
	{{- else if .IsNumeric }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%v", action.{{.FieldName}}))
	{{- else if eq .Type "SHA" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:\"%x\"", action.{{.FieldName}}))
	{{- else if eq .FieldGoType "[]byte" }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#x", action.{{.FieldName}}))
	{{- else }}
	vals = append(vals, fmt.Sprintf("{{.FieldName}}:%#+v", action.{{.FieldName}}))
	{{- end }}{{ end }}

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
{{ end -}}
