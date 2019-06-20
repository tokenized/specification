package protocol

import (
	"encoding/hex"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

{{range .}}
func Test{{.Name}}(t *testing.T) {
	// Create a randomized object
	initialMessage := {{.Name}}{}

	{{- range $i, $field := .PayloadFields }}
	// {{ $field.FieldName }} ({{ $field.Type }})
	{{- if ne (len .IncludeIfTrue) 0 }}
	if initialMessage.{{ .IncludeIfTrue }} {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
	if {{ range $j, $include := .IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} initialMessage.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }} {
	{{- else }}
	{
	{{- end }}
		{{- if eq $field.Type "bool" }}
		initialMessage.{{ $field.FieldName }} = true
		{{- else if $field.IsVarChar }}
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
			{{- else if .IsInternalTypeArray }}
		for i := 0; i < 2; i++ {
			initialMessage.{{ $field.FieldName }} = append(initialMessage.{{ $field.FieldName }}, {{.SingularType}}{})
		}
		{{- else if .IsNativeTypeArray }}
		for i := 0; i < 5; i++ {
			var item {{.SingularType}}
			initialMessage.{{ $field.FieldName }} = append(initialMessage.{{ $field.FieldName }}, item)
		}
		{{- else if .IsInternalType }}
		initialMessage.{{ $field.FieldName }} = {{.FieldGoType}}{}
		{{- else }}
		// {{ $field.Type }} test not setup
		{{- end }}
	}
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

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	{{- range $i, $field := .PayloadFields }}
	// {{ $field.FieldName }} ({{ $field.Type }})
		{{- if $field.IsVarChar }}
	if initialMessage.{{ $field.FieldName }} != decodedMessage.{{ $field.FieldName }} {
		t.Errorf("{{ $field.FieldName }} doesn't match : %s != %s", initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }})
	}
		{{- else if $field.IsFixedChar }}
	if initialMessage.{{ $field.FieldName }} != decodedMessage.{{ $field.FieldName }} {
		t.Errorf("{{ $field.FieldName }} doesn't match : %s != %s", initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }})
	}
		{{- else if $field.IsBytes }}
	if !bytes.Equal(initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }}) {
		t.Errorf("{{ $field.FieldName }} doesn't match : %x != %x", initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }})
	}
		{{- else if eq $field.Type "Polity" }}
	if len(initialMessage.{{ $field.FieldName }}.Items) != len(decodedMessage.{{ $field.FieldName }}.Items) {
		t.Errorf("{{ $field.FieldName }} length doesn't match : %d != %d", initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }})
	}
	for i, value := range initialMessage.{{ $field.FieldName }}.Items {
		if !bytes.Equal(value[:], decodedMessage.{{ $field.FieldName }}.Items[i][:]) {
			t.Errorf("{{ $field.FieldName }}.Items[%d] doesn't match : %s != %s", i, string(value[:]), string(decodedMessage.{{ $field.FieldName }}.Items[i][:]))
		}
	}
		{{- else if .IsInternalTypeArray }}
	if len(initialMessage.{{ $field.FieldName }}) != len(decodedMessage.{{ $field.FieldName }}) {
		t.Errorf("{{ $field.FieldName }} lengths don't match : %d != %d", len(initialMessage.{{ $field.FieldName }}), len(decodedMessage.{{ $field.FieldName }}))
	}
		{{- else if .IsNativeTypeArray }}
	if len(initialMessage.{{ $field.FieldName }}) != len(decodedMessage.{{ $field.FieldName }}) {
		t.Errorf("{{ $field.FieldName }} lengths don't match : %d != %d", len(initialMessage.{{ $field.FieldName }}), len(decodedMessage.{{ $field.FieldName }}))
	}
	for i, value := range initialMessage.{{ $field.FieldName }} {
		if value != decodedMessage.{{ $field.FieldName }}[i] {
			t.Errorf("{{ $field.FieldName }} value %d doesn't match : %v != %v", i, value, decodedMessage.{{ $field.FieldName }}[i])
		}
	}
		{{- else if .IsNativeType }}
	if initialMessage.{{ $field.FieldName }} != decodedMessage.{{ $field.FieldName }} {
		t.Errorf("{{ $field.FieldName }} doesn't match : %v != %v", initialMessage.{{ $field.FieldName }}, decodedMessage.{{ $field.FieldName }})
	}
		{{- else }}
	// {{ $field.Type }} test compare not setup
		{{- end }}
	{{- end }}
}

{{- end}}
