package strategy

type OperationMultiply struct {
}

func (o *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1 * num2
}
