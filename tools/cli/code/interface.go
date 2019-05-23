package code

import (
	"io"

	"github.com/huandu/xstrings"
)

type Interface struct {
	module *Module
	Name   string
	Camel  string
}

func NewInterface(module *Module, name string) *Interface {
	return &Interface{
		module: module,
		Name:   name,
		Camel:  xstrings.ToCamelCase(name),
	}
}

func (m *Interface) FileName() string {
	return m.Name + goSuffix
}

func (m *Interface) Generate(w io.Writer, useGitHub bool) error {

	return tplInterface.Execute(
		w,
		map[string]interface{}{
			"Package": m.module.Module,
			"Name":    m.Camel,
			"GitHub":  useGithubGroup(useGitHub),
		})
}
