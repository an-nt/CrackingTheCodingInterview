package stacksandqueues

import (
	"fmt"
)

type FixedMultiStack struct {
	numberOfStacks int
	stackCapacity  int
	values         []int
	currentSize    []int
}

type IFixedMultiStack interface {
	isStackNumOutRange(stackNum int) bool
	Pop(stackNum int) (int, error)
	Push(stackNum int, value int) error
	Peek(stackNum int) (int, error)
	isEmpty(stackNum int) bool
	isFull(stackNum int) bool
}

func CreateFixedMultiStack(numberOfStack int, stackSize int) IFixedMultiStack {
	return &FixedMultiStack{
		numberOfStacks: numberOfStack,
		stackCapacity:  stackSize,
		values:         make([]int, stackSize*numberOfStack),
		currentSize:    make([]int, numberOfStack),
	}
}

func (f *FixedMultiStack) isStackNumOutRange(stackNum int) bool {
	if stackNum > f.numberOfStacks || stackNum < 0 {
		return true
	}
	return false
}
func (f *FixedMultiStack) Pop(stackNum int) (int, error) {
	if f.isStackNumOutRange(stackNum) {
		return 0, fmt.Errorf("Stacknum out of range, %d out of 0-%d", stackNum, f.numberOfStacks)
	}

	if f.isEmpty(stackNum) {
		return 0, fmt.Errorf("Stack number %d is empty", stackNum)
	}

	atIndex := (stackNum-1)*f.stackCapacity + f.currentSize[stackNum-1] - 1
	f.currentSize[stackNum-1]--
	return f.values[atIndex], nil
}

func (f *FixedMultiStack) Push(stackNum int, value int) error {
	if f.isStackNumOutRange(stackNum) {
		return fmt.Errorf("Stacknum out of range, %d out of 0-%d", stackNum, f.numberOfStacks)
	}

	if f.isFull(stackNum) {
		return fmt.Errorf("Stack number %d is full", stackNum)
	}

	atIndex := (stackNum-1)*f.stackCapacity + f.currentSize[stackNum-1]
	f.values[atIndex] = value
	f.currentSize[stackNum-1]++
	return nil
}

func (f *FixedMultiStack) Peek(stackNum int) (int, error) {
	if f.isStackNumOutRange(stackNum) {
		return 0, fmt.Errorf("Stacknum out of range, %d out of 0-%d", stackNum, f.numberOfStacks)
	}

	atIndex := (stackNum-1)*f.stackCapacity + f.currentSize[stackNum-1] - 1
	return f.values[atIndex], nil
}

func (f *FixedMultiStack) isEmpty(stackNum int) bool {
	return f.currentSize[stackNum-1] == 0
}

func (f *FixedMultiStack) isFull(stackNum int) bool {
	return f.currentSize[stackNum-1] == f.stackCapacity
}
