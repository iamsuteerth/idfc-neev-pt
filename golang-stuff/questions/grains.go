package questions

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("invalid square number")
	}

	return 1 << (number - 1), nil // Bitwise left shift for power of 2
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		sq, err := Square(i)
		if err != nil {
			// In a real application, you might want to log this error
			// or handle it differently.  For this example, we'll
			// just ignore the error and continue.
			continue // Skip to the next iteration if there's an error
		}
		sum += sq
	}
	return sum
}
