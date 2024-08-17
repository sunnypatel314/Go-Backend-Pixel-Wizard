package handlers

import (
	"context"
	"time"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/database"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/models"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LogInHandler(c *fiber.Ctx) error {
	type LogInRequest struct {
		Identifier string `json:"identifier"` // Can be email or username
		Password   string `json:"password"`
	}

	var req LogInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Initialize the user repository
	userRepo := repository.NewUserRepository(database.DB)

	// Find the user by email or username
	user, err := userRepo.FindUserByEmailOrUsername(context.Background(), req.Identifier)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid identifier or password"})
		}
		log.Printf("Error finding user: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid identifier or password"})
	}

	secretKey := os.Getenv("JWT_SECRET") // Store your secret key in an environment variable

	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),                         // Store the user's ID in the token
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expiration time (e.g., 72 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Return the token to the user
	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}

func SignUpHandler(c *fiber.Ctx) error {
	type SignUpRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Initialize the user repository
	userRepo := repository.NewUserRepository(database.DB)

	_, err := userRepo.FindUserByEmailOrUsername(context.Background(), req.Email)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email is already taken"})
	}

	_, err = userRepo.FindUserByEmailOrUsername(context.Background(), req.Username)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Username is already taken"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	newUser := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	_, err = userRepo.CreateUser(context.Background(), newUser)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Return success response
	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}
