package protocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

{{range .}}
func Test{{.Name}}(t *testing.T) {
	// Create a randomized object
	initialMessage := {{.Name}}{}

	{{- range $i, $field := .PayloadFields }}
	// {{ $field.FieldName }} ({{ $field.Type }})
		{{- if $field.IsVarChar }}
	initialMessage.{{ $field.FieldName }} = "Text {{ $i }}"
		{{- else if $field.IsFixedChar }}
	{
		text := make([]byte, 0, {{ $field.Length }})
		for i := uint64(0); i < {{ $field.Length }}; i++ {
			text = append(text, byte(65 + i + {{ $i }}))
		}
		initialMessage.{{ $field.FieldName }} = string(text)
	}
		{{- else }}
	// {{ $field.Type }} test not setup
		{{- end }}
	{{ end }}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := {{.Name}}{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}
}
{{end}}
