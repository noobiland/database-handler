package handler

import (
	"database-handler/util"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ExpensesDbHandler struct {
	DbPath     string
	BackupPath string
	InitPath   string
}

func (eDb ExpensesDbHandler) InitDb() {
	slog.Info("Init expenses db")
	// validate db
	if eDb.validateDb() {
		// TODO: stop function if error
		// backup if db exist.
		eDb.backupDb()
	}
	// create db
	eDb.createDb()
}

func (eDb ExpensesDbHandler) validateDb() bool {
	slog.Info("Check if the database file exists")
	if _, err := os.Stat(eDb.DbPath); os.IsNotExist(err) {
		slog.Info("Database does not exist. Will Skip backup steps.")
		return false
	}
	return true
}

func (eDb ExpensesDbHandler) backupDb() {
	slog.Info("Backup db by moving file")
	currentTime := time.Now()
	dst := fmt.Sprintf("%s.%s", eDb.BackupPath, currentTime.Format("20060102-15-04"))

	// Move (rename) the file
	err := os.Rename(eDb.DbPath, dst)
	if err != nil {
		util.Logger.Error("Failed to rename (move) file", "error", err)
	}
	slog.Info("File moved successfully")
}

func (eDb ExpensesDbHandler) createDb() {
	slog.Info("create db")
	slog.Info("Connect to SQLite database to create it)")
	db, err := sql.Open("sqlite3", eDb.DbPath)
	if err != nil {
		util.Logger.Error("Failed to connect to database", "error", err)
	}
	defer db.Close()

	err = util.RunQueryFromFile(db, eDb.InitPath)
	if err != nil {
		util.Logger.Error("Can't create table", "error", err)
	}
}
