package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler interface {
	CreateOrder(*fiber.Ctx) error
	CreateCustomer(*fiber.Ctx) error
	GetOrders(*fiber.Ctx) error
	GetCustomers(*fiber.Ctx) error
	UpdateCustomer(*fiber.Ctx) error
	UpdateOrder(*fiber.Ctx) error
	DeleteCustomer(*fiber.Ctx) error
}

type handler struct {
	pg *gorm.DB
}

func New(db *gorm.DB) Handler {
	return &handler{
		pg: db,
	}
}
