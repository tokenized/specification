package bsor

import (
	"reflect"

	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/{{ .Package }}"
)

{{- $package := .Package }}

func Generate{{ $package }}() (bsor.Definitions, error) {
	return bsor.BuildDefinitions(
{{- range .Messages }}
		reflect.TypeOf({{ $package }}.{{.Name}}{}),
{{- end }}
	)
}
