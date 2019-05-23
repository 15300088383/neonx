package bff

import "html/template"

var Impls = template.Must(
	template.New("").Parse(`package impls

import (
"errors"
"{{.PrePath}}/rpc"

"github.com/geekymedic/neonx/utils/context"

user "protocol/user_system"

"github.com/geekymedic/neonx/bff"
)

type {{.CamelName}}Request struct {
}

type {{.CamelName}}Response struct {
}

func {{.CamelName}}Handler(state *bff.State) {
	var (
		ask = &{{.CamelName}}Request{}
		ack   = &{{.CamelName}}Response{}
	)
	if err := state.BindJSON(ask); err != nil {
		state.Error(bff.CodeRequestBodyError, err)
	}
	
	state.Success(ack)
}`))
