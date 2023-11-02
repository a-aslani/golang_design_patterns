package prototype

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = iota
	Black
	Blue
)

func GetShirtCloner() ShirtCloner {
	return new(ShirtsCache)
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float64
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype = &Shirt{
	Price: 100.0,
	SKU:   "white",
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 120.0,
	SKU:   "black",
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 150.0,
	SKU:   "blue",
	Color: Blue,
}

func (s *Shirt) GetPrice() float64 {
	return s.Price
}
