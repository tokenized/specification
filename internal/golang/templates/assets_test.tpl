package {{ .Package }}

import (
    "testing"
)


func TestEmptyDeserialize(t *testing.T) {
    var ok bool
{{- range .Messages }}
	// {{.Name}} identifies a payload as a {{.Name}} message.
	asset{{.Name}}, err := Deserialize([]byte(Code{{.Name}}), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for {{.Name}} : %s", err)
	}
	_, ok = asset{{.Name}}.(*{{.Name}})
	if !ok {
		t.Fatalf("Failed deserialize type check for {{.Name}}")
	}
{{ end }}
}
