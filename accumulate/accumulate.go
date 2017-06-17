//accumulate offers support for mapping functions
package accumulate

const testVersion = 1

// Accumulate maps a []string into another []string of same size by applying the mapper function on each element of the original slice.
func Accumulate(slice []string, mapper func(string) string) []string {
	result := make([]string, 0, len(slice))
	for _, value := range slice {
		result = append(result, mapper(value))
	}
	return result
}
