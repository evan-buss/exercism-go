package acronym

import "strings"

// Abbreviate returns the abreviation of the given string
func Abbreviate(s string) string {
	s = strings.Replace(s, " - ", " ", -1)
	s = strings.Replace(s, "-", " ", -1)
	words := strings.Split(s, " ")
	abrev := ""
	for _, word := range words {
		abrev += strings.ToUpper(string(word[0]))
	}
	return abrev
}
