{{$assets := .Messages -}}
{{$fieldTypes := .FieldTypes -}}
{{$fieldAliases := .FieldAliases -}}

{{- define "render_field"}}
    <tr>
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
        <td>
            {{.Description}}
            {{.Notes}}
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
