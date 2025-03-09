package payment_method

// PaymentStrategy interface
type PaymentStrategy interface {
	Pay(amount float64) string
	TopUp(amount float64) string
	GetMethodName() string
}
