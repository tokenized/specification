{{$messages := .Messages -}}

{{ define "ListValues" -}}
{{- range $i, $v := . -}}
{{ if gt $i 0 }}, {{ end }}{{ $v }}
{{- end -}}
{{- end }}

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
    {{- end}}
</table>

{{ end }}
