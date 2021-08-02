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

var ModuleUpdateExample string = `Update the description of the module using id
    %[1]v module update --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb --id 38b6635a-d38e-46f2-a95e-d00a416de4fd -d "new description" `

var ModuleId string
var ModuleUpdateDescription string
var ModuleUpdateName string
var ModuleUpdateOrgId string
var ModuleUpdateSource string
var ModuleUpdateProvider string

var updateModuleCmd = &cobra.Command{
	Use:   "update",
	Short: "update a module",
	Run: func(cmd *cobra.Command, args []string) {
		updateModule()
	},
	Example: fmt.Sprintf(ModuleUpdateExample, rootCmd.Use),
}

func init() {
	moduleCmd.AddCommand(updateModuleCmd)
	updateModuleCmd.AddCommand(updateOrganizationCmd)
	updateModuleCmd.Flags().StringVarP(&ModuleId, "id", "", "", "Id of the module (required)")
	_ = updateModuleCmd.MarkFlagRequired("id")
	updateModuleCmd.Flags().StringVarP(&ModuleUpdateName, "name", "n", "", "Name of the module")
	updateModuleCmd.Flags().StringVarP(&ModuleUpdateDescription, "description", "d", "", "Description of the module")
	updateModuleCmd.Flags().StringVarP(&ModuleUpdateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = updateModuleCmd.MarkFlagRequired("organization-id")
	updateModuleCmd.Flags().StringVarP(&ModuleUpdateSource, "source", "s", "", "Source of the module")
	updateModuleCmd.Flags().StringVarP(&ModuleUpdateProvider, "provider", "p", "", "Provider of the module")
}

func updateModule() {
	client := newClient()

	module := models.Module{
		Attributes: &models.ModuleAttributes{
			Name:        ModuleUpdateName,
			Description: ModuleUpdateDescription,
			Source:      ModuleUpdateSource,
			Provider:    ModuleUpdateProvider,
		},
		ID:   ModuleId,
		Type: "module",
	}
	err := client.Module.Update(ModuleUpdateOrgId, module)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Updated")

}
