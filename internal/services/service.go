package services

import (
	"log"
	"time"

	"github.com/theborzet/test_task-binance-/internal/config"
	"github.com/theborzet/test_task-binance-/internal/db/repository"
	"github.com/theborzet/test_task-binance-/pkg/binance"
)

// Константы форматов времени для лучшей читаемости кода
const (
	DateFormat      = "02.01.06 15:04:05"
	TimestampFormat = "2006-01-02 15:04:05"
)

type TickerService struct {
	repo   repository.Repository
	cfg    *config.Config
	client *binance.BinanceClient
}

func NewTickerService(repo repository.Repository, cfg *config.Config) *TickerService {
	client := binance.NewBinanceClient(cfg.Binance.APIURL)
	return &TickerService{
		repo:   repo,
		cfg:    cfg,
		client: client,
	}
}

func (s *TickerService) AddTicker(tiсker string) error {
	return s.repo.AddTicker(tiсker)
}

func (s *TickerService) FetchPrice(ticker, dateFrom, dateTo string) (float64, float64, error) {
	dateFromTime, err := time.Parse(DateFormat, dateFrom)
	if err != nil {
		return 0, 0, err
	}

	dateToTime, err := time.Parse(DateFormat, dateTo)
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

// Данную функцию я перенес из app.go в service.go для соблюдения принципа
// единственной ответственности. Эта функция явно относится к уровню
// бизнес-логики и к работе с данными коинов.
func (s *TickerService) AdditionTickers() {
	for {
		tickers, err := s.repo.GetAllTickers()
		if err != nil {
			log.Printf("Some problems with tickers: %e", err)
		}

		for _, ticker := range tickers {
			price, err := s.client.GetTickerPrice(ticker)

			if err != nil {
				log.Printf("Error fetching price for ticker %s: %v\n", ticker, err)
				continue
			}

			tickeID, err := s.repo.GetTickerID(ticker)

			if err != nil {
				log.Printf("Error fetching price for ticker %s: %v\n", ticker, err)
				continue
			}

			err = s.repo.AddTickerPrice(tickeID, price, time.Now().Format(TimestampFormat))
			if err != nil {
				log.Printf("Error adding ticker price for ticker %s: %v\n", ticker, err)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
