package questions

import "math"

type Kind int

const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		k = NaT
		// Cover -Inf and Inf both
	case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) || math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1):
		k = NaT
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT
	case a+b < c || a+c < b || b+c < a:
		k = NaT
	case a == b && a == c:
		k = Equ
	case (a == b && a != c) || (a == c && a != b) || (b == c && a != b):
		k = Iso
	case a != b && a != c:
		k = Sca
	}
	return k
}
