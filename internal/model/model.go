package model

type Market struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Submarket struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	MarketID int    `json:"market_id"`
}

type Order struct {
	MarketID    int `json:"market_id"`
	SubmarketID int `json:"submarket_id"`
}
