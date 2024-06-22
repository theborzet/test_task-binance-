package handlers

import (
	"github.com/theborzet/test_task-binance-/internal/services"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	Ticker string `json:"ticker"`
}

type response struct {
	Ticker     string
	Price      float64
	Difference float64
}

type TickerHandler struct {
	serv *services.TickerService
}

func NewTickerHandler(serv *services.TickerService) *TickerHandler {
	return &TickerHandler{serv: serv}
}

func (h *TickerHandler) AddTicker(ctx *fiber.Ctx) error {

	var req request

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

	resp := response{
		Ticker:     ticker,
		Price:      price,
		Difference: difference,
	}

	response := fiber.Map{
		"Ticker":     resp.Ticker + "/USDT",
		"Price":      resp.Price,
		"Difference": resp.Difference,
	}

	return ctx.JSON(response)
}
