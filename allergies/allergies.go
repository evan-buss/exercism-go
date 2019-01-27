package allergies

import "fmt"

var allergens = [...]string{"eggs", "peanuts", "shellfish", "strawberries",
	"tomatoes", "chocolate", "pollen", "cats"}

var allergies = map[string]byte{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128}

// AllergicTo checks whether the score indicates that they are allergic to a substance
func AllergicTo(score uint, substance string) bool {
	// Parse the score to find the allergies
	return byte(score)&allergies[substance] > 0
}

// Allergies returns a list of strings that you score indicates you are allergic to
func Allergies(score uint) []string {
	ret := make([]string, 0, 8)
	code := byte(score)
	fmt.Println("Starting Code:", code)
	for i := 0; code > 0; i++ {
		if code&1 == 1 {
			fmt.Println("if:", allergens[i], "code:", code)
			ret = append(ret, allergens[i])
		}
		code >>= 1
		fmt.Println("new code:", code)
	}
	return ret
}
