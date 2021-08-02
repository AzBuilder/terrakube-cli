package cmd

import (
	"azb/client/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var EnvironmentCreateExample string = `Create a new environment
    %[1]v workspace environment create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd -k ARM_CLIENT_ID -v "38b6635a-d38e-46f2-a95e-d00a416de4fd" `

var EnvironmentCreateKey string
var EnvironmentCreateValue string
var EnvironmentCreateOrgId string
var EnvironmentCreateWorkspaceId string

var createEnvironmentCmd = &cobra.Command{
	Use:   "create",
	Short: "create an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		createEnvironment()
	},
	Example: fmt.Sprintf(EnvironmentCreateExample, rootCmd.Use),
}

func init() {
	environmentCmd.AddCommand(createEnvironmentCmd)
	createEnvironmentCmd.Flags().StringVarP(&EnvironmentCreateKey, "key", "k", "", "Key of the new environment variable (required)")
	_ = createEnvironmentCmd.MarkFlagRequired("key")
	createEnvironmentCmd.Flags().StringVarP(&EnvironmentCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createEnvironmentCmd.MarkFlagRequired("organization-id")
	createEnvironmentCmd.Flags().StringVarP(&EnvironmentCreateValue, "value", "v", "", "Value for the new environment variable")
	_ = createEnvironmentCmd.MarkFlagRequired("value")
	createEnvironmentCmd.Flags().StringVarP(&EnvironmentCreateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createEnvironmentCmd.MarkFlagRequired("workspace-id")
}

func createEnvironment() {
	client := newClient()

	environment := models.Environment{
		Attributes: &models.EnvironmentAttributes{
			Key:   EnvironmentCreateKey,
			Value: EnvironmentCreateValue,
		},
		Type: "environment",
	}

	resp, err := client.Environment.Create(EnvironmentCreateOrgId, EnvironmentCreateWorkspaceId, environment)

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
