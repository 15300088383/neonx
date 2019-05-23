package system

import "github.com/spf13/cobra"

func CmdSystem() *cobra.Command {

	_system := &cobra.Command{
		Use:   "system",
		Short: "system相关命令",
		SuggestFor: []string{
			"i",
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	_system.AddCommand(
		cmdGenerate(),
	)

	return _system
}
