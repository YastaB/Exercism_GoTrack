package isogram

import "unicode"

//IsIsogram  checks if the characters in strings occurred multiple times
// if occurred multiple times, then it is not a Isogram
func IsIsogram(testStr string) bool {
	occurredChars := map[rune]bool{}
	for _, ch := range testStr {
		if ch != '-' && ch != ' ' {
			// convert character to lower case
			upperCh := unicode.ToUpper(ch)
			if _, ok := occurredChars[upperCh]; !ok {
				occurredChars[upperCh] = true
			} else {
				return false
			}
		}
	}
	return true
}
