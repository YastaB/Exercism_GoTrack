// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isContainAnyLetter(sentence string) bool {
	for _, letter := range sentence {
		if unicode.IsLetter(letter) {
			return true
		}
	}
	return false
}

func isContainNothingSpecial(sentence string) bool {
	for _, letter := range sentence {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			return false
		}
	}
	return true
}

func isQuestion(sentence string) bool {
	sentence = strings.TrimRight(sentence, " ")
	return len(sentence) > 0 && sentence[len(sentence)-1] == '?'
}

// Hey Detects the type of the sentences
// according to place of question mark, capital letter occurrence etc.
func Hey(remark string) string {
	var response string

	if isQuestion(remark) && remark == strings.ToUpper(remark) && isContainAnyLetter(remark) {
		response = "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		response = "Sure."
	} else if isContainAnyLetter(remark) && remark == strings.ToUpper(remark) {
		response = "Whoa, chill out!"
	} else if isContainNothingSpecial(remark) {
		response = "Fine. Be that way!"
	} else {
		response = "Whatever."
	}
	return response
}
