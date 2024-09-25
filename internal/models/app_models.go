package models

import (
	"fmt"
)

type MovingAverage200Weeks struct {
	Stock     string  `json:"stock"`
	Value     float64 `json:"value"`
	CurrValue float64 `json:"currValue"`
	From      string  `json:"from"`
	To        string  `json:"to"`
}

func (m MovingAverage200Weeks) String() string {
	return fmt.Sprintf("Stock: %s,\nValue: %f,\nFrom: %s,\nTo: %s,\nCurrent Value: %f",
		m.Stock, m.Value, m.From, m.To, m.CurrValue)
}
