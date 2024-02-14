package strategy

type OperationAdd struct{}

func (o *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}
