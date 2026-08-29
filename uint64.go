package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryUint64 struct {
	rand     *rand.Rand
	from, to uint64
}

// ArbitraryUint64 - return generator for arbitrary uint64 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryUint64(rnd *rand.Rand, from, to uint64) Arbitrary[uint64] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint64) Generate() iter.Seq[uint64] {
	return func(yield func(uint64) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.Uint64N(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (a arbitraryUint64) Shrink(value uint64) iter.Seq[uint64] {
	return shrink.Halving[uint64](0)(value)
}
