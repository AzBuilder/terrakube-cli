package cmd

import (
	"azb/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var WorkspaceUpdateExample string = `Update Terraform version in workspace
    %[1]v workspace update --organization-id 312b4415-806b-47a9-9452-b71f0753136e --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -t 0.14.0 `

var WorkspaceUpdateName string
var WorkspaceUpdateSource string
var WorkspaceUpdateBranch string
var WorkspaceUpdateTerraformV string
var WorkspaceUpdateOrgId string
var WorkspaceUpdateId string

var updateWorkspaceCmd = &cobra.Command{
	Use:   "update",
	Short: "update a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		updateWorkspace()
	},
	Example: fmt.Sprintf(WorkspaceUpdateExample, rootCmd.Use),
}

func init() {
	workspaceCmd.AddCommand(updateWorkspaceCmd)
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateName, "name", "n", "", "Name of the workspace (required)")
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateOrgId, "organization-id", "", "", "Id of the organization (required)")
	_ = updateWorkspaceCmd.MarkFlagRequired("organization-id")
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateId, "id", "", "", "Id of the workspace (required)")
	_ = updateWorkspaceCmd.MarkFlagRequired("id")
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateBranch, "branch", "b", "", "Branch of the workspace")
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateSource, "source", "s", "", "Source of the workspace")
	updateWorkspaceCmd.Flags().StringVarP(&WorkspaceUpdateTerraformV, "terraform-version", "t", "", "Terraform Version use in the workspace")
}

func updateWorkspace() {
	client := newClient()

	workspace := models.Workspace{
		Attributes: &models.WorkspaceAttributes{
			Name:             WorkspaceUpdateName,
			Branch:           WorkspaceUpdateBranch,
			Source:           WorkspaceUpdateSource,
			TerraformVersion: WorkspaceUpdateTerraformV,
		},
		Type: "workspace",
	}

	err := client.Workspace.Update(WorkspaceUpdateOrgId, workspace)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")
}
