package app

import (
	"log"

	"github.com/theborzet/test_task-binance-/internal/handlers"
	"github.com/theborzet/test_task-binance-/internal/services"

	"github.com/theborzet/test_task-binance-/internal/config"
	"github.com/theborzet/test_task-binance-/internal/db"
	"github.com/theborzet/test_task-binance-/internal/db/repository"
	"github.com/theborzet/test_task-binance-/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Some problems with config: %v", err)
	}

	database := db.Init(config)
	defer func() {
		if err := db.Close(database); err != nil {
			log.Printf("Failed to close DB: %v", err)
		}
	}()

	repo := repository.NewSQLRepository(database)

	service := services.NewTickerService(repo, config)

	handler := handlers.NewTickerHandler(service)

	app := fiber.New()

	routes.RegistrationRoutess(app, handler)

	go service.AdditionTickers()

	if err := app.Listen(config.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
