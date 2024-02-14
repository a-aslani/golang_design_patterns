package strategy

type OperationSubstract struct {
}

func (o *OperationSubstract) DoOperation(num1, num2 int) int {
	return num1 - num2
}
