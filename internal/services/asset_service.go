package services

import (
	"database/sql"
	"fmt"
	"log"
	"stock_tracker/internal/models"
)

// Asset Service provides business logic for managing stock data
type AssetService struct {
	DB *sql.DB
}

func NewAssetService(db *sql.DB) *AssetService {
	return &AssetService{DB: db}
}

// CreateAsset inserts a new asset into the database.
func (s *AssetService) CreateAsset(asset *models.Asset) error {

	// Check if asset exists
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM asset WHERE symbol = ?)`
	err := s.DB.QueryRow(query, asset.Symbol).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if asset exists: %v", err)
		return err
	}

	if exists {
		log.Printf("Asset for %s already exists.", asset.Symbol)
		return fmt.Errorf("Asset for %s already exists", asset.Symbol)
	}

	// Insert the asset if it does not exist
	query = `INSERT INTO asset (symbol, description, asset_type) 
			  VALUES (?, ? ,?)`
	_, err = s.DB.Exec(query, asset.Symbol, asset.Description, asset.AssetType)
	if err != nil {
		log.Printf("Error inserting asset %v", err)
		return err
	}

	return nil
}

// GetAssetBySymbol retrieves the asset by symbol.
func (s *AssetService) GetAssetBySymbol(symbol string) (*models.Asset, error) {
	query := `SELECT a.id, a.symbol, a.description, a.asset_type FROM asset a
              WHERE a.symbol = ?`
	log.Printf("Query: %s", query)
	row := s.DB.QueryRow(query, symbol)

	var stockSnapshot models.Asset
	err := row.Scan(&stockSnapshot.ID, &stockSnapshot.Symbol, &stockSnapshot.Description, &stockSnapshot.AssetType)
	if err != nil {
		return nil, err
	}
	return &stockSnapshot, nil
}
