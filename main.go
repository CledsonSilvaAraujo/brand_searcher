package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	database.Connect()
}

// Handler function that Vercel will use
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()

	// Configure CORS to allow requests from the frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URI")},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Define the /api/identify-key-words route
	router.POST("/api/identify-key-words", handlers.IdentifyKeyWords)

	router.ServeHTTP(w, r)
}

func main() {
	// Use the PORT environment variable to define the server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(Handler)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
