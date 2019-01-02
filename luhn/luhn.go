package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(code string) bool {
	// First remove the spaces
	fmt.Println(code)
	code = strings.Replace(code, " ", "", -1)
	output := ""
	for i, val := range code {
		// Apply to every other index
		if i%2 != 0 {
			newVal, _ := strconv.Atoi(string(val))
			newVal *= 2
			output = output + string(newVal)
		} else {
			output = output + string(val)
		}
	}
	fmt.Println("Output: ", output)
	return false
}
