package payment

import (
	"errors"

	"github.com/google/uuid"
)

type PaymentContext struct {
	paymentId string
	strategy  PaymentStrategy
}

func NewPaymentContext() *PaymentContext {
	return &PaymentContext{
		paymentId: uuid.New().String(),
		strategy:  nil,
	}
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) (string, error) {
	if pc.strategy == nil {
		return "", errors.New("no payment strategy choosen")
	}

	return pc.strategy.Pay(amount), nil
}
