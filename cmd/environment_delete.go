package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EnvironmentDeleteExample string = `Delete a environment variable
    %[1]v workspace environment delete --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb  -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd `

var EnvironmentDeleteId string
var EnvironmentDeleteOrgId string
var EnvironmentDeleteWorkspaceId string

var deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		deleteEnvironment()
	},
	Example: fmt.Sprintf(EnvironmentDeleteExample, rootCmd.Use),
}

func init() {
	environmentCmd.AddCommand(deleteEnvironmentCmd)
	deleteEnvironmentCmd.Flags().StringVarP(&EnvironmentDeleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteEnvironmentCmd.MarkFlagRequired("organization-id")
	deleteEnvironmentCmd.Flags().StringVarP(&EnvironmentDeleteId, "id", "", "", "Id of the variable (required)")
	_ = deleteEnvironmentCmd.MarkFlagRequired("id")
	deleteEnvironmentCmd.Flags().StringVarP(&EnvironmentDeleteWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = deleteEnvironmentCmd.MarkFlagRequired("workspace-id")
}

func deleteEnvironment() {
	client := newClient()

	err := client.Variable.Delete(VariableDeleteOrgId, VariableDeleteWorkspaceId, VariableDeleteId)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("deleted")
}
