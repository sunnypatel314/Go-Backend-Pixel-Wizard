package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/cloudinary"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/database"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/models"
	"github.com/sunnypatel314/Go-Backend-Pixel-Wizard/repository"
)

func GetAllPostsHandler(c *fiber.Ctx) error {
	// Initialize the post repository
	postRepo := repository.NewPostRepository(database.DB)

	// Retrieve all posts from the database
	posts, err := postRepo.GetAllPosts(context.Background())
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch posts", "success": false})
	}

	// Return the posts in the response
	return c.Status(200).JSON(fiber.Map{"data": posts, "success": true})
}

func CreatePostHandler(c *fiber.Ctx) error {
	type CreatePostRequest struct {
		Username string `json:"username"`
		Prompt   string `json:"prompt"`
		Photo    string `json:"photo"`
	}

	var req CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input", "success": false})
	}

	// Upload photo to Cloudinary
	photoURL, err := cloudinary.UploadImage(req.Photo)
	if err != nil {
		log.Printf("Error uploading photo: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error", "success": false})
	}

	// Initialize the post repository
	postRepo := repository.NewPostRepository(database.DB)

	// Initialize the user repository
	userRepo := repository.NewUserRepository(database.DB)

	// Find the user by username (assuming the username is unique)
	user, err := userRepo.FindUserByEmailOrUsername(context.Background(), req.Username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found", "success": false})
	}

	// Create a new post
	newPost := &models.Post{
		CreatorID:   user.ID,
		CreatorName: req.Username,
		Prompt:      req.Prompt,
		PhotoURL:    photoURL,
		CreatedDate: time.Now(),
	}

	_, err = postRepo.CreatePost(context.Background(), newPost)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error", "success": false})
	}

	// Return success response
	return c.Status(201).JSON(fiber.Map{"message": "Post created successfully", "success": true})
}

func DeletePostHandler(c *fiber.Ctx) error {
	postID := c.Params("id")
	if postID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Post ID is required", "success": false})
	}

	// Initialize the post repository
	postRepo := repository.NewPostRepository(database.DB)

	err := postRepo.DeletePost(context.Background(), postID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error", "success": false})
	}
	return c.Status(204).Send(nil)
}
