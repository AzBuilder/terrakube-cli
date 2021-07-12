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
	"azb/client/client/organization"
	"fmt"

	"github.com/spf13/cobra"
)

var OrganizationDeleteId string
var OrganizationDeleteExample string = `Delete an organization using id
    %[1]v organization delete --id 38b6635a-d38e-46f2-a95e-d00a416de4fd`

var deleteOrganizationCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an organization",
	Run: func(cmd *cobra.Command, args []string) {
		deleteOrganization()
	},
	Example: fmt.Sprintf(OrganizationDeleteExample, rootCmd.Use),
}

func init() {
	organizationCmd.AddCommand(deleteOrganizationCmd)
	deleteOrganizationCmd.Flags().StringVarP(&OrganizationDeleteId, "id", "", "", "Id of the organization (required)")
<<<<<<< HEAD
	_ = deleteOrganizationCmd.MarkFlagRequired("id")
=======
	//lint:ignore
	deleteOrganizationCmd.MarkFlagRequired("id")
>>>>>>> bf3c7991b9958288281f167626fcb5fcdb065b91
}

func deleteOrganization() {
	client := newClient()

	_, err := client.Organization.DeleteOrganizationOrganizationID(organization.NewDeleteOrganizationOrganizationIDParams().WithOrganizationID(OrganizationDeleteId))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("deleted")

}
