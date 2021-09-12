package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"

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
var Path string
var Scheme string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to azb server",
	Long:  loginLong,
	Run: func(cmd *cobra.Command, args []string) {
		acquireAccessToken(ClientID, TenantID, Server, Scheme, Path)
	},
	Example: fmt.Sprintf(loginExamples, rootCmd.Use),
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&Server, "server", "s", "", "AZB Server url (required)")
	_ = loginCmd.MarkFlagRequired("server")
	_ = viper.BindEnv("server")
	loginCmd.Flags().StringVarP(&ClientID, "client-id", "c", "", "Azure application Client id (required)")
	_ = loginCmd.MarkFlagRequired("client-id")
	_ = viper.BindEnv("client-id", "AZB_CLIENT_ID")
	loginCmd.Flags().StringVarP(&TenantID, "tenant-id", "t", "", "Azure tenant Id (required)")
	_ = loginCmd.MarkFlagRequired("tenant-id")
	_ = viper.BindEnv("tenant-id", "AZB_TENANT_ID")
	loginCmd.Flags().StringVarP(&Path, "path", "p", "", "AZB Server path")
	_ = viper.BindEnv("path")
	loginCmd.Flags().StringVarP(&Scheme, "scheme", "", "http", "AZB Server scheme: http or https")
	_ = viper.BindEnv("scheme")
}

func acquireAccessToken(clientID string, tenantID string, url string, scheme string, path string) {

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
	color.Yellow(devCode.Result.Message)
	result, err := devCode.AuthenticationResult(ctx)
	if err != nil {
		panic(fmt.Sprintf("got error while waiting for user to input the device code: %s", err))
	}

	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.SetConfigFile(home + "/" + "azb.yml")
	viper.SetConfigType("yaml")
	viper.Set("token", result.AccessToken)
	viper.Set("server", url)
	viper.Set("scheme", url)
	if err := viper.WriteConfig(); err != nil {
		fmt.Println(err)

	}
}
