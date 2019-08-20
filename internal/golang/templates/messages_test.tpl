package {{ .Package }}

import (
	"bytes"
	"encoding/binary"
	"testing"
)


func codeBytes(code int) []byte {
    var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, uint16(code))
    return buf.Bytes()
}

func TestEmptyDeserialize(t *testing.T) {
    var ok bool
{{- range .Messages }}
	// {{.Name}} identifies a payload as a {{.Name}} message.
	message{{.Name}}, err := Deserialize(codeBytes(Code{{.Name}}), nil)
	if err != nil {
		t.Fatalf("Failed deserialize for {{.Name}} : %s", err)
	}
	_, ok = message{{.Name}}.(*{{.Name}})
	if !ok {
		t.Fatalf("Failed deserialize type check for {{.Name}}")
	}
{{ end }}
}
