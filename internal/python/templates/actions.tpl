# Package protocol provides base level structs and validation for
# the protocol.
#
# The code in this file is auto-generated. Do not edit it by hand as it will
# be overwritten when code is regenerated.

{{ range . }}
{{comment .Metadata.Description "#"}}

class Action_{{.Name}}({{ range .Fields }}{{ if eq .Name "Version" }}Versioned{{ end }}{{ end }}ActionBase):
    ActionPrefix = '{{.Code}}'
{{ range .Fields }}{{ if eq .Name "Version" }}    ActiveVersion = 0
{{ end }}{{ end }}
    schema = {
        {{ range $i, $v := .Fields }}{{ if gt $i 1 }}{{ if gt $i 2 }},
        {{ end }}'{{$v.Name}}': {{ padding $v.Name 30 }} [{{ minus $i 2 }}, DAT_{{$v.Type}}, {{$v.Length}}]{{ end }}{{ end }}
    }

    rules = {
        'contractFee': {{.Rules.Fee}},
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
{{ range $i, $v := .Fields }}{{ if gt $i 2 }}        {{ if eq $v.Name "Version" }}self.Version = self.ActiveVersion{{ else }}self.{{$v.Name}} = None{{ end }}
{{ end }}{{ end }}
{{ end }}

ActionClassMap = {
    {{ range $i, $v := . }}{{ if $i }},
    {{ end }}'{{$v.Code}}': Action_{{$v.Name}}{{ end }}
}
