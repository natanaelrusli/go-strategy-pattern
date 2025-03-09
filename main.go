package main

import (
	"fmt"

	"github.com/natanaelrusli/go-strategy-pattern/payment"
	payment_method "github.com/natanaelrusli/go-strategy-pattern/paymentmethod"
)

func main() {
	fmt.Println("Payment Gateway: Strategy Pattern with Observer Example")

	// Create payment context
	paymentContext := payment.NewPaymentContext()

	// Add observers
	emailNotifier := payment.NewEmailNotifier("customer@example.com")
	smsNotifier := payment.NewSMSNotifier("+1234567890")
	webhookNotifier := payment.NewWebhookNotifier("https://api.example.com/webhook")

	paymentContext.AddObserver(emailNotifier)
	paymentContext.AddObserver(smsNotifier)
	paymentContext.AddObserver(webhookNotifier)

	// Credit Card Payment
	creditCard := &payment_method.CreditCardPayment{
		CardNumber: "1234-5678-9012-3456",
		CVV:        "123",
		ExpiryDate: "12/25",
	}
	paymentContext.SetStrategy(creditCard)
	result, _ := paymentContext.ExecutePayment(100.50)
	fmt.Println(result)

	fmt.Println("\n--- Removing SMS notifications ---")
	paymentContext.RemoveObserver(smsNotifier)

	// PayPal Payment
	paypal := &payment_method.PayPalPayment{
		Email: "user@example.com",
	}
	paymentContext.SetStrategy(paypal)
	result, _ = paymentContext.ExecutePayment(50.75)
	fmt.Println(result)
}
