package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt16 struct {
	rand     *rand.Rand
	from, to int16
}

// ArbitraryInt16 - return generator for arbitrary int16 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryInt16(rnd *rand.Rand, from, to int16) Arbitrary[int16] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt16) Generate() iter.Seq[int16] {
	return func(yield func(int16) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := int16(a.rand.IntN(int(x))) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryInt16) Shrink(value int16) iter.Seq[int16] {
	return shrink.Halving[int16](0)(value)
}
