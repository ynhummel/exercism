package luhn

import (
	"strings"
)

func Valid(id string) bool {
	idWithoutSpaces := strings.ReplaceAll(id, " ", "")

	if len(idWithoutSpaces) <= 1 {
		return false
	}

	acc := 0
	for i := 0; i < len(idWithoutSpaces); i++ {
		idx := len(idWithoutSpaces) - 1 - i

		if idWithoutSpaces[idx] < '0' || idWithoutSpaces[idx] > '9' {
			return false
		}

		if i%2 != 0 {
			doubleDigit := int(idWithoutSpaces[idx]-'0') * 2
			if doubleDigit > 9 {
				doubleDigit = doubleDigit - 9
			}
			acc += doubleDigit
		} else {
			acc += int(idWithoutSpaces[idx] - '0')
		}
	}

	if acc%10 == 0 {
		return true
	}

	return false
}
