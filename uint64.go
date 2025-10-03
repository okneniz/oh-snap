package ohsnap

import (
	"math/rand/v2"
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

func (a arbitraryUint64) Generate() uint64 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Uint64N(x) + a.from
}

func (a arbitraryUint64) Shrink(value uint64) []uint64 {
	var results []uint64

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
