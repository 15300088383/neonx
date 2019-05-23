package bff

import "github.com/spf13/cobra"

func CmdBFF() *cobra.Command {

	_bbf := &cobra.Command{
		Use:   "bff",
		Short: "BFF层相关命令",
		SuggestFor: []string{
			"i",
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	_bbf.AddCommand(
		cmdGenerate(),
		cmdFlushRouter(),
		cmdAddInterface(),
	)

	return _bbf
}
