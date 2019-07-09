package reverse

// Reverse reverses a string
func Reverse(str string) string {
	reversed := ""
	for _, ch := range str {
		reversed = string(ch) + reversed
	}
	return reversed
}
