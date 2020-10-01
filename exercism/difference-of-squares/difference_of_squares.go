package diffsquares

// SquareOfSum adds all sequential numbers from 1 to num (inclusive)
func SquareOfSum(num int) int {
	sum := (num * (num + 1)) / 2
	return sum * sum
}

// SumOfSquares adds square of sequential numbers from 1 to num (inclusive)
func SumOfSquares(num int) int {
	return (num * (num + 1) * (2*num + 1)) / 6
}

// Difference subs SumOfSquares from SquareOfSum
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
