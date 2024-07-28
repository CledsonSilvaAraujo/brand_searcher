package main

import (
	"backend/email"
	"backend/handlers"
	"log"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Monitoramento de Concorrentes
// @version 1.0
// @description Esta é a API para monitorar concorrentes usando termos de marca no Google.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.POST("/identify-key-words", handlers.IdentifyKeyWords)

	router.POST("/send-email", func(c *gin.Context) {
		var emailRequest struct {
			To      string `json:"to"`
			Subject string `json:"subject"`
			Body    string `json:"body"`
		}
		if err := c.BindJSON(&emailRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := email.DefaultEmailSender.Send(emailRequest.To, emailRequest.Subject, emailRequest.Body); err != nil {
			log.Printf("Could not send email: %v", err)
			c.JSON(500, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(200, gin.H{"message": "Email sent successfully"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
