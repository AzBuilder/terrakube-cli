package cmd

import (
	"terrakube/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var SecretCreateExample string = `Create a new secret
    %[1]v workspace secret create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd -k mySecret -v "secretvalue" `

var SecretCreateKey string
var SecretCreateValue string
var SecretCreateOrgId string
var SecretCreateWorkspaceId string

var createSecretCmd = &cobra.Command{
	Use:   "create",
	Short: "create a secret",
	Run: func(cmd *cobra.Command, args []string) {
		createSecret()
	},
	Example: fmt.Sprintf(SecretCreateExample, rootCmd.Use),
}

func init() {
	secretCmd.AddCommand(createSecretCmd)
	createSecretCmd.Flags().StringVarP(&SecretCreateKey, "key", "k", "", "Key of the new secret (required)")
	_ = createSecretCmd.MarkFlagRequired("key")
	createSecretCmd.Flags().StringVarP(&SecretCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createSecretCmd.MarkFlagRequired("organization-id")
	createSecretCmd.Flags().StringVarP(&SecretCreateValue, "value", "v", "", "Value for the new secret")
	_ = createSecretCmd.MarkFlagRequired("value")
	createSecretCmd.Flags().StringVarP(&SecretCreateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createSecretCmd.MarkFlagRequired("workspace-id")
}

func createSecret() {
	client := newClient()

	Secret := models.Secret{
		Attributes: &models.SecretAttributes{
			Key:   SecretCreateKey,
			Value: SecretCreateValue,
		},
		Type: "secret",
	}

	resp, err := client.Secret.Create(SecretCreateOrgId, SecretCreateWorkspaceId, Secret)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
