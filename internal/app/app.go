package app

import (
	"database/sql"
	"log"

	"github.com/Manuel1Aguilar/portf_mger/internal/db"
	"github.com/Manuel1Aguilar/portf_mger/internal/services"
)

type App struct {
	AssetService *services.AssetService
	DB           *sql.DB
}

func NewApp() (*App, error) {
	// Initialize the database
	database, err := db.InitializeSQLite("portfolio.db")
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := db.RunMigrations(database); err != nil {
		return nil, err
	}

	//Initialize services
	stockService := services.NewAssetService(database)

	return &App{
		AssetService: stockService,
		DB:           database,
	}, nil
}

func (a *App) Close() {
	log.Println("Closing database connection...")
	a.DB.Close()
}
