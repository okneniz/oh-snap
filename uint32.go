package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint32 struct {
	rand     *rand.Rand
	from, to uint32
}

// ArbitraryUint32 - return generator for arbitrary uint32 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryUint32(rnd *rand.Rand, from, to uint32) Arbitrary[uint32] {
	return &arbitraryUint32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint32) Generate() uint32 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Uint32N(x) + a.from
}

func (arbitraryUint32) Shrink(value uint32) []uint32 {
	var results []uint32

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
