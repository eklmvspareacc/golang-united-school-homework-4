package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	input = reduseSigns(input)

	operands := []string{}
	index := strings.LastIndexAny(input, "+-")
	for index != -1 {
		left, right := stringSplitByRuneIndex(input, index)
		operands = append(operands, right)
		input = left
		index = strings.LastIndexAny(input, "+-")
	}
	if input != "" {
		operands = append(operands, input)
	}
	if len(operands) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	sum := 0
	for _, op := range operands {
		v, err := strconv.Atoi(op)
		if err != nil {
			return "", fmt.Errorf("invalid input expression: %w", err)
		}
		sum += v
	}

	return strconv.Itoa(sum), nil
}

// reduse multiply signs like +-+-+3 to +3
func reduseSigns(input string) (output string) {
	sign := 1
	atSign := false
	for _, ch := range input {
		switch ch {
		case '+':
			atSign = true
		case '-':
			sign *= -1
			atSign = true
		default:
			if atSign {
				if sign > 0 {
					output += "+"
				}
				if sign < 0 {
					output += "-"
				}
				atSign = false
				sign = 1
			}
			if !atSign {
				output += string(ch)
			}
		}
	}

	//Trailing sign meaningless but still for error
	if atSign {
		output += "+"
	}

	return
}

func stringSplitByRuneIndex(input string, index int) (left string, righ string) {
	runes := []rune(input)
	left = string(runes[:index])
	righ = string(runes[index:])
	return
}
