package factory_method

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCashPaymentMethod(t *testing.T) {

	t.Run("get payment method", func(t *testing.T) {
		cash, err := GetPaymentMethod(Cash)
		require.Nil(t, err)
		require.NotNil(t, cash)
	})

	t.Run("check cash payment method", func(t *testing.T) {
		cash, _ := GetPaymentMethod(Cash)
		msg, err := cash.Pay(10.5)
		require.Nil(t, err)
		require.NotEmpty(t, msg)
		require.Contains(t, msg, "paid using cash")
	})

}

func TestDebitCardPaymentMethod(t *testing.T) {

	t.Run("get payment method", func(t *testing.T) {
		debitCard, err := GetPaymentMethod(DebitCard)
		require.Nil(t, err)
		require.NotNil(t, debitCard)
	})

	t.Run("check debit card payment method", func(t *testing.T) {
		debitCard, _ := GetPaymentMethod(DebitCard)
		msg, err := debitCard.Pay(10.5)
		require.Nil(t, err)
		require.NotEmpty(t, msg)
		require.Contains(t, msg, "paid using debit card")
	})
}

func TestPaymentMethodNotExist(t *testing.T) {

	pm, err := GetPaymentMethod(100)
	require.NotNil(t, err)
	require.Nil(t, pm)
	require.Contains(t, err.Error(), PaymentMethodNotExistErr)
}
