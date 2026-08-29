package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryUint32 struct {
	rand     *rand.Rand
	from, to uint32
}

// ArbitraryUint32 - return generator for arbitrary uint32 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryUint32(rnd *rand.Rand, from, to uint32) Arbitrary[uint32] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint32) Generate() iter.Seq[uint32] {
	return func(yield func(uint32) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.Uint32N(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryUint32) Shrink(value uint32) iter.Seq[uint32] {
	return func(yield func(uint32) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
