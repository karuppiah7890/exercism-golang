package diffsquares

// SquareOfSum returns the square of the sum of first n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of first n natural numbers
func SumOfSquares(n int) int {
	sumOfSquares := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

// Difference returns the difference between the square of the sum and
// the sum of the squares of the first N natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
