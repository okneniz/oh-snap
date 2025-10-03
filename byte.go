package ohsnap

import (
	"math/rand/v2"
)

type arbitraryByte struct {
	rand     *rand.Rand
	from, to byte
}

// ArbitraryByte - return generator for arbitrary byte values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryByte(rnd *rand.Rand, from, to byte) Arbitrary[byte] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryByte{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryByte) Generate() byte {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return byte(a.rand.UintN(x)) + a.from
}

func (arbitraryByte) Shrink(value byte) []byte {
	var results []byte

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
