package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryInt64 struct {
	rand     *rand.Rand
	from, to int64
}

// ArbitraryInt64 - return generator for arbitrary int64 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryInt64(rnd *rand.Rand, from, to int64) Arbitrary[int64] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt64) Generate() iter.Seq[int64] {
	return func(yield func(int64) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.Int64N(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryInt64) Shrink(value int64) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
