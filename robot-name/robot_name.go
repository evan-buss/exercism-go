package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// Robot 's name was Robert Paulsen
type Robot struct {
	name string
}

// Set a random seed on initialization
func init() {
	rand.Seed(time.Now().Unix())
}

var names = make(map[string]bool)
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Name returns the robots name, or assigns a new one if the name does not exist
func (robot *Robot) Name() (string, error) {

	if robot.name == "" {

		newName := generateName()

		// Loop until a new name is generated
		_, exists := names[newName]
		for exists {
			newName = generateName()
			_, exists = names[newName]
		}

		names[newName] = true
		robot.name = newName
		return newName, nil
	}

	return robot.name, nil
}

// Reset removes a robot's name
func (robot *Robot) Reset() {
	robot.name = ""
}

func generateName() string {
	b := make([]rune, 2)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	// Generate new number string
	nums := ""
	for i := 0; i < 3; i++ {
		nums += strconv.Itoa(rand.Intn(10))
	}
	return string(b) + nums
}
