package main

import (
	"backend/database"
	"backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conecte-se ao MongoDB
	database.Connect()

	router := gin.Default()

	// Configurar CORS para permitir requisições do frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/identify-key-words", handlers.IdentifyKeyWords)

	router.Run(":8080")
}
