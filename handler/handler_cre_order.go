package handler

import "github.com/gofiber/fiber/v2"

type CreateOrderReq struct {
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Email    string `json:"e-mail"`
	FilmNo   int    `json:"filmnumber"`
	FilmName string `json:"filmname"`
	FilmType string `json:"filmtype"`
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
	return nil
}
