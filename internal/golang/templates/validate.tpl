package {{ .Package }}

import (
    "fmt"
    "regexp"

    "github.com/tokenized/pkg/bitcoin"

    "github.com/pkg/errors"
)

const (
    max1ByteInteger = 255
    max2ByteInteger = 65535
    max4ByteInteger = 4294967295

    maxArticleDepth = 4
)

{{ define "ListValues" -}}
{{- range $i, $v := . -}}
{{ if gt $i 0 }}, {{ end }}{{ $v }}
{{- end -}}
{{- end }}

{{ define "ListValuesString" -}}
{{- range $i, $v := . -}}
{{ if gt $i 0 }}, {{ end }}"{{ $v }}"
{{- end -}}
{{- end }}

{{ define "SetFieldIsEmpty" -}}
    {{- if or .IsList (eq .BaseType "fixedchar" "bin" "varbin" "varchar") }}
    {{ .Name }}FieldIsEmpty := len(a.{{ .Name }}) == 0
    {{- else if eq .BaseType "uint" }}
    {{ .Name }}FieldIsEmpty := a.{{ .Name }} == 0
    {{- else if not .IsPrimitive }}
    {{ .Name }}FieldIsEmpty := a.{{ .Name }} == nil
    {{- else }}
    var {{ .Name }}FieldIsEmpty bool
    {
        var z {{ .GoType }} // zero value {{ .BaseType }}
        {{ .Name }}FieldIsEmpty = a.{{ .Name }} == z
    }
    {{- end }}
{{- end }}

{{ define "IsEmptyCheck" -}}
    {{- if or .IsList (eq .BaseType "fixedchar" "bin" "varbin" "varchar") -}}
    len(a.{{ .Name }}) == 0
    {{- else if eq .BaseType "uint" -}}
    a.{{ .Name }} == 0
    {{- else if eq .BaseType "bool" -}}
    a.{{ .Name }} == false
    {{- else if not .IsPrimitive -}}
    a.{{ .Name }} == nil
    {{- else -}}
    Unsupported type for empty check {{ .BaseType }}
    {{- end -}}
{{- end }}

{{ define "IsNotEmptyCheck" -}}
    {{- if or .IsList (eq .BaseType "fixedchar" "bin" "varbin" "varchar") -}}
    len(a.{{ .Name }}) != 0
    {{- else if eq .BaseType "uint" -}}
    a.{{ .Name }} != 0
    {{- else if eq .BaseType "bool" -}}
    a.{{ .Name }} != false
    {{- else if not .IsPrimitive -}}
    a.{{ .Name }} != nil
    {{- else -}}
    Unsupported type for not empty check {{ .BaseType }}
    {{- end -}}
{{- end }}

{{ define "ValidateField" -}}
    // Field {{ .Name }} - {{ .BaseType }} {{ if gt (len .BaseResource) 0 }} ({{ .BaseResource }} Resource) {{ end }}
    {{- if .IsList }}
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
        {{- if and (gt (len .BaseResource) 0) (ne .BaseResource "LegalSystems") (ne .BaseResource "Polities") }}
            if {{ .BaseResource }}Data(v) == nil {
                return fmt.Errorf("{{ .Name }}[%d] resource {{ .BaseResource }} value not defined : %v", i, v)
            }
        {{- else if .IsAlias }}
            {{- if eq .AliasField.Name "Address" }}
            if len(v) > 0 {
                if err := AddressIsValid(v); err != nil {
                    return errors.Wrap(err, fmt.Sprintf("{{ .Name }}[%d]", i))
                }
            }
            {{- else if eq .AliasField.Name "PublicKey" }}
            if len(v) > 0 {
                if err := PublicKeyIsValid(v); err != nil {
                    return errors.Wrap(err, fmt.Sprintf("{{ .Name }}[%d]", i))
                }
            }
            {{- else if eq .AliasField.Name "Signature" }}
            if len(v) > 0 {
                if err := SignatureIsValid(v); err != nil {
                    return errors.Wrap(err, fmt.Sprintf("{{ .Name }}[%d]", i))
                }
            }
            {{- end }}
        {{- end }}
        {{- if eq .BaseType "fixedchar" "bin" }}
            if len(v) != 0 && len(v) != {{ .BaseSize }} {
                return fmt.Errorf("{{ .Name }}[%d] fixed width element wrong size : %d should be %d",
                    i, len(v), {{ .BaseSize }})
            }
            {{- if and (gt (len .BaseOptions) 0) (eq .BaseType "fixedchar") }}
            found{{ .Name }} := false
            for _, o := range []{{ .GoSingularType }}{ {{ template "ListValuesString" .BaseOptions }} } {
                if v == o {
                    found{{ .Name }} = true
                    break
                }
            }
            if !found{{ .Name }} {
                return fmt.Errorf("{{ .Name }}[%d] value not within options {{ .BaseOptions }} : %s", i, a.{{ .Name }})
            }
            {{- end }}
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
        {{- if and (gt (len .BaseResource) 0) (ne .BaseResource "LegalSystems") (ne .BaseResource "Polities") }}
            if {{ .BaseResource }}Data(a.{{ .Name }}) == nil {
                return fmt.Errorf("{{ .Name }} resource {{ .BaseResource }} value not defined : %v", a.{{ .Name }})
            }
        {{- else if .IsAlias }}
            {{- if eq .AliasField.Name "Address" }}
        if len(a.{{ .Name }}) > 0 {
            if err := AddressIsValid(a.{{ .Name }}); err != nil {
                return errors.Wrap(err, "{{ .Name }}")
            }
        }
            {{- else if eq .AliasField.Name "PublicKey" }}
        if len(a.{{ .Name }}) > 0 {
            if err := PublicKeyIsValid(a.{{ .Name }}); err != nil {
                return errors.Wrap(err, "{{ .Name }}")
            }
        }
            {{- else if eq .AliasField.Name "Signature" }}
        if len(a.{{ .Name }}) > 0 {
            if err := SignatureIsValid(a.{{ .Name }}); err != nil {
                return errors.Wrap(err, "{{ .Name }}")
            }
        }
            {{- end }}
        {{- end }}

        {{- if eq .BaseType "fixedchar" "bin" }}
        if len(a.{{ .Name }}) != 0 && len(a.{{ .Name }}) != {{ .BaseSize }} {
            return fmt.Errorf("{{ .Name }} fixed width field wrong size : %d should be %d",
                len(a.{{ .Name }}), {{ .BaseSize }})
        }
            {{- if and (gt (len .BaseOptions) 0) (eq .BaseType "fixedchar") }}
        found{{ .Name }} := false
        for _, o := range []{{ .GoSingularType }}{ {{ template "ListValuesString" .BaseOptions }} } {
            if a.{{ .Name }} == o {
                found{{ .Name }} = true
                break
            }
        }
        if !found{{ .Name }} {
            return fmt.Errorf("{{ .Name }} value not within options {{ .BaseOptions }} : %s", a.{{ .Name }})
        }
            {{- end }}
        {{- else if eq .BaseType "varbin" "varchar" }}
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
        if err := a.{{ .Name }}.Validate(); err != nil {
            return errors.Wrap(err, "{{ .Name }}")
        }
        {{- end }}
    {{- end }}
    {{- if .HasOnlyValidWhen }}
        {{- if (eq (len .OnlyValidWhen.Values) 0) }}
    if {{ .OnlyValidWhen.FieldName }}FieldIsEmpty && {{ template "IsNotEmptyCheck" . }} {
        return fmt.Errorf("{{ .Name }} is only allowed when {{ .OnlyValidWhen.FieldName }} is specified : %v", a.{{ .OnlyValidWhen.FieldName }})
    }
        {{- else }}
    validValueFound{{ .Name }} := false
    {{ if eq .OnlyValidWhen.FieldGoType "string" }}
    for _, v := range []{{ .OnlyValidWhen.FieldGoType }}{ {{ template "ListValuesString" .OnlyValidWhen.Values }} } {
        if a.{{ .OnlyValidWhen.FieldName }} == v {
            validValueFound{{ .Name }} = true
            break
        }
    }
    {{- else -}}
    for _, v := range []{{ .OnlyValidWhen.FieldGoType }}{ {{ template "ListValues" .OnlyValidWhen.Values }} } {
        if a.{{ .OnlyValidWhen.FieldName }} == v {
            validValueFound{{ .Name }} = true
            break
        }
    }
    {{- end }}
    if !validValueFound{{ .Name }} && {{ template "IsNotEmptyCheck" . }} {
        return fmt.Errorf("{{ .Name }} is only allowed when {{ .OnlyValidWhen.FieldName }} value is within values {{ .OnlyValidWhen.Values }} : %v", a.{{ .OnlyValidWhen.FieldName }})
    }
        {{- end }}
    {{- end }}
    {{- if .HasRequiredWhen }}
        {{- if (eq (len .RequiredWhen.Values) 0) }}
    if !{{ .RequiredWhen.FieldName }}FieldIsEmpty && {{ template "IsEmptyCheck" . }} {
        return fmt.Errorf("{{ .Name }} is required when {{ .RequiredWhen.FieldName }} is specified : %v", a.{{ .RequiredWhen.FieldName }})
    }
        {{- else }}
    requiredValueFound{{ .Name }} := false
    {{ if eq .RequiredWhen.FieldGoType "string" }}
    for _, v := range []{{ .RequiredWhen.FieldGoType }}{ {{ template "ListValuesString" .RequiredWhen.Values }} } {
        if a.{{ .RequiredWhen.FieldName }} == v {
            requiredValueFound{{ .Name }} = true
            break
        }
    }
    {{- else -}}
    for _, v := range []{{ .RequiredWhen.FieldGoType }}{ {{ template "ListValues" .RequiredWhen.Values }} } {
        if a.{{ .RequiredWhen.FieldName }} == v {
            requiredValueFound{{ .Name }} = true
            break
        }
    }
    {{- end }}
    if requiredValueFound{{ .Name }} && {{ template "IsEmptyCheck" . }} {
        return fmt.Errorf("{{ .Name }} is required when {{ .RequiredWhen.FieldName }} value is within values {{ .RequiredWhen.Values }} : %v", a.{{ .RequiredWhen.FieldName }})
    }
        {{- end }}
    {{- end }}
    {{- if .Required }}
    if {{ template "IsEmptyCheck" . }} {
        return fmt.Errorf("{{ .Name }} required")
    }
    {{- end }}
{{ end }}

{{- $termsNeeded := false }}

{{ range .Messages }}
func (a *{{.Name}}) Validate() error {
    if a == nil {
        return errors.New("Empty")
    }
    {{- if eq .Name "BodyOfAgreementOffer" "BodyOfAgreementFormation" }}
        {{- $termsNeeded = true }}
    // Check depth or articles. Articles, Sections, Subsections, Paragraphs, Subparagraphs.
    terms := make(map[string]int)
    for _, chapter := range a.Chapters {
        terms = chapter.Terms(terms)
    }

    if err := checkTerms(a.Definitions, terms); err != nil {
        return err
    }
    {{- end }}
    {{ $fields := .Fields }}
    {{ range $i, $field := $fields }}
        {{- if ne $field.Type "deprecated" }}
            {{- $has_empty_check := false }}
            {{ range $j, $subfield := $fields }}
                {{- if ne $subfield.Type "deprecated" }}
                    {{- if $subfield.HasOnlyValidWhen }}
                        {{- if and (eq $subfield.OnlyValidWhen.FieldName $field.Name) (eq (len $subfield.OnlyValidWhen.Values) 0) }}
                            {{- $has_empty_check = true }}
                        {{- end }}
                    {{- end }}
                    {{- if $subfield.HasRequiredWhen }}
                        {{- if and (eq $subfield.RequiredWhen.FieldName $field.Name) (eq (len $subfield.RequiredWhen.Values) 0) }}
                            {{- $has_empty_check = true }}
                        {{- end }}
                    {{- end }}
                {{- end }}
            {{- end }}
            {{- if $has_empty_check }}
    {{ template "SetFieldIsEmpty" $field -}}
            {{- end }}
        {{- end }}
    {{- end }}
    {{ range $fields }}
        {{- if ne .Type "deprecated" }}
        {{ template "ValidateField" . -}}
        {{- end }}
    {{- end }}

    return nil
}
{{ end }}

{{- $clauseDepthNeeded := false }}

{{ range .FieldTypes }}
func (a *{{.Name}}Field) Validate() error {
    if a == nil {
        return nil
    }
    {{- if eq .Name "Chapter" }}
        {{- $clauseDepthNeeded = true }}
    // Check depth or articles. Articles, Sections, Subsections, Paragraphs, Subparagraphs.
    for i, article := range a.Articles {
        if article.Depth() > maxArticleDepth {
            return fmt.Errorf("Article %d over max depth : %d > %d", i, article.Depth(),
                maxArticleDepth)
        }
    }
    {{- end }}
    {{ range .Fields }}
        {{- if ne .Type "deprecated" }}
        {{ template "ValidateField" . -}}
        {{- end }}
    {{- end }}

    return nil
}
{{ end }}

// AddressIsValid returns true if an "Address" alias field is valid.
func AddressIsValid(b []byte) error {
    _, err := bitcoin.DecodeRawAddress(b)
    return err
}

// PublicKeyIsValid returns true if a "PublicKey" alias field is valid.
func PublicKeyIsValid(b []byte) error {
    _, err := bitcoin.PublicKeyFromBytes(b)
    return err
}

// SignatureIsValid returns true if a "Signature" alias field is valid.
func SignatureIsValid(b []byte) error {
    _, err := bitcoin.SignatureFromBytes(b)
    return err
}

{{- if $clauseDepthNeeded }}
func (c *ClauseField) Depth() int {
    if len(c.Children) == 0 {
        return 0 // no children
    }

    depth := 0
    for _, child := range c.Children {
        childDepth := child.Depth()
        if childDepth > depth {
            depth = childDepth
        }
    }

    return depth + 1
}
{{- end }}

{{- if $termsNeeded }}
// escape backslashes to get single backslashes in regex
var termRegEx = regexp.MustCompile("(?:\\[)(.+?)(?:\\]\\(\\))")

// findTerms adds any terms contained in the text to the map.
func findTerms(text string, terms map[string]int) map[string]int {
    matches := termRegEx.FindAllStringSubmatch(text, -1)
    for _, match := range matches {
        // index 1 for first regex group
        terms[match[1]] = 1
    }
    return terms
}

// Terms adds any terms in the chapter or its articles to the map.
func (c *ChapterField) Terms(terms map[string]int) map[string]int {
    terms = findTerms(c.Preamble, terms)

    for _, article := range c.Articles {
        terms = article.Terms(terms)
    }

    return terms
}

// Terms adds any terms in the clause or its children to the map.
func (c *ClauseField) Terms(terms map[string]int) map[string]int {
    terms = findTerms(c.Body, terms)

    for _, child := range c.Children {
        terms = child.Terms(terms)
    }

    return terms
}

// checkTerms returns an error if any of the defined terms are not found in the referenced terms
// map or if any of the referenced terms are not in the defined terms.
func checkTerms(defined []*DefinedTermField, referenced map[string]int) error {
    // Check that all referenced terms are defined.
    var undefined []string
    for term, _ := range referenced {
        found := false
        for _, defined := range defined {
            if term == defined.Term {
                found = true
                break
            }
        }

        if !found {
            undefined = append(undefined, term)
        }
    }

    if len(undefined) > 0 {
        return fmt.Errorf("Undefined terms : %v", undefined)
    }

    // Check that all defined terms are referenced
    var unreferenced []string
    for _, defined := range defined {
        if _, exists := referenced[defined.Term]; !exists {
            unreferenced = append(unreferenced, defined.Term)
        }
    }

    if len(unreferenced) > 0 {
        return fmt.Errorf("Unreferenced defined terms : %v", unreferenced)
    }

    return nil
}
{{- end }}
