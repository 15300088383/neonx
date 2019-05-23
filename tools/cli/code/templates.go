package code

import "html/template"

var (
	textInterface = `package {{.Package}}

import (
	"{{.GitHub}}neonx/bff"

	"github.com/gin-gonic/gin"
)

type {{.Name}}Request struct {
	Name  string
	Value string
}

type {{.Name}}Response struct {
}

func {{.Name}}Handler(ctx *gin.Context) {

	var (
		state = bff.NewState(ctx)
		ask   = &{{.Name}}Request{}
		ack   = &{{.Name}}Response{}
	)

	if err := ctx.BindJSON(ask); err != nil {
		state.Error(bff.CodeRequestBodyError, err)
		return
	}

	state.Success(ack)

}

`
	tplInterface = template.Must(
		template.New("").Parse(textInterface),
	)
)

var (
	txtRouter = `package {{.Package}}

import (
	"{{.GitHub}}neonx/bff"
)

const (
	{{range .Interfaces}}Interface{{.Camel}} = "/{{$.Group}}/{{.Name}}"
	{{end}}
)

var (
	engine = bff.Engine()
	group  = engine.Group("/{{.Group}}")
)

func init() {
	{{range .Interfaces}}group.POST("/{{.Name}}", {{.Camel}}Handler)
	{{end}}
}`
	tplRouter = template.Must(
		template.New("").Parse(txtRouter))
)
