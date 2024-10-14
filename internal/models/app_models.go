package models

import (
	"fmt"
	"time"
)

type MovingAverage200Weeks struct {
	Stock     string  `json:"stock"`
	MAValue   float64 `json:"value"`
	CurrValue float64 `json:"currValue"`
	From      string  `json:"from"`
	To        string  `json:"to"`
}

func (m MovingAverage200Weeks) String() string {
	return fmt.Sprintf("200 Weeks Moving Average:\nStock: %s,\nMAValue: %f,\nFrom: %s,\nTo: %s,\nCurrent Value: %f",
		m.Stock, m.MAValue, m.From, m.To, m.CurrValue)
}

type AssetObjectiveCreate struct {
	Symbol                     string  `json:"symbol"`
	TargetAllocationPercentage float64 `json:"target_allocation_percentage"`
}

type HoldingModel struct {
	ID          int       `json:"id" db:"id"`
	Symbol      string    `json:"symbol" db:"symbol"`
	AssetType   string    `json:"asset_type" db:"asset_type"`
	UnitsHeld   float64   `json:"units_held" db:"units_held"`
	USDValue    float64   `json:"usd_value" db:"usd_value"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
}

type PortfolioEntry struct {
	Symbol              string  `json:"symbol" db:"symbol"`
	USDValue            float64 `json:"usd_value" db:"usd_value"`
	Units               float64 `json:"units" db:"units"`
	TotalPercentage     float64 `json:"total_percentage" db:"total_percentage"`
	ObjectivePercentage float64 `json:"objective_percentage" db:"objective_percentage"`
}

type Portfolio struct {
	Entries      []PortfolioEntry `json:"entries"`
	TotalHolding float64          `json:"total_holding"`
}

type AssetTransactionCreate struct {
	Symbol    string  `json:"symbol"`
	Type      string  `json:"type"`
	ValueUSD  float64 `json:"valueUSD"`
	Units     float64 `json:"units"`
	UnitPrice float64 `json:"unit_price"`
}

type AssetLatestValue struct {
	Symbol string  `json:"symbol"`
	Value  float64 `json:"value"`
}

func (p Portfolio) String() string {
	return fmt.Sprintf("Portfolio:\n Entries: %s,\n Total: %.2f", p.Entries, p.TotalHolding)
}
func (pe PortfolioEntry) String() string {
	return fmt.Sprintf("PortfolioEntry: [Symbol: %s, USDValue: %.2f, Units: %.2f, TotalPercentage: %.2f, ObjectivePercentage: %.2f",
		pe.Symbol, pe.USDValue, pe.Units, pe.TotalPercentage, pe.ObjectivePercentage)
}

func (alv AssetLatestValue) String() string {
	return fmt.Sprintf("AssetLatestValue:\nSymbol: %s,\nValue: %.2f",
		alv.Symbol, alv.Value)
}

func (aoc AssetObjectiveCreate) String() string {
	return fmt.Sprintf("AssetObjectiveCreate:\nSymbol: %s,\nTarget Allocation percentage: %f",
		aoc.Symbol, aoc.TargetAllocationPercentage)
}

func (atc AssetTransactionCreate) String() string {
	return fmt.Sprintf("AssetTransactionCreate:\nSymbol: %s,\nType: %s,\nValueUSD: %.2f,\nUnits: %.2f,\n"+
		"UnitPrice: %.2f",
		atc.Symbol, atc.Type, atc.ValueUSD, atc.Units, atc.UnitPrice)
}

func (hm HoldingModel) String() string {
	return fmt.Sprintf("HoldingModel:\nId: %d,\nSymbol: %s,\nAssetType: %s,\nUnitsHeld: %.2f,\nUSDValue: %.2f,\nLastUpdated: %s",
		hm.ID, hm.Symbol, hm.AssetType, hm.UnitsHeld, hm.USDValue, hm.LastUpdated)
}
