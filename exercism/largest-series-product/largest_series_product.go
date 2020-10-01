package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct calculates the largest product for a contiguous
// substring of digits of length span
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return -1, errors.New("span is larger than digits length")
	}
	if span < 0 {
		return -1, errors.New("span is smaller than zero")
	}
	if len(digits) == 0 || span == 0 {
		return 1, nil
	}

	max := 0
	i := 0
	j := 0
	for i+span <= len(digits) {
		curProduct := 1
		j = i
		for j < len(digits) && j < i+span {
			x, err := strconv.Atoi(string(digits[j]))
			if err != nil {
				return -1, errors.New("non valid digit is encountered")
			}
			curProduct *= x
			j++
		}
		if j == i+span && curProduct > max {
			max = curProduct
		}
		i++
	}
	return max, nil
}
