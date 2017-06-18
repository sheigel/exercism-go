//Package triangle offers support for identifying the type of a triangle
package triangle

import "math"

const testVersion = 3

type Kind string

var NaT Kind = "NaT" // not a triangle
var Equ Kind = "Equ" // equilateral
var Iso Kind = "Iso" // isosceles
var Sca Kind = "Sca" // scalene

var notAllowedTriangleSideValue = []func(side float64) bool{
	math.IsNaN,
	func(side float64) bool {
		return math.IsInf(side, -1)
	},
	func(side float64) bool {
		return math.IsInf(side, 1)
	},
	func(side float64) bool {
		return side == 0
	},
}

//KindFromSides identifies the Kind of a triangle based on the values of its sides:
//"NaT" // not a triangle
//"Equ" // equilateral
//"Iso" // isosceles
//"Sca" // scalene
func KindFromSides(a, b, c float64) Kind {
	if isNotATriangle(a, b, c) {
		return NaT
	}
	switch {
	case isEquilateral(a, b, c):
		return Equ
	case isIsosceles(a, b, c):
		return Iso
	default:
		return Sca
	}
}

func isNotATriangle(a, b, c float64) bool {
	for _, value := range notAllowedTriangleSideValue {
		if value(a) || value(b) || value(c) {
			return true
		}
	}
	if a+b < c || a+c < b || b+c < a {
		return true
	}
	return false
}

func isIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}
