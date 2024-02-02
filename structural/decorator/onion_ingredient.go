package decorator

import (
	"errors"
	"fmt"
)

var OnionErr = errors.New("pizzaer not found")

const OnionMsg = "Onion"

type Onion struct {
	Pizzaer Pizzaer
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Pizzaer == nil {
		return "", OnionErr
	}

	s, err := o.Pizzaer.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", s, OnionMsg), nil
}
