package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return count, errors.New("DNA strands must have the same size")
	}

	for ind, run := range a {
		if run != []rune(b)[ind] {
			count++
		}
	}
	return count, nil
}
