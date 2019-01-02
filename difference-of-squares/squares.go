package diffsquares

// SquareOfSum returns the square of the sum of numbers up to `num`
func SquareOfSum(num int) int {
	squareOfSum := 0
	for i := 1; i <= num; i++ {
		squareOfSum += i
	}
	return squareOfSum * squareOfSum
}

// SumOfSquares returns the sum of the squares of the numbers up to `num`
func SumOfSquares(num int) int {
	sumOfSquares := 0

	for i := 1; i <= num; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

// Difference returns the difference between SquareOfSum and SumOfSquares
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
