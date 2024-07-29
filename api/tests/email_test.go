package tests

import (
	"backend/email"
	"testing"
)

// MockEmailSender is a mock implementation of EmailSender
type MockEmailSender struct{}

// Send is the mock implementation of Send
func (m *MockEmailSender) Send(to string, subject string, body string) error {
	return nil
}

func TestSendEmail(t *testing.T) {
	// Use the mock sender for testing
	email.DefaultEmailSender = &MockEmailSender{}

	err := email.DefaultEmailSender.Send("test@gmail.com", "Test Subject", "Test Body")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
