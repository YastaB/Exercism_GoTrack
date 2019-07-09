package hamming

import "errors"

// Distance calculate hamming distance between a and b
// if lengths of strings dont match return error and -1 as distance
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("string lengths should be equal")
	}

	var distance = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
