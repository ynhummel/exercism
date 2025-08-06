package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("The number must be a non-zero positive integer")
	}

	steps := 0
	for n != 1 {
		steps += 1
		if n%2 == 0 {
			n = n / 2
			continue
		}
		n = 3*n + 1
	}
	return steps, nil
}
