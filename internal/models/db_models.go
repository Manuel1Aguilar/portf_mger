package models

import (
	"fmt"
	"time"
)

// Asset model
type Asset struct {
	ID          int    `json:"id" db:"id"`
	Symbol      string `json:"symbol" db:"symbol"`
	Description string `json:"description" db:"description"`
	AssetType   string `json:"asset_type" db:"asset_type"`
}

func (a Asset) String() string {
	return fmt.Sprintf("Asset[ID: %d, Symbol: %s, Description: %s, Type: %s]", a.ID, a.Symbol, a.Description, a.AssetType)
}

// MetricType model
type MetricType struct {
	ID          int    `json:"id" db:"id"`
	TypeName    string `json:"type_name" db:"type_name"` // E.g., 'ma_200w', 'ma_100w'
	Description string `json:"description" db:"description"`
}

func (mt MetricType) String() string {
	return fmt.Sprintf("MetricType[ID: %d, TypeName: %s, Description: %s]", mt.ID, mt.TypeName, mt.Description)
}

// AssetSnapshot model
type AssetSnapshot struct {
	ID           int       `json:"id" db:"id"`
	AssetID      int       `json:"asset_id" db:"asset_id"` // Foreign key to Asset table
	MetricValue  float64   `json:"metric_value" db:"metric_value"`
	MetricTypeID int       `json:"metric_type_id" db:"metric_type_id"` // Foreign key to MetricType table
	CurrentValue float64   `json:"current_value" db:"current_value"`
	DateTaken    time.Time `json:"date_taken" db:"date_taken"`
}

func (as AssetSnapshot) String() string {
	return fmt.Sprintf("AssetSnapshot[ID: %d, AssetID: %d, MetricValue: %.2f, MetricTypeID: %d, CurrentValue: %.2f, DateTaken: %s]",
		as.ID, as.AssetID, as.MetricValue, as.MetricTypeID, as.CurrentValue, as.DateTaken.Format("2006-01-02 15:04:05"))
}

// AssetObjective model
type AssetObjective struct {
	ID                         int     `json:"id" db:"id"`
	AssetID                    int     `json:"asset_id" db:"asset_id"` // Foreign key to Asset table
	TargetAllocationPercentage float64 `json:"target_allocation_percentage" db:"target_allocation_percentage"`
}

func (ao AssetObjective) String() string {
	return fmt.Sprintf("AssetObjective[ID: %d, AssetID: %d, TargetAllocationPercentage: %.2f%%]",
		ao.ID, ao.AssetID, ao.TargetAllocationPercentage)
}

// PortfolioEntry model
type PortfolioEntry struct {
	ID              int     `json:"id" db:"id"`
	AssetID         int     `json:"asset_id" db:"asset_id"` // Foreign key to Asset table
	UnitsHeld       float64 `json:"units_held" db:"units_held"`
	CurrentPosition float64 `json:"current_position" db:"current_position"`
}

func (pe PortfolioEntry) String() string {
	return fmt.Sprintf("PortfolioEntry[ID: %d, AssetID: %d, UnitsHeld: %.2f, CurrentPosition: %.2f]",
		pe.ID, pe.AssetID, pe.UnitsHeld, pe.CurrentPosition)
}

// AssetTransaction model
type AssetTransaction struct {
	ID              int       `json:"id" db:"id"`
	AssetID         int       `json:"asset_id" db:"asset_id"`                 // Foreign key to Asset table
	TransactionType string    `json:"transaction_type" db:"transaction_type"` // BUY or SELL
	ValueUSD        float64   `json:"valueUSD" db:"valueUSD"`
	Units           float64   `json:"units" db:"units"`
	UnitPrice       float64   `json:"unit_price" db:"unit_price"`
	DateTransacted  time.Time `json:"date_transacted" db:"date_transacted"`
}

func (t AssetTransaction) String() string {
	return fmt.Sprintf("AssetTransaction[ID: %d, AssetID: %d, Type: %s, ValueUSD: %.2f, Units: %.2f, UnitPrice: %.2f, DateTransacted: %s]",
		t.ID, t.AssetID, t.TransactionType, t.ValueUSD, t.Units, t.UnitPrice, t.DateTransacted.Format("2006-01-02 15:04:05"))
}
