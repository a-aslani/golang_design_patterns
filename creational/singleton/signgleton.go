package singleton

type Counter interface {
	Increment() int
	Decrement() int
	GetCount() int
}

type counter struct {
	count int
}

var instance *counter

func NewCounter() Counter {

	if instance == nil {
		instance = new(counter)
		return instance
	}

	return instance
}

func (c *counter) Increment() int {
	c.count++
	return c.count
}

func (c *counter) Decrement() int {
	c.count--
	return c.count
}

func (c *counter) GetCount() int {
	return c.count
}
