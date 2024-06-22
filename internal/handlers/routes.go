package handlers

import (
	"github.com/theborzet/test_task-binance-/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RegistrationRoutess(app *fiber.App, serv *services.TickerService) {
	handler := NewTickerHandler(serv)

	app.Post("/add_ticker", handler.AddTicker)

	app.Get("/fetch", handler.FetchPrice)
}
