package questions

func DartsScore(x, y float64) int {
	// For a circle : x^2 + y^2 = r^2
	x2 := x * x
	y2 := y * y
	if x2+y2 <= 1.0 {
		return 10
	}
	if x2+y2 > 1.0 && x2+y2 <= 25.0 {
		return 5
	}
	if x2+y2 > 25.0 && x2+y2 <= 100.0 {
		return 1
	}
	return 0
}
