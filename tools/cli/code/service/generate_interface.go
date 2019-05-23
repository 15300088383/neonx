package service

import (
	"os"

	"github.com/geekymedic/neonx/tools/cli/templates/service"
)

func GenerateInterface(absPath string, systemName string, interfaceName string) error {

	// add go file
	file, err := os.Create(absPath + "/impls/" + interfaceName + ".go")
	defer file.Close()
	if err != nil {
		return err
	}
	impls := service.Impls
	err = impls.Execute(file, map[string]interface{}{
		"Name":            interfaceName,
		"CamelName":       camelString(interfaceName),
		"SystemName":      systemName,
		"CamelSystemName": camelString(systemName),
	})
	if err != nil {
		return err
	}
	addStruct(absPath, systemName)

	return nil
}

func addStruct(absPath string, systemName string) error {
	filePth := absPath + "/impls/" + systemName + "_server.go"
	exist, err := pathExists(filePth)
	if exist {
		return nil
	}
	file, err := os.Create(filePth)
	defer file.Close()
	if err != nil {
		return err
	}
	impls := service.ServerStruct
	err = impls.Execute(file, map[string]interface{}{
		"SystemName":      systemName,
		"CamelSystemName": camelString(systemName),
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
