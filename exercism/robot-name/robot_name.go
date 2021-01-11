package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot struct
type Robot struct {
	name string
}

const maxNum = 26 * 26 * 10 * 10 * 10

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var usedNames = map[string]bool{}

func newName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

// Name returns robot's name, create if not exists
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	// number of items in the map represents used namespace
	if len(usedNames) == maxNum {
		return "", errors.New("namespace is exhausted")
	}

	r.name = newName()
	for usedNames[r.name] {
		r.name = newName()
	}

	return r.name, nil
}

// Reset the robot's name
func (r *Robot) Reset() {
	r.name = ""
}
