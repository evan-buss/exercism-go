package isogram

import (
	"unicode"
)

// IsIsogram determines if the given string is an isogram
func IsIsogram(s string) bool {

	prevLetters := make(map[rune]bool)

	for _, char := range s {
		char = unicode.ToLower(char)

		// If the rune is a space or hyphen, skip it
		if !unicode.IsLetter(char) {
			continue
		}

		// Letter not in the map yet, add it
		if _, ok := prevLetters[char]; !ok {
			prevLetters[char] = true
		} else {
			// Letter already in the map, not an isogram
			return false
		}
	}
	return true
}
