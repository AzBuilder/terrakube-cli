package cmd

import (
	"fmt"
	"terrakube/client/models"

	"github.com/spf13/cobra"
)

var TeamCreateExample string = `Create a new Team
    %[1]v team create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -n AZB_USER --manage-workspace=true --manage-module=true --manage-provider=true`

var TeamCreateName string
var TeamCreateOrgId string
var TeamCreateManageProvider bool
var TeamCreateManageModule bool
var TeamCreateManageWorkspace bool

var createTeamCmd = &cobra.Command{
	Use:   "create",
	Short: "create a Team",
	Run: func(cmd *cobra.Command, args []string) {
		createTeam()
	},
	Example: fmt.Sprintf(TeamCreateExample, rootCmd.Use),
}

func init() {
	teamCmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().StringVarP(&TeamCreateName, "name", "n", "", "Name of the new Team (required)")
	_ = createTeamCmd.MarkFlagRequired("name")
	createTeamCmd.Flags().StringVarP(&TeamCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createTeamCmd.MarkFlagRequired("organization-id")
	createTeamCmd.Flags().BoolVarP(&TeamCreateManageProvider, "manage-provider", "", false, "Manage Provider Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-provider")
	createTeamCmd.Flags().BoolVarP(&TeamCreateManageModule, "manage-module", "", false, "Manage Module Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-module")
	createTeamCmd.Flags().BoolVarP(&TeamCreateManageWorkspace, "manage-workspace", "", false, "Manage Workspaces Permissions")
	_ = createTeamCmd.MarkFlagRequired("manage-workspace")

}

func createTeam() {
	client := newClient()

	team := models.Team{
		Attributes: &models.TeamAttributes{
			Name:            TeamCreateName,
			ManageWorkspace: TeamCreateManageWorkspace,
			ManageModule:    TeamCreateManageModule,
			ManageProvider:  TeamCreateManageProvider,
		},
		Type: "team",
	}

	resp, err := client.Team.Create(TeamCreateOrgId, team)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
