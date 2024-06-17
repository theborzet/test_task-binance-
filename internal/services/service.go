package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test_api/internal/db/repository"
	"time"
)

type TickerService struct {
	repo repository.Repository
}

func NewTickerService(repo repository.Repository) *TickerService {
	return &TickerService{repo: repo}
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

	return s.repo.GetPriceAndDifference(ticker, dateFromTime, dateToTime)
}

func (s *TickerService) GetTickerPriceFromBinance(ticker string) (float64, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", ticker)
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
