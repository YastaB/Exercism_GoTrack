package pangram

import "unicode"

const alphabetCount = 26

// IsPangram check if the input sentence contains every character in English alphabet
func IsPangram(input string) bool {
	var alphabetMap = make(map[rune]bool)
	for _, rune := range input {
		if unicode.IsLetter(rune) {
			alphabetMap[unicode.ToLower(rune)] = true
		}
	}
	if len(alphabetMap) != alphabetCount {
		return false
	}
	return true
}
