package questions

func Factors(n int64) []int64 {
	res := make([]int64, 0)
	// Theory used: A number is a product of it's prime factors only
	// 12 = 2 x 2 x 3
	for i := int64(2); n > 1; { // Start off with 2, until we reach n which will directly divide and give 1
		if n%i == 0 {
			res = append(res, i)
			n /= i
		} else {
			i++
		}
	}
	return res
}
