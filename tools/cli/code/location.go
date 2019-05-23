package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type LocationEx struct {
	System    string
	Version   string
	Module    string
	Interface string
}

func (m *LocationEx) String() string {
	return strings.Join([]string{
		m.System,
		m.Version,
		m.Module,
		m.Interface,
	}, "/")
}

const (
	StopImpls = "impls"
)

type FileLocation struct {
	LocationEx
	base     string
	fileName string
}

func (m *FileLocation) SetFileName(name string) {
	m.fileName = name
	m.Interface = strings.TrimSuffix(name, ".go")
}

func (m *FileLocation) SetInterface(name string) {

	m.fileName = name + ".go"
	m.Interface = name
}

func NewFileLocationFromCWD(separator string) (*FileLocation, error) {

	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return NewFileLocation(dir, separator)
}

func NewFileLocation(dir, separator string) (loc *FileLocation, err error) {

	dir, err = filepath.Abs(dir)

	if err != nil {
		return
	}

	loc = &FileLocation{}

	var (
		points = strings.Split(dir, string(filepath.Separator))
		index  = 0
		point  string
	)

	for index, point = range points {
		if point == separator {
			loc.base = filepath.Join(points[0 : index+1]...)
			points = points[index+1:]
			break
		}
	}

	if len(points) == 0 {
		return nil, errors.Errorf("separator [%s] not found.", separator)
	}

	switch len(points) {
	case 4:
		loc.SetFileName(points[3])
		fallthrough
	case 3:
		loc.Module = points[2]
		loc.Version = points[1]
		loc.System = points[0]
	default:
		return nil, errors.Errorf("file path [%s] is no a interface path.", dir)

	}

	if len(loc.fileName) == 0 {
		loc.SetInterface("unknown")
	}

	return loc, nil

}

func (m *FileLocation) GetLocation() *LocationEx {
	return &m.LocationEx
}

func (m *FileLocation) String() string {
	return filepath.Join(string(filepath.Separator), m.base, m.System, m.Version, m.Module, m.fileName)
}

func (m *FileLocation) FileName() string {
	return m.fileName
}

func (m *FileLocation) FullName() string {
	return m.String()
}

func (m *FileLocation) Dir() string {

	return filepath.Join(string(filepath.Separator), m.base, m.System, m.Version, m.Module)
}

func (m *FileLocation) IsFileExists() bool {

	var (
		file   = fmt.Sprint(m)
		_, err = os.Stat(file)
	)

	return err == nil || os.IsExist(err)

}

type Location interface {
}
