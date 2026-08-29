package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryUint16 struct {
	rand     *rand.Rand
	from, to uint16
}

// ArbitraryUint16 - return generator for arbitrary uint16 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryUint16(rnd *rand.Rand, from, to uint16) Arbitrary[uint16] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint16) Generate() iter.Seq[uint16] {
	return func(yield func(uint16) bool) {
		x := uint(a.to - a.from)
		if x == 0 {
			x++
		}

		for {
			value := uint16(a.rand.UintN(x)) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryUint16) Shrink(value uint16) iter.Seq[uint16] {
	return func(yield func(uint16) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
