package models

import (
	"fmt"
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

func (aoc AssetObjectiveCreate) String() string {
	return fmt.Sprintf("AssetObjectiveCreate:\nSymbol: %s,\nTarget Allocation percentage: %f",
		aoc.Symbol, aoc.TargetAllocationPercentage)
}

type AssetTransactionCreate struct {
	Symbol    string  `json:"symbol"`
	Type      string  `json:"type"`
	ValueUSD  float64 `json:"valueUSD"`
	Units     float64 `json:"units"`
	UnitPrice float64 `json:"unit_price"`
}

func (atc AssetTransactionCreate) String() string {
	return fmt.Sprintf("AssetTransactionCreate:\nSymbol: %s,\nType: %s,\nValueUSD: %.2f,\nUnits: %.2f,\n"+
		"UnitPrice: %.2f",
		atc.Symbol, atc.Type, atc.ValueUSD, atc.Units, atc.UnitPrice)
}
