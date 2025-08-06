package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("The number must be a non-zero positive integer")
	}
	return collatz(n, 0), nil
}

func collatz(n, step int) int {
	if n == 1 {
		return step
	}

	if n%2 == 0 {
		return collatz(n/2, step+1)
	}

	return collatz(3*n+1, step+1)
}
