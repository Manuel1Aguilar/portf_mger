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
