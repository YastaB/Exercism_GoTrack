package brackets

// Bracket checks if the brackets match each other in a string
func Bracket(s string) bool {
	var chs = []rune{}
	var reverseBrackets = map[rune]rune{'}': '{', ')': '(', ']': '['}
	var brackets = map[rune]bool{'{': true, '(': true, '[': true, '}': true, ')': true, ']': true}
	for _, ch := range s {
		if _, ok := brackets[ch]; ok {
			if val, isIn := reverseBrackets[ch]; isIn && len(chs) > 0 && chs[len(chs)-1] == val {
				chs = chs[:len(chs)-1]
			} else {
				chs = append(chs, ch)
			}
		}
	}
	if len(chs) != 0 {
		return false
	}
	return true
}
