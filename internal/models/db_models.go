package models

type Stock struct {
	ID          int     `json:"id"`
	StockSymbol string  `json:"stock_symbol"`
	LatestMA    float64 `json:"latest_ma"`
	LastUpdated string  `json:"last_updated"`
}
