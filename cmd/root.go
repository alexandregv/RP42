package cmd

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"strings"
	"sync"

	"github.com/alexandregv/RP42/pkg/core"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "RP42",
	Short: "Discord Rich Presence for 42",
	// 	Long: `Usage of RP42:
	// 	-i, --id Client ID of your API app (required)
	// 	-s, --secret Client Secret of your API app (required)
	//
	// If you don't have an API app yet, create one here: https://profile.intra.42.fr/oauth/applications/new
	// /!\ Do NOT share your credentials to someone else, or on GitHub, etc. /!\
	//
	// `,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		osUser, err := user.Current()
		if err != nil {
			panic(err)
		}
		login := strings.ToLower(osUser.Username)

		apiClient, err := cmd.Flags().GetString("id")
		apiSecret, err := cmd.Flags().GetString("secret")

		if apiClient == "" || apiSecret == "" {
			fmt.Println("Please provide Intra API credentials with --id and --secret. See --help for help.")
			os.Exit(2)
		}

		err = core.Run(ctx, login, apiClient, apiSecret)
		if err != nil {
			fmt.Println("Error while trying to send Rich Presence: %s", err.Error())
			os.Exit(1)
		}

		fmt.Println("Sleeping... Press CTRL+C to stop.")
		m := sync.Mutex{}
		m.Lock()
		m.Lock()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().StringP("id", "i", "", "Client ID from API settings")
	rootCmd.MarkFlagRequired("id")

	rootCmd.Flags().StringP("secret", "s", "", "Client Secret from API settings")
	rootCmd.MarkFlagRequired("secret")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
