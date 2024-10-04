package models

import (
	"fmt"
)

type StockApiResponseAdjusted struct {
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

func (s StockApiResponseAdjusted) String() string {
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
