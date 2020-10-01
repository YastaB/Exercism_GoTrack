package pythagorean

import (
	"testing"
)

var testCases = map[int]bool{
	10: false,
	100: true,
	15: false,
	25: true,

}

func TestIsSquare(t *testing.T){
	for val, result := range testCases{
		isSqr := isSquare(val)
		if isSqr != result{
			t.Fatalf("IsSquare(%d) = %v, want %v",
				val, isSqr, result)
		}
	}
}