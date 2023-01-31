package main

import (
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
		result.SetStack(0, List.PLUS_LEXEME, 1)
		*str = (*str)[1:]
	} else if (*str)[0] == '-' {
		result.SetStack(0, List.MINUS_LEXEME, 1)
		*str = (*str)[1:]
	} else if (*str)[0] == '*' {
		result.SetStack(0, List.MUL_LEXEME, 2)
		*str = (*str)[1:]
	} else if (*str)[0] == '/' {
		result.SetStack(0, List.DIV_LEXEME, 2)
		*str = (*str)[1:]
	} else if (*str)[0] == '^' {
		result.SetStack(0, List.POW_LEXEME, 3)
		*str = (*str)[1:]
	} else if (*str)[0] == '(' {
		result.SetStack(0, List.LEFTScobe_LEXEME, -1)
		*str = (*str)[1:]
	} else if (*str)[0] == ')' {
		result.SetStack(0, List.RIGHTScobe_LEXEME, -1)
		*str = (*str)[1:]
	} else if (*str)[0:3] == "mod" {
		result.SetStack(0, List.MOD_LEXEME, 3)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "sin" {
		result.SetStack(0, List.SIN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "cos" {
		result.SetStack(0, List.COS_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "tan" {
		result.SetStack(0, List.TAN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:3] == "log" {
		result.SetStack(0, List.LOG_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:2] == "ln" {
		result.SetStack(0, List.LN_LEXEME, 4)
		*str = (*str)[3:]
	} else if (*str)[0:4] == "asin" {
		result.SetStack(0, List.ASIN_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "acos" {
		result.SetStack(0, List.ACOS_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "atan" {
		result.SetStack(0, List.ATAN_LEXEME, 4)
		*str = (*str)[4:]
	} else if (*str)[0:4] == "sqrt" {
		result.SetStack(0, List.SQRT_LEXEME, 4)
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
	//operations := List.Stack{}
	rangeStr := len(str)
	for i := 0; i <= rangeStr; i++ {
		str = strings.ReplaceAll(str, " ", "")
		// parse numbers
		value, err := findAndRemoveNumber(&str)
		if err != nil {
			fmt.Println(err)
		} else {
			numbers.Push(value, List.DIG_LEXEME, 0)
		}
		// parse lexemes
		if len(str) >= 1 {
			tmp, _ := parseLexemes(&str)
			fmt.Printf("Tmp: %v\n", tmp)
		}
		if len(str) == 0 {
			break
		}
	}
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
