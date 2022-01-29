package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

type UpdateCustomerReq struct {
	Tel    string `json:"tel"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Social string `json:"social_network"`
}

func (h *handler) UpdateCustomer(c *fiber.Ctx) error {
	target := c.Params("tel")
	var updateCust UpdateCustomerReq
	if err := c.BodyParser(&updateCust); err != nil {
		return c.SendStatus(404)
	}
	update := h.pg.WithContext(c.Context()).Model(datamodel.Customer{}).Where("tel = ?", target).Updates(map[string]interface{}{
		"tel":            updateCust.Tel,
		"name":           updateCust.Name,
		"email":          updateCust.Email,
		"social_network": updateCust.Social,
	})

	return c.Status(200).JSON(update)
}

// func updateCust(u UpdateCustomerReq)
