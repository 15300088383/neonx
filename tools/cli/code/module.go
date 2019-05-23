package code

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	goSuffix    = ".go"
	separator   = string(filepath.Separator)
	githubGroup = "github.com/geekymedic/"
	routerFile  = "router.go"
)

func useGithubGroup(use bool) string {

	if use {
		return githubGroup
	}

	return ""
}

var (
	_skips = map[string]bool{
		"code.go":  true,
		routerFile: true,
	}
)

type Module struct {
	System     string
	Version    string
	Module     string
	interfaces []*Interface
}

func NewModule(relative string) (*Module, error) {

	splits := strings.Split(relative, "/")

	if len(splits) != 3 {
		return nil, errors.Errorf("relative [%s] format error, use [<bff>/<version>/<module>]", relative)
	}

	return &Module{
		System:     splits[0],
		Version:    splits[1],
		Module:     splits[2],
		interfaces: make([]*Interface, 0),
	}, nil
}

func (m *Module) GroupName() string {
	return strings.Join([]string{m.System, m.Version, m.Module}, "/")
}

func (m *Module) Dir(base string) string {
	return filepath.Join(base, m.System, m.Version, m.Module)
}

func (m *Module) FileName(base, file string) string {
	return filepath.Join(m.Dir(base), file)
}

func (m *Module) load(base string) (err error) {

	var (
		dir = m.Dir(base)
	)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return
	}

	m.interfaces = make([]*Interface, 0)

	for _, file := range files {

		_, skip := _skips[file.Name()]

		if skip {
			continue
		}

		m.interfaces = append(
			m.interfaces,
			NewInterface(m, strings.TrimSuffix(file.Name(), goSuffix)))

	}

	return
}

func LoadModule(dir, stop string) (base string, module *Module, err error) {

	dir, err = filepath.Abs(dir)

	if err != nil {
		return
	}

	var (
		points = strings.Split(dir, separator)
	)

	for index, point := range points {
		if point == stop {
			base = separator + filepath.Join(points[0:index+1]...)
			points = points[index+1:]
		}
	}

	module, err = NewModule(strings.Join(points, "/"))

	if err != nil {
		return
	}

	err = module.load(base)

	return
}

func (m *Module) GenerateRouter(base string, useGithub bool) error {

	var (
		fileName = m.FileName(base, routerFile)
	)

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	return GenerateRouter(file, m, useGithub)

}

func (m *Module) GenerateInterface(base, name string, useGithub bool) (in *Interface, err error) {

	in = NewInterface(m, name)
	err = m.load(base)

	fileName := m.FileName(base, in.FileName())

	if err != nil {
		return nil, err
	}

	_, err = os.Stat(fileName)

	if err == nil || os.IsExist(err) {
		return nil, errors.Errorf("file [%s] already exists.", fileName)
	}

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = in.Generate(file, useGithub)

	if err != nil {
		return
	}

	m.interfaces = append(m.interfaces, in)

	return in, m.GenerateRouter(base, useGithub)

}
