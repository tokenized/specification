package protocol

// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

// NewFormByCode returns a new ProtocolForm by code.
//
// An error will be returned if there is no matching Form.
func NewFormByCode(code string) (Form, error) {
{{range .Messages}}
	if code == Code{{.Name}} {
		return &{{.FormName}}{}, nil
	}
{{end}}

	return nil, ErrUnknownMessage
}

{{range .Messages}}
// {{.FormName}} is the JSON friendly version of a {{.Name}}.
type {{.FormName}} struct {
	BaseForm
	{{range .DataFields}}
	{{ .Name}} {{.FormType}} {{.Tags}}{{end}}
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (f *{{.FormName}}) Write(b []byte) (n int, err error) {
	if err := json.Unmarshal(b, f); err != nil {
		return 0, err
	}

	return len(b), nil
}

// Validate returns an error if validations fails.
func (f {{.FormName}}) Validate() error {
	return nil
}

// PayloadForm returns a PayloadForm for this Form, or nil if this Form has
// no matching PayloadForm.
func (f {{.FormName}}) PayloadForm() (PayloadForm, error) {
{{.PayloadFormPartial}}
}

// BuildMessage returns an OpReturnMessage from the form.
func (f {{.FormName}}) BuildMessage() (OpReturnMessage, error) {
	var err error

	m := New{{.Name}}()

	{{range .DataFields}}
	{{ .AssignFormValueToField }}
	{{end}}

	{{ if .HasPayloadMessage }}
		pf, err := f.PayloadForm()
	if err != nil {
		return nil, err
	}

	if pf == nil {
		return &m, nil
	}

	b, err := pf.Bytes()
	if err != nil {
		return nil, err
	}

	m.Payload = b
	{{end }}

	return &m, nil
}
{{end}}
