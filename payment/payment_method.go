package payment

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

// Credit Card concrete strategy
type CreditCardPayment struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func NewCreditCardPaymentMethod(cardNumber, cvv, expiryDate string) PaymentStrategy {
	return &CreditCardPayment{
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
	}
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, c.CardNumber)
}

// QRIS concrete strategy
type QRISPayment struct {
	Amount float64
}

func NewQRISPaymentMethod(amount float64) PaymentStrategy {
	return &QRISPayment{
		Amount: amount,
	}
}

func (q *QRISPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using QRIS", amount)
}

// PayPal concrete strategy
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account: %s", amount, p.Email)
}

// Bank Transfer concrete strategy
type BankTransferPayment struct {
	AccountNumber string
	BankName      string
}

func (b *BankTransferPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Bank Transfer from %s account: %s", amount, b.BankName, b.AccountNumber)
}
