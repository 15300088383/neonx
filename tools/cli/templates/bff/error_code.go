package bff

import "html/template"

var Codes = template.Must(
	template.New("").Parse(`package codes

import "github.com/geekymedic/neonx/bff"

const (
	
)

var (
	_codes = bff.Codes{
	}
)

func GetMessage(code int) string {
	return _codes[code]
}

func init() {
	bff.MergeCodes(_codes)
}`))
