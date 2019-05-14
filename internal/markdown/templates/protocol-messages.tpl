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

{{ end }}
