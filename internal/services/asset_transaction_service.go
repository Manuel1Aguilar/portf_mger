package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Manuel1Aguilar/portf_mger/internal/models"
)

type AssetTransactionService struct {
	DB                      *sql.DB
	PortfolioHoldingService PortfolioHoldingService
}

func NewAssetTransactionService(db *sql.DB, portfolioHoldingService *PortfolioHoldingService) *AssetTransactionService {
	return &AssetTransactionService{
		DB:                      db,
		PortfolioHoldingService: *portfolioHoldingService,
	}
}

// CreateAsssetTransaction inserts a new AssetTransaction into the database
func (s *AssetTransactionService) CreateAssetTransaction(transaction *models.AssetTransaction) error {
	query := `INSERT INTO asset_transaction(asset_id, transaction_type, valueUSD, units, unit_price, date_transacted)
				VALUES	(?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, transaction.AssetID, transaction.TransactionType, transaction.ValueUSD, transaction.Units,
		transaction.UnitPrice, transaction.DateTransacted)
	if err != nil {
		return err
	}
	return nil
}

// SanitizeAssetTransactionCreationModel checks the creation model has a valid asset,
// verifies that units * unit_price = valueUSD, adds the date_transacted from time now
// and that transaction_type is either BUY or SELL
func (s *AssetTransactionService) SanitizeAssetTransactionCreationModel(
	tModel *models.AssetTransactionCreate) (*models.AssetTransaction, error) {
	// Get Asset id
	query := `SELECT TOP 1 id FROM asset WHERE symbol = ?`
	var assetId int
	err := s.DB.QueryRow(query, tModel.Symbol).Scan(&assetId)
	if err != nil {
		return nil, fmt.Errorf("Error getting asset by symbol: %v", err)
	}
	// If valueUSD != units * unit_price return err
	if tModel.ValueUSD != tModel.Units*tModel.UnitPrice {
		return nil, fmt.Errorf("Check that the total amount of the transaction corresponds with the\n" +
			"number of units and unit price inputted")
	}
	// If transaction_type != BUY || SELL return err
	if tModel.Type != "BUY" && tModel.Type != "SELL" {
		return nil, fmt.Errorf("Couldn't parse type, should be either BUY or SELL")
	}

	// add date_transacted from time now
	entity := &models.AssetTransaction{
		AssetID:         assetId,
		TransactionType: tModel.Type,
		ValueUSD:        tModel.ValueUSD,
		Units:           tModel.Units,
		UnitPrice:       tModel.UnitPrice,
		DateTransacted:  time.Now(),
	}
	return entity, nil
}

// SaveAssetTransaction sanitizes the creation model, inserts the transaction in the db and
// calls the portfolio_service to update the portfolio
func (s *AssetTransactionService) SaveAssetTransaction(
	tModel *models.AssetTransactionCreate) error {
	// Open db transaction
	// Sanitize the transaction
	transaction, err := s.SanitizeAssetTransactionCreationModel(tModel)
	if err != nil {
		log.Printf("Error sanitizing the transaction: %v", err)
		// Rollback transaction
		return err
	}

	// Insert the transaction
	err = s.CreateAssetTransaction(transaction)
	if err != nil {
		log.Printf("Error inserting the transaction: %v", err)
		// Rollback transaction
		return err
	}

	// Let the portfolio service know
	err = s.PortfolioHoldingService.FinishTransaction(transaction)
	if err != nil {
		log.Printf("Error updating the portfolio holdings: %v \n", err)
		// Rollback transaction
		return err
	}
	// Commit transaction
	return nil
}
