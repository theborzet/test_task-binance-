package handlers

import (
	"test_api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RegistrationRoutess(app *fiber.App, serv *services.TickerService) {
	handler := NewTickerHandler(serv)

	app.Post("/add_ticker", handler.AddTicker)

	app.Get("/fetch", handler.FetchPrice)
}
