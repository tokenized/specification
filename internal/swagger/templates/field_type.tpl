
{{- define "FieldType" -}}
  {{- if .IsList }}
    type: array
  {{- else if eq .BaseType "bin" "varbin" }}
    type: string
  {{- else if eq .BaseType "fixedchar" "varchar" }}
    type: string
  {{- else if eq .BaseType "uint" }}
    type: number
  {{- else if eq .BaseType "bool" }}
    type: boolean
  {{- else }}
    type: object
  {{- end -}}
{{- end -}}

{{- define "FieldRef" -}}
  {{- if .IsList }}
    items:
    {{- if eq .BaseType "fixedchar" "bin" }}
      type: string
    {{- else if eq .BaseType "varbin" "varchar" }}
      type: string
    {{- else if eq .BaseType "uint" }}
      type: number
    {{- else if eq .BaseType "bool" }}
      type: boolean
    {{- else if not .IsPrimitive }}
      type: object
      $ref: "#/components/schemas/{{ .BaseType }}Field"
    {{- end }}
  {{- else if eq .BaseType "bin" "varbin" -}}
  {{- else if eq .BaseType "fixedchar" "varchar" -}}
  {{- else if eq .BaseType "uint" -}}
  {{- else if eq .BaseType "bool" -}}
  {{- else }}
    $ref: "#/components/schemas/{{ .BaseType }}Field"
  {{- end -}}
{{- end -}}

{{- range .FieldTypes }}
description: {{ yamldescription .Description "  " }}
type: object
properties:
  {{- range .Fields }}
    {{- if ne .Type "deprecated" }}
  {{ .Name }}:
      {{- if gt (len .Description) 0 }}
    description: {{ yamldescription .Description "      " }}
      {{- end }}
      {{- template "FieldType" . }}
      {{- template "FieldRef" . }}
      {{- if gt (len .Example) 0 }}
    example: {{ .Example }}
      {{- end -}}
    {{- end }}
  {{- end }}
{{ end }}
