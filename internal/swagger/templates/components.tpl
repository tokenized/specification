components:
  tokenized:
{{- range .Messages }}
    {{ .Name }}:
      $ref: ./_components/schemas/tokenized/{{ $.Package }}/{{ .Name }}.yaml
{{- end }}
{{- range .FieldTypes }}
    {{ .Name }}Field:
      $ref: ./_components/schemas/tokenized/{{ $.Package }}/{{ .Name }}Field.yaml
{{- end }}
