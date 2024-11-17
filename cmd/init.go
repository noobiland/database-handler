/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"database-handler/handler"
	"database-handler/util"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init <database name>",
	Long:  `Tool to initialize databases. Use get command to get database information`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Init called")
		if len(args) == 0 {
			util.Logger.Error("no database name specified", "error", nil)
		}
		for _, database := range args {
			slog.Info(fmt.Sprintf("Initialization for argument: %s", database))
			switch database {
			case "expenses":
				h := handler.ExpensesDbHandler{
					DbPath:     "./databases/db/expenses.db",
					BackupPath: "./databases/backup/expenses.db.bkp",
					InitPath:   "./sql/expenses/init.sql",
				}
				h.InitDb()
			case "users":
				h := handler.UsersDbHandler{
					DbPath:       "./databases/db/users.db",
					BackupPath:   "./databases/backup/users.db.bkp",
					InitPath:     "./sql/users/init.sql",
					InitDataPath: "./secret-data/users.csv",
				}
				h.InitDb()
			default:
				util.Logger.Error("arg is not supported", "error", nil)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
