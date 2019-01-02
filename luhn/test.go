package luhn

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(changer("222"))
}

func changer(code string) string {
	output := ""
	for i := range code {
		temp, _ := strconv.Atoi(string(code[i]))
	}
	return output
}
