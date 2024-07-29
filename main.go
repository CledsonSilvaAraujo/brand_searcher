package main

import (
	"backend/database"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conecte-se ao MongoDB
	database.Connect()

	router := gin.Default()

	router.POST("/identify-key-words", handlers.IdentifyKeyWords)

	router.Run(":8080")
}
