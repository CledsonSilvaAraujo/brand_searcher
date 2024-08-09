package email

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

type EmailSender interface {
	Send(to string, subject string, body string) error
}

type SMTPEmailSender struct{}

func (s *SMTPEmailSender) Send(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_ADDRESS"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_APP_PASSWORD"))

	return d.DialAndSend(m)
}

var DefaultEmailSender EmailSender = &SMTPEmailSender{}

func SendEmail(to string, subject string, body string) error {
	return DefaultEmailSender.Send(to, subject, body)
}
