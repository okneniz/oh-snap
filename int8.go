package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryInt8 struct {
	rand     *rand.Rand
	from, to int8
}

// ArbitraryInt8 - return generator for arbitrary int8 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryInt8(rnd *rand.Rand, from, to int8) Arbitrary[int8] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt8) Generate() iter.Seq[int8] {
	return func(yield func(int8) bool) {
		x := int(a.to - a.from)
		if x == 0 {
			x++
		}

		for {
			value := int8(a.rand.IntN(x)) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryInt8) Shrink(value int8) iter.Seq[int8] {
	return func(yield func(int8) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
