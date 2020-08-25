{{$assets := .Messages -}}
{{$fieldTypes := .FieldTypes -}}
{{$fieldAliases := .FieldAliases -}}

{{ define "ListValues" -}}
{{- range $i, $v := . -}}
{{ if gt $i 0 }}, {{ end }}{{ $v }}
{{- end -}}
{{- end }}

{{- define "render_field"}}
    <tr>
    {{- if eq .Type "deprecated" }}
        <td>(Deprecated){{.Name}}</td>
        <td>deprecated</td>
    {{- else }}
        <td>{{.Name}}</td>
        <td>
        {{- if .IsList }}
          {{- if .IsAlias }}
            <a href="#alias-{{kebabcase .BaseType}}">{{.BaseType}}[{{.Size}}]</a>
          {{- else if .IsCompoundType }}
            <a href="#type-{{kebabcase .BaseType}}">{{.BaseType}}[{{.Size}}]</a>
          {{- else}}
            {{.MarkdownType}}
          {{- end}}
        {{- else}}
          {{- if .IsAlias }}
            <a href="#alias-{{kebabcase .BaseType}}">{{.Type}}</a>{{ if ne .Size 0 }}({{.Size}}){{ end }}
          {{- else if .IsCompoundType }}
            <a href="#type-{{kebabcase .BaseType}}">{{.Type}}</a>{{ if ne .Size 0 }}({{.Size}}){{ end }}
          {{- else}}
            {{.MarkdownType}}
          {{- end}}
        {{- end}}
        </td>
    {{- end }}
        <td>
            {{.Description}}
            {{.Notes}}
            {{- if .Required }} This field is always required. {{ end }}
            {{- if .HasRequiredWhen }}
                {{- if (eq (len .RequiredWhen.Values) 0) }}
            This field is required when the field {{ .RequiredWhen.FieldName }} is specified. 
                {{- else if (eq (len .RequiredWhen.Values) 1) }}
            This field is required when the field {{ .RequiredWhen.FieldName }} equals {{ template "ListValues" .RequiredWhen.Values }}. 
                {{- else }}
            This field is required when the field {{ .RequiredWhen.FieldName }} is within the values {{ template "ListValues" .RequiredWhen.Values }}. 
                {{- end }}
            {{- end }}
            {{- if .HasOnlyValidWhen }}
                {{- if (eq (len .OnlyValidWhen.Values) 0) }}
            This field is only valid when the field {{ .OnlyValidWhen.FieldName }} is specified. 
                {{- else if (eq (len .OnlyValidWhen.Values) 1) }}
            This field is only valid when the field {{ .OnlyValidWhen.FieldName }} equals {{ template "ListValues" .OnlyValidWhen.Values }}. 
                {{- else }}
            This field is only valid when the field {{ .OnlyValidWhen.FieldName }} is within the values {{ template "ListValues" .OnlyValidWhen.Values }}. 
                {{- end }}
            {{- end }}
            {{- if .Example }} Example: {{.Example}}{{ end }}
        </td>
    </tr>
{{end -}}

# Protocol Assets

- [Introduction](#introduction)
- [Available Assets](#all-assets)
- [Field Types](#field-types)
- [Field Aliases](#field-aliases)

<a name="introduction"></a>
## Introduction

Asset Types are used with reference to the `AssetPayload` field found in the Asset Definition, Asset Creation and Asset Modification actions.

<a name="all-assets"></a>
## Available Assets

<div class="content-list collection-method-list" markdown="1">
{{- range $assets }}
- [{{.Label}}](#{{kebabcase .Name}})
{{- end }}
</div>


{{- range $assets}}

<a name="{{kebabcase .Name}}"></a>
#### {{.Label}}

{{.Description}}

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range .Fields}}
        {{- template "render_field" . -}}
    {{- end}}
</table>

{{ end }}

<a name="field-types"></a>
## Field Types

<div class="content-list collection-method-list" markdown="1">
{{- range $fieldTypes }}
- [{{.Label}}](#type-{{kebabcase .Name}})
{{- end }}
</div>

{{range $fieldTypes}}

<a name="type-{{kebabcase .Name}}"></a>
### {{.Label}}

{{.Description}}

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range .Fields}}
        {{- template "render_field" . -}}
    {{- end}}
</table>

{{end}}

<a name="field-aliases"></a>
## Field Aliases

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range $fieldAliases}}
        <tr id="alias-{{kebabcase .Name}}">
            <td>{{.Name}}</td>
            <td>
            {{- if .IsList }}
              {{- if .IsAlias }}
                <a href="#alias-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
              {{- else if .IsCompoundType }}
                <a href="#type-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
              {{- else}}
                {{.MarkdownType}}
              {{- end}}
            {{- else}}
              {{- if .IsAlias }}
                <a href="#alias-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
              {{- else if .IsCompoundType }}
                <a href="#type-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
              {{- else}}
                {{.MarkdownType}}
              {{- end}}
            {{- end}}
            </td>
            <td>
                {{.Description}}
                {{.Notes}}
                {{- if .Example }} Example: {{.Example}}{{ end }}
            </td>
        </tr>
    {{- end}}
</table>
