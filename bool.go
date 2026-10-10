package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryBool struct {
	rand *rand.Rand
}

// ArbitraryBool - returns generator for arbitrary bool values.
// rnd - pseudo-random number generator.
func ArbitraryBool(rnd *rand.Rand) Arbitrary[bool] {
	return &arbitraryBool{
		rand: rnd,
	}
}

func (a arbitraryBool) Generate() bool {
	return a.rand.Int()%2 == 0
}

func (arbitraryBool) Shrink(bool) iter.Seq[bool] {
	return Empty[bool]()
}
