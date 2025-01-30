package questions

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	if n < 9 {
		return true
	}
	numString := strconv.Itoa(n)
	order := len(numString)
	sum := 0
	for _, digit := range numString {
		sum += int(math.Pow(float64(digit-'0'), float64(order)))
	}
	return sum == n
}
