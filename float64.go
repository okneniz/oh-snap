package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryFloat64 struct {
	rand     *rand.Rand
	from, to float64
}

// ArbitraryFloat64 - return generator for arbitrary float64 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryFloat64(rnd *rand.Rand, from, to float64) Arbitrary[float64] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryFloat64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryFloat64) Generate() iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for {
			value := a.from + a.rand.Float64()*(a.to-a.from)
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryFloat64) Shrink(value float64) iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for value != 0 {
			value /= 2.0
			if !yield(value) {
				return
			}
		}
	}
}
