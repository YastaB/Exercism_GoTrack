package luhn

import "strconv"

// Valid check if given number digit is valid according to luhn algorithm
func Valid(num string) bool {
	numSum := 0
	numCount := 0

	for i := len(num) - 1; i >= 0; i-- {
		ch := num[i]
		if ch == ' ' {
			continue
		}
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			// reject other than space characters
			return false
		}
		// only increase numCount if it is a valid num
		numCount++

		if numCount%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		numSum += num
	}

	if numCount <= 1 {
		return false
	}

	return numSum%10 == 0
}
