package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

func (h *handler) GetCustomers(c *fiber.Ctx) error {
	var custs []*datamodel.Customer
	h.pg.WithContext(c.Context()).Find(&custs)
	if len(custs) == 0 {
		return c.Status(404).JSON(map[string]interface{}{
			"error messege": "empty customers",
		})
	}
	return c.Status(200).JSON(custs)
}
