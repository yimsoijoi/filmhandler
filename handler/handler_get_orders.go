package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

func (h *handler) GetOrders(c *fiber.Ctx) error {
	var orders []*datamodel.Order
	h.pg.WithContext(c.Context()).Preload("Customer").Find(&orders)
	if len(orders) == 0 {
		return c.Status(404).JSON(map[string]interface{}{
			"error messege": "empty orders",
		})
	}
	return c.Status(200).JSON(orders)
}
