package repository

import (
	"fmt"

	"github.com/thongsoi/testd/internal/model"
)

var markets = []model.Market{
	{ID: 1, Name: "Market 1"},
	{ID: 2, Name: "Market 2"},
}

var submarkets = []model.Submarket{
	{ID: 1, Name: "Submarket 1", MarketID: 1},
	{ID: 2, Name: "Submarket 2", MarketID: 1},
	{ID: 3, Name: "Submarket 3", MarketID: 2},
}

type Repository struct{}

func (r *Repository) GetMarkets() []model.Market {
	return markets
}

func (r *Repository) GetSubmarketsByMarketID(marketID int) []model.Submarket {
	var result []model.Submarket
	for _, submarket := range submarkets {
		if submarket.MarketID == marketID {
			result = append(result, submarket)
		}
	}
	return result
}

func (r *Repository) CreateOrder(order model.Order) error {
	fmt.Printf("Order created: %+v\n", order)
	return nil
}
