package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides function returns triangle type as integer type Kind value
// Nat = 0, Equ = 1 Iso = 2 and Sca = 3
//
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a+b < c || a+c < b || b+c < a || a == 0 || b == 0 || c == 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || c == a {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
