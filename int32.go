package ohsnap

import (
	"math/rand/v2"
)

type arbitraryInt32 struct {
	rand     *rand.Rand
	from, to int32
}

// ArbitraryInt32 - return generator for arbitray int32 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryInt32(rnd *rand.Rand, from, to int32) Arbitrary[int32] {
	return &arbitraryInt32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt32) Generate() int32 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Int32N(x) + a.from
}

func (arbitraryInt32) Shrink(value int32) []int32 {
	var results []int32

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
