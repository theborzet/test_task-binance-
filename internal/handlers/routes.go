package handlers

import (
	"test_api/internal/db/repository"
	"test_api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RegistrationRoutess(app *fiber.App, db *sqlx.DB) {
	repo := repository.NewSQLRepository(db)
	tickerServ := *services.NewTickerService(repo)

	handler := NewTickerHandler(tickerServ)

	app.Post("/add_ticker", handler.AddTicker)

	app.Get("/fetch", handler.FetchPrice)
}
