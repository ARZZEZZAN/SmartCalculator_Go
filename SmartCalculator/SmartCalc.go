package main

import (
	"Calculations"
	"List"
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseLexemes(str *string) (List.Stack, error) {
	err := errors.New("error: wrong operation input, please go to the school, it's not a joke, i'm serious")
	result := List.Stack{}
	// very big if statements, sorry!
	if (*str)[0] == '+' {
		result.SetStack(0, Calculations.PLUS_LEXEME, 1)
		*str = (*str)[1:]
	} else if (*str)[0] == '-' {
		result.SetStack(0, Calculations.MINUS_LEXEME, 1)
		*str = (*str)[1:]
	} else if (*str)[0] == '*' {
		result.SetStack(0, Calculations.MUL_LEXEME, 2)
		*str = (*str)[1:]
	} else if (*str)[0] == '/' {
		result.SetStack(0, Calculations.DIV_LEXEME, 2)
		*str = (*str)[1:]
	} else if (*str)[0] == '^' {
		result.SetStack(0, Calculations.POW_LEXEME, 3)
		*str = (*str)[1:]
	} else if (*str)[0] == '(' {
		result.SetStack(0, Calculations.LEFTScobe_LEXEME, -1)
		*str = (*str)[1:]
	} else if (*str)[0] == ')' {
		result.SetStack(0, Calculations.RIGHTScobe_LEXEME, -1)
		*str = (*str)[1:]
	} else if (*str)[0:3] == "mod" {
		result.SetStack(0, Calculations.MOD_LEXEME, 3)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "sin" {
		result.SetStack(0, Calculations.SIN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "cos" {
		result.SetStack(0, Calculations.COS_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "tan" {
		result.SetStack(0, Calculations.TAN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "log" {
		result.SetStack(0, Calculations.LOG_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:2] == "ln" {
		result.SetStack(0, Calculations.LN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:4] == "asin" {
		result.SetStack(0, Calculations.ASIN_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "acos" {
		result.SetStack(0, Calculations.ACOS_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "atan" {
		result.SetStack(0, Calculations.ATAN_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "sqrt" {
		result.SetStack(0, Calculations.SQRT_LEXEME, 4)
		*str = (*str)[4:]
	}
	return result, err
}
func findAndRemoveNumber(s *string) (float64, error) {
	// Use a regular expression to find the first number in the string
	re := regexp.MustCompile(`\d+`)
	// Find number
	match := re.FindString(*s)
	if match == "" {
		return 0, fmt.Errorf("no number found in string")
	}

	// Convert the matching number string to an int
	num, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, err
	}

	// Remove the number and following non-whitespace character from the original string
	*s = strings.Replace(*s, match, "", 1)

	// return
	return num, nil
}
func smartCalc(str string) {
	numbers := List.Stack{}
	operations := List.Stack{}
	rangeStr := len(str)
	for i := 0; i <= rangeStr; i++ {
		str = strings.ReplaceAll(str, " ", "")
		// parse numbers
		value, err := findAndRemoveNumber(&str)
		if err != nil {
			fmt.Println(err)
		} else {
			numbers.Push(value, Calculations.DIG_LEXEME, 0)
		}
		// parse lexemes
		if len(str) >= 1 {
			tmp, _ := parseLexemes(&str)
			priority(&numbers, &operations, tmp)
		}
		if len(str) == 0 {
			break
		}
	}
	for operations.Next != nil && numbers.Next != nil {
		calculation(&operations, &numbers)
	}
	fmt.Printf("The result of this expression: %v", numbers.Next.GetValue())
}
func priority(numbers *List.Stack, operations *List.Stack, tmp List.Stack) {
	topStack := &List.Stack{}
	if operations.Next != nil {
		topStack = operations.Top()
	} else {
		topStack.Push(0, 0, -1)
	}
	if tmp.GetType() == Calculations.LEFTScobe_LEXEME {
		operations.Push(0, Calculations.LEFTScobe_LEXEME, -1)
	} else if tmp.GetPriority() > topStack.GetPriority() &&
		tmp.GetType() != Calculations.RIGHTScobe_LEXEME {
		operations.Push(0, tmp.GetType(), tmp.GetPriority())
	} else if tmp.GetPriority() <= topStack.GetPriority() &&
		tmp.GetType() != Calculations.RIGHTScobe_LEXEME && tmp.GetType() != Calculations.LEFTScobe_LEXEME {
		calculation(operations, numbers)
		operations.Push(0, tmp.GetType(), tmp.GetPriority())
	} else if tmp.GetType() == Calculations.RIGHTScobe_LEXEME {
		for (operations != nil) && operations.GetType() != Calculations.LEFTScobe_LEXEME {
			calculation(operations, numbers)
		}
		operations.Pop()
	}
}
func calculation(operations *List.Stack, numbers *List.Stack) {
	operation := &List.Stack{}
	operation = operations.Top()
	operations.Pop()

	operand1 := &List.Stack{}
	operand2 := &List.Stack{}
	// For Understanding the operation because of my structure
	if operation.GetType() > 8 {
		operand2 = numbers.Top()
		numbers.Pop()
	} else {
		operand1 = numbers.Top()
		numbers.Pop()
		operand2 = numbers.Top()
		numbers.Pop()
	}
	result := &List.Stack{}
	if operation.GetType() <= 8 {
		result.Context.SetStrategy(Calculations.NewStrategy("operation"))
		result.SetValue(result.Context.Calculate(operand2.GetValue(), operand1.GetValue(), operation.GetType()))
	} else {
		result.Context.SetStrategy(Calculations.NewStrategy("function"))
		result.SetValue(result.Context.Calculate(operand2.GetValue(), operand1.GetValue(), operation.GetType()))
	}

	numbers.Push(result.GetValue(), result.GetType(), result.GetPriority())
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Wrong input variables:", err)
	}
	str := scanner.Text()
	smartCalc(str)
}
