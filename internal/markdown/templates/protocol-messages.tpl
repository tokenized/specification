{{$messages := . -}}

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
- [{{.Metadata.Label}}](#{{.URLCode}})
{{- end }}
</div>

{{- range $messages}}

<a name="{{.URLCode}}"></a>
#### {{.Metadata.Label}}

{{.Metadata.Description}}

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
        {{- if .IsResource }}
            <a href="resources#{{.TypeURLCode}}">{{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}</a>
        {{- else if .IsComplexType }}
            <a href="field-types#{{.TypeURLCode}}">{{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}</a>
        {{- else}}
            {{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}
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
