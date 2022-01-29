package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

func (h *handler) DeleteCustomer(c *fiber.Ctx) error {
	target := c.Params("tel")
	var dummy datamodel.Customer
	tx := h.pg.WithContext(c.Context()).Where("tel = ?", target).Delete(&dummy)
	if tx.Error != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
