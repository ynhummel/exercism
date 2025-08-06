// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 != 0 {
		return true
	}

	if year%400 == 0 {
		return true
	}
	return false
}
