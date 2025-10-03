package ohsnap

import (
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

func (a arbitraryInt64) Generate() int64 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Int64N(x) + a.from
}

func (arbitraryInt64) Shrink(value int64) []int64 {
	var results []int64

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
