package main

import (
	"fmt"

	"github.com/geekymedic/neonx/tools/cli/cmd/service"

	"github.com/geekymedic/neonx/tools/cli/cmd/system"

	"github.com/geekymedic/neonx/tools/cli/cmd/bff"

	"github.com/spf13/cobra"
)

func main() {

	cmd := cobra.Command{
		Use: "micro",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		bff.CmdBFF(),
		system.CmdSystem(),
		service.CmdService(),
	)

	err := cmd.Execute()

	if err != nil {
		fmt.Print(err)
	}
}

// go build -o C:\workspace\go\bin\micro.exe
