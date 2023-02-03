package Calculations

import "math"

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

type StrategyCalculation interface {
	Calculate(num, num2 float64, typeLex int64) float64
}
type Operations struct {
}
type Functions struct {
}

func (oper *Operations) Calculate(num, num2 float64, typeLex int64) float64 {
	var result float64
	if typeLex == PLUS_LEXEME {
		result = num + num2
	} else if typeLex == MINUS_LEXEME {
		result = num - num2
	} else if typeLex == DIV_LEXEME {
		result = num / num2
	} else if typeLex == MUL_LEXEME {
		result = num * num2
	} else if typeLex == POW_LEXEME {
		result = math.Pow(num, num2)
	} else if typeLex == MOD_LEXEME {
		result = math.Mod(num, num2)
	}
	return result
}
func (oper *Functions) Calculate(num, num2 float64, typeLex int64) float64 {
	var result float64
	if typeLex == SIN_LEXEME {
		result = math.Sin(num)
	} else if typeLex == COS_LEXEME {
		result = math.Cos(num)
	} else if typeLex == SQRT_LEXEME {
		result = math.Sqrt(num)
	} else if typeLex == TAN_LEXEME {
		result = math.Tan(num)
	} else if typeLex == ASIN_LEXEME {
		result = math.Asin(num)
	} else if typeLex == ACOS_LEXEME {
		result = math.Acos(num)
	} else if typeLex == ATAN_LEXEME {
		result = math.Atan(num)
	} else if typeLex == LN_LEXEME {
		result = math.Log(num)
	} else if typeLex == LOG_LEXEME {
		result = math.Log10(num)
	}
	return result
}

type Context struct {
	strategy StrategyCalculation
}

func (c *Context) SetStrategy(a StrategyCalculation) {
	c.strategy = a
}
func (c *Context) Calculate(num1, num2 float64, typeLex int64) float64 {
	return c.strategy.Calculate(num1, num2, typeLex)
}
func NewStrategy(strategyType string) StrategyCalculation {
	switch strategyType {
	case "operation":
		return &Operations{}
	case "function":
		return &Functions{}
	default:
		panic("unsupported strategy type")
	}
}
