package payment_method

// Debit Card concrete strategy
type DebitCardPayment struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func NewDebitCardPayment(cardNumber, cvv, expiryDate string) PaymentStrategy {
	return &DebitCardPayment{
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
	}
}

func (d *DebitCardPayment) Pay(amount float64) string {
	return "Paid " + d.CardNumber + " using Debit Card"
}

func (d *DebitCardPayment) TopUp(amount float64) string {
	return "Top-up " + d.CardNumber + " using Debit Card"
}

func (d *DebitCardPayment) GetMethodName() string {
	return "Debit Card"
}
