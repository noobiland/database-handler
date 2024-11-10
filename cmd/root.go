/*
Copyright © 2024 Dmitry Telitsyn <dmitry.telitsyn@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var version = "0.0.1"
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "database-handler",
	Version: version,
	Short: "App to handle all databases for homelab",
	Long: `Application to setup and manage databases`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.database-handler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("all", "a", false,"perform for all databases")
}


