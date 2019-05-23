package bff

import (
	"fmt"
	"os"

	"github.com/geekymedic/neonx/tools/cli/code"

	"github.com/spf13/cobra"
)

func cmdFlushRouter() *cobra.Command {

	var (
		useGitHub = true
	)

	cmd := &cobra.Command{
		Use:     "flush",
		Short:   "刷新路由表",
		Example: "micro bff flush",
		Run: func(cmd *cobra.Command, args []string) {

			var (
				wd, err = os.Getwd()
			)

			defer func() {
				if err != nil {
					fmt.Println(err)
				}
			}()

			if err != nil {
				return
			}

			base, module, err := code.LoadModule(wd, code.StopImpls)

			if err != nil {
				return
			}

			err = module.GenerateRouter(base, useGitHub)

		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&useGitHub, "github", true, "bbf包是否使用github地址")

	return cmd
}
