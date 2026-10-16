package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryBool struct {
	rand *rand.Rand
}

// Bool - returns generator for arbitrary bool values.
// rnd - pseudo-random number generator.
func Bool(rnd *rand.Rand) ohsnap.Arbitrary[bool] {
	return &arbitraryBool{
		rand: rnd,
	}
}

func (a arbitraryBool) Generate() bool {
	return a.rand.Int()%2 == 0
}

func (arbitraryBool) Shrink(value bool) iter.Seq[bool] {
	return shrink.NoShrink[bool]()(value)
}
