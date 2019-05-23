package bff

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/geekymedic/neonx/tools/cli/templates/bff"
)

func GenerateInterface(absPath string, codePrePath string, systemName string, interfaceName string) error {

	// add go file
	file, err := os.Create(absPath + "/impls/" + interfaceName + ".go")
	defer file.Close()
	if err != nil {
		return err
	}
	impls := bff.Impls
	err = impls.Execute(file, map[string]interface{}{
		"Name":      interfaceName,
		"CamelName": camelString(interfaceName),
		"PrePath":   codePrePath,
	})
	if err != nil {
		return err
	}
	err = addToRouter(absPath, codePrePath, systemName, interfaceName)
	if err != nil {
		return err
	}
	return nil
}

func addToRouter(absPath string, prePath string, systemName string, interfaceName string) error {
	filePth := absPath + "/router/router.go"
	exist, err := pathExists(filePth)
	if exist {
		f, err := os.OpenFile(filePth, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		all, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		fileText := string(all)
		index := strings.LastIndex(fileText, "}")
		newText := fileText[0:index] + "\tgroup.POST(\"/" + interfaceName + "\", bff.HttpHandler(impls." + camelString(interfaceName) + "Handler))\n}"
		_, err = f.WriteString(newText)
		if err != nil {
			return err
		}
		defer f.Close()
		return nil
	}
	file, err := os.Create(filePth)
	defer file.Close()
	if err != nil {
		return err
	}
	impls := bff.Router
	err = impls.Execute(file, map[string]interface{}{
		"Name":       interfaceName,
		"CamelName":  camelString(interfaceName),
		"PrePath":    prePath,
		"SystemName": systemName,
	})
	if err != nil {
		return err
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
