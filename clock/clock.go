package clock

import (
	"fmt"
	"math"
)

const testVersion = 4

// Clock represents a simple time representation.
// It is limited to storing only hours and minutes.
type Clock struct {
	Hour, Minute int
}

// New creates a Clock structure based on the provided hours and minutes.
// Given values that are smaller than 0 or bigger than 24 and 60 respectively it will trim these.
// e.g. 25, 61->02:01, -02, -70->20:50
func New(hours, minutes int) Clock {
	c := Clock{
		moduloHour(hours + hoursFromMinutes(minutes)),
		moduloMinutes(minutes),
	}
	return c
}

func hoursFromMinutes(minutes int) int {
	return int(math.Floor(float64(minutes) / 60))
}

func moduloHour(hours int) int {
	trimmedHours := (24 + hours%24) % 24
	return abs(trimmedHours)
}

func moduloMinutes(minutes int) int {
	trimmedMinutes := (60 + minutes%60) % 60
	return abs(trimmedMinutes)
}

func abs(i int) int {
	return (60 + i%60) % 60
}

// String returns the string representation of a Clock structure with hours and minutes padded with 0 to 2 decimal places. The hour is in 24h format.
// 8 hours and 2 minutes -> 08:02
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add returns a new clock that adds the provided minutes value to the receiver clock.
// Resulting values that are smaller than 0 or bigger than 24 and 60 respectively will be trimmed
// e.g. 25, 61->02:01, -02, -70->20:50
func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute+minutes)
}
