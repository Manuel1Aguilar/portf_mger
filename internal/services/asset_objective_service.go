package services

import (
	"database/sql"
	"log"

	"github.com/Manuel1Aguilar/portf_mger/internal/models"
)

type AssetObjectiveService struct {
	DB *sql.DB
}

func NewAssetObjectiveService(db *sql.DB) *AssetObjectiveService {
	return &AssetObjectiveService{DB: db}
}

// CreateAssetObjective inserts a new AssetObjective into the database
func (s *AssetObjectiveService) CreateAssetObjective(objective *models.AssetObjectiveCreate) error {

	// Get asset id
	var assetId int
	query := `SELECT a.id from asset a where a.symbol = ?`
	err := s.DB.QueryRow(query, objective.Symbol).Scan(&assetId)
	if err != nil {
		log.Printf("Error getting asset by symbol: %v", err)
		return err
	}
	// Check if asset already has an objective
	var exists bool
	query = `SELECT EXISTS(SELECT 1 FROM asset_objective WHERE asset_id = ?)`
	err = s.DB.QueryRow(query, assetId).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if asset already has an objective: %v", err)
		return err
	}
	// Save new asset objective
	if exists {
		query = `UPDATE asset_objective SET target_allocation_percentage = ? WHERE asset_id = ?`
	}
	query = `INSERT INTO asset_objective (target_allocation_percentage, asset_id) VALUES (?, ?)`
	_, err = s.DB.Exec(query, objective.TargetAllocationPercentage, objective.Symbol)
	if err != nil {
		log.Printf("Error saving asset_objective: %v", err)
		return err
	}
	return nil
}
