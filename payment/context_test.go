package payment

import (
	"testing"

	payment_method "github.com/natanaelrusli/go-strategy-pattern/paymentmethod"
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
		paymentContext := NewPaymentContext()

		creditCard := payment_method.NewCreditCardPayment("1234-5678-9012-3456", "123", "12/25")
		paymentContext.SetStrategy(creditCard)
		result, err := paymentContext.ExecutePayment(100.50)

		assert.NoError(t, err)
		assert.Equal(t, "Paid 1234-5678-9012-3456 using Credit Card", result)

		paypal := payment_method.NewPayPalPayment("user@example.com")
		paymentContext.SetStrategy(paypal)
		result, err = paymentContext.ExecutePayment(50.75)

		assert.NoError(t, err)
		assert.Equal(t, "Paid user@example.com using PayPal account", result)

		debitCard := payment_method.NewDebitCardPayment("1234-5678-9012-3456", "123", "12/25")
		paymentContext.SetStrategy(debitCard)
		result, err = paymentContext.ExecutePayment(25.25)

		assert.NoError(t, err)
		assert.Equal(t, "Paid 1234-5678-9012-3456 using Debit Card", result)
	})
}
