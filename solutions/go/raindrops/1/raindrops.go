package raindrops

import "fmt"

func Convert(number int) string {
	drops := ""
	match := false

	if number%3 == 0 {
		drops = drops + "Pling"
		match = true
	}

	if number%5 == 0 {
		drops = drops + "Plang"
		match = true
	}

	if number%7 == 0 {
		drops = drops + "Plong"
		match = true
	}

	if !match {
		drops = fmt.Sprint(number)
	}
	return drops
}
