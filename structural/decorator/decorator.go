package decorator

const PizzaMsg = "pizza:"

type Pizza struct{}

func (p *Pizza) AddIngredient() (string, error) {
	return PizzaMsg, nil
}
