package service

import "html/template"

var Main = template.Must(
	template.New("").Parse(`package main

import (
	"fmt"

	"github.com/geekymedic/neonx/service"

	_ "{{.SystemName}}_system/services/{{.ServiceName}}/impls"
)

func main() {
	service.Main()
}
`))

// go build -o C:\workspace\go\bin\neonx.exe
