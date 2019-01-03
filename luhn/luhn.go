package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid checks if the given string is a valid credit card number
func Valid(code string) bool {
	// First remove the spaces
	code = strings.Replace(code, " ", "", -1)

	if len(code) == 1 {
		return false
	}

	output, isDouble := "", false
	// Loop backwards through the code
	for i := len(code) - 1; i >= 0; i-- {
		// If the value is not a number, the code is invalid
		if !unicode.IsNumber(rune(code[i])) {
			return false
		}

		// Boolean that is true every two characters
		if isDouble {
			newVal := int(code[i] - '0')
			newVal *= 2
			if newVal > 9 {
				newVal -= 9
			}
			output = strconv.Itoa(newVal) + output
		} else {
			output = string(code[i]) + output
		}
		isDouble = !isDouble
	}

	// Loop through the new code, adding each value
	total := 0
	for _, val := range output {
		total += (int(val) - 48)
	}
	return total%10 == 0
}
