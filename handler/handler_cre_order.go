package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

type CreateOrderReq struct {
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	FilmNo   int    `json:"film_number"`
	FilmName string `json:"film_name"`
	FilmType string `json:"film_type"`
	Status   string `json:"status"`
	Keep     bool   `json:"keep"`
}

func (h *handler) CreateOrder(c *fiber.Ctx) error {
	var req CreateOrderReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error message": "bad request",
			"error":         err.Error(),
		})
	}
	order := datamodel.NewOrder(
		datamodel.Info(req.Tel),
		req.FilmNo,
		req.FilmName,
		req.FilmType,
		req.Status,
		req.Keep)

	h.pg.WithContext(c.Context()).Create(&order)
	return c.Status(201).JSON(order)
}
