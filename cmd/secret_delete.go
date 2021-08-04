package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SecretDeleteExample string = `Delete a secret
    %[1]v workspace secret delete --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb  -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --id 38b6635a-d38e-46f2-a95e-d00a416de4fd `

var SecretDeleteId string
var SecretDeleteOrgId string
var SecretDeleteWorkspaceId string

var deleteSecretCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a secret",
	Run: func(cmd *cobra.Command, args []string) {
		deleteSecret()
	},
	Example: fmt.Sprintf(SecretDeleteExample, rootCmd.Use),
}

func init() {
	secretCmd.AddCommand(deleteSecretCmd)
	deleteSecretCmd.Flags().StringVarP(&SecretDeleteOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = deleteSecretCmd.MarkFlagRequired("organization-id")
	deleteSecretCmd.Flags().StringVarP(&SecretDeleteId, "id", "", "", "Id of the Secret (required)")
	_ = deleteSecretCmd.MarkFlagRequired("id")
	deleteSecretCmd.Flags().StringVarP(&SecretDeleteWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = deleteSecretCmd.MarkFlagRequired("workspace-id")
}

func deleteSecret() {
	client := newClient()

	err := client.Secret.Delete(SecretDeleteOrgId, SecretDeleteWorkspaceId, SecretDeleteId)

	if err != nil {
		fmt.Println(err)
		return
	}
}
