package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitrationString struct {
	rand     *rand.Rand
	letters  string
	from, to int
}

// ArbitraryString - return generator for arbitrary strings.
// rnd - pseudo-random number generator.
// letters - string with allowed runes.
// from and to - bounds of length of strings.
func ArbitraryString(rnd *rand.Rand, letters string, from, to int) Arbitrary[string] {
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

func (a arbitrationString) Generate() iter.Seq[string] {
	return func(yield func(string) bool) {
		for {
			length := a.rand.IntN(a.to-a.from+1) + int(a.from)
			result := make([]byte, length)

			for i := range result {
				result[i] = a.letters[a.rand.IntN(len(a.letters))]
			}

			value := string(result)
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitrationString) Shrink(value string) iter.Seq[string] {
	return shrink.Peel()(value)
}
