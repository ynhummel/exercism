package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
const (
	MAXPREFIX = 26 * 26
	MAXSUFIX  = 10 * 10 * 10
)

var Letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var usedName = make(map[string]struct{}, MAXPREFIX*MAXSUFIX)

// var usedPrefix = make(map[string]struct{}, MAXPREFIX)
// var usedSufix = make(map[int]struct{}, MAXSUFIX)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}
	prefix := []rune{Letters[rand.Intn(len(Letters))], Letters[rand.Intn(len(Letters))]}
	sufix := rand.Intn(1000)

	newName := fmt.Sprintf("%s%03d", string(prefix), sufix)

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
