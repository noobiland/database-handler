/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database-handler/util"
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

// getAllCmd represents the getAll command
var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info(fmt.Sprintf("Supported DBs: %s", util.GetSupportedDbs()))
	},
}

// TODO:
func init() {
	rootCmd.AddCommand(getAllCmd)
}
