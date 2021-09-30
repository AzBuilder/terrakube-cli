package cmd

import (
	"terrakube/client/models"
	"fmt"

	"github.com/spf13/cobra"
)

var OrganizationUpdateExample string = `Update the description of the organization using id
    %[1]v organization update --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -d "new description" `

var updateOrganizationCmd = &cobra.Command{
	Use:   "update",
	Short: "update an organization",
	Run: func(cmd *cobra.Command, args []string) {
		updateOrganization()
	},
	Example: fmt.Sprintf(OrganizationUpdateExample, rootCmd.Use),
}

var OrganizationId string
var OrganizationUpdateDescription string
var OrganizationUpdateName string

func init() {
	organizationCmd.AddCommand(updateOrganizationCmd)
	updateOrganizationCmd.Flags().StringVarP(&OrganizationId, "id", "", "", "Id of the organization (required)")
	_ = updateOrganizationCmd.MarkFlagRequired("id")
	updateOrganizationCmd.Flags().StringVarP(&OrganizationUpdateName, "name", "n", "", "Name of the organization")
	updateOrganizationCmd.Flags().StringVarP(&OrganizationUpdateDescription, "description", "d", "", "Description of the organization")
}

func updateOrganization() {
	client := newClient()

	organization := models.Organization{
		Attributes: &models.OrganizationAttributes{
			Name:        OrganizationUpdateName,
			Description: OrganizationUpdateDescription,
		},
		Type: "organization",
		ID:   OrganizationId,
	}
	err := client.Organization.Update(organization)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
