package diffsquares

// SquareOfSum adds all sequential numbers from 1 to num (inclusive)
func SquareOfSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares adds square of sequential numbers from 1 to num (inclusive)
func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}

// Difference subs SumOfSquares from SquareOfSum
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
