package questions

import "math"

type Number struct {
	r float64
	i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	z := Number{
		r: n1.r + n2.r,
		i: n1.i + n2.i,
	}
	return z
}

func (n1 Number) Subtract(n2 Number) Number {
	z := Number{
		r: n1.r - n2.r,
		i: n1.i - n2.i,
	}
	return z
}

func (n1 Number) Multiply(n2 Number) Number {
	z := Number{
		r: (n1.r * n2.r) - (n1.i * n2.i),
		i: (n1.r * n2.i) + (n1.i * n2.r),
	}
	return z
}

func (n Number) Times(factor float64) Number {
	return Number{
		r: n.r * factor,
		i: n.i * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	denominator := n2.r*n2.r + n2.i*n2.i
	if denominator == 0 {
		panic("Division by zero")
	}
	realPart := (n1.r*n2.r + n1.i*n2.i) / denominator
	imaginaryPart := (n1.i*n2.r - n1.r*n2.i) / denominator
	return Number{r: realPart, i: imaginaryPart}
}

func (n Number) Conjugate() Number {
	return Number{r: n.r, i: (-1) * n.i}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.r*n.r + n.i*n.i)
}

func (n Number) Exp() Number {
	realPart := math.Exp(n.r)
	cosB := math.Cos(n.i)
	sinB := math.Sin(n.i)
	return Number{
		r: realPart * cosB,
		i: realPart * sinB,
	}
}
