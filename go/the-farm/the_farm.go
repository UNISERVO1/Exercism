package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fd FodderCalculator, numOfCows int) (float64, error) {
	fattener, errFattner := fd.FatteningFactor()
	if errFattner != nil {
		return 0, errFattner
	}
	fodder, errFodder := fd.FodderAmount(numOfCows)
	if errFodder != nil {
		return 0, errFodder
	}
	return fattener * fodder / float64(numOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fd FodderCalculator, numOfCows int) (float64, error) {
	if numOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fd, numOfCows)
}

type InvalidCowsError struct {
	numOfCows int
}

func (e *InvalidCowsError) Error() string {
	numCowsMessage := "there are no negative cows"
	if e.numOfCows == 0 {
		numCowsMessage = "no cows don't need food"
	}
	return fmt.Sprintf("%d cows are invalid: %s", e.numOfCows, numCowsMessage)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numOfCows int) error {
	if numOfCows <= 0 {
		return &InvalidCowsError{numOfCows: numOfCows}
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
