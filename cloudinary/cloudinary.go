package cloudinary

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Cloudinary instance
var CLD *cloudinary.Cloudinary

// Init initializes the Cloudinary client
func Init() {
	var err error
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	// Initialize Cloudinary
	CLD, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Error initializing Cloudinary: %v", err)
	}
}

// UploadImage uploads an image to Cloudinary and returns the URL
func UploadImage(filePath string) (string, error) {
	resp, err := CLD.Upload.Upload(context.Background(), filePath, uploader.UploadParams{})
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
