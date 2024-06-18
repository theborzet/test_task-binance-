package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test_api/internal/config"
	"test_api/internal/db/repository"
	"time"
)

type TickerService struct {
	repo repository.Repository
	cfg  *config.Config
}

func NewTickerService(repo repository.Repository, cfg *config.Config) *TickerService {
	return &TickerService{
		repo: repo,
		cfg:  cfg}
}

func (s *TickerService) AddTicker(tiker string) error {
	return s.repo.AddTicker(tiker)
}

func (s *TickerService) FetchPrice(ticker, dateFrom, dateTo string) (float64, float64, error) {
	dateFromTime, err := time.Parse("02.01.06 15:04:05", dateFrom)
	if err != nil {
		return 0, 0, err
	}

	dateToTime, err := time.Parse("02.01.06 15:04:05", dateTo)
	if err != nil {
		return 0, 0, err
	}

	last_price, first_price, err := s.repo.GetPriceAndDifference(ticker, dateFromTime, dateToTime)
	if err != nil {
		return 0, 0, err
	}

	difference := ((last_price - first_price) / first_price) * 100

	return last_price, difference, nil

}

func (s *TickerService) GetTickerPriceFromBinance(ticker string) (float64, error) {
	url := fmt.Sprintf("%s%sUSDT", s.cfg.BinanceAPIURL, ticker)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Price string `json:"price"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
