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
	//if (lexem[i] == 'x') {
	//	s21_set_structLexeme(&result, X_LEXEME, 0, 0);
	//	i++;
	//} else if (lexem[i] == '+') {
	//	s21_set_structLexeme(&result, PLUS_LEXEME, 1, 0);
	//	i++;
	//} else if (lexem[i] == '-') {
	//	s21_set_structLexeme(&result, MINUS_LEXEME, 1, 0);
	//	i++;
	//} else if (lexem[i] == '/') {
	//	s21_set_structLexeme(&result, DIV_LEXEME, 2, 0);
	//	i++;
	//} else if (lexem[i] == '*') {
	//	s21_set_structLexeme(&result, MUL_LEXEME, 2, 0);
	//	i++;
	//} else if (lexem[i] == '^') {
	//	s21_set_structLexeme(&result, POW_LEXEME, 3, 0);
	//	i++;
	//} else if (lexem[i] == 'm' && lexem[i + 1] == 'o' && lexem[i + 2] == 'd') {
	//	s21_set_structLexeme(&result, MOD_LEXEME, 3, 0);
	//	i += 3;
	//} else if (lexem[i] == 's' && lexem[i + 1] == 'i' && lexem[i + 2] == 'n') {
	//	s21_set_structLexeme(&result, SIN_LEXEME, 4, 0);
	//	i += 3;
	//} else if (lexem[i] == 'c' && lexem[i + 1] == 'o' && lexem[i + 2] == 's') {
	//	s21_set_structLexeme(&result, COS_LEXEME, 4, 0);
	//	i += 3;
	//} else if (lexem[i] == 't' && lexem[i + 1] == 'a' && lexem[i + 2] == 'n') {
	//	s21_set_structLexeme(&result, TAN_LEXEME, 4, 0);
	//	i += 3;
	//} else if (lexem[i] == 'l' && lexem[i + 1] == 'o' && lexem[i + 2] == 'g') {
	//	s21_set_structLexeme(&result, LOG_LEXEME, 4, 0);
	//	i += 3;
	//} else if (lexem[i] == 'l' && lexem[i + 1] == 'n') {
	//	s21_set_structLexeme(&result, LN_LEXEME, 4, 0);
	//	i += 2;
	//} else if (lexem[i] == 'a' && lexem[i + 1] == 's' && lexem[i + 2] == 'i' &&
	//	lexem[i + 3] == 'n') {
	//	s21_set_structLexeme(&result, ASIN_LEXEME, 4, 0);
	//	i += 4;
	//} else if (lexem[i] == 'a' && lexem[i + 1] == 'c' && lexem[i + 2] == 'o' &&
	//	lexem[i + 3] == 's') {
	//	s21_set_structLexeme(&result, ACOS_LEXEME, 4, 0);
	//	i += 4;
	//} else if (lexem[i] == 'a' && lexem[i + 1] == 't' && lexem[i + 2] == 'a' &&
	//	lexem[i + 3] == 'n') {
	//	s21_set_structLexeme(&result, ATAN_LEXEME, 4, 0);
	//	i += 4;
	//} else if (lexem[i] == 's' && lexem[i + 1] == 'q' && lexem[i + 2] == 'r' &&
	//	lexem[i + 3] == 't') {
	//	s21_set_structLexeme(&result, SQRT_LEXEME, 4, 0);
	//	i += 4;
	//} else if (lexem[i] == '(') {
	//	s21_set_structLexeme(&result, LEFTScobe_LEXEME, -1, 0);
	//	i++;
	//} else if (lexem[i] == ')') {
	//	s21_set_structLexeme(&result, RIGHTScobe_LEXEME, -1, 0);
	//	i++;
	//}
	//*end = (char *)&lexem[i];
	//return result;
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
	//tmp := List.Stack{}

	//var end string
	var i int
	for i < len(str) {
		str = strings.ReplaceAll(str, " ", "")
		// parse numbers
		value, err := findAndRemoveNumber(&str)
		if err != nil {
			fmt.Println(err)
		} else {
			numbers.Push(value, List.DIG_LEXEME, 0)
		}
		// parse lexemes
		tmp, err := parseLexemes(&str)
		fmt.Printf("Tmp: %v\n", tmp)
		fmt.Printf("Str: %v\n", str)
		break
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
