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
	"fmt"

	"github.com/spf13/cobra"
)

const logoutLong = `
Remove credentials stored for a remote azb server.
`

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from azb server",
	Long:  logoutLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout ok")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
