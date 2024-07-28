package email

import (
	"gopkg.in/gomail.v2"
)

// EmailSender interface to allow mocking
type EmailSender interface {
	Send(to string, subject string, body string) error
}

// SMTPEmailSender is the real implementation of EmailSender
type SMTPEmailSender struct{}

// Send sends an email with the specified subject and body to the given recipient.
func (s *SMTPEmailSender) Send(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "cledson.silva.teste@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "ciebinho.embetant@gmail.com", "jacz rqjo zxbd pcat")

	return d.DialAndSend(m)
}

// DefaultEmailSender is the default EmailSender used in the application
var DefaultEmailSender EmailSender = &SMTPEmailSender{}

// SendEmail sends an email with the specified subject and body to the given recipient.
func SendEmail(to string, subject string, body string) error {
	return DefaultEmailSender.Send(to, subject, body)
}
