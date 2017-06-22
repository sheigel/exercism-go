package leap

const testVersion = 3

//IsLeapYear is a predicate that determines if the provided year is a leap year.
//There are 2 important conditions for a year to be leap:
// - if year is a multiple of 4 or 400 it is leap year
// - if year is a multiple of 100 but not of 400 it is not leap year
func IsLeapYear(year int) bool {
	switch {
	default:
		return false
	case year%100 == 0 && year%400 != 0:
		return false
	case year%4 == 0:
		return true
	}
}
