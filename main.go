package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/database"
)

func main() {

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connect()
	defer func() {
		if err := database.DB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	PORT := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	app.Listen(":" + PORT)
}
