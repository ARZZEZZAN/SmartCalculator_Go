package List

import "Calculations"

const (
	DIG_LEXEME = iota + 1
	MINUS_LEXEME
	PLUS_LEXEME
	DIV_LEXEME
	MUL_LEXEME
	POW_LEXEME
	MOD_LEXEME
	UNARM_LEXEME
	UNARP_LEXEME
	COS_LEXEME
	SIN_LEXEME
	TAN_LEXEME
	ACOS_LEXEME
	ASIN_LEXEME
	ATAN_LEXEME
	SQRT_LEXEME
	LN_LEXEME
	LOG_LEXEME
	LEFTScobe_LEXEME
	RIGHTScobe_LEXEME
)

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
