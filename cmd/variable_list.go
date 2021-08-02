package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var VariableFilter string
var VariableOrgId string
var VariableWorkspaceId string
var VariableListExample string = `List all existing variables for a workspace
    %[1]v workspace variable list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd
List specific variable applying a filter
    %[1]v workspace variable list --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w 38b6635a-d38e-46f2-a95e-d00a416de4fd --filter key==myvariable `

var listVariablesCmd = &cobra.Command{
	Use:   "list",
	Short: "list variables",
	Run: func(cmd *cobra.Command, args []string) {
		listVariables()
	},
	Example: fmt.Sprintf(VariableListExample, rootCmd.Use),
}

func init() {
	variableCmd.AddCommand(listVariablesCmd)
	listVariablesCmd.Flags().StringVarP(&VariableFilter, "filter", "f", "", "Filter")
	listVariablesCmd.Flags().StringVarP(&VariableOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = listVariablesCmd.MarkFlagRequired("organization-id")
	listVariablesCmd.Flags().StringVarP(&VariableWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = listVariablesCmd.MarkFlagRequired("workspace-id")
}

func listVariables() {
	client := newClient()
	resp, err := client.Variable.List(VariableOrgId, VariableWorkspaceId, VariableFilter)

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
