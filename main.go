package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/cloudinary"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/database"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/handlers"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/middleware"
)

func main() {

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize Cloudinary
	cloudinary.Init()

	database.Connect()
	defer func() {
		if err := database.DB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// AUTH ROUTES
	app.Post("/api/v1/auth/log-in", handlers.LogInHandler)
	app.Post("/api/v1/auth/sign-up", handlers.SignUpHandler)

	// POST ROUTES
	auth := app.Group("/api/v1", middleware.AuthMiddleware) // Auth middleware
	auth.Post("/posts", handlers.CreatePostHandler)
	auth.Get("/posts", handlers.GetAllPostsHandler)
	auth.Delete("/posts/:id", handlers.DeletePostHandler)

	// DALLE ROUTES

	PORT := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	app.Listen(":" + PORT)
}
