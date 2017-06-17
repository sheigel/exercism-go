package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

var gsDuration = time.Duration(math.Pow(10, 9)) * time.Second

// AddGigasecond determines a gigasecond has passed given a start Time.
// Given a t time.Time it adds a gigasecond to the new time.Time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gsDuration)
}
