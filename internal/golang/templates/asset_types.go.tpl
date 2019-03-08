package protocol

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152
{{range .Models }}
	// Code{{.FullName}} identifies data as a {{.Name}} message.
	Code{{.FullName}} = "{{.Code}}"
{{end}})

{{range .Models}}
// {{.FullName}} asset type.
type {{.FullName}} struct {
	BaseMessage
	{{range .Fields }}
	{{.Name}} {{.GoType}}{{end}}
}

// New{{.FullName}} returns a new {{.FullName}}.
func New{{.FullName}}() *{{.FullName}} {
	return &{{.FullName}}{}
}

// Type returns the type identifer for this message.
func (m {{.FullName}}) Type() string {
	return Code{{.FullName}}
}

// Len returns the byte size of this message.
func (m {{.FullName}}) Len() int64 {
	return AssetTypeLen
}

// Bytes returns the message in bytes.
func (m {{.FullName}}) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	{{range .Fields }}
	{{.Bytes}}{{end}}

	return buf.Bytes(), nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *{{.FullName}}) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	{{range .Fields }}
	{{.Write}}{{end}}

	return int(m.Len()), nil
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m {{.FullName}}) Read(b []byte) (int, error) {
	data, err := m.Bytes()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}
{{end}}
