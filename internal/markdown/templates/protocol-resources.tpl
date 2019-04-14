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
- [{{.Metadata.Name}}](#{{.Metadata.Type}})
{{- end }}
</div>


{{- range $resourceTypes}}

<a name="{{.Metadata.Type}}"></a>
#### {{.Metadata.Name}}

{{.Metadata.Description}}

Source: https://github.com/tokenized/specification/blob/master/src/resources/develop/{{.Metadata.Name}}.yaml

<div><button onclick="showHideYaml('{{.Metadata.Name}}Yaml')">Show/Hide {{.Metadata.Name}}</button></div>
<code id="{{.Metadata.Name}}Yaml">
{{ .Data }}
</code>

{{ end }}

<script type="text/javascript">
function showHideYaml(id) {
  var x = document.getElementById(id);
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}
</script>