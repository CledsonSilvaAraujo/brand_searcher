package handlers

import (
	"backend/database"
	"backend/email"
	"backend/google_crawler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IdentifyKeyWords is the handler for identifying key words and sending an email.
// @Summary Identifica palavras-chave de concorrentes
// @Description Recebe termos de marca e um email, e retorna links patrocinados do Google que usam esses termos.
// @Accept  json
// @Produce  json
// @Param   terms  body  string  true  "Termos de Marca"
// @Param   email  body  string  true  "Email para envio do relatório"
// @Success 200 {string} string "Obrigado! O diagnóstico será processado e enviado para o seu email."
// @Failure 400 {string} string "Erro na solicitação"
// @Failure 500 {string} string "Erro no processamento"
// @Router /identify-key-words [post]
func IdentifyKeyWords(c *gin.Context) {
	var request struct {
		Terms string `json:"terms"`
		Email string `json:"email"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := google_crawler.CrawlGoogle(request.Terms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to crawl Google"})
		return
	}

	// Process and send email
	if err := email.DefaultEmailSender.Send(request.Email, "Search Results", results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	// Save results to MongoDB
	database.SaveResultsToMongo(request.Terms, results)

	c.JSON(http.StatusOK, gin.H{"message": "Thank you! The diagnosis will be processed and sent to your email."})
}
