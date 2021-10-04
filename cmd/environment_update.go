package cmd

import (
	"terrakube/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var EnvironmentUpdateExample string = `Update the value of the environment variable using id
    %[1]v workspace environment update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -v "new value" `

var EnvironmentId string
var EnvironmentUpdateKey string
var EnvironmentUpdateValue string
var EnvironmentUpdateOrgId string
var EnvironmentUpdateWorkspaceId string

var updateEnvironmentCmd = &cobra.Command{
	Use:   "update",
	Short: "update an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		updateEnvironment()
	},
	Example: fmt.Sprintf(EnvironmentUpdateExample, rootCmd.Use),
}

func init() {
	environmentCmd.AddCommand(updateEnvironmentCmd)
	updateEnvironmentCmd.AddCommand(updateOrganizationCmd)
	updateEnvironmentCmd.Flags().StringVarP(&EnvironmentId, "id", "", "", "Id of the environment variable (required)")
	_ = updateEnvironmentCmd.MarkFlagRequired("id")
	updateEnvironmentCmd.Flags().StringVarP(&EnvironmentUpdateKey, "key", "k", "", "Key of the environment variable")
	updateEnvironmentCmd.Flags().StringVarP(&EnvironmentUpdateValue, "value", "d", "", "Value of the variable")
	updateEnvironmentCmd.Flags().StringVarP(&EnvironmentUpdateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateEnvironmentCmd.MarkFlagRequired("organization-id")
	updateEnvironmentCmd.Flags().StringVarP(&EnvironmentUpdateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = updateEnvironmentCmd.MarkFlagRequired("workspace-id")

}

func updateEnvironment() {
	client := newClient()

	environment := models.Environment{
		Attributes: &models.EnvironmentAttributes{
			Key:   EnvironmentUpdateKey,
			Value: EnvironmentUpdateValue,
		},
		ID:   EnvironmentId,
		Type: "environment",
	}
	err := client.Environment.Update(EnvironmentUpdateOrgId, EnvironmentUpdateWorkspaceId, environment)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
