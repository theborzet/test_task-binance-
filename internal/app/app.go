package app

import (
	"log"
	"test_api/internal/config"
	"test_api/internal/db"
	"test_api/internal/db/repository"
	"test_api/internal/handlers"
	"test_api/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("Some problems with config: %v", err)
	}

	database := db.Init(config)
	defer func() {
		if err := db.Close(database); err != nil {
			log.Printf("Failed to close DB: %v", err)
		}
	}()

	repo := repository.NewSQLRepository(database)
	service := services.NewTickerService(repo)

	app := fiber.New()

	handlers.RegistrationRoutess(app, database)

	go func() {
		for {
			tickers, err := repo.GetAllTickers()
			if err != nil {
				log.Printf("Some problems with tickers: %e", err)
			}

			for _, ticker := range tickers {
				price, err := service.GetTickerPriceFromBinance(ticker)

				if err != nil {
					log.Printf("Error fetching price for ticker %s: %v\n", ticker, err)
					continue
				}

				tickeID, err := repo.GetTickerID(ticker)

				if err != nil {
					log.Printf("Error fetching price for ticker %s: %v\n", ticker, err)
					continue
				}

				err = repo.AddTickerPrice(tickeID, price, time.Now().Format("2006-01-02 15:04:05"))
				if err != nil {
					log.Printf("Error adding ticker price for ticker %s: %v\n", ticker, err)
				}
			}

			time.Sleep(1 * time.Minute)
		}
	}()

	if err := app.Listen(config.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
