package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
)

type UpdateOrderReq struct {
	OrderUuid string `json:"order_uuid"`
	Tel       string `json:"tel"`
	Email     string `json:"email"`
	FilmNo    int    `json:"film_number"`
	FilmName  string `json:"film_name"`
	FilmType  string `json:"film_type"`
	Status    string `json:"status"`
	Keep      bool   `json:"keep"`
}

func (h *handler) UpdateOrder(c *fiber.Ctx) error {
	var updateOrder UpdateOrderReq
	if err := c.BodyParser(&updateOrder); err != nil {
		return c.SendStatus(404)
	}
	update := h.pg.WithContext(c.Context()).Model(datamodel.Order{}).Where("order_uuid = ?", updateOrder.OrderUuid).Omit("").Updates(map[string]interface{}{
		"tel":         updateOrder.Tel,
		"email":       updateOrder.Email,
		"film_number": updateOrder.FilmNo,
		"film_name":   updateOrder.FilmName,
		"film_type":   updateOrder.FilmType,
		"status":      updateOrder.Status,
		"keep":        updateOrder.Keep,
	})

	return c.Status(200).JSON(update)
}
