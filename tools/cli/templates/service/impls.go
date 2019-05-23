package service

import "html/template"

var Impls = template.Must(
	template.New("").Parse(`package impls

import (
	{{.SystemName}} "protocol/{{.SystemName}}_system"

	"golang.org/x/net/context"
)

func (m *{{.CamelSystemName}}Server) {{.CamelName}}(ctx context.Context, in *{{.SystemName}}.{{.CamelName}}Request) (*{{.SystemName}}.{{.CamelName}}Response, error) {
	return &{{.SystemName}}.{{.CamelName}}Response{}, nil
}`))
