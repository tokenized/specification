{{$messages := .Messages -}}

# Protocol Messages

- [Introduction](#introduction)
- [Available Messages](#all-messages)

<a name="introduction"></a>
## Introduction

The Tokenized protocol features a complete messaging suite for all types of messaging including general purpose private and public messaging, as well as commercial, financial and legal messaging in accordance with a variety of established Electronic Data Interchange (EDI) standards. They are also used to allow smart contracts to share information and orchestrate multiple signature transactions, such as atomic swaps.

<a name="all-messages"></a>
## Available Messages

<div class="content-list collection-method-list" markdown="1">
{{- range $messages }}
- [{{.Label}}](#{{kebabcase .Name}})
{{- end }}
</div>

{{- range $messages}}

<a name="{{kebabcase .Name}}"></a>
#### {{.Label}}

{{.Description}}

<table>
    <tr>
        <th style="width:15%">Message Code</th>
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
            {{- if .IsList }}
              {{- if .IsAlias }}
                <a href="#alias-{{kebabcase .BaseType}}">{{.BaseType}}[{{.Size}}]</a>
              {{- else if .IsCompoundType }}
                <a href="#type-{{kebabcase .BaseType}}">{{.BaseType}}[{{.Size}}]</a>
              {{- else}}
                {{.BaseType}}[{{.Size}}]
              {{- end}}
            {{- else}}
              {{- if .IsAlias }}
                <a href="#alias-{{kebabcase .BaseType}}">{{.Type}}</a>{{ if ne .Size 0 }}({{.Size}}){{ end }}
              {{- else if .IsCompoundType }}
                <a href="#type-{{kebabcase .BaseType}}">{{.Type}}</a>{{ if ne .Size 0 }}({{.Size}}){{ end }}
              {{- else}}
                {{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}
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

{{ end }}
