package cmd

import (
	"azb/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var SecretUpdateExample string = `Update the value of the secret using id
    %[1]v workspace secret update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -v "new value" `

var SecretId string
var SecretUpdateKey string
var SecretUpdateValue string
var SecretUpdateOrgId string
var SecretUpdateWorkspaceId string

var updateSecretCmd = &cobra.Command{
	Use:   "update",
	Short: "update a secret",
	Run: func(cmd *cobra.Command, args []string) {
		updateSecret()
	},
	Example: fmt.Sprintf(SecretUpdateExample, rootCmd.Use),
}

func init() {
	secretCmd.AddCommand(updateSecretCmd)
	updateSecretCmd.AddCommand(updateOrganizationCmd)
	updateSecretCmd.Flags().StringVarP(&SecretId, "id", "", "", "Id of the secret (required)")
	_ = updateSecretCmd.MarkFlagRequired("id")
	updateSecretCmd.Flags().StringVarP(&SecretUpdateKey, "key", "k", "", "Key of the secret")
	updateSecretCmd.Flags().StringVarP(&SecretUpdateValue, "value", "d", "", "Value of the variable")
	updateSecretCmd.Flags().StringVarP(&SecretUpdateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateSecretCmd.MarkFlagRequired("organization-id")
	updateSecretCmd.Flags().StringVarP(&SecretUpdateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = updateSecretCmd.MarkFlagRequired("workspace-id")

}

func updateSecret() {
	client := newClient()

	Secret := models.Secret{
		Attributes: &models.SecretAttributes{
			Key:   SecretUpdateKey,
			Value: SecretUpdateValue,
		},
		ID:   SecretId,
		Type: "secret",
	}
	err := client.Secret.Update(SecretUpdateOrgId, SecretUpdateWorkspaceId, Secret)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
