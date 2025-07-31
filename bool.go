package ohsnap

import (
	"math/rand"
)

type arbitraryBool struct {
	rand *rand.Rand
}

func ArbitraryBool(rnd *rand.Rand) Arbitrary[bool] {
	return &arbitraryBool{
		rand: rnd,
	}
}

func (a arbitraryBool) Generate() bool {
	return a.rand.Int()%2 == 0
}

func (arbitraryBool) Shrink(value bool) []bool {
	return nil
}
