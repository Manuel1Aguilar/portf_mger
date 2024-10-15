package app

import (
	"database/sql"
	"io"
	"log"
	"os"

	"github.com/Manuel1Aguilar/portf_mger/internal/db"
	"github.com/Manuel1Aguilar/portf_mger/internal/services"
	"github.com/joho/godotenv"
)

type App struct {
	AssetService            *services.AssetService
	AssetTransactionService *services.AssetTransactionService
	PortfolioHoldingService *services.PortfolioHoldingService
	DB                      *sql.DB
}

func NewApp() (*App, error) {
	configureLogging()
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
	portfolioHoldingService := services.NewPortfolioHoldingService(database)
	assetTransactionService := services.NewAssetTransactionService(database, portfolioHoldingService)

	return &App{
		AssetService:            stockService,
		AssetTransactionService: assetTransactionService,
		PortfolioHoldingService: portfolioHoldingService,
		DB:                      database,
	}, nil
}

func (a *App) Close() {
	log.Println("Closing database connection...")
	a.DB.Close()
}

func configureLogging() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	logOutput := os.Getenv("LOG_OUTPUT")
	switch logOutput {
	case "discard":
		log.SetOutput(io.Discard)
	case "stdout":
		log.SetOutput(os.Stdout)
	case "file":
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(file)
	default:
		log.SetOutput(os.Stdout)
	}
}
