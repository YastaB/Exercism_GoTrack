package pythagorean

import (
	"math"
)

// Triplet pythagorean triplet where [0] < [1] < [2]
type Triplet [3]int

// Range finds pythagorean triplets between [min,max]
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for i := min; i <= max; i++ {
		for j := i + 1; j <= max; j++ {
			k := i*i + j*j
			l := int(math.Sqrt(float64(k)))
			if isSquare(k) && l <= max {
				triplets = append(triplets, Triplet{i, j, l})
			}
		}
	}
	return triplets
}

// Sum finds pythagorean triplets between where sum of them is p
func Sum(p int) []Triplet {
	var triplets []Triplet
	for i := 1; i < p/2; i++ {
		for j := i + 1; j < p/2; j++ {
			k := i*i + j*j
			l := int(math.Sqrt(float64(k)))
			if isSquare(k) && i+j+l == p {
				triplets = append(triplets, Triplet{i, j, l})
			}
		}
	}
	return triplets
}

func isSquare(a int) bool {
	var intRoot = int(math.Sqrt(float64(a)))
	return (intRoot * intRoot) == a
}
