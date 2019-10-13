package {{ .Package }}

import (
    "errors"
    "fmt"
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
        {{- if eq .BaseListSize "tiny" "" }}
        if len(a.{{ .Name }}) > max1ByteInteger {
            return fmt.Errorf("List over max length : %d > %d", len(a.{{ .Name }}), max1ByteInteger)
        }
        {{- else if eq .BaseListSize "small" }}
        if len(a.{{ .Name }}) > max2ByteInteger {
            return fmt.Errorf("List over max length : %d > %d", len(a.{{ .Name }}), max2ByteInteger)
        }
        {{- else if eq .BaseListSize "medium" }}
        if len(a.{{ .Name }}) > max4ByteInteger {
            return fmt.Errorf("List over max length : %d > %d", len(a.{{ .Name }}), max4ByteInteger)
        }
        {{- else if ne .BaseListSize "large" }}
        INVALID LIST SIZE : {{ .BaseListSize }}
        {{- end }}
        {{- if or (ne .BaseType "uint") (and (ne .BaseSize 4) (ne .BaseSize 8)) }}
        for i, v := range a.{{ .Name }} {
        {{- if gt (len .BaseResource) 0 }}
            if {{ .BaseResource }}Data(v) == nil {
                return fmt.Errorf("{{ .BaseResource }} resource value not defined : %v", v)
            }
        {{- end }}
        {{- if eq .BaseType "fixedchar" "bin" }}
            if len(v) != 0 && len(v) != {{ .BaseSize }} {
                return fmt.Errorf("Fixed width element {{ .Name }}[%d] wrong size : %d should be %d",
                    i, len(v), {{ .BaseSize }})
            }
        {{- else if eq .BaseType "varbin" "varchar" }}
            {{- if eq .BaseVarSize "tiny" "" }}
            if len(v) > max1ByteInteger {
                return fmt.Errorf("variable size over max value : %d > %d", len(v), max1ByteInteger)
            }
            {{- else if eq .BaseVarSize "small" }}
            if len(v) > max2ByteInteger {
                return fmt.Errorf("variable size over max value : %d > %d", len(v), max2ByteInteger)
            }
            {{- else if eq .BaseVarSize "medium" }}
            if len(v) > max4ByteInteger {
                return fmt.Errorf("variable size over max value : %d > %d", len(v), max4ByteInteger)
            }
            {{- else if ne .BaseVarSize "large" }}
            INVALID VARIABLE SIZE : {{ .BaseVarSize }}
            {{- end }}
        {{- else if eq .BaseType "uint" }}
            {{- if le .BaseSize 1 }}
            if v > uint32(max1ByteInteger) {
                return fmt.Errorf("{{ .Name }}[%d] uint over max value : %d > %d", i, v, max1ByteInteger)
            }
            {{- else if eq .BaseSize 2 }}
            if v > uint32(max2ByteInteger) {
                return fmt.Errorf("{{ .Name }}[%d] uint over max value : %d > %d", i, v, max2ByteInteger)
            }
            {{- end }}
        {{- else if not .IsPrimitive }}
            if err := v.Validate(); err != nil {
                return fmt.Errorf("{{ .Name }}[%d] invalid : %s", i, err)
            }
        {{- end }}
        }
        {{- end }}
    {{- else }}
        {{- if gt (len .BaseResource) 0 }}
            if {{ .BaseResource }}Data(a.{{ .Name }}) == nil {
                return fmt.Errorf("{{ .BaseResource }} resource value not defined : %v", a.{{ .Name }})
            }
        {{- end }}
        {{- if eq .BaseType "fixedchar" "bin" }}
        if len(a.{{ .Name }}) != 0 && len(a.{{ .Name }}) != {{ .BaseSize }} {
            return fmt.Errorf("Fixed width field {{ .Name }} wrong size : %d should be %d",
                len(a.{{ .Name }}), {{ .BaseSize }})
        }
        {{- else if eq .BaseType "varbin" "varchar" }}
            {{- if eq .BaseVarSize "tiny" "" }}
        if len(a.{{ .Name }}) > max1ByteInteger {
            return fmt.Errorf("variable size over max value : %d > %d", len(a.{{ .Name }}), max1ByteInteger)
        }
            {{- else if eq .BaseVarSize "small" }}
        if len(a.{{ .Name }}) > max2ByteInteger {
            return fmt.Errorf("variable size over max value : %d > %d", len(a.{{ .Name }}), max2ByteInteger)
        }
            {{- else if eq .BaseVarSize "medium" }}
        if len(a.{{ .Name }}) > max4ByteInteger {
            return fmt.Errorf("variable size over max value : %d > %d", len(a.{{ .Name }}), max4ByteInteger)
        }
            {{- else if ne .BaseVarSize "large" }}
        INVALID VARIABLE SIZE : {{ .BaseVarSize }}
            {{- end }}
        {{- else if eq .BaseType "uint" }}
            {{- if gt (len .Options) 0 }}
        found{{ .Name }} := false
        for _, v := range []{{ .GoType }}{ {{ template "ListValues" .Options }} } {
            if a.{{ .Name }} == v {
                found{{ .Name }} = true
                break
            }
        }
        if !found{{ .Name }} {
            return fmt.Errorf("{{ .Name }} value not within options {{ .Options }} : %d", a.{{ .Name }})
        }
            {{- else if le .BaseSize 1 }}
        if a.{{ .Name }} > uint32(max1ByteInteger) {
            return fmt.Errorf("uint over max value : %d > %d", a.{{ .Name }}, max1ByteInteger)
        }
            {{- else if eq .BaseSize 2 }}
        if a.{{ .Name }} > uint32(max2ByteInteger) {
            return fmt.Errorf("uint over max value : %d > %d", a.{{ .Name }}, max2ByteInteger)
        }
            {{- end }}
        {{- else if not .IsPrimitive }}
            if err := a.{{ .Name }}.Validate(); err != nil {
                return fmt.Errorf("{{ .Name }} invalid : %s", err)
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
        {{ template "ValidateField" . -}}
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
        {{ template "ValidateField" . -}}
    {{- end }}

    return nil
}
{{ end }}
