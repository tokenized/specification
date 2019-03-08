// Package {{.Package}} provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.
package protocol

const (
{{- range .ProtocolMessages}}
{{.CodeNameComment}}
	{{.CodeName}} = "{{.Code}}"
{{ end -}}
)

// TypeMapping holds a mapping of message codes to message types.
var (
	TypeMapping = map[string]OpReturnMessage{
{{- range .ProtocolMessages }}
		Code{{.Name}}: &{{.Name}}{},
{{- end }}
	}

	// ProtocolID is the current protocol ID
	ProtocolID = []byte("tokenized.com")
)
