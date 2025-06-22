/*
Copyright © 2024 Telitsyn Dmitry <dmitry.telitsyn@gmail.com>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log/slog"

	"database-handler/handler"
	"database-handler/util"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
		migrationFlag, _ := cmd.Flags().GetBool("migration")
		slog.Info("Starting arguments verification")
		verified := true
		for _, database := range args {
			verified = verified && util.CheckDbIsSupported(database)
		}
		if verified {
			slog.Info("All the databases are confirmed, starting the operation")
		} else {
			slog.Error("Not all dbs are supported, check your arguments and try again")
			return
		}

		for _, database := range args {
			h, err := getDbHandlerByName(database)
			if err != nil {
				panic(err)
			}
			h.InitDb(migrationFlag)
		}
	},
}

type DBHandler interface {
	InitDb(migrationFlag bool)
	ValidateDb() bool
	BackupDb() string
	CreateDb()
	MigrateData(backupLoc string)
	ImportInitialDataFromCsv(db *sql.DB)
}

func getDbHandlerByName(name string) (DBHandler, error) {
	var cfg util.Config
	if err := yaml.Unmarshal(util.GetConfigs(), &cfg); err != nil {
		panic(err)
	}
	switch name {
	case "expenses":
		return handler.ExpensesDbHandler{
			DbPath:     cfg.Databases[name].DBPath,
			BackupPath: cfg.Databases[name].BackupPath,
			InitPath:   cfg.Databases[name].InitPath,
		}, nil
	case "users":
		return handler.ExpensesDbHandler{
			DbPath:       cfg.Databases[name].DBPath,
			BackupPath:   cfg.Databases[name].BackupPath,
			InitPath:     cfg.Databases[name].InitPath,
			InitDataPath: *cfg.Databases[name].InitDataPath,
		}, nil
	default:
		util.Logger.Error("arg is not supported", "error", nil)
		return nil, fmt.Errorf("unknown handler type: %s", name)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("migration", "m", false, "init with data migration from existed db")
}
