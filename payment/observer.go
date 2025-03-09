package payment

import (
	"fmt"
	"time"
)

// PaymentEvent represents the payment event data
type PaymentEvent struct {
	PaymentID string
	Amount    float64
	Method    string
	Status    string
	TimeStamp time.Time
}

// Observer interface for payment notifications
type PaymentObserver interface {
	Update(event PaymentEvent)
}

// Email Observer
type EmailNotifier struct {
	EmailAddress string
}

func NewEmailNotifier(email string) *EmailNotifier {
	return &EmailNotifier{EmailAddress: email}
}

func (e *EmailNotifier) Update(event PaymentEvent) {
	fmt.Printf("[EMAIL] Payment notification sent to %s: Payment %s of %.2f via %s at %v\n",
		e.EmailAddress, event.Status, event.Amount, event.Method, event.TimeStamp)
}

// SMS Observer
type SMSNotifier struct {
	PhoneNumber string
}

func NewSMSNotifier(phone string) *SMSNotifier {
	return &SMSNotifier{PhoneNumber: phone}
}

func (s *SMSNotifier) Update(event PaymentEvent) {
	fmt.Printf("[SMS] Payment notification sent to %s: Payment %s of %.2f via %s at %v\n",
		s.PhoneNumber, event.Status, event.Amount, event.Method, event.TimeStamp)
}

// Webhook Observer
type WebhookNotifier struct {
	URL string
}

func NewWebhookNotifier(url string) *WebhookNotifier {
	return &WebhookNotifier{URL: url}
}

func (w *WebhookNotifier) Update(event PaymentEvent) {
	fmt.Printf("[WEBHOOK] Payment notification sent to %s: Payment %s of %.2f via %s at %v\n",
		w.URL, event.Status, event.Amount, event.Method, event.TimeStamp)
}
