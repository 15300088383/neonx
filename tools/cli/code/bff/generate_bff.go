package bff

import (
	"os"

	"github.com/geekymedic/neonx/tools/cli/templates/bff"
)

func GenerateBff(absPath string, systemName string, bffName string, interfaceNames []string) error {
	bffPath := absPath + "/bff/" + bffName
	err := os.MkdirAll(bffPath, os.ModePerm)
	if err != nil {
		return err
	}
	// 创建codes
	err = os.MkdirAll(bffPath+"/codes", os.ModePerm)
	if err != nil {
		return err
	}
	file, error := os.Create(bffPath + "/codes/error_code.go")
	if error != nil {
		return err
	}
	codesTpl := bff.Codes
	err = codesTpl.Execute(file, map[string]interface{}{})
	if err != nil {
		return err
	}
	defer file.Close()

	// create config
	err = os.MkdirAll(bffPath+"/config", os.ModePerm)
	if err != nil {
		return err
	}
	file, error = os.Create(bffPath + "/config/config.yml")
	if error != nil {
		return err
	}
	configTpl := bff.Config
	err = configTpl.Execute(file, map[string]interface{}{})
	if err != nil {
		return err
	}
	defer file.Close()

	// impls
	err = os.MkdirAll(bffPath+"/impls", os.ModePerm)
	if err != nil {
		return err
	}

	// router
	err = os.MkdirAll(bffPath+"/router", os.ModePerm)
	if err != nil {
		return err
	}

	// rpc
	err = os.MkdirAll(bffPath+"/rpc", os.ModePerm)
	if err != nil {
		return err
	}

	// main.go
	file, error = os.Create(bffPath + "/main.go")
	if error != nil {
		return err
	}
	mainTpl := bff.BffMain
	err = mainTpl.Execute(file, map[string]interface{}{"PrePath": systemName + "_system/bff/" + bffName})
	if err != nil {
		return err
	}
	defer file.Close()

	for _, name := range interfaceNames {
		err = GenerateInterface(bffPath, systemName+"_system/bff/"+bffName, systemName, name)
		return err
	}

	return nil
}
