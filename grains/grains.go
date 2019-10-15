package grains

import (
	"errors"
)

const maxInput = 64

// Square calculates number of grains in given sequence ,
// grains start from 1 and doubles each time
func Square(input int) (uint64, error) {
	if input <= 0 || input > maxInput {
		return 0, errors.New("input should be between 0 and 65")
	}
	return 1 << (input - 1), nil
}

// Total calculate sum of grains in chess board
func Total() uint64 {
	return 1<<maxInput - 1
}
