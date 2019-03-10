package protocol

import "bytes"

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152
{{range . }}
	// Code{{.Name}} identifies data as a {{.Name}} message.
	Code{{.Name}} = "{{.Code}}"
{{end}})
{{range .}}
// {{.Name}} asset type.
type {{.Name}} struct {
	BaseMessage
{{range .Fields }}	{{.Name}} {{.GoType}}
{{ end -}}
}

// New{{.Name}} returns a new {{.Name}}.
func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{}
}

// Type returns the type identifer for this message.
func (m {{.Name}}) Type() string {
	return Code{{.Name}}
}

// Len returns the byte size of this message.
func (m {{.Name}}) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m {{.Name}}) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)
{{range .Fields }}{{.Bytes}}
{{end}}
	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *{{.Name}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)
{{range .Fields }}{{.Write}}
{{end}}
	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m {{.Name}}) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}
{{end}}