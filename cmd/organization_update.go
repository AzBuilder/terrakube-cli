/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"azb/client/models"
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
