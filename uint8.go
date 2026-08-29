package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryUint8 struct {
	rand     *rand.Rand
	from, to uint8
}

// ArbitraryUint8 - return generator for arbitrary uint8 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryUint8(rnd *rand.Rand, from, to uint8) Arbitrary[uint8] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint8) Generate() iter.Seq[uint8] {
	return func(yield func(uint8) bool) {
		x := uint(a.to - a.from)
		if x == 0 {
			x++
		}

		for {
			value := uint8(a.rand.UintN(x)) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryUint8) Shrink(value uint8) iter.Seq[uint8] {
	return func(yield func(uint8) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
