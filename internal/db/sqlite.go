package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitializeSQLite opens the SQLite database and returns the DB connection.
func InitializeSQLite(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Check if the database connection is alive
	if err = db.Ping(); err != nil {
		return nil, err
	}

	DB = db
	log.Println("Database connection initialized.")
	return db, nil
}
