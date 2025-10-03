package ohsnap

import (
	"math/rand/v2"
)

type arbitraryInt struct {
	rand     *rand.Rand
	from, to int
}

// ArbitraryInt - return generator for arbitrary int values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryInt(rnd *rand.Rand, from, to int) Arbitrary[int] {
	return &arbitraryInt{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt) Generate() int {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.IntN(x) + a.from
}

func (arbitraryInt) Shrink(value int) []int {
	var results []int

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
