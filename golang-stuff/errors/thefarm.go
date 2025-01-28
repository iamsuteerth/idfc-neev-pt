package errors

import (
	"errors"
	"fmt"
)

// Just to remove the errors in VSCode, dummy struct and functions.

type FodderCalculator struct{}

func (fc FodderCalculator) FodderAmount(cows int) (int, error) {
	return 0, nil
}
func (fc FodderCalculator) FatteningFactor() (int, error) {
	return 0, nil
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	totalFodder, okFodder := fc.FodderAmount(cows)
	factor, okFactor := fc.FatteningFactor()
	if okFodder != nil || okFactor != nil {
		if okFodder == nil {
			return 0, okFactor
		} else {
			return 0, okFodder
		}
	}
	totalFodder *= factor
	return float64(totalFodder) / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fc, cows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return errors.New(fmt.Sprintf("%d cows are invalid: there are no negative cows", cows))
	} else if cows == 0 {
		return errors.New(fmt.Sprintf("%d cows are invalid: no cows don't need food", cows))
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
