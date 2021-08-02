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
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var ModuleCreateExample string = `Create a new module
    %[1]v module create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -n myModule -d "module description" -p azurerm -s https://github.com/AzBuilder/terraform-sample-repository.git `

var ModuleCreateName string
var ModuleCreateDescription string
var ModuleCreateOrgId string
var ModuleCreateSource string
var ModuleCreateProvider string

var createModuleCmd = &cobra.Command{
	Use:   "create",
	Short: "create a module",
	Run: func(cmd *cobra.Command, args []string) {
		createModule()
	},
	Example: fmt.Sprintf(ModuleCreateExample, rootCmd.Use),
}

func init() {
	moduleCmd.AddCommand(createModuleCmd)
	createModuleCmd.Flags().StringVarP(&ModuleCreateName, "name", "n", "", "Name of the new module (required)")
	_ = createModuleCmd.MarkFlagRequired("name")
	createModuleCmd.Flags().StringVarP(&ModuleCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createModuleCmd.MarkFlagRequired("organization-id")
	createModuleCmd.Flags().StringVarP(&ModuleCreateDescription, "description", "d", "", "Description of the new module")
	createModuleCmd.Flags().StringVarP(&ModuleCreateSource, "source", "s", "", "Source of the new module")
	createModuleCmd.Flags().StringVarP(&ModuleCreateProvider, "provider", "p", "", "Provider of the new module")

}

func createModule() {
	client := newClient()

	module := models.Module{
		Attributes: &models.ModuleAttributes{
			Description:  ModuleCreateDescription,
			Name:         ModuleCreateName,
			Source:       ModuleCreateSource,
			SourceSample: ModuleCreateSource,
			Provider:     ModuleCreateProvider,
		},
		Type: "module",
	}

	resp, err := client.Module.Create(ModuleCreateOrgId, module)

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
