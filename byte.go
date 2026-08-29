package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryByte struct {
	rand     *rand.Rand
	from, to byte
}

// ArbitraryByte - return generator for arbitrary byte values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryByte(rnd *rand.Rand, from, to byte) Arbitrary[byte] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryByte{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryByte) Generate() iter.Seq[byte] {
	return func(yield func(byte) bool) {
		x := uint(a.to - a.from)
		if x == 0 {
			x++
		}

		for {
			value := byte(a.rand.UintN(x)) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryByte) Shrink(value byte) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
