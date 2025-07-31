package ohsnap

import (
	"math/rand"
)

type arbitraryInt struct {
	rand     *rand.Rand
	from, to int
}

func ArbitraryInt(rnd *rand.Rand, from, to int) Arbitrary[int] {
	return &arbitraryInt{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt) Generate() int {
	return a.rand.Intn(a.to-a.from+1) + int(a.from)
}

func (arbitraryInt) Shrink(value int) []int {
	var results []int
	for value != 0 {
		value /= 2
		results = append(results, value)
	}
	return results
}
