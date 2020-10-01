package cryptosquare

import (
	"fmt"
	"regexp"
	"strings"
)

func prepareString(message string) string{
	// remove alfa numeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil{
		fmt.Errorf("Regexp compile error")
		return ""
	}

	message = reg.ReplaceAllString(message, "")
	// convert to lowercase
	return strings.ToLower(message)
}

func findRectSize(size int) (int,int){
	//x := int(math.Sqrt(float64(size)))
	diff := size
	var sizeC, sizeR int
	for c:=1; c<size; c++{
		for r:=1; r<=c; r++{
			curDiff := (c * r) - size
			if diff > curDiff && curDiff >= 0 && c - r <= 1{
				diff = curDiff
				sizeC = c
				sizeR = r
			}
		}
	}

	return sizeC, sizeR
}

// Encode given message according to square code
func Encode(message string) string{
	message = prepareString(message)

	/*
	msgLen := len(message)




	for _, char := range message{

	}*/

	return ""
}