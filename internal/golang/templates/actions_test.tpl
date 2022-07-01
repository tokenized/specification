package {{ .Package }}

import (
    "testing"
)


func TestEmptyDeserialize(t *testing.T) {
    var ok bool
{{- range .Messages }}
	// {{.Name}} identifies a payload as a {{.Name}} message.
	action{{.Name}}ScriptItems := bitcoin.ScriptItems{
		bitcoin.NewPushDataScriptItem([]byte(Code{{.Name}})),
	}
	action{{.Name}}, err := DeserializeV1(action{{.Name}}ScriptItems)
	if err != nil {
		t.Fatalf("Failed deserialize for {{.Name}} : %s", err)
	}
	_, ok = action{{.Name}}.(*{{.Name}})
	if !ok {
		t.Fatalf("Failed deserialize type check for {{.Name}}")
	}
{{ end }}
}
