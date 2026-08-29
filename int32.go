package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt32 struct {
	rand     *rand.Rand
	from, to int32
}

// ArbitraryInt32 - return generator for arbitray int32 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryInt32(rnd *rand.Rand, from, to int32) Arbitrary[int32] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt32) Generate() iter.Seq[int32] {
	return func(yield func(int32) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.Int32N(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryInt32) Shrink(value int32) iter.Seq[int32] {
	return shrink.Halving[int32](0)(value)
}
