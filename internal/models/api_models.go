package models

import (
	"fmt"
)

type WeeklyAdjustedResponse struct {
	MetaData                 MetaData                            `json:"Meta Data"`
	WeeklyTimeSeriesAdjusted map[string]WeeklyTimeSeriesAdjusted `json:"Weekly Adjusted Time Series"`
}

type MetaData struct {
	Info          string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type WeeklyTimeSeriesAdjusted struct {
	Open          string `json:"1. open"`
	High          string `json:"2. high"`
	Low           string `json:"3. low"`
	Close         string `json:"4. close"`
	AdjustedClose string `json:"5. adjusted close"`
	Volume        string `json:"6. volume"`
	DividendAmt   string `json:"7. dividend amount"`
}
type WeeklyTimeSeries struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type GlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

func (gq GlobalQuote) String() string {
	return fmt.Sprintf("GlobalQuote:\nSymbol: %s,\nOpen: %s,\nHigh: %s,\nLow: %s,\nPrice: %s,\n"+
		"Volume: %s,\nLatestTradingDay: %s,\nPreviousClose: %s,\nChange: %s,\nChangePercent: %s",
		gq.Symbol, gq.Open, gq.High, gq.Low, gq.Price, gq.Volume, gq.LatestTradingDay, gq.PreviousClose,
		gq.Change, gq.ChangePercent)
}

func (s WeeklyAdjustedResponse) String() string {
	return fmt.Sprintf("Meta Data:\n %s,\n"+
		"WeeklyTimeSeriesAdjusted: %d entries,\nLast Entry: \n %s",
		s.MetaData, len(s.WeeklyTimeSeriesAdjusted), s.WeeklyTimeSeriesAdjusted[s.MetaData.LastRefreshed])
}
func (m MetaData) String() string {
	return fmt.Sprintf("Info: %s,\n Symbol: %s,\n Last Refreshed: %s,\n Time Zone: %s",
		m.Info, m.Symbol, m.LastRefreshed, m.TimeZone)
}
func (w WeeklyTimeSeries) String() string {
	return fmt.Sprintf("Open: %s,\n High: %s,\n Low: %s,\n Close: %s,\n Volume: %s",
		w.Open, w.High, w.Low, w.Close, w.Volume)
}

func (w WeeklyTimeSeriesAdjusted) String() string {
	return fmt.Sprintf("Open: %s,\n High: %s,\n Low: %s,\n Close: %s,\n Adjusted Close: %s,\n Volume: %s,\n Dividend Amount: %s",
		w.Open, w.High, w.Low, w.Close, w.AdjustedClose, w.Volume, w.DividendAmt)
}
