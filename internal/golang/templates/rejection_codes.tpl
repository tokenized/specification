package protocol

// ------------------------------------------------------------------------------------------------
//                                       DO NOT EDIT
// This file is now generated from "resources/develop/Rejections.yaml"
// ------------------------------------------------------------------------------------------------

const (
{{- range $code, $value := .Values}}
	// {{$value.Name}} - {{$value.Description}}
	Reject{{$value.Name}} = {{ $code }}
{{ end }}
)