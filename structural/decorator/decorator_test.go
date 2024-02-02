package decorator

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPizza_AddIngredient(t *testing.T) {

	pizza := &Pizza{}

	s, err := pizza.AddIngredient()
	require.NoError(t, err)
	require.Equal(t, s, PizzaMsg)
}

func TestOnion_AddIngredient(t *testing.T) {

	pizza := &Pizza{}

	t.Run("should showing error when call method without sending pizza", func(t *testing.T) {

		onion := Onion{}
		s, err := onion.AddIngredient()
		require.Error(t, err, OnionErr.Error())
		require.Empty(t, s)

	})

	t.Run("should run without error by sending pizza", func(t *testing.T) {

		onion := Onion{Pizzaer: pizza}

		ps, _ := pizza.AddIngredient()

		s, err := onion.AddIngredient()
		require.NoError(t, err)
		require.Equal(t, s, fmt.Sprintf("%s %s", ps, OnionMsg))
	})

}

func TestMeat_AddIngredient(t *testing.T) {

	pizza := &Pizza{}

	t.Run("should showing error when call method without sending pizza", func(t *testing.T) {

		meat := Meat{}
		s, err := meat.AddIngredient()
		require.Error(t, err, MeatErr.Error())
		require.Empty(t, s)

	})

	t.Run("should run without error by sending pizza", func(t *testing.T) {

		meat := Meat{Pizzaer: pizza}

		ps, _ := pizza.AddIngredient()

		s, err := meat.AddIngredient()
		require.NoError(t, err)
		require.Equal(t, s, fmt.Sprintf("%s %s", ps, MeatMsg))
	})

}
