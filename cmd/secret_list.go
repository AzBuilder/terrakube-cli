package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var SecretFilter string
var SecretOrgId string
var SecretWorkspaceId string
var SecretListExample string = `List all existing secrets for a workspace
    %[1]v workspace secret list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd
List specific secrets applying a filter
    %[1]v workspace secret list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --filter key==mysecret `

var listSecretsCmd = &cobra.Command{
	Use:   "list",
	Short: "list secrets",
	Run: func(cmd *cobra.Command, args []string) {
		listSecrets()
	},
	Example: fmt.Sprintf(SecretListExample, rootCmd.Use),
}

func init() {
	secretCmd.AddCommand(listSecretsCmd)
	listSecretsCmd.Flags().StringVarP(&SecretFilter, "filter", "f", "", "Filter")
	listSecretsCmd.Flags().StringVarP(&SecretOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = listSecretsCmd.MarkFlagRequired("organization-id")
	listSecretsCmd.Flags().StringVarP(&SecretWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = listSecretsCmd.MarkFlagRequired("workspace-id")
}

func listSecrets() {
	client := newClient()
	resp, err := client.Secret.List(SecretOrgId, SecretWorkspaceId, SecretFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("%s\n", string(prettyJSON))
}
