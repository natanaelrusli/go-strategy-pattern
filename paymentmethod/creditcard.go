package payment_method

// Credit Card concrete strategy
type CreditCardPayment struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func NewCreditCardPayment(cardNumber, cvv, expiryDate string) PaymentStrategy {
	return &CreditCardPayment{
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
	}
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return "Paid " + c.CardNumber + " using Credit Card"
}

func (c *CreditCardPayment) TopUp(amount float64) string {
	return "Top-up " + c.CardNumber + " using Credit Card"
}

func (c *CreditCardPayment) GetMethodName() string {
	return "Credit Card"
}
