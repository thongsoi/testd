package service

import (
	"github.com/thongsoi/testd/internal/repository"

	"github.com/thongsoi/testd/internal/model"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMarkets() []model.Market {
	return s.repo.GetMarkets()
}

func (s *Service) GetSubmarketsByMarketID(marketID int) []model.Submarket {
	return s.repo.GetSubmarketsByMarketID(marketID)
}

func (s *Service) CreateOrder(order model.Order) error {
	return s.repo.CreateOrder(order)
}
