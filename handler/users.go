package handler

import (
	"database-handler/util"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type UsersDbHandler struct {
	DbPath       string
	BackupPath   string
	InitPath     string
	InitDataPath string
}

func (uDb UsersDbHandler) InitDb() {
	slog.Info("Init users db")
	// validate db
	if uDb.validateDb() {
		// TODO: stop function if error
		// backup if db exist.
		uDb.backupDb()
	}
	// create db
	uDb.createDb()
}

func (uDb UsersDbHandler) validateDb() bool {
	slog.Info("Check if the database file exists")
	if _, err := os.Stat(uDb.DbPath); os.IsNotExist(err) {
		slog.Info("Database does not exist. Will Skip backup steps.")
		return false
	}
	return true
}

func (uDb UsersDbHandler) backupDb() {
	slog.Info("Backup db by moving file")
	currentTime := time.Now()
	dst := fmt.Sprintf("%s.%s", uDb.BackupPath, currentTime.Format("20060102-15-04"))

	// Move (rename) the file
	err := os.Rename(uDb.DbPath, dst)
	if err != nil {
		util.Logger.Error("Failed to rename (move) file", "error", err)
	}
	slog.Info("File moved successfully")
}

func (uDb UsersDbHandler) createDb() {
	slog.Info("create db")
	slog.Info("Connect to SQLite database to create it)")
	db, err := sql.Open("sqlite3", uDb.DbPath)
	if err != nil {
		util.Logger.Error("Failed to connect to database", "error", err)
	}
	defer db.Close()

	err = util.RunQueryFromFile(db, uDb.InitPath)
	if err != nil {
		util.Logger.Error("Can't create table", "error", err)
	}
	uDb.importInitialDataFromCsv(db)

}

func (uDb UsersDbHandler) importInitialDataFromCsv(db *sql.DB) {
	// Open CSV file
	file, err := os.Open(uDb.InitDataPath)
	if err != nil {
		util.Logger.Error("Can't open initial data file", "error", err)
	}
	defer file.Close()

	// Read CSV content
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		util.Logger.Error("Can't read file content", "error", err)
	}

	// Insert data into the database
	for i := 1; i < len(records); i++ { // Skip header row
		_, err := db.Exec("INSERT INTO users (id, name, telegram_chat_id) VALUES (?, ?, ?)", i, records[i][1], records[i][0])
		if err != nil {
			util.Logger.Error("Can't insert lines into db", "error", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}
