package payment_method

// Bank Transfer concrete strategy
type BankTransferPayment struct {
	AccountNumber string
	BankName      string
}

func NewBankTransferPayment(accountNumber, bankName string) PaymentStrategy {
	return &BankTransferPayment{
		AccountNumber: accountNumber,
		BankName:      bankName,
	}
}

func (b *BankTransferPayment) Pay(amount float64) string {
	return "Paid " + b.AccountNumber + " using Bank Transfer from " + b.BankName + " account"
}

func (b *BankTransferPayment) TopUp(amount float64) string {
	return "Top-up " + b.AccountNumber + " using Bank Transfer from " + b.BankName + " account"
}

func (b *BankTransferPayment) GetMethodName() string {
	return "Bank Transfer"
}
