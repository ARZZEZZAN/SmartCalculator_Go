package Calculations

type StrategyCalculation interface {
	Calculate(num, num2 float64, typeLex int64) float64
}
type Operations struct {
}
type Functions struct {
}

func (oper *Operations) Calculate(num, num2 float64, typeLex int64) float64 {
	var result float64
	result = 5.0
	//if (result.GetType() == List.PLUS_LEXEME) {
	//	result.GetValue() = num1 + num2
	//} else if (typeLex == MINUS_LEXEME) {
	//	result.value = num1 - num2;
	//} else if (typeLex == DIV_LEXEME) {
	//	result.value = num1 / num2;
	//} else if (typeLex == MUL_LEXEME) {
	//	result.value = num1 * num2;
	//} else if (typeLex == POW_LEXEME) {
	//	result.value = pow(num1, num2);
	//} else if (typeLex == MOD_LEXEME) {
	//	result.value = fmod(num1, num2);
	//}
	return result
}
func (oper *Functions) Calculate(num, num2 float64, typeLex int64) float64 {
	var result float64
	result = 6.0
	return result
}

// lexeme s21_calculate_withFunctions(double num, lexeme_enum type) {
// lexeme result = {0};
// s21_set_structLexeme(&result, DIG_LEXEME, 0, 0);
// if (type == SIN_LEXEME) {
// result.value = sin(num);
// } else if (type == COS_LEXEME) {
// result.value = cos(num);
// } else if (type == SQRT_LEXEME) {
// result.value = sqrt(num);
// } else if (type == TAN_LEXEME) {
// result.value = tan(num);
// } else if (type == ASIN_LEXEME) {
// result.value = asin(num);
// } else if (type == ACOS_LEXEME) {
// result.value = acos(num);
// } else if (type == ATAN_LEXEME) {
// result.value = atan(num);
// } else if (type == LN_LEXEME) {
// result.value = log(num);
// } else if (type == LOG_LEXEME) {
// result.value = log10(num);
// }
// return result;
// }
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
