package protocol

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

{{range .Models}}
// {{.FormName}} is a JSON friendly model for an asset type.
type {{.FormName}} struct {
	{{range .Fields }}
	{{.Name}} {{.FormType}} {{.Tags}}{{end}}
}

// New{{.FormName}} returns a new {{.FormName}}.
func New{{.FormName}}() *{{.FormName}} {
	return &{{.FormName}}{}
}

// Validate returns an error if validation fails, nil otherwise.
func (f {{.FormName}}) Validate() error {
	return nil
}

func (f {{.FormName}}) pad(b []byte, l int) []byte {
	if len(b) == l {
		return b
	}

	padding := []byte{}
	c := l - len(b)

	for i := 0; i < c; i++ {
		padding = append(padding, 0)
	}

	return append(b, padding...)
}

// write writes the value to the buffer.
func (f {{.FormName}}) write(buf *bytes.Buffer, v interface{}) error {

	return binary.Write(buf, binary.BigEndian, v)
}

// writeBytes writes a string of fixed length to the buffer. If the string
// is longer that the length an error will be returned.
func (f {{.FormName}}) writeBytes(buf *bytes.Buffer, s string, l int) error {

	if len(s) > l {
		return fmt.Errorf("length exceeds %v", l)
	}

	b := f.pad([]byte(s), l)

	return f.write(buf, b)
}

// Bytes returns the form as a []byte that can be read by a protocol message.
func (f {{.FormName}}) Bytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	{{range .Fields }}
	{{.ToBuffer}}
	{{end}}
	return buf.Bytes(), nil
}

// PayloadMessage returns a PayloadMessage from the form, if any.
func (f {{.FormName}}) PayloadMessage(code []byte) (PayloadMessage, error) {
	m, err := NewPayloadMessageFromCode(code)
	if err != nil {
		return nil, err
	}

	if m != nil {
		return m, nil
	}

	return nil, errors.New("Not implemented")
}
{{end}}
