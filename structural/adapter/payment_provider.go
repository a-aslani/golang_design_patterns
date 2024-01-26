package adapter

import "fmt"

type PaymentProvider interface {
	ExchangeCurrency(amount float64, currency string) (float64, error)
	GeneratePaymentLink(customerId string, customerEmail string) (string, error)
	Callback(transactionId string) error
}

type BetatransferAdapter struct {
	Betatransfer  Betatransfer
	amount        float64
	TransactionId string
	Currency      string
	Email         string
}

func (b *BetatransferAdapter) ExchangeCurrency(amount float64, currency string) (float64, error) {
	usdAmount, err := b.Betatransfer.ExchangeToUSD(amount, currency)
	b.amount = usdAmount
	return b.amount, err
}

func (b *BetatransferAdapter) GeneratePaymentLink(customerId string, customerEmail string) (string, error) {
	return b.Betatransfer.CreatePaymentUrl(b.amount, b.TransactionId, b.Currency, customerEmail)
}

func (b *BetatransferAdapter) Callback(transactionId string) error {
	url := fmt.Sprintf("http://anydomain/%s", transactionId)
	return b.Betatransfer.VerifyPayment(url, transactionId, b.Email)
}

type EcommpayAdapter struct {
	Ecommpay Ecommpay
	amount   float64
	Currency string
}

func (e *EcommpayAdapter) ExchangeCurrency(amount float64, currency string) (float64, error) {
	return amount, nil
}

func (e *EcommpayAdapter) GeneratePaymentLink(customerId string, customerEmail string) (string, error) {
	return e.Ecommpay.PaymentUrl(e.amount, e.Currency, customerEmail)
}

func (e *EcommpayAdapter) Callback(transactionId string) error {
	url := fmt.Sprintf("http://anydomain/%s", transactionId)
	return e.Ecommpay.CheckCallbackUrl(transactionId, url)
}
