package etl

import "strings"

// Transform convert map[int][]string type to map[string]int
func Transform(input map[int][]string) map[string]int {
	var transformed = map[string]int{}
	for k, v := range input {
		for _, str := range v {
			transformed[strings.ToLower(str)] = k
		}
	}
	return transformed
}
