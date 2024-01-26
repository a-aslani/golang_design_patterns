package adapter

import "fmt"

type Betatransfer struct {
	Public        string
	Private       string
	PaymentSystem string
}

func (b *Betatransfer) CreatePaymentUrl(amount float64, transactionId string, currency string, email string) (string, error) {
	fmt.Println("generating payment url for the Betatransfer")

	// send params to the Betatransfer to generate payment url

	return "payment_url", nil
}

func (b *Betatransfer) VerifyPayment(url string, transactionId string, email string) error {
	fmt.Println("checking the callback url from the Betatransfer")

	// checking

	return nil
}

func (b *Betatransfer) ExchangeToUSD(amount float64, currency string) (float64, error) {
	fmt.Println("exchanging currency in the Betatransfer")

	// do exchange

	return amount, nil
}
