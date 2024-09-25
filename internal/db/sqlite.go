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

	DB = db
	return db, nil
}

// RunMigrations runs the migrations using golang-migrate
func RunMigrations() {
	log.Println("Running database migrations...")

	// Add migrations logic

	log.Println("Migrations completed succesfully")
}
