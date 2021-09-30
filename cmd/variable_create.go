package cmd

import (
	"terrakube/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var VariableCreateExample string = `Create a new variable
    %[1]v workspace variable create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd -k tag_name -v "Hola mundo" `

var VariableCreateKey string
var VariableCreateValue string
var VariableCreateOrgId string
var VariableCreateWorkspaceId string

var createVariableCmd = &cobra.Command{
	Use:   "create",
	Short: "create a variable",
	Run: func(cmd *cobra.Command, args []string) {
		createVariable()
	},
	Example: fmt.Sprintf(VariableCreateExample, rootCmd.Use),
}

func init() {
	variableCmd.AddCommand(createVariableCmd)
	createVariableCmd.Flags().StringVarP(&VariableCreateKey, "key", "k", "", "Key of the new variable (required)")
	_ = createVariableCmd.MarkFlagRequired("key")
	createVariableCmd.Flags().StringVarP(&VariableCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createVariableCmd.MarkFlagRequired("organization-id")
	createVariableCmd.Flags().StringVarP(&VariableCreateValue, "value", "v", "", "Value for the new variable")
	_ = createVariableCmd.MarkFlagRequired("value")
	createVariableCmd.Flags().StringVarP(&VariableCreateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createVariableCmd.MarkFlagRequired("workspace-id")
}

func createVariable() {
	client := newClient()

	variable := models.Variable{
		Attributes: &models.VariableAttributes{
			Key:   VariableCreateKey,
			Value: VariableCreateValue,
		},
		Type: "variable",
	}

	resp, err := client.Variable.Create(VariableCreateOrgId, VariableCreateWorkspaceId, variable)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
