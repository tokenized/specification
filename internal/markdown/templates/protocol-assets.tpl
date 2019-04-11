{{$assetTypes := . -}}

# Protocol Assets

- [Introduction](#introduction)
- [Available Actions](#all-assets)

<a name="introduction"></a>
## Introduction

Asset Types are used with reference to the `AssetPayload` field found in the Asset Definition, Asset Creation and Asset Modification actions.

<a name="all-assets"></a>
## Available Assets

<div class="content-list collection-method-list" markdown="1">
{{- range $assetTypes }}
- [{{.Metadata.Label}}](#{{.URLCode}})
{{- end }}
</div>


{{- range $assetTypes}}

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
        {{- if .IsComplexType }}
            <a href="#{{.TypeURLCode}}">{{.Type}}</a>
        {{- else}}
            {{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}
        {{- end}}
        </td>
        <td>{{.Description}} {{.Notes}}</td>
    </tr>
    {{- end}}
</table>

{{ end }}
