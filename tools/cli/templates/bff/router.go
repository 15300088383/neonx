package bff

import "html/template"

var Router = template.Must(
	template.New("").Parse(`package router

import (
	"{{.PrePath}}/impls"

	"github.com/geekymedic/neonx/bff"
)

var (
	engine = bff.Engine()
	group  = engine.Group("/{{.SystemName}}/v1")
)

func init() {
	group.POST("/{{.Name}}", bff.HttpHandler(impls.{{.CamelName}}Handler))
}`))
