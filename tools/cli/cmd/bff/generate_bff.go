package bff

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/geekymedic/neonx/tools/cli/code/bff"

	"github.com/spf13/cobra"
)

func cmdGenerate() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "生成一个bff系统",
		Example: "micro bff generate <name>",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) <= 0 {
				fmt.Println(
					"micro bff generate <name>")
				cmd.Help()
				return
			}
			var err error
			defer func() {
				if err != nil {
					fmt.Println(err)
				}
			}()
			interfaceNames, err := cmd.Flags().GetStringSlice("interfaceNames")
			if err != nil {
				fmt.Println(
					"micro bff generate <name>")
				cmd.Help()
				return
			}
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			pathSlice := strings.Split(dir, string(filepath.Separator))
			systemName := strings.Split(pathSlice[len(pathSlice)-1], "_")[0]
			err = bff.GenerateBff(dir, systemName, args[0], interfaceNames)
		},
	}

	cmd.Flags().StringSliceP("interfaceNames", "i", []string{}, "add interface for bff,example:-i \"get_user,get_phone\",default nil ")
	return cmd
}
