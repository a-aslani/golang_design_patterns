package adapter

import "fmt"

type Ecommpay struct {
	ID     string
	Secret string
}

func (e *Ecommpay) PaymentUrl(amount float64, currency string, email string) (string, error) {
	fmt.Println("generating payment url for the Ecommpay")

	// send params to the Ecommpay to generate payment url

	return "payment_url", nil
}

func (e *Ecommpay) CheckCallbackUrl(transactionId string, url string) error {
	fmt.Println("checking the callback url from the Ecommpay")

	// checking

	return nil
}
