package List

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
	next     *Stack
}

func (stack *Stack) Push(value float64, typeLex int64, priority int64) {
	var tmp Stack
	tmp.value = value
	tmp.typeLex = typeLex
	tmp.priority = priority
	tmp.next = stack.next
	stack.next = &tmp
}

func (stack *Stack) Pop() {
	if stack.next != nil {
		stack.next = stack.next.next
	}
}

func (stack *Stack) Top() *Stack {
	tmp := stack.next
	return tmp
}
func (stack *Stack) SetStack(value float64, typeLex int64, priority int64) {
	*stack = Stack{
		value:    value,
		typeLex:  typeLex,
		priority: priority,
	}

}
