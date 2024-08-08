package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"net/http"
	"os"

	_ "backend/docs" // Certifique-se de que este caminho está correto e que os arquivos foram gerados pelo swag init

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	database.Connect()

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Handler function that Render will use
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
	router.POST("/identify-key-words", handlers.IdentifyKeyWords)

	// Define the root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go Gin Brand Searcher API!",
		})
	})

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
