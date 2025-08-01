package ohsnap

import (
	"math/rand/v2"
)

type arbitraryUint64 struct {
	rand     *rand.Rand
	from, to uint64
}

func ArbitraryUint64(rnd *rand.Rand, from, to uint64) Arbitrary[uint64] {
	return &arbitraryUint64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint64) Generate() uint64 {
	return a.rand.Uint64N(a.to-a.from) + a.from
}

func (a arbitraryUint64) Shrink(value uint64) []uint64 {
	var results []uint64

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
