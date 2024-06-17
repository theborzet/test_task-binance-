package handlers

import (
	"test_api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TickerHandler struct {
	serv services.TickerService
}

func NewTickerHandler(serv services.TickerService) *TickerHandler {
	return &TickerHandler{serv: serv}
}

func (h *TickerHandler) AddTicker(ctx *fiber.Ctx) error {
	var req struct {
		Ticker string `json:"ticker"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON payload",
		})
	}

	if err := h.serv.AddTicker(req.Ticker); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add ticker",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Ticker added successfully",
	})
}

func (h *TickerHandler) FetchPrice(ctx *fiber.Ctx) error {
	ticker := ctx.Query("ticker")
	dateFrom := ctx.Query("date_from")
	dateTo := ctx.Query("date_to")

	price, difference, err := h.serv.FetchPrice(ticker, dateFrom, dateTo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch price",
		})
	}

	response := fiber.Map{
		"Ticker":     ticker + "/USDT",
		"Price":      price,
		"Difference": difference,
	}

	return ctx.JSON(response)
}
