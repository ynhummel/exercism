package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	str := strings.ToLower(word)

	for i, ch := range str {
		if unicode.IsLetter(ch) && strings.ContainsRune(str[i+1:], ch) {
			return false
		}
	}
	return true
}
