package List

import "Calculations"

type Stack struct {
	value    float64
	priority int64
	typeLex  int64
	Next     *Stack
	Context  Calculations.Context
}

func (stack *Stack) Push(value float64, typeLex int64, priority int64) {
	var tmp Stack
	tmp.value = value
	tmp.typeLex = typeLex
	tmp.priority = priority
	tmp.Next = stack.Next
	stack.Next = &tmp
}

func (stack *Stack) Pop() {
	if stack.Next != nil {
		stack.Next = stack.Next.Next
	}
}

func (stack *Stack) Top() *Stack {
	tmp := stack.Next
	return tmp
}
func (stack *Stack) SetStack(value float64, typeLex int64, priority int64) {
	*stack = Stack{
		value:    value,
		typeLex:  typeLex,
		priority: priority,
	}
}
func (stack *Stack) GetType() int64 {
	return stack.typeLex
}
func (stack *Stack) GetPriority() int64 {
	return stack.priority
}
func (stack *Stack) GetValue() float64 {
	return stack.value
}

func (stack *Stack) SetValue(value float64) {
	stack.value = value
}
