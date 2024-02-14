package strategy

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) execute(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}
