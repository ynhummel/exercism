package scrabble

import (
	"strings"
)

func Score(word string) int {
	points := 0
	word = strings.ToLower(word)
	for i := range word {
		switch word[i] {
		case 'q', 'z':
			points += 10
		case 'j', 'x':
			points += 8
		case 'k':
			points += 5
		case 'f', 'h', 'v', 'w', 'y':
			points += 4
		case 'b', 'c', 'm', 'p':
			points += 3
		case 'd', 'g':
			points += 2
		default:
			points += 1
		}
	}
	return points
}
