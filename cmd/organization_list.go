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
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"azb/client/client/organization"
)

var OrganizationFilter string
var listOrganizationsCmd = &cobra.Command{
	Use:   "list",
	Short: "list organizations",
	Run: func(cmd *cobra.Command, args []string) {
		listOrganizations()
	},
	Example: fmt.Sprintf(OrganizationListExample, rootCmd.Use),
}

var OrganizationListExample string = `List all existing organizations
    %[1]v organization list
List specific organizations applying a filter
    %[1]v organization list --filter name==myorg `

func init() {
	organizationCmd.AddCommand(listOrganizationsCmd)
	listOrganizationsCmd.Flags().StringVarP(&OrganizationFilter, "filter", "f", "", "Filter")
}

func listOrganizations() {
	client := newClient()
	resp, err := client.Organization.GetOrganization(organization.NewGetOrganizationParams().WithFilterOrganization(&OrganizationFilter))

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
