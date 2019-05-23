package system

import (
	"os"

	"github.com/geekymedic/neonx/tools/cli/code/service"

	"github.com/geekymedic/neonx/tools/cli/code/bff"
)

func Generate(absPath string, name string, bffNames []string, serviceNames []string) error {
	if len(serviceNames) == 0 {
		serviceNames = append(serviceNames, name)
	}
	systemName := name + "_system"
	systemPath := absPath + "/" + systemName
	//创建主目录
	err := os.MkdirAll(systemPath, os.ModePerm)
	if err != nil {
		return err
	}
	//bff
	err = os.MkdirAll(systemPath+"/bff", os.ModePerm)
	if err != nil {
		return err
	}
	for _, bName := range bffNames {
		err = bff.GenerateBff(systemPath, name, bName, []string{})
		if err != nil {
			return err
		}
	}

	// services
	err = os.MkdirAll(systemPath+"/service", os.ModePerm)
	if err != nil {
		return err
	}
	for _, sName := range serviceNames {
		err = service.GenerateService(systemPath, name, sName, []string{})
		if err != nil {
			return err
		}
	}

	return nil
}
