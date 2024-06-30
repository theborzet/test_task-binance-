package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theborzet/test_task-binance-/internal/handlers"
)

func RegistrationRoutess(app *fiber.App, handler *handlers.TickerHandler) {
	app.Post("/add_ticker", handler.AddTicker)

	app.Get("/fetch", handler.FetchPrice)
}
