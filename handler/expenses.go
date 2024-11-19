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

func (eDb ExpensesDbHandler) InitDb(migrationFlag bool) {
	slog.Info("Init expenses db")
	// validate db
	var backupLoc string
	if eDb.validateDb() {
		// TODO: stop function if error
		// backup if db exist.
		backupLoc = eDb.backupDb()
	}
	// create db
	eDb.createDb()
	if backupLoc != "" && migrationFlag {
		eDb.migrateData(backupLoc)
	}
}

func (eDb ExpensesDbHandler) validateDb() bool {
	slog.Info("Check if the database file exists")
	if _, err := os.Stat(eDb.DbPath); os.IsNotExist(err) {
		slog.Info("Database does not exist. Will Skip backup steps.")
		return false
	}
	return true
}

func (eDb ExpensesDbHandler) backupDb() string {
	slog.Info("Backup db by moving file")
	currentTime := time.Now()
	dst := fmt.Sprintf("%s.%s", eDb.BackupPath, currentTime.Format("20060102-15-04"))

	// Move (rename) the file
	err := os.Rename(eDb.DbPath, dst)
	if err != nil {
		util.Logger.Error("Failed to rename (move) file", "error", err)
	}
	slog.Info("File moved successfully")
	return dst
}

func (eDb ExpensesDbHandler) createDb() {
	slog.Info("create db")
	slog.Info("Connect to SQLite database to create it")
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

func (eDb ExpensesDbHandler) migrateData(backupLoc string) {
	// TODO: make a config with tables which should be migrated
	tables := map[string]string{"expenses": "(timestamp, user, amount, category, payment) VALUES (?, ?, ?, ?, ?)"}

	db1, err := sql.Open("sqlite3", backupLoc)
	if err != nil {
		util.Logger.Error("Failed to open backup db", "error", err)
	}
	defer db1.Close()

	db2, err := sql.Open("sqlite3", eDb.DbPath)
	if err != nil {
		util.Logger.Error("Failed to open new db", "error", err)
	}
	defer db2.Close()
	for migrationTable, tableFields := range tables {
		rows, err := db1.Query(fmt.Sprintf("SELECT * FROM %s", migrationTable))
		if err != nil {
			util.Logger.Error("Failed to backup db", "error", err)
		}
		defer rows.Close()

		tx, err := db2.Begin()
		if err != nil {
			util.Logger.Error("Failed to begin transaction on new db", "error", err)
		}

		stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s %s", migrationTable, tableFields))
		if err != nil {
			util.Logger.Error("Failed to prepare statement:", "error", err)
		}
		defer stmt.Close()

		for rows.Next() {
			// TODO: make generic scheme insertion
			var timestamp, amount int
			var user, category, payment string
			err = rows.Scan(&timestamp, &user, &amount, &category, &payment)
			if err != nil {
				util.Logger.Error("Failed to scan row:", "error", err)
			}

			_, err = stmt.Exec(timestamp, user, amount, category, payment)
			if err != nil {
				util.Logger.Error("Failed to insert row into new db", "error", err)
			}
		}

		err = tx.Commit()
		if err != nil {
			util.Logger.Error("Failed to commit transaction:", "error", err)
		}
	}
	slog.Info("Data migration completed successfully!")
}
