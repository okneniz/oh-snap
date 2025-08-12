package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint struct {
	rand     *rand.Rand
	from, to uint
}

// ArbitraryUint - return generator for arbitrary uint values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryUint(rnd *rand.Rand, from, to uint) Arbitrary[uint] {
	return &arbitraryUint{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint) Generate() uint {
	return a.rand.UintN(a.to-a.from) + a.from
}

func (arbitraryUint) Shrink(value uint) []uint {
	var results []uint

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
