package factory_method

import (
	"errors"
	"fmt"
)

const (
	PaymentMethodNotExistErr = "payment method not exist"
)

type PaymentMethod interface {
	Pay(amount float64) (msg string, err error)
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(method int) (PaymentMethod, error) {

	switch method {
	case Cash:
		return new(CashPaymentMethod), nil
	case DebitCard:
		return new(DebitCardPaymentMethod), nil
	}

	return nil, errors.New(fmt.Sprintf("%d %s", method, PaymentMethodNotExistErr))
}

type CashPaymentMethod struct{}

func (c *CashPaymentMethod) Pay(amount float64) (msg string, err error) {
	return fmt.Sprintf("%0.2f paid using cash", amount), nil
}

type DebitCardPaymentMethod struct{}

func (d *DebitCardPaymentMethod) Pay(amount float64) (msg string, err error) {
	return fmt.Sprintf("%0.2f paid using debit card", amount), nil
}
