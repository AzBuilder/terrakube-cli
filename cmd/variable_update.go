package cmd

import (
	"azb/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var VariableUpdateExample string = `Update the value of the variable using id
    %[1]v workspace variable update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -v "new value" `

var VariableId string
var VariableUpdateKey string
var VariableUpdateValue string
var VariableUpdateOrgId string
var VariableUpdateWorkspaceId string

var updateVariableCmd = &cobra.Command{
	Use:   "update",
	Short: "update a variable",
	Run: func(cmd *cobra.Command, args []string) {
		updateVariable()
	},
	Example: fmt.Sprintf(VariableUpdateExample, rootCmd.Use),
}

func init() {
	variableCmd.AddCommand(updateVariableCmd)
	updateVariableCmd.AddCommand(updateOrganizationCmd)
	updateVariableCmd.Flags().StringVarP(&VariableId, "id", "", "", "Id of the variable (required)")
	_ = updateVariableCmd.MarkFlagRequired("id")
	updateVariableCmd.Flags().StringVarP(&VariableUpdateKey, "key", "k", "", "Key of the variable")
	updateVariableCmd.Flags().StringVarP(&VariableUpdateValue, "value", "d", "", "Value of the variable")
	updateVariableCmd.Flags().StringVarP(&VariableUpdateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateVariableCmd.MarkFlagRequired("organization-id")
	updateVariableCmd.Flags().StringVarP(&VariableUpdateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = updateVariableCmd.MarkFlagRequired("workspace-id")

}

func updateVariable() {
	client := newClient()

	variable := models.Variable{
		Attributes: &models.VariableAttributes{
			Key:   VariableUpdateKey,
			Value: VariableUpdateValue,
		},
		ID:   VariableId,
		Type: "variable",
	}
	err := client.Variable.Update(VariableUpdateOrgId, VariableUpdateWorkspaceId, variable)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
