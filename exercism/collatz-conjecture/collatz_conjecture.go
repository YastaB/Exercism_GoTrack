package collatzconjecture

import "errors"

// CollatzConjecture throws error on 0 or negative values
// returns num of steps to reach to value 1
func CollatzConjecture(num int) (int, error) {
	steps := 0
	if num <= 0 {
		return -1, errors.New("invalid num is given, negative or zero values are are ")
	}
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
		steps++
	}
	return steps, nil
}
