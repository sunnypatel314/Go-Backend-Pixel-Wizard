# Go-Backend-Pixel-Wizard

A backend API for handling user authentication, AI image generation, and post management, built with Golang, Fiber, MongoDB, and integrated with Cloudinary for image storage and the OpenAI API for AI image generation.

Github repository for the React frontend: 
    ```bash
    https://github.com/sunnypatel314/React-Frontend-Pixel-Wizard


## Features
- User authentication (sign up, log in)
- AI image generation using OpenAI's DALL-E
- Create, read, and manage posts
- Image upload and storage with Cloudinary

## Tech Stack
- **Golang**: Modern, compiled programming language used for backend development.
- **Fiber**: Web framework for building the REST API.
- **MongoDB**: NoSQL database to store user and post data.
- **Cloudinary**: Cloud service for managing image storage.
- **OpenAI API**: For generating AI images based on prompts.

## Getting Started

### Prerequisites
Before running the application, ensure you have the following installed:
- Go 
- MongoDB 
- Cloudinary account (for image uploads)
- OpenAI API key (for AI image generation)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/sunnypatel314/Go-Backend-Pixel-Wizard.git
   cd Go-Backend-Pixel-Wizard
2. **Set up environment variables:**

   Create a `.env` file in the root directory and add the following environment variables:

   ```bash
   PORT=8080
   MONGO_URI=your_mongo_connection_string
   MONGO_DB_NAME=your_mongo_db_name
   CLOUDINARY_CLOUD_NAME=your_cloudinary_cloud_name
   CLOUDINARY_API_KEY=your_cloudinary_api_key
   CLOUDINARY_API_SECRET=your_cloudinary_api_secret
   OPENAI_API_KEY=your_openai_api_key
3. **Install dependencies:**

   Ensure all dependencies are installed by running the following command:

   ```bash
   go mod tidy
4. **Run the application:**

   Start the API with the following command:

   ```bash
   go run main.go

### API Endpoints

#### Authentication
- **Sign Up**: `POST /api/v1/auth/sign-up`
  - Request body: 
    ```json
    {
      "username": "string",
      "email": "string",
      "password": "string"
    }
    ```
  - Response: 201 Created

- **Log In**: `POST /api/v1/auth/log-in`
  - Request body: 
    ```json
    {
      "identifier": "string",
      "password": "string"
    }
    ```
  - Response: 200 OK with JWT token

#### Posts
- **Create Post**: `POST /api/v1/posts`
  - Requires authentication (Bearer token)
  - Request body: 
    ```json
    {
      "username": "string",
      "prompt": "string",
      "photo": "string"
    }
    ```
  - Response: 201 Created

- **Delete Post by ID**: `DELETE /api/v1/posts/:id`
  - Response: 204 No Content

- **Get All Posts**: `GET /api/v1/posts`
  - Response: 200 OK with a list of posts

#### AI Image Generation
- **Generate Image**: `POST /api/v1/dalle`
  - Requires authentication (Bearer token)
  - Request body: 
    ```json
    {
      "prompt": "string"
    }
    ```
  - Response: 200 OK with generated image


