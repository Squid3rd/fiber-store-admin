package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/squid3rd/fiber-store-admin/internal/db"
	"github.com/squid3rd/fiber-store-admin/internal/handler"
	"github.com/squid3rd/fiber-store-admin/internal/repository"
	"github.com/squid3rd/fiber-store-admin/internal/router"
	"github.com/squid3rd/fiber-store-admin/internal/service"
)

func main() {
	_ = godotenv.Load()

	mongo, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongo.Close(); err != nil {
			log.Printf("Error closing MongoDB: %v", err)
		}
	}()

	// Repositories
	productRepo := repository.NewProductRepo(mongo.DB.Collection("products"))

	// Services
	productService := service.NewProductService(productRepo)

	// Handlers
	productHandler := handler.NewProductHandler(productService)
	healthHandler := handler.NewHealthHandler()

	// New
	app := fiber.New(
		fiber.Config{
			AppName:   "Fiber Store Admin API",
			BodyLimit: 1024 * 1024 * 10, // 10MB
		},
	)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Register(app, router.Deps{
		Product: productHandler,
		Health:  healthHandler,
	})

	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Fatal(err)
		}
	}()

}
