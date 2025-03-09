package payment

import (
	"errors"
	"time"

	"github.com/google/uuid"
	payment_method "github.com/natanaelrusli/go-strategy-pattern/paymentmethod"
)

type PaymentContext struct {
	paymentId string
	strategy  payment_method.PaymentStrategy
	observers []PaymentObserver
}

func NewPaymentContext() *PaymentContext {
	return &PaymentContext{
		paymentId: uuid.New().String(),
		strategy:  nil,
		observers: make([]PaymentObserver, 0),
	}
}

func (pc *PaymentContext) SetStrategy(strategy payment_method.PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) AddObserver(observer PaymentObserver) {
	pc.observers = append(pc.observers, observer)
}

func (pc *PaymentContext) RemoveObserver(observer PaymentObserver) {
	for i, obs := range pc.observers {
		if obs == observer {
			pc.observers = append(pc.observers[:i], pc.observers[i+1:]...)
			break
		}
	}
}

func (pc *PaymentContext) notifyObservers(amount float64, method string, status string) {
	event := PaymentEvent{
		PaymentID: pc.paymentId,
		Amount:    amount,
		Method:    method,
		Status:    status,
		TimeStamp: time.Now(),
	}

	for _, observer := range pc.observers {
		observer.Update(event)
	}
}

func (pc *PaymentContext) ExecutePayment(amount float64) (string, error) {
	if pc.strategy == nil {
		return "", errors.New("no payment strategy chosen")
	}

	result := pc.strategy.Pay(amount)

	// Notify observers about successful payment
	pc.notifyObservers(amount, pc.strategy.Pay(amount), "SUCCESS")

	return result, nil
}
