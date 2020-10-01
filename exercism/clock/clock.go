package clock

import (
	"fmt"
)

// Clock struct represent time
type Clock struct {
	minute int
}

// New gets time as hour and minutes and returns adjusted clock
func New(hour int, minute int) Clock {
	totalMin := minute + hour*60
	totalMin %= 24 * 60
	if totalMin < 0 {
		totalMin += 24 * 60
	}
	return Clock{totalMin}
}

// string representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

// Add minutes to the clock
func (c Clock) Add(minute int) Clock {
	return New(0, c.minute+minute)
}

// Subtract minutes from the clock
func (c Clock) Subtract(minute int) Clock {
	return New(0, c.minute-minute)
}
