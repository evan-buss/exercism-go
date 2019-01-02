package raindrops

import "strconv"

// Convert does some wierd shit
func Convert(num int) string {
	output := ""
	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}
	if len(output) > 0 {
		return output
	}
	return strconv.Itoa(num)
}
