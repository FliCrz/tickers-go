package models

// Ticker ...
type Ticker struct {
	Coin string `json:"coin"`
	Currency string `json:"currency"`
	Symbol   string `json:"symbol"`
	BidPrice float64 `json:"bidPrice"`
	BidQty   float64 `json:"bidQty"`
	AskPrice float64 `json:"askPrice"`
	AskQty   float64 `json:"askQty"`
	Exchange string `json:"exchange"`
	Timestamp int `json:"timestamp"`
}