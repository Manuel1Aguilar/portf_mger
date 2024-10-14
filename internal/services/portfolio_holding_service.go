package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Manuel1Aguilar/portf_mger/internal/models"
)

type PortfolioHoldingService struct {
	DB *sql.DB
}

func NewPortfolioHoldingService(db *sql.DB) *PortfolioHoldingService {
	return &PortfolioHoldingService{DB: db}
}

// Finish Asset Transaction
// from transaction model create a new PfolioEntry model and insert it
func (s *PortfolioHoldingService) FinishTransactionWithTx(tx *sql.Tx, tModel *models.AssetTransaction) error {
	// Check if portfolio_holding for asset exists
	exists, err := s.ExistsByAssetIdWithTx(tx, tModel.AssetID)
	if err != nil {
		return err
	}
	if exists {
		// if exists update
		holding, err := s.GetByAssetIdWithTx(tx, tModel.AssetID)
		if err != nil {
			return err
		}
		//TODO update holding value
		if tModel.TransactionType == models.TransactionTypeBuy {
			holding.UnitsHeld += tModel.Units
			holding.USDValue += tModel.ValueUSD
		} else {
			holding.UnitsHeld -= tModel.Units
			holding.USDValue -= tModel.ValueUSD
		}
		holding.LastUpdated = time.Now()
		err = s.UpdateWithTx(tx, holding)
		if err != nil {
			return err
		}

		return nil
	}
	// if not created
	holding := &models.PortfolioHolding{
		AssetID:     tModel.AssetID,
		UnitsHeld:   tModel.Units,
		USDValue:    tModel.ValueUSD,
		LastUpdated: time.Now(),
		TargetPp:    0,
	}

	err = s.AddWithTx(tx, holding)
	if err != nil {
		return err
	}

	return nil
}

// Check if holding exists by asset id with transaction
func (s *PortfolioHoldingService) ExistsByAssetIdWithTx(tx *sql.Tx, assetId int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM portfolio_holding WHERE asset_id = ?)`
	var exists bool
	err := tx.QueryRow(query, assetId).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if holding exists: %v \n", err)
		return false, err
	}

	return exists, nil
}

// Check if holding exists by asset id
func (s *PortfolioHoldingService) ExistsByAssetId(assetId int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM portfolio_holding WHERE asset_id = ?)`
	var exists bool
	err := s.DB.QueryRow(query, assetId).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if holding exists: %v \n", err)
		return false, err
	}

	return exists, nil
}

// Get holding by asset id with tx
func (s *PortfolioHoldingService) GetByAssetIdWithTx(tx *sql.Tx, assetId int) (*models.PortfolioHolding, error) {
	query := `SELECT p.id, p.asset_id, p.units_held, p.usd_value, p.last_updated, p.target_pp FROM portfolio_holding p WHERE p.asset_id = ? LIMIT 1`
	var holding models.PortfolioHolding
	err := tx.QueryRow(query, assetId).Scan(&holding.ID, &holding.AssetID, &holding.UnitsHeld, &holding.USDValue, &holding.LastUpdated, &holding.TargetPp)
	if err != nil {
		log.Printf("Error getting holding by asset id: %v \n", err)
		return nil, err
	}
	return &holding, nil
}

// Get holding by asset id
func (s *PortfolioHoldingService) GetByAssetId(assetId int) (*models.PortfolioHolding, error) {
	query := `SELECT p.id, p.asset_id, p.units_held, p.usd_value, p.last_updated, p.target_pp FROM portfolio_holding p WHERE p.asset_id = ? LIMIT 1`
	var holding models.PortfolioHolding
	err := s.DB.QueryRow(query, assetId).Scan(&holding.ID, &holding.AssetID, &holding.UnitsHeld, &holding.USDValue, &holding.LastUpdated, &holding.TargetPp)
	if err != nil {
		log.Printf("Error getting holding by asset id: %v \n", err)
		return nil, err
	}
	return &holding, nil
}

// Add holding with Transaction
func (s *PortfolioHoldingService) AddWithTx(tx *sql.Tx, pModel *models.PortfolioHolding) error {
	query := `INSERT INTO portfolio_holding (asset_id, units_held, usd_value, last_updated, target_pp) values (?, ?, ?, ?, ?, ?)`

	_, err := tx.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.USDValue, pModel.LastUpdated, pModel.TargetPp)
	if err != nil {
		log.Printf("Error inserting portfolio_holding: %v \n", err)
		return err
	}
	return nil
}

// Add holding
func (s *PortfolioHoldingService) Add(pModel *models.PortfolioHolding) error {
	query := `INSERT INTO portfolio_holding (asset_id, units_held, usd_value, last_updated, target_pp) values (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.USDValue, pModel.LastUpdated, pModel.TargetPp)
	if err != nil {
		log.Printf("Error inserting portfolio_holding: %v \n", err)
		return err
	}
	return nil
}

// Update holding with transaction
func (s *PortfolioHoldingService) UpdateWithTx(tx *sql.Tx, pModel *models.PortfolioHolding) error {
	query := `UPDATE portfolio_holding SET asset_id = ?, units_held = ?, usd_value = ?, last_updated = ?, target_pp = ? WHERE id = ?`
	_, err := tx.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.USDValue, pModel.LastUpdated, pModel.TargetPp, pModel.ID)
	if err != nil {
		log.Printf("Error updating portfolio_holding: %v \n", err)
		return err
	}
	return nil
}

// Update holding
func (s *PortfolioHoldingService) Update(pModel *models.PortfolioHolding) error {
	query := `UPDATE portfolio_holding SET asset_id = ?, units_held = ?, usd_value = ?, last_updated = ?, target_pp = ? WHERE id = ?`
	_, err := s.DB.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.USDValue, pModel.LastUpdated, pModel.TargetPp, pModel.ID)
	if err != nil {
		log.Printf("Error updating portfolio_holding: %v \n", err)
		return err
	}
	return nil
}

func (s *PortfolioHoldingService) GetUpdatedPortfolio() (*models.Portfolio, error) {
	// Get all holdings
	holdings, err := s.GetAllHoldings()
	if err != nil {
		return nil, err
	}
	// Calculate their value
	var USDtotal float64
	USDtotal = 0
	for _, holding := range holdings {
		if holding.USDValue > 0 && holding.LastUpdated.Sub(time.Now().AddDate(0, 0, -1)) < 0 {
			// Update
		}
		USDtotal += holding.USDValue
	}

	// Calculate the % of the portfolio they represent
	var entries []models.PortfolioEntry
	for _, holding := range holdings {
		entry := models.PortfolioEntry{
			Symbol:          holding.Symbol,
			USDValue:        holding.USDValue,
			Units:           holding.UnitsHeld,
			TotalPercentage: holding.USDValue / (USDtotal / 100),
			TargetPp:        holding.TargetPp,
		}
		entries = append(entries, entry)
	}
	portfolio := &models.Portfolio{
		Entries:      entries,
		TotalHolding: USDtotal,
	}
	// Return Symbol, Units, Value, Percentage, TargetPercentage
	return portfolio, nil
}

func (s *PortfolioHoldingService) GetAllHoldings() ([]models.HoldingModel, error) {
	query := `SELECT p.id, a.symbol, a.asset_type, p.units_held, p.usd_value, p.last_updated, p.target_pp FROM portfolio_holding p
			  INNER JOIN asset a ON a.id = p.asset_id`

	var holdings []models.HoldingModel

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query portfolio holdings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var h models.HoldingModel
		if err := rows.Scan(&h.ID, &h.Symbol, &h.AssetType, &h.UnitsHeld, &h.USDValue, &h.LastUpdated, &h.TargetPp); err != nil {
			return nil, fmt.Errorf("Failed to scan row: %v", err)
		}
		holdings = append(holdings, h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %v", err)
	}

	return holdings, nil
}

// CreateAssetObjective inserts a new empty portfolio holding with a percentage objective into the database
func (s *PortfolioHoldingService) CreateAssetObjective(objective *models.AssetObjectiveCreate) error {

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
	query = `SELECT EXISTS(SELECT 1 FROM portfolio_holding WHERE asset_id = ?)`
	err = s.DB.QueryRow(query, assetId).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if asset already has an objective: %v", err)
		return err
	}
	// Save new asset objective
	if exists {
		query = `UPDATE portfolio_holding SET target_pp = ? WHERE asset_id = ?`
		_, err = s.DB.Exec(query, objective.TargetAllocationPercentage, assetId)
		if err != nil {
			log.Printf("Error updating porfolio holding with asset objective: %v", err)
			return err
		}
	} else {
		query = `INSERT INTO portfolio_holding (asset_id, units_held, usd_value, last_updated, target_pp) VALUES (?, ?, ?, ?, ?)`
		_, err = s.DB.Exec(query, assetId, 0, 0, time.Now(), objective.TargetAllocationPercentage)
		if err != nil {
			log.Printf("Error saving asset_objective: %v", err)
			return err
		}
	}
	return nil
}
