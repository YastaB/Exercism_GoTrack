package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot struct
type RobotE struct {
	name string
}

const maxNumE = 26 * 26 * 10 * 10 * 10
var randomE = rand.New(rand.NewSource(time.Now().UnixNano()))

// namesArr holds numbers between 0 and maxNum
// in each name generation random index is selected
// and removed from the array
var namesArrE []int

// genRandomName generates a random name for robots without collision
func genRandomNameE() (string, error) {
	// fill the namesArr with equal to indexes
	if namesArrE == nil {
		namesArrE = make([]int, maxNum)
		for i := 0; i < maxNum; i++ {
			namesArrE[i] = i
		}
	}
	// remaining items in the array represents non-used namespace
	if len(namesArrE) == 0 {
		return "", errors.New("namespace is exhausted")
	}

	// get a random number within range of length of the array
	rndInd := random.Intn(len(namesArrE))
	rndNum := namesArrE[rndInd]

	// remove used number from the array in an efficient way
	namesArrE[rndInd] = namesArrE[len(namesArrE)-1]
	namesArrE = namesArrE[:len(namesArrE)-1]

	// convert random number to letters and numbers
	// rndNum = ch1 * 26^1 * 10^3 + ch2 * 26^0 * 10^3 + num1 * 10^2 + num2 * 10^1 + num3 * 10^0
	ch1 := rndNum / 26000
	ch2 := (rndNum % 26000) / 1000
	// the last 3 digits
	num := rndNum - (ch1*26000 + ch2*1000)

	return fmt.Sprintf("%c%c%03d", ch1+'A', ch2+'A', num), nil
}

// Name returns robot's name, create if not exists
func (r *Robot) NameE() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	n, err := genRandomNameE()
	if err != nil {
		return "", err
	}
	r.name = n
	return n, nil
}

// Reset the robot's name
func (r *Robot) ResetE() {
	r.name = ""
}
