package main

import (
	"context"
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
	// Initialize new Fiber instance
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB limit
	})

	// Loads .env file if in development mode; production environment loads env vars automatically
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	// Initialize Cloudinary
	cloudinary.Init()

	// Connect to Mongo Database
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
	app.Post("/api/v1/auth/log-in", handlers.LogInHandler)   // log-in route
	app.Post("/api/v1/auth/sign-up", handlers.SignUpHandler) // sign-in route

	// POST & DALLE ROUTES
	auth := app.Group("/api/v1", middleware.AuthMiddleware) // Auth middleware for Json web token verification
	auth.Post("/posts", handlers.CreatePostHandler)         // creating post route
	auth.Get("/posts", handlers.GetAllPostsHandler)         // fetching all posts route
	auth.Delete("/posts/:id", handlers.DeletePostHandler)   // deleting post route
	auth.Post("/dalle", handlers.GenerateImageHandler)      // generating image route

	app.Listen(":" + os.Getenv("PORT")) // listening to whatever port you defined as the env var (5000, 8000, 8080, etc)

}
