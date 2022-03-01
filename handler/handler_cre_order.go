package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
	"gorm.io/gorm"
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
	var checkCust datamodel.Customer
	if tx := h.pg.WithContext(c.Context()).Where("tel = ?", req.Tel).First(&checkCust); errors.Is(gorm.ErrRecordNotFound, tx.Error) {
		return c.Status(400).Send([]byte("customer not found"))
	}
	order := datamodel.NewOrder(
		(req.Tel),
		req.FilmNo,
		req.FilmName,
		req.FilmType,
		req.Status,
		req.Keep,
	)

	if tx := h.pg.WithContext(c.Context()).Create(&order); tx.Error != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"error": "failed to create order",
		})
	}
	return c.Status(201).JSON(order)
}
