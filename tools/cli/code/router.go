package code

import (
	"io"
)

func GenerateRouter(w io.Writer, module *Module, useGithub bool) error {

	return tplRouter.Execute(
		w,
		map[string]interface{}{
			"Package":    module.Module,
			"GitHub":     useGithubGroup(useGithub),
			"Interfaces": module.interfaces,
			"Group":      module.GroupName(),
		})
}
