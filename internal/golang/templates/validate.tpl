package {{ .Package }}

import (
    "fmt"

    "github.com/pkg/errors"
)

const (
    max1ByteInteger = 255
    max2ByteInteger = 65535
    max4ByteInteger = 4294967295
)

{{ define "ListValues" -}}
{{- range $i, $v := . -}}
{{ if gt $i 0 }}, {{ end }}{{ $v }}
{{- end -}}
{{- end }}

{{ define "ValidateField" -}}
    // Field {{ .Name }} - {{ .BaseType }}
    {{- if .IsList }}
        {{- if .HasOnlyValidWhen }}
        validValueFound{{ .Name }} := false
        for _, v := range []{{ .OnlyValidWhen.FieldGoType }}{ {{ template "ListValues" .OnlyValidWhen.Values }} } {
            if a.{{ .OnlyValidWhen.FieldName }} == v {
                validValueFound{{ .Name }} = true
                break
            }
        }
        if !validValueFound{{ .Name }} && len(a.{{ .Name }}) != 0 {
            return fmt.Errorf("{{ .Name }} not allowed. {{ .OnlyValidWhen.FieldName }} value not within values {{ .OnlyValidWhen.Values }} : %v", a.{{ .OnlyValidWhen.FieldName }})
        }
        {{- end }}
        {{- if eq .BaseListSize "tiny" "" }}
        if len(a.{{ .Name }}) > max1ByteInteger {
            return fmt.Errorf("{{ .Name }} list over max length : %d > %d", len(a.{{ .Name }}), max1ByteInteger)
        }
        {{- else if eq .BaseListSize "small" }}
        if len(a.{{ .Name }}) > max2ByteInteger {
            return fmt.Errorf("{{ .Name }} list over max length : %d > %d", len(a.{{ .Name }}), max2ByteInteger)
        }
        {{- else if eq .BaseListSize "medium" }}
        if len(a.{{ .Name }}) > max4ByteInteger {
            return fmt.Errorf("{{ .Name }} list over max length : %d > %d", len(a.{{ .Name }}), max4ByteInteger)
        }
        {{- else if ne .BaseListSize "large" }}
        INVALID LIST SIZE : {{ .BaseListSize }}
        {{- end }}
        {{- if and (or (ne .BaseType "uint") (and (ne .BaseSize 4) (ne .BaseSize 8))) (or (ne .BaseType "varbin") (ne .BaseVarSize "large")) }}
        for i, v := range a.{{ .Name }} {
        {{- if gt (len .BaseResource) 0 }}
            if {{ .BaseResource }}Data(v) == nil {
                return fmt.Errorf("{{ .Name }}[%d] resource {{ .BaseResource }} value not defined : %v", i, v)
            }
        {{- end }}
        {{- if eq .BaseType "fixedchar" "bin" }}
            if len(v) != 0 && len(v) != {{ .BaseSize }} {
                return fmt.Errorf("{{ .Name }}[%d] fixed width element wrong size : %d should be %d",
                    i, len(v), {{ .BaseSize }})
            }
        {{- else if eq .BaseType "varbin" "varchar" }}
            {{- if eq .BaseVarSize "tiny" "" }}
            if len(v) > max1ByteInteger {
                return fmt.Errorf("[%d] {{ .Name }} size over max value : %d > %d", i, len(v), max1ByteInteger)
            }
            {{- else if eq .BaseVarSize "small" }}
            if len(v) > max2ByteInteger {
                return fmt.Errorf("[%d] {{ .Name }} size over max value : %d > %d", i, len(v), max2ByteInteger)
            }
            {{- else if eq .BaseVarSize "medium" }}
            if len(v) > max4ByteInteger {
                return fmt.Errorf("[%d] {{ .Name }} size over max value : %d > %d", i, len(v), max4ByteInteger)
            }
            {{- else if ne .BaseVarSize "large" }}
            INVALID VARIABLE SIZE : {{ .BaseVarSize }}
            {{- end }}
        {{- else if eq .BaseType "uint" }}
            {{- if gt (len .BaseOptions) 0 }}
            found{{ .Name }} := false
            for _, o := range []{{ .GoSingularType }}{ {{ template "ListValues" .BaseOptions }} } {
                if v == o {
                    found{{ .Name }} = true
                    break
                }
            }
            if !found{{ .Name }} {
                return fmt.Errorf("{{ .Name }}[%d] value not within options {{ .BaseOptions }} : %d", i, a.{{ .Name }})
            }
            {{- else if le .BaseSize 1 }}
            if v > uint32(max1ByteInteger) {
                return fmt.Errorf("{{ .Name }}[%d] over max value : %d > %d", i, v, max1ByteInteger)
            }
            {{- else if eq .BaseSize 2 }}
            if v > uint32(max2ByteInteger) {
                return fmt.Errorf("{{ .Name }}[%d] over max value : %d > %d", i, v, max2ByteInteger)
            }
            {{- end }}
        {{- else if not .IsPrimitive }}
            if err := v.Validate(); err != nil {
                return errors.Wrap(err, fmt.Sprintf("{{ .Name }}[%d]", i))
            }
        {{- end }}
        }
        {{- end }}
    {{- else }}
        {{- if gt (len .BaseResource) 0 }}
            if {{ .BaseResource }}Data(a.{{ .Name }}) == nil {
                return fmt.Errorf("{{ .Name }} resource {{ .BaseResource }} value not defined : %v", a.{{ .Name }})
            }
        {{- end }}
        {{- if eq .BaseType "fixedchar" "bin" }}
        if len(a.{{ .Name }}) != 0 && len(a.{{ .Name }}) != {{ .BaseSize }} {
            return fmt.Errorf("{{ .Name }} fixed width field wrong size : %d should be %d",
                len(a.{{ .Name }}), {{ .BaseSize }})
        }
        {{- else if eq .BaseType "varbin" "varchar" }}
            {{- if .HasOnlyValidWhen }}
        validValueFound{{ .Name }} := false
        for _, v := range []{{ .OnlyValidWhen.FieldGoType }}{ {{ template "ListValues" .OnlyValidWhen.Values }} } {
            if a.{{ .OnlyValidWhen.FieldName }} == v {
                validValueFound{{ .Name }} = true
                break
            }
        }
        if !validValueFound{{ .Name }} && len(a.{{ .Name }}) != 0 {
            return fmt.Errorf("{{ .Name }} not allowed. {{ .OnlyValidWhen.FieldName }} value not within values {{ .OnlyValidWhen.Values }} : %v", a.{{ .OnlyValidWhen.FieldName }})
        }
            {{- end }}
            {{- if .HasRequiredWhen }}
        requiredValueFound{{ .Name }} := false
        for _, v := range []{{ .RequiredWhen.FieldGoType }}{ {{ template "ListValues" .RequiredWhen.Values }} } {
            if a.{{ .RequiredWhen.FieldName }} == v {
                requiredValueFound{{ .Name }} = true
                break
            }
        }
        if requiredValueFound{{ .Name }} && len(a.{{ .Name }}) == 0 {
            return fmt.Errorf("{{ .Name }} required. {{ .RequiredWhen.FieldName }} value within values {{ .RequiredWhen.Values }} : %v", a.{{ .RequiredWhen.FieldName }})
        }
            {{- end }}
            {{- if eq .BaseVarSize "tiny" "" }}
        if len(a.{{ .Name }}) > max1ByteInteger {
            return fmt.Errorf("{{ .Name }} over max size : %d > %d", len(a.{{ .Name }}), max1ByteInteger)
        }
            {{- else if eq .BaseVarSize "small" }}
        if len(a.{{ .Name }}) > max2ByteInteger {
            return fmt.Errorf("{{ .Name }} over max size : %d > %d", len(a.{{ .Name }}), max2ByteInteger)
        }
            {{- else if eq .BaseVarSize "medium" }}
        if len(a.{{ .Name }}) > max4ByteInteger {
            return fmt.Errorf("{{ .Name }} over max size : %d > %d", len(a.{{ .Name }}), max4ByteInteger)
        }
            {{- else if ne .BaseVarSize "large" }}
        INVALID VARIABLE SIZE : {{ .BaseVarSize }}
            {{- end }}
        {{- else if eq .BaseType "uint" }}
            {{- if gt (len .BaseOptions) 0 }}
        found{{ .Name }} := false
        for _, v := range []{{ .GoType }}{ {{ template "ListValues" .BaseOptions }} } {
            if a.{{ .Name }} == v {
                found{{ .Name }} = true
                break
            }
        }
        if !found{{ .Name }} {
            return fmt.Errorf("{{ .Name }} value not within options {{ .BaseOptions }} : %d", a.{{ .Name }})
        }
            {{- else if le .BaseSize 1 }}
        if a.{{ .Name }} > uint32(max1ByteInteger) {
            return fmt.Errorf("{{ .Name }} over max value : %d > %d", a.{{ .Name }}, max1ByteInteger)
        }
            {{- else if eq .BaseSize 2 }}
        if a.{{ .Name }} > uint32(max2ByteInteger) {
            return fmt.Errorf("{{ .Name }} over max value : %d > %d", a.{{ .Name }}, max2ByteInteger)
        }
            {{- end }}
        {{- else if not .IsPrimitive }}
            {{- if .HasOnlyValidWhen }}
        validValueFound{{ .Name }} := false
        for _, v := range []{{ .OnlyValidWhen.FieldGoType }}{ {{ template "ListValues" .OnlyValidWhen.Values }} } {
            if a.{{ .OnlyValidWhen.FieldName }} == v {
                validValueFound{{ .Name }} = true
                break
            }
        }
        if !validValueFound{{ .Name }} && a.{{ .Name }} != nil {
            return fmt.Errorf("{{ .Name }} not allowed. {{ .OnlyValidWhen.FieldName }} value not within values {{ .OnlyValidWhen.Values }} : %v", a.{{ .OnlyValidWhen.FieldName }})
        }
            {{- end }}
            {{- if .HasRequiredWhen }}
        requiredValueFound{{ .Name }} := false
        for _, v := range []{{ .RequiredWhen.FieldGoType }}{ {{ template "ListValues" .RequiredWhen.Values }} } {
            if a.{{ .RequiredWhen.FieldName }} == v {
                requiredValueFound{{ .Name }} = true
                break
            }
        }
        if requiredValueFound{{ .Name }} && a.{{ .Name }} == nil {
            return fmt.Errorf("{{ .Name }} required. {{ .RequiredWhen.FieldName }} value within values {{ .RequiredWhen.Values }} : %v", a.{{ .RequiredWhen.FieldName }})
        }
            {{- end }}
        if err := a.{{ .Name }}.Validate(); err != nil {
            return errors.Wrap(err, "{{ .Name }}")
        }
        {{- end }}
    {{- end }}
{{ end }}

{{ range .Messages }}
func (a *{{.Name}}) Validate() error {
    if a == nil {
        return nil
    }
    {{ range .Fields }}
        {{- if ne .Type "deprecated" }}
        {{ template "ValidateField" . -}}
        {{- end }}
    {{- end }}

    return nil
}
{{ end }}

{{ range .FieldTypes }}
func (a *{{.Name}}Field) Validate() error {
    if a == nil {
        return nil
    }
    {{ range .Fields }}
        {{- if ne .Type "deprecated" }}
        {{ template "ValidateField" . -}}
        {{- end }}
    {{- end }}

    return nil
}
{{ end }}
