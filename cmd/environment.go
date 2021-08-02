package cmd

import (
	"github.com/spf13/cobra"
)

var environmentLong = `
This command consists of multiple subcommands to interact with environment variables.
It can be used to create, update, delete and list environment variables within a workspace.
`

var environmentCmd = &cobra.Command{
	Use:   "environment create|update|delete|list [ARGS]",
	Short: "create, update, delete and list environment variables",
	Long:  environmentLong,
}

func init() {
	workspaceCmd.AddCommand(environmentCmd)
}
