package cmd

import (
	"terrakube/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var WorkspaceCreateExample string = `Create a new workspace
    %[1]v workspace create --organization-id 312b4415-806b-47a9-9452-b71f0753136e -n myWorkspace -s https://github.com/terrakube-io/terraform-sample-repository.git -b master -t 0.15.0`

var WorkspaceCreateName string
var WorkspaceCreateSource string
var WorkspaceCreateBranch string
var WorkspaceCreateTerraformV string
var WorkspaceCreateOrgId string
var createWorkspaceCmd = &cobra.Command{
	Use:   "create",
	Short: "create a workspace",
	Run: func(cmd *cobra.Command, args []string) {
		createWorkspace()
	},
	Example: fmt.Sprintf(WorkspaceCreateExample, rootCmd.Use),
}

func init() {
	workspaceCmd.AddCommand(createWorkspaceCmd)
	createWorkspaceCmd.Flags().StringVarP(&WorkspaceCreateName, "name", "n", "", "Name of the new workspace (required)")
	_ = createWorkspaceCmd.MarkFlagRequired("name")
	createWorkspaceCmd.Flags().StringVarP(&WorkspaceCreateOrgId, "organization-id", "", "", "Id of the organization (required)")
	_ = createWorkspaceCmd.MarkFlagRequired("organization-id")
	createWorkspaceCmd.Flags().StringVarP(&WorkspaceCreateBranch, "branch", "b", "", "Branch of the new workspace")
	createWorkspaceCmd.Flags().StringVarP(&WorkspaceCreateSource, "source", "s", "", "Source of the new workspace")
	createWorkspaceCmd.Flags().StringVarP(&WorkspaceCreateTerraformV, "terraform-version", "t", "", "Terraform Version use in the new workspace")
}

func createWorkspace() {
	client := newClient()

	workspace := models.Workspace{
		Attributes: &models.WorkspaceAttributes{
			Name:             WorkspaceCreateName,
			Source:           WorkspaceCreateSource,
			Branch:           WorkspaceCreateBranch,
			TerraformVersion: WorkspaceCreateTerraformV,
		},
		Type: "workspace",
	}

	resp, err := client.Workspace.Create(WorkspaceCreateOrgId, workspace)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
