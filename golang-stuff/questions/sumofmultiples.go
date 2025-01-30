package questions

func SumMultiples(limit int, divisors ...int) int { // find sum of multiples of given divisors
	// limit is the excluded upper limit
	// starting from
	sum := 0
	for i := 1; i < limit; i++ {
		// add the number if its a multiple of at least one of the divisors
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				sum += i
				break // to add each multiple only once
			}
		}
	}
	return sum
}
