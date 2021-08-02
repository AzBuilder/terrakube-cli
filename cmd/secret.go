package cmd

import (
	"github.com/spf13/cobra"
)

var secretLong = `
This command consists of multiple subcommands to interact with secrets.
It can be used to create, update, delete and list secrets within a workspace.
`

var secretCmd = &cobra.Command{
	Use:   "secret create|update|delete|list [ARGS]",
	Short: "create, update, delete and list secrets",
	Long:  secretLong,
}

func init() {
	workspaceCmd.AddCommand(secretCmd)
}
