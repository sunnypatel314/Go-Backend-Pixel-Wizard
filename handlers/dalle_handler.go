package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
)

// var client *openai.Client

// func init() {
// 	apiKey := os.Getenv("OPENAI_API_KEY")
// 	if apiKey == "" {
// 		log.Fatal("OPENAI_API_KEY environment variable is not set")
// 	}
// 	client = openai.NewClient(apiKey)
// }

// GenerateImageHandler generates an image using OpenAI's DALL-E
func GenerateImageHandler(c *fiber.Ctx) error {
	var req struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input", "success": false})
	}

	// Initialize OpenAI client (assuming your API key is in an environment variable)
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	// Call OpenAI API to generate an image
	resp, err := client.CreateImage(c.Context(), openai.ImageRequest{
		Model:          "dall-e-2",
		Prompt:         req.Prompt,
		N:              1,
		Size:           "1024x1024",
		ResponseFormat: "b64_json",
		Quality:        "standard",
	})

	if err != nil {
		log.Printf("OpenAI API error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate image", "success": false})
	}

	// Access image data from response
	if len(resp.Data) == 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "No images generated", "success": true})
	}

	image := resp.Data[0].B64JSON
	return c.Status(200).JSON(fiber.Map{"photo": image, "success": true})
}
