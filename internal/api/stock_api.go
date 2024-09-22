package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"stock_tracker/internal/models"
)

func FetchHistoricalData(symbol, apiKey string) ([]byte, error) {

	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s",
		symbol, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body %v", err)
	}

	return responseBody, nil
}
func FetchStockData(symbol string) (*models.StockData, error) {

	response, err := FetchHistoricalData(symbol, "apikey")
	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}
	var stockData models.StockData

	err = json.Unmarshal(response, &stockData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &stockData, nil
}
