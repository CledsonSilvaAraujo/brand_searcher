package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Conecte-se ao MongoDB
	database.Connect()

	router := gin.Default()

	// Configurar CORS para permitir requisições do frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URI")},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/identify-key-words", handlers.IdentifyKeyWords)

	// Usar a variável de ambiente PORT para definir a porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
