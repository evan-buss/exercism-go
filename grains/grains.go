package grains

import (
	"errors"
)

// Square calculates the number of grains on a given chessboard square
func Square(num int) (uint64, error) {

	if num > 0 && num <= 64 {
		prev := 1
		for i := 1; i < num; i++ {
			prev = prev * 2
		}
		return uint64(prev), nil
	}
	return 0, errors.New("invalid input")
}

// Total calculates the total number of grains on the entire chessboard
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		total += sq
	}

	return total
}
