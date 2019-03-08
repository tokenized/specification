package {{.Package}}

import (
	"fmt"
	"reflect"
	"testing"
)

{{range .ProtocolMessages}}
func Test{{.Name}}(t *testing.T) {
     // the hex is the body of the message
     body := "{{.Hex}}"

     b, err := hex.DecodeString(body)
     if err != nil {
          t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
     }

     // create a valid header for the body
     m := New{{.Name}}()

     header, err := NewHeaderForCode({{.CodeName}}, len(b))
     if err != nil {
          t.Fatal(err)
     }

     m.Header = *header

     headerBytes, err := m.Header.Serialize()
     if err != nil {
        t.Fatal(err)
     }

     // this is the target byte payload
     want := headerBytes
     want = append(want, b...)

     {{/*
     fmt.Printf("header : %v\n", header)
     fmt.Printf("body   : %v\n", b)
     fmt.Printf("total  : %v\n", want)
     */}}

     n, err := m.Write(want)
     if err != nil {
        t.Fatal(err)
     }

     if n != len(want) {
        t.Fatalf("got %v, want %v", n, len(want))
     }

	// serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}
{{end}}
