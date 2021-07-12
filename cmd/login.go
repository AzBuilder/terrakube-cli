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
	"context"
	"fmt"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const loginLong = `
Authenticate to a remote azb server.
`

var loginExamples = `
Login to a local azb server 
  %v login -s localhost:8080 -c 853b26d6-1849-4c00-8543-da5805b0e593 -t 0e6427af-ab9e-4af6-9f6f-bc098f470d75
`

var Server string
var ClientID string
var TenantID string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to azb server",
	Long:  loginLong,
	Run: func(cmd *cobra.Command, args []string) {
		acquireAccessToken(ClientID, TenantID, Server)
	},
	Example: fmt.Sprintf(loginExamples, rootCmd.Use),
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&Server, "server", "s", "", "AZB Server url (required)")
	loginCmd.MarkFlagRequired("server")
	loginCmd.Flags().StringVarP(&ClientID, "client-id", "c", "", "Azure application Client id (required)")
	loginCmd.MarkFlagRequired("client-id")
	loginCmd.Flags().StringVarP(&TenantID, "tenant-id", "t", "", "Azure tenant Id (required)")
	loginCmd.MarkFlagRequired("tenant-id")
}

func acquireAccessToken(clientID string, tenantID string, url string) {

	app, err := public.New(clientID, public.WithAuthority(fmt.Sprintf("https://login.microsoftonline.com/%v", tenantID)))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	devCode, err := app.AcquireTokenByDeviceCode(ctx, []string{"api://azbuilder/Builder.Default"})
	if err != nil {
		panic(err)
	}
	fmt.Printf(devCode.Result.Message)
	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		panic(fmt.Sprintf("got error while waiting for user to input the device code: %s", err))
	}

	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.SetConfigFile(home + "/" + "azb.yml")
	viper.SetConfigType("yaml")
	viper.Set("Token", result.AccessToken)
	viper.Set("Server", url)
	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)

	}
}
