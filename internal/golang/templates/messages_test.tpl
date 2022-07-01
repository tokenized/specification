package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"testing"
)


func TestEmptyDeserialize(t *testing.T) {
    var ok bool
{{- range .Messages }}
	// {{.Name}} identifies a payload as a {{.Name}} message.
	message{{.Name}}, err := DeserializeV1(Code{{.Name}}, nil)
	if err != nil {
		t.Fatalf("Failed deserialize for {{.Name}} : %s", err)
	}
	_, ok = message{{.Name}}.(*{{.Name}})
	if !ok {
		t.Fatalf("Failed deserialize type check for {{.Name}}")
	}
{{ end }}
}
