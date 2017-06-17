package hamming

import (
	"errors"
)

const testVersion = 6

//Distance calculates the number of differences between two DNA strands.
//Distance fails if the provided DNA strands are of different length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Different length DNA strands")
	}

	var differenceCount int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differenceCount++
		}
	}
	return differenceCount, nil
}
