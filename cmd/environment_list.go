package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var EnvironmentFilter string
var EnvironmentOrgId string
var EnvironmentWorkspaceId string
var EnvironmentListExample string = `List all existing environment variables for a workspace
    %[1]v workspace environment list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd
List specific environments variables applying a filter
    %[1]v workspace environment list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --filter key==myvariable `

var listEnvironmentsCmd = &cobra.Command{
	Use:   "list",
	Short: "list environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		listEnvironments()
	},
	Example: fmt.Sprintf(EnvironmentListExample, rootCmd.Use),
}

func init() {
	environmentCmd.AddCommand(listEnvironmentsCmd)
	listEnvironmentsCmd.Flags().StringVarP(&EnvironmentFilter, "filter", "f", "", "Filter")
	listEnvironmentsCmd.Flags().StringVarP(&EnvironmentOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = listEnvironmentsCmd.MarkFlagRequired("organization-id")
	listEnvironmentsCmd.Flags().StringVarP(&EnvironmentWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = listEnvironmentsCmd.MarkFlagRequired("workspace-id")
}

func listEnvironments() {
	client := newClient()
	resp, err := client.Environment.List(EnvironmentOrgId, EnvironmentWorkspaceId, EnvironmentFilter)

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
