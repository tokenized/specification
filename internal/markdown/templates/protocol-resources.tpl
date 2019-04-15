{{$resourceTypes := . -}}

# Protocol Resources

- [Introduction](#introduction)
- [Available Resources](#all-resources)

<a name="introduction"></a>
## Introduction

Resources are used to define lists of values, like the definitions of the possible values for a type field.

<a name="all-resources"></a>
## Available Resources

<div class="content-list collection-method-list" markdown="1">
{{- range $resourceTypes }}
- [{{.Metadata.Name}}](#{{.URLCode}})
{{- end }}
</div>


{{- range $resourceTypes}}

<a name="{{.URLCode}}"></a>
#### {{.Metadata.Name}}

{{.Metadata.Description}}

[View Source File](https://github.com/tokenized/specification/blob/master/src/resources/develop/{{.Metadata.Name}}.yaml)

<div class="content-list collection-method-list" markdown="1">
{{- range .Values}}
- {{ if .Name }}{{.Name}}{{ else }}{{.Label}}{{ end }}
{{- end }}
</div>
{{ end }}
