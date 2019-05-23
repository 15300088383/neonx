package service

import "github.com/spf13/cobra"

func CmdService() *cobra.Command {

	_bbf := &cobra.Command{
		Use:   "service",
		Short: "service层相关命令",
		SuggestFor: []string{
			"i",
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	_bbf.AddCommand(
		cmdGenerate(),
	)

	return _bbf
}
