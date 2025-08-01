package ohsnap

import (
	"math/rand/v2"
)

type arbitraryInt32 struct {
	rand     *rand.Rand
	from, to int32
}

func ArbitraryInt32(rnd *rand.Rand, from, to int32) Arbitrary[int32] {
	return &arbitraryInt32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt32) Generate() int32 {
	return a.rand.Int32N(a.to-a.from) + a.from
}

func (arbitraryInt32) Shrink(value int32) []int32 {
	var results []int32

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
