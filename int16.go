package ohsnap

import (
	"math/rand/v2"
)

type arbitraryInt16 struct {
	rand     *rand.Rand
	from, to int16
}

// ArbitraryInt16 - return generator for arbitrary int16 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryInt16(rnd *rand.Rand, from, to int16) Arbitrary[int16] {
	return &arbitraryInt16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt16) Generate() int16 {
	return int16(a.rand.IntN(int(a.to-a.from))) + a.from
}

func (arbitraryInt16) Shrink(value int16) []int16 {
	var results []int16

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
