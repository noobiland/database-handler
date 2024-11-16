package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func RunQueryFromFile(db *sql.DB, sqlFilePath string) error {
    // Read the SQL file
    sqlContent, err := os.ReadFile(sqlFilePath)
    if err != nil {
        log.Fatalf("Failed to read SQL file: %v", err)
		return err
    }

    // Split the content by semicolon to handle each query separately
    queries := strings.Split(string(sqlContent), ";")
    for _, query := range queries {
        // Trim whitespace
        query = strings.TrimSpace(query)
        
        if query == "" {
            continue
        }
        
        // Execute each query
        _, err := db.Exec(query)
        if err != nil {
            log.Printf("Failed to execute query: %v\nError: %v", query, err)
			return err
        } else {
            fmt.Printf("Executed query: %v\n", query)
        }
    }

    log.Println("All queries executed successfully")
	return nil
}
