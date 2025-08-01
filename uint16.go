package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint16 struct {
	rand     *rand.Rand
	from, to uint16
}

func ArbitraryUint16(rnd *rand.Rand, from, to uint16) Arbitrary[uint16] {
	return &arbitraryUint16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint16) Generate() uint16 {
	return uint16(a.rand.UintN(uint(a.to-a.from))) + a.from
}

func (arbitraryUint16) Shrink(value uint16) []uint16 {
	var results []uint16

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
