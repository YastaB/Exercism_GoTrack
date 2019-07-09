package raindrops

import "strconv"

// Convert converts input value to the string
// according to the factors that is defined.
func Convert(input int) string {
	var output = ""
	if input%3 == 0 {
		output += "Pling"
	}

	if input%5 == 0 {
		output += "Plang"
	}

	if input%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		output = strconv.Itoa(input)
	}

	return output
}
