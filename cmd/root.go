package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCommand returns a new instance of the root command
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "argocd-custom-plugin",
		Short: "This is a plugin to annotate namespaces with metadata",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewAnnotateCommand())

	return command
}
