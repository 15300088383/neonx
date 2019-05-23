package system

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/geekymedic/neonx/tools/cli/code/system"
	"github.com/spf13/cobra"
)

func cmdGenerate() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "生成一个系统",
		Example: "micro system generate <name>",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) <= 0 {
				fmt.Println(
					"micro system generate <name>")
				cmd.Help()
				return
			}
			var err error
			defer func() {
				if err != nil {
					fmt.Println(err)
				}
			}()

			bffNames, err := cmd.Flags().GetStringSlice("bffNames")
			if err != nil {
				fmt.Println(
					"micro system generate <name>")
				cmd.Help()
				return
			}
			serviceNames, err := cmd.Flags().GetStringSlice("bffNames")
			if err != nil {
				fmt.Println(
					"micro system generate <name>")
				cmd.Help()
				return
			}
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				fmt.Println(
					"micro system generate <name>")
				cmd.Help()
				return
			}
			err = system.Generate(dir, args[0], bffNames, serviceNames)
		},
	}

	cmd.Flags().StringSliceP("bffNames", "b", []string{}, "name for bff,example:-b \"admin,app\",default nil ")
	cmd.Flags().StringSliceP("serviceNames", "s", []string{}, "name for bff,example:-s \"account,user\",default bff inl ")
	return cmd
}
