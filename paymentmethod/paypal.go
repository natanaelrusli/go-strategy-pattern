package payment_method

// PayPal concrete strategy
type PayPalPayment struct {
	Email string
}

func NewPayPalPayment(email string) PaymentStrategy {
	return &PayPalPayment{
		Email: email,
	}
}

func (p *PayPalPayment) Pay(amount float64) string {
	return "Paid " + p.Email + " using PayPal account"
}

func (p *PayPalPayment) TopUp(amount float64) string {
	return "Top-up " + p.Email + " using PayPal account"
}

func (p *PayPalPayment) GetMethodName() string {
	return "PayPal"
}
