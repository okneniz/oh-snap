package ohsnap

import (
	"math/rand/v2"
)

type arbitraryByte struct {
	rand     *rand.Rand
	from, to byte
}

func ArbitraryByte(rnd *rand.Rand, from, to byte) Arbitrary[byte] {
	return &arbitraryByte{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryByte) Generate() byte {
	return byte(a.rand.UintN(uint(a.to-a.from))) + a.from
}

func (arbitraryByte) Shrink(value byte) []byte {
	var results []byte

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
