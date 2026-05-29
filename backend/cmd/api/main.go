package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/squid3rd/fiber-store-admin/docs"
	"github.com/squid3rd/fiber-store-admin/internal/db"
	"github.com/squid3rd/fiber-store-admin/internal/handler"
	"github.com/squid3rd/fiber-store-admin/internal/repository"
	"github.com/squid3rd/fiber-store-admin/internal/router"
	"github.com/squid3rd/fiber-store-admin/internal/service"
)

// @title Fiber Store Admin API
// @version 1.0
// @description Fiber Store Admin API
// @BasePath /api/v1
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
	authRepo := repository.NewAuthRepo(mongo.DB.Collection("users"))

	// Services
	productService := service.NewProductService(productRepo)
	authService := service.NewAuthService(authRepo)
	// Handlers
	productHandler := handler.NewProductHandler(productService)
	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(authService)

	// New
	app := fiber.New(
		fiber.Config{
			AppName:   "Fiber Store Admin API",
			BodyLimit: 1024 * 1024 * 10, // 10MB
		},
	)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,http://localhost:3000,http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Register(app, router.Deps{
		Product: productHandler,
		Health:  healthHandler,
		Auth:    authHandler,
	})

	app.Get("/api/v1/*", swagger.HandlerDefault)
	docs.SwaggerInfo.BasePath = "/api/v1"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
