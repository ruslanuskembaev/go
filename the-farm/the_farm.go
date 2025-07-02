package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	fooderAmount, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (fooderAmount * fatteningFactor) / float64(numberOfCows), err
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		var errInvalidNumverOfCows = errors.New("invalid number of cows")
		return 0, errInvalidNumverOfCows
	}
	fooderAmount, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (fooderAmount * fatteningFactor) / float64(numberOfCows), nil
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numOfCows int) error {
	if numOfCows < 0 {
		errInvalidNumberOfCows := fmt.Errorf("%d cows are invalid: there are no negative cows", numOfCows)
		return errInvalidNumberOfCows
	}
	if numOfCows == 0 {
		errInvalidNumberOfCows := fmt.Errorf("%d cows are invalid: no cows don't need food", numOfCows)
		return errInvalidNumberOfCows
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
