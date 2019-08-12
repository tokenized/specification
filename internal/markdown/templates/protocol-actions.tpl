{{$actions := .Messages -}}

# Protocol Actions

- [Introduction](#introduction)
- [Header Fields](#header-fields)
- [Available Actions](#all-actions)

<a name="introduction"></a>
## Introduction

The following actions break down the construction of a Tokenized protocol message. The action is constructed by building a single string from each of the elements in order. Each field within the action is given a specific type, including standard types and compound types.

See the [Transactions article](../concepts/transactions) for details on how to construct a complete transaction.

<a name="header-fields"></a>
## Header Fields

Every protocol action is prepended with a header that specifies the necessary details about the subsequent action contents.

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>Version</td>
        <td>
            uint(1)
        </td>
        <td>
            The version number that determines the structure of the action payload and field definitions.
        </td>
    </tr>
    <tr>
        <td>ActionCode</td>
        <td>
            fixedchar(2)
        </td>
        <td>
            The Action Code that determines the expected contents of the action payload. Example: C1
        </td>
    </tr>
</table>

<a name="all-actions"></a>
## Available Actions

<div class="content-list collection-method-list" markdown="1">
{{- range $actions }}
- [{{.Metadata.Label}}](#{{.URLCode}})
{{- end }}
</div>

{{- range $actions}}

<a name="{{.URLCode}}"></a>
#### {{.Metadata.Label}}

{{.Metadata.Description}}

<table>
    <tr>
        <th style="width:15%">Action Code</th>
        <td>{{ .Code }}</td>
    </tr>
</table>


<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range .Fields}}
    <tr>
        <td>{{.Name}}</td>
        <td>
        {{- if .IsArrayType }}
          {{- if .IsResource }}
            <a href="resources#{{.TypeURLCode}}">{{.SingularDisplayType}}[{{.Length}}]</a>
          {{- else if .IsComplexType }}
            <a href="field-types#{{.TypeURLCode}}">{{.SingularDisplayType}}[{{.Length}}]</a>
          {{- else}}
            {{.SingularDisplayType}}[{{.Length}}]
          {{- end}}
        {{- else}}
          {{- if .IsResource }}
            <a href="resources#{{.TypeURLCode}}">{{.Type}}</a>{{ if ne .Length 0 }}({{.Length}}){{ end }}
          {{- else if .IsComplexType }}
            <a href="field-types#{{.TypeURLCode}}">{{.Type}}</a>{{ if ne .Length 0 }}({{.Length}}){{ end }}
          {{- else}}
            {{.Type}}{{ if ne .Length 0 }}({{.Length}}){{ end }}
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

##### Transaction Summary

{{if .Rules.Inputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Input</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .Rules.Inputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

{{if .Rules.Outputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Output</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .Rules.Outputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

{{- if ne .Rules.Fee 0 }}
> **Note**: This action requires a Smart Contract Operator Fee of at least <strong>{{.Rules.Fee}} satoshis</strong>
{{- end }}

<hr />

{{end}}
