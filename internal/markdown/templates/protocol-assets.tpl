{{$assetTypes := . -}}

# Protocol Assets

- [Introduction](#introduction)
- [Available Assets](#all-assets)

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
