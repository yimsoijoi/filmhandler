package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/datamodel"
	"gorm.io/gorm"
)

type CreateCustomerReq struct {
	Name   string `json:"name"`
	Tel    string `json:"tel"`
	Email  string `json:"email"`
	Social string `json:"social_network"`
}

//ใช

func (h *handler) CreateCustomer(c *fiber.Ctx) error {
	var req CreateCustomerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error message": "bad request",
			"error":         err.Error(),
		})
	}

	var checkCust datamodel.Customer
	if tx := h.pg.WithContext(c.Context()).Where("tel = ?", req.Tel).First(&checkCust); errors.Is(
		gorm.ErrRecordNotFound,
		tx.Error) {
		// Try to parse social JSON
		var socialData datamodel.SocialData
		if err := json.Unmarshal([]byte(req.Social), &socialData); err != nil {
			log.Println("failed to unmarshal socialData JSON:", req.Social)
		}
		fmt.Println(socialData)
		customer := datamodel.NewCustomer(
			(req.Name),
			(req.Tel),
			(req.Email),
			(req.Social))
		h.pg.WithContext(c.Context()).Create(&customer)
		return c.Status(201).JSON(customer)
	}

	return c.Status(400).JSON(map[string]interface{}{
		"error": "duplicate tel",
	})
}

// // Want ALL datamodel.Customers - find ALL customers
// customers := []*datamodel.Customer{}
// h.pg.WithContext(c.Context()).Find(&customers)

// // Loop through customers
// for _, cust := range customers {
// 	if cust.Tel == datamodel.Info(req.Tel) {
// 		return c.Status(400).JSON(map[string]interface{}{
// 			"error": "duplicate tel",
// 		})
// 	}
// }
