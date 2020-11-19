package stacksandqueues

type FixedMultiStack struct {
	numberOfStacks int
	stackCapacity  int
	values         []int
	size           []int
}

type IFixedMultiStack interface {
	Pop(stackNum int) (int, error)
	Push(stackNum int, value int) error
	Peek(stackNum int) (int, error)
	IsEmpty(stackNum int) bool
	IsFull(stackNum int) bool
}

func CreateFixedMultiStack(numberOfStack int) {

}
