package tests

import (
	"backend/email"
	"testing"
)

type MockEmailSender struct{}

func (m *MockEmailSender) Send(to string, subject string, body string) error {
	return nil
}

func TestSendEmail(t *testing.T) {
	email.DefaultEmailSender = &MockEmailSender{}

	err := email.DefaultEmailSender.Send("test@gmail.com", "Test Subject", "Test Body")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
