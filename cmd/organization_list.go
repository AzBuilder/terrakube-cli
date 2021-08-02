package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
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
	organizations, err := client.Organization.List(OrganizationFilter)

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyJSON, err := json.MarshalIndent(organizations, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("%s\n", string(prettyJSON))
}
