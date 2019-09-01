package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the given string is a valid credit card number
func Valid(code string) bool {
	// First remove the spaces
	code = strings.ReplaceAll(code, " ", "")

	if len(code) == 1 {
		return false
	}

	// Loop backwards through the code
	var total = 0
  isDouble := false
	for i := len(code) - 1; i >= 0; i-- {
		// If the value is not a number, the code is invalid
		if !unicode.IsNumber(rune(code[i])) {
			return false
		}

		newVal := int(code[i] - '0')
		// Boolean that is true every two characters
		if isDouble {

			newVal *= 2
			if newVal > 9 {
				newVal -= 9
			}
		}
    isDouble = !isDouble
		total += newVal
	}

	return total%10 == 0
}
