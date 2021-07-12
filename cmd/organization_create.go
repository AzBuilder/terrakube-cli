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
	"azb/client/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var OrganizationCreateExample string = `Create a new organization
    %[1]v organization create -n myorg -d "org description" `

var OrganizationCreateName string
var OrganizationCreateDescription string
var createOrganizationCmd = &cobra.Command{
	Use:   "create",
	Short: "create an organization",
	Run: func(cmd *cobra.Command, args []string) {
		createOrganization()
	},
	Example: fmt.Sprintf(OrganizationCreateExample, rootCmd.Use),
}

func init() {
	organizationCmd.AddCommand(createOrganizationCmd)
	createOrganizationCmd.Flags().StringVarP(&OrganizationCreateName, "name", "n", "", "Name of the new organization (required)")
	createOrganizationCmd.MarkFlagRequired("name")
	createOrganizationCmd.Flags().StringVarP(&OrganizationCreateDescription, "description", "d", "", "Description of the new organization")

}

func createOrganization() {
	client := newClient()

	organizationBody := organization.PostOrganizationBody{
		Data: &models.Organization{
			Attributes: &models.OrganizationAttributes{
				Description: OrganizationCreateDescription,
				Name:        OrganizationCreateName,
			},
			Type: "organization",
		},
	}

	organizationParams := organization.NewPostOrganizationParams().WithOrganization(organizationBody)
	resp, err := client.Organization.PostOrganization(organizationParams)

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyJSON, err := json.MarshalIndent(resp.Payload.Data, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("%s\n", string(prettyJSON))

}
