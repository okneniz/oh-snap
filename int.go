package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt struct {
	rand     *rand.Rand
	from, to int
}

// ArbitraryInt - return generator for arbitrary int values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryInt(rnd *rand.Rand, from, to int) Arbitrary[int] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt) Generate() iter.Seq[int] {
	return func(yield func(int) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.IntN(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryInt) Shrink(value int) iter.Seq[int] {
	return shrink.Halving[int](0)(value)
}
