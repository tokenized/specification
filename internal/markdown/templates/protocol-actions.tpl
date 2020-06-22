{{$actions := .Messages -}}
{{$fieldTypes := .FieldTypes -}}
{{$fieldAliases := .FieldAliases -}}

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
            <a href="#alias-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
          {{- else if .IsCompoundType }}
            <a href="#type-{{kebabcase .BaseType}}">{{.MarkdownType}}</a>
          {{- else}}
            {{.MarkdownType}}
          {{- end}}
        {{- else}}
          {{- if .IsAlias }}
            <a href="#alias-{{kebabcase .BaseTypeRaw}}">{{.MarkdownType}}</a>
          {{- else if .IsCompoundType }}
            <a href="#type-{{kebabcase .BaseTypeRaw}}">{{.MarkdownType}}</a>
          {{- else}}
            {{.MarkdownType}}
          {{- end}}
        {{- end}}
        </td>
    {{- end }}
        <td>
            {{.Description}}
            {{.Notes}}
            {{- if .Example }} Example: {{.Example}}{{ end }}
        </td>
    </tr>
{{end -}}

# Protocol Actions

- [Introduction](#introduction)
- [Available Actions](#all-actions)
- [Field Types](#field-types)
- [Field Aliases](#field-aliases)

<a name="introduction"></a>
## Introduction

The following actions break down the construction of a Tokenized protocol message. The action is constructed by building a single string from each of the elements in order. Each field within the action is given a specific type, including standard types and compound types.

See the [Transactions article](../concepts/transactions) for details on how to construct a complete transaction.

<a name="all-actions"></a>
## Available Actions

<div class="content-list collection-method-list" markdown="1">
{{- range $actions }}
- [{{.Label}}](#action-{{kebabcase .Name}})
{{- end }}
</div>

{{- range $actions}}

<a name="action-{{kebabcase .Name}}"></a>
#### {{.Label}}

{{.Description}}

<table>
    <tr>
        <th style="width:15%">Action Code</th>
        <td>{{.Code}}</td>
    </tr>
</table>


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

##### Transaction Summary

{{if .MetaData.Inputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Input</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .MetaData.Inputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

{{if .MetaData.Outputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Output</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .MetaData.Outputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

<hr />

{{end}}

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

