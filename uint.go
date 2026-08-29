package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryUint struct {
	rand     *rand.Rand
	from, to uint
}

// ArbitraryUint - return generator for arbitrary uint values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryUint(rnd *rand.Rand, from, to uint) Arbitrary[uint] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint) Generate() iter.Seq[uint] {
	return func(yield func(uint) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.UintN(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryUint) Shrink(value uint) iter.Seq[uint] {
	return shrink.Halving[uint](0)(value)
}
