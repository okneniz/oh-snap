package ohsnap

import (
	"math/rand/v2"
)

type arbitraryInt8 struct {
	rand     *rand.Rand
	from, to int8
}

func ArbitraryInt8(rnd *rand.Rand, from, to int8) Arbitrary[int8] {
	return &arbitraryInt8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt8) Generate() int8 {
	return int8(a.rand.IntN(int(a.to-a.from))) + a.from
}

func (arbitraryInt8) Shrink(value int8) []int8 {
	var results []int8

	for value != 0 {
		value /= 2
		results = append(results, value)
	}

	return results
}
