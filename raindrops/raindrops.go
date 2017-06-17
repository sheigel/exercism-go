//Package raindrops provide conversion method from a number to raindrop-speak
package raindrops

import "strconv"

const testVersion = 3

// Convert translates the given integer to raindrop-speak
func Convert(num int) string {
	translation := pling(num) + plang(num) + plong(num)
	if len(translation) > 0 {
		return translation
	}
	return strconv.Itoa(num)
}

func pling(num int) string {
	if num%3 == 0 {
		return "Pling"
	}
	return ""
}

func plang(num int) string {
	if num%5 == 0 {
		return "Plang"
	}
	return ""
}
func plong(num int) string {
	if num%7 == 0 {
		return "Plong"
	}
	return ""
}
