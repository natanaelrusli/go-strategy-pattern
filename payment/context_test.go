package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentContext(t *testing.T) {
	t.Run("should generate unique payment id on PaymentContext construct", func(t *testing.T) {
		paymentContext := NewPaymentContext()
		assert.NotEmpty(t, paymentContext.paymentId)
	})

	t.Run("should return error if no payment strategy choosen", func(t *testing.T) {
		paymentContext := NewPaymentContext()
		result, err := paymentContext.ExecutePayment(100.50)

		assert.Error(t, err)
		assert.Empty(t, result)
	})

	t.Run("should be able to execute payment using different payment strategies", func(t *testing.T) {
		// Create payment context
		paymentContext := NewPaymentContext()

		// Credit Card Payment
		creditCard := &CreditCardPayment{
			CardNumber: "1234-5678-9012-3456",
			CVV:        "123",
			ExpiryDate: "12/25",
		}
		paymentContext.SetStrategy(creditCard)
		result, _ := paymentContext.ExecutePayment(100.50)
		if result != "Paid 100.50 using Credit Card 1234-5678-9012-3456" {
			t.Errorf("Expected 'Paid 100.50 using Credit Card 1234-5678-9012-3456', got '%s'", result)
		}

		// PayPal Payment
		paypal := &PayPalPayment{
			Email: "user@example.com",
		}
		paymentContext.SetStrategy(paypal)
		result, _ = paymentContext.ExecutePayment(50.75)
		if result != "Paid 50.75 using PayPal account: user@example.com" {
			t.Errorf("Expected 'Paid 50.75 using PayPal account: user@example.com', got '%s'", result)
		}

		// Bank Transfer Payment
		bankTransfer := &BankTransferPayment{
			AccountNumber: "987654321",
			BankName:      "Example Bank",
		}
		paymentContext.SetStrategy(bankTransfer)
		result, _ = paymentContext.ExecutePayment(75.25)
		if result != "Paid 75.25 using Bank Transfer from Example Bank account: 987654321" {
			t.Errorf("Expected 'Paid 75.25 using Bank Transfer from Example Bank account: 987654321', got '%s'", result)
		}

		// QRIS Payment
		qris := NewQRISPaymentMethod(20000)
		paymentContext.SetStrategy(qris)
		result, _ = paymentContext.ExecutePayment(20000)
		if result != "Paid 20000.00 using QRIS" {
			t.Errorf("Expected 'Paid 20000.00 using QRIS', got '%s'", result)
		}
		assert.Equal(t, "Paid 20000.00 using QRIS", result)
	})
}
