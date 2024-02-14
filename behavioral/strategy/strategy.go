package strategy

type Strategy interface {
	DoOperation(num1, num2 int) int
}
