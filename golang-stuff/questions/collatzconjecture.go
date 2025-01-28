package questions

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		return 0, errors.New("Only positive integers are allowed.")
	}
	for ; n != 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n += 1
		}

	}
	return steps, nil
}
