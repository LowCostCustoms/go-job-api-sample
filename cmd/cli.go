package cmd

import "github.com/spf13/cobra"

func NewCli() *cobra.Command {
	command := &cobra.Command{}
	command.AddCommand(
		NewServeCommand(),
		NewMigrateCommand(),
	)

	return command
}
