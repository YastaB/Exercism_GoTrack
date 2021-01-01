package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Robot struct
type Robot struct {
	name string
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const maxNum = 26 * 26 * 10 * 10 * 10

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// namesArr holds numbers between 0 and maxNum
// in each name generation random index is selected
// and removed from the array
var namesArr []int

// genRandomName generates a random name for robots without collision
func genRandomName() (string, error) {
	// fill the namesArr with equal to indexes
	if namesArr == nil {
		namesArr = make([]int, maxNum)
		for i := 0; i < maxNum; i++ {
			namesArr[i] = i
		}
	}
	// remaining items in the array represents non-used namespace
	if len(namesArr) == 0 {
		return "", errors.New("namespace is exhausted")
	}

	// get a random number within range of length of the array
	rndInd := random.Intn(len(namesArr))
	rndNum := namesArr[rndInd]

	// remove used number from the array in an efficient way
	namesArr[rndInd] = namesArr[len(namesArr)-1]
	namesArr = namesArr[:len(namesArr)-1]

	// convert random number to letters and numbers
	// rndNum = ch1 * 26^1 * 10^3 + ch2 * 26^0 * 10^3 + num1 * 10^2 + num2 * 10^1 + num3 * 10^0
	ch1 := rndNum / 26000
	ch2 := (rndNum % 26000) / 1000
	// the last 3 digits
	nums := rndNum - (ch1*26000 + ch2*1000)

	strBuild := strings.Builder{}
	strBuild.WriteByte(letters[ch1])
	strBuild.WriteByte(letters[ch2])
	strBuild.WriteString(fmt.Sprintf("%03d", nums))
	return strBuild.String(), nil
}

// Name returns robot's name, create if not exists
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	n, err := genRandomName()
	if err != nil {
		return "", err
	}
	r.name = n
	return n, nil
}

// Reset the robot's name
func (r *Robot) Reset() {
	r.name = ""
}
