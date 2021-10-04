package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const logoutLong = `
Remove credentials stored for a remote Terrakube server.
`

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from Terrakube server",
	Long:  logoutLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout ok")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
