package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
const (
	MAXNAMES = 26 * 26 * 10 * 10 * 10
)

var Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var usedName = make(map[string]struct{}, MAXNAMES)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}

	letters := []rune{Letters[rand.Intn(len(Letters))], Letters[rand.Intn(len(Letters))]}
	numbers := rand.Intn(1000)

	newName := fmt.Sprintf("%s%03d", string(letters), numbers)
	if _, ok := usedName[newName]; ok {
		return r.Name()
	}

	usedName[newName] = struct{}{}
	r.name = newName
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
