package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint8 struct {
	rand     *rand.Rand
	from, to uint8
}

// ArbitraryUint8 - return generator for arbitrary uint8 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryUint8(rnd *rand.Rand, from, to uint8) Arbitrary[uint8] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint8) Generate() uint8 {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return uint8(a.rand.UintN(x)) + a.from
}

func (arbitraryUint8) Shrink(value uint8) []uint8 {
	var results []uint8

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
