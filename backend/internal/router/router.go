package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/squid3rd/fiber-store-admin/internal/handler"
)

type Deps struct {
	Product *handler.ProductHandler
	Health  *handler.HealthHandler
	Auth    *handler.AuthHandler
}

// s
func Register(app *fiber.App, d Deps) {

	api := app.Group("/api/v1")
	api.Get("/health", d.Health.Check)
	api.Get("/products", d.Product.FindAll)
	api.Post("/products", d.Product.Create)
	api.Post("/users", d.Auth.CreateUser)
}
