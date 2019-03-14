package protocol

import (
	"encoding/hex"
	"reflect"
	"testing"
	"time"
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
		{{- else if $field.IsBytes }}
		initialMessage.{{ $field.FieldName }} = make([]byte, 0, {{ $field.Length }})
		for i := uint64(0); i < {{ $field.Length }}; i++ {
			initialMessage.{{ $field.FieldName }} = append(initialMessage.{{ $field.FieldName }}, byte(65 + i + {{ $i }}))
		}
		{{- else if eq $field.Type "timestamp" }}
		initialMessage.{{ $field.FieldName }} = uint64(time.Now().UnixNano())
		{{- else }}
	// {{ $field.Type }} test not setup
		{{- end }}
	{{ end }}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := {{.Name}}{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}
{{end}}
