package {{ .Package }}

import (
    "testing"
)


func TestEmptyDeserialize(t *testing.T) {
    var ok bool
{{- range .Messages }}
	// {{.Name}} identifies a payload as a {{.Name}} message.
	action{{.Name}}, err := Deserialize([]byte(Code{{.Name}}), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for {{.Name}} : %s", err)
	}
	_, ok = action{{.Name}}.(*{{.Name}})
	if !ok {
		t.Fatalf("Failed deserialize type check for {{.Name}}")
	}
    err = action{{.Name}}.Validate()
    if err != nil && !strings.Contains(err.Error(), "not within options") {
        // default values aren't always valid 
		t.Fatalf("Failed validate for {{.Name}} : %s", err)
    }
{{ end }}
}
