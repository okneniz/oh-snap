package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt struct {
	rand     *rand.Rand
	from, to int
}

// Int - return generator for arbitrary int values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Int(rnd *rand.Rand, from, to int) ohsnap.Arbitrary[int] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt) Generate() int {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.IntN(x) + a.from
}

func (arbitraryInt) Shrink(value int) iter.Seq[int] {
	return shrink.Halving[int](0)(value)
}
