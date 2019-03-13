package protocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

{{range .}}
func Test{{.Name}}(t *testing.T) {
	// The hex is the body of the message
	body := "{{.Hex}}"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewEmpty{{.Name}}()

	header, err := NewHeaderForCode([]byte({{.CodeName}}), len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}
{{end}}
