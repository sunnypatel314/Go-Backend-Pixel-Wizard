package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/cloudinary"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/database"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/handlers"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/middleware"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB limit
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		fmt.Println(err)
	}

	// Initialize Cloudinary
	cloudinary.Init()

	database.Connect()
	defer func() {
		if err := database.DB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// AUTH ROUTES
	app.Post("/api/v1/auth/log-in", handlers.LogInHandler)
	app.Post("/api/v1/auth/sign-up", handlers.SignUpHandler)

	// POST & DALLE ROUTES
	auth := app.Group("/api/v1", middleware.AuthMiddleware) // Auth middleware
	auth.Post("/posts", handlers.CreatePostHandler)
	auth.Get("/posts", handlers.GetAllPostsHandler)
	auth.Delete("/posts/:id", handlers.DeletePostHandler)
	auth.Post("/dalle", handlers.GenerateImageHandler)

	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)
	app.Listen(":" + os.Getenv("PORT"))

}
