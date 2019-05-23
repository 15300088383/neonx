package service

import (
	"os"

	"github.com/geekymedic/neonx/tools/cli/templates/service"
)

func GenerateService(absPath string, systemName string, serviceName string, interfaceNames []string) error {
	servicePath := absPath + "/services/" + serviceName
	err := os.MkdirAll(servicePath, os.ModePerm)
	if err != nil {
		return err
	}

	// create config
	err = os.MkdirAll(servicePath+"/config", os.ModePerm)
	if err != nil {
		return err
	}
	file, error := os.Create(servicePath + "/config/config.yml")
	if error != nil {
		return err
	}
	configTpl := service.Config
	err = configTpl.Execute(file, map[string]interface{}{})
	if err != nil {
		return err
	}
	defer file.Close()

	// impls
	err = os.MkdirAll(servicePath+"/impls", os.ModePerm)
	if err != nil {
		return err
	}

	// main.go
	file, error = os.Create(servicePath + "/main.go")
	if error != nil {
		return err
	}
	mainTpl := service.Main
	err = mainTpl.Execute(file, map[string]interface{}{"SystemName": systemName, "ServiceName": serviceName})
	if err != nil {
		return err
	}
	defer file.Close()

	for _, name := range interfaceNames {
		err = GenerateInterface(servicePath, systemName, name)
		return err
	}

	return nil
}
