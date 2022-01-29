package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/filmhandler/config"
	"github.com/yimsoijoi/filmhandler/handler"
	"github.com/yimsoijoi/filmhandler/lib/postgres"
)

func main() {
	conf, err := config.Load()
	db, err := postgres.New(conf.Postgres)
	if err != nil {
		log.Println("db connect failed:", err.Error())
	}
	migrate(db, conf.Postgres)

	handler := handler.New(db)
	app := fiber.New()
	orderAPI := app.Group("/orders")
	custAPI := app.Group("/customers")
	custAPI.Post("/", handler.CreateCustomer)
	custAPI.Get("/", handler.GetCustomers)
	custAPI.Patch("/:tel", handler.UpdateCustomer)
	custAPI.Delete("/:tel", handler.DeleteCustomer)
	orderAPI.Post("/", handler.CreateOrder)
	orderAPI.Get("/", handler.GetOrders)
	orderAPI.Patch("/:uuid", handler.UpdateOrder)

	log.Fatal(app.Listen(":8000"))
}
