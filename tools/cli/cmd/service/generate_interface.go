package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geekymedic/neonx/tools/cli/code/service"

	"github.com/spf13/cobra"
)

func cmdAddInterface() *cobra.Command {

	var (
		useGitHub = true
	)

	cmd := &cobra.Command{
		Use:     "add-interface",
		Short:   "生成一个service接口",
		Example: "micro service add-interface <name>",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) <= 0 {
				fmt.Println(
					"micro service add-interface <name>")
				cmd.Help()
				return
			}

			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			pathSlice := strings.Split(dir, string(filepath.Separator))
			systemName := pathSlice[len(pathSlice)-3]

			defer func() {
				if err != nil {
					fmt.Println(err)
				}
			}()
			err = service.GenerateInterface(dir, strings.Split(systemName, "_")[0], args[0])
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&useGitHub, "github", true, "bbf包是否使用github地址")

	return cmd
}
