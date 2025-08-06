package isogram

import "strings"

func IsIsogram(word string) bool {
	byteWord := []byte(strings.ToLower(word))
	for i, ch := range byteWord {
		for _, ch2 := range byteWord[i+1:] {
			if ch != ' ' && ch != '-' && ch == ch2 {
				return false
			}
		}
	}
	return true
}
