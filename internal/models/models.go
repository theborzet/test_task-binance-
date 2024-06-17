package models

type Ticker struct {
	ID     int    `json:"id"`
	Ticker string `json:"ticker"`
}

type TickerPrice struct {
	ID        int     `json:"id"`
	TickerID  int     `json:"ticker_id"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
}
