package decorator

import (
	"errors"
	"fmt"
)

var MeatErr = errors.New("pizzaer not found")

const MeatMsg = "Meat"

type Meat struct {
	Pizzaer Pizzaer
}

func (m *Meat) AddIngredient() (string, error) {

	if m.Pizzaer == nil {
		return "", MeatErr
	}

	s, err := m.Pizzaer.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", s, MeatMsg), nil
}
