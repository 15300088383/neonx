package bff

import "html/template"

var BffMain = template.Must(
	template.New("").Parse(`package main

import (
	"fmt"
	_ "{{.PrePath}}/router"

	"github.com/geekymedic/neonx/bff"
	_ "github.com/geekymedic/neonx/plugin/rpc"
)

func main() {
	err := bff.Main()
	if err != nil {
		fmt.Println(err)
	}
}`))
