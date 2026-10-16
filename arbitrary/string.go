package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitrationString struct {
	rand     *rand.Rand
	letters  string
	from, to int
}

// String - return generator for arbitrary strings.
// rnd - pseudo-random number generator.
// letters - string with allowed runes.
// from and to - bounds of length of strings.
func String(rnd *rand.Rand, letters string, from, to int) ohsnap.Arbitrary[string] {
	if from > to {
		from, to = to, from
	}

	return &arbitrationString{
		rand:    rnd,
		letters: letters,
		from:    from,
		to:      to,
	}
}

func (a arbitrationString) Generate() string {
	length := a.rand.IntN(a.to-a.from+1) + int(a.from)
	result := make([]byte, length)

	for i := range result {
		result[i] = a.letters[a.rand.IntN(len(a.letters))]
	}

	return string(result)
}

func (arbitrationString) Shrink(value string) iter.Seq[string] {
	return shrink.Peel()(value)
}
