package adapter

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBetatransfer(t *testing.T) {

	bt := BetatransferAdapter{
		Betatransfer: Betatransfer{
			Public:        uuid.New().String(),
			Private:       uuid.New().String(),
			PaymentSystem: faker.Name(),
		},
		TransactionId: uuid.New().String(),
		Currency:      "USD",
		Email:         faker.Email(),
	}

	amount, err := bt.ExchangeCurrency(100, bt.Currency)
	require.NoError(t, err)
	require.NotEmpty(t, amount)

	url, err := bt.GeneratePaymentLink(uuid.New().String(), bt.Email)
	require.NoError(t, err)
	require.NotEmpty(t, url)

	err = bt.Callback(bt.TransactionId)
	require.NoError(t, err)
}

func TestEcommpay(t *testing.T) {

	ep := EcommpayAdapter{
		Ecommpay: Ecommpay{
			ID:     uuid.New().String(),
			Secret: uuid.New().String(),
		},
		Currency: "USD",
	}

	amount, err := ep.ExchangeCurrency(100, ep.Currency)
	require.NoError(t, err)
	require.NotEmpty(t, amount)

	url, err := ep.GeneratePaymentLink(uuid.New().String(), faker.Email())
	require.NoError(t, err)
	require.NotEmpty(t, url)

	err = ep.Callback(uuid.New().String())
	require.NoError(t, err)
}
