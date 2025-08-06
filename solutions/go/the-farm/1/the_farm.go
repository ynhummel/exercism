package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

func (ce *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ce.cows, ce.message)
}

func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	amount, err := calc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fattening, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amount * fattening) / float64(cows), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calc, cows)
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	default:
		return nil
	}
}
