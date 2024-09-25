package models

import (
	"fmt"
)

type StockApiResponse struct {
	MetaData         MetaData                    `json:"Meta Data"`
	WeeklyTimeSeries map[string]WeeklyTimeSeries `json:"Weekly Time Series"`
}

type MetaData struct {
	Info          string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type WeeklyTimeSeries struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func (s StockApiResponse) String() string {
	return fmt.Sprintf("Meta Data:\n %s,\n"+
		"WeeklyTimeSeries: %d entries,\nLast Entry: \n %s",
		s.MetaData, len(s.WeeklyTimeSeries), s.WeeklyTimeSeries[s.MetaData.LastRefreshed])
}
func (m MetaData) String() string {
	return fmt.Sprintf("Info: %s,\n Symbol: %s,\n Last Refreshed: %s,\n Time Zone: %s",
		m.Info, m.Symbol, m.LastRefreshed, m.TimeZone)
}
func (w WeeklyTimeSeries) String() string {
	return fmt.Sprintf("Open: %s,\n High: %s,\n Low: %s,\n Close: %s,\n Volume: %s",
		w.Open, w.High, w.Low, w.Close, w.Volume)
}
