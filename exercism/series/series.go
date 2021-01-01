package series

// All finds and returns all substrings which are n chars long
func All(n int, s string) []string {
	substrings := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

// UnsafeFirst returns first n chars long substring
func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return ""
	}
	return s[:n]
}

// First returns  n chars long substring with additional status return
func First(n int, s string) (first string, ok bool) {
	retStr := UnsafeFirst(n, s)
	if retStr == "" {
		return "", false
	}
	return retStr, true

}
