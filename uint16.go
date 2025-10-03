package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint16 struct {
	rand     *rand.Rand
	from, to uint16
}

// ArbitraryUint16 - return generator for arbitrary uint16 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryUint16(rnd *rand.Rand, from, to uint16) Arbitrary[uint16] {
	return &arbitraryUint16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint16) Generate() uint16 {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return uint16(a.rand.UintN(x)) + a.from
}

func (arbitraryUint16) Shrink(value uint16) []uint16 {
	var results []uint16

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
