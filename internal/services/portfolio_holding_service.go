package services

import (
	"database/sql"
	"log"

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
func (s *PortfolioHoldingService) FinishTransaction(tModel *models.AssetTransaction) error {
	// Check if portfolio_holding for asset exists
	exists, err := s.ExistsByAssetId(tModel.AssetID)
	if err != nil {
		return err
	}
	if exists {
		// if exists update
		holding, err := s.GetByAssetId(tModel.AssetID)
		if err != nil {
			return err
		}
		if tModel.TransactionType == models.TransactionTypeBuy {
			holding.UnitsHeld += tModel.Units
			holding.UsdValue += tModel.ValueUSD
		} else {
			holding.UnitsHeld -= tModel.Units
			holding.UsdValue -= tModel.ValueUSD
		}
		err = s.Update(holding)
		if err != nil {
			return err
		}

		return nil
	}

	//if not create
	holding := &models.PortfolioHolding{
		AssetID:   tModel.AssetID,
		UnitsHeld: tModel.Units,
		UsdValue:  tModel.ValueUSD,
	}

	err = s.Add(holding)
	if err != nil {
		return err
	}

	return nil
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

// Get holding by asset id
func (s *PortfolioHoldingService) GetByAssetId(assetId int) (*models.PortfolioHolding, error) {
	query := `SELECT p.id, p.asset_id, p.units_held, p.usd_value FROM portfolio_holding p WHERE p.asset_id = ? LIMIT 1`
	var holding models.PortfolioHolding
	err := s.DB.QueryRow(query, assetId).Scan(&holding.ID, &holding.AssetID, &holding.UnitsHeld, &holding.UsdValue)
	if err != nil {
		log.Printf("Error getting holding by asset id: %v \n", err)
		return nil, err
	}
	return &holding, nil
}

// Add holding
func (s *PortfolioHoldingService) Add(pModel *models.PortfolioHolding) error {
	query := `INSERT INTO portfolio_holding (asset_id, units_held, usd_value) values (?, ?, ?)`

	_, err := s.DB.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.UsdValue)
	if err != nil {
		log.Printf("Error inserting portfolio_holding: %v \n", err)
		return err
	}
	return nil
}

// Update holding
func (s *PortfolioHoldingService) Update(pModel *models.PortfolioHolding) error {
	query := `UPDATE portfolio_holding SET asset_id = ?, units_held = ?, usd_value = ? WHERE id = ?`
	_, err := s.DB.Exec(query, pModel.AssetID, pModel.UnitsHeld, pModel.UsdValue, pModel.ID)
	if err != nil {
		log.Printf("Error updating portfolio_holding: %v \n", err)
		return err
	}
	return nil
}
