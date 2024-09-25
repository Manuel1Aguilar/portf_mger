package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"

	"stock_tracker/internal/config"
	"stock_tracker/internal/models"

	"github.com/joho/godotenv"
)

func FetchHistoricalData(symbol string) ([]byte, error) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")

	url := fmt.Sprintf(
		"%s?function=TIME_SERIES_WEEKLY&symbol=%s&apikey=%s",
		config.AlphaVantageAPIBaseUrl, symbol, apiKey,
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

func FetchStockData(symbol string) (*models.StockApiResponse, error) {

	response, err := FetchHistoricalData(symbol)
	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}
	var stockData models.StockApiResponse

	err = json.Unmarshal(response, &stockData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &stockData, nil
}

func Get200WeekMovingAverage(symbol string) (*models.MovingAverage200Weeks, error) {
	data, err := FetchStockData(symbol)
	if err != nil {
		return nil, err
	}

	var dates []string

	for date := range data.WeeklyTimeSeries {
		dates = append(dates, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	count := 0
	sum := 0.0
	for _, date := range dates {
		if count >= 200 {
			break
		}
		closeStr := data.WeeklyTimeSeries[date].Close
		closePrice, err := strconv.ParseFloat(closeStr, 64)
		if err != nil {
			fmt.Printf("Error parsing close price: %s", closeStr)
			continue
		}

		sum += closePrice
		count++
	}

	currentValueStr := data.WeeklyTimeSeries[dates[0]].Close
	currentValue, err := strconv.ParseFloat(currentValueStr, 64)
	if err != nil {
		fmt.Printf("Error while parsing closing value: %s \n", currentValueStr)
	}
	ma := sum / float64(count)
	res := &models.MovingAverage200Weeks{
		Stock:     symbol,
		Value:     ma,
		From:      dates[199],
		To:        dates[0],
		CurrValue: currentValue,
	}
	return res, nil
}
