package service

import "html/template"

var ServerStruct = template.Must(
	template.New("").Parse(`package impls

import (
	{{.SystemName}} "protocol/{{.SystemName}}_system"

	"github.com/geekymedic/neonx/service"
)

type {{.CamelSystemName}}Server struct {
}

func init() {
	{{.SystemName}}.Register{{.CamelSystemName}}Server(service.Server(), &{{.CamelSystemName}}Server{})
}`))
