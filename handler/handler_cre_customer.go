package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
	"gorm.io/gorm"
)

type CreateCustomerReq struct {
	Name   string `json:"name"`
	Tel    string `json:"tel"`
	Email  string `json:"e-mail"`
	Social string `json:"social_network"`
}

func (h *handler) CreateCustomer(c *fiber.Ctx) error {
	var req CreateCustomerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error message": "bad request",
			"error":         err.Error(),
		})
	}

	customer := datamodel.NewCustomer(datamodel.Info(req.Name), datamodel.Info(req.Tel), datamodel.Info(req.Email), datamodel.Info(req.Social))
	customers := []*datamodel.Customer{}
	for _, tel := range customers {
		var _tel datamodel.Customer
		findtel := h.pg.WithContext(c.Context()).Where("tel = ?", tel.Tel).First(&_tel)
		if errors.Is(gorm.ErrRecordNotFound, findtel.Error) {
			h.pg.WithContext(c.Context()).Create(&customer)
		} else {
			return c.Status(400).JSON(map[string]interface{}{
				"error message": "duplicate tel",
			})
		}

	}

	return c.Status(201).JSON(customer)
}
