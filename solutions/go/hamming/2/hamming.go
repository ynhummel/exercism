package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return count, errors.New("DNA strands must have the same size")
	}
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
