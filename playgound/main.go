package main

import "fmt"

func IsIsogram(testStr string) bool {
	occurredChars := map[rune]bool{}
	for _, ch := range testStr {
		if ch != '-' && ch != ' ' {
			//
			if _, ok := occurredChars[ch]; !ok {
				occurredChars[ch] = true
			} else {
				return false
			}
		}
	}
	return true
}

// Reverse reverses a string
func Reverse(str string) string {
	reversed := ""
	for _, ch := range str {
		reversed = string(ch) + reversed
	}
	return reversed
}

func main() {
	fmt.Println(Reverse("Alphabet"))
}
